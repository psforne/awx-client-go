// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobNodesWorkflowJobNodesSuccessNodesListReader is a Reader for the WorkflowJobNodesWorkflowJobNodesSuccessNodesList structure.
type WorkflowJobNodesWorkflowJobNodesSuccessNodesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobNodesWorkflowJobNodesSuccessNodesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobNodesWorkflowJobNodesSuccessNodesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobNodesWorkflowJobNodesSuccessNodesListOK creates a WorkflowJobNodesWorkflowJobNodesSuccessNodesListOK with default headers values
func NewWorkflowJobNodesWorkflowJobNodesSuccessNodesListOK() *WorkflowJobNodesWorkflowJobNodesSuccessNodesListOK {
	return &WorkflowJobNodesWorkflowJobNodesSuccessNodesListOK{}
}

/*
WorkflowJobNodesWorkflowJobNodesSuccessNodesListOK describes a response with status code 200, with default header values.

WorkflowJobNodesWorkflowJobNodesSuccessNodesListOK workflow job nodes workflow job nodes success nodes list o k
*/
type WorkflowJobNodesWorkflowJobNodesSuccessNodesListOK struct {
}

// IsSuccess returns true when this workflow job nodes workflow job nodes success nodes list o k response has a 2xx status code
func (o *WorkflowJobNodesWorkflowJobNodesSuccessNodesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow job nodes workflow job nodes success nodes list o k response has a 3xx status code
func (o *WorkflowJobNodesWorkflowJobNodesSuccessNodesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow job nodes workflow job nodes success nodes list o k response has a 4xx status code
func (o *WorkflowJobNodesWorkflowJobNodesSuccessNodesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow job nodes workflow job nodes success nodes list o k response has a 5xx status code
func (o *WorkflowJobNodesWorkflowJobNodesSuccessNodesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow job nodes workflow job nodes success nodes list o k response a status code equal to that given
func (o *WorkflowJobNodesWorkflowJobNodesSuccessNodesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *WorkflowJobNodesWorkflowJobNodesSuccessNodesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_nodes/{id}/success_nodes/][%d] workflowJobNodesWorkflowJobNodesSuccessNodesListOK ", 200)
}

func (o *WorkflowJobNodesWorkflowJobNodesSuccessNodesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_nodes/{id}/success_nodes/][%d] workflowJobNodesWorkflowJobNodesSuccessNodesListOK ", 200)
}

func (o *WorkflowJobNodesWorkflowJobNodesSuccessNodesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}