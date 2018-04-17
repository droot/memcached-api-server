package v1alpha1

import (
	v1alpha1 "github.com/droot/memcached-api-server/pkg/apis/myapps/v1alpha1"
	"github.com/droot/memcached-api-server/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type MyappsV1alpha1Interface interface {
	RESTClient() rest.Interface
	MemcachedsGetter
}

// MyappsV1alpha1Client is used to interact with features provided by the myapps.memcached.example.com group.
type MyappsV1alpha1Client struct {
	restClient rest.Interface
}

func (c *MyappsV1alpha1Client) Memcacheds(namespace string) MemcachedInterface {
	return newMemcacheds(c, namespace)
}

// NewForConfig creates a new MyappsV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*MyappsV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MyappsV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new MyappsV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *MyappsV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MyappsV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *MyappsV1alpha1Client {
	return &MyappsV1alpha1Client{c}
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
func (c *MyappsV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
