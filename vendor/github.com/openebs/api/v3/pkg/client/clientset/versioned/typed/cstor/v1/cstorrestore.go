/*
Copyright 2021 The OpenEBS Authors

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

package v1

import (
	"context"
	"time"

	v1 "github.com/openebs/api/v3/pkg/apis/cstor/v1"
	scheme "github.com/openebs/api/v3/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CStorRestoresGetter has a method to return a CStorRestoreInterface.
// A group's client should implement this interface.
type CStorRestoresGetter interface {
	CStorRestores(namespace string) CStorRestoreInterface
}

// CStorRestoreInterface has methods to work with CStorRestore resources.
type CStorRestoreInterface interface {
	Create(ctx context.Context, cStorRestore *v1.CStorRestore, opts metav1.CreateOptions) (*v1.CStorRestore, error)
	Update(ctx context.Context, cStorRestore *v1.CStorRestore, opts metav1.UpdateOptions) (*v1.CStorRestore, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CStorRestore, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CStorRestoreList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CStorRestore, err error)
	CStorRestoreExpansion
}

// cStorRestores implements CStorRestoreInterface
type cStorRestores struct {
	client rest.Interface
	ns     string
}

// newCStorRestores returns a CStorRestores
func newCStorRestores(c *CstorV1Client, namespace string) *cStorRestores {
	return &cStorRestores{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cStorRestore, and returns the corresponding cStorRestore object, and an error if there is any.
func (c *cStorRestores) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CStorRestore, err error) {
	result = &v1.CStorRestore{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cstorrestores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CStorRestores that match those selectors.
func (c *cStorRestores) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CStorRestoreList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CStorRestoreList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cstorrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cStorRestores.
func (c *cStorRestores) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cstorrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cStorRestore and creates it.  Returns the server's representation of the cStorRestore, and an error, if there is any.
func (c *cStorRestores) Create(ctx context.Context, cStorRestore *v1.CStorRestore, opts metav1.CreateOptions) (result *v1.CStorRestore, err error) {
	result = &v1.CStorRestore{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cstorrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cStorRestore).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cStorRestore and updates it. Returns the server's representation of the cStorRestore, and an error, if there is any.
func (c *cStorRestores) Update(ctx context.Context, cStorRestore *v1.CStorRestore, opts metav1.UpdateOptions) (result *v1.CStorRestore, err error) {
	result = &v1.CStorRestore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cstorrestores").
		Name(cStorRestore.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cStorRestore).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cStorRestore and deletes it. Returns an error if one occurs.
func (c *cStorRestores) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cstorrestores").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cStorRestores) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cstorrestores").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cStorRestore.
func (c *cStorRestores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CStorRestore, err error) {
	result = &v1.CStorRestore{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cstorrestores").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
