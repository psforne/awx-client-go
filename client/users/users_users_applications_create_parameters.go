// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUsersUsersApplicationsCreateParams creates a new UsersUsersApplicationsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersUsersApplicationsCreateParams() *UsersUsersApplicationsCreateParams {
	return &UsersUsersApplicationsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersUsersApplicationsCreateParamsWithTimeout creates a new UsersUsersApplicationsCreateParams object
// with the ability to set a timeout on a request.
func NewUsersUsersApplicationsCreateParamsWithTimeout(timeout time.Duration) *UsersUsersApplicationsCreateParams {
	return &UsersUsersApplicationsCreateParams{
		timeout: timeout,
	}
}

// NewUsersUsersApplicationsCreateParamsWithContext creates a new UsersUsersApplicationsCreateParams object
// with the ability to set a context for a request.
func NewUsersUsersApplicationsCreateParamsWithContext(ctx context.Context) *UsersUsersApplicationsCreateParams {
	return &UsersUsersApplicationsCreateParams{
		Context: ctx,
	}
}

// NewUsersUsersApplicationsCreateParamsWithHTTPClient creates a new UsersUsersApplicationsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersUsersApplicationsCreateParamsWithHTTPClient(client *http.Client) *UsersUsersApplicationsCreateParams {
	return &UsersUsersApplicationsCreateParams{
		HTTPClient: client,
	}
}

/*
UsersUsersApplicationsCreateParams contains all the parameters to send to the API endpoint

	for the users users applications create operation.

	Typically these are written to a http.Request.
*/
type UsersUsersApplicationsCreateParams struct {

	// Data.
	Data UsersUsersApplicationsCreateBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users users applications create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersApplicationsCreateParams) WithDefaults() *UsersUsersApplicationsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users users applications create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersApplicationsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users users applications create params
func (o *UsersUsersApplicationsCreateParams) WithTimeout(timeout time.Duration) *UsersUsersApplicationsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users users applications create params
func (o *UsersUsersApplicationsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users users applications create params
func (o *UsersUsersApplicationsCreateParams) WithContext(ctx context.Context) *UsersUsersApplicationsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users users applications create params
func (o *UsersUsersApplicationsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users users applications create params
func (o *UsersUsersApplicationsCreateParams) WithHTTPClient(client *http.Client) *UsersUsersApplicationsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users users applications create params
func (o *UsersUsersApplicationsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users users applications create params
func (o *UsersUsersApplicationsCreateParams) WithData(data UsersUsersApplicationsCreateBody) *UsersUsersApplicationsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users users applications create params
func (o *UsersUsersApplicationsCreateParams) SetData(data UsersUsersApplicationsCreateBody) {
	o.Data = data
}

// WithID adds the id to the users users applications create params
func (o *UsersUsersApplicationsCreateParams) WithID(id string) *UsersUsersApplicationsCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users users applications create params
func (o *UsersUsersApplicationsCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersUsersApplicationsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
