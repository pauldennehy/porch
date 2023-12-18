// Copyright 2023 The kpt and Nephio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/GoogleContainerTools/kpt/porch/api/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Functions returns a FunctionInformer.
	Functions() FunctionInformer
	// Packages returns a PackageInformer.
	Packages() PackageInformer
	// PackageRevisions returns a PackageRevisionInformer.
	PackageRevisions() PackageRevisionInformer
	// PackageRevisionResources returns a PackageRevisionResourcesInformer.
	PackageRevisionResources() PackageRevisionResourcesInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Functions returns a FunctionInformer.
func (v *version) Functions() FunctionInformer {
	return &functionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Packages returns a PackageInformer.
func (v *version) Packages() PackageInformer {
	return &packageInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PackageRevisions returns a PackageRevisionInformer.
func (v *version) PackageRevisions() PackageRevisionInformer {
	return &packageRevisionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PackageRevisionResources returns a PackageRevisionResourcesInformer.
func (v *version) PackageRevisionResources() PackageRevisionResourcesInformer {
	return &packageRevisionResourcesInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
