// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersUsersListReader is a Reader for the UsersUsersList structure.
type UsersUsersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUsersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersUsersListOK creates a UsersUsersListOK with default headers values
func NewUsersUsersListOK() *UsersUsersListOK {
	return &UsersUsersListOK{}
}

/*
UsersUsersListOK describes a response with status code 200, with default header values.

UsersUsersListOK users users list o k
*/
type UsersUsersListOK struct {
}

// IsSuccess returns true when this users users list o k response has a 2xx status code
func (o *UsersUsersListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users users list o k response has a 3xx status code
func (o *UsersUsersListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users list o k response has a 4xx status code
func (o *UsersUsersListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users users list o k response has a 5xx status code
func (o *UsersUsersListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users users list o k response a status code equal to that given
func (o *UsersUsersListOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersUsersListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/][%d] usersUsersListOK ", 200)
}

func (o *UsersUsersListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/users/][%d] usersUsersListOK ", 200)
}

func (o *UsersUsersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
