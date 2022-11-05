// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationsOrganizationsAdminsCreateReader is a Reader for the OrganizationsOrganizationsAdminsCreate structure.
type OrganizationsOrganizationsAdminsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationsAdminsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewOrganizationsOrganizationsAdminsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewOrganizationsOrganizationsAdminsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsOrganizationsAdminsCreateCreated creates a OrganizationsOrganizationsAdminsCreateCreated with default headers values
func NewOrganizationsOrganizationsAdminsCreateCreated() *OrganizationsOrganizationsAdminsCreateCreated {
	return &OrganizationsOrganizationsAdminsCreateCreated{}
}

/*
OrganizationsOrganizationsAdminsCreateCreated describes a response with status code 201, with default header values.

OrganizationsOrganizationsAdminsCreateCreated organizations organizations admins create created
*/
type OrganizationsOrganizationsAdminsCreateCreated struct {
}

// IsSuccess returns true when this organizations organizations admins create created response has a 2xx status code
func (o *OrganizationsOrganizationsAdminsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations organizations admins create created response has a 3xx status code
func (o *OrganizationsOrganizationsAdminsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations organizations admins create created response has a 4xx status code
func (o *OrganizationsOrganizationsAdminsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations organizations admins create created response has a 5xx status code
func (o *OrganizationsOrganizationsAdminsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations organizations admins create created response a status code equal to that given
func (o *OrganizationsOrganizationsAdminsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *OrganizationsOrganizationsAdminsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/organizations/{id}/admins/][%d] organizationsOrganizationsAdminsCreateCreated ", 201)
}

func (o *OrganizationsOrganizationsAdminsCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/organizations/{id}/admins/][%d] organizationsOrganizationsAdminsCreateCreated ", 201)
}

func (o *OrganizationsOrganizationsAdminsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrganizationsOrganizationsAdminsCreateForbidden creates a OrganizationsOrganizationsAdminsCreateForbidden with default headers values
func NewOrganizationsOrganizationsAdminsCreateForbidden() *OrganizationsOrganizationsAdminsCreateForbidden {
	return &OrganizationsOrganizationsAdminsCreateForbidden{}
}

/*
OrganizationsOrganizationsAdminsCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type OrganizationsOrganizationsAdminsCreateForbidden struct {
}

// IsSuccess returns true when this organizations organizations admins create forbidden response has a 2xx status code
func (o *OrganizationsOrganizationsAdminsCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations organizations admins create forbidden response has a 3xx status code
func (o *OrganizationsOrganizationsAdminsCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations organizations admins create forbidden response has a 4xx status code
func (o *OrganizationsOrganizationsAdminsCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations organizations admins create forbidden response has a 5xx status code
func (o *OrganizationsOrganizationsAdminsCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations organizations admins create forbidden response a status code equal to that given
func (o *OrganizationsOrganizationsAdminsCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OrganizationsOrganizationsAdminsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/organizations/{id}/admins/][%d] organizationsOrganizationsAdminsCreateForbidden ", 403)
}

func (o *OrganizationsOrganizationsAdminsCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/organizations/{id}/admins/][%d] organizationsOrganizationsAdminsCreateForbidden ", 403)
}

func (o *OrganizationsOrganizationsAdminsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}