// Code generated by go-swagger; DO NOT EDIT.

package job_events

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

// NewJobEventsJobEventsReadParams creates a new JobEventsJobEventsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewJobEventsJobEventsReadParams() *JobEventsJobEventsReadParams {
	return &JobEventsJobEventsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewJobEventsJobEventsReadParamsWithTimeout creates a new JobEventsJobEventsReadParams object
// with the ability to set a timeout on a request.
func NewJobEventsJobEventsReadParamsWithTimeout(timeout time.Duration) *JobEventsJobEventsReadParams {
	return &JobEventsJobEventsReadParams{
		timeout: timeout,
	}
}

// NewJobEventsJobEventsReadParamsWithContext creates a new JobEventsJobEventsReadParams object
// with the ability to set a context for a request.
func NewJobEventsJobEventsReadParamsWithContext(ctx context.Context) *JobEventsJobEventsReadParams {
	return &JobEventsJobEventsReadParams{
		Context: ctx,
	}
}

// NewJobEventsJobEventsReadParamsWithHTTPClient creates a new JobEventsJobEventsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewJobEventsJobEventsReadParamsWithHTTPClient(client *http.Client) *JobEventsJobEventsReadParams {
	return &JobEventsJobEventsReadParams{
		HTTPClient: client,
	}
}

/*
JobEventsJobEventsReadParams contains all the parameters to send to the API endpoint

	for the job events job events read operation.

	Typically these are written to a http.Request.
*/
type JobEventsJobEventsReadParams struct {

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

// WithDefaults hydrates default values in the job events job events read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobEventsJobEventsReadParams) WithDefaults() *JobEventsJobEventsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the job events job events read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobEventsJobEventsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the job events job events read params
func (o *JobEventsJobEventsReadParams) WithTimeout(timeout time.Duration) *JobEventsJobEventsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job events job events read params
func (o *JobEventsJobEventsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job events job events read params
func (o *JobEventsJobEventsReadParams) WithContext(ctx context.Context) *JobEventsJobEventsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job events job events read params
func (o *JobEventsJobEventsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job events job events read params
func (o *JobEventsJobEventsReadParams) WithHTTPClient(client *http.Client) *JobEventsJobEventsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job events job events read params
func (o *JobEventsJobEventsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the job events job events read params
func (o *JobEventsJobEventsReadParams) WithID(id string) *JobEventsJobEventsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the job events job events read params
func (o *JobEventsJobEventsReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the job events job events read params
func (o *JobEventsJobEventsReadParams) WithSearch(search *string) *JobEventsJobEventsReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the job events job events read params
func (o *JobEventsJobEventsReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *JobEventsJobEventsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
