// Code generated by go-swagger; DO NOT EDIT.

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CredentialsCredentialsDeleteReader is a Reader for the CredentialsCredentialsDelete structure.
type CredentialsCredentialsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialsCredentialsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCredentialsCredentialsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCredentialsCredentialsDeleteNoContent creates a CredentialsCredentialsDeleteNoContent with default headers values
func NewCredentialsCredentialsDeleteNoContent() *CredentialsCredentialsDeleteNoContent {
	return &CredentialsCredentialsDeleteNoContent{}
}

/*
CredentialsCredentialsDeleteNoContent describes a response with status code 204, with default header values.

CredentialsCredentialsDeleteNoContent credentials credentials delete no content
*/
type CredentialsCredentialsDeleteNoContent struct {
}

// IsSuccess returns true when this credentials credentials delete no content response has a 2xx status code
func (o *CredentialsCredentialsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this credentials credentials delete no content response has a 3xx status code
func (o *CredentialsCredentialsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this credentials credentials delete no content response has a 4xx status code
func (o *CredentialsCredentialsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this credentials credentials delete no content response has a 5xx status code
func (o *CredentialsCredentialsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this credentials credentials delete no content response a status code equal to that given
func (o *CredentialsCredentialsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *CredentialsCredentialsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/credentials/{id}/][%d] credentialsCredentialsDeleteNoContent ", 204)
}

func (o *CredentialsCredentialsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/credentials/{id}/][%d] credentialsCredentialsDeleteNoContent ", 204)
}

func (o *CredentialsCredentialsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
