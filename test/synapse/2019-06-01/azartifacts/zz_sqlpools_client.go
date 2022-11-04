//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

type sqlPoolsClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newSQLPoolsClient creates a new instance of sqlPoolsClient with the specified values.
//   - endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
//   - pl - the pipeline used for sending requests and handling responses.
func newSQLPoolsClient(endpoint string, pl runtime.Pipeline) *sqlPoolsClient {
	client := &sqlPoolsClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// Get - Get Sql Pool
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - sqlPoolName - The Sql Pool name
//   - options - SqlPoolsClientGetOptions contains the optional parameters for the sqlPoolsClient.Get method.
func (client *sqlPoolsClient) Get(ctx context.Context, sqlPoolName string, options *SqlPoolsClientGetOptions) (SqlPoolsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, sqlPoolName, options)
	if err != nil {
		return SqlPoolsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SqlPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SqlPoolsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *sqlPoolsClient) getCreateRequest(ctx context.Context, sqlPoolName string, options *SqlPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/sqlPools/{sqlPoolName}"
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *sqlPoolsClient) getHandleResponse(resp *http.Response) (SqlPoolsClientGetResponse, error) {
	result := SqlPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLPool); err != nil {
		return SqlPoolsClientGetResponse{}, err
	}
	return result, nil
}

// List - List Sql Pools
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - options - SqlPoolsClientListOptions contains the optional parameters for the sqlPoolsClient.List method.
func (client *sqlPoolsClient) List(ctx context.Context, options *SqlPoolsClientListOptions) (SqlPoolsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return SqlPoolsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SqlPoolsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SqlPoolsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *sqlPoolsClient) listCreateRequest(ctx context.Context, options *SqlPoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/sqlPools"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *sqlPoolsClient) listHandleResponse(resp *http.Response) (SqlPoolsClientListResponse, error) {
	result := SqlPoolsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLPoolInfoListResult); err != nil {
		return SqlPoolsClientListResponse{}, err
	}
	return result, nil
}