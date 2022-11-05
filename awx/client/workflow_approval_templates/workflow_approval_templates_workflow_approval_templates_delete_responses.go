// Code generated by go-swagger; DO NOT EDIT.

package workflow_approval_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteReader is a Reader for the WorkflowApprovalTemplatesWorkflowApprovalTemplatesDelete structure.
type WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent creates a WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent with default headers values
func NewWorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent() *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent {
	return &WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent{}
}

/*
WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent describes a response with status code 204, with default header values.

WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent workflow approval templates workflow approval templates delete no content
*/
type WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent struct {
}

// IsSuccess returns true when this workflow approval templates workflow approval templates delete no content response has a 2xx status code
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow approval templates workflow approval templates delete no content response has a 3xx status code
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow approval templates workflow approval templates delete no content response has a 4xx status code
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow approval templates workflow approval templates delete no content response has a 5xx status code
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow approval templates workflow approval templates delete no content response a status code equal to that given
func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workflow_approval_templates/{id}/][%d] workflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent ", 204)
}

func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/workflow_approval_templates/{id}/][%d] workflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent ", 204)
}

func (o *WorkflowApprovalTemplatesWorkflowApprovalTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
