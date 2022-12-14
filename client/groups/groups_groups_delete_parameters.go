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

// NewGroupsGroupsDeleteParams creates a new GroupsGroupsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupsGroupsDeleteParams() *GroupsGroupsDeleteParams {
	return &GroupsGroupsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupsGroupsDeleteParamsWithTimeout creates a new GroupsGroupsDeleteParams object
// with the ability to set a timeout on a request.
func NewGroupsGroupsDeleteParamsWithTimeout(timeout time.Duration) *GroupsGroupsDeleteParams {
	return &GroupsGroupsDeleteParams{
		timeout: timeout,
	}
}

// NewGroupsGroupsDeleteParamsWithContext creates a new GroupsGroupsDeleteParams object
// with the ability to set a context for a request.
func NewGroupsGroupsDeleteParamsWithContext(ctx context.Context) *GroupsGroupsDeleteParams {
	return &GroupsGroupsDeleteParams{
		Context: ctx,
	}
}

// NewGroupsGroupsDeleteParamsWithHTTPClient creates a new GroupsGroupsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupsGroupsDeleteParamsWithHTTPClient(client *http.Client) *GroupsGroupsDeleteParams {
	return &GroupsGroupsDeleteParams{
		HTTPClient: client,
	}
}

/*
GroupsGroupsDeleteParams contains all the parameters to send to the API endpoint

	for the groups groups delete operation.

	Typically these are written to a http.Request.
*/
type GroupsGroupsDeleteParams struct {

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

// WithDefaults hydrates default values in the groups groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupsGroupsDeleteParams) WithDefaults() *GroupsGroupsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the groups groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupsGroupsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the groups groups delete params
func (o *GroupsGroupsDeleteParams) WithTimeout(timeout time.Duration) *GroupsGroupsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the groups groups delete params
func (o *GroupsGroupsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the groups groups delete params
func (o *GroupsGroupsDeleteParams) WithContext(ctx context.Context) *GroupsGroupsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the groups groups delete params
func (o *GroupsGroupsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the groups groups delete params
func (o *GroupsGroupsDeleteParams) WithHTTPClient(client *http.Client) *GroupsGroupsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the groups groups delete params
func (o *GroupsGroupsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the groups groups delete params
func (o *GroupsGroupsDeleteParams) WithID(id string) *GroupsGroupsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the groups groups delete params
func (o *GroupsGroupsDeleteParams) SetID(id string) {
	o.ID = id
}

// WithSearch adds the search to the groups groups delete params
func (o *GroupsGroupsDeleteParams) WithSearch(search *string) *GroupsGroupsDeleteParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the groups groups delete params
func (o *GroupsGroupsDeleteParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GroupsGroupsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
