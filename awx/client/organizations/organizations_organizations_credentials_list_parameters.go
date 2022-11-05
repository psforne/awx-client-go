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

// NewOrganizationsOrganizationsCredentialsListParams creates a new OrganizationsOrganizationsCredentialsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationsOrganizationsCredentialsListParams() *OrganizationsOrganizationsCredentialsListParams {
	return &OrganizationsOrganizationsCredentialsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationsOrganizationsCredentialsListParamsWithTimeout creates a new OrganizationsOrganizationsCredentialsListParams object
// with the ability to set a timeout on a request.
func NewOrganizationsOrganizationsCredentialsListParamsWithTimeout(timeout time.Duration) *OrganizationsOrganizationsCredentialsListParams {
	return &OrganizationsOrganizationsCredentialsListParams{
		timeout: timeout,
	}
}

// NewOrganizationsOrganizationsCredentialsListParamsWithContext creates a new OrganizationsOrganizationsCredentialsListParams object
// with the ability to set a context for a request.
func NewOrganizationsOrganizationsCredentialsListParamsWithContext(ctx context.Context) *OrganizationsOrganizationsCredentialsListParams {
	return &OrganizationsOrganizationsCredentialsListParams{
		Context: ctx,
	}
}

// NewOrganizationsOrganizationsCredentialsListParamsWithHTTPClient creates a new OrganizationsOrganizationsCredentialsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrganizationsOrganizationsCredentialsListParamsWithHTTPClient(client *http.Client) *OrganizationsOrganizationsCredentialsListParams {
	return &OrganizationsOrganizationsCredentialsListParams{
		HTTPClient: client,
	}
}

/*
OrganizationsOrganizationsCredentialsListParams contains all the parameters to send to the API endpoint

	for the organizations organizations credentials list operation.

	Typically these are written to a http.Request.
*/
type OrganizationsOrganizationsCredentialsListParams struct {

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

// WithDefaults hydrates default values in the organizations organizations credentials list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsOrganizationsCredentialsListParams) WithDefaults() *OrganizationsOrganizationsCredentialsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organizations organizations credentials list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsOrganizationsCredentialsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the organizations organizations credentials list params
func (o *OrganizationsOrganizationsCredentialsListParams) WithTimeout(timeout time.Duration) *OrganizationsOrganizationsCredentialsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organizations organizations credentials list params
func (o *OrganizationsOrganizationsCredentialsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organizations organizations credentials list params
func (o *OrganizationsOrganizationsCredentialsListParams) WithContext(ctx context.Context) *OrganizationsOrganizationsCredentialsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organizations organizations credentials list params
func (o *OrganizationsOrganizationsCredentialsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organizations organizations credentials list params
func (o *OrganizationsOrganizationsCredentialsListParams) WithHTTPClient(client *http.Client) *OrganizationsOrganizationsCredentialsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organizations organizations credentials list params
func (o *OrganizationsOrganizationsCredentialsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the organizations organizations credentials list params
func (o *OrganizationsOrganizationsCredentialsListParams) WithID(id string) *OrganizationsOrganizationsCredentialsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the organizations organizations credentials list params
func (o *OrganizationsOrganizationsCredentialsListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the organizations organizations credentials list params
func (o *OrganizationsOrganizationsCredentialsListParams) WithPage(page *int64) *OrganizationsOrganizationsCredentialsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the organizations organizations credentials list params
func (o *OrganizationsOrganizationsCredentialsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the organizations organizations credentials list params
func (o *OrganizationsOrganizationsCredentialsListParams) WithPageSize(pageSize *int64) *OrganizationsOrganizationsCredentialsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the organizations organizations credentials list params
func (o *OrganizationsOrganizationsCredentialsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the organizations organizations credentials list params
func (o *OrganizationsOrganizationsCredentialsListParams) WithSearch(search *string) *OrganizationsOrganizationsCredentialsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the organizations organizations credentials list params
func (o *OrganizationsOrganizationsCredentialsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationsOrganizationsCredentialsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
