// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersUsersRolesListReader is a Reader for the UsersUsersRolesList structure.
type UsersUsersRolesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersRolesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUsersRolesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersUsersRolesListOK creates a UsersUsersRolesListOK with default headers values
func NewUsersUsersRolesListOK() *UsersUsersRolesListOK {
	return &UsersUsersRolesListOK{}
}

/*
UsersUsersRolesListOK describes a response with status code 200, with default header values.

UsersUsersRolesListOK users users roles list o k
*/
type UsersUsersRolesListOK struct {
}

// IsSuccess returns true when this users users roles list o k response has a 2xx status code
func (o *UsersUsersRolesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users users roles list o k response has a 3xx status code
func (o *UsersUsersRolesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users roles list o k response has a 4xx status code
func (o *UsersUsersRolesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users users roles list o k response has a 5xx status code
func (o *UsersUsersRolesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users users roles list o k response a status code equal to that given
func (o *UsersUsersRolesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersUsersRolesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{id}/roles/][%d] usersUsersRolesListOK ", 200)
}

func (o *UsersUsersRolesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{id}/roles/][%d] usersUsersRolesListOK ", 200)
}

func (o *UsersUsersRolesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
