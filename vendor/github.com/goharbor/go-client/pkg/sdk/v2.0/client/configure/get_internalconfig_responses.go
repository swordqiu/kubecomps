// Code generated by go-swagger; DO NOT EDIT.

package configure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// GetInternalconfigReader is a Reader for the GetInternalconfig structure.
type GetInternalconfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInternalconfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInternalconfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetInternalconfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInternalconfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInternalconfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInternalconfigOK creates a GetInternalconfigOK with default headers values
func NewGetInternalconfigOK() *GetInternalconfigOK {
	return &GetInternalconfigOK{}
}

/*
GetInternalconfigOK describes a response with status code 200, with default header values.

Get system configurations successfully. The response body is a map.
*/
type GetInternalconfigOK struct {
	Payload models.InternalConfigurationsResponse
}

// IsSuccess returns true when this get internalconfig o k response has a 2xx status code
func (o *GetInternalconfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get internalconfig o k response has a 3xx status code
func (o *GetInternalconfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get internalconfig o k response has a 4xx status code
func (o *GetInternalconfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get internalconfig o k response has a 5xx status code
func (o *GetInternalconfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get internalconfig o k response a status code equal to that given
func (o *GetInternalconfigOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInternalconfigOK) Error() string {
	return fmt.Sprintf("[GET /internalconfig][%d] getInternalconfigOK  %+v", 200, o.Payload)
}

func (o *GetInternalconfigOK) String() string {
	return fmt.Sprintf("[GET /internalconfig][%d] getInternalconfigOK  %+v", 200, o.Payload)
}

func (o *GetInternalconfigOK) GetPayload() models.InternalConfigurationsResponse {
	return o.Payload
}

func (o *GetInternalconfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternalconfigUnauthorized creates a GetInternalconfigUnauthorized with default headers values
func NewGetInternalconfigUnauthorized() *GetInternalconfigUnauthorized {
	return &GetInternalconfigUnauthorized{}
}

/*
GetInternalconfigUnauthorized describes a response with status code 401, with default header values.

User need to log in first.
*/
type GetInternalconfigUnauthorized struct {
}

// IsSuccess returns true when this get internalconfig unauthorized response has a 2xx status code
func (o *GetInternalconfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get internalconfig unauthorized response has a 3xx status code
func (o *GetInternalconfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get internalconfig unauthorized response has a 4xx status code
func (o *GetInternalconfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get internalconfig unauthorized response has a 5xx status code
func (o *GetInternalconfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get internalconfig unauthorized response a status code equal to that given
func (o *GetInternalconfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetInternalconfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /internalconfig][%d] getInternalconfigUnauthorized ", 401)
}

func (o *GetInternalconfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /internalconfig][%d] getInternalconfigUnauthorized ", 401)
}

func (o *GetInternalconfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInternalconfigForbidden creates a GetInternalconfigForbidden with default headers values
func NewGetInternalconfigForbidden() *GetInternalconfigForbidden {
	return &GetInternalconfigForbidden{}
}

/*
GetInternalconfigForbidden describes a response with status code 403, with default header values.

User does not have permission of admin role.
*/
type GetInternalconfigForbidden struct {
}

// IsSuccess returns true when this get internalconfig forbidden response has a 2xx status code
func (o *GetInternalconfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get internalconfig forbidden response has a 3xx status code
func (o *GetInternalconfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get internalconfig forbidden response has a 4xx status code
func (o *GetInternalconfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get internalconfig forbidden response has a 5xx status code
func (o *GetInternalconfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get internalconfig forbidden response a status code equal to that given
func (o *GetInternalconfigForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetInternalconfigForbidden) Error() string {
	return fmt.Sprintf("[GET /internalconfig][%d] getInternalconfigForbidden ", 403)
}

func (o *GetInternalconfigForbidden) String() string {
	return fmt.Sprintf("[GET /internalconfig][%d] getInternalconfigForbidden ", 403)
}

func (o *GetInternalconfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInternalconfigInternalServerError creates a GetInternalconfigInternalServerError with default headers values
func NewGetInternalconfigInternalServerError() *GetInternalconfigInternalServerError {
	return &GetInternalconfigInternalServerError{}
}

/*
GetInternalconfigInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type GetInternalconfigInternalServerError struct {
}

// IsSuccess returns true when this get internalconfig internal server error response has a 2xx status code
func (o *GetInternalconfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get internalconfig internal server error response has a 3xx status code
func (o *GetInternalconfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get internalconfig internal server error response has a 4xx status code
func (o *GetInternalconfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get internalconfig internal server error response has a 5xx status code
func (o *GetInternalconfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get internalconfig internal server error response a status code equal to that given
func (o *GetInternalconfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetInternalconfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /internalconfig][%d] getInternalconfigInternalServerError ", 500)
}

func (o *GetInternalconfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /internalconfig][%d] getInternalconfigInternalServerError ", 500)
}

func (o *GetInternalconfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
