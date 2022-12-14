// Code generated by go-swagger; DO NOT EDIT.

package instance_groups

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

// NewInstanceGroupsInstanceGroupsCreateParams creates a new InstanceGroupsInstanceGroupsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInstanceGroupsInstanceGroupsCreateParams() *InstanceGroupsInstanceGroupsCreateParams {
	return &InstanceGroupsInstanceGroupsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInstanceGroupsInstanceGroupsCreateParamsWithTimeout creates a new InstanceGroupsInstanceGroupsCreateParams object
// with the ability to set a timeout on a request.
func NewInstanceGroupsInstanceGroupsCreateParamsWithTimeout(timeout time.Duration) *InstanceGroupsInstanceGroupsCreateParams {
	return &InstanceGroupsInstanceGroupsCreateParams{
		timeout: timeout,
	}
}

// NewInstanceGroupsInstanceGroupsCreateParamsWithContext creates a new InstanceGroupsInstanceGroupsCreateParams object
// with the ability to set a context for a request.
func NewInstanceGroupsInstanceGroupsCreateParamsWithContext(ctx context.Context) *InstanceGroupsInstanceGroupsCreateParams {
	return &InstanceGroupsInstanceGroupsCreateParams{
		Context: ctx,
	}
}

// NewInstanceGroupsInstanceGroupsCreateParamsWithHTTPClient creates a new InstanceGroupsInstanceGroupsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewInstanceGroupsInstanceGroupsCreateParamsWithHTTPClient(client *http.Client) *InstanceGroupsInstanceGroupsCreateParams {
	return &InstanceGroupsInstanceGroupsCreateParams{
		HTTPClient: client,
	}
}

/*
InstanceGroupsInstanceGroupsCreateParams contains all the parameters to send to the API endpoint

	for the instance groups instance groups create operation.

	Typically these are written to a http.Request.
*/
type InstanceGroupsInstanceGroupsCreateParams struct {

	// Data.
	Data InstanceGroupsInstanceGroupsCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the instance groups instance groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstanceGroupsInstanceGroupsCreateParams) WithDefaults() *InstanceGroupsInstanceGroupsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the instance groups instance groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstanceGroupsInstanceGroupsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the instance groups instance groups create params
func (o *InstanceGroupsInstanceGroupsCreateParams) WithTimeout(timeout time.Duration) *InstanceGroupsInstanceGroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the instance groups instance groups create params
func (o *InstanceGroupsInstanceGroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the instance groups instance groups create params
func (o *InstanceGroupsInstanceGroupsCreateParams) WithContext(ctx context.Context) *InstanceGroupsInstanceGroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the instance groups instance groups create params
func (o *InstanceGroupsInstanceGroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the instance groups instance groups create params
func (o *InstanceGroupsInstanceGroupsCreateParams) WithHTTPClient(client *http.Client) *InstanceGroupsInstanceGroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the instance groups instance groups create params
func (o *InstanceGroupsInstanceGroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the instance groups instance groups create params
func (o *InstanceGroupsInstanceGroupsCreateParams) WithData(data InstanceGroupsInstanceGroupsCreateBody) *InstanceGroupsInstanceGroupsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the instance groups instance groups create params
func (o *InstanceGroupsInstanceGroupsCreateParams) SetData(data InstanceGroupsInstanceGroupsCreateBody) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *InstanceGroupsInstanceGroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
