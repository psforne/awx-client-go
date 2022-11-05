// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationsOrganizationsInventoriesListReader is a Reader for the OrganizationsOrganizationsInventoriesList structure.
type OrganizationsOrganizationsInventoriesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationsInventoriesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsOrganizationsInventoriesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewOrganizationsOrganizationsInventoriesListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsOrganizationsInventoriesListOK creates a OrganizationsOrganizationsInventoriesListOK with default headers values
func NewOrganizationsOrganizationsInventoriesListOK() *OrganizationsOrganizationsInventoriesListOK {
	return &OrganizationsOrganizationsInventoriesListOK{}
}

/*
OrganizationsOrganizationsInventoriesListOK describes a response with status code 200, with default header values.

OrganizationsOrganizationsInventoriesListOK organizations organizations inventories list o k
*/
type OrganizationsOrganizationsInventoriesListOK struct {
}

// IsSuccess returns true when this organizations organizations inventories list o k response has a 2xx status code
func (o *OrganizationsOrganizationsInventoriesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations organizations inventories list o k response has a 3xx status code
func (o *OrganizationsOrganizationsInventoriesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations organizations inventories list o k response has a 4xx status code
func (o *OrganizationsOrganizationsInventoriesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations organizations inventories list o k response has a 5xx status code
func (o *OrganizationsOrganizationsInventoriesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations organizations inventories list o k response a status code equal to that given
func (o *OrganizationsOrganizationsInventoriesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *OrganizationsOrganizationsInventoriesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/inventories/][%d] organizationsOrganizationsInventoriesListOK ", 200)
}

func (o *OrganizationsOrganizationsInventoriesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/inventories/][%d] organizationsOrganizationsInventoriesListOK ", 200)
}

func (o *OrganizationsOrganizationsInventoriesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrganizationsOrganizationsInventoriesListForbidden creates a OrganizationsOrganizationsInventoriesListForbidden with default headers values
func NewOrganizationsOrganizationsInventoriesListForbidden() *OrganizationsOrganizationsInventoriesListForbidden {
	return &OrganizationsOrganizationsInventoriesListForbidden{}
}

/*
OrganizationsOrganizationsInventoriesListForbidden describes a response with status code 403, with default header values.

forbidden
*/
type OrganizationsOrganizationsInventoriesListForbidden struct {
}

// IsSuccess returns true when this organizations organizations inventories list forbidden response has a 2xx status code
func (o *OrganizationsOrganizationsInventoriesListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations organizations inventories list forbidden response has a 3xx status code
func (o *OrganizationsOrganizationsInventoriesListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations organizations inventories list forbidden response has a 4xx status code
func (o *OrganizationsOrganizationsInventoriesListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations organizations inventories list forbidden response has a 5xx status code
func (o *OrganizationsOrganizationsInventoriesListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations organizations inventories list forbidden response a status code equal to that given
func (o *OrganizationsOrganizationsInventoriesListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OrganizationsOrganizationsInventoriesListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/inventories/][%d] organizationsOrganizationsInventoriesListForbidden ", 403)
}

func (o *OrganizationsOrganizationsInventoriesListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/organizations/{id}/inventories/][%d] organizationsOrganizationsInventoriesListForbidden ", 403)
}

func (o *OrganizationsOrganizationsInventoriesListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}