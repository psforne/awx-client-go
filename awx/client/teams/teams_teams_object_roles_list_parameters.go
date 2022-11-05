// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewTeamsTeamsObjectRolesListParams creates a new TeamsTeamsObjectRolesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTeamsTeamsObjectRolesListParams() *TeamsTeamsObjectRolesListParams {
	return &TeamsTeamsObjectRolesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTeamsTeamsObjectRolesListParamsWithTimeout creates a new TeamsTeamsObjectRolesListParams object
// with the ability to set a timeout on a request.
func NewTeamsTeamsObjectRolesListParamsWithTimeout(timeout time.Duration) *TeamsTeamsObjectRolesListParams {
	return &TeamsTeamsObjectRolesListParams{
		timeout: timeout,
	}
}

// NewTeamsTeamsObjectRolesListParamsWithContext creates a new TeamsTeamsObjectRolesListParams object
// with the ability to set a context for a request.
func NewTeamsTeamsObjectRolesListParamsWithContext(ctx context.Context) *TeamsTeamsObjectRolesListParams {
	return &TeamsTeamsObjectRolesListParams{
		Context: ctx,
	}
}

// NewTeamsTeamsObjectRolesListParamsWithHTTPClient creates a new TeamsTeamsObjectRolesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewTeamsTeamsObjectRolesListParamsWithHTTPClient(client *http.Client) *TeamsTeamsObjectRolesListParams {
	return &TeamsTeamsObjectRolesListParams{
		HTTPClient: client,
	}
}

/*
TeamsTeamsObjectRolesListParams contains all the parameters to send to the API endpoint

	for the teams teams object roles list operation.

	Typically these are written to a http.Request.
*/
type TeamsTeamsObjectRolesListParams struct {

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

// WithDefaults hydrates default values in the teams teams object roles list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeamsTeamsObjectRolesListParams) WithDefaults() *TeamsTeamsObjectRolesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the teams teams object roles list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeamsTeamsObjectRolesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the teams teams object roles list params
func (o *TeamsTeamsObjectRolesListParams) WithTimeout(timeout time.Duration) *TeamsTeamsObjectRolesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the teams teams object roles list params
func (o *TeamsTeamsObjectRolesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the teams teams object roles list params
func (o *TeamsTeamsObjectRolesListParams) WithContext(ctx context.Context) *TeamsTeamsObjectRolesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the teams teams object roles list params
func (o *TeamsTeamsObjectRolesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the teams teams object roles list params
func (o *TeamsTeamsObjectRolesListParams) WithHTTPClient(client *http.Client) *TeamsTeamsObjectRolesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the teams teams object roles list params
func (o *TeamsTeamsObjectRolesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the teams teams object roles list params
func (o *TeamsTeamsObjectRolesListParams) WithID(id string) *TeamsTeamsObjectRolesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the teams teams object roles list params
func (o *TeamsTeamsObjectRolesListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the teams teams object roles list params
func (o *TeamsTeamsObjectRolesListParams) WithPage(page *int64) *TeamsTeamsObjectRolesListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the teams teams object roles list params
func (o *TeamsTeamsObjectRolesListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the teams teams object roles list params
func (o *TeamsTeamsObjectRolesListParams) WithPageSize(pageSize *int64) *TeamsTeamsObjectRolesListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the teams teams object roles list params
func (o *TeamsTeamsObjectRolesListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the teams teams object roles list params
func (o *TeamsTeamsObjectRolesListParams) WithSearch(search *string) *TeamsTeamsObjectRolesListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the teams teams object roles list params
func (o *TeamsTeamsObjectRolesListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *TeamsTeamsObjectRolesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
