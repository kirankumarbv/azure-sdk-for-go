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

package functions

import original "github.com/Azure/azure-sdk-for-go/service/streamanalytics/management/2016-03-01/functions"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type GroupClient = original.GroupClient
type UdfType = original.UdfType

const (
	Scalar UdfType = original.Scalar
)

type AzureMachineLearningWebServiceFunctionBinding = original.AzureMachineLearningWebServiceFunctionBinding
type AzureMachineLearningWebServiceFunctionBindingProperties = original.AzureMachineLearningWebServiceFunctionBindingProperties
type AzureMachineLearningWebServiceFunctionBindingRetrievalProperties = original.AzureMachineLearningWebServiceFunctionBindingRetrievalProperties
type AzureMachineLearningWebServiceFunctionRetrieveDefaultDefinitionParameters = original.AzureMachineLearningWebServiceFunctionRetrieveDefaultDefinitionParameters
type AzureMachineLearningWebServiceInputColumn = original.AzureMachineLearningWebServiceInputColumn
type AzureMachineLearningWebServiceInputs = original.AzureMachineLearningWebServiceInputs
type AzureMachineLearningWebServiceOutputColumn = original.AzureMachineLearningWebServiceOutputColumn
type Binding = original.Binding
type ErrorResponse = original.ErrorResponse
type Function = original.Function
type Input = original.Input
type JavaScriptFunctionBinding = original.JavaScriptFunctionBinding
type JavaScriptFunctionBindingProperties = original.JavaScriptFunctionBindingProperties
type JavaScriptFunctionBindingRetrievalProperties = original.JavaScriptFunctionBindingRetrievalProperties
type JavaScriptFunctionRetrieveDefaultDefinitionParameters = original.JavaScriptFunctionRetrieveDefaultDefinitionParameters
type ListResult = original.ListResult
type Output = original.Output
type Properties = original.Properties
type ResourceTestStatus = original.ResourceTestStatus
type RetrieveDefaultDefinitionParameters = original.RetrieveDefaultDefinitionParameters
type ScalarFunctionConfiguration = original.ScalarFunctionConfiguration
type ScalarFunctionProperties = original.ScalarFunctionProperties
type SubResource = original.SubResource

func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewGroupClient(subscriptionID string) GroupClient {
	return original.NewGroupClient(subscriptionID)
}
func NewGroupClientWithBaseURI(baseURI string, subscriptionID string) GroupClient {
	return original.NewGroupClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
