// Code generated by go-swagger; DO NOT EDIT.

package project_updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectUpdatesProjectUpdatesCancelCreateReader is a Reader for the ProjectUpdatesProjectUpdatesCancelCreate structure.
type ProjectUpdatesProjectUpdatesCancelCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectUpdatesProjectUpdatesCancelCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewProjectUpdatesProjectUpdatesCancelCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectUpdatesProjectUpdatesCancelCreateCreated creates a ProjectUpdatesProjectUpdatesCancelCreateCreated with default headers values
func NewProjectUpdatesProjectUpdatesCancelCreateCreated() *ProjectUpdatesProjectUpdatesCancelCreateCreated {
	return &ProjectUpdatesProjectUpdatesCancelCreateCreated{}
}

/*
ProjectUpdatesProjectUpdatesCancelCreateCreated describes a response with status code 201, with default header values.

ProjectUpdatesProjectUpdatesCancelCreateCreated project updates project updates cancel create created
*/
type ProjectUpdatesProjectUpdatesCancelCreateCreated struct {
}

// IsSuccess returns true when this project updates project updates cancel create created response has a 2xx status code
func (o *ProjectUpdatesProjectUpdatesCancelCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project updates project updates cancel create created response has a 3xx status code
func (o *ProjectUpdatesProjectUpdatesCancelCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project updates project updates cancel create created response has a 4xx status code
func (o *ProjectUpdatesProjectUpdatesCancelCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this project updates project updates cancel create created response has a 5xx status code
func (o *ProjectUpdatesProjectUpdatesCancelCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this project updates project updates cancel create created response a status code equal to that given
func (o *ProjectUpdatesProjectUpdatesCancelCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *ProjectUpdatesProjectUpdatesCancelCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/project_updates/{id}/cancel/][%d] projectUpdatesProjectUpdatesCancelCreateCreated ", 201)
}

func (o *ProjectUpdatesProjectUpdatesCancelCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/project_updates/{id}/cancel/][%d] projectUpdatesProjectUpdatesCancelCreateCreated ", 201)
}

func (o *ProjectUpdatesProjectUpdatesCancelCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}