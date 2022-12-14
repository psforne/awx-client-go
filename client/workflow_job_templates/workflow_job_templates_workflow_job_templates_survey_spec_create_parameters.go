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

// NewWorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams creates a new WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams() *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParamsWithTimeout creates a new WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams object
// with the ability to set a timeout on a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParamsWithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams{
		timeout: timeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParamsWithContext creates a new WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams object
// with the ability to set a context for a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParamsWithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams{
		Context: ctx,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParamsWithHTTPClient creates a new WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParamsWithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams{
		HTTPClient: client,
	}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams contains all the parameters to send to the API endpoint

	for the workflow job templates workflow job templates survey spec create operation.

	Typically these are written to a http.Request.
*/
type WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams struct {

	// Data.
	Data interface{}

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the workflow job templates workflow job templates survey spec create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams) WithDefaults() *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow job templates workflow job templates survey spec create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow job templates workflow job templates survey spec create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams) WithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow job templates workflow job templates survey spec create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow job templates workflow job templates survey spec create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams) WithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow job templates workflow job templates survey spec create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow job templates workflow job templates survey spec create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams) WithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow job templates workflow job templates survey spec create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the workflow job templates workflow job templates survey spec create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams) WithData(data interface{}) *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the workflow job templates workflow job templates survey spec create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the workflow job templates workflow job templates survey spec create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams) WithID(id string) *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow job templates workflow job templates survey spec create params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSurveySpecCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
