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

// NewHostsHostsCreateParams creates a new HostsHostsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHostsHostsCreateParams() *HostsHostsCreateParams {
	return &HostsHostsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHostsHostsCreateParamsWithTimeout creates a new HostsHostsCreateParams object
// with the ability to set a timeout on a request.
func NewHostsHostsCreateParamsWithTimeout(timeout time.Duration) *HostsHostsCreateParams {
	return &HostsHostsCreateParams{
		timeout: timeout,
	}
}

// NewHostsHostsCreateParamsWithContext creates a new HostsHostsCreateParams object
// with the ability to set a context for a request.
func NewHostsHostsCreateParamsWithContext(ctx context.Context) *HostsHostsCreateParams {
	return &HostsHostsCreateParams{
		Context: ctx,
	}
}

// NewHostsHostsCreateParamsWithHTTPClient creates a new HostsHostsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewHostsHostsCreateParamsWithHTTPClient(client *http.Client) *HostsHostsCreateParams {
	return &HostsHostsCreateParams{
		HTTPClient: client,
	}
}

/*
HostsHostsCreateParams contains all the parameters to send to the API endpoint

	for the hosts hosts create operation.

	Typically these are written to a http.Request.
*/
type HostsHostsCreateParams struct {

	// Data.
	Data HostsHostsCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hosts hosts create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsCreateParams) WithDefaults() *HostsHostsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hosts hosts create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hosts hosts create params
func (o *HostsHostsCreateParams) WithTimeout(timeout time.Duration) *HostsHostsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hosts hosts create params
func (o *HostsHostsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hosts hosts create params
func (o *HostsHostsCreateParams) WithContext(ctx context.Context) *HostsHostsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hosts hosts create params
func (o *HostsHostsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hosts hosts create params
func (o *HostsHostsCreateParams) WithHTTPClient(client *http.Client) *HostsHostsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hosts hosts create params
func (o *HostsHostsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the hosts hosts create params
func (o *HostsHostsCreateParams) WithData(data HostsHostsCreateBody) *HostsHostsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the hosts hosts create params
func (o *HostsHostsCreateParams) SetData(data HostsHostsCreateBody) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *HostsHostsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
