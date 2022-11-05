// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationsOrganizationsUsersListReader is a Reader for the OrganizationsOrganizationsUsersList structure.
type OrganizationsOrganizationsUsersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationsUsersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsOrganizationsUsersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsOrganizationsUsersListOK creates a OrganizationsOrganizationsUsersListOK with default headers values
func NewOrganizationsOrganizationsUsersListOK() *OrganizationsOrganizationsUsersListOK {
	return &OrganizationsOrganizationsUsersListOK{}
}

/*
OrganizationsOrganizationsUsersListOK describes a response with status code 200, with default header values.

OrganizationsOrganizationsUsersListOK organizations organizations users list o k
*/
type OrganizationsOrganizationsUsersListOK struct {
}

// IsSuccess returns true when this organizations organizations users list o k response has a 2xx status code
func (o *OrganizationsOrganizationsUsersListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations organizations users list o k response has a 3xx status code
func (o *OrganizationsOrganizationsUsersListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations organizations users list o k response has a 4xx status code
func (o *OrganizationsOrganizationsUsersListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations organizations users list o k response has a 5xx status code
func (o *OrganizationsOrganizationsUsersListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations organizations users list o k response a status code equal to that given
func (o *OrganizationsOrganizationsUsersListOK) IsCode(code int) bool {
	return code == 200
}

func (o *OrganizationsOrganizationsUsersListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/users/][%d] organizationsOrganizationsUsersListOK ", 200)
}

func (o *OrganizationsOrganizationsUsersListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/users/][%d] organizationsOrganizationsUsersListOK ", 200)
}

func (o *OrganizationsOrganizationsUsersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}