// Code generated by go-swagger; DO NOT EDIT.

package workflow_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobsWorkflowJobsReadReader is a Reader for the WorkflowJobsWorkflowJobsRead structure.
type WorkflowJobsWorkflowJobsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobsWorkflowJobsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobsWorkflowJobsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobsWorkflowJobsReadOK creates a WorkflowJobsWorkflowJobsReadOK with default headers values
func NewWorkflowJobsWorkflowJobsReadOK() *WorkflowJobsWorkflowJobsReadOK {
	return &WorkflowJobsWorkflowJobsReadOK{}
}

/*
WorkflowJobsWorkflowJobsReadOK describes a response with status code 200, with default header values.

WorkflowJobsWorkflowJobsReadOK workflow jobs workflow jobs read o k
*/
type WorkflowJobsWorkflowJobsReadOK struct {
}

// IsSuccess returns true when this workflow jobs workflow jobs read o k response has a 2xx status code
func (o *WorkflowJobsWorkflowJobsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow jobs workflow jobs read o k response has a 3xx status code
func (o *WorkflowJobsWorkflowJobsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow jobs workflow jobs read o k response has a 4xx status code
func (o *WorkflowJobsWorkflowJobsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow jobs workflow jobs read o k response has a 5xx status code
func (o *WorkflowJobsWorkflowJobsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow jobs workflow jobs read o k response a status code equal to that given
func (o *WorkflowJobsWorkflowJobsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *WorkflowJobsWorkflowJobsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_jobs/{id}/][%d] workflowJobsWorkflowJobsReadOK ", 200)
}

func (o *WorkflowJobsWorkflowJobsReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workflow_jobs/{id}/][%d] workflowJobsWorkflowJobsReadOK ", 200)
}

func (o *WorkflowJobsWorkflowJobsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
