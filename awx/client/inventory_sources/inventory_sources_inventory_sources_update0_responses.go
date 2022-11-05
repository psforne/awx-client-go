// Code generated by go-swagger; DO NOT EDIT.

package inventory_sources

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

// InventorySourcesInventorySourcesUpdate0Reader is a Reader for the InventorySourcesInventorySourcesUpdate0 structure.
type InventorySourcesInventorySourcesUpdate0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventorySourcesInventorySourcesUpdate0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventorySourcesInventorySourcesUpdate0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInventorySourcesInventorySourcesUpdate0OK creates a InventorySourcesInventorySourcesUpdate0OK with default headers values
func NewInventorySourcesInventorySourcesUpdate0OK() *InventorySourcesInventorySourcesUpdate0OK {
	return &InventorySourcesInventorySourcesUpdate0OK{}
}

/*
InventorySourcesInventorySourcesUpdate0OK describes a response with status code 200, with default header values.

InventorySourcesInventorySourcesUpdate0OK inventory sources inventory sources update0 o k
*/
type InventorySourcesInventorySourcesUpdate0OK struct {
}

// IsSuccess returns true when this inventory sources inventory sources update0 o k response has a 2xx status code
func (o *InventorySourcesInventorySourcesUpdate0OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventory sources inventory sources update0 o k response has a 3xx status code
func (o *InventorySourcesInventorySourcesUpdate0OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventory sources inventory sources update0 o k response has a 4xx status code
func (o *InventorySourcesInventorySourcesUpdate0OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventory sources inventory sources update0 o k response has a 5xx status code
func (o *InventorySourcesInventorySourcesUpdate0OK) IsServerError() bool {
	return false
}

// IsCode returns true when this inventory sources inventory sources update0 o k response a status code equal to that given
func (o *InventorySourcesInventorySourcesUpdate0OK) IsCode(code int) bool {
	return code == 200
}

func (o *InventorySourcesInventorySourcesUpdate0OK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/inventory_sources/{id}/][%d] inventorySourcesInventorySourcesUpdate0OK ", 200)
}

func (o *InventorySourcesInventorySourcesUpdate0OK) String() string {
	return fmt.Sprintf("[PUT /api/v2/inventory_sources/{id}/][%d] inventorySourcesInventorySourcesUpdate0OK ", 200)
}

func (o *InventorySourcesInventorySourcesUpdate0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
InventorySourcesInventorySourcesUpdate0Body inventory sources inventory sources update0 body
swagger:model InventorySourcesInventorySourcesUpdate0Body
*/
type InventorySourcesInventorySourcesUpdate0Body struct {

	// Cloud credential to use for inventory updates.
	Credential int64 `json:"credential,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var="status.power_state"and enabled_value="powered_on" with host variables:{   "status": {     "power_state": "powered_on",     "created": "2018-02-01T08:00:00.000000Z:00",     "healthy": true    },    "name": "foobar",    "ip_address": "192.168.2.1"}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled
	EnabledValue string `json:"enabled_value,omitempty"`

	// Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as "foo.bar", in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get("foo", {}).get("bar", default)
	EnabledVar string `json:"enabled_var,omitempty"`

	// The container image to be used for execution.
	ExecutionEnvironment int64 `json:"execution_environment,omitempty"`

	// Regex where only matching hosts will be imported.
	HostFilter string `json:"host_filter,omitempty"`

	// inventory
	// Required: true
	Inventory *int64 `json:"inventory"`

	// name
	// Required: true
	Name *string `json:"name"`

	// Overwrite local groups and hosts from remote inventory source.
	Overwrite bool `json:"overwrite,omitempty"`

	// Overwrite local variables from remote inventory source.
	OverwriteVars bool `json:"overwrite_vars,omitempty"`

	// source
	Source string `json:"source,omitempty"`

	// source path
	SourcePath string `json:"source_path,omitempty"`

	// Project containing inventory file used as source.
	SourceProject string `json:"source_project,omitempty"`

	// Inventory source variables in YAML or JSON format.
	SourceVars string `json:"source_vars,omitempty"`

	// The amount of time (in seconds) to run before the task is canceled.
	Timeout int64 `json:"timeout,omitempty"`

	// update cache timeout
	UpdateCacheTimeout int64 `json:"update_cache_timeout,omitempty"`

	// update on launch
	UpdateOnLaunch bool `json:"update_on_launch,omitempty"`

	// update on project update
	UpdateOnProjectUpdate bool `json:"update_on_project_update,omitempty"`

	// verbosity
	Verbosity string `json:"verbosity,omitempty"`
}

// Validate validates this inventory sources inventory sources update0 body
func (o *InventorySourcesInventorySourcesUpdate0Body) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInventory(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InventorySourcesInventorySourcesUpdate0Body) validateInventory(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"inventory", "body", o.Inventory); err != nil {
		return err
	}

	return nil
}

func (o *InventorySourcesInventorySourcesUpdate0Body) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this inventory sources inventory sources update0 body based on context it is used
func (o *InventorySourcesInventorySourcesUpdate0Body) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InventorySourcesInventorySourcesUpdate0Body) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InventorySourcesInventorySourcesUpdate0Body) UnmarshalBinary(b []byte) error {
	var res InventorySourcesInventorySourcesUpdate0Body
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}