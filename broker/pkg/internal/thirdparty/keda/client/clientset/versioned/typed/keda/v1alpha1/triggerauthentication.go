/*
Copyright 2020 the original author or authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	scheme "knative.dev/eventing-rabbitmq/broker/pkg/internal/thirdparty/keda/client/clientset/versioned/scheme"
	v1alpha1 "knative.dev/eventing-rabbitmq/broker/pkg/internal/thirdparty/keda/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TriggerAuthenticationsGetter has a method to return a TriggerAuthenticationInterface.
// A group's client should implement this interface.
type TriggerAuthenticationsGetter interface {
	TriggerAuthentications(namespace string) TriggerAuthenticationInterface
}

// TriggerAuthenticationInterface has methods to work with TriggerAuthentication resources.
type TriggerAuthenticationInterface interface {
	Create(*v1alpha1.TriggerAuthentication) (*v1alpha1.TriggerAuthentication, error)
	Update(*v1alpha1.TriggerAuthentication) (*v1alpha1.TriggerAuthentication, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.TriggerAuthentication, error)
	List(opts v1.ListOptions) (*v1alpha1.TriggerAuthenticationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TriggerAuthentication, err error)
	TriggerAuthenticationExpansion
}

// triggerAuthentications implements TriggerAuthenticationInterface
type triggerAuthentications struct {
	client rest.Interface
	ns     string
}

// newTriggerAuthentications returns a TriggerAuthentications
func newTriggerAuthentications(c *KedaV1alpha1Client, namespace string) *triggerAuthentications {
	return &triggerAuthentications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the triggerAuthentication, and returns the corresponding triggerAuthentication object, and an error if there is any.
func (c *triggerAuthentications) Get(name string, options v1.GetOptions) (result *v1alpha1.TriggerAuthentication, err error) {
	result = &v1alpha1.TriggerAuthentication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("triggerauthentications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TriggerAuthentications that match those selectors.
func (c *triggerAuthentications) List(opts v1.ListOptions) (result *v1alpha1.TriggerAuthenticationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TriggerAuthenticationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("triggerauthentications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested triggerAuthentications.
func (c *triggerAuthentications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("triggerauthentications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a triggerAuthentication and creates it.  Returns the server's representation of the triggerAuthentication, and an error, if there is any.
func (c *triggerAuthentications) Create(triggerAuthentication *v1alpha1.TriggerAuthentication) (result *v1alpha1.TriggerAuthentication, err error) {
	result = &v1alpha1.TriggerAuthentication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("triggerauthentications").
		Body(triggerAuthentication).
		Do().
		Into(result)
	return
}

// Update takes the representation of a triggerAuthentication and updates it. Returns the server's representation of the triggerAuthentication, and an error, if there is any.
func (c *triggerAuthentications) Update(triggerAuthentication *v1alpha1.TriggerAuthentication) (result *v1alpha1.TriggerAuthentication, err error) {
	result = &v1alpha1.TriggerAuthentication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("triggerauthentications").
		Name(triggerAuthentication.Name).
		Body(triggerAuthentication).
		Do().
		Into(result)
	return
}

// Delete takes name of the triggerAuthentication and deletes it. Returns an error if one occurs.
func (c *triggerAuthentications) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("triggerauthentications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *triggerAuthentications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("triggerauthentications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched triggerAuthentication.
func (c *triggerAuthentications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TriggerAuthentication, err error) {
	result = &v1alpha1.TriggerAuthentication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("triggerauthentications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
