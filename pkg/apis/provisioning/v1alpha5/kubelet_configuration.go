/*
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

package v1alpha5

// KubeletConfiguration defines args to be used when configuring kubelet on provisioned nodes.
// They are a subset of the upstream types, recognizing not all options may be supported.
// Wherever possible, the types and names should reflect the upstream kubelet types.
type KubeletConfiguration struct {
	// clusterDNS is a list of IP addresses for the cluster DNS server.
	// Note that not all providers may use all addresses.
	//+optional
	ClusterDNS []string `json:"clusterDNS,omitempty"`
	// MaxPods is the maximum number of Pods that can run on the Kubelet.
	//+optional
	MaxPods *int32 `json:"maxPods,omitempty"`
	// PodsPerCore is the maximum number of pods per core. Cannot exceed maxPods. The value must be a non-negative integer.
	// If 0, there is no limit on the number of Pods.
	//+optional
	PodsPerCore *int32 `json:"podsPerCore,omitempty"`
}
