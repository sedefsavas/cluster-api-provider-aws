/*
Copyright 2021 The Kubernetes Authors.

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

package controllerprincipalcreator

import (
	"context"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	infrav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
	"sigs.k8s.io/cluster-api-provider-aws/controllers"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/scope"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// AWSControllerPrincipalReconciler reconciles a AWSClusterControllerPrincipal object
type AWSControllerPrincipalReconciler struct {
	client.Client
	Log       logr.Logger
	Endpoints []scope.ServiceEndpoint
}

// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=awsclustercontrollerprincipals,verbs=get;list;watch;create

func (r *AWSControllerPrincipalReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.TODO()
	log := r.Log.WithValues("namespace", req.NamespacedName, "awsControllerPrincipal", req.Name)

	// Fetch the AWSClusterControllerPrincipal instance
	controllerPrincipal := &infrav1.AWSClusterControllerPrincipal{}
	err := r.Get(ctx, req.NamespacedName, controllerPrincipal)
	// If AWSClusterControllerPrincipal instance already exists, then do not update it.
	if err == nil {
		return ctrl.Result{}, nil
	}
	if apierrors.IsNotFound(err) {
		log.Info("AWSClusterControllerPrincipal instance not found, creating a new instance", "cluster", req.Name)
		// Fetch the AWSClusterControllerPrincipal instance
		controllerPrincipal = &infrav1.AWSClusterControllerPrincipal{
			TypeMeta: metav1.TypeMeta{
				APIVersion: infrav1.GroupVersion.String(),
				Kind:       "AWSClusterControllerPrincipal",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: infrav1.AWSClusterControllerPrincipalName,
			},
			Spec: infrav1.AWSClusterControllerPrincipalSpec{
				AWSClusterPrincipalSpec: infrav1.AWSClusterPrincipalSpec{
					AllowedNamespaces: &infrav1.AllowedNamespaces{},
				},
			},
		}
		err := r.Create(ctx, controllerPrincipal)
		if err != nil {
			if apierrors.IsAlreadyExists(err) {
				return reconcile.Result{}, nil
			}
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}
	return reconcile.Result{}, err
}

func (r *AWSControllerPrincipalReconciler) SetupWithManager(mgr ctrl.Manager, options controller.Options) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrav1.AWSCluster{}).
		WithOptions(options).
		WithEventFilter(controllers.PausedPredicates(r.Log)).
		Complete(r)
}
