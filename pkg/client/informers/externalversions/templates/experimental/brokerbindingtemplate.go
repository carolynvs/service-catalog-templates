// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// This file was automatically generated by informer-gen

package experimental

import (
	time "time"

	templates_experimental "github.com/Azure/service-catalog-templates/pkg/apis/templates/experimental"
	versioned "github.com/Azure/service-catalog-templates/pkg/client/clientset/versioned"
	internalinterfaces "github.com/Azure/service-catalog-templates/pkg/client/informers/externalversions/internalinterfaces"
	experimental "github.com/Azure/service-catalog-templates/pkg/client/listers/templates/experimental"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BrokerBindingTemplateInformer provides access to a shared informer and lister for
// BrokerBindingTemplates.
type BrokerBindingTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() experimental.BrokerBindingTemplateLister
}

type brokerBindingTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewBrokerBindingTemplateInformer constructs a new informer for BrokerBindingTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBrokerBindingTemplateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBrokerBindingTemplateInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredBrokerBindingTemplateInformer constructs a new informer for BrokerBindingTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBrokerBindingTemplateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TemplatesExperimental().BrokerBindingTemplates().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TemplatesExperimental().BrokerBindingTemplates().Watch(options)
			},
		},
		&templates_experimental.BrokerBindingTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *brokerBindingTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBrokerBindingTemplateInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *brokerBindingTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&templates_experimental.BrokerBindingTemplate{}, f.defaultInformer)
}

func (f *brokerBindingTemplateInformer) Lister() experimental.BrokerBindingTemplateLister {
	return experimental.NewBrokerBindingTemplateLister(f.Informer().GetIndexer())
}