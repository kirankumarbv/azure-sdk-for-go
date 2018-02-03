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
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
)

// NodeReportsClient is the automation Client
type NodeReportsClient struct {
	ManagementClient
}

// NewNodeReportsClient creates an instance of the NodeReportsClient client.
func NewNodeReportsClient(p pipeline.Pipeline) NodeReportsClient {
	return NodeReportsClient{NewManagementClient(p)}
}

// Get retrieve the Dsc node report data by node id and report id.
//
// automationAccountName is the automation account name. nodeID is the Dsc node id. reportID is the report id.
func (client NodeReportsClient) Get(ctx context.Context, automationAccountName string, nodeID string, reportID string) (*DscNodeReport, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.getPreparer(automationAccountName, nodeID, reportID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*DscNodeReport), err
}

// getPreparer prepares the Get request.
func (client NodeReportsClient) getPreparer(automationAccountName string, nodeID string, reportID string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}/reports/{reportId}"
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
func (client NodeReportsClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &DscNodeReport{rawResponse: resp.Response()}
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

// GetContent retrieve the Dsc node reports by node id and report id.
//
// automationAccountName is the automation account name. nodeID is the Dsc node id. reportID is the report id.
func (client NodeReportsClient) GetContent(ctx context.Context, automationAccountName string, nodeID string, reportID string) (*GetContentResponse, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.getContentPreparer(automationAccountName, nodeID, reportID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getContentResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*GetContentResponse), err
}

// getContentPreparer prepares the GetContent request.
func (client NodeReportsClient) getContentPreparer(automationAccountName string, nodeID string, reportID string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}/reports/{reportId}/content"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2015-10-31")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getContentResponder handles the response to the GetContent request.
func (client NodeReportsClient) getContentResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	return &GetContentResponse{rawResponse: resp.Response()}, err
}

// ListByNode retrieve the Dsc node report list by node id.
//
// automationAccountName is the automation account name. nodeID is the parameters supplied to the list operation.
// filter is the filter to apply on the operation.
func (client NodeReportsClient) ListByNode(ctx context.Context, automationAccountName string, nodeID string, filter *string) (*DscNodeReportListResult, error) {
	if err := validate([]validation{
		{targetValue: client.ResourceGroupName,
			constraints: []constraint{{target: "client.ResourceGroupName", name: pattern, rule: `^[-\w\._]+$`, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.listByNodePreparer(automationAccountName, nodeID, filter)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listByNodeResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*DscNodeReportListResult), err
}

// listByNodePreparer prepares the ListByNode request.
func (client NodeReportsClient) listByNodePreparer(automationAccountName string, nodeID string, filter *string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}/reports"
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

// listByNodeResponder handles the response to the ListByNode request.
func (client NodeReportsClient) listByNodeResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &DscNodeReportListResult{rawResponse: resp.Response()}
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
