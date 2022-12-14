// Code generated by go-swagger; DO NOT EDIT.

package inventory_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventorySourcesInventorySourcesPartialUpdateReader is a Reader for the InventorySourcesInventorySourcesPartialUpdate structure.
type InventorySourcesInventorySourcesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventorySourcesInventorySourcesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventorySourcesInventorySourcesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInventorySourcesInventorySourcesPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventorySourcesInventorySourcesPartialUpdateOK creates a InventorySourcesInventorySourcesPartialUpdateOK with default headers values
func NewInventorySourcesInventorySourcesPartialUpdateOK() *InventorySourcesInventorySourcesPartialUpdateOK {
	return &InventorySourcesInventorySourcesPartialUpdateOK{}
}

/*
InventorySourcesInventorySourcesPartialUpdateOK describes a response with status code 200, with default header values.

InventorySourcesInventorySourcesPartialUpdateOK inventory sources inventory sources partial update o k
*/
type InventorySourcesInventorySourcesPartialUpdateOK struct {
}

// IsSuccess returns true when this inventory sources inventory sources partial update o k response has a 2xx status code
func (o *InventorySourcesInventorySourcesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventory sources inventory sources partial update o k response has a 3xx status code
func (o *InventorySourcesInventorySourcesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventory sources inventory sources partial update o k response has a 4xx status code
func (o *InventorySourcesInventorySourcesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventory sources inventory sources partial update o k response has a 5xx status code
func (o *InventorySourcesInventorySourcesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this inventory sources inventory sources partial update o k response a status code equal to that given
func (o *InventorySourcesInventorySourcesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *InventorySourcesInventorySourcesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/inventory_sources/{id}/][%d] inventorySourcesInventorySourcesPartialUpdateOK ", 200)
}

func (o *InventorySourcesInventorySourcesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/inventory_sources/{id}/][%d] inventorySourcesInventorySourcesPartialUpdateOK ", 200)
}

func (o *InventorySourcesInventorySourcesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInventorySourcesInventorySourcesPartialUpdateBadRequest creates a InventorySourcesInventorySourcesPartialUpdateBadRequest with default headers values
func NewInventorySourcesInventorySourcesPartialUpdateBadRequest() *InventorySourcesInventorySourcesPartialUpdateBadRequest {
	return &InventorySourcesInventorySourcesPartialUpdateBadRequest{}
}

/*
InventorySourcesInventorySourcesPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type InventorySourcesInventorySourcesPartialUpdateBadRequest struct {
}

// IsSuccess returns true when this inventory sources inventory sources partial update bad request response has a 2xx status code
func (o *InventorySourcesInventorySourcesPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this inventory sources inventory sources partial update bad request response has a 3xx status code
func (o *InventorySourcesInventorySourcesPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventory sources inventory sources partial update bad request response has a 4xx status code
func (o *InventorySourcesInventorySourcesPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this inventory sources inventory sources partial update bad request response has a 5xx status code
func (o *InventorySourcesInventorySourcesPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this inventory sources inventory sources partial update bad request response a status code equal to that given
func (o *InventorySourcesInventorySourcesPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *InventorySourcesInventorySourcesPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/inventory_sources/{id}/][%d] inventorySourcesInventorySourcesPartialUpdateBadRequest ", 400)
}

func (o *InventorySourcesInventorySourcesPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/inventory_sources/{id}/][%d] inventorySourcesInventorySourcesPartialUpdateBadRequest ", 400)
}

func (o *InventorySourcesInventorySourcesPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
