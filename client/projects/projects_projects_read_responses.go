// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectsProjectsReadReader is a Reader for the ProjectsProjectsRead structure.
type ProjectsProjectsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsProjectsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsProjectsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsProjectsReadOK creates a ProjectsProjectsReadOK with default headers values
func NewProjectsProjectsReadOK() *ProjectsProjectsReadOK {
	return &ProjectsProjectsReadOK{}
}

/*
ProjectsProjectsReadOK describes a response with status code 200, with default header values.

ProjectsProjectsReadOK projects projects read o k
*/
type ProjectsProjectsReadOK struct {
}

// IsSuccess returns true when this projects projects read o k response has a 2xx status code
func (o *ProjectsProjectsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects projects read o k response has a 3xx status code
func (o *ProjectsProjectsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects projects read o k response has a 4xx status code
func (o *ProjectsProjectsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects projects read o k response has a 5xx status code
func (o *ProjectsProjectsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects projects read o k response a status code equal to that given
func (o *ProjectsProjectsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsProjectsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{id}/][%d] projectsProjectsReadOK ", 200)
}

func (o *ProjectsProjectsReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{id}/][%d] projectsProjectsReadOK ", 200)
}

func (o *ProjectsProjectsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
