// Code generated by go-swagger; DO NOT EDIT.

package job_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new job events API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for job events API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	JobEventsJobEventsChildrenList(params *JobEventsJobEventsChildrenListParams, opts ...ClientOption) (*JobEventsJobEventsChildrenListOK, error)

	JobEventsJobEventsRead(params *JobEventsJobEventsReadParams, opts ...ClientOption) (*JobEventsJobEventsReadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	JobEventsJobEventsChildrenList lists job events for a job event

Make a GET request to this resource to retrieve a list of
job events associated with the selected
job event.

The resulting data structure contains:

	{
	    "count": 99,
	    "next": null,
	    "previous": null,
	    "results": [
	        ...
	    ]
	}

The `count` field indicates the total number of job events
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more job event records.

## Results

Each job event data structure includes the following fields:

* `id`: Database ID for this job event. (integer)
* `type`: Data type for this job event. (choice)
* `url`: URL for this job event. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this job event was created. (datetime)
* `modified`: Timestamp when this job event was last modified. (datetime)
* `job`:  (id)
* `event`:  (choice)
  - `runner_on_failed`: Host Failed
  - `runner_on_start`: Host Started
  - `runner_on_ok`: Host OK
  - `runner_on_error`: Host Failure
  - `runner_on_skipped`: Host Skipped
  - `runner_on_unreachable`: Host Unreachable
  - `runner_on_no_hosts`: No Hosts Remaining
  - `runner_on_async_poll`: Host Polling
  - `runner_on_async_ok`: Host Async OK
  - `runner_on_async_failed`: Host Async Failure
  - `runner_item_on_ok`: Item OK
  - `runner_item_on_failed`: Item Failed
  - `runner_item_on_skipped`: Item Skipped
  - `runner_retry`: Host Retry
  - `runner_on_file_diff`: File Difference
  - `playbook_on_start`: Playbook Started
  - `playbook_on_notify`: Running Handlers
  - `playbook_on_include`: Including File
  - `playbook_on_no_hosts_matched`: No Hosts Matched
  - `playbook_on_no_hosts_remaining`: No Hosts Remaining
  - `playbook_on_task_start`: Task Started
  - `playbook_on_vars_prompt`: Variables Prompted
  - `playbook_on_setup`: Gathering Facts
  - `playbook_on_import_for_host`: internal: on Import for Host
  - `playbook_on_not_import_for_host`: internal: on Not Import for Host
  - `playbook_on_play_start`: Play Started
  - `playbook_on_stats`: Playbook Complete
  - `debug`: Debug
  - `verbose`: Verbose
  - `deprecated`: Deprecated
  - `warning`: Warning
  - `system_warning`: System Warning
  - `error`: Error

* `counter`:  (integer)
* `event_display`:  (string)
* `event_data`:  (json)
* `event_level`:  (integer)
* `failed`:  (boolean)
* `changed`:  (boolean)
* `uuid`:  (string)
* `parent_uuid`:  (string)
* `host`:  (id)
* `host_name`:  (string)
* `playbook`:  (string)
* `play`:  (string)
* `task`:  (string)
* `role`:  (string)
* `stdout`:  (string)
* `start_line`:  (integer)
* `end_line`:  (integer)
* `verbosity`:  (integer)

## Sorting

To specify that job events are returned in a particular
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
func (a *Client) JobEventsJobEventsChildrenList(params *JobEventsJobEventsChildrenListParams, opts ...ClientOption) (*JobEventsJobEventsChildrenListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobEventsJobEventsChildrenListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Job Events_job_events_children_list",
		Method:             "GET",
		PathPattern:        "/api/v2/job_events/{id}/children/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JobEventsJobEventsChildrenListReader{formats: a.formats},
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
	success, ok := result.(*JobEventsJobEventsChildrenListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Job Events_job_events_children_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	JobEventsJobEventsRead retrieves a job event

Make GET request to this resource to retrieve a single job event
record containing the following fields:

* `id`: Database ID for this job event. (integer)
* `type`: Data type for this job event. (choice)
* `url`: URL for this job event. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this job event was created. (datetime)
* `modified`: Timestamp when this job event was last modified. (datetime)
* `job`:  (id)
* `event`:  (choice)
  - `runner_on_failed`: Host Failed
  - `runner_on_start`: Host Started
  - `runner_on_ok`: Host OK
  - `runner_on_error`: Host Failure
  - `runner_on_skipped`: Host Skipped
  - `runner_on_unreachable`: Host Unreachable
  - `runner_on_no_hosts`: No Hosts Remaining
  - `runner_on_async_poll`: Host Polling
  - `runner_on_async_ok`: Host Async OK
  - `runner_on_async_failed`: Host Async Failure
  - `runner_item_on_ok`: Item OK
  - `runner_item_on_failed`: Item Failed
  - `runner_item_on_skipped`: Item Skipped
  - `runner_retry`: Host Retry
  - `runner_on_file_diff`: File Difference
  - `playbook_on_start`: Playbook Started
  - `playbook_on_notify`: Running Handlers
  - `playbook_on_include`: Including File
  - `playbook_on_no_hosts_matched`: No Hosts Matched
  - `playbook_on_no_hosts_remaining`: No Hosts Remaining
  - `playbook_on_task_start`: Task Started
  - `playbook_on_vars_prompt`: Variables Prompted
  - `playbook_on_setup`: Gathering Facts
  - `playbook_on_import_for_host`: internal: on Import for Host
  - `playbook_on_not_import_for_host`: internal: on Not Import for Host
  - `playbook_on_play_start`: Play Started
  - `playbook_on_stats`: Playbook Complete
  - `debug`: Debug
  - `verbose`: Verbose
  - `deprecated`: Deprecated
  - `warning`: Warning
  - `system_warning`: System Warning
  - `error`: Error

* `counter`:  (integer)
* `event_display`:  (string)
* `event_data`:  (json)
* `event_level`:  (integer)
* `failed`:  (boolean)
* `changed`:  (boolean)
* `uuid`:  (string)
* `parent_uuid`:  (string)
* `host`:  (id)
* `host_name`:  (string)
* `playbook`:  (string)
* `play`:  (string)
* `task`:  (string)
* `role`:  (string)
* `stdout`:  (string)
* `start_line`:  (integer)
* `end_line`:  (integer)
* `verbosity`:  (integer)
*/
func (a *Client) JobEventsJobEventsRead(params *JobEventsJobEventsReadParams, opts ...ClientOption) (*JobEventsJobEventsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobEventsJobEventsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Job Events_job_events_read",
		Method:             "GET",
		PathPattern:        "/api/v2/job_events/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JobEventsJobEventsReadReader{formats: a.formats},
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
	success, ok := result.(*JobEventsJobEventsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Job Events_job_events_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
