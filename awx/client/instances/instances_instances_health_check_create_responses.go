// Code generated by go-swagger; DO NOT EDIT.

package instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InstancesInstancesHealthCheckCreateReader is a Reader for the InstancesInstancesHealthCheckCreate structure.
type InstancesInstancesHealthCheckCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstancesInstancesHealthCheckCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstancesInstancesHealthCheckCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewInstancesInstancesHealthCheckCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewInstancesInstancesHealthCheckCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInstancesInstancesHealthCheckCreateOK creates a InstancesInstancesHealthCheckCreateOK with default headers values
func NewInstancesInstancesHealthCheckCreateOK() *InstancesInstancesHealthCheckCreateOK {
	return &InstancesInstancesHealthCheckCreateOK{}
}

/*
InstancesInstancesHealthCheckCreateOK describes a response with status code 200, with default header values.

InstancesInstancesHealthCheckCreateOK instances instances health check create o k
*/
type InstancesInstancesHealthCheckCreateOK struct {
}

// IsSuccess returns true when this instances instances health check create o k response has a 2xx status code
func (o *InstancesInstancesHealthCheckCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this instances instances health check create o k response has a 3xx status code
func (o *InstancesInstancesHealthCheckCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instances instances health check create o k response has a 4xx status code
func (o *InstancesInstancesHealthCheckCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this instances instances health check create o k response has a 5xx status code
func (o *InstancesInstancesHealthCheckCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this instances instances health check create o k response a status code equal to that given
func (o *InstancesInstancesHealthCheckCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *InstancesInstancesHealthCheckCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/instances/{id}/health_check/][%d] instancesInstancesHealthCheckCreateOK ", 200)
}

func (o *InstancesInstancesHealthCheckCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v2/instances/{id}/health_check/][%d] instancesInstancesHealthCheckCreateOK ", 200)
}

func (o *InstancesInstancesHealthCheckCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstancesInstancesHealthCheckCreateCreated creates a InstancesInstancesHealthCheckCreateCreated with default headers values
func NewInstancesInstancesHealthCheckCreateCreated() *InstancesInstancesHealthCheckCreateCreated {
	return &InstancesInstancesHealthCheckCreateCreated{}
}

/*
InstancesInstancesHealthCheckCreateCreated describes a response with status code 201, with default header values.

InstancesInstancesHealthCheckCreateCreated instances instances health check create created
*/
type InstancesInstancesHealthCheckCreateCreated struct {
}

// IsSuccess returns true when this instances instances health check create created response has a 2xx status code
func (o *InstancesInstancesHealthCheckCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this instances instances health check create created response has a 3xx status code
func (o *InstancesInstancesHealthCheckCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instances instances health check create created response has a 4xx status code
func (o *InstancesInstancesHealthCheckCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this instances instances health check create created response has a 5xx status code
func (o *InstancesInstancesHealthCheckCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this instances instances health check create created response a status code equal to that given
func (o *InstancesInstancesHealthCheckCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *InstancesInstancesHealthCheckCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/instances/{id}/health_check/][%d] instancesInstancesHealthCheckCreateCreated ", 201)
}

func (o *InstancesInstancesHealthCheckCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/instances/{id}/health_check/][%d] instancesInstancesHealthCheckCreateCreated ", 201)
}

func (o *InstancesInstancesHealthCheckCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstancesInstancesHealthCheckCreateForbidden creates a InstancesInstancesHealthCheckCreateForbidden with default headers values
func NewInstancesInstancesHealthCheckCreateForbidden() *InstancesInstancesHealthCheckCreateForbidden {
	return &InstancesInstancesHealthCheckCreateForbidden{}
}

/*
InstancesInstancesHealthCheckCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type InstancesInstancesHealthCheckCreateForbidden struct {
}

// IsSuccess returns true when this instances instances health check create forbidden response has a 2xx status code
func (o *InstancesInstancesHealthCheckCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this instances instances health check create forbidden response has a 3xx status code
func (o *InstancesInstancesHealthCheckCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instances instances health check create forbidden response has a 4xx status code
func (o *InstancesInstancesHealthCheckCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this instances instances health check create forbidden response has a 5xx status code
func (o *InstancesInstancesHealthCheckCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this instances instances health check create forbidden response a status code equal to that given
func (o *InstancesInstancesHealthCheckCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *InstancesInstancesHealthCheckCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/instances/{id}/health_check/][%d] instancesInstancesHealthCheckCreateForbidden ", 403)
}

func (o *InstancesInstancesHealthCheckCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/instances/{id}/health_check/][%d] instancesInstancesHealthCheckCreateForbidden ", 403)
}

func (o *InstancesInstancesHealthCheckCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
