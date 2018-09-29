///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// DeleteEndpointReader is a Reader for the DeleteEndpoint structure.
type DeleteEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteEndpointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteEndpointBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteEndpointUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteEndpointForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteEndpointNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteEndpointDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteEndpointOK creates a DeleteEndpointOK with default headers values
func NewDeleteEndpointOK() *DeleteEndpointOK {
	return &DeleteEndpointOK{}
}

/*DeleteEndpointOK handles this case with default header values.

Successful operation
*/
type DeleteEndpointOK struct {
	Payload *v1.Endpoint
}

func (o *DeleteEndpointOK) Error() string {
	return fmt.Sprintf("[DELETE /{endpoint}][%d] deleteEndpointOK  %+v", 200, o.Payload)
}

func (o *DeleteEndpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Endpoint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEndpointBadRequest creates a DeleteEndpointBadRequest with default headers values
func NewDeleteEndpointBadRequest() *DeleteEndpointBadRequest {
	return &DeleteEndpointBadRequest{}
}

/*DeleteEndpointBadRequest handles this case with default header values.

Invalid Name supplied
*/
type DeleteEndpointBadRequest struct {
	Payload *v1.Error
}

func (o *DeleteEndpointBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /{endpoint}][%d] deleteEndpointBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteEndpointBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEndpointUnauthorized creates a DeleteEndpointUnauthorized with default headers values
func NewDeleteEndpointUnauthorized() *DeleteEndpointUnauthorized {
	return &DeleteEndpointUnauthorized{}
}

/*DeleteEndpointUnauthorized handles this case with default header values.

Unauthorized Request
*/
type DeleteEndpointUnauthorized struct {
	Payload *v1.Error
}

func (o *DeleteEndpointUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /{endpoint}][%d] deleteEndpointUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteEndpointUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEndpointForbidden creates a DeleteEndpointForbidden with default headers values
func NewDeleteEndpointForbidden() *DeleteEndpointForbidden {
	return &DeleteEndpointForbidden{}
}

/*DeleteEndpointForbidden handles this case with default header values.

access to this resource is forbidden
*/
type DeleteEndpointForbidden struct {
	Payload *v1.Error
}

func (o *DeleteEndpointForbidden) Error() string {
	return fmt.Sprintf("[DELETE /{endpoint}][%d] deleteEndpointForbidden  %+v", 403, o.Payload)
}

func (o *DeleteEndpointForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEndpointNotFound creates a DeleteEndpointNotFound with default headers values
func NewDeleteEndpointNotFound() *DeleteEndpointNotFound {
	return &DeleteEndpointNotFound{}
}

/*DeleteEndpointNotFound handles this case with default header values.

Endpoint not found
*/
type DeleteEndpointNotFound struct {
	Payload *v1.Error
}

func (o *DeleteEndpointNotFound) Error() string {
	return fmt.Sprintf("[DELETE /{endpoint}][%d] deleteEndpointNotFound  %+v", 404, o.Payload)
}

func (o *DeleteEndpointNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEndpointDefault creates a DeleteEndpointDefault with default headers values
func NewDeleteEndpointDefault(code int) *DeleteEndpointDefault {
	return &DeleteEndpointDefault{
		_statusCode: code,
	}
}

/*DeleteEndpointDefault handles this case with default header values.

Unknown error
*/
type DeleteEndpointDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the delete endpoint default response
func (o *DeleteEndpointDefault) Code() int {
	return o._statusCode
}

func (o *DeleteEndpointDefault) Error() string {
	return fmt.Sprintf("[DELETE /{endpoint}][%d] deleteEndpoint default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteEndpointDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}