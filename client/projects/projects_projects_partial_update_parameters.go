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

// NewProjectsProjectsPartialUpdateParams creates a new ProjectsProjectsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectsProjectsPartialUpdateParams() *ProjectsProjectsPartialUpdateParams {
	return &ProjectsProjectsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsProjectsPartialUpdateParamsWithTimeout creates a new ProjectsProjectsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewProjectsProjectsPartialUpdateParamsWithTimeout(timeout time.Duration) *ProjectsProjectsPartialUpdateParams {
	return &ProjectsProjectsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewProjectsProjectsPartialUpdateParamsWithContext creates a new ProjectsProjectsPartialUpdateParams object
// with the ability to set a context for a request.
func NewProjectsProjectsPartialUpdateParamsWithContext(ctx context.Context) *ProjectsProjectsPartialUpdateParams {
	return &ProjectsProjectsPartialUpdateParams{
		Context: ctx,
	}
}

// NewProjectsProjectsPartialUpdateParamsWithHTTPClient creates a new ProjectsProjectsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectsProjectsPartialUpdateParamsWithHTTPClient(client *http.Client) *ProjectsProjectsPartialUpdateParams {
	return &ProjectsProjectsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
ProjectsProjectsPartialUpdateParams contains all the parameters to send to the API endpoint

	for the projects projects partial update operation.

	Typically these are written to a http.Request.
*/
type ProjectsProjectsPartialUpdateParams struct {

	// Data.
	Data interface{}

	// ID.
	ID string

	/* Search.

	   A search term.
	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the projects projects partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsProjectsPartialUpdateParams) WithDefaults() *ProjectsProjectsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the projects projects partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsProjectsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the projects projects partial update params
func (o *ProjectsProjectsPartialUpdateParams) WithTimeout(timeout time.Duration) *ProjectsProjectsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects projects partial update params
func (o *ProjectsProjectsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects projects partial update params
func (o *ProjectsProjectsPartialUpdateParams) WithContext(ctx context.Context) *ProjectsProjectsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects projects partial update params
func (o *ProjectsProjectsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects projects partial update params
func (o *ProjectsProjectsPartialUpdateParams) WithHTTPClient(client *http.Client) *ProjectsProjectsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects projects partial update params
func (o *ProjectsProjectsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the projects projects partial update params
func (o *ProjectsProjectsPartialUpdateParams) WithData(data interface{}) *ProjectsProjectsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the projects projects partial update params
func (o *ProjectsProjectsPartialUpdateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the projects projects partial update params
func (o *ProjectsProjectsPartialUpdateParams) WithID(id string) *ProjectsProjectsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects projects partial update params
func (o *ProjectsProjectsPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the projects projects partial update params
func (o *ProjectsProjectsPartialUpdateParams) WithSearch(search *string) *ProjectsProjectsPartialUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the projects projects partial update params
func (o *ProjectsProjectsPartialUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsProjectsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Search != nil {

		// query param search
		var qrSearch string

		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {

			if err := r.SetQueryParam("search", qSearch); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
