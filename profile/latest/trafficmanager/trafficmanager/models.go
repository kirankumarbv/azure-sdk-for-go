// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: b19c730e3a5c747d9055c95884b5c0310f7f2f16

package trafficmanager

import original "github.com/Azure/azure-sdk-for-go/service/trafficmanager/management/2017-03-01/trafficmanager"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type EndpointsClient = original.EndpointsClient
type GeographicHierarchiesClient = original.GeographicHierarchiesClient
type CheckTrafficManagerRelativeDNSNameAvailabilityParameters = original.CheckTrafficManagerRelativeDNSNameAvailabilityParameters
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type DeleteOperationResult = original.DeleteOperationResult
type DNSConfig = original.DNSConfig
type Endpoint = original.Endpoint
type EndpointProperties = original.EndpointProperties
type GeographicHierarchy = original.GeographicHierarchy
type GeographicHierarchyProperties = original.GeographicHierarchyProperties
type MonitorConfig = original.MonitorConfig
type NameAvailability = original.NameAvailability
type Profile = original.Profile
type ProfileListResult = original.ProfileListResult
type ProfileProperties = original.ProfileProperties
type Region = original.Region
type Resource = original.Resource
type SubResource = original.SubResource
type ProfilesClient = original.ProfilesClient

func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewEndpointsClient(subscriptionID string) EndpointsClient {
	return original.NewEndpointsClient(subscriptionID)
}
func NewEndpointsClientWithBaseURI(baseURI string, subscriptionID string) EndpointsClient {
	return original.NewEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewGeographicHierarchiesClient(subscriptionID string) GeographicHierarchiesClient {
	return original.NewGeographicHierarchiesClient(subscriptionID)
}
func NewGeographicHierarchiesClientWithBaseURI(baseURI string, subscriptionID string) GeographicHierarchiesClient {
	return original.NewGeographicHierarchiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewProfilesClient(subscriptionID string) ProfilesClient {
	return original.NewProfilesClient(subscriptionID)
}
func NewProfilesClientWithBaseURI(baseURI string, subscriptionID string) ProfilesClient {
	return original.NewProfilesClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
