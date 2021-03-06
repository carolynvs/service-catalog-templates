// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.
package experimental

import (
	experimental "github.com/Azure/service-catalog-templates/pkg/apis/templates/experimental"
	scheme "github.com/Azure/service-catalog-templates/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BrokerBindingTemplatesGetter has a method to return a BrokerBindingTemplateInterface.
// A group's client should implement this interface.
type BrokerBindingTemplatesGetter interface {
	BrokerBindingTemplates() BrokerBindingTemplateInterface
}

// BrokerBindingTemplateInterface has methods to work with BrokerBindingTemplate resources.
type BrokerBindingTemplateInterface interface {
	Create(*experimental.BrokerBindingTemplate) (*experimental.BrokerBindingTemplate, error)
	Update(*experimental.BrokerBindingTemplate) (*experimental.BrokerBindingTemplate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*experimental.BrokerBindingTemplate, error)
	List(opts v1.ListOptions) (*experimental.BrokerBindingTemplateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *experimental.BrokerBindingTemplate, err error)
	BrokerBindingTemplateExpansion
}

// brokerBindingTemplates implements BrokerBindingTemplateInterface
type brokerBindingTemplates struct {
	client rest.Interface
}

// newBrokerBindingTemplates returns a BrokerBindingTemplates
func newBrokerBindingTemplates(c *TemplatesExperimentalClient) *brokerBindingTemplates {
	return &brokerBindingTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the brokerBindingTemplate, and returns the corresponding brokerBindingTemplate object, and an error if there is any.
func (c *brokerBindingTemplates) Get(name string, options v1.GetOptions) (result *experimental.BrokerBindingTemplate, err error) {
	result = &experimental.BrokerBindingTemplate{}
	err = c.client.Get().
		Resource("brokerbindingtemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BrokerBindingTemplates that match those selectors.
func (c *brokerBindingTemplates) List(opts v1.ListOptions) (result *experimental.BrokerBindingTemplateList, err error) {
	result = &experimental.BrokerBindingTemplateList{}
	err = c.client.Get().
		Resource("brokerbindingtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested brokerBindingTemplates.
func (c *brokerBindingTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("brokerbindingtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a brokerBindingTemplate and creates it.  Returns the server's representation of the brokerBindingTemplate, and an error, if there is any.
func (c *brokerBindingTemplates) Create(brokerBindingTemplate *experimental.BrokerBindingTemplate) (result *experimental.BrokerBindingTemplate, err error) {
	result = &experimental.BrokerBindingTemplate{}
	err = c.client.Post().
		Resource("brokerbindingtemplates").
		Body(brokerBindingTemplate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a brokerBindingTemplate and updates it. Returns the server's representation of the brokerBindingTemplate, and an error, if there is any.
func (c *brokerBindingTemplates) Update(brokerBindingTemplate *experimental.BrokerBindingTemplate) (result *experimental.BrokerBindingTemplate, err error) {
	result = &experimental.BrokerBindingTemplate{}
	err = c.client.Put().
		Resource("brokerbindingtemplates").
		Name(brokerBindingTemplate.Name).
		Body(brokerBindingTemplate).
		Do().
		Into(result)
	return
}

// Delete takes name of the brokerBindingTemplate and deletes it. Returns an error if one occurs.
func (c *brokerBindingTemplates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("brokerbindingtemplates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *brokerBindingTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("brokerbindingtemplates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched brokerBindingTemplate.
func (c *brokerBindingTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *experimental.BrokerBindingTemplate, err error) {
	result = &experimental.BrokerBindingTemplate{}
	err = c.client.Patch(pt).
		Resource("brokerbindingtemplates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
