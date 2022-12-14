// Code generated by go-swagger; DO NOT EDIT.

package ad_hoc_commands

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

// NewAdHocCommandsAdHocCommandsCancelCreateParams creates a new AdHocCommandsAdHocCommandsCancelCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdHocCommandsAdHocCommandsCancelCreateParams() *AdHocCommandsAdHocCommandsCancelCreateParams {
	return &AdHocCommandsAdHocCommandsCancelCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdHocCommandsAdHocCommandsCancelCreateParamsWithTimeout creates a new AdHocCommandsAdHocCommandsCancelCreateParams object
// with the ability to set a timeout on a request.
func NewAdHocCommandsAdHocCommandsCancelCreateParamsWithTimeout(timeout time.Duration) *AdHocCommandsAdHocCommandsCancelCreateParams {
	return &AdHocCommandsAdHocCommandsCancelCreateParams{
		timeout: timeout,
	}
}

// NewAdHocCommandsAdHocCommandsCancelCreateParamsWithContext creates a new AdHocCommandsAdHocCommandsCancelCreateParams object
// with the ability to set a context for a request.
func NewAdHocCommandsAdHocCommandsCancelCreateParamsWithContext(ctx context.Context) *AdHocCommandsAdHocCommandsCancelCreateParams {
	return &AdHocCommandsAdHocCommandsCancelCreateParams{
		Context: ctx,
	}
}

// NewAdHocCommandsAdHocCommandsCancelCreateParamsWithHTTPClient creates a new AdHocCommandsAdHocCommandsCancelCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdHocCommandsAdHocCommandsCancelCreateParamsWithHTTPClient(client *http.Client) *AdHocCommandsAdHocCommandsCancelCreateParams {
	return &AdHocCommandsAdHocCommandsCancelCreateParams{
		HTTPClient: client,
	}
}

/*
AdHocCommandsAdHocCommandsCancelCreateParams contains all the parameters to send to the API endpoint

	for the ad hoc commands ad hoc commands cancel create operation.

	Typically these are written to a http.Request.
*/
type AdHocCommandsAdHocCommandsCancelCreateParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ad hoc commands ad hoc commands cancel create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdHocCommandsAdHocCommandsCancelCreateParams) WithDefaults() *AdHocCommandsAdHocCommandsCancelCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ad hoc commands ad hoc commands cancel create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdHocCommandsAdHocCommandsCancelCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ad hoc commands ad hoc commands cancel create params
func (o *AdHocCommandsAdHocCommandsCancelCreateParams) WithTimeout(timeout time.Duration) *AdHocCommandsAdHocCommandsCancelCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ad hoc commands ad hoc commands cancel create params
func (o *AdHocCommandsAdHocCommandsCancelCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ad hoc commands ad hoc commands cancel create params
func (o *AdHocCommandsAdHocCommandsCancelCreateParams) WithContext(ctx context.Context) *AdHocCommandsAdHocCommandsCancelCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ad hoc commands ad hoc commands cancel create params
func (o *AdHocCommandsAdHocCommandsCancelCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ad hoc commands ad hoc commands cancel create params
func (o *AdHocCommandsAdHocCommandsCancelCreateParams) WithHTTPClient(client *http.Client) *AdHocCommandsAdHocCommandsCancelCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ad hoc commands ad hoc commands cancel create params
func (o *AdHocCommandsAdHocCommandsCancelCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ad hoc commands ad hoc commands cancel create params
func (o *AdHocCommandsAdHocCommandsCancelCreateParams) WithID(id string) *AdHocCommandsAdHocCommandsCancelCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ad hoc commands ad hoc commands cancel create params
func (o *AdHocCommandsAdHocCommandsCancelCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AdHocCommandsAdHocCommandsCancelCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
