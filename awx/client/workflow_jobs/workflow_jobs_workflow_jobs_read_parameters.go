// Code generated by go-swagger; DO NOT EDIT.

package workflow_jobs

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

// NewWorkflowJobsWorkflowJobsReadParams creates a new WorkflowJobsWorkflowJobsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowJobsWorkflowJobsReadParams() *WorkflowJobsWorkflowJobsReadParams {
	return &WorkflowJobsWorkflowJobsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobsWorkflowJobsReadParamsWithTimeout creates a new WorkflowJobsWorkflowJobsReadParams object
// with the ability to set a timeout on a request.
func NewWorkflowJobsWorkflowJobsReadParamsWithTimeout(timeout time.Duration) *WorkflowJobsWorkflowJobsReadParams {
	return &WorkflowJobsWorkflowJobsReadParams{
		timeout: timeout,
	}
}

// NewWorkflowJobsWorkflowJobsReadParamsWithContext creates a new WorkflowJobsWorkflowJobsReadParams object
// with the ability to set a context for a request.
func NewWorkflowJobsWorkflowJobsReadParamsWithContext(ctx context.Context) *WorkflowJobsWorkflowJobsReadParams {
	return &WorkflowJobsWorkflowJobsReadParams{
		Context: ctx,
	}
}

// NewWorkflowJobsWorkflowJobsReadParamsWithHTTPClient creates a new WorkflowJobsWorkflowJobsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowJobsWorkflowJobsReadParamsWithHTTPClient(client *http.Client) *WorkflowJobsWorkflowJobsReadParams {
	return &WorkflowJobsWorkflowJobsReadParams{
		HTTPClient: client,
	}
}

/*
WorkflowJobsWorkflowJobsReadParams contains all the parameters to send to the API endpoint

	for the workflow jobs workflow jobs read operation.

	Typically these are written to a http.Request.
*/
type WorkflowJobsWorkflowJobsReadParams struct {

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

// WithDefaults hydrates default values in the workflow jobs workflow jobs read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobsWorkflowJobsReadParams) WithDefaults() *WorkflowJobsWorkflowJobsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow jobs workflow jobs read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobsWorkflowJobsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow jobs workflow jobs read params
func (o *WorkflowJobsWorkflowJobsReadParams) WithTimeout(timeout time.Duration) *WorkflowJobsWorkflowJobsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow jobs workflow jobs read params
func (o *WorkflowJobsWorkflowJobsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow jobs workflow jobs read params
func (o *WorkflowJobsWorkflowJobsReadParams) WithContext(ctx context.Context) *WorkflowJobsWorkflowJobsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow jobs workflow jobs read params
func (o *WorkflowJobsWorkflowJobsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow jobs workflow jobs read params
func (o *WorkflowJobsWorkflowJobsReadParams) WithHTTPClient(client *http.Client) *WorkflowJobsWorkflowJobsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow jobs workflow jobs read params
func (o *WorkflowJobsWorkflowJobsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the workflow jobs workflow jobs read params
func (o *WorkflowJobsWorkflowJobsReadParams) WithID(id string) *WorkflowJobsWorkflowJobsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow jobs workflow jobs read params
func (o *WorkflowJobsWorkflowJobsReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the workflow jobs workflow jobs read params
func (o *WorkflowJobsWorkflowJobsReadParams) WithSearch(search *string) *WorkflowJobsWorkflowJobsReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the workflow jobs workflow jobs read params
func (o *WorkflowJobsWorkflowJobsReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobsWorkflowJobsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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