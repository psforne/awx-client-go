// Code generated by go-swagger; DO NOT EDIT.

package users

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

// UsersUsersUpdateReader is a Reader for the UsersUsersUpdate structure.
type UsersUsersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUsersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersUsersUpdateOK creates a UsersUsersUpdateOK with default headers values
func NewUsersUsersUpdateOK() *UsersUsersUpdateOK {
	return &UsersUsersUpdateOK{}
}

/*
UsersUsersUpdateOK describes a response with status code 200, with default header values.

UsersUsersUpdateOK users users update o k
*/
type UsersUsersUpdateOK struct {
}

// IsSuccess returns true when this users users update o k response has a 2xx status code
func (o *UsersUsersUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users users update o k response has a 3xx status code
func (o *UsersUsersUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users update o k response has a 4xx status code
func (o *UsersUsersUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users users update o k response has a 5xx status code
func (o *UsersUsersUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users users update o k response a status code equal to that given
func (o *UsersUsersUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersUsersUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{id}/][%d] usersUsersUpdateOK ", 200)
}

func (o *UsersUsersUpdateOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/users/{id}/][%d] usersUsersUpdateOK ", 200)
}

func (o *UsersUsersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
UsersUsersUpdateBody users users update body
swagger:model UsersUsersUpdateBody
*/
type UsersUsersUpdateBody struct {

	// email
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// Designates that this user has all permissions without explicitly assigning them.
	IsSuperuser bool `json:"is_superuser,omitempty"`

	// is system auditor
	IsSystemAuditor bool `json:"is_system_auditor,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// Write-only field used to change the password.
	Password string `json:"password,omitempty"`

	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this users users update body
func (o *UsersUsersUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UsersUsersUpdateBody) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this users users update body based on context it is used
func (o *UsersUsersUpdateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UsersUsersUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersUsersUpdateBody) UnmarshalBinary(b []byte) error {
	var res UsersUsersUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
