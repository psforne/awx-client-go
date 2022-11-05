// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SettingsSettingsPartialUpdateReader is a Reader for the SettingsSettingsPartialUpdate structure.
type SettingsSettingsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SettingsSettingsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSettingsSettingsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSettingsSettingsPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSettingsSettingsPartialUpdateOK creates a SettingsSettingsPartialUpdateOK with default headers values
func NewSettingsSettingsPartialUpdateOK() *SettingsSettingsPartialUpdateOK {
	return &SettingsSettingsPartialUpdateOK{}
}

/*
SettingsSettingsPartialUpdateOK describes a response with status code 200, with default header values.

SettingsSettingsPartialUpdateOK settings settings partial update o k
*/
type SettingsSettingsPartialUpdateOK struct {
}

// IsSuccess returns true when this settings settings partial update o k response has a 2xx status code
func (o *SettingsSettingsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this settings settings partial update o k response has a 3xx status code
func (o *SettingsSettingsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this settings settings partial update o k response has a 4xx status code
func (o *SettingsSettingsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this settings settings partial update o k response has a 5xx status code
func (o *SettingsSettingsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this settings settings partial update o k response a status code equal to that given
func (o *SettingsSettingsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *SettingsSettingsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/settings/{category_slug}/][%d] settingsSettingsPartialUpdateOK ", 200)
}

func (o *SettingsSettingsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/settings/{category_slug}/][%d] settingsSettingsPartialUpdateOK ", 200)
}

func (o *SettingsSettingsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSettingsSettingsPartialUpdateBadRequest creates a SettingsSettingsPartialUpdateBadRequest with default headers values
func NewSettingsSettingsPartialUpdateBadRequest() *SettingsSettingsPartialUpdateBadRequest {
	return &SettingsSettingsPartialUpdateBadRequest{}
}

/*
SettingsSettingsPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SettingsSettingsPartialUpdateBadRequest struct {
}

// IsSuccess returns true when this settings settings partial update bad request response has a 2xx status code
func (o *SettingsSettingsPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this settings settings partial update bad request response has a 3xx status code
func (o *SettingsSettingsPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this settings settings partial update bad request response has a 4xx status code
func (o *SettingsSettingsPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this settings settings partial update bad request response has a 5xx status code
func (o *SettingsSettingsPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this settings settings partial update bad request response a status code equal to that given
func (o *SettingsSettingsPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SettingsSettingsPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/settings/{category_slug}/][%d] settingsSettingsPartialUpdateBadRequest ", 400)
}

func (o *SettingsSettingsPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/settings/{category_slug}/][%d] settingsSettingsPartialUpdateBadRequest ", 400)
}

func (o *SettingsSettingsPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}