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

package fake

import (
	v1alpha1 "github.com/networkmachinery/networkmachinery-operators/pkg/apis/networkmachinery/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkMonitors implements NetworkMonitorInterface
type FakeNetworkMonitors struct {
	Fake *FakeNetworkmachineryV1alpha1
}

var networkmonitorsResource = schema.GroupVersionResource{Group: "networkmachinery.io", Version: "v1alpha1", Resource: "networkmonitors"}

var networkmonitorsKind = schema.GroupVersionKind{Group: "networkmachinery.io", Version: "v1alpha1", Kind: "NetworkMonitor"}

// Get takes name of the networkMonitor, and returns the corresponding networkMonitor object, and an error if there is any.
func (c *FakeNetworkMonitors) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(networkmonitorsResource, name), &v1alpha1.NetworkMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkMonitor), err
}

// List takes label and field selectors, and returns the list of NetworkMonitors that match those selectors.
func (c *FakeNetworkMonitors) List(opts v1.ListOptions) (result *v1alpha1.NetworkMonitorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(networkmonitorsResource, networkmonitorsKind, opts), &v1alpha1.NetworkMonitorList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkMonitorList{ListMeta: obj.(*v1alpha1.NetworkMonitorList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkMonitorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkMonitors.
func (c *FakeNetworkMonitors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(networkmonitorsResource, opts))
}

// Create takes the representation of a networkMonitor and creates it.  Returns the server's representation of the networkMonitor, and an error, if there is any.
func (c *FakeNetworkMonitors) Create(networkMonitor *v1alpha1.NetworkMonitor) (result *v1alpha1.NetworkMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(networkmonitorsResource, networkMonitor), &v1alpha1.NetworkMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkMonitor), err
}

// Update takes the representation of a networkMonitor and updates it. Returns the server's representation of the networkMonitor, and an error, if there is any.
func (c *FakeNetworkMonitors) Update(networkMonitor *v1alpha1.NetworkMonitor) (result *v1alpha1.NetworkMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(networkmonitorsResource, networkMonitor), &v1alpha1.NetworkMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkMonitor), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkMonitors) UpdateStatus(networkMonitor *v1alpha1.NetworkMonitor) (*v1alpha1.NetworkMonitor, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(networkmonitorsResource, "status", networkMonitor), &v1alpha1.NetworkMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkMonitor), err
}

// Delete takes name of the networkMonitor and deletes it. Returns an error if one occurs.
func (c *FakeNetworkMonitors) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(networkmonitorsResource, name), &v1alpha1.NetworkMonitor{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkMonitors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(networkmonitorsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkMonitorList{})
	return err
}

// Patch applies the patch and returns the patched networkMonitor.
func (c *FakeNetworkMonitors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(networkmonitorsResource, name, pt, data, subresources...), &v1alpha1.NetworkMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkMonitor), err
}
