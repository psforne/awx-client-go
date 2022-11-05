// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectsProjectsNotificationTemplatesErrorListReader is a Reader for the ProjectsProjectsNotificationTemplatesErrorList structure.
type ProjectsProjectsNotificationTemplatesErrorListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsProjectsNotificationTemplatesErrorListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsProjectsNotificationTemplatesErrorListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsProjectsNotificationTemplatesErrorListOK creates a ProjectsProjectsNotificationTemplatesErrorListOK with default headers values
func NewProjectsProjectsNotificationTemplatesErrorListOK() *ProjectsProjectsNotificationTemplatesErrorListOK {
	return &ProjectsProjectsNotificationTemplatesErrorListOK{}
}

/*
ProjectsProjectsNotificationTemplatesErrorListOK describes a response with status code 200, with default header values.

ProjectsProjectsNotificationTemplatesErrorListOK projects projects notification templates error list o k
*/
type ProjectsProjectsNotificationTemplatesErrorListOK struct {
}

// IsSuccess returns true when this projects projects notification templates error list o k response has a 2xx status code
func (o *ProjectsProjectsNotificationTemplatesErrorListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects projects notification templates error list o k response has a 3xx status code
func (o *ProjectsProjectsNotificationTemplatesErrorListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects projects notification templates error list o k response has a 4xx status code
func (o *ProjectsProjectsNotificationTemplatesErrorListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects projects notification templates error list o k response has a 5xx status code
func (o *ProjectsProjectsNotificationTemplatesErrorListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects projects notification templates error list o k response a status code equal to that given
func (o *ProjectsProjectsNotificationTemplatesErrorListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsProjectsNotificationTemplatesErrorListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{id}/notification_templates_error/][%d] projectsProjectsNotificationTemplatesErrorListOK ", 200)
}

func (o *ProjectsProjectsNotificationTemplatesErrorListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{id}/notification_templates_error/][%d] projectsProjectsNotificationTemplatesErrorListOK ", 200)
}

func (o *ProjectsProjectsNotificationTemplatesErrorListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
