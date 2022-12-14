// Code generated by go-swagger; DO NOT EDIT.

package schedules

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

// SchedulesSchedulesUpdateReader is a Reader for the SchedulesSchedulesUpdate structure.
type SchedulesSchedulesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchedulesSchedulesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchedulesSchedulesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSchedulesSchedulesUpdateOK creates a SchedulesSchedulesUpdateOK with default headers values
func NewSchedulesSchedulesUpdateOK() *SchedulesSchedulesUpdateOK {
	return &SchedulesSchedulesUpdateOK{}
}

/*
SchedulesSchedulesUpdateOK describes a response with status code 200, with default header values.

SchedulesSchedulesUpdateOK schedules schedules update o k
*/
type SchedulesSchedulesUpdateOK struct {
}

// IsSuccess returns true when this schedules schedules update o k response has a 2xx status code
func (o *SchedulesSchedulesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schedules schedules update o k response has a 3xx status code
func (o *SchedulesSchedulesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedules schedules update o k response has a 4xx status code
func (o *SchedulesSchedulesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this schedules schedules update o k response has a 5xx status code
func (o *SchedulesSchedulesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this schedules schedules update o k response a status code equal to that given
func (o *SchedulesSchedulesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *SchedulesSchedulesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/schedules/{id}/][%d] schedulesSchedulesUpdateOK ", 200)
}

func (o *SchedulesSchedulesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/schedules/{id}/][%d] schedulesSchedulesUpdateOK ", 200)
}

func (o *SchedulesSchedulesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
SchedulesSchedulesUpdateBody schedules schedules update body
swagger:model SchedulesSchedulesUpdateBody
*/
type SchedulesSchedulesUpdateBody struct {

	// description
	Description string `json:"description,omitempty"`

	// diff mode
	DiffMode bool `json:"diff_mode,omitempty"`

	// Enables processing of this schedule.
	Enabled bool `json:"enabled,omitempty"`

	// extra data
	ExtraData string `json:"extra_data,omitempty"`

	// Inventory applied as a prompt, assuming job template prompts for inventory
	Inventory int64 `json:"inventory,omitempty"`

	// job tags
	JobTags string `json:"job_tags,omitempty"`

	// job type
	JobType string `json:"job_type,omitempty"`

	// limit
	Limit string `json:"limit,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// A value representing the schedules iCal recurrence rule.
	// Required: true
	Rrule *string `json:"rrule"`

	// scm branch
	ScmBranch string `json:"scm_branch,omitempty"`

	// skip tags
	SkipTags string `json:"skip_tags,omitempty"`

	// unified job template
	// Required: true
	UnifiedJobTemplate *int64 `json:"unified_job_template"`

	// verbosity
	Verbosity string `json:"verbosity,omitempty"`
}

// Validate validates this schedules schedules update body
func (o *SchedulesSchedulesUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRrule(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUnifiedJobTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SchedulesSchedulesUpdateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *SchedulesSchedulesUpdateBody) validateRrule(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"rrule", "body", o.Rrule); err != nil {
		return err
	}

	return nil
}

func (o *SchedulesSchedulesUpdateBody) validateUnifiedJobTemplate(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"unified_job_template", "body", o.UnifiedJobTemplate); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this schedules schedules update body based on context it is used
func (o *SchedulesSchedulesUpdateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SchedulesSchedulesUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SchedulesSchedulesUpdateBody) UnmarshalBinary(b []byte) error {
	var res SchedulesSchedulesUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
