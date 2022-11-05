// Code generated by go-swagger; DO NOT EDIT.

package credential_input_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CredentialInputSourcesCredentialInputSourcesPartialUpdateReader is a Reader for the CredentialInputSourcesCredentialInputSourcesPartialUpdate structure.
type CredentialInputSourcesCredentialInputSourcesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCredentialInputSourcesCredentialInputSourcesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCredentialInputSourcesCredentialInputSourcesPartialUpdateOK creates a CredentialInputSourcesCredentialInputSourcesPartialUpdateOK with default headers values
func NewCredentialInputSourcesCredentialInputSourcesPartialUpdateOK() *CredentialInputSourcesCredentialInputSourcesPartialUpdateOK {
	return &CredentialInputSourcesCredentialInputSourcesPartialUpdateOK{}
}

/*
CredentialInputSourcesCredentialInputSourcesPartialUpdateOK describes a response with status code 200, with default header values.

CredentialInputSourcesCredentialInputSourcesPartialUpdateOK credential input sources credential input sources partial update o k
*/
type CredentialInputSourcesCredentialInputSourcesPartialUpdateOK struct {
}

// IsSuccess returns true when this credential input sources credential input sources partial update o k response has a 2xx status code
func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this credential input sources credential input sources partial update o k response has a 3xx status code
func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this credential input sources credential input sources partial update o k response has a 4xx status code
func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this credential input sources credential input sources partial update o k response has a 5xx status code
func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this credential input sources credential input sources partial update o k response a status code equal to that given
func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/credential_input_sources/{id}/][%d] credentialInputSourcesCredentialInputSourcesPartialUpdateOK ", 200)
}

func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/credential_input_sources/{id}/][%d] credentialInputSourcesCredentialInputSourcesPartialUpdateOK ", 200)
}

func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden creates a CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden with default headers values
func NewCredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden() *CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden {
	return &CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden{}
}

/*
CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden struct {
}

// IsSuccess returns true when this credential input sources credential input sources partial update forbidden response has a 2xx status code
func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this credential input sources credential input sources partial update forbidden response has a 3xx status code
func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this credential input sources credential input sources partial update forbidden response has a 4xx status code
func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this credential input sources credential input sources partial update forbidden response has a 5xx status code
func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this credential input sources credential input sources partial update forbidden response a status code equal to that given
func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/credential_input_sources/{id}/][%d] credentialInputSourcesCredentialInputSourcesPartialUpdateForbidden ", 403)
}

func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/credential_input_sources/{id}/][%d] credentialInputSourcesCredentialInputSourcesPartialUpdateForbidden ", 403)
}

func (o *CredentialInputSourcesCredentialInputSourcesPartialUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}