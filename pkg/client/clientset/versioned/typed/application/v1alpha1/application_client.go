package v1alpha1

import (
	v1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type ArgoprojV1alpha1Interface interface {
	RESTClient() rest.Interface
	ApplicationsGetter
	ClustersGetter
}

// ArgoprojV1alpha1Client is used to interact with features provided by the argoproj.io group.
type ArgoprojV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ArgoprojV1alpha1Client) Applications(namespace string) ApplicationInterface {
	return newApplications(c, namespace)
}

func (c *ArgoprojV1alpha1Client) Clusters(namespace string) ClusterInterface {
	return newClusters(c, namespace)
}

// NewForConfig creates a new ArgoprojV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*ArgoprojV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ArgoprojV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ArgoprojV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ArgoprojV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ArgoprojV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ArgoprojV1alpha1Client {
	return &ArgoprojV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ArgoprojV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}