// Code generated by go-swagger; DO NOT EDIT.

package inventory_updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoryUpdatesInventoryUpdatesNotificationsListReader is a Reader for the InventoryUpdatesInventoryUpdatesNotificationsList structure.
type InventoryUpdatesInventoryUpdatesNotificationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryUpdatesInventoryUpdatesNotificationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoryUpdatesInventoryUpdatesNotificationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventoryUpdatesInventoryUpdatesNotificationsListOK creates a InventoryUpdatesInventoryUpdatesNotificationsListOK with default headers values
func NewInventoryUpdatesInventoryUpdatesNotificationsListOK() *InventoryUpdatesInventoryUpdatesNotificationsListOK {
	return &InventoryUpdatesInventoryUpdatesNotificationsListOK{}
}

/*
InventoryUpdatesInventoryUpdatesNotificationsListOK describes a response with status code 200, with default header values.

InventoryUpdatesInventoryUpdatesNotificationsListOK inventory updates inventory updates notifications list o k
*/
type InventoryUpdatesInventoryUpdatesNotificationsListOK struct {
}

// IsSuccess returns true when this inventory updates inventory updates notifications list o k response has a 2xx status code
func (o *InventoryUpdatesInventoryUpdatesNotificationsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventory updates inventory updates notifications list o k response has a 3xx status code
func (o *InventoryUpdatesInventoryUpdatesNotificationsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventory updates inventory updates notifications list o k response has a 4xx status code
func (o *InventoryUpdatesInventoryUpdatesNotificationsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventory updates inventory updates notifications list o k response has a 5xx status code
func (o *InventoryUpdatesInventoryUpdatesNotificationsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this inventory updates inventory updates notifications list o k response a status code equal to that given
func (o *InventoryUpdatesInventoryUpdatesNotificationsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *InventoryUpdatesInventoryUpdatesNotificationsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventory_updates/{id}/notifications/][%d] inventoryUpdatesInventoryUpdatesNotificationsListOK ", 200)
}

func (o *InventoryUpdatesInventoryUpdatesNotificationsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/inventory_updates/{id}/notifications/][%d] inventoryUpdatesInventoryUpdatesNotificationsListOK ", 200)
}

func (o *InventoryUpdatesInventoryUpdatesNotificationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}