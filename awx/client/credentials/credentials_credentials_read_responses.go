// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CredentialsCredentialsReadReader is a Reader for the CredentialsCredentialsRead structure.
type CredentialsCredentialsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialsCredentialsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCredentialsCredentialsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCredentialsCredentialsReadOK creates a CredentialsCredentialsReadOK with default headers values
func NewCredentialsCredentialsReadOK() *CredentialsCredentialsReadOK {
	return &CredentialsCredentialsReadOK{}
}

/*
CredentialsCredentialsReadOK describes a response with status code 200, with default header values.

CredentialsCredentialsReadOK credentials credentials read o k
*/
type CredentialsCredentialsReadOK struct {
}

// IsSuccess returns true when this credentials credentials read o k response has a 2xx status code
func (o *CredentialsCredentialsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this credentials credentials read o k response has a 3xx status code
func (o *CredentialsCredentialsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this credentials credentials read o k response has a 4xx status code
func (o *CredentialsCredentialsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this credentials credentials read o k response has a 5xx status code
func (o *CredentialsCredentialsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this credentials credentials read o k response a status code equal to that given
func (o *CredentialsCredentialsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *CredentialsCredentialsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/credentials/{id}/][%d] credentialsCredentialsReadOK ", 200)
}

func (o *CredentialsCredentialsReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/credentials/{id}/][%d] credentialsCredentialsReadOK ", 200)
}

func (o *CredentialsCredentialsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
