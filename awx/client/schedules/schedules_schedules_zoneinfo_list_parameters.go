// Code generated by go-swagger; DO NOT EDIT.

package schedules

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

// NewSchedulesSchedulesZoneinfoListParams creates a new SchedulesSchedulesZoneinfoListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSchedulesSchedulesZoneinfoListParams() *SchedulesSchedulesZoneinfoListParams {
	return &SchedulesSchedulesZoneinfoListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSchedulesSchedulesZoneinfoListParamsWithTimeout creates a new SchedulesSchedulesZoneinfoListParams object
// with the ability to set a timeout on a request.
func NewSchedulesSchedulesZoneinfoListParamsWithTimeout(timeout time.Duration) *SchedulesSchedulesZoneinfoListParams {
	return &SchedulesSchedulesZoneinfoListParams{
		timeout: timeout,
	}
}

// NewSchedulesSchedulesZoneinfoListParamsWithContext creates a new SchedulesSchedulesZoneinfoListParams object
// with the ability to set a context for a request.
func NewSchedulesSchedulesZoneinfoListParamsWithContext(ctx context.Context) *SchedulesSchedulesZoneinfoListParams {
	return &SchedulesSchedulesZoneinfoListParams{
		Context: ctx,
	}
}

// NewSchedulesSchedulesZoneinfoListParamsWithHTTPClient creates a new SchedulesSchedulesZoneinfoListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSchedulesSchedulesZoneinfoListParamsWithHTTPClient(client *http.Client) *SchedulesSchedulesZoneinfoListParams {
	return &SchedulesSchedulesZoneinfoListParams{
		HTTPClient: client,
	}
}

/*
SchedulesSchedulesZoneinfoListParams contains all the parameters to send to the API endpoint

	for the schedules schedules zoneinfo list operation.

	Typically these are written to a http.Request.
*/
type SchedulesSchedulesZoneinfoListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the schedules schedules zoneinfo list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulesSchedulesZoneinfoListParams) WithDefaults() *SchedulesSchedulesZoneinfoListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the schedules schedules zoneinfo list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulesSchedulesZoneinfoListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the schedules schedules zoneinfo list params
func (o *SchedulesSchedulesZoneinfoListParams) WithTimeout(timeout time.Duration) *SchedulesSchedulesZoneinfoListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schedules schedules zoneinfo list params
func (o *SchedulesSchedulesZoneinfoListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schedules schedules zoneinfo list params
func (o *SchedulesSchedulesZoneinfoListParams) WithContext(ctx context.Context) *SchedulesSchedulesZoneinfoListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schedules schedules zoneinfo list params
func (o *SchedulesSchedulesZoneinfoListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schedules schedules zoneinfo list params
func (o *SchedulesSchedulesZoneinfoListParams) WithHTTPClient(client *http.Client) *SchedulesSchedulesZoneinfoListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schedules schedules zoneinfo list params
func (o *SchedulesSchedulesZoneinfoListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SchedulesSchedulesZoneinfoListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
