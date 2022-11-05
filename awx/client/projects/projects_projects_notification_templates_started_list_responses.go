// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectsProjectsNotificationTemplatesStartedListReader is a Reader for the ProjectsProjectsNotificationTemplatesStartedList structure.
type ProjectsProjectsNotificationTemplatesStartedListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsProjectsNotificationTemplatesStartedListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsProjectsNotificationTemplatesStartedListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsProjectsNotificationTemplatesStartedListOK creates a ProjectsProjectsNotificationTemplatesStartedListOK with default headers values
func NewProjectsProjectsNotificationTemplatesStartedListOK() *ProjectsProjectsNotificationTemplatesStartedListOK {
	return &ProjectsProjectsNotificationTemplatesStartedListOK{}
}

/*
ProjectsProjectsNotificationTemplatesStartedListOK describes a response with status code 200, with default header values.

ProjectsProjectsNotificationTemplatesStartedListOK projects projects notification templates started list o k
*/
type ProjectsProjectsNotificationTemplatesStartedListOK struct {
}

// IsSuccess returns true when this projects projects notification templates started list o k response has a 2xx status code
func (o *ProjectsProjectsNotificationTemplatesStartedListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects projects notification templates started list o k response has a 3xx status code
func (o *ProjectsProjectsNotificationTemplatesStartedListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects projects notification templates started list o k response has a 4xx status code
func (o *ProjectsProjectsNotificationTemplatesStartedListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects projects notification templates started list o k response has a 5xx status code
func (o *ProjectsProjectsNotificationTemplatesStartedListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects projects notification templates started list o k response a status code equal to that given
func (o *ProjectsProjectsNotificationTemplatesStartedListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsProjectsNotificationTemplatesStartedListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{id}/notification_templates_started/][%d] projectsProjectsNotificationTemplatesStartedListOK ", 200)
}

func (o *ProjectsProjectsNotificationTemplatesStartedListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{id}/notification_templates_started/][%d] projectsProjectsNotificationTemplatesStartedListOK ", 200)
}

func (o *ProjectsProjectsNotificationTemplatesStartedListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
