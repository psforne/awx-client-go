// Code generated by go-swagger; DO NOT EDIT.

package inventory_updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoryUpdatesInventoryUpdatesCredentialsListReader is a Reader for the InventoryUpdatesInventoryUpdatesCredentialsList structure.
type InventoryUpdatesInventoryUpdatesCredentialsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryUpdatesInventoryUpdatesCredentialsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoryUpdatesInventoryUpdatesCredentialsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventoryUpdatesInventoryUpdatesCredentialsListOK creates a InventoryUpdatesInventoryUpdatesCredentialsListOK with default headers values
func NewInventoryUpdatesInventoryUpdatesCredentialsListOK() *InventoryUpdatesInventoryUpdatesCredentialsListOK {
	return &InventoryUpdatesInventoryUpdatesCredentialsListOK{}
}

/*
InventoryUpdatesInventoryUpdatesCredentialsListOK describes a response with status code 200, with default header values.

InventoryUpdatesInventoryUpdatesCredentialsListOK inventory updates inventory updates credentials list o k
*/
type InventoryUpdatesInventoryUpdatesCredentialsListOK struct {
}

// IsSuccess returns true when this inventory updates inventory updates credentials list o k response has a 2xx status code
func (o *InventoryUpdatesInventoryUpdatesCredentialsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventory updates inventory updates credentials list o k response has a 3xx status code
func (o *InventoryUpdatesInventoryUpdatesCredentialsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventory updates inventory updates credentials list o k response has a 4xx status code
func (o *InventoryUpdatesInventoryUpdatesCredentialsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventory updates inventory updates credentials list o k response has a 5xx status code
func (o *InventoryUpdatesInventoryUpdatesCredentialsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this inventory updates inventory updates credentials list o k response a status code equal to that given
func (o *InventoryUpdatesInventoryUpdatesCredentialsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *InventoryUpdatesInventoryUpdatesCredentialsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventory_updates/{id}/credentials/][%d] inventoryUpdatesInventoryUpdatesCredentialsListOK ", 200)
}

func (o *InventoryUpdatesInventoryUpdatesCredentialsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/inventory_updates/{id}/credentials/][%d] inventoryUpdatesInventoryUpdatesCredentialsListOK ", 200)
}

func (o *InventoryUpdatesInventoryUpdatesCredentialsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
