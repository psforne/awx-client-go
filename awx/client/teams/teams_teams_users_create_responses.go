// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// TeamsTeamsUsersCreateReader is a Reader for the TeamsTeamsUsersCreate structure.
type TeamsTeamsUsersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamsTeamsUsersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTeamsTeamsUsersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTeamsTeamsUsersCreateCreated creates a TeamsTeamsUsersCreateCreated with default headers values
func NewTeamsTeamsUsersCreateCreated() *TeamsTeamsUsersCreateCreated {
	return &TeamsTeamsUsersCreateCreated{}
}

/*
TeamsTeamsUsersCreateCreated describes a response with status code 201, with default header values.

TeamsTeamsUsersCreateCreated teams teams users create created
*/
type TeamsTeamsUsersCreateCreated struct {
}

// IsSuccess returns true when this teams teams users create created response has a 2xx status code
func (o *TeamsTeamsUsersCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this teams teams users create created response has a 3xx status code
func (o *TeamsTeamsUsersCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this teams teams users create created response has a 4xx status code
func (o *TeamsTeamsUsersCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this teams teams users create created response has a 5xx status code
func (o *TeamsTeamsUsersCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this teams teams users create created response a status code equal to that given
func (o *TeamsTeamsUsersCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *TeamsTeamsUsersCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams/{id}/users/][%d] teamsTeamsUsersCreateCreated ", 201)
}

func (o *TeamsTeamsUsersCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/teams/{id}/users/][%d] teamsTeamsUsersCreateCreated ", 201)
}

func (o *TeamsTeamsUsersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
TeamsTeamsUsersCreateBody teams teams users create body
swagger:model TeamsTeamsUsersCreateBody
*/
type TeamsTeamsUsersCreateBody struct {

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

// Validate validates this teams teams users create body
func (o *TeamsTeamsUsersCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TeamsTeamsUsersCreateBody) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this teams teams users create body based on context it is used
func (o *TeamsTeamsUsersCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TeamsTeamsUsersCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TeamsTeamsUsersCreateBody) UnmarshalBinary(b []byte) error {
	var res TeamsTeamsUsersCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
