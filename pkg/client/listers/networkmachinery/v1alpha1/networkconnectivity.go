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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/networkmachinery/networkmachinery-operators/pkg/apis/networkmachinery/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NetworkConnectivityLister helps list NetworkConnectivities.
type NetworkConnectivityLister interface {
	// List lists all NetworkConnectivities in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkConnectivityTest, err error)
	// Get retrieves the NetworkConnectivityTest from the index for a given name.
	Get(name string) (*v1alpha1.NetworkConnectivityTest, error)
	NetworkConnectivityListerExpansion
}

// networkConnectivityLister implements the NetworkConnectivityLister interface.
type networkConnectivityLister struct {
	indexer cache.Indexer
}

// NewNetworkConnectivityLister returns a new NetworkConnectivityLister.
func NewNetworkConnectivityLister(indexer cache.Indexer) NetworkConnectivityLister {
	return &networkConnectivityLister{indexer: indexer}
}

// List lists all NetworkConnectivities in the indexer.
func (s *networkConnectivityLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkConnectivityTest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkConnectivityTest))
	})
	return ret, err
}

// Get retrieves the NetworkConnectivityTest from the index for a given name.
func (s *networkConnectivityLister) Get(name string) (*v1alpha1.NetworkConnectivityTest, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("networkconnectivity"), name)
	}
	return obj.(*v1alpha1.NetworkConnectivityTest), nil
}
