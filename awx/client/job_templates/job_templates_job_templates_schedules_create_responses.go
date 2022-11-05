// Code generated by go-swagger; DO NOT EDIT.

package job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobTemplatesJobTemplatesSchedulesCreateReader is a Reader for the JobTemplatesJobTemplatesSchedulesCreate structure.
type JobTemplatesJobTemplatesSchedulesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobTemplatesJobTemplatesSchedulesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewJobTemplatesJobTemplatesSchedulesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewJobTemplatesJobTemplatesSchedulesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJobTemplatesJobTemplatesSchedulesCreateCreated creates a JobTemplatesJobTemplatesSchedulesCreateCreated with default headers values
func NewJobTemplatesJobTemplatesSchedulesCreateCreated() *JobTemplatesJobTemplatesSchedulesCreateCreated {
	return &JobTemplatesJobTemplatesSchedulesCreateCreated{}
}

/*
JobTemplatesJobTemplatesSchedulesCreateCreated describes a response with status code 201, with default header values.

JobTemplatesJobTemplatesSchedulesCreateCreated job templates job templates schedules create created
*/
type JobTemplatesJobTemplatesSchedulesCreateCreated struct {
}

// IsSuccess returns true when this job templates job templates schedules create created response has a 2xx status code
func (o *JobTemplatesJobTemplatesSchedulesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this job templates job templates schedules create created response has a 3xx status code
func (o *JobTemplatesJobTemplatesSchedulesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this job templates job templates schedules create created response has a 4xx status code
func (o *JobTemplatesJobTemplatesSchedulesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this job templates job templates schedules create created response has a 5xx status code
func (o *JobTemplatesJobTemplatesSchedulesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this job templates job templates schedules create created response a status code equal to that given
func (o *JobTemplatesJobTemplatesSchedulesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *JobTemplatesJobTemplatesSchedulesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/job_templates/{id}/schedules/][%d] jobTemplatesJobTemplatesSchedulesCreateCreated ", 201)
}

func (o *JobTemplatesJobTemplatesSchedulesCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/job_templates/{id}/schedules/][%d] jobTemplatesJobTemplatesSchedulesCreateCreated ", 201)
}

func (o *JobTemplatesJobTemplatesSchedulesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewJobTemplatesJobTemplatesSchedulesCreateBadRequest creates a JobTemplatesJobTemplatesSchedulesCreateBadRequest with default headers values
func NewJobTemplatesJobTemplatesSchedulesCreateBadRequest() *JobTemplatesJobTemplatesSchedulesCreateBadRequest {
	return &JobTemplatesJobTemplatesSchedulesCreateBadRequest{}
}

/*
JobTemplatesJobTemplatesSchedulesCreateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type JobTemplatesJobTemplatesSchedulesCreateBadRequest struct {
}

// IsSuccess returns true when this job templates job templates schedules create bad request response has a 2xx status code
func (o *JobTemplatesJobTemplatesSchedulesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this job templates job templates schedules create bad request response has a 3xx status code
func (o *JobTemplatesJobTemplatesSchedulesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this job templates job templates schedules create bad request response has a 4xx status code
func (o *JobTemplatesJobTemplatesSchedulesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this job templates job templates schedules create bad request response has a 5xx status code
func (o *JobTemplatesJobTemplatesSchedulesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this job templates job templates schedules create bad request response a status code equal to that given
func (o *JobTemplatesJobTemplatesSchedulesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *JobTemplatesJobTemplatesSchedulesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/job_templates/{id}/schedules/][%d] jobTemplatesJobTemplatesSchedulesCreateBadRequest ", 400)
}

func (o *JobTemplatesJobTemplatesSchedulesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/job_templates/{id}/schedules/][%d] jobTemplatesJobTemplatesSchedulesCreateBadRequest ", 400)
}

func (o *JobTemplatesJobTemplatesSchedulesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
