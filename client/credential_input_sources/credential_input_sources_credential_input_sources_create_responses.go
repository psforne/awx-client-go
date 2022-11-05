// Code generated by go-swagger; DO NOT EDIT.

package credential_input_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CredentialInputSourcesCredentialInputSourcesCreateReader is a Reader for the CredentialInputSourcesCredentialInputSourcesCreate structure.
type CredentialInputSourcesCredentialInputSourcesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialInputSourcesCredentialInputSourcesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCredentialInputSourcesCredentialInputSourcesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCredentialInputSourcesCredentialInputSourcesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCredentialInputSourcesCredentialInputSourcesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCredentialInputSourcesCredentialInputSourcesCreateCreated creates a CredentialInputSourcesCredentialInputSourcesCreateCreated with default headers values
func NewCredentialInputSourcesCredentialInputSourcesCreateCreated() *CredentialInputSourcesCredentialInputSourcesCreateCreated {
	return &CredentialInputSourcesCredentialInputSourcesCreateCreated{}
}

/*
CredentialInputSourcesCredentialInputSourcesCreateCreated describes a response with status code 201, with default header values.

CredentialInputSourcesCredentialInputSourcesCreateCreated credential input sources credential input sources create created
*/
type CredentialInputSourcesCredentialInputSourcesCreateCreated struct {
}

// IsSuccess returns true when this credential input sources credential input sources create created response has a 2xx status code
func (o *CredentialInputSourcesCredentialInputSourcesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this credential input sources credential input sources create created response has a 3xx status code
func (o *CredentialInputSourcesCredentialInputSourcesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this credential input sources credential input sources create created response has a 4xx status code
func (o *CredentialInputSourcesCredentialInputSourcesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this credential input sources credential input sources create created response has a 5xx status code
func (o *CredentialInputSourcesCredentialInputSourcesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this credential input sources credential input sources create created response a status code equal to that given
func (o *CredentialInputSourcesCredentialInputSourcesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CredentialInputSourcesCredentialInputSourcesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/credential_input_sources/][%d] credentialInputSourcesCredentialInputSourcesCreateCreated ", 201)
}

func (o *CredentialInputSourcesCredentialInputSourcesCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/credential_input_sources/][%d] credentialInputSourcesCredentialInputSourcesCreateCreated ", 201)
}

func (o *CredentialInputSourcesCredentialInputSourcesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCredentialInputSourcesCredentialInputSourcesCreateBadRequest creates a CredentialInputSourcesCredentialInputSourcesCreateBadRequest with default headers values
func NewCredentialInputSourcesCredentialInputSourcesCreateBadRequest() *CredentialInputSourcesCredentialInputSourcesCreateBadRequest {
	return &CredentialInputSourcesCredentialInputSourcesCreateBadRequest{}
}

/*
CredentialInputSourcesCredentialInputSourcesCreateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CredentialInputSourcesCredentialInputSourcesCreateBadRequest struct {
}

// IsSuccess returns true when this credential input sources credential input sources create bad request response has a 2xx status code
func (o *CredentialInputSourcesCredentialInputSourcesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this credential input sources credential input sources create bad request response has a 3xx status code
func (o *CredentialInputSourcesCredentialInputSourcesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this credential input sources credential input sources create bad request response has a 4xx status code
func (o *CredentialInputSourcesCredentialInputSourcesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this credential input sources credential input sources create bad request response has a 5xx status code
func (o *CredentialInputSourcesCredentialInputSourcesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this credential input sources credential input sources create bad request response a status code equal to that given
func (o *CredentialInputSourcesCredentialInputSourcesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CredentialInputSourcesCredentialInputSourcesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/credential_input_sources/][%d] credentialInputSourcesCredentialInputSourcesCreateBadRequest ", 400)
}

func (o *CredentialInputSourcesCredentialInputSourcesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/credential_input_sources/][%d] credentialInputSourcesCredentialInputSourcesCreateBadRequest ", 400)
}

func (o *CredentialInputSourcesCredentialInputSourcesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCredentialInputSourcesCredentialInputSourcesCreateForbidden creates a CredentialInputSourcesCredentialInputSourcesCreateForbidden with default headers values
func NewCredentialInputSourcesCredentialInputSourcesCreateForbidden() *CredentialInputSourcesCredentialInputSourcesCreateForbidden {
	return &CredentialInputSourcesCredentialInputSourcesCreateForbidden{}
}

/*
CredentialInputSourcesCredentialInputSourcesCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type CredentialInputSourcesCredentialInputSourcesCreateForbidden struct {
}

// IsSuccess returns true when this credential input sources credential input sources create forbidden response has a 2xx status code
func (o *CredentialInputSourcesCredentialInputSourcesCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this credential input sources credential input sources create forbidden response has a 3xx status code
func (o *CredentialInputSourcesCredentialInputSourcesCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this credential input sources credential input sources create forbidden response has a 4xx status code
func (o *CredentialInputSourcesCredentialInputSourcesCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this credential input sources credential input sources create forbidden response has a 5xx status code
func (o *CredentialInputSourcesCredentialInputSourcesCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this credential input sources credential input sources create forbidden response a status code equal to that given
func (o *CredentialInputSourcesCredentialInputSourcesCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CredentialInputSourcesCredentialInputSourcesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/credential_input_sources/][%d] credentialInputSourcesCredentialInputSourcesCreateForbidden ", 403)
}

func (o *CredentialInputSourcesCredentialInputSourcesCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/credential_input_sources/][%d] credentialInputSourcesCredentialInputSourcesCreateForbidden ", 403)
}

func (o *CredentialInputSourcesCredentialInputSourcesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}