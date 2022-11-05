// Code generated by go-swagger; DO NOT EDIT.

package credential_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CredentialTypesCredentialTypesActivityStreamListReader is a Reader for the CredentialTypesCredentialTypesActivityStreamList structure.
type CredentialTypesCredentialTypesActivityStreamListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialTypesCredentialTypesActivityStreamListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCredentialTypesCredentialTypesActivityStreamListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCredentialTypesCredentialTypesActivityStreamListOK creates a CredentialTypesCredentialTypesActivityStreamListOK with default headers values
func NewCredentialTypesCredentialTypesActivityStreamListOK() *CredentialTypesCredentialTypesActivityStreamListOK {
	return &CredentialTypesCredentialTypesActivityStreamListOK{}
}

/*
CredentialTypesCredentialTypesActivityStreamListOK describes a response with status code 200, with default header values.

CredentialTypesCredentialTypesActivityStreamListOK credential types credential types activity stream list o k
*/
type CredentialTypesCredentialTypesActivityStreamListOK struct {
}

// IsSuccess returns true when this credential types credential types activity stream list o k response has a 2xx status code
func (o *CredentialTypesCredentialTypesActivityStreamListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this credential types credential types activity stream list o k response has a 3xx status code
func (o *CredentialTypesCredentialTypesActivityStreamListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this credential types credential types activity stream list o k response has a 4xx status code
func (o *CredentialTypesCredentialTypesActivityStreamListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this credential types credential types activity stream list o k response has a 5xx status code
func (o *CredentialTypesCredentialTypesActivityStreamListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this credential types credential types activity stream list o k response a status code equal to that given
func (o *CredentialTypesCredentialTypesActivityStreamListOK) IsCode(code int) bool {
	return code == 200
}

func (o *CredentialTypesCredentialTypesActivityStreamListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/credential_types/{id}/activity_stream/][%d] credentialTypesCredentialTypesActivityStreamListOK ", 200)
}

func (o *CredentialTypesCredentialTypesActivityStreamListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/credential_types/{id}/activity_stream/][%d] credentialTypesCredentialTypesActivityStreamListOK ", 200)
}

func (o *CredentialTypesCredentialTypesActivityStreamListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}