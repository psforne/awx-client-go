// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesGithubCreate structure.
type WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated creates a WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated() *WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated {
	return &WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated{}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated describes a response with status code 201, with default header values.

WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated workflow job templates workflow job templates github create created
*/
type WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated struct {
}

// IsSuccess returns true when this workflow job templates workflow job templates github create created response has a 2xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow job templates workflow job templates github create created response has a 3xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow job templates workflow job templates github create created response has a 4xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow job templates workflow job templates github create created response has a 5xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow job templates workflow job templates github create created response a status code equal to that given
func (o *WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/{id}/github/][%d] workflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated ", 201)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/{id}/github/][%d] workflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated ", 201)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesGithubCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
