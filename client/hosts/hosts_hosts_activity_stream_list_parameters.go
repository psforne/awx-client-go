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

// NewHostsHostsActivityStreamListParams creates a new HostsHostsActivityStreamListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHostsHostsActivityStreamListParams() *HostsHostsActivityStreamListParams {
	return &HostsHostsActivityStreamListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHostsHostsActivityStreamListParamsWithTimeout creates a new HostsHostsActivityStreamListParams object
// with the ability to set a timeout on a request.
func NewHostsHostsActivityStreamListParamsWithTimeout(timeout time.Duration) *HostsHostsActivityStreamListParams {
	return &HostsHostsActivityStreamListParams{
		timeout: timeout,
	}
}

// NewHostsHostsActivityStreamListParamsWithContext creates a new HostsHostsActivityStreamListParams object
// with the ability to set a context for a request.
func NewHostsHostsActivityStreamListParamsWithContext(ctx context.Context) *HostsHostsActivityStreamListParams {
	return &HostsHostsActivityStreamListParams{
		Context: ctx,
	}
}

// NewHostsHostsActivityStreamListParamsWithHTTPClient creates a new HostsHostsActivityStreamListParams object
// with the ability to set a custom HTTPClient for a request.
func NewHostsHostsActivityStreamListParamsWithHTTPClient(client *http.Client) *HostsHostsActivityStreamListParams {
	return &HostsHostsActivityStreamListParams{
		HTTPClient: client,
	}
}

/*
HostsHostsActivityStreamListParams contains all the parameters to send to the API endpoint

	for the hosts hosts activity stream list operation.

	Typically these are written to a http.Request.
*/
type HostsHostsActivityStreamListParams struct {

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

// WithDefaults hydrates default values in the hosts hosts activity stream list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsActivityStreamListParams) WithDefaults() *HostsHostsActivityStreamListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hosts hosts activity stream list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsActivityStreamListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hosts hosts activity stream list params
func (o *HostsHostsActivityStreamListParams) WithTimeout(timeout time.Duration) *HostsHostsActivityStreamListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hosts hosts activity stream list params
func (o *HostsHostsActivityStreamListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hosts hosts activity stream list params
func (o *HostsHostsActivityStreamListParams) WithContext(ctx context.Context) *HostsHostsActivityStreamListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hosts hosts activity stream list params
func (o *HostsHostsActivityStreamListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hosts hosts activity stream list params
func (o *HostsHostsActivityStreamListParams) WithHTTPClient(client *http.Client) *HostsHostsActivityStreamListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hosts hosts activity stream list params
func (o *HostsHostsActivityStreamListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the hosts hosts activity stream list params
func (o *HostsHostsActivityStreamListParams) WithID(id string) *HostsHostsActivityStreamListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the hosts hosts activity stream list params
func (o *HostsHostsActivityStreamListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the hosts hosts activity stream list params
func (o *HostsHostsActivityStreamListParams) WithPage(page *int64) *HostsHostsActivityStreamListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the hosts hosts activity stream list params
func (o *HostsHostsActivityStreamListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the hosts hosts activity stream list params
func (o *HostsHostsActivityStreamListParams) WithPageSize(pageSize *int64) *HostsHostsActivityStreamListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the hosts hosts activity stream list params
func (o *HostsHostsActivityStreamListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the hosts hosts activity stream list params
func (o *HostsHostsActivityStreamListParams) WithSearch(search *string) *HostsHostsActivityStreamListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the hosts hosts activity stream list params
func (o *HostsHostsActivityStreamListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *HostsHostsActivityStreamListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
