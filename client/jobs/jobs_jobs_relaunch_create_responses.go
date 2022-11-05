// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobsJobsRelaunchCreateReader is a Reader for the JobsJobsRelaunchCreate structure.
type JobsJobsRelaunchCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobsJobsRelaunchCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewJobsJobsRelaunchCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewJobsJobsRelaunchCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJobsJobsRelaunchCreateCreated creates a JobsJobsRelaunchCreateCreated with default headers values
func NewJobsJobsRelaunchCreateCreated() *JobsJobsRelaunchCreateCreated {
	return &JobsJobsRelaunchCreateCreated{}
}

/*
JobsJobsRelaunchCreateCreated describes a response with status code 201, with default header values.

JobsJobsRelaunchCreateCreated jobs jobs relaunch create created
*/
type JobsJobsRelaunchCreateCreated struct {
}

// IsSuccess returns true when this jobs jobs relaunch create created response has a 2xx status code
func (o *JobsJobsRelaunchCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this jobs jobs relaunch create created response has a 3xx status code
func (o *JobsJobsRelaunchCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this jobs jobs relaunch create created response has a 4xx status code
func (o *JobsJobsRelaunchCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this jobs jobs relaunch create created response has a 5xx status code
func (o *JobsJobsRelaunchCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this jobs jobs relaunch create created response a status code equal to that given
func (o *JobsJobsRelaunchCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *JobsJobsRelaunchCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/jobs/{id}/relaunch/][%d] jobsJobsRelaunchCreateCreated ", 201)
}

func (o *JobsJobsRelaunchCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/jobs/{id}/relaunch/][%d] jobsJobsRelaunchCreateCreated ", 201)
}

func (o *JobsJobsRelaunchCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewJobsJobsRelaunchCreateForbidden creates a JobsJobsRelaunchCreateForbidden with default headers values
func NewJobsJobsRelaunchCreateForbidden() *JobsJobsRelaunchCreateForbidden {
	return &JobsJobsRelaunchCreateForbidden{}
}

/*
JobsJobsRelaunchCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type JobsJobsRelaunchCreateForbidden struct {
}

// IsSuccess returns true when this jobs jobs relaunch create forbidden response has a 2xx status code
func (o *JobsJobsRelaunchCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this jobs jobs relaunch create forbidden response has a 3xx status code
func (o *JobsJobsRelaunchCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this jobs jobs relaunch create forbidden response has a 4xx status code
func (o *JobsJobsRelaunchCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this jobs jobs relaunch create forbidden response has a 5xx status code
func (o *JobsJobsRelaunchCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this jobs jobs relaunch create forbidden response a status code equal to that given
func (o *JobsJobsRelaunchCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *JobsJobsRelaunchCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/jobs/{id}/relaunch/][%d] jobsJobsRelaunchCreateForbidden ", 403)
}

func (o *JobsJobsRelaunchCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/jobs/{id}/relaunch/][%d] jobsJobsRelaunchCreateForbidden ", 403)
}

func (o *JobsJobsRelaunchCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}