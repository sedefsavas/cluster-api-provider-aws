module sigs.k8s.io/cluster-api-provider-aws

go 1.13

require (
	github.com/aws/aws-sdk-go v1.35.0
	github.com/awslabs/goformation/v4 v4.15.0
	github.com/blang/semver v3.5.1+incompatible
	github.com/docker/docker v0.7.3-0.20190327010347-be7ac8be2ae0 // indirect
	github.com/go-logr/logr v0.2.1
	github.com/golang/mock v1.4.4
	github.com/google/goexpect v0.0.0-20200816234442-b5b77125c2c5
	github.com/google/goterm v0.0.0-20200907032337-555d40f16ae2 // indirect
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1
	github.com/sedefsavas/cluster-api v0.2.6-0.20201119122002-b7effa1c4f68
	github.com/sergi/go-diff v1.1.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	golang.org/x/crypto v0.0.0-20200930160638-afb6bcd081ae
	golang.org/x/net v0.0.0-20200930145003-4acb6c075d10
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.19.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v0.19.2
	k8s.io/component-base v0.19.2
	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20200912215256-4140de9c8800
	sigs.k8s.io/aws-iam-authenticator v0.5.1
	sigs.k8s.io/cluster-api v0.3.10
	//sigs.k8s.io/cluster-api v0.3.10 // indirect
	sigs.k8s.io/controller-runtime v0.7.0-alpha.5
	sigs.k8s.io/structured-merge-diff/v2 v2.0.1 // indirect
	sigs.k8s.io/yaml v1.2.0
)
