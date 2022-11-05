// Code generated by go-swagger; DO NOT EDIT.

package instance_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InstanceGroupsInstanceGroupsInstancesCreateReader is a Reader for the InstanceGroupsInstanceGroupsInstancesCreate structure.
type InstanceGroupsInstanceGroupsInstancesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstanceGroupsInstanceGroupsInstancesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInstanceGroupsInstanceGroupsInstancesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInstanceGroupsInstanceGroupsInstancesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInstanceGroupsInstanceGroupsInstancesCreateCreated creates a InstanceGroupsInstanceGroupsInstancesCreateCreated with default headers values
func NewInstanceGroupsInstanceGroupsInstancesCreateCreated() *InstanceGroupsInstanceGroupsInstancesCreateCreated {
	return &InstanceGroupsInstanceGroupsInstancesCreateCreated{}
}

/*
InstanceGroupsInstanceGroupsInstancesCreateCreated describes a response with status code 201, with default header values.

InstanceGroupsInstanceGroupsInstancesCreateCreated instance groups instance groups instances create created
*/
type InstanceGroupsInstanceGroupsInstancesCreateCreated struct {
}

// IsSuccess returns true when this instance groups instance groups instances create created response has a 2xx status code
func (o *InstanceGroupsInstanceGroupsInstancesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this instance groups instance groups instances create created response has a 3xx status code
func (o *InstanceGroupsInstanceGroupsInstancesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instance groups instance groups instances create created response has a 4xx status code
func (o *InstanceGroupsInstanceGroupsInstancesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this instance groups instance groups instances create created response has a 5xx status code
func (o *InstanceGroupsInstanceGroupsInstancesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this instance groups instance groups instances create created response a status code equal to that given
func (o *InstanceGroupsInstanceGroupsInstancesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *InstanceGroupsInstanceGroupsInstancesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/instance_groups/{id}/instances/][%d] instanceGroupsInstanceGroupsInstancesCreateCreated ", 201)
}

func (o *InstanceGroupsInstanceGroupsInstancesCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/instance_groups/{id}/instances/][%d] instanceGroupsInstanceGroupsInstancesCreateCreated ", 201)
}

func (o *InstanceGroupsInstanceGroupsInstancesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstanceGroupsInstanceGroupsInstancesCreateBadRequest creates a InstanceGroupsInstanceGroupsInstancesCreateBadRequest with default headers values
func NewInstanceGroupsInstanceGroupsInstancesCreateBadRequest() *InstanceGroupsInstanceGroupsInstancesCreateBadRequest {
	return &InstanceGroupsInstanceGroupsInstancesCreateBadRequest{}
}

/*
InstanceGroupsInstanceGroupsInstancesCreateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type InstanceGroupsInstanceGroupsInstancesCreateBadRequest struct {
}

// IsSuccess returns true when this instance groups instance groups instances create bad request response has a 2xx status code
func (o *InstanceGroupsInstanceGroupsInstancesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this instance groups instance groups instances create bad request response has a 3xx status code
func (o *InstanceGroupsInstanceGroupsInstancesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instance groups instance groups instances create bad request response has a 4xx status code
func (o *InstanceGroupsInstanceGroupsInstancesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this instance groups instance groups instances create bad request response has a 5xx status code
func (o *InstanceGroupsInstanceGroupsInstancesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this instance groups instance groups instances create bad request response a status code equal to that given
func (o *InstanceGroupsInstanceGroupsInstancesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *InstanceGroupsInstanceGroupsInstancesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/instance_groups/{id}/instances/][%d] instanceGroupsInstanceGroupsInstancesCreateBadRequest ", 400)
}

func (o *InstanceGroupsInstanceGroupsInstancesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/instance_groups/{id}/instances/][%d] instanceGroupsInstanceGroupsInstancesCreateBadRequest ", 400)
}

func (o *InstanceGroupsInstanceGroupsInstancesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
