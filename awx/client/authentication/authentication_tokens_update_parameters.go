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

// NewAuthenticationTokensUpdateParams creates a new AuthenticationTokensUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthenticationTokensUpdateParams() *AuthenticationTokensUpdateParams {
	return &AuthenticationTokensUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthenticationTokensUpdateParamsWithTimeout creates a new AuthenticationTokensUpdateParams object
// with the ability to set a timeout on a request.
func NewAuthenticationTokensUpdateParamsWithTimeout(timeout time.Duration) *AuthenticationTokensUpdateParams {
	return &AuthenticationTokensUpdateParams{
		timeout: timeout,
	}
}

// NewAuthenticationTokensUpdateParamsWithContext creates a new AuthenticationTokensUpdateParams object
// with the ability to set a context for a request.
func NewAuthenticationTokensUpdateParamsWithContext(ctx context.Context) *AuthenticationTokensUpdateParams {
	return &AuthenticationTokensUpdateParams{
		Context: ctx,
	}
}

// NewAuthenticationTokensUpdateParamsWithHTTPClient creates a new AuthenticationTokensUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthenticationTokensUpdateParamsWithHTTPClient(client *http.Client) *AuthenticationTokensUpdateParams {
	return &AuthenticationTokensUpdateParams{
		HTTPClient: client,
	}
}

/*
AuthenticationTokensUpdateParams contains all the parameters to send to the API endpoint

	for the authentication tokens update operation.

	Typically these are written to a http.Request.
*/
type AuthenticationTokensUpdateParams struct {

	// Data.
	Data AuthenticationTokensUpdateBody

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

// WithDefaults hydrates default values in the authentication tokens update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticationTokensUpdateParams) WithDefaults() *AuthenticationTokensUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the authentication tokens update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticationTokensUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the authentication tokens update params
func (o *AuthenticationTokensUpdateParams) WithTimeout(timeout time.Duration) *AuthenticationTokensUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authentication tokens update params
func (o *AuthenticationTokensUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authentication tokens update params
func (o *AuthenticationTokensUpdateParams) WithContext(ctx context.Context) *AuthenticationTokensUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authentication tokens update params
func (o *AuthenticationTokensUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authentication tokens update params
func (o *AuthenticationTokensUpdateParams) WithHTTPClient(client *http.Client) *AuthenticationTokensUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authentication tokens update params
func (o *AuthenticationTokensUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the authentication tokens update params
func (o *AuthenticationTokensUpdateParams) WithData(data AuthenticationTokensUpdateBody) *AuthenticationTokensUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the authentication tokens update params
func (o *AuthenticationTokensUpdateParams) SetData(data AuthenticationTokensUpdateBody) {
	o.Data = data
}

// WithID adds the id to the authentication tokens update params
func (o *AuthenticationTokensUpdateParams) WithID(id string) *AuthenticationTokensUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the authentication tokens update params
func (o *AuthenticationTokensUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the authentication tokens update params
func (o *AuthenticationTokensUpdateParams) WithSearch(search *string) *AuthenticationTokensUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the authentication tokens update params
func (o *AuthenticationTokensUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *AuthenticationTokensUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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