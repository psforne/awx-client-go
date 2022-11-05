// Code generated by go-swagger; DO NOT EDIT.

package workflow_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobsWorkflowJobsWorkflowNodesListReader is a Reader for the WorkflowJobsWorkflowJobsWorkflowNodesList structure.
type WorkflowJobsWorkflowJobsWorkflowNodesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobsWorkflowJobsWorkflowNodesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobsWorkflowJobsWorkflowNodesListOK creates a WorkflowJobsWorkflowJobsWorkflowNodesListOK with default headers values
func NewWorkflowJobsWorkflowJobsWorkflowNodesListOK() *WorkflowJobsWorkflowJobsWorkflowNodesListOK {
	return &WorkflowJobsWorkflowJobsWorkflowNodesListOK{}
}

/*
WorkflowJobsWorkflowJobsWorkflowNodesListOK describes a response with status code 200, with default header values.

WorkflowJobsWorkflowJobsWorkflowNodesListOK workflow jobs workflow jobs workflow nodes list o k
*/
type WorkflowJobsWorkflowJobsWorkflowNodesListOK struct {
}

// IsSuccess returns true when this workflow jobs workflow jobs workflow nodes list o k response has a 2xx status code
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow jobs workflow jobs workflow nodes list o k response has a 3xx status code
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow jobs workflow jobs workflow nodes list o k response has a 4xx status code
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow jobs workflow jobs workflow nodes list o k response has a 5xx status code
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow jobs workflow jobs workflow nodes list o k response a status code equal to that given
func (o *WorkflowJobsWorkflowJobsWorkflowNodesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *WorkflowJobsWorkflowJobsWorkflowNodesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_jobs/{id}/workflow_nodes/][%d] workflowJobsWorkflowJobsWorkflowNodesListOK ", 200)
}

func (o *WorkflowJobsWorkflowJobsWorkflowNodesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workflow_jobs/{id}/workflow_nodes/][%d] workflowJobsWorkflowJobsWorkflowNodesListOK ", 200)
}

func (o *WorkflowJobsWorkflowJobsWorkflowNodesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
