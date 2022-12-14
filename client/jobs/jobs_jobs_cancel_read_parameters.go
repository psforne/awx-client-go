// Code generated by go-swagger; DO NOT EDIT.

package jobs

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

// NewJobsJobsCancelReadParams creates a new JobsJobsCancelReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewJobsJobsCancelReadParams() *JobsJobsCancelReadParams {
	return &JobsJobsCancelReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewJobsJobsCancelReadParamsWithTimeout creates a new JobsJobsCancelReadParams object
// with the ability to set a timeout on a request.
func NewJobsJobsCancelReadParamsWithTimeout(timeout time.Duration) *JobsJobsCancelReadParams {
	return &JobsJobsCancelReadParams{
		timeout: timeout,
	}
}

// NewJobsJobsCancelReadParamsWithContext creates a new JobsJobsCancelReadParams object
// with the ability to set a context for a request.
func NewJobsJobsCancelReadParamsWithContext(ctx context.Context) *JobsJobsCancelReadParams {
	return &JobsJobsCancelReadParams{
		Context: ctx,
	}
}

// NewJobsJobsCancelReadParamsWithHTTPClient creates a new JobsJobsCancelReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewJobsJobsCancelReadParamsWithHTTPClient(client *http.Client) *JobsJobsCancelReadParams {
	return &JobsJobsCancelReadParams{
		HTTPClient: client,
	}
}

/*
JobsJobsCancelReadParams contains all the parameters to send to the API endpoint

	for the jobs jobs cancel read operation.

	Typically these are written to a http.Request.
*/
type JobsJobsCancelReadParams struct {

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

// WithDefaults hydrates default values in the jobs jobs cancel read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobsJobsCancelReadParams) WithDefaults() *JobsJobsCancelReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the jobs jobs cancel read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobsJobsCancelReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the jobs jobs cancel read params
func (o *JobsJobsCancelReadParams) WithTimeout(timeout time.Duration) *JobsJobsCancelReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the jobs jobs cancel read params
func (o *JobsJobsCancelReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the jobs jobs cancel read params
func (o *JobsJobsCancelReadParams) WithContext(ctx context.Context) *JobsJobsCancelReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the jobs jobs cancel read params
func (o *JobsJobsCancelReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the jobs jobs cancel read params
func (o *JobsJobsCancelReadParams) WithHTTPClient(client *http.Client) *JobsJobsCancelReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the jobs jobs cancel read params
func (o *JobsJobsCancelReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the jobs jobs cancel read params
func (o *JobsJobsCancelReadParams) WithID(id string) *JobsJobsCancelReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the jobs jobs cancel read params
func (o *JobsJobsCancelReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the jobs jobs cancel read params
func (o *JobsJobsCancelReadParams) WithSearch(search *string) *JobsJobsCancelReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the jobs jobs cancel read params
func (o *JobsJobsCancelReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *JobsJobsCancelReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
