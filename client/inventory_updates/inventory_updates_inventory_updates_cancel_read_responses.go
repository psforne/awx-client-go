// Code generated by go-swagger; DO NOT EDIT.

package inventory_updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoryUpdatesInventoryUpdatesCancelReadReader is a Reader for the InventoryUpdatesInventoryUpdatesCancelRead structure.
type InventoryUpdatesInventoryUpdatesCancelReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryUpdatesInventoryUpdatesCancelReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoryUpdatesInventoryUpdatesCancelReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventoryUpdatesInventoryUpdatesCancelReadOK creates a InventoryUpdatesInventoryUpdatesCancelReadOK with default headers values
func NewInventoryUpdatesInventoryUpdatesCancelReadOK() *InventoryUpdatesInventoryUpdatesCancelReadOK {
	return &InventoryUpdatesInventoryUpdatesCancelReadOK{}
}

/*
InventoryUpdatesInventoryUpdatesCancelReadOK describes a response with status code 200, with default header values.

InventoryUpdatesInventoryUpdatesCancelReadOK inventory updates inventory updates cancel read o k
*/
type InventoryUpdatesInventoryUpdatesCancelReadOK struct {
}

// IsSuccess returns true when this inventory updates inventory updates cancel read o k response has a 2xx status code
func (o *InventoryUpdatesInventoryUpdatesCancelReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventory updates inventory updates cancel read o k response has a 3xx status code
func (o *InventoryUpdatesInventoryUpdatesCancelReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventory updates inventory updates cancel read o k response has a 4xx status code
func (o *InventoryUpdatesInventoryUpdatesCancelReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventory updates inventory updates cancel read o k response has a 5xx status code
func (o *InventoryUpdatesInventoryUpdatesCancelReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this inventory updates inventory updates cancel read o k response a status code equal to that given
func (o *InventoryUpdatesInventoryUpdatesCancelReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *InventoryUpdatesInventoryUpdatesCancelReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventory_updates/{id}/cancel/][%d] inventoryUpdatesInventoryUpdatesCancelReadOK ", 200)
}

func (o *InventoryUpdatesInventoryUpdatesCancelReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/inventory_updates/{id}/cancel/][%d] inventoryUpdatesInventoryUpdatesCancelReadOK ", 200)
}

func (o *InventoryUpdatesInventoryUpdatesCancelReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
