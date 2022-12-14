// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewProjectsProjectsUpdateCreateParams creates a new ProjectsProjectsUpdateCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectsProjectsUpdateCreateParams() *ProjectsProjectsUpdateCreateParams {
	return &ProjectsProjectsUpdateCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsProjectsUpdateCreateParamsWithTimeout creates a new ProjectsProjectsUpdateCreateParams object
// with the ability to set a timeout on a request.
func NewProjectsProjectsUpdateCreateParamsWithTimeout(timeout time.Duration) *ProjectsProjectsUpdateCreateParams {
	return &ProjectsProjectsUpdateCreateParams{
		timeout: timeout,
	}
}

// NewProjectsProjectsUpdateCreateParamsWithContext creates a new ProjectsProjectsUpdateCreateParams object
// with the ability to set a context for a request.
func NewProjectsProjectsUpdateCreateParamsWithContext(ctx context.Context) *ProjectsProjectsUpdateCreateParams {
	return &ProjectsProjectsUpdateCreateParams{
		Context: ctx,
	}
}

// NewProjectsProjectsUpdateCreateParamsWithHTTPClient creates a new ProjectsProjectsUpdateCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectsProjectsUpdateCreateParamsWithHTTPClient(client *http.Client) *ProjectsProjectsUpdateCreateParams {
	return &ProjectsProjectsUpdateCreateParams{
		HTTPClient: client,
	}
}

/*
ProjectsProjectsUpdateCreateParams contains all the parameters to send to the API endpoint

	for the projects projects update create operation.

	Typically these are written to a http.Request.
*/
type ProjectsProjectsUpdateCreateParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the projects projects update create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsProjectsUpdateCreateParams) WithDefaults() *ProjectsProjectsUpdateCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the projects projects update create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsProjectsUpdateCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the projects projects update create params
func (o *ProjectsProjectsUpdateCreateParams) WithTimeout(timeout time.Duration) *ProjectsProjectsUpdateCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects projects update create params
func (o *ProjectsProjectsUpdateCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects projects update create params
func (o *ProjectsProjectsUpdateCreateParams) WithContext(ctx context.Context) *ProjectsProjectsUpdateCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects projects update create params
func (o *ProjectsProjectsUpdateCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects projects update create params
func (o *ProjectsProjectsUpdateCreateParams) WithHTTPClient(client *http.Client) *ProjectsProjectsUpdateCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects projects update create params
func (o *ProjectsProjectsUpdateCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the projects projects update create params
func (o *ProjectsProjectsUpdateCreateParams) WithID(id string) *ProjectsProjectsUpdateCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects projects update create params
func (o *ProjectsProjectsUpdateCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsProjectsUpdateCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
