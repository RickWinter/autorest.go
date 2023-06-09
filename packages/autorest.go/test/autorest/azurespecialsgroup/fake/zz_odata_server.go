//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/azurespecialsgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
)

// ODataServer is a fake server for instances of the azurespecialsgroup.ODataClient type.
type ODataServer struct {
	// GetWithFilter is the fake for method ODataClient.GetWithFilter
	// HTTP status codes to indicate success: http.StatusOK
	GetWithFilter func(ctx context.Context, options *azurespecialsgroup.ODataClientGetWithFilterOptions) (resp azfake.Responder[azurespecialsgroup.ODataClientGetWithFilterResponse], errResp azfake.ErrorResponder)
}

// NewODataServerTransport creates a new instance of ODataServerTransport with the provided implementation.
// The returned ODataServerTransport instance is connected to an instance of azurespecialsgroup.ODataClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewODataServerTransport(srv *ODataServer) *ODataServerTransport {
	return &ODataServerTransport{srv: srv}
}

// ODataServerTransport connects instances of azurespecialsgroup.ODataClient to instances of ODataServer.
// Don't use this type directly, use NewODataServerTransport instead.
type ODataServerTransport struct {
	srv *ODataServer
}

// Do implements the policy.Transporter interface for ODataServerTransport.
func (o *ODataServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ODataClient.GetWithFilter":
		resp, err = o.dispatchGetWithFilter(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *ODataServerTransport) dispatchGetWithFilter(req *http.Request) (*http.Response, error) {
	if o.srv.GetWithFilter == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetWithFilter not implemented")}
	}
	qp := req.URL.Query()
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
	if err != nil {
		return nil, err
	}
	topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
	if err != nil {
		return nil, err
	}
	orderbyParam := getOptional(orderbyUnescaped)
	var options *azurespecialsgroup.ODataClientGetWithFilterOptions
	if filterParam != nil || topParam != nil || orderbyParam != nil {
		options = &azurespecialsgroup.ODataClientGetWithFilterOptions{
			Filter:  filterParam,
			Top:     topParam,
			Orderby: orderbyParam,
		}
	}
	respr, errRespr := o.srv.GetWithFilter(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}