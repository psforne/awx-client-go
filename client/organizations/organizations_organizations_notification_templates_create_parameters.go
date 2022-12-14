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

// NewOrganizationsOrganizationsNotificationTemplatesCreateParams creates a new OrganizationsOrganizationsNotificationTemplatesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationsOrganizationsNotificationTemplatesCreateParams() *OrganizationsOrganizationsNotificationTemplatesCreateParams {
	return &OrganizationsOrganizationsNotificationTemplatesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationsOrganizationsNotificationTemplatesCreateParamsWithTimeout creates a new OrganizationsOrganizationsNotificationTemplatesCreateParams object
// with the ability to set a timeout on a request.
func NewOrganizationsOrganizationsNotificationTemplatesCreateParamsWithTimeout(timeout time.Duration) *OrganizationsOrganizationsNotificationTemplatesCreateParams {
	return &OrganizationsOrganizationsNotificationTemplatesCreateParams{
		timeout: timeout,
	}
}

// NewOrganizationsOrganizationsNotificationTemplatesCreateParamsWithContext creates a new OrganizationsOrganizationsNotificationTemplatesCreateParams object
// with the ability to set a context for a request.
func NewOrganizationsOrganizationsNotificationTemplatesCreateParamsWithContext(ctx context.Context) *OrganizationsOrganizationsNotificationTemplatesCreateParams {
	return &OrganizationsOrganizationsNotificationTemplatesCreateParams{
		Context: ctx,
	}
}

// NewOrganizationsOrganizationsNotificationTemplatesCreateParamsWithHTTPClient creates a new OrganizationsOrganizationsNotificationTemplatesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrganizationsOrganizationsNotificationTemplatesCreateParamsWithHTTPClient(client *http.Client) *OrganizationsOrganizationsNotificationTemplatesCreateParams {
	return &OrganizationsOrganizationsNotificationTemplatesCreateParams{
		HTTPClient: client,
	}
}

/*
OrganizationsOrganizationsNotificationTemplatesCreateParams contains all the parameters to send to the API endpoint

	for the organizations organizations notification templates create operation.

	Typically these are written to a http.Request.
*/
type OrganizationsOrganizationsNotificationTemplatesCreateParams struct {

	// Data.
	Data OrganizationsOrganizationsNotificationTemplatesCreateBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the organizations organizations notification templates create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsOrganizationsNotificationTemplatesCreateParams) WithDefaults() *OrganizationsOrganizationsNotificationTemplatesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organizations organizations notification templates create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsOrganizationsNotificationTemplatesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the organizations organizations notification templates create params
func (o *OrganizationsOrganizationsNotificationTemplatesCreateParams) WithTimeout(timeout time.Duration) *OrganizationsOrganizationsNotificationTemplatesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organizations organizations notification templates create params
func (o *OrganizationsOrganizationsNotificationTemplatesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organizations organizations notification templates create params
func (o *OrganizationsOrganizationsNotificationTemplatesCreateParams) WithContext(ctx context.Context) *OrganizationsOrganizationsNotificationTemplatesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organizations organizations notification templates create params
func (o *OrganizationsOrganizationsNotificationTemplatesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organizations organizations notification templates create params
func (o *OrganizationsOrganizationsNotificationTemplatesCreateParams) WithHTTPClient(client *http.Client) *OrganizationsOrganizationsNotificationTemplatesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organizations organizations notification templates create params
func (o *OrganizationsOrganizationsNotificationTemplatesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the organizations organizations notification templates create params
func (o *OrganizationsOrganizationsNotificationTemplatesCreateParams) WithData(data OrganizationsOrganizationsNotificationTemplatesCreateBody) *OrganizationsOrganizationsNotificationTemplatesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the organizations organizations notification templates create params
func (o *OrganizationsOrganizationsNotificationTemplatesCreateParams) SetData(data OrganizationsOrganizationsNotificationTemplatesCreateBody) {
	o.Data = data
}

// WithID adds the id to the organizations organizations notification templates create params
func (o *OrganizationsOrganizationsNotificationTemplatesCreateParams) WithID(id string) *OrganizationsOrganizationsNotificationTemplatesCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the organizations organizations notification templates create params
func (o *OrganizationsOrganizationsNotificationTemplatesCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationsOrganizationsNotificationTemplatesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
