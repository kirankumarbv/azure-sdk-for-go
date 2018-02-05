package account

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
)

// ComputePoliciesClient is the creates an Azure Data Lake Analytics account management client.
type ComputePoliciesClient struct {
	ManagementClient
}

// NewComputePoliciesClient creates an instance of the ComputePoliciesClient client.
func NewComputePoliciesClient(p pipeline.Pipeline) ComputePoliciesClient {
	return ComputePoliciesClient{NewManagementClient(p)}
}

// CreateOrUpdate creates or updates the specified compute policy. During update, the compute policy with the specified
// name will be replaced with this new compute policy. An account supports, at most, 50 policies
//
// resourceGroupName is the name of the Azure resource group. accountName is the name of the Data Lake Analytics
// account. computePolicyName is the name of the compute policy to create or update. parameters is parameters supplied
// to create or update the compute policy. The max degree of parallelism per job property, min priority per job
// property, or both must be present.
func (client ComputePoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, computePolicyName string, parameters CreateOrUpdateComputePolicyParameters) (*ComputePolicy, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.CreateOrUpdateComputePolicyProperties", name: null, rule: true,
				chain: []constraint{{target: "parameters.CreateOrUpdateComputePolicyProperties.ObjectID", name: null, rule: true, chain: nil},
					{target: "parameters.CreateOrUpdateComputePolicyProperties.MaxDegreeOfParallelismPerJob", name: null, rule: false,
						chain: []constraint{{target: "parameters.CreateOrUpdateComputePolicyProperties.MaxDegreeOfParallelismPerJob", name: inclusiveMinimum, rule: 1, chain: nil}}},
					{target: "parameters.CreateOrUpdateComputePolicyProperties.MinPriorityPerJob", name: null, rule: false,
						chain: []constraint{{target: "parameters.CreateOrUpdateComputePolicyProperties.MinPriorityPerJob", name: inclusiveMinimum, rule: 1, chain: nil}}},
				}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createOrUpdatePreparer(resourceGroupName, accountName, computePolicyName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ComputePolicy), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ComputePoliciesClient) createOrUpdatePreparer(resourceGroupName string, accountName string, computePolicyName string, parameters CreateOrUpdateComputePolicyParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(parameters)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// createOrUpdateResponder handles the response to the CreateOrUpdate request.
func (client ComputePoliciesClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ComputePolicy{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Delete deletes the specified compute policy from the specified Data Lake Analytics account
//
// resourceGroupName is the name of the Azure resource group. accountName is the name of the Data Lake Analytics
// account. computePolicyName is the name of the compute policy to delete.
func (client ComputePoliciesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, computePolicyName string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, accountName, computePolicyName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.Response(), err
}

// deletePreparer prepares the Delete request.
func (client ComputePoliciesClient) deletePreparer(resourceGroupName string, accountName string, computePolicyName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client ComputePoliciesClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusNoContent)
	if resp == nil {
		return nil, err
	}
	return resp, err
}

// Get gets the specified Data Lake Analytics compute policy.
//
// resourceGroupName is the name of the Azure resource group. accountName is the name of the Data Lake Analytics
// account. computePolicyName is the name of the compute policy to retrieve.
func (client ComputePoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, computePolicyName string) (*ComputePolicy, error) {
	req, err := client.getPreparer(resourceGroupName, accountName, computePolicyName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ComputePolicy), err
}

// getPreparer prepares the Get request.
func (client ComputePoliciesClient) getPreparer(resourceGroupName string, accountName string, computePolicyName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client ComputePoliciesClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ComputePolicy{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// ListByAccount lists the Data Lake Analytics compute policies within the specified Data Lake Analytics account. An
// account supports, at most, 50 policies
//
// resourceGroupName is the name of the Azure resource group. accountName is the name of the Data Lake Analytics
// account.
func (client ComputePoliciesClient) ListByAccount(ctx context.Context, resourceGroupName string, accountName string) (*ComputePolicyListResult, error) {
	req, err := client.listByAccountPreparer(resourceGroupName, accountName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listByAccountResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ComputePolicyListResult), err
}

// listByAccountPreparer prepares the ListByAccount request.
func (client ComputePoliciesClient) listByAccountPreparer(resourceGroupName string, accountName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listByAccountResponder handles the response to the ListByAccount request.
func (client ComputePoliciesClient) listByAccountResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ComputePolicyListResult{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Update updates the specified compute policy.
//
// resourceGroupName is the name of the Azure resource group. accountName is the name of the Data Lake Analytics
// account. computePolicyName is the name of the compute policy to update. parameters is parameters supplied to update
// the compute policy.
func (client ComputePoliciesClient) Update(ctx context.Context, resourceGroupName string, accountName string, computePolicyName string, parameters *UpdateComputePolicyParameters) (*ComputePolicy, error) {
	req, err := client.updatePreparer(resourceGroupName, accountName, computePolicyName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.updateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ComputePolicy), err
}

// updatePreparer prepares the Update request.
func (client ComputePoliciesClient) updatePreparer(resourceGroupName string, accountName string, computePolicyName string, parameters *UpdateComputePolicyParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}"
	req, err := pipeline.NewRequest("PATCH", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(parameters)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// updateResponder handles the response to the Update request.
func (client ComputePoliciesClient) updateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ComputePolicy{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}