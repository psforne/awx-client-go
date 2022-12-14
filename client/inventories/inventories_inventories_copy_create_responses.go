// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InventoriesInventoriesCopyCreateReader is a Reader for the InventoriesInventoriesCopyCreate structure.
type InventoriesInventoriesCopyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoriesInventoriesCopyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInventoriesInventoriesCopyCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventoriesInventoriesCopyCreateCreated creates a InventoriesInventoriesCopyCreateCreated with default headers values
func NewInventoriesInventoriesCopyCreateCreated() *InventoriesInventoriesCopyCreateCreated {
	return &InventoriesInventoriesCopyCreateCreated{}
}

/*
InventoriesInventoriesCopyCreateCreated describes a response with status code 201, with default header values.

InventoriesInventoriesCopyCreateCreated inventories inventories copy create created
*/
type InventoriesInventoriesCopyCreateCreated struct {
}

// IsSuccess returns true when this inventories inventories copy create created response has a 2xx status code
func (o *InventoriesInventoriesCopyCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventories inventories copy create created response has a 3xx status code
func (o *InventoriesInventoriesCopyCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventories inventories copy create created response has a 4xx status code
func (o *InventoriesInventoriesCopyCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventories inventories copy create created response has a 5xx status code
func (o *InventoriesInventoriesCopyCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this inventories inventories copy create created response a status code equal to that given
func (o *InventoriesInventoriesCopyCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *InventoriesInventoriesCopyCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/inventories/{id}/copy/][%d] inventoriesInventoriesCopyCreateCreated ", 201)
}

func (o *InventoriesInventoriesCopyCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/inventories/{id}/copy/][%d] inventoriesInventoriesCopyCreateCreated ", 201)
}

func (o *InventoriesInventoriesCopyCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
InventoriesInventoriesCopyCreateBody inventories inventories copy create body
swagger:model InventoriesInventoriesCopyCreateBody
*/
type InventoriesInventoriesCopyCreateBody struct {

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this inventories inventories copy create body
func (o *InventoriesInventoriesCopyCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InventoriesInventoriesCopyCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this inventories inventories copy create body based on context it is used
func (o *InventoriesInventoriesCopyCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InventoriesInventoriesCopyCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InventoriesInventoriesCopyCreateBody) UnmarshalBinary(b []byte) error {
	var res InventoriesInventoriesCopyCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
