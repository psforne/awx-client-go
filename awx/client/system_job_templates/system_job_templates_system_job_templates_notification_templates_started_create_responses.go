// Code generated by go-swagger; DO NOT EDIT.

package system_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateReader is a Reader for the SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreate structure.
type SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated creates a SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated with default headers values
func NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated() *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated {
	return &SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated{}
}

/*
SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated describes a response with status code 201, with default header values.

SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated system job templates system job templates notification templates started create created
*/
type SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated struct {
}

// IsSuccess returns true when this system job templates system job templates notification templates started create created response has a 2xx status code
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this system job templates system job templates notification templates started create created response has a 3xx status code
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system job templates system job templates notification templates started create created response has a 4xx status code
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this system job templates system job templates notification templates started create created response has a 5xx status code
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this system job templates system job templates notification templates started create created response a status code equal to that given
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/system_job_templates/{id}/notification_templates_started/][%d] systemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated ", 201)
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/system_job_templates/{id}/notification_templates_started/][%d] systemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated ", 201)
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesStartedCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
