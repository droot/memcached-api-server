package fake

import (
	v1alpha1 "github.com/droot/memcached-api-server/pkg/client/clientset/versioned/typed/myapps/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMyappsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMyappsV1alpha1) Memcacheds(namespace string) v1alpha1.MemcachedInterface {
	return &FakeMemcacheds{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMyappsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
