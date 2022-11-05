// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HostsHostsPartialUpdateReader is a Reader for the HostsHostsPartialUpdate structure.
type HostsHostsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostsHostsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHostsHostsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHostsHostsPartialUpdateOK creates a HostsHostsPartialUpdateOK with default headers values
func NewHostsHostsPartialUpdateOK() *HostsHostsPartialUpdateOK {
	return &HostsHostsPartialUpdateOK{}
}

/*
HostsHostsPartialUpdateOK describes a response with status code 200, with default header values.

HostsHostsPartialUpdateOK hosts hosts partial update o k
*/
type HostsHostsPartialUpdateOK struct {
}

// IsSuccess returns true when this hosts hosts partial update o k response has a 2xx status code
func (o *HostsHostsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hosts hosts partial update o k response has a 3xx status code
func (o *HostsHostsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hosts hosts partial update o k response has a 4xx status code
func (o *HostsHostsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hosts hosts partial update o k response has a 5xx status code
func (o *HostsHostsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hosts hosts partial update o k response a status code equal to that given
func (o *HostsHostsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *HostsHostsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/hosts/{id}/][%d] hostsHostsPartialUpdateOK ", 200)
}

func (o *HostsHostsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/hosts/{id}/][%d] hostsHostsPartialUpdateOK ", 200)
}

func (o *HostsHostsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
HostsHostsPartialUpdateBody hosts hosts partial update body
swagger:model HostsHostsPartialUpdateBody
*/
type HostsHostsPartialUpdateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// Is this host online and available for running jobs?
	Enabled bool `json:"enabled,omitempty"`

	// The value used by the remote inventory source to uniquely identify the host
	InstanceID string `json:"instance_id,omitempty"`

	// inventory
	Inventory int64 `json:"inventory,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Host variables in JSON or YAML format.
	Variables string `json:"variables,omitempty"`
}

// Validate validates this hosts hosts partial update body
func (o *HostsHostsPartialUpdateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hosts hosts partial update body based on context it is used
func (o *HostsHostsPartialUpdateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *HostsHostsPartialUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *HostsHostsPartialUpdateBody) UnmarshalBinary(b []byte) error {
	var res HostsHostsPartialUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
