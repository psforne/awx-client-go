// Code generated by go-swagger; DO NOT EDIT.

package hosts

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

// NewHostsHostsAdHocCommandsCreateParams creates a new HostsHostsAdHocCommandsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHostsHostsAdHocCommandsCreateParams() *HostsHostsAdHocCommandsCreateParams {
	return &HostsHostsAdHocCommandsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHostsHostsAdHocCommandsCreateParamsWithTimeout creates a new HostsHostsAdHocCommandsCreateParams object
// with the ability to set a timeout on a request.
func NewHostsHostsAdHocCommandsCreateParamsWithTimeout(timeout time.Duration) *HostsHostsAdHocCommandsCreateParams {
	return &HostsHostsAdHocCommandsCreateParams{
		timeout: timeout,
	}
}

// NewHostsHostsAdHocCommandsCreateParamsWithContext creates a new HostsHostsAdHocCommandsCreateParams object
// with the ability to set a context for a request.
func NewHostsHostsAdHocCommandsCreateParamsWithContext(ctx context.Context) *HostsHostsAdHocCommandsCreateParams {
	return &HostsHostsAdHocCommandsCreateParams{
		Context: ctx,
	}
}

// NewHostsHostsAdHocCommandsCreateParamsWithHTTPClient creates a new HostsHostsAdHocCommandsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewHostsHostsAdHocCommandsCreateParamsWithHTTPClient(client *http.Client) *HostsHostsAdHocCommandsCreateParams {
	return &HostsHostsAdHocCommandsCreateParams{
		HTTPClient: client,
	}
}

/*
HostsHostsAdHocCommandsCreateParams contains all the parameters to send to the API endpoint

	for the hosts hosts ad hoc commands create operation.

	Typically these are written to a http.Request.
*/
type HostsHostsAdHocCommandsCreateParams struct {

	// Data.
	Data HostsHostsAdHocCommandsCreateBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hosts hosts ad hoc commands create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsAdHocCommandsCreateParams) WithDefaults() *HostsHostsAdHocCommandsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hosts hosts ad hoc commands create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsAdHocCommandsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hosts hosts ad hoc commands create params
func (o *HostsHostsAdHocCommandsCreateParams) WithTimeout(timeout time.Duration) *HostsHostsAdHocCommandsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hosts hosts ad hoc commands create params
func (o *HostsHostsAdHocCommandsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hosts hosts ad hoc commands create params
func (o *HostsHostsAdHocCommandsCreateParams) WithContext(ctx context.Context) *HostsHostsAdHocCommandsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hosts hosts ad hoc commands create params
func (o *HostsHostsAdHocCommandsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hosts hosts ad hoc commands create params
func (o *HostsHostsAdHocCommandsCreateParams) WithHTTPClient(client *http.Client) *HostsHostsAdHocCommandsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hosts hosts ad hoc commands create params
func (o *HostsHostsAdHocCommandsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the hosts hosts ad hoc commands create params
func (o *HostsHostsAdHocCommandsCreateParams) WithData(data HostsHostsAdHocCommandsCreateBody) *HostsHostsAdHocCommandsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the hosts hosts ad hoc commands create params
func (o *HostsHostsAdHocCommandsCreateParams) SetData(data HostsHostsAdHocCommandsCreateBody) {
	o.Data = data
}

// WithID adds the id to the hosts hosts ad hoc commands create params
func (o *HostsHostsAdHocCommandsCreateParams) WithID(id string) *HostsHostsAdHocCommandsCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the hosts hosts ad hoc commands create params
func (o *HostsHostsAdHocCommandsCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *HostsHostsAdHocCommandsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
