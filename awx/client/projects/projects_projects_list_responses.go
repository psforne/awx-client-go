// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectsProjectsListReader is a Reader for the ProjectsProjectsList structure.
type ProjectsProjectsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsProjectsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsProjectsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsProjectsListOK creates a ProjectsProjectsListOK with default headers values
func NewProjectsProjectsListOK() *ProjectsProjectsListOK {
	return &ProjectsProjectsListOK{}
}

/*
ProjectsProjectsListOK describes a response with status code 200, with default header values.

ProjectsProjectsListOK projects projects list o k
*/
type ProjectsProjectsListOK struct {
}

// IsSuccess returns true when this projects projects list o k response has a 2xx status code
func (o *ProjectsProjectsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects projects list o k response has a 3xx status code
func (o *ProjectsProjectsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects projects list o k response has a 4xx status code
func (o *ProjectsProjectsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects projects list o k response has a 5xx status code
func (o *ProjectsProjectsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects projects list o k response a status code equal to that given
func (o *ProjectsProjectsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsProjectsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/][%d] projectsProjectsListOK ", 200)
}

func (o *ProjectsProjectsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/][%d] projectsProjectsListOK ", 200)
}

func (o *ProjectsProjectsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
