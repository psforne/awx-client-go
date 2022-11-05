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

// NewHostsHostsVariableDataReadParams creates a new HostsHostsVariableDataReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHostsHostsVariableDataReadParams() *HostsHostsVariableDataReadParams {
	return &HostsHostsVariableDataReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHostsHostsVariableDataReadParamsWithTimeout creates a new HostsHostsVariableDataReadParams object
// with the ability to set a timeout on a request.
func NewHostsHostsVariableDataReadParamsWithTimeout(timeout time.Duration) *HostsHostsVariableDataReadParams {
	return &HostsHostsVariableDataReadParams{
		timeout: timeout,
	}
}

// NewHostsHostsVariableDataReadParamsWithContext creates a new HostsHostsVariableDataReadParams object
// with the ability to set a context for a request.
func NewHostsHostsVariableDataReadParamsWithContext(ctx context.Context) *HostsHostsVariableDataReadParams {
	return &HostsHostsVariableDataReadParams{
		Context: ctx,
	}
}

// NewHostsHostsVariableDataReadParamsWithHTTPClient creates a new HostsHostsVariableDataReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewHostsHostsVariableDataReadParamsWithHTTPClient(client *http.Client) *HostsHostsVariableDataReadParams {
	return &HostsHostsVariableDataReadParams{
		HTTPClient: client,
	}
}

/*
HostsHostsVariableDataReadParams contains all the parameters to send to the API endpoint

	for the hosts hosts variable data read operation.

	Typically these are written to a http.Request.
*/
type HostsHostsVariableDataReadParams struct {

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

// WithDefaults hydrates default values in the hosts hosts variable data read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsVariableDataReadParams) WithDefaults() *HostsHostsVariableDataReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hosts hosts variable data read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsVariableDataReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hosts hosts variable data read params
func (o *HostsHostsVariableDataReadParams) WithTimeout(timeout time.Duration) *HostsHostsVariableDataReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hosts hosts variable data read params
func (o *HostsHostsVariableDataReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hosts hosts variable data read params
func (o *HostsHostsVariableDataReadParams) WithContext(ctx context.Context) *HostsHostsVariableDataReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hosts hosts variable data read params
func (o *HostsHostsVariableDataReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hosts hosts variable data read params
func (o *HostsHostsVariableDataReadParams) WithHTTPClient(client *http.Client) *HostsHostsVariableDataReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hosts hosts variable data read params
func (o *HostsHostsVariableDataReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the hosts hosts variable data read params
func (o *HostsHostsVariableDataReadParams) WithID(id string) *HostsHostsVariableDataReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the hosts hosts variable data read params
func (o *HostsHostsVariableDataReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the hosts hosts variable data read params
func (o *HostsHostsVariableDataReadParams) WithSearch(search *string) *HostsHostsVariableDataReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the hosts hosts variable data read params
func (o *HostsHostsVariableDataReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *HostsHostsVariableDataReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
