// Code generated by go-swagger; DO NOT EDIT.

package notifications

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

// NewNotificationsNotificationsListParams creates a new NotificationsNotificationsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNotificationsNotificationsListParams() *NotificationsNotificationsListParams {
	return &NotificationsNotificationsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNotificationsNotificationsListParamsWithTimeout creates a new NotificationsNotificationsListParams object
// with the ability to set a timeout on a request.
func NewNotificationsNotificationsListParamsWithTimeout(timeout time.Duration) *NotificationsNotificationsListParams {
	return &NotificationsNotificationsListParams{
		timeout: timeout,
	}
}

// NewNotificationsNotificationsListParamsWithContext creates a new NotificationsNotificationsListParams object
// with the ability to set a context for a request.
func NewNotificationsNotificationsListParamsWithContext(ctx context.Context) *NotificationsNotificationsListParams {
	return &NotificationsNotificationsListParams{
		Context: ctx,
	}
}

// NewNotificationsNotificationsListParamsWithHTTPClient creates a new NotificationsNotificationsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewNotificationsNotificationsListParamsWithHTTPClient(client *http.Client) *NotificationsNotificationsListParams {
	return &NotificationsNotificationsListParams{
		HTTPClient: client,
	}
}

/*
NotificationsNotificationsListParams contains all the parameters to send to the API endpoint

	for the notifications notifications list operation.

	Typically these are written to a http.Request.
*/
type NotificationsNotificationsListParams struct {

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

// WithDefaults hydrates default values in the notifications notifications list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotificationsNotificationsListParams) WithDefaults() *NotificationsNotificationsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the notifications notifications list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotificationsNotificationsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the notifications notifications list params
func (o *NotificationsNotificationsListParams) WithTimeout(timeout time.Duration) *NotificationsNotificationsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notifications notifications list params
func (o *NotificationsNotificationsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notifications notifications list params
func (o *NotificationsNotificationsListParams) WithContext(ctx context.Context) *NotificationsNotificationsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notifications notifications list params
func (o *NotificationsNotificationsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notifications notifications list params
func (o *NotificationsNotificationsListParams) WithHTTPClient(client *http.Client) *NotificationsNotificationsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notifications notifications list params
func (o *NotificationsNotificationsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the notifications notifications list params
func (o *NotificationsNotificationsListParams) WithPage(page *int64) *NotificationsNotificationsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the notifications notifications list params
func (o *NotificationsNotificationsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the notifications notifications list params
func (o *NotificationsNotificationsListParams) WithPageSize(pageSize *int64) *NotificationsNotificationsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the notifications notifications list params
func (o *NotificationsNotificationsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the notifications notifications list params
func (o *NotificationsNotificationsListParams) WithSearch(search *string) *NotificationsNotificationsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the notifications notifications list params
func (o *NotificationsNotificationsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *NotificationsNotificationsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
