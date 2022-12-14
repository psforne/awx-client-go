// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersUsersAccessListListReader is a Reader for the UsersUsersAccessListList structure.
type UsersUsersAccessListListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersAccessListListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUsersAccessListListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersUsersAccessListListOK creates a UsersUsersAccessListListOK with default headers values
func NewUsersUsersAccessListListOK() *UsersUsersAccessListListOK {
	return &UsersUsersAccessListListOK{}
}

/*
UsersUsersAccessListListOK describes a response with status code 200, with default header values.

UsersUsersAccessListListOK users users access list list o k
*/
type UsersUsersAccessListListOK struct {
}

// IsSuccess returns true when this users users access list list o k response has a 2xx status code
func (o *UsersUsersAccessListListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users users access list list o k response has a 3xx status code
func (o *UsersUsersAccessListListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users access list list o k response has a 4xx status code
func (o *UsersUsersAccessListListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users users access list list o k response has a 5xx status code
func (o *UsersUsersAccessListListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users users access list list o k response a status code equal to that given
func (o *UsersUsersAccessListListOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersUsersAccessListListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{id}/access_list/][%d] usersUsersAccessListListOK ", 200)
}

func (o *UsersUsersAccessListListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{id}/access_list/][%d] usersUsersAccessListListOK ", 200)
}

func (o *UsersUsersAccessListListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
