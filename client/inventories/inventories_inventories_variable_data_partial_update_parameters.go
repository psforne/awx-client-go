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

// NewInventoriesInventoriesVariableDataPartialUpdateParams creates a new InventoriesInventoriesVariableDataPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInventoriesInventoriesVariableDataPartialUpdateParams() *InventoriesInventoriesVariableDataPartialUpdateParams {
	return &InventoriesInventoriesVariableDataPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInventoriesInventoriesVariableDataPartialUpdateParamsWithTimeout creates a new InventoriesInventoriesVariableDataPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewInventoriesInventoriesVariableDataPartialUpdateParamsWithTimeout(timeout time.Duration) *InventoriesInventoriesVariableDataPartialUpdateParams {
	return &InventoriesInventoriesVariableDataPartialUpdateParams{
		timeout: timeout,
	}
}

// NewInventoriesInventoriesVariableDataPartialUpdateParamsWithContext creates a new InventoriesInventoriesVariableDataPartialUpdateParams object
// with the ability to set a context for a request.
func NewInventoriesInventoriesVariableDataPartialUpdateParamsWithContext(ctx context.Context) *InventoriesInventoriesVariableDataPartialUpdateParams {
	return &InventoriesInventoriesVariableDataPartialUpdateParams{
		Context: ctx,
	}
}

// NewInventoriesInventoriesVariableDataPartialUpdateParamsWithHTTPClient creates a new InventoriesInventoriesVariableDataPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewInventoriesInventoriesVariableDataPartialUpdateParamsWithHTTPClient(client *http.Client) *InventoriesInventoriesVariableDataPartialUpdateParams {
	return &InventoriesInventoriesVariableDataPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
InventoriesInventoriesVariableDataPartialUpdateParams contains all the parameters to send to the API endpoint

	for the inventories inventories variable data partial update operation.

	Typically these are written to a http.Request.
*/
type InventoriesInventoriesVariableDataPartialUpdateParams struct {

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

// WithDefaults hydrates default values in the inventories inventories variable data partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventoriesInventoriesVariableDataPartialUpdateParams) WithDefaults() *InventoriesInventoriesVariableDataPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the inventories inventories variable data partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventoriesInventoriesVariableDataPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the inventories inventories variable data partial update params
func (o *InventoriesInventoriesVariableDataPartialUpdateParams) WithTimeout(timeout time.Duration) *InventoriesInventoriesVariableDataPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventories inventories variable data partial update params
func (o *InventoriesInventoriesVariableDataPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventories inventories variable data partial update params
func (o *InventoriesInventoriesVariableDataPartialUpdateParams) WithContext(ctx context.Context) *InventoriesInventoriesVariableDataPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventories inventories variable data partial update params
func (o *InventoriesInventoriesVariableDataPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventories inventories variable data partial update params
func (o *InventoriesInventoriesVariableDataPartialUpdateParams) WithHTTPClient(client *http.Client) *InventoriesInventoriesVariableDataPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventories inventories variable data partial update params
func (o *InventoriesInventoriesVariableDataPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the inventories inventories variable data partial update params
func (o *InventoriesInventoriesVariableDataPartialUpdateParams) WithData(data interface{}) *InventoriesInventoriesVariableDataPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the inventories inventories variable data partial update params
func (o *InventoriesInventoriesVariableDataPartialUpdateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the inventories inventories variable data partial update params
func (o *InventoriesInventoriesVariableDataPartialUpdateParams) WithID(id string) *InventoriesInventoriesVariableDataPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventories inventories variable data partial update params
func (o *InventoriesInventoriesVariableDataPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the inventories inventories variable data partial update params
func (o *InventoriesInventoriesVariableDataPartialUpdateParams) WithSearch(search *string) *InventoriesInventoriesVariableDataPartialUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the inventories inventories variable data partial update params
func (o *InventoriesInventoriesVariableDataPartialUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *InventoriesInventoriesVariableDataPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
