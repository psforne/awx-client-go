// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// OrganizationsOrganizationsCredentialsCreateReader is a Reader for the OrganizationsOrganizationsCredentialsCreate structure.
type OrganizationsOrganizationsCredentialsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationsCredentialsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewOrganizationsOrganizationsCredentialsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsOrganizationsCredentialsCreateCreated creates a OrganizationsOrganizationsCredentialsCreateCreated with default headers values
func NewOrganizationsOrganizationsCredentialsCreateCreated() *OrganizationsOrganizationsCredentialsCreateCreated {
	return &OrganizationsOrganizationsCredentialsCreateCreated{}
}

/*
OrganizationsOrganizationsCredentialsCreateCreated describes a response with status code 201, with default header values.

OrganizationsOrganizationsCredentialsCreateCreated organizations organizations credentials create created
*/
type OrganizationsOrganizationsCredentialsCreateCreated struct {
}

// IsSuccess returns true when this organizations organizations credentials create created response has a 2xx status code
func (o *OrganizationsOrganizationsCredentialsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations organizations credentials create created response has a 3xx status code
func (o *OrganizationsOrganizationsCredentialsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations organizations credentials create created response has a 4xx status code
func (o *OrganizationsOrganizationsCredentialsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations organizations credentials create created response has a 5xx status code
func (o *OrganizationsOrganizationsCredentialsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations organizations credentials create created response a status code equal to that given
func (o *OrganizationsOrganizationsCredentialsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *OrganizationsOrganizationsCredentialsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/organizations/{id}/credentials/][%d] organizationsOrganizationsCredentialsCreateCreated ", 201)
}

func (o *OrganizationsOrganizationsCredentialsCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/organizations/{id}/credentials/][%d] organizationsOrganizationsCredentialsCreateCreated ", 201)
}

func (o *OrganizationsOrganizationsCredentialsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
OrganizationsOrganizationsCredentialsCreateBody organizations organizations credentials create body
swagger:model OrganizationsOrganizationsCredentialsCreateBody
*/
type OrganizationsOrganizationsCredentialsCreateBody struct {

	// Specify the type of credential you want to create. Refer to the documentation for details on each type.
	// Required: true
	CredentialType *int64 `json:"credential_type"`

	// description
	Description string `json:"description,omitempty"`

	// Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax.
	Inputs interface{} `json:"inputs,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// Inherit permissions from organization roles. If provided on creation, do not give either user or team.
	Organization int64 `json:"organization,omitempty"`
}

// Validate validates this organizations organizations credentials create body
func (o *OrganizationsOrganizationsCredentialsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCredentialType(formats); err != nil {
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

func (o *OrganizationsOrganizationsCredentialsCreateBody) validateCredentialType(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"credential_type", "body", o.CredentialType); err != nil {
		return err
	}

	return nil
}

func (o *OrganizationsOrganizationsCredentialsCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this organizations organizations credentials create body based on context it is used
func (o *OrganizationsOrganizationsCredentialsCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *OrganizationsOrganizationsCredentialsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OrganizationsOrganizationsCredentialsCreateBody) UnmarshalBinary(b []byte) error {
	var res OrganizationsOrganizationsCredentialsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}