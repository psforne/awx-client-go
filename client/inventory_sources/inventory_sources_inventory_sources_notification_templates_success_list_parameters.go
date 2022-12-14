// Code generated by go-swagger; DO NOT EDIT.

package inventory_sources

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

// NewInventorySourcesInventorySourcesNotificationTemplatesSuccessListParams creates a new InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInventorySourcesInventorySourcesNotificationTemplatesSuccessListParams() *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams {
	return &InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInventorySourcesInventorySourcesNotificationTemplatesSuccessListParamsWithTimeout creates a new InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams object
// with the ability to set a timeout on a request.
func NewInventorySourcesInventorySourcesNotificationTemplatesSuccessListParamsWithTimeout(timeout time.Duration) *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams {
	return &InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams{
		timeout: timeout,
	}
}

// NewInventorySourcesInventorySourcesNotificationTemplatesSuccessListParamsWithContext creates a new InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams object
// with the ability to set a context for a request.
func NewInventorySourcesInventorySourcesNotificationTemplatesSuccessListParamsWithContext(ctx context.Context) *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams {
	return &InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams{
		Context: ctx,
	}
}

// NewInventorySourcesInventorySourcesNotificationTemplatesSuccessListParamsWithHTTPClient creates a new InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams object
// with the ability to set a custom HTTPClient for a request.
func NewInventorySourcesInventorySourcesNotificationTemplatesSuccessListParamsWithHTTPClient(client *http.Client) *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams {
	return &InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams{
		HTTPClient: client,
	}
}

/*
InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams contains all the parameters to send to the API endpoint

	for the inventory sources inventory sources notification templates success list operation.

	Typically these are written to a http.Request.
*/
type InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams struct {

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

// WithDefaults hydrates default values in the inventory sources inventory sources notification templates success list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) WithDefaults() *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the inventory sources inventory sources notification templates success list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the inventory sources inventory sources notification templates success list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) WithTimeout(timeout time.Duration) *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventory sources inventory sources notification templates success list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventory sources inventory sources notification templates success list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) WithContext(ctx context.Context) *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventory sources inventory sources notification templates success list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventory sources inventory sources notification templates success list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) WithHTTPClient(client *http.Client) *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventory sources inventory sources notification templates success list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the inventory sources inventory sources notification templates success list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) WithID(id string) *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventory sources inventory sources notification templates success list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the inventory sources inventory sources notification templates success list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) WithPage(page *int64) *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the inventory sources inventory sources notification templates success list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the inventory sources inventory sources notification templates success list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) WithPageSize(pageSize *int64) *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the inventory sources inventory sources notification templates success list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the inventory sources inventory sources notification templates success list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) WithSearch(search *string) *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the inventory sources inventory sources notification templates success list params
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *InventorySourcesInventorySourcesNotificationTemplatesSuccessListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
