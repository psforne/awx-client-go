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

// HostsHostsVariableDataPartialUpdateReader is a Reader for the HostsHostsVariableDataPartialUpdate structure.
type HostsHostsVariableDataPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostsHostsVariableDataPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHostsHostsVariableDataPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHostsHostsVariableDataPartialUpdateOK creates a HostsHostsVariableDataPartialUpdateOK with default headers values
func NewHostsHostsVariableDataPartialUpdateOK() *HostsHostsVariableDataPartialUpdateOK {
	return &HostsHostsVariableDataPartialUpdateOK{}
}

/*
HostsHostsVariableDataPartialUpdateOK describes a response with status code 200, with default header values.

HostsHostsVariableDataPartialUpdateOK hosts hosts variable data partial update o k
*/
type HostsHostsVariableDataPartialUpdateOK struct {
}

// IsSuccess returns true when this hosts hosts variable data partial update o k response has a 2xx status code
func (o *HostsHostsVariableDataPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hosts hosts variable data partial update o k response has a 3xx status code
func (o *HostsHostsVariableDataPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hosts hosts variable data partial update o k response has a 4xx status code
func (o *HostsHostsVariableDataPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hosts hosts variable data partial update o k response has a 5xx status code
func (o *HostsHostsVariableDataPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hosts hosts variable data partial update o k response a status code equal to that given
func (o *HostsHostsVariableDataPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *HostsHostsVariableDataPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/hosts/{id}/variable_data/][%d] hostsHostsVariableDataPartialUpdateOK ", 200)
}

func (o *HostsHostsVariableDataPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/hosts/{id}/variable_data/][%d] hostsHostsVariableDataPartialUpdateOK ", 200)
}

func (o *HostsHostsVariableDataPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
HostsHostsVariableDataPartialUpdateBody hosts hosts variable data partial update body
swagger:model HostsHostsVariableDataPartialUpdateBody
*/
type HostsHostsVariableDataPartialUpdateBody struct {

	// Host variables in JSON or YAML format.
	Variables string `json:"variables,omitempty"`
}

// Validate validates this hosts hosts variable data partial update body
func (o *HostsHostsVariableDataPartialUpdateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hosts hosts variable data partial update body based on context it is used
func (o *HostsHostsVariableDataPartialUpdateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *HostsHostsVariableDataPartialUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *HostsHostsVariableDataPartialUpdateBody) UnmarshalBinary(b []byte) error {
	var res HostsHostsVariableDataPartialUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
