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

package v1beta1

import (
	"context"
	time "time"

	softwarecompositionv1beta1 "github.com/Aryaman6492/storage/pkg/apis/softwarecomposition/v1beta1"
	versioned "github.com/Aryaman6492/storage/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/Aryaman6492/storage/pkg/generated/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/Aryaman6492/storage/pkg/generated/listers/softwarecomposition/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GeneratedNetworkPolicyInformer provides access to a shared informer and lister for
// GeneratedNetworkPolicies.
type GeneratedNetworkPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.GeneratedNetworkPolicyLister
}

type generatedNetworkPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGeneratedNetworkPolicyInformer constructs a new informer for GeneratedNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGeneratedNetworkPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGeneratedNetworkPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGeneratedNetworkPolicyInformer constructs a new informer for GeneratedNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGeneratedNetworkPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SpdxV1beta1().GeneratedNetworkPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SpdxV1beta1().GeneratedNetworkPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&softwarecompositionv1beta1.GeneratedNetworkPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *generatedNetworkPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGeneratedNetworkPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *generatedNetworkPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&softwarecompositionv1beta1.GeneratedNetworkPolicy{}, f.defaultInformer)
}

func (f *generatedNetworkPolicyInformer) Lister() v1beta1.GeneratedNetworkPolicyLister {
	return v1beta1.NewGeneratedNetworkPolicyLister(f.Informer().GetIndexer())
}
