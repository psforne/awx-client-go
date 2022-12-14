// Code generated by go-swagger; DO NOT EDIT.

package versioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new versioning API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for versioning API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	VersioningList(params *VersioningListParams, opts ...ClientOption) (*VersioningListOK, error)

	VersioningRead(params *VersioningReadParams, opts ...ClientOption) (*VersioningReadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
VersioningList lists supported API versions
*/
func (a *Client) VersioningList(params *VersioningListParams, opts ...ClientOption) (*VersioningListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVersioningListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Versioning_list",
		Method:             "GET",
		PathPattern:        "/api/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VersioningListReader{formats: a.formats},
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
	success, ok := result.(*VersioningListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Versioning_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VersioningRead lists top level resources
*/
func (a *Client) VersioningRead(params *VersioningReadParams, opts ...ClientOption) (*VersioningReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVersioningReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Versioning_read",
		Method:             "GET",
		PathPattern:        "/api/v2/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VersioningReadReader{formats: a.formats},
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
	success, ok := result.(*VersioningReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Versioning_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
