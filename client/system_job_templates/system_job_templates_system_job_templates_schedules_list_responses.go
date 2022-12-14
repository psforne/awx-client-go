// Code generated by go-swagger; DO NOT EDIT.

package system_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SystemJobTemplatesSystemJobTemplatesSchedulesListReader is a Reader for the SystemJobTemplatesSystemJobTemplatesSchedulesList structure.
type SystemJobTemplatesSystemJobTemplatesSchedulesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemJobTemplatesSystemJobTemplatesSchedulesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemJobTemplatesSystemJobTemplatesSchedulesListOK creates a SystemJobTemplatesSystemJobTemplatesSchedulesListOK with default headers values
func NewSystemJobTemplatesSystemJobTemplatesSchedulesListOK() *SystemJobTemplatesSystemJobTemplatesSchedulesListOK {
	return &SystemJobTemplatesSystemJobTemplatesSchedulesListOK{}
}

/*
SystemJobTemplatesSystemJobTemplatesSchedulesListOK describes a response with status code 200, with default header values.

SystemJobTemplatesSystemJobTemplatesSchedulesListOK system job templates system job templates schedules list o k
*/
type SystemJobTemplatesSystemJobTemplatesSchedulesListOK struct {
}

// IsSuccess returns true when this system job templates system job templates schedules list o k response has a 2xx status code
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this system job templates system job templates schedules list o k response has a 3xx status code
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system job templates system job templates schedules list o k response has a 4xx status code
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this system job templates system job templates schedules list o k response has a 5xx status code
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this system job templates system job templates schedules list o k response a status code equal to that given
func (o *SystemJobTemplatesSystemJobTemplatesSchedulesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SystemJobTemplatesSystemJobTemplatesSchedulesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/system_job_templates/{id}/schedules/][%d] systemJobTemplatesSystemJobTemplatesSchedulesListOK ", 200)
}

func (o *SystemJobTemplatesSystemJobTemplatesSchedulesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/system_job_templates/{id}/schedules/][%d] systemJobTemplatesSystemJobTemplatesSchedulesListOK ", 200)
}

func (o *SystemJobTemplatesSystemJobTemplatesSchedulesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
