// Code generated by go-swagger; DO NOT EDIT.

package project_updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectUpdatesProjectUpdatesScmInventoryUpdatesListReader is a Reader for the ProjectUpdatesProjectUpdatesScmInventoryUpdatesList structure.
type ProjectUpdatesProjectUpdatesScmInventoryUpdatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectUpdatesProjectUpdatesScmInventoryUpdatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK creates a ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK with default headers values
func NewProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK() *ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK {
	return &ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK{}
}

/*
ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK describes a response with status code 200, with default header values.

ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK project updates project updates scm inventory updates list o k
*/
type ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK struct {
}

// IsSuccess returns true when this project updates project updates scm inventory updates list o k response has a 2xx status code
func (o *ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project updates project updates scm inventory updates list o k response has a 3xx status code
func (o *ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project updates project updates scm inventory updates list o k response has a 4xx status code
func (o *ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project updates project updates scm inventory updates list o k response has a 5xx status code
func (o *ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project updates project updates scm inventory updates list o k response a status code equal to that given
func (o *ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/project_updates/{id}/scm_inventory_updates/][%d] projectUpdatesProjectUpdatesScmInventoryUpdatesListOK ", 200)
}

func (o *ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/project_updates/{id}/scm_inventory_updates/][%d] projectUpdatesProjectUpdatesScmInventoryUpdatesListOK ", 200)
}

func (o *ProjectUpdatesProjectUpdatesScmInventoryUpdatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}