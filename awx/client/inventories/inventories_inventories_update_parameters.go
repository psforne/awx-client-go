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

// NewInventoriesInventoriesUpdateParams creates a new InventoriesInventoriesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInventoriesInventoriesUpdateParams() *InventoriesInventoriesUpdateParams {
	return &InventoriesInventoriesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInventoriesInventoriesUpdateParamsWithTimeout creates a new InventoriesInventoriesUpdateParams object
// with the ability to set a timeout on a request.
func NewInventoriesInventoriesUpdateParamsWithTimeout(timeout time.Duration) *InventoriesInventoriesUpdateParams {
	return &InventoriesInventoriesUpdateParams{
		timeout: timeout,
	}
}

// NewInventoriesInventoriesUpdateParamsWithContext creates a new InventoriesInventoriesUpdateParams object
// with the ability to set a context for a request.
func NewInventoriesInventoriesUpdateParamsWithContext(ctx context.Context) *InventoriesInventoriesUpdateParams {
	return &InventoriesInventoriesUpdateParams{
		Context: ctx,
	}
}

// NewInventoriesInventoriesUpdateParamsWithHTTPClient creates a new InventoriesInventoriesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewInventoriesInventoriesUpdateParamsWithHTTPClient(client *http.Client) *InventoriesInventoriesUpdateParams {
	return &InventoriesInventoriesUpdateParams{
		HTTPClient: client,
	}
}

/*
InventoriesInventoriesUpdateParams contains all the parameters to send to the API endpoint

	for the inventories inventories update operation.

	Typically these are written to a http.Request.
*/
type InventoriesInventoriesUpdateParams struct {

	// Data.
	Data interface{}

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

// WithDefaults hydrates default values in the inventories inventories update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventoriesInventoriesUpdateParams) WithDefaults() *InventoriesInventoriesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the inventories inventories update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventoriesInventoriesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the inventories inventories update params
func (o *InventoriesInventoriesUpdateParams) WithTimeout(timeout time.Duration) *InventoriesInventoriesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventories inventories update params
func (o *InventoriesInventoriesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventories inventories update params
func (o *InventoriesInventoriesUpdateParams) WithContext(ctx context.Context) *InventoriesInventoriesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventories inventories update params
func (o *InventoriesInventoriesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventories inventories update params
func (o *InventoriesInventoriesUpdateParams) WithHTTPClient(client *http.Client) *InventoriesInventoriesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventories inventories update params
func (o *InventoriesInventoriesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the inventories inventories update params
func (o *InventoriesInventoriesUpdateParams) WithData(data interface{}) *InventoriesInventoriesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the inventories inventories update params
func (o *InventoriesInventoriesUpdateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the inventories inventories update params
func (o *InventoriesInventoriesUpdateParams) WithID(id string) *InventoriesInventoriesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventories inventories update params
func (o *InventoriesInventoriesUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the inventories inventories update params
func (o *InventoriesInventoriesUpdateParams) WithSearch(search *string) *InventoriesInventoriesUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the inventories inventories update params
func (o *InventoriesInventoriesUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *InventoriesInventoriesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

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
