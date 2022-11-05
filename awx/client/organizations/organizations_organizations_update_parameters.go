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

// NewOrganizationsOrganizationsUpdateParams creates a new OrganizationsOrganizationsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationsOrganizationsUpdateParams() *OrganizationsOrganizationsUpdateParams {
	return &OrganizationsOrganizationsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationsOrganizationsUpdateParamsWithTimeout creates a new OrganizationsOrganizationsUpdateParams object
// with the ability to set a timeout on a request.
func NewOrganizationsOrganizationsUpdateParamsWithTimeout(timeout time.Duration) *OrganizationsOrganizationsUpdateParams {
	return &OrganizationsOrganizationsUpdateParams{
		timeout: timeout,
	}
}

// NewOrganizationsOrganizationsUpdateParamsWithContext creates a new OrganizationsOrganizationsUpdateParams object
// with the ability to set a context for a request.
func NewOrganizationsOrganizationsUpdateParamsWithContext(ctx context.Context) *OrganizationsOrganizationsUpdateParams {
	return &OrganizationsOrganizationsUpdateParams{
		Context: ctx,
	}
}

// NewOrganizationsOrganizationsUpdateParamsWithHTTPClient creates a new OrganizationsOrganizationsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrganizationsOrganizationsUpdateParamsWithHTTPClient(client *http.Client) *OrganizationsOrganizationsUpdateParams {
	return &OrganizationsOrganizationsUpdateParams{
		HTTPClient: client,
	}
}

/*
OrganizationsOrganizationsUpdateParams contains all the parameters to send to the API endpoint

	for the organizations organizations update operation.

	Typically these are written to a http.Request.
*/
type OrganizationsOrganizationsUpdateParams struct {

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

// WithDefaults hydrates default values in the organizations organizations update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsOrganizationsUpdateParams) WithDefaults() *OrganizationsOrganizationsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organizations organizations update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsOrganizationsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the organizations organizations update params
func (o *OrganizationsOrganizationsUpdateParams) WithTimeout(timeout time.Duration) *OrganizationsOrganizationsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organizations organizations update params
func (o *OrganizationsOrganizationsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organizations organizations update params
func (o *OrganizationsOrganizationsUpdateParams) WithContext(ctx context.Context) *OrganizationsOrganizationsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organizations organizations update params
func (o *OrganizationsOrganizationsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organizations organizations update params
func (o *OrganizationsOrganizationsUpdateParams) WithHTTPClient(client *http.Client) *OrganizationsOrganizationsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organizations organizations update params
func (o *OrganizationsOrganizationsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the organizations organizations update params
func (o *OrganizationsOrganizationsUpdateParams) WithData(data interface{}) *OrganizationsOrganizationsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the organizations organizations update params
func (o *OrganizationsOrganizationsUpdateParams) SetData(data interface{}) {
	o.Data = data
}

// WithID adds the id to the organizations organizations update params
func (o *OrganizationsOrganizationsUpdateParams) WithID(id string) *OrganizationsOrganizationsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the organizations organizations update params
func (o *OrganizationsOrganizationsUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the organizations organizations update params
func (o *OrganizationsOrganizationsUpdateParams) WithSearch(search *string) *OrganizationsOrganizationsUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the organizations organizations update params
func (o *OrganizationsOrganizationsUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationsOrganizationsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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