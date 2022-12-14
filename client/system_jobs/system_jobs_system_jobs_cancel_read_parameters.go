// Code generated by go-swagger; DO NOT EDIT.

package system_jobs

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

// NewSystemJobsSystemJobsCancelReadParams creates a new SystemJobsSystemJobsCancelReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemJobsSystemJobsCancelReadParams() *SystemJobsSystemJobsCancelReadParams {
	return &SystemJobsSystemJobsCancelReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSystemJobsSystemJobsCancelReadParamsWithTimeout creates a new SystemJobsSystemJobsCancelReadParams object
// with the ability to set a timeout on a request.
func NewSystemJobsSystemJobsCancelReadParamsWithTimeout(timeout time.Duration) *SystemJobsSystemJobsCancelReadParams {
	return &SystemJobsSystemJobsCancelReadParams{
		timeout: timeout,
	}
}

// NewSystemJobsSystemJobsCancelReadParamsWithContext creates a new SystemJobsSystemJobsCancelReadParams object
// with the ability to set a context for a request.
func NewSystemJobsSystemJobsCancelReadParamsWithContext(ctx context.Context) *SystemJobsSystemJobsCancelReadParams {
	return &SystemJobsSystemJobsCancelReadParams{
		Context: ctx,
	}
}

// NewSystemJobsSystemJobsCancelReadParamsWithHTTPClient creates a new SystemJobsSystemJobsCancelReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemJobsSystemJobsCancelReadParamsWithHTTPClient(client *http.Client) *SystemJobsSystemJobsCancelReadParams {
	return &SystemJobsSystemJobsCancelReadParams{
		HTTPClient: client,
	}
}

/*
SystemJobsSystemJobsCancelReadParams contains all the parameters to send to the API endpoint

	for the system jobs system jobs cancel read operation.

	Typically these are written to a http.Request.
*/
type SystemJobsSystemJobsCancelReadParams struct {

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

// WithDefaults hydrates default values in the system jobs system jobs cancel read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemJobsSystemJobsCancelReadParams) WithDefaults() *SystemJobsSystemJobsCancelReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the system jobs system jobs cancel read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemJobsSystemJobsCancelReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system jobs system jobs cancel read params
func (o *SystemJobsSystemJobsCancelReadParams) WithTimeout(timeout time.Duration) *SystemJobsSystemJobsCancelReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system jobs system jobs cancel read params
func (o *SystemJobsSystemJobsCancelReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system jobs system jobs cancel read params
func (o *SystemJobsSystemJobsCancelReadParams) WithContext(ctx context.Context) *SystemJobsSystemJobsCancelReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system jobs system jobs cancel read params
func (o *SystemJobsSystemJobsCancelReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system jobs system jobs cancel read params
func (o *SystemJobsSystemJobsCancelReadParams) WithHTTPClient(client *http.Client) *SystemJobsSystemJobsCancelReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system jobs system jobs cancel read params
func (o *SystemJobsSystemJobsCancelReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the system jobs system jobs cancel read params
func (o *SystemJobsSystemJobsCancelReadParams) WithID(id string) *SystemJobsSystemJobsCancelReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the system jobs system jobs cancel read params
func (o *SystemJobsSystemJobsCancelReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the system jobs system jobs cancel read params
func (o *SystemJobsSystemJobsCancelReadParams) WithSearch(search *string) *SystemJobsSystemJobsCancelReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the system jobs system jobs cancel read params
func (o *SystemJobsSystemJobsCancelReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *SystemJobsSystemJobsCancelReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
