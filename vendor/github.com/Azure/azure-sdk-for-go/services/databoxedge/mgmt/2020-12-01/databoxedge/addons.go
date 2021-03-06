package databoxedge

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AddonsClient is the client for the Addons methods of the Databoxedge service.
type AddonsClient struct {
	BaseClient
}

// NewAddonsClient creates an instance of the AddonsClient client.
func NewAddonsClient(subscriptionID string) AddonsClient {
	return NewAddonsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAddonsClientWithBaseURI creates an instance of the AddonsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAddonsClientWithBaseURI(baseURI string, subscriptionID string) AddonsClient {
	return AddonsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a addon.
// Parameters:
// deviceName - the device name.
// roleName - the role name.
// addonName - the addon name.
// addon - the addon properties.
// resourceGroupName - the resource group name.
func (client AddonsClient) CreateOrUpdate(ctx context.Context, deviceName string, roleName string, addonName string, addon BasicAddon, resourceGroupName string) (result AddonsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddonsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, deviceName, roleName, addonName, addon, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.AddonsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.AddonsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AddonsClient) CreateOrUpdatePreparer(ctx context.Context, deviceName string, roleName string, addonName string, addon BasicAddon, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"addonName":         autorest.Encode("path", addonName),
		"deviceName":        autorest.Encode("path", deviceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"roleName":          autorest.Encode("path", roleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/addons/{addonName}", pathParameters),
		autorest.WithJSON(addon),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AddonsClient) CreateOrUpdateSender(req *http.Request) (future AddonsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AddonsClient) CreateOrUpdateResponder(resp *http.Response) (result AddonModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the addon on the device.
// Parameters:
// deviceName - the device name.
// roleName - the role name.
// addonName - the addon name.
// resourceGroupName - the resource group name.
func (client AddonsClient) Delete(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string) (result AddonsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddonsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, deviceName, roleName, addonName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.AddonsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.AddonsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AddonsClient) DeletePreparer(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"addonName":         autorest.Encode("path", addonName),
		"deviceName":        autorest.Encode("path", deviceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"roleName":          autorest.Encode("path", roleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/addons/{addonName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AddonsClient) DeleteSender(req *http.Request) (future AddonsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AddonsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a specific addon by name.
// Parameters:
// deviceName - the device name.
// roleName - the role name.
// addonName - the addon name.
// resourceGroupName - the resource group name.
func (client AddonsClient) Get(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string) (result AddonModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddonsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, deviceName, roleName, addonName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.AddonsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databoxedge.AddonsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.AddonsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AddonsClient) GetPreparer(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"addonName":         autorest.Encode("path", addonName),
		"deviceName":        autorest.Encode("path", deviceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"roleName":          autorest.Encode("path", roleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/addons/{addonName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AddonsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AddonsClient) GetResponder(resp *http.Response) (result AddonModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByRole lists all the addons configured in the role.
// Parameters:
// deviceName - the device name.
// roleName - the role name.
// resourceGroupName - the resource group name.
func (client AddonsClient) ListByRole(ctx context.Context, deviceName string, roleName string, resourceGroupName string) (result AddonListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddonsClient.ListByRole")
		defer func() {
			sc := -1
			if result.al.Response.Response != nil {
				sc = result.al.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByRoleNextResults
	req, err := client.ListByRolePreparer(ctx, deviceName, roleName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.AddonsClient", "ListByRole", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByRoleSender(req)
	if err != nil {
		result.al.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databoxedge.AddonsClient", "ListByRole", resp, "Failure sending request")
		return
	}

	result.al, err = client.ListByRoleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.AddonsClient", "ListByRole", resp, "Failure responding to request")
		return
	}
	if result.al.hasNextLink() && result.al.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByRolePreparer prepares the ListByRole request.
func (client AddonsClient) ListByRolePreparer(ctx context.Context, deviceName string, roleName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"roleName":          autorest.Encode("path", roleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/addons", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByRoleSender sends the ListByRole request. The method will close the
// http.Response Body if it receives an error.
func (client AddonsClient) ListByRoleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByRoleResponder handles the response to the ListByRole request. The method always
// closes the http.Response Body.
func (client AddonsClient) ListByRoleResponder(resp *http.Response) (result AddonList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByRoleNextResults retrieves the next set of results, if any.
func (client AddonsClient) listByRoleNextResults(ctx context.Context, lastResults AddonList) (result AddonList, err error) {
	req, err := lastResults.addonListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "databoxedge.AddonsClient", "listByRoleNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByRoleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "databoxedge.AddonsClient", "listByRoleNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByRoleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.AddonsClient", "listByRoleNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByRoleComplete enumerates all values, automatically crossing page boundaries as required.
func (client AddonsClient) ListByRoleComplete(ctx context.Context, deviceName string, roleName string, resourceGroupName string) (result AddonListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddonsClient.ListByRole")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByRole(ctx, deviceName, roleName, resourceGroupName)
	return
}
