// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoriesInventoriesVariableDataReadReader is a Reader for the InventoriesInventoriesVariableDataRead structure.
type InventoriesInventoriesVariableDataReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoriesInventoriesVariableDataReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoriesInventoriesVariableDataReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventoriesInventoriesVariableDataReadOK creates a InventoriesInventoriesVariableDataReadOK with default headers values
func NewInventoriesInventoriesVariableDataReadOK() *InventoriesInventoriesVariableDataReadOK {
	return &InventoriesInventoriesVariableDataReadOK{}
}

/*
InventoriesInventoriesVariableDataReadOK describes a response with status code 200, with default header values.

InventoriesInventoriesVariableDataReadOK inventories inventories variable data read o k
*/
type InventoriesInventoriesVariableDataReadOK struct {
}

// IsSuccess returns true when this inventories inventories variable data read o k response has a 2xx status code
func (o *InventoriesInventoriesVariableDataReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventories inventories variable data read o k response has a 3xx status code
func (o *InventoriesInventoriesVariableDataReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventories inventories variable data read o k response has a 4xx status code
func (o *InventoriesInventoriesVariableDataReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventories inventories variable data read o k response has a 5xx status code
func (o *InventoriesInventoriesVariableDataReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this inventories inventories variable data read o k response a status code equal to that given
func (o *InventoriesInventoriesVariableDataReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *InventoriesInventoriesVariableDataReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventories/{id}/variable_data/][%d] inventoriesInventoriesVariableDataReadOK ", 200)
}

func (o *InventoriesInventoriesVariableDataReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/inventories/{id}/variable_data/][%d] inventoriesInventoriesVariableDataReadOK ", 200)
}

func (o *InventoriesInventoriesVariableDataReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
