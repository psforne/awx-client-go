// Code generated by go-swagger; DO NOT EDIT.

package system_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SystemConfigurationConfigAttachCreateReader is a Reader for the SystemConfigurationConfigAttachCreate structure.
type SystemConfigurationConfigAttachCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemConfigurationConfigAttachCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSystemConfigurationConfigAttachCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemConfigurationConfigAttachCreateCreated creates a SystemConfigurationConfigAttachCreateCreated with default headers values
func NewSystemConfigurationConfigAttachCreateCreated() *SystemConfigurationConfigAttachCreateCreated {
	return &SystemConfigurationConfigAttachCreateCreated{}
}

/*
SystemConfigurationConfigAttachCreateCreated describes a response with status code 201, with default header values.

SystemConfigurationConfigAttachCreateCreated system configuration config attach create created
*/
type SystemConfigurationConfigAttachCreateCreated struct {
}

// IsSuccess returns true when this system configuration config attach create created response has a 2xx status code
func (o *SystemConfigurationConfigAttachCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this system configuration config attach create created response has a 3xx status code
func (o *SystemConfigurationConfigAttachCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system configuration config attach create created response has a 4xx status code
func (o *SystemConfigurationConfigAttachCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this system configuration config attach create created response has a 5xx status code
func (o *SystemConfigurationConfigAttachCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this system configuration config attach create created response a status code equal to that given
func (o *SystemConfigurationConfigAttachCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *SystemConfigurationConfigAttachCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/config/attach/][%d] systemConfigurationConfigAttachCreateCreated ", 201)
}

func (o *SystemConfigurationConfigAttachCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/config/attach/][%d] systemConfigurationConfigAttachCreateCreated ", 201)
}

func (o *SystemConfigurationConfigAttachCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}