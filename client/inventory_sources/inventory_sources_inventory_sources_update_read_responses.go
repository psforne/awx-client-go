// Code generated by go-swagger; DO NOT EDIT.

package inventory_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventorySourcesInventorySourcesUpdateReadReader is a Reader for the InventorySourcesInventorySourcesUpdateRead structure.
type InventorySourcesInventorySourcesUpdateReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventorySourcesInventorySourcesUpdateReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventorySourcesInventorySourcesUpdateReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventorySourcesInventorySourcesUpdateReadOK creates a InventorySourcesInventorySourcesUpdateReadOK with default headers values
func NewInventorySourcesInventorySourcesUpdateReadOK() *InventorySourcesInventorySourcesUpdateReadOK {
	return &InventorySourcesInventorySourcesUpdateReadOK{}
}

/*
InventorySourcesInventorySourcesUpdateReadOK describes a response with status code 200, with default header values.

InventorySourcesInventorySourcesUpdateReadOK inventory sources inventory sources update read o k
*/
type InventorySourcesInventorySourcesUpdateReadOK struct {
}

// IsSuccess returns true when this inventory sources inventory sources update read o k response has a 2xx status code
func (o *InventorySourcesInventorySourcesUpdateReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventory sources inventory sources update read o k response has a 3xx status code
func (o *InventorySourcesInventorySourcesUpdateReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventory sources inventory sources update read o k response has a 4xx status code
func (o *InventorySourcesInventorySourcesUpdateReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventory sources inventory sources update read o k response has a 5xx status code
func (o *InventorySourcesInventorySourcesUpdateReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this inventory sources inventory sources update read o k response a status code equal to that given
func (o *InventorySourcesInventorySourcesUpdateReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *InventorySourcesInventorySourcesUpdateReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventory_sources/{id}/update/][%d] inventorySourcesInventorySourcesUpdateReadOK ", 200)
}

func (o *InventorySourcesInventorySourcesUpdateReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/inventory_sources/{id}/update/][%d] inventorySourcesInventorySourcesUpdateReadOK ", 200)
}

func (o *InventorySourcesInventorySourcesUpdateReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
