// Code generated by go-swagger; DO NOT EDIT.

package pxc_clusters

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

// CreatePXCClusterReader is a Reader for the CreatePXCCluster structure.
type CreatePXCClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePXCClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePXCClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreatePXCClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePXCClusterOK creates a CreatePXCClusterOK with default headers values
func NewCreatePXCClusterOK() *CreatePXCClusterOK {
	return &CreatePXCClusterOK{}
}

/*CreatePXCClusterOK handles this case with default header values.

A successful response.
*/
type CreatePXCClusterOK struct {
	Payload interface{}
}

func (o *CreatePXCClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PXCCluster/Create][%d] createPxcClusterOk  %+v", 200, o.Payload)
}

func (o *CreatePXCClusterOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CreatePXCClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePXCClusterDefault creates a CreatePXCClusterDefault with default headers values
func NewCreatePXCClusterDefault(code int) *CreatePXCClusterDefault {
	return &CreatePXCClusterDefault{
		_statusCode: code,
	}
}

/*CreatePXCClusterDefault handles this case with default header values.

An unexpected error response.
*/
type CreatePXCClusterDefault struct {
	_statusCode int

	Payload *CreatePXCClusterDefaultBody
}

// Code gets the status code for the create PXC cluster default response
func (o *CreatePXCClusterDefault) Code() int {
	return o._statusCode
}

func (o *CreatePXCClusterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PXCCluster/Create][%d] CreatePXCCluster default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePXCClusterDefault) GetPayload() *CreatePXCClusterDefaultBody {
	return o.Payload
}

func (o *CreatePXCClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreatePXCClusterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreatePXCClusterBody create PXC cluster body
swagger:model CreatePXCClusterBody
*/
type CreatePXCClusterBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// PXC cluster name.
	// a DNS-1035 label must consist of lower case alphanumeric characters or '-',
	// start with an alphabetic character, and end with an alphanumeric character
	// (e.g. 'my-name',  or 'abc-123', regex used for validation is '[a-z]([-a-z0-9]*[a-z0-9])?')
	Name string `json:"name,omitempty"`

	// Make DB cluster accessible outside of K8s cluster.
	Expose bool `json:"expose,omitempty"`

	// params
	Params *CreatePXCClusterParamsBodyParams `json:"params,omitempty"`
}

// Validate validates this create PXC cluster body
func (o *CreatePXCClusterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreatePXCClusterBody) validateParams(formats strfmt.Registry) error {

	if swag.IsZero(o.Params) { // not required
		return nil
	}

	if o.Params != nil {
		if err := o.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreatePXCClusterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePXCClusterBody) UnmarshalBinary(b []byte) error {
	var res CreatePXCClusterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreatePXCClusterDefaultBody create PXC cluster default body
swagger:model CreatePXCClusterDefaultBody
*/
type CreatePXCClusterDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this create PXC cluster default body
func (o *CreatePXCClusterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreatePXCClusterDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("CreatePXCCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreatePXCClusterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePXCClusterDefaultBody) UnmarshalBinary(b []byte) error {
	var res CreatePXCClusterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreatePXCClusterParamsBodyParams PXCClusterParams represents PXC cluster parameters that can be updated.
swagger:model CreatePXCClusterParamsBodyParams
*/
type CreatePXCClusterParamsBodyParams struct {

	// Cluster size.
	ClusterSize int32 `json:"cluster_size,omitempty"`

	// haproxy
	Haproxy *CreatePXCClusterParamsBodyParamsHaproxy `json:"haproxy,omitempty"`

	// proxysql
	Proxysql *CreatePXCClusterParamsBodyParamsProxysql `json:"proxysql,omitempty"`

	// pxc
	PXC *CreatePXCClusterParamsBodyParamsPXC `json:"pxc,omitempty"`
}

// Validate validates this create PXC cluster params body params
func (o *CreatePXCClusterParamsBodyParams) Validate(formats strfmt.Registry) error {
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

func (o *CreatePXCClusterParamsBodyParams) validateHaproxy(formats strfmt.Registry) error {

	if swag.IsZero(o.Haproxy) { // not required
		return nil
	}

	if o.Haproxy != nil {
		if err := o.Haproxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "haproxy")
			}
			return err
		}
	}

	return nil
}

func (o *CreatePXCClusterParamsBodyParams) validateProxysql(formats strfmt.Registry) error {

	if swag.IsZero(o.Proxysql) { // not required
		return nil
	}

	if o.Proxysql != nil {
		if err := o.Proxysql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "proxysql")
			}
			return err
		}
	}

	return nil
}

func (o *CreatePXCClusterParamsBodyParams) validatePXC(formats strfmt.Registry) error {

	if swag.IsZero(o.PXC) { // not required
		return nil
	}

	if o.PXC != nil {
		if err := o.PXC.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "pxc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreatePXCClusterParamsBodyParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePXCClusterParamsBodyParams) UnmarshalBinary(b []byte) error {
	var res CreatePXCClusterParamsBodyParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreatePXCClusterParamsBodyParamsHaproxy HAProxy container parameters.
// NOTE: HAProxy does not need disk size as ProxySQL does because the container does not require it.
swagger:model CreatePXCClusterParamsBodyParamsHaproxy
*/
type CreatePXCClusterParamsBodyParamsHaproxy struct {

	// Docker image used for HAProxy.
	Image string `json:"image,omitempty"`

	// compute resources
	ComputeResources *CreatePXCClusterParamsBodyParamsHaproxyComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this create PXC cluster params body params haproxy
func (o *CreatePXCClusterParamsBodyParamsHaproxy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreatePXCClusterParamsBodyParamsHaproxy) validateComputeResources(formats strfmt.Registry) error {

	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "haproxy" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreatePXCClusterParamsBodyParamsHaproxy) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePXCClusterParamsBodyParamsHaproxy) UnmarshalBinary(b []byte) error {
	var res CreatePXCClusterParamsBodyParamsHaproxy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreatePXCClusterParamsBodyParamsHaproxyComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model CreatePXCClusterParamsBodyParamsHaproxyComputeResources
*/
type CreatePXCClusterParamsBodyParamsHaproxyComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this create PXC cluster params body params haproxy compute resources
func (o *CreatePXCClusterParamsBodyParamsHaproxyComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreatePXCClusterParamsBodyParamsHaproxyComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePXCClusterParamsBodyParamsHaproxyComputeResources) UnmarshalBinary(b []byte) error {
	var res CreatePXCClusterParamsBodyParamsHaproxyComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreatePXCClusterParamsBodyParamsPXC PXC container parameters.
// TODO Do not use inner messages in all public APIs (for consistency).
swagger:model CreatePXCClusterParamsBodyParamsPXC
*/
type CreatePXCClusterParamsBodyParamsPXC struct {

	// Docker image used for PXC.
	Image string `json:"image,omitempty"`

	// Disk size in bytes.
	DiskSize string `json:"disk_size,omitempty"`

	// compute resources
	ComputeResources *CreatePXCClusterParamsBodyParamsPXCComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this create PXC cluster params body params PXC
func (o *CreatePXCClusterParamsBodyParamsPXC) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreatePXCClusterParamsBodyParamsPXC) validateComputeResources(formats strfmt.Registry) error {

	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "pxc" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreatePXCClusterParamsBodyParamsPXC) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePXCClusterParamsBodyParamsPXC) UnmarshalBinary(b []byte) error {
	var res CreatePXCClusterParamsBodyParamsPXC
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreatePXCClusterParamsBodyParamsPXCComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model CreatePXCClusterParamsBodyParamsPXCComputeResources
*/
type CreatePXCClusterParamsBodyParamsPXCComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this create PXC cluster params body params PXC compute resources
func (o *CreatePXCClusterParamsBodyParamsPXCComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreatePXCClusterParamsBodyParamsPXCComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePXCClusterParamsBodyParamsPXCComputeResources) UnmarshalBinary(b []byte) error {
	var res CreatePXCClusterParamsBodyParamsPXCComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreatePXCClusterParamsBodyParamsProxysql ProxySQL container parameters.
// TODO Do not use inner messages in all public APIs (for consistency).
swagger:model CreatePXCClusterParamsBodyParamsProxysql
*/
type CreatePXCClusterParamsBodyParamsProxysql struct {

	// Docker image used for ProxySQL.
	Image string `json:"image,omitempty"`

	// Disk size in bytes.
	DiskSize string `json:"disk_size,omitempty"`

	// compute resources
	ComputeResources *CreatePXCClusterParamsBodyParamsProxysqlComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this create PXC cluster params body params proxysql
func (o *CreatePXCClusterParamsBodyParamsProxysql) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreatePXCClusterParamsBodyParamsProxysql) validateComputeResources(formats strfmt.Registry) error {

	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "proxysql" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreatePXCClusterParamsBodyParamsProxysql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePXCClusterParamsBodyParamsProxysql) UnmarshalBinary(b []byte) error {
	var res CreatePXCClusterParamsBodyParamsProxysql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreatePXCClusterParamsBodyParamsProxysqlComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model CreatePXCClusterParamsBodyParamsProxysqlComputeResources
*/
type CreatePXCClusterParamsBodyParamsProxysqlComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this create PXC cluster params body params proxysql compute resources
func (o *CreatePXCClusterParamsBodyParamsProxysqlComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreatePXCClusterParamsBodyParamsProxysqlComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePXCClusterParamsBodyParamsProxysqlComputeResources) UnmarshalBinary(b []byte) error {
	var res CreatePXCClusterParamsBodyParamsProxysqlComputeResources
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

	// at type
	AtType string `json:"@type,omitempty"`
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
