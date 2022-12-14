// Code generated by go-swagger; DO NOT EDIT.

package workflow_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new workflow jobs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workflow jobs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	WorkflowJobsWorkflowJobsActivityStreamList(params *WorkflowJobsWorkflowJobsActivityStreamListParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsActivityStreamListOK, error)

	WorkflowJobsWorkflowJobsCancelCreate(params *WorkflowJobsWorkflowJobsCancelCreateParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsCancelCreateCreated, error)

	WorkflowJobsWorkflowJobsCancelRead(params *WorkflowJobsWorkflowJobsCancelReadParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsCancelReadOK, error)

	WorkflowJobsWorkflowJobsDelete(params *WorkflowJobsWorkflowJobsDeleteParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsDeleteNoContent, error)

	WorkflowJobsWorkflowJobsLabelsList(params *WorkflowJobsWorkflowJobsLabelsListParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsLabelsListOK, error)

	WorkflowJobsWorkflowJobsList(params *WorkflowJobsWorkflowJobsListParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsListOK, error)

	WorkflowJobsWorkflowJobsNotificationsList(params *WorkflowJobsWorkflowJobsNotificationsListParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsNotificationsListOK, error)

	WorkflowJobsWorkflowJobsRead(params *WorkflowJobsWorkflowJobsReadParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsReadOK, error)

	WorkflowJobsWorkflowJobsRelaunchCreate(params *WorkflowJobsWorkflowJobsRelaunchCreateParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsRelaunchCreateCreated, error)

	WorkflowJobsWorkflowJobsRelaunchList(params *WorkflowJobsWorkflowJobsRelaunchListParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsRelaunchListOK, error)

	WorkflowJobsWorkflowJobsWorkflowNodesList(params *WorkflowJobsWorkflowJobsWorkflowNodesListParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsWorkflowNodesListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	WorkflowJobsWorkflowJobsActivityStreamList lists activity streams for a workflow job

Make a GET request to this resource to retrieve a list of
activity streams associated with the selected
workflow job.

The resulting data structure contains:

	{
	    "count": 99,
	    "next": null,
	    "previous": null,
	    "results": [
	        ...
	    ]
	}

The `count` field indicates the total number of activity streams
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more activity stream records.

## Results

Each activity stream data structure includes the following fields:

* `id`: Database ID for this activity stream. (integer)
* `type`: Data type for this activity stream. (choice)
* `url`: URL for this activity stream. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `timestamp`:  (datetime)
* `operation`: The action taken with respect to the given object(s). (choice)
  - `create`: Entity Created
  - `update`: Entity Updated
  - `delete`: Entity Deleted
  - `associate`: Entity Associated with another Entity
  - `disassociate`: Entity was Disassociated with another Entity

* `changes`: A summary of the new and changed values when an object is created, updated, or deleted (json)
* `object1`: For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. (string)
* `object2`: Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. (string)
* `object_association`: When present, shows the field name of the role or relationship that changed. (field)
* `action_node`: The cluster node the activity took place on. (string)
* `object_type`: When present, shows the model on which the role or relationship was defined. (field)

## Sorting

To specify that activity streams are returned in a particular
order, use the `order_by` query string parameter on the GET request.

	?order_by=name

Prefix the field name with a dash `-` to sort in reverse:

	?order_by=-name

Multiple sorting fields may be specified by separating the field names with a
comma `,`:

	?order_by=name,some_other_field

## Pagination

Use the `page_size` query string parameter to change the number of results
returned for each request.  Use the `page` query string parameter to retrieve
a particular page of results.

	?page_size=100&page=2

The `previous` and `next` links returned with the results will set these query
string parameters automatically.

## Searching

Use the `search` query string parameter to perform a case-insensitive search
within all designated text fields of a model.

	?search=findme

(_Added in Ansible Tower 3.1.0_) Search across related fields:

	?related__search=findme
*/
func (a *Client) WorkflowJobsWorkflowJobsActivityStreamList(params *WorkflowJobsWorkflowJobsActivityStreamListParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsActivityStreamListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowJobsWorkflowJobsActivityStreamListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Workflow Jobs_workflow_jobs_activity_stream_list",
		Method:             "GET",
		PathPattern:        "/api/v2/workflow_jobs/{id}/activity_stream/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowJobsWorkflowJobsActivityStreamListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WorkflowJobsWorkflowJobsActivityStreamListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Workflow Jobs_workflow_jobs_activity_stream_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	WorkflowJobsWorkflowJobsCancelCreate cancels workflow job

Make a GET request to this resource to determine if the workflow job can be
canceled. The response will include the following field:

  - `can_cancel`: Indicates whether this workflow job is in a state that can
    be canceled (boolean, read-only)

Make a POST request to this endpoint to submit a request to cancel a pending
or running workflow job.  The response status code will be 202 if the
request to cancel was successfully submitted, or 405 if the workflow job
cannot be canceled.
*/
func (a *Client) WorkflowJobsWorkflowJobsCancelCreate(params *WorkflowJobsWorkflowJobsCancelCreateParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsCancelCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowJobsWorkflowJobsCancelCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Workflow Jobs_workflow_jobs_cancel_create",
		Method:             "POST",
		PathPattern:        "/api/v2/workflow_jobs/{id}/cancel/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowJobsWorkflowJobsCancelCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WorkflowJobsWorkflowJobsCancelCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Workflow Jobs_workflow_jobs_cancel_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	WorkflowJobsWorkflowJobsCancelRead cancels workflow job

Make a GET request to this resource to determine if the workflow job can be
canceled. The response will include the following field:

  - `can_cancel`: Indicates whether this workflow job is in a state that can
    be canceled (boolean, read-only)

Make a POST request to this endpoint to submit a request to cancel a pending
or running workflow job.  The response status code will be 202 if the
request to cancel was successfully submitted, or 405 if the workflow job
cannot be canceled.
*/
func (a *Client) WorkflowJobsWorkflowJobsCancelRead(params *WorkflowJobsWorkflowJobsCancelReadParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsCancelReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowJobsWorkflowJobsCancelReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Workflow Jobs_workflow_jobs_cancel_read",
		Method:             "GET",
		PathPattern:        "/api/v2/workflow_jobs/{id}/cancel/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowJobsWorkflowJobsCancelReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WorkflowJobsWorkflowJobsCancelReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Workflow Jobs_workflow_jobs_cancel_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	WorkflowJobsWorkflowJobsDelete deletes a workflow job

Make a DELETE request to this resource to delete this workflow job.
*/
func (a *Client) WorkflowJobsWorkflowJobsDelete(params *WorkflowJobsWorkflowJobsDeleteParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowJobsWorkflowJobsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Workflow Jobs_workflow_jobs_delete",
		Method:             "DELETE",
		PathPattern:        "/api/v2/workflow_jobs/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowJobsWorkflowJobsDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WorkflowJobsWorkflowJobsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Workflow Jobs_workflow_jobs_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	WorkflowJobsWorkflowJobsLabelsList lists labels for a workflow job

Make a GET request to this resource to retrieve a list of
labels associated with the selected
workflow job.

The resulting data structure contains:

	{
	    "count": 99,
	    "next": null,
	    "previous": null,
	    "results": [
	        ...
	    ]
	}

The `count` field indicates the total number of labels
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more label records.

## Results

Each label data structure includes the following fields:

* `id`: Database ID for this label. (integer)
* `type`: Data type for this label. (choice)
* `url`: URL for this label. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this label was created. (datetime)
* `modified`: Timestamp when this label was last modified. (datetime)
* `name`: Name of this label. (string)
* `organization`: Organization this label belongs to. (id)

## Sorting

To specify that labels are returned in a particular
order, use the `order_by` query string parameter on the GET request.

	?order_by=name

Prefix the field name with a dash `-` to sort in reverse:

	?order_by=-name

Multiple sorting fields may be specified by separating the field names with a
comma `,`:

	?order_by=name,some_other_field

## Pagination

Use the `page_size` query string parameter to change the number of results
returned for each request.  Use the `page` query string parameter to retrieve
a particular page of results.

	?page_size=100&page=2

The `previous` and `next` links returned with the results will set these query
string parameters automatically.

## Searching

Use the `search` query string parameter to perform a case-insensitive search
within all designated text fields of a model.

	?search=findme

(_Added in Ansible Tower 3.1.0_) Search across related fields:

	?related__search=findme
*/
func (a *Client) WorkflowJobsWorkflowJobsLabelsList(params *WorkflowJobsWorkflowJobsLabelsListParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsLabelsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowJobsWorkflowJobsLabelsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Workflow Jobs_workflow_jobs_labels_list",
		Method:             "GET",
		PathPattern:        "/api/v2/workflow_jobs/{id}/labels/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowJobsWorkflowJobsLabelsListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WorkflowJobsWorkflowJobsLabelsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Workflow Jobs_workflow_jobs_labels_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	WorkflowJobsWorkflowJobsList lists workflow jobs

Make a GET request to this resource to retrieve the list of
workflow jobs.

The resulting data structure contains:

	{
	    "count": 99,
	    "next": null,
	    "previous": null,
	    "results": [
	        ...
	    ]
	}

The `count` field indicates the total number of workflow jobs
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more workflow job records.

## Results

Each workflow job data structure includes the following fields:

* `id`: Database ID for this workflow job. (integer)
* `type`: Data type for this workflow job. (choice)
* `url`: URL for this workflow job. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this workflow job was created. (datetime)
* `modified`: Timestamp when this workflow job was last modified. (datetime)
* `name`: Name of this workflow job. (string)
* `description`: Optional description of this workflow job. (string)
* `unified_job_template`:  (id)
* `launch_type`:  (choice)
  - `manual`: Manual
  - `relaunch`: Relaunch
  - `callback`: Callback
  - `scheduled`: Scheduled
  - `dependency`: Dependency
  - `workflow`: Workflow
  - `webhook`: Webhook
  - `sync`: Sync
  - `scm`: SCM Update

* `status`:  (choice)
  - `new`: New
  - `pending`: Pending
  - `waiting`: Waiting
  - `running`: Running
  - `successful`: Successful
  - `failed`: Failed
  - `error`: Error
  - `canceled`: Canceled

* `failed`:  (boolean)
* `started`: The date and time the job was queued for starting. (datetime)
* `finished`: The date and time the job finished execution. (datetime)
* `canceled_on`: The date and time when the cancel request was sent. (datetime)
* `elapsed`: Elapsed time in seconds that the job ran. (decimal)
* `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string)
* `launched_by`:  (field)
* `work_unit_id`: The Receptor work unit ID associated with this job. (string)
* `workflow_job_template`:  (id)
* `extra_vars`:  (json)
* `allow_simultaneous`:  (boolean)
* `job_template`: If automatically created for a sliced job run, the job template the workflow job was created from. (id)
* `is_sliced_job`:  (boolean)
* `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id)
* `limit`:  (string)
* `scm_branch`:  (string)
* `webhook_service`: Service that webhook requests will be accepted from (choice)
  - `""`: ---------
  - `github`: GitHub
  - `gitlab`: GitLab

* `webhook_credential`: Personal Access Token for posting back the status to the service API (id)
* `webhook_guid`: Unique identifier of the event that triggered this webhook (string)

## Sorting

To specify that workflow jobs are returned in a particular
order, use the `order_by` query string parameter on the GET request.

	?order_by=name

Prefix the field name with a dash `-` to sort in reverse:

	?order_by=-name

Multiple sorting fields may be specified by separating the field names with a
comma `,`:

	?order_by=name,some_other_field

## Pagination

Use the `page_size` query string parameter to change the number of results
returned for each request.  Use the `page` query string parameter to retrieve
a particular page of results.

	?page_size=100&page=2

The `previous` and `next` links returned with the results will set these query
string parameters automatically.

## Searching

Use the `search` query string parameter to perform a case-insensitive search
within all designated text fields of a model.

	?search=findme

(_Added in Ansible Tower 3.1.0_) Search across related fields:

	?related__search=findme
*/
func (a *Client) WorkflowJobsWorkflowJobsList(params *WorkflowJobsWorkflowJobsListParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowJobsWorkflowJobsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Workflow Jobs_workflow_jobs_list",
		Method:             "GET",
		PathPattern:        "/api/v2/workflow_jobs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowJobsWorkflowJobsListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WorkflowJobsWorkflowJobsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Workflow Jobs_workflow_jobs_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	WorkflowJobsWorkflowJobsNotificationsList lists notifications for a workflow job

Make a GET request to this resource to retrieve a list of
notifications associated with the selected
workflow job.

The resulting data structure contains:

	{
	    "count": 99,
	    "next": null,
	    "previous": null,
	    "results": [
	        ...
	    ]
	}

The `count` field indicates the total number of notifications
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more notification records.

## Results

Each notification data structure includes the following fields:

* `id`: Database ID for this notification. (integer)
* `type`: Data type for this notification. (choice)
* `url`: URL for this notification. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this notification was created. (datetime)
* `modified`: Timestamp when this notification was last modified. (datetime)
* `notification_template`:  (id)
* `error`:  (string)
* `status`:  (choice)
  - `pending`: Pending
  - `successful`: Successful
  - `failed`: Failed

* `notifications_sent`:  (integer)
* `notification_type`:  (choice)
  - `email`: Email
  - `grafana`: Grafana
  - `irc`: IRC
  - `mattermost`: Mattermost
  - `pagerduty`: Pagerduty
  - `rocketchat`: Rocket.Chat
  - `slack`: Slack
  - `twilio`: Twilio
  - `webhook`: Webhook

* `recipients`:  (string)
* `subject`:  (string)
* `body`: Notification body (json)

## Sorting

To specify that notifications are returned in a particular
order, use the `order_by` query string parameter on the GET request.

	?order_by=name

Prefix the field name with a dash `-` to sort in reverse:

	?order_by=-name

Multiple sorting fields may be specified by separating the field names with a
comma `,`:

	?order_by=name,some_other_field

## Pagination

Use the `page_size` query string parameter to change the number of results
returned for each request.  Use the `page` query string parameter to retrieve
a particular page of results.

	?page_size=100&page=2

The `previous` and `next` links returned with the results will set these query
string parameters automatically.

## Searching

Use the `search` query string parameter to perform a case-insensitive search
within all designated text fields of a model.

	?search=findme

(_Added in Ansible Tower 3.1.0_) Search across related fields:

	?related__search=findme
*/
func (a *Client) WorkflowJobsWorkflowJobsNotificationsList(params *WorkflowJobsWorkflowJobsNotificationsListParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsNotificationsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowJobsWorkflowJobsNotificationsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Workflow Jobs_workflow_jobs_notifications_list",
		Method:             "GET",
		PathPattern:        "/api/v2/workflow_jobs/{id}/notifications/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowJobsWorkflowJobsNotificationsListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WorkflowJobsWorkflowJobsNotificationsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Workflow Jobs_workflow_jobs_notifications_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	WorkflowJobsWorkflowJobsRead retrieves a workflow job

Make GET request to this resource to retrieve a single workflow job
record containing the following fields:

* `id`: Database ID for this workflow job. (integer)
* `type`: Data type for this workflow job. (choice)
* `url`: URL for this workflow job. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this workflow job was created. (datetime)
* `modified`: Timestamp when this workflow job was last modified. (datetime)
* `name`: Name of this workflow job. (string)
* `description`: Optional description of this workflow job. (string)
* `unified_job_template`:  (id)
* `launch_type`:  (choice)
  - `manual`: Manual
  - `relaunch`: Relaunch
  - `callback`: Callback
  - `scheduled`: Scheduled
  - `dependency`: Dependency
  - `workflow`: Workflow
  - `webhook`: Webhook
  - `sync`: Sync
  - `scm`: SCM Update

* `status`:  (choice)
  - `new`: New
  - `pending`: Pending
  - `waiting`: Waiting
  - `running`: Running
  - `successful`: Successful
  - `failed`: Failed
  - `error`: Error
  - `canceled`: Canceled

* `failed`:  (boolean)
* `started`: The date and time the job was queued for starting. (datetime)
* `finished`: The date and time the job finished execution. (datetime)
* `canceled_on`: The date and time when the cancel request was sent. (datetime)
* `elapsed`: Elapsed time in seconds that the job ran. (decimal)
* `job_args`:  (string)
* `job_cwd`:  (string)
* `job_env`:  (json)
* `job_explanation`: A status field to indicate the state of the job if it wasn&#39;t able to run and capture stdout (string)
* `result_traceback`:  (string)
* `launched_by`:  (field)
* `work_unit_id`: The Receptor work unit ID associated with this job. (string)
* `workflow_job_template`:  (id)
* `extra_vars`:  (json)
* `allow_simultaneous`:  (boolean)
* `job_template`: If automatically created for a sliced job run, the job template the workflow job was created from. (id)
* `is_sliced_job`:  (boolean)
* `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id)
* `limit`:  (string)
* `scm_branch`:  (string)
* `webhook_service`: Service that webhook requests will be accepted from (choice)
  - `""`: ---------
  - `github`: GitHub
  - `gitlab`: GitLab

* `webhook_credential`: Personal Access Token for posting back the status to the service API (id)
* `webhook_guid`: Unique identifier of the event that triggered this webhook (string)
*/
func (a *Client) WorkflowJobsWorkflowJobsRead(params *WorkflowJobsWorkflowJobsReadParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowJobsWorkflowJobsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Workflow Jobs_workflow_jobs_read",
		Method:             "GET",
		PathPattern:        "/api/v2/workflow_jobs/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowJobsWorkflowJobsReadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WorkflowJobsWorkflowJobsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Workflow Jobs_workflow_jobs_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	WorkflowJobsWorkflowJobsRelaunchCreate relaunches a workflow job

Make a POST request to this endpoint to launch a workflow job identical to the parent workflow job. This will spawn jobs, project updates, or inventory updates based on the unified job templates referenced in the workflow nodes in the workflow job. No POST data is accepted for this action.

If successful, the response status code will be 201 and serialized data of the new workflow job will be returned.
*/
func (a *Client) WorkflowJobsWorkflowJobsRelaunchCreate(params *WorkflowJobsWorkflowJobsRelaunchCreateParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsRelaunchCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowJobsWorkflowJobsRelaunchCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Workflow Jobs_workflow_jobs_relaunch_create",
		Method:             "POST",
		PathPattern:        "/api/v2/workflow_jobs/{id}/relaunch/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowJobsWorkflowJobsRelaunchCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WorkflowJobsWorkflowJobsRelaunchCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Workflow Jobs_workflow_jobs_relaunch_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	WorkflowJobsWorkflowJobsRelaunchList relaunches a workflow job

Make a POST request to this endpoint to launch a workflow job identical to the parent workflow job. This will spawn jobs, project updates, or inventory updates based on the unified job templates referenced in the workflow nodes in the workflow job. No POST data is accepted for this action.

If successful, the response status code will be 201 and serialized data of the new workflow job will be returned.
*/
func (a *Client) WorkflowJobsWorkflowJobsRelaunchList(params *WorkflowJobsWorkflowJobsRelaunchListParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsRelaunchListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowJobsWorkflowJobsRelaunchListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Workflow Jobs_workflow_jobs_relaunch_list",
		Method:             "GET",
		PathPattern:        "/api/v2/workflow_jobs/{id}/relaunch/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowJobsWorkflowJobsRelaunchListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WorkflowJobsWorkflowJobsRelaunchListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Workflow Jobs_workflow_jobs_relaunch_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	WorkflowJobsWorkflowJobsWorkflowNodesList lists workflow job nodes for a workflow job

Make a GET request to this resource to retrieve a list of
workflow job nodes associated with the selected
workflow job.

The resulting data structure contains:

	{
	    "count": 99,
	    "next": null,
	    "previous": null,
	    "results": [
	        ...
	    ]
	}

The `count` field indicates the total number of workflow job nodes
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more workflow job node records.

## Results

Each workflow job node data structure includes the following fields:

* `id`: Database ID for this workflow job node. (integer)
* `type`: Data type for this workflow job node. (choice)
* `url`: URL for this workflow job node. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this workflow job node was created. (datetime)
* `modified`: Timestamp when this workflow job node was last modified. (datetime)
* `extra_data`:  (json)
* `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id)
* `scm_branch`:  (string)
* `job_type`:  (choice)
  - `None`: ---------
  - `""`: ---------
  - `run`: Run
  - `check`: Check

* `job_tags`:  (string)
* `skip_tags`:  (string)
* `limit`:  (string)
* `diff_mode`:  (boolean)
* `verbosity`:  (choice)
  - `None`: ---------
  - `0`: 0 (Normal)
  - `1`: 1 (Verbose)
  - `2`: 2 (More Verbose)
  - `3`: 3 (Debug)
  - `4`: 4 (Connection Debug)
  - `5`: 5 (WinRM Debug)

* `job`:  (id)
* `workflow_job`:  (id)
* `unified_job_template`:  (id)
* `success_nodes`:  (field)
* `failure_nodes`:  (field)
* `always_nodes`:  (field)
* `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean)
* `do_not_run`: Indicates that a job will not be created when True. Workflow runtime semantics will mark this True if the node is in a path that will decidedly not be ran. A value of False means the node may not run. (boolean)
* `identifier`: An identifier coresponding to the workflow job template node that this node was created from. (string)

## Sorting

To specify that workflow job nodes are returned in a particular
order, use the `order_by` query string parameter on the GET request.

	?order_by=name

Prefix the field name with a dash `-` to sort in reverse:

	?order_by=-name

Multiple sorting fields may be specified by separating the field names with a
comma `,`:

	?order_by=name,some_other_field

## Pagination

Use the `page_size` query string parameter to change the number of results
returned for each request.  Use the `page` query string parameter to retrieve
a particular page of results.

	?page_size=100&page=2

The `previous` and `next` links returned with the results will set these query
string parameters automatically.

## Searching

Use the `search` query string parameter to perform a case-insensitive search
within all designated text fields of a model.

	?search=findme

(_Added in Ansible Tower 3.1.0_) Search across related fields:

	?related__search=findme
*/
func (a *Client) WorkflowJobsWorkflowJobsWorkflowNodesList(params *WorkflowJobsWorkflowJobsWorkflowNodesListParams, opts ...ClientOption) (*WorkflowJobsWorkflowJobsWorkflowNodesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkflowJobsWorkflowJobsWorkflowNodesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Workflow Jobs_workflow_jobs_workflow_nodes_list",
		Method:             "GET",
		PathPattern:        "/api/v2/workflow_jobs/{id}/workflow_nodes/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkflowJobsWorkflowJobsWorkflowNodesListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WorkflowJobsWorkflowJobsWorkflowNodesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Workflow Jobs_workflow_jobs_workflow_nodes_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
