//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// MetastoreClient contains the methods for the Metastore group.
// Don't use this type directly, use a constructor function instead.
type MetastoreClient struct {
	internal *azcore.Client
	endpoint string
}

// Delete - Remove files in Syms
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01-preview
//   - options - MetastoreClientDeleteOptions contains the optional parameters for the MetastoreClient.Delete method.
func (client *MetastoreClient) Delete(ctx context.Context, id string, options *MetastoreClientDeleteOptions) (MetastoreClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, id, options)
	if err != nil {
		return MetastoreClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetastoreClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MetastoreClientDeleteResponse{}, err
	}
	return MetastoreClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MetastoreClient) deleteCreateRequest(ctx context.Context, id string, options *MetastoreClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/metastore/databases/{id}"
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetDatabaseOperations - Gets status of the database
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01-preview
//   - options - MetastoreClientGetDatabaseOperationsOptions contains the optional parameters for the MetastoreClient.GetDatabaseOperations
//     method.
func (client *MetastoreClient) GetDatabaseOperations(ctx context.Context, id string, options *MetastoreClientGetDatabaseOperationsOptions) (MetastoreClientGetDatabaseOperationsResponse, error) {
	var err error
	req, err := client.getDatabaseOperationsCreateRequest(ctx, id, options)
	if err != nil {
		return MetastoreClientGetDatabaseOperationsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetastoreClientGetDatabaseOperationsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetastoreClientGetDatabaseOperationsResponse{}, err
	}
	resp, err := client.getDatabaseOperationsHandleResponse(httpResp)
	return resp, err
}

// getDatabaseOperationsCreateRequest creates the GetDatabaseOperations request.
func (client *MetastoreClient) getDatabaseOperationsCreateRequest(ctx context.Context, id string, options *MetastoreClientGetDatabaseOperationsOptions) (*policy.Request, error) {
	urlPath := "/metastore/create-database-operations/{id}"
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDatabaseOperationsHandleResponse handles the GetDatabaseOperations response.
func (client *MetastoreClient) getDatabaseOperationsHandleResponse(resp *http.Response) (MetastoreClientGetDatabaseOperationsResponse, error) {
	result := MetastoreClientGetDatabaseOperationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetastoreRequestSuccessResponse); err != nil {
		return MetastoreClientGetDatabaseOperationsResponse{}, err
	}
	return result, nil
}

// Register - Register files in Syms
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01-preview
//   - id - The name of the database to be created. The name can contain only alphanumeric characters and should not exceed 24
//     characters
//   - registerBody - The body for the register request
//   - options - MetastoreClientRegisterOptions contains the optional parameters for the MetastoreClient.Register method.
func (client *MetastoreClient) Register(ctx context.Context, id string, registerBody MetastoreRegisterObject, options *MetastoreClientRegisterOptions) (MetastoreClientRegisterResponse, error) {
	var err error
	req, err := client.registerCreateRequest(ctx, id, registerBody, options)
	if err != nil {
		return MetastoreClientRegisterResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetastoreClientRegisterResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return MetastoreClientRegisterResponse{}, err
	}
	resp, err := client.registerHandleResponse(httpResp)
	return resp, err
}

// registerCreateRequest creates the Register request.
func (client *MetastoreClient) registerCreateRequest(ctx context.Context, id string, registerBody MetastoreRegisterObject, options *MetastoreClientRegisterOptions) (*policy.Request, error) {
	urlPath := "/metastore/create-database-operations/{id}"
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, registerBody); err != nil {
		return nil, err
	}
	return req, nil
}

// registerHandleResponse handles the Register response.
func (client *MetastoreClient) registerHandleResponse(resp *http.Response) (MetastoreClientRegisterResponse, error) {
	result := MetastoreClientRegisterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetastoreRegistrationResponse); err != nil {
		return MetastoreClientRegisterResponse{}, err
	}
	return result, nil
}

// Update - Update files in Syms
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01-preview
//   - id - The name of the database to be updated
//   - updateBody - The body for the update request
//   - options - MetastoreClientUpdateOptions contains the optional parameters for the MetastoreClient.Update method.
func (client *MetastoreClient) Update(ctx context.Context, id string, updateBody MetastoreUpdateObject, options *MetastoreClientUpdateOptions) (MetastoreClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, id, updateBody, options)
	if err != nil {
		return MetastoreClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetastoreClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return MetastoreClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *MetastoreClient) updateCreateRequest(ctx context.Context, id string, updateBody MetastoreUpdateObject, options *MetastoreClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/metastore/update-database-operations/{id}"
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateBody); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *MetastoreClient) updateHandleResponse(resp *http.Response) (MetastoreClientUpdateResponse, error) {
	result := MetastoreClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetastoreUpdationResponse); err != nil {
		return MetastoreClientUpdateResponse{}, err
	}
	return result, nil
}
