// Code generated by go-swagger; DO NOT EDIT.

package ha_proxy

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

// AddHAProxyReader is a Reader for the AddHAProxy structure.
type AddHAProxyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddHAProxyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddHAProxyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddHAProxyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddHAProxyOK creates a AddHAProxyOK with default headers values
func NewAddHAProxyOK() *AddHAProxyOK {
	return &AddHAProxyOK{}
}

/*AddHAProxyOK handles this case with default header values.

A successful response.
*/
type AddHAProxyOK struct {
	Payload *AddHAProxyOKBody
}

func (o *AddHAProxyOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/HAProxy/Add][%d] addHaProxyOk  %+v", 200, o.Payload)
}

func (o *AddHAProxyOK) GetPayload() *AddHAProxyOKBody {
	return o.Payload
}

func (o *AddHAProxyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddHAProxyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddHAProxyDefault creates a AddHAProxyDefault with default headers values
func NewAddHAProxyDefault(code int) *AddHAProxyDefault {
	return &AddHAProxyDefault{
		_statusCode: code,
	}
}

/*AddHAProxyDefault handles this case with default header values.

An unexpected error response.
*/
type AddHAProxyDefault struct {
	_statusCode int

	Payload *AddHAProxyDefaultBody
}

// Code gets the status code for the add HA proxy default response
func (o *AddHAProxyDefault) Code() int {
	return o._statusCode
}

func (o *AddHAProxyDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/HAProxy/Add][%d] AddHAProxy default  %+v", o._statusCode, o.Payload)
}

func (o *AddHAProxyDefault) GetPayload() *AddHAProxyDefaultBody {
	return o.Payload
}

func (o *AddHAProxyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddHAProxyDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddHAProxyBody add HA proxy body
swagger:model AddHAProxyBody
*/
type AddHAProxyBody struct {

	// Node identifier on which an external exporter is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeID string `json:"nodeId,omitempty"`

	// Node name on which a service and node is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeName string `json:"nodeName,omitempty"`

	// Node and Exporter access address (DNS name or IP).
	// address always should be passed with add_node.
	Address string `json:"address,omitempty"`

	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"serviceName,omitempty"`

	// HTTP basic auth username for collecting metrics.
	Username string `json:"username,omitempty"`

	// HTTP basic auth password for collecting metrics.
	Password string `json:"password,omitempty"`

	// Scheme to generate URI to exporter metrics endpoints.
	Scheme string `json:"scheme,omitempty"`

	// Path under which metrics are exposed, used to generate URI.
	MetricsPath string `json:"metricsPath,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listenPort,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replicationSet,omitempty"`

	// Custom user-assigned labels for Service.
	CustomLabels map[string]string `json:"customLabels,omitempty"`

	// MetricsMode defines desired metrics mode for agent,
	// it can be pull, push or auto mode chosen by server.
	// Enum: [AUTO PULL PUSH]
	MetricsMode *string `json:"metricsMode,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skipConnectionCheck,omitempty"`

	// add node
	AddNode *AddHAProxyParamsBodyAddNode `json:"addNode,omitempty"`
}

// Validate validates this add HA proxy body
func (o *AddHAProxyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMetricsMode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAddNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addHaProxyBodyTypeMetricsModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AUTO","PULL","PUSH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addHaProxyBodyTypeMetricsModePropEnum = append(addHaProxyBodyTypeMetricsModePropEnum, v)
	}
}

const (

	// AddHAProxyBodyMetricsModeAUTO captures enum value "AUTO"
	AddHAProxyBodyMetricsModeAUTO string = "AUTO"

	// AddHAProxyBodyMetricsModePULL captures enum value "PULL"
	AddHAProxyBodyMetricsModePULL string = "PULL"

	// AddHAProxyBodyMetricsModePUSH captures enum value "PUSH"
	AddHAProxyBodyMetricsModePUSH string = "PUSH"
)

// prop value enum
func (o *AddHAProxyBody) validateMetricsModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addHaProxyBodyTypeMetricsModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddHAProxyBody) validateMetricsMode(formats strfmt.Registry) error {

	if swag.IsZero(o.MetricsMode) { // not required
		return nil
	}

	// value enum
	if err := o.validateMetricsModeEnum("body"+"."+"metricsMode", "body", *o.MetricsMode); err != nil {
		return err
	}

	return nil
}

func (o *AddHAProxyBody) validateAddNode(formats strfmt.Registry) error {

	if swag.IsZero(o.AddNode) { // not required
		return nil
	}

	if o.AddNode != nil {
		if err := o.AddNode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "addNode")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddHAProxyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddHAProxyBody) UnmarshalBinary(b []byte) error {
	var res AddHAProxyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddHAProxyDefaultBody add HA proxy default body
swagger:model AddHAProxyDefaultBody
*/
type AddHAProxyDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this add HA proxy default body
func (o *AddHAProxyDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddHAProxyDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("AddHAProxy default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddHAProxyDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddHAProxyDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddHAProxyDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddHAProxyOKBody add HA proxy OK body
swagger:model AddHAProxyOKBody
*/
type AddHAProxyOKBody struct {

	// external exporter
	ExternalExporter *AddHAProxyOKBodyExternalExporter `json:"externalExporter,omitempty"`

	// service
	Service *AddHAProxyOKBodyService `json:"service,omitempty"`
}

// Validate validates this add HA proxy OK body
func (o *AddHAProxyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExternalExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddHAProxyOKBody) validateExternalExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.ExternalExporter) { // not required
		return nil
	}

	if o.ExternalExporter != nil {
		if err := o.ExternalExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addHaProxyOk" + "." + "externalExporter")
			}
			return err
		}
	}

	return nil
}

func (o *AddHAProxyOKBody) validateService(formats strfmt.Registry) error {

	if swag.IsZero(o.Service) { // not required
		return nil
	}

	if o.Service != nil {
		if err := o.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addHaProxyOk" + "." + "service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddHAProxyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddHAProxyOKBody) UnmarshalBinary(b []byte) error {
	var res AddHAProxyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddHAProxyOKBodyExternalExporter ExternalExporter runs on any Node type, including Remote Node.
swagger:model AddHAProxyOKBodyExternalExporter
*/
type AddHAProxyOKBodyExternalExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agentId,omitempty"`

	// Node identifier where this instance runs.
	RunsOnNodeID string `json:"runsOnNodeId,omitempty"`

	// If disabled, metrics from this exporter will not be collected.
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"serviceId,omitempty"`

	// HTTP basic auth username for collecting metrics.
	Username string `json:"username,omitempty"`

	// Scheme to generate URI to exporter metrics endpoints.
	Scheme string `json:"scheme,omitempty"`

	// Path under which metrics are exposed, used to generate URI.
	MetricsPath string `json:"metricsPath,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"customLabels,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listenPort,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetricsEnabled bool `json:"pushMetricsEnabled,omitempty"`
}

// Validate validates this add HA proxy OK body external exporter
func (o *AddHAProxyOKBodyExternalExporter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddHAProxyOKBodyExternalExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddHAProxyOKBodyExternalExporter) UnmarshalBinary(b []byte) error {
	var res AddHAProxyOKBodyExternalExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddHAProxyOKBodyService HAProxyService represents a generic HAProxy service instance.
swagger:model AddHAProxyOKBodyService
*/
type AddHAProxyOKBodyService struct {

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
}

// Validate validates this add HA proxy OK body service
func (o *AddHAProxyOKBodyService) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddHAProxyOKBodyService) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddHAProxyOKBodyService) UnmarshalBinary(b []byte) error {
	var res AddHAProxyOKBodyService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddHAProxyParamsBodyAddNode AddNodeParams is a params to add new node to inventory while adding new service.
swagger:model AddHAProxyParamsBodyAddNode
*/
type AddHAProxyParamsBodyAddNode struct {

	// NodeType describes supported Node types.
	// Enum: [NODE_TYPE_INVALID GENERIC_NODE CONTAINER_NODE REMOTE_NODE REMOTE_RDS_NODE REMOTE_AZURE_DATABASE_NODE]
	NodeType *string `json:"nodeType,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"nodeName,omitempty"`

	// Linux machine-id.
	MachineID string `json:"machineId,omitempty"`

	// Linux distribution name and version.
	Distro string `json:"distro,omitempty"`

	// Container identifier. If specified, must be a unique Docker container identifier.
	ContainerID string `json:"containerId,omitempty"`

	// Container name.
	ContainerName string `json:"containerName,omitempty"`

	// Node model.
	NodeModel string `json:"nodeModel,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels for Node.
	CustomLabels map[string]string `json:"customLabels,omitempty"`
}

// Validate validates this add HA proxy params body add node
func (o *AddHAProxyParamsBodyAddNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNodeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addHaProxyParamsBodyAddNodeTypeNodeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NODE_TYPE_INVALID","GENERIC_NODE","CONTAINER_NODE","REMOTE_NODE","REMOTE_RDS_NODE","REMOTE_AZURE_DATABASE_NODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addHaProxyParamsBodyAddNodeTypeNodeTypePropEnum = append(addHaProxyParamsBodyAddNodeTypeNodeTypePropEnum, v)
	}
}

const (

	// AddHAProxyParamsBodyAddNodeNodeTypeNODETYPEINVALID captures enum value "NODE_TYPE_INVALID"
	AddHAProxyParamsBodyAddNodeNodeTypeNODETYPEINVALID string = "NODE_TYPE_INVALID"

	// AddHAProxyParamsBodyAddNodeNodeTypeGENERICNODE captures enum value "GENERIC_NODE"
	AddHAProxyParamsBodyAddNodeNodeTypeGENERICNODE string = "GENERIC_NODE"

	// AddHAProxyParamsBodyAddNodeNodeTypeCONTAINERNODE captures enum value "CONTAINER_NODE"
	AddHAProxyParamsBodyAddNodeNodeTypeCONTAINERNODE string = "CONTAINER_NODE"

	// AddHAProxyParamsBodyAddNodeNodeTypeREMOTENODE captures enum value "REMOTE_NODE"
	AddHAProxyParamsBodyAddNodeNodeTypeREMOTENODE string = "REMOTE_NODE"

	// AddHAProxyParamsBodyAddNodeNodeTypeREMOTERDSNODE captures enum value "REMOTE_RDS_NODE"
	AddHAProxyParamsBodyAddNodeNodeTypeREMOTERDSNODE string = "REMOTE_RDS_NODE"

	// AddHAProxyParamsBodyAddNodeNodeTypeREMOTEAZUREDATABASENODE captures enum value "REMOTE_AZURE_DATABASE_NODE"
	AddHAProxyParamsBodyAddNodeNodeTypeREMOTEAZUREDATABASENODE string = "REMOTE_AZURE_DATABASE_NODE"
)

// prop value enum
func (o *AddHAProxyParamsBodyAddNode) validateNodeTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addHaProxyParamsBodyAddNodeTypeNodeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddHAProxyParamsBodyAddNode) validateNodeType(formats strfmt.Registry) error {

	if swag.IsZero(o.NodeType) { // not required
		return nil
	}

	// value enum
	if err := o.validateNodeTypeEnum("body"+"."+"addNode"+"."+"nodeType", "body", *o.NodeType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddHAProxyParamsBodyAddNode) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddHAProxyParamsBodyAddNode) UnmarshalBinary(b []byte) error {
	var res AddHAProxyParamsBodyAddNode
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
