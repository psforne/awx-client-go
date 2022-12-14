// Code generated by go-swagger; DO NOT EDIT.

package execution_environments

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

// NewExecutionEnvironmentsExecutionEnvironmentsUpdateParams creates a new ExecutionEnvironmentsExecutionEnvironmentsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecutionEnvironmentsExecutionEnvironmentsUpdateParams() *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams {
	return &ExecutionEnvironmentsExecutionEnvironmentsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecutionEnvironmentsExecutionEnvironmentsUpdateParamsWithTimeout creates a new ExecutionEnvironmentsExecutionEnvironmentsUpdateParams object
// with the ability to set a timeout on a request.
func NewExecutionEnvironmentsExecutionEnvironmentsUpdateParamsWithTimeout(timeout time.Duration) *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams {
	return &ExecutionEnvironmentsExecutionEnvironmentsUpdateParams{
		timeout: timeout,
	}
}

// NewExecutionEnvironmentsExecutionEnvironmentsUpdateParamsWithContext creates a new ExecutionEnvironmentsExecutionEnvironmentsUpdateParams object
// with the ability to set a context for a request.
func NewExecutionEnvironmentsExecutionEnvironmentsUpdateParamsWithContext(ctx context.Context) *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams {
	return &ExecutionEnvironmentsExecutionEnvironmentsUpdateParams{
		Context: ctx,
	}
}

// NewExecutionEnvironmentsExecutionEnvironmentsUpdateParamsWithHTTPClient creates a new ExecutionEnvironmentsExecutionEnvironmentsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecutionEnvironmentsExecutionEnvironmentsUpdateParamsWithHTTPClient(client *http.Client) *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams {
	return &ExecutionEnvironmentsExecutionEnvironmentsUpdateParams{
		HTTPClient: client,
	}
}

/*
ExecutionEnvironmentsExecutionEnvironmentsUpdateParams contains all the parameters to send to the API endpoint

	for the execution environments execution environments update operation.

	Typically these are written to a http.Request.
*/
type ExecutionEnvironmentsExecutionEnvironmentsUpdateParams struct {

	// Data.
	Data ExecutionEnvironmentsExecutionEnvironmentsUpdateBody

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

// WithDefaults hydrates default values in the execution environments execution environments update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams) WithDefaults() *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the execution environments execution environments update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the execution environments execution environments update params
func (o *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams) WithTimeout(timeout time.Duration) *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execution environments execution environments update params
func (o *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execution environments execution environments update params
func (o *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams) WithContext(ctx context.Context) *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execution environments execution environments update params
func (o *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execution environments execution environments update params
func (o *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams) WithHTTPClient(client *http.Client) *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execution environments execution environments update params
func (o *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the execution environments execution environments update params
func (o *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams) WithData(data ExecutionEnvironmentsExecutionEnvironmentsUpdateBody) *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the execution environments execution environments update params
func (o *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams) SetData(data ExecutionEnvironmentsExecutionEnvironmentsUpdateBody) {
	o.Data = data
}

// WithID adds the id to the execution environments execution environments update params
func (o *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams) WithID(id string) *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the execution environments execution environments update params
func (o *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the execution environments execution environments update params
func (o *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams) WithSearch(search *string) *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the execution environments execution environments update params
func (o *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *ExecutionEnvironmentsExecutionEnvironmentsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
