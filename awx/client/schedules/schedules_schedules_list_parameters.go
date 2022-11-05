// Code generated by go-swagger; DO NOT EDIT.

package schedules

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

// NewSchedulesSchedulesListParams creates a new SchedulesSchedulesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSchedulesSchedulesListParams() *SchedulesSchedulesListParams {
	return &SchedulesSchedulesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSchedulesSchedulesListParamsWithTimeout creates a new SchedulesSchedulesListParams object
// with the ability to set a timeout on a request.
func NewSchedulesSchedulesListParamsWithTimeout(timeout time.Duration) *SchedulesSchedulesListParams {
	return &SchedulesSchedulesListParams{
		timeout: timeout,
	}
}

// NewSchedulesSchedulesListParamsWithContext creates a new SchedulesSchedulesListParams object
// with the ability to set a context for a request.
func NewSchedulesSchedulesListParamsWithContext(ctx context.Context) *SchedulesSchedulesListParams {
	return &SchedulesSchedulesListParams{
		Context: ctx,
	}
}

// NewSchedulesSchedulesListParamsWithHTTPClient creates a new SchedulesSchedulesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSchedulesSchedulesListParamsWithHTTPClient(client *http.Client) *SchedulesSchedulesListParams {
	return &SchedulesSchedulesListParams{
		HTTPClient: client,
	}
}

/*
SchedulesSchedulesListParams contains all the parameters to send to the API endpoint

	for the schedules schedules list operation.

	Typically these are written to a http.Request.
*/
type SchedulesSchedulesListParams struct {

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

// WithDefaults hydrates default values in the schedules schedules list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulesSchedulesListParams) WithDefaults() *SchedulesSchedulesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the schedules schedules list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulesSchedulesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the schedules schedules list params
func (o *SchedulesSchedulesListParams) WithTimeout(timeout time.Duration) *SchedulesSchedulesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schedules schedules list params
func (o *SchedulesSchedulesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schedules schedules list params
func (o *SchedulesSchedulesListParams) WithContext(ctx context.Context) *SchedulesSchedulesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schedules schedules list params
func (o *SchedulesSchedulesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schedules schedules list params
func (o *SchedulesSchedulesListParams) WithHTTPClient(client *http.Client) *SchedulesSchedulesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schedules schedules list params
func (o *SchedulesSchedulesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the schedules schedules list params
func (o *SchedulesSchedulesListParams) WithPage(page *int64) *SchedulesSchedulesListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the schedules schedules list params
func (o *SchedulesSchedulesListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the schedules schedules list params
func (o *SchedulesSchedulesListParams) WithPageSize(pageSize *int64) *SchedulesSchedulesListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the schedules schedules list params
func (o *SchedulesSchedulesListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the schedules schedules list params
func (o *SchedulesSchedulesListParams) WithSearch(search *string) *SchedulesSchedulesListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the schedules schedules list params
func (o *SchedulesSchedulesListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *SchedulesSchedulesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
