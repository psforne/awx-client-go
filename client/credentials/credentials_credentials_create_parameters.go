// Code generated by go-swagger; DO NOT EDIT.

package credentials

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

// NewCredentialsCredentialsCreateParams creates a new CredentialsCredentialsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCredentialsCredentialsCreateParams() *CredentialsCredentialsCreateParams {
	return &CredentialsCredentialsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCredentialsCredentialsCreateParamsWithTimeout creates a new CredentialsCredentialsCreateParams object
// with the ability to set a timeout on a request.
func NewCredentialsCredentialsCreateParamsWithTimeout(timeout time.Duration) *CredentialsCredentialsCreateParams {
	return &CredentialsCredentialsCreateParams{
		timeout: timeout,
	}
}

// NewCredentialsCredentialsCreateParamsWithContext creates a new CredentialsCredentialsCreateParams object
// with the ability to set a context for a request.
func NewCredentialsCredentialsCreateParamsWithContext(ctx context.Context) *CredentialsCredentialsCreateParams {
	return &CredentialsCredentialsCreateParams{
		Context: ctx,
	}
}

// NewCredentialsCredentialsCreateParamsWithHTTPClient creates a new CredentialsCredentialsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCredentialsCredentialsCreateParamsWithHTTPClient(client *http.Client) *CredentialsCredentialsCreateParams {
	return &CredentialsCredentialsCreateParams{
		HTTPClient: client,
	}
}

/*
CredentialsCredentialsCreateParams contains all the parameters to send to the API endpoint

	for the credentials credentials create operation.

	Typically these are written to a http.Request.
*/
type CredentialsCredentialsCreateParams struct {

	// Data.
	Data interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the credentials credentials create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CredentialsCredentialsCreateParams) WithDefaults() *CredentialsCredentialsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the credentials credentials create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CredentialsCredentialsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the credentials credentials create params
func (o *CredentialsCredentialsCreateParams) WithTimeout(timeout time.Duration) *CredentialsCredentialsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the credentials credentials create params
func (o *CredentialsCredentialsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the credentials credentials create params
func (o *CredentialsCredentialsCreateParams) WithContext(ctx context.Context) *CredentialsCredentialsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the credentials credentials create params
func (o *CredentialsCredentialsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the credentials credentials create params
func (o *CredentialsCredentialsCreateParams) WithHTTPClient(client *http.Client) *CredentialsCredentialsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the credentials credentials create params
func (o *CredentialsCredentialsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the credentials credentials create params
func (o *CredentialsCredentialsCreateParams) WithData(data interface{}) *CredentialsCredentialsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the credentials credentials create params
func (o *CredentialsCredentialsCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CredentialsCredentialsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
