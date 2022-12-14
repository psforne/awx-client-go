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

// NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams creates a new WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams() *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParamsWithTimeout creates a new WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams object
// with the ability to set a timeout on a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParamsWithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams{
		timeout: timeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParamsWithContext creates a new WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams object
// with the ability to set a context for a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParamsWithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams{
		Context: ctx,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParamsWithHTTPClient creates a new WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParamsWithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams{
		HTTPClient: client,
	}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams contains all the parameters to send to the API endpoint

	for the workflow job templates workflow job templates webhook key create operation.

	Typically these are written to a http.Request.
*/
type WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the workflow job templates workflow job templates webhook key create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams) WithDefaults() *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow job templates workflow job templates webhook key create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow job templates workflow job templates webhook key create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams) WithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow job templates workflow job templates webhook key create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow job templates workflow job templates webhook key create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams) WithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow job templates workflow job templates webhook key create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow job templates workflow job templates webhook key create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams) WithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow job templates workflow job templates webhook key create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the workflow job templates workflow job templates webhook key create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams) WithID(id string) *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow job templates workflow job templates webhook key create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
