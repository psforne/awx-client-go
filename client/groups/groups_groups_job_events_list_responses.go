// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GroupsGroupsJobEventsListReader is a Reader for the GroupsGroupsJobEventsList structure.
type GroupsGroupsJobEventsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsGroupsJobEventsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupsGroupsJobEventsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGroupsGroupsJobEventsListOK creates a GroupsGroupsJobEventsListOK with default headers values
func NewGroupsGroupsJobEventsListOK() *GroupsGroupsJobEventsListOK {
	return &GroupsGroupsJobEventsListOK{}
}

/*
GroupsGroupsJobEventsListOK describes a response with status code 200, with default header values.

GroupsGroupsJobEventsListOK groups groups job events list o k
*/
type GroupsGroupsJobEventsListOK struct {
}

// IsSuccess returns true when this groups groups job events list o k response has a 2xx status code
func (o *GroupsGroupsJobEventsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this groups groups job events list o k response has a 3xx status code
func (o *GroupsGroupsJobEventsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this groups groups job events list o k response has a 4xx status code
func (o *GroupsGroupsJobEventsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this groups groups job events list o k response has a 5xx status code
func (o *GroupsGroupsJobEventsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this groups groups job events list o k response a status code equal to that given
func (o *GroupsGroupsJobEventsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *GroupsGroupsJobEventsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{id}/job_events/][%d] groupsGroupsJobEventsListOK ", 200)
}

func (o *GroupsGroupsJobEventsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/groups/{id}/job_events/][%d] groupsGroupsJobEventsListOK ", 200)
}

func (o *GroupsGroupsJobEventsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
