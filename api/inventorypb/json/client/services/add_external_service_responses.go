// Code generated by go-swagger; DO NOT EDIT.

package services

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

// AddExternalServiceReader is a Reader for the AddExternalService structure.
type AddExternalServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddExternalServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddExternalServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddExternalServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddExternalServiceOK creates a AddExternalServiceOK with default headers values
func NewAddExternalServiceOK() *AddExternalServiceOK {
	return &AddExternalServiceOK{}
}

/*AddExternalServiceOK handles this case with default header values.

A successful response.
*/
type AddExternalServiceOK struct {
	Payload *AddExternalServiceOKBody
}

func (o *AddExternalServiceOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/AddExternalService][%d] addExternalServiceOk  %+v", 200, o.Payload)
}

func (o *AddExternalServiceOK) GetPayload() *AddExternalServiceOKBody {
	return o.Payload
}

func (o *AddExternalServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddExternalServiceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddExternalServiceDefault creates a AddExternalServiceDefault with default headers values
func NewAddExternalServiceDefault(code int) *AddExternalServiceDefault {
	return &AddExternalServiceDefault{
		_statusCode: code,
	}
}

/*AddExternalServiceDefault handles this case with default header values.

An unexpected error response.
*/
type AddExternalServiceDefault struct {
	_statusCode int

	Payload *AddExternalServiceDefaultBody
}

// Code gets the status code for the add external service default response
func (o *AddExternalServiceDefault) Code() int {
	return o._statusCode
}

func (o *AddExternalServiceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/AddExternalService][%d] AddExternalService default  %+v", o._statusCode, o.Payload)
}

func (o *AddExternalServiceDefault) GetPayload() *AddExternalServiceDefaultBody {
	return o.Payload
}

func (o *AddExternalServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddExternalServiceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddExternalServiceBody add external service body
swagger:model AddExternalServiceBody
*/
type AddExternalServiceBody struct {

	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"serviceName,omitempty"`

	// Node identifier where this instance runs. Required.
	NodeID string `json:"nodeId,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replicationSet,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"customLabels,omitempty"`

	// Group name of external service.
	Group string `json:"group,omitempty"`
}

// Validate validates this add external service body
func (o *AddExternalServiceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddExternalServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddExternalServiceBody) UnmarshalBinary(b []byte) error {
	var res AddExternalServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddExternalServiceDefaultBody add external service default body
swagger:model AddExternalServiceDefaultBody
*/
type AddExternalServiceDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this add external service default body
func (o *AddExternalServiceDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddExternalServiceDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("AddExternalService default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddExternalServiceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddExternalServiceDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddExternalServiceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddExternalServiceOKBody add external service OK body
swagger:model AddExternalServiceOKBody
*/
type AddExternalServiceOKBody struct {

	// external
	External *AddExternalServiceOKBodyExternal `json:"external,omitempty"`
}

// Validate validates this add external service OK body
func (o *AddExternalServiceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExternal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddExternalServiceOKBody) validateExternal(formats strfmt.Registry) error {

	if swag.IsZero(o.External) { // not required
		return nil
	}

	if o.External != nil {
		if err := o.External.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addExternalServiceOk" + "." + "external")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddExternalServiceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddExternalServiceOKBody) UnmarshalBinary(b []byte) error {
	var res AddExternalServiceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddExternalServiceOKBodyExternal ExternalService represents a generic External service instance.
swagger:model AddExternalServiceOKBodyExternal
*/
type AddExternalServiceOKBodyExternal struct {

	// Unique randomly generated instance identifier.
	ServiceID string `json:"serviceId,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"serviceName,omitempty"`

	// Node identifier where this service instance runs.
	NodeID string `json:"nodeId,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replicationSet,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"customLabels,omitempty"`

	// Group name of external service.
	Group string `json:"group,omitempty"`
}

// Validate validates this add external service OK body external
func (o *AddExternalServiceOKBodyExternal) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddExternalServiceOKBodyExternal) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddExternalServiceOKBodyExternal) UnmarshalBinary(b []byte) error {
	var res AddExternalServiceOKBodyExternal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
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
