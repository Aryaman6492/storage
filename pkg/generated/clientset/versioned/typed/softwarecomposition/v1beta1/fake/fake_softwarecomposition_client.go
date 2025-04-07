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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/Aryaman6492/storage/pkg/generated/clientset/versioned/typed/softwarecomposition/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSpdxV1beta1 struct {
	*testing.Fake
}

func (c *FakeSpdxV1beta1) ApplicationActivities(namespace string) v1beta1.ApplicationActivityInterface {
	return &FakeApplicationActivities{c, namespace}
}

func (c *FakeSpdxV1beta1) ApplicationProfiles(namespace string) v1beta1.ApplicationProfileInterface {
	return &FakeApplicationProfiles{c, namespace}
}

func (c *FakeSpdxV1beta1) ConfigurationScanSummaries(namespace string) v1beta1.ConfigurationScanSummaryInterface {
	return &FakeConfigurationScanSummaries{c, namespace}
}

func (c *FakeSpdxV1beta1) GeneratedNetworkPolicies(namespace string) v1beta1.GeneratedNetworkPolicyInterface {
	return &FakeGeneratedNetworkPolicies{c, namespace}
}

func (c *FakeSpdxV1beta1) KnownServers(namespace string) v1beta1.KnownServerInterface {
	return &FakeKnownServers{c, namespace}
}

func (c *FakeSpdxV1beta1) NetworkNeighborhoods(namespace string) v1beta1.NetworkNeighborhoodInterface {
	return &FakeNetworkNeighborhoods{c, namespace}
}

func (c *FakeSpdxV1beta1) OpenVulnerabilityExchangeContainers(namespace string) v1beta1.OpenVulnerabilityExchangeContainerInterface {
	return &FakeOpenVulnerabilityExchangeContainers{c, namespace}
}

func (c *FakeSpdxV1beta1) SBOMSyfts(namespace string) v1beta1.SBOMSyftInterface {
	return &FakeSBOMSyfts{c, namespace}
}

func (c *FakeSpdxV1beta1) SBOMSyftFiltereds(namespace string) v1beta1.SBOMSyftFilteredInterface {
	return &FakeSBOMSyftFiltereds{c, namespace}
}

func (c *FakeSpdxV1beta1) SeccompProfiles(namespace string) v1beta1.SeccompProfileInterface {
	return &FakeSeccompProfiles{c, namespace}
}

func (c *FakeSpdxV1beta1) VulnerabilityManifests(namespace string) v1beta1.VulnerabilityManifestInterface {
	return &FakeVulnerabilityManifests{c, namespace}
}

func (c *FakeSpdxV1beta1) VulnerabilityManifestSummaries(namespace string) v1beta1.VulnerabilityManifestSummaryInterface {
	return &FakeVulnerabilityManifestSummaries{c, namespace}
}

func (c *FakeSpdxV1beta1) VulnerabilitySummaries(namespace string) v1beta1.VulnerabilitySummaryInterface {
	return &FakeVulnerabilitySummaries{c, namespace}
}

func (c *FakeSpdxV1beta1) WorkloadConfigurationScans(namespace string) v1beta1.WorkloadConfigurationScanInterface {
	return &FakeWorkloadConfigurationScans{c, namespace}
}

func (c *FakeSpdxV1beta1) WorkloadConfigurationScanSummaries(namespace string) v1beta1.WorkloadConfigurationScanSummaryInterface {
	return &FakeWorkloadConfigurationScanSummaries{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSpdxV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
