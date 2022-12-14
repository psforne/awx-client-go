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

// NewHostsHostsAdHocCommandEventsListParams creates a new HostsHostsAdHocCommandEventsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHostsHostsAdHocCommandEventsListParams() *HostsHostsAdHocCommandEventsListParams {
	return &HostsHostsAdHocCommandEventsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHostsHostsAdHocCommandEventsListParamsWithTimeout creates a new HostsHostsAdHocCommandEventsListParams object
// with the ability to set a timeout on a request.
func NewHostsHostsAdHocCommandEventsListParamsWithTimeout(timeout time.Duration) *HostsHostsAdHocCommandEventsListParams {
	return &HostsHostsAdHocCommandEventsListParams{
		timeout: timeout,
	}
}

// NewHostsHostsAdHocCommandEventsListParamsWithContext creates a new HostsHostsAdHocCommandEventsListParams object
// with the ability to set a context for a request.
func NewHostsHostsAdHocCommandEventsListParamsWithContext(ctx context.Context) *HostsHostsAdHocCommandEventsListParams {
	return &HostsHostsAdHocCommandEventsListParams{
		Context: ctx,
	}
}

// NewHostsHostsAdHocCommandEventsListParamsWithHTTPClient creates a new HostsHostsAdHocCommandEventsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewHostsHostsAdHocCommandEventsListParamsWithHTTPClient(client *http.Client) *HostsHostsAdHocCommandEventsListParams {
	return &HostsHostsAdHocCommandEventsListParams{
		HTTPClient: client,
	}
}

/*
HostsHostsAdHocCommandEventsListParams contains all the parameters to send to the API endpoint

	for the hosts hosts ad hoc command events list operation.

	Typically these are written to a http.Request.
*/
type HostsHostsAdHocCommandEventsListParams struct {

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

// WithDefaults hydrates default values in the hosts hosts ad hoc command events list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsAdHocCommandEventsListParams) WithDefaults() *HostsHostsAdHocCommandEventsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hosts hosts ad hoc command events list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HostsHostsAdHocCommandEventsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hosts hosts ad hoc command events list params
func (o *HostsHostsAdHocCommandEventsListParams) WithTimeout(timeout time.Duration) *HostsHostsAdHocCommandEventsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hosts hosts ad hoc command events list params
func (o *HostsHostsAdHocCommandEventsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hosts hosts ad hoc command events list params
func (o *HostsHostsAdHocCommandEventsListParams) WithContext(ctx context.Context) *HostsHostsAdHocCommandEventsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hosts hosts ad hoc command events list params
func (o *HostsHostsAdHocCommandEventsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hosts hosts ad hoc command events list params
func (o *HostsHostsAdHocCommandEventsListParams) WithHTTPClient(client *http.Client) *HostsHostsAdHocCommandEventsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hosts hosts ad hoc command events list params
func (o *HostsHostsAdHocCommandEventsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the hosts hosts ad hoc command events list params
func (o *HostsHostsAdHocCommandEventsListParams) WithID(id string) *HostsHostsAdHocCommandEventsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the hosts hosts ad hoc command events list params
func (o *HostsHostsAdHocCommandEventsListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the hosts hosts ad hoc command events list params
func (o *HostsHostsAdHocCommandEventsListParams) WithPage(page *int64) *HostsHostsAdHocCommandEventsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the hosts hosts ad hoc command events list params
func (o *HostsHostsAdHocCommandEventsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the hosts hosts ad hoc command events list params
func (o *HostsHostsAdHocCommandEventsListParams) WithPageSize(pageSize *int64) *HostsHostsAdHocCommandEventsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the hosts hosts ad hoc command events list params
func (o *HostsHostsAdHocCommandEventsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the hosts hosts ad hoc command events list params
func (o *HostsHostsAdHocCommandEventsListParams) WithSearch(search *string) *HostsHostsAdHocCommandEventsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the hosts hosts ad hoc command events list params
func (o *HostsHostsAdHocCommandEventsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *HostsHostsAdHocCommandEventsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
