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

// NewProjectsProjectsProjectUpdatesListParams creates a new ProjectsProjectsProjectUpdatesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectsProjectsProjectUpdatesListParams() *ProjectsProjectsProjectUpdatesListParams {
	return &ProjectsProjectsProjectUpdatesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsProjectsProjectUpdatesListParamsWithTimeout creates a new ProjectsProjectsProjectUpdatesListParams object
// with the ability to set a timeout on a request.
func NewProjectsProjectsProjectUpdatesListParamsWithTimeout(timeout time.Duration) *ProjectsProjectsProjectUpdatesListParams {
	return &ProjectsProjectsProjectUpdatesListParams{
		timeout: timeout,
	}
}

// NewProjectsProjectsProjectUpdatesListParamsWithContext creates a new ProjectsProjectsProjectUpdatesListParams object
// with the ability to set a context for a request.
func NewProjectsProjectsProjectUpdatesListParamsWithContext(ctx context.Context) *ProjectsProjectsProjectUpdatesListParams {
	return &ProjectsProjectsProjectUpdatesListParams{
		Context: ctx,
	}
}

// NewProjectsProjectsProjectUpdatesListParamsWithHTTPClient creates a new ProjectsProjectsProjectUpdatesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectsProjectsProjectUpdatesListParamsWithHTTPClient(client *http.Client) *ProjectsProjectsProjectUpdatesListParams {
	return &ProjectsProjectsProjectUpdatesListParams{
		HTTPClient: client,
	}
}

/*
ProjectsProjectsProjectUpdatesListParams contains all the parameters to send to the API endpoint

	for the projects projects project updates list operation.

	Typically these are written to a http.Request.
*/
type ProjectsProjectsProjectUpdatesListParams struct {

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

// WithDefaults hydrates default values in the projects projects project updates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsProjectsProjectUpdatesListParams) WithDefaults() *ProjectsProjectsProjectUpdatesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the projects projects project updates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsProjectsProjectUpdatesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the projects projects project updates list params
func (o *ProjectsProjectsProjectUpdatesListParams) WithTimeout(timeout time.Duration) *ProjectsProjectsProjectUpdatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects projects project updates list params
func (o *ProjectsProjectsProjectUpdatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects projects project updates list params
func (o *ProjectsProjectsProjectUpdatesListParams) WithContext(ctx context.Context) *ProjectsProjectsProjectUpdatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects projects project updates list params
func (o *ProjectsProjectsProjectUpdatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects projects project updates list params
func (o *ProjectsProjectsProjectUpdatesListParams) WithHTTPClient(client *http.Client) *ProjectsProjectsProjectUpdatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects projects project updates list params
func (o *ProjectsProjectsProjectUpdatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the projects projects project updates list params
func (o *ProjectsProjectsProjectUpdatesListParams) WithID(id string) *ProjectsProjectsProjectUpdatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects projects project updates list params
func (o *ProjectsProjectsProjectUpdatesListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the projects projects project updates list params
func (o *ProjectsProjectsProjectUpdatesListParams) WithPage(page *int64) *ProjectsProjectsProjectUpdatesListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the projects projects project updates list params
func (o *ProjectsProjectsProjectUpdatesListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the projects projects project updates list params
func (o *ProjectsProjectsProjectUpdatesListParams) WithPageSize(pageSize *int64) *ProjectsProjectsProjectUpdatesListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the projects projects project updates list params
func (o *ProjectsProjectsProjectUpdatesListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the projects projects project updates list params
func (o *ProjectsProjectsProjectUpdatesListParams) WithSearch(search *string) *ProjectsProjectsProjectUpdatesListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the projects projects project updates list params
func (o *ProjectsProjectsProjectUpdatesListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsProjectsProjectUpdatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
