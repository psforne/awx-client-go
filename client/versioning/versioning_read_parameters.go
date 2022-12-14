// Code generated by go-swagger; DO NOT EDIT.

package versioning

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

// NewVersioningReadParams creates a new VersioningReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVersioningReadParams() *VersioningReadParams {
	return &VersioningReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVersioningReadParamsWithTimeout creates a new VersioningReadParams object
// with the ability to set a timeout on a request.
func NewVersioningReadParamsWithTimeout(timeout time.Duration) *VersioningReadParams {
	return &VersioningReadParams{
		timeout: timeout,
	}
}

// NewVersioningReadParamsWithContext creates a new VersioningReadParams object
// with the ability to set a context for a request.
func NewVersioningReadParamsWithContext(ctx context.Context) *VersioningReadParams {
	return &VersioningReadParams{
		Context: ctx,
	}
}

// NewVersioningReadParamsWithHTTPClient creates a new VersioningReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewVersioningReadParamsWithHTTPClient(client *http.Client) *VersioningReadParams {
	return &VersioningReadParams{
		HTTPClient: client,
	}
}

/*
VersioningReadParams contains all the parameters to send to the API endpoint

	for the versioning read operation.

	Typically these are written to a http.Request.
*/
type VersioningReadParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the versioning read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VersioningReadParams) WithDefaults() *VersioningReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the versioning read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VersioningReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the versioning read params
func (o *VersioningReadParams) WithTimeout(timeout time.Duration) *VersioningReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the versioning read params
func (o *VersioningReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the versioning read params
func (o *VersioningReadParams) WithContext(ctx context.Context) *VersioningReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the versioning read params
func (o *VersioningReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the versioning read params
func (o *VersioningReadParams) WithHTTPClient(client *http.Client) *VersioningReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the versioning read params
func (o *VersioningReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *VersioningReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
