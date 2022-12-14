// Code generated by go-swagger; DO NOT EDIT.

package system_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SystemConfigurationConfigSubscriptionsCreateReader is a Reader for the SystemConfigurationConfigSubscriptionsCreate structure.
type SystemConfigurationConfigSubscriptionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemConfigurationConfigSubscriptionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSystemConfigurationConfigSubscriptionsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemConfigurationConfigSubscriptionsCreateCreated creates a SystemConfigurationConfigSubscriptionsCreateCreated with default headers values
func NewSystemConfigurationConfigSubscriptionsCreateCreated() *SystemConfigurationConfigSubscriptionsCreateCreated {
	return &SystemConfigurationConfigSubscriptionsCreateCreated{}
}

/*
SystemConfigurationConfigSubscriptionsCreateCreated describes a response with status code 201, with default header values.

SystemConfigurationConfigSubscriptionsCreateCreated system configuration config subscriptions create created
*/
type SystemConfigurationConfigSubscriptionsCreateCreated struct {
}

// IsSuccess returns true when this system configuration config subscriptions create created response has a 2xx status code
func (o *SystemConfigurationConfigSubscriptionsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this system configuration config subscriptions create created response has a 3xx status code
func (o *SystemConfigurationConfigSubscriptionsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system configuration config subscriptions create created response has a 4xx status code
func (o *SystemConfigurationConfigSubscriptionsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this system configuration config subscriptions create created response has a 5xx status code
func (o *SystemConfigurationConfigSubscriptionsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this system configuration config subscriptions create created response a status code equal to that given
func (o *SystemConfigurationConfigSubscriptionsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *SystemConfigurationConfigSubscriptionsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/config/subscriptions/][%d] systemConfigurationConfigSubscriptionsCreateCreated ", 201)
}

func (o *SystemConfigurationConfigSubscriptionsCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/config/subscriptions/][%d] systemConfigurationConfigSubscriptionsCreateCreated ", 201)
}

func (o *SystemConfigurationConfigSubscriptionsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
