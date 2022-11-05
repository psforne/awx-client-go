// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_template_nodes

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

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams() *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParamsWithTimeout creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams object
// with the ability to set a timeout on a request.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParamsWithTimeout(timeout time.Duration) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams{
		timeout: timeout,
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParamsWithContext creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams object
// with the ability to set a context for a request.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParamsWithContext(ctx context.Context) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams{
		Context: ctx,
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParamsWithHTTPClient creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParamsWithHTTPClient(client *http.Client) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams{
		HTTPClient: client,
	}
}

/*
WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams contains all the parameters to send to the API endpoint

	for the workflow job template nodes workflow job template nodes update operation.

	Typically these are written to a http.Request.
*/
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams struct {

	// Data.
	Data WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateBody

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

// WithDefaults hydrates default values in the workflow job template nodes workflow job template nodes update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams) WithDefaults() *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow job template nodes workflow job template nodes update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow job template nodes workflow job template nodes update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams) WithTimeout(timeout time.Duration) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow job template nodes workflow job template nodes update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow job template nodes workflow job template nodes update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams) WithContext(ctx context.Context) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow job template nodes workflow job template nodes update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow job template nodes workflow job template nodes update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams) WithHTTPClient(client *http.Client) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow job template nodes workflow job template nodes update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the workflow job template nodes workflow job template nodes update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams) WithData(data WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateBody) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the workflow job template nodes workflow job template nodes update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams) SetData(data WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateBody) {
	o.Data = data
}

// WithID adds the id to the workflow job template nodes workflow job template nodes update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams) WithID(id string) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow job template nodes workflow job template nodes update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the workflow job template nodes workflow job template nodes update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams) WithSearch(search *string) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the workflow job template nodes workflow job template nodes update params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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