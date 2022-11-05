// Code generated by go-swagger; DO NOT EDIT.

package job_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobEventsJobEventsChildrenListReader is a Reader for the JobEventsJobEventsChildrenList structure.
type JobEventsJobEventsChildrenListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobEventsJobEventsChildrenListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobEventsJobEventsChildrenListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJobEventsJobEventsChildrenListOK creates a JobEventsJobEventsChildrenListOK with default headers values
func NewJobEventsJobEventsChildrenListOK() *JobEventsJobEventsChildrenListOK {
	return &JobEventsJobEventsChildrenListOK{}
}

/*
JobEventsJobEventsChildrenListOK describes a response with status code 200, with default header values.

JobEventsJobEventsChildrenListOK job events job events children list o k
*/
type JobEventsJobEventsChildrenListOK struct {
}

// IsSuccess returns true when this job events job events children list o k response has a 2xx status code
func (o *JobEventsJobEventsChildrenListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this job events job events children list o k response has a 3xx status code
func (o *JobEventsJobEventsChildrenListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this job events job events children list o k response has a 4xx status code
func (o *JobEventsJobEventsChildrenListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this job events job events children list o k response has a 5xx status code
func (o *JobEventsJobEventsChildrenListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this job events job events children list o k response a status code equal to that given
func (o *JobEventsJobEventsChildrenListOK) IsCode(code int) bool {
	return code == 200
}

func (o *JobEventsJobEventsChildrenListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/job_events/{id}/children/][%d] jobEventsJobEventsChildrenListOK ", 200)
}

func (o *JobEventsJobEventsChildrenListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/job_events/{id}/children/][%d] jobEventsJobEventsChildrenListOK ", 200)
}

func (o *JobEventsJobEventsChildrenListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}