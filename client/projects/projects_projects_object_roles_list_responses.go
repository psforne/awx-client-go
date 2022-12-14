// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectsProjectsObjectRolesListReader is a Reader for the ProjectsProjectsObjectRolesList structure.
type ProjectsProjectsObjectRolesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsProjectsObjectRolesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsProjectsObjectRolesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsProjectsObjectRolesListOK creates a ProjectsProjectsObjectRolesListOK with default headers values
func NewProjectsProjectsObjectRolesListOK() *ProjectsProjectsObjectRolesListOK {
	return &ProjectsProjectsObjectRolesListOK{}
}

/*
ProjectsProjectsObjectRolesListOK describes a response with status code 200, with default header values.

ProjectsProjectsObjectRolesListOK projects projects object roles list o k
*/
type ProjectsProjectsObjectRolesListOK struct {
}

// IsSuccess returns true when this projects projects object roles list o k response has a 2xx status code
func (o *ProjectsProjectsObjectRolesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects projects object roles list o k response has a 3xx status code
func (o *ProjectsProjectsObjectRolesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects projects object roles list o k response has a 4xx status code
func (o *ProjectsProjectsObjectRolesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects projects object roles list o k response has a 5xx status code
func (o *ProjectsProjectsObjectRolesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects projects object roles list o k response a status code equal to that given
func (o *ProjectsProjectsObjectRolesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsProjectsObjectRolesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{id}/object_roles/][%d] projectsProjectsObjectRolesListOK ", 200)
}

func (o *ProjectsProjectsObjectRolesListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{id}/object_roles/][%d] projectsProjectsObjectRolesListOK ", 200)
}

func (o *ProjectsProjectsObjectRolesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
