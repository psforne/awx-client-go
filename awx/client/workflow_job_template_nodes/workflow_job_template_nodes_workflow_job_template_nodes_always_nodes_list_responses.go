// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_template_nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListReader is a Reader for the WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesList structure.
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK creates a WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK with default headers values
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK() *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK{}
}

/*
WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK describes a response with status code 200, with default header values.

WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK workflow job template nodes workflow job template nodes always nodes list o k
*/
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK struct {
}

// IsSuccess returns true when this workflow job template nodes workflow job template nodes always nodes list o k response has a 2xx status code
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow job template nodes workflow job template nodes always nodes list o k response has a 3xx status code
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow job template nodes workflow job template nodes always nodes list o k response has a 4xx status code
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow job template nodes workflow job template nodes always nodes list o k response has a 5xx status code
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow job template nodes workflow job template nodes always nodes list o k response a status code equal to that given
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_template_nodes/{id}/always_nodes/][%d] workflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK ", 200)
}

func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_template_nodes/{id}/always_nodes/][%d] workflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK ", 200)
}

func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesAlwaysNodesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
