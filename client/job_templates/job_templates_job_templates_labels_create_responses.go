// Code generated by go-swagger; DO NOT EDIT.

package job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JobTemplatesJobTemplatesLabelsCreateReader is a Reader for the JobTemplatesJobTemplatesLabelsCreate structure.
type JobTemplatesJobTemplatesLabelsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobTemplatesJobTemplatesLabelsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewJobTemplatesJobTemplatesLabelsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJobTemplatesJobTemplatesLabelsCreateCreated creates a JobTemplatesJobTemplatesLabelsCreateCreated with default headers values
func NewJobTemplatesJobTemplatesLabelsCreateCreated() *JobTemplatesJobTemplatesLabelsCreateCreated {
	return &JobTemplatesJobTemplatesLabelsCreateCreated{}
}

/*
JobTemplatesJobTemplatesLabelsCreateCreated describes a response with status code 201, with default header values.

JobTemplatesJobTemplatesLabelsCreateCreated job templates job templates labels create created
*/
type JobTemplatesJobTemplatesLabelsCreateCreated struct {
}

// IsSuccess returns true when this job templates job templates labels create created response has a 2xx status code
func (o *JobTemplatesJobTemplatesLabelsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this job templates job templates labels create created response has a 3xx status code
func (o *JobTemplatesJobTemplatesLabelsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this job templates job templates labels create created response has a 4xx status code
func (o *JobTemplatesJobTemplatesLabelsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this job templates job templates labels create created response has a 5xx status code
func (o *JobTemplatesJobTemplatesLabelsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this job templates job templates labels create created response a status code equal to that given
func (o *JobTemplatesJobTemplatesLabelsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *JobTemplatesJobTemplatesLabelsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/job_templates/{id}/labels/][%d] jobTemplatesJobTemplatesLabelsCreateCreated ", 201)
}

func (o *JobTemplatesJobTemplatesLabelsCreateCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/job_templates/{id}/labels/][%d] jobTemplatesJobTemplatesLabelsCreateCreated ", 201)
}

func (o *JobTemplatesJobTemplatesLabelsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
JobTemplatesJobTemplatesLabelsCreateBody job templates job templates labels create body
swagger:model JobTemplatesJobTemplatesLabelsCreateBody
*/
type JobTemplatesJobTemplatesLabelsCreateBody struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// Organization this label belongs to.
	// Required: true
	Organization *int64 `json:"organization"`
}

// Validate validates this job templates job templates labels create body
func (o *JobTemplatesJobTemplatesLabelsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *JobTemplatesJobTemplatesLabelsCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *JobTemplatesJobTemplatesLabelsCreateBody) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"organization", "body", o.Organization); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this job templates job templates labels create body based on context it is used
func (o *JobTemplatesJobTemplatesLabelsCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *JobTemplatesJobTemplatesLabelsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *JobTemplatesJobTemplatesLabelsCreateBody) UnmarshalBinary(b []byte) error {
	var res JobTemplatesJobTemplatesLabelsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
