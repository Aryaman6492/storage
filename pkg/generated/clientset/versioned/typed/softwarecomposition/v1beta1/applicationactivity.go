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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/Aryaman6492/storage/pkg/apis/softwarecomposition/v1beta1"
	scheme "github.com/Aryaman6492/storage/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApplicationActivitiesGetter has a method to return a ApplicationActivityInterface.
// A group's client should implement this interface.
type ApplicationActivitiesGetter interface {
	ApplicationActivities(namespace string) ApplicationActivityInterface
}

// ApplicationActivityInterface has methods to work with ApplicationActivity resources.
type ApplicationActivityInterface interface {
	Create(ctx context.Context, applicationActivity *v1beta1.ApplicationActivity, opts v1.CreateOptions) (*v1beta1.ApplicationActivity, error)
	Update(ctx context.Context, applicationActivity *v1beta1.ApplicationActivity, opts v1.UpdateOptions) (*v1beta1.ApplicationActivity, error)
	UpdateStatus(ctx context.Context, applicationActivity *v1beta1.ApplicationActivity, opts v1.UpdateOptions) (*v1beta1.ApplicationActivity, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ApplicationActivity, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ApplicationActivityList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ApplicationActivity, err error)
	ApplicationActivityExpansion
}

// applicationActivities implements ApplicationActivityInterface
type applicationActivities struct {
	client rest.Interface
	ns     string
}

// newApplicationActivities returns a ApplicationActivities
func newApplicationActivities(c *SpdxV1beta1Client, namespace string) *applicationActivities {
	return &applicationActivities{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the applicationActivity, and returns the corresponding applicationActivity object, and an error if there is any.
func (c *applicationActivities) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ApplicationActivity, err error) {
	result = &v1beta1.ApplicationActivity{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationactivities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApplicationActivities that match those selectors.
func (c *applicationActivities) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ApplicationActivityList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ApplicationActivityList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationactivities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested applicationActivities.
func (c *applicationActivities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("applicationactivities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a applicationActivity and creates it.  Returns the server's representation of the applicationActivity, and an error, if there is any.
func (c *applicationActivities) Create(ctx context.Context, applicationActivity *v1beta1.ApplicationActivity, opts v1.CreateOptions) (result *v1beta1.ApplicationActivity, err error) {
	result = &v1beta1.ApplicationActivity{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("applicationactivities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(applicationActivity).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a applicationActivity and updates it. Returns the server's representation of the applicationActivity, and an error, if there is any.
func (c *applicationActivities) Update(ctx context.Context, applicationActivity *v1beta1.ApplicationActivity, opts v1.UpdateOptions) (result *v1beta1.ApplicationActivity, err error) {
	result = &v1beta1.ApplicationActivity{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationactivities").
		Name(applicationActivity.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(applicationActivity).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *applicationActivities) UpdateStatus(ctx context.Context, applicationActivity *v1beta1.ApplicationActivity, opts v1.UpdateOptions) (result *v1beta1.ApplicationActivity, err error) {
	result = &v1beta1.ApplicationActivity{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationactivities").
		Name(applicationActivity.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(applicationActivity).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the applicationActivity and deletes it. Returns an error if one occurs.
func (c *applicationActivities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationactivities").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *applicationActivities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationactivities").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched applicationActivity.
func (c *applicationActivities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ApplicationActivity, err error) {
	result = &v1beta1.ApplicationActivity{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("applicationactivities").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
