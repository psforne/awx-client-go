// Code generated by go-swagger; DO NOT EDIT.

package activity_streams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new activity streams API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for activity streams API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ActivityStreamsActivityStreamList(params *ActivityStreamsActivityStreamListParams, opts ...ClientOption) (*ActivityStreamsActivityStreamListOK, error)

	ActivityStreamsActivityStreamRead(params *ActivityStreamsActivityStreamReadParams, opts ...ClientOption) (*ActivityStreamsActivityStreamReadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	ActivityStreamsActivityStreamList lists activity streams

Make a GET request to this resource to retrieve the list of
activity streams.

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

	?order_by=

Prefix the field name with a dash `-` to sort in reverse:

	?order_by=-

Multiple sorting fields may be specified by separating the field names with a
comma `,`:

	?order_by=,some_other_field

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
func (a *Client) ActivityStreamsActivityStreamList(params *ActivityStreamsActivityStreamListParams, opts ...ClientOption) (*ActivityStreamsActivityStreamListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActivityStreamsActivityStreamListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Activity Streams_activity_stream_list",
		Method:             "GET",
		PathPattern:        "/api/v2/activity_stream/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActivityStreamsActivityStreamListReader{formats: a.formats},
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
	success, ok := result.(*ActivityStreamsActivityStreamListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Activity Streams_activity_stream_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ActivityStreamsActivityStreamRead retrieves an activity stream

Make GET request to this resource to retrieve a single activity stream
record containing the following fields:

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
*/
func (a *Client) ActivityStreamsActivityStreamRead(params *ActivityStreamsActivityStreamReadParams, opts ...ClientOption) (*ActivityStreamsActivityStreamReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActivityStreamsActivityStreamReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Activity Streams_activity_stream_read",
		Method:             "GET",
		PathPattern:        "/api/v2/activity_stream/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActivityStreamsActivityStreamReadReader{formats: a.formats},
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
	success, ok := result.(*ActivityStreamsActivityStreamReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Activity Streams_activity_stream_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
