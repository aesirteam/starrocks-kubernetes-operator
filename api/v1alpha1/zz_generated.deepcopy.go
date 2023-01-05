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

package v1alpha1

import (
	"k8s.io/api/autoscaling/v2beta2"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingPolicy) DeepCopyInto(out *AutoScalingPolicy) {
	*out = *in
	if in.HPAPolicy != nil {
		in, out := &in.HPAPolicy, &out.HPAPolicy
		*out = new(HPAPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingPolicy.
func (in *AutoScalingPolicy) DeepCopy() *AutoScalingPolicy {
	if in == nil {
		return nil
	}
	out := new(AutoScalingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnInfo) DeepCopyInto(out *CnInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnInfo.
func (in *CnInfo) DeepCopy() *CnInfo {
	if in == nil {
		return nil
	}
	out := new(CnInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeNodeGroup) DeepCopyInto(out *ComputeNodeGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeNodeGroup.
func (in *ComputeNodeGroup) DeepCopy() *ComputeNodeGroup {
	if in == nil {
		return nil
	}
	out := new(ComputeNodeGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputeNodeGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeNodeGroupList) DeepCopyInto(out *ComputeNodeGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComputeNodeGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeNodeGroupList.
func (in *ComputeNodeGroupList) DeepCopy() *ComputeNodeGroupList {
	if in == nil {
		return nil
	}
	out := new(ComputeNodeGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputeNodeGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeNodeGroupSpec) DeepCopyInto(out *ComputeNodeGroupSpec) {
	*out = *in
	in.FeInfo.DeepCopyInto(&out.FeInfo)
	out.CnInfo = in.CnInfo
	out.Images = in.Images
	if in.AutoScalingPolicy != nil {
		in, out := &in.AutoScalingPolicy, &out.AutoScalingPolicy
		*out = new(AutoScalingPolicy)
		(*in).DeepCopyInto(*out)
	}
	in.PodPolicy.DeepCopyInto(&out.PodPolicy)
	out.CronJobPolicy = in.CronJobPolicy
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeNodeGroupSpec.
func (in *ComputeNodeGroupSpec) DeepCopy() *ComputeNodeGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ComputeNodeGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeNodeGroupStatus) DeepCopyInto(out *ComputeNodeGroupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(map[CnComponent]ResourceCondition, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	out.Servers = in.Servers
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeNodeGroupStatus.
func (in *ComputeNodeGroupStatus) DeepCopy() *ComputeNodeGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ComputeNodeGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapInfo) DeepCopyInto(out *ConfigMapInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapInfo.
func (in *ConfigMapInfo) DeepCopy() *ConfigMapInfo {
	if in == nil {
		return nil
	}
	out := new(ConfigMapInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJobPolicy) DeepCopyInto(out *CronJobPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJobPolicy.
func (in *CronJobPolicy) DeepCopy() *CronJobPolicy {
	if in == nil {
		return nil
	}
	out := new(CronJobPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeInfo) DeepCopyInto(out *FeInfo) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeInfo.
func (in *FeInfo) DeepCopy() *FeInfo {
	if in == nil {
		return nil
	}
	out := new(FeInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPAPolicy) DeepCopyInto(out *HPAPolicy) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]v2beta2.MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Behavior != nil {
		in, out := &in.Behavior, &out.Behavior
		*out = new(v2beta2.HorizontalPodAutoscalerBehavior)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPAPolicy.
func (in *HPAPolicy) DeepCopy() *HPAPolicy {
	if in == nil {
		return nil
	}
	out := new(HPAPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Images) DeepCopyInto(out *Images) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Images.
func (in *Images) DeepCopy() *Images {
	if in == nil {
		return nil
	}
	out := new(Images)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodPolicy) DeepCopyInto(out *PodPolicy) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodPolicy.
func (in *PodPolicy) DeepCopy() *PodPolicy {
	if in == nil {
		return nil
	}
	out := new(PodPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceCondition) DeepCopyInto(out *ResourceCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceCondition.
func (in *ResourceCondition) DeepCopy() *ResourceCondition {
	if in == nil {
		return nil
	}
	out := new(ResourceCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServersStatus) DeepCopyInto(out *ServersStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServersStatus.
func (in *ServersStatus) DeepCopy() *ServersStatus {
	if in == nil {
		return nil
	}
	out := new(ServersStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksBeSpec) DeepCopyInto(out *StarRocksBeSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(StarRocksService)
		(*in).DeepCopyInto(*out)
	}
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	out.ConfigMapInfo = in.ConfigMapInfo
	if in.Probe != nil {
		in, out := &in.Probe, &out.Probe
		*out = new(StarRocksProbe)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageVolumes != nil {
		in, out := &in.StorageVolumes, &out.StorageVolumes
		*out = make([]StorageVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReplicaInstances != nil {
		in, out := &in.ReplicaInstances, &out.ReplicaInstances
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksBeSpec.
func (in *StarRocksBeSpec) DeepCopy() *StarRocksBeSpec {
	if in == nil {
		return nil
	}
	out := new(StarRocksBeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksBeStatus) DeepCopyInto(out *StarRocksBeStatus) {
	*out = *in
	if in.FailedInstances != nil {
		in, out := &in.FailedInstances, &out.FailedInstances
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CreatingInstances != nil {
		in, out := &in.CreatingInstances, &out.CreatingInstances
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RunningInstances != nil {
		in, out := &in.RunningInstances, &out.RunningInstances
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceNames != nil {
		in, out := &in.ResourceNames, &out.ResourceNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksBeStatus.
func (in *StarRocksBeStatus) DeepCopy() *StarRocksBeStatus {
	if in == nil {
		return nil
	}
	out := new(StarRocksBeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksCluster) DeepCopyInto(out *StarRocksCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksCluster.
func (in *StarRocksCluster) DeepCopy() *StarRocksCluster {
	if in == nil {
		return nil
	}
	out := new(StarRocksCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StarRocksCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksClusterList) DeepCopyInto(out *StarRocksClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StarRocksCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksClusterList.
func (in *StarRocksClusterList) DeepCopy() *StarRocksClusterList {
	if in == nil {
		return nil
	}
	out := new(StarRocksClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StarRocksClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksClusterSpec) DeepCopyInto(out *StarRocksClusterSpec) {
	*out = *in
	if in.StarRocksFeSpec != nil {
		in, out := &in.StarRocksFeSpec, &out.StarRocksFeSpec
		*out = new(StarRocksFeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.StarRocksBeSpec != nil {
		in, out := &in.StarRocksBeSpec, &out.StarRocksBeSpec
		*out = new(StarRocksBeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.StarRocksCnSpec != nil {
		in, out := &in.StarRocksCnSpec, &out.StarRocksCnSpec
		*out = new(StarRocksCnSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksClusterSpec.
func (in *StarRocksClusterSpec) DeepCopy() *StarRocksClusterSpec {
	if in == nil {
		return nil
	}
	out := new(StarRocksClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksClusterStatus) DeepCopyInto(out *StarRocksClusterStatus) {
	*out = *in
	if in.StarRocksFeStatus != nil {
		in, out := &in.StarRocksFeStatus, &out.StarRocksFeStatus
		*out = new(StarRocksFeStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.StarRocksBeStatus != nil {
		in, out := &in.StarRocksBeStatus, &out.StarRocksBeStatus
		*out = new(StarRocksBeStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.StarRocksCnStatus != nil {
		in, out := &in.StarRocksCnStatus, &out.StarRocksCnStatus
		*out = new(StarRocksCnStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksClusterStatus.
func (in *StarRocksClusterStatus) DeepCopy() *StarRocksClusterStatus {
	if in == nil {
		return nil
	}
	out := new(StarRocksClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksCnSpec) DeepCopyInto(out *StarRocksCnSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(StarRocksService)
		(*in).DeepCopyInto(*out)
	}
	out.ConfigMapInfo = in.ConfigMapInfo
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	if in.Probe != nil {
		in, out := &in.Probe, &out.Probe
		*out = new(StarRocksProbe)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoScalingPolicy != nil {
		in, out := &in.AutoScalingPolicy, &out.AutoScalingPolicy
		*out = new(AutoScalingPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksCnSpec.
func (in *StarRocksCnSpec) DeepCopy() *StarRocksCnSpec {
	if in == nil {
		return nil
	}
	out := new(StarRocksCnSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksCnStatus) DeepCopyInto(out *StarRocksCnStatus) {
	*out = *in
	if in.FailedInstances != nil {
		in, out := &in.FailedInstances, &out.FailedInstances
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CreatingInstances != nil {
		in, out := &in.CreatingInstances, &out.CreatingInstances
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RunningInstances != nil {
		in, out := &in.RunningInstances, &out.RunningInstances
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceNames != nil {
		in, out := &in.ResourceNames, &out.ResourceNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksCnStatus.
func (in *StarRocksCnStatus) DeepCopy() *StarRocksCnStatus {
	if in == nil {
		return nil
	}
	out := new(StarRocksCnStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksFeSpec) DeepCopyInto(out *StarRocksFeSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(StarRocksService)
		(*in).DeepCopyInto(*out)
	}
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	out.ConfigMapInfo = in.ConfigMapInfo
	if in.Probe != nil {
		in, out := &in.Probe, &out.Probe
		*out = new(StarRocksProbe)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageVolumes != nil {
		in, out := &in.StorageVolumes, &out.StorageVolumes
		*out = make([]StorageVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksFeSpec.
func (in *StarRocksFeSpec) DeepCopy() *StarRocksFeSpec {
	if in == nil {
		return nil
	}
	out := new(StarRocksFeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksFeStatus) DeepCopyInto(out *StarRocksFeStatus) {
	*out = *in
	if in.FailedInstances != nil {
		in, out := &in.FailedInstances, &out.FailedInstances
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CreatingInstances != nil {
		in, out := &in.CreatingInstances, &out.CreatingInstances
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RunningInstances != nil {
		in, out := &in.RunningInstances, &out.RunningInstances
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceNames != nil {
		in, out := &in.ResourceNames, &out.ResourceNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksFeStatus.
func (in *StarRocksFeStatus) DeepCopy() *StarRocksFeStatus {
	if in == nil {
		return nil
	}
	out := new(StarRocksFeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksProbe) DeepCopyInto(out *StarRocksProbe) {
	*out = *in
	if in.InitialDelaySeconds != nil {
		in, out := &in.InitialDelaySeconds, &out.InitialDelaySeconds
		*out = new(int32)
		**out = **in
	}
	if in.PeriodSeconds != nil {
		in, out := &in.PeriodSeconds, &out.PeriodSeconds
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksProbe.
func (in *StarRocksProbe) DeepCopy() *StarRocksProbe {
	if in == nil {
		return nil
	}
	out := new(StarRocksProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksService) DeepCopyInto(out *StarRocksService) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]StarRocksServicePort, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksService.
func (in *StarRocksService) DeepCopy() *StarRocksService {
	if in == nil {
		return nil
	}
	out := new(StarRocksService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StarRocksServicePort) DeepCopyInto(out *StarRocksServicePort) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StarRocksServicePort.
func (in *StarRocksServicePort) DeepCopy() *StarRocksServicePort {
	if in == nil {
		return nil
	}
	out := new(StarRocksServicePort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageVolume) DeepCopyInto(out *StorageVolume) {
	*out = *in
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageVolume.
func (in *StorageVolume) DeepCopy() *StorageVolume {
	if in == nil {
		return nil
	}
	out := new(StorageVolume)
	in.DeepCopyInto(out)
	return out
}
