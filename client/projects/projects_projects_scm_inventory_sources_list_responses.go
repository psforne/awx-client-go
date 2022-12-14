// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectsProjectsScmInventorySourcesListReader is a Reader for the ProjectsProjectsScmInventorySourcesList structure.
type ProjectsProjectsScmInventorySourcesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsProjectsScmInventorySourcesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsProjectsScmInventorySourcesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsProjectsScmInventorySourcesListOK creates a ProjectsProjectsScmInventorySourcesListOK with default headers values
func NewProjectsProjectsScmInventorySourcesListOK() *ProjectsProjectsScmInventorySourcesListOK {
	return &ProjectsProjectsScmInventorySourcesListOK{}
}

/*
ProjectsProjectsScmInventorySourcesListOK describes a response with status code 200, with default header values.

ProjectsProjectsScmInventorySourcesListOK projects projects scm inventory sources list o k
*/
type ProjectsProjectsScmInventorySourcesListOK struct {
}

// IsSuccess returns true when this projects projects scm inventory sources list o k response has a 2xx status code
func (o *ProjectsProjectsScmInventorySourcesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects projects scm inventory sources list o k response has a 3xx status code
func (o *ProjectsProjectsScmInventorySourcesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects projects scm inventory sources list o k response has a 4xx status code
func (o *ProjectsProjectsScmInventorySourcesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects projects scm inventory sources list o k response has a 5xx status code
func (o *ProjectsProjectsScmInventorySourcesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects projects scm inventory sources list o k response a status code equal to that given
func (o *ProjectsProjectsScmInventorySourcesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsProjectsScmInventorySourcesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{id}/scm_inventory_sources/][%d] projectsProjectsScmInventorySourcesListOK ", 200)
}

func (o *ProjectsProjectsScmInventorySourcesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{id}/scm_inventory_sources/][%d] projectsProjectsScmInventorySourcesListOK ", 200)
}

func (o *ProjectsProjectsScmInventorySourcesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
