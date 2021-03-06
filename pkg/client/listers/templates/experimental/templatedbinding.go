// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// This file was automatically generated by lister-gen

package experimental

import (
	experimental "github.com/Azure/service-catalog-templates/pkg/apis/templates/experimental"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TemplatedBindingLister helps list TemplatedBindings.
type TemplatedBindingLister interface {
	// List lists all TemplatedBindings in the indexer.
	List(selector labels.Selector) (ret []*experimental.TemplatedBinding, err error)
	// TemplatedBindings returns an object that can list and get TemplatedBindings.
	TemplatedBindings(namespace string) TemplatedBindingNamespaceLister
	TemplatedBindingListerExpansion
}

// templatedBindingLister implements the TemplatedBindingLister interface.
type templatedBindingLister struct {
	indexer cache.Indexer
}

// NewTemplatedBindingLister returns a new TemplatedBindingLister.
func NewTemplatedBindingLister(indexer cache.Indexer) TemplatedBindingLister {
	return &templatedBindingLister{indexer: indexer}
}

// List lists all TemplatedBindings in the indexer.
func (s *templatedBindingLister) List(selector labels.Selector) (ret []*experimental.TemplatedBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*experimental.TemplatedBinding))
	})
	return ret, err
}

// TemplatedBindings returns an object that can list and get TemplatedBindings.
func (s *templatedBindingLister) TemplatedBindings(namespace string) TemplatedBindingNamespaceLister {
	return templatedBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TemplatedBindingNamespaceLister helps list and get TemplatedBindings.
type TemplatedBindingNamespaceLister interface {
	// List lists all TemplatedBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*experimental.TemplatedBinding, err error)
	// Get retrieves the TemplatedBinding from the indexer for a given namespace and name.
	Get(name string) (*experimental.TemplatedBinding, error)
	TemplatedBindingNamespaceListerExpansion
}

// templatedBindingNamespaceLister implements the TemplatedBindingNamespaceLister
// interface.
type templatedBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TemplatedBindings in the indexer for a given namespace.
func (s templatedBindingNamespaceLister) List(selector labels.Selector) (ret []*experimental.TemplatedBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*experimental.TemplatedBinding))
	})
	return ret, err
}

// Get retrieves the TemplatedBinding from the indexer for a given namespace and name.
func (s templatedBindingNamespaceLister) Get(name string) (*experimental.TemplatedBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(experimental.Resource("templatedbinding"), name)
	}
	return obj.(*experimental.TemplatedBinding), nil
}
