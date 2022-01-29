// Code generated by go-swagger; DO NOT EDIT.

package agent_local

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

// ReloadReader is a Reader for the Reload structure.
type ReloadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReloadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReloadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReloadOK creates a ReloadOK with default headers values
func NewReloadOK() *ReloadOK {
	return &ReloadOK{}
}

/*ReloadOK handles this case with default header values.

A successful response.
*/
type ReloadOK struct {
	Payload interface{}
}

func (o *ReloadOK) Error() string {
	return fmt.Sprintf("[POST /local/Reload][%d] reloadOk  %+v", 200, o.Payload)
}

func (o *ReloadOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ReloadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReloadDefault creates a ReloadDefault with default headers values
func NewReloadDefault(code int) *ReloadDefault {
	return &ReloadDefault{
		_statusCode: code,
	}
}

/*ReloadDefault handles this case with default header values.

An unexpected error response.
*/
type ReloadDefault struct {
	_statusCode int

	Payload *ReloadDefaultBody
}

// Code gets the status code for the reload default response
func (o *ReloadDefault) Code() int {
	return o._statusCode
}

func (o *ReloadDefault) Error() string {
	return fmt.Sprintf("[POST /local/Reload][%d] Reload default  %+v", o._statusCode, o.Payload)
}

func (o *ReloadDefault) GetPayload() *ReloadDefaultBody {
	return o.Payload
}

func (o *ReloadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ReloadDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DetailsItems0 details items0
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// type Url
	TypeURL string `json:"typeUrl,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ReloadDefaultBody reload default body
swagger:model ReloadDefaultBody
*/
type ReloadDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this reload default body
func (o *ReloadDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ReloadDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("Reload default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ReloadDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReloadDefaultBody) UnmarshalBinary(b []byte) error {
	var res ReloadDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
