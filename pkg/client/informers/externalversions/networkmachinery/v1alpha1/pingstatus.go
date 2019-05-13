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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	networkmachineryv1alpha1 "github.com/networkmachinery/networkmachinery-operators/pkg/apis/networkmachinery/v1alpha1"
	versioned "github.com/networkmachinery/networkmachinery-operators/pkg/client/clientset/versioned"
	internalinterfaces "github.com/networkmachinery/networkmachinery-operators/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/networkmachinery/networkmachinery-operators/pkg/client/listers/networkmachinery/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PingStatusInformer provides access to a shared informer and lister for
// PingStatuses.
type PingStatusInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PingStatusLister
}

type pingStatusInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPingStatusInformer constructs a new informer for PingStatus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPingStatusInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPingStatusInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPingStatusInformer constructs a new informer for PingStatus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPingStatusInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkmachineryV1alpha1().PingStatuses().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkmachineryV1alpha1().PingStatuses().Watch(options)
			},
		},
		&networkmachineryv1alpha1.PingStatus{},
		resyncPeriod,
		indexers,
	)
}

func (f *pingStatusInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPingStatusInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *pingStatusInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkmachineryv1alpha1.PingStatus{}, f.defaultInformer)
}

func (f *pingStatusInformer) Lister() v1alpha1.PingStatusLister {
	return v1alpha1.NewPingStatusLister(f.Informer().GetIndexer())
}
