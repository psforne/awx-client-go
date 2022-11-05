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

// NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams creates a new WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams() *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParamsWithTimeout creates a new WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams object
// with the ability to set a timeout on a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParamsWithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams{
		timeout: timeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParamsWithContext creates a new WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams object
// with the ability to set a context for a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParamsWithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams{
		Context: ctx,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParamsWithHTTPClient creates a new WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParamsWithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams{
		HTTPClient: client,
	}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams contains all the parameters to send to the API endpoint

	for the workflow job templates workflow job templates notification templates success create operation.

	Typically these are written to a http.Request.
*/
type WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams struct {

	// Data.
	Data WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the workflow job templates workflow job templates notification templates success create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams) WithDefaults() *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow job templates workflow job templates notification templates success create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow job templates workflow job templates notification templates success create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams) WithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow job templates workflow job templates notification templates success create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow job templates workflow job templates notification templates success create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams) WithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow job templates workflow job templates notification templates success create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow job templates workflow job templates notification templates success create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams) WithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow job templates workflow job templates notification templates success create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the workflow job templates workflow job templates notification templates success create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams) WithData(data WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateBody) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the workflow job templates workflow job templates notification templates success create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams) SetData(data WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateBody) {
	o.Data = data
}

// WithID adds the id to the workflow job templates workflow job templates notification templates success create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams) WithID(id string) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow job templates workflow job templates notification templates success create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesSuccessCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
