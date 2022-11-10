//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package synchronizer

import (
	"github.com/wso2/apk/adapter/internal/operator/apis/dp/v1alpha1"
	"sigs.k8s.io/gateway-api/apis/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIState) DeepCopyInto(out *APIState) {
	*out = *in
	if in.APIDefinition != nil {
		in, out := &in.APIDefinition, &out.APIDefinition
		*out = new(v1alpha1.API)
		(*in).DeepCopyInto(*out)
	}
	if in.ProdHTTPRoute != nil {
		in, out := &in.ProdHTTPRoute, &out.ProdHTTPRoute
		*out = new(v1beta1.HTTPRoute)
		(*in).DeepCopyInto(*out)
	}
	if in.SandHTTPRoute != nil {
		in, out := &in.SandHTTPRoute, &out.SandHTTPRoute
		*out = new(v1beta1.HTTPRoute)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIState.
func (in *APIState) DeepCopy() *APIState {
	if in == nil {
		return nil
	}
	out := new(APIState)
	in.DeepCopyInto(out)
	return out
}
