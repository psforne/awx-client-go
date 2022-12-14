// Code generated by go-swagger; DO NOT EDIT.

package ad_hoc_commands

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

// NewAdHocCommandsAdHocCommandsListParams creates a new AdHocCommandsAdHocCommandsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdHocCommandsAdHocCommandsListParams() *AdHocCommandsAdHocCommandsListParams {
	return &AdHocCommandsAdHocCommandsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdHocCommandsAdHocCommandsListParamsWithTimeout creates a new AdHocCommandsAdHocCommandsListParams object
// with the ability to set a timeout on a request.
func NewAdHocCommandsAdHocCommandsListParamsWithTimeout(timeout time.Duration) *AdHocCommandsAdHocCommandsListParams {
	return &AdHocCommandsAdHocCommandsListParams{
		timeout: timeout,
	}
}

// NewAdHocCommandsAdHocCommandsListParamsWithContext creates a new AdHocCommandsAdHocCommandsListParams object
// with the ability to set a context for a request.
func NewAdHocCommandsAdHocCommandsListParamsWithContext(ctx context.Context) *AdHocCommandsAdHocCommandsListParams {
	return &AdHocCommandsAdHocCommandsListParams{
		Context: ctx,
	}
}

// NewAdHocCommandsAdHocCommandsListParamsWithHTTPClient creates a new AdHocCommandsAdHocCommandsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdHocCommandsAdHocCommandsListParamsWithHTTPClient(client *http.Client) *AdHocCommandsAdHocCommandsListParams {
	return &AdHocCommandsAdHocCommandsListParams{
		HTTPClient: client,
	}
}

/*
AdHocCommandsAdHocCommandsListParams contains all the parameters to send to the API endpoint

	for the ad hoc commands ad hoc commands list operation.

	Typically these are written to a http.Request.
*/
type AdHocCommandsAdHocCommandsListParams struct {

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

// WithDefaults hydrates default values in the ad hoc commands ad hoc commands list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdHocCommandsAdHocCommandsListParams) WithDefaults() *AdHocCommandsAdHocCommandsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ad hoc commands ad hoc commands list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdHocCommandsAdHocCommandsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ad hoc commands ad hoc commands list params
func (o *AdHocCommandsAdHocCommandsListParams) WithTimeout(timeout time.Duration) *AdHocCommandsAdHocCommandsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ad hoc commands ad hoc commands list params
func (o *AdHocCommandsAdHocCommandsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ad hoc commands ad hoc commands list params
func (o *AdHocCommandsAdHocCommandsListParams) WithContext(ctx context.Context) *AdHocCommandsAdHocCommandsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ad hoc commands ad hoc commands list params
func (o *AdHocCommandsAdHocCommandsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ad hoc commands ad hoc commands list params
func (o *AdHocCommandsAdHocCommandsListParams) WithHTTPClient(client *http.Client) *AdHocCommandsAdHocCommandsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ad hoc commands ad hoc commands list params
func (o *AdHocCommandsAdHocCommandsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the ad hoc commands ad hoc commands list params
func (o *AdHocCommandsAdHocCommandsListParams) WithPage(page *int64) *AdHocCommandsAdHocCommandsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the ad hoc commands ad hoc commands list params
func (o *AdHocCommandsAdHocCommandsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the ad hoc commands ad hoc commands list params
func (o *AdHocCommandsAdHocCommandsListParams) WithPageSize(pageSize *int64) *AdHocCommandsAdHocCommandsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the ad hoc commands ad hoc commands list params
func (o *AdHocCommandsAdHocCommandsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the ad hoc commands ad hoc commands list params
func (o *AdHocCommandsAdHocCommandsListParams) WithSearch(search *string) *AdHocCommandsAdHocCommandsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the ad hoc commands ad hoc commands list params
func (o *AdHocCommandsAdHocCommandsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *AdHocCommandsAdHocCommandsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
