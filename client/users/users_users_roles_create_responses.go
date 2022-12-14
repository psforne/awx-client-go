// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersUsersRolesCreateReader is a Reader for the UsersUsersRolesCreate structure.
type UsersUsersRolesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersRolesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersUsersRolesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersUsersRolesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersUsersRolesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersUsersRolesCreateCreated creates a UsersUsersRolesCreateCreated with default headers values
func NewUsersUsersRolesCreateCreated() *UsersUsersRolesCreateCreated {
	return &UsersUsersRolesCreateCreated{}
}

/*
UsersUsersRolesCreateCreated describes a response with status code 201, with default header values.

UsersUsersRolesCreateCreated users users roles create created
*/
type UsersUsersRolesCreateCreated struct {
}

// IsSuccess returns true when this users users roles create created response has a 2xx status code
func (o *UsersUsersRolesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users users roles create created response has a 3xx status code
func (o *UsersUsersRolesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users roles create created response has a 4xx status code
func (o *UsersUsersRolesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this users users roles create created response has a 5xx status code
func (o *UsersUsersRolesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this users users roles create created response a status code equal to that given
func (o *UsersUsersRolesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *UsersUsersRolesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{id}/roles/][%d] usersUsersRolesCreateCreated ", 201)
}

func (o *UsersUsersRolesCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{id}/roles/][%d] usersUsersRolesCreateCreated ", 201)
}

func (o *UsersUsersRolesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersUsersRolesCreateBadRequest creates a UsersUsersRolesCreateBadRequest with default headers values
func NewUsersUsersRolesCreateBadRequest() *UsersUsersRolesCreateBadRequest {
	return &UsersUsersRolesCreateBadRequest{}
}

/*
UsersUsersRolesCreateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UsersUsersRolesCreateBadRequest struct {
}

// IsSuccess returns true when this users users roles create bad request response has a 2xx status code
func (o *UsersUsersRolesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users users roles create bad request response has a 3xx status code
func (o *UsersUsersRolesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users roles create bad request response has a 4xx status code
func (o *UsersUsersRolesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users users roles create bad request response has a 5xx status code
func (o *UsersUsersRolesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users users roles create bad request response a status code equal to that given
func (o *UsersUsersRolesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersUsersRolesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{id}/roles/][%d] usersUsersRolesCreateBadRequest ", 400)
}

func (o *UsersUsersRolesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{id}/roles/][%d] usersUsersRolesCreateBadRequest ", 400)
}

func (o *UsersUsersRolesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersUsersRolesCreateForbidden creates a UsersUsersRolesCreateForbidden with default headers values
func NewUsersUsersRolesCreateForbidden() *UsersUsersRolesCreateForbidden {
	return &UsersUsersRolesCreateForbidden{}
}

/*
UsersUsersRolesCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type UsersUsersRolesCreateForbidden struct {
}

// IsSuccess returns true when this users users roles create forbidden response has a 2xx status code
func (o *UsersUsersRolesCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users users roles create forbidden response has a 3xx status code
func (o *UsersUsersRolesCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users roles create forbidden response has a 4xx status code
func (o *UsersUsersRolesCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this users users roles create forbidden response has a 5xx status code
func (o *UsersUsersRolesCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this users users roles create forbidden response a status code equal to that given
func (o *UsersUsersRolesCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UsersUsersRolesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{id}/roles/][%d] usersUsersRolesCreateForbidden ", 403)
}

func (o *UsersUsersRolesCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{id}/roles/][%d] usersUsersRolesCreateForbidden ", 403)
}

func (o *UsersUsersRolesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
