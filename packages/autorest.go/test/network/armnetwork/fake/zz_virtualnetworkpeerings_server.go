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

// VirtualNetworkPeeringsServer is a fake server for instances of the armnetwork.VirtualNetworkPeeringsClient type.
type VirtualNetworkPeeringsServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualNetworkPeeringsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, virtualNetworkPeeringParameters armnetwork.VirtualNetworkPeering, options *armnetwork.VirtualNetworkPeeringsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworkPeeringsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualNetworkPeeringsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, options *armnetwork.VirtualNetworkPeeringsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.VirtualNetworkPeeringsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualNetworkPeeringsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, options *armnetwork.VirtualNetworkPeeringsClientGetOptions) (resp azfake.Responder[armnetwork.VirtualNetworkPeeringsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualNetworkPeeringsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, virtualNetworkName string, options *armnetwork.VirtualNetworkPeeringsClientListOptions) (resp azfake.PagerResponder[armnetwork.VirtualNetworkPeeringsClientListResponse])
}

// NewVirtualNetworkPeeringsServerTransport creates a new instance of VirtualNetworkPeeringsServerTransport with the provided implementation.
// The returned VirtualNetworkPeeringsServerTransport instance is connected to an instance of armnetwork.VirtualNetworkPeeringsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualNetworkPeeringsServerTransport(srv *VirtualNetworkPeeringsServer) *VirtualNetworkPeeringsServerTransport {
	return &VirtualNetworkPeeringsServerTransport{srv: srv}
}

// VirtualNetworkPeeringsServerTransport connects instances of armnetwork.VirtualNetworkPeeringsClient to instances of VirtualNetworkPeeringsServer.
// Don't use this type directly, use NewVirtualNetworkPeeringsServerTransport instead.
type VirtualNetworkPeeringsServerTransport struct {
	srv                 *VirtualNetworkPeeringsServer
	beginCreateOrUpdate *azfake.PollerResponder[armnetwork.VirtualNetworkPeeringsClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armnetwork.VirtualNetworkPeeringsClientDeleteResponse]
	newListPager        *azfake.PagerResponder[armnetwork.VirtualNetworkPeeringsClientListResponse]
}

// Do implements the policy.Transporter interface for VirtualNetworkPeeringsServerTransport.
func (v *VirtualNetworkPeeringsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualNetworkPeeringsClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualNetworkPeeringsClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualNetworkPeeringsClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualNetworkPeeringsClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualNetworkPeeringsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	if v.beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkPeerings/(?P<virtualNetworkPeeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VirtualNetworkPeering](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkPeeringNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkPeeringName")])
		if err != nil {
			return nil, err
		}
		syncRemoteAddressSpaceUnescaped, err := url.QueryUnescape(qp.Get("syncRemoteAddressSpace"))
		if err != nil {
			return nil, err
		}
		syncRemoteAddressSpaceParam := getOptional(armnetwork.SyncRemoteAddressSpace(syncRemoteAddressSpaceUnescaped))
		var options *armnetwork.VirtualNetworkPeeringsClientBeginCreateOrUpdateOptions
		if syncRemoteAddressSpaceParam != nil {
			options = &armnetwork.VirtualNetworkPeeringsClientBeginCreateOrUpdateOptions{
				SyncRemoteAddressSpace: syncRemoteAddressSpaceParam,
			}
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, virtualNetworkNameUnescaped, virtualNetworkPeeringNameUnescaped, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(v.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginCreateOrUpdate) {
		v.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (v *VirtualNetworkPeeringsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if v.beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkPeerings/(?P<virtualNetworkPeeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkPeeringNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkPeeringName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, virtualNetworkNameUnescaped, virtualNetworkPeeringNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		v.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(v.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(v.beginDelete) {
		v.beginDelete = nil
	}

	return resp, nil
}

func (v *VirtualNetworkPeeringsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkPeerings/(?P<virtualNetworkPeeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
	if err != nil {
		return nil, err
	}
	virtualNetworkPeeringNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkPeeringName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameUnescaped, virtualNetworkNameUnescaped, virtualNetworkPeeringNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualNetworkPeering, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualNetworkPeeringsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if v.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkPeerings`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListPager(resourceGroupNameUnescaped, virtualNetworkNameUnescaped, nil)
		v.newListPager = &resp
		server.PagerResponderInjectNextLinks(v.newListPager, req, func(page *armnetwork.VirtualNetworkPeeringsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(v.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(v.newListPager) {
		v.newListPager = nil
	}
	return resp, nil
}