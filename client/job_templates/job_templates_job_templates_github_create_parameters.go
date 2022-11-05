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

// NewJobTemplatesJobTemplatesGithubCreateParams creates a new JobTemplatesJobTemplatesGithubCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewJobTemplatesJobTemplatesGithubCreateParams() *JobTemplatesJobTemplatesGithubCreateParams {
	return &JobTemplatesJobTemplatesGithubCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewJobTemplatesJobTemplatesGithubCreateParamsWithTimeout creates a new JobTemplatesJobTemplatesGithubCreateParams object
// with the ability to set a timeout on a request.
func NewJobTemplatesJobTemplatesGithubCreateParamsWithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesGithubCreateParams {
	return &JobTemplatesJobTemplatesGithubCreateParams{
		timeout: timeout,
	}
}

// NewJobTemplatesJobTemplatesGithubCreateParamsWithContext creates a new JobTemplatesJobTemplatesGithubCreateParams object
// with the ability to set a context for a request.
func NewJobTemplatesJobTemplatesGithubCreateParamsWithContext(ctx context.Context) *JobTemplatesJobTemplatesGithubCreateParams {
	return &JobTemplatesJobTemplatesGithubCreateParams{
		Context: ctx,
	}
}

// NewJobTemplatesJobTemplatesGithubCreateParamsWithHTTPClient creates a new JobTemplatesJobTemplatesGithubCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewJobTemplatesJobTemplatesGithubCreateParamsWithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesGithubCreateParams {
	return &JobTemplatesJobTemplatesGithubCreateParams{
		HTTPClient: client,
	}
}

/*
JobTemplatesJobTemplatesGithubCreateParams contains all the parameters to send to the API endpoint

	for the job templates job templates github create operation.

	Typically these are written to a http.Request.
*/
type JobTemplatesJobTemplatesGithubCreateParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the job templates job templates github create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobTemplatesJobTemplatesGithubCreateParams) WithDefaults() *JobTemplatesJobTemplatesGithubCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the job templates job templates github create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobTemplatesJobTemplatesGithubCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the job templates job templates github create params
func (o *JobTemplatesJobTemplatesGithubCreateParams) WithTimeout(timeout time.Duration) *JobTemplatesJobTemplatesGithubCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job templates job templates github create params
func (o *JobTemplatesJobTemplatesGithubCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job templates job templates github create params
func (o *JobTemplatesJobTemplatesGithubCreateParams) WithContext(ctx context.Context) *JobTemplatesJobTemplatesGithubCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job templates job templates github create params
func (o *JobTemplatesJobTemplatesGithubCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job templates job templates github create params
func (o *JobTemplatesJobTemplatesGithubCreateParams) WithHTTPClient(client *http.Client) *JobTemplatesJobTemplatesGithubCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job templates job templates github create params
func (o *JobTemplatesJobTemplatesGithubCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the job templates job templates github create params
func (o *JobTemplatesJobTemplatesGithubCreateParams) WithID(id string) *JobTemplatesJobTemplatesGithubCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the job templates job templates github create params
func (o *JobTemplatesJobTemplatesGithubCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JobTemplatesJobTemplatesGithubCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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