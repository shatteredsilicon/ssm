// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StartMongoDBExplainActionReader is a Reader for the StartMongoDBExplainAction structure.
type StartMongoDBExplainActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartMongoDBExplainActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartMongoDBExplainActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartMongoDBExplainActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartMongoDBExplainActionOK creates a StartMongoDBExplainActionOK with default headers values
func NewStartMongoDBExplainActionOK() *StartMongoDBExplainActionOK {
	return &StartMongoDBExplainActionOK{}
}

/*StartMongoDBExplainActionOK handles this case with default header values.

A successful response.
*/
type StartMongoDBExplainActionOK struct {
	Payload *StartMongoDBExplainActionOKBody
}

func (o *StartMongoDBExplainActionOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartMongoDBExplain][%d] startMongoDbExplainActionOk  %+v", 200, o.Payload)
}

func (o *StartMongoDBExplainActionOK) GetPayload() *StartMongoDBExplainActionOKBody {
	return o.Payload
}

func (o *StartMongoDBExplainActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartMongoDBExplainActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartMongoDBExplainActionDefault creates a StartMongoDBExplainActionDefault with default headers values
func NewStartMongoDBExplainActionDefault(code int) *StartMongoDBExplainActionDefault {
	return &StartMongoDBExplainActionDefault{
		_statusCode: code,
	}
}

/*StartMongoDBExplainActionDefault handles this case with default header values.

An unexpected error response.
*/
type StartMongoDBExplainActionDefault struct {
	_statusCode int

	Payload *StartMongoDBExplainActionDefaultBody
}

// Code gets the status code for the start mongo DB explain action default response
func (o *StartMongoDBExplainActionDefault) Code() int {
	return o._statusCode
}

func (o *StartMongoDBExplainActionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartMongoDBExplain][%d] StartMongoDBExplainAction default  %+v", o._statusCode, o.Payload)
}

func (o *StartMongoDBExplainActionDefault) GetPayload() *StartMongoDBExplainActionDefaultBody {
	return o.Payload
}

func (o *StartMongoDBExplainActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartMongoDBExplainActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartMongoDBExplainActionBody start mongo DB explain action body
swagger:model StartMongoDBExplainActionBody
*/
type StartMongoDBExplainActionBody struct {

	// pmm-agent ID where to run this Action.
	PMMAgentID string `json:"pmmAgentId,omitempty"`

	// Service ID for this Action. Required.
	ServiceID string `json:"serviceId,omitempty"`

	// Query. Required.
	Query string `json:"query,omitempty"`
}

// Validate validates this start mongo DB explain action body
func (o *StartMongoDBExplainActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMongoDBExplainActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMongoDBExplainActionBody) UnmarshalBinary(b []byte) error {
	var res StartMongoDBExplainActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartMongoDBExplainActionDefaultBody start mongo DB explain action default body
swagger:model StartMongoDBExplainActionDefaultBody
*/
type StartMongoDBExplainActionDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this start mongo DB explain action default body
func (o *StartMongoDBExplainActionDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartMongoDBExplainActionDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StartMongoDBExplainAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *StartMongoDBExplainActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMongoDBExplainActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartMongoDBExplainActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartMongoDBExplainActionOKBody start mongo DB explain action OK body
swagger:model StartMongoDBExplainActionOKBody
*/
type StartMongoDBExplainActionOKBody struct {

	// Unique Action ID.
	ActionID string `json:"actionId,omitempty"`

	// pmm-agent ID where to this Action was started.
	PMMAgentID string `json:"pmmAgentId,omitempty"`
}

// Validate validates this start mongo DB explain action OK body
func (o *StartMongoDBExplainActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMongoDBExplainActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMongoDBExplainActionOKBody) UnmarshalBinary(b []byte) error {
	var res StartMongoDBExplainActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
