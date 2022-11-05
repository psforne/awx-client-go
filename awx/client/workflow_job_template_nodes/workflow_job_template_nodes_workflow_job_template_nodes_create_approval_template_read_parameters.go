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

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams() *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParamsWithTimeout creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams object
// with the ability to set a timeout on a request.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParamsWithTimeout(timeout time.Duration) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams{
		timeout: timeout,
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParamsWithContext creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams object
// with the ability to set a context for a request.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParamsWithContext(ctx context.Context) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams{
		Context: ctx,
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParamsWithHTTPClient creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParamsWithHTTPClient(client *http.Client) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams{
		HTTPClient: client,
	}
}

/*
WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams contains all the parameters to send to the API endpoint

	for the workflow job template nodes workflow job template nodes create approval template read operation.

	Typically these are written to a http.Request.
*/
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams struct {

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

// WithDefaults hydrates default values in the workflow job template nodes workflow job template nodes create approval template read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams) WithDefaults() *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow job template nodes workflow job template nodes create approval template read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow job template nodes workflow job template nodes create approval template read params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams) WithTimeout(timeout time.Duration) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow job template nodes workflow job template nodes create approval template read params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow job template nodes workflow job template nodes create approval template read params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams) WithContext(ctx context.Context) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow job template nodes workflow job template nodes create approval template read params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow job template nodes workflow job template nodes create approval template read params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams) WithHTTPClient(client *http.Client) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow job template nodes workflow job template nodes create approval template read params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the workflow job template nodes workflow job template nodes create approval template read params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams) WithID(id string) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow job template nodes workflow job template nodes create approval template read params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the workflow job template nodes workflow job template nodes create approval template read params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams) WithSearch(search *string) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the workflow job template nodes workflow job template nodes create approval template read params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
