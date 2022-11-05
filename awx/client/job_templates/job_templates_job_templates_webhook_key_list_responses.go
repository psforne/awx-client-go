// Code generated by go-swagger; DO NOT EDIT.

package job_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// JobTemplatesJobTemplatesWebhookKeyListReader is a Reader for the JobTemplatesJobTemplatesWebhookKeyList structure.
type JobTemplatesJobTemplatesWebhookKeyListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobTemplatesJobTemplatesWebhookKeyListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobTemplatesJobTemplatesWebhookKeyListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewJobTemplatesJobTemplatesWebhookKeyListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJobTemplatesJobTemplatesWebhookKeyListOK creates a JobTemplatesJobTemplatesWebhookKeyListOK with default headers values
func NewJobTemplatesJobTemplatesWebhookKeyListOK() *JobTemplatesJobTemplatesWebhookKeyListOK {
	return &JobTemplatesJobTemplatesWebhookKeyListOK{}
}

/*
JobTemplatesJobTemplatesWebhookKeyListOK describes a response with status code 200, with default header values.

JobTemplatesJobTemplatesWebhookKeyListOK job templates job templates webhook key list o k
*/
type JobTemplatesJobTemplatesWebhookKeyListOK struct {
}

// IsSuccess returns true when this job templates job templates webhook key list o k response has a 2xx status code
func (o *JobTemplatesJobTemplatesWebhookKeyListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this job templates job templates webhook key list o k response has a 3xx status code
func (o *JobTemplatesJobTemplatesWebhookKeyListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this job templates job templates webhook key list o k response has a 4xx status code
func (o *JobTemplatesJobTemplatesWebhookKeyListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this job templates job templates webhook key list o k response has a 5xx status code
func (o *JobTemplatesJobTemplatesWebhookKeyListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this job templates job templates webhook key list o k response a status code equal to that given
func (o *JobTemplatesJobTemplatesWebhookKeyListOK) IsCode(code int) bool {
	return code == 200
}

func (o *JobTemplatesJobTemplatesWebhookKeyListOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/job_templates/{id}/webhook_key/][%d] jobTemplatesJobTemplatesWebhookKeyListOK ", 200)
}

func (o *JobTemplatesJobTemplatesWebhookKeyListOK) String() string {
	return fmt.Sprintf("[GET /api/v2/job_templates/{id}/webhook_key/][%d] jobTemplatesJobTemplatesWebhookKeyListOK ", 200)
}

func (o *JobTemplatesJobTemplatesWebhookKeyListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewJobTemplatesJobTemplatesWebhookKeyListForbidden creates a JobTemplatesJobTemplatesWebhookKeyListForbidden with default headers values
func NewJobTemplatesJobTemplatesWebhookKeyListForbidden() *JobTemplatesJobTemplatesWebhookKeyListForbidden {
	return &JobTemplatesJobTemplatesWebhookKeyListForbidden{}
}

/*
JobTemplatesJobTemplatesWebhookKeyListForbidden describes a response with status code 403, with default header values.

forbidden
*/
type JobTemplatesJobTemplatesWebhookKeyListForbidden struct {
}

// IsSuccess returns true when this job templates job templates webhook key list forbidden response has a 2xx status code
func (o *JobTemplatesJobTemplatesWebhookKeyListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this job templates job templates webhook key list forbidden response has a 3xx status code
func (o *JobTemplatesJobTemplatesWebhookKeyListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this job templates job templates webhook key list forbidden response has a 4xx status code
func (o *JobTemplatesJobTemplatesWebhookKeyListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this job templates job templates webhook key list forbidden response has a 5xx status code
func (o *JobTemplatesJobTemplatesWebhookKeyListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this job templates job templates webhook key list forbidden response a status code equal to that given
func (o *JobTemplatesJobTemplatesWebhookKeyListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *JobTemplatesJobTemplatesWebhookKeyListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/job_templates/{id}/webhook_key/][%d] jobTemplatesJobTemplatesWebhookKeyListForbidden ", 403)
}

func (o *JobTemplatesJobTemplatesWebhookKeyListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/job_templates/{id}/webhook_key/][%d] jobTemplatesJobTemplatesWebhookKeyListForbidden ", 403)
}

func (o *JobTemplatesJobTemplatesWebhookKeyListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
