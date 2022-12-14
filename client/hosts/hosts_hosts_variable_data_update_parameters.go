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

// NewHostsHostsVariableDataUpdateParams creates a new HostsHostsVariableDataUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHostsHostsVariableDataUpdateParams() *HostsHostsVariableDataUpdateParams {
	return &HostsHostsVariableDataUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHostsHostsVariableDataUpdateParamsWithTimeout creates a new HostsHostsVariableDataUpdateParams object
// with the ability to set a timeout on a request.
func NewHostsHostsVariableDataUpdateParamsWithTimeout(timeout time.Duration) *HostsHostsVariableDataUpdateParams {
	return &HostsHostsVariableDataUpdateParams{
		timeout: timeout,
	}
}

// NewHostsHostsVariableDataUpdateParamsWithContext creates a new HostsHostsVariableDataUpdateParams object
// with the ability to set a context for a request.
func NewHostsHostsVariableDataUpdateParamsWithContext(ctx context.Context) *HostsHostsVariableDataUpdateParams {
	return &HostsHostsVariableDataUpdateParams{
		Context: ctx,
	}
}

// NewHostsHostsVariableDataUpdateParamsWithHTTPClient creates a new HostsHostsVariableDataUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewHostsHostsVariableDataUpdateParamsWithHTTPClient(client *http.Client) *HostsHostsVariableDataUpdateParams {
	return &HostsHostsVariableDataUpdateParams{
		HTTPClient: client,
	}
}

/*
HostsHostsVariableDataUpdateParams contains all the parameters to send to the API endpoint

	for the hosts hosts variable data update operation.

	Typically these are written to a http.Request.
*/
type HostsHostsVariableDataUpdateParams struct {

	// Data.
	Data HostsHostsVariableDataUpdateBody

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

// WithDefaults hydrates default values in the hosts hosts variable data update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsVariableDataUpdateParams) WithDefaults() *HostsHostsVariableDataUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hosts hosts variable data update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsVariableDataUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hosts hosts variable data update params
func (o *HostsHostsVariableDataUpdateParams) WithTimeout(timeout time.Duration) *HostsHostsVariableDataUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hosts hosts variable data update params
func (o *HostsHostsVariableDataUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hosts hosts variable data update params
func (o *HostsHostsVariableDataUpdateParams) WithContext(ctx context.Context) *HostsHostsVariableDataUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hosts hosts variable data update params
func (o *HostsHostsVariableDataUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hosts hosts variable data update params
func (o *HostsHostsVariableDataUpdateParams) WithHTTPClient(client *http.Client) *HostsHostsVariableDataUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hosts hosts variable data update params
func (o *HostsHostsVariableDataUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the hosts hosts variable data update params
func (o *HostsHostsVariableDataUpdateParams) WithData(data HostsHostsVariableDataUpdateBody) *HostsHostsVariableDataUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the hosts hosts variable data update params
func (o *HostsHostsVariableDataUpdateParams) SetData(data HostsHostsVariableDataUpdateBody) {
	o.Data = data
}

// WithID adds the id to the hosts hosts variable data update params
func (o *HostsHostsVariableDataUpdateParams) WithID(id string) *HostsHostsVariableDataUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the hosts hosts variable data update params
func (o *HostsHostsVariableDataUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the hosts hosts variable data update params
func (o *HostsHostsVariableDataUpdateParams) WithSearch(search *string) *HostsHostsVariableDataUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the hosts hosts variable data update params
func (o *HostsHostsVariableDataUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *HostsHostsVariableDataUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
