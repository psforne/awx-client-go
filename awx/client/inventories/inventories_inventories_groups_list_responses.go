// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoriesInventoriesGroupsListReader is a Reader for the InventoriesInventoriesGroupsList structure.
type InventoriesInventoriesGroupsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoriesInventoriesGroupsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoriesInventoriesGroupsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventoriesInventoriesGroupsListOK creates a InventoriesInventoriesGroupsListOK with default headers values
func NewInventoriesInventoriesGroupsListOK() *InventoriesInventoriesGroupsListOK {
	return &InventoriesInventoriesGroupsListOK{}
}

/*
InventoriesInventoriesGroupsListOK describes a response with status code 200, with default header values.

InventoriesInventoriesGroupsListOK inventories inventories groups list o k
*/
type InventoriesInventoriesGroupsListOK struct {
}

// IsSuccess returns true when this inventories inventories groups list o k response has a 2xx status code
func (o *InventoriesInventoriesGroupsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventories inventories groups list o k response has a 3xx status code
func (o *InventoriesInventoriesGroupsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventories inventories groups list o k response has a 4xx status code
func (o *InventoriesInventoriesGroupsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventories inventories groups list o k response has a 5xx status code
func (o *InventoriesInventoriesGroupsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this inventories inventories groups list o k response a status code equal to that given
func (o *InventoriesInventoriesGroupsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *InventoriesInventoriesGroupsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventories/{id}/groups/][%d] inventoriesInventoriesGroupsListOK ", 200)
}

func (o *InventoriesInventoriesGroupsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/inventories/{id}/groups/][%d] inventoriesInventoriesGroupsListOK ", 200)
}

func (o *InventoriesInventoriesGroupsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
