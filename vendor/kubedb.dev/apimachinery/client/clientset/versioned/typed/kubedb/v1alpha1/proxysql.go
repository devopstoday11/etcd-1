/*
Copyright 2019 The KubeDB Authors.

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

package v1alpha1

import (
	"time"

	v1alpha1 "kubedb.dev/apimachinery/apis/kubedb/v1alpha1"
	scheme "kubedb.dev/apimachinery/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ProxySQLsGetter has a method to return a ProxySQLInterface.
// A group's client should implement this interface.
type ProxySQLsGetter interface {
	ProxySQLs(namespace string) ProxySQLInterface
}

// ProxySQLInterface has methods to work with ProxySQL resources.
type ProxySQLInterface interface {
	Create(*v1alpha1.ProxySQL) (*v1alpha1.ProxySQL, error)
	Update(*v1alpha1.ProxySQL) (*v1alpha1.ProxySQL, error)
	UpdateStatus(*v1alpha1.ProxySQL) (*v1alpha1.ProxySQL, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ProxySQL, error)
	List(opts v1.ListOptions) (*v1alpha1.ProxySQLList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ProxySQL, err error)
	ProxySQLExpansion
}

// proxySQLs implements ProxySQLInterface
type proxySQLs struct {
	client rest.Interface
	ns     string
}

// newProxySQLs returns a ProxySQLs
func newProxySQLs(c *KubedbV1alpha1Client, namespace string) *proxySQLs {
	return &proxySQLs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the proxySQL, and returns the corresponding proxySQL object, and an error if there is any.
func (c *proxySQLs) Get(name string, options v1.GetOptions) (result *v1alpha1.ProxySQL, err error) {
	result = &v1alpha1.ProxySQL{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("proxysqls").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ProxySQLs that match those selectors.
func (c *proxySQLs) List(opts v1.ListOptions) (result *v1alpha1.ProxySQLList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ProxySQLList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("proxysqls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested proxySQLs.
func (c *proxySQLs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("proxysqls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a proxySQL and creates it.  Returns the server's representation of the proxySQL, and an error, if there is any.
func (c *proxySQLs) Create(proxySQL *v1alpha1.ProxySQL) (result *v1alpha1.ProxySQL, err error) {
	result = &v1alpha1.ProxySQL{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("proxysqls").
		Body(proxySQL).
		Do().
		Into(result)
	return
}

// Update takes the representation of a proxySQL and updates it. Returns the server's representation of the proxySQL, and an error, if there is any.
func (c *proxySQLs) Update(proxySQL *v1alpha1.ProxySQL) (result *v1alpha1.ProxySQL, err error) {
	result = &v1alpha1.ProxySQL{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("proxysqls").
		Name(proxySQL.Name).
		Body(proxySQL).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *proxySQLs) UpdateStatus(proxySQL *v1alpha1.ProxySQL) (result *v1alpha1.ProxySQL, err error) {
	result = &v1alpha1.ProxySQL{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("proxysqls").
		Name(proxySQL.Name).
		SubResource("status").
		Body(proxySQL).
		Do().
		Into(result)
	return
}

// Delete takes name of the proxySQL and deletes it. Returns an error if one occurs.
func (c *proxySQLs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("proxysqls").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *proxySQLs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("proxysqls").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched proxySQL.
func (c *proxySQLs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ProxySQL, err error) {
	result = &v1alpha1.ProxySQL{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("proxysqls").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
