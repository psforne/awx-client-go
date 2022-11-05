// Code generated by go-swagger; DO NOT EDIT.

package system_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SystemJobsSystemJobsCancelCreateReader is a Reader for the SystemJobsSystemJobsCancelCreate structure.
type SystemJobsSystemJobsCancelCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemJobsSystemJobsCancelCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSystemJobsSystemJobsCancelCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemJobsSystemJobsCancelCreateCreated creates a SystemJobsSystemJobsCancelCreateCreated with default headers values
func NewSystemJobsSystemJobsCancelCreateCreated() *SystemJobsSystemJobsCancelCreateCreated {
	return &SystemJobsSystemJobsCancelCreateCreated{}
}

/*
SystemJobsSystemJobsCancelCreateCreated describes a response with status code 201, with default header values.

SystemJobsSystemJobsCancelCreateCreated system jobs system jobs cancel create created
*/
type SystemJobsSystemJobsCancelCreateCreated struct {
}

// IsSuccess returns true when this system jobs system jobs cancel create created response has a 2xx status code
func (o *SystemJobsSystemJobsCancelCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this system jobs system jobs cancel create created response has a 3xx status code
func (o *SystemJobsSystemJobsCancelCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system jobs system jobs cancel create created response has a 4xx status code
func (o *SystemJobsSystemJobsCancelCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this system jobs system jobs cancel create created response has a 5xx status code
func (o *SystemJobsSystemJobsCancelCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this system jobs system jobs cancel create created response a status code equal to that given
func (o *SystemJobsSystemJobsCancelCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *SystemJobsSystemJobsCancelCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/system_jobs/{id}/cancel/][%d] systemJobsSystemJobsCancelCreateCreated ", 201)
}

func (o *SystemJobsSystemJobsCancelCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/system_jobs/{id}/cancel/][%d] systemJobsSystemJobsCancelCreateCreated ", 201)
}

func (o *SystemJobsSystemJobsCancelCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}