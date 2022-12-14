// Code generated by go-swagger; DO NOT EDIT.

package notification_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// NotificationTemplatesNotificationTemplatesCopyListReader is a Reader for the NotificationTemplatesNotificationTemplatesCopyList structure.
type NotificationTemplatesNotificationTemplatesCopyListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationTemplatesNotificationTemplatesCopyListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationTemplatesNotificationTemplatesCopyListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNotificationTemplatesNotificationTemplatesCopyListOK creates a NotificationTemplatesNotificationTemplatesCopyListOK with default headers values
func NewNotificationTemplatesNotificationTemplatesCopyListOK() *NotificationTemplatesNotificationTemplatesCopyListOK {
	return &NotificationTemplatesNotificationTemplatesCopyListOK{}
}

/*
NotificationTemplatesNotificationTemplatesCopyListOK describes a response with status code 200, with default header values.

NotificationTemplatesNotificationTemplatesCopyListOK notification templates notification templates copy list o k
*/
type NotificationTemplatesNotificationTemplatesCopyListOK struct {
}

// IsSuccess returns true when this notification templates notification templates copy list o k response has a 2xx status code
func (o *NotificationTemplatesNotificationTemplatesCopyListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notification templates notification templates copy list o k response has a 3xx status code
func (o *NotificationTemplatesNotificationTemplatesCopyListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification templates notification templates copy list o k response has a 4xx status code
func (o *NotificationTemplatesNotificationTemplatesCopyListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this notification templates notification templates copy list o k response has a 5xx status code
func (o *NotificationTemplatesNotificationTemplatesCopyListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this notification templates notification templates copy list o k response a status code equal to that given
func (o *NotificationTemplatesNotificationTemplatesCopyListOK) IsCode(code int) bool {
	return code == 200
}

func (o *NotificationTemplatesNotificationTemplatesCopyListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/notification_templates/{id}/copy/][%d] notificationTemplatesNotificationTemplatesCopyListOK ", 200)
}

func (o *NotificationTemplatesNotificationTemplatesCopyListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/notification_templates/{id}/copy/][%d] notificationTemplatesNotificationTemplatesCopyListOK ", 200)
}

func (o *NotificationTemplatesNotificationTemplatesCopyListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
