// Code generated by go-swagger; DO NOT EDIT.

package workflow_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobsWorkflowJobsRelaunchListReader is a Reader for the WorkflowJobsWorkflowJobsRelaunchList structure.
type WorkflowJobsWorkflowJobsRelaunchListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobsWorkflowJobsRelaunchListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobsWorkflowJobsRelaunchListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobsWorkflowJobsRelaunchListOK creates a WorkflowJobsWorkflowJobsRelaunchListOK with default headers values
func NewWorkflowJobsWorkflowJobsRelaunchListOK() *WorkflowJobsWorkflowJobsRelaunchListOK {
	return &WorkflowJobsWorkflowJobsRelaunchListOK{}
}

/*
WorkflowJobsWorkflowJobsRelaunchListOK describes a response with status code 200, with default header values.

WorkflowJobsWorkflowJobsRelaunchListOK workflow jobs workflow jobs relaunch list o k
*/
type WorkflowJobsWorkflowJobsRelaunchListOK struct {
}

// IsSuccess returns true when this workflow jobs workflow jobs relaunch list o k response has a 2xx status code
func (o *WorkflowJobsWorkflowJobsRelaunchListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow jobs workflow jobs relaunch list o k response has a 3xx status code
func (o *WorkflowJobsWorkflowJobsRelaunchListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow jobs workflow jobs relaunch list o k response has a 4xx status code
func (o *WorkflowJobsWorkflowJobsRelaunchListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow jobs workflow jobs relaunch list o k response has a 5xx status code
func (o *WorkflowJobsWorkflowJobsRelaunchListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow jobs workflow jobs relaunch list o k response a status code equal to that given
func (o *WorkflowJobsWorkflowJobsRelaunchListOK) IsCode(code int) bool {
	return code == 200
}

func (o *WorkflowJobsWorkflowJobsRelaunchListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_jobs/{id}/relaunch/][%d] workflowJobsWorkflowJobsRelaunchListOK ", 200)
}

func (o *WorkflowJobsWorkflowJobsRelaunchListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workflow_jobs/{id}/relaunch/][%d] workflowJobsWorkflowJobsRelaunchListOK ", 200)
}

func (o *WorkflowJobsWorkflowJobsRelaunchListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
