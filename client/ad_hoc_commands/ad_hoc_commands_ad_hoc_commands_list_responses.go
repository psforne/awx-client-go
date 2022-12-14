// Code generated by go-swagger; DO NOT EDIT.

package ad_hoc_commands

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AdHocCommandsAdHocCommandsListReader is a Reader for the AdHocCommandsAdHocCommandsList structure.
type AdHocCommandsAdHocCommandsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdHocCommandsAdHocCommandsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdHocCommandsAdHocCommandsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdHocCommandsAdHocCommandsListOK creates a AdHocCommandsAdHocCommandsListOK with default headers values
func NewAdHocCommandsAdHocCommandsListOK() *AdHocCommandsAdHocCommandsListOK {
	return &AdHocCommandsAdHocCommandsListOK{}
}

/*
AdHocCommandsAdHocCommandsListOK describes a response with status code 200, with default header values.

AdHocCommandsAdHocCommandsListOK ad hoc commands ad hoc commands list o k
*/
type AdHocCommandsAdHocCommandsListOK struct {
}

// IsSuccess returns true when this ad hoc commands ad hoc commands list o k response has a 2xx status code
func (o *AdHocCommandsAdHocCommandsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ad hoc commands ad hoc commands list o k response has a 3xx status code
func (o *AdHocCommandsAdHocCommandsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ad hoc commands ad hoc commands list o k response has a 4xx status code
func (o *AdHocCommandsAdHocCommandsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ad hoc commands ad hoc commands list o k response has a 5xx status code
func (o *AdHocCommandsAdHocCommandsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ad hoc commands ad hoc commands list o k response a status code equal to that given
func (o *AdHocCommandsAdHocCommandsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *AdHocCommandsAdHocCommandsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/ad_hoc_commands/][%d] adHocCommandsAdHocCommandsListOK ", 200)
}

func (o *AdHocCommandsAdHocCommandsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/ad_hoc_commands/][%d] adHocCommandsAdHocCommandsListOK ", 200)
}

func (o *AdHocCommandsAdHocCommandsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
