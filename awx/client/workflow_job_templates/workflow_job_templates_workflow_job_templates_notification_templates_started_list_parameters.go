// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

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

// NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams creates a new WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams() *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParamsWithTimeout creates a new WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams object
// with the ability to set a timeout on a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParamsWithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams{
		timeout: timeout,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParamsWithContext creates a new WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams object
// with the ability to set a context for a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParamsWithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams{
		Context: ctx,
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParamsWithHTTPClient creates a new WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParamsWithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams {
	return &WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams{
		HTTPClient: client,
	}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams contains all the parameters to send to the API endpoint

	for the workflow job templates workflow job templates notification templates started list operation.

	Typically these are written to a http.Request.
*/
type WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams struct {

	// ID.
	ID string

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

// WithDefaults hydrates default values in the workflow job templates workflow job templates notification templates started list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) WithDefaults() *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow job templates workflow job templates notification templates started list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow job templates workflow job templates notification templates started list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) WithTimeout(timeout time.Duration) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow job templates workflow job templates notification templates started list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow job templates workflow job templates notification templates started list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) WithContext(ctx context.Context) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow job templates workflow job templates notification templates started list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow job templates workflow job templates notification templates started list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) WithHTTPClient(client *http.Client) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow job templates workflow job templates notification templates started list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the workflow job templates workflow job templates notification templates started list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) WithID(id string) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow job templates workflow job templates notification templates started list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the workflow job templates workflow job templates notification templates started list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) WithPage(page *int64) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the workflow job templates workflow job templates notification templates started list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the workflow job templates workflow job templates notification templates started list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) WithPageSize(pageSize *int64) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the workflow job templates workflow job templates notification templates started list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the workflow job templates workflow job templates notification templates started list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) WithSearch(search *string) *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the workflow job templates workflow job templates notification templates started list params
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesStartedListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

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
