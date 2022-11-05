// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreate structure.
type WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated creates a WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated() *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated {
	return &WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated{}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated describes a response with status code 201, with default header values.

WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated workflow job templates workflow job templates launch create created
*/
type WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated struct {
}

// IsSuccess returns true when this workflow job templates workflow job templates launch create created response has a 2xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow job templates workflow job templates launch create created response has a 3xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow job templates workflow job templates launch create created response has a 4xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow job templates workflow job templates launch create created response has a 5xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow job templates workflow job templates launch create created response a status code equal to that given
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/{id}/launch/][%d] workflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated ", 201)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/{id}/launch/][%d] workflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated ", 201)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest creates a WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest() *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest {
	return &WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest{}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest struct {
}

// IsSuccess returns true when this workflow job templates workflow job templates launch create bad request response has a 2xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this workflow job templates workflow job templates launch create bad request response has a 3xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow job templates workflow job templates launch create bad request response has a 4xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this workflow job templates workflow job templates launch create bad request response has a 5xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow job templates workflow job templates launch create bad request response a status code equal to that given
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/{id}/launch/][%d] workflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest ", 400)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/{id}/launch/][%d] workflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest ", 400)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBody workflow job templates workflow job templates launch create body
swagger:model WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBody
*/
type WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBody struct {

	// ask limit on launch
	AskLimitOnLaunch bool `json:"ask_limit_on_launch,omitempty"`

	// ask scm branch on launch
	AskScmBranchOnLaunch bool `json:"ask_scm_branch_on_launch,omitempty"`

	// extra vars
	ExtraVars string `json:"extra_vars,omitempty"`

	// inventory
	Inventory int64 `json:"inventory,omitempty"`

	// limit
	Limit string `json:"limit,omitempty"`

	// scm branch
	ScmBranch string `json:"scm_branch,omitempty"`
}

// Validate validates this workflow job templates workflow job templates launch create body
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this workflow job templates workflow job templates launch create body based on context it is used
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBody) UnmarshalBinary(b []byte) error {
	var res WorkflowJobTemplatesWorkflowJobTemplatesLaunchCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}