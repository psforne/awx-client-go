// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewOrganizationsOrganizationsUsersCreateParams creates a new OrganizationsOrganizationsUsersCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationsOrganizationsUsersCreateParams() *OrganizationsOrganizationsUsersCreateParams {
	return &OrganizationsOrganizationsUsersCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationsOrganizationsUsersCreateParamsWithTimeout creates a new OrganizationsOrganizationsUsersCreateParams object
// with the ability to set a timeout on a request.
func NewOrganizationsOrganizationsUsersCreateParamsWithTimeout(timeout time.Duration) *OrganizationsOrganizationsUsersCreateParams {
	return &OrganizationsOrganizationsUsersCreateParams{
		timeout: timeout,
	}
}

// NewOrganizationsOrganizationsUsersCreateParamsWithContext creates a new OrganizationsOrganizationsUsersCreateParams object
// with the ability to set a context for a request.
func NewOrganizationsOrganizationsUsersCreateParamsWithContext(ctx context.Context) *OrganizationsOrganizationsUsersCreateParams {
	return &OrganizationsOrganizationsUsersCreateParams{
		Context: ctx,
	}
}

// NewOrganizationsOrganizationsUsersCreateParamsWithHTTPClient creates a new OrganizationsOrganizationsUsersCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrganizationsOrganizationsUsersCreateParamsWithHTTPClient(client *http.Client) *OrganizationsOrganizationsUsersCreateParams {
	return &OrganizationsOrganizationsUsersCreateParams{
		HTTPClient: client,
	}
}

/*
OrganizationsOrganizationsUsersCreateParams contains all the parameters to send to the API endpoint

	for the organizations organizations users create operation.

	Typically these are written to a http.Request.
*/
type OrganizationsOrganizationsUsersCreateParams struct {

	// Data.
	Data interface{}

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the organizations organizations users create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsOrganizationsUsersCreateParams) WithDefaults() *OrganizationsOrganizationsUsersCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organizations organizations users create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsOrganizationsUsersCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the organizations organizations users create params
func (o *OrganizationsOrganizationsUsersCreateParams) WithTimeout(timeout time.Duration) *OrganizationsOrganizationsUsersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organizations organizations users create params
func (o *OrganizationsOrganizationsUsersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organizations organizations users create params
func (o *OrganizationsOrganizationsUsersCreateParams) WithContext(ctx context.Context) *OrganizationsOrganizationsUsersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organizations organizations users create params
func (o *OrganizationsOrganizationsUsersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organizations organizations users create params
func (o *OrganizationsOrganizationsUsersCreateParams) WithHTTPClient(client *http.Client) *OrganizationsOrganizationsUsersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organizations organizations users create params
func (o *OrganizationsOrganizationsUsersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the organizations organizations users create params
func (o *OrganizationsOrganizationsUsersCreateParams) WithData(data interface{}) *OrganizationsOrganizationsUsersCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the organizations organizations users create params
func (o *OrganizationsOrganizationsUsersCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the organizations organizations users create params
func (o *OrganizationsOrganizationsUsersCreateParams) WithID(id string) *OrganizationsOrganizationsUsersCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the organizations organizations users create params
func (o *OrganizationsOrganizationsUsersCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationsOrganizationsUsersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
