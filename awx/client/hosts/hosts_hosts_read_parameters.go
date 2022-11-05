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

// NewHostsHostsReadParams creates a new HostsHostsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHostsHostsReadParams() *HostsHostsReadParams {
	return &HostsHostsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHostsHostsReadParamsWithTimeout creates a new HostsHostsReadParams object
// with the ability to set a timeout on a request.
func NewHostsHostsReadParamsWithTimeout(timeout time.Duration) *HostsHostsReadParams {
	return &HostsHostsReadParams{
		timeout: timeout,
	}
}

// NewHostsHostsReadParamsWithContext creates a new HostsHostsReadParams object
// with the ability to set a context for a request.
func NewHostsHostsReadParamsWithContext(ctx context.Context) *HostsHostsReadParams {
	return &HostsHostsReadParams{
		Context: ctx,
	}
}

// NewHostsHostsReadParamsWithHTTPClient creates a new HostsHostsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewHostsHostsReadParamsWithHTTPClient(client *http.Client) *HostsHostsReadParams {
	return &HostsHostsReadParams{
		HTTPClient: client,
	}
}

/*
HostsHostsReadParams contains all the parameters to send to the API endpoint

	for the hosts hosts read operation.

	Typically these are written to a http.Request.
*/
type HostsHostsReadParams struct {

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

// WithDefaults hydrates default values in the hosts hosts read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsReadParams) WithDefaults() *HostsHostsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hosts hosts read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hosts hosts read params
func (o *HostsHostsReadParams) WithTimeout(timeout time.Duration) *HostsHostsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hosts hosts read params
func (o *HostsHostsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hosts hosts read params
func (o *HostsHostsReadParams) WithContext(ctx context.Context) *HostsHostsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hosts hosts read params
func (o *HostsHostsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hosts hosts read params
func (o *HostsHostsReadParams) WithHTTPClient(client *http.Client) *HostsHostsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hosts hosts read params
func (o *HostsHostsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the hosts hosts read params
func (o *HostsHostsReadParams) WithID(id string) *HostsHostsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the hosts hosts read params
func (o *HostsHostsReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the hosts hosts read params
func (o *HostsHostsReadParams) WithSearch(search *string) *HostsHostsReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the hosts hosts read params
func (o *HostsHostsReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *HostsHostsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
