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

// StartPTMySQLSummaryActionReader is a Reader for the StartPTMySQLSummaryAction structure.
type StartPTMySQLSummaryActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartPTMySQLSummaryActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartPTMySQLSummaryActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartPTMySQLSummaryActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartPTMySQLSummaryActionOK creates a StartPTMySQLSummaryActionOK with default headers values
func NewStartPTMySQLSummaryActionOK() *StartPTMySQLSummaryActionOK {
	return &StartPTMySQLSummaryActionOK{}
}

/*StartPTMySQLSummaryActionOK handles this case with default header values.

A successful response.
*/
type StartPTMySQLSummaryActionOK struct {
	Payload *StartPTMySQLSummaryActionOKBody
}

func (o *StartPTMySQLSummaryActionOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartPTMySQLSummary][%d] startPtMySqlSummaryActionOk  %+v", 200, o.Payload)
}

func (o *StartPTMySQLSummaryActionOK) GetPayload() *StartPTMySQLSummaryActionOKBody {
	return o.Payload
}

func (o *StartPTMySQLSummaryActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartPTMySQLSummaryActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartPTMySQLSummaryActionDefault creates a StartPTMySQLSummaryActionDefault with default headers values
func NewStartPTMySQLSummaryActionDefault(code int) *StartPTMySQLSummaryActionDefault {
	return &StartPTMySQLSummaryActionDefault{
		_statusCode: code,
	}
}

/*StartPTMySQLSummaryActionDefault handles this case with default header values.

An unexpected error response.
*/
type StartPTMySQLSummaryActionDefault struct {
	_statusCode int

	Payload *StartPTMySQLSummaryActionDefaultBody
}

// Code gets the status code for the start PT my SQL summary action default response
func (o *StartPTMySQLSummaryActionDefault) Code() int {
	return o._statusCode
}

func (o *StartPTMySQLSummaryActionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartPTMySQLSummary][%d] StartPTMySQLSummaryAction default  %+v", o._statusCode, o.Payload)
}

func (o *StartPTMySQLSummaryActionDefault) GetPayload() *StartPTMySQLSummaryActionDefaultBody {
	return o.Payload
}

func (o *StartPTMySQLSummaryActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartPTMySQLSummaryActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartPTMySQLSummaryActionBody Message to prepare pt-mysql-summary data
swagger:model StartPTMySQLSummaryActionBody
*/
type StartPTMySQLSummaryActionBody struct {

	// pmm-agent ID where to run this Action.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service ID for this Action.
	ServiceID string `json:"service_id,omitempty"`
}

// Validate validates this start PT my SQL summary action body
func (o *StartPTMySQLSummaryActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPTMySQLSummaryActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTMySQLSummaryActionBody) UnmarshalBinary(b []byte) error {
	var res StartPTMySQLSummaryActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartPTMySQLSummaryActionDefaultBody start PT my SQL summary action default body
swagger:model StartPTMySQLSummaryActionDefaultBody
*/
type StartPTMySQLSummaryActionDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this start PT my SQL summary action default body
func (o *StartPTMySQLSummaryActionDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartPTMySQLSummaryActionDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("StartPTMySQLSummaryAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *StartPTMySQLSummaryActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTMySQLSummaryActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartPTMySQLSummaryActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartPTMySQLSummaryActionOKBody Message to retrieve the prepared pt-mysql-summary data
swagger:model StartPTMySQLSummaryActionOKBody
*/
type StartPTMySQLSummaryActionOKBody struct {

	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`

	// pmm-agent ID where to this Action was started.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this start PT my SQL summary action OK body
func (o *StartPTMySQLSummaryActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPTMySQLSummaryActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPTMySQLSummaryActionOKBody) UnmarshalBinary(b []byte) error {
	var res StartPTMySQLSummaryActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
