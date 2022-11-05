// Code generated by go-swagger; DO NOT EDIT.

package job_templates

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

// NewJobTemplatesJobTemplatesUpdateParams creates a new JobTemplatesJobTemplatesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewJobTemplatesJobTemplatesUpdateParams() *JobTemplatesJobTemplatesUpdateParams {
	return &JobTemplatesJobTemplatesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewJobTemplatesJobTemplatesUpdateParamsWithTimeout creates a new JobTemplatesJobTemplatesUpdateParams object
// with the ability to set a timeout on a request.
func NewJobTemplatesJobTemplatesUpdateParamsWithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesUpdateParams {
	return &JobTemplatesJobTemplatesUpdateParams{
		timeout: timeout,
	}
}

// NewJobTemplatesJobTemplatesUpdateParamsWithContext creates a new JobTemplatesJobTemplatesUpdateParams object
// with the ability to set a context for a request.
func NewJobTemplatesJobTemplatesUpdateParamsWithContext(ctx context.Context) *JobTemplatesJobTemplatesUpdateParams {
	return &JobTemplatesJobTemplatesUpdateParams{
		Context: ctx,
	}
}

// NewJobTemplatesJobTemplatesUpdateParamsWithHTTPClient creates a new JobTemplatesJobTemplatesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewJobTemplatesJobTemplatesUpdateParamsWithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesUpdateParams {
	return &JobTemplatesJobTemplatesUpdateParams{
		HTTPClient: client,
	}
}

/*
JobTemplatesJobTemplatesUpdateParams contains all the parameters to send to the API endpoint

	for the job templates job templates update operation.

	Typically these are written to a http.Request.
*/
type JobTemplatesJobTemplatesUpdateParams struct {

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

// WithDefaults hydrates default values in the job templates job templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobTemplatesJobTemplatesUpdateParams) WithDefaults() *JobTemplatesJobTemplatesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the job templates job templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobTemplatesJobTemplatesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the job templates job templates update params
func (o *JobTemplatesJobTemplatesUpdateParams) WithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job templates job templates update params
func (o *JobTemplatesJobTemplatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job templates job templates update params
func (o *JobTemplatesJobTemplatesUpdateParams) WithContext(ctx context.Context) *JobTemplatesJobTemplatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job templates job templates update params
func (o *JobTemplatesJobTemplatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job templates job templates update params
func (o *JobTemplatesJobTemplatesUpdateParams) WithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job templates job templates update params
func (o *JobTemplatesJobTemplatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the job templates job templates update params
func (o *JobTemplatesJobTemplatesUpdateParams) WithData(data interface{}) *JobTemplatesJobTemplatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the job templates job templates update params
func (o *JobTemplatesJobTemplatesUpdateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the job templates job templates update params
func (o *JobTemplatesJobTemplatesUpdateParams) WithID(id string) *JobTemplatesJobTemplatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the job templates job templates update params
func (o *JobTemplatesJobTemplatesUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the job templates job templates update params
func (o *JobTemplatesJobTemplatesUpdateParams) WithSearch(search *string) *JobTemplatesJobTemplatesUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the job templates job templates update params
func (o *JobTemplatesJobTemplatesUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *JobTemplatesJobTemplatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
