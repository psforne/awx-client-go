// Code generated by go-swagger; DO NOT EDIT.

package credentials

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

// NewCredentialsCredentialsUpdateParams creates a new CredentialsCredentialsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCredentialsCredentialsUpdateParams() *CredentialsCredentialsUpdateParams {
	return &CredentialsCredentialsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCredentialsCredentialsUpdateParamsWithTimeout creates a new CredentialsCredentialsUpdateParams object
// with the ability to set a timeout on a request.
func NewCredentialsCredentialsUpdateParamsWithTimeout(timeout time.Duration) *CredentialsCredentialsUpdateParams {
	return &CredentialsCredentialsUpdateParams{
		timeout: timeout,
	}
}

// NewCredentialsCredentialsUpdateParamsWithContext creates a new CredentialsCredentialsUpdateParams object
// with the ability to set a context for a request.
func NewCredentialsCredentialsUpdateParamsWithContext(ctx context.Context) *CredentialsCredentialsUpdateParams {
	return &CredentialsCredentialsUpdateParams{
		Context: ctx,
	}
}

// NewCredentialsCredentialsUpdateParamsWithHTTPClient creates a new CredentialsCredentialsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCredentialsCredentialsUpdateParamsWithHTTPClient(client *http.Client) *CredentialsCredentialsUpdateParams {
	return &CredentialsCredentialsUpdateParams{
		HTTPClient: client,
	}
}

/*
CredentialsCredentialsUpdateParams contains all the parameters to send to the API endpoint

	for the credentials credentials update operation.

	Typically these are written to a http.Request.
*/
type CredentialsCredentialsUpdateParams struct {

	// Data.
	Data interface{}

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

// WithDefaults hydrates default values in the credentials credentials update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CredentialsCredentialsUpdateParams) WithDefaults() *CredentialsCredentialsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the credentials credentials update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CredentialsCredentialsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the credentials credentials update params
func (o *CredentialsCredentialsUpdateParams) WithTimeout(timeout time.Duration) *CredentialsCredentialsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the credentials credentials update params
func (o *CredentialsCredentialsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the credentials credentials update params
func (o *CredentialsCredentialsUpdateParams) WithContext(ctx context.Context) *CredentialsCredentialsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the credentials credentials update params
func (o *CredentialsCredentialsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the credentials credentials update params
func (o *CredentialsCredentialsUpdateParams) WithHTTPClient(client *http.Client) *CredentialsCredentialsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the credentials credentials update params
func (o *CredentialsCredentialsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the credentials credentials update params
func (o *CredentialsCredentialsUpdateParams) WithData(data interface{}) *CredentialsCredentialsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the credentials credentials update params
func (o *CredentialsCredentialsUpdateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the credentials credentials update params
func (o *CredentialsCredentialsUpdateParams) WithID(id string) *CredentialsCredentialsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the credentials credentials update params
func (o *CredentialsCredentialsUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the credentials credentials update params
func (o *CredentialsCredentialsUpdateParams) WithSearch(search *string) *CredentialsCredentialsUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the credentials credentials update params
func (o *CredentialsCredentialsUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *CredentialsCredentialsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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