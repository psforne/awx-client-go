// Code generated by go-swagger; DO NOT EDIT.

package workflow_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WorkflowJobsWorkflowJobsDeleteReader is a Reader for the WorkflowJobsWorkflowJobsDelete structure.
type WorkflowJobsWorkflowJobsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobsWorkflowJobsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewWorkflowJobsWorkflowJobsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewWorkflowJobsWorkflowJobsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobsWorkflowJobsDeleteNoContent creates a WorkflowJobsWorkflowJobsDeleteNoContent with default headers values
func NewWorkflowJobsWorkflowJobsDeleteNoContent() *WorkflowJobsWorkflowJobsDeleteNoContent {
	return &WorkflowJobsWorkflowJobsDeleteNoContent{}
}

/*
WorkflowJobsWorkflowJobsDeleteNoContent describes a response with status code 204, with default header values.

WorkflowJobsWorkflowJobsDeleteNoContent workflow jobs workflow jobs delete no content
*/
type WorkflowJobsWorkflowJobsDeleteNoContent struct {
}

// IsSuccess returns true when this workflow jobs workflow jobs delete no content response has a 2xx status code
func (o *WorkflowJobsWorkflowJobsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow jobs workflow jobs delete no content response has a 3xx status code
func (o *WorkflowJobsWorkflowJobsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow jobs workflow jobs delete no content response has a 4xx status code
func (o *WorkflowJobsWorkflowJobsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow jobs workflow jobs delete no content response has a 5xx status code
func (o *WorkflowJobsWorkflowJobsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow jobs workflow jobs delete no content response a status code equal to that given
func (o *WorkflowJobsWorkflowJobsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *WorkflowJobsWorkflowJobsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workflow_jobs/{id}/][%d] workflowJobsWorkflowJobsDeleteNoContent ", 204)
}

func (o *WorkflowJobsWorkflowJobsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/workflow_jobs/{id}/][%d] workflowJobsWorkflowJobsDeleteNoContent ", 204)
}

func (o *WorkflowJobsWorkflowJobsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWorkflowJobsWorkflowJobsDeleteForbidden creates a WorkflowJobsWorkflowJobsDeleteForbidden with default headers values
func NewWorkflowJobsWorkflowJobsDeleteForbidden() *WorkflowJobsWorkflowJobsDeleteForbidden {
	return &WorkflowJobsWorkflowJobsDeleteForbidden{}
}

/*
WorkflowJobsWorkflowJobsDeleteForbidden describes a response with status code 403, with default header values.

forbidden
*/
type WorkflowJobsWorkflowJobsDeleteForbidden struct {
}

// IsSuccess returns true when this workflow jobs workflow jobs delete forbidden response has a 2xx status code
func (o *WorkflowJobsWorkflowJobsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this workflow jobs workflow jobs delete forbidden response has a 3xx status code
func (o *WorkflowJobsWorkflowJobsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow jobs workflow jobs delete forbidden response has a 4xx status code
func (o *WorkflowJobsWorkflowJobsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this workflow jobs workflow jobs delete forbidden response has a 5xx status code
func (o *WorkflowJobsWorkflowJobsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow jobs workflow jobs delete forbidden response a status code equal to that given
func (o *WorkflowJobsWorkflowJobsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *WorkflowJobsWorkflowJobsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workflow_jobs/{id}/][%d] workflowJobsWorkflowJobsDeleteForbidden ", 403)
}

func (o *WorkflowJobsWorkflowJobsDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/workflow_jobs/{id}/][%d] workflowJobsWorkflowJobsDeleteForbidden ", 403)
}

func (o *WorkflowJobsWorkflowJobsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}