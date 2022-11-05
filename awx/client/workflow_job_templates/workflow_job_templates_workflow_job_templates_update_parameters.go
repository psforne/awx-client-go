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

// NewWorkflowJobTemplatesWorkflowJobTemplatesUpdateParams creates a new WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowJobTemplatesWorkflowJobTemplatesUpdateParams() *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesUpdateParamsWithTimeout creates a new WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams object
// with the ability to set a timeout on a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesUpdateParamsWithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams{
		timeout: timeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesUpdateParamsWithContext creates a new WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams object
// with the ability to set a context for a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesUpdateParamsWithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams{
		Context: ctx,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesUpdateParamsWithHTTPClient creates a new WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesUpdateParamsWithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams{
		HTTPClient: client,
	}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams contains all the parameters to send to the API endpoint

	for the workflow job templates workflow job templates update operation.

	Typically these are written to a http.Request.
*/
type WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams struct {

	// Data.
	Data WorkflowJobTemplatesWorkflowJobTemplatesUpdateBody

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

// WithDefaults hydrates default values in the workflow job templates workflow job templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams) WithDefaults() *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow job templates workflow job templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow job templates workflow job templates update params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams) WithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow job templates workflow job templates update params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow job templates workflow job templates update params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams) WithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow job templates workflow job templates update params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow job templates workflow job templates update params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams) WithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow job templates workflow job templates update params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the workflow job templates workflow job templates update params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams) WithData(data WorkflowJobTemplatesWorkflowJobTemplatesUpdateBody) *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the workflow job templates workflow job templates update params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams) SetData(data WorkflowJobTemplatesWorkflowJobTemplatesUpdateBody) {
	o.Data = data
}

// WithID adds the id to the workflow job templates workflow job templates update params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams) WithID(id string) *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow job templates workflow job templates update params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the workflow job templates workflow job templates update params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams) WithSearch(search *string) *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the workflow job templates workflow job templates update params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobTemplatesWorkflowJobTemplatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Data); err != nil {
		return err
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
