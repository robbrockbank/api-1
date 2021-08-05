// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	scheme "github.com/tigera/api/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// StagedKubernetesNetworkPoliciesGetter has a method to return a StagedKubernetesNetworkPolicyInterface.
// A group's client should implement this interface.
type StagedKubernetesNetworkPoliciesGetter interface {
	StagedKubernetesNetworkPolicies(namespace string) StagedKubernetesNetworkPolicyInterface
}

// StagedKubernetesNetworkPolicyInterface has methods to work with StagedKubernetesNetworkPolicy resources.
type StagedKubernetesNetworkPolicyInterface interface {
	Create(ctx context.Context, stagedKubernetesNetworkPolicy *v3.StagedKubernetesNetworkPolicy, opts v1.CreateOptions) (*v3.StagedKubernetesNetworkPolicy, error)
	Update(ctx context.Context, stagedKubernetesNetworkPolicy *v3.StagedKubernetesNetworkPolicy, opts v1.UpdateOptions) (*v3.StagedKubernetesNetworkPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.StagedKubernetesNetworkPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.StagedKubernetesNetworkPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.StagedKubernetesNetworkPolicy, err error)
	StagedKubernetesNetworkPolicyExpansion
}

// stagedKubernetesNetworkPolicies implements StagedKubernetesNetworkPolicyInterface
type stagedKubernetesNetworkPolicies struct {
	client rest.Interface
	ns     string
}

// newStagedKubernetesNetworkPolicies returns a StagedKubernetesNetworkPolicies
func newStagedKubernetesNetworkPolicies(c *ProjectcalicoV3Client, namespace string) *stagedKubernetesNetworkPolicies {
	return &stagedKubernetesNetworkPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the stagedKubernetesNetworkPolicy, and returns the corresponding stagedKubernetesNetworkPolicy object, and an error if there is any.
func (c *stagedKubernetesNetworkPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.StagedKubernetesNetworkPolicy, err error) {
	result = &v3.StagedKubernetesNetworkPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("stagedkubernetesnetworkpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StagedKubernetesNetworkPolicies that match those selectors.
func (c *stagedKubernetesNetworkPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v3.StagedKubernetesNetworkPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.StagedKubernetesNetworkPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("stagedkubernetesnetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested stagedKubernetesNetworkPolicies.
func (c *stagedKubernetesNetworkPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("stagedkubernetesnetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a stagedKubernetesNetworkPolicy and creates it.  Returns the server's representation of the stagedKubernetesNetworkPolicy, and an error, if there is any.
func (c *stagedKubernetesNetworkPolicies) Create(ctx context.Context, stagedKubernetesNetworkPolicy *v3.StagedKubernetesNetworkPolicy, opts v1.CreateOptions) (result *v3.StagedKubernetesNetworkPolicy, err error) {
	result = &v3.StagedKubernetesNetworkPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("stagedkubernetesnetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(stagedKubernetesNetworkPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a stagedKubernetesNetworkPolicy and updates it. Returns the server's representation of the stagedKubernetesNetworkPolicy, and an error, if there is any.
func (c *stagedKubernetesNetworkPolicies) Update(ctx context.Context, stagedKubernetesNetworkPolicy *v3.StagedKubernetesNetworkPolicy, opts v1.UpdateOptions) (result *v3.StagedKubernetesNetworkPolicy, err error) {
	result = &v3.StagedKubernetesNetworkPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("stagedkubernetesnetworkpolicies").
		Name(stagedKubernetesNetworkPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(stagedKubernetesNetworkPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the stagedKubernetesNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *stagedKubernetesNetworkPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("stagedkubernetesnetworkpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *stagedKubernetesNetworkPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("stagedkubernetesnetworkpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched stagedKubernetesNetworkPolicy.
func (c *stagedKubernetesNetworkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.StagedKubernetesNetworkPolicy, err error) {
	result = &v3.StagedKubernetesNetworkPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("stagedkubernetesnetworkpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}