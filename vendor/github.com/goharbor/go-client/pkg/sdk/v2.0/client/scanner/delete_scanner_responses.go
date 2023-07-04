// Code generated by go-swagger; DO NOT EDIT.

package scanner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// DeleteScannerReader is a Reader for the DeleteScanner structure.
type DeleteScannerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteScannerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteScannerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteScannerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteScannerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteScannerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteScannerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteScannerOK creates a DeleteScannerOK with default headers values
func NewDeleteScannerOK() *DeleteScannerOK {
	return &DeleteScannerOK{}
}

/*
DeleteScannerOK describes a response with status code 200, with default header values.

Deleted successfully and return the deleted registration
*/
type DeleteScannerOK struct {
	Payload *models.ScannerRegistration
}

// IsSuccess returns true when this delete scanner o k response has a 2xx status code
func (o *DeleteScannerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete scanner o k response has a 3xx status code
func (o *DeleteScannerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete scanner o k response has a 4xx status code
func (o *DeleteScannerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete scanner o k response has a 5xx status code
func (o *DeleteScannerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete scanner o k response a status code equal to that given
func (o *DeleteScannerOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteScannerOK) Error() string {
	return fmt.Sprintf("[DELETE /scanners/{registration_id}][%d] deleteScannerOK  %+v", 200, o.Payload)
}

func (o *DeleteScannerOK) String() string {
	return fmt.Sprintf("[DELETE /scanners/{registration_id}][%d] deleteScannerOK  %+v", 200, o.Payload)
}

func (o *DeleteScannerOK) GetPayload() *models.ScannerRegistration {
	return o.Payload
}

func (o *DeleteScannerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScannerRegistration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScannerUnauthorized creates a DeleteScannerUnauthorized with default headers values
func NewDeleteScannerUnauthorized() *DeleteScannerUnauthorized {
	return &DeleteScannerUnauthorized{}
}

/*
DeleteScannerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteScannerUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete scanner unauthorized response has a 2xx status code
func (o *DeleteScannerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete scanner unauthorized response has a 3xx status code
func (o *DeleteScannerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete scanner unauthorized response has a 4xx status code
func (o *DeleteScannerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete scanner unauthorized response has a 5xx status code
func (o *DeleteScannerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete scanner unauthorized response a status code equal to that given
func (o *DeleteScannerUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteScannerUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /scanners/{registration_id}][%d] deleteScannerUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteScannerUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /scanners/{registration_id}][%d] deleteScannerUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteScannerUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteScannerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScannerForbidden creates a DeleteScannerForbidden with default headers values
func NewDeleteScannerForbidden() *DeleteScannerForbidden {
	return &DeleteScannerForbidden{}
}

/*
DeleteScannerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteScannerForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete scanner forbidden response has a 2xx status code
func (o *DeleteScannerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete scanner forbidden response has a 3xx status code
func (o *DeleteScannerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete scanner forbidden response has a 4xx status code
func (o *DeleteScannerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete scanner forbidden response has a 5xx status code
func (o *DeleteScannerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete scanner forbidden response a status code equal to that given
func (o *DeleteScannerForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteScannerForbidden) Error() string {
	return fmt.Sprintf("[DELETE /scanners/{registration_id}][%d] deleteScannerForbidden  %+v", 403, o.Payload)
}

func (o *DeleteScannerForbidden) String() string {
	return fmt.Sprintf("[DELETE /scanners/{registration_id}][%d] deleteScannerForbidden  %+v", 403, o.Payload)
}

func (o *DeleteScannerForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteScannerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScannerNotFound creates a DeleteScannerNotFound with default headers values
func NewDeleteScannerNotFound() *DeleteScannerNotFound {
	return &DeleteScannerNotFound{}
}

/*
DeleteScannerNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteScannerNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete scanner not found response has a 2xx status code
func (o *DeleteScannerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete scanner not found response has a 3xx status code
func (o *DeleteScannerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete scanner not found response has a 4xx status code
func (o *DeleteScannerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete scanner not found response has a 5xx status code
func (o *DeleteScannerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete scanner not found response a status code equal to that given
func (o *DeleteScannerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteScannerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /scanners/{registration_id}][%d] deleteScannerNotFound  %+v", 404, o.Payload)
}

func (o *DeleteScannerNotFound) String() string {
	return fmt.Sprintf("[DELETE /scanners/{registration_id}][%d] deleteScannerNotFound  %+v", 404, o.Payload)
}

func (o *DeleteScannerNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteScannerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScannerInternalServerError creates a DeleteScannerInternalServerError with default headers values
func NewDeleteScannerInternalServerError() *DeleteScannerInternalServerError {
	return &DeleteScannerInternalServerError{}
}

/*
DeleteScannerInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteScannerInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete scanner internal server error response has a 2xx status code
func (o *DeleteScannerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete scanner internal server error response has a 3xx status code
func (o *DeleteScannerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete scanner internal server error response has a 4xx status code
func (o *DeleteScannerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete scanner internal server error response has a 5xx status code
func (o *DeleteScannerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete scanner internal server error response a status code equal to that given
func (o *DeleteScannerInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteScannerInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /scanners/{registration_id}][%d] deleteScannerInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteScannerInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /scanners/{registration_id}][%d] deleteScannerInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteScannerInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteScannerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}