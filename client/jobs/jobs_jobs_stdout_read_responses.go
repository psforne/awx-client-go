// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobsJobsStdoutReadReader is a Reader for the JobsJobsStdoutRead structure.
type JobsJobsStdoutReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobsJobsStdoutReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobsJobsStdoutReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJobsJobsStdoutReadOK creates a JobsJobsStdoutReadOK with default headers values
func NewJobsJobsStdoutReadOK() *JobsJobsStdoutReadOK {
	return &JobsJobsStdoutReadOK{}
}

/*
JobsJobsStdoutReadOK describes a response with status code 200, with default header values.

JobsJobsStdoutReadOK jobs jobs stdout read o k
*/
type JobsJobsStdoutReadOK struct {
}

// IsSuccess returns true when this jobs jobs stdout read o k response has a 2xx status code
func (o *JobsJobsStdoutReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this jobs jobs stdout read o k response has a 3xx status code
func (o *JobsJobsStdoutReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this jobs jobs stdout read o k response has a 4xx status code
func (o *JobsJobsStdoutReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this jobs jobs stdout read o k response has a 5xx status code
func (o *JobsJobsStdoutReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this jobs jobs stdout read o k response a status code equal to that given
func (o *JobsJobsStdoutReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *JobsJobsStdoutReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/jobs/{id}/stdout/][%d] jobsJobsStdoutReadOK ", 200)
}

func (o *JobsJobsStdoutReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/jobs/{id}/stdout/][%d] jobsJobsStdoutReadOK ", 200)
}

func (o *JobsJobsStdoutReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
