// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TeamsTeamsDeleteReader is a Reader for the TeamsTeamsDelete structure.
type TeamsTeamsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamsTeamsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTeamsTeamsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTeamsTeamsDeleteNoContent creates a TeamsTeamsDeleteNoContent with default headers values
func NewTeamsTeamsDeleteNoContent() *TeamsTeamsDeleteNoContent {
	return &TeamsTeamsDeleteNoContent{}
}

/*
TeamsTeamsDeleteNoContent describes a response with status code 204, with default header values.

TeamsTeamsDeleteNoContent teams teams delete no content
*/
type TeamsTeamsDeleteNoContent struct {
}

// IsSuccess returns true when this teams teams delete no content response has a 2xx status code
func (o *TeamsTeamsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this teams teams delete no content response has a 3xx status code
func (o *TeamsTeamsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this teams teams delete no content response has a 4xx status code
func (o *TeamsTeamsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this teams teams delete no content response has a 5xx status code
func (o *TeamsTeamsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this teams teams delete no content response a status code equal to that given
func (o *TeamsTeamsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *TeamsTeamsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{id}/][%d] teamsTeamsDeleteNoContent ", 204)
}

func (o *TeamsTeamsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/teams/{id}/][%d] teamsTeamsDeleteNoContent ", 204)
}

func (o *TeamsTeamsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
