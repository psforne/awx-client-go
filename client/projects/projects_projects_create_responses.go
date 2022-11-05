// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectsProjectsCreateReader is a Reader for the ProjectsProjectsCreate structure.
type ProjectsProjectsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsProjectsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewProjectsProjectsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 415:
		result := NewProjectsProjectsCreateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsProjectsCreateCreated creates a ProjectsProjectsCreateCreated with default headers values
func NewProjectsProjectsCreateCreated() *ProjectsProjectsCreateCreated {
	return &ProjectsProjectsCreateCreated{}
}

/*
ProjectsProjectsCreateCreated describes a response with status code 201, with default header values.

ProjectsProjectsCreateCreated projects projects create created
*/
type ProjectsProjectsCreateCreated struct {
}

// IsSuccess returns true when this projects projects create created response has a 2xx status code
func (o *ProjectsProjectsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects projects create created response has a 3xx status code
func (o *ProjectsProjectsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects projects create created response has a 4xx status code
func (o *ProjectsProjectsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects projects create created response has a 5xx status code
func (o *ProjectsProjectsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this projects projects create created response a status code equal to that given
func (o *ProjectsProjectsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *ProjectsProjectsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/][%d] projectsProjectsCreateCreated ", 201)
}

func (o *ProjectsProjectsCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/][%d] projectsProjectsCreateCreated ", 201)
}

func (o *ProjectsProjectsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsProjectsCreateUnsupportedMediaType creates a ProjectsProjectsCreateUnsupportedMediaType with default headers values
func NewProjectsProjectsCreateUnsupportedMediaType() *ProjectsProjectsCreateUnsupportedMediaType {
	return &ProjectsProjectsCreateUnsupportedMediaType{}
}

/*
ProjectsProjectsCreateUnsupportedMediaType describes a response with status code 415, with default header values.

ProjectsProjectsCreateUnsupportedMediaType projects projects create unsupported media type
*/
type ProjectsProjectsCreateUnsupportedMediaType struct {
}

// IsSuccess returns true when this projects projects create unsupported media type response has a 2xx status code
func (o *ProjectsProjectsCreateUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects projects create unsupported media type response has a 3xx status code
func (o *ProjectsProjectsCreateUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects projects create unsupported media type response has a 4xx status code
func (o *ProjectsProjectsCreateUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects projects create unsupported media type response has a 5xx status code
func (o *ProjectsProjectsCreateUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this projects projects create unsupported media type response a status code equal to that given
func (o *ProjectsProjectsCreateUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *ProjectsProjectsCreateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/][%d] projectsProjectsCreateUnsupportedMediaType ", 415)
}

func (o *ProjectsProjectsCreateUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/projects/][%d] projectsProjectsCreateUnsupportedMediaType ", 415)
}

func (o *ProjectsProjectsCreateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}