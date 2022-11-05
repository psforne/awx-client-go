// Code generated by go-swagger; DO NOT EDIT.

package roles

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

// NewRolesRolesTeamsCreateParams creates a new RolesRolesTeamsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRolesRolesTeamsCreateParams() *RolesRolesTeamsCreateParams {
	return &RolesRolesTeamsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRolesRolesTeamsCreateParamsWithTimeout creates a new RolesRolesTeamsCreateParams object
// with the ability to set a timeout on a request.
func NewRolesRolesTeamsCreateParamsWithTimeout(timeout time.Duration) *RolesRolesTeamsCreateParams {
	return &RolesRolesTeamsCreateParams{
		timeout: timeout,
	}
}

// NewRolesRolesTeamsCreateParamsWithContext creates a new RolesRolesTeamsCreateParams object
// with the ability to set a context for a request.
func NewRolesRolesTeamsCreateParamsWithContext(ctx context.Context) *RolesRolesTeamsCreateParams {
	return &RolesRolesTeamsCreateParams{
		Context: ctx,
	}
}

// NewRolesRolesTeamsCreateParamsWithHTTPClient creates a new RolesRolesTeamsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewRolesRolesTeamsCreateParamsWithHTTPClient(client *http.Client) *RolesRolesTeamsCreateParams {
	return &RolesRolesTeamsCreateParams{
		HTTPClient: client,
	}
}

/*
RolesRolesTeamsCreateParams contains all the parameters to send to the API endpoint

	for the roles roles teams create operation.

	Typically these are written to a http.Request.
*/
type RolesRolesTeamsCreateParams struct {

	// Data.
	Data interface{}

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the roles roles teams create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RolesRolesTeamsCreateParams) WithDefaults() *RolesRolesTeamsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the roles roles teams create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RolesRolesTeamsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the roles roles teams create params
func (o *RolesRolesTeamsCreateParams) WithTimeout(timeout time.Duration) *RolesRolesTeamsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the roles roles teams create params
func (o *RolesRolesTeamsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the roles roles teams create params
func (o *RolesRolesTeamsCreateParams) WithContext(ctx context.Context) *RolesRolesTeamsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the roles roles teams create params
func (o *RolesRolesTeamsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the roles roles teams create params
func (o *RolesRolesTeamsCreateParams) WithHTTPClient(client *http.Client) *RolesRolesTeamsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the roles roles teams create params
func (o *RolesRolesTeamsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the roles roles teams create params
func (o *RolesRolesTeamsCreateParams) WithData(data interface{}) *RolesRolesTeamsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the roles roles teams create params
func (o *RolesRolesTeamsCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the roles roles teams create params
func (o *RolesRolesTeamsCreateParams) WithID(id string) *RolesRolesTeamsCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the roles roles teams create params
func (o *RolesRolesTeamsCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RolesRolesTeamsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
