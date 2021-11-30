//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdomainservices

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// OuContainerClient contains the methods for the OuContainer group.
// Don't use this type directly, use NewOuContainerClient() instead.
type OuContainerClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewOuContainerClient creates a new instance of OuContainerClient with the specified values.
func NewOuContainerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *OuContainerClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &OuContainerClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreate - The Create OuContainer operation creates a new OuContainer under the specified Domain Service instance.
// If the operation fails it returns the *CloudError error type.
func (client *OuContainerClient) BeginCreate(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount ContainerAccount, options *OuContainerBeginCreateOptions) (OuContainerCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, domainServiceName, ouContainerName, containerAccount, options)
	if err != nil {
		return OuContainerCreatePollerResponse{}, err
	}
	result := OuContainerCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OuContainerClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return OuContainerCreatePollerResponse{}, err
	}
	result.Poller = &OuContainerCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - The Create OuContainer operation creates a new OuContainer under the specified Domain Service instance.
// If the operation fails it returns the *CloudError error type.
func (client *OuContainerClient) create(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount ContainerAccount, options *OuContainerBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, domainServiceName, ouContainerName, containerAccount, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *OuContainerClient) createCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount ContainerAccount, options *OuContainerBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Aad/domainServices/{domainServiceName}/ouContainer/{ouContainerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	if ouContainerName == "" {
		return nil, errors.New("parameter ouContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ouContainerName}", url.PathEscape(ouContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, containerAccount)
}

// createHandleError handles the Create error response.
func (client *OuContainerClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - The Delete OuContainer operation deletes specified OuContainer.
// If the operation fails it returns the *CloudError error type.
func (client *OuContainerClient) BeginDelete(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, options *OuContainerBeginDeleteOptions) (OuContainerDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, domainServiceName, ouContainerName, options)
	if err != nil {
		return OuContainerDeletePollerResponse{}, err
	}
	result := OuContainerDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OuContainerClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return OuContainerDeletePollerResponse{}, err
	}
	result.Poller = &OuContainerDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - The Delete OuContainer operation deletes specified OuContainer.
// If the operation fails it returns the *CloudError error type.
func (client *OuContainerClient) deleteOperation(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, options *OuContainerBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, domainServiceName, ouContainerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OuContainerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, options *OuContainerBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Aad/domainServices/{domainServiceName}/ouContainer/{ouContainerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	if ouContainerName == "" {
		return nil, errors.New("parameter ouContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ouContainerName}", url.PathEscape(ouContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *OuContainerClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get OuContainer in DomainService instance.
// If the operation fails it returns the *CloudError error type.
func (client *OuContainerClient) Get(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, options *OuContainerGetOptions) (OuContainerGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, domainServiceName, ouContainerName, options)
	if err != nil {
		return OuContainerGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OuContainerGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OuContainerGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OuContainerClient) getCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, options *OuContainerGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Aad/domainServices/{domainServiceName}/ouContainer/{ouContainerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	if ouContainerName == "" {
		return nil, errors.New("parameter ouContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ouContainerName}", url.PathEscape(ouContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OuContainerClient) getHandleResponse(resp *http.Response) (OuContainerGetResponse, error) {
	result := OuContainerGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OuContainer); err != nil {
		return OuContainerGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *OuContainerClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - The List of OuContainers in DomainService instance.
// If the operation fails it returns the *CloudError error type.
func (client *OuContainerClient) List(resourceGroupName string, domainServiceName string, options *OuContainerListOptions) *OuContainerListPager {
	return &OuContainerListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, domainServiceName, options)
		},
		advancer: func(ctx context.Context, resp OuContainerListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OuContainerListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *OuContainerClient) listCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, options *OuContainerListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Aad/domainServices/{domainServiceName}/ouContainer"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OuContainerClient) listHandleResponse(resp *http.Response) (OuContainerListResponse, error) {
	result := OuContainerListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OuContainerListResult); err != nil {
		return OuContainerListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *OuContainerClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - The Update OuContainer operation can be used to update the existing OuContainers.
// If the operation fails it returns the *CloudError error type.
func (client *OuContainerClient) BeginUpdate(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount ContainerAccount, options *OuContainerBeginUpdateOptions) (OuContainerUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, domainServiceName, ouContainerName, containerAccount, options)
	if err != nil {
		return OuContainerUpdatePollerResponse{}, err
	}
	result := OuContainerUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OuContainerClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return OuContainerUpdatePollerResponse{}, err
	}
	result.Poller = &OuContainerUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - The Update OuContainer operation can be used to update the existing OuContainers.
// If the operation fails it returns the *CloudError error type.
func (client *OuContainerClient) update(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount ContainerAccount, options *OuContainerBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, domainServiceName, ouContainerName, containerAccount, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *OuContainerClient) updateCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount ContainerAccount, options *OuContainerBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Aad/domainServices/{domainServiceName}/ouContainer/{ouContainerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	if ouContainerName == "" {
		return nil, errors.New("parameter ouContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ouContainerName}", url.PathEscape(ouContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, containerAccount)
}

// updateHandleError handles the Update error response.
func (client *OuContainerClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}