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

// InventoriesInventoriesRootGroupsCreateReader is a Reader for the InventoriesInventoriesRootGroupsCreate structure.
type InventoriesInventoriesRootGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoriesInventoriesRootGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInventoriesInventoriesRootGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventoriesInventoriesRootGroupsCreateCreated creates a InventoriesInventoriesRootGroupsCreateCreated with default headers values
func NewInventoriesInventoriesRootGroupsCreateCreated() *InventoriesInventoriesRootGroupsCreateCreated {
	return &InventoriesInventoriesRootGroupsCreateCreated{}
}

/*
InventoriesInventoriesRootGroupsCreateCreated describes a response with status code 201, with default header values.

InventoriesInventoriesRootGroupsCreateCreated inventories inventories root groups create created
*/
type InventoriesInventoriesRootGroupsCreateCreated struct {
}

// IsSuccess returns true when this inventories inventories root groups create created response has a 2xx status code
func (o *InventoriesInventoriesRootGroupsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventories inventories root groups create created response has a 3xx status code
func (o *InventoriesInventoriesRootGroupsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventories inventories root groups create created response has a 4xx status code
func (o *InventoriesInventoriesRootGroupsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventories inventories root groups create created response has a 5xx status code
func (o *InventoriesInventoriesRootGroupsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this inventories inventories root groups create created response a status code equal to that given
func (o *InventoriesInventoriesRootGroupsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *InventoriesInventoriesRootGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/inventories/{id}/root_groups/][%d] inventoriesInventoriesRootGroupsCreateCreated ", 201)
}

func (o *InventoriesInventoriesRootGroupsCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/inventories/{id}/root_groups/][%d] inventoriesInventoriesRootGroupsCreateCreated ", 201)
}

func (o *InventoriesInventoriesRootGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
InventoriesInventoriesRootGroupsCreateBody inventories inventories root groups create body
swagger:model InventoriesInventoriesRootGroupsCreateBody
*/
type InventoriesInventoriesRootGroupsCreateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// inventory
	// Required: true
	Inventory *int64 `json:"inventory"`

	// name
	// Required: true
	Name *string `json:"name"`

	// Group variables in JSON or YAML format.
	Variables string `json:"variables,omitempty"`
}

// Validate validates this inventories inventories root groups create body
func (o *InventoriesInventoriesRootGroupsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInventory(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InventoriesInventoriesRootGroupsCreateBody) validateInventory(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"inventory", "body", o.Inventory); err != nil {
		return err
	}

	return nil
}

func (o *InventoriesInventoriesRootGroupsCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this inventories inventories root groups create body based on context it is used
func (o *InventoriesInventoriesRootGroupsCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InventoriesInventoriesRootGroupsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InventoriesInventoriesRootGroupsCreateBody) UnmarshalBinary(b []byte) error {
	var res InventoriesInventoriesRootGroupsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
