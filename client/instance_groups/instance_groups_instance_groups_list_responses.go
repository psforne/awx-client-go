// Code generated by go-swagger; DO NOT EDIT.

package instance_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InstanceGroupsInstanceGroupsListReader is a Reader for the InstanceGroupsInstanceGroupsList structure.
type InstanceGroupsInstanceGroupsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstanceGroupsInstanceGroupsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstanceGroupsInstanceGroupsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInstanceGroupsInstanceGroupsListOK creates a InstanceGroupsInstanceGroupsListOK with default headers values
func NewInstanceGroupsInstanceGroupsListOK() *InstanceGroupsInstanceGroupsListOK {
	return &InstanceGroupsInstanceGroupsListOK{}
}

/*
InstanceGroupsInstanceGroupsListOK describes a response with status code 200, with default header values.

InstanceGroupsInstanceGroupsListOK instance groups instance groups list o k
*/
type InstanceGroupsInstanceGroupsListOK struct {
}

// IsSuccess returns true when this instance groups instance groups list o k response has a 2xx status code
func (o *InstanceGroupsInstanceGroupsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this instance groups instance groups list o k response has a 3xx status code
func (o *InstanceGroupsInstanceGroupsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instance groups instance groups list o k response has a 4xx status code
func (o *InstanceGroupsInstanceGroupsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this instance groups instance groups list o k response has a 5xx status code
func (o *InstanceGroupsInstanceGroupsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this instance groups instance groups list o k response a status code equal to that given
func (o *InstanceGroupsInstanceGroupsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *InstanceGroupsInstanceGroupsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/instance_groups/][%d] instanceGroupsInstanceGroupsListOK ", 200)
}

func (o *InstanceGroupsInstanceGroupsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/instance_groups/][%d] instanceGroupsInstanceGroupsListOK ", 200)
}

func (o *InstanceGroupsInstanceGroupsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
