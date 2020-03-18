// Code generated by go-swagger; DO NOT EDIT.

package engines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ory/keto/internal/httpclient/models"
)

// GetOryAccessControlPolicyReader is a Reader for the GetOryAccessControlPolicy structure.
type GetOryAccessControlPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOryAccessControlPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOryAccessControlPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetOryAccessControlPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOryAccessControlPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOryAccessControlPolicyOK creates a GetOryAccessControlPolicyOK with default headers values
func NewGetOryAccessControlPolicyOK() *GetOryAccessControlPolicyOK {
	return &GetOryAccessControlPolicyOK{}
}

/*GetOryAccessControlPolicyOK handles this case with default header values.

oryAccessControlPolicy
*/
type GetOryAccessControlPolicyOK struct {
	Payload *models.OryAccessControlPolicy
}

func (o *GetOryAccessControlPolicyOK) Error() string {
	return fmt.Sprintf("[GET /engines/acp/ory/{flavor}/policies/{id}][%d] getOryAccessControlPolicyOK  %+v", 200, o.Payload)
}

func (o *GetOryAccessControlPolicyOK) GetPayload() *models.OryAccessControlPolicy {
	return o.Payload
}

func (o *GetOryAccessControlPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OryAccessControlPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOryAccessControlPolicyNotFound creates a GetOryAccessControlPolicyNotFound with default headers values
func NewGetOryAccessControlPolicyNotFound() *GetOryAccessControlPolicyNotFound {
	return &GetOryAccessControlPolicyNotFound{}
}

/*GetOryAccessControlPolicyNotFound handles this case with default header values.

The standard error format
*/
type GetOryAccessControlPolicyNotFound struct {
	Payload *GetOryAccessControlPolicyNotFoundBody
}

func (o *GetOryAccessControlPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /engines/acp/ory/{flavor}/policies/{id}][%d] getOryAccessControlPolicyNotFound  %+v", 404, o.Payload)
}

func (o *GetOryAccessControlPolicyNotFound) GetPayload() *GetOryAccessControlPolicyNotFoundBody {
	return o.Payload
}

func (o *GetOryAccessControlPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOryAccessControlPolicyNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOryAccessControlPolicyInternalServerError creates a GetOryAccessControlPolicyInternalServerError with default headers values
func NewGetOryAccessControlPolicyInternalServerError() *GetOryAccessControlPolicyInternalServerError {
	return &GetOryAccessControlPolicyInternalServerError{}
}

/*GetOryAccessControlPolicyInternalServerError handles this case with default header values.

The standard error format
*/
type GetOryAccessControlPolicyInternalServerError struct {
	Payload *GetOryAccessControlPolicyInternalServerErrorBody
}

func (o *GetOryAccessControlPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /engines/acp/ory/{flavor}/policies/{id}][%d] getOryAccessControlPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOryAccessControlPolicyInternalServerError) GetPayload() *GetOryAccessControlPolicyInternalServerErrorBody {
	return o.Payload
}

func (o *GetOryAccessControlPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOryAccessControlPolicyInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOryAccessControlPolicyInternalServerErrorBody get ory access control policy internal server error body
swagger:model GetOryAccessControlPolicyInternalServerErrorBody
*/
type GetOryAccessControlPolicyInternalServerErrorBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []map[string]interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this get ory access control policy internal server error body
func (o *GetOryAccessControlPolicyInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOryAccessControlPolicyInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOryAccessControlPolicyInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetOryAccessControlPolicyInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetOryAccessControlPolicyNotFoundBody get ory access control policy not found body
swagger:model GetOryAccessControlPolicyNotFoundBody
*/
type GetOryAccessControlPolicyNotFoundBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []map[string]interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this get ory access control policy not found body
func (o *GetOryAccessControlPolicyNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOryAccessControlPolicyNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOryAccessControlPolicyNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetOryAccessControlPolicyNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
