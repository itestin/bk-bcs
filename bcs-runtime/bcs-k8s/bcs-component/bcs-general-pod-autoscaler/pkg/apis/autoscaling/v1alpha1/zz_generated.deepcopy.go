//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1beta1 "k8s.io/api/admissionregistration/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingDrivenMode) DeepCopyInto(out *AutoScalingDrivenMode) {
	*out = *in
	if in.MetricMode != nil {
		in, out := &in.MetricMode, &out.MetricMode
		*out = new(MetricMode)
		(*in).DeepCopyInto(*out)
	}
	if in.WebhookMode != nil {
		in, out := &in.WebhookMode, &out.WebhookMode
		*out = new(WebhookMode)
		(*in).DeepCopyInto(*out)
	}
	if in.TimeMode != nil {
		in, out := &in.TimeMode, &out.TimeMode
		*out = new(TimeMode)
		(*in).DeepCopyInto(*out)
	}
	if in.EventMode != nil {
		in, out := &in.EventMode, &out.EventMode
		*out = new(EventMode)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingDrivenMode.
func (in *AutoScalingDrivenMode) DeepCopy() *AutoScalingDrivenMode {
	if in == nil {
		return nil
	}
	out := new(AutoScalingDrivenMode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerResourceMetricSource) DeepCopyInto(out *ContainerResourceMetricSource) {
	*out = *in
	in.Target.DeepCopyInto(&out.Target)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerResourceMetricSource.
func (in *ContainerResourceMetricSource) DeepCopy() *ContainerResourceMetricSource {
	if in == nil {
		return nil
	}
	out := new(ContainerResourceMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerResourceMetricStatus) DeepCopyInto(out *ContainerResourceMetricStatus) {
	*out = *in
	in.Current.DeepCopyInto(&out.Current)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerResourceMetricStatus.
func (in *ContainerResourceMetricStatus) DeepCopy() *ContainerResourceMetricStatus {
	if in == nil {
		return nil
	}
	out := new(ContainerResourceMetricStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrossVersionObjectReference) DeepCopyInto(out *CrossVersionObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrossVersionObjectReference.
func (in *CrossVersionObjectReference) DeepCopy() *CrossVersionObjectReference {
	if in == nil {
		return nil
	}
	out := new(CrossVersionObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventMode) DeepCopyInto(out *EventMode) {
	*out = *in
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]ScaleTriggers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventMode.
func (in *EventMode) DeepCopy() *EventMode {
	if in == nil {
		return nil
	}
	out := new(EventMode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalMetricSource) DeepCopyInto(out *ExternalMetricSource) {
	*out = *in
	in.Metric.DeepCopyInto(&out.Metric)
	in.Target.DeepCopyInto(&out.Target)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalMetricSource.
func (in *ExternalMetricSource) DeepCopy() *ExternalMetricSource {
	if in == nil {
		return nil
	}
	out := new(ExternalMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalMetricStatus) DeepCopyInto(out *ExternalMetricStatus) {
	*out = *in
	in.Metric.DeepCopyInto(&out.Metric)
	in.Current.DeepCopyInto(&out.Current)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalMetricStatus.
func (in *ExternalMetricStatus) DeepCopy() *ExternalMetricStatus {
	if in == nil {
		return nil
	}
	out := new(ExternalMetricStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GPAScalingPolicy) DeepCopyInto(out *GPAScalingPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GPAScalingPolicy.
func (in *GPAScalingPolicy) DeepCopy() *GPAScalingPolicy {
	if in == nil {
		return nil
	}
	out := new(GPAScalingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GPAScalingRules) DeepCopyInto(out *GPAScalingRules) {
	*out = *in
	if in.StabilizationWindowSeconds != nil {
		in, out := &in.StabilizationWindowSeconds, &out.StabilizationWindowSeconds
		*out = new(int32)
		**out = **in
	}
	if in.SelectPolicy != nil {
		in, out := &in.SelectPolicy, &out.SelectPolicy
		*out = new(ScalingPolicySelect)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]GPAScalingPolicy, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GPAScalingRules.
func (in *GPAScalingRules) DeepCopy() *GPAScalingRules {
	if in == nil {
		return nil
	}
	out := new(GPAScalingRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeneralPodAutoscaler) DeepCopyInto(out *GeneralPodAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeneralPodAutoscaler.
func (in *GeneralPodAutoscaler) DeepCopy() *GeneralPodAutoscaler {
	if in == nil {
		return nil
	}
	out := new(GeneralPodAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GeneralPodAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeneralPodAutoscalerBehavior) DeepCopyInto(out *GeneralPodAutoscalerBehavior) {
	*out = *in
	if in.ScaleUp != nil {
		in, out := &in.ScaleUp, &out.ScaleUp
		*out = new(GPAScalingRules)
		(*in).DeepCopyInto(*out)
	}
	if in.ScaleDown != nil {
		in, out := &in.ScaleDown, &out.ScaleDown
		*out = new(GPAScalingRules)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeneralPodAutoscalerBehavior.
func (in *GeneralPodAutoscalerBehavior) DeepCopy() *GeneralPodAutoscalerBehavior {
	if in == nil {
		return nil
	}
	out := new(GeneralPodAutoscalerBehavior)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeneralPodAutoscalerCondition) DeepCopyInto(out *GeneralPodAutoscalerCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeneralPodAutoscalerCondition.
func (in *GeneralPodAutoscalerCondition) DeepCopy() *GeneralPodAutoscalerCondition {
	if in == nil {
		return nil
	}
	out := new(GeneralPodAutoscalerCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeneralPodAutoscalerList) DeepCopyInto(out *GeneralPodAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GeneralPodAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeneralPodAutoscalerList.
func (in *GeneralPodAutoscalerList) DeepCopy() *GeneralPodAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(GeneralPodAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GeneralPodAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeneralPodAutoscalerSpec) DeepCopyInto(out *GeneralPodAutoscalerSpec) {
	*out = *in
	in.AutoScalingDrivenMode.DeepCopyInto(&out.AutoScalingDrivenMode)
	out.ScaleTargetRef = in.ScaleTargetRef
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.Behavior != nil {
		in, out := &in.Behavior, &out.Behavior
		*out = new(GeneralPodAutoscalerBehavior)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeneralPodAutoscalerSpec.
func (in *GeneralPodAutoscalerSpec) DeepCopy() *GeneralPodAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(GeneralPodAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeneralPodAutoscalerStatus) DeepCopyInto(out *GeneralPodAutoscalerStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.LastScaleTime != nil {
		in, out := &in.LastScaleTime, &out.LastScaleTime
		*out = (*in).DeepCopy()
	}
	if in.CurrentMetrics != nil {
		in, out := &in.CurrentMetrics, &out.CurrentMetrics
		*out = make([]MetricStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]GeneralPodAutoscalerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastCronScheduleTime != nil {
		in, out := &in.LastCronScheduleTime, &out.LastCronScheduleTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeneralPodAutoscalerStatus.
func (in *GeneralPodAutoscalerStatus) DeepCopy() *GeneralPodAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(GeneralPodAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricIdentifier) DeepCopyInto(out *MetricIdentifier) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricIdentifier.
func (in *MetricIdentifier) DeepCopy() *MetricIdentifier {
	if in == nil {
		return nil
	}
	out := new(MetricIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricMode) DeepCopyInto(out *MetricMode) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricMode.
func (in *MetricMode) DeepCopy() *MetricMode {
	if in == nil {
		return nil
	}
	out := new(MetricMode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricSpec) DeepCopyInto(out *MetricSpec) {
	*out = *in
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(ObjectMetricSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = new(PodsMetricSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(ResourceMetricSource)
		(*in).DeepCopyInto(*out)
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ExternalMetricSource)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerResource != nil {
		in, out := &in.ContainerResource, &out.ContainerResource
		*out = new(ContainerResourceMetricSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricSpec.
func (in *MetricSpec) DeepCopy() *MetricSpec {
	if in == nil {
		return nil
	}
	out := new(MetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricStatus) DeepCopyInto(out *MetricStatus) {
	*out = *in
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(ObjectMetricStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = new(PodsMetricStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(ResourceMetricStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ExternalMetricStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerResource != nil {
		in, out := &in.ContainerResource, &out.ContainerResource
		*out = new(ContainerResourceMetricStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricStatus.
func (in *MetricStatus) DeepCopy() *MetricStatus {
	if in == nil {
		return nil
	}
	out := new(MetricStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricTarget) DeepCopyInto(out *MetricTarget) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.AverageValue != nil {
		in, out := &in.AverageValue, &out.AverageValue
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.AverageUtilization != nil {
		in, out := &in.AverageUtilization, &out.AverageUtilization
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricTarget.
func (in *MetricTarget) DeepCopy() *MetricTarget {
	if in == nil {
		return nil
	}
	out := new(MetricTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricValueStatus) DeepCopyInto(out *MetricValueStatus) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.AverageValue != nil {
		in, out := &in.AverageValue, &out.AverageValue
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.AverageUtilization != nil {
		in, out := &in.AverageUtilization, &out.AverageUtilization
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricValueStatus.
func (in *MetricValueStatus) DeepCopy() *MetricValueStatus {
	if in == nil {
		return nil
	}
	out := new(MetricValueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMetricSource) DeepCopyInto(out *ObjectMetricSource) {
	*out = *in
	out.DescribedObject = in.DescribedObject
	in.Target.DeepCopyInto(&out.Target)
	in.Metric.DeepCopyInto(&out.Metric)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMetricSource.
func (in *ObjectMetricSource) DeepCopy() *ObjectMetricSource {
	if in == nil {
		return nil
	}
	out := new(ObjectMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMetricStatus) DeepCopyInto(out *ObjectMetricStatus) {
	*out = *in
	in.Metric.DeepCopyInto(&out.Metric)
	in.Current.DeepCopyInto(&out.Current)
	out.DescribedObject = in.DescribedObject
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMetricStatus.
func (in *ObjectMetricStatus) DeepCopy() *ObjectMetricStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectMetricStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodsMetricSource) DeepCopyInto(out *PodsMetricSource) {
	*out = *in
	in.Metric.DeepCopyInto(&out.Metric)
	in.Target.DeepCopyInto(&out.Target)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodsMetricSource.
func (in *PodsMetricSource) DeepCopy() *PodsMetricSource {
	if in == nil {
		return nil
	}
	out := new(PodsMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodsMetricStatus) DeepCopyInto(out *PodsMetricStatus) {
	*out = *in
	in.Metric.DeepCopyInto(&out.Metric)
	in.Current.DeepCopyInto(&out.Current)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodsMetricStatus.
func (in *PodsMetricStatus) DeepCopy() *PodsMetricStatus {
	if in == nil {
		return nil
	}
	out := new(PodsMetricStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMetricSource) DeepCopyInto(out *ResourceMetricSource) {
	*out = *in
	in.Target.DeepCopyInto(&out.Target)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMetricSource.
func (in *ResourceMetricSource) DeepCopy() *ResourceMetricSource {
	if in == nil {
		return nil
	}
	out := new(ResourceMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMetricStatus) DeepCopyInto(out *ResourceMetricStatus) {
	*out = *in
	in.Current.DeepCopyInto(&out.Current)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMetricStatus.
func (in *ResourceMetricStatus) DeepCopy() *ResourceMetricStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceMetricStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleTriggers) DeepCopyInto(out *ScaleTriggers) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleTriggers.
func (in *ScaleTriggers) DeepCopy() *ScaleTriggers {
	if in == nil {
		return nil
	}
	out := new(ScaleTriggers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeMode) DeepCopyInto(out *TimeMode) {
	*out = *in
	if in.TimeRanges != nil {
		in, out := &in.TimeRanges, &out.TimeRanges
		*out = make([]TimeRange, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeMode.
func (in *TimeMode) DeepCopy() *TimeMode {
	if in == nil {
		return nil
	}
	out := new(TimeMode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeRange) DeepCopyInto(out *TimeRange) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeRange.
func (in *TimeRange) DeepCopy() *TimeRange {
	if in == nil {
		return nil
	}
	out := new(TimeRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookMode) DeepCopyInto(out *WebhookMode) {
	*out = *in
	if in.WebhookClientConfig != nil {
		in, out := &in.WebhookClientConfig, &out.WebhookClientConfig
		*out = new(v1beta1.WebhookClientConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookMode.
func (in *WebhookMode) DeepCopy() *WebhookMode {
	if in == nil {
		return nil
	}
	out := new(WebhookMode)
	in.DeepCopyInto(out)
	return out
}
