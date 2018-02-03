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

// SourceControlClient is the automation Client
type SourceControlClient struct {
	ManagementClient
}

// NewSourceControlClient creates an instance of the SourceControlClient client.
func NewSourceControlClient(p pipeline.Pipeline) SourceControlClient {
	return SourceControlClient{NewManagementClient(p)}
}

// CreateOrUpdate create a source control.
//
// automationAccountName is the automation account name. sourceControlName is the source control name. parameters is
// the parameters supplied to the create or update source control operation.
func (client SourceControlClient) CreateOrUpdate(ctx context.Context, automationAccountName string, sourceControlName string, parameters SourceControlCreateOrUpdateParameters) (*SourceControl, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}},
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.SourceControlCreateOrUpdateProperties", name: null, rule: true,
				chain: []constraint{{target: "parameters.SourceControlCreateOrUpdateProperties.RepoURL", name: null, rule: false,
					chain: []constraint{{target: "parameters.SourceControlCreateOrUpdateProperties.RepoURL", name: maxLength, rule: 2000, chain: nil}}},
					{target: "parameters.SourceControlCreateOrUpdateProperties.Branch", name: null, rule: false,
						chain: []constraint{{target: "parameters.SourceControlCreateOrUpdateProperties.Branch", name: maxLength, rule: 255, chain: nil}}},
					{target: "parameters.SourceControlCreateOrUpdateProperties.FolderPath", name: null, rule: false,
						chain: []constraint{{target: "parameters.SourceControlCreateOrUpdateProperties.FolderPath", name: maxLength, rule: 255, chain: nil}}},
					{target: "parameters.SourceControlCreateOrUpdateProperties.SecurityToken", name: null, rule: false,
						chain: []constraint{{target: "parameters.SourceControlCreateOrUpdateProperties.SecurityToken", name: maxLength, rule: 1024, chain: nil}}},
					{target: "parameters.SourceControlCreateOrUpdateProperties.Description", name: null, rule: false,
						chain: []constraint{{target: "parameters.SourceControlCreateOrUpdateProperties.Description", name: maxLength, rule: 512, chain: nil}}},
				}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createOrUpdatePreparer(automationAccountName, sourceControlName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*SourceControl), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SourceControlClient) createOrUpdatePreparer(automationAccountName string, sourceControlName string, parameters SourceControlCreateOrUpdateParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}"
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
func (client SourceControlClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	result := &SourceControl{rawResponse: resp.Response()}
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

// Delete delete the source control.
//
// automationAccountName is the automation account name. sourceControlName is the name of source control.
func (client SourceControlClient) Delete(ctx context.Context, automationAccountName string, sourceControlName string) (*http.Response, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.deletePreparer(automationAccountName, sourceControlName)
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
func (client SourceControlClient) deletePreparer(automationAccountName string, sourceControlName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}"
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
func (client SourceControlClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	return resp, err
}

// Get retrieve the source control identified by source control name.
//
// automationAccountName is the automation account name. sourceControlName is the name of source control.
func (client SourceControlClient) Get(ctx context.Context, automationAccountName string, sourceControlName string) (*SourceControl, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.getPreparer(automationAccountName, sourceControlName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*SourceControl), err
}

// getPreparer prepares the Get request.
func (client SourceControlClient) getPreparer(automationAccountName string, sourceControlName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}"
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
func (client SourceControlClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &SourceControl{rawResponse: resp.Response()}
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

// ListByAutomationAccount retrieve a list of source controls.
//
// automationAccountName is the automation account name. filter is the filter to apply on the operation.
func (client SourceControlClient) ListByAutomationAccount(ctx context.Context, automationAccountName string, filter *string) (*SourceControlListResult, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.listByAutomationAccountPreparer(automationAccountName, filter)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listByAutomationAccountResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*SourceControlListResult), err
}

// listByAutomationAccountPreparer prepares the ListByAutomationAccount request.
func (client SourceControlClient) listByAutomationAccountPreparer(automationAccountName string, filter *string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if filter != nil {
		params.Set("$filter", *filter)
	}
	params.Set("api-version", "2015-10-31")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listByAutomationAccountResponder handles the response to the ListByAutomationAccount request.
func (client SourceControlClient) listByAutomationAccountResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &SourceControlListResult{rawResponse: resp.Response()}
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

// Update update a source control.
//
// automationAccountName is the automation account name. sourceControlName is the source control name. parameters is
// the parameters supplied to the update source control operation.
func (client SourceControlClient) Update(ctx context.Context, automationAccountName string, sourceControlName string, parameters SourceControlUpdateParameters) (*SourceControl, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.updatePreparer(automationAccountName, sourceControlName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.updateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*SourceControl), err
}

// updatePreparer prepares the Update request.
func (client SourceControlClient) updatePreparer(automationAccountName string, sourceControlName string, parameters SourceControlUpdateParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}"
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
func (client SourceControlClient) updateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &SourceControl{rawResponse: resp.Response()}
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
