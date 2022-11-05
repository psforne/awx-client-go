// Code generated by go-swagger; DO NOT EDIT.

package instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InstancesInstancesJobsListReader is a Reader for the InstancesInstancesJobsList structure.
type InstancesInstancesJobsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstancesInstancesJobsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstancesInstancesJobsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInstancesInstancesJobsListOK creates a InstancesInstancesJobsListOK with default headers values
func NewInstancesInstancesJobsListOK() *InstancesInstancesJobsListOK {
	return &InstancesInstancesJobsListOK{}
}

/*
InstancesInstancesJobsListOK describes a response with status code 200, with default header values.

InstancesInstancesJobsListOK instances instances jobs list o k
*/
type InstancesInstancesJobsListOK struct {
}

// IsSuccess returns true when this instances instances jobs list o k response has a 2xx status code
func (o *InstancesInstancesJobsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this instances instances jobs list o k response has a 3xx status code
func (o *InstancesInstancesJobsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instances instances jobs list o k response has a 4xx status code
func (o *InstancesInstancesJobsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this instances instances jobs list o k response has a 5xx status code
func (o *InstancesInstancesJobsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this instances instances jobs list o k response a status code equal to that given
func (o *InstancesInstancesJobsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *InstancesInstancesJobsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/instances/{id}/jobs/][%d] instancesInstancesJobsListOK ", 200)
}

func (o *InstancesInstancesJobsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/instances/{id}/jobs/][%d] instancesInstancesJobsListOK ", 200)
}

func (o *InstancesInstancesJobsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}