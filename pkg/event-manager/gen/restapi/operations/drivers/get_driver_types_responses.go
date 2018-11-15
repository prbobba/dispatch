///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetDriverTypesOKCode is the HTTP code returned for type GetDriverTypesOK
const GetDriverTypesOKCode int = 200

/*GetDriverTypesOK Successful operation

swagger:response getDriverTypesOK
*/
type GetDriverTypesOK struct {

	/*
	  In: Body
	*/
	Payload []*v1.EventDriverType `json:"body,omitempty"`
}

// NewGetDriverTypesOK creates GetDriverTypesOK with default headers values
func NewGetDriverTypesOK() *GetDriverTypesOK {

	return &GetDriverTypesOK{}
}

// WithPayload adds the payload to the get driver types o k response
func (o *GetDriverTypesOK) WithPayload(payload []*v1.EventDriverType) *GetDriverTypesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get driver types o k response
func (o *GetDriverTypesOK) SetPayload(payload []*v1.EventDriverType) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriverTypesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*v1.EventDriverType, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetDriverTypesUnauthorizedCode is the HTTP code returned for type GetDriverTypesUnauthorized
const GetDriverTypesUnauthorizedCode int = 401

/*GetDriverTypesUnauthorized Unauthorized Request

swagger:response getDriverTypesUnauthorized
*/
type GetDriverTypesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetDriverTypesUnauthorized creates GetDriverTypesUnauthorized with default headers values
func NewGetDriverTypesUnauthorized() *GetDriverTypesUnauthorized {

	return &GetDriverTypesUnauthorized{}
}

// WithPayload adds the payload to the get driver types unauthorized response
func (o *GetDriverTypesUnauthorized) WithPayload(payload *v1.Error) *GetDriverTypesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get driver types unauthorized response
func (o *GetDriverTypesUnauthorized) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriverTypesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDriverTypesForbiddenCode is the HTTP code returned for type GetDriverTypesForbidden
const GetDriverTypesForbiddenCode int = 403

/*GetDriverTypesForbidden access to this resource is forbidden

swagger:response getDriverTypesForbidden
*/
type GetDriverTypesForbidden struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetDriverTypesForbidden creates GetDriverTypesForbidden with default headers values
func NewGetDriverTypesForbidden() *GetDriverTypesForbidden {

	return &GetDriverTypesForbidden{}
}

// WithPayload adds the payload to the get driver types forbidden response
func (o *GetDriverTypesForbidden) WithPayload(payload *v1.Error) *GetDriverTypesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get driver types forbidden response
func (o *GetDriverTypesForbidden) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriverTypesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetDriverTypesDefault Unknown error

swagger:response getDriverTypesDefault
*/
type GetDriverTypesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetDriverTypesDefault creates GetDriverTypesDefault with default headers values
func NewGetDriverTypesDefault(code int) *GetDriverTypesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDriverTypesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get driver types default response
func (o *GetDriverTypesDefault) WithStatusCode(code int) *GetDriverTypesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get driver types default response
func (o *GetDriverTypesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get driver types default response
func (o *GetDriverTypesDefault) WithPayload(payload *v1.Error) *GetDriverTypesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get driver types default response
func (o *GetDriverTypesDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriverTypesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}