// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for fed.enterprise.gloo.solo.io/v1 resources

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for FederatedAuthConfig

func (in *FederatedAuthConfig) DeepCopyInto(out *FederatedAuthConfig) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *FederatedAuthConfig) DeepCopy() *FederatedAuthConfig {
	if in == nil {
		return nil
	}
	out := new(FederatedAuthConfig)
	in.DeepCopyInto(out)
	return out
}

func (in *FederatedAuthConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *FederatedAuthConfigList) DeepCopyInto(out *FederatedAuthConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedAuthConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *FederatedAuthConfigList) DeepCopy() *FederatedAuthConfigList {
	if in == nil {
		return nil
	}
	out := new(FederatedAuthConfigList)
	in.DeepCopyInto(out)
	return out
}

func (in *FederatedAuthConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}