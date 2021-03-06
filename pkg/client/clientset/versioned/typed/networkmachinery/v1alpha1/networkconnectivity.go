/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/networkmachinery/networkmachinery-operators/pkg/apis/networkmachinery/v1alpha1"
	scheme "github.com/networkmachinery/networkmachinery-operators/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NetworkConnectivitiesGetter has a method to return a NetworkConnectivityInterface.
// A group's client should implement this interface.
type NetworkConnectivitiesGetter interface {
	NetworkConnectivities() NetworkConnectivityInterface
}

// NetworkConnectivityInterface has methods to work with NetworkConnectivityTest resources.
type NetworkConnectivityInterface interface {
	Create(*v1alpha1.NetworkConnectivityTest) (*v1alpha1.NetworkConnectivityTest, error)
	Update(*v1alpha1.NetworkConnectivityTest) (*v1alpha1.NetworkConnectivityTest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NetworkConnectivityTest, error)
	List(opts v1.ListOptions) (*v1alpha1.NetworkConnectivityList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkConnectivityTest, err error)
	NetworkConnectivityExpansion
}

// networkConnectivities implements NetworkConnectivityInterface
type networkConnectivities struct {
	client rest.Interface
}

// newNetworkConnectivities returns a NetworkConnectivities
func newNetworkConnectivities(c *NetworkmachineryV1alpha1Client) *networkConnectivities {
	return &networkConnectivities{
		client: c.RESTClient(),
	}
}

// Get takes name of the networkConnectivity, and returns the corresponding networkConnectivity object, and an error if there is any.
func (c *networkConnectivities) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkConnectivityTest, err error) {
	result = &v1alpha1.NetworkConnectivityTest{}
	err = c.client.Get().
		Resource("networkconnectivities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkConnectivities that match those selectors.
func (c *networkConnectivities) List(opts v1.ListOptions) (result *v1alpha1.NetworkConnectivityList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NetworkConnectivityList{}
	err = c.client.Get().
		Resource("networkconnectivities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkConnectivities.
func (c *networkConnectivities) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("networkconnectivities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a networkConnectivity and creates it.  Returns the server's representation of the networkConnectivity, and an error, if there is any.
func (c *networkConnectivities) Create(networkConnectivity *v1alpha1.NetworkConnectivityTest) (result *v1alpha1.NetworkConnectivityTest, err error) {
	result = &v1alpha1.NetworkConnectivityTest{}
	err = c.client.Post().
		Resource("networkconnectivities").
		Body(networkConnectivity).
		Do().
		Into(result)
	return
}

// Update takes the representation of a networkConnectivity and updates it. Returns the server's representation of the networkConnectivity, and an error, if there is any.
func (c *networkConnectivities) Update(networkConnectivity *v1alpha1.NetworkConnectivityTest) (result *v1alpha1.NetworkConnectivityTest, err error) {
	result = &v1alpha1.NetworkConnectivityTest{}
	err = c.client.Put().
		Resource("networkconnectivities").
		Name(networkConnectivity.Name).
		Body(networkConnectivity).
		Do().
		Into(result)
	return
}

// Delete takes name of the networkConnectivity and deletes it. Returns an error if one occurs.
func (c *networkConnectivities) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("networkconnectivities").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkConnectivities) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("networkconnectivities").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched networkConnectivity.
func (c *networkConnectivities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkConnectivityTest, err error) {
	result = &v1alpha1.NetworkConnectivityTest{}
	err = c.client.Patch(pt).
		Resource("networkconnectivities").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
