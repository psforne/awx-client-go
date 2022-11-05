// Code generated by go-swagger; DO NOT EDIT.

package workflow_approvals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowApprovalsWorkflowApprovalsDeleteReader is a Reader for the WorkflowApprovalsWorkflowApprovalsDelete structure.
type WorkflowApprovalsWorkflowApprovalsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowApprovalsWorkflowApprovalsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewWorkflowApprovalsWorkflowApprovalsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowApprovalsWorkflowApprovalsDeleteNoContent creates a WorkflowApprovalsWorkflowApprovalsDeleteNoContent with default headers values
func NewWorkflowApprovalsWorkflowApprovalsDeleteNoContent() *WorkflowApprovalsWorkflowApprovalsDeleteNoContent {
	return &WorkflowApprovalsWorkflowApprovalsDeleteNoContent{}
}

/*
WorkflowApprovalsWorkflowApprovalsDeleteNoContent describes a response with status code 204, with default header values.

WorkflowApprovalsWorkflowApprovalsDeleteNoContent workflow approvals workflow approvals delete no content
*/
type WorkflowApprovalsWorkflowApprovalsDeleteNoContent struct {
}

// IsSuccess returns true when this workflow approvals workflow approvals delete no content response has a 2xx status code
func (o *WorkflowApprovalsWorkflowApprovalsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow approvals workflow approvals delete no content response has a 3xx status code
func (o *WorkflowApprovalsWorkflowApprovalsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow approvals workflow approvals delete no content response has a 4xx status code
func (o *WorkflowApprovalsWorkflowApprovalsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow approvals workflow approvals delete no content response has a 5xx status code
func (o *WorkflowApprovalsWorkflowApprovalsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow approvals workflow approvals delete no content response a status code equal to that given
func (o *WorkflowApprovalsWorkflowApprovalsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *WorkflowApprovalsWorkflowApprovalsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workflow_approvals/{id}/][%d] workflowApprovalsWorkflowApprovalsDeleteNoContent ", 204)
}

func (o *WorkflowApprovalsWorkflowApprovalsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/workflow_approvals/{id}/][%d] workflowApprovalsWorkflowApprovalsDeleteNoContent ", 204)
}

func (o *WorkflowApprovalsWorkflowApprovalsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
