// Code generated by go-swagger; DO NOT EDIT.

package job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobTemplatesJobTemplatesReadReader is a Reader for the JobTemplatesJobTemplatesRead structure.
type JobTemplatesJobTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobTemplatesJobTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobTemplatesJobTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJobTemplatesJobTemplatesReadOK creates a JobTemplatesJobTemplatesReadOK with default headers values
func NewJobTemplatesJobTemplatesReadOK() *JobTemplatesJobTemplatesReadOK {
	return &JobTemplatesJobTemplatesReadOK{}
}

/*
JobTemplatesJobTemplatesReadOK describes a response with status code 200, with default header values.

JobTemplatesJobTemplatesReadOK job templates job templates read o k
*/
type JobTemplatesJobTemplatesReadOK struct {
}

// IsSuccess returns true when this job templates job templates read o k response has a 2xx status code
func (o *JobTemplatesJobTemplatesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this job templates job templates read o k response has a 3xx status code
func (o *JobTemplatesJobTemplatesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this job templates job templates read o k response has a 4xx status code
func (o *JobTemplatesJobTemplatesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this job templates job templates read o k response has a 5xx status code
func (o *JobTemplatesJobTemplatesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this job templates job templates read o k response a status code equal to that given
func (o *JobTemplatesJobTemplatesReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *JobTemplatesJobTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/job_templates/{id}/][%d] jobTemplatesJobTemplatesReadOK ", 200)
}

func (o *JobTemplatesJobTemplatesReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/job_templates/{id}/][%d] jobTemplatesJobTemplatesReadOK ", 200)
}

func (o *JobTemplatesJobTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
