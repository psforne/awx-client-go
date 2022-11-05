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

// NewInventoriesInventoriesCopyCreateParams creates a new InventoriesInventoriesCopyCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInventoriesInventoriesCopyCreateParams() *InventoriesInventoriesCopyCreateParams {
	return &InventoriesInventoriesCopyCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInventoriesInventoriesCopyCreateParamsWithTimeout creates a new InventoriesInventoriesCopyCreateParams object
// with the ability to set a timeout on a request.
func NewInventoriesInventoriesCopyCreateParamsWithTimeout(timeout time.Duration) *InventoriesInventoriesCopyCreateParams {
	return &InventoriesInventoriesCopyCreateParams{
		timeout: timeout,
	}
}

// NewInventoriesInventoriesCopyCreateParamsWithContext creates a new InventoriesInventoriesCopyCreateParams object
// with the ability to set a context for a request.
func NewInventoriesInventoriesCopyCreateParamsWithContext(ctx context.Context) *InventoriesInventoriesCopyCreateParams {
	return &InventoriesInventoriesCopyCreateParams{
		Context: ctx,
	}
}

// NewInventoriesInventoriesCopyCreateParamsWithHTTPClient creates a new InventoriesInventoriesCopyCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewInventoriesInventoriesCopyCreateParamsWithHTTPClient(client *http.Client) *InventoriesInventoriesCopyCreateParams {
	return &InventoriesInventoriesCopyCreateParams{
		HTTPClient: client,
	}
}

/*
InventoriesInventoriesCopyCreateParams contains all the parameters to send to the API endpoint

	for the inventories inventories copy create operation.

	Typically these are written to a http.Request.
*/
type InventoriesInventoriesCopyCreateParams struct {

	// Data.
	Data InventoriesInventoriesCopyCreateBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the inventories inventories copy create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventoriesInventoriesCopyCreateParams) WithDefaults() *InventoriesInventoriesCopyCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the inventories inventories copy create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventoriesInventoriesCopyCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the inventories inventories copy create params
func (o *InventoriesInventoriesCopyCreateParams) WithTimeout(timeout time.Duration) *InventoriesInventoriesCopyCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventories inventories copy create params
func (o *InventoriesInventoriesCopyCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventories inventories copy create params
func (o *InventoriesInventoriesCopyCreateParams) WithContext(ctx context.Context) *InventoriesInventoriesCopyCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventories inventories copy create params
func (o *InventoriesInventoriesCopyCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventories inventories copy create params
func (o *InventoriesInventoriesCopyCreateParams) WithHTTPClient(client *http.Client) *InventoriesInventoriesCopyCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventories inventories copy create params
func (o *InventoriesInventoriesCopyCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the inventories inventories copy create params
func (o *InventoriesInventoriesCopyCreateParams) WithData(data InventoriesInventoriesCopyCreateBody) *InventoriesInventoriesCopyCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the inventories inventories copy create params
func (o *InventoriesInventoriesCopyCreateParams) SetData(data InventoriesInventoriesCopyCreateBody) {
	o.Data = data
}

// WithID adds the id to the inventories inventories copy create params
func (o *InventoriesInventoriesCopyCreateParams) WithID(id string) *InventoriesInventoriesCopyCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventories inventories copy create params
func (o *InventoriesInventoriesCopyCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InventoriesInventoriesCopyCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
