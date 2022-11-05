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
	"github.com/go-openapi/swag"
)

// NewAuthenticationTokensList0Params creates a new AuthenticationTokensList0Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthenticationTokensList0Params() *AuthenticationTokensList0Params {
	return &AuthenticationTokensList0Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthenticationTokensList0ParamsWithTimeout creates a new AuthenticationTokensList0Params object
// with the ability to set a timeout on a request.
func NewAuthenticationTokensList0ParamsWithTimeout(timeout time.Duration) *AuthenticationTokensList0Params {
	return &AuthenticationTokensList0Params{
		timeout: timeout,
	}
}

// NewAuthenticationTokensList0ParamsWithContext creates a new AuthenticationTokensList0Params object
// with the ability to set a context for a request.
func NewAuthenticationTokensList0ParamsWithContext(ctx context.Context) *AuthenticationTokensList0Params {
	return &AuthenticationTokensList0Params{
		Context: ctx,
	}
}

// NewAuthenticationTokensList0ParamsWithHTTPClient creates a new AuthenticationTokensList0Params object
// with the ability to set a custom HTTPClient for a request.
func NewAuthenticationTokensList0ParamsWithHTTPClient(client *http.Client) *AuthenticationTokensList0Params {
	return &AuthenticationTokensList0Params{
		HTTPClient: client,
	}
}

/*
AuthenticationTokensList0Params contains all the parameters to send to the API endpoint

	for the authentication tokens list 0 operation.

	Typically these are written to a http.Request.
*/
type AuthenticationTokensList0Params struct {

	/* Page.

	   A page number within the paginated result set.
	*/
	Page *int64

	/* PageSize.

	   Number of results to return per page.
	*/
	PageSize *int64

	/* Search.

	   A search term.
	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the authentication tokens list 0 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticationTokensList0Params) WithDefaults() *AuthenticationTokensList0Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the authentication tokens list 0 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticationTokensList0Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the authentication tokens list 0 params
func (o *AuthenticationTokensList0Params) WithTimeout(timeout time.Duration) *AuthenticationTokensList0Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authentication tokens list 0 params
func (o *AuthenticationTokensList0Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authentication tokens list 0 params
func (o *AuthenticationTokensList0Params) WithContext(ctx context.Context) *AuthenticationTokensList0Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authentication tokens list 0 params
func (o *AuthenticationTokensList0Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authentication tokens list 0 params
func (o *AuthenticationTokensList0Params) WithHTTPClient(client *http.Client) *AuthenticationTokensList0Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authentication tokens list 0 params
func (o *AuthenticationTokensList0Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the authentication tokens list 0 params
func (o *AuthenticationTokensList0Params) WithPage(page *int64) *AuthenticationTokensList0Params {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the authentication tokens list 0 params
func (o *AuthenticationTokensList0Params) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the authentication tokens list 0 params
func (o *AuthenticationTokensList0Params) WithPageSize(pageSize *int64) *AuthenticationTokensList0Params {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the authentication tokens list 0 params
func (o *AuthenticationTokensList0Params) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the authentication tokens list 0 params
func (o *AuthenticationTokensList0Params) WithSearch(search *string) *AuthenticationTokensList0Params {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the authentication tokens list 0 params
func (o *AuthenticationTokensList0Params) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *AuthenticationTokensList0Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
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
