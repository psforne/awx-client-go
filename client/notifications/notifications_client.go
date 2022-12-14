// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new notifications API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for notifications API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	NotificationsNotificationsList(params *NotificationsNotificationsListParams, opts ...ClientOption) (*NotificationsNotificationsListOK, error)

	NotificationsNotificationsRead(params *NotificationsNotificationsReadParams, opts ...ClientOption) (*NotificationsNotificationsReadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	NotificationsNotificationsList lists notifications

Make a GET request to this resource to retrieve the list of
notifications.

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
func (a *Client) NotificationsNotificationsList(params *NotificationsNotificationsListParams, opts ...ClientOption) (*NotificationsNotificationsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationsNotificationsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Notifications_notifications_list",
		Method:             "GET",
		PathPattern:        "/api/v2/notifications/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationsNotificationsListReader{formats: a.formats},
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
	success, ok := result.(*NotificationsNotificationsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Notifications_notifications_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	NotificationsNotificationsRead retrieves a notification

Make GET request to this resource to retrieve a single notification
record containing the following fields:

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
*/
func (a *Client) NotificationsNotificationsRead(params *NotificationsNotificationsReadParams, opts ...ClientOption) (*NotificationsNotificationsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNotificationsNotificationsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Notifications_notifications_read",
		Method:             "GET",
		PathPattern:        "/api/v2/notifications/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NotificationsNotificationsReadReader{formats: a.formats},
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
	success, ok := result.(*NotificationsNotificationsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Notifications_notifications_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
