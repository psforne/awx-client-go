// Code generated by go-swagger; DO NOT EDIT.

package system_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListReader is a Reader for the SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorList structure.
type SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK creates a SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK with default headers values
func NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK() *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK {
	return &SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK{}
}

/*
SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK describes a response with status code 200, with default header values.

SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK system job templates system job templates notification templates error list o k
*/
type SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK struct {
}

// IsSuccess returns true when this system job templates system job templates notification templates error list o k response has a 2xx status code
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this system job templates system job templates notification templates error list o k response has a 3xx status code
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system job templates system job templates notification templates error list o k response has a 4xx status code
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this system job templates system job templates notification templates error list o k response has a 5xx status code
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this system job templates system job templates notification templates error list o k response a status code equal to that given
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/system_job_templates/{id}/notification_templates_error/][%d] systemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK ", 200)
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/system_job_templates/{id}/notification_templates_error/][%d] systemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK ", 200)
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
