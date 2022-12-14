// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationsOrganizationsJobTemplatesListReader is a Reader for the OrganizationsOrganizationsJobTemplatesList structure.
type OrganizationsOrganizationsJobTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationsJobTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsOrganizationsJobTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsOrganizationsJobTemplatesListOK creates a OrganizationsOrganizationsJobTemplatesListOK with default headers values
func NewOrganizationsOrganizationsJobTemplatesListOK() *OrganizationsOrganizationsJobTemplatesListOK {
	return &OrganizationsOrganizationsJobTemplatesListOK{}
}

/*
OrganizationsOrganizationsJobTemplatesListOK describes a response with status code 200, with default header values.

OrganizationsOrganizationsJobTemplatesListOK organizations organizations job templates list o k
*/
type OrganizationsOrganizationsJobTemplatesListOK struct {
}

// IsSuccess returns true when this organizations organizations job templates list o k response has a 2xx status code
func (o *OrganizationsOrganizationsJobTemplatesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations organizations job templates list o k response has a 3xx status code
func (o *OrganizationsOrganizationsJobTemplatesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations organizations job templates list o k response has a 4xx status code
func (o *OrganizationsOrganizationsJobTemplatesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations organizations job templates list o k response has a 5xx status code
func (o *OrganizationsOrganizationsJobTemplatesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations organizations job templates list o k response a status code equal to that given
func (o *OrganizationsOrganizationsJobTemplatesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *OrganizationsOrganizationsJobTemplatesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/job_templates/][%d] organizationsOrganizationsJobTemplatesListOK ", 200)
}

func (o *OrganizationsOrganizationsJobTemplatesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/job_templates/][%d] organizationsOrganizationsJobTemplatesListOK ", 200)
}

func (o *OrganizationsOrganizationsJobTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
