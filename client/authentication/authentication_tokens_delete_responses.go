// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AuthenticationTokensDeleteReader is a Reader for the AuthenticationTokensDelete structure.
type AuthenticationTokensDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthenticationTokensDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAuthenticationTokensDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAuthenticationTokensDeleteNoContent creates a AuthenticationTokensDeleteNoContent with default headers values
func NewAuthenticationTokensDeleteNoContent() *AuthenticationTokensDeleteNoContent {
	return &AuthenticationTokensDeleteNoContent{}
}

/*
AuthenticationTokensDeleteNoContent describes a response with status code 204, with default header values.

AuthenticationTokensDeleteNoContent authentication tokens delete no content
*/
type AuthenticationTokensDeleteNoContent struct {
}

// IsSuccess returns true when this authentication tokens delete no content response has a 2xx status code
func (o *AuthenticationTokensDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this authentication tokens delete no content response has a 3xx status code
func (o *AuthenticationTokensDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authentication tokens delete no content response has a 4xx status code
func (o *AuthenticationTokensDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this authentication tokens delete no content response has a 5xx status code
func (o *AuthenticationTokensDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this authentication tokens delete no content response a status code equal to that given
func (o *AuthenticationTokensDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *AuthenticationTokensDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/tokens/{id}/][%d] authenticationTokensDeleteNoContent ", 204)
}

func (o *AuthenticationTokensDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/tokens/{id}/][%d] authenticationTokensDeleteNoContent ", 204)
}

func (o *AuthenticationTokensDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
