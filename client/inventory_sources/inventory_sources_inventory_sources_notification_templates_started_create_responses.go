// Code generated by go-swagger; DO NOT EDIT.

package inventory_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventorySourcesInventorySourcesNotificationTemplatesStartedCreateReader is a Reader for the InventorySourcesInventorySourcesNotificationTemplatesStartedCreate structure.
type InventorySourcesInventorySourcesNotificationTemplatesStartedCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated creates a InventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated with default headers values
func NewInventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated() *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated {
	return &InventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated{}
}

/*
InventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated describes a response with status code 201, with default header values.

InventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated inventory sources inventory sources notification templates started create created
*/
type InventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated struct {
}

// IsSuccess returns true when this inventory sources inventory sources notification templates started create created response has a 2xx status code
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventory sources inventory sources notification templates started create created response has a 3xx status code
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventory sources inventory sources notification templates started create created response has a 4xx status code
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventory sources inventory sources notification templates started create created response has a 5xx status code
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this inventory sources inventory sources notification templates started create created response a status code equal to that given
func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/inventory_sources/{id}/notification_templates_started/][%d] inventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated ", 201)
}

func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/inventory_sources/{id}/notification_templates_started/][%d] inventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated ", 201)
}

func (o *InventorySourcesInventorySourcesNotificationTemplatesStartedCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
