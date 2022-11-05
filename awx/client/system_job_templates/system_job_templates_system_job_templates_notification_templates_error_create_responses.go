// Code generated by go-swagger; DO NOT EDIT.

package system_job_templates

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

// SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateReader is a Reader for the SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreate structure.
type SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated creates a SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated with default headers values
func NewSystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated() *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated {
	return &SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated{}
}

/*
SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated describes a response with status code 201, with default header values.

SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated system job templates system job templates notification templates error create created
*/
type SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated struct {
}

// IsSuccess returns true when this system job templates system job templates notification templates error create created response has a 2xx status code
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this system job templates system job templates notification templates error create created response has a 3xx status code
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system job templates system job templates notification templates error create created response has a 4xx status code
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this system job templates system job templates notification templates error create created response has a 5xx status code
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this system job templates system job templates notification templates error create created response a status code equal to that given
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/system_job_templates/{id}/notification_templates_error/][%d] systemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated ", 201)
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/system_job_templates/{id}/notification_templates_error/][%d] systemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated ", 201)
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateBody system job templates system job templates notification templates error create body
swagger:model SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateBody
*/
type SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateBody struct {

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

// Validate validates this system job templates system job templates notification templates error create body
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateBody) Validate(formats strfmt.Registry) error {
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

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateBody) validateNotificationType(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"notification_type", "body", o.NotificationType); err != nil {
		return err
	}

	return nil
}

func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateBody) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"organization", "body", o.Organization); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this system job templates system job templates notification templates error create body based on context it is used
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateBody) UnmarshalBinary(b []byte) error {
	var res SystemJobTemplatesSystemJobTemplatesNotificationTemplatesErrorCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}