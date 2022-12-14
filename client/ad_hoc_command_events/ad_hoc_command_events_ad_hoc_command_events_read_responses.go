// Code generated by go-swagger; DO NOT EDIT.

package ad_hoc_command_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AdHocCommandEventsAdHocCommandEventsReadReader is a Reader for the AdHocCommandEventsAdHocCommandEventsRead structure.
type AdHocCommandEventsAdHocCommandEventsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdHocCommandEventsAdHocCommandEventsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdHocCommandEventsAdHocCommandEventsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdHocCommandEventsAdHocCommandEventsReadOK creates a AdHocCommandEventsAdHocCommandEventsReadOK with default headers values
func NewAdHocCommandEventsAdHocCommandEventsReadOK() *AdHocCommandEventsAdHocCommandEventsReadOK {
	return &AdHocCommandEventsAdHocCommandEventsReadOK{}
}

/*
AdHocCommandEventsAdHocCommandEventsReadOK describes a response with status code 200, with default header values.

AdHocCommandEventsAdHocCommandEventsReadOK ad hoc command events ad hoc command events read o k
*/
type AdHocCommandEventsAdHocCommandEventsReadOK struct {
}

// IsSuccess returns true when this ad hoc command events ad hoc command events read o k response has a 2xx status code
func (o *AdHocCommandEventsAdHocCommandEventsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ad hoc command events ad hoc command events read o k response has a 3xx status code
func (o *AdHocCommandEventsAdHocCommandEventsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ad hoc command events ad hoc command events read o k response has a 4xx status code
func (o *AdHocCommandEventsAdHocCommandEventsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ad hoc command events ad hoc command events read o k response has a 5xx status code
func (o *AdHocCommandEventsAdHocCommandEventsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ad hoc command events ad hoc command events read o k response a status code equal to that given
func (o *AdHocCommandEventsAdHocCommandEventsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *AdHocCommandEventsAdHocCommandEventsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/ad_hoc_command_events/{id}/][%d] adHocCommandEventsAdHocCommandEventsReadOK ", 200)
}

func (o *AdHocCommandEventsAdHocCommandEventsReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/ad_hoc_command_events/{id}/][%d] adHocCommandEventsAdHocCommandEventsReadOK ", 200)
}

func (o *AdHocCommandEventsAdHocCommandEventsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
