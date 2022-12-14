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
	"github.com/go-openapi/swag"
)

// NewExecutionEnvironmentsExecutionEnvironmentsListParams creates a new ExecutionEnvironmentsExecutionEnvironmentsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecutionEnvironmentsExecutionEnvironmentsListParams() *ExecutionEnvironmentsExecutionEnvironmentsListParams {
	return &ExecutionEnvironmentsExecutionEnvironmentsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecutionEnvironmentsExecutionEnvironmentsListParamsWithTimeout creates a new ExecutionEnvironmentsExecutionEnvironmentsListParams object
// with the ability to set a timeout on a request.
func NewExecutionEnvironmentsExecutionEnvironmentsListParamsWithTimeout(timeout time.Duration) *ExecutionEnvironmentsExecutionEnvironmentsListParams {
	return &ExecutionEnvironmentsExecutionEnvironmentsListParams{
		timeout: timeout,
	}
}

// NewExecutionEnvironmentsExecutionEnvironmentsListParamsWithContext creates a new ExecutionEnvironmentsExecutionEnvironmentsListParams object
// with the ability to set a context for a request.
func NewExecutionEnvironmentsExecutionEnvironmentsListParamsWithContext(ctx context.Context) *ExecutionEnvironmentsExecutionEnvironmentsListParams {
	return &ExecutionEnvironmentsExecutionEnvironmentsListParams{
		Context: ctx,
	}
}

// NewExecutionEnvironmentsExecutionEnvironmentsListParamsWithHTTPClient creates a new ExecutionEnvironmentsExecutionEnvironmentsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecutionEnvironmentsExecutionEnvironmentsListParamsWithHTTPClient(client *http.Client) *ExecutionEnvironmentsExecutionEnvironmentsListParams {
	return &ExecutionEnvironmentsExecutionEnvironmentsListParams{
		HTTPClient: client,
	}
}

/*
ExecutionEnvironmentsExecutionEnvironmentsListParams contains all the parameters to send to the API endpoint

	for the execution environments execution environments list operation.

	Typically these are written to a http.Request.
*/
type ExecutionEnvironmentsExecutionEnvironmentsListParams struct {

	/* Page.

	   A page number within the paginated result set.
	*/
	Page *int64

	/* PageSize.

	   Number of results to return per page.
	*/
	PageSize *int64

	/* Search.

	   A search term.
	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the execution environments execution environments list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecutionEnvironmentsExecutionEnvironmentsListParams) WithDefaults() *ExecutionEnvironmentsExecutionEnvironmentsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the execution environments execution environments list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecutionEnvironmentsExecutionEnvironmentsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the execution environments execution environments list params
func (o *ExecutionEnvironmentsExecutionEnvironmentsListParams) WithTimeout(timeout time.Duration) *ExecutionEnvironmentsExecutionEnvironmentsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execution environments execution environments list params
func (o *ExecutionEnvironmentsExecutionEnvironmentsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execution environments execution environments list params
func (o *ExecutionEnvironmentsExecutionEnvironmentsListParams) WithContext(ctx context.Context) *ExecutionEnvironmentsExecutionEnvironmentsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execution environments execution environments list params
func (o *ExecutionEnvironmentsExecutionEnvironmentsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execution environments execution environments list params
func (o *ExecutionEnvironmentsExecutionEnvironmentsListParams) WithHTTPClient(client *http.Client) *ExecutionEnvironmentsExecutionEnvironmentsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execution environments execution environments list params
func (o *ExecutionEnvironmentsExecutionEnvironmentsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the execution environments execution environments list params
func (o *ExecutionEnvironmentsExecutionEnvironmentsListParams) WithPage(page *int64) *ExecutionEnvironmentsExecutionEnvironmentsListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the execution environments execution environments list params
func (o *ExecutionEnvironmentsExecutionEnvironmentsListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the execution environments execution environments list params
func (o *ExecutionEnvironmentsExecutionEnvironmentsListParams) WithPageSize(pageSize *int64) *ExecutionEnvironmentsExecutionEnvironmentsListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the execution environments execution environments list params
func (o *ExecutionEnvironmentsExecutionEnvironmentsListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the execution environments execution environments list params
func (o *ExecutionEnvironmentsExecutionEnvironmentsListParams) WithSearch(search *string) *ExecutionEnvironmentsExecutionEnvironmentsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the execution environments execution environments list params
func (o *ExecutionEnvironmentsExecutionEnvironmentsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *ExecutionEnvironmentsExecutionEnvironmentsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
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
