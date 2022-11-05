// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RolesRolesUsersListReader is a Reader for the RolesRolesUsersList structure.
type RolesRolesUsersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RolesRolesUsersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRolesRolesUsersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRolesRolesUsersListOK creates a RolesRolesUsersListOK with default headers values
func NewRolesRolesUsersListOK() *RolesRolesUsersListOK {
	return &RolesRolesUsersListOK{}
}

/*
RolesRolesUsersListOK describes a response with status code 200, with default header values.

RolesRolesUsersListOK roles roles users list o k
*/
type RolesRolesUsersListOK struct {
}

// IsSuccess returns true when this roles roles users list o k response has a 2xx status code
func (o *RolesRolesUsersListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this roles roles users list o k response has a 3xx status code
func (o *RolesRolesUsersListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this roles roles users list o k response has a 4xx status code
func (o *RolesRolesUsersListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this roles roles users list o k response has a 5xx status code
func (o *RolesRolesUsersListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this roles roles users list o k response a status code equal to that given
func (o *RolesRolesUsersListOK) IsCode(code int) bool {
	return code == 200
}

func (o *RolesRolesUsersListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/roles/{id}/users/][%d] rolesRolesUsersListOK ", 200)
}

func (o *RolesRolesUsersListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/roles/{id}/users/][%d] rolesRolesUsersListOK ", 200)
}

func (o *RolesRolesUsersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
