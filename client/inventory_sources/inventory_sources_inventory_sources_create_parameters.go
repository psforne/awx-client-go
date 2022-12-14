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
)

// NewInventorySourcesInventorySourcesCreateParams creates a new InventorySourcesInventorySourcesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInventorySourcesInventorySourcesCreateParams() *InventorySourcesInventorySourcesCreateParams {
	return &InventorySourcesInventorySourcesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInventorySourcesInventorySourcesCreateParamsWithTimeout creates a new InventorySourcesInventorySourcesCreateParams object
// with the ability to set a timeout on a request.
func NewInventorySourcesInventorySourcesCreateParamsWithTimeout(timeout time.Duration) *InventorySourcesInventorySourcesCreateParams {
	return &InventorySourcesInventorySourcesCreateParams{
		timeout: timeout,
	}
}

// NewInventorySourcesInventorySourcesCreateParamsWithContext creates a new InventorySourcesInventorySourcesCreateParams object
// with the ability to set a context for a request.
func NewInventorySourcesInventorySourcesCreateParamsWithContext(ctx context.Context) *InventorySourcesInventorySourcesCreateParams {
	return &InventorySourcesInventorySourcesCreateParams{
		Context: ctx,
	}
}

// NewInventorySourcesInventorySourcesCreateParamsWithHTTPClient creates a new InventorySourcesInventorySourcesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewInventorySourcesInventorySourcesCreateParamsWithHTTPClient(client *http.Client) *InventorySourcesInventorySourcesCreateParams {
	return &InventorySourcesInventorySourcesCreateParams{
		HTTPClient: client,
	}
}

/*
InventorySourcesInventorySourcesCreateParams contains all the parameters to send to the API endpoint

	for the inventory sources inventory sources create operation.

	Typically these are written to a http.Request.
*/
type InventorySourcesInventorySourcesCreateParams struct {

	// Data.
	Data interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the inventory sources inventory sources create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventorySourcesInventorySourcesCreateParams) WithDefaults() *InventorySourcesInventorySourcesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the inventory sources inventory sources create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventorySourcesInventorySourcesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the inventory sources inventory sources create params
func (o *InventorySourcesInventorySourcesCreateParams) WithTimeout(timeout time.Duration) *InventorySourcesInventorySourcesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventory sources inventory sources create params
func (o *InventorySourcesInventorySourcesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventory sources inventory sources create params
func (o *InventorySourcesInventorySourcesCreateParams) WithContext(ctx context.Context) *InventorySourcesInventorySourcesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventory sources inventory sources create params
func (o *InventorySourcesInventorySourcesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventory sources inventory sources create params
func (o *InventorySourcesInventorySourcesCreateParams) WithHTTPClient(client *http.Client) *InventorySourcesInventorySourcesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventory sources inventory sources create params
func (o *InventorySourcesInventorySourcesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the inventory sources inventory sources create params
func (o *InventorySourcesInventorySourcesCreateParams) WithData(data interface{}) *InventorySourcesInventorySourcesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the inventory sources inventory sources create params
func (o *InventorySourcesInventorySourcesCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *InventorySourcesInventorySourcesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
