// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RolesRolesReadReader is a Reader for the RolesRolesRead structure.
type RolesRolesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RolesRolesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRolesRolesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRolesRolesReadOK creates a RolesRolesReadOK with default headers values
func NewRolesRolesReadOK() *RolesRolesReadOK {
	return &RolesRolesReadOK{}
}

/*
RolesRolesReadOK describes a response with status code 200, with default header values.

RolesRolesReadOK roles roles read o k
*/
type RolesRolesReadOK struct {
}

// IsSuccess returns true when this roles roles read o k response has a 2xx status code
func (o *RolesRolesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this roles roles read o k response has a 3xx status code
func (o *RolesRolesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this roles roles read o k response has a 4xx status code
func (o *RolesRolesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this roles roles read o k response has a 5xx status code
func (o *RolesRolesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this roles roles read o k response a status code equal to that given
func (o *RolesRolesReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *RolesRolesReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/roles/{id}/][%d] rolesRolesReadOK ", 200)
}

func (o *RolesRolesReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/roles/{id}/][%d] rolesRolesReadOK ", 200)
}

func (o *RolesRolesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}