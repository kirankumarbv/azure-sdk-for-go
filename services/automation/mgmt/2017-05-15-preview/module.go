package automation

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

// ModuleClient is the automation Client
type ModuleClient struct {
	ManagementClient
}

// NewModuleClient creates an instance of the ModuleClient client.
func NewModuleClient(p pipeline.Pipeline) ModuleClient {
	return ModuleClient{NewManagementClient(p)}
}

// CreateOrUpdate create or Update the module identified by module name.
//
// automationAccountName is the automation account name. moduleName is the name of module. parameters is the create or
// update parameters for module.
func (client ModuleClient) CreateOrUpdate(ctx context.Context, automationAccountName string, moduleName string, parameters ModuleCreateOrUpdateParameters) (*Module, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}},
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.ModuleCreateOrUpdateProperties", name: null, rule: true,
				chain: []constraint{{target: "parameters.ModuleCreateOrUpdateProperties.ContentLink", name: null, rule: true,
					chain: []constraint{{target: "parameters.ModuleCreateOrUpdateProperties.ContentLink.ContentHash", name: null, rule: false,
						chain: []constraint{{target: "parameters.ModuleCreateOrUpdateProperties.ContentLink.ContentHash.Algorithm", name: null, rule: true, chain: nil},
							{target: "parameters.ModuleCreateOrUpdateProperties.ContentLink.ContentHash.Value", name: null, rule: true, chain: nil},
						}},
					}},
				}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createOrUpdatePreparer(automationAccountName, moduleName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Module), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ModuleClient) createOrUpdatePreparer(automationAccountName string, moduleName string, parameters ModuleCreateOrUpdateParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2015-10-31")
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
func (client ModuleClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusCreated, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Module{rawResponse: resp.Response()}
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

// Delete delete the module by name.
//
// automationAccountName is the automation account name. moduleName is the module name.
func (client ModuleClient) Delete(ctx context.Context, automationAccountName string, moduleName string) (*http.Response, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.deletePreparer(automationAccountName, moduleName)
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
func (client ModuleClient) deletePreparer(automationAccountName string, moduleName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2015-10-31")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client ModuleClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	return resp, err
}

// Get retrieve the module identified by module name.
//
// automationAccountName is the automation account name. moduleName is the module name.
func (client ModuleClient) Get(ctx context.Context, automationAccountName string, moduleName string) (*Module, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.getPreparer(automationAccountName, moduleName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Module), err
}

// getPreparer prepares the Get request.
func (client ModuleClient) getPreparer(automationAccountName string, moduleName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2015-10-31")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client ModuleClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Module{rawResponse: resp.Response()}
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

// ListByAutomationAccount retrieve a list of modules.
//
// automationAccountName is the automation account name.
func (client ModuleClient) ListByAutomationAccount(ctx context.Context, automationAccountName string) (*ModuleListResult, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.listByAutomationAccountPreparer(automationAccountName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listByAutomationAccountResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ModuleListResult), err
}

// listByAutomationAccountPreparer prepares the ListByAutomationAccount request.
func (client ModuleClient) listByAutomationAccountPreparer(automationAccountName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2015-10-31")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listByAutomationAccountResponder handles the response to the ListByAutomationAccount request.
func (client ModuleClient) listByAutomationAccountResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ModuleListResult{rawResponse: resp.Response()}
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

// Update update the module identified by module name.
//
// automationAccountName is the automation account name. moduleName is the name of module. parameters is the update
// parameters for module.
func (client ModuleClient) Update(ctx context.Context, automationAccountName string, moduleName string, parameters ModuleUpdateParameters) (*Module, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.updatePreparer(automationAccountName, moduleName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.updateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Module), err
}

// updatePreparer prepares the Update request.
func (client ModuleClient) updatePreparer(automationAccountName string, moduleName string, parameters ModuleUpdateParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}"
	req, err := pipeline.NewRequest("PATCH", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2015-10-31")
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
func (client ModuleClient) updateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Module{rawResponse: resp.Response()}
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
