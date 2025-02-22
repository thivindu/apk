//go:build !ignore_autogenerated

/*
 *  Copyright (c) 2023, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha3

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIProvider) DeepCopyInto(out *AIProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIProvider.
func (in *AIProvider) DeepCopy() *AIProvider {
	if in == nil {
		return nil
	}
	out := new(AIProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AIProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIProviderList) DeepCopyInto(out *AIProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AIProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIProviderList.
func (in *AIProviderList) DeepCopy() *AIProviderList {
	if in == nil {
		return nil
	}
	out := new(AIProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AIProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIProviderReference) DeepCopyInto(out *AIProviderReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIProviderReference.
func (in *AIProviderReference) DeepCopy() *AIProviderReference {
	if in == nil {
		return nil
	}
	out := new(AIProviderReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIProviderSpec) DeepCopyInto(out *AIProviderSpec) {
	*out = *in
	out.Model = in.Model
	out.RateLimitFields = in.RateLimitFields
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIProviderSpec.
func (in *AIProviderSpec) DeepCopy() *AIProviderSpec {
	if in == nil {
		return nil
	}
	out := new(AIProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIProviderStatus) DeepCopyInto(out *AIProviderStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIProviderStatus.
func (in *AIProviderStatus) DeepCopy() *AIProviderStatus {
	if in == nil {
		return nil
	}
	out := new(AIProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIRateLimit) DeepCopyInto(out *AIRateLimit) {
	*out = *in
	if in.TokenCount != nil {
		in, out := &in.TokenCount, &out.TokenCount
		*out = new(TokenCount)
		**out = **in
	}
	if in.RequestCount != nil {
		in, out := &in.RequestCount, &out.RequestCount
		*out = new(RequestCount)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIRateLimit.
func (in *AIRateLimit) DeepCopy() *AIRateLimit {
	if in == nil {
		return nil
	}
	out := new(AIRateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIRateLimitPolicy) DeepCopyInto(out *AIRateLimitPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIRateLimitPolicy.
func (in *AIRateLimitPolicy) DeepCopy() *AIRateLimitPolicy {
	if in == nil {
		return nil
	}
	out := new(AIRateLimitPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AIRateLimitPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIRateLimitPolicyList) DeepCopyInto(out *AIRateLimitPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AIRateLimitPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIRateLimitPolicyList.
func (in *AIRateLimitPolicyList) DeepCopy() *AIRateLimitPolicyList {
	if in == nil {
		return nil
	}
	out := new(AIRateLimitPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AIRateLimitPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIRateLimitPolicySpec) DeepCopyInto(out *AIRateLimitPolicySpec) {
	*out = *in
	if in.Override != nil {
		in, out := &in.Override, &out.Override
		*out = new(AIRateLimit)
		(*in).DeepCopyInto(*out)
	}
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(AIRateLimit)
		(*in).DeepCopyInto(*out)
	}
	in.TargetRef.DeepCopyInto(&out.TargetRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIRateLimitPolicySpec.
func (in *AIRateLimitPolicySpec) DeepCopy() *AIRateLimitPolicySpec {
	if in == nil {
		return nil
	}
	out := new(AIRateLimitPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIRateLimitPolicyStatus) DeepCopyInto(out *AIRateLimitPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIRateLimitPolicyStatus.
func (in *AIRateLimitPolicyStatus) DeepCopy() *AIRateLimitPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(AIRateLimitPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *API) DeepCopyInto(out *API) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new API.
func (in *API) DeepCopy() *API {
	if in == nil {
		return nil
	}
	out := new(API)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *API) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIList) DeepCopyInto(out *APIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]API, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIList.
func (in *APIList) DeepCopy() *APIList {
	if in == nil {
		return nil
	}
	out := new(APIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPolicy) DeepCopyInto(out *APIPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPolicy.
func (in *APIPolicy) DeepCopy() *APIPolicy {
	if in == nil {
		return nil
	}
	out := new(APIPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPolicyList) DeepCopyInto(out *APIPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPolicyList.
func (in *APIPolicyList) DeepCopy() *APIPolicyList {
	if in == nil {
		return nil
	}
	out := new(APIPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPolicySpec) DeepCopyInto(out *APIPolicySpec) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(PolicySpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Override != nil {
		in, out := &in.Override, &out.Override
		*out = new(PolicySpec)
		(*in).DeepCopyInto(*out)
	}
	in.TargetRef.DeepCopyInto(&out.TargetRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPolicySpec.
func (in *APIPolicySpec) DeepCopy() *APIPolicySpec {
	if in == nil {
		return nil
	}
	out := new(APIPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPolicyStatus) DeepCopyInto(out *APIPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPolicyStatus.
func (in *APIPolicyStatus) DeepCopy() *APIPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(APIPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIRateLimitPolicy) DeepCopyInto(out *APIRateLimitPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIRateLimitPolicy.
func (in *APIRateLimitPolicy) DeepCopy() *APIRateLimitPolicy {
	if in == nil {
		return nil
	}
	out := new(APIRateLimitPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APISpec) DeepCopyInto(out *APISpec) {
	*out = *in
	if in.Production != nil {
		in, out := &in.Production, &out.Production
		*out = make([]EnvConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sandbox != nil {
		in, out := &in.Sandbox, &out.Sandbox
		*out = make([]EnvConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.APIProperties != nil {
		in, out := &in.APIProperties, &out.APIProperties
		*out = make([]Property, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APISpec.
func (in *APISpec) DeepCopy() *APISpec {
	if in == nil {
		return nil
	}
	out := new(APISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIStatus) DeepCopyInto(out *APIStatus) {
	*out = *in
	in.DeploymentStatus.DeepCopyInto(&out.DeploymentStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIStatus.
func (in *APIStatus) DeepCopy() *APIStatus {
	if in == nil {
		return nil
	}
	out := new(APIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendJWTToken) DeepCopyInto(out *BackendJWTToken) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendJWTToken.
func (in *BackendJWTToken) DeepCopy() *BackendJWTToken {
	if in == nil {
		return nil
	}
	out := new(BackendJWTToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BurstControl) DeepCopyInto(out *BurstControl) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BurstControl.
func (in *BurstControl) DeepCopy() *BurstControl {
	if in == nil {
		return nil
	}
	out := new(BurstControl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CORSPolicy) DeepCopyInto(out *CORSPolicy) {
	*out = *in
	if in.AccessControlAllowHeaders != nil {
		in, out := &in.AccessControlAllowHeaders, &out.AccessControlAllowHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AccessControlAllowMethods != nil {
		in, out := &in.AccessControlAllowMethods, &out.AccessControlAllowMethods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AccessControlAllowOrigins != nil {
		in, out := &in.AccessControlAllowOrigins, &out.AccessControlAllowOrigins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AccessControlExposeHeaders != nil {
		in, out := &in.AccessControlExposeHeaders, &out.AccessControlExposeHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AccessControlMaxAge != nil {
		in, out := &in.AccessControlMaxAge, &out.AccessControlMaxAge
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CORSPolicy.
func (in *CORSPolicy) DeepCopy() *CORSPolicy {
	if in == nil {
		return nil
	}
	out := new(CORSPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRateLimitPolicy) DeepCopyInto(out *CustomRateLimitPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRateLimitPolicy.
func (in *CustomRateLimitPolicy) DeepCopy() *CustomRateLimitPolicy {
	if in == nil {
		return nil
	}
	out := new(CustomRateLimitPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStatus) DeepCopyInto(out *DeploymentStatus) {
	*out = *in
	if in.TransitionTime != nil {
		in, out := &in.TransitionTime, &out.TransitionTime
		*out = (*in).DeepCopy()
	}
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStatus.
func (in *DeploymentStatus) DeepCopy() *DeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvConfig) DeepCopyInto(out *EnvConfig) {
	*out = *in
	if in.RouteRefs != nil {
		in, out := &in.RouteRefs, &out.RouteRefs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvConfig.
func (in *EnvConfig) DeepCopy() *EnvConfig {
	if in == nil {
		return nil
	}
	out := new(EnvConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterceptorReference) DeepCopyInto(out *InterceptorReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterceptorReference.
func (in *InterceptorReference) DeepCopy() *InterceptorReference {
	if in == nil {
		return nil
	}
	out := new(InterceptorReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	if in.RequestInterceptors != nil {
		in, out := &in.RequestInterceptors, &out.RequestInterceptors
		*out = make([]InterceptorReference, len(*in))
		copy(*out, *in)
	}
	if in.ResponseInterceptors != nil {
		in, out := &in.ResponseInterceptors, &out.ResponseInterceptors
		*out = make([]InterceptorReference, len(*in))
		copy(*out, *in)
	}
	if in.BackendJWTPolicy != nil {
		in, out := &in.BackendJWTPolicy, &out.BackendJWTPolicy
		*out = new(BackendJWTToken)
		**out = **in
	}
	if in.CORSPolicy != nil {
		in, out := &in.CORSPolicy, &out.CORSPolicy
		*out = new(CORSPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.AIProvider != nil {
		in, out := &in.AIProvider, &out.AIProvider
		*out = new(AIProviderReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (in *PolicySpec) DeepCopy() *PolicySpec {
	if in == nil {
		return nil
	}
	out := new(PolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Property) DeepCopyInto(out *Property) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Property.
func (in *Property) DeepCopy() *Property {
	if in == nil {
		return nil
	}
	out := new(Property)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitAPIPolicy) DeepCopyInto(out *RateLimitAPIPolicy) {
	*out = *in
	if in.API != nil {
		in, out := &in.API, &out.API
		*out = new(APIRateLimitPolicy)
		**out = **in
	}
	if in.Subscription != nil {
		in, out := &in.Subscription, &out.Subscription
		*out = new(SubscriptionRateLimitPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = new(CustomRateLimitPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitAPIPolicy.
func (in *RateLimitAPIPolicy) DeepCopy() *RateLimitAPIPolicy {
	if in == nil {
		return nil
	}
	out := new(RateLimitAPIPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitFields) DeepCopyInto(out *RateLimitFields) {
	*out = *in
	out.PromptTokens = in.PromptTokens
	out.CompletionToken = in.CompletionToken
	out.TotalToken = in.TotalToken
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitFields.
func (in *RateLimitFields) DeepCopy() *RateLimitFields {
	if in == nil {
		return nil
	}
	out := new(RateLimitFields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicy) DeepCopyInto(out *RateLimitPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicy.
func (in *RateLimitPolicy) DeepCopy() *RateLimitPolicy {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RateLimitPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicyList) DeepCopyInto(out *RateLimitPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RateLimitPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicyList.
func (in *RateLimitPolicyList) DeepCopy() *RateLimitPolicyList {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RateLimitPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicySpec) DeepCopyInto(out *RateLimitPolicySpec) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(RateLimitAPIPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Override != nil {
		in, out := &in.Override, &out.Override
		*out = new(RateLimitAPIPolicy)
		(*in).DeepCopyInto(*out)
	}
	in.TargetRef.DeepCopyInto(&out.TargetRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicySpec.
func (in *RateLimitPolicySpec) DeepCopy() *RateLimitPolicySpec {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicyStatus) DeepCopyInto(out *RateLimitPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicyStatus.
func (in *RateLimitPolicyStatus) DeepCopy() *RateLimitPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestCount) DeepCopyInto(out *RequestCount) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestCount.
func (in *RequestCount) DeepCopy() *RequestCount {
	if in == nil {
		return nil
	}
	out := new(RequestCount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolveBurstControl) DeepCopyInto(out *ResolveBurstControl) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolveBurstControl.
func (in *ResolveBurstControl) DeepCopy() *ResolveBurstControl {
	if in == nil {
		return nil
	}
	out := new(ResolveBurstControl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolveRequestCount) DeepCopyInto(out *ResolveRequestCount) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolveRequestCount.
func (in *ResolveRequestCount) DeepCopy() *ResolveRequestCount {
	if in == nil {
		return nil
	}
	out := new(ResolveRequestCount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolveSubscriptionRatelimitPolicy) DeepCopyInto(out *ResolveSubscriptionRatelimitPolicy) {
	*out = *in
	out.RequestCount = in.RequestCount
	out.BurstControl = in.BurstControl
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolveSubscriptionRatelimitPolicy.
func (in *ResolveSubscriptionRatelimitPolicy) DeepCopy() *ResolveSubscriptionRatelimitPolicy {
	if in == nil {
		return nil
	}
	out := new(ResolveSubscriptionRatelimitPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionRateLimitPolicy) DeepCopyInto(out *SubscriptionRateLimitPolicy) {
	*out = *in
	if in.RequestCount != nil {
		in, out := &in.RequestCount, &out.RequestCount
		*out = new(RequestCount)
		**out = **in
	}
	if in.BurstControl != nil {
		in, out := &in.BurstControl, &out.BurstControl
		*out = new(BurstControl)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionRateLimitPolicy.
func (in *SubscriptionRateLimitPolicy) DeepCopy() *SubscriptionRateLimitPolicy {
	if in == nil {
		return nil
	}
	out := new(SubscriptionRateLimitPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenCount) DeepCopyInto(out *TokenCount) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenCount.
func (in *TokenCount) DeepCopy() *TokenCount {
	if in == nil {
		return nil
	}
	out := new(TokenCount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueDetails) DeepCopyInto(out *ValueDetails) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueDetails.
func (in *ValueDetails) DeepCopy() *ValueDetails {
	if in == nil {
		return nil
	}
	out := new(ValueDetails)
	in.DeepCopyInto(out)
	return out
}
