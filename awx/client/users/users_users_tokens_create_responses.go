// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersUsersTokensCreateReader is a Reader for the UsersUsersTokensCreate structure.
type UsersUsersTokensCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersTokensCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersUsersTokensCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersUsersTokensCreateCreated creates a UsersUsersTokensCreateCreated with default headers values
func NewUsersUsersTokensCreateCreated() *UsersUsersTokensCreateCreated {
	return &UsersUsersTokensCreateCreated{}
}

/*
UsersUsersTokensCreateCreated describes a response with status code 201, with default header values.

UsersUsersTokensCreateCreated users users tokens create created
*/
type UsersUsersTokensCreateCreated struct {
}

// IsSuccess returns true when this users users tokens create created response has a 2xx status code
func (o *UsersUsersTokensCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users users tokens create created response has a 3xx status code
func (o *UsersUsersTokensCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users tokens create created response has a 4xx status code
func (o *UsersUsersTokensCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this users users tokens create created response has a 5xx status code
func (o *UsersUsersTokensCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this users users tokens create created response a status code equal to that given
func (o *UsersUsersTokensCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *UsersUsersTokensCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{id}/tokens/][%d] usersUsersTokensCreateCreated ", 201)
}

func (o *UsersUsersTokensCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{id}/tokens/][%d] usersUsersTokensCreateCreated ", 201)
}

func (o *UsersUsersTokensCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
