// Code generated by go-swagger; DO NOT EDIT.

package credentials

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

// NewCredentialsCredentialsOwnerTeamsListParams creates a new CredentialsCredentialsOwnerTeamsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCredentialsCredentialsOwnerTeamsListParams() *CredentialsCredentialsOwnerTeamsListParams {
	return &CredentialsCredentialsOwnerTeamsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCredentialsCredentialsOwnerTeamsListParamsWithTimeout creates a new CredentialsCredentialsOwnerTeamsListParams object
// with the ability to set a timeout on a request.
func NewCredentialsCredentialsOwnerTeamsListParamsWithTimeout(timeout time.Duration) *CredentialsCredentialsOwnerTeamsListParams {
	return &CredentialsCredentialsOwnerTeamsListParams{
		timeout: timeout,
	}
}

// NewCredentialsCredentialsOwnerTeamsListParamsWithContext creates a new CredentialsCredentialsOwnerTeamsListParams object
// with the ability to set a context for a request.
func NewCredentialsCredentialsOwnerTeamsListParamsWithContext(ctx context.Context) *CredentialsCredentialsOwnerTeamsListParams {
	return &CredentialsCredentialsOwnerTeamsListParams{
		Context: ctx,
	}
}

// NewCredentialsCredentialsOwnerTeamsListParamsWithHTTPClient creates a new CredentialsCredentialsOwnerTeamsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewCredentialsCredentialsOwnerTeamsListParamsWithHTTPClient(client *http.Client) *CredentialsCredentialsOwnerTeamsListParams {
	return &CredentialsCredentialsOwnerTeamsListParams{
		HTTPClient: client,
	}
}

/*
CredentialsCredentialsOwnerTeamsListParams contains all the parameters to send to the API endpoint

	for the credentials credentials owner teams list operation.

	Typically these are written to a http.Request.
*/
type CredentialsCredentialsOwnerTeamsListParams struct {

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

// WithDefaults hydrates default values in the credentials credentials owner teams list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CredentialsCredentialsOwnerTeamsListParams) WithDefaults() *CredentialsCredentialsOwnerTeamsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the credentials credentials owner teams list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CredentialsCredentialsOwnerTeamsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the credentials credentials owner teams list params
func (o *CredentialsCredentialsOwnerTeamsListParams) WithTimeout(timeout time.Duration) *CredentialsCredentialsOwnerTeamsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the credentials credentials owner teams list params
func (o *CredentialsCredentialsOwnerTeamsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the credentials credentials owner teams list params
func (o *CredentialsCredentialsOwnerTeamsListParams) WithContext(ctx context.Context) *CredentialsCredentialsOwnerTeamsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the credentials credentials owner teams list params
func (o *CredentialsCredentialsOwnerTeamsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the credentials credentials owner teams list params
func (o *CredentialsCredentialsOwnerTeamsListParams) WithHTTPClient(client *http.Client) *CredentialsCredentialsOwnerTeamsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the credentials credentials owner teams list params
func (o *CredentialsCredentialsOwnerTeamsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the credentials credentials owner teams list params
func (o *CredentialsCredentialsOwnerTeamsListParams) WithID(id string) *CredentialsCredentialsOwnerTeamsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the credentials credentials owner teams list params
func (o *CredentialsCredentialsOwnerTeamsListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the credentials credentials owner teams list params
func (o *CredentialsCredentialsOwnerTeamsListParams) WithPage(page *int64) *CredentialsCredentialsOwnerTeamsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the credentials credentials owner teams list params
func (o *CredentialsCredentialsOwnerTeamsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the credentials credentials owner teams list params
func (o *CredentialsCredentialsOwnerTeamsListParams) WithPageSize(pageSize *int64) *CredentialsCredentialsOwnerTeamsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the credentials credentials owner teams list params
func (o *CredentialsCredentialsOwnerTeamsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the credentials credentials owner teams list params
func (o *CredentialsCredentialsOwnerTeamsListParams) WithSearch(search *string) *CredentialsCredentialsOwnerTeamsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the credentials credentials owner teams list params
func (o *CredentialsCredentialsOwnerTeamsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *CredentialsCredentialsOwnerTeamsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
