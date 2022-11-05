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

// NewExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams creates a new ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams() *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams {
	return &ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecutionEnvironmentsExecutionEnvironmentsCopyCreateParamsWithTimeout creates a new ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams object
// with the ability to set a timeout on a request.
func NewExecutionEnvironmentsExecutionEnvironmentsCopyCreateParamsWithTimeout(timeout time.Duration) *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams {
	return &ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams{
		timeout: timeout,
	}
}

// NewExecutionEnvironmentsExecutionEnvironmentsCopyCreateParamsWithContext creates a new ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams object
// with the ability to set a context for a request.
func NewExecutionEnvironmentsExecutionEnvironmentsCopyCreateParamsWithContext(ctx context.Context) *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams {
	return &ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams{
		Context: ctx,
	}
}

// NewExecutionEnvironmentsExecutionEnvironmentsCopyCreateParamsWithHTTPClient creates a new ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecutionEnvironmentsExecutionEnvironmentsCopyCreateParamsWithHTTPClient(client *http.Client) *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams {
	return &ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams{
		HTTPClient: client,
	}
}

/*
ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams contains all the parameters to send to the API endpoint

	for the execution environments execution environments copy create operation.

	Typically these are written to a http.Request.
*/
type ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams struct {

	// Data.
	Data ExecutionEnvironmentsExecutionEnvironmentsCopyCreateBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the execution environments execution environments copy create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams) WithDefaults() *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the execution environments execution environments copy create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the execution environments execution environments copy create params
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams) WithTimeout(timeout time.Duration) *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execution environments execution environments copy create params
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execution environments execution environments copy create params
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams) WithContext(ctx context.Context) *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execution environments execution environments copy create params
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execution environments execution environments copy create params
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams) WithHTTPClient(client *http.Client) *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execution environments execution environments copy create params
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the execution environments execution environments copy create params
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams) WithData(data ExecutionEnvironmentsExecutionEnvironmentsCopyCreateBody) *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the execution environments execution environments copy create params
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams) SetData(data ExecutionEnvironmentsExecutionEnvironmentsCopyCreateBody) {
	o.Data = data
}

// WithID adds the id to the execution environments execution environments copy create params
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams) WithID(id string) *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the execution environments execution environments copy create params
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}