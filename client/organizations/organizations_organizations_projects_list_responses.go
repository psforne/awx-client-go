// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationsOrganizationsProjectsListReader is a Reader for the OrganizationsOrganizationsProjectsList structure.
type OrganizationsOrganizationsProjectsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationsProjectsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsOrganizationsProjectsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewOrganizationsOrganizationsProjectsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsOrganizationsProjectsListOK creates a OrganizationsOrganizationsProjectsListOK with default headers values
func NewOrganizationsOrganizationsProjectsListOK() *OrganizationsOrganizationsProjectsListOK {
	return &OrganizationsOrganizationsProjectsListOK{}
}

/*
OrganizationsOrganizationsProjectsListOK describes a response with status code 200, with default header values.

OrganizationsOrganizationsProjectsListOK organizations organizations projects list o k
*/
type OrganizationsOrganizationsProjectsListOK struct {
}

// IsSuccess returns true when this organizations organizations projects list o k response has a 2xx status code
func (o *OrganizationsOrganizationsProjectsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations organizations projects list o k response has a 3xx status code
func (o *OrganizationsOrganizationsProjectsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations organizations projects list o k response has a 4xx status code
func (o *OrganizationsOrganizationsProjectsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations organizations projects list o k response has a 5xx status code
func (o *OrganizationsOrganizationsProjectsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations organizations projects list o k response a status code equal to that given
func (o *OrganizationsOrganizationsProjectsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *OrganizationsOrganizationsProjectsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/projects/][%d] organizationsOrganizationsProjectsListOK ", 200)
}

func (o *OrganizationsOrganizationsProjectsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/projects/][%d] organizationsOrganizationsProjectsListOK ", 200)
}

func (o *OrganizationsOrganizationsProjectsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrganizationsOrganizationsProjectsListForbidden creates a OrganizationsOrganizationsProjectsListForbidden with default headers values
func NewOrganizationsOrganizationsProjectsListForbidden() *OrganizationsOrganizationsProjectsListForbidden {
	return &OrganizationsOrganizationsProjectsListForbidden{}
}

/*
OrganizationsOrganizationsProjectsListForbidden describes a response with status code 403, with default header values.

forbidden
*/
type OrganizationsOrganizationsProjectsListForbidden struct {
}

// IsSuccess returns true when this organizations organizations projects list forbidden response has a 2xx status code
func (o *OrganizationsOrganizationsProjectsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations organizations projects list forbidden response has a 3xx status code
func (o *OrganizationsOrganizationsProjectsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations organizations projects list forbidden response has a 4xx status code
func (o *OrganizationsOrganizationsProjectsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations organizations projects list forbidden response has a 5xx status code
func (o *OrganizationsOrganizationsProjectsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations organizations projects list forbidden response a status code equal to that given
func (o *OrganizationsOrganizationsProjectsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OrganizationsOrganizationsProjectsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/projects/][%d] organizationsOrganizationsProjectsListForbidden ", 403)
}

func (o *OrganizationsOrganizationsProjectsListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/projects/][%d] organizationsOrganizationsProjectsListForbidden ", 403)
}

func (o *OrganizationsOrganizationsProjectsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
