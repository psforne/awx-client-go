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
	"github.com/go-openapi/swag"
)

// NewHostsHostsListParams creates a new HostsHostsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHostsHostsListParams() *HostsHostsListParams {
	return &HostsHostsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHostsHostsListParamsWithTimeout creates a new HostsHostsListParams object
// with the ability to set a timeout on a request.
func NewHostsHostsListParamsWithTimeout(timeout time.Duration) *HostsHostsListParams {
	return &HostsHostsListParams{
		timeout: timeout,
	}
}

// NewHostsHostsListParamsWithContext creates a new HostsHostsListParams object
// with the ability to set a context for a request.
func NewHostsHostsListParamsWithContext(ctx context.Context) *HostsHostsListParams {
	return &HostsHostsListParams{
		Context: ctx,
	}
}

// NewHostsHostsListParamsWithHTTPClient creates a new HostsHostsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewHostsHostsListParamsWithHTTPClient(client *http.Client) *HostsHostsListParams {
	return &HostsHostsListParams{
		HTTPClient: client,
	}
}

/*
HostsHostsListParams contains all the parameters to send to the API endpoint

	for the hosts hosts list operation.

	Typically these are written to a http.Request.
*/
type HostsHostsListParams struct {

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

// WithDefaults hydrates default values in the hosts hosts list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsListParams) WithDefaults() *HostsHostsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hosts hosts list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hosts hosts list params
func (o *HostsHostsListParams) WithTimeout(timeout time.Duration) *HostsHostsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hosts hosts list params
func (o *HostsHostsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hosts hosts list params
func (o *HostsHostsListParams) WithContext(ctx context.Context) *HostsHostsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hosts hosts list params
func (o *HostsHostsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hosts hosts list params
func (o *HostsHostsListParams) WithHTTPClient(client *http.Client) *HostsHostsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hosts hosts list params
func (o *HostsHostsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the hosts hosts list params
func (o *HostsHostsListParams) WithPage(page *int64) *HostsHostsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the hosts hosts list params
func (o *HostsHostsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the hosts hosts list params
func (o *HostsHostsListParams) WithPageSize(pageSize *int64) *HostsHostsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the hosts hosts list params
func (o *HostsHostsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the hosts hosts list params
func (o *HostsHostsListParams) WithSearch(search *string) *HostsHostsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the hosts hosts list params
func (o *HostsHostsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *HostsHostsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
