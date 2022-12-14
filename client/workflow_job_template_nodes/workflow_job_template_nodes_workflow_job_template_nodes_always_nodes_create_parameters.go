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

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams() *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParamsWithTimeout creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams object
// with the ability to set a timeout on a request.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParamsWithTimeout(timeout time.Duration) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams{
		timeout: timeout,
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParamsWithContext creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams object
// with the ability to set a context for a request.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParamsWithContext(ctx context.Context) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams{
		Context: ctx,
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParamsWithHTTPClient creates a new WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParamsWithHTTPClient(client *http.Client) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams{
		HTTPClient: client,
	}
}

/*
WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams contains all the parameters to send to the API endpoint

	for the workflow job template nodes workflow job template nodes always nodes create operation.

	Typically these are written to a http.Request.
*/
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams struct {

	// Data.
	Data interface{}

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the workflow job template nodes workflow job template nodes always nodes create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams) WithDefaults() *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow job template nodes workflow job template nodes always nodes create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow job template nodes workflow job template nodes always nodes create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams) WithTimeout(timeout time.Duration) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow job template nodes workflow job template nodes always nodes create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow job template nodes workflow job template nodes always nodes create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams) WithContext(ctx context.Context) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow job template nodes workflow job template nodes always nodes create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow job template nodes workflow job template nodes always nodes create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams) WithHTTPClient(client *http.Client) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow job template nodes workflow job template nodes always nodes create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the workflow job template nodes workflow job template nodes always nodes create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams) WithData(data interface{}) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the workflow job template nodes workflow job template nodes always nodes create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the workflow job template nodes workflow job template nodes always nodes create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams) WithID(id string) *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow job template nodes workflow job template nodes always nodes create params
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
