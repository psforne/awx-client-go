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

// NewInventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams creates a new InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams() *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams {
	return &InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInventorySourcesInventorySourcesNotificationTemplatesStartedCreateParamsWithTimeout creates a new InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams object
// with the ability to set a timeout on a request.
func NewInventorySourcesInventorySourcesNotificationTemplatesStartedCreateParamsWithTimeout(timeout time.Duration) *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams {
	return &InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams{
		timeout: timeout,
	}
}

// NewInventorySourcesInventorySourcesNotificationTemplatesStartedCreateParamsWithContext creates a new InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams object
// with the ability to set a context for a request.
func NewInventorySourcesInventorySourcesNotificationTemplatesStartedCreateParamsWithContext(ctx context.Context) *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams {
	return &InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams{
		Context: ctx,
	}
}

// NewInventorySourcesInventorySourcesNotificationTemplatesStartedCreateParamsWithHTTPClient creates a new InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewInventorySourcesInventorySourcesNotificationTemplatesStartedCreateParamsWithHTTPClient(client *http.Client) *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams {
	return &InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams{
		HTTPClient: client,
	}
}

/*
InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams contains all the parameters to send to the API endpoint

	for the inventory sources inventory sources notification templates started create operation.

	Typically these are written to a http.Request.
*/
type InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams struct {

	// Data.
	Data interface{}

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the inventory sources inventory sources notification templates started create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams) WithDefaults() *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the inventory sources inventory sources notification templates started create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the inventory sources inventory sources notification templates started create params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams) WithTimeout(timeout time.Duration) *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventory sources inventory sources notification templates started create params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventory sources inventory sources notification templates started create params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams) WithContext(ctx context.Context) *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventory sources inventory sources notification templates started create params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventory sources inventory sources notification templates started create params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams) WithHTTPClient(client *http.Client) *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventory sources inventory sources notification templates started create params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the inventory sources inventory sources notification templates started create params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams) WithData(data interface{}) *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the inventory sources inventory sources notification templates started create params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the inventory sources inventory sources notification templates started create params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams) WithID(id string) *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the inventory sources inventory sources notification templates started create params
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
