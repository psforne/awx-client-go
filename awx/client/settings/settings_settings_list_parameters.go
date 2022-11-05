// Code generated by go-swagger; DO NOT EDIT.

package settings

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

// NewSettingsSettingsListParams creates a new SettingsSettingsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSettingsSettingsListParams() *SettingsSettingsListParams {
	return &SettingsSettingsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSettingsSettingsListParamsWithTimeout creates a new SettingsSettingsListParams object
// with the ability to set a timeout on a request.
func NewSettingsSettingsListParamsWithTimeout(timeout time.Duration) *SettingsSettingsListParams {
	return &SettingsSettingsListParams{
		timeout: timeout,
	}
}

// NewSettingsSettingsListParamsWithContext creates a new SettingsSettingsListParams object
// with the ability to set a context for a request.
func NewSettingsSettingsListParamsWithContext(ctx context.Context) *SettingsSettingsListParams {
	return &SettingsSettingsListParams{
		Context: ctx,
	}
}

// NewSettingsSettingsListParamsWithHTTPClient creates a new SettingsSettingsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSettingsSettingsListParamsWithHTTPClient(client *http.Client) *SettingsSettingsListParams {
	return &SettingsSettingsListParams{
		HTTPClient: client,
	}
}

/*
SettingsSettingsListParams contains all the parameters to send to the API endpoint

	for the settings settings list operation.

	Typically these are written to a http.Request.
*/
type SettingsSettingsListParams struct {

	/* Page.

	   A page number within the paginated result set.
	*/
	Page *int64

	/* PageSize.

	   Number of results to return per page.
	*/
	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the settings settings list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingsSettingsListParams) WithDefaults() *SettingsSettingsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the settings settings list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingsSettingsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the settings settings list params
func (o *SettingsSettingsListParams) WithTimeout(timeout time.Duration) *SettingsSettingsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the settings settings list params
func (o *SettingsSettingsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the settings settings list params
func (o *SettingsSettingsListParams) WithContext(ctx context.Context) *SettingsSettingsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the settings settings list params
func (o *SettingsSettingsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the settings settings list params
func (o *SettingsSettingsListParams) WithHTTPClient(client *http.Client) *SettingsSettingsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the settings settings list params
func (o *SettingsSettingsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the settings settings list params
func (o *SettingsSettingsListParams) WithPage(page *int64) *SettingsSettingsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the settings settings list params
func (o *SettingsSettingsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the settings settings list params
func (o *SettingsSettingsListParams) WithPageSize(pageSize *int64) *SettingsSettingsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the settings settings list params
func (o *SettingsSettingsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *SettingsSettingsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
