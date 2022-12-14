// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobsJobsJobEventsListReader is a Reader for the JobsJobsJobEventsList structure.
type JobsJobsJobEventsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobsJobsJobEventsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobsJobsJobEventsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJobsJobsJobEventsListOK creates a JobsJobsJobEventsListOK with default headers values
func NewJobsJobsJobEventsListOK() *JobsJobsJobEventsListOK {
	return &JobsJobsJobEventsListOK{}
}

/*
JobsJobsJobEventsListOK describes a response with status code 200, with default header values.

JobsJobsJobEventsListOK jobs jobs job events list o k
*/
type JobsJobsJobEventsListOK struct {
}

// IsSuccess returns true when this jobs jobs job events list o k response has a 2xx status code
func (o *JobsJobsJobEventsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this jobs jobs job events list o k response has a 3xx status code
func (o *JobsJobsJobEventsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this jobs jobs job events list o k response has a 4xx status code
func (o *JobsJobsJobEventsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this jobs jobs job events list o k response has a 5xx status code
func (o *JobsJobsJobEventsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this jobs jobs job events list o k response a status code equal to that given
func (o *JobsJobsJobEventsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *JobsJobsJobEventsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/jobs/{id}/job_events/][%d] jobsJobsJobEventsListOK ", 200)
}

func (o *JobsJobsJobEventsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/jobs/{id}/job_events/][%d] jobsJobsJobEventsListOK ", 200)
}

func (o *JobsJobsJobEventsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
