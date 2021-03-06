/*
Copyright wbsnail.com.

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

	stablewbsnailcomv1 "github.com/wbsnail/articles/archive/dive-into-kubernetes-informer/11-crd-informer/api/stable.wbsnail.com/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRabbits implements RabbitInterface
type FakeRabbits struct {
	Fake *FakeStableV1
	ns   string
}

var rabbitsResource = schema.GroupVersionResource{Group: "stable.wbsnail.com", Version: "v1", Resource: "rabbits"}

var rabbitsKind = schema.GroupVersionKind{Group: "stable.wbsnail.com", Version: "v1", Kind: "Rabbit"}

// Get takes name of the rabbit, and returns the corresponding rabbit object, and an error if there is any.
func (c *FakeRabbits) Get(ctx context.Context, name string, options v1.GetOptions) (result *stablewbsnailcomv1.Rabbit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rabbitsResource, c.ns, name), &stablewbsnailcomv1.Rabbit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*stablewbsnailcomv1.Rabbit), err
}

// List takes label and field selectors, and returns the list of Rabbits that match those selectors.
func (c *FakeRabbits) List(ctx context.Context, opts v1.ListOptions) (result *stablewbsnailcomv1.RabbitList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rabbitsResource, rabbitsKind, c.ns, opts), &stablewbsnailcomv1.RabbitList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &stablewbsnailcomv1.RabbitList{ListMeta: obj.(*stablewbsnailcomv1.RabbitList).ListMeta}
	for _, item := range obj.(*stablewbsnailcomv1.RabbitList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rabbits.
func (c *FakeRabbits) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rabbitsResource, c.ns, opts))

}

// Create takes the representation of a rabbit and creates it.  Returns the server's representation of the rabbit, and an error, if there is any.
func (c *FakeRabbits) Create(ctx context.Context, rabbit *stablewbsnailcomv1.Rabbit, opts v1.CreateOptions) (result *stablewbsnailcomv1.Rabbit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rabbitsResource, c.ns, rabbit), &stablewbsnailcomv1.Rabbit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*stablewbsnailcomv1.Rabbit), err
}

// Update takes the representation of a rabbit and updates it. Returns the server's representation of the rabbit, and an error, if there is any.
func (c *FakeRabbits) Update(ctx context.Context, rabbit *stablewbsnailcomv1.Rabbit, opts v1.UpdateOptions) (result *stablewbsnailcomv1.Rabbit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rabbitsResource, c.ns, rabbit), &stablewbsnailcomv1.Rabbit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*stablewbsnailcomv1.Rabbit), err
}

// Delete takes name of the rabbit and deletes it. Returns an error if one occurs.
func (c *FakeRabbits) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rabbitsResource, c.ns, name), &stablewbsnailcomv1.Rabbit{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRabbits) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rabbitsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &stablewbsnailcomv1.RabbitList{})
	return err
}

// Patch applies the patch and returns the patched rabbit.
func (c *FakeRabbits) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *stablewbsnailcomv1.Rabbit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rabbitsResource, c.ns, name, pt, data, subresources...), &stablewbsnailcomv1.Rabbit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*stablewbsnailcomv1.Rabbit), err
}
