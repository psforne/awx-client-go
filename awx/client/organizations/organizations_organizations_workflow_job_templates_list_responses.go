// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationsOrganizationsWorkflowJobTemplatesListReader is a Reader for the OrganizationsOrganizationsWorkflowJobTemplatesList structure.
type OrganizationsOrganizationsWorkflowJobTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationsWorkflowJobTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsOrganizationsWorkflowJobTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsOrganizationsWorkflowJobTemplatesListOK creates a OrganizationsOrganizationsWorkflowJobTemplatesListOK with default headers values
func NewOrganizationsOrganizationsWorkflowJobTemplatesListOK() *OrganizationsOrganizationsWorkflowJobTemplatesListOK {
	return &OrganizationsOrganizationsWorkflowJobTemplatesListOK{}
}

/*
OrganizationsOrganizationsWorkflowJobTemplatesListOK describes a response with status code 200, with default header values.

OrganizationsOrganizationsWorkflowJobTemplatesListOK organizations organizations workflow job templates list o k
*/
type OrganizationsOrganizationsWorkflowJobTemplatesListOK struct {
}

// IsSuccess returns true when this organizations organizations workflow job templates list o k response has a 2xx status code
func (o *OrganizationsOrganizationsWorkflowJobTemplatesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations organizations workflow job templates list o k response has a 3xx status code
func (o *OrganizationsOrganizationsWorkflowJobTemplatesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations organizations workflow job templates list o k response has a 4xx status code
func (o *OrganizationsOrganizationsWorkflowJobTemplatesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations organizations workflow job templates list o k response has a 5xx status code
func (o *OrganizationsOrganizationsWorkflowJobTemplatesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations organizations workflow job templates list o k response a status code equal to that given
func (o *OrganizationsOrganizationsWorkflowJobTemplatesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *OrganizationsOrganizationsWorkflowJobTemplatesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/workflow_job_templates/][%d] organizationsOrganizationsWorkflowJobTemplatesListOK ", 200)
}

func (o *OrganizationsOrganizationsWorkflowJobTemplatesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/workflow_job_templates/][%d] organizationsOrganizationsWorkflowJobTemplatesListOK ", 200)
}

func (o *OrganizationsOrganizationsWorkflowJobTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}