// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesSchedulesList structure.
type WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK creates a WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK() *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK {
	return &WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK{}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK describes a response with status code 200, with default header values.

WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK workflow job templates workflow job templates schedules list o k
*/
type WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK struct {
}

// IsSuccess returns true when this workflow job templates workflow job templates schedules list o k response has a 2xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow job templates workflow job templates schedules list o k response has a 3xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow job templates workflow job templates schedules list o k response has a 4xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow job templates workflow job templates schedules list o k response has a 5xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow job templates workflow job templates schedules list o k response a status code equal to that given
func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_templates/{id}/schedules/][%d] workflowJobTemplatesWorkflowJobTemplatesSchedulesListOK ", 200)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workflow_job_templates/{id}/schedules/][%d] workflowJobTemplatesWorkflowJobTemplatesSchedulesListOK ", 200)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesSchedulesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
