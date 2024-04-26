//go:build !ignore_autogenerated

/*
Copyright 2024.

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

package v1alpha1

import (
	"github.com/nukleros/operator-builder-tools/pkg/status"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapability) DeepCopyInto(out *CertificatesCapability) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapability.
func (in *CertificatesCapability) DeepCopy() *CertificatesCapability {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificatesCapability) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilityList) DeepCopyInto(out *CertificatesCapabilityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertificatesCapability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilityList.
func (in *CertificatesCapabilityList) DeepCopy() *CertificatesCapabilityList {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificatesCapabilityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpec) DeepCopyInto(out *CertificatesCapabilitySpec) {
	*out = *in
	out.Aws = in.Aws
	out.CertManager = in.CertManager
	out.TrustManager = in.TrustManager
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpec.
func (in *CertificatesCapabilitySpec) DeepCopy() *CertificatesCapabilitySpec {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecAws) DeepCopyInto(out *CertificatesCapabilitySpecAws) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecAws.
func (in *CertificatesCapabilitySpecAws) DeepCopy() *CertificatesCapabilitySpecAws {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecAws)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecCertManager) DeepCopyInto(out *CertificatesCapabilitySpecCertManager) {
	*out = *in
	out.Injector = in.Injector
	out.Controller = in.Controller
	out.Webhook = in.Webhook
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecCertManager.
func (in *CertificatesCapabilitySpecCertManager) DeepCopy() *CertificatesCapabilitySpecCertManager {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecCertManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecCertManagerController) DeepCopyInto(out *CertificatesCapabilitySpecCertManagerController) {
	*out = *in
	out.Resources = in.Resources
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecCertManagerController.
func (in *CertificatesCapabilitySpecCertManagerController) DeepCopy() *CertificatesCapabilitySpecCertManagerController {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecCertManagerController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecCertManagerControllerResources) DeepCopyInto(out *CertificatesCapabilitySpecCertManagerControllerResources) {
	*out = *in
	out.Requests = in.Requests
	out.Limits = in.Limits
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecCertManagerControllerResources.
func (in *CertificatesCapabilitySpecCertManagerControllerResources) DeepCopy() *CertificatesCapabilitySpecCertManagerControllerResources {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecCertManagerControllerResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecCertManagerControllerResourcesLimits) DeepCopyInto(out *CertificatesCapabilitySpecCertManagerControllerResourcesLimits) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecCertManagerControllerResourcesLimits.
func (in *CertificatesCapabilitySpecCertManagerControllerResourcesLimits) DeepCopy() *CertificatesCapabilitySpecCertManagerControllerResourcesLimits {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecCertManagerControllerResourcesLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecCertManagerControllerResourcesRequests) DeepCopyInto(out *CertificatesCapabilitySpecCertManagerControllerResourcesRequests) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecCertManagerControllerResourcesRequests.
func (in *CertificatesCapabilitySpecCertManagerControllerResourcesRequests) DeepCopy() *CertificatesCapabilitySpecCertManagerControllerResourcesRequests {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecCertManagerControllerResourcesRequests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecCertManagerInjector) DeepCopyInto(out *CertificatesCapabilitySpecCertManagerInjector) {
	*out = *in
	out.Resources = in.Resources
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecCertManagerInjector.
func (in *CertificatesCapabilitySpecCertManagerInjector) DeepCopy() *CertificatesCapabilitySpecCertManagerInjector {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecCertManagerInjector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecCertManagerInjectorResources) DeepCopyInto(out *CertificatesCapabilitySpecCertManagerInjectorResources) {
	*out = *in
	out.Requests = in.Requests
	out.Limits = in.Limits
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecCertManagerInjectorResources.
func (in *CertificatesCapabilitySpecCertManagerInjectorResources) DeepCopy() *CertificatesCapabilitySpecCertManagerInjectorResources {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecCertManagerInjectorResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecCertManagerInjectorResourcesLimits) DeepCopyInto(out *CertificatesCapabilitySpecCertManagerInjectorResourcesLimits) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecCertManagerInjectorResourcesLimits.
func (in *CertificatesCapabilitySpecCertManagerInjectorResourcesLimits) DeepCopy() *CertificatesCapabilitySpecCertManagerInjectorResourcesLimits {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecCertManagerInjectorResourcesLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecCertManagerInjectorResourcesRequests) DeepCopyInto(out *CertificatesCapabilitySpecCertManagerInjectorResourcesRequests) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecCertManagerInjectorResourcesRequests.
func (in *CertificatesCapabilitySpecCertManagerInjectorResourcesRequests) DeepCopy() *CertificatesCapabilitySpecCertManagerInjectorResourcesRequests {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecCertManagerInjectorResourcesRequests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecCertManagerWebhook) DeepCopyInto(out *CertificatesCapabilitySpecCertManagerWebhook) {
	*out = *in
	out.Resources = in.Resources
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecCertManagerWebhook.
func (in *CertificatesCapabilitySpecCertManagerWebhook) DeepCopy() *CertificatesCapabilitySpecCertManagerWebhook {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecCertManagerWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecCertManagerWebhookResources) DeepCopyInto(out *CertificatesCapabilitySpecCertManagerWebhookResources) {
	*out = *in
	out.Requests = in.Requests
	out.Limits = in.Limits
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecCertManagerWebhookResources.
func (in *CertificatesCapabilitySpecCertManagerWebhookResources) DeepCopy() *CertificatesCapabilitySpecCertManagerWebhookResources {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecCertManagerWebhookResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecCertManagerWebhookResourcesLimits) DeepCopyInto(out *CertificatesCapabilitySpecCertManagerWebhookResourcesLimits) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecCertManagerWebhookResourcesLimits.
func (in *CertificatesCapabilitySpecCertManagerWebhookResourcesLimits) DeepCopy() *CertificatesCapabilitySpecCertManagerWebhookResourcesLimits {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecCertManagerWebhookResourcesLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecCertManagerWebhookResourcesRequests) DeepCopyInto(out *CertificatesCapabilitySpecCertManagerWebhookResourcesRequests) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecCertManagerWebhookResourcesRequests.
func (in *CertificatesCapabilitySpecCertManagerWebhookResourcesRequests) DeepCopy() *CertificatesCapabilitySpecCertManagerWebhookResourcesRequests {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecCertManagerWebhookResourcesRequests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecTrustManager) DeepCopyInto(out *CertificatesCapabilitySpecTrustManager) {
	*out = *in
	out.Controller = in.Controller
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecTrustManager.
func (in *CertificatesCapabilitySpecTrustManager) DeepCopy() *CertificatesCapabilitySpecTrustManager {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecTrustManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecTrustManagerController) DeepCopyInto(out *CertificatesCapabilitySpecTrustManagerController) {
	*out = *in
	out.Resources = in.Resources
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecTrustManagerController.
func (in *CertificatesCapabilitySpecTrustManagerController) DeepCopy() *CertificatesCapabilitySpecTrustManagerController {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecTrustManagerController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecTrustManagerControllerResources) DeepCopyInto(out *CertificatesCapabilitySpecTrustManagerControllerResources) {
	*out = *in
	out.Requests = in.Requests
	out.Limits = in.Limits
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecTrustManagerControllerResources.
func (in *CertificatesCapabilitySpecTrustManagerControllerResources) DeepCopy() *CertificatesCapabilitySpecTrustManagerControllerResources {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecTrustManagerControllerResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecTrustManagerControllerResourcesLimits) DeepCopyInto(out *CertificatesCapabilitySpecTrustManagerControllerResourcesLimits) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecTrustManagerControllerResourcesLimits.
func (in *CertificatesCapabilitySpecTrustManagerControllerResourcesLimits) DeepCopy() *CertificatesCapabilitySpecTrustManagerControllerResourcesLimits {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecTrustManagerControllerResourcesLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilitySpecTrustManagerControllerResourcesRequests) DeepCopyInto(out *CertificatesCapabilitySpecTrustManagerControllerResourcesRequests) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilitySpecTrustManagerControllerResourcesRequests.
func (in *CertificatesCapabilitySpecTrustManagerControllerResourcesRequests) DeepCopy() *CertificatesCapabilitySpecTrustManagerControllerResourcesRequests {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilitySpecTrustManagerControllerResourcesRequests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificatesCapabilityStatus) DeepCopyInto(out *CertificatesCapabilityStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*status.PhaseCondition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(status.PhaseCondition)
				**out = **in
			}
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]*status.ChildResource, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(status.ChildResource)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificatesCapabilityStatus.
func (in *CertificatesCapabilityStatus) DeepCopy() *CertificatesCapabilityStatus {
	if in == nil {
		return nil
	}
	out := new(CertificatesCapabilityStatus)
	in.DeepCopyInto(out)
	return out
}
