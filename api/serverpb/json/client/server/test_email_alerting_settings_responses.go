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

// TestEmailAlertingSettingsReader is a Reader for the TestEmailAlertingSettings structure.
type TestEmailAlertingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestEmailAlertingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestEmailAlertingSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTestEmailAlertingSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTestEmailAlertingSettingsOK creates a TestEmailAlertingSettingsOK with default headers values
func NewTestEmailAlertingSettingsOK() *TestEmailAlertingSettingsOK {
	return &TestEmailAlertingSettingsOK{}
}

/*TestEmailAlertingSettingsOK handles this case with default header values.

A successful response.
*/
type TestEmailAlertingSettingsOK struct {
	Payload interface{}
}

func (o *TestEmailAlertingSettingsOK) Error() string {
	return fmt.Sprintf("[POST /v1/Settings/TestEmailAlertingSettings][%d] testEmailAlertingSettingsOk  %+v", 200, o.Payload)
}

func (o *TestEmailAlertingSettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *TestEmailAlertingSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestEmailAlertingSettingsDefault creates a TestEmailAlertingSettingsDefault with default headers values
func NewTestEmailAlertingSettingsDefault(code int) *TestEmailAlertingSettingsDefault {
	return &TestEmailAlertingSettingsDefault{
		_statusCode: code,
	}
}

/*TestEmailAlertingSettingsDefault handles this case with default header values.

An unexpected error response.
*/
type TestEmailAlertingSettingsDefault struct {
	_statusCode int

	Payload *TestEmailAlertingSettingsDefaultBody
}

// Code gets the status code for the test email alerting settings default response
func (o *TestEmailAlertingSettingsDefault) Code() int {
	return o._statusCode
}

func (o *TestEmailAlertingSettingsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/Settings/TestEmailAlertingSettings][%d] TestEmailAlertingSettings default  %+v", o._statusCode, o.Payload)
}

func (o *TestEmailAlertingSettingsDefault) GetPayload() *TestEmailAlertingSettingsDefaultBody {
	return o.Payload
}

func (o *TestEmailAlertingSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TestEmailAlertingSettingsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TestEmailAlertingSettingsBody test email alerting settings body
swagger:model TestEmailAlertingSettingsBody
*/
type TestEmailAlertingSettingsBody struct {

	// Target email address to send the email to.
	EmailTo string `json:"emailTo,omitempty"`

	// email alerting settings
	EmailAlertingSettings *TestEmailAlertingSettingsParamsBodyEmailAlertingSettings `json:"emailAlertingSettings,omitempty"`
}

// Validate validates this test email alerting settings body
func (o *TestEmailAlertingSettingsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmailAlertingSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TestEmailAlertingSettingsBody) validateEmailAlertingSettings(formats strfmt.Registry) error {

	if swag.IsZero(o.EmailAlertingSettings) { // not required
		return nil
	}

	if o.EmailAlertingSettings != nil {
		if err := o.EmailAlertingSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "emailAlertingSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TestEmailAlertingSettingsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TestEmailAlertingSettingsBody) UnmarshalBinary(b []byte) error {
	var res TestEmailAlertingSettingsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TestEmailAlertingSettingsDefaultBody test email alerting settings default body
swagger:model TestEmailAlertingSettingsDefaultBody
*/
type TestEmailAlertingSettingsDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this test email alerting settings default body
func (o *TestEmailAlertingSettingsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TestEmailAlertingSettingsDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("TestEmailAlertingSettings default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *TestEmailAlertingSettingsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TestEmailAlertingSettingsDefaultBody) UnmarshalBinary(b []byte) error {
	var res TestEmailAlertingSettingsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TestEmailAlertingSettingsParamsBodyEmailAlertingSettings EmailAlertingSettings represents email (SMTP) configuration for Alerting.
swagger:model TestEmailAlertingSettingsParamsBodyEmailAlertingSettings
*/
type TestEmailAlertingSettingsParamsBodyEmailAlertingSettings struct {

	// SMTP From header field.
	From string `json:"from,omitempty"`

	// SMTP host and port.
	Smarthost string `json:"smarthost,omitempty"`

	// Hostname to identify to the SMTP server.
	Hello string `json:"hello,omitempty"`

	// Auth using CRAM-MD5, LOGIN and PLAIN.
	Username string `json:"username,omitempty"`

	// Auth using LOGIN and PLAIN.
	Password string `json:"password,omitempty"`

	// Auth using PLAIN.
	Identity string `json:"identity,omitempty"`

	// Auth using CRAM-MD5.
	Secret string `json:"secret,omitempty"`

	// Require TLS.
	RequireTLS bool `json:"requireTls,omitempty"`
}

// Validate validates this test email alerting settings params body email alerting settings
func (o *TestEmailAlertingSettingsParamsBodyEmailAlertingSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TestEmailAlertingSettingsParamsBodyEmailAlertingSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TestEmailAlertingSettingsParamsBodyEmailAlertingSettings) UnmarshalBinary(b []byte) error {
	var res TestEmailAlertingSettingsParamsBodyEmailAlertingSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
