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

// WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateReader is a Reader for the WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreate structure.
type WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated creates a WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated with default headers values
func NewWorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated() *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated {
	return &WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated{}
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated describes a response with status code 201, with default header values.

WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated workflow job templates workflow job templates notification templates error create created
*/
type WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated struct {
}

// IsSuccess returns true when this workflow job templates workflow job templates notification templates error create created response has a 2xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workflow job templates workflow job templates notification templates error create created response has a 3xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workflow job templates workflow job templates notification templates error create created response has a 4xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this workflow job templates workflow job templates notification templates error create created response has a 5xx status code
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this workflow job templates workflow job templates notification templates error create created response a status code equal to that given
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/{id}/notification_templates_error/][%d] workflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated ", 201)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/workflow_job_templates/{id}/notification_templates_error/][%d] workflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated ", 201)
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateBody workflow job templates workflow job templates notification templates error create body
swagger:model WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateBody
*/
type WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// Optional custom messages for notification template.
	Messages string `json:"messages,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// notification configuration
	NotificationConfiguration string `json:"notification_configuration,omitempty"`

	// notification type
	// Required: true
	NotificationType *string `json:"notification_type"`

	// organization
	// Required: true
	Organization *int64 `json:"organization"`
}

// Validate validates this workflow job templates workflow job templates notification templates error create body
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNotificationType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateBody) validateNotificationType(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"notification_type", "body", o.NotificationType); err != nil {
		return err
	}

	return nil
}

func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateBody) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"organization", "body", o.Organization); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this workflow job templates workflow job templates notification templates error create body based on context it is used
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateBody) UnmarshalBinary(b []byte) error {
	var res WorkflowJobTemplatesWorkflowJobTemplatesNotificationTemplatesErrorCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
