// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_template_nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplateNodesWorkflowJobTemplateNodesListReader is a Reader for the WorkflowJobTemplateNodesWorkflowJobTemplateNodesList structure.
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK creates a WorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK with default headers values
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK() *WorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK{}
}

/*
WorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK describes a response with status code 200, with default header values.

WorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK workflow job template nodes workflow job template nodes list o k
*/
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK struct {
}

// IsSuccess returns true when this workflow job template nodes workflow job template nodes list o k response has a 2xx status code
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow job template nodes workflow job template nodes list o k response has a 3xx status code
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow job template nodes workflow job template nodes list o k response has a 4xx status code
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow job template nodes workflow job template nodes list o k response has a 5xx status code
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow job template nodes workflow job template nodes list o k response a status code equal to that given
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_template_nodes/][%d] workflowJobTemplateNodesWorkflowJobTemplateNodesListOK ", 200)
}

func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_template_nodes/][%d] workflowJobTemplateNodesWorkflowJobTemplateNodesListOK ", 200)
}

func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
