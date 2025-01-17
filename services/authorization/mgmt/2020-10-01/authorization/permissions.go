package authorization

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PermissionsClient is the client for the Permissions methods of the Authorization service.
type PermissionsClient struct {
	BaseClient
}

// NewPermissionsClient creates an instance of the PermissionsClient client.
func NewPermissionsClient(subscriptionID string) PermissionsClient {
	return NewPermissionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPermissionsClientWithBaseURI creates an instance of the PermissionsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPermissionsClientWithBaseURI(baseURI string, subscriptionID string) PermissionsClient {
	return PermissionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListForResource gets all permissions the caller has for a resource.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceProviderNamespace - the namespace of the resource provider.
// parentResourcePath - the parent resource identity.
// resourceType - the resource type of the resource.
// resourceName - the name of the resource to get the permissions for.
func (client PermissionsClient) ListForResource(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string) (result PermissionGetResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PermissionsClient.ListForResource")
		defer func() {
			sc := -1
			if result.pgr.Response.Response != nil {
				sc = result.pgr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("authorization.PermissionsClient", "ListForResource", err.Error())
	}

	result.fn = client.listForResourceNextResults
	req, err := client.ListForResourcePreparer(ctx, resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.PermissionsClient", "ListForResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListForResourceSender(req)
	if err != nil {
		result.pgr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.PermissionsClient", "ListForResource", resp, "Failure sending request")
		return
	}

	result.pgr, err = client.ListForResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.PermissionsClient", "ListForResource", resp, "Failure responding to request")
		return
	}
	if result.pgr.hasNextLink() && result.pgr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListForResourcePreparer prepares the ListForResource request.
func (client PermissionsClient) ListForResourcePreparer(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"resourceName":              autorest.Encode("path", resourceName),
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/permissions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListForResourceSender sends the ListForResource request. The method will close the
// http.Response Body if it receives an error.
func (client PermissionsClient) ListForResourceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListForResourceResponder handles the response to the ListForResource request. The method always
// closes the http.Response Body.
func (client PermissionsClient) ListForResourceResponder(resp *http.Response) (result PermissionGetResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listForResourceNextResults retrieves the next set of results, if any.
func (client PermissionsClient) listForResourceNextResults(ctx context.Context, lastResults PermissionGetResult) (result PermissionGetResult, err error) {
	req, err := lastResults.permissionGetResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.PermissionsClient", "listForResourceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListForResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization.PermissionsClient", "listForResourceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListForResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.PermissionsClient", "listForResourceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListForResourceComplete enumerates all values, automatically crossing page boundaries as required.
func (client PermissionsClient) ListForResourceComplete(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string) (result PermissionGetResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PermissionsClient.ListForResource")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListForResource(ctx, resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName)
	return
}

// ListForResourceGroup gets all permissions the caller has for a resource group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
func (client PermissionsClient) ListForResourceGroup(ctx context.Context, resourceGroupName string) (result PermissionGetResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PermissionsClient.ListForResourceGroup")
		defer func() {
			sc := -1
			if result.pgr.Response.Response != nil {
				sc = result.pgr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("authorization.PermissionsClient", "ListForResourceGroup", err.Error())
	}

	result.fn = client.listForResourceGroupNextResults
	req, err := client.ListForResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.PermissionsClient", "ListForResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListForResourceGroupSender(req)
	if err != nil {
		result.pgr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.PermissionsClient", "ListForResourceGroup", resp, "Failure sending request")
		return
	}

	result.pgr, err = client.ListForResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.PermissionsClient", "ListForResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.pgr.hasNextLink() && result.pgr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListForResourceGroupPreparer prepares the ListForResourceGroup request.
func (client PermissionsClient) ListForResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Authorization/permissions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListForResourceGroupSender sends the ListForResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client PermissionsClient) ListForResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListForResourceGroupResponder handles the response to the ListForResourceGroup request. The method always
// closes the http.Response Body.
func (client PermissionsClient) ListForResourceGroupResponder(resp *http.Response) (result PermissionGetResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listForResourceGroupNextResults retrieves the next set of results, if any.
func (client PermissionsClient) listForResourceGroupNextResults(ctx context.Context, lastResults PermissionGetResult) (result PermissionGetResult, err error) {
	req, err := lastResults.permissionGetResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.PermissionsClient", "listForResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListForResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization.PermissionsClient", "listForResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListForResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.PermissionsClient", "listForResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListForResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client PermissionsClient) ListForResourceGroupComplete(ctx context.Context, resourceGroupName string) (result PermissionGetResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PermissionsClient.ListForResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListForResourceGroup(ctx, resourceGroupName)
	return
}
