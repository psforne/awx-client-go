// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CredentialsCredentialsInputSourcesCreateReader is a Reader for the CredentialsCredentialsInputSourcesCreate structure.
type CredentialsCredentialsInputSourcesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialsCredentialsInputSourcesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCredentialsCredentialsInputSourcesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCredentialsCredentialsInputSourcesCreateCreated creates a CredentialsCredentialsInputSourcesCreateCreated with default headers values
func NewCredentialsCredentialsInputSourcesCreateCreated() *CredentialsCredentialsInputSourcesCreateCreated {
	return &CredentialsCredentialsInputSourcesCreateCreated{}
}

/*
CredentialsCredentialsInputSourcesCreateCreated describes a response with status code 201, with default header values.

CredentialsCredentialsInputSourcesCreateCreated credentials credentials input sources create created
*/
type CredentialsCredentialsInputSourcesCreateCreated struct {
}

// IsSuccess returns true when this credentials credentials input sources create created response has a 2xx status code
func (o *CredentialsCredentialsInputSourcesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this credentials credentials input sources create created response has a 3xx status code
func (o *CredentialsCredentialsInputSourcesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this credentials credentials input sources create created response has a 4xx status code
func (o *CredentialsCredentialsInputSourcesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this credentials credentials input sources create created response has a 5xx status code
func (o *CredentialsCredentialsInputSourcesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this credentials credentials input sources create created response a status code equal to that given
func (o *CredentialsCredentialsInputSourcesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CredentialsCredentialsInputSourcesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/credentials/{id}/input_sources/][%d] credentialsCredentialsInputSourcesCreateCreated ", 201)
}

func (o *CredentialsCredentialsInputSourcesCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/credentials/{id}/input_sources/][%d] credentialsCredentialsInputSourcesCreateCreated ", 201)
}

func (o *CredentialsCredentialsInputSourcesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
