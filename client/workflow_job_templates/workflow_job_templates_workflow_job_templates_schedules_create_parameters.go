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

// NewWorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams creates a new WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams() *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParamsWithTimeout creates a new WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams object
// with the ability to set a timeout on a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParamsWithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams{
		timeout: timeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParamsWithContext creates a new WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams object
// with the ability to set a context for a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParamsWithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams{
		Context: ctx,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParamsWithHTTPClient creates a new WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParamsWithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams{
		HTTPClient: client,
	}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams contains all the parameters to send to the API endpoint

	for the workflow job templates workflow job templates schedules create operation.

	Typically these are written to a http.Request.
*/
type WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams struct {

	// Data.
	Data interface{}

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the workflow job templates workflow job templates schedules create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams) WithDefaults() *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow job templates workflow job templates schedules create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow job templates workflow job templates schedules create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams) WithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow job templates workflow job templates schedules create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow job templates workflow job templates schedules create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams) WithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow job templates workflow job templates schedules create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow job templates workflow job templates schedules create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams) WithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow job templates workflow job templates schedules create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the workflow job templates workflow job templates schedules create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams) WithData(data interface{}) *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the workflow job templates workflow job templates schedules create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the workflow job templates workflow job templates schedules create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams) WithID(id string) *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow job templates workflow job templates schedules create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
