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

package v1beta1

import (
	"net/http"

	v1beta1 "github.com/kubescape/storage/pkg/apis/softwarecomposition/v1beta1"
	"github.com/kubescape/storage/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type SpdxV1beta1Interface interface {
	RESTClient() rest.Interface
	ApplicationActivitiesGetter
	ApplicationProfilesGetter
	ConfigurationScanSummariesGetter
	GeneratedNetworkPoliciesGetter
	KnownServersGetter
	NetworkNeighborhoodsGetter
	NetworkNeighborsesGetter
	OpenVulnerabilityExchangeContainersGetter
	SBOMSyftsGetter
	SBOMSyftFilteredsGetter
	SeccompProfilesGetter
	VulnerabilityManifestsGetter
	VulnerabilityManifestSummariesGetter
	VulnerabilitySummariesGetter
	WorkloadConfigurationScansGetter
	WorkloadConfigurationScanSummariesGetter
}

// SpdxV1beta1Client is used to interact with features provided by the spdx.softwarecomposition.seclogic.io group.
type SpdxV1beta1Client struct {
	restClient rest.Interface
}

func (c *SpdxV1beta1Client) ApplicationActivities(namespace string) ApplicationActivityInterface {
	return newApplicationActivities(c, namespace)
}

func (c *SpdxV1beta1Client) ApplicationProfiles(namespace string) ApplicationProfileInterface {
	return newApplicationProfiles(c, namespace)
}

func (c *SpdxV1beta1Client) ConfigurationScanSummaries(namespace string) ConfigurationScanSummaryInterface {
	return newConfigurationScanSummaries(c, namespace)
}

func (c *SpdxV1beta1Client) GeneratedNetworkPolicies(namespace string) GeneratedNetworkPolicyInterface {
	return newGeneratedNetworkPolicies(c, namespace)
}

func (c *SpdxV1beta1Client) KnownServers(namespace string) KnownServerInterface {
	return newKnownServers(c, namespace)
}

func (c *SpdxV1beta1Client) NetworkNeighborhoods(namespace string) NetworkNeighborhoodInterface {
	return newNetworkNeighborhoods(c, namespace)
}

func (c *SpdxV1beta1Client) NetworkNeighborses(namespace string) NetworkNeighborsInterface {
	return newNetworkNeighborses(c, namespace)
}

func (c *SpdxV1beta1Client) OpenVulnerabilityExchangeContainers(namespace string) OpenVulnerabilityExchangeContainerInterface {
	return newOpenVulnerabilityExchangeContainers(c, namespace)
}

func (c *SpdxV1beta1Client) SBOMSyfts(namespace string) SBOMSyftInterface {
	return newSBOMSyfts(c, namespace)
}

func (c *SpdxV1beta1Client) SBOMSyftFiltereds(namespace string) SBOMSyftFilteredInterface {
	return newSBOMSyftFiltereds(c, namespace)
}

func (c *SpdxV1beta1Client) SeccompProfiles(namespace string) SeccompProfileInterface {
	return newSeccompProfiles(c, namespace)
}

func (c *SpdxV1beta1Client) VulnerabilityManifests(namespace string) VulnerabilityManifestInterface {
	return newVulnerabilityManifests(c, namespace)
}

func (c *SpdxV1beta1Client) VulnerabilityManifestSummaries(namespace string) VulnerabilityManifestSummaryInterface {
	return newVulnerabilityManifestSummaries(c, namespace)
}

func (c *SpdxV1beta1Client) VulnerabilitySummaries(namespace string) VulnerabilitySummaryInterface {
	return newVulnerabilitySummaries(c, namespace)
}

func (c *SpdxV1beta1Client) WorkloadConfigurationScans(namespace string) WorkloadConfigurationScanInterface {
	return newWorkloadConfigurationScans(c, namespace)
}

func (c *SpdxV1beta1Client) WorkloadConfigurationScanSummaries(namespace string) WorkloadConfigurationScanSummaryInterface {
	return newWorkloadConfigurationScanSummaries(c, namespace)
}

// NewForConfig creates a new SpdxV1beta1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*SpdxV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new SpdxV1beta1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*SpdxV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &SpdxV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new SpdxV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *SpdxV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SpdxV1beta1Client for the given RESTClient.
func New(c rest.Interface) *SpdxV1beta1Client {
	return &SpdxV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *SpdxV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}