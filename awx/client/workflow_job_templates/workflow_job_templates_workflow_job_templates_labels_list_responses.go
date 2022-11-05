// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplatesWorkflowJobTemplatesLabelsListReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesLabelsList structure.
type WorkflowJobTemplatesWorkflowJobTemplatesLabelsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK creates a WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK() *WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK {
	return &WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK{}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK describes a response with status code 200, with default header values.

WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK workflow job templates workflow job templates labels list o k
*/
type WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK struct {
}

// IsSuccess returns true when this workflow job templates workflow job templates labels list o k response has a 2xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow job templates workflow job templates labels list o k response has a 3xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow job templates workflow job templates labels list o k response has a 4xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow job templates workflow job templates labels list o k response has a 5xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow job templates workflow job templates labels list o k response a status code equal to that given
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_templates/{id}/labels/][%d] workflowJobTemplatesWorkflowJobTemplatesLabelsListOK ", 200)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_templates/{id}/labels/][%d] workflowJobTemplatesWorkflowJobTemplatesLabelsListOK ", 200)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLabelsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
