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

// RouteTablesServer is a fake server for instances of the armnetwork.RouteTablesClient type.
type RouteTablesServer struct {
	// BeginCreateOrUpdate is the fake for method RouteTablesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, routeTableName string, parameters armnetwork.RouteTable, options *armnetwork.RouteTablesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.RouteTablesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method RouteTablesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, routeTableName string, options *armnetwork.RouteTablesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.RouteTablesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RouteTablesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, routeTableName string, options *armnetwork.RouteTablesClientGetOptions) (resp azfake.Responder[armnetwork.RouteTablesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RouteTablesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.RouteTablesClientListOptions) (resp azfake.PagerResponder[armnetwork.RouteTablesClientListResponse])

	// NewListAllPager is the fake for method RouteTablesClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armnetwork.RouteTablesClientListAllOptions) (resp azfake.PagerResponder[armnetwork.RouteTablesClientListAllResponse])

	// UpdateTags is the fake for method RouteTablesClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, routeTableName string, parameters armnetwork.TagsObject, options *armnetwork.RouteTablesClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.RouteTablesClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewRouteTablesServerTransport creates a new instance of RouteTablesServerTransport with the provided implementation.
// The returned RouteTablesServerTransport instance is connected to an instance of armnetwork.RouteTablesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRouteTablesServerTransport(srv *RouteTablesServer) *RouteTablesServerTransport {
	return &RouteTablesServerTransport{srv: srv}
}

// RouteTablesServerTransport connects instances of armnetwork.RouteTablesClient to instances of RouteTablesServer.
// Don't use this type directly, use NewRouteTablesServerTransport instead.
type RouteTablesServerTransport struct {
	srv                 *RouteTablesServer
	beginCreateOrUpdate *azfake.PollerResponder[armnetwork.RouteTablesClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armnetwork.RouteTablesClientDeleteResponse]
	newListPager        *azfake.PagerResponder[armnetwork.RouteTablesClientListResponse]
	newListAllPager     *azfake.PagerResponder[armnetwork.RouteTablesClientListAllResponse]
}

// Do implements the policy.Transporter interface for RouteTablesServerTransport.
func (r *RouteTablesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RouteTablesClient.BeginCreateOrUpdate":
		resp, err = r.dispatchBeginCreateOrUpdate(req)
	case "RouteTablesClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "RouteTablesClient.Get":
		resp, err = r.dispatchGet(req)
	case "RouteTablesClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	case "RouteTablesClient.NewListAllPager":
		resp, err = r.dispatchNewListAllPager(req)
	case "RouteTablesClient.UpdateTags":
		resp, err = r.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RouteTablesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if r.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/routeTables/(?P<routeTableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.RouteTable](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		routeTableNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("routeTableName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, routeTableNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		r.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(r.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(r.beginCreateOrUpdate) {
		r.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (r *RouteTablesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if r.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/routeTables/(?P<routeTableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		routeTableNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("routeTableName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, routeTableNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		r.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(r.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(r.beginDelete) {
		r.beginDelete = nil
	}

	return resp, nil
}

func (r *RouteTablesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/routeTables/(?P<routeTableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	routeTableNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("routeTableName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armnetwork.RouteTablesClientGetOptions
	if expandParam != nil {
		options = &armnetwork.RouteTablesClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameUnescaped, routeTableNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RouteTable, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RouteTablesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if r.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/routeTables`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListPager(resourceGroupNameUnescaped, nil)
		r.newListPager = &resp
		server.PagerResponderInjectNextLinks(r.newListPager, req, func(page *armnetwork.RouteTablesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(r.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(r.newListPager) {
		r.newListPager = nil
	}
	return resp, nil
}

func (r *RouteTablesServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	if r.newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/routeTables`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := r.srv.NewListAllPager(nil)
		r.newListAllPager = &resp
		server.PagerResponderInjectNextLinks(r.newListAllPager, req, func(page *armnetwork.RouteTablesClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(r.newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(r.newListAllPager) {
		r.newListAllPager = nil
	}
	return resp, nil
}

func (r *RouteTablesServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if r.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/routeTables/(?P<routeTableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	routeTableNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("routeTableName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.UpdateTags(req.Context(), resourceGroupNameUnescaped, routeTableNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RouteTable, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
