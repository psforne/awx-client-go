// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RolesRolesTeamsCreateReader is a Reader for the RolesRolesTeamsCreate structure.
type RolesRolesTeamsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RolesRolesTeamsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRolesRolesTeamsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRolesRolesTeamsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRolesRolesTeamsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRolesRolesTeamsCreateCreated creates a RolesRolesTeamsCreateCreated with default headers values
func NewRolesRolesTeamsCreateCreated() *RolesRolesTeamsCreateCreated {
	return &RolesRolesTeamsCreateCreated{}
}

/*
RolesRolesTeamsCreateCreated describes a response with status code 201, with default header values.

RolesRolesTeamsCreateCreated roles roles teams create created
*/
type RolesRolesTeamsCreateCreated struct {
}

// IsSuccess returns true when this roles roles teams create created response has a 2xx status code
func (o *RolesRolesTeamsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this roles roles teams create created response has a 3xx status code
func (o *RolesRolesTeamsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this roles roles teams create created response has a 4xx status code
func (o *RolesRolesTeamsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this roles roles teams create created response has a 5xx status code
func (o *RolesRolesTeamsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this roles roles teams create created response a status code equal to that given
func (o *RolesRolesTeamsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *RolesRolesTeamsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/roles/{id}/teams/][%d] rolesRolesTeamsCreateCreated ", 201)
}

func (o *RolesRolesTeamsCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/roles/{id}/teams/][%d] rolesRolesTeamsCreateCreated ", 201)
}

func (o *RolesRolesTeamsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRolesRolesTeamsCreateBadRequest creates a RolesRolesTeamsCreateBadRequest with default headers values
func NewRolesRolesTeamsCreateBadRequest() *RolesRolesTeamsCreateBadRequest {
	return &RolesRolesTeamsCreateBadRequest{}
}

/*
RolesRolesTeamsCreateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type RolesRolesTeamsCreateBadRequest struct {
}

// IsSuccess returns true when this roles roles teams create bad request response has a 2xx status code
func (o *RolesRolesTeamsCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this roles roles teams create bad request response has a 3xx status code
func (o *RolesRolesTeamsCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this roles roles teams create bad request response has a 4xx status code
func (o *RolesRolesTeamsCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this roles roles teams create bad request response has a 5xx status code
func (o *RolesRolesTeamsCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this roles roles teams create bad request response a status code equal to that given
func (o *RolesRolesTeamsCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *RolesRolesTeamsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/roles/{id}/teams/][%d] rolesRolesTeamsCreateBadRequest ", 400)
}

func (o *RolesRolesTeamsCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/roles/{id}/teams/][%d] rolesRolesTeamsCreateBadRequest ", 400)
}

func (o *RolesRolesTeamsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRolesRolesTeamsCreateForbidden creates a RolesRolesTeamsCreateForbidden with default headers values
func NewRolesRolesTeamsCreateForbidden() *RolesRolesTeamsCreateForbidden {
	return &RolesRolesTeamsCreateForbidden{}
}

/*
RolesRolesTeamsCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type RolesRolesTeamsCreateForbidden struct {
}

// IsSuccess returns true when this roles roles teams create forbidden response has a 2xx status code
func (o *RolesRolesTeamsCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this roles roles teams create forbidden response has a 3xx status code
func (o *RolesRolesTeamsCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this roles roles teams create forbidden response has a 4xx status code
func (o *RolesRolesTeamsCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this roles roles teams create forbidden response has a 5xx status code
func (o *RolesRolesTeamsCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this roles roles teams create forbidden response a status code equal to that given
func (o *RolesRolesTeamsCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *RolesRolesTeamsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/roles/{id}/teams/][%d] rolesRolesTeamsCreateForbidden ", 403)
}

func (o *RolesRolesTeamsCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/roles/{id}/teams/][%d] rolesRolesTeamsCreateForbidden ", 403)
}

func (o *RolesRolesTeamsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
