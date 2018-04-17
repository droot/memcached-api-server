// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/droot/memcached-api-server/pkg/apis/myapps/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MemcachedLister helps list Memcacheds.
type MemcachedLister interface {
	// List lists all Memcacheds in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Memcached, err error)
	// Memcacheds returns an object that can list and get Memcacheds.
	Memcacheds(namespace string) MemcachedNamespaceLister
	MemcachedListerExpansion
}

// memcachedLister implements the MemcachedLister interface.
type memcachedLister struct {
	indexer cache.Indexer
}

// NewMemcachedLister returns a new MemcachedLister.
func NewMemcachedLister(indexer cache.Indexer) MemcachedLister {
	return &memcachedLister{indexer: indexer}
}

// List lists all Memcacheds in the indexer.
func (s *memcachedLister) List(selector labels.Selector) (ret []*v1alpha1.Memcached, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Memcached))
	})
	return ret, err
}

// Memcacheds returns an object that can list and get Memcacheds.
func (s *memcachedLister) Memcacheds(namespace string) MemcachedNamespaceLister {
	return memcachedNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MemcachedNamespaceLister helps list and get Memcacheds.
type MemcachedNamespaceLister interface {
	// List lists all Memcacheds in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Memcached, err error)
	// Get retrieves the Memcached from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Memcached, error)
	MemcachedNamespaceListerExpansion
}

// memcachedNamespaceLister implements the MemcachedNamespaceLister
// interface.
type memcachedNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Memcacheds in the indexer for a given namespace.
func (s memcachedNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Memcached, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Memcached))
	})
	return ret, err
}

// Get retrieves the Memcached from the indexer for a given namespace and name.
func (s memcachedNamespaceLister) Get(name string) (*v1alpha1.Memcached, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("memcached"), name)
	}
	return obj.(*v1alpha1.Memcached), nil
}
