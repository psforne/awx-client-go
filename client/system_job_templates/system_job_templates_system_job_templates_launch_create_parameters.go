// Code generated by go-swagger; DO NOT EDIT.

package system_job_templates

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

// NewSystemJobTemplatesSystemJobTemplatesLaunchCreateParams creates a new SystemJobTemplatesSystemJobTemplatesLaunchCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemJobTemplatesSystemJobTemplatesLaunchCreateParams() *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams {
	return &SystemJobTemplatesSystemJobTemplatesLaunchCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSystemJobTemplatesSystemJobTemplatesLaunchCreateParamsWithTimeout creates a new SystemJobTemplatesSystemJobTemplatesLaunchCreateParams object
// with the ability to set a timeout on a request.
func NewSystemJobTemplatesSystemJobTemplatesLaunchCreateParamsWithTimeout(timeout time.Duration) *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams {
	return &SystemJobTemplatesSystemJobTemplatesLaunchCreateParams{
		timeout: timeout,
	}
}

// NewSystemJobTemplatesSystemJobTemplatesLaunchCreateParamsWithContext creates a new SystemJobTemplatesSystemJobTemplatesLaunchCreateParams object
// with the ability to set a context for a request.
func NewSystemJobTemplatesSystemJobTemplatesLaunchCreateParamsWithContext(ctx context.Context) *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams {
	return &SystemJobTemplatesSystemJobTemplatesLaunchCreateParams{
		Context: ctx,
	}
}

// NewSystemJobTemplatesSystemJobTemplatesLaunchCreateParamsWithHTTPClient creates a new SystemJobTemplatesSystemJobTemplatesLaunchCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemJobTemplatesSystemJobTemplatesLaunchCreateParamsWithHTTPClient(client *http.Client) *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams {
	return &SystemJobTemplatesSystemJobTemplatesLaunchCreateParams{
		HTTPClient: client,
	}
}

/*
SystemJobTemplatesSystemJobTemplatesLaunchCreateParams contains all the parameters to send to the API endpoint

	for the system job templates system job templates launch create operation.

	Typically these are written to a http.Request.
*/
type SystemJobTemplatesSystemJobTemplatesLaunchCreateParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the system job templates system job templates launch create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams) WithDefaults() *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the system job templates system job templates launch create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system job templates system job templates launch create params
func (o *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams) WithTimeout(timeout time.Duration) *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system job templates system job templates launch create params
func (o *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system job templates system job templates launch create params
func (o *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams) WithContext(ctx context.Context) *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system job templates system job templates launch create params
func (o *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system job templates system job templates launch create params
func (o *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams) WithHTTPClient(client *http.Client) *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system job templates system job templates launch create params
func (o *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the system job templates system job templates launch create params
func (o *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams) WithID(id string) *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the system job templates system job templates launch create params
func (o *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SystemJobTemplatesSystemJobTemplatesLaunchCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
