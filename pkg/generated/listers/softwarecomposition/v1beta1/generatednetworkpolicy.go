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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/kubescape/storage/pkg/apis/softwarecomposition/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GeneratedNetworkPolicyLister helps list GeneratedNetworkPolicies.
// All objects returned here must be treated as read-only.
type GeneratedNetworkPolicyLister interface {
	// List lists all GeneratedNetworkPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.GeneratedNetworkPolicy, err error)
	// GeneratedNetworkPolicies returns an object that can list and get GeneratedNetworkPolicies.
	GeneratedNetworkPolicies(namespace string) GeneratedNetworkPolicyNamespaceLister
	GeneratedNetworkPolicyListerExpansion
}

// generatedNetworkPolicyLister implements the GeneratedNetworkPolicyLister interface.
type generatedNetworkPolicyLister struct {
	indexer cache.Indexer
}

// NewGeneratedNetworkPolicyLister returns a new GeneratedNetworkPolicyLister.
func NewGeneratedNetworkPolicyLister(indexer cache.Indexer) GeneratedNetworkPolicyLister {
	return &generatedNetworkPolicyLister{indexer: indexer}
}

// List lists all GeneratedNetworkPolicies in the indexer.
func (s *generatedNetworkPolicyLister) List(selector labels.Selector) (ret []*v1beta1.GeneratedNetworkPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.GeneratedNetworkPolicy))
	})
	return ret, err
}

// GeneratedNetworkPolicies returns an object that can list and get GeneratedNetworkPolicies.
func (s *generatedNetworkPolicyLister) GeneratedNetworkPolicies(namespace string) GeneratedNetworkPolicyNamespaceLister {
	return generatedNetworkPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GeneratedNetworkPolicyNamespaceLister helps list and get GeneratedNetworkPolicies.
// All objects returned here must be treated as read-only.
type GeneratedNetworkPolicyNamespaceLister interface {
	// List lists all GeneratedNetworkPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.GeneratedNetworkPolicy, err error)
	// Get retrieves the GeneratedNetworkPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.GeneratedNetworkPolicy, error)
	GeneratedNetworkPolicyNamespaceListerExpansion
}

// generatedNetworkPolicyNamespaceLister implements the GeneratedNetworkPolicyNamespaceLister
// interface.
type generatedNetworkPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GeneratedNetworkPolicies in the indexer for a given namespace.
func (s generatedNetworkPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.GeneratedNetworkPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.GeneratedNetworkPolicy))
	})
	return ret, err
}

// Get retrieves the GeneratedNetworkPolicy from the indexer for a given namespace and name.
func (s generatedNetworkPolicyNamespaceLister) Get(name string) (*v1beta1.GeneratedNetworkPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("generatednetworkpolicy"), name)
	}
	return obj.(*v1beta1.GeneratedNetworkPolicy), nil
}
