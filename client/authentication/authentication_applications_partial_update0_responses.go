// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AuthenticationApplicationsPartialUpdate0Reader is a Reader for the AuthenticationApplicationsPartialUpdate0 structure.
type AuthenticationApplicationsPartialUpdate0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthenticationApplicationsPartialUpdate0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthenticationApplicationsPartialUpdate0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAuthenticationApplicationsPartialUpdate0OK creates a AuthenticationApplicationsPartialUpdate0OK with default headers values
func NewAuthenticationApplicationsPartialUpdate0OK() *AuthenticationApplicationsPartialUpdate0OK {
	return &AuthenticationApplicationsPartialUpdate0OK{}
}

/*
AuthenticationApplicationsPartialUpdate0OK describes a response with status code 200, with default header values.

AuthenticationApplicationsPartialUpdate0OK authentication applications partial update0 o k
*/
type AuthenticationApplicationsPartialUpdate0OK struct {
}

// IsSuccess returns true when this authentication applications partial update0 o k response has a 2xx status code
func (o *AuthenticationApplicationsPartialUpdate0OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this authentication applications partial update0 o k response has a 3xx status code
func (o *AuthenticationApplicationsPartialUpdate0OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authentication applications partial update0 o k response has a 4xx status code
func (o *AuthenticationApplicationsPartialUpdate0OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this authentication applications partial update0 o k response has a 5xx status code
func (o *AuthenticationApplicationsPartialUpdate0OK) IsServerError() bool {
	return false
}

// IsCode returns true when this authentication applications partial update0 o k response a status code equal to that given
func (o *AuthenticationApplicationsPartialUpdate0OK) IsCode(code int) bool {
	return code == 200
}

func (o *AuthenticationApplicationsPartialUpdate0OK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/applications/{id}/][%d] authenticationApplicationsPartialUpdate0OK ", 200)
}

func (o *AuthenticationApplicationsPartialUpdate0OK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/applications/{id}/][%d] authenticationApplicationsPartialUpdate0OK ", 200)
}

func (o *AuthenticationApplicationsPartialUpdate0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
