/*


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
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/tilt-dev/tilt-apiserver/pkg/apis/core/v1alpha1"
	corev1alpha1 "github.com/tilt-dev/tilt-apiserver/pkg/generated/applyconfiguration/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManifests implements ManifestInterface
type FakeManifests struct {
	Fake *FakeCoreV1alpha1
}

var manifestsResource = v1alpha1.SchemeGroupVersion.WithResource("manifests")

var manifestsKind = v1alpha1.SchemeGroupVersion.WithKind("Manifest")

// Get takes name of the manifest, and returns the corresponding manifest object, and an error if there is any.
func (c *FakeManifests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Manifest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(manifestsResource, name), &v1alpha1.Manifest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Manifest), err
}

// List takes label and field selectors, and returns the list of Manifests that match those selectors.
func (c *FakeManifests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManifestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(manifestsResource, manifestsKind, opts), &v1alpha1.ManifestList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManifestList{ListMeta: obj.(*v1alpha1.ManifestList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManifestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested manifests.
func (c *FakeManifests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(manifestsResource, opts))
}

// Create takes the representation of a manifest and creates it.  Returns the server's representation of the manifest, and an error, if there is any.
func (c *FakeManifests) Create(ctx context.Context, manifest *v1alpha1.Manifest, opts v1.CreateOptions) (result *v1alpha1.Manifest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(manifestsResource, manifest), &v1alpha1.Manifest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Manifest), err
}

// Update takes the representation of a manifest and updates it. Returns the server's representation of the manifest, and an error, if there is any.
func (c *FakeManifests) Update(ctx context.Context, manifest *v1alpha1.Manifest, opts v1.UpdateOptions) (result *v1alpha1.Manifest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(manifestsResource, manifest), &v1alpha1.Manifest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Manifest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManifests) UpdateStatus(ctx context.Context, manifest *v1alpha1.Manifest, opts v1.UpdateOptions) (*v1alpha1.Manifest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(manifestsResource, "status", manifest), &v1alpha1.Manifest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Manifest), err
}

// Delete takes name of the manifest and deletes it. Returns an error if one occurs.
func (c *FakeManifests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(manifestsResource, name, opts), &v1alpha1.Manifest{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManifests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(manifestsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManifestList{})
	return err
}

// Patch applies the patch and returns the patched manifest.
func (c *FakeManifests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Manifest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(manifestsResource, name, pt, data, subresources...), &v1alpha1.Manifest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Manifest), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied manifest.
func (c *FakeManifests) Apply(ctx context.Context, manifest *corev1alpha1.ManifestApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Manifest, err error) {
	if manifest == nil {
		return nil, fmt.Errorf("manifest provided to Apply must not be nil")
	}
	data, err := json.Marshal(manifest)
	if err != nil {
		return nil, err
	}
	name := manifest.Name
	if name == nil {
		return nil, fmt.Errorf("manifest.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(manifestsResource, *name, types.ApplyPatchType, data), &v1alpha1.Manifest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Manifest), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeManifests) ApplyStatus(ctx context.Context, manifest *corev1alpha1.ManifestApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Manifest, err error) {
	if manifest == nil {
		return nil, fmt.Errorf("manifest provided to Apply must not be nil")
	}
	data, err := json.Marshal(manifest)
	if err != nil {
		return nil, err
	}
	name := manifest.Name
	if name == nil {
		return nil, fmt.Errorf("manifest.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(manifestsResource, *name, types.ApplyPatchType, data, "status"), &v1alpha1.Manifest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Manifest), err
}
