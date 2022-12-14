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

// NewGroupsGroupsVariableDataReadParams creates a new GroupsGroupsVariableDataReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupsGroupsVariableDataReadParams() *GroupsGroupsVariableDataReadParams {
	return &GroupsGroupsVariableDataReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupsGroupsVariableDataReadParamsWithTimeout creates a new GroupsGroupsVariableDataReadParams object
// with the ability to set a timeout on a request.
func NewGroupsGroupsVariableDataReadParamsWithTimeout(timeout time.Duration) *GroupsGroupsVariableDataReadParams {
	return &GroupsGroupsVariableDataReadParams{
		timeout: timeout,
	}
}

// NewGroupsGroupsVariableDataReadParamsWithContext creates a new GroupsGroupsVariableDataReadParams object
// with the ability to set a context for a request.
func NewGroupsGroupsVariableDataReadParamsWithContext(ctx context.Context) *GroupsGroupsVariableDataReadParams {
	return &GroupsGroupsVariableDataReadParams{
		Context: ctx,
	}
}

// NewGroupsGroupsVariableDataReadParamsWithHTTPClient creates a new GroupsGroupsVariableDataReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupsGroupsVariableDataReadParamsWithHTTPClient(client *http.Client) *GroupsGroupsVariableDataReadParams {
	return &GroupsGroupsVariableDataReadParams{
		HTTPClient: client,
	}
}

/*
GroupsGroupsVariableDataReadParams contains all the parameters to send to the API endpoint

	for the groups groups variable data read operation.

	Typically these are written to a http.Request.
*/
type GroupsGroupsVariableDataReadParams struct {

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

// WithDefaults hydrates default values in the groups groups variable data read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupsGroupsVariableDataReadParams) WithDefaults() *GroupsGroupsVariableDataReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the groups groups variable data read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupsGroupsVariableDataReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the groups groups variable data read params
func (o *GroupsGroupsVariableDataReadParams) WithTimeout(timeout time.Duration) *GroupsGroupsVariableDataReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the groups groups variable data read params
func (o *GroupsGroupsVariableDataReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the groups groups variable data read params
func (o *GroupsGroupsVariableDataReadParams) WithContext(ctx context.Context) *GroupsGroupsVariableDataReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the groups groups variable data read params
func (o *GroupsGroupsVariableDataReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the groups groups variable data read params
func (o *GroupsGroupsVariableDataReadParams) WithHTTPClient(client *http.Client) *GroupsGroupsVariableDataReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the groups groups variable data read params
func (o *GroupsGroupsVariableDataReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the groups groups variable data read params
func (o *GroupsGroupsVariableDataReadParams) WithID(id string) *GroupsGroupsVariableDataReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the groups groups variable data read params
func (o *GroupsGroupsVariableDataReadParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the groups groups variable data read params
func (o *GroupsGroupsVariableDataReadParams) WithSearch(search *string) *GroupsGroupsVariableDataReadParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the groups groups variable data read params
func (o *GroupsGroupsVariableDataReadParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GroupsGroupsVariableDataReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
