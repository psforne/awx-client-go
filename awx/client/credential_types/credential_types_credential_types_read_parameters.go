// Code generated by go-swagger; DO NOT EDIT.

package credential_types

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

// NewCredentialTypesCredentialTypesReadParams creates a new CredentialTypesCredentialTypesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCredentialTypesCredentialTypesReadParams() *CredentialTypesCredentialTypesReadParams {
	return &CredentialTypesCredentialTypesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCredentialTypesCredentialTypesReadParamsWithTimeout creates a new CredentialTypesCredentialTypesReadParams object
// with the ability to set a timeout on a request.
func NewCredentialTypesCredentialTypesReadParamsWithTimeout(timeout time.Duration) *CredentialTypesCredentialTypesReadParams {
	return &CredentialTypesCredentialTypesReadParams{
		timeout: timeout,
	}
}

// NewCredentialTypesCredentialTypesReadParamsWithContext creates a new CredentialTypesCredentialTypesReadParams object
// with the ability to set a context for a request.
func NewCredentialTypesCredentialTypesReadParamsWithContext(ctx context.Context) *CredentialTypesCredentialTypesReadParams {
	return &CredentialTypesCredentialTypesReadParams{
		Context: ctx,
	}
}

// NewCredentialTypesCredentialTypesReadParamsWithHTTPClient creates a new CredentialTypesCredentialTypesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewCredentialTypesCredentialTypesReadParamsWithHTTPClient(client *http.Client) *CredentialTypesCredentialTypesReadParams {
	return &CredentialTypesCredentialTypesReadParams{
		HTTPClient: client,
	}
}

/*
CredentialTypesCredentialTypesReadParams contains all the parameters to send to the API endpoint

	for the credential types credential types read operation.

	Typically these are written to a http.Request.
*/
type CredentialTypesCredentialTypesReadParams struct {

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

// WithDefaults hydrates default values in the credential types credential types read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CredentialTypesCredentialTypesReadParams) WithDefaults() *CredentialTypesCredentialTypesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the credential types credential types read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CredentialTypesCredentialTypesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the credential types credential types read params
func (o *CredentialTypesCredentialTypesReadParams) WithTimeout(timeout time.Duration) *CredentialTypesCredentialTypesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the credential types credential types read params
func (o *CredentialTypesCredentialTypesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the credential types credential types read params
func (o *CredentialTypesCredentialTypesReadParams) WithContext(ctx context.Context) *CredentialTypesCredentialTypesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the credential types credential types read params
func (o *CredentialTypesCredentialTypesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the credential types credential types read params
func (o *CredentialTypesCredentialTypesReadParams) WithHTTPClient(client *http.Client) *CredentialTypesCredentialTypesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the credential types credential types read params
func (o *CredentialTypesCredentialTypesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the credential types credential types read params
func (o *CredentialTypesCredentialTypesReadParams) WithID(id string) *CredentialTypesCredentialTypesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the credential types credential types read params
func (o *CredentialTypesCredentialTypesReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the credential types credential types read params
func (o *CredentialTypesCredentialTypesReadParams) WithSearch(search *string) *CredentialTypesCredentialTypesReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the credential types credential types read params
func (o *CredentialTypesCredentialTypesReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *CredentialTypesCredentialTypesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
