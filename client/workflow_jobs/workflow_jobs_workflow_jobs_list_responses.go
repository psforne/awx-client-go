// Code generated by go-swagger; DO NOT EDIT.

package workflow_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobsWorkflowJobsListReader is a Reader for the WorkflowJobsWorkflowJobsList structure.
type WorkflowJobsWorkflowJobsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobsWorkflowJobsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobsWorkflowJobsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobsWorkflowJobsListOK creates a WorkflowJobsWorkflowJobsListOK with default headers values
func NewWorkflowJobsWorkflowJobsListOK() *WorkflowJobsWorkflowJobsListOK {
	return &WorkflowJobsWorkflowJobsListOK{}
}

/*
WorkflowJobsWorkflowJobsListOK describes a response with status code 200, with default header values.

WorkflowJobsWorkflowJobsListOK workflow jobs workflow jobs list o k
*/
type WorkflowJobsWorkflowJobsListOK struct {
}

// IsSuccess returns true when this workflow jobs workflow jobs list o k response has a 2xx status code
func (o *WorkflowJobsWorkflowJobsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow jobs workflow jobs list o k response has a 3xx status code
func (o *WorkflowJobsWorkflowJobsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow jobs workflow jobs list o k response has a 4xx status code
func (o *WorkflowJobsWorkflowJobsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow jobs workflow jobs list o k response has a 5xx status code
func (o *WorkflowJobsWorkflowJobsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow jobs workflow jobs list o k response a status code equal to that given
func (o *WorkflowJobsWorkflowJobsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *WorkflowJobsWorkflowJobsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_jobs/][%d] workflowJobsWorkflowJobsListOK ", 200)
}

func (o *WorkflowJobsWorkflowJobsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workflow_jobs/][%d] workflowJobsWorkflowJobsListOK ", 200)
}

func (o *WorkflowJobsWorkflowJobsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}