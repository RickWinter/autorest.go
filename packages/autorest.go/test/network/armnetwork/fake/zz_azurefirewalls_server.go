//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armnetwork"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"net/url"
	"regexp"
)

// AzureFirewallsServer is a fake server for instances of the armnetwork.AzureFirewallsClient type.
type AzureFirewallsServer struct {
	// BeginCreateOrUpdate is the fake for method AzureFirewallsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters armnetwork.AzureFirewall, options *armnetwork.AzureFirewallsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.AzureFirewallsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AzureFirewallsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, azureFirewallName string, options *armnetwork.AzureFirewallsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.AzureFirewallsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AzureFirewallsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, azureFirewallName string, options *armnetwork.AzureFirewallsClientGetOptions) (resp azfake.Responder[armnetwork.AzureFirewallsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AzureFirewallsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.AzureFirewallsClientListOptions) (resp azfake.PagerResponder[armnetwork.AzureFirewallsClientListResponse])

	// NewListAllPager is the fake for method AzureFirewallsClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armnetwork.AzureFirewallsClientListAllOptions) (resp azfake.PagerResponder[armnetwork.AzureFirewallsClientListAllResponse])

	// BeginListLearnedPrefixes is the fake for method AzureFirewallsClient.BeginListLearnedPrefixes
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginListLearnedPrefixes func(ctx context.Context, resourceGroupName string, azureFirewallName string, options *armnetwork.AzureFirewallsClientBeginListLearnedPrefixesOptions) (resp azfake.PollerResponder[armnetwork.AzureFirewallsClientListLearnedPrefixesResponse], errResp azfake.ErrorResponder)

	// BeginUpdateTags is the fake for method AzureFirewallsClient.BeginUpdateTags
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateTags func(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters armnetwork.TagsObject, options *armnetwork.AzureFirewallsClientBeginUpdateTagsOptions) (resp azfake.PollerResponder[armnetwork.AzureFirewallsClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewAzureFirewallsServerTransport creates a new instance of AzureFirewallsServerTransport with the provided implementation.
// The returned AzureFirewallsServerTransport instance is connected to an instance of armnetwork.AzureFirewallsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAzureFirewallsServerTransport(srv *AzureFirewallsServer) *AzureFirewallsServerTransport {
	return &AzureFirewallsServerTransport{srv: srv}
}

// AzureFirewallsServerTransport connects instances of armnetwork.AzureFirewallsClient to instances of AzureFirewallsServer.
// Don't use this type directly, use NewAzureFirewallsServerTransport instead.
type AzureFirewallsServerTransport struct {
	srv                      *AzureFirewallsServer
	beginCreateOrUpdate      *azfake.PollerResponder[armnetwork.AzureFirewallsClientCreateOrUpdateResponse]
	beginDelete              *azfake.PollerResponder[armnetwork.AzureFirewallsClientDeleteResponse]
	newListPager             *azfake.PagerResponder[armnetwork.AzureFirewallsClientListResponse]
	newListAllPager          *azfake.PagerResponder[armnetwork.AzureFirewallsClientListAllResponse]
	beginListLearnedPrefixes *azfake.PollerResponder[armnetwork.AzureFirewallsClientListLearnedPrefixesResponse]
	beginUpdateTags          *azfake.PollerResponder[armnetwork.AzureFirewallsClientUpdateTagsResponse]
}

// Do implements the policy.Transporter interface for AzureFirewallsServerTransport.
func (a *AzureFirewallsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AzureFirewallsClient.BeginCreateOrUpdate":
		resp, err = a.dispatchBeginCreateOrUpdate(req)
	case "AzureFirewallsClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AzureFirewallsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AzureFirewallsClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	case "AzureFirewallsClient.NewListAllPager":
		resp, err = a.dispatchNewListAllPager(req)
	case "AzureFirewallsClient.BeginListLearnedPrefixes":
		resp, err = a.dispatchBeginListLearnedPrefixes(req)
	case "AzureFirewallsClient.BeginUpdateTags":
		resp, err = a.dispatchBeginUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AzureFirewallsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if a.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/azureFirewalls/(?P<azureFirewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.AzureFirewall](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureFirewallNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("azureFirewallName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, azureFirewallNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(a.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginCreateOrUpdate) {
		a.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (a *AzureFirewallsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if a.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/azureFirewalls/(?P<azureFirewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureFirewallNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("azureFirewallName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, azureFirewallNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(a.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginDelete) {
		a.beginDelete = nil
	}

	return resp, nil
}

func (a *AzureFirewallsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/azureFirewalls/(?P<azureFirewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	azureFirewallNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("azureFirewallName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameUnescaped, azureFirewallNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureFirewall, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AzureFirewallsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if a.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/azureFirewalls`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(resourceGroupNameUnescaped, nil)
		a.newListPager = &resp
		server.PagerResponderInjectNextLinks(a.newListPager, req, func(page *armnetwork.AzureFirewallsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(a.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(a.newListPager) {
		a.newListPager = nil
	}
	return resp, nil
}

func (a *AzureFirewallsServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	if a.newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/azureFirewalls`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListAllPager(nil)
		a.newListAllPager = &resp
		server.PagerResponderInjectNextLinks(a.newListAllPager, req, func(page *armnetwork.AzureFirewallsClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(a.newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(a.newListAllPager) {
		a.newListAllPager = nil
	}
	return resp, nil
}

func (a *AzureFirewallsServerTransport) dispatchBeginListLearnedPrefixes(req *http.Request) (*http.Response, error) {
	if a.srv.BeginListLearnedPrefixes == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginListLearnedPrefixes not implemented")}
	}
	if a.beginListLearnedPrefixes == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/azureFirewalls/(?P<azureFirewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/learnedIPPrefixes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureFirewallNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("azureFirewallName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginListLearnedPrefixes(req.Context(), resourceGroupNameUnescaped, azureFirewallNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginListLearnedPrefixes = &respr
	}

	resp, err := server.PollerResponderNext(a.beginListLearnedPrefixes, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginListLearnedPrefixes) {
		a.beginListLearnedPrefixes = nil
	}

	return resp, nil
}

func (a *AzureFirewallsServerTransport) dispatchBeginUpdateTags(req *http.Request) (*http.Response, error) {
	if a.srv.BeginUpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateTags not implemented")}
	}
	if a.beginUpdateTags == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/azureFirewalls/(?P<azureFirewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureFirewallNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("azureFirewallName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginUpdateTags(req.Context(), resourceGroupNameUnescaped, azureFirewallNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		a.beginUpdateTags = &respr
	}

	resp, err := server.PollerResponderNext(a.beginUpdateTags, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(a.beginUpdateTags) {
		a.beginUpdateTags = nil
	}

	return resp, nil
}
