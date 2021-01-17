// +build !ignore_autogenerated

/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hubble) DeepCopyInto(out *Hubble) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hubble.
func (in *Hubble) DeepCopy() *Hubble {
	if in == nil {
		return nil
	}
	out := new(Hubble)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxy) DeepCopyInto(out *KubeProxy) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ServiceHost != nil {
		in, out := &in.ServiceHost, &out.ServiceHost
		*out = new(string)
		**out = **in
	}
	if in.ServicePort != nil {
		in, out := &in.ServicePort, &out.ServicePort
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxy.
func (in *KubeProxy) DeepCopy() *KubeProxy {
	if in == nil {
		return nil
	}
	out := new(KubeProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfig) DeepCopyInto(out *NetworkConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Debug != nil {
		in, out := &in.Debug, &out.Debug
		*out = new(bool)
		**out = **in
	}
	if in.PSPEnabled != nil {
		in, out := &in.PSPEnabled, &out.PSPEnabled
		*out = new(bool)
		**out = **in
	}
	if in.KubeProxy != nil {
		in, out := &in.KubeProxy, &out.KubeProxy
		*out = new(KubeProxy)
		(*in).DeepCopyInto(*out)
	}
	if in.Hubble != nil {
		in, out := &in.Hubble, &out.Hubble
		*out = new(Hubble)
		**out = **in
	}
	if in.TunnelMode != nil {
		in, out := &in.TunnelMode, &out.TunnelMode
		*out = new(TunnelMode)
		**out = **in
	}
	if in.Store != nil {
		in, out := &in.Store, &out.Store
		*out = new(Store)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfig.
func (in *NetworkConfig) DeepCopy() *NetworkConfig {
	if in == nil {
		return nil
	}
	out := new(NetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Nodeport) DeepCopyInto(out *Nodeport) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Nodeport.
func (in *Nodeport) DeepCopy() *Nodeport {
	if in == nil {
		return nil
	}
	out := new(Nodeport)
	in.DeepCopyInto(out)
	return out
}
