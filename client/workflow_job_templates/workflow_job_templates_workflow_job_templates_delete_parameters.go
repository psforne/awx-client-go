// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

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

// NewWorkflowJobTemplatesWorkflowJobTemplatesDeleteParams creates a new WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowJobTemplatesWorkflowJobTemplatesDeleteParams() *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesDeleteParamsWithTimeout creates a new WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams object
// with the ability to set a timeout on a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesDeleteParamsWithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams{
		timeout: timeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesDeleteParamsWithContext creates a new WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams object
// with the ability to set a context for a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesDeleteParamsWithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams{
		Context: ctx,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesDeleteParamsWithHTTPClient creates a new WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesDeleteParamsWithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams{
		HTTPClient: client,
	}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams contains all the parameters to send to the API endpoint

	for the workflow job templates workflow job templates delete operation.

	Typically these are written to a http.Request.
*/
type WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams struct {

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

// WithDefaults hydrates default values in the workflow job templates workflow job templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams) WithDefaults() *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow job templates workflow job templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow job templates workflow job templates delete params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams) WithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow job templates workflow job templates delete params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow job templates workflow job templates delete params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams) WithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow job templates workflow job templates delete params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow job templates workflow job templates delete params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams) WithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow job templates workflow job templates delete params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the workflow job templates workflow job templates delete params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams) WithID(id string) *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow job templates workflow job templates delete params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the workflow job templates workflow job templates delete params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams) WithSearch(search *string) *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the workflow job templates workflow job templates delete params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobTemplatesWorkflowJobTemplatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
