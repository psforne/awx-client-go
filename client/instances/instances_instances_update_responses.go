// Code generated by go-swagger; DO NOT EDIT.

package instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InstancesInstancesUpdateReader is a Reader for the InstancesInstancesUpdate structure.
type InstancesInstancesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstancesInstancesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstancesInstancesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInstancesInstancesUpdateOK creates a InstancesInstancesUpdateOK with default headers values
func NewInstancesInstancesUpdateOK() *InstancesInstancesUpdateOK {
	return &InstancesInstancesUpdateOK{}
}

/*
InstancesInstancesUpdateOK describes a response with status code 200, with default header values.

InstancesInstancesUpdateOK instances instances update o k
*/
type InstancesInstancesUpdateOK struct {
}

// IsSuccess returns true when this instances instances update o k response has a 2xx status code
func (o *InstancesInstancesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this instances instances update o k response has a 3xx status code
func (o *InstancesInstancesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instances instances update o k response has a 4xx status code
func (o *InstancesInstancesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this instances instances update o k response has a 5xx status code
func (o *InstancesInstancesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this instances instances update o k response a status code equal to that given
func (o *InstancesInstancesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *InstancesInstancesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/instances/{id}/][%d] instancesInstancesUpdateOK ", 200)
}

func (o *InstancesInstancesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/instances/{id}/][%d] instancesInstancesUpdateOK ", 200)
}

func (o *InstancesInstancesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
InstancesInstancesUpdateBody instances instances update body
swagger:model InstancesInstancesUpdateBody
*/
type InstancesInstancesUpdateBody struct {

	// capacity adjustment
	CapacityAdjustment float64 `json:"capacity_adjustment,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// managed by policy
	ManagedByPolicy bool `json:"managed_by_policy,omitempty"`
}

// Validate validates this instances instances update body
func (o *InstancesInstancesUpdateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this instances instances update body based on context it is used
func (o *InstancesInstancesUpdateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InstancesInstancesUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InstancesInstancesUpdateBody) UnmarshalBinary(b []byte) error {
	var res InstancesInstancesUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}