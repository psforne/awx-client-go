// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreate structure.
type WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated creates a WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated() *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated {
	return &WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated{}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated describes a response with status code 201, with default header values.

WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated workflow job templates workflow job templates webhook key create created
*/
type WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated struct {
}

// IsSuccess returns true when this workflow job templates workflow job templates webhook key create created response has a 2xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow job templates workflow job templates webhook key create created response has a 3xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow job templates workflow job templates webhook key create created response has a 4xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow job templates workflow job templates webhook key create created response has a 5xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow job templates workflow job templates webhook key create created response a status code equal to that given
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/{id}/webhook_key/][%d] workflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated ", 201)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/{id}/webhook_key/][%d] workflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated ", 201)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden creates a WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden() *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden {
	return &WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden{}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden struct {
}

// IsSuccess returns true when this workflow job templates workflow job templates webhook key create forbidden response has a 2xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this workflow job templates workflow job templates webhook key create forbidden response has a 3xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow job templates workflow job templates webhook key create forbidden response has a 4xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this workflow job templates workflow job templates webhook key create forbidden response has a 5xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow job templates workflow job templates webhook key create forbidden response a status code equal to that given
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/{id}/webhook_key/][%d] workflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden ", 403)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/{id}/webhook_key/][%d] workflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden ", 403)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesWebhookKeyCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}