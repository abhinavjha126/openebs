module github.com/openebs/cstor-operators

go 1.13

require (
	github.com/cespare/xxhash v1.1.0
	github.com/davecgh/go-spew v1.1.1
	github.com/ghodss/yaml v1.0.0
	github.com/hashicorp/go-version v1.2.1
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.9.0
	github.com/openebs/api/v3 v3.0.0-20211116062351-ecd9a8a61d3e
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1
	github.com/ugorji/go/codec v0.0.0-20181204163529-d75b2dcb6bc8
	go.uber.org/zap v1.13.0
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	google.golang.org/grpc v1.27.1
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	k8s.io/klog v1.0.0
)

replace (
	k8s.io/api => k8s.io/api v0.20.2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.2

	k8s.io/apimachinery => k8s.io/apimachinery v0.20.2

	k8s.io/apiserver => k8s.io/apiserver v0.20.2

	k8s.io/cli-runtime => k8s.io/cli-runtime v0.20.2

	k8s.io/client-go => k8s.io/client-go v0.20.2

	k8s.io/cloud-provider => k8s.io/cloud-provider v0.20.2

	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.20.2

	k8s.io/code-generator => k8s.io/code-generator v0.20.2

	k8s.io/component-base => k8s.io/component-base v0.20.2
	k8s.io/component-helpers => k8s.io/component-helpers v0.20.2
	k8s.io/cri-api => k8s.io/cri-api v0.20.2

	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.20.2

	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.20.2

	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.20.2

	k8s.io/kube-proxy => k8s.io/kube-proxy v0.20.2

	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.20.2

	k8s.io/kubectl => k8s.io/kubectl v0.20.2

	k8s.io/kubelet => k8s.io/kubelet v0.20.2

	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.20.2

	k8s.io/metrics => k8s.io/metrics v0.20.2

	k8s.io/node-api => k8s.io/node-api v0.20.2

	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.20.2

	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.20.2

	k8s.io/sample-controller => k8s.io/sample-controller v0.20.2
)
