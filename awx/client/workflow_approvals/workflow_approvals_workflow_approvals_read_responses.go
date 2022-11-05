// Code generated by go-swagger; DO NOT EDIT.

package workflow_approvals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowApprovalsWorkflowApprovalsReadReader is a Reader for the WorkflowApprovalsWorkflowApprovalsRead structure.
type WorkflowApprovalsWorkflowApprovalsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowApprovalsWorkflowApprovalsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowApprovalsWorkflowApprovalsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowApprovalsWorkflowApprovalsReadOK creates a WorkflowApprovalsWorkflowApprovalsReadOK with default headers values
func NewWorkflowApprovalsWorkflowApprovalsReadOK() *WorkflowApprovalsWorkflowApprovalsReadOK {
	return &WorkflowApprovalsWorkflowApprovalsReadOK{}
}

/*
WorkflowApprovalsWorkflowApprovalsReadOK describes a response with status code 200, with default header values.

WorkflowApprovalsWorkflowApprovalsReadOK workflow approvals workflow approvals read o k
*/
type WorkflowApprovalsWorkflowApprovalsReadOK struct {
}

// IsSuccess returns true when this workflow approvals workflow approvals read o k response has a 2xx status code
func (o *WorkflowApprovalsWorkflowApprovalsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow approvals workflow approvals read o k response has a 3xx status code
func (o *WorkflowApprovalsWorkflowApprovalsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow approvals workflow approvals read o k response has a 4xx status code
func (o *WorkflowApprovalsWorkflowApprovalsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow approvals workflow approvals read o k response has a 5xx status code
func (o *WorkflowApprovalsWorkflowApprovalsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow approvals workflow approvals read o k response a status code equal to that given
func (o *WorkflowApprovalsWorkflowApprovalsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *WorkflowApprovalsWorkflowApprovalsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_approvals/{id}/][%d] workflowApprovalsWorkflowApprovalsReadOK ", 200)
}

func (o *WorkflowApprovalsWorkflowApprovalsReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workflow_approvals/{id}/][%d] workflowApprovalsWorkflowApprovalsReadOK ", 200)
}

func (o *WorkflowApprovalsWorkflowApprovalsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
