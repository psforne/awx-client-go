// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationsOrganizationsInstanceGroupsCreateReader is a Reader for the OrganizationsOrganizationsInstanceGroupsCreate structure.
type OrganizationsOrganizationsInstanceGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationsInstanceGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewOrganizationsOrganizationsInstanceGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsOrganizationsInstanceGroupsCreateCreated creates a OrganizationsOrganizationsInstanceGroupsCreateCreated with default headers values
func NewOrganizationsOrganizationsInstanceGroupsCreateCreated() *OrganizationsOrganizationsInstanceGroupsCreateCreated {
	return &OrganizationsOrganizationsInstanceGroupsCreateCreated{}
}

/*
OrganizationsOrganizationsInstanceGroupsCreateCreated describes a response with status code 201, with default header values.

OrganizationsOrganizationsInstanceGroupsCreateCreated organizations organizations instance groups create created
*/
type OrganizationsOrganizationsInstanceGroupsCreateCreated struct {
}

// IsSuccess returns true when this organizations organizations instance groups create created response has a 2xx status code
func (o *OrganizationsOrganizationsInstanceGroupsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations organizations instance groups create created response has a 3xx status code
func (o *OrganizationsOrganizationsInstanceGroupsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations organizations instance groups create created response has a 4xx status code
func (o *OrganizationsOrganizationsInstanceGroupsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations organizations instance groups create created response has a 5xx status code
func (o *OrganizationsOrganizationsInstanceGroupsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations organizations instance groups create created response a status code equal to that given
func (o *OrganizationsOrganizationsInstanceGroupsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *OrganizationsOrganizationsInstanceGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/organizations/{id}/instance_groups/][%d] organizationsOrganizationsInstanceGroupsCreateCreated ", 201)
}

func (o *OrganizationsOrganizationsInstanceGroupsCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/organizations/{id}/instance_groups/][%d] organizationsOrganizationsInstanceGroupsCreateCreated ", 201)
}

func (o *OrganizationsOrganizationsInstanceGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
