// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoriesInventoriesAdHocCommandsListReader is a Reader for the InventoriesInventoriesAdHocCommandsList structure.
type InventoriesInventoriesAdHocCommandsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoriesInventoriesAdHocCommandsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoriesInventoriesAdHocCommandsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewInventoriesInventoriesAdHocCommandsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventoriesInventoriesAdHocCommandsListOK creates a InventoriesInventoriesAdHocCommandsListOK with default headers values
func NewInventoriesInventoriesAdHocCommandsListOK() *InventoriesInventoriesAdHocCommandsListOK {
	return &InventoriesInventoriesAdHocCommandsListOK{}
}

/*
InventoriesInventoriesAdHocCommandsListOK describes a response with status code 200, with default header values.

InventoriesInventoriesAdHocCommandsListOK inventories inventories ad hoc commands list o k
*/
type InventoriesInventoriesAdHocCommandsListOK struct {
}

// IsSuccess returns true when this inventories inventories ad hoc commands list o k response has a 2xx status code
func (o *InventoriesInventoriesAdHocCommandsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventories inventories ad hoc commands list o k response has a 3xx status code
func (o *InventoriesInventoriesAdHocCommandsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventories inventories ad hoc commands list o k response has a 4xx status code
func (o *InventoriesInventoriesAdHocCommandsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventories inventories ad hoc commands list o k response has a 5xx status code
func (o *InventoriesInventoriesAdHocCommandsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this inventories inventories ad hoc commands list o k response a status code equal to that given
func (o *InventoriesInventoriesAdHocCommandsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *InventoriesInventoriesAdHocCommandsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventories/{id}/ad_hoc_commands/][%d] inventoriesInventoriesAdHocCommandsListOK ", 200)
}

func (o *InventoriesInventoriesAdHocCommandsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/inventories/{id}/ad_hoc_commands/][%d] inventoriesInventoriesAdHocCommandsListOK ", 200)
}

func (o *InventoriesInventoriesAdHocCommandsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInventoriesInventoriesAdHocCommandsListForbidden creates a InventoriesInventoriesAdHocCommandsListForbidden with default headers values
func NewInventoriesInventoriesAdHocCommandsListForbidden() *InventoriesInventoriesAdHocCommandsListForbidden {
	return &InventoriesInventoriesAdHocCommandsListForbidden{}
}

/*
InventoriesInventoriesAdHocCommandsListForbidden describes a response with status code 403, with default header values.

forbidden
*/
type InventoriesInventoriesAdHocCommandsListForbidden struct {
}

// IsSuccess returns true when this inventories inventories ad hoc commands list forbidden response has a 2xx status code
func (o *InventoriesInventoriesAdHocCommandsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this inventories inventories ad hoc commands list forbidden response has a 3xx status code
func (o *InventoriesInventoriesAdHocCommandsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventories inventories ad hoc commands list forbidden response has a 4xx status code
func (o *InventoriesInventoriesAdHocCommandsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this inventories inventories ad hoc commands list forbidden response has a 5xx status code
func (o *InventoriesInventoriesAdHocCommandsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this inventories inventories ad hoc commands list forbidden response a status code equal to that given
func (o *InventoriesInventoriesAdHocCommandsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *InventoriesInventoriesAdHocCommandsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/inventories/{id}/ad_hoc_commands/][%d] inventoriesInventoriesAdHocCommandsListForbidden ", 403)
}

func (o *InventoriesInventoriesAdHocCommandsListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/inventories/{id}/ad_hoc_commands/][%d] inventoriesInventoriesAdHocCommandsListForbidden ", 403)
}

func (o *InventoriesInventoriesAdHocCommandsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
