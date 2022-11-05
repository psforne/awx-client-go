// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoriesInventoriesInventorySourcesCreateReader is a Reader for the InventoriesInventoriesInventorySourcesCreate structure.
type InventoriesInventoriesInventorySourcesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoriesInventoriesInventorySourcesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInventoriesInventoriesInventorySourcesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInventoriesInventoriesInventorySourcesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInventoriesInventoriesInventorySourcesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventoriesInventoriesInventorySourcesCreateCreated creates a InventoriesInventoriesInventorySourcesCreateCreated with default headers values
func NewInventoriesInventoriesInventorySourcesCreateCreated() *InventoriesInventoriesInventorySourcesCreateCreated {
	return &InventoriesInventoriesInventorySourcesCreateCreated{}
}

/*
InventoriesInventoriesInventorySourcesCreateCreated describes a response with status code 201, with default header values.

InventoriesInventoriesInventorySourcesCreateCreated inventories inventories inventory sources create created
*/
type InventoriesInventoriesInventorySourcesCreateCreated struct {
}

// IsSuccess returns true when this inventories inventories inventory sources create created response has a 2xx status code
func (o *InventoriesInventoriesInventorySourcesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventories inventories inventory sources create created response has a 3xx status code
func (o *InventoriesInventoriesInventorySourcesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventories inventories inventory sources create created response has a 4xx status code
func (o *InventoriesInventoriesInventorySourcesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventories inventories inventory sources create created response has a 5xx status code
func (o *InventoriesInventoriesInventorySourcesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this inventories inventories inventory sources create created response a status code equal to that given
func (o *InventoriesInventoriesInventorySourcesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *InventoriesInventoriesInventorySourcesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/inventories/{id}/inventory_sources/][%d] inventoriesInventoriesInventorySourcesCreateCreated ", 201)
}

func (o *InventoriesInventoriesInventorySourcesCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/inventories/{id}/inventory_sources/][%d] inventoriesInventoriesInventorySourcesCreateCreated ", 201)
}

func (o *InventoriesInventoriesInventorySourcesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInventoriesInventoriesInventorySourcesCreateBadRequest creates a InventoriesInventoriesInventorySourcesCreateBadRequest with default headers values
func NewInventoriesInventoriesInventorySourcesCreateBadRequest() *InventoriesInventoriesInventorySourcesCreateBadRequest {
	return &InventoriesInventoriesInventorySourcesCreateBadRequest{}
}

/*
InventoriesInventoriesInventorySourcesCreateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type InventoriesInventoriesInventorySourcesCreateBadRequest struct {
}

// IsSuccess returns true when this inventories inventories inventory sources create bad request response has a 2xx status code
func (o *InventoriesInventoriesInventorySourcesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this inventories inventories inventory sources create bad request response has a 3xx status code
func (o *InventoriesInventoriesInventorySourcesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventories inventories inventory sources create bad request response has a 4xx status code
func (o *InventoriesInventoriesInventorySourcesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this inventories inventories inventory sources create bad request response has a 5xx status code
func (o *InventoriesInventoriesInventorySourcesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this inventories inventories inventory sources create bad request response a status code equal to that given
func (o *InventoriesInventoriesInventorySourcesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *InventoriesInventoriesInventorySourcesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/inventories/{id}/inventory_sources/][%d] inventoriesInventoriesInventorySourcesCreateBadRequest ", 400)
}

func (o *InventoriesInventoriesInventorySourcesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/inventories/{id}/inventory_sources/][%d] inventoriesInventoriesInventorySourcesCreateBadRequest ", 400)
}

func (o *InventoriesInventoriesInventorySourcesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInventoriesInventoriesInventorySourcesCreateForbidden creates a InventoriesInventoriesInventorySourcesCreateForbidden with default headers values
func NewInventoriesInventoriesInventorySourcesCreateForbidden() *InventoriesInventoriesInventorySourcesCreateForbidden {
	return &InventoriesInventoriesInventorySourcesCreateForbidden{}
}

/*
InventoriesInventoriesInventorySourcesCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type InventoriesInventoriesInventorySourcesCreateForbidden struct {
}

// IsSuccess returns true when this inventories inventories inventory sources create forbidden response has a 2xx status code
func (o *InventoriesInventoriesInventorySourcesCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this inventories inventories inventory sources create forbidden response has a 3xx status code
func (o *InventoriesInventoriesInventorySourcesCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventories inventories inventory sources create forbidden response has a 4xx status code
func (o *InventoriesInventoriesInventorySourcesCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this inventories inventories inventory sources create forbidden response has a 5xx status code
func (o *InventoriesInventoriesInventorySourcesCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this inventories inventories inventory sources create forbidden response a status code equal to that given
func (o *InventoriesInventoriesInventorySourcesCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *InventoriesInventoriesInventorySourcesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/inventories/{id}/inventory_sources/][%d] inventoriesInventoriesInventorySourcesCreateForbidden ", 403)
}

func (o *InventoriesInventoriesInventorySourcesCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/inventories/{id}/inventory_sources/][%d] inventoriesInventoriesInventorySourcesCreateForbidden ", 403)
}

func (o *InventoriesInventoriesInventorySourcesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
