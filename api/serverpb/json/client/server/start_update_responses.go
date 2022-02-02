// Code generated by go-swagger; DO NOT EDIT.

package server

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

// StartUpdateReader is a Reader for the StartUpdate structure.
type StartUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartUpdateOK creates a StartUpdateOK with default headers values
func NewStartUpdateOK() *StartUpdateOK {
	return &StartUpdateOK{}
}

/*StartUpdateOK handles this case with default header values.

A successful response.
*/
type StartUpdateOK struct {
	Payload *StartUpdateOKBody
}

func (o *StartUpdateOK) Error() string {
	return fmt.Sprintf("[POST /v1/Updates/Start][%d] startUpdateOk  %+v", 200, o.Payload)
}

func (o *StartUpdateOK) GetPayload() *StartUpdateOKBody {
	return o.Payload
}

func (o *StartUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartUpdateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartUpdateDefault creates a StartUpdateDefault with default headers values
func NewStartUpdateDefault(code int) *StartUpdateDefault {
	return &StartUpdateDefault{
		_statusCode: code,
	}
}

/*StartUpdateDefault handles this case with default header values.

An unexpected error response.
*/
type StartUpdateDefault struct {
	_statusCode int

	Payload *StartUpdateDefaultBody
}

// Code gets the status code for the start update default response
func (o *StartUpdateDefault) Code() int {
	return o._statusCode
}

func (o *StartUpdateDefault) Error() string {
	return fmt.Sprintf("[POST /v1/Updates/Start][%d] StartUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *StartUpdateDefault) GetPayload() *StartUpdateDefaultBody {
	return o.Payload
}

func (o *StartUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartUpdateDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartUpdateDefaultBody start update default body
swagger:model StartUpdateDefaultBody
*/
type StartUpdateDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this start update default body
func (o *StartUpdateDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartUpdateDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("StartUpdate default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *StartUpdateDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartUpdateDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartUpdateDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartUpdateOKBody start update OK body
swagger:model StartUpdateOKBody
*/
type StartUpdateOKBody struct {

	// Authentication token for getting update statuses.
	AuthToken string `json:"auth_token,omitempty"`

	// Progress log offset.
	LogOffset int64 `json:"log_offset,omitempty"`
}

// Validate validates this start update OK body
func (o *StartUpdateOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartUpdateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartUpdateOKBody) UnmarshalBinary(b []byte) error {
	var res StartUpdateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
