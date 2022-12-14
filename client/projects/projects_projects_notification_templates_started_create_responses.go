// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectsProjectsNotificationTemplatesStartedCreateReader is a Reader for the ProjectsProjectsNotificationTemplatesStartedCreate structure.
type ProjectsProjectsNotificationTemplatesStartedCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsProjectsNotificationTemplatesStartedCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewProjectsProjectsNotificationTemplatesStartedCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsProjectsNotificationTemplatesStartedCreateCreated creates a ProjectsProjectsNotificationTemplatesStartedCreateCreated with default headers values
func NewProjectsProjectsNotificationTemplatesStartedCreateCreated() *ProjectsProjectsNotificationTemplatesStartedCreateCreated {
	return &ProjectsProjectsNotificationTemplatesStartedCreateCreated{}
}

/*
ProjectsProjectsNotificationTemplatesStartedCreateCreated describes a response with status code 201, with default header values.

ProjectsProjectsNotificationTemplatesStartedCreateCreated projects projects notification templates started create created
*/
type ProjectsProjectsNotificationTemplatesStartedCreateCreated struct {
}

// IsSuccess returns true when this projects projects notification templates started create created response has a 2xx status code
func (o *ProjectsProjectsNotificationTemplatesStartedCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects projects notification templates started create created response has a 3xx status code
func (o *ProjectsProjectsNotificationTemplatesStartedCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects projects notification templates started create created response has a 4xx status code
func (o *ProjectsProjectsNotificationTemplatesStartedCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects projects notification templates started create created response has a 5xx status code
func (o *ProjectsProjectsNotificationTemplatesStartedCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this projects projects notification templates started create created response a status code equal to that given
func (o *ProjectsProjectsNotificationTemplatesStartedCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *ProjectsProjectsNotificationTemplatesStartedCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{id}/notification_templates_started/][%d] projectsProjectsNotificationTemplatesStartedCreateCreated ", 201)
}

func (o *ProjectsProjectsNotificationTemplatesStartedCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/{id}/notification_templates_started/][%d] projectsProjectsNotificationTemplatesStartedCreateCreated ", 201)
}

func (o *ProjectsProjectsNotificationTemplatesStartedCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
