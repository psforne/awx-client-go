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

// NewAuthenticationTokensPartialUpdateParams creates a new AuthenticationTokensPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthenticationTokensPartialUpdateParams() *AuthenticationTokensPartialUpdateParams {
	return &AuthenticationTokensPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthenticationTokensPartialUpdateParamsWithTimeout creates a new AuthenticationTokensPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewAuthenticationTokensPartialUpdateParamsWithTimeout(timeout time.Duration) *AuthenticationTokensPartialUpdateParams {
	return &AuthenticationTokensPartialUpdateParams{
		timeout: timeout,
	}
}

// NewAuthenticationTokensPartialUpdateParamsWithContext creates a new AuthenticationTokensPartialUpdateParams object
// with the ability to set a context for a request.
func NewAuthenticationTokensPartialUpdateParamsWithContext(ctx context.Context) *AuthenticationTokensPartialUpdateParams {
	return &AuthenticationTokensPartialUpdateParams{
		Context: ctx,
	}
}

// NewAuthenticationTokensPartialUpdateParamsWithHTTPClient creates a new AuthenticationTokensPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthenticationTokensPartialUpdateParamsWithHTTPClient(client *http.Client) *AuthenticationTokensPartialUpdateParams {
	return &AuthenticationTokensPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
AuthenticationTokensPartialUpdateParams contains all the parameters to send to the API endpoint

	for the authentication tokens partial update operation.

	Typically these are written to a http.Request.
*/
type AuthenticationTokensPartialUpdateParams struct {

	// Data.
	Data interface{}

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

// WithDefaults hydrates default values in the authentication tokens partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticationTokensPartialUpdateParams) WithDefaults() *AuthenticationTokensPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the authentication tokens partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticationTokensPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the authentication tokens partial update params
func (o *AuthenticationTokensPartialUpdateParams) WithTimeout(timeout time.Duration) *AuthenticationTokensPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authentication tokens partial update params
func (o *AuthenticationTokensPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authentication tokens partial update params
func (o *AuthenticationTokensPartialUpdateParams) WithContext(ctx context.Context) *AuthenticationTokensPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authentication tokens partial update params
func (o *AuthenticationTokensPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authentication tokens partial update params
func (o *AuthenticationTokensPartialUpdateParams) WithHTTPClient(client *http.Client) *AuthenticationTokensPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authentication tokens partial update params
func (o *AuthenticationTokensPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the authentication tokens partial update params
func (o *AuthenticationTokensPartialUpdateParams) WithData(data interface{}) *AuthenticationTokensPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the authentication tokens partial update params
func (o *AuthenticationTokensPartialUpdateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the authentication tokens partial update params
func (o *AuthenticationTokensPartialUpdateParams) WithID(id string) *AuthenticationTokensPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the authentication tokens partial update params
func (o *AuthenticationTokensPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the authentication tokens partial update params
func (o *AuthenticationTokensPartialUpdateParams) WithSearch(search *string) *AuthenticationTokensPartialUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the authentication tokens partial update params
func (o *AuthenticationTokensPartialUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *AuthenticationTokensPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
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
