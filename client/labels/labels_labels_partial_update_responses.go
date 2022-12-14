// Code generated by go-swagger; DO NOT EDIT.

package labels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LabelsLabelsPartialUpdateReader is a Reader for the LabelsLabelsPartialUpdate structure.
type LabelsLabelsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LabelsLabelsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLabelsLabelsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLabelsLabelsPartialUpdateOK creates a LabelsLabelsPartialUpdateOK with default headers values
func NewLabelsLabelsPartialUpdateOK() *LabelsLabelsPartialUpdateOK {
	return &LabelsLabelsPartialUpdateOK{}
}

/*
LabelsLabelsPartialUpdateOK describes a response with status code 200, with default header values.

LabelsLabelsPartialUpdateOK labels labels partial update o k
*/
type LabelsLabelsPartialUpdateOK struct {
}

// IsSuccess returns true when this labels labels partial update o k response has a 2xx status code
func (o *LabelsLabelsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this labels labels partial update o k response has a 3xx status code
func (o *LabelsLabelsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this labels labels partial update o k response has a 4xx status code
func (o *LabelsLabelsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this labels labels partial update o k response has a 5xx status code
func (o *LabelsLabelsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this labels labels partial update o k response a status code equal to that given
func (o *LabelsLabelsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *LabelsLabelsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/labels/{id}/][%d] labelsLabelsPartialUpdateOK ", 200)
}

func (o *LabelsLabelsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/labels/{id}/][%d] labelsLabelsPartialUpdateOK ", 200)
}

func (o *LabelsLabelsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
LabelsLabelsPartialUpdateBody labels labels partial update body
swagger:model LabelsLabelsPartialUpdateBody
*/
type LabelsLabelsPartialUpdateBody struct {

	// name
	Name string `json:"name,omitempty"`

	// Organization this label belongs to.
	Organization int64 `json:"organization,omitempty"`
}

// Validate validates this labels labels partial update body
func (o *LabelsLabelsPartialUpdateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this labels labels partial update body based on context it is used
func (o *LabelsLabelsPartialUpdateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LabelsLabelsPartialUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LabelsLabelsPartialUpdateBody) UnmarshalBinary(b []byte) error {
	var res LabelsLabelsPartialUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
