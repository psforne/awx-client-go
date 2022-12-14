// Code generated by go-swagger; DO NOT EDIT.

package workflow_approvals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowApprovalsWorkflowApprovalsApproveReadReader is a Reader for the WorkflowApprovalsWorkflowApprovalsApproveRead structure.
type WorkflowApprovalsWorkflowApprovalsApproveReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowApprovalsWorkflowApprovalsApproveReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowApprovalsWorkflowApprovalsApproveReadOK creates a WorkflowApprovalsWorkflowApprovalsApproveReadOK with default headers values
func NewWorkflowApprovalsWorkflowApprovalsApproveReadOK() *WorkflowApprovalsWorkflowApprovalsApproveReadOK {
	return &WorkflowApprovalsWorkflowApprovalsApproveReadOK{}
}

/*
WorkflowApprovalsWorkflowApprovalsApproveReadOK describes a response with status code 200, with default header values.

WorkflowApprovalsWorkflowApprovalsApproveReadOK workflow approvals workflow approvals approve read o k
*/
type WorkflowApprovalsWorkflowApprovalsApproveReadOK struct {
}

// IsSuccess returns true when this workflow approvals workflow approvals approve read o k response has a 2xx status code
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow approvals workflow approvals approve read o k response has a 3xx status code
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow approvals workflow approvals approve read o k response has a 4xx status code
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow approvals workflow approvals approve read o k response has a 5xx status code
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow approvals workflow approvals approve read o k response a status code equal to that given
func (o *WorkflowApprovalsWorkflowApprovalsApproveReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *WorkflowApprovalsWorkflowApprovalsApproveReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_approvals/{id}/approve/][%d] workflowApprovalsWorkflowApprovalsApproveReadOK ", 200)
}

func (o *WorkflowApprovalsWorkflowApprovalsApproveReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workflow_approvals/{id}/approve/][%d] workflowApprovalsWorkflowApprovalsApproveReadOK ", 200)
}

func (o *WorkflowApprovalsWorkflowApprovalsApproveReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
