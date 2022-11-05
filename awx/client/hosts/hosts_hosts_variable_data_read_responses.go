// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HostsHostsVariableDataReadReader is a Reader for the HostsHostsVariableDataRead structure.
type HostsHostsVariableDataReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostsHostsVariableDataReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHostsHostsVariableDataReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHostsHostsVariableDataReadOK creates a HostsHostsVariableDataReadOK with default headers values
func NewHostsHostsVariableDataReadOK() *HostsHostsVariableDataReadOK {
	return &HostsHostsVariableDataReadOK{}
}

/*
HostsHostsVariableDataReadOK describes a response with status code 200, with default header values.

HostsHostsVariableDataReadOK hosts hosts variable data read o k
*/
type HostsHostsVariableDataReadOK struct {
}

// IsSuccess returns true when this hosts hosts variable data read o k response has a 2xx status code
func (o *HostsHostsVariableDataReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hosts hosts variable data read o k response has a 3xx status code
func (o *HostsHostsVariableDataReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hosts hosts variable data read o k response has a 4xx status code
func (o *HostsHostsVariableDataReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hosts hosts variable data read o k response has a 5xx status code
func (o *HostsHostsVariableDataReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hosts hosts variable data read o k response a status code equal to that given
func (o *HostsHostsVariableDataReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *HostsHostsVariableDataReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/hosts/{id}/variable_data/][%d] hostsHostsVariableDataReadOK ", 200)
}

func (o *HostsHostsVariableDataReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/hosts/{id}/variable_data/][%d] hostsHostsVariableDataReadOK ", 200)
}

func (o *HostsHostsVariableDataReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
