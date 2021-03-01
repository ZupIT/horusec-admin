// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/tiagoangelozup/horusec-admin/pkg/api/install/v1alpha1"
	scheme "github.com/tiagoangelozup/horusec-admin/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HorusecManagerSpecsGetter has a method to return a HorusecManagerSpecInterface.
// A group's client should implement this interface.
type HorusecManagerSpecsGetter interface {
	HorusecManagerSpecs(namespace string) HorusecManagerSpecInterface
}

// HorusecManagerSpecInterface has methods to work with HorusecManagerSpec resources.
type HorusecManagerSpecInterface interface {
	Create(ctx context.Context, horusecManagerSpec *v1alpha1.HorusecManagerSpec, opts v1.CreateOptions) (*v1alpha1.HorusecManagerSpec, error)
	Update(ctx context.Context, horusecManagerSpec *v1alpha1.HorusecManagerSpec, opts v1.UpdateOptions) (*v1alpha1.HorusecManagerSpec, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.HorusecManagerSpec, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.HorusecManagerSpecList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HorusecManagerSpec, err error)
	HorusecManagerSpecExpansion
}

// horusecManagerSpecs implements HorusecManagerSpecInterface
type horusecManagerSpecs struct {
	client rest.Interface
	ns     string
}

// newHorusecManagerSpecs returns a HorusecManagerSpecs
func newHorusecManagerSpecs(c *InstallV1alpha1Client, namespace string) *horusecManagerSpecs {
	return &horusecManagerSpecs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the horusecManagerSpec, and returns the corresponding horusecManagerSpec object, and an error if there is any.
func (c *horusecManagerSpecs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HorusecManagerSpec, err error) {
	result = &v1alpha1.HorusecManagerSpec{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("horusecmanagerspecs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HorusecManagerSpecs that match those selectors.
func (c *horusecManagerSpecs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HorusecManagerSpecList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.HorusecManagerSpecList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("horusecmanagerspecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested horusecManagerSpecs.
func (c *horusecManagerSpecs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("horusecmanagerspecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a horusecManagerSpec and creates it.  Returns the server's representation of the horusecManagerSpec, and an error, if there is any.
func (c *horusecManagerSpecs) Create(ctx context.Context, horusecManagerSpec *v1alpha1.HorusecManagerSpec, opts v1.CreateOptions) (result *v1alpha1.HorusecManagerSpec, err error) {
	result = &v1alpha1.HorusecManagerSpec{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("horusecmanagerspecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(horusecManagerSpec).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a horusecManagerSpec and updates it. Returns the server's representation of the horusecManagerSpec, and an error, if there is any.
func (c *horusecManagerSpecs) Update(ctx context.Context, horusecManagerSpec *v1alpha1.HorusecManagerSpec, opts v1.UpdateOptions) (result *v1alpha1.HorusecManagerSpec, err error) {
	result = &v1alpha1.HorusecManagerSpec{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("horusecmanagerspecs").
		Name(horusecManagerSpec.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(horusecManagerSpec).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the horusecManagerSpec and deletes it. Returns an error if one occurs.
func (c *horusecManagerSpecs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("horusecmanagerspecs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *horusecManagerSpecs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("horusecmanagerspecs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched horusecManagerSpec.
func (c *horusecManagerSpecs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HorusecManagerSpec, err error) {
	result = &v1alpha1.HorusecManagerSpec{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("horusecmanagerspecs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
