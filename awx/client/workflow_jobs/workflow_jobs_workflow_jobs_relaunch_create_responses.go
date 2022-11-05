// Code generated by go-swagger; DO NOT EDIT.

package workflow_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobsWorkflowJobsRelaunchCreateReader is a Reader for the WorkflowJobsWorkflowJobsRelaunchCreate structure.
type WorkflowJobsWorkflowJobsRelaunchCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobsWorkflowJobsRelaunchCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewWorkflowJobsWorkflowJobsRelaunchCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobsWorkflowJobsRelaunchCreateCreated creates a WorkflowJobsWorkflowJobsRelaunchCreateCreated with default headers values
func NewWorkflowJobsWorkflowJobsRelaunchCreateCreated() *WorkflowJobsWorkflowJobsRelaunchCreateCreated {
	return &WorkflowJobsWorkflowJobsRelaunchCreateCreated{}
}

/*
WorkflowJobsWorkflowJobsRelaunchCreateCreated describes a response with status code 201, with default header values.

WorkflowJobsWorkflowJobsRelaunchCreateCreated workflow jobs workflow jobs relaunch create created
*/
type WorkflowJobsWorkflowJobsRelaunchCreateCreated struct {
}

// IsSuccess returns true when this workflow jobs workflow jobs relaunch create created response has a 2xx status code
func (o *WorkflowJobsWorkflowJobsRelaunchCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow jobs workflow jobs relaunch create created response has a 3xx status code
func (o *WorkflowJobsWorkflowJobsRelaunchCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow jobs workflow jobs relaunch create created response has a 4xx status code
func (o *WorkflowJobsWorkflowJobsRelaunchCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow jobs workflow jobs relaunch create created response has a 5xx status code
func (o *WorkflowJobsWorkflowJobsRelaunchCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow jobs workflow jobs relaunch create created response a status code equal to that given
func (o *WorkflowJobsWorkflowJobsRelaunchCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *WorkflowJobsWorkflowJobsRelaunchCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workflow_jobs/{id}/relaunch/][%d] workflowJobsWorkflowJobsRelaunchCreateCreated ", 201)
}

func (o *WorkflowJobsWorkflowJobsRelaunchCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/workflow_jobs/{id}/relaunch/][%d] workflowJobsWorkflowJobsRelaunchCreateCreated ", 201)
}

func (o *WorkflowJobsWorkflowJobsRelaunchCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
