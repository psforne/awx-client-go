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

// NewSchedulesSchedulesPartialUpdateParams creates a new SchedulesSchedulesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSchedulesSchedulesPartialUpdateParams() *SchedulesSchedulesPartialUpdateParams {
	return &SchedulesSchedulesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSchedulesSchedulesPartialUpdateParamsWithTimeout creates a new SchedulesSchedulesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewSchedulesSchedulesPartialUpdateParamsWithTimeout(timeout time.Duration) *SchedulesSchedulesPartialUpdateParams {
	return &SchedulesSchedulesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewSchedulesSchedulesPartialUpdateParamsWithContext creates a new SchedulesSchedulesPartialUpdateParams object
// with the ability to set a context for a request.
func NewSchedulesSchedulesPartialUpdateParamsWithContext(ctx context.Context) *SchedulesSchedulesPartialUpdateParams {
	return &SchedulesSchedulesPartialUpdateParams{
		Context: ctx,
	}
}

// NewSchedulesSchedulesPartialUpdateParamsWithHTTPClient creates a new SchedulesSchedulesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSchedulesSchedulesPartialUpdateParamsWithHTTPClient(client *http.Client) *SchedulesSchedulesPartialUpdateParams {
	return &SchedulesSchedulesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
SchedulesSchedulesPartialUpdateParams contains all the parameters to send to the API endpoint

	for the schedules schedules partial update operation.

	Typically these are written to a http.Request.
*/
type SchedulesSchedulesPartialUpdateParams struct {

	// Data.
	Data interface{}

	// ID.
	ID string

	/* Search.

	   A search term.
	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the schedules schedules partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulesSchedulesPartialUpdateParams) WithDefaults() *SchedulesSchedulesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the schedules schedules partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulesSchedulesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the schedules schedules partial update params
func (o *SchedulesSchedulesPartialUpdateParams) WithTimeout(timeout time.Duration) *SchedulesSchedulesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schedules schedules partial update params
func (o *SchedulesSchedulesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schedules schedules partial update params
func (o *SchedulesSchedulesPartialUpdateParams) WithContext(ctx context.Context) *SchedulesSchedulesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schedules schedules partial update params
func (o *SchedulesSchedulesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schedules schedules partial update params
func (o *SchedulesSchedulesPartialUpdateParams) WithHTTPClient(client *http.Client) *SchedulesSchedulesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schedules schedules partial update params
func (o *SchedulesSchedulesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the schedules schedules partial update params
func (o *SchedulesSchedulesPartialUpdateParams) WithData(data interface{}) *SchedulesSchedulesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the schedules schedules partial update params
func (o *SchedulesSchedulesPartialUpdateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the schedules schedules partial update params
func (o *SchedulesSchedulesPartialUpdateParams) WithID(id string) *SchedulesSchedulesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the schedules schedules partial update params
func (o *SchedulesSchedulesPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the schedules schedules partial update params
func (o *SchedulesSchedulesPartialUpdateParams) WithSearch(search *string) *SchedulesSchedulesPartialUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the schedules schedules partial update params
func (o *SchedulesSchedulesPartialUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *SchedulesSchedulesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Search != nil {

		// query param search
		var qrSearch string

		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {

			if err := r.SetQueryParam("search", qSearch); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
