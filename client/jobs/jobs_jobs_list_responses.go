// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobsJobsListReader is a Reader for the JobsJobsList structure.
type JobsJobsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobsJobsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobsJobsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewJobsJobsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJobsJobsListOK creates a JobsJobsListOK with default headers values
func NewJobsJobsListOK() *JobsJobsListOK {
	return &JobsJobsListOK{}
}

/*
JobsJobsListOK describes a response with status code 200, with default header values.

JobsJobsListOK jobs jobs list o k
*/
type JobsJobsListOK struct {
}

// IsSuccess returns true when this jobs jobs list o k response has a 2xx status code
func (o *JobsJobsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this jobs jobs list o k response has a 3xx status code
func (o *JobsJobsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this jobs jobs list o k response has a 4xx status code
func (o *JobsJobsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this jobs jobs list o k response has a 5xx status code
func (o *JobsJobsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this jobs jobs list o k response a status code equal to that given
func (o *JobsJobsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *JobsJobsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/jobs/][%d] jobsJobsListOK ", 200)
}

func (o *JobsJobsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/jobs/][%d] jobsJobsListOK ", 200)
}

func (o *JobsJobsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewJobsJobsListForbidden creates a JobsJobsListForbidden with default headers values
func NewJobsJobsListForbidden() *JobsJobsListForbidden {
	return &JobsJobsListForbidden{}
}

/*
JobsJobsListForbidden describes a response with status code 403, with default header values.

forbidden
*/
type JobsJobsListForbidden struct {
}

// IsSuccess returns true when this jobs jobs list forbidden response has a 2xx status code
func (o *JobsJobsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this jobs jobs list forbidden response has a 3xx status code
func (o *JobsJobsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this jobs jobs list forbidden response has a 4xx status code
func (o *JobsJobsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this jobs jobs list forbidden response has a 5xx status code
func (o *JobsJobsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this jobs jobs list forbidden response a status code equal to that given
func (o *JobsJobsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *JobsJobsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/jobs/][%d] jobsJobsListForbidden ", 403)
}

func (o *JobsJobsListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/jobs/][%d] jobsJobsListForbidden ", 403)
}

func (o *JobsJobsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
