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

// StartMySQLShowIndexActionReader is a Reader for the StartMySQLShowIndexAction structure.
type StartMySQLShowIndexActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartMySQLShowIndexActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartMySQLShowIndexActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartMySQLShowIndexActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartMySQLShowIndexActionOK creates a StartMySQLShowIndexActionOK with default headers values
func NewStartMySQLShowIndexActionOK() *StartMySQLShowIndexActionOK {
	return &StartMySQLShowIndexActionOK{}
}

/*StartMySQLShowIndexActionOK handles this case with default header values.

A successful response.
*/
type StartMySQLShowIndexActionOK struct {
	Payload *StartMySQLShowIndexActionOKBody
}

func (o *StartMySQLShowIndexActionOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartMySQLShowIndex][%d] startMySqlShowIndexActionOk  %+v", 200, o.Payload)
}

func (o *StartMySQLShowIndexActionOK) GetPayload() *StartMySQLShowIndexActionOKBody {
	return o.Payload
}

func (o *StartMySQLShowIndexActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartMySQLShowIndexActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartMySQLShowIndexActionDefault creates a StartMySQLShowIndexActionDefault with default headers values
func NewStartMySQLShowIndexActionDefault(code int) *StartMySQLShowIndexActionDefault {
	return &StartMySQLShowIndexActionDefault{
		_statusCode: code,
	}
}

/*StartMySQLShowIndexActionDefault handles this case with default header values.

An unexpected error response.
*/
type StartMySQLShowIndexActionDefault struct {
	_statusCode int

	Payload *StartMySQLShowIndexActionDefaultBody
}

// Code gets the status code for the start my SQL show index action default response
func (o *StartMySQLShowIndexActionDefault) Code() int {
	return o._statusCode
}

func (o *StartMySQLShowIndexActionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartMySQLShowIndex][%d] StartMySQLShowIndexAction default  %+v", o._statusCode, o.Payload)
}

func (o *StartMySQLShowIndexActionDefault) GetPayload() *StartMySQLShowIndexActionDefaultBody {
	return o.Payload
}

func (o *StartMySQLShowIndexActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartMySQLShowIndexActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartMySQLShowIndexActionBody start my SQL show index action body
swagger:model StartMySQLShowIndexActionBody
*/
type StartMySQLShowIndexActionBody struct {

	// pmm-agent ID where to run this Action.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service ID for this Action. Required.
	ServiceID string `json:"service_id,omitempty"`

	// Table name. Required. May additionally contain a database name.
	TableName string `json:"table_name,omitempty"`

	// Database name. Required if not given in the table_name field.
	Database string `json:"database,omitempty"`
}

// Validate validates this start my SQL show index action body
func (o *StartMySQLShowIndexActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLShowIndexActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLShowIndexActionBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLShowIndexActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartMySQLShowIndexActionDefaultBody start my SQL show index action default body
swagger:model StartMySQLShowIndexActionDefaultBody
*/
type StartMySQLShowIndexActionDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this start my SQL show index action default body
func (o *StartMySQLShowIndexActionDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartMySQLShowIndexActionDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("StartMySQLShowIndexAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLShowIndexActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLShowIndexActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLShowIndexActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartMySQLShowIndexActionOKBody start my SQL show index action OK body
swagger:model StartMySQLShowIndexActionOKBody
*/
type StartMySQLShowIndexActionOKBody struct {

	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`

	// pmm-agent ID where to this Action was started.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this start my SQL show index action OK body
func (o *StartMySQLShowIndexActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLShowIndexActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLShowIndexActionOKBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLShowIndexActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
