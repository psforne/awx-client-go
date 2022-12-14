// Code generated by go-swagger; DO NOT EDIT.

package job_templates

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

// NewJobTemplatesJobTemplatesCallbackCreateParams creates a new JobTemplatesJobTemplatesCallbackCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewJobTemplatesJobTemplatesCallbackCreateParams() *JobTemplatesJobTemplatesCallbackCreateParams {
	return &JobTemplatesJobTemplatesCallbackCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewJobTemplatesJobTemplatesCallbackCreateParamsWithTimeout creates a new JobTemplatesJobTemplatesCallbackCreateParams object
// with the ability to set a timeout on a request.
func NewJobTemplatesJobTemplatesCallbackCreateParamsWithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesCallbackCreateParams {
	return &JobTemplatesJobTemplatesCallbackCreateParams{
		timeout: timeout,
	}
}

// NewJobTemplatesJobTemplatesCallbackCreateParamsWithContext creates a new JobTemplatesJobTemplatesCallbackCreateParams object
// with the ability to set a context for a request.
func NewJobTemplatesJobTemplatesCallbackCreateParamsWithContext(ctx context.Context) *JobTemplatesJobTemplatesCallbackCreateParams {
	return &JobTemplatesJobTemplatesCallbackCreateParams{
		Context: ctx,
	}
}

// NewJobTemplatesJobTemplatesCallbackCreateParamsWithHTTPClient creates a new JobTemplatesJobTemplatesCallbackCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewJobTemplatesJobTemplatesCallbackCreateParamsWithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesCallbackCreateParams {
	return &JobTemplatesJobTemplatesCallbackCreateParams{
		HTTPClient: client,
	}
}

/*
JobTemplatesJobTemplatesCallbackCreateParams contains all the parameters to send to the API endpoint

	for the job templates job templates callback create operation.

	Typically these are written to a http.Request.
*/
type JobTemplatesJobTemplatesCallbackCreateParams struct {

	// Data.
	Data interface{}

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the job templates job templates callback create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobTemplatesJobTemplatesCallbackCreateParams) WithDefaults() *JobTemplatesJobTemplatesCallbackCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the job templates job templates callback create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobTemplatesJobTemplatesCallbackCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the job templates job templates callback create params
func (o *JobTemplatesJobTemplatesCallbackCreateParams) WithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesCallbackCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job templates job templates callback create params
func (o *JobTemplatesJobTemplatesCallbackCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job templates job templates callback create params
func (o *JobTemplatesJobTemplatesCallbackCreateParams) WithContext(ctx context.Context) *JobTemplatesJobTemplatesCallbackCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job templates job templates callback create params
func (o *JobTemplatesJobTemplatesCallbackCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job templates job templates callback create params
func (o *JobTemplatesJobTemplatesCallbackCreateParams) WithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesCallbackCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job templates job templates callback create params
func (o *JobTemplatesJobTemplatesCallbackCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the job templates job templates callback create params
func (o *JobTemplatesJobTemplatesCallbackCreateParams) WithData(data interface{}) *JobTemplatesJobTemplatesCallbackCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the job templates job templates callback create params
func (o *JobTemplatesJobTemplatesCallbackCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the job templates job templates callback create params
func (o *JobTemplatesJobTemplatesCallbackCreateParams) WithID(id string) *JobTemplatesJobTemplatesCallbackCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the job templates job templates callback create params
func (o *JobTemplatesJobTemplatesCallbackCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JobTemplatesJobTemplatesCallbackCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
