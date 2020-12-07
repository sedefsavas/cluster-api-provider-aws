/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ekscontrolplanev1 "sigs.k8s.io/cluster-api-provider-aws/controlplane/eks/api/v1alpha4"
	infrav1exp "sigs.k8s.io/cluster-api-provider-aws/exp/api/v1alpha4"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha4"
	"sigs.k8s.io/cluster-api/util"
	"sigs.k8s.io/cluster-api/util/patch"
	"sigs.k8s.io/cluster-api/util/predicates"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// AWSManagedClusterReconciler reconciles AWSManagedCluster
type AWSManagedClusterReconciler struct {
	client.Client
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=awsmanagedclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=awsmanagedclusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=controlplane.cluster.x-k8s.io,resources=awsmanagedcontrolplanes;awsmanagedcontrolplanes/status,verbs=get;list;watch
// +kubebuilder:rbac:groups=cluster.x-k8s.io,resources=clusters;clusters/status,verbs=get;list;watch
// +kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch;create;update;patch;delete

func (r *AWSManagedClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (_ ctrl.Result, reterr error) {
	log := ctrl.LoggerFrom(ctx)

	// Fetch the AWSManagedCluster instance
	awsManagedCluster := &infrav1exp.AWSManagedCluster{}
	err := r.Get(ctx, req.NamespacedName, awsManagedCluster)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	// Fetch the Cluster.
	cluster, err := util.GetOwnerCluster(ctx, r.Client, awsManagedCluster.ObjectMeta)
	if err != nil {
		return reconcile.Result{}, err
	}
	if cluster == nil {
		log.Info("Cluster Controller has not yet set OwnerRef")
		return reconcile.Result{}, nil
	}

	if util.IsPaused(cluster, awsManagedCluster) {
		log.Info("AWSManagedCluster or linked Cluster is marked as paused. Won't reconcile")
		return reconcile.Result{}, nil
	}

	log = log.WithValues("cluster", cluster.Name)

	controlPlane := &ekscontrolplanev1.AWSManagedControlPlane{}
	controlPlaneRef := types.NamespacedName{
		Name:      cluster.Spec.ControlPlaneRef.Name,
		Namespace: cluster.Spec.ControlPlaneRef.Namespace,
	}

	if err := r.Get(ctx, controlPlaneRef, controlPlane); err != nil {
		return reconcile.Result{}, fmt.Errorf("failed to get control plane ref: %w", err)
	}

	log = log.WithValues("controlPlane", controlPlaneRef.Name)

	patchHelper, err := patch.NewHelper(awsManagedCluster, r.Client)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("failed to init patch helper: %w", err)
	}

	// Set the values from the managed control plane
	awsManagedCluster.Status.Ready = controlPlane.Status.Ready
	awsManagedCluster.Spec.ControlPlaneEndpoint = controlPlane.Spec.ControlPlaneEndpoint
	awsManagedCluster.Status.FailureDomains = controlPlane.Status.FailureDomains

	if err := patchHelper.Patch(ctx, awsManagedCluster); err != nil {
		return reconcile.Result{}, fmt.Errorf("failed to patch AWSManagedCluster: %w", err)
	}

	log.Info("Successfully reconciled AWSManagedCluster")

	return reconcile.Result{}, nil
}

func (r *AWSManagedClusterReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, options controller.Options) error {
	awsManagedCluster := &infrav1exp.AWSManagedCluster{}

	log := ctrl.LoggerFrom(ctx)

	controller, err := ctrl.NewControllerManagedBy(mgr).
		WithOptions(options).
		For(awsManagedCluster).
		WithEventFilter(predicates.ResourceNotPaused(log)).
		Build(r)

	if err != nil {
		return fmt.Errorf("error creating controller: %w", err)
	}

	// Add a watch for clusterv1.Cluster unpaise
	if err = controller.Watch(
		&source.Kind{Type: &clusterv1.Cluster{}},
		handler.EnqueueRequestsFromMapFunc(util.ClusterToInfrastructureMapFunc(awsManagedCluster.GroupVersionKind())),
		predicates.ClusterUnpaused(log),
	); err != nil {
		return fmt.Errorf("failed adding a watch for ready clusters: %w", err)
	}

	// Add a watch for AWSManagedControlPlane
	if err = controller.Watch(
		&source.Kind{Type: &ekscontrolplanev1.AWSManagedControlPlane{}},
		handler.EnqueueRequestsFromMapFunc(r.managedControlPlaneToManagedCluster(log)),
	); err != nil {
		return fmt.Errorf("failed adding watch on AWSManagedControlPlane: %w", err)
	}

	return nil
}

func (r *AWSManagedClusterReconciler) managedControlPlaneToManagedCluster(log logr.Logger) handler.MapFunc {
	return func(o client.Object) []ctrl.Request {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		awsManagedControlPlane, ok := o.(*ekscontrolplanev1.AWSManagedControlPlane)
		if !ok {
			panic(fmt.Sprintf("Expected a AWSManagedControlPlane but got a %T", o))
		}

		if !awsManagedControlPlane.ObjectMeta.DeletionTimestamp.IsZero() {
			log.V(4).Info("AWSManagedControlPlane has a deletion timestamp, skipping mapping")
			return nil
		}

		cluster, err := util.GetOwnerCluster(ctx, r.Client, awsManagedControlPlane.ObjectMeta)
		if err != nil {
			log.Error(err, "failed to get owning cluster")
			return nil
		}
		if cluster == nil {
			log.Info("no owning cluster, skipping mapping")
			return nil
		}

		managedClusterRef := cluster.Spec.InfrastructureRef
		if managedClusterRef == nil || managedClusterRef.Kind != "AWSManagedCluster" {
			log.V(4).Info("InfrastructureRef is nil or not AWSManagedCluster, skipping mapping")
			return nil
		}

		return []ctrl.Request{
			{
				NamespacedName: types.NamespacedName{
					Name:      managedClusterRef.Name,
					Namespace: managedClusterRef.Namespace,
				},
			},
		}
	}
}
