// This file was automatically generated by informer-gen

package v1alpha1

import (
	myapps_v1alpha1 "github.com/droot/memcached-api-server/pkg/apis/myapps/v1alpha1"
	versioned "github.com/droot/memcached-api-server/pkg/client/clientset/versioned"
	internalinterfaces "github.com/droot/memcached-api-server/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/droot/memcached-api-server/pkg/client/listers/myapps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// MemcachedInformer provides access to a shared informer and lister for
// Memcacheds.
type MemcachedInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MemcachedLister
}

type memcachedInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMemcachedInformer constructs a new informer for Memcached type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMemcachedInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMemcachedInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMemcachedInformer constructs a new informer for Memcached type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMemcachedInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MyappsV1alpha1().Memcacheds(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MyappsV1alpha1().Memcacheds(namespace).Watch(options)
			},
		},
		&myapps_v1alpha1.Memcached{},
		resyncPeriod,
		indexers,
	)
}

func (f *memcachedInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMemcachedInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *memcachedInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&myapps_v1alpha1.Memcached{}, f.defaultInformer)
}

func (f *memcachedInformer) Lister() v1alpha1.MemcachedLister {
	return v1alpha1.NewMemcachedLister(f.Informer().GetIndexer())
}
