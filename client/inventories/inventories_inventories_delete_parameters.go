// Code generated by go-swagger; DO NOT EDIT.

package inventories

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
)

// NewInventoriesInventoriesDeleteParams creates a new InventoriesInventoriesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInventoriesInventoriesDeleteParams() *InventoriesInventoriesDeleteParams {
	return &InventoriesInventoriesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInventoriesInventoriesDeleteParamsWithTimeout creates a new InventoriesInventoriesDeleteParams object
// with the ability to set a timeout on a request.
func NewInventoriesInventoriesDeleteParamsWithTimeout(timeout time.Duration) *InventoriesInventoriesDeleteParams {
	return &InventoriesInventoriesDeleteParams{
		timeout: timeout,
	}
}

// NewInventoriesInventoriesDeleteParamsWithContext creates a new InventoriesInventoriesDeleteParams object
// with the ability to set a context for a request.
func NewInventoriesInventoriesDeleteParamsWithContext(ctx context.Context) *InventoriesInventoriesDeleteParams {
	return &InventoriesInventoriesDeleteParams{
		Context: ctx,
	}
}

// NewInventoriesInventoriesDeleteParamsWithHTTPClient creates a new InventoriesInventoriesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewInventoriesInventoriesDeleteParamsWithHTTPClient(client *http.Client) *InventoriesInventoriesDeleteParams {
	return &InventoriesInventoriesDeleteParams{
		HTTPClient: client,
	}
}

/*
InventoriesInventoriesDeleteParams contains all the parameters to send to the API endpoint

	for the inventories inventories delete operation.

	Typically these are written to a http.Request.
*/
type InventoriesInventoriesDeleteParams struct {

	// ID.
	ID string

	/* Search.

	   A search term.
	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the inventories inventories delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventoriesInventoriesDeleteParams) WithDefaults() *InventoriesInventoriesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the inventories inventories delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventoriesInventoriesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the inventories inventories delete params
func (o *InventoriesInventoriesDeleteParams) WithTimeout(timeout time.Duration) *InventoriesInventoriesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventories inventories delete params
func (o *InventoriesInventoriesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventories inventories delete params
func (o *InventoriesInventoriesDeleteParams) WithContext(ctx context.Context) *InventoriesInventoriesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventories inventories delete params
func (o *InventoriesInventoriesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventories inventories delete params
func (o *InventoriesInventoriesDeleteParams) WithHTTPClient(client *http.Client) *InventoriesInventoriesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventories inventories delete params
func (o *InventoriesInventoriesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the inventories inventories delete params
func (o *InventoriesInventoriesDeleteParams) WithID(id string) *InventoriesInventoriesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventories inventories delete params
func (o *InventoriesInventoriesDeleteParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the inventories inventories delete params
func (o *InventoriesInventoriesDeleteParams) WithSearch(search *string) *InventoriesInventoriesDeleteParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the inventories inventories delete params
func (o *InventoriesInventoriesDeleteParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *InventoriesInventoriesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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
