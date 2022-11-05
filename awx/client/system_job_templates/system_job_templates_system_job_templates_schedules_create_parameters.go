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

// NewSystemJobTemplatesSystemJobTemplatesSchedulesCreateParams creates a new SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemJobTemplatesSystemJobTemplatesSchedulesCreateParams() *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams {
	return &SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSystemJobTemplatesSystemJobTemplatesSchedulesCreateParamsWithTimeout creates a new SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams object
// with the ability to set a timeout on a request.
func NewSystemJobTemplatesSystemJobTemplatesSchedulesCreateParamsWithTimeout(timeout time.Duration) *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams {
	return &SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams{
		timeout: timeout,
	}
}

// NewSystemJobTemplatesSystemJobTemplatesSchedulesCreateParamsWithContext creates a new SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams object
// with the ability to set a context for a request.
func NewSystemJobTemplatesSystemJobTemplatesSchedulesCreateParamsWithContext(ctx context.Context) *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams {
	return &SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams{
		Context: ctx,
	}
}

// NewSystemJobTemplatesSystemJobTemplatesSchedulesCreateParamsWithHTTPClient creates a new SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemJobTemplatesSystemJobTemplatesSchedulesCreateParamsWithHTTPClient(client *http.Client) *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams {
	return &SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams{
		HTTPClient: client,
	}
}

/*
SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams contains all the parameters to send to the API endpoint

	for the system job templates system job templates schedules create operation.

	Typically these are written to a http.Request.
*/
type SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams struct {

	// Data.
	Data SystemJobTemplatesSystemJobTemplatesSchedulesCreateBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the system job templates system job templates schedules create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams) WithDefaults() *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the system job templates system job templates schedules create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system job templates system job templates schedules create params
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams) WithTimeout(timeout time.Duration) *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system job templates system job templates schedules create params
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system job templates system job templates schedules create params
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams) WithContext(ctx context.Context) *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system job templates system job templates schedules create params
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system job templates system job templates schedules create params
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams) WithHTTPClient(client *http.Client) *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system job templates system job templates schedules create params
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the system job templates system job templates schedules create params
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams) WithData(data SystemJobTemplatesSystemJobTemplatesSchedulesCreateBody) *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the system job templates system job templates schedules create params
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams) SetData(data SystemJobTemplatesSystemJobTemplatesSchedulesCreateBody) {
	o.Data = data
}

// WithID adds the id to the system job templates system job templates schedules create params
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams) WithID(id string) *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the system job templates system job templates schedules create params
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
