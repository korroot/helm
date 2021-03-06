/*
Copyright 2017 The Kubernetes Authors All rights reserved.
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

package chartutil

import (
	"k8s.io/apimachinery/pkg/version"
	tversion "k8s.io/helm/pkg/proto/hapi/version"
)

// DefaultVersionSet is the default version set, which includes only Core V1 ("v1").
var DefaultVersionSet = NewVersionSet("v1")

// Capabilities describes the capabilities of the Kubernetes cluster that Tiller is attached to.
type Capabilities struct {
	// List of all supported API versions
	APIVersions VersionSet
	// KubeVerison is the Kubernetes version
	KubeVersion *version.Info
	// TillerVersion is the Tiller version
	//
	// This always comes from pkg/version.GetVersionProto().
	TillerVersion *tversion.Version
}

// VersionSet is a set of Kubernetes API versions.
type VersionSet map[string]interface{}

// NewVersionSet creates a new version set from a list of strings.
func NewVersionSet(apiVersions ...string) VersionSet {
	vs := VersionSet{}
	for _, v := range apiVersions {
		vs[v] = struct{}{}
	}
	return vs
}

// Has returns true if the version string is in the set.
//
//	vs.Has("extensions/v1beta1")
func (v VersionSet) Has(apiVersion string) bool {
	_, ok := v[apiVersion]
	return ok
}
