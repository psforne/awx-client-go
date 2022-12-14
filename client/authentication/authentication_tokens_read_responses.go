// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AuthenticationTokensReadReader is a Reader for the AuthenticationTokensRead structure.
type AuthenticationTokensReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthenticationTokensReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthenticationTokensReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAuthenticationTokensReadOK creates a AuthenticationTokensReadOK with default headers values
func NewAuthenticationTokensReadOK() *AuthenticationTokensReadOK {
	return &AuthenticationTokensReadOK{}
}

/*
AuthenticationTokensReadOK describes a response with status code 200, with default header values.

AuthenticationTokensReadOK authentication tokens read o k
*/
type AuthenticationTokensReadOK struct {
}

// IsSuccess returns true when this authentication tokens read o k response has a 2xx status code
func (o *AuthenticationTokensReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this authentication tokens read o k response has a 3xx status code
func (o *AuthenticationTokensReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authentication tokens read o k response has a 4xx status code
func (o *AuthenticationTokensReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this authentication tokens read o k response has a 5xx status code
func (o *AuthenticationTokensReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this authentication tokens read o k response a status code equal to that given
func (o *AuthenticationTokensReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *AuthenticationTokensReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/tokens/{id}/][%d] authenticationTokensReadOK ", 200)
}

func (o *AuthenticationTokensReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/tokens/{id}/][%d] authenticationTokensReadOK ", 200)
}

func (o *AuthenticationTokensReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
