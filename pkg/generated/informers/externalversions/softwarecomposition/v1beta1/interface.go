/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/Aryaman6492/storage/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ApplicationActivities returns a ApplicationActivityInformer.
	ApplicationActivities() ApplicationActivityInformer
	// ApplicationProfiles returns a ApplicationProfileInformer.
	ApplicationProfiles() ApplicationProfileInformer
	// ConfigurationScanSummaries returns a ConfigurationScanSummaryInformer.
	ConfigurationScanSummaries() ConfigurationScanSummaryInformer
	// GeneratedNetworkPolicies returns a GeneratedNetworkPolicyInformer.
	GeneratedNetworkPolicies() GeneratedNetworkPolicyInformer
	// KnownServers returns a KnownServerInformer.
	KnownServers() KnownServerInformer
	// NetworkNeighborhoods returns a NetworkNeighborhoodInformer.
	NetworkNeighborhoods() NetworkNeighborhoodInformer
	// NetworkNeighborses returns a NetworkNeighborsInformer.
	NetworkNeighborses() NetworkNeighborsInformer
	// OpenVulnerabilityExchangeContainers returns a OpenVulnerabilityExchangeContainerInformer.
	OpenVulnerabilityExchangeContainers() OpenVulnerabilityExchangeContainerInformer
	// SBOMSyfts returns a SBOMSyftInformer.
	SBOMSyfts() SBOMSyftInformer
	// SBOMSyftFiltereds returns a SBOMSyftFilteredInformer.
	SBOMSyftFiltereds() SBOMSyftFilteredInformer
	// SeccompProfiles returns a SeccompProfileInformer.
	SeccompProfiles() SeccompProfileInformer
	// VulnerabilityManifests returns a VulnerabilityManifestInformer.
	VulnerabilityManifests() VulnerabilityManifestInformer
	// VulnerabilityManifestSummaries returns a VulnerabilityManifestSummaryInformer.
	VulnerabilityManifestSummaries() VulnerabilityManifestSummaryInformer
	// VulnerabilitySummaries returns a VulnerabilitySummaryInformer.
	VulnerabilitySummaries() VulnerabilitySummaryInformer
	// WorkloadConfigurationScans returns a WorkloadConfigurationScanInformer.
	WorkloadConfigurationScans() WorkloadConfigurationScanInformer
	// WorkloadConfigurationScanSummaries returns a WorkloadConfigurationScanSummaryInformer.
	WorkloadConfigurationScanSummaries() WorkloadConfigurationScanSummaryInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ApplicationActivities returns a ApplicationActivityInformer.
func (v *version) ApplicationActivities() ApplicationActivityInformer {
	return &applicationActivityInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApplicationProfiles returns a ApplicationProfileInformer.
func (v *version) ApplicationProfiles() ApplicationProfileInformer {
	return &applicationProfileInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConfigurationScanSummaries returns a ConfigurationScanSummaryInformer.
func (v *version) ConfigurationScanSummaries() ConfigurationScanSummaryInformer {
	return &configurationScanSummaryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GeneratedNetworkPolicies returns a GeneratedNetworkPolicyInformer.
func (v *version) GeneratedNetworkPolicies() GeneratedNetworkPolicyInformer {
	return &generatedNetworkPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// KnownServers returns a KnownServerInformer.
func (v *version) KnownServers() KnownServerInformer {
	return &knownServerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NetworkNeighborhoods returns a NetworkNeighborhoodInformer.
func (v *version) NetworkNeighborhoods() NetworkNeighborhoodInformer {
	return &networkNeighborhoodInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NetworkNeighborses returns a NetworkNeighborsInformer.
func (v *version) NetworkNeighborses() NetworkNeighborsInformer {
	return &networkNeighborsInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OpenVulnerabilityExchangeContainers returns a OpenVulnerabilityExchangeContainerInformer.
func (v *version) OpenVulnerabilityExchangeContainers() OpenVulnerabilityExchangeContainerInformer {
	return &openVulnerabilityExchangeContainerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SBOMSyfts returns a SBOMSyftInformer.
func (v *version) SBOMSyfts() SBOMSyftInformer {
	return &sBOMSyftInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SBOMSyftFiltereds returns a SBOMSyftFilteredInformer.
func (v *version) SBOMSyftFiltereds() SBOMSyftFilteredInformer {
	return &sBOMSyftFilteredInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SeccompProfiles returns a SeccompProfileInformer.
func (v *version) SeccompProfiles() SeccompProfileInformer {
	return &seccompProfileInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VulnerabilityManifests returns a VulnerabilityManifestInformer.
func (v *version) VulnerabilityManifests() VulnerabilityManifestInformer {
	return &vulnerabilityManifestInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VulnerabilityManifestSummaries returns a VulnerabilityManifestSummaryInformer.
func (v *version) VulnerabilityManifestSummaries() VulnerabilityManifestSummaryInformer {
	return &vulnerabilityManifestSummaryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VulnerabilitySummaries returns a VulnerabilitySummaryInformer.
func (v *version) VulnerabilitySummaries() VulnerabilitySummaryInformer {
	return &vulnerabilitySummaryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// WorkloadConfigurationScans returns a WorkloadConfigurationScanInformer.
func (v *version) WorkloadConfigurationScans() WorkloadConfigurationScanInformer {
	return &workloadConfigurationScanInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// WorkloadConfigurationScanSummaries returns a WorkloadConfigurationScanSummaryInformer.
func (v *version) WorkloadConfigurationScanSummaries() WorkloadConfigurationScanSummaryInformer {
	return &workloadConfigurationScanSummaryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
