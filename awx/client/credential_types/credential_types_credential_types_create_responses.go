// Code generated by go-swagger; DO NOT EDIT.

package credential_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CredentialTypesCredentialTypesCreateReader is a Reader for the CredentialTypesCredentialTypesCreate structure.
type CredentialTypesCredentialTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialTypesCredentialTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCredentialTypesCredentialTypesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCredentialTypesCredentialTypesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCredentialTypesCredentialTypesCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCredentialTypesCredentialTypesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCredentialTypesCredentialTypesCreateCreated creates a CredentialTypesCredentialTypesCreateCreated with default headers values
func NewCredentialTypesCredentialTypesCreateCreated() *CredentialTypesCredentialTypesCreateCreated {
	return &CredentialTypesCredentialTypesCreateCreated{}
}

/*
CredentialTypesCredentialTypesCreateCreated describes a response with status code 201, with default header values.

CredentialTypesCredentialTypesCreateCreated credential types credential types create created
*/
type CredentialTypesCredentialTypesCreateCreated struct {
}

// IsSuccess returns true when this credential types credential types create created response has a 2xx status code
func (o *CredentialTypesCredentialTypesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this credential types credential types create created response has a 3xx status code
func (o *CredentialTypesCredentialTypesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this credential types credential types create created response has a 4xx status code
func (o *CredentialTypesCredentialTypesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this credential types credential types create created response has a 5xx status code
func (o *CredentialTypesCredentialTypesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this credential types credential types create created response a status code equal to that given
func (o *CredentialTypesCredentialTypesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CredentialTypesCredentialTypesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/credential_types/][%d] credentialTypesCredentialTypesCreateCreated ", 201)
}

func (o *CredentialTypesCredentialTypesCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/credential_types/][%d] credentialTypesCredentialTypesCreateCreated ", 201)
}

func (o *CredentialTypesCredentialTypesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCredentialTypesCredentialTypesCreateBadRequest creates a CredentialTypesCredentialTypesCreateBadRequest with default headers values
func NewCredentialTypesCredentialTypesCreateBadRequest() *CredentialTypesCredentialTypesCreateBadRequest {
	return &CredentialTypesCredentialTypesCreateBadRequest{}
}

/*
CredentialTypesCredentialTypesCreateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CredentialTypesCredentialTypesCreateBadRequest struct {
}

// IsSuccess returns true when this credential types credential types create bad request response has a 2xx status code
func (o *CredentialTypesCredentialTypesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this credential types credential types create bad request response has a 3xx status code
func (o *CredentialTypesCredentialTypesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this credential types credential types create bad request response has a 4xx status code
func (o *CredentialTypesCredentialTypesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this credential types credential types create bad request response has a 5xx status code
func (o *CredentialTypesCredentialTypesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this credential types credential types create bad request response a status code equal to that given
func (o *CredentialTypesCredentialTypesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CredentialTypesCredentialTypesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/credential_types/][%d] credentialTypesCredentialTypesCreateBadRequest ", 400)
}

func (o *CredentialTypesCredentialTypesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/credential_types/][%d] credentialTypesCredentialTypesCreateBadRequest ", 400)
}

func (o *CredentialTypesCredentialTypesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCredentialTypesCredentialTypesCreateUnauthorized creates a CredentialTypesCredentialTypesCreateUnauthorized with default headers values
func NewCredentialTypesCredentialTypesCreateUnauthorized() *CredentialTypesCredentialTypesCreateUnauthorized {
	return &CredentialTypesCredentialTypesCreateUnauthorized{}
}

/*
CredentialTypesCredentialTypesCreateUnauthorized describes a response with status code 401, with default header values.

CredentialTypesCredentialTypesCreateUnauthorized credential types credential types create unauthorized
*/
type CredentialTypesCredentialTypesCreateUnauthorized struct {
}

// IsSuccess returns true when this credential types credential types create unauthorized response has a 2xx status code
func (o *CredentialTypesCredentialTypesCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this credential types credential types create unauthorized response has a 3xx status code
func (o *CredentialTypesCredentialTypesCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this credential types credential types create unauthorized response has a 4xx status code
func (o *CredentialTypesCredentialTypesCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this credential types credential types create unauthorized response has a 5xx status code
func (o *CredentialTypesCredentialTypesCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this credential types credential types create unauthorized response a status code equal to that given
func (o *CredentialTypesCredentialTypesCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CredentialTypesCredentialTypesCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/credential_types/][%d] credentialTypesCredentialTypesCreateUnauthorized ", 401)
}

func (o *CredentialTypesCredentialTypesCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/credential_types/][%d] credentialTypesCredentialTypesCreateUnauthorized ", 401)
}

func (o *CredentialTypesCredentialTypesCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCredentialTypesCredentialTypesCreateForbidden creates a CredentialTypesCredentialTypesCreateForbidden with default headers values
func NewCredentialTypesCredentialTypesCreateForbidden() *CredentialTypesCredentialTypesCreateForbidden {
	return &CredentialTypesCredentialTypesCreateForbidden{}
}

/*
CredentialTypesCredentialTypesCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type CredentialTypesCredentialTypesCreateForbidden struct {
}

// IsSuccess returns true when this credential types credential types create forbidden response has a 2xx status code
func (o *CredentialTypesCredentialTypesCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this credential types credential types create forbidden response has a 3xx status code
func (o *CredentialTypesCredentialTypesCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this credential types credential types create forbidden response has a 4xx status code
func (o *CredentialTypesCredentialTypesCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this credential types credential types create forbidden response has a 5xx status code
func (o *CredentialTypesCredentialTypesCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this credential types credential types create forbidden response a status code equal to that given
func (o *CredentialTypesCredentialTypesCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CredentialTypesCredentialTypesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/credential_types/][%d] credentialTypesCredentialTypesCreateForbidden ", 403)
}

func (o *CredentialTypesCredentialTypesCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/credential_types/][%d] credentialTypesCredentialTypesCreateForbidden ", 403)
}

func (o *CredentialTypesCredentialTypesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
