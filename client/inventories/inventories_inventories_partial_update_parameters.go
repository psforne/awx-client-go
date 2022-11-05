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

// NewInventoriesInventoriesPartialUpdateParams creates a new InventoriesInventoriesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInventoriesInventoriesPartialUpdateParams() *InventoriesInventoriesPartialUpdateParams {
	return &InventoriesInventoriesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInventoriesInventoriesPartialUpdateParamsWithTimeout creates a new InventoriesInventoriesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewInventoriesInventoriesPartialUpdateParamsWithTimeout(timeout time.Duration) *InventoriesInventoriesPartialUpdateParams {
	return &InventoriesInventoriesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewInventoriesInventoriesPartialUpdateParamsWithContext creates a new InventoriesInventoriesPartialUpdateParams object
// with the ability to set a context for a request.
func NewInventoriesInventoriesPartialUpdateParamsWithContext(ctx context.Context) *InventoriesInventoriesPartialUpdateParams {
	return &InventoriesInventoriesPartialUpdateParams{
		Context: ctx,
	}
}

// NewInventoriesInventoriesPartialUpdateParamsWithHTTPClient creates a new InventoriesInventoriesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewInventoriesInventoriesPartialUpdateParamsWithHTTPClient(client *http.Client) *InventoriesInventoriesPartialUpdateParams {
	return &InventoriesInventoriesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
InventoriesInventoriesPartialUpdateParams contains all the parameters to send to the API endpoint

	for the inventories inventories partial update operation.

	Typically these are written to a http.Request.
*/
type InventoriesInventoriesPartialUpdateParams struct {

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

// WithDefaults hydrates default values in the inventories inventories partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventoriesInventoriesPartialUpdateParams) WithDefaults() *InventoriesInventoriesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the inventories inventories partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventoriesInventoriesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the inventories inventories partial update params
func (o *InventoriesInventoriesPartialUpdateParams) WithTimeout(timeout time.Duration) *InventoriesInventoriesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventories inventories partial update params
func (o *InventoriesInventoriesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventories inventories partial update params
func (o *InventoriesInventoriesPartialUpdateParams) WithContext(ctx context.Context) *InventoriesInventoriesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventories inventories partial update params
func (o *InventoriesInventoriesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventories inventories partial update params
func (o *InventoriesInventoriesPartialUpdateParams) WithHTTPClient(client *http.Client) *InventoriesInventoriesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventories inventories partial update params
func (o *InventoriesInventoriesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the inventories inventories partial update params
func (o *InventoriesInventoriesPartialUpdateParams) WithData(data interface{}) *InventoriesInventoriesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the inventories inventories partial update params
func (o *InventoriesInventoriesPartialUpdateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the inventories inventories partial update params
func (o *InventoriesInventoriesPartialUpdateParams) WithID(id string) *InventoriesInventoriesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventories inventories partial update params
func (o *InventoriesInventoriesPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the inventories inventories partial update params
func (o *InventoriesInventoriesPartialUpdateParams) WithSearch(search *string) *InventoriesInventoriesPartialUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the inventories inventories partial update params
func (o *InventoriesInventoriesPartialUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *InventoriesInventoriesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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