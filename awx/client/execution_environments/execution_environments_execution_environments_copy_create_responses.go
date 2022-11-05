// Code generated by go-swagger; DO NOT EDIT.

package execution_environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExecutionEnvironmentsExecutionEnvironmentsCopyCreateReader is a Reader for the ExecutionEnvironmentsExecutionEnvironmentsCopyCreate structure.
type ExecutionEnvironmentsExecutionEnvironmentsCopyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated creates a ExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated with default headers values
func NewExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated() *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated {
	return &ExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated{}
}

/*
ExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated describes a response with status code 201, with default header values.

ExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated execution environments execution environments copy create created
*/
type ExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated struct {
}

// IsSuccess returns true when this execution environments execution environments copy create created response has a 2xx status code
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this execution environments execution environments copy create created response has a 3xx status code
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this execution environments execution environments copy create created response has a 4xx status code
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this execution environments execution environments copy create created response has a 5xx status code
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this execution environments execution environments copy create created response a status code equal to that given
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/execution_environments/{id}/copy/][%d] executionEnvironmentsExecutionEnvironmentsCopyCreateCreated ", 201)
}

func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/execution_environments/{id}/copy/][%d] executionEnvironmentsExecutionEnvironmentsCopyCreateCreated ", 201)
}

func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
ExecutionEnvironmentsExecutionEnvironmentsCopyCreateBody execution environments execution environments copy create body
swagger:model ExecutionEnvironmentsExecutionEnvironmentsCopyCreateBody
*/
type ExecutionEnvironmentsExecutionEnvironmentsCopyCreateBody struct {

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this execution environments execution environments copy create body
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this execution environments execution environments copy create body based on context it is used
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExecutionEnvironmentsExecutionEnvironmentsCopyCreateBody) UnmarshalBinary(b []byte) error {
	var res ExecutionEnvironmentsExecutionEnvironmentsCopyCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}