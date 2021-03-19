/*
Copyright 2020 The Kubernetes Authors.

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

package resource

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	autoscalingv1 "k8s.io/api/autoscaling/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Object must be implemented by all resources served by the apiserver.
type Object interface {
	// Object allows the apiserver libraries to operate on the Object
	runtime.Object

	// GetObjectMeta returns the object meta reference.
	GetObjectMeta() *metav1.ObjectMeta

	// Scoper is used to qualify the resource as either namespace scoped or non-namespace scoped.
	rest.Scoper

	// New returns a new instance of the resource -- e.g. &v1.Deployment{}
	New() runtime.Object

	// NewList return a new list instance of the resource -- e.g. &v1.DeploymentList{}
	NewList() runtime.Object

	// GetGroupVersionResource returns the GroupVersionResource for this resource.  The resource should
	// be the all lowercase and pluralized kind.s
	GetGroupVersionResource() schema.GroupVersionResource

	// IsStorageVersion returns true if the object is also the internal version -- i.e. is the type defined
	// for the API group an alias to this object.
	// If false, the resource is expected to implement MultiVersionObject interface.
	IsStorageVersion() bool
}

// ObjectList must be implemented by all resources' list object.
type ObjectList interface {
	// Object allows the apiserver libraries to operate on the Object
	runtime.Object

	// GetListMeta returns the list meta reference.
	GetListMeta() *metav1.ListMeta
}

// MultiVersionObject should be implemented if the resource is not storage version and has multiple versions serving
// at the server.
type MultiVersionObject interface {
	// NewStorageVersionObject returns a new empty instance of storage version.
	NewStorageVersionObject() runtime.Object

	// ConvertToStorageVersion receives an new instance of storage version object as the conversion target
	// and overwrites it to the equal form of the current resource version.
	ConvertToStorageVersion(storageObj runtime.Object) error

	// ConvertFromStorageVersion receives an instance of storage version as the conversion source and
	// in-place mutates the current object to the equal form of the storage version object.
	ConvertFromStorageVersion(storageObj runtime.Object) error
}

// ObjectWithStatusSubResource defines an interface for getting and setting the status sub-resource for a resource.
type ObjectWithStatusSubResource interface {
	Object
	GetStatus() (statusSubResource StatusSubResource)
}

// ObjectWithScaleSubResource defines an interface for getting and setting the scale sub-resource for a resource.
type ObjectWithScaleSubResource interface {
	Object
	SetScale(scaleSubResource *autoscalingv1.Scale)
	GetScale() (scaleSubResource *autoscalingv1.Scale)
}
