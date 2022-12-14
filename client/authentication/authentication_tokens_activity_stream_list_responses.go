// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AuthenticationTokensActivityStreamListReader is a Reader for the AuthenticationTokensActivityStreamList structure.
type AuthenticationTokensActivityStreamListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthenticationTokensActivityStreamListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthenticationTokensActivityStreamListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAuthenticationTokensActivityStreamListOK creates a AuthenticationTokensActivityStreamListOK with default headers values
func NewAuthenticationTokensActivityStreamListOK() *AuthenticationTokensActivityStreamListOK {
	return &AuthenticationTokensActivityStreamListOK{}
}

/*
AuthenticationTokensActivityStreamListOK describes a response with status code 200, with default header values.

AuthenticationTokensActivityStreamListOK authentication tokens activity stream list o k
*/
type AuthenticationTokensActivityStreamListOK struct {
}

// IsSuccess returns true when this authentication tokens activity stream list o k response has a 2xx status code
func (o *AuthenticationTokensActivityStreamListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this authentication tokens activity stream list o k response has a 3xx status code
func (o *AuthenticationTokensActivityStreamListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authentication tokens activity stream list o k response has a 4xx status code
func (o *AuthenticationTokensActivityStreamListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this authentication tokens activity stream list o k response has a 5xx status code
func (o *AuthenticationTokensActivityStreamListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this authentication tokens activity stream list o k response a status code equal to that given
func (o *AuthenticationTokensActivityStreamListOK) IsCode(code int) bool {
	return code == 200
}

func (o *AuthenticationTokensActivityStreamListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/tokens/{id}/activity_stream/][%d] authenticationTokensActivityStreamListOK ", 200)
}

func (o *AuthenticationTokensActivityStreamListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/tokens/{id}/activity_stream/][%d] authenticationTokensActivityStreamListOK ", 200)
}

func (o *AuthenticationTokensActivityStreamListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
