// Code generated by go-swagger; DO NOT EDIT.

package system_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SystemJobTemplatesSystemJobTemplatesListReader is a Reader for the SystemJobTemplatesSystemJobTemplatesList structure.
type SystemJobTemplatesSystemJobTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemJobTemplatesSystemJobTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemJobTemplatesSystemJobTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemJobTemplatesSystemJobTemplatesListOK creates a SystemJobTemplatesSystemJobTemplatesListOK with default headers values
func NewSystemJobTemplatesSystemJobTemplatesListOK() *SystemJobTemplatesSystemJobTemplatesListOK {
	return &SystemJobTemplatesSystemJobTemplatesListOK{}
}

/*
SystemJobTemplatesSystemJobTemplatesListOK describes a response with status code 200, with default header values.

SystemJobTemplatesSystemJobTemplatesListOK system job templates system job templates list o k
*/
type SystemJobTemplatesSystemJobTemplatesListOK struct {
}

// IsSuccess returns true when this system job templates system job templates list o k response has a 2xx status code
func (o *SystemJobTemplatesSystemJobTemplatesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this system job templates system job templates list o k response has a 3xx status code
func (o *SystemJobTemplatesSystemJobTemplatesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system job templates system job templates list o k response has a 4xx status code
func (o *SystemJobTemplatesSystemJobTemplatesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this system job templates system job templates list o k response has a 5xx status code
func (o *SystemJobTemplatesSystemJobTemplatesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this system job templates system job templates list o k response a status code equal to that given
func (o *SystemJobTemplatesSystemJobTemplatesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SystemJobTemplatesSystemJobTemplatesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/system_job_templates/][%d] systemJobTemplatesSystemJobTemplatesListOK ", 200)
}

func (o *SystemJobTemplatesSystemJobTemplatesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/system_job_templates/][%d] systemJobTemplatesSystemJobTemplatesListOK ", 200)
}

func (o *SystemJobTemplatesSystemJobTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
