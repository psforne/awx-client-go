// Code generated by go-swagger; DO NOT EDIT.

package ad_hoc_command_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ad hoc command events API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ad hoc command events API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AdHocCommandEventsAdHocCommandEventsRead(params *AdHocCommandEventsAdHocCommandEventsReadParams, opts ...ClientOption) (*AdHocCommandEventsAdHocCommandEventsReadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	AdHocCommandEventsAdHocCommandEventsRead retrieves an ad hoc command event

Make GET request to this resource to retrieve a single ad hoc command event
record containing the following fields:

* `id`: Database ID for this ad hoc command event. (integer)
* `type`: Data type for this ad hoc command event. (choice)
* `url`: URL for this ad hoc command event. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this ad hoc command event was created. (datetime)
* `modified`: Timestamp when this ad hoc command event was last modified. (datetime)
* `ad_hoc_command`:  (id)
* `event`:  (choice)
  - `runner_on_failed`: Host Failed
  - `runner_on_ok`: Host OK
  - `runner_on_unreachable`: Host Unreachable
  - `runner_on_skipped`: Host Skipped
  - `debug`: Debug
  - `verbose`: Verbose
  - `deprecated`: Deprecated
  - `warning`: Warning
  - `system_warning`: System Warning
  - `error`: Error

* `counter`:  (integer)
* `event_display`:  (string)
* `event_data`:  (json)
* `failed`:  (boolean)
* `changed`:  (boolean)
* `uuid`:  (string)
* `host`:  (id)
* `host_name`:  (string)
* `stdout`:  (string)
* `start_line`:  (integer)
* `end_line`:  (integer)
* `verbosity`:  (integer)
*/
func (a *Client) AdHocCommandEventsAdHocCommandEventsRead(params *AdHocCommandEventsAdHocCommandEventsReadParams, opts ...ClientOption) (*AdHocCommandEventsAdHocCommandEventsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdHocCommandEventsAdHocCommandEventsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Ad Hoc Command Events_ad_hoc_command_events_read",
		Method:             "GET",
		PathPattern:        "/api/v2/ad_hoc_command_events/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AdHocCommandEventsAdHocCommandEventsReadReader{formats: a.formats},
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
	success, ok := result.(*AdHocCommandEventsAdHocCommandEventsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Ad Hoc Command Events_ad_hoc_command_events_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
