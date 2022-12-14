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

// NewHostsHostsVariableDataPartialUpdateParams creates a new HostsHostsVariableDataPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHostsHostsVariableDataPartialUpdateParams() *HostsHostsVariableDataPartialUpdateParams {
	return &HostsHostsVariableDataPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHostsHostsVariableDataPartialUpdateParamsWithTimeout creates a new HostsHostsVariableDataPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewHostsHostsVariableDataPartialUpdateParamsWithTimeout(timeout time.Duration) *HostsHostsVariableDataPartialUpdateParams {
	return &HostsHostsVariableDataPartialUpdateParams{
		timeout: timeout,
	}
}

// NewHostsHostsVariableDataPartialUpdateParamsWithContext creates a new HostsHostsVariableDataPartialUpdateParams object
// with the ability to set a context for a request.
func NewHostsHostsVariableDataPartialUpdateParamsWithContext(ctx context.Context) *HostsHostsVariableDataPartialUpdateParams {
	return &HostsHostsVariableDataPartialUpdateParams{
		Context: ctx,
	}
}

// NewHostsHostsVariableDataPartialUpdateParamsWithHTTPClient creates a new HostsHostsVariableDataPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewHostsHostsVariableDataPartialUpdateParamsWithHTTPClient(client *http.Client) *HostsHostsVariableDataPartialUpdateParams {
	return &HostsHostsVariableDataPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
HostsHostsVariableDataPartialUpdateParams contains all the parameters to send to the API endpoint

	for the hosts hosts variable data partial update operation.

	Typically these are written to a http.Request.
*/
type HostsHostsVariableDataPartialUpdateParams struct {

	// Data.
	Data HostsHostsVariableDataPartialUpdateBody

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

// WithDefaults hydrates default values in the hosts hosts variable data partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsVariableDataPartialUpdateParams) WithDefaults() *HostsHostsVariableDataPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hosts hosts variable data partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsVariableDataPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hosts hosts variable data partial update params
func (o *HostsHostsVariableDataPartialUpdateParams) WithTimeout(timeout time.Duration) *HostsHostsVariableDataPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hosts hosts variable data partial update params
func (o *HostsHostsVariableDataPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hosts hosts variable data partial update params
func (o *HostsHostsVariableDataPartialUpdateParams) WithContext(ctx context.Context) *HostsHostsVariableDataPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hosts hosts variable data partial update params
func (o *HostsHostsVariableDataPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hosts hosts variable data partial update params
func (o *HostsHostsVariableDataPartialUpdateParams) WithHTTPClient(client *http.Client) *HostsHostsVariableDataPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hosts hosts variable data partial update params
func (o *HostsHostsVariableDataPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the hosts hosts variable data partial update params
func (o *HostsHostsVariableDataPartialUpdateParams) WithData(data HostsHostsVariableDataPartialUpdateBody) *HostsHostsVariableDataPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the hosts hosts variable data partial update params
func (o *HostsHostsVariableDataPartialUpdateParams) SetData(data HostsHostsVariableDataPartialUpdateBody) {
	o.Data = data
}

// WithID adds the id to the hosts hosts variable data partial update params
func (o *HostsHostsVariableDataPartialUpdateParams) WithID(id string) *HostsHostsVariableDataPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the hosts hosts variable data partial update params
func (o *HostsHostsVariableDataPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the hosts hosts variable data partial update params
func (o *HostsHostsVariableDataPartialUpdateParams) WithSearch(search *string) *HostsHostsVariableDataPartialUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the hosts hosts variable data partial update params
func (o *HostsHostsVariableDataPartialUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *HostsHostsVariableDataPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
