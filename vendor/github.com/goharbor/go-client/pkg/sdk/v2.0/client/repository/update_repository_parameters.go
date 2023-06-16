// Code generated by go-swagger; DO NOT EDIT.

package repository

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

// NewUpdateRepositoryParams creates a new UpdateRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRepositoryParams() *UpdateRepositoryParams {
	return &UpdateRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRepositoryParamsWithTimeout creates a new UpdateRepositoryParams object
// with the ability to set a timeout on a request.
func NewUpdateRepositoryParamsWithTimeout(timeout time.Duration) *UpdateRepositoryParams {
	return &UpdateRepositoryParams{
		timeout: timeout,
	}
}

// NewUpdateRepositoryParamsWithContext creates a new UpdateRepositoryParams object
// with the ability to set a context for a request.
func NewUpdateRepositoryParamsWithContext(ctx context.Context) *UpdateRepositoryParams {
	return &UpdateRepositoryParams{
		Context: ctx,
	}
}

// NewUpdateRepositoryParamsWithHTTPClient creates a new UpdateRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRepositoryParamsWithHTTPClient(client *http.Client) *UpdateRepositoryParams {
	return &UpdateRepositoryParams{
		HTTPClient: client,
	}
}

/*
UpdateRepositoryParams contains all the parameters to send to the API endpoint

	for the update repository operation.

	Typically these are written to a http.Request.
*/
type UpdateRepositoryParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* ProjectName.

	   The name of the project
	*/
	ProjectName string

	/* Repository.

	   The JSON object of repository.
	*/
	Repository *models.Repository

	/* RepositoryName.

	   The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -> a%252Fb
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepositoryParams) WithDefaults() *UpdateRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRepositoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update repository params
func (o *UpdateRepositoryParams) WithTimeout(timeout time.Duration) *UpdateRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update repository params
func (o *UpdateRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update repository params
func (o *UpdateRepositoryParams) WithContext(ctx context.Context) *UpdateRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update repository params
func (o *UpdateRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update repository params
func (o *UpdateRepositoryParams) WithHTTPClient(client *http.Client) *UpdateRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update repository params
func (o *UpdateRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update repository params
func (o *UpdateRepositoryParams) WithXRequestID(xRequestID *string) *UpdateRepositoryParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update repository params
func (o *UpdateRepositoryParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectName adds the projectName to the update repository params
func (o *UpdateRepositoryParams) WithProjectName(projectName string) *UpdateRepositoryParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the update repository params
func (o *UpdateRepositoryParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithRepository adds the repository to the update repository params
func (o *UpdateRepositoryParams) WithRepository(repository *models.Repository) *UpdateRepositoryParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the update repository params
func (o *UpdateRepositoryParams) SetRepository(repository *models.Repository) {
	o.Repository = repository
}

// WithRepositoryName adds the repositoryName to the update repository params
func (o *UpdateRepositoryParams) WithRepositoryName(repositoryName string) *UpdateRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the update repository params
func (o *UpdateRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param project_name
	if err := r.SetPathParam("project_name", o.ProjectName); err != nil {
		return err
	}
	if o.Repository != nil {
		if err := r.SetBodyParam(o.Repository); err != nil {
			return err
		}
	}

	// path param repository_name
	if err := r.SetPathParam("repository_name", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
