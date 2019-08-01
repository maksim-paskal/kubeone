// +build !ignore_autogenerated

/*
Copyright The KubeOne Authors.

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
	json "encoding/json"

	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIEndpoint) DeepCopyInto(out *APIEndpoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIEndpoint.
func (in *APIEndpoint) DeepCopy() *APIEndpoint {
	if in == nil {
		return nil
	}
	out := new(APIEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CNI) DeepCopyInto(out *CNI) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNI.
func (in *CNI) DeepCopy() *CNI {
	if in == nil {
		return nil
	}
	out := new(CNI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudProviderSpec) DeepCopyInto(out *CloudProviderSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudProviderSpec.
func (in *CloudProviderSpec) DeepCopy() *CloudProviderSpec {
	if in == nil {
		return nil
	}
	out := new(CloudProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetworkConfig) DeepCopyInto(out *ClusterNetworkConfig) {
	*out = *in
	if in.CNI != nil {
		in, out := &in.CNI, &out.CNI
		*out = new(CNI)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetworkConfig.
func (in *ClusterNetworkConfig) DeepCopy() *ClusterNetworkConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSConfig) DeepCopyInto(out *DNSConfig) {
	*out = *in
	if in.Servers != nil {
		in, out := &in.Servers, &out.Servers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSConfig.
func (in *DNSConfig) DeepCopy() *DNSConfig {
	if in == nil {
		return nil
	}
	out := new(DNSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicAuditLog) DeepCopyInto(out *DynamicAuditLog) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicAuditLog.
func (in *DynamicAuditLog) DeepCopy() *DynamicAuditLog {
	if in == nil {
		return nil
	}
	out := new(DynamicAuditLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Features) DeepCopyInto(out *Features) {
	*out = *in
	if in.PodSecurityPolicy != nil {
		in, out := &in.PodSecurityPolicy, &out.PodSecurityPolicy
		*out = new(PodSecurityPolicy)
		**out = **in
	}
	if in.DynamicAuditLog != nil {
		in, out := &in.DynamicAuditLog, &out.DynamicAuditLog
		*out = new(DynamicAuditLog)
		**out = **in
	}
	if in.MetricsServer != nil {
		in, out := &in.MetricsServer, &out.MetricsServer
		*out = new(MetricsServer)
		**out = **in
	}
	if in.OpenIDConnect != nil {
		in, out := &in.OpenIDConnect, &out.OpenIDConnect
		*out = new(OpenIDConnect)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Features.
func (in *Features) DeepCopy() *Features {
	if in == nil {
		return nil
	}
	out := new(Features)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostConfig) DeepCopyInto(out *HostConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostConfig.
func (in *HostConfig) DeepCopy() *HostConfig {
	if in == nil {
		return nil
	}
	out := new(HostConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeOneCluster) DeepCopyInto(out *KubeOneCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]HostConfig, len(*in))
		copy(*out, *in)
	}
	out.APIEndpoint = in.APIEndpoint
	out.CloudProvider = in.CloudProvider
	out.Versions = in.Versions
	in.ClusterNetwork.DeepCopyInto(&out.ClusterNetwork)
	out.Proxy = in.Proxy
	if in.Workers != nil {
		in, out := &in.Workers, &out.Workers
		*out = make([]WorkerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MachineController != nil {
		in, out := &in.MachineController, &out.MachineController
		*out = new(MachineControllerConfig)
		**out = **in
	}
	in.Features.DeepCopyInto(&out.Features)
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeOneCluster.
func (in *KubeOneCluster) DeepCopy() *KubeOneCluster {
	if in == nil {
		return nil
	}
	out := new(KubeOneCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeOneCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeOneSecrets) DeepCopyInto(out *KubeOneSecrets) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeOneSecrets.
func (in *KubeOneSecrets) DeepCopy() *KubeOneSecrets {
	if in == nil {
		return nil
	}
	out := new(KubeOneSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeOneSecrets) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineControllerConfig) DeepCopyInto(out *MachineControllerConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineControllerConfig.
func (in *MachineControllerConfig) DeepCopy() *MachineControllerConfig {
	if in == nil {
		return nil
	}
	out := new(MachineControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsServer) DeepCopyInto(out *MetricsServer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsServer.
func (in *MetricsServer) DeepCopy() *MetricsServer {
	if in == nil {
		return nil
	}
	out := new(MetricsServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfig) DeepCopyInto(out *NetworkConfig) {
	*out = *in
	in.DNS.DeepCopyInto(&out.DNS)
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenIDConnect) DeepCopyInto(out *OpenIDConnect) {
	*out = *in
	out.Config = in.Config
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenIDConnect.
func (in *OpenIDConnect) DeepCopy() *OpenIDConnect {
	if in == nil {
		return nil
	}
	out := new(OpenIDConnect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenIDConnectConfig) DeepCopyInto(out *OpenIDConnectConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenIDConnectConfig.
func (in *OpenIDConnectConfig) DeepCopy() *OpenIDConnectConfig {
	if in == nil {
		return nil
	}
	out := new(OpenIDConnectConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicy) DeepCopyInto(out *PodSecurityPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicy.
func (in *PodSecurityPolicy) DeepCopy() *PodSecurityPolicy {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSpec) DeepCopyInto(out *ProviderSpec) {
	*out = *in
	if in.CloudProviderSpec != nil {
		in, out := &in.CloudProviderSpec, &out.CloudProviderSpec
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SSHPublicKeys != nil {
		in, out := &in.SSHPublicKeys, &out.SSHPublicKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OperatingSystemSpec != nil {
		in, out := &in.OperatingSystemSpec, &out.OperatingSystemSpec
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(NetworkConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.OverwriteCloudConfig != nil {
		in, out := &in.OverwriteCloudConfig, &out.OverwriteCloudConfig
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSpec.
func (in *ProviderSpec) DeepCopy() *ProviderSpec {
	if in == nil {
		return nil
	}
	out := new(ProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyConfig) DeepCopyInto(out *ProxyConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyConfig.
func (in *ProxyConfig) DeepCopy() *ProxyConfig {
	if in == nil {
		return nil
	}
	out := new(ProxyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionConfig) DeepCopyInto(out *VersionConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionConfig.
func (in *VersionConfig) DeepCopy() *VersionConfig {
	if in == nil {
		return nil
	}
	out := new(VersionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerConfig) DeepCopyInto(out *WorkerConfig) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int)
		**out = **in
	}
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerConfig.
func (in *WorkerConfig) DeepCopy() *WorkerConfig {
	if in == nil {
		return nil
	}
	out := new(WorkerConfig)
	in.DeepCopyInto(out)
	return out
}
