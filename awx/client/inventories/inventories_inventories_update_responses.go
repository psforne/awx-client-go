// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InventoriesInventoriesUpdateReader is a Reader for the InventoriesInventoriesUpdate structure.
type InventoriesInventoriesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoriesInventoriesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoriesInventoriesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewInventoriesInventoriesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventoriesInventoriesUpdateOK creates a InventoriesInventoriesUpdateOK with default headers values
func NewInventoriesInventoriesUpdateOK() *InventoriesInventoriesUpdateOK {
	return &InventoriesInventoriesUpdateOK{}
}

/*
InventoriesInventoriesUpdateOK describes a response with status code 200, with default header values.

InventoriesInventoriesUpdateOK inventories inventories update o k
*/
type InventoriesInventoriesUpdateOK struct {
}

// IsSuccess returns true when this inventories inventories update o k response has a 2xx status code
func (o *InventoriesInventoriesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventories inventories update o k response has a 3xx status code
func (o *InventoriesInventoriesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventories inventories update o k response has a 4xx status code
func (o *InventoriesInventoriesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventories inventories update o k response has a 5xx status code
func (o *InventoriesInventoriesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this inventories inventories update o k response a status code equal to that given
func (o *InventoriesInventoriesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *InventoriesInventoriesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/inventories/{id}/][%d] inventoriesInventoriesUpdateOK ", 200)
}

func (o *InventoriesInventoriesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/inventories/{id}/][%d] inventoriesInventoriesUpdateOK ", 200)
}

func (o *InventoriesInventoriesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInventoriesInventoriesUpdateForbidden creates a InventoriesInventoriesUpdateForbidden with default headers values
func NewInventoriesInventoriesUpdateForbidden() *InventoriesInventoriesUpdateForbidden {
	return &InventoriesInventoriesUpdateForbidden{}
}

/*
InventoriesInventoriesUpdateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type InventoriesInventoriesUpdateForbidden struct {
}

// IsSuccess returns true when this inventories inventories update forbidden response has a 2xx status code
func (o *InventoriesInventoriesUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this inventories inventories update forbidden response has a 3xx status code
func (o *InventoriesInventoriesUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventories inventories update forbidden response has a 4xx status code
func (o *InventoriesInventoriesUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this inventories inventories update forbidden response has a 5xx status code
func (o *InventoriesInventoriesUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this inventories inventories update forbidden response a status code equal to that given
func (o *InventoriesInventoriesUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *InventoriesInventoriesUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/inventories/{id}/][%d] inventoriesInventoriesUpdateForbidden ", 403)
}

func (o *InventoriesInventoriesUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/inventories/{id}/][%d] inventoriesInventoriesUpdateForbidden ", 403)
}

func (o *InventoriesInventoriesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
