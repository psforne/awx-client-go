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

// NewCredentialsCredentialsPartialUpdateParams creates a new CredentialsCredentialsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCredentialsCredentialsPartialUpdateParams() *CredentialsCredentialsPartialUpdateParams {
	return &CredentialsCredentialsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCredentialsCredentialsPartialUpdateParamsWithTimeout creates a new CredentialsCredentialsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewCredentialsCredentialsPartialUpdateParamsWithTimeout(timeout time.Duration) *CredentialsCredentialsPartialUpdateParams {
	return &CredentialsCredentialsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewCredentialsCredentialsPartialUpdateParamsWithContext creates a new CredentialsCredentialsPartialUpdateParams object
// with the ability to set a context for a request.
func NewCredentialsCredentialsPartialUpdateParamsWithContext(ctx context.Context) *CredentialsCredentialsPartialUpdateParams {
	return &CredentialsCredentialsPartialUpdateParams{
		Context: ctx,
	}
}

// NewCredentialsCredentialsPartialUpdateParamsWithHTTPClient creates a new CredentialsCredentialsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCredentialsCredentialsPartialUpdateParamsWithHTTPClient(client *http.Client) *CredentialsCredentialsPartialUpdateParams {
	return &CredentialsCredentialsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
CredentialsCredentialsPartialUpdateParams contains all the parameters to send to the API endpoint

	for the credentials credentials partial update operation.

	Typically these are written to a http.Request.
*/
type CredentialsCredentialsPartialUpdateParams struct {

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

// WithDefaults hydrates default values in the credentials credentials partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CredentialsCredentialsPartialUpdateParams) WithDefaults() *CredentialsCredentialsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the credentials credentials partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CredentialsCredentialsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the credentials credentials partial update params
func (o *CredentialsCredentialsPartialUpdateParams) WithTimeout(timeout time.Duration) *CredentialsCredentialsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the credentials credentials partial update params
func (o *CredentialsCredentialsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the credentials credentials partial update params
func (o *CredentialsCredentialsPartialUpdateParams) WithContext(ctx context.Context) *CredentialsCredentialsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the credentials credentials partial update params
func (o *CredentialsCredentialsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the credentials credentials partial update params
func (o *CredentialsCredentialsPartialUpdateParams) WithHTTPClient(client *http.Client) *CredentialsCredentialsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the credentials credentials partial update params
func (o *CredentialsCredentialsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the credentials credentials partial update params
func (o *CredentialsCredentialsPartialUpdateParams) WithData(data interface{}) *CredentialsCredentialsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the credentials credentials partial update params
func (o *CredentialsCredentialsPartialUpdateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the credentials credentials partial update params
func (o *CredentialsCredentialsPartialUpdateParams) WithID(id string) *CredentialsCredentialsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the credentials credentials partial update params
func (o *CredentialsCredentialsPartialUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the credentials credentials partial update params
func (o *CredentialsCredentialsPartialUpdateParams) WithSearch(search *string) *CredentialsCredentialsPartialUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the credentials credentials partial update params
func (o *CredentialsCredentialsPartialUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *CredentialsCredentialsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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