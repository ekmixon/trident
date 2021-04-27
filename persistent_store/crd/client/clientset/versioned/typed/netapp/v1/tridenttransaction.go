// Copyright 2021 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	scheme "github.com/netapp/trident/persistent_store/crd/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TridentTransactionsGetter has a method to return a TridentTransactionInterface.
// A group's client should implement this interface.
type TridentTransactionsGetter interface {
	TridentTransactions(namespace string) TridentTransactionInterface
}

// TridentTransactionInterface has methods to work with TridentTransaction resources.
type TridentTransactionInterface interface {
	Create(ctx context.Context, tridentTransaction *v1.TridentTransaction, opts metav1.CreateOptions) (*v1.TridentTransaction, error)
	Update(ctx context.Context, tridentTransaction *v1.TridentTransaction, opts metav1.UpdateOptions) (*v1.TridentTransaction, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.TridentTransaction, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TridentTransactionList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentTransaction, err error)
	TridentTransactionExpansion
}

// tridentTransactions implements TridentTransactionInterface
type tridentTransactions struct {
	client rest.Interface
	ns     string
}

// newTridentTransactions returns a TridentTransactions
func newTridentTransactions(c *TridentV1Client, namespace string) *tridentTransactions {
	return &tridentTransactions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tridentTransaction, and returns the corresponding tridentTransaction object, and an error if there is any.
func (c *tridentTransactions) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.TridentTransaction, err error) {
	result = &v1.TridentTransaction{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridenttransactions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TridentTransactions that match those selectors.
func (c *tridentTransactions) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TridentTransactionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TridentTransactionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridenttransactions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tridentTransactions.
func (c *tridentTransactions) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tridenttransactions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tridentTransaction and creates it.  Returns the server's representation of the tridentTransaction, and an error, if there is any.
func (c *tridentTransactions) Create(ctx context.Context, tridentTransaction *v1.TridentTransaction, opts metav1.CreateOptions) (result *v1.TridentTransaction, err error) {
	result = &v1.TridentTransaction{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tridenttransactions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentTransaction).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tridentTransaction and updates it. Returns the server's representation of the tridentTransaction, and an error, if there is any.
func (c *tridentTransactions) Update(ctx context.Context, tridentTransaction *v1.TridentTransaction, opts metav1.UpdateOptions) (result *v1.TridentTransaction, err error) {
	result = &v1.TridentTransaction{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tridenttransactions").
		Name(tridentTransaction.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentTransaction).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tridentTransaction and deletes it. Returns an error if one occurs.
func (c *tridentTransactions) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridenttransactions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tridentTransactions) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridenttransactions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tridentTransaction.
func (c *tridentTransactions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentTransaction, err error) {
	result = &v1.TridentTransaction{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tridenttransactions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
