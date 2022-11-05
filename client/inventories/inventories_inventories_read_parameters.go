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

// NewInventoriesInventoriesReadParams creates a new InventoriesInventoriesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInventoriesInventoriesReadParams() *InventoriesInventoriesReadParams {
	return &InventoriesInventoriesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInventoriesInventoriesReadParamsWithTimeout creates a new InventoriesInventoriesReadParams object
// with the ability to set a timeout on a request.
func NewInventoriesInventoriesReadParamsWithTimeout(timeout time.Duration) *InventoriesInventoriesReadParams {
	return &InventoriesInventoriesReadParams{
		timeout: timeout,
	}
}

// NewInventoriesInventoriesReadParamsWithContext creates a new InventoriesInventoriesReadParams object
// with the ability to set a context for a request.
func NewInventoriesInventoriesReadParamsWithContext(ctx context.Context) *InventoriesInventoriesReadParams {
	return &InventoriesInventoriesReadParams{
		Context: ctx,
	}
}

// NewInventoriesInventoriesReadParamsWithHTTPClient creates a new InventoriesInventoriesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewInventoriesInventoriesReadParamsWithHTTPClient(client *http.Client) *InventoriesInventoriesReadParams {
	return &InventoriesInventoriesReadParams{
		HTTPClient: client,
	}
}

/*
InventoriesInventoriesReadParams contains all the parameters to send to the API endpoint

	for the inventories inventories read operation.

	Typically these are written to a http.Request.
*/
type InventoriesInventoriesReadParams struct {

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

// WithDefaults hydrates default values in the inventories inventories read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventoriesInventoriesReadParams) WithDefaults() *InventoriesInventoriesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the inventories inventories read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventoriesInventoriesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the inventories inventories read params
func (o *InventoriesInventoriesReadParams) WithTimeout(timeout time.Duration) *InventoriesInventoriesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventories inventories read params
func (o *InventoriesInventoriesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventories inventories read params
func (o *InventoriesInventoriesReadParams) WithContext(ctx context.Context) *InventoriesInventoriesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventories inventories read params
func (o *InventoriesInventoriesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventories inventories read params
func (o *InventoriesInventoriesReadParams) WithHTTPClient(client *http.Client) *InventoriesInventoriesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventories inventories read params
func (o *InventoriesInventoriesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the inventories inventories read params
func (o *InventoriesInventoriesReadParams) WithID(id string) *InventoriesInventoriesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventories inventories read params
func (o *InventoriesInventoriesReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the inventories inventories read params
func (o *InventoriesInventoriesReadParams) WithSearch(search *string) *InventoriesInventoriesReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the inventories inventories read params
func (o *InventoriesInventoriesReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *InventoriesInventoriesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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