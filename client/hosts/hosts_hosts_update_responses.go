// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HostsHostsUpdateReader is a Reader for the HostsHostsUpdate structure.
type HostsHostsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostsHostsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHostsHostsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewHostsHostsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHostsHostsUpdateOK creates a HostsHostsUpdateOK with default headers values
func NewHostsHostsUpdateOK() *HostsHostsUpdateOK {
	return &HostsHostsUpdateOK{}
}

/*
HostsHostsUpdateOK describes a response with status code 200, with default header values.

HostsHostsUpdateOK hosts hosts update o k
*/
type HostsHostsUpdateOK struct {
}

// IsSuccess returns true when this hosts hosts update o k response has a 2xx status code
func (o *HostsHostsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hosts hosts update o k response has a 3xx status code
func (o *HostsHostsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hosts hosts update o k response has a 4xx status code
func (o *HostsHostsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hosts hosts update o k response has a 5xx status code
func (o *HostsHostsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hosts hosts update o k response a status code equal to that given
func (o *HostsHostsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *HostsHostsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/hosts/{id}/][%d] hostsHostsUpdateOK ", 200)
}

func (o *HostsHostsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/hosts/{id}/][%d] hostsHostsUpdateOK ", 200)
}

func (o *HostsHostsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHostsHostsUpdateForbidden creates a HostsHostsUpdateForbidden with default headers values
func NewHostsHostsUpdateForbidden() *HostsHostsUpdateForbidden {
	return &HostsHostsUpdateForbidden{}
}

/*
HostsHostsUpdateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type HostsHostsUpdateForbidden struct {
}

// IsSuccess returns true when this hosts hosts update forbidden response has a 2xx status code
func (o *HostsHostsUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hosts hosts update forbidden response has a 3xx status code
func (o *HostsHostsUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hosts hosts update forbidden response has a 4xx status code
func (o *HostsHostsUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this hosts hosts update forbidden response has a 5xx status code
func (o *HostsHostsUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this hosts hosts update forbidden response a status code equal to that given
func (o *HostsHostsUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *HostsHostsUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/hosts/{id}/][%d] hostsHostsUpdateForbidden ", 403)
}

func (o *HostsHostsUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/hosts/{id}/][%d] hostsHostsUpdateForbidden ", 403)
}

func (o *HostsHostsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
