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
	"strconv"
	"strings"
)

// PipelineClient contains the methods for the Pipeline group.
// Don't use this type directly, use a constructor function instead.
type PipelineClient struct {
	internal *azcore.Client
	endpoint string
}

// BeginCreateOrUpdatePipeline - Creates or updates a pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - pipelineName - The pipeline name.
//   - pipeline - Pipeline resource definition.
//   - options - PipelineClientBeginCreateOrUpdatePipelineOptions contains the optional parameters for the PipelineClient.BeginCreateOrUpdatePipeline
//     method.
func (client *PipelineClient) BeginCreateOrUpdatePipeline(ctx context.Context, pipelineName string, pipeline PipelineResource, options *PipelineClientBeginCreateOrUpdatePipelineOptions) (*runtime.Poller[PipelineClientCreateOrUpdatePipelineResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdatePipeline(ctx, pipelineName, pipeline, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[PipelineClientCreateOrUpdatePipelineResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PipelineClientCreateOrUpdatePipelineResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdatePipeline - Creates or updates a pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *PipelineClient) createOrUpdatePipeline(ctx context.Context, pipelineName string, pipeline PipelineResource, options *PipelineClientBeginCreateOrUpdatePipelineOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdatePipelineCreateRequest(ctx, pipelineName, pipeline, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdatePipelineCreateRequest creates the CreateOrUpdatePipeline request.
func (client *PipelineClient) createOrUpdatePipelineCreateRequest(ctx context.Context, pipelineName string, pipeline PipelineResource, options *PipelineClientBeginCreateOrUpdatePipelineOptions) (*policy.Request, error) {
	urlPath := "/pipelines/{pipelineName}"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, pipeline); err != nil {
		return nil, err
	}
	return req, nil
}

// CreatePipelineRun - Creates a run of a pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - pipelineName - The pipeline name.
//   - options - PipelineClientCreatePipelineRunOptions contains the optional parameters for the PipelineClient.CreatePipelineRun
//     method.
func (client *PipelineClient) CreatePipelineRun(ctx context.Context, pipelineName string, options *PipelineClientCreatePipelineRunOptions) (PipelineClientCreatePipelineRunResponse, error) {
	var err error
	req, err := client.createPipelineRunCreateRequest(ctx, pipelineName, options)
	if err != nil {
		return PipelineClientCreatePipelineRunResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelineClientCreatePipelineRunResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return PipelineClientCreatePipelineRunResponse{}, err
	}
	resp, err := client.createPipelineRunHandleResponse(httpResp)
	return resp, err
}

// createPipelineRunCreateRequest creates the CreatePipelineRun request.
func (client *PipelineClient) createPipelineRunCreateRequest(ctx context.Context, pipelineName string, options *PipelineClientCreatePipelineRunOptions) (*policy.Request, error) {
	urlPath := "/pipelines/{pipelineName}/createRun"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	if options != nil && options.ReferencePipelineRunID != nil {
		reqQP.Set("referencePipelineRunId", *options.ReferencePipelineRunID)
	}
	if options != nil && options.IsRecovery != nil {
		reqQP.Set("isRecovery", strconv.FormatBool(*options.IsRecovery))
	}
	if options != nil && options.StartActivityName != nil {
		reqQP.Set("startActivityName", *options.StartActivityName)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		if err := runtime.MarshalAsJSON(req, options.Parameters); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// createPipelineRunHandleResponse handles the CreatePipelineRun response.
func (client *PipelineClient) createPipelineRunHandleResponse(resp *http.Response) (PipelineClientCreatePipelineRunResponse, error) {
	result := PipelineClientCreatePipelineRunResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CreateRunResponse); err != nil {
		return PipelineClientCreatePipelineRunResponse{}, err
	}
	return result, nil
}

// BeginDeletePipeline - Deletes a pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - pipelineName - The pipeline name.
//   - options - PipelineClientBeginDeletePipelineOptions contains the optional parameters for the PipelineClient.BeginDeletePipeline
//     method.
func (client *PipelineClient) BeginDeletePipeline(ctx context.Context, pipelineName string, options *PipelineClientBeginDeletePipelineOptions) (*runtime.Poller[PipelineClientDeletePipelineResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deletePipeline(ctx, pipelineName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[PipelineClientDeletePipelineResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PipelineClientDeletePipelineResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// DeletePipeline - Deletes a pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *PipelineClient) deletePipeline(ctx context.Context, pipelineName string, options *PipelineClientBeginDeletePipelineOptions) (*http.Response, error) {
	var err error
	req, err := client.deletePipelineCreateRequest(ctx, pipelineName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deletePipelineCreateRequest creates the DeletePipeline request.
func (client *PipelineClient) deletePipelineCreateRequest(ctx context.Context, pipelineName string, options *PipelineClientBeginDeletePipelineOptions) (*policy.Request, error) {
	urlPath := "/pipelines/{pipelineName}"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetPipeline - Gets a pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - pipelineName - The pipeline name.
//   - options - PipelineClientGetPipelineOptions contains the optional parameters for the PipelineClient.GetPipeline method.
func (client *PipelineClient) GetPipeline(ctx context.Context, pipelineName string, options *PipelineClientGetPipelineOptions) (PipelineClientGetPipelineResponse, error) {
	var err error
	req, err := client.getPipelineCreateRequest(ctx, pipelineName, options)
	if err != nil {
		return PipelineClientGetPipelineResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelineClientGetPipelineResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotModified) {
		err = runtime.NewResponseError(httpResp)
		return PipelineClientGetPipelineResponse{}, err
	}
	resp, err := client.getPipelineHandleResponse(httpResp)
	return resp, err
}

// getPipelineCreateRequest creates the GetPipeline request.
func (client *PipelineClient) getPipelineCreateRequest(ctx context.Context, pipelineName string, options *PipelineClientGetPipelineOptions) (*policy.Request, error) {
	urlPath := "/pipelines/{pipelineName}"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getPipelineHandleResponse handles the GetPipeline response.
func (client *PipelineClient) getPipelineHandleResponse(resp *http.Response) (PipelineClientGetPipelineResponse, error) {
	result := PipelineClientGetPipelineResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineResource); err != nil {
		return PipelineClientGetPipelineResponse{}, err
	}
	return result, nil
}

// NewGetPipelinesByWorkspacePager - Lists pipelines.
//
// Generated from API version 2020-12-01
//   - options - PipelineClientGetPipelinesByWorkspaceOptions contains the optional parameters for the PipelineClient.NewGetPipelinesByWorkspacePager
//     method.
func (client *PipelineClient) NewGetPipelinesByWorkspacePager(options *PipelineClientGetPipelinesByWorkspaceOptions) *runtime.Pager[PipelineClientGetPipelinesByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[PipelineClientGetPipelinesByWorkspaceResponse]{
		More: func(page PipelineClientGetPipelinesByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PipelineClientGetPipelinesByWorkspaceResponse) (PipelineClientGetPipelinesByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getPipelinesByWorkspaceCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PipelineClientGetPipelinesByWorkspaceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PipelineClientGetPipelinesByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PipelineClientGetPipelinesByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.getPipelinesByWorkspaceHandleResponse(resp)
		},
	})
}

// getPipelinesByWorkspaceCreateRequest creates the GetPipelinesByWorkspace request.
func (client *PipelineClient) getPipelinesByWorkspaceCreateRequest(ctx context.Context, options *PipelineClientGetPipelinesByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/pipelines"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getPipelinesByWorkspaceHandleResponse handles the GetPipelinesByWorkspace response.
func (client *PipelineClient) getPipelinesByWorkspaceHandleResponse(resp *http.Response) (PipelineClientGetPipelinesByWorkspaceResponse, error) {
	result := PipelineClientGetPipelinesByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineListResponse); err != nil {
		return PipelineClientGetPipelinesByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginRenamePipeline - Renames a pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - pipelineName - The pipeline name.
//   - request - proposed new name.
//   - options - PipelineClientBeginRenamePipelineOptions contains the optional parameters for the PipelineClient.BeginRenamePipeline
//     method.
func (client *PipelineClient) BeginRenamePipeline(ctx context.Context, pipelineName string, request ArtifactRenameRequest, options *PipelineClientBeginRenamePipelineOptions) (*runtime.Poller[PipelineClientRenamePipelineResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.renamePipeline(ctx, pipelineName, request, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[PipelineClientRenamePipelineResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PipelineClientRenamePipelineResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// RenamePipeline - Renames a pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
func (client *PipelineClient) renamePipeline(ctx context.Context, pipelineName string, request ArtifactRenameRequest, options *PipelineClientBeginRenamePipelineOptions) (*http.Response, error) {
	var err error
	req, err := client.renamePipelineCreateRequest(ctx, pipelineName, request, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// renamePipelineCreateRequest creates the RenamePipeline request.
func (client *PipelineClient) renamePipelineCreateRequest(ctx context.Context, pipelineName string, request ArtifactRenameRequest, options *PipelineClientBeginRenamePipelineOptions) (*policy.Request, error) {
	urlPath := "/pipelines/{pipelineName}/rename"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, request); err != nil {
		return nil, err
	}
	return req, nil
}