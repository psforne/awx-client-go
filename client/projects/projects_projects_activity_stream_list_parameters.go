// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewProjectsProjectsActivityStreamListParams creates a new ProjectsProjectsActivityStreamListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectsProjectsActivityStreamListParams() *ProjectsProjectsActivityStreamListParams {
	return &ProjectsProjectsActivityStreamListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsProjectsActivityStreamListParamsWithTimeout creates a new ProjectsProjectsActivityStreamListParams object
// with the ability to set a timeout on a request.
func NewProjectsProjectsActivityStreamListParamsWithTimeout(timeout time.Duration) *ProjectsProjectsActivityStreamListParams {
	return &ProjectsProjectsActivityStreamListParams{
		timeout: timeout,
	}
}

// NewProjectsProjectsActivityStreamListParamsWithContext creates a new ProjectsProjectsActivityStreamListParams object
// with the ability to set a context for a request.
func NewProjectsProjectsActivityStreamListParamsWithContext(ctx context.Context) *ProjectsProjectsActivityStreamListParams {
	return &ProjectsProjectsActivityStreamListParams{
		Context: ctx,
	}
}

// NewProjectsProjectsActivityStreamListParamsWithHTTPClient creates a new ProjectsProjectsActivityStreamListParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectsProjectsActivityStreamListParamsWithHTTPClient(client *http.Client) *ProjectsProjectsActivityStreamListParams {
	return &ProjectsProjectsActivityStreamListParams{
		HTTPClient: client,
	}
}

/*
ProjectsProjectsActivityStreamListParams contains all the parameters to send to the API endpoint

	for the projects projects activity stream list operation.

	Typically these are written to a http.Request.
*/
type ProjectsProjectsActivityStreamListParams struct {

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

// WithDefaults hydrates default values in the projects projects activity stream list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsProjectsActivityStreamListParams) WithDefaults() *ProjectsProjectsActivityStreamListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the projects projects activity stream list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsProjectsActivityStreamListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the projects projects activity stream list params
func (o *ProjectsProjectsActivityStreamListParams) WithTimeout(timeout time.Duration) *ProjectsProjectsActivityStreamListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects projects activity stream list params
func (o *ProjectsProjectsActivityStreamListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects projects activity stream list params
func (o *ProjectsProjectsActivityStreamListParams) WithContext(ctx context.Context) *ProjectsProjectsActivityStreamListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects projects activity stream list params
func (o *ProjectsProjectsActivityStreamListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects projects activity stream list params
func (o *ProjectsProjectsActivityStreamListParams) WithHTTPClient(client *http.Client) *ProjectsProjectsActivityStreamListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects projects activity stream list params
func (o *ProjectsProjectsActivityStreamListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the projects projects activity stream list params
func (o *ProjectsProjectsActivityStreamListParams) WithID(id string) *ProjectsProjectsActivityStreamListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects projects activity stream list params
func (o *ProjectsProjectsActivityStreamListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the projects projects activity stream list params
func (o *ProjectsProjectsActivityStreamListParams) WithPage(page *int64) *ProjectsProjectsActivityStreamListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the projects projects activity stream list params
func (o *ProjectsProjectsActivityStreamListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the projects projects activity stream list params
func (o *ProjectsProjectsActivityStreamListParams) WithPageSize(pageSize *int64) *ProjectsProjectsActivityStreamListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the projects projects activity stream list params
func (o *ProjectsProjectsActivityStreamListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the projects projects activity stream list params
func (o *ProjectsProjectsActivityStreamListParams) WithSearch(search *string) *ProjectsProjectsActivityStreamListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the projects projects activity stream list params
func (o *ProjectsProjectsActivityStreamListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsProjectsActivityStreamListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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