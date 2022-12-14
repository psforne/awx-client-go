// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewTeamsTeamsCredentialsCreateParams creates a new TeamsTeamsCredentialsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTeamsTeamsCredentialsCreateParams() *TeamsTeamsCredentialsCreateParams {
	return &TeamsTeamsCredentialsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTeamsTeamsCredentialsCreateParamsWithTimeout creates a new TeamsTeamsCredentialsCreateParams object
// with the ability to set a timeout on a request.
func NewTeamsTeamsCredentialsCreateParamsWithTimeout(timeout time.Duration) *TeamsTeamsCredentialsCreateParams {
	return &TeamsTeamsCredentialsCreateParams{
		timeout: timeout,
	}
}

// NewTeamsTeamsCredentialsCreateParamsWithContext creates a new TeamsTeamsCredentialsCreateParams object
// with the ability to set a context for a request.
func NewTeamsTeamsCredentialsCreateParamsWithContext(ctx context.Context) *TeamsTeamsCredentialsCreateParams {
	return &TeamsTeamsCredentialsCreateParams{
		Context: ctx,
	}
}

// NewTeamsTeamsCredentialsCreateParamsWithHTTPClient creates a new TeamsTeamsCredentialsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTeamsTeamsCredentialsCreateParamsWithHTTPClient(client *http.Client) *TeamsTeamsCredentialsCreateParams {
	return &TeamsTeamsCredentialsCreateParams{
		HTTPClient: client,
	}
}

/*
TeamsTeamsCredentialsCreateParams contains all the parameters to send to the API endpoint

	for the teams teams credentials create operation.

	Typically these are written to a http.Request.
*/
type TeamsTeamsCredentialsCreateParams struct {

	// Data.
	Data interface{}

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the teams teams credentials create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeamsTeamsCredentialsCreateParams) WithDefaults() *TeamsTeamsCredentialsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the teams teams credentials create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeamsTeamsCredentialsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the teams teams credentials create params
func (o *TeamsTeamsCredentialsCreateParams) WithTimeout(timeout time.Duration) *TeamsTeamsCredentialsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the teams teams credentials create params
func (o *TeamsTeamsCredentialsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the teams teams credentials create params
func (o *TeamsTeamsCredentialsCreateParams) WithContext(ctx context.Context) *TeamsTeamsCredentialsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the teams teams credentials create params
func (o *TeamsTeamsCredentialsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the teams teams credentials create params
func (o *TeamsTeamsCredentialsCreateParams) WithHTTPClient(client *http.Client) *TeamsTeamsCredentialsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the teams teams credentials create params
func (o *TeamsTeamsCredentialsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the teams teams credentials create params
func (o *TeamsTeamsCredentialsCreateParams) WithData(data interface{}) *TeamsTeamsCredentialsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the teams teams credentials create params
func (o *TeamsTeamsCredentialsCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the teams teams credentials create params
func (o *TeamsTeamsCredentialsCreateParams) WithID(id string) *TeamsTeamsCredentialsCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the teams teams credentials create params
func (o *TeamsTeamsCredentialsCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TeamsTeamsCredentialsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
