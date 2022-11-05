// Code generated by go-swagger; DO NOT EDIT.

package instance_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InstanceGroupsInstanceGroupsInstancesListReader is a Reader for the InstanceGroupsInstanceGroupsInstancesList structure.
type InstanceGroupsInstanceGroupsInstancesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstanceGroupsInstanceGroupsInstancesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstanceGroupsInstanceGroupsInstancesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInstanceGroupsInstanceGroupsInstancesListOK creates a InstanceGroupsInstanceGroupsInstancesListOK with default headers values
func NewInstanceGroupsInstanceGroupsInstancesListOK() *InstanceGroupsInstanceGroupsInstancesListOK {
	return &InstanceGroupsInstanceGroupsInstancesListOK{}
}

/*
InstanceGroupsInstanceGroupsInstancesListOK describes a response with status code 200, with default header values.

InstanceGroupsInstanceGroupsInstancesListOK instance groups instance groups instances list o k
*/
type InstanceGroupsInstanceGroupsInstancesListOK struct {
}

// IsSuccess returns true when this instance groups instance groups instances list o k response has a 2xx status code
func (o *InstanceGroupsInstanceGroupsInstancesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this instance groups instance groups instances list o k response has a 3xx status code
func (o *InstanceGroupsInstanceGroupsInstancesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instance groups instance groups instances list o k response has a 4xx status code
func (o *InstanceGroupsInstanceGroupsInstancesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this instance groups instance groups instances list o k response has a 5xx status code
func (o *InstanceGroupsInstanceGroupsInstancesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this instance groups instance groups instances list o k response a status code equal to that given
func (o *InstanceGroupsInstanceGroupsInstancesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *InstanceGroupsInstanceGroupsInstancesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/instance_groups/{id}/instances/][%d] instanceGroupsInstanceGroupsInstancesListOK ", 200)
}

func (o *InstanceGroupsInstanceGroupsInstancesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/instance_groups/{id}/instances/][%d] instanceGroupsInstanceGroupsInstancesListOK ", 200)
}

func (o *InstanceGroupsInstanceGroupsInstancesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
