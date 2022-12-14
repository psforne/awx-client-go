// Code generated by go-swagger; DO NOT EDIT.

package workflow_approval_templates

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

// NewWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams creates a new WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams() *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams {
	return &WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParamsWithTimeout creates a new WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams object
// with the ability to set a timeout on a request.
func NewWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParamsWithTimeout(timeout time.Duration) *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams {
	return &WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams{
		timeout: timeout,
	}
}

// NewWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParamsWithContext creates a new WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams object
// with the ability to set a context for a request.
func NewWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParamsWithContext(ctx context.Context) *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams {
	return &WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams{
		Context: ctx,
	}
}

// NewWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParamsWithHTTPClient creates a new WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParamsWithHTTPClient(client *http.Client) *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams {
	return &WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams{
		HTTPClient: client,
	}
}

/*
WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams contains all the parameters to send to the API endpoint

	for the workflow approval templates workflow approval templates delete operation.

	Typically these are written to a http.Request.
*/
type WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams struct {

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

// WithDefaults hydrates default values in the workflow approval templates workflow approval templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams) WithDefaults() *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow approval templates workflow approval templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow approval templates workflow approval templates delete params
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams) WithTimeout(timeout time.Duration) *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow approval templates workflow approval templates delete params
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow approval templates workflow approval templates delete params
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams) WithContext(ctx context.Context) *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow approval templates workflow approval templates delete params
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow approval templates workflow approval templates delete params
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams) WithHTTPClient(client *http.Client) *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow approval templates workflow approval templates delete params
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the workflow approval templates workflow approval templates delete params
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams) WithID(id string) *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow approval templates workflow approval templates delete params
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the workflow approval templates workflow approval templates delete params
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams) WithSearch(search *string) *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the workflow approval templates workflow approval templates delete params
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
