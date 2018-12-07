// +build !ignore_autogenerated

/*
Copyright 2018 The Service Fabrik Authors.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Context) DeepCopyInto(out *Context) {
	*out = *in
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceFabrik != nil {
		in, out := &in.ServiceFabrik, &out.ServiceFabrik
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Context.
func (in *Context) DeepCopy() *Context {
	if in == nil {
		return nil
	}
	out := new(Context)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardClient) DeepCopyInto(out *DashboardClient) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardClient.
func (in *DashboardClient) DeepCopy() *DashboardClient {
	if in == nil {
		return nil
	}
	out := new(DashboardClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Schema) DeepCopyInto(out *Schema) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Schema.
func (in *Schema) DeepCopy() *Schema {
	if in == nil {
		return nil
	}
	out := new(Schema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingSchema) DeepCopyInto(out *ServiceBindingSchema) {
	*out = *in
	in.Create.DeepCopyInto(&out.Create)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingSchema.
func (in *ServiceBindingSchema) DeepCopy() *ServiceBindingSchema {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceInstanceSchema) DeepCopyInto(out *ServiceInstanceSchema) {
	*out = *in
	in.Create.DeepCopyInto(&out.Create)
	in.Update.DeepCopyInto(&out.Update)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceInstanceSchema.
func (in *ServiceInstanceSchema) DeepCopy() *ServiceInstanceSchema {
	if in == nil {
		return nil
	}
	out := new(ServiceInstanceSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMetadata) DeepCopyInto(out *ServiceMetadata) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMetadata.
func (in *ServiceMetadata) DeepCopy() *ServiceMetadata {
	if in == nil {
		return nil
	}
	out := new(ServiceMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSchemas) DeepCopyInto(out *ServiceSchemas) {
	*out = *in
	in.Instance.DeepCopyInto(&out.Instance)
	in.Binding.DeepCopyInto(&out.Binding)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSchemas.
func (in *ServiceSchemas) DeepCopy() *ServiceSchemas {
	if in == nil {
		return nil
	}
	out := new(ServiceSchemas)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SfPlan) DeepCopyInto(out *SfPlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SfPlan.
func (in *SfPlan) DeepCopy() *SfPlan {
	if in == nil {
		return nil
	}
	out := new(SfPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SfPlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SfPlanList) DeepCopyInto(out *SfPlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SfPlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SfPlanList.
func (in *SfPlanList) DeepCopy() *SfPlanList {
	if in == nil {
		return nil
	}
	out := new(SfPlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SfPlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SfPlanSpec) DeepCopyInto(out *SfPlanSpec) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Schemas != nil {
		in, out := &in.Schemas, &out.Schemas
		*out = new(ServiceSchemas)
		(*in).DeepCopyInto(*out)
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make([]TemplateSpec, len(*in))
		copy(*out, *in)
	}
	if in.RawContext != nil {
		in, out := &in.RawContext, &out.RawContext
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Manager != nil {
		in, out := &in.Manager, &out.Manager
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SfPlanSpec.
func (in *SfPlanSpec) DeepCopy() *SfPlanSpec {
	if in == nil {
		return nil
	}
	out := new(SfPlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SfPlanStatus) DeepCopyInto(out *SfPlanStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SfPlanStatus.
func (in *SfPlanStatus) DeepCopy() *SfPlanStatus {
	if in == nil {
		return nil
	}
	out := new(SfPlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SfService) DeepCopyInto(out *SfService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SfService.
func (in *SfService) DeepCopy() *SfService {
	if in == nil {
		return nil
	}
	out := new(SfService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SfService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SfServiceBinding) DeepCopyInto(out *SfServiceBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SfServiceBinding.
func (in *SfServiceBinding) DeepCopy() *SfServiceBinding {
	if in == nil {
		return nil
	}
	out := new(SfServiceBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SfServiceBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SfServiceBindingList) DeepCopyInto(out *SfServiceBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SfServiceBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SfServiceBindingList.
func (in *SfServiceBindingList) DeepCopy() *SfServiceBindingList {
	if in == nil {
		return nil
	}
	out := new(SfServiceBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SfServiceBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SfServiceBindingSpec) DeepCopyInto(out *SfServiceBindingSpec) {
	*out = *in
	if in.BindResource != nil {
		in, out := &in.BindResource, &out.BindResource
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.RawContext != nil {
		in, out := &in.RawContext, &out.RawContext
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.RawParameters != nil {
		in, out := &in.RawParameters, &out.RawParameters
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SfServiceBindingSpec.
func (in *SfServiceBindingSpec) DeepCopy() *SfServiceBindingSpec {
	if in == nil {
		return nil
	}
	out := new(SfServiceBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SfServiceBindingStatus) DeepCopyInto(out *SfServiceBindingStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SfServiceBindingStatus.
func (in *SfServiceBindingStatus) DeepCopy() *SfServiceBindingStatus {
	if in == nil {
		return nil
	}
	out := new(SfServiceBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SfServiceInstance) DeepCopyInto(out *SfServiceInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SfServiceInstance.
func (in *SfServiceInstance) DeepCopy() *SfServiceInstance {
	if in == nil {
		return nil
	}
	out := new(SfServiceInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SfServiceInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SfServiceInstanceList) DeepCopyInto(out *SfServiceInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SfServiceInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SfServiceInstanceList.
func (in *SfServiceInstanceList) DeepCopy() *SfServiceInstanceList {
	if in == nil {
		return nil
	}
	out := new(SfServiceInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SfServiceInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SfServiceInstanceSpec) DeepCopyInto(out *SfServiceInstanceSpec) {
	*out = *in
	if in.RawContext != nil {
		in, out := &in.RawContext, &out.RawContext
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.RawParameters != nil {
		in, out := &in.RawParameters, &out.RawParameters
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SfServiceInstanceSpec.
func (in *SfServiceInstanceSpec) DeepCopy() *SfServiceInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(SfServiceInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SfServiceInstanceStatus) DeepCopyInto(out *SfServiceInstanceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SfServiceInstanceStatus.
func (in *SfServiceInstanceStatus) DeepCopy() *SfServiceInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(SfServiceInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SfServiceList) DeepCopyInto(out *SfServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SfService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SfServiceList.
func (in *SfServiceList) DeepCopy() *SfServiceList {
	if in == nil {
		return nil
	}
	out := new(SfServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SfServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SfServiceSpec) DeepCopyInto(out *SfServiceSpec) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Requires != nil {
		in, out := &in.Requires, &out.Requires
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	out.DashboardClient = in.DashboardClient
	if in.RawContext != nil {
		in, out := &in.RawContext, &out.RawContext
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SfServiceSpec.
func (in *SfServiceSpec) DeepCopy() *SfServiceSpec {
	if in == nil {
		return nil
	}
	out := new(SfServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SfServiceStatus) DeepCopyInto(out *SfServiceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SfServiceStatus.
func (in *SfServiceStatus) DeepCopy() *SfServiceStatus {
	if in == nil {
		return nil
	}
	out := new(SfServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpec) DeepCopyInto(out *TemplateSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpec.
func (in *TemplateSpec) DeepCopy() *TemplateSpec {
	if in == nil {
		return nil
	}
	out := new(TemplateSpec)
	in.DeepCopyInto(out)
	return out
}
