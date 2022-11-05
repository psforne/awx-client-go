// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

// NewAuthenticationTokensDeleteParams creates a new AuthenticationTokensDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthenticationTokensDeleteParams() *AuthenticationTokensDeleteParams {
	return &AuthenticationTokensDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthenticationTokensDeleteParamsWithTimeout creates a new AuthenticationTokensDeleteParams object
// with the ability to set a timeout on a request.
func NewAuthenticationTokensDeleteParamsWithTimeout(timeout time.Duration) *AuthenticationTokensDeleteParams {
	return &AuthenticationTokensDeleteParams{
		timeout: timeout,
	}
}

// NewAuthenticationTokensDeleteParamsWithContext creates a new AuthenticationTokensDeleteParams object
// with the ability to set a context for a request.
func NewAuthenticationTokensDeleteParamsWithContext(ctx context.Context) *AuthenticationTokensDeleteParams {
	return &AuthenticationTokensDeleteParams{
		Context: ctx,
	}
}

// NewAuthenticationTokensDeleteParamsWithHTTPClient creates a new AuthenticationTokensDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthenticationTokensDeleteParamsWithHTTPClient(client *http.Client) *AuthenticationTokensDeleteParams {
	return &AuthenticationTokensDeleteParams{
		HTTPClient: client,
	}
}

/*
AuthenticationTokensDeleteParams contains all the parameters to send to the API endpoint

	for the authentication tokens delete operation.

	Typically these are written to a http.Request.
*/
type AuthenticationTokensDeleteParams struct {

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

// WithDefaults hydrates default values in the authentication tokens delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticationTokensDeleteParams) WithDefaults() *AuthenticationTokensDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the authentication tokens delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticationTokensDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the authentication tokens delete params
func (o *AuthenticationTokensDeleteParams) WithTimeout(timeout time.Duration) *AuthenticationTokensDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authentication tokens delete params
func (o *AuthenticationTokensDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authentication tokens delete params
func (o *AuthenticationTokensDeleteParams) WithContext(ctx context.Context) *AuthenticationTokensDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authentication tokens delete params
func (o *AuthenticationTokensDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authentication tokens delete params
func (o *AuthenticationTokensDeleteParams) WithHTTPClient(client *http.Client) *AuthenticationTokensDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authentication tokens delete params
func (o *AuthenticationTokensDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the authentication tokens delete params
func (o *AuthenticationTokensDeleteParams) WithID(id string) *AuthenticationTokensDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the authentication tokens delete params
func (o *AuthenticationTokensDeleteParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the authentication tokens delete params
func (o *AuthenticationTokensDeleteParams) WithSearch(search *string) *AuthenticationTokensDeleteParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the authentication tokens delete params
func (o *AuthenticationTokensDeleteParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *AuthenticationTokensDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
