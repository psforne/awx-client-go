// Code generated by go-swagger; DO NOT EDIT.

package instances

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

// NewInstancesInstancesUpdateParams creates a new InstancesInstancesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInstancesInstancesUpdateParams() *InstancesInstancesUpdateParams {
	return &InstancesInstancesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInstancesInstancesUpdateParamsWithTimeout creates a new InstancesInstancesUpdateParams object
// with the ability to set a timeout on a request.
func NewInstancesInstancesUpdateParamsWithTimeout(timeout time.Duration) *InstancesInstancesUpdateParams {
	return &InstancesInstancesUpdateParams{
		timeout: timeout,
	}
}

// NewInstancesInstancesUpdateParamsWithContext creates a new InstancesInstancesUpdateParams object
// with the ability to set a context for a request.
func NewInstancesInstancesUpdateParamsWithContext(ctx context.Context) *InstancesInstancesUpdateParams {
	return &InstancesInstancesUpdateParams{
		Context: ctx,
	}
}

// NewInstancesInstancesUpdateParamsWithHTTPClient creates a new InstancesInstancesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewInstancesInstancesUpdateParamsWithHTTPClient(client *http.Client) *InstancesInstancesUpdateParams {
	return &InstancesInstancesUpdateParams{
		HTTPClient: client,
	}
}

/*
InstancesInstancesUpdateParams contains all the parameters to send to the API endpoint

	for the instances instances update operation.

	Typically these are written to a http.Request.
*/
type InstancesInstancesUpdateParams struct {

	// Data.
	Data InstancesInstancesUpdateBody

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

// WithDefaults hydrates default values in the instances instances update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstancesInstancesUpdateParams) WithDefaults() *InstancesInstancesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the instances instances update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstancesInstancesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the instances instances update params
func (o *InstancesInstancesUpdateParams) WithTimeout(timeout time.Duration) *InstancesInstancesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the instances instances update params
func (o *InstancesInstancesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the instances instances update params
func (o *InstancesInstancesUpdateParams) WithContext(ctx context.Context) *InstancesInstancesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the instances instances update params
func (o *InstancesInstancesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the instances instances update params
func (o *InstancesInstancesUpdateParams) WithHTTPClient(client *http.Client) *InstancesInstancesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the instances instances update params
func (o *InstancesInstancesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the instances instances update params
func (o *InstancesInstancesUpdateParams) WithData(data InstancesInstancesUpdateBody) *InstancesInstancesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the instances instances update params
func (o *InstancesInstancesUpdateParams) SetData(data InstancesInstancesUpdateBody) {
	o.Data = data
}

// WithID adds the id to the instances instances update params
func (o *InstancesInstancesUpdateParams) WithID(id string) *InstancesInstancesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the instances instances update params
func (o *InstancesInstancesUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the instances instances update params
func (o *InstancesInstancesUpdateParams) WithSearch(search *string) *InstancesInstancesUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the instances instances update params
func (o *InstancesInstancesUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *InstancesInstancesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
