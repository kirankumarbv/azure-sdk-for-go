package authorization

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

// RoleAssignmentsClient is the role based access control provides you a way to apply granular level policy
// administration down to individual resources or resource groups. These operations enable you to manage role
// definitions and role assignments. A role definition describes the set of actions that can be performed on resources.
// A role assignment grants access to Azure Active Directory users.
type RoleAssignmentsClient struct {
	ManagementClient
}

// NewRoleAssignmentsClient creates an instance of the RoleAssignmentsClient client.
func NewRoleAssignmentsClient(p pipeline.Pipeline) RoleAssignmentsClient {
	return RoleAssignmentsClient{NewManagementClient(p)}
}

// Create creates a role assignment.
//
// scope is the scope of the role assignment to create. The scope can be any REST resource instance. For example, use
// '/subscriptions/{subscription-id}/' for a subscription,
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}'
// for a resource. roleAssignmentName is the name of the role assignment to create. It can be any valid GUID.
// parameters is parameters for the role assignment.
func (client RoleAssignmentsClient) Create(ctx context.Context, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters) (*RoleAssignment, error) {
	req, err := client.createPreparer(scope, roleAssignmentName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RoleAssignment), err
}

// createPreparer prepares the Create request.
func (client RoleAssignmentsClient) createPreparer(scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
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

// createResponder handles the response to the Create request.
func (client RoleAssignmentsClient) createResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	result := &RoleAssignment{rawResponse: resp.Response()}
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

// CreateByID creates a role assignment by ID.
//
// roleAssignmentID is the fully qualified ID of the role assignment, including the scope, resource name and resource
// type. Use the format, /{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}. Example:
// /subscriptions/{subId}/resourcegroups/{rgname}//providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}.
// parameters is parameters for the role assignment.
func (client RoleAssignmentsClient) CreateByID(ctx context.Context, roleAssignmentID string, parameters RoleAssignmentCreateParameters) (*RoleAssignment, error) {
	req, err := client.createByIDPreparer(roleAssignmentID, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createByIDResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RoleAssignment), err
}

// createByIDPreparer prepares the CreateByID request.
func (client RoleAssignmentsClient) createByIDPreparer(roleAssignmentID string, parameters RoleAssignmentCreateParameters) (pipeline.Request, error) {
	u := client.url
	u.Path = "/{roleAssignmentId}"
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

// createByIDResponder handles the response to the CreateByID request.
func (client RoleAssignmentsClient) createByIDResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	result := &RoleAssignment{rawResponse: resp.Response()}
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

// Delete deletes a role assignment.
//
// scope is the scope of the role assignment to delete. roleAssignmentName is the name of the role assignment to
// delete.
func (client RoleAssignmentsClient) Delete(ctx context.Context, scope string, roleAssignmentName string) (*RoleAssignment, error) {
	req, err := client.deletePreparer(scope, roleAssignmentName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RoleAssignment), err
}

// deletePreparer prepares the Delete request.
func (client RoleAssignmentsClient) deletePreparer(scope string, roleAssignmentName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
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
func (client RoleAssignmentsClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RoleAssignment{rawResponse: resp.Response()}
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

// DeleteByID deletes a role assignment.
//
// roleAssignmentID is the fully qualified ID of the role assignment, including the scope, resource name and resource
// type. Use the format, /{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}. Example:
// /subscriptions/{subId}/resourcegroups/{rgname}//providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}.
func (client RoleAssignmentsClient) DeleteByID(ctx context.Context, roleAssignmentID string) (*RoleAssignment, error) {
	req, err := client.deleteByIDPreparer(roleAssignmentID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteByIDResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RoleAssignment), err
}

// deleteByIDPreparer prepares the DeleteByID request.
func (client RoleAssignmentsClient) deleteByIDPreparer(roleAssignmentID string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/{roleAssignmentId}"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deleteByIDResponder handles the response to the DeleteByID request.
func (client RoleAssignmentsClient) deleteByIDResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RoleAssignment{rawResponse: resp.Response()}
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

// Get get the specified role assignment.
//
// scope is the scope of the role assignment. roleAssignmentName is the name of the role assignment to get.
func (client RoleAssignmentsClient) Get(ctx context.Context, scope string, roleAssignmentName string) (*RoleAssignment, error) {
	req, err := client.getPreparer(scope, roleAssignmentName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RoleAssignment), err
}

// getPreparer prepares the Get request.
func (client RoleAssignmentsClient) getPreparer(scope string, roleAssignmentName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
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
func (client RoleAssignmentsClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RoleAssignment{rawResponse: resp.Response()}
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

// GetByID gets a role assignment by ID.
//
// roleAssignmentID is the fully qualified ID of the role assignment, including the scope, resource name and resource
// type. Use the format, /{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}. Example:
// /subscriptions/{subId}/resourcegroups/{rgname}//providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}.
func (client RoleAssignmentsClient) GetByID(ctx context.Context, roleAssignmentID string) (*RoleAssignment, error) {
	req, err := client.getByIDPreparer(roleAssignmentID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getByIDResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RoleAssignment), err
}

// getByIDPreparer prepares the GetByID request.
func (client RoleAssignmentsClient) getByIDPreparer(roleAssignmentID string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/{roleAssignmentId}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getByIDResponder handles the response to the GetByID request.
func (client RoleAssignmentsClient) getByIDResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RoleAssignment{rawResponse: resp.Response()}
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

// List gets all role assignments for the subscription.
//
// filter is the filter to apply on the operation. Use $filter=atScope() to return all role assignments at or above the
// scope. Use $filter=principalId eq {id} to return all role assignments at, above or below the scope for the specified
// principal.
func (client RoleAssignmentsClient) List(ctx context.Context, filter *string) (*RoleAssignmentListResult, error) {
	req, err := client.listPreparer(filter)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RoleAssignmentListResult), err
}

// listPreparer prepares the List request.
func (client RoleAssignmentsClient) listPreparer(filter *string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/roleAssignments"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if filter != nil {
		params.Set("$filter", *filter)
	}
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listResponder handles the response to the List request.
func (client RoleAssignmentsClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RoleAssignmentListResult{rawResponse: resp.Response()}
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

// ListForResource gets role assignments for a resource.
//
// resourceGroupName is the name of the resource group. resourceProviderNamespace is the namespace of the resource
// provider. parentResourcePath is the parent resource identity. resourceType is the resource type of the resource.
// resourceName is the name of the resource to get role assignments for. filter is the filter to apply on the
// operation. Use $filter=atScope() to return all role assignments at or above the scope. Use $filter=principalId eq
// {id} to return all role assignments at, above or below the scope for the specified principal.
func (client RoleAssignmentsClient) ListForResource(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter *string) (*RoleAssignmentListResult, error) {
	req, err := client.listForResourcePreparer(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, filter)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listForResourceResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RoleAssignmentListResult), err
}

// listForResourcePreparer prepares the ListForResource request.
func (client RoleAssignmentsClient) listForResourcePreparer(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter *string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/roleAssignments"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if filter != nil {
		params.Set("$filter", *filter)
	}
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listForResourceResponder handles the response to the ListForResource request.
func (client RoleAssignmentsClient) listForResourceResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RoleAssignmentListResult{rawResponse: resp.Response()}
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

// ListForResourceGroup gets role assignments for a resource group.
//
// resourceGroupName is the name of the resource group. filter is the filter to apply on the operation. Use
// $filter=atScope() to return all role assignments at or above the scope. Use $filter=principalId eq {id} to return
// all role assignments at, above or below the scope for the specified principal.
func (client RoleAssignmentsClient) ListForResourceGroup(ctx context.Context, resourceGroupName string, filter *string) (*RoleAssignmentListResult, error) {
	req, err := client.listForResourceGroupPreparer(resourceGroupName, filter)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listForResourceGroupResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RoleAssignmentListResult), err
}

// listForResourceGroupPreparer prepares the ListForResourceGroup request.
func (client RoleAssignmentsClient) listForResourceGroupPreparer(resourceGroupName string, filter *string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/roleAssignments"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if filter != nil {
		params.Set("$filter", *filter)
	}
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listForResourceGroupResponder handles the response to the ListForResourceGroup request.
func (client RoleAssignmentsClient) listForResourceGroupResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RoleAssignmentListResult{rawResponse: resp.Response()}
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

// ListForScope gets role assignments for a scope.
//
// scope is the scope of the role assignments. filter is the filter to apply on the operation. Use $filter=atScope() to
// return all role assignments at or above the scope. Use $filter=principalId eq {id} to return all role assignments
// at, above or below the scope for the specified principal.
func (client RoleAssignmentsClient) ListForScope(ctx context.Context, scope string, filter *string) (*RoleAssignmentListResult, error) {
	req, err := client.listForScopePreparer(scope, filter)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listForScopeResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*RoleAssignmentListResult), err
}

// listForScopePreparer prepares the ListForScope request.
func (client RoleAssignmentsClient) listForScopePreparer(scope string, filter *string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/{scope}/providers/Microsoft.Authorization/roleAssignments"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if filter != nil {
		params.Set("$filter", *filter)
	}
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listForScopeResponder handles the response to the ListForScope request.
func (client RoleAssignmentsClient) listForScopeResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &RoleAssignmentListResult{rawResponse: resp.Response()}
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
