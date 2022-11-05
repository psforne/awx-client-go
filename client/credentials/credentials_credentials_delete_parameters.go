// Code generated by go-swagger; DO NOT EDIT.

package credentials

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

// NewCredentialsCredentialsDeleteParams creates a new CredentialsCredentialsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCredentialsCredentialsDeleteParams() *CredentialsCredentialsDeleteParams {
	return &CredentialsCredentialsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCredentialsCredentialsDeleteParamsWithTimeout creates a new CredentialsCredentialsDeleteParams object
// with the ability to set a timeout on a request.
func NewCredentialsCredentialsDeleteParamsWithTimeout(timeout time.Duration) *CredentialsCredentialsDeleteParams {
	return &CredentialsCredentialsDeleteParams{
		timeout: timeout,
	}
}

// NewCredentialsCredentialsDeleteParamsWithContext creates a new CredentialsCredentialsDeleteParams object
// with the ability to set a context for a request.
func NewCredentialsCredentialsDeleteParamsWithContext(ctx context.Context) *CredentialsCredentialsDeleteParams {
	return &CredentialsCredentialsDeleteParams{
		Context: ctx,
	}
}

// NewCredentialsCredentialsDeleteParamsWithHTTPClient creates a new CredentialsCredentialsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCredentialsCredentialsDeleteParamsWithHTTPClient(client *http.Client) *CredentialsCredentialsDeleteParams {
	return &CredentialsCredentialsDeleteParams{
		HTTPClient: client,
	}
}

/*
CredentialsCredentialsDeleteParams contains all the parameters to send to the API endpoint

	for the credentials credentials delete operation.

	Typically these are written to a http.Request.
*/
type CredentialsCredentialsDeleteParams struct {

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

// WithDefaults hydrates default values in the credentials credentials delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CredentialsCredentialsDeleteParams) WithDefaults() *CredentialsCredentialsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the credentials credentials delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CredentialsCredentialsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the credentials credentials delete params
func (o *CredentialsCredentialsDeleteParams) WithTimeout(timeout time.Duration) *CredentialsCredentialsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the credentials credentials delete params
func (o *CredentialsCredentialsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the credentials credentials delete params
func (o *CredentialsCredentialsDeleteParams) WithContext(ctx context.Context) *CredentialsCredentialsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the credentials credentials delete params
func (o *CredentialsCredentialsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the credentials credentials delete params
func (o *CredentialsCredentialsDeleteParams) WithHTTPClient(client *http.Client) *CredentialsCredentialsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the credentials credentials delete params
func (o *CredentialsCredentialsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the credentials credentials delete params
func (o *CredentialsCredentialsDeleteParams) WithID(id string) *CredentialsCredentialsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the credentials credentials delete params
func (o *CredentialsCredentialsDeleteParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the credentials credentials delete params
func (o *CredentialsCredentialsDeleteParams) WithSearch(search *string) *CredentialsCredentialsDeleteParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the credentials credentials delete params
func (o *CredentialsCredentialsDeleteParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *CredentialsCredentialsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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