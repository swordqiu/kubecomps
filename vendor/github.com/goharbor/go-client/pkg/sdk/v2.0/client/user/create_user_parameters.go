// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// NewCreateUserParams creates a new CreateUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUserParams() *CreateUserParams {
	return &CreateUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserParamsWithTimeout creates a new CreateUserParams object
// with the ability to set a timeout on a request.
func NewCreateUserParamsWithTimeout(timeout time.Duration) *CreateUserParams {
	return &CreateUserParams{
		timeout: timeout,
	}
}

// NewCreateUserParamsWithContext creates a new CreateUserParams object
// with the ability to set a context for a request.
func NewCreateUserParamsWithContext(ctx context.Context) *CreateUserParams {
	return &CreateUserParams{
		Context: ctx,
	}
}

// NewCreateUserParamsWithHTTPClient creates a new CreateUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUserParamsWithHTTPClient(client *http.Client) *CreateUserParams {
	return &CreateUserParams{
		HTTPClient: client,
	}
}

/*
CreateUserParams contains all the parameters to send to the API endpoint

	for the create user operation.

	Typically these are written to a http.Request.
*/
type CreateUserParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* UserReq.

	   The new user
	*/
	UserReq *models.UserCreationReq

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserParams) WithDefaults() *CreateUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create user params
func (o *CreateUserParams) WithTimeout(timeout time.Duration) *CreateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create user params
func (o *CreateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create user params
func (o *CreateUserParams) WithContext(ctx context.Context) *CreateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create user params
func (o *CreateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create user params
func (o *CreateUserParams) WithHTTPClient(client *http.Client) *CreateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create user params
func (o *CreateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create user params
func (o *CreateUserParams) WithXRequestID(xRequestID *string) *CreateUserParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create user params
func (o *CreateUserParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithUserReq adds the userReq to the create user params
func (o *CreateUserParams) WithUserReq(userReq *models.UserCreationReq) *CreateUserParams {
	o.SetUserReq(userReq)
	return o
}

// SetUserReq adds the userReq to the create user params
func (o *CreateUserParams) SetUserReq(userReq *models.UserCreationReq) {
	o.UserReq = userReq
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}
	if o.UserReq != nil {
		if err := r.SetBodyParam(o.UserReq); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
