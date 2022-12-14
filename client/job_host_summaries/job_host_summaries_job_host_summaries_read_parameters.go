// Code generated by go-swagger; DO NOT EDIT.

package job_host_summaries

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

// NewJobHostSummariesJobHostSummariesReadParams creates a new JobHostSummariesJobHostSummariesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewJobHostSummariesJobHostSummariesReadParams() *JobHostSummariesJobHostSummariesReadParams {
	return &JobHostSummariesJobHostSummariesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewJobHostSummariesJobHostSummariesReadParamsWithTimeout creates a new JobHostSummariesJobHostSummariesReadParams object
// with the ability to set a timeout on a request.
func NewJobHostSummariesJobHostSummariesReadParamsWithTimeout(timeout time.Duration) *JobHostSummariesJobHostSummariesReadParams {
	return &JobHostSummariesJobHostSummariesReadParams{
		timeout: timeout,
	}
}

// NewJobHostSummariesJobHostSummariesReadParamsWithContext creates a new JobHostSummariesJobHostSummariesReadParams object
// with the ability to set a context for a request.
func NewJobHostSummariesJobHostSummariesReadParamsWithContext(ctx context.Context) *JobHostSummariesJobHostSummariesReadParams {
	return &JobHostSummariesJobHostSummariesReadParams{
		Context: ctx,
	}
}

// NewJobHostSummariesJobHostSummariesReadParamsWithHTTPClient creates a new JobHostSummariesJobHostSummariesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewJobHostSummariesJobHostSummariesReadParamsWithHTTPClient(client *http.Client) *JobHostSummariesJobHostSummariesReadParams {
	return &JobHostSummariesJobHostSummariesReadParams{
		HTTPClient: client,
	}
}

/*
JobHostSummariesJobHostSummariesReadParams contains all the parameters to send to the API endpoint

	for the job host summaries job host summaries read operation.

	Typically these are written to a http.Request.
*/
type JobHostSummariesJobHostSummariesReadParams struct {

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

// WithDefaults hydrates default values in the job host summaries job host summaries read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobHostSummariesJobHostSummariesReadParams) WithDefaults() *JobHostSummariesJobHostSummariesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the job host summaries job host summaries read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobHostSummariesJobHostSummariesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the job host summaries job host summaries read params
func (o *JobHostSummariesJobHostSummariesReadParams) WithTimeout(timeout time.Duration) *JobHostSummariesJobHostSummariesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job host summaries job host summaries read params
func (o *JobHostSummariesJobHostSummariesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job host summaries job host summaries read params
func (o *JobHostSummariesJobHostSummariesReadParams) WithContext(ctx context.Context) *JobHostSummariesJobHostSummariesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job host summaries job host summaries read params
func (o *JobHostSummariesJobHostSummariesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job host summaries job host summaries read params
func (o *JobHostSummariesJobHostSummariesReadParams) WithHTTPClient(client *http.Client) *JobHostSummariesJobHostSummariesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job host summaries job host summaries read params
func (o *JobHostSummariesJobHostSummariesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the job host summaries job host summaries read params
func (o *JobHostSummariesJobHostSummariesReadParams) WithID(id string) *JobHostSummariesJobHostSummariesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the job host summaries job host summaries read params
func (o *JobHostSummariesJobHostSummariesReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the job host summaries job host summaries read params
func (o *JobHostSummariesJobHostSummariesReadParams) WithSearch(search *string) *JobHostSummariesJobHostSummariesReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the job host summaries job host summaries read params
func (o *JobHostSummariesJobHostSummariesReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *JobHostSummariesJobHostSummariesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
