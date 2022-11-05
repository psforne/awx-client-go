// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationsOrganizationsActivityStreamListReader is a Reader for the OrganizationsOrganizationsActivityStreamList structure.
type OrganizationsOrganizationsActivityStreamListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationsActivityStreamListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsOrganizationsActivityStreamListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsOrganizationsActivityStreamListOK creates a OrganizationsOrganizationsActivityStreamListOK with default headers values
func NewOrganizationsOrganizationsActivityStreamListOK() *OrganizationsOrganizationsActivityStreamListOK {
	return &OrganizationsOrganizationsActivityStreamListOK{}
}

/*
OrganizationsOrganizationsActivityStreamListOK describes a response with status code 200, with default header values.

OrganizationsOrganizationsActivityStreamListOK organizations organizations activity stream list o k
*/
type OrganizationsOrganizationsActivityStreamListOK struct {
}

// IsSuccess returns true when this organizations organizations activity stream list o k response has a 2xx status code
func (o *OrganizationsOrganizationsActivityStreamListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations organizations activity stream list o k response has a 3xx status code
func (o *OrganizationsOrganizationsActivityStreamListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations organizations activity stream list o k response has a 4xx status code
func (o *OrganizationsOrganizationsActivityStreamListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations organizations activity stream list o k response has a 5xx status code
func (o *OrganizationsOrganizationsActivityStreamListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations organizations activity stream list o k response a status code equal to that given
func (o *OrganizationsOrganizationsActivityStreamListOK) IsCode(code int) bool {
	return code == 200
}

func (o *OrganizationsOrganizationsActivityStreamListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/activity_stream/][%d] organizationsOrganizationsActivityStreamListOK ", 200)
}

func (o *OrganizationsOrganizationsActivityStreamListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/activity_stream/][%d] organizationsOrganizationsActivityStreamListOK ", 200)
}

func (o *OrganizationsOrganizationsActivityStreamListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}