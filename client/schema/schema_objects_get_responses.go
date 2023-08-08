// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// SchemaObjectsGetReader is a Reader for the SchemaObjectsGet structure.
type SchemaObjectsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchemaObjectsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchemaObjectsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSchemaObjectsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSchemaObjectsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSchemaObjectsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSchemaObjectsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSchemaObjectsGetOK creates a SchemaObjectsGetOK with default headers values
func NewSchemaObjectsGetOK() *SchemaObjectsGetOK {
	return &SchemaObjectsGetOK{}
}

/*
SchemaObjectsGetOK describes a response with status code 200, with default header values.

Found the Class, returned as body
*/
type SchemaObjectsGetOK struct {
	Payload *models.Class
}

// IsSuccess returns true when this schema objects get o k response has a 2xx status code
func (o *SchemaObjectsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schema objects get o k response has a 3xx status code
func (o *SchemaObjectsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema objects get o k response has a 4xx status code
func (o *SchemaObjectsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this schema objects get o k response has a 5xx status code
func (o *SchemaObjectsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this schema objects get o k response a status code equal to that given
func (o *SchemaObjectsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the schema objects get o k response
func (o *SchemaObjectsGetOK) Code() int {
	return 200
}

func (o *SchemaObjectsGetOK) Error() string {
	return fmt.Sprintf("[GET /schema/{className}][%d] schemaObjectsGetOK  %+v", 200, o.Payload)
}

func (o *SchemaObjectsGetOK) String() string {
	return fmt.Sprintf("[GET /schema/{className}][%d] schemaObjectsGetOK  %+v", 200, o.Payload)
}

func (o *SchemaObjectsGetOK) GetPayload() *models.Class {
	return o.Payload
}

func (o *SchemaObjectsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Class)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsGetUnauthorized creates a SchemaObjectsGetUnauthorized with default headers values
func NewSchemaObjectsGetUnauthorized() *SchemaObjectsGetUnauthorized {
	return &SchemaObjectsGetUnauthorized{}
}

/*
SchemaObjectsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type SchemaObjectsGetUnauthorized struct {
}

// IsSuccess returns true when this schema objects get unauthorized response has a 2xx status code
func (o *SchemaObjectsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema objects get unauthorized response has a 3xx status code
func (o *SchemaObjectsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema objects get unauthorized response has a 4xx status code
func (o *SchemaObjectsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this schema objects get unauthorized response has a 5xx status code
func (o *SchemaObjectsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this schema objects get unauthorized response a status code equal to that given
func (o *SchemaObjectsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the schema objects get unauthorized response
func (o *SchemaObjectsGetUnauthorized) Code() int {
	return 401
}

func (o *SchemaObjectsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /schema/{className}][%d] schemaObjectsGetUnauthorized ", 401)
}

func (o *SchemaObjectsGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /schema/{className}][%d] schemaObjectsGetUnauthorized ", 401)
}

func (o *SchemaObjectsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaObjectsGetForbidden creates a SchemaObjectsGetForbidden with default headers values
func NewSchemaObjectsGetForbidden() *SchemaObjectsGetForbidden {
	return &SchemaObjectsGetForbidden{}
}

/*
SchemaObjectsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SchemaObjectsGetForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this schema objects get forbidden response has a 2xx status code
func (o *SchemaObjectsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema objects get forbidden response has a 3xx status code
func (o *SchemaObjectsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema objects get forbidden response has a 4xx status code
func (o *SchemaObjectsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this schema objects get forbidden response has a 5xx status code
func (o *SchemaObjectsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this schema objects get forbidden response a status code equal to that given
func (o *SchemaObjectsGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the schema objects get forbidden response
func (o *SchemaObjectsGetForbidden) Code() int {
	return 403
}

func (o *SchemaObjectsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /schema/{className}][%d] schemaObjectsGetForbidden  %+v", 403, o.Payload)
}

func (o *SchemaObjectsGetForbidden) String() string {
	return fmt.Sprintf("[GET /schema/{className}][%d] schemaObjectsGetForbidden  %+v", 403, o.Payload)
}

func (o *SchemaObjectsGetForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsGetNotFound creates a SchemaObjectsGetNotFound with default headers values
func NewSchemaObjectsGetNotFound() *SchemaObjectsGetNotFound {
	return &SchemaObjectsGetNotFound{}
}

/*
SchemaObjectsGetNotFound describes a response with status code 404, with default header values.

This class does not exist
*/
type SchemaObjectsGetNotFound struct {
}

// IsSuccess returns true when this schema objects get not found response has a 2xx status code
func (o *SchemaObjectsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema objects get not found response has a 3xx status code
func (o *SchemaObjectsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema objects get not found response has a 4xx status code
func (o *SchemaObjectsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this schema objects get not found response has a 5xx status code
func (o *SchemaObjectsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this schema objects get not found response a status code equal to that given
func (o *SchemaObjectsGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the schema objects get not found response
func (o *SchemaObjectsGetNotFound) Code() int {
	return 404
}

func (o *SchemaObjectsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /schema/{className}][%d] schemaObjectsGetNotFound ", 404)
}

func (o *SchemaObjectsGetNotFound) String() string {
	return fmt.Sprintf("[GET /schema/{className}][%d] schemaObjectsGetNotFound ", 404)
}

func (o *SchemaObjectsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaObjectsGetInternalServerError creates a SchemaObjectsGetInternalServerError with default headers values
func NewSchemaObjectsGetInternalServerError() *SchemaObjectsGetInternalServerError {
	return &SchemaObjectsGetInternalServerError{}
}

/*
SchemaObjectsGetInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type SchemaObjectsGetInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this schema objects get internal server error response has a 2xx status code
func (o *SchemaObjectsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema objects get internal server error response has a 3xx status code
func (o *SchemaObjectsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema objects get internal server error response has a 4xx status code
func (o *SchemaObjectsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this schema objects get internal server error response has a 5xx status code
func (o *SchemaObjectsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this schema objects get internal server error response a status code equal to that given
func (o *SchemaObjectsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the schema objects get internal server error response
func (o *SchemaObjectsGetInternalServerError) Code() int {
	return 500
}

func (o *SchemaObjectsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /schema/{className}][%d] schemaObjectsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaObjectsGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /schema/{className}][%d] schemaObjectsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaObjectsGetInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
