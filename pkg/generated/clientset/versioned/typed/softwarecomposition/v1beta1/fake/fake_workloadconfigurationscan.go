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

// FakeWorkloadConfigurationScans implements WorkloadConfigurationScanInterface
type FakeWorkloadConfigurationScans struct {
	Fake *FakeSpdxV1beta1
	ns   string
}

var workloadconfigurationscansResource = v1beta1.SchemeGroupVersion.WithResource("workloadconfigurationscans")

var workloadconfigurationscansKind = v1beta1.SchemeGroupVersion.WithKind("WorkloadConfigurationScan")

// Get takes name of the workloadConfigurationScan, and returns the corresponding workloadConfigurationScan object, and an error if there is any.
func (c *FakeWorkloadConfigurationScans) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.WorkloadConfigurationScan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(workloadconfigurationscansResource, c.ns, name), &v1beta1.WorkloadConfigurationScan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.WorkloadConfigurationScan), err
}

// List takes label and field selectors, and returns the list of WorkloadConfigurationScans that match those selectors.
func (c *FakeWorkloadConfigurationScans) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.WorkloadConfigurationScanList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(workloadconfigurationscansResource, workloadconfigurationscansKind, c.ns, opts), &v1beta1.WorkloadConfigurationScanList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.WorkloadConfigurationScanList{ListMeta: obj.(*v1beta1.WorkloadConfigurationScanList).ListMeta}
	for _, item := range obj.(*v1beta1.WorkloadConfigurationScanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workloadConfigurationScans.
func (c *FakeWorkloadConfigurationScans) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(workloadconfigurationscansResource, c.ns, opts))

}

// Create takes the representation of a workloadConfigurationScan and creates it.  Returns the server's representation of the workloadConfigurationScan, and an error, if there is any.
func (c *FakeWorkloadConfigurationScans) Create(ctx context.Context, workloadConfigurationScan *v1beta1.WorkloadConfigurationScan, opts v1.CreateOptions) (result *v1beta1.WorkloadConfigurationScan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(workloadconfigurationscansResource, c.ns, workloadConfigurationScan), &v1beta1.WorkloadConfigurationScan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.WorkloadConfigurationScan), err
}

// Update takes the representation of a workloadConfigurationScan and updates it. Returns the server's representation of the workloadConfigurationScan, and an error, if there is any.
func (c *FakeWorkloadConfigurationScans) Update(ctx context.Context, workloadConfigurationScan *v1beta1.WorkloadConfigurationScan, opts v1.UpdateOptions) (result *v1beta1.WorkloadConfigurationScan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(workloadconfigurationscansResource, c.ns, workloadConfigurationScan), &v1beta1.WorkloadConfigurationScan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.WorkloadConfigurationScan), err
}

// Delete takes name of the workloadConfigurationScan and deletes it. Returns an error if one occurs.
func (c *FakeWorkloadConfigurationScans) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(workloadconfigurationscansResource, c.ns, name, opts), &v1beta1.WorkloadConfigurationScan{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkloadConfigurationScans) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(workloadconfigurationscansResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.WorkloadConfigurationScanList{})
	return err
}

// Patch applies the patch and returns the patched workloadConfigurationScan.
func (c *FakeWorkloadConfigurationScans) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.WorkloadConfigurationScan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(workloadconfigurationscansResource, c.ns, name, pt, data, subresources...), &v1beta1.WorkloadConfigurationScan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.WorkloadConfigurationScan), err
}
