// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GroupsGroupsPartialUpdateReader is a Reader for the GroupsGroupsPartialUpdate structure.
type GroupsGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupsGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGroupsGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGroupsGroupsPartialUpdateOK creates a GroupsGroupsPartialUpdateOK with default headers values
func NewGroupsGroupsPartialUpdateOK() *GroupsGroupsPartialUpdateOK {
	return &GroupsGroupsPartialUpdateOK{}
}

/*
GroupsGroupsPartialUpdateOK describes a response with status code 200, with default header values.

GroupsGroupsPartialUpdateOK groups groups partial update o k
*/
type GroupsGroupsPartialUpdateOK struct {
}

// IsSuccess returns true when this groups groups partial update o k response has a 2xx status code
func (o *GroupsGroupsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this groups groups partial update o k response has a 3xx status code
func (o *GroupsGroupsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this groups groups partial update o k response has a 4xx status code
func (o *GroupsGroupsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this groups groups partial update o k response has a 5xx status code
func (o *GroupsGroupsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this groups groups partial update o k response a status code equal to that given
func (o *GroupsGroupsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *GroupsGroupsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/groups/{id}/][%d] groupsGroupsPartialUpdateOK ", 200)
}

func (o *GroupsGroupsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/groups/{id}/][%d] groupsGroupsPartialUpdateOK ", 200)
}

func (o *GroupsGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
GroupsGroupsPartialUpdateBody groups groups partial update body
swagger:model GroupsGroupsPartialUpdateBody
*/
type GroupsGroupsPartialUpdateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// inventory
	Inventory int64 `json:"inventory,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Group variables in JSON or YAML format.
	Variables string `json:"variables,omitempty"`
}

// Validate validates this groups groups partial update body
func (o *GroupsGroupsPartialUpdateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this groups groups partial update body based on context it is used
func (o *GroupsGroupsPartialUpdateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GroupsGroupsPartialUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GroupsGroupsPartialUpdateBody) UnmarshalBinary(b []byte) error {
	var res GroupsGroupsPartialUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
