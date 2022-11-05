// Code generated by go-swagger; DO NOT EDIT.

package dashboard

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

// NewDashboardDashboardListParams creates a new DashboardDashboardListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDashboardDashboardListParams() *DashboardDashboardListParams {
	return &DashboardDashboardListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDashboardDashboardListParamsWithTimeout creates a new DashboardDashboardListParams object
// with the ability to set a timeout on a request.
func NewDashboardDashboardListParamsWithTimeout(timeout time.Duration) *DashboardDashboardListParams {
	return &DashboardDashboardListParams{
		timeout: timeout,
	}
}

// NewDashboardDashboardListParamsWithContext creates a new DashboardDashboardListParams object
// with the ability to set a context for a request.
func NewDashboardDashboardListParamsWithContext(ctx context.Context) *DashboardDashboardListParams {
	return &DashboardDashboardListParams{
		Context: ctx,
	}
}

// NewDashboardDashboardListParamsWithHTTPClient creates a new DashboardDashboardListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDashboardDashboardListParamsWithHTTPClient(client *http.Client) *DashboardDashboardListParams {
	return &DashboardDashboardListParams{
		HTTPClient: client,
	}
}

/*
DashboardDashboardListParams contains all the parameters to send to the API endpoint

	for the dashboard dashboard list operation.

	Typically these are written to a http.Request.
*/
type DashboardDashboardListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dashboard dashboard list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DashboardDashboardListParams) WithDefaults() *DashboardDashboardListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dashboard dashboard list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DashboardDashboardListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dashboard dashboard list params
func (o *DashboardDashboardListParams) WithTimeout(timeout time.Duration) *DashboardDashboardListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dashboard dashboard list params
func (o *DashboardDashboardListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dashboard dashboard list params
func (o *DashboardDashboardListParams) WithContext(ctx context.Context) *DashboardDashboardListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dashboard dashboard list params
func (o *DashboardDashboardListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dashboard dashboard list params
func (o *DashboardDashboardListParams) WithHTTPClient(client *http.Client) *DashboardDashboardListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dashboard dashboard list params
func (o *DashboardDashboardListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DashboardDashboardListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}