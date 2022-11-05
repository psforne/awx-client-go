// Code generated by go-swagger; DO NOT EDIT.

package groups

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

// NewGroupsGroupsCreateParams creates a new GroupsGroupsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupsGroupsCreateParams() *GroupsGroupsCreateParams {
	return &GroupsGroupsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupsGroupsCreateParamsWithTimeout creates a new GroupsGroupsCreateParams object
// with the ability to set a timeout on a request.
func NewGroupsGroupsCreateParamsWithTimeout(timeout time.Duration) *GroupsGroupsCreateParams {
	return &GroupsGroupsCreateParams{
		timeout: timeout,
	}
}

// NewGroupsGroupsCreateParamsWithContext creates a new GroupsGroupsCreateParams object
// with the ability to set a context for a request.
func NewGroupsGroupsCreateParamsWithContext(ctx context.Context) *GroupsGroupsCreateParams {
	return &GroupsGroupsCreateParams{
		Context: ctx,
	}
}

// NewGroupsGroupsCreateParamsWithHTTPClient creates a new GroupsGroupsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupsGroupsCreateParamsWithHTTPClient(client *http.Client) *GroupsGroupsCreateParams {
	return &GroupsGroupsCreateParams{
		HTTPClient: client,
	}
}

/*
GroupsGroupsCreateParams contains all the parameters to send to the API endpoint

	for the groups groups create operation.

	Typically these are written to a http.Request.
*/
type GroupsGroupsCreateParams struct {

	// Data.
	Data GroupsGroupsCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the groups groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupsGroupsCreateParams) WithDefaults() *GroupsGroupsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the groups groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupsGroupsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the groups groups create params
func (o *GroupsGroupsCreateParams) WithTimeout(timeout time.Duration) *GroupsGroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the groups groups create params
func (o *GroupsGroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the groups groups create params
func (o *GroupsGroupsCreateParams) WithContext(ctx context.Context) *GroupsGroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the groups groups create params
func (o *GroupsGroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the groups groups create params
func (o *GroupsGroupsCreateParams) WithHTTPClient(client *http.Client) *GroupsGroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the groups groups create params
func (o *GroupsGroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the groups groups create params
func (o *GroupsGroupsCreateParams) WithData(data GroupsGroupsCreateBody) *GroupsGroupsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the groups groups create params
func (o *GroupsGroupsCreateParams) SetData(data GroupsGroupsCreateBody) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *GroupsGroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
