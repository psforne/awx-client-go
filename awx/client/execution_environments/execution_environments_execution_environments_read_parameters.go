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

// NewExecutionEnvironmentsExecutionEnvironmentsReadParams creates a new ExecutionEnvironmentsExecutionEnvironmentsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecutionEnvironmentsExecutionEnvironmentsReadParams() *ExecutionEnvironmentsExecutionEnvironmentsReadParams {
	return &ExecutionEnvironmentsExecutionEnvironmentsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecutionEnvironmentsExecutionEnvironmentsReadParamsWithTimeout creates a new ExecutionEnvironmentsExecutionEnvironmentsReadParams object
// with the ability to set a timeout on a request.
func NewExecutionEnvironmentsExecutionEnvironmentsReadParamsWithTimeout(timeout time.Duration) *ExecutionEnvironmentsExecutionEnvironmentsReadParams {
	return &ExecutionEnvironmentsExecutionEnvironmentsReadParams{
		timeout: timeout,
	}
}

// NewExecutionEnvironmentsExecutionEnvironmentsReadParamsWithContext creates a new ExecutionEnvironmentsExecutionEnvironmentsReadParams object
// with the ability to set a context for a request.
func NewExecutionEnvironmentsExecutionEnvironmentsReadParamsWithContext(ctx context.Context) *ExecutionEnvironmentsExecutionEnvironmentsReadParams {
	return &ExecutionEnvironmentsExecutionEnvironmentsReadParams{
		Context: ctx,
	}
}

// NewExecutionEnvironmentsExecutionEnvironmentsReadParamsWithHTTPClient creates a new ExecutionEnvironmentsExecutionEnvironmentsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecutionEnvironmentsExecutionEnvironmentsReadParamsWithHTTPClient(client *http.Client) *ExecutionEnvironmentsExecutionEnvironmentsReadParams {
	return &ExecutionEnvironmentsExecutionEnvironmentsReadParams{
		HTTPClient: client,
	}
}

/*
ExecutionEnvironmentsExecutionEnvironmentsReadParams contains all the parameters to send to the API endpoint

	for the execution environments execution environments read operation.

	Typically these are written to a http.Request.
*/
type ExecutionEnvironmentsExecutionEnvironmentsReadParams struct {

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

// WithDefaults hydrates default values in the execution environments execution environments read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecutionEnvironmentsExecutionEnvironmentsReadParams) WithDefaults() *ExecutionEnvironmentsExecutionEnvironmentsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the execution environments execution environments read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecutionEnvironmentsExecutionEnvironmentsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the execution environments execution environments read params
func (o *ExecutionEnvironmentsExecutionEnvironmentsReadParams) WithTimeout(timeout time.Duration) *ExecutionEnvironmentsExecutionEnvironmentsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execution environments execution environments read params
func (o *ExecutionEnvironmentsExecutionEnvironmentsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execution environments execution environments read params
func (o *ExecutionEnvironmentsExecutionEnvironmentsReadParams) WithContext(ctx context.Context) *ExecutionEnvironmentsExecutionEnvironmentsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execution environments execution environments read params
func (o *ExecutionEnvironmentsExecutionEnvironmentsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execution environments execution environments read params
func (o *ExecutionEnvironmentsExecutionEnvironmentsReadParams) WithHTTPClient(client *http.Client) *ExecutionEnvironmentsExecutionEnvironmentsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execution environments execution environments read params
func (o *ExecutionEnvironmentsExecutionEnvironmentsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the execution environments execution environments read params
func (o *ExecutionEnvironmentsExecutionEnvironmentsReadParams) WithID(id string) *ExecutionEnvironmentsExecutionEnvironmentsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the execution environments execution environments read params
func (o *ExecutionEnvironmentsExecutionEnvironmentsReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the execution environments execution environments read params
func (o *ExecutionEnvironmentsExecutionEnvironmentsReadParams) WithSearch(search *string) *ExecutionEnvironmentsExecutionEnvironmentsReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the execution environments execution environments read params
func (o *ExecutionEnvironmentsExecutionEnvironmentsReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *ExecutionEnvironmentsExecutionEnvironmentsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
