// Code generated by go-swagger; DO NOT EDIT.

package notification_templates

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

// NewNotificationTemplatesNotificationTemplatesTestCreateParams creates a new NotificationTemplatesNotificationTemplatesTestCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNotificationTemplatesNotificationTemplatesTestCreateParams() *NotificationTemplatesNotificationTemplatesTestCreateParams {
	return &NotificationTemplatesNotificationTemplatesTestCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNotificationTemplatesNotificationTemplatesTestCreateParamsWithTimeout creates a new NotificationTemplatesNotificationTemplatesTestCreateParams object
// with the ability to set a timeout on a request.
func NewNotificationTemplatesNotificationTemplatesTestCreateParamsWithTimeout(timeout time.Duration) *NotificationTemplatesNotificationTemplatesTestCreateParams {
	return &NotificationTemplatesNotificationTemplatesTestCreateParams{
		timeout: timeout,
	}
}

// NewNotificationTemplatesNotificationTemplatesTestCreateParamsWithContext creates a new NotificationTemplatesNotificationTemplatesTestCreateParams object
// with the ability to set a context for a request.
func NewNotificationTemplatesNotificationTemplatesTestCreateParamsWithContext(ctx context.Context) *NotificationTemplatesNotificationTemplatesTestCreateParams {
	return &NotificationTemplatesNotificationTemplatesTestCreateParams{
		Context: ctx,
	}
}

// NewNotificationTemplatesNotificationTemplatesTestCreateParamsWithHTTPClient creates a new NotificationTemplatesNotificationTemplatesTestCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewNotificationTemplatesNotificationTemplatesTestCreateParamsWithHTTPClient(client *http.Client) *NotificationTemplatesNotificationTemplatesTestCreateParams {
	return &NotificationTemplatesNotificationTemplatesTestCreateParams{
		HTTPClient: client,
	}
}

/*
NotificationTemplatesNotificationTemplatesTestCreateParams contains all the parameters to send to the API endpoint

	for the notification templates notification templates test create operation.

	Typically these are written to a http.Request.
*/
type NotificationTemplatesNotificationTemplatesTestCreateParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the notification templates notification templates test create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotificationTemplatesNotificationTemplatesTestCreateParams) WithDefaults() *NotificationTemplatesNotificationTemplatesTestCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the notification templates notification templates test create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotificationTemplatesNotificationTemplatesTestCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the notification templates notification templates test create params
func (o *NotificationTemplatesNotificationTemplatesTestCreateParams) WithTimeout(timeout time.Duration) *NotificationTemplatesNotificationTemplatesTestCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notification templates notification templates test create params
func (o *NotificationTemplatesNotificationTemplatesTestCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notification templates notification templates test create params
func (o *NotificationTemplatesNotificationTemplatesTestCreateParams) WithContext(ctx context.Context) *NotificationTemplatesNotificationTemplatesTestCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notification templates notification templates test create params
func (o *NotificationTemplatesNotificationTemplatesTestCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notification templates notification templates test create params
func (o *NotificationTemplatesNotificationTemplatesTestCreateParams) WithHTTPClient(client *http.Client) *NotificationTemplatesNotificationTemplatesTestCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notification templates notification templates test create params
func (o *NotificationTemplatesNotificationTemplatesTestCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the notification templates notification templates test create params
func (o *NotificationTemplatesNotificationTemplatesTestCreateParams) WithID(id string) *NotificationTemplatesNotificationTemplatesTestCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the notification templates notification templates test create params
func (o *NotificationTemplatesNotificationTemplatesTestCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *NotificationTemplatesNotificationTemplatesTestCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
