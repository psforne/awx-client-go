// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesList structure.
type WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK creates a WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK() *WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK {
	return &WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK{}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK describes a response with status code 200, with default header values.

WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK workflow job templates workflow job templates workflow nodes list o k
*/
type WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK struct {
}

// IsSuccess returns true when this workflow job templates workflow job templates workflow nodes list o k response has a 2xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow job templates workflow job templates workflow nodes list o k response has a 3xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow job templates workflow job templates workflow nodes list o k response has a 4xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow job templates workflow job templates workflow nodes list o k response has a 5xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow job templates workflow job templates workflow nodes list o k response a status code equal to that given
func (o *WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_templates/{id}/workflow_nodes/][%d] workflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK ", 200)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_templates/{id}/workflow_nodes/][%d] workflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK ", 200)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesWorkflowNodesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
