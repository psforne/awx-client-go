// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUsersUsersTokensListParams creates a new UsersUsersTokensListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersUsersTokensListParams() *UsersUsersTokensListParams {
	return &UsersUsersTokensListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersUsersTokensListParamsWithTimeout creates a new UsersUsersTokensListParams object
// with the ability to set a timeout on a request.
func NewUsersUsersTokensListParamsWithTimeout(timeout time.Duration) *UsersUsersTokensListParams {
	return &UsersUsersTokensListParams{
		timeout: timeout,
	}
}

// NewUsersUsersTokensListParamsWithContext creates a new UsersUsersTokensListParams object
// with the ability to set a context for a request.
func NewUsersUsersTokensListParamsWithContext(ctx context.Context) *UsersUsersTokensListParams {
	return &UsersUsersTokensListParams{
		Context: ctx,
	}
}

// NewUsersUsersTokensListParamsWithHTTPClient creates a new UsersUsersTokensListParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersUsersTokensListParamsWithHTTPClient(client *http.Client) *UsersUsersTokensListParams {
	return &UsersUsersTokensListParams{
		HTTPClient: client,
	}
}

/*
UsersUsersTokensListParams contains all the parameters to send to the API endpoint

	for the users users tokens list operation.

	Typically these are written to a http.Request.
*/
type UsersUsersTokensListParams struct {

	// ID.
	ID string

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

// WithDefaults hydrates default values in the users users tokens list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersTokensListParams) WithDefaults() *UsersUsersTokensListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users users tokens list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersTokensListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users users tokens list params
func (o *UsersUsersTokensListParams) WithTimeout(timeout time.Duration) *UsersUsersTokensListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users users tokens list params
func (o *UsersUsersTokensListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users users tokens list params
func (o *UsersUsersTokensListParams) WithContext(ctx context.Context) *UsersUsersTokensListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users users tokens list params
func (o *UsersUsersTokensListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users users tokens list params
func (o *UsersUsersTokensListParams) WithHTTPClient(client *http.Client) *UsersUsersTokensListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users users tokens list params
func (o *UsersUsersTokensListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users users tokens list params
func (o *UsersUsersTokensListParams) WithID(id string) *UsersUsersTokensListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users users tokens list params
func (o *UsersUsersTokensListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the users users tokens list params
func (o *UsersUsersTokensListParams) WithPage(page *int64) *UsersUsersTokensListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the users users tokens list params
func (o *UsersUsersTokensListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the users users tokens list params
func (o *UsersUsersTokensListParams) WithPageSize(pageSize *int64) *UsersUsersTokensListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the users users tokens list params
func (o *UsersUsersTokensListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the users users tokens list params
func (o *UsersUsersTokensListParams) WithSearch(search *string) *UsersUsersTokensListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the users users tokens list params
func (o *UsersUsersTokensListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *UsersUsersTokensListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

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