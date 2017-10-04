///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/function-manager/gen/models"
)

// UpdateFunctionReader is a Reader for the UpdateFunction structure.
type UpdateFunctionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFunctionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateFunctionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateFunctionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateFunctionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateFunctionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateFunctionOK creates a UpdateFunctionOK with default headers values
func NewUpdateFunctionOK() *UpdateFunctionOK {
	return &UpdateFunctionOK{}
}

/*UpdateFunctionOK handles this case with default header values.

Successful update
*/
type UpdateFunctionOK struct {
	Payload *models.Function
}

func (o *UpdateFunctionOK) Error() string {
	return fmt.Sprintf("[PUT /{functionName}][%d] updateFunctionOK  %+v", 200, o.Payload)
}

func (o *UpdateFunctionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Function)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFunctionBadRequest creates a UpdateFunctionBadRequest with default headers values
func NewUpdateFunctionBadRequest() *UpdateFunctionBadRequest {
	return &UpdateFunctionBadRequest{}
}

/*UpdateFunctionBadRequest handles this case with default header values.

Invalid input
*/
type UpdateFunctionBadRequest struct {
	Payload *models.Error
}

func (o *UpdateFunctionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /{functionName}][%d] updateFunctionBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateFunctionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFunctionNotFound creates a UpdateFunctionNotFound with default headers values
func NewUpdateFunctionNotFound() *UpdateFunctionNotFound {
	return &UpdateFunctionNotFound{}
}

/*UpdateFunctionNotFound handles this case with default header values.

Function not found
*/
type UpdateFunctionNotFound struct {
	Payload *models.Error
}

func (o *UpdateFunctionNotFound) Error() string {
	return fmt.Sprintf("[PUT /{functionName}][%d] updateFunctionNotFound  %+v", 404, o.Payload)
}

func (o *UpdateFunctionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFunctionInternalServerError creates a UpdateFunctionInternalServerError with default headers values
func NewUpdateFunctionInternalServerError() *UpdateFunctionInternalServerError {
	return &UpdateFunctionInternalServerError{}
}

/*UpdateFunctionInternalServerError handles this case with default header values.

Internal error
*/
type UpdateFunctionInternalServerError struct {
	Payload *models.Error
}

func (o *UpdateFunctionInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /{functionName}][%d] updateFunctionInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateFunctionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
