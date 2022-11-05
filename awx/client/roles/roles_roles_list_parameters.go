// Code generated by go-swagger; DO NOT EDIT.

package roles

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

// NewRolesRolesListParams creates a new RolesRolesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRolesRolesListParams() *RolesRolesListParams {
	return &RolesRolesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRolesRolesListParamsWithTimeout creates a new RolesRolesListParams object
// with the ability to set a timeout on a request.
func NewRolesRolesListParamsWithTimeout(timeout time.Duration) *RolesRolesListParams {
	return &RolesRolesListParams{
		timeout: timeout,
	}
}

// NewRolesRolesListParamsWithContext creates a new RolesRolesListParams object
// with the ability to set a context for a request.
func NewRolesRolesListParamsWithContext(ctx context.Context) *RolesRolesListParams {
	return &RolesRolesListParams{
		Context: ctx,
	}
}

// NewRolesRolesListParamsWithHTTPClient creates a new RolesRolesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewRolesRolesListParamsWithHTTPClient(client *http.Client) *RolesRolesListParams {
	return &RolesRolesListParams{
		HTTPClient: client,
	}
}

/*
RolesRolesListParams contains all the parameters to send to the API endpoint

	for the roles roles list operation.

	Typically these are written to a http.Request.
*/
type RolesRolesListParams struct {

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

// WithDefaults hydrates default values in the roles roles list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RolesRolesListParams) WithDefaults() *RolesRolesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the roles roles list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RolesRolesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the roles roles list params
func (o *RolesRolesListParams) WithTimeout(timeout time.Duration) *RolesRolesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the roles roles list params
func (o *RolesRolesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the roles roles list params
func (o *RolesRolesListParams) WithContext(ctx context.Context) *RolesRolesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the roles roles list params
func (o *RolesRolesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the roles roles list params
func (o *RolesRolesListParams) WithHTTPClient(client *http.Client) *RolesRolesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the roles roles list params
func (o *RolesRolesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the roles roles list params
func (o *RolesRolesListParams) WithPage(page *int64) *RolesRolesListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the roles roles list params
func (o *RolesRolesListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the roles roles list params
func (o *RolesRolesListParams) WithPageSize(pageSize *int64) *RolesRolesListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the roles roles list params
func (o *RolesRolesListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the roles roles list params
func (o *RolesRolesListParams) WithSearch(search *string) *RolesRolesListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the roles roles list params
func (o *RolesRolesListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *RolesRolesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
