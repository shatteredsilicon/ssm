// Code generated by go-swagger; DO NOT EDIT.

package db_clusters

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

// RestartDBClusterReader is a Reader for the RestartDBCluster structure.
type RestartDBClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartDBClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestartDBClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRestartDBClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestartDBClusterOK creates a RestartDBClusterOK with default headers values
func NewRestartDBClusterOK() *RestartDBClusterOK {
	return &RestartDBClusterOK{}
}

/*RestartDBClusterOK handles this case with default header values.

A successful response.
*/
type RestartDBClusterOK struct {
	Payload interface{}
}

func (o *RestartDBClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/DBClusters/Restart][%d] restartDbClusterOk  %+v", 200, o.Payload)
}

func (o *RestartDBClusterOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RestartDBClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartDBClusterDefault creates a RestartDBClusterDefault with default headers values
func NewRestartDBClusterDefault(code int) *RestartDBClusterDefault {
	return &RestartDBClusterDefault{
		_statusCode: code,
	}
}

/*RestartDBClusterDefault handles this case with default header values.

An unexpected error response.
*/
type RestartDBClusterDefault struct {
	_statusCode int

	Payload *RestartDBClusterDefaultBody
}

// Code gets the status code for the restart DB cluster default response
func (o *RestartDBClusterDefault) Code() int {
	return o._statusCode
}

func (o *RestartDBClusterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/DBClusters/Restart][%d] RestartDBCluster default  %+v", o._statusCode, o.Payload)
}

func (o *RestartDBClusterDefault) GetPayload() *RestartDBClusterDefaultBody {
	return o.Payload
}

func (o *RestartDBClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RestartDBClusterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RestartDBClusterBody restart DB cluster body
swagger:model RestartDBClusterBody
*/
type RestartDBClusterBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// PXC cluster name.
	Name string `json:"name,omitempty"`

	// DBClusterType represents database cluster type.
	//
	//  - DB_CLUSTER_TYPE_INVALID: DB_CLUSTER_TYPE_INVALID represents unknown cluster type.
	//  - DB_CLUSTER_TYPE_PXC: DB_CLUSTER_TYPE_PXC represents pxc cluster type.
	//  - DB_CLUSTER_TYPE_PSMDB: DB_CLUSTER_TYPE_PSMDB represents psmdb cluster type.
	// Enum: [DB_CLUSTER_TYPE_INVALID DB_CLUSTER_TYPE_PXC DB_CLUSTER_TYPE_PSMDB]
	ClusterType *string `json:"cluster_type,omitempty"`
}

// Validate validates this restart DB cluster body
func (o *RestartDBClusterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClusterType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var restartDbClusterBodyTypeClusterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DB_CLUSTER_TYPE_INVALID","DB_CLUSTER_TYPE_PXC","DB_CLUSTER_TYPE_PSMDB"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		restartDbClusterBodyTypeClusterTypePropEnum = append(restartDbClusterBodyTypeClusterTypePropEnum, v)
	}
}

const (

	// RestartDBClusterBodyClusterTypeDBCLUSTERTYPEINVALID captures enum value "DB_CLUSTER_TYPE_INVALID"
	RestartDBClusterBodyClusterTypeDBCLUSTERTYPEINVALID string = "DB_CLUSTER_TYPE_INVALID"

	// RestartDBClusterBodyClusterTypeDBCLUSTERTYPEPXC captures enum value "DB_CLUSTER_TYPE_PXC"
	RestartDBClusterBodyClusterTypeDBCLUSTERTYPEPXC string = "DB_CLUSTER_TYPE_PXC"

	// RestartDBClusterBodyClusterTypeDBCLUSTERTYPEPSMDB captures enum value "DB_CLUSTER_TYPE_PSMDB"
	RestartDBClusterBodyClusterTypeDBCLUSTERTYPEPSMDB string = "DB_CLUSTER_TYPE_PSMDB"
)

// prop value enum
func (o *RestartDBClusterBody) validateClusterTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, restartDbClusterBodyTypeClusterTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *RestartDBClusterBody) validateClusterType(formats strfmt.Registry) error {

	if swag.IsZero(o.ClusterType) { // not required
		return nil
	}

	// value enum
	if err := o.validateClusterTypeEnum("body"+"."+"cluster_type", "body", *o.ClusterType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RestartDBClusterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RestartDBClusterBody) UnmarshalBinary(b []byte) error {
	var res RestartDBClusterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RestartDBClusterDefaultBody restart DB cluster default body
swagger:model RestartDBClusterDefaultBody
*/
type RestartDBClusterDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this restart DB cluster default body
func (o *RestartDBClusterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RestartDBClusterDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("RestartDBCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *RestartDBClusterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RestartDBClusterDefaultBody) UnmarshalBinary(b []byte) error {
	var res RestartDBClusterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
