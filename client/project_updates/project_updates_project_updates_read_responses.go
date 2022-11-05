// Code generated by go-swagger; DO NOT EDIT.

package project_updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectUpdatesProjectUpdatesReadReader is a Reader for the ProjectUpdatesProjectUpdatesRead structure.
type ProjectUpdatesProjectUpdatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectUpdatesProjectUpdatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectUpdatesProjectUpdatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectUpdatesProjectUpdatesReadOK creates a ProjectUpdatesProjectUpdatesReadOK with default headers values
func NewProjectUpdatesProjectUpdatesReadOK() *ProjectUpdatesProjectUpdatesReadOK {
	return &ProjectUpdatesProjectUpdatesReadOK{}
}

/*
ProjectUpdatesProjectUpdatesReadOK describes a response with status code 200, with default header values.

ProjectUpdatesProjectUpdatesReadOK project updates project updates read o k
*/
type ProjectUpdatesProjectUpdatesReadOK struct {
}

// IsSuccess returns true when this project updates project updates read o k response has a 2xx status code
func (o *ProjectUpdatesProjectUpdatesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project updates project updates read o k response has a 3xx status code
func (o *ProjectUpdatesProjectUpdatesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project updates project updates read o k response has a 4xx status code
func (o *ProjectUpdatesProjectUpdatesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project updates project updates read o k response has a 5xx status code
func (o *ProjectUpdatesProjectUpdatesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project updates project updates read o k response a status code equal to that given
func (o *ProjectUpdatesProjectUpdatesReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectUpdatesProjectUpdatesReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/project_updates/{id}/][%d] projectUpdatesProjectUpdatesReadOK ", 200)
}

func (o *ProjectUpdatesProjectUpdatesReadOK) String() string {
	return fmt.Sprintf("[GET /api/v2/project_updates/{id}/][%d] projectUpdatesProjectUpdatesReadOK ", 200)
}

func (o *ProjectUpdatesProjectUpdatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}