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

// NewAuthenticationApplicationsDelete0Params creates a new AuthenticationApplicationsDelete0Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthenticationApplicationsDelete0Params() *AuthenticationApplicationsDelete0Params {
	return &AuthenticationApplicationsDelete0Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthenticationApplicationsDelete0ParamsWithTimeout creates a new AuthenticationApplicationsDelete0Params object
// with the ability to set a timeout on a request.
func NewAuthenticationApplicationsDelete0ParamsWithTimeout(timeout time.Duration) *AuthenticationApplicationsDelete0Params {
	return &AuthenticationApplicationsDelete0Params{
		timeout: timeout,
	}
}

// NewAuthenticationApplicationsDelete0ParamsWithContext creates a new AuthenticationApplicationsDelete0Params object
// with the ability to set a context for a request.
func NewAuthenticationApplicationsDelete0ParamsWithContext(ctx context.Context) *AuthenticationApplicationsDelete0Params {
	return &AuthenticationApplicationsDelete0Params{
		Context: ctx,
	}
}

// NewAuthenticationApplicationsDelete0ParamsWithHTTPClient creates a new AuthenticationApplicationsDelete0Params object
// with the ability to set a custom HTTPClient for a request.
func NewAuthenticationApplicationsDelete0ParamsWithHTTPClient(client *http.Client) *AuthenticationApplicationsDelete0Params {
	return &AuthenticationApplicationsDelete0Params{
		HTTPClient: client,
	}
}

/*
AuthenticationApplicationsDelete0Params contains all the parameters to send to the API endpoint

	for the authentication applications delete 0 operation.

	Typically these are written to a http.Request.
*/
type AuthenticationApplicationsDelete0Params struct {

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

// WithDefaults hydrates default values in the authentication applications delete 0 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticationApplicationsDelete0Params) WithDefaults() *AuthenticationApplicationsDelete0Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the authentication applications delete 0 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticationApplicationsDelete0Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the authentication applications delete 0 params
func (o *AuthenticationApplicationsDelete0Params) WithTimeout(timeout time.Duration) *AuthenticationApplicationsDelete0Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authentication applications delete 0 params
func (o *AuthenticationApplicationsDelete0Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authentication applications delete 0 params
func (o *AuthenticationApplicationsDelete0Params) WithContext(ctx context.Context) *AuthenticationApplicationsDelete0Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authentication applications delete 0 params
func (o *AuthenticationApplicationsDelete0Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authentication applications delete 0 params
func (o *AuthenticationApplicationsDelete0Params) WithHTTPClient(client *http.Client) *AuthenticationApplicationsDelete0Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authentication applications delete 0 params
func (o *AuthenticationApplicationsDelete0Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the authentication applications delete 0 params
func (o *AuthenticationApplicationsDelete0Params) WithID(id string) *AuthenticationApplicationsDelete0Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the authentication applications delete 0 params
func (o *AuthenticationApplicationsDelete0Params) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the authentication applications delete 0 params
func (o *AuthenticationApplicationsDelete0Params) WithSearch(search *string) *AuthenticationApplicationsDelete0Params {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the authentication applications delete 0 params
func (o *AuthenticationApplicationsDelete0Params) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *AuthenticationApplicationsDelete0Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
