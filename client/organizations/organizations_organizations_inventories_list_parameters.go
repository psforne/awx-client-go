// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewOrganizationsOrganizationsInventoriesListParams creates a new OrganizationsOrganizationsInventoriesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationsOrganizationsInventoriesListParams() *OrganizationsOrganizationsInventoriesListParams {
	return &OrganizationsOrganizationsInventoriesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationsOrganizationsInventoriesListParamsWithTimeout creates a new OrganizationsOrganizationsInventoriesListParams object
// with the ability to set a timeout on a request.
func NewOrganizationsOrganizationsInventoriesListParamsWithTimeout(timeout time.Duration) *OrganizationsOrganizationsInventoriesListParams {
	return &OrganizationsOrganizationsInventoriesListParams{
		timeout: timeout,
	}
}

// NewOrganizationsOrganizationsInventoriesListParamsWithContext creates a new OrganizationsOrganizationsInventoriesListParams object
// with the ability to set a context for a request.
func NewOrganizationsOrganizationsInventoriesListParamsWithContext(ctx context.Context) *OrganizationsOrganizationsInventoriesListParams {
	return &OrganizationsOrganizationsInventoriesListParams{
		Context: ctx,
	}
}

// NewOrganizationsOrganizationsInventoriesListParamsWithHTTPClient creates a new OrganizationsOrganizationsInventoriesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrganizationsOrganizationsInventoriesListParamsWithHTTPClient(client *http.Client) *OrganizationsOrganizationsInventoriesListParams {
	return &OrganizationsOrganizationsInventoriesListParams{
		HTTPClient: client,
	}
}

/*
OrganizationsOrganizationsInventoriesListParams contains all the parameters to send to the API endpoint

	for the organizations organizations inventories list operation.

	Typically these are written to a http.Request.
*/
type OrganizationsOrganizationsInventoriesListParams struct {

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

// WithDefaults hydrates default values in the organizations organizations inventories list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsOrganizationsInventoriesListParams) WithDefaults() *OrganizationsOrganizationsInventoriesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organizations organizations inventories list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsOrganizationsInventoriesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the organizations organizations inventories list params
func (o *OrganizationsOrganizationsInventoriesListParams) WithTimeout(timeout time.Duration) *OrganizationsOrganizationsInventoriesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organizations organizations inventories list params
func (o *OrganizationsOrganizationsInventoriesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organizations organizations inventories list params
func (o *OrganizationsOrganizationsInventoriesListParams) WithContext(ctx context.Context) *OrganizationsOrganizationsInventoriesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organizations organizations inventories list params
func (o *OrganizationsOrganizationsInventoriesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organizations organizations inventories list params
func (o *OrganizationsOrganizationsInventoriesListParams) WithHTTPClient(client *http.Client) *OrganizationsOrganizationsInventoriesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organizations organizations inventories list params
func (o *OrganizationsOrganizationsInventoriesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the organizations organizations inventories list params
func (o *OrganizationsOrganizationsInventoriesListParams) WithID(id string) *OrganizationsOrganizationsInventoriesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the organizations organizations inventories list params
func (o *OrganizationsOrganizationsInventoriesListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the organizations organizations inventories list params
func (o *OrganizationsOrganizationsInventoriesListParams) WithPage(page *int64) *OrganizationsOrganizationsInventoriesListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the organizations organizations inventories list params
func (o *OrganizationsOrganizationsInventoriesListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the organizations organizations inventories list params
func (o *OrganizationsOrganizationsInventoriesListParams) WithPageSize(pageSize *int64) *OrganizationsOrganizationsInventoriesListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the organizations organizations inventories list params
func (o *OrganizationsOrganizationsInventoriesListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the organizations organizations inventories list params
func (o *OrganizationsOrganizationsInventoriesListParams) WithSearch(search *string) *OrganizationsOrganizationsInventoriesListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the organizations organizations inventories list params
func (o *OrganizationsOrganizationsInventoriesListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationsOrganizationsInventoriesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
