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

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams() *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParamsWithTimeout creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams object
// with the ability to set a timeout on a request.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParamsWithTimeout(timeout time.Duration) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams{
		timeout: timeout,
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParamsWithContext creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams object
// with the ability to set a context for a request.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParamsWithContext(ctx context.Context) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams{
		Context: ctx,
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParamsWithHTTPClient creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParamsWithHTTPClient(client *http.Client) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams{
		HTTPClient: client,
	}
}

/*
WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams contains all the parameters to send to the API endpoint

	for the workflow job template nodes workflow job template nodes create approval template create operation.

	Typically these are written to a http.Request.
*/
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams struct {

	// Data.
	Data interface{}

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the workflow job template nodes workflow job template nodes create approval template create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams) WithDefaults() *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow job template nodes workflow job template nodes create approval template create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow job template nodes workflow job template nodes create approval template create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams) WithTimeout(timeout time.Duration) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow job template nodes workflow job template nodes create approval template create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow job template nodes workflow job template nodes create approval template create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams) WithContext(ctx context.Context) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow job template nodes workflow job template nodes create approval template create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow job template nodes workflow job template nodes create approval template create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams) WithHTTPClient(client *http.Client) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow job template nodes workflow job template nodes create approval template create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the workflow job template nodes workflow job template nodes create approval template create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams) WithData(data interface{}) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the workflow job template nodes workflow job template nodes create approval template create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the workflow job template nodes workflow job template nodes create approval template create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams) WithID(id string) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow job template nodes workflow job template nodes create approval template create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesCreateApprovalTemplateCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
