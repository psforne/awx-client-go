// Code generated by go-swagger; DO NOT EDIT.

package project_updates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProjectUpdatesProjectUpdatesEventsListReader is a Reader for the ProjectUpdatesProjectUpdatesEventsList structure.
type ProjectUpdatesProjectUpdatesEventsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectUpdatesProjectUpdatesEventsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectUpdatesProjectUpdatesEventsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectUpdatesProjectUpdatesEventsListOK creates a ProjectUpdatesProjectUpdatesEventsListOK with default headers values
func NewProjectUpdatesProjectUpdatesEventsListOK() *ProjectUpdatesProjectUpdatesEventsListOK {
	return &ProjectUpdatesProjectUpdatesEventsListOK{}
}

/*
ProjectUpdatesProjectUpdatesEventsListOK describes a response with status code 200, with default header values.

ProjectUpdatesProjectUpdatesEventsListOK project updates project updates events list o k
*/
type ProjectUpdatesProjectUpdatesEventsListOK struct {
}

// IsSuccess returns true when this project updates project updates events list o k response has a 2xx status code
func (o *ProjectUpdatesProjectUpdatesEventsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project updates project updates events list o k response has a 3xx status code
func (o *ProjectUpdatesProjectUpdatesEventsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project updates project updates events list o k response has a 4xx status code
func (o *ProjectUpdatesProjectUpdatesEventsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project updates project updates events list o k response has a 5xx status code
func (o *ProjectUpdatesProjectUpdatesEventsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project updates project updates events list o k response a status code equal to that given
func (o *ProjectUpdatesProjectUpdatesEventsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectUpdatesProjectUpdatesEventsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/project_updates/{id}/events/][%d] projectUpdatesProjectUpdatesEventsListOK ", 200)
}

func (o *ProjectUpdatesProjectUpdatesEventsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/project_updates/{id}/events/][%d] projectUpdatesProjectUpdatesEventsListOK ", 200)
}

func (o *ProjectUpdatesProjectUpdatesEventsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}