// Code generated by go-swagger; DO NOT EDIT.

package workflow_approvals

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

// NewWorkflowApprovalsWorkflowApprovalsApproveReadParams creates a new WorkflowApprovalsWorkflowApprovalsApproveReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowApprovalsWorkflowApprovalsApproveReadParams() *WorkflowApprovalsWorkflowApprovalsApproveReadParams {
	return &WorkflowApprovalsWorkflowApprovalsApproveReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowApprovalsWorkflowApprovalsApproveReadParamsWithTimeout creates a new WorkflowApprovalsWorkflowApprovalsApproveReadParams object
// with the ability to set a timeout on a request.
func NewWorkflowApprovalsWorkflowApprovalsApproveReadParamsWithTimeout(timeout time.Duration) *WorkflowApprovalsWorkflowApprovalsApproveReadParams {
	return &WorkflowApprovalsWorkflowApprovalsApproveReadParams{
		timeout: timeout,
	}
}

// NewWorkflowApprovalsWorkflowApprovalsApproveReadParamsWithContext creates a new WorkflowApprovalsWorkflowApprovalsApproveReadParams object
// with the ability to set a context for a request.
func NewWorkflowApprovalsWorkflowApprovalsApproveReadParamsWithContext(ctx context.Context) *WorkflowApprovalsWorkflowApprovalsApproveReadParams {
	return &WorkflowApprovalsWorkflowApprovalsApproveReadParams{
		Context: ctx,
	}
}

// NewWorkflowApprovalsWorkflowApprovalsApproveReadParamsWithHTTPClient creates a new WorkflowApprovalsWorkflowApprovalsApproveReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowApprovalsWorkflowApprovalsApproveReadParamsWithHTTPClient(client *http.Client) *WorkflowApprovalsWorkflowApprovalsApproveReadParams {
	return &WorkflowApprovalsWorkflowApprovalsApproveReadParams{
		HTTPClient: client,
	}
}

/*
WorkflowApprovalsWorkflowApprovalsApproveReadParams contains all the parameters to send to the API endpoint

	for the workflow approvals workflow approvals approve read operation.

	Typically these are written to a http.Request.
*/
type WorkflowApprovalsWorkflowApprovalsApproveReadParams struct {

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

// WithDefaults hydrates default values in the workflow approvals workflow approvals approve read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadParams) WithDefaults() *WorkflowApprovalsWorkflowApprovalsApproveReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow approvals workflow approvals approve read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow approvals workflow approvals approve read params
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadParams) WithTimeout(timeout time.Duration) *WorkflowApprovalsWorkflowApprovalsApproveReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow approvals workflow approvals approve read params
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow approvals workflow approvals approve read params
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadParams) WithContext(ctx context.Context) *WorkflowApprovalsWorkflowApprovalsApproveReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow approvals workflow approvals approve read params
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow approvals workflow approvals approve read params
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadParams) WithHTTPClient(client *http.Client) *WorkflowApprovalsWorkflowApprovalsApproveReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow approvals workflow approvals approve read params
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the workflow approvals workflow approvals approve read params
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadParams) WithID(id string) *WorkflowApprovalsWorkflowApprovalsApproveReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow approvals workflow approvals approve read params
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the workflow approvals workflow approvals approve read params
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadParams) WithSearch(search *string) *WorkflowApprovalsWorkflowApprovalsApproveReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the workflow approvals workflow approvals approve read params
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
