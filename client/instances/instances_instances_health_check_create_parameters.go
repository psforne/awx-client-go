// Code generated by go-swagger; DO NOT EDIT.

package instances

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

// NewInstancesInstancesHealthCheckCreateParams creates a new InstancesInstancesHealthCheckCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInstancesInstancesHealthCheckCreateParams() *InstancesInstancesHealthCheckCreateParams {
	return &InstancesInstancesHealthCheckCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInstancesInstancesHealthCheckCreateParamsWithTimeout creates a new InstancesInstancesHealthCheckCreateParams object
// with the ability to set a timeout on a request.
func NewInstancesInstancesHealthCheckCreateParamsWithTimeout(timeout time.Duration) *InstancesInstancesHealthCheckCreateParams {
	return &InstancesInstancesHealthCheckCreateParams{
		timeout: timeout,
	}
}

// NewInstancesInstancesHealthCheckCreateParamsWithContext creates a new InstancesInstancesHealthCheckCreateParams object
// with the ability to set a context for a request.
func NewInstancesInstancesHealthCheckCreateParamsWithContext(ctx context.Context) *InstancesInstancesHealthCheckCreateParams {
	return &InstancesInstancesHealthCheckCreateParams{
		Context: ctx,
	}
}

// NewInstancesInstancesHealthCheckCreateParamsWithHTTPClient creates a new InstancesInstancesHealthCheckCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewInstancesInstancesHealthCheckCreateParamsWithHTTPClient(client *http.Client) *InstancesInstancesHealthCheckCreateParams {
	return &InstancesInstancesHealthCheckCreateParams{
		HTTPClient: client,
	}
}

/*
InstancesInstancesHealthCheckCreateParams contains all the parameters to send to the API endpoint

	for the instances instances health check create operation.

	Typically these are written to a http.Request.
*/
type InstancesInstancesHealthCheckCreateParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the instances instances health check create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstancesInstancesHealthCheckCreateParams) WithDefaults() *InstancesInstancesHealthCheckCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the instances instances health check create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstancesInstancesHealthCheckCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the instances instances health check create params
func (o *InstancesInstancesHealthCheckCreateParams) WithTimeout(timeout time.Duration) *InstancesInstancesHealthCheckCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the instances instances health check create params
func (o *InstancesInstancesHealthCheckCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the instances instances health check create params
func (o *InstancesInstancesHealthCheckCreateParams) WithContext(ctx context.Context) *InstancesInstancesHealthCheckCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the instances instances health check create params
func (o *InstancesInstancesHealthCheckCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the instances instances health check create params
func (o *InstancesInstancesHealthCheckCreateParams) WithHTTPClient(client *http.Client) *InstancesInstancesHealthCheckCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the instances instances health check create params
func (o *InstancesInstancesHealthCheckCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the instances instances health check create params
func (o *InstancesInstancesHealthCheckCreateParams) WithID(id string) *InstancesInstancesHealthCheckCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the instances instances health check create params
func (o *InstancesInstancesHealthCheckCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InstancesInstancesHealthCheckCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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