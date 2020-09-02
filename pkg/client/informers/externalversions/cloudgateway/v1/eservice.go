/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	cloudgatewayv1 "k8s.io/kubernetes/pkg/apis/cloudgateway/v1"
	versioned "k8s.io/kubernetes/pkg/client/clientset/versioned"
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/externalversions/internalinterfaces"
	v1 "k8s.io/kubernetes/pkg/client/listers/cloudgateway/v1"
)

// EServiceInformer provides access to a shared informer and lister for
// EServices.
type EServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.EServiceLister
}

type eServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
	tenant           string
}

// NewEServiceInformer constructs a new informer for EService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEServiceInformer(client, namespace, resyncPeriod, indexers, nil)
}

func NewEServiceInformerWithMultiTenancy(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tenant string) cache.SharedIndexInformer {
	return NewFilteredEServiceInformerWithMultiTenancy(client, namespace, resyncPeriod, indexers, nil, tenant)
}

// NewFilteredEServiceInformer constructs a new informer for EService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return NewFilteredEServiceInformerWithMultiTenancy(client, namespace, resyncPeriod, indexers, tweakListOptions, "system")
}

func NewFilteredEServiceInformerWithMultiTenancy(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc, tenant string) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudgatewayV1().EServicesWithMultiTenancy(namespace, tenant).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) watch.AggregatedWatchInterface {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudgatewayV1().EServicesWithMultiTenancy(namespace, tenant).Watch(options)
			},
		},
		&cloudgatewayv1.EService{},
		resyncPeriod,
		indexers,
	)
}

func (f *eServiceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEServiceInformerWithMultiTenancy(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions, f.tenant)
}

func (f *eServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cloudgatewayv1.EService{}, f.defaultInformer)
}

func (f *eServiceInformer) Lister() v1.EServiceLister {
	return v1.NewEServiceLister(f.Informer().GetIndexer())
}
