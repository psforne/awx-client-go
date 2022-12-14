// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HostsHostsActivityStreamListReader is a Reader for the HostsHostsActivityStreamList structure.
type HostsHostsActivityStreamListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HostsHostsActivityStreamListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHostsHostsActivityStreamListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHostsHostsActivityStreamListOK creates a HostsHostsActivityStreamListOK with default headers values
func NewHostsHostsActivityStreamListOK() *HostsHostsActivityStreamListOK {
	return &HostsHostsActivityStreamListOK{}
}

/*
HostsHostsActivityStreamListOK describes a response with status code 200, with default header values.

HostsHostsActivityStreamListOK hosts hosts activity stream list o k
*/
type HostsHostsActivityStreamListOK struct {
}

// IsSuccess returns true when this hosts hosts activity stream list o k response has a 2xx status code
func (o *HostsHostsActivityStreamListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hosts hosts activity stream list o k response has a 3xx status code
func (o *HostsHostsActivityStreamListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hosts hosts activity stream list o k response has a 4xx status code
func (o *HostsHostsActivityStreamListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hosts hosts activity stream list o k response has a 5xx status code
func (o *HostsHostsActivityStreamListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hosts hosts activity stream list o k response a status code equal to that given
func (o *HostsHostsActivityStreamListOK) IsCode(code int) bool {
	return code == 200
}

func (o *HostsHostsActivityStreamListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/hosts/{id}/activity_stream/][%d] hostsHostsActivityStreamListOK ", 200)
}

func (o *HostsHostsActivityStreamListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/hosts/{id}/activity_stream/][%d] hostsHostsActivityStreamListOK ", 200)
}

func (o *HostsHostsActivityStreamListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
