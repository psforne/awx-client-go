// Code generated by go-swagger; DO NOT EDIT.

package inventory_updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoryUpdatesInventoryUpdatesEventsListReader is a Reader for the InventoryUpdatesInventoryUpdatesEventsList structure.
type InventoryUpdatesInventoryUpdatesEventsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryUpdatesInventoryUpdatesEventsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoryUpdatesInventoryUpdatesEventsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventoryUpdatesInventoryUpdatesEventsListOK creates a InventoryUpdatesInventoryUpdatesEventsListOK with default headers values
func NewInventoryUpdatesInventoryUpdatesEventsListOK() *InventoryUpdatesInventoryUpdatesEventsListOK {
	return &InventoryUpdatesInventoryUpdatesEventsListOK{}
}

/*
InventoryUpdatesInventoryUpdatesEventsListOK describes a response with status code 200, with default header values.

InventoryUpdatesInventoryUpdatesEventsListOK inventory updates inventory updates events list o k
*/
type InventoryUpdatesInventoryUpdatesEventsListOK struct {
}

// IsSuccess returns true when this inventory updates inventory updates events list o k response has a 2xx status code
func (o *InventoryUpdatesInventoryUpdatesEventsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventory updates inventory updates events list o k response has a 3xx status code
func (o *InventoryUpdatesInventoryUpdatesEventsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventory updates inventory updates events list o k response has a 4xx status code
func (o *InventoryUpdatesInventoryUpdatesEventsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventory updates inventory updates events list o k response has a 5xx status code
func (o *InventoryUpdatesInventoryUpdatesEventsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this inventory updates inventory updates events list o k response a status code equal to that given
func (o *InventoryUpdatesInventoryUpdatesEventsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *InventoryUpdatesInventoryUpdatesEventsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventory_updates/{id}/events/][%d] inventoryUpdatesInventoryUpdatesEventsListOK ", 200)
}

func (o *InventoryUpdatesInventoryUpdatesEventsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/inventory_updates/{id}/events/][%d] inventoryUpdatesInventoryUpdatesEventsListOK ", 200)
}

func (o *InventoryUpdatesInventoryUpdatesEventsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}