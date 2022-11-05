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

// NewHostsHostsPartialUpdateParams creates a new HostsHostsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHostsHostsPartialUpdateParams() *HostsHostsPartialUpdateParams {
	return &HostsHostsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHostsHostsPartialUpdateParamsWithTimeout creates a new HostsHostsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewHostsHostsPartialUpdateParamsWithTimeout(timeout time.Duration) *HostsHostsPartialUpdateParams {
	return &HostsHostsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewHostsHostsPartialUpdateParamsWithContext creates a new HostsHostsPartialUpdateParams object
// with the ability to set a context for a request.
func NewHostsHostsPartialUpdateParamsWithContext(ctx context.Context) *HostsHostsPartialUpdateParams {
	return &HostsHostsPartialUpdateParams{
		Context: ctx,
	}
}

// NewHostsHostsPartialUpdateParamsWithHTTPClient creates a new HostsHostsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewHostsHostsPartialUpdateParamsWithHTTPClient(client *http.Client) *HostsHostsPartialUpdateParams {
	return &HostsHostsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
HostsHostsPartialUpdateParams contains all the parameters to send to the API endpoint

	for the hosts hosts partial update operation.

	Typically these are written to a http.Request.
*/
type HostsHostsPartialUpdateParams struct {

	// Data.
	Data HostsHostsPartialUpdateBody

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

// WithDefaults hydrates default values in the hosts hosts partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsPartialUpdateParams) WithDefaults() *HostsHostsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hosts hosts partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hosts hosts partial update params
func (o *HostsHostsPartialUpdateParams) WithTimeout(timeout time.Duration) *HostsHostsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hosts hosts partial update params
func (o *HostsHostsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hosts hosts partial update params
func (o *HostsHostsPartialUpdateParams) WithContext(ctx context.Context) *HostsHostsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hosts hosts partial update params
func (o *HostsHostsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hosts hosts partial update params
func (o *HostsHostsPartialUpdateParams) WithHTTPClient(client *http.Client) *HostsHostsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hosts hosts partial update params
func (o *HostsHostsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the hosts hosts partial update params
func (o *HostsHostsPartialUpdateParams) WithData(data HostsHostsPartialUpdateBody) *HostsHostsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the hosts hosts partial update params
func (o *HostsHostsPartialUpdateParams) SetData(data HostsHostsPartialUpdateBody) {
	o.Data = data
}

// WithID adds the id to the hosts hosts partial update params
func (o *HostsHostsPartialUpdateParams) WithID(id string) *HostsHostsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the hosts hosts partial update params
func (o *HostsHostsPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the hosts hosts partial update params
func (o *HostsHostsPartialUpdateParams) WithSearch(search *string) *HostsHostsPartialUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the hosts hosts partial update params
func (o *HostsHostsPartialUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *HostsHostsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
