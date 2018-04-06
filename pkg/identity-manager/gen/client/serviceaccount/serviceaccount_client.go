///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package serviceaccount

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new serviceaccount API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for serviceaccount API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddServiceAccount adds a new service account
*/
func (a *Client) AddServiceAccount(params *AddServiceAccountParams, authInfo runtime.ClientAuthInfoWriter) (*AddServiceAccountCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddServiceAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addServiceAccount",
		Method:             "POST",
		PathPattern:        "/v1/iam/serviceaccount",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddServiceAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddServiceAccountCreated), nil

}

/*
DeleteServiceAccount deletes an service account
*/
func (a *Client) DeleteServiceAccount(params *DeleteServiceAccountParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServiceAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServiceAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteServiceAccount",
		Method:             "DELETE",
		PathPattern:        "/v1/iam/serviceaccount/{serviceAccountName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteServiceAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteServiceAccountOK), nil

}

/*
GetServiceAccount finds service account by name

get a Service Account by name
*/
func (a *Client) GetServiceAccount(params *GetServiceAccountParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getServiceAccount",
		Method:             "GET",
		PathPattern:        "/v1/iam/serviceaccount/{serviceAccountName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServiceAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetServiceAccountOK), nil

}

/*
GetServiceAccounts lists all existing service accounts
*/
func (a *Client) GetServiceAccounts(params *GetServiceAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getServiceAccounts",
		Method:             "GET",
		PathPattern:        "/v1/iam/serviceaccount",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServiceAccountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetServiceAccountsOK), nil

}

/*
UpdateServiceAccount updates a service account
*/
func (a *Client) UpdateServiceAccount(params *UpdateServiceAccountParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServiceAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateServiceAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateServiceAccount",
		Method:             "PUT",
		PathPattern:        "/v1/iam/serviceaccount/{serviceAccountName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateServiceAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateServiceAccountOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
