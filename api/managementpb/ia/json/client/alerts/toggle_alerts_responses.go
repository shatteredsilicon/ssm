// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ToggleAlertsReader is a Reader for the ToggleAlerts structure.
type ToggleAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ToggleAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewToggleAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewToggleAlertsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewToggleAlertsOK creates a ToggleAlertsOK with default headers values
func NewToggleAlertsOK() *ToggleAlertsOK {
	return &ToggleAlertsOK{}
}

/*ToggleAlertsOK handles this case with default header values.

A successful response.
*/
type ToggleAlertsOK struct {
	Payload interface{}
}

func (o *ToggleAlertsOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Alerts/Toggle][%d] toggleAlertsOk  %+v", 200, o.Payload)
}

func (o *ToggleAlertsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ToggleAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewToggleAlertsDefault creates a ToggleAlertsDefault with default headers values
func NewToggleAlertsDefault(code int) *ToggleAlertsDefault {
	return &ToggleAlertsDefault{
		_statusCode: code,
	}
}

/*ToggleAlertsDefault handles this case with default header values.

An unexpected error response.
*/
type ToggleAlertsDefault struct {
	_statusCode int

	Payload *ToggleAlertsDefaultBody
}

// Code gets the status code for the toggle alerts default response
func (o *ToggleAlertsDefault) Code() int {
	return o._statusCode
}

func (o *ToggleAlertsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Alerts/Toggle][%d] ToggleAlerts default  %+v", o._statusCode, o.Payload)
}

func (o *ToggleAlertsDefault) GetPayload() *ToggleAlertsDefaultBody {
	return o.Payload
}

func (o *ToggleAlertsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ToggleAlertsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ToggleAlertsBody toggle alerts body
swagger:model ToggleAlertsBody
*/
type ToggleAlertsBody struct {

<<<<<<< HEAD:api/managementpb/ia/json/client/alerts/toggle_alerts_responses.go
	// List of alerts that silence state should be switched. If provided array is empty than all
	// existing alerts are switched.
	AlertIds []string `json:"alert_ids"`
=======
	// alert id
	AlertID string `json:"alert_id,omitempty"`
>>>>>>> d7959adc (PMM-9377 preserve snake case in field names):api/managementpb/ia/json/client/alerts/toggle_alert_responses.go

	// BooleanFlag represent a command to set some boolean property to true,
	// to false, or avoid changing that property.
	//
	//  - DO_NOT_CHANGE: Do not change boolean property. Default value.
	//  - TRUE: True.
	//  - FALSE: False.
	// Enum: [DO_NOT_CHANGE TRUE FALSE]
	Silenced *string `json:"silenced,omitempty"`
}

// Validate validates this toggle alerts body
func (o *ToggleAlertsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSilenced(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var toggleAlertsBodyTypeSilencedPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DO_NOT_CHANGE","TRUE","FALSE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		toggleAlertsBodyTypeSilencedPropEnum = append(toggleAlertsBodyTypeSilencedPropEnum, v)
	}
}

const (

	// ToggleAlertsBodySilencedDONOTCHANGE captures enum value "DO_NOT_CHANGE"
	ToggleAlertsBodySilencedDONOTCHANGE string = "DO_NOT_CHANGE"

	// ToggleAlertsBodySilencedTRUE captures enum value "TRUE"
	ToggleAlertsBodySilencedTRUE string = "TRUE"

	// ToggleAlertsBodySilencedFALSE captures enum value "FALSE"
	ToggleAlertsBodySilencedFALSE string = "FALSE"
)

// prop value enum
func (o *ToggleAlertsBody) validateSilencedEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, toggleAlertsBodyTypeSilencedPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ToggleAlertsBody) validateSilenced(formats strfmt.Registry) error {

	if swag.IsZero(o.Silenced) { // not required
		return nil
	}

	// value enum
	if err := o.validateSilencedEnum("body"+"."+"silenced", "body", *o.Silenced); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ToggleAlertsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ToggleAlertsBody) UnmarshalBinary(b []byte) error {
	var res ToggleAlertsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ToggleAlertsDefaultBody toggle alerts default body
swagger:model ToggleAlertsDefaultBody
*/
type ToggleAlertsDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this toggle alerts default body
func (o *ToggleAlertsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ToggleAlertsDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ToggleAlerts default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ToggleAlertsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ToggleAlertsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ToggleAlertsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
