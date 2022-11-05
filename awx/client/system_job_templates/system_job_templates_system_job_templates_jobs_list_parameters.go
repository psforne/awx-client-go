// Code generated by go-swagger; DO NOT EDIT.

package system_job_templates

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

// NewSystemJobTemplatesSystemJobTemplatesJobsListParams creates a new SystemJobTemplatesSystemJobTemplatesJobsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemJobTemplatesSystemJobTemplatesJobsListParams() *SystemJobTemplatesSystemJobTemplatesJobsListParams {
	return &SystemJobTemplatesSystemJobTemplatesJobsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSystemJobTemplatesSystemJobTemplatesJobsListParamsWithTimeout creates a new SystemJobTemplatesSystemJobTemplatesJobsListParams object
// with the ability to set a timeout on a request.
func NewSystemJobTemplatesSystemJobTemplatesJobsListParamsWithTimeout(timeout time.Duration) *SystemJobTemplatesSystemJobTemplatesJobsListParams {
	return &SystemJobTemplatesSystemJobTemplatesJobsListParams{
		timeout: timeout,
	}
}

// NewSystemJobTemplatesSystemJobTemplatesJobsListParamsWithContext creates a new SystemJobTemplatesSystemJobTemplatesJobsListParams object
// with the ability to set a context for a request.
func NewSystemJobTemplatesSystemJobTemplatesJobsListParamsWithContext(ctx context.Context) *SystemJobTemplatesSystemJobTemplatesJobsListParams {
	return &SystemJobTemplatesSystemJobTemplatesJobsListParams{
		Context: ctx,
	}
}

// NewSystemJobTemplatesSystemJobTemplatesJobsListParamsWithHTTPClient creates a new SystemJobTemplatesSystemJobTemplatesJobsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemJobTemplatesSystemJobTemplatesJobsListParamsWithHTTPClient(client *http.Client) *SystemJobTemplatesSystemJobTemplatesJobsListParams {
	return &SystemJobTemplatesSystemJobTemplatesJobsListParams{
		HTTPClient: client,
	}
}

/*
SystemJobTemplatesSystemJobTemplatesJobsListParams contains all the parameters to send to the API endpoint

	for the system job templates system job templates jobs list operation.

	Typically these are written to a http.Request.
*/
type SystemJobTemplatesSystemJobTemplatesJobsListParams struct {

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

// WithDefaults hydrates default values in the system job templates system job templates jobs list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) WithDefaults() *SystemJobTemplatesSystemJobTemplatesJobsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the system job templates system job templates jobs list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system job templates system job templates jobs list params
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) WithTimeout(timeout time.Duration) *SystemJobTemplatesSystemJobTemplatesJobsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system job templates system job templates jobs list params
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system job templates system job templates jobs list params
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) WithContext(ctx context.Context) *SystemJobTemplatesSystemJobTemplatesJobsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system job templates system job templates jobs list params
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system job templates system job templates jobs list params
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) WithHTTPClient(client *http.Client) *SystemJobTemplatesSystemJobTemplatesJobsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system job templates system job templates jobs list params
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the system job templates system job templates jobs list params
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) WithID(id string) *SystemJobTemplatesSystemJobTemplatesJobsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the system job templates system job templates jobs list params
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the system job templates system job templates jobs list params
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) WithPage(page *int64) *SystemJobTemplatesSystemJobTemplatesJobsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the system job templates system job templates jobs list params
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the system job templates system job templates jobs list params
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) WithPageSize(pageSize *int64) *SystemJobTemplatesSystemJobTemplatesJobsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the system job templates system job templates jobs list params
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the system job templates system job templates jobs list params
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) WithSearch(search *string) *SystemJobTemplatesSystemJobTemplatesJobsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the system job templates system job templates jobs list params
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *SystemJobTemplatesSystemJobTemplatesJobsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
