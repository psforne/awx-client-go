// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GroupsGroupsVariableDataReadReader is a Reader for the GroupsGroupsVariableDataRead structure.
type GroupsGroupsVariableDataReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsGroupsVariableDataReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupsGroupsVariableDataReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGroupsGroupsVariableDataReadOK creates a GroupsGroupsVariableDataReadOK with default headers values
func NewGroupsGroupsVariableDataReadOK() *GroupsGroupsVariableDataReadOK {
	return &GroupsGroupsVariableDataReadOK{}
}

/*
GroupsGroupsVariableDataReadOK describes a response with status code 200, with default header values.

GroupsGroupsVariableDataReadOK groups groups variable data read o k
*/
type GroupsGroupsVariableDataReadOK struct {
}

// IsSuccess returns true when this groups groups variable data read o k response has a 2xx status code
func (o *GroupsGroupsVariableDataReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this groups groups variable data read o k response has a 3xx status code
func (o *GroupsGroupsVariableDataReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this groups groups variable data read o k response has a 4xx status code
func (o *GroupsGroupsVariableDataReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this groups groups variable data read o k response has a 5xx status code
func (o *GroupsGroupsVariableDataReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this groups groups variable data read o k response a status code equal to that given
func (o *GroupsGroupsVariableDataReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *GroupsGroupsVariableDataReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{id}/variable_data/][%d] groupsGroupsVariableDataReadOK ", 200)
}

func (o *GroupsGroupsVariableDataReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/groups/{id}/variable_data/][%d] groupsGroupsVariableDataReadOK ", 200)
}

func (o *GroupsGroupsVariableDataReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
