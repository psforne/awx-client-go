// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkflowJobTemplatesWorkflowJobTemplatesCreateReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesCreate structure.
type WorkflowJobTemplatesWorkflowJobTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesCreateCreated creates a WorkflowJobTemplatesWorkflowJobTemplatesCreateCreated with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesCreateCreated() *WorkflowJobTemplatesWorkflowJobTemplatesCreateCreated {
	return &WorkflowJobTemplatesWorkflowJobTemplatesCreateCreated{}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesCreateCreated describes a response with status code 201, with default header values.

WorkflowJobTemplatesWorkflowJobTemplatesCreateCreated workflow job templates workflow job templates create created
*/
type WorkflowJobTemplatesWorkflowJobTemplatesCreateCreated struct {
}

// IsSuccess returns true when this workflow job templates workflow job templates create created response has a 2xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow job templates workflow job templates create created response has a 3xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow job templates workflow job templates create created response has a 4xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow job templates workflow job templates create created response has a 5xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow job templates workflow job templates create created response a status code equal to that given
func (o *WorkflowJobTemplatesWorkflowJobTemplatesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/][%d] workflowJobTemplatesWorkflowJobTemplatesCreateCreated ", 201)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/][%d] workflowJobTemplatesWorkflowJobTemplatesCreateCreated ", 201)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesCreateBody workflow job templates workflow job templates create body
swagger:model WorkflowJobTemplatesWorkflowJobTemplatesCreateBody
*/
type WorkflowJobTemplatesWorkflowJobTemplatesCreateBody struct {

	// allow simultaneous
	AllowSimultaneous bool `json:"allow_simultaneous,omitempty"`

	// ask inventory on launch
	AskInventoryOnLaunch bool `json:"ask_inventory_on_launch,omitempty"`

	// ask limit on launch
	AskLimitOnLaunch bool `json:"ask_limit_on_launch,omitempty"`

	// ask scm branch on launch
	AskScmBranchOnLaunch bool `json:"ask_scm_branch_on_launch,omitempty"`

	// ask variables on launch
	AskVariablesOnLaunch bool `json:"ask_variables_on_launch,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// extra vars
	ExtraVars string `json:"extra_vars,omitempty"`

	// Inventory applied as a prompt, assuming job template prompts for inventory
	Inventory int64 `json:"inventory,omitempty"`

	// limit
	Limit string `json:"limit,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// The organization used to determine access to this template.
	Organization int64 `json:"organization,omitempty"`

	// scm branch
	ScmBranch string `json:"scm_branch,omitempty"`

	// survey enabled
	SurveyEnabled bool `json:"survey_enabled,omitempty"`

	// Personal Access Token for posting back the status to the service API
	WebhookCredential int64 `json:"webhook_credential,omitempty"`

	// Service that webhook requests will be accepted from
	WebhookService string `json:"webhook_service,omitempty"`
}

// Validate validates this workflow job templates workflow job templates create body
func (o *WorkflowJobTemplatesWorkflowJobTemplatesCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this workflow job templates workflow job templates create body based on context it is used
func (o *WorkflowJobTemplatesWorkflowJobTemplatesCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *WorkflowJobTemplatesWorkflowJobTemplatesCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WorkflowJobTemplatesWorkflowJobTemplatesCreateBody) UnmarshalBinary(b []byte) error {
	var res WorkflowJobTemplatesWorkflowJobTemplatesCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
