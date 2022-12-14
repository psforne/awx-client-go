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

// NewTeamsTeamsCredentialsListParams creates a new TeamsTeamsCredentialsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTeamsTeamsCredentialsListParams() *TeamsTeamsCredentialsListParams {
	return &TeamsTeamsCredentialsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTeamsTeamsCredentialsListParamsWithTimeout creates a new TeamsTeamsCredentialsListParams object
// with the ability to set a timeout on a request.
func NewTeamsTeamsCredentialsListParamsWithTimeout(timeout time.Duration) *TeamsTeamsCredentialsListParams {
	return &TeamsTeamsCredentialsListParams{
		timeout: timeout,
	}
}

// NewTeamsTeamsCredentialsListParamsWithContext creates a new TeamsTeamsCredentialsListParams object
// with the ability to set a context for a request.
func NewTeamsTeamsCredentialsListParamsWithContext(ctx context.Context) *TeamsTeamsCredentialsListParams {
	return &TeamsTeamsCredentialsListParams{
		Context: ctx,
	}
}

// NewTeamsTeamsCredentialsListParamsWithHTTPClient creates a new TeamsTeamsCredentialsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewTeamsTeamsCredentialsListParamsWithHTTPClient(client *http.Client) *TeamsTeamsCredentialsListParams {
	return &TeamsTeamsCredentialsListParams{
		HTTPClient: client,
	}
}

/*
TeamsTeamsCredentialsListParams contains all the parameters to send to the API endpoint

	for the teams teams credentials list operation.

	Typically these are written to a http.Request.
*/
type TeamsTeamsCredentialsListParams struct {

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

// WithDefaults hydrates default values in the teams teams credentials list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeamsTeamsCredentialsListParams) WithDefaults() *TeamsTeamsCredentialsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the teams teams credentials list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeamsTeamsCredentialsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the teams teams credentials list params
func (o *TeamsTeamsCredentialsListParams) WithTimeout(timeout time.Duration) *TeamsTeamsCredentialsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the teams teams credentials list params
func (o *TeamsTeamsCredentialsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the teams teams credentials list params
func (o *TeamsTeamsCredentialsListParams) WithContext(ctx context.Context) *TeamsTeamsCredentialsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the teams teams credentials list params
func (o *TeamsTeamsCredentialsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the teams teams credentials list params
func (o *TeamsTeamsCredentialsListParams) WithHTTPClient(client *http.Client) *TeamsTeamsCredentialsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the teams teams credentials list params
func (o *TeamsTeamsCredentialsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the teams teams credentials list params
func (o *TeamsTeamsCredentialsListParams) WithID(id string) *TeamsTeamsCredentialsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the teams teams credentials list params
func (o *TeamsTeamsCredentialsListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the teams teams credentials list params
func (o *TeamsTeamsCredentialsListParams) WithPage(page *int64) *TeamsTeamsCredentialsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the teams teams credentials list params
func (o *TeamsTeamsCredentialsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the teams teams credentials list params
func (o *TeamsTeamsCredentialsListParams) WithPageSize(pageSize *int64) *TeamsTeamsCredentialsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the teams teams credentials list params
func (o *TeamsTeamsCredentialsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the teams teams credentials list params
func (o *TeamsTeamsCredentialsListParams) WithSearch(search *string) *TeamsTeamsCredentialsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the teams teams credentials list params
func (o *TeamsTeamsCredentialsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *TeamsTeamsCredentialsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
