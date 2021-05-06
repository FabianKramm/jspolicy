/*
Copyright 2021 Loft Labs Inc.

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

	v1beta1 "github.com/loft-sh/jspolicy/pkg/apis/policy/v1beta1"
	scheme "github.com/loft-sh/jspolicy/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// JsPoliciesGetter has a method to return a JsPolicyInterface.
// A group's client should implement this interface.
type JsPoliciesGetter interface {
	JsPolicies() JsPolicyInterface
}

// JsPolicyInterface has methods to work with JsPolicy resources.
type JsPolicyInterface interface {
	Create(ctx context.Context, jsPolicy *v1beta1.JsPolicy, opts v1.CreateOptions) (*v1beta1.JsPolicy, error)
	Update(ctx context.Context, jsPolicy *v1beta1.JsPolicy, opts v1.UpdateOptions) (*v1beta1.JsPolicy, error)
	UpdateStatus(ctx context.Context, jsPolicy *v1beta1.JsPolicy, opts v1.UpdateOptions) (*v1beta1.JsPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.JsPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.JsPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.JsPolicy, err error)
	JsPolicyExpansion
}

// jsPolicies implements JsPolicyInterface
type jsPolicies struct {
	client rest.Interface
}

// newJsPolicies returns a JsPolicies
func newJsPolicies(c *PolicyV1beta1Client) *jsPolicies {
	return &jsPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the jsPolicy, and returns the corresponding jsPolicy object, and an error if there is any.
func (c *jsPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.JsPolicy, err error) {
	result = &v1beta1.JsPolicy{}
	err = c.client.Get().
		Resource("jspolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of JsPolicies that match those selectors.
func (c *jsPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.JsPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.JsPolicyList{}
	err = c.client.Get().
		Resource("jspolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested jsPolicies.
func (c *jsPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("jspolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a jsPolicy and creates it.  Returns the server's representation of the jsPolicy, and an error, if there is any.
func (c *jsPolicies) Create(ctx context.Context, jsPolicy *v1beta1.JsPolicy, opts v1.CreateOptions) (result *v1beta1.JsPolicy, err error) {
	result = &v1beta1.JsPolicy{}
	err = c.client.Post().
		Resource("jspolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jsPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a jsPolicy and updates it. Returns the server's representation of the jsPolicy, and an error, if there is any.
func (c *jsPolicies) Update(ctx context.Context, jsPolicy *v1beta1.JsPolicy, opts v1.UpdateOptions) (result *v1beta1.JsPolicy, err error) {
	result = &v1beta1.JsPolicy{}
	err = c.client.Put().
		Resource("jspolicies").
		Name(jsPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jsPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *jsPolicies) UpdateStatus(ctx context.Context, jsPolicy *v1beta1.JsPolicy, opts v1.UpdateOptions) (result *v1beta1.JsPolicy, err error) {
	result = &v1beta1.JsPolicy{}
	err = c.client.Put().
		Resource("jspolicies").
		Name(jsPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jsPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the jsPolicy and deletes it. Returns an error if one occurs.
func (c *jsPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("jspolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *jsPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("jspolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched jsPolicy.
func (c *jsPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.JsPolicy, err error) {
	result = &v1beta1.JsPolicy{}
	err = c.client.Patch(pt).
		Resource("jspolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}