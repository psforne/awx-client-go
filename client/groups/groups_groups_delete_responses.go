// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GroupsGroupsDeleteReader is a Reader for the GroupsGroupsDelete structure.
type GroupsGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewGroupsGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGroupsGroupsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGroupsGroupsDeleteNoContent creates a GroupsGroupsDeleteNoContent with default headers values
func NewGroupsGroupsDeleteNoContent() *GroupsGroupsDeleteNoContent {
	return &GroupsGroupsDeleteNoContent{}
}

/*
GroupsGroupsDeleteNoContent describes a response with status code 204, with default header values.

GroupsGroupsDeleteNoContent groups groups delete no content
*/
type GroupsGroupsDeleteNoContent struct {
}

// IsSuccess returns true when this groups groups delete no content response has a 2xx status code
func (o *GroupsGroupsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this groups groups delete no content response has a 3xx status code
func (o *GroupsGroupsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this groups groups delete no content response has a 4xx status code
func (o *GroupsGroupsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this groups groups delete no content response has a 5xx status code
func (o *GroupsGroupsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this groups groups delete no content response a status code equal to that given
func (o *GroupsGroupsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *GroupsGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{id}/][%d] groupsGroupsDeleteNoContent ", 204)
}

func (o *GroupsGroupsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{id}/][%d] groupsGroupsDeleteNoContent ", 204)
}

func (o *GroupsGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGroupsGroupsDeleteForbidden creates a GroupsGroupsDeleteForbidden with default headers values
func NewGroupsGroupsDeleteForbidden() *GroupsGroupsDeleteForbidden {
	return &GroupsGroupsDeleteForbidden{}
}

/*
GroupsGroupsDeleteForbidden describes a response with status code 403, with default header values.

forbidden
*/
type GroupsGroupsDeleteForbidden struct {
}

// IsSuccess returns true when this groups groups delete forbidden response has a 2xx status code
func (o *GroupsGroupsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this groups groups delete forbidden response has a 3xx status code
func (o *GroupsGroupsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this groups groups delete forbidden response has a 4xx status code
func (o *GroupsGroupsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this groups groups delete forbidden response has a 5xx status code
func (o *GroupsGroupsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this groups groups delete forbidden response a status code equal to that given
func (o *GroupsGroupsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GroupsGroupsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{id}/][%d] groupsGroupsDeleteForbidden ", 403)
}

func (o *GroupsGroupsDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/groups/{id}/][%d] groupsGroupsDeleteForbidden ", 403)
}

func (o *GroupsGroupsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
