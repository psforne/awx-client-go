// Code generated by go-swagger; DO NOT EDIT.

package workflow_job_template_nodes

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

// WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateReader is a Reader for the WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdate structure.
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK creates a WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK with default headers values
func NewWorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK() *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK {
	return &WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK{}
}

/*
WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK describes a response with status code 200, with default header values.

WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK workflow job template nodes workflow job template nodes update o k
*/
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK struct {
}

// IsSuccess returns true when this workflow job template nodes workflow job template nodes update o k response has a 2xx status code
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow job template nodes workflow job template nodes update o k response has a 3xx status code
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow job template nodes workflow job template nodes update o k response has a 4xx status code
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow job template nodes workflow job template nodes update o k response has a 5xx status code
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow job template nodes workflow job template nodes update o k response a status code equal to that given
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/workflow_job_template_nodes/{id}/][%d] workflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK ", 200)
}

func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/workflow_job_template_nodes/{id}/][%d] workflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK ", 200)
}

func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateBody workflow job template nodes workflow job template nodes update body
swagger:model WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateBody
*/
type WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateBody struct {

	// If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node
	AllParentsMustConverge bool `json:"all_parents_must_converge,omitempty"`

	// diff mode
	DiffMode bool `json:"diff_mode,omitempty"`

	// extra data
	ExtraData string `json:"extra_data,omitempty"`

	// An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node.
	Identifier string `json:"identifier,omitempty"`

	// Inventory applied as a prompt, assuming job template prompts for inventory
	Inventory int64 `json:"inventory,omitempty"`

	// job tags
	JobTags string `json:"job_tags,omitempty"`

	// job type
	JobType string `json:"job_type,omitempty"`

	// limit
	Limit string `json:"limit,omitempty"`

	// scm branch
	ScmBranch string `json:"scm_branch,omitempty"`

	// skip tags
	SkipTags string `json:"skip_tags,omitempty"`

	// unified job template
	UnifiedJobTemplate int64 `json:"unified_job_template,omitempty"`

	// verbosity
	Verbosity string `json:"verbosity,omitempty"`

	// workflow job template
	// Required: true
	WorkflowJobTemplate *string `json:"workflow_job_template"`
}

// Validate validates this workflow job template nodes workflow job template nodes update body
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateWorkflowJobTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateBody) validateWorkflowJobTemplate(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"workflow_job_template", "body", o.WorkflowJobTemplate); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this workflow job template nodes workflow job template nodes update body based on context it is used
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateBody) UnmarshalBinary(b []byte) error {
	var res WorkflowJobTemplateNodesWorkflowJobTemplateNodesUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
