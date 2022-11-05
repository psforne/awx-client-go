// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GroupsGroupsHostsCreateReader is a Reader for the GroupsGroupsHostsCreate structure.
type GroupsGroupsHostsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsGroupsHostsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGroupsGroupsHostsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGroupsGroupsHostsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGroupsGroupsHostsCreateCreated creates a GroupsGroupsHostsCreateCreated with default headers values
func NewGroupsGroupsHostsCreateCreated() *GroupsGroupsHostsCreateCreated {
	return &GroupsGroupsHostsCreateCreated{}
}

/*
GroupsGroupsHostsCreateCreated describes a response with status code 201, with default header values.

GroupsGroupsHostsCreateCreated groups groups hosts create created
*/
type GroupsGroupsHostsCreateCreated struct {
}

// IsSuccess returns true when this groups groups hosts create created response has a 2xx status code
func (o *GroupsGroupsHostsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this groups groups hosts create created response has a 3xx status code
func (o *GroupsGroupsHostsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this groups groups hosts create created response has a 4xx status code
func (o *GroupsGroupsHostsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this groups groups hosts create created response has a 5xx status code
func (o *GroupsGroupsHostsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this groups groups hosts create created response a status code equal to that given
func (o *GroupsGroupsHostsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *GroupsGroupsHostsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{id}/hosts/][%d] groupsGroupsHostsCreateCreated ", 201)
}

func (o *GroupsGroupsHostsCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/{id}/hosts/][%d] groupsGroupsHostsCreateCreated ", 201)
}

func (o *GroupsGroupsHostsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGroupsGroupsHostsCreateForbidden creates a GroupsGroupsHostsCreateForbidden with default headers values
func NewGroupsGroupsHostsCreateForbidden() *GroupsGroupsHostsCreateForbidden {
	return &GroupsGroupsHostsCreateForbidden{}
}

/*
GroupsGroupsHostsCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type GroupsGroupsHostsCreateForbidden struct {
}

// IsSuccess returns true when this groups groups hosts create forbidden response has a 2xx status code
func (o *GroupsGroupsHostsCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this groups groups hosts create forbidden response has a 3xx status code
func (o *GroupsGroupsHostsCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this groups groups hosts create forbidden response has a 4xx status code
func (o *GroupsGroupsHostsCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this groups groups hosts create forbidden response has a 5xx status code
func (o *GroupsGroupsHostsCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this groups groups hosts create forbidden response a status code equal to that given
func (o *GroupsGroupsHostsCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GroupsGroupsHostsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/{id}/hosts/][%d] groupsGroupsHostsCreateForbidden ", 403)
}

func (o *GroupsGroupsHostsCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/{id}/hosts/][%d] groupsGroupsHostsCreateForbidden ", 403)
}

func (o *GroupsGroupsHostsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
