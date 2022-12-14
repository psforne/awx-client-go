// Code generated by go-swagger; DO NOT EDIT.

package instance_groups

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

// NewInstanceGroupsInstanceGroupsDeleteParams creates a new InstanceGroupsInstanceGroupsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInstanceGroupsInstanceGroupsDeleteParams() *InstanceGroupsInstanceGroupsDeleteParams {
	return &InstanceGroupsInstanceGroupsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInstanceGroupsInstanceGroupsDeleteParamsWithTimeout creates a new InstanceGroupsInstanceGroupsDeleteParams object
// with the ability to set a timeout on a request.
func NewInstanceGroupsInstanceGroupsDeleteParamsWithTimeout(timeout time.Duration) *InstanceGroupsInstanceGroupsDeleteParams {
	return &InstanceGroupsInstanceGroupsDeleteParams{
		timeout: timeout,
	}
}

// NewInstanceGroupsInstanceGroupsDeleteParamsWithContext creates a new InstanceGroupsInstanceGroupsDeleteParams object
// with the ability to set a context for a request.
func NewInstanceGroupsInstanceGroupsDeleteParamsWithContext(ctx context.Context) *InstanceGroupsInstanceGroupsDeleteParams {
	return &InstanceGroupsInstanceGroupsDeleteParams{
		Context: ctx,
	}
}

// NewInstanceGroupsInstanceGroupsDeleteParamsWithHTTPClient creates a new InstanceGroupsInstanceGroupsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewInstanceGroupsInstanceGroupsDeleteParamsWithHTTPClient(client *http.Client) *InstanceGroupsInstanceGroupsDeleteParams {
	return &InstanceGroupsInstanceGroupsDeleteParams{
		HTTPClient: client,
	}
}

/*
InstanceGroupsInstanceGroupsDeleteParams contains all the parameters to send to the API endpoint

	for the instance groups instance groups delete operation.

	Typically these are written to a http.Request.
*/
type InstanceGroupsInstanceGroupsDeleteParams struct {

	// ID.
	ID string

	/* Search.

	   A search term.
	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the instance groups instance groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstanceGroupsInstanceGroupsDeleteParams) WithDefaults() *InstanceGroupsInstanceGroupsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the instance groups instance groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstanceGroupsInstanceGroupsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the instance groups instance groups delete params
func (o *InstanceGroupsInstanceGroupsDeleteParams) WithTimeout(timeout time.Duration) *InstanceGroupsInstanceGroupsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the instance groups instance groups delete params
func (o *InstanceGroupsInstanceGroupsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the instance groups instance groups delete params
func (o *InstanceGroupsInstanceGroupsDeleteParams) WithContext(ctx context.Context) *InstanceGroupsInstanceGroupsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the instance groups instance groups delete params
func (o *InstanceGroupsInstanceGroupsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the instance groups instance groups delete params
func (o *InstanceGroupsInstanceGroupsDeleteParams) WithHTTPClient(client *http.Client) *InstanceGroupsInstanceGroupsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the instance groups instance groups delete params
func (o *InstanceGroupsInstanceGroupsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the instance groups instance groups delete params
func (o *InstanceGroupsInstanceGroupsDeleteParams) WithID(id string) *InstanceGroupsInstanceGroupsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the instance groups instance groups delete params
func (o *InstanceGroupsInstanceGroupsDeleteParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the instance groups instance groups delete params
func (o *InstanceGroupsInstanceGroupsDeleteParams) WithSearch(search *string) *InstanceGroupsInstanceGroupsDeleteParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the instance groups instance groups delete params
func (o *InstanceGroupsInstanceGroupsDeleteParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *InstanceGroupsInstanceGroupsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Search != nil {

		// query param search
		var qrSearch string

		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {

			if err := r.SetQueryParam("search", qSearch); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
