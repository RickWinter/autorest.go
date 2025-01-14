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

// ExpressRouteCircuitsServer is a fake server for instances of the armnetwork.ExpressRouteCircuitsClient type.
type ExpressRouteCircuitsServer struct {
	// BeginCreateOrUpdate is the fake for method ExpressRouteCircuitsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, circuitName string, parameters armnetwork.ExpressRouteCircuit, options *armnetwork.ExpressRouteCircuitsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteCircuitsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ExpressRouteCircuitsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, circuitName string, options *armnetwork.ExpressRouteCircuitsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteCircuitsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ExpressRouteCircuitsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, circuitName string, options *armnetwork.ExpressRouteCircuitsClientGetOptions) (resp azfake.Responder[armnetwork.ExpressRouteCircuitsClientGetResponse], errResp azfake.ErrorResponder)

	// GetPeeringStats is the fake for method ExpressRouteCircuitsClient.GetPeeringStats
	// HTTP status codes to indicate success: http.StatusOK
	GetPeeringStats func(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *armnetwork.ExpressRouteCircuitsClientGetPeeringStatsOptions) (resp azfake.Responder[armnetwork.ExpressRouteCircuitsClientGetPeeringStatsResponse], errResp azfake.ErrorResponder)

	// GetStats is the fake for method ExpressRouteCircuitsClient.GetStats
	// HTTP status codes to indicate success: http.StatusOK
	GetStats func(ctx context.Context, resourceGroupName string, circuitName string, options *armnetwork.ExpressRouteCircuitsClientGetStatsOptions) (resp azfake.Responder[armnetwork.ExpressRouteCircuitsClientGetStatsResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ExpressRouteCircuitsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.ExpressRouteCircuitsClientListOptions) (resp azfake.PagerResponder[armnetwork.ExpressRouteCircuitsClientListResponse])

	// NewListAllPager is the fake for method ExpressRouteCircuitsClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armnetwork.ExpressRouteCircuitsClientListAllOptions) (resp azfake.PagerResponder[armnetwork.ExpressRouteCircuitsClientListAllResponse])

	// BeginListArpTable is the fake for method ExpressRouteCircuitsClient.BeginListArpTable
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginListArpTable func(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *armnetwork.ExpressRouteCircuitsClientBeginListArpTableOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteCircuitsClientListArpTableResponse], errResp azfake.ErrorResponder)

	// BeginListRoutesTable is the fake for method ExpressRouteCircuitsClient.BeginListRoutesTable
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginListRoutesTable func(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *armnetwork.ExpressRouteCircuitsClientBeginListRoutesTableOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteCircuitsClientListRoutesTableResponse], errResp azfake.ErrorResponder)

	// BeginListRoutesTableSummary is the fake for method ExpressRouteCircuitsClient.BeginListRoutesTableSummary
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginListRoutesTableSummary func(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *armnetwork.ExpressRouteCircuitsClientBeginListRoutesTableSummaryOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteCircuitsClientListRoutesTableSummaryResponse], errResp azfake.ErrorResponder)

	// UpdateTags is the fake for method ExpressRouteCircuitsClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, circuitName string, parameters armnetwork.TagsObject, options *armnetwork.ExpressRouteCircuitsClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.ExpressRouteCircuitsClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewExpressRouteCircuitsServerTransport creates a new instance of ExpressRouteCircuitsServerTransport with the provided implementation.
// The returned ExpressRouteCircuitsServerTransport instance is connected to an instance of armnetwork.ExpressRouteCircuitsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExpressRouteCircuitsServerTransport(srv *ExpressRouteCircuitsServer) *ExpressRouteCircuitsServerTransport {
	return &ExpressRouteCircuitsServerTransport{srv: srv}
}

// ExpressRouteCircuitsServerTransport connects instances of armnetwork.ExpressRouteCircuitsClient to instances of ExpressRouteCircuitsServer.
// Don't use this type directly, use NewExpressRouteCircuitsServerTransport instead.
type ExpressRouteCircuitsServerTransport struct {
	srv                         *ExpressRouteCircuitsServer
	beginCreateOrUpdate         *azfake.PollerResponder[armnetwork.ExpressRouteCircuitsClientCreateOrUpdateResponse]
	beginDelete                 *azfake.PollerResponder[armnetwork.ExpressRouteCircuitsClientDeleteResponse]
	newListPager                *azfake.PagerResponder[armnetwork.ExpressRouteCircuitsClientListResponse]
	newListAllPager             *azfake.PagerResponder[armnetwork.ExpressRouteCircuitsClientListAllResponse]
	beginListArpTable           *azfake.PollerResponder[armnetwork.ExpressRouteCircuitsClientListArpTableResponse]
	beginListRoutesTable        *azfake.PollerResponder[armnetwork.ExpressRouteCircuitsClientListRoutesTableResponse]
	beginListRoutesTableSummary *azfake.PollerResponder[armnetwork.ExpressRouteCircuitsClientListRoutesTableSummaryResponse]
}

// Do implements the policy.Transporter interface for ExpressRouteCircuitsServerTransport.
func (e *ExpressRouteCircuitsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExpressRouteCircuitsClient.BeginCreateOrUpdate":
		resp, err = e.dispatchBeginCreateOrUpdate(req)
	case "ExpressRouteCircuitsClient.BeginDelete":
		resp, err = e.dispatchBeginDelete(req)
	case "ExpressRouteCircuitsClient.Get":
		resp, err = e.dispatchGet(req)
	case "ExpressRouteCircuitsClient.GetPeeringStats":
		resp, err = e.dispatchGetPeeringStats(req)
	case "ExpressRouteCircuitsClient.GetStats":
		resp, err = e.dispatchGetStats(req)
	case "ExpressRouteCircuitsClient.NewListPager":
		resp, err = e.dispatchNewListPager(req)
	case "ExpressRouteCircuitsClient.NewListAllPager":
		resp, err = e.dispatchNewListAllPager(req)
	case "ExpressRouteCircuitsClient.BeginListArpTable":
		resp, err = e.dispatchBeginListArpTable(req)
	case "ExpressRouteCircuitsClient.BeginListRoutesTable":
		resp, err = e.dispatchBeginListRoutesTable(req)
	case "ExpressRouteCircuitsClient.BeginListRoutesTableSummary":
		resp, err = e.dispatchBeginListRoutesTableSummary(req)
	case "ExpressRouteCircuitsClient.UpdateTags":
		resp, err = e.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExpressRouteCircuitsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if e.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ExpressRouteCircuit](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		circuitNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, circuitNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		e.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(e.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(e.beginCreateOrUpdate) {
		e.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (e *ExpressRouteCircuitsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if e.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if e.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		circuitNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, circuitNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		e.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(e.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(e.beginDelete) {
		e.beginDelete = nil
	}

	return resp, nil
}

func (e *ExpressRouteCircuitsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	circuitNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameUnescaped, circuitNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRouteCircuit, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExpressRouteCircuitsServerTransport) dispatchGetPeeringStats(req *http.Request) (*http.Response, error) {
	if e.srv.GetPeeringStats == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetPeeringStats not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stats`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	circuitNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
	if err != nil {
		return nil, err
	}
	peeringNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.GetPeeringStats(req.Context(), resourceGroupNameUnescaped, circuitNameUnescaped, peeringNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRouteCircuitStats, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExpressRouteCircuitsServerTransport) dispatchGetStats(req *http.Request) (*http.Response, error) {
	if e.srv.GetStats == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetStats not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stats`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	circuitNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.GetStats(req.Context(), resourceGroupNameUnescaped, circuitNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRouteCircuitStats, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExpressRouteCircuitsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if e.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCircuits`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListPager(resourceGroupNameUnescaped, nil)
		e.newListPager = &resp
		server.PagerResponderInjectNextLinks(e.newListPager, req, func(page *armnetwork.ExpressRouteCircuitsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(e.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(e.newListPager) {
		e.newListPager = nil
	}
	return resp, nil
}

func (e *ExpressRouteCircuitsServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	if e.newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCircuits`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := e.srv.NewListAllPager(nil)
		e.newListAllPager = &resp
		server.PagerResponderInjectNextLinks(e.newListAllPager, req, func(page *armnetwork.ExpressRouteCircuitsClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(e.newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(e.newListAllPager) {
		e.newListAllPager = nil
	}
	return resp, nil
}

func (e *ExpressRouteCircuitsServerTransport) dispatchBeginListArpTable(req *http.Request) (*http.Response, error) {
	if e.srv.BeginListArpTable == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginListArpTable not implemented")}
	}
	if e.beginListArpTable == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/arpTables/(?P<devicePath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		circuitNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
		if err != nil {
			return nil, err
		}
		peeringNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
		if err != nil {
			return nil, err
		}
		devicePathUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("devicePath")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginListArpTable(req.Context(), resourceGroupNameUnescaped, circuitNameUnescaped, peeringNameUnescaped, devicePathUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		e.beginListArpTable = &respr
	}

	resp, err := server.PollerResponderNext(e.beginListArpTable, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(e.beginListArpTable) {
		e.beginListArpTable = nil
	}

	return resp, nil
}

func (e *ExpressRouteCircuitsServerTransport) dispatchBeginListRoutesTable(req *http.Request) (*http.Response, error) {
	if e.srv.BeginListRoutesTable == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginListRoutesTable not implemented")}
	}
	if e.beginListRoutesTable == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/routeTables/(?P<devicePath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		circuitNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
		if err != nil {
			return nil, err
		}
		peeringNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
		if err != nil {
			return nil, err
		}
		devicePathUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("devicePath")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginListRoutesTable(req.Context(), resourceGroupNameUnescaped, circuitNameUnescaped, peeringNameUnescaped, devicePathUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		e.beginListRoutesTable = &respr
	}

	resp, err := server.PollerResponderNext(e.beginListRoutesTable, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(e.beginListRoutesTable) {
		e.beginListRoutesTable = nil
	}

	return resp, nil
}

func (e *ExpressRouteCircuitsServerTransport) dispatchBeginListRoutesTableSummary(req *http.Request) (*http.Response, error) {
	if e.srv.BeginListRoutesTableSummary == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginListRoutesTableSummary not implemented")}
	}
	if e.beginListRoutesTableSummary == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/routeTablesSummary/(?P<devicePath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		circuitNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
		if err != nil {
			return nil, err
		}
		peeringNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
		if err != nil {
			return nil, err
		}
		devicePathUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("devicePath")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginListRoutesTableSummary(req.Context(), resourceGroupNameUnescaped, circuitNameUnescaped, peeringNameUnescaped, devicePathUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		e.beginListRoutesTableSummary = &respr
	}

	resp, err := server.PollerResponderNext(e.beginListRoutesTableSummary, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(e.beginListRoutesTableSummary) {
		e.beginListRoutesTableSummary = nil
	}

	return resp, nil
}

func (e *ExpressRouteCircuitsServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if e.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	circuitNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.UpdateTags(req.Context(), resourceGroupNameUnescaped, circuitNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRouteCircuit, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
