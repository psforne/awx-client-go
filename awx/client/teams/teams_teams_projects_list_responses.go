// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TeamsTeamsProjectsListReader is a Reader for the TeamsTeamsProjectsList structure.
type TeamsTeamsProjectsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamsTeamsProjectsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTeamsTeamsProjectsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTeamsTeamsProjectsListOK creates a TeamsTeamsProjectsListOK with default headers values
func NewTeamsTeamsProjectsListOK() *TeamsTeamsProjectsListOK {
	return &TeamsTeamsProjectsListOK{}
}

/*
TeamsTeamsProjectsListOK describes a response with status code 200, with default header values.

TeamsTeamsProjectsListOK teams teams projects list o k
*/
type TeamsTeamsProjectsListOK struct {
}

// IsSuccess returns true when this teams teams projects list o k response has a 2xx status code
func (o *TeamsTeamsProjectsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this teams teams projects list o k response has a 3xx status code
func (o *TeamsTeamsProjectsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this teams teams projects list o k response has a 4xx status code
func (o *TeamsTeamsProjectsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this teams teams projects list o k response has a 5xx status code
func (o *TeamsTeamsProjectsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this teams teams projects list o k response a status code equal to that given
func (o *TeamsTeamsProjectsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *TeamsTeamsProjectsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/teams/{id}/projects/][%d] teamsTeamsProjectsListOK ", 200)
}

func (o *TeamsTeamsProjectsListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/teams/{id}/projects/][%d] teamsTeamsProjectsListOK ", 200)
}

func (o *TeamsTeamsProjectsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}