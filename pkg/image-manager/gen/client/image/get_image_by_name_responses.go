///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/image-manager/gen/models"
)

// GetImageByNameReader is a Reader for the GetImageByName structure.
type GetImageByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImageByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetImageByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetImageByNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetImageByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetImageByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetImageByNameOK creates a GetImageByNameOK with default headers values
func NewGetImageByNameOK() *GetImageByNameOK {
	return &GetImageByNameOK{}
}

/*GetImageByNameOK handles this case with default header values.

successful operation
*/
type GetImageByNameOK struct {
	Payload *models.Image
}

func (o *GetImageByNameOK) Error() string {
	return fmt.Sprintf("[GET /{imageName}][%d] getImageByNameOK  %+v", 200, o.Payload)
}

func (o *GetImageByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageByNameBadRequest creates a GetImageByNameBadRequest with default headers values
func NewGetImageByNameBadRequest() *GetImageByNameBadRequest {
	return &GetImageByNameBadRequest{}
}

/*GetImageByNameBadRequest handles this case with default header values.

Invalid ID supplied
*/
type GetImageByNameBadRequest struct {
	Payload *models.Error
}

func (o *GetImageByNameBadRequest) Error() string {
	return fmt.Sprintf("[GET /{imageName}][%d] getImageByNameBadRequest  %+v", 400, o.Payload)
}

func (o *GetImageByNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageByNameNotFound creates a GetImageByNameNotFound with default headers values
func NewGetImageByNameNotFound() *GetImageByNameNotFound {
	return &GetImageByNameNotFound{}
}

/*GetImageByNameNotFound handles this case with default header values.

Image not found
*/
type GetImageByNameNotFound struct {
	Payload *models.Error
}

func (o *GetImageByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /{imageName}][%d] getImageByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetImageByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageByNameDefault creates a GetImageByNameDefault with default headers values
func NewGetImageByNameDefault(code int) *GetImageByNameDefault {
	return &GetImageByNameDefault{
		_statusCode: code,
	}
}

/*GetImageByNameDefault handles this case with default header values.

Generic error response
*/
type GetImageByNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get image by name default response
func (o *GetImageByNameDefault) Code() int {
	return o._statusCode
}

func (o *GetImageByNameDefault) Error() string {
	return fmt.Sprintf("[GET /{imageName}][%d] getImageByName default  %+v", o._statusCode, o.Payload)
}

func (o *GetImageByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}