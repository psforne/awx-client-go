// Code generated by go-swagger; DO NOT EDIT.

package execution_environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListReader is a Reader for the ExecutionEnvironmentsExecutionEnvironmentsActivityStreamList structure.
type ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK creates a ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK with default headers values
func NewExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK() *ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK {
	return &ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK{}
}

/*
ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK describes a response with status code 200, with default header values.

ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK execution environments execution environments activity stream list o k
*/
type ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK struct {
}

// IsSuccess returns true when this execution environments execution environments activity stream list o k response has a 2xx status code
func (o *ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this execution environments execution environments activity stream list o k response has a 3xx status code
func (o *ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this execution environments execution environments activity stream list o k response has a 4xx status code
func (o *ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this execution environments execution environments activity stream list o k response has a 5xx status code
func (o *ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this execution environments execution environments activity stream list o k response a status code equal to that given
func (o *ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/execution_environments/{id}/activity_stream/][%d] executionEnvironmentsExecutionEnvironmentsActivityStreamListOK ", 200)
}

func (o *ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/execution_environments/{id}/activity_stream/][%d] executionEnvironmentsExecutionEnvironmentsActivityStreamListOK ", 200)
}

func (o *ExecutionEnvironmentsExecutionEnvironmentsActivityStreamListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}