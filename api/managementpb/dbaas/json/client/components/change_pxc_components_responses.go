// Code generated by go-swagger; DO NOT EDIT.

package components

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

// ChangePXCComponentsReader is a Reader for the ChangePXCComponents structure.
type ChangePXCComponentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangePXCComponentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangePXCComponentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangePXCComponentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangePXCComponentsOK creates a ChangePXCComponentsOK with default headers values
func NewChangePXCComponentsOK() *ChangePXCComponentsOK {
	return &ChangePXCComponentsOK{}
}

/*ChangePXCComponentsOK handles this case with default header values.

A successful response.
*/
type ChangePXCComponentsOK struct {
	Payload interface{}
}

func (o *ChangePXCComponentsOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/ChangePXC][%d] changePxcComponentsOk  %+v", 200, o.Payload)
}

func (o *ChangePXCComponentsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ChangePXCComponentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangePXCComponentsDefault creates a ChangePXCComponentsDefault with default headers values
func NewChangePXCComponentsDefault(code int) *ChangePXCComponentsDefault {
	return &ChangePXCComponentsDefault{
		_statusCode: code,
	}
}

/*ChangePXCComponentsDefault handles this case with default header values.

An unexpected error response.
*/
type ChangePXCComponentsDefault struct {
	_statusCode int

	Payload *ChangePXCComponentsDefaultBody
}

// Code gets the status code for the change PXC components default response
func (o *ChangePXCComponentsDefault) Code() int {
	return o._statusCode
}

func (o *ChangePXCComponentsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/ChangePXC][%d] ChangePXCComponents default  %+v", o._statusCode, o.Payload)
}

func (o *ChangePXCComponentsDefault) GetPayload() *ChangePXCComponentsDefaultBody {
	return o.Payload
}

func (o *ChangePXCComponentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangePXCComponentsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangePXCComponentsBody change PXC components body
swagger:model ChangePXCComponentsBody
*/
type ChangePXCComponentsBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// haproxy
	Haproxy *ChangePXCComponentsParamsBodyHaproxy `json:"haproxy,omitempty"`

	// proxysql
	Proxysql *ChangePXCComponentsParamsBodyProxysql `json:"proxysql,omitempty"`

	// pxc
	PXC *ChangePXCComponentsParamsBodyPXC `json:"pxc,omitempty"`
}

// Validate validates this change PXC components body
func (o *ChangePXCComponentsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHaproxy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProxysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePXC(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangePXCComponentsBody) validateHaproxy(formats strfmt.Registry) error {

	if swag.IsZero(o.Haproxy) { // not required
		return nil
	}

	if o.Haproxy != nil {
		if err := o.Haproxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "haproxy")
			}
			return err
		}
	}

	return nil
}

func (o *ChangePXCComponentsBody) validateProxysql(formats strfmt.Registry) error {

	if swag.IsZero(o.Proxysql) { // not required
		return nil
	}

	if o.Proxysql != nil {
		if err := o.Proxysql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "proxysql")
			}
			return err
		}
	}

	return nil
}

func (o *ChangePXCComponentsBody) validatePXC(formats strfmt.Registry) error {

	if swag.IsZero(o.PXC) { // not required
		return nil
	}

	if o.PXC != nil {
		if err := o.PXC.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "pxc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangePXCComponentsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePXCComponentsBody) UnmarshalBinary(b []byte) error {
	var res ChangePXCComponentsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangePXCComponentsDefaultBody change PXC components default body
swagger:model ChangePXCComponentsDefaultBody
*/
type ChangePXCComponentsDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this change PXC components default body
func (o *ChangePXCComponentsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangePXCComponentsDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ChangePXCComponents default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangePXCComponentsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePXCComponentsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangePXCComponentsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangePXCComponentsParamsBodyHaproxy ChangeComponent contains fields to manage components.
swagger:model ChangePXCComponentsParamsBodyHaproxy
*/
type ChangePXCComponentsParamsBodyHaproxy struct {

	// default version
	DefaultVersion string `json:"default_version,omitempty"`

	// versions
	Versions []*ChangePXCComponentsParamsBodyHaproxyVersionsItems0 `json:"versions"`
}

// Validate validates this change PXC components params body haproxy
func (o *ChangePXCComponentsParamsBodyHaproxy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangePXCComponentsParamsBodyHaproxy) validateVersions(formats strfmt.Registry) error {

	if swag.IsZero(o.Versions) { // not required
		return nil
	}

	for i := 0; i < len(o.Versions); i++ {
		if swag.IsZero(o.Versions[i]) { // not required
			continue
		}

		if o.Versions[i] != nil {
			if err := o.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "haproxy" + "." + "versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangePXCComponentsParamsBodyHaproxy) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePXCComponentsParamsBodyHaproxy) UnmarshalBinary(b []byte) error {
	var res ChangePXCComponentsParamsBodyHaproxy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangePXCComponentsParamsBodyHaproxyVersionsItems0 ComponentVersion contains operations which should be done with component version.
swagger:model ChangePXCComponentsParamsBodyHaproxyVersionsItems0
*/
type ChangePXCComponentsParamsBodyHaproxyVersionsItems0 struct {

	// version
	Version string `json:"version,omitempty"`

	// disable
	Disable bool `json:"disable,omitempty"`

	// enable
	Enable bool `json:"enable,omitempty"`
}

// Validate validates this change PXC components params body haproxy versions items0
func (o *ChangePXCComponentsParamsBodyHaproxyVersionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangePXCComponentsParamsBodyHaproxyVersionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePXCComponentsParamsBodyHaproxyVersionsItems0) UnmarshalBinary(b []byte) error {
	var res ChangePXCComponentsParamsBodyHaproxyVersionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangePXCComponentsParamsBodyPXC ChangeComponent contains fields to manage components.
swagger:model ChangePXCComponentsParamsBodyPXC
*/
type ChangePXCComponentsParamsBodyPXC struct {

	// default version
	DefaultVersion string `json:"default_version,omitempty"`

	// versions
	Versions []*ChangePXCComponentsParamsBodyPXCVersionsItems0 `json:"versions"`
}

// Validate validates this change PXC components params body PXC
func (o *ChangePXCComponentsParamsBodyPXC) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangePXCComponentsParamsBodyPXC) validateVersions(formats strfmt.Registry) error {

	if swag.IsZero(o.Versions) { // not required
		return nil
	}

	for i := 0; i < len(o.Versions); i++ {
		if swag.IsZero(o.Versions[i]) { // not required
			continue
		}

		if o.Versions[i] != nil {
			if err := o.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "pxc" + "." + "versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangePXCComponentsParamsBodyPXC) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePXCComponentsParamsBodyPXC) UnmarshalBinary(b []byte) error {
	var res ChangePXCComponentsParamsBodyPXC
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangePXCComponentsParamsBodyPXCVersionsItems0 ComponentVersion contains operations which should be done with component version.
swagger:model ChangePXCComponentsParamsBodyPXCVersionsItems0
*/
type ChangePXCComponentsParamsBodyPXCVersionsItems0 struct {

	// version
	Version string `json:"version,omitempty"`

	// disable
	Disable bool `json:"disable,omitempty"`

	// enable
	Enable bool `json:"enable,omitempty"`
}

// Validate validates this change PXC components params body PXC versions items0
func (o *ChangePXCComponentsParamsBodyPXCVersionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangePXCComponentsParamsBodyPXCVersionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePXCComponentsParamsBodyPXCVersionsItems0) UnmarshalBinary(b []byte) error {
	var res ChangePXCComponentsParamsBodyPXCVersionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangePXCComponentsParamsBodyProxysql ChangeComponent contains fields to manage components.
swagger:model ChangePXCComponentsParamsBodyProxysql
*/
type ChangePXCComponentsParamsBodyProxysql struct {

	// default version
	DefaultVersion string `json:"default_version,omitempty"`

	// versions
	Versions []*ChangePXCComponentsParamsBodyProxysqlVersionsItems0 `json:"versions"`
}

// Validate validates this change PXC components params body proxysql
func (o *ChangePXCComponentsParamsBodyProxysql) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangePXCComponentsParamsBodyProxysql) validateVersions(formats strfmt.Registry) error {

	if swag.IsZero(o.Versions) { // not required
		return nil
	}

	for i := 0; i < len(o.Versions); i++ {
		if swag.IsZero(o.Versions[i]) { // not required
			continue
		}

		if o.Versions[i] != nil {
			if err := o.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "proxysql" + "." + "versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangePXCComponentsParamsBodyProxysql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePXCComponentsParamsBodyProxysql) UnmarshalBinary(b []byte) error {
	var res ChangePXCComponentsParamsBodyProxysql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangePXCComponentsParamsBodyProxysqlVersionsItems0 ComponentVersion contains operations which should be done with component version.
swagger:model ChangePXCComponentsParamsBodyProxysqlVersionsItems0
*/
type ChangePXCComponentsParamsBodyProxysqlVersionsItems0 struct {

	// version
	Version string `json:"version,omitempty"`

	// disable
	Disable bool `json:"disable,omitempty"`

	// enable
	Enable bool `json:"enable,omitempty"`
}

// Validate validates this change PXC components params body proxysql versions items0
func (o *ChangePXCComponentsParamsBodyProxysqlVersionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangePXCComponentsParamsBodyProxysqlVersionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePXCComponentsParamsBodyProxysqlVersionsItems0) UnmarshalBinary(b []byte) error {
	var res ChangePXCComponentsParamsBodyProxysqlVersionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
