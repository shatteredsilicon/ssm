// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// AddMongoDBExporterReader is a Reader for the AddMongoDBExporter structure.
type AddMongoDBExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddMongoDBExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddMongoDBExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddMongoDBExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddMongoDBExporterOK creates a AddMongoDBExporterOK with default headers values
func NewAddMongoDBExporterOK() *AddMongoDBExporterOK {
	return &AddMongoDBExporterOK{}
}

/*AddMongoDBExporterOK handles this case with default header values.

A successful response.
*/
type AddMongoDBExporterOK struct {
	Payload *AddMongoDBExporterOKBody
}

func (o *AddMongoDBExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddMongoDBExporter][%d] addMongoDbExporterOk  %+v", 200, o.Payload)
}

func (o *AddMongoDBExporterOK) GetPayload() *AddMongoDBExporterOKBody {
	return o.Payload
}

func (o *AddMongoDBExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMongoDBExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddMongoDBExporterDefault creates a AddMongoDBExporterDefault with default headers values
func NewAddMongoDBExporterDefault(code int) *AddMongoDBExporterDefault {
	return &AddMongoDBExporterDefault{
		_statusCode: code,
	}
}

/*AddMongoDBExporterDefault handles this case with default header values.

An unexpected error response.
*/
type AddMongoDBExporterDefault struct {
	_statusCode int

	Payload *AddMongoDBExporterDefaultBody
}

// Code gets the status code for the add mongo DB exporter default response
func (o *AddMongoDBExporterDefault) Code() int {
	return o._statusCode
}

func (o *AddMongoDBExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddMongoDBExporter][%d] AddMongoDBExporter default  %+v", o._statusCode, o.Payload)
}

func (o *AddMongoDBExporterDefault) GetPayload() *AddMongoDBExporterDefaultBody {
	return o.Payload
}

func (o *AddMongoDBExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMongoDBExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddMongoDBExporterBody add mongo DB exporter body
swagger:model AddMongoDBExporterBody
*/
type AddMongoDBExporterBody struct {

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MongoDB username for scraping metrics.
	Username string `json:"username,omitempty"`

	// MongoDB password for scraping metrics.
	Password string `json:"password,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Client certificate and key.
	TLSCertificateKey string `json:"tls_certificate_key,omitempty"`

	// Password for decrypting tls_certificate_key.
	TLSCertificateKeyFilePassword string `json:"tls_certificate_key_file_password,omitempty"`

	// Certificate Authority certificate chain.
	TLSCa string `json:"tls_ca,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Enables push metrics mode for exporter.
	PushMetrics bool `json:"push_metrics,omitempty"`

	// List of collector names to disable in this exporter.
	DisableCollectors []string `json:"disable_collectors"`

	// Authentication mechanism.
	// See https://docs.mongodb.com/manual/reference/connection-string/#mongodb-urioption-urioption.authMechanism
	// for details.
	AuthenticationMechanism string `json:"authentication_mechanism,omitempty"`

	// Authentication database.
	AuthenticationDatabase string `json:"authentication_database,omitempty"`

	// Custom password for exporter endpoint /metrics.
	AgentPassword string `json:"agent_password,omitempty"`

	// List of colletions to get stats from. Can use *
	StatsCollections []string `json:"stats_collections"`

	// Collections limit. Only get Databases and collection stats if the total number of collections in the server
	// is less than this value. 0: no limit
	CollectionsLimit int32 `json:"collections_limit,omitempty"`
}

// Validate validates this add mongo DB exporter body
func (o *AddMongoDBExporterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBExporterBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBExporterDefaultBody add mongo DB exporter default body
swagger:model AddMongoDBExporterDefaultBody
*/
type AddMongoDBExporterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this add mongo DB exporter default body
func (o *AddMongoDBExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMongoDBExporterDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("AddMongoDBExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBExporterOKBody add mongo DB exporter OK body
swagger:model AddMongoDBExporterOKBody
*/
type AddMongoDBExporterOKBody struct {

	// mongodb exporter
	MongodbExporter *AddMongoDBExporterOKBodyMongodbExporter `json:"mongodb_exporter,omitempty"`
}

// Validate validates this add mongo DB exporter OK body
func (o *AddMongoDBExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMongodbExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMongoDBExporterOKBody) validateMongodbExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.MongodbExporter) { // not required
		return nil
	}

	if o.MongodbExporter != nil {
		if err := o.MongodbExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMongoDbExporterOk" + "." + "mongodb_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBExporterOKBody) UnmarshalBinary(b []byte) error {
	var res AddMongoDBExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMongoDBExporterOKBodyMongodbExporter MongoDBExporter runs on Generic or Container Node and exposes MongoDB Service metrics.
swagger:model AddMongoDBExporterOKBodyMongodbExporter
*/
type AddMongoDBExporterOKBodyMongodbExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MongoDB username for scraping metrics.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetricsEnabled bool `json:"push_metrics_enabled,omitempty"`

	// List of disabled collector names.
	DisabledCollectors []string `json:"disabled_collectors"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	//  - UNKNOWN: Agent is not connected, we don't know anything about it's state.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE UNKNOWN]
	Status *string `json:"status,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`
}

// Validate validates this add mongo DB exporter OK body mongodb exporter
func (o *AddMongoDBExporterOKBodyMongodbExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum = append(addMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddMongoDBExporterOKBodyMongodbExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddMongoDBExporterOKBodyMongodbExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddMongoDBExporterOKBodyMongodbExporterStatusSTARTING captures enum value "STARTING"
	AddMongoDBExporterOKBodyMongodbExporterStatusSTARTING string = "STARTING"

	// AddMongoDBExporterOKBodyMongodbExporterStatusRUNNING captures enum value "RUNNING"
	AddMongoDBExporterOKBodyMongodbExporterStatusRUNNING string = "RUNNING"

	// AddMongoDBExporterOKBodyMongodbExporterStatusWAITING captures enum value "WAITING"
	AddMongoDBExporterOKBodyMongodbExporterStatusWAITING string = "WAITING"

	// AddMongoDBExporterOKBodyMongodbExporterStatusSTOPPING captures enum value "STOPPING"
	AddMongoDBExporterOKBodyMongodbExporterStatusSTOPPING string = "STOPPING"

	// AddMongoDBExporterOKBodyMongodbExporterStatusDONE captures enum value "DONE"
	AddMongoDBExporterOKBodyMongodbExporterStatusDONE string = "DONE"

	// AddMongoDBExporterOKBodyMongodbExporterStatusUNKNOWN captures enum value "UNKNOWN"
	AddMongoDBExporterOKBodyMongodbExporterStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *AddMongoDBExporterOKBodyMongodbExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddMongoDBExporterOKBodyMongodbExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addMongoDbExporterOk"+"."+"mongodb_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMongoDBExporterOKBodyMongodbExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMongoDBExporterOKBodyMongodbExporter) UnmarshalBinary(b []byte) error {
	var res AddMongoDBExporterOKBodyMongodbExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
