// Code generated by go-swagger; DO NOT EDIT.

package inventory_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventorySourcesInventorySourcesGroupsDeleteReader is a Reader for the InventorySourcesInventorySourcesGroupsDelete structure.
type InventorySourcesInventorySourcesGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventorySourcesInventorySourcesGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewInventorySourcesInventorySourcesGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventorySourcesInventorySourcesGroupsDeleteNoContent creates a InventorySourcesInventorySourcesGroupsDeleteNoContent with default headers values
func NewInventorySourcesInventorySourcesGroupsDeleteNoContent() *InventorySourcesInventorySourcesGroupsDeleteNoContent {
	return &InventorySourcesInventorySourcesGroupsDeleteNoContent{}
}

/*
InventorySourcesInventorySourcesGroupsDeleteNoContent describes a response with status code 204, with default header values.

InventorySourcesInventorySourcesGroupsDeleteNoContent inventory sources inventory sources groups delete no content
*/
type InventorySourcesInventorySourcesGroupsDeleteNoContent struct {
}

// IsSuccess returns true when this inventory sources inventory sources groups delete no content response has a 2xx status code
func (o *InventorySourcesInventorySourcesGroupsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventory sources inventory sources groups delete no content response has a 3xx status code
func (o *InventorySourcesInventorySourcesGroupsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventory sources inventory sources groups delete no content response has a 4xx status code
func (o *InventorySourcesInventorySourcesGroupsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventory sources inventory sources groups delete no content response has a 5xx status code
func (o *InventorySourcesInventorySourcesGroupsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this inventory sources inventory sources groups delete no content response a status code equal to that given
func (o *InventorySourcesInventorySourcesGroupsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *InventorySourcesInventorySourcesGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/inventory_sources/{id}/groups/][%d] inventorySourcesInventorySourcesGroupsDeleteNoContent ", 204)
}

func (o *InventorySourcesInventorySourcesGroupsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/inventory_sources/{id}/groups/][%d] inventorySourcesInventorySourcesGroupsDeleteNoContent ", 204)
}

func (o *InventorySourcesInventorySourcesGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
