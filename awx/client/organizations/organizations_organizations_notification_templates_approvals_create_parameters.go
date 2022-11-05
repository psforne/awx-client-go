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

// NewOrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams creates a new OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams() *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams {
	return &OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationsOrganizationsNotificationTemplatesApprovalsCreateParamsWithTimeout creates a new OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams object
// with the ability to set a timeout on a request.
func NewOrganizationsOrganizationsNotificationTemplatesApprovalsCreateParamsWithTimeout(timeout time.Duration) *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams {
	return &OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams{
		timeout: timeout,
	}
}

// NewOrganizationsOrganizationsNotificationTemplatesApprovalsCreateParamsWithContext creates a new OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams object
// with the ability to set a context for a request.
func NewOrganizationsOrganizationsNotificationTemplatesApprovalsCreateParamsWithContext(ctx context.Context) *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams {
	return &OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams{
		Context: ctx,
	}
}

// NewOrganizationsOrganizationsNotificationTemplatesApprovalsCreateParamsWithHTTPClient creates a new OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrganizationsOrganizationsNotificationTemplatesApprovalsCreateParamsWithHTTPClient(client *http.Client) *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams {
	return &OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams{
		HTTPClient: client,
	}
}

/*
OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams contains all the parameters to send to the API endpoint

	for the organizations organizations notification templates approvals create operation.

	Typically these are written to a http.Request.
*/
type OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams struct {

	// Data.
	Data interface{}

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the organizations organizations notification templates approvals create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams) WithDefaults() *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organizations organizations notification templates approvals create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the organizations organizations notification templates approvals create params
func (o *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams) WithTimeout(timeout time.Duration) *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organizations organizations notification templates approvals create params
func (o *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organizations organizations notification templates approvals create params
func (o *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams) WithContext(ctx context.Context) *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organizations organizations notification templates approvals create params
func (o *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organizations organizations notification templates approvals create params
func (o *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams) WithHTTPClient(client *http.Client) *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organizations organizations notification templates approvals create params
func (o *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the organizations organizations notification templates approvals create params
func (o *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams) WithData(data interface{}) *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the organizations organizations notification templates approvals create params
func (o *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the organizations organizations notification templates approvals create params
func (o *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams) WithID(id string) *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the organizations organizations notification templates approvals create params
func (o *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationsOrganizationsNotificationTemplatesApprovalsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
