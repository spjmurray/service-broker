/* Copyright (C) Couchbase, Inc 2020 - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	servicebroker "github.com/couchbase/service-broker/generated/clientset/servicebroker"
	internalinterfaces "github.com/couchbase/service-broker/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/couchbase/service-broker/generated/listers/servicebroker/v1alpha1"
	servicebrokerv1alpha1 "github.com/couchbase/service-broker/pkg/apis/servicebroker/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceBrokerConfigInformer provides access to a shared informer and lister for
// ServiceBrokerConfigs.
type ServiceBrokerConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ServiceBrokerConfigLister
}

type serviceBrokerConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceBrokerConfigInformer constructs a new informer for ServiceBrokerConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceBrokerConfigInformer(client servicebroker.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceBrokerConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceBrokerConfigInformer constructs a new informer for ServiceBrokerConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceBrokerConfigInformer(client servicebroker.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServicebrokerV1alpha1().ServiceBrokerConfigs(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServicebrokerV1alpha1().ServiceBrokerConfigs(namespace).Watch(options)
			},
		},
		&servicebrokerv1alpha1.ServiceBrokerConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceBrokerConfigInformer) defaultInformer(client servicebroker.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceBrokerConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceBrokerConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&servicebrokerv1alpha1.ServiceBrokerConfig{}, f.defaultInformer)
}

func (f *serviceBrokerConfigInformer) Lister() v1alpha1.ServiceBrokerConfigLister {
	return v1alpha1.NewServiceBrokerConfigLister(f.Informer().GetIndexer())
}