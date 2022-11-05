// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoriesInventoriesObjectRolesListReader is a Reader for the InventoriesInventoriesObjectRolesList structure.
type InventoriesInventoriesObjectRolesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoriesInventoriesObjectRolesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoriesInventoriesObjectRolesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventoriesInventoriesObjectRolesListOK creates a InventoriesInventoriesObjectRolesListOK with default headers values
func NewInventoriesInventoriesObjectRolesListOK() *InventoriesInventoriesObjectRolesListOK {
	return &InventoriesInventoriesObjectRolesListOK{}
}

/*
InventoriesInventoriesObjectRolesListOK describes a response with status code 200, with default header values.

InventoriesInventoriesObjectRolesListOK inventories inventories object roles list o k
*/
type InventoriesInventoriesObjectRolesListOK struct {
}

// IsSuccess returns true when this inventories inventories object roles list o k response has a 2xx status code
func (o *InventoriesInventoriesObjectRolesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventories inventories object roles list o k response has a 3xx status code
func (o *InventoriesInventoriesObjectRolesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventories inventories object roles list o k response has a 4xx status code
func (o *InventoriesInventoriesObjectRolesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventories inventories object roles list o k response has a 5xx status code
func (o *InventoriesInventoriesObjectRolesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this inventories inventories object roles list o k response a status code equal to that given
func (o *InventoriesInventoriesObjectRolesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *InventoriesInventoriesObjectRolesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventories/{id}/object_roles/][%d] inventoriesInventoriesObjectRolesListOK ", 200)
}

func (o *InventoriesInventoriesObjectRolesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/inventories/{id}/object_roles/][%d] inventoriesInventoriesObjectRolesListOK ", 200)
}

func (o *InventoriesInventoriesObjectRolesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}