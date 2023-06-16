// Code generated by go-swagger; DO NOT EDIT.

package chart_repository

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
)

// NewDeleteChartrepoRepoChartsNameParams creates a new DeleteChartrepoRepoChartsNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteChartrepoRepoChartsNameParams() *DeleteChartrepoRepoChartsNameParams {
	return &DeleteChartrepoRepoChartsNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteChartrepoRepoChartsNameParamsWithTimeout creates a new DeleteChartrepoRepoChartsNameParams object
// with the ability to set a timeout on a request.
func NewDeleteChartrepoRepoChartsNameParamsWithTimeout(timeout time.Duration) *DeleteChartrepoRepoChartsNameParams {
	return &DeleteChartrepoRepoChartsNameParams{
		timeout: timeout,
	}
}

// NewDeleteChartrepoRepoChartsNameParamsWithContext creates a new DeleteChartrepoRepoChartsNameParams object
// with the ability to set a context for a request.
func NewDeleteChartrepoRepoChartsNameParamsWithContext(ctx context.Context) *DeleteChartrepoRepoChartsNameParams {
	return &DeleteChartrepoRepoChartsNameParams{
		Context: ctx,
	}
}

// NewDeleteChartrepoRepoChartsNameParamsWithHTTPClient creates a new DeleteChartrepoRepoChartsNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteChartrepoRepoChartsNameParamsWithHTTPClient(client *http.Client) *DeleteChartrepoRepoChartsNameParams {
	return &DeleteChartrepoRepoChartsNameParams{
		HTTPClient: client,
	}
}

/*
DeleteChartrepoRepoChartsNameParams contains all the parameters to send to the API endpoint

	for the delete chartrepo repo charts name operation.

	Typically these are written to a http.Request.
*/
type DeleteChartrepoRepoChartsNameParams struct {

	/* Name.

	   The chart name
	*/
	Name string

	/* Repo.

	   The project name
	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete chartrepo repo charts name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteChartrepoRepoChartsNameParams) WithDefaults() *DeleteChartrepoRepoChartsNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete chartrepo repo charts name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteChartrepoRepoChartsNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete chartrepo repo charts name params
func (o *DeleteChartrepoRepoChartsNameParams) WithTimeout(timeout time.Duration) *DeleteChartrepoRepoChartsNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete chartrepo repo charts name params
func (o *DeleteChartrepoRepoChartsNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete chartrepo repo charts name params
func (o *DeleteChartrepoRepoChartsNameParams) WithContext(ctx context.Context) *DeleteChartrepoRepoChartsNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete chartrepo repo charts name params
func (o *DeleteChartrepoRepoChartsNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete chartrepo repo charts name params
func (o *DeleteChartrepoRepoChartsNameParams) WithHTTPClient(client *http.Client) *DeleteChartrepoRepoChartsNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete chartrepo repo charts name params
func (o *DeleteChartrepoRepoChartsNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete chartrepo repo charts name params
func (o *DeleteChartrepoRepoChartsNameParams) WithName(name string) *DeleteChartrepoRepoChartsNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete chartrepo repo charts name params
func (o *DeleteChartrepoRepoChartsNameParams) SetName(name string) {
	o.Name = name
}

// WithRepo adds the repo to the delete chartrepo repo charts name params
func (o *DeleteChartrepoRepoChartsNameParams) WithRepo(repo string) *DeleteChartrepoRepoChartsNameParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the delete chartrepo repo charts name params
func (o *DeleteChartrepoRepoChartsNameParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteChartrepoRepoChartsNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
