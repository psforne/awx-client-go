// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreate structure.
type WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated creates a WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated() *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated {
	return &WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated{}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated describes a response with status code 201, with default header values.

WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated workflow job templates workflow job templates notification templates approvals create created
*/
type WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated struct {
}

// IsSuccess returns true when this workflow job templates workflow job templates notification templates approvals create created response has a 2xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow job templates workflow job templates notification templates approvals create created response has a 3xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow job templates workflow job templates notification templates approvals create created response has a 4xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow job templates workflow job templates notification templates approvals create created response has a 5xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow job templates workflow job templates notification templates approvals create created response a status code equal to that given
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/{id}/notification_templates_approvals/][%d] workflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated ", 201)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/{id}/notification_templates_approvals/][%d] workflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated ", 201)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesApprovalsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
