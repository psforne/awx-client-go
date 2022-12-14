// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GroupsGroupsHostsListReader is a Reader for the GroupsGroupsHostsList structure.
type GroupsGroupsHostsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsGroupsHostsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupsGroupsHostsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGroupsGroupsHostsListOK creates a GroupsGroupsHostsListOK with default headers values
func NewGroupsGroupsHostsListOK() *GroupsGroupsHostsListOK {
	return &GroupsGroupsHostsListOK{}
}

/*
GroupsGroupsHostsListOK describes a response with status code 200, with default header values.

GroupsGroupsHostsListOK groups groups hosts list o k
*/
type GroupsGroupsHostsListOK struct {
}

// IsSuccess returns true when this groups groups hosts list o k response has a 2xx status code
func (o *GroupsGroupsHostsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this groups groups hosts list o k response has a 3xx status code
func (o *GroupsGroupsHostsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this groups groups hosts list o k response has a 4xx status code
func (o *GroupsGroupsHostsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this groups groups hosts list o k response has a 5xx status code
func (o *GroupsGroupsHostsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this groups groups hosts list o k response a status code equal to that given
func (o *GroupsGroupsHostsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *GroupsGroupsHostsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{id}/hosts/][%d] groupsGroupsHostsListOK ", 200)
}

func (o *GroupsGroupsHostsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/groups/{id}/hosts/][%d] groupsGroupsHostsListOK ", 200)
}

func (o *GroupsGroupsHostsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
