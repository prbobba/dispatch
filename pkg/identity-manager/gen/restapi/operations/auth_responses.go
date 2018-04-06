///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware/dispatch/pkg/identity-manager/gen/models"
)

// AuthAcceptedCode is the HTTP code returned for type AuthAccepted
const AuthAcceptedCode int = 202

/*AuthAccepted default response if authorized

swagger:response authAccepted
*/
type AuthAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.Message `json:"body,omitempty"`
}

// NewAuthAccepted creates AuthAccepted with default headers values
func NewAuthAccepted() *AuthAccepted {

	return &AuthAccepted{}
}

// WithPayload adds the payload to the auth accepted response
func (o *AuthAccepted) WithPayload(payload *models.Message) *AuthAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the auth accepted response
func (o *AuthAccepted) SetPayload(payload *models.Message) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AuthUnauthorizedCode is the HTTP code returned for type AuthUnauthorized
const AuthUnauthorizedCode int = 401

/*AuthUnauthorized Unauthorized

swagger:response authUnauthorized
*/
type AuthUnauthorized struct {
}

// NewAuthUnauthorized creates AuthUnauthorized with default headers values
func NewAuthUnauthorized() *AuthUnauthorized {

	return &AuthUnauthorized{}
}

// WriteResponse to the client
func (o *AuthUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// AuthForbiddenCode is the HTTP code returned for type AuthForbidden
const AuthForbiddenCode int = 403

/*AuthForbidden Forbidden

swagger:response authForbidden
*/
type AuthForbidden struct {
}

// NewAuthForbidden creates AuthForbidden with default headers values
func NewAuthForbidden() *AuthForbidden {

	return &AuthForbidden{}
}

// WriteResponse to the client
func (o *AuthForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

/*AuthDefault error

swagger:response authDefault
*/
type AuthDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAuthDefault creates AuthDefault with default headers values
func NewAuthDefault(code int) *AuthDefault {
	if code <= 0 {
		code = 500
	}

	return &AuthDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the auth default response
func (o *AuthDefault) WithStatusCode(code int) *AuthDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the auth default response
func (o *AuthDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the auth default response
func (o *AuthDefault) WithPayload(payload *models.Error) *AuthDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the auth default response
func (o *AuthDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
