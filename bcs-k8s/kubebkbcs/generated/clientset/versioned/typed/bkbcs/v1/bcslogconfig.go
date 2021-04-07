/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/Tencent/bk-bcs/bcs-k8s/kubebkbcs/apis/bkbcs/v1"
	scheme "github.com/Tencent/bk-bcs/bcs-k8s/kubebkbcs/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BcsLogConfigsGetter has a method to return a BcsLogConfigInterface.
// A group's client should implement this interface.
type BcsLogConfigsGetter interface {
	BcsLogConfigs(namespace string) BcsLogConfigInterface
}

// BcsLogConfigInterface has methods to work with BcsLogConfig resources.
type BcsLogConfigInterface interface {
	Create(ctx context.Context, bcsLogConfig *v1.BcsLogConfig, opts metav1.CreateOptions) (*v1.BcsLogConfig, error)
	Update(ctx context.Context, bcsLogConfig *v1.BcsLogConfig, opts metav1.UpdateOptions) (*v1.BcsLogConfig, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.BcsLogConfig, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.BcsLogConfigList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.BcsLogConfig, err error)
	BcsLogConfigExpansion
}

// bcsLogConfigs implements BcsLogConfigInterface
type bcsLogConfigs struct {
	client rest.Interface
	ns     string
}

// newBcsLogConfigs returns a BcsLogConfigs
func newBcsLogConfigs(c *BkbcsV1Client, namespace string) *bcsLogConfigs {
	return &bcsLogConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bcsLogConfig, and returns the corresponding bcsLogConfig object, and an error if there is any.
func (c *bcsLogConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.BcsLogConfig, err error) {
	result = &v1.BcsLogConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bcslogconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BcsLogConfigs that match those selectors.
func (c *bcsLogConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.BcsLogConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.BcsLogConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bcslogconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bcsLogConfigs.
func (c *bcsLogConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bcslogconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bcsLogConfig and creates it.  Returns the server's representation of the bcsLogConfig, and an error, if there is any.
func (c *bcsLogConfigs) Create(ctx context.Context, bcsLogConfig *v1.BcsLogConfig, opts metav1.CreateOptions) (result *v1.BcsLogConfig, err error) {
	result = &v1.BcsLogConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bcslogconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bcsLogConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bcsLogConfig and updates it. Returns the server's representation of the bcsLogConfig, and an error, if there is any.
func (c *bcsLogConfigs) Update(ctx context.Context, bcsLogConfig *v1.BcsLogConfig, opts metav1.UpdateOptions) (result *v1.BcsLogConfig, err error) {
	result = &v1.BcsLogConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bcslogconfigs").
		Name(bcsLogConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bcsLogConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bcsLogConfig and deletes it. Returns an error if one occurs.
func (c *bcsLogConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bcslogconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bcsLogConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bcslogconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bcsLogConfig.
func (c *bcsLogConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.BcsLogConfig, err error) {
	result = &v1.BcsLogConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bcslogconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
