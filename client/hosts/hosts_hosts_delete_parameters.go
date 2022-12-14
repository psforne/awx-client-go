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

// NewHostsHostsDeleteParams creates a new HostsHostsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHostsHostsDeleteParams() *HostsHostsDeleteParams {
	return &HostsHostsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHostsHostsDeleteParamsWithTimeout creates a new HostsHostsDeleteParams object
// with the ability to set a timeout on a request.
func NewHostsHostsDeleteParamsWithTimeout(timeout time.Duration) *HostsHostsDeleteParams {
	return &HostsHostsDeleteParams{
		timeout: timeout,
	}
}

// NewHostsHostsDeleteParamsWithContext creates a new HostsHostsDeleteParams object
// with the ability to set a context for a request.
func NewHostsHostsDeleteParamsWithContext(ctx context.Context) *HostsHostsDeleteParams {
	return &HostsHostsDeleteParams{
		Context: ctx,
	}
}

// NewHostsHostsDeleteParamsWithHTTPClient creates a new HostsHostsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewHostsHostsDeleteParamsWithHTTPClient(client *http.Client) *HostsHostsDeleteParams {
	return &HostsHostsDeleteParams{
		HTTPClient: client,
	}
}

/*
HostsHostsDeleteParams contains all the parameters to send to the API endpoint

	for the hosts hosts delete operation.

	Typically these are written to a http.Request.
*/
type HostsHostsDeleteParams struct {

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

// WithDefaults hydrates default values in the hosts hosts delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsDeleteParams) WithDefaults() *HostsHostsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hosts hosts delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hosts hosts delete params
func (o *HostsHostsDeleteParams) WithTimeout(timeout time.Duration) *HostsHostsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hosts hosts delete params
func (o *HostsHostsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hosts hosts delete params
func (o *HostsHostsDeleteParams) WithContext(ctx context.Context) *HostsHostsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hosts hosts delete params
func (o *HostsHostsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hosts hosts delete params
func (o *HostsHostsDeleteParams) WithHTTPClient(client *http.Client) *HostsHostsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hosts hosts delete params
func (o *HostsHostsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the hosts hosts delete params
func (o *HostsHostsDeleteParams) WithID(id string) *HostsHostsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the hosts hosts delete params
func (o *HostsHostsDeleteParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the hosts hosts delete params
func (o *HostsHostsDeleteParams) WithSearch(search *string) *HostsHostsDeleteParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the hosts hosts delete params
func (o *HostsHostsDeleteParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *HostsHostsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
