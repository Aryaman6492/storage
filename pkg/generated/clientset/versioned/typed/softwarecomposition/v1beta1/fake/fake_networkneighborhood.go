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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/Aryaman6492/storage/pkg/apis/softwarecomposition/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkNeighborhoods implements NetworkNeighborhoodInterface
type FakeNetworkNeighborhoods struct {
	Fake *FakeSpdxV1beta1
	ns   string
}

var networkneighborhoodsResource = v1beta1.SchemeGroupVersion.WithResource("networkneighborhoods")

var networkneighborhoodsKind = v1beta1.SchemeGroupVersion.WithKind("NetworkNeighborhood")

// Get takes name of the networkNeighborhood, and returns the corresponding networkNeighborhood object, and an error if there is any.
func (c *FakeNetworkNeighborhoods) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.NetworkNeighborhood, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkneighborhoodsResource, c.ns, name), &v1beta1.NetworkNeighborhood{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkNeighborhood), err
}

// List takes label and field selectors, and returns the list of NetworkNeighborhoods that match those selectors.
func (c *FakeNetworkNeighborhoods) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.NetworkNeighborhoodList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkneighborhoodsResource, networkneighborhoodsKind, c.ns, opts), &v1beta1.NetworkNeighborhoodList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.NetworkNeighborhoodList{ListMeta: obj.(*v1beta1.NetworkNeighborhoodList).ListMeta}
	for _, item := range obj.(*v1beta1.NetworkNeighborhoodList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkNeighborhoods.
func (c *FakeNetworkNeighborhoods) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkneighborhoodsResource, c.ns, opts))

}

// Create takes the representation of a networkNeighborhood and creates it.  Returns the server's representation of the networkNeighborhood, and an error, if there is any.
func (c *FakeNetworkNeighborhoods) Create(ctx context.Context, networkNeighborhood *v1beta1.NetworkNeighborhood, opts v1.CreateOptions) (result *v1beta1.NetworkNeighborhood, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkneighborhoodsResource, c.ns, networkNeighborhood), &v1beta1.NetworkNeighborhood{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkNeighborhood), err
}

// Update takes the representation of a networkNeighborhood and updates it. Returns the server's representation of the networkNeighborhood, and an error, if there is any.
func (c *FakeNetworkNeighborhoods) Update(ctx context.Context, networkNeighborhood *v1beta1.NetworkNeighborhood, opts v1.UpdateOptions) (result *v1beta1.NetworkNeighborhood, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkneighborhoodsResource, c.ns, networkNeighborhood), &v1beta1.NetworkNeighborhood{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkNeighborhood), err
}

// Delete takes name of the networkNeighborhood and deletes it. Returns an error if one occurs.
func (c *FakeNetworkNeighborhoods) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkneighborhoodsResource, c.ns, name, opts), &v1beta1.NetworkNeighborhood{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkNeighborhoods) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkneighborhoodsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.NetworkNeighborhoodList{})
	return err
}

// Patch applies the patch and returns the patched networkNeighborhood.
func (c *FakeNetworkNeighborhoods) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NetworkNeighborhood, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkneighborhoodsResource, c.ns, name, pt, data, subresources...), &v1beta1.NetworkNeighborhood{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkNeighborhood), err
}
