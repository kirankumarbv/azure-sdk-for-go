package management

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// GroupSubscriptionsClient is the the Azure Management Groups API enables consolidation of multiple
// subscriptions/resources into an organizational hierarchy and centrally
// manage access control, policies, alerting and reporting for those resources.
type GroupSubscriptionsClient struct {
	BaseClient
}

// NewGroupSubscriptionsClient creates an instance of the GroupSubscriptionsClient client.
func NewGroupSubscriptionsClient(groupID string, subscriptionID string) GroupSubscriptionsClient {
	return NewGroupSubscriptionsClientWithBaseURI(DefaultBaseURI, groupID, subscriptionID)
}

// NewGroupSubscriptionsClientWithBaseURI creates an instance of the GroupSubscriptionsClient client.
func NewGroupSubscriptionsClientWithBaseURI(baseURI string, groupID string, subscriptionID string) GroupSubscriptionsClient {
	return GroupSubscriptionsClient{NewWithBaseURI(baseURI, groupID, subscriptionID)}
}

// Create associates existing subscription with the management group.
func (client GroupSubscriptionsClient) Create(ctx context.Context) (result autorest.Response, err error) {
	req, err := client.CreatePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupSubscriptionsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "management.GroupSubscriptionsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupSubscriptionsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client GroupSubscriptionsClient) CreatePreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":        autorest.Encode("path", client.GroupID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}/subscriptions/{subscriptionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.CacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(client.CacheControl)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client GroupSubscriptionsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client GroupSubscriptionsClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete de-associates subscription from the management group.
func (client GroupSubscriptionsClient) Delete(ctx context.Context) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupSubscriptionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "management.GroupSubscriptionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "management.GroupSubscriptionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client GroupSubscriptionsClient) DeletePreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":        autorest.Encode("path", client.GroupID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}/subscriptions/{subscriptionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.CacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(client.CacheControl)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client GroupSubscriptionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client GroupSubscriptionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
