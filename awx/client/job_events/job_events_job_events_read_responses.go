// Code generated by go-swagger; DO NOT EDIT.

package job_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobEventsJobEventsReadReader is a Reader for the JobEventsJobEventsRead structure.
type JobEventsJobEventsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobEventsJobEventsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobEventsJobEventsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJobEventsJobEventsReadOK creates a JobEventsJobEventsReadOK with default headers values
func NewJobEventsJobEventsReadOK() *JobEventsJobEventsReadOK {
	return &JobEventsJobEventsReadOK{}
}

/*
JobEventsJobEventsReadOK describes a response with status code 200, with default header values.

JobEventsJobEventsReadOK job events job events read o k
*/
type JobEventsJobEventsReadOK struct {
}

// IsSuccess returns true when this job events job events read o k response has a 2xx status code
func (o *JobEventsJobEventsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this job events job events read o k response has a 3xx status code
func (o *JobEventsJobEventsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this job events job events read o k response has a 4xx status code
func (o *JobEventsJobEventsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this job events job events read o k response has a 5xx status code
func (o *JobEventsJobEventsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this job events job events read o k response a status code equal to that given
func (o *JobEventsJobEventsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *JobEventsJobEventsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/job_events/{id}/][%d] jobEventsJobEventsReadOK ", 200)
}

func (o *JobEventsJobEventsReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/job_events/{id}/][%d] jobEventsJobEventsReadOK ", 200)
}

func (o *JobEventsJobEventsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}