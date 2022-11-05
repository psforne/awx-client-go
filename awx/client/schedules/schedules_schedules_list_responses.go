// Code generated by go-swagger; DO NOT EDIT.

package schedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SchedulesSchedulesListReader is a Reader for the SchedulesSchedulesList structure.
type SchedulesSchedulesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchedulesSchedulesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchedulesSchedulesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSchedulesSchedulesListOK creates a SchedulesSchedulesListOK with default headers values
func NewSchedulesSchedulesListOK() *SchedulesSchedulesListOK {
	return &SchedulesSchedulesListOK{}
}

/*
SchedulesSchedulesListOK describes a response with status code 200, with default header values.

SchedulesSchedulesListOK schedules schedules list o k
*/
type SchedulesSchedulesListOK struct {
}

// IsSuccess returns true when this schedules schedules list o k response has a 2xx status code
func (o *SchedulesSchedulesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schedules schedules list o k response has a 3xx status code
func (o *SchedulesSchedulesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedules schedules list o k response has a 4xx status code
func (o *SchedulesSchedulesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this schedules schedules list o k response has a 5xx status code
func (o *SchedulesSchedulesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this schedules schedules list o k response a status code equal to that given
func (o *SchedulesSchedulesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SchedulesSchedulesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/schedules/][%d] schedulesSchedulesListOK ", 200)
}

func (o *SchedulesSchedulesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/schedules/][%d] schedulesSchedulesListOK ", 200)
}

func (o *SchedulesSchedulesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
