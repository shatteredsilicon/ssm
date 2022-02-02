// Code generated by go-swagger; DO NOT EDIT.

package channels

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

// AddChannelReader is a Reader for the AddChannel structure.
type AddChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddChannelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddChannelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddChannelOK creates a AddChannelOK with default headers values
func NewAddChannelOK() *AddChannelOK {
	return &AddChannelOK{}
}

/*AddChannelOK handles this case with default header values.

A successful response.
*/
type AddChannelOK struct {
	Payload *AddChannelOKBody
}

func (o *AddChannelOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Channels/Add][%d] addChannelOk  %+v", 200, o.Payload)
}

func (o *AddChannelOK) GetPayload() *AddChannelOKBody {
	return o.Payload
}

func (o *AddChannelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddChannelOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddChannelDefault creates a AddChannelDefault with default headers values
func NewAddChannelDefault(code int) *AddChannelDefault {
	return &AddChannelDefault{
		_statusCode: code,
	}
}

/*AddChannelDefault handles this case with default header values.

An unexpected error response.
*/
type AddChannelDefault struct {
	_statusCode int

	Payload *AddChannelDefaultBody
}

// Code gets the status code for the add channel default response
func (o *AddChannelDefault) Code() int {
	return o._statusCode
}

func (o *AddChannelDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Channels/Add][%d] AddChannel default  %+v", o._statusCode, o.Payload)
}

func (o *AddChannelDefault) GetPayload() *AddChannelDefaultBody {
	return o.Payload
}

func (o *AddChannelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddChannelDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddChannelBody add channel body
swagger:model AddChannelBody
*/
type AddChannelBody struct {

	// Short human-readable summary.
	Summary string `json:"summary,omitempty"`

	// New channel status.
	Disabled bool `json:"disabled,omitempty"`

	// email config
	EmailConfig *AddChannelParamsBodyEmailConfig `json:"email_config,omitempty"`

	// pagerduty config
	PagerdutyConfig *AddChannelParamsBodyPagerdutyConfig `json:"pagerduty_config,omitempty"`

	// slack config
	SlackConfig *AddChannelParamsBodySlackConfig `json:"slack_config,omitempty"`

	// webhook config
	WebhookConfig *AddChannelParamsBodyWebhookConfig `json:"webhook_config,omitempty"`
}

// Validate validates this add channel body
func (o *AddChannelBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmailConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePagerdutyConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSlackConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWebhookConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddChannelBody) validateEmailConfig(formats strfmt.Registry) error {

	if swag.IsZero(o.EmailConfig) { // not required
		return nil
	}

	if o.EmailConfig != nil {
		if err := o.EmailConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "email_config")
			}
			return err
		}
	}

	return nil
}

func (o *AddChannelBody) validatePagerdutyConfig(formats strfmt.Registry) error {

	if swag.IsZero(o.PagerdutyConfig) { // not required
		return nil
	}

	if o.PagerdutyConfig != nil {
		if err := o.PagerdutyConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "pagerduty_config")
			}
			return err
		}
	}

	return nil
}

func (o *AddChannelBody) validateSlackConfig(formats strfmt.Registry) error {

	if swag.IsZero(o.SlackConfig) { // not required
		return nil
	}

	if o.SlackConfig != nil {
		if err := o.SlackConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "slack_config")
			}
			return err
		}
	}

	return nil
}

func (o *AddChannelBody) validateWebhookConfig(formats strfmt.Registry) error {

	if swag.IsZero(o.WebhookConfig) { // not required
		return nil
	}

	if o.WebhookConfig != nil {
		if err := o.WebhookConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "webhook_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddChannelBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddChannelBody) UnmarshalBinary(b []byte) error {
	var res AddChannelBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddChannelDefaultBody add channel default body
swagger:model AddChannelDefaultBody
*/
type AddChannelDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this add channel default body
func (o *AddChannelDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddChannelDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("AddChannel default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddChannelDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddChannelDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddChannelDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddChannelOKBody add channel OK body
swagger:model AddChannelOKBody
*/
type AddChannelOKBody struct {

	// Machine-readable ID.
	ChannelID string `json:"channel_id,omitempty"`
}

// Validate validates this add channel OK body
func (o *AddChannelOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddChannelOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddChannelOKBody) UnmarshalBinary(b []byte) error {
	var res AddChannelOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddChannelParamsBodyEmailConfig EmailConfig represents email configuration.
swagger:model AddChannelParamsBodyEmailConfig
*/
type AddChannelParamsBodyEmailConfig struct {

	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// to
	To []string `json:"to"`
}

// Validate validates this add channel params body email config
func (o *AddChannelParamsBodyEmailConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddChannelParamsBodyEmailConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddChannelParamsBodyEmailConfig) UnmarshalBinary(b []byte) error {
	var res AddChannelParamsBodyEmailConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddChannelParamsBodyPagerdutyConfig PagerDutyConfig represents PagerDuty configuration.
swagger:model AddChannelParamsBodyPagerdutyConfig
*/
type AddChannelParamsBodyPagerdutyConfig struct {

	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// The PagerDuty key for "Events API v2" integration type. Exactly one key should be set.
	RoutingKey string `json:"routing_key,omitempty"`

	// The PagerDuty key for "Prometheus" integration type. Exactly one key should be set.
	ServiceKey string `json:"service_key,omitempty"`
}

// Validate validates this add channel params body pagerduty config
func (o *AddChannelParamsBodyPagerdutyConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddChannelParamsBodyPagerdutyConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddChannelParamsBodyPagerdutyConfig) UnmarshalBinary(b []byte) error {
	var res AddChannelParamsBodyPagerdutyConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddChannelParamsBodySlackConfig SlackConfig represents Slack configuration.
swagger:model AddChannelParamsBodySlackConfig
*/
type AddChannelParamsBodySlackConfig struct {

	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// channel
	Channel string `json:"channel,omitempty"`
}

// Validate validates this add channel params body slack config
func (o *AddChannelParamsBodySlackConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddChannelParamsBodySlackConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddChannelParamsBodySlackConfig) UnmarshalBinary(b []byte) error {
	var res AddChannelParamsBodySlackConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddChannelParamsBodyWebhookConfig WebhookConfig represents webhook configuration.
swagger:model AddChannelParamsBodyWebhookConfig
*/
type AddChannelParamsBodyWebhookConfig struct {

	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// max alerts
	MaxAlerts int32 `json:"max_alerts,omitempty"`

	// http config
	HTTPConfig *AddChannelParamsBodyWebhookConfigHTTPConfig `json:"http_config,omitempty"`
}

// Validate validates this add channel params body webhook config
func (o *AddChannelParamsBodyWebhookConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHTTPConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddChannelParamsBodyWebhookConfig) validateHTTPConfig(formats strfmt.Registry) error {

	if swag.IsZero(o.HTTPConfig) { // not required
		return nil
	}

	if o.HTTPConfig != nil {
		if err := o.HTTPConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "webhook_config" + "." + "http_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddChannelParamsBodyWebhookConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddChannelParamsBodyWebhookConfig) UnmarshalBinary(b []byte) error {
	var res AddChannelParamsBodyWebhookConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddChannelParamsBodyWebhookConfigHTTPConfig HTTPConfig represents HTTP client configuration.
swagger:model AddChannelParamsBodyWebhookConfigHTTPConfig
*/
type AddChannelParamsBodyWebhookConfigHTTPConfig struct {

	// bearer token
	BearerToken string `json:"bearer_token,omitempty"`

	// bearer token file
	BearerTokenFile string `json:"bearer_token_file,omitempty"`

	// proxy url
	ProxyURL string `json:"proxy_url,omitempty"`

	// basic auth
	BasicAuth *AddChannelParamsBodyWebhookConfigHTTPConfigBasicAuth `json:"basic_auth,omitempty"`

	// tls config
	TLSConfig *AddChannelParamsBodyWebhookConfigHTTPConfigTLSConfig `json:"tls_config,omitempty"`
}

// Validate validates this add channel params body webhook config HTTP config
func (o *AddChannelParamsBodyWebhookConfigHTTPConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBasicAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTLSConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddChannelParamsBodyWebhookConfigHTTPConfig) validateBasicAuth(formats strfmt.Registry) error {

	if swag.IsZero(o.BasicAuth) { // not required
		return nil
	}

	if o.BasicAuth != nil {
		if err := o.BasicAuth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "webhook_config" + "." + "http_config" + "." + "basic_auth")
			}
			return err
		}
	}

	return nil
}

func (o *AddChannelParamsBodyWebhookConfigHTTPConfig) validateTLSConfig(formats strfmt.Registry) error {

	if swag.IsZero(o.TLSConfig) { // not required
		return nil
	}

	if o.TLSConfig != nil {
		if err := o.TLSConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "webhook_config" + "." + "http_config" + "." + "tls_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddChannelParamsBodyWebhookConfigHTTPConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddChannelParamsBodyWebhookConfigHTTPConfig) UnmarshalBinary(b []byte) error {
	var res AddChannelParamsBodyWebhookConfigHTTPConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddChannelParamsBodyWebhookConfigHTTPConfigBasicAuth BasicAuth represents basic HTTP auth configuration.
swagger:model AddChannelParamsBodyWebhookConfigHTTPConfigBasicAuth
*/
type AddChannelParamsBodyWebhookConfigHTTPConfigBasicAuth struct {

	// username
	Username string `json:"username,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// password file
	PasswordFile string `json:"password_file,omitempty"`
}

// Validate validates this add channel params body webhook config HTTP config basic auth
func (o *AddChannelParamsBodyWebhookConfigHTTPConfigBasicAuth) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddChannelParamsBodyWebhookConfigHTTPConfigBasicAuth) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddChannelParamsBodyWebhookConfigHTTPConfigBasicAuth) UnmarshalBinary(b []byte) error {
	var res AddChannelParamsBodyWebhookConfigHTTPConfigBasicAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddChannelParamsBodyWebhookConfigHTTPConfigTLSConfig TLSConfig represents TLS configuration for alertmanager
// https://prometheus.io/docs/alerting/latest/configuration/#tls_config
swagger:model AddChannelParamsBodyWebhookConfigHTTPConfigTLSConfig
*/
type AddChannelParamsBodyWebhookConfigHTTPConfigTLSConfig struct {

	// A path to the CA certificate file to validate the server certificate with.
	// ca_file and ca_file_content should not be set at the same time.
	CaFile string `json:"ca_file,omitempty"`

	// A path to the certificate file for client cert authentication to the server.
	// cert_file and cert_file_content should not be set at the same time.
	CertFile string `json:"cert_file,omitempty"`

	// A path to the key file for client cert authentication to the server.
	// key_file and key_file_content should not be set at the same time.
	KeyFile string `json:"key_file,omitempty"`

	// Name of the server.
	ServerName string `json:"server_name,omitempty"`

	// Disable validation of the server certificate.
	InsecureSkipVerify bool `json:"insecure_skip_verify,omitempty"`

	// CA certificate to validate the server certificate with.
	// ca_file and ca_file_content should not be set at the same time.
	CaFileContent string `json:"ca_file_content,omitempty"`

	// A certificate for client cert authentication to the server.
	// cert_file and cert_file_content should not be set at the same time.
	CertFileContent string `json:"cert_file_content,omitempty"`

	// A key for client cert authentication to the server.
	// key_file and key_file_content should not be set at the same time.
	KeyFileContent string `json:"key_file_content,omitempty"`
}

// Validate validates this add channel params body webhook config HTTP config TLS config
func (o *AddChannelParamsBodyWebhookConfigHTTPConfigTLSConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddChannelParamsBodyWebhookConfigHTTPConfigTLSConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddChannelParamsBodyWebhookConfigHTTPConfigTLSConfig) UnmarshalBinary(b []byte) error {
	var res AddChannelParamsBodyWebhookConfigHTTPConfigTLSConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//     Foo foo = ...;
//     Any any;
//     any.PackFrom(foo);
//     ...
//     if (any.UnpackTo(&foo)) {
//       ...
//     }
//
// Example 2: Pack and unpack a message in Java.
//
//     Foo foo = ...;
//     Any any = Any.pack(foo);
//     ...
//     if (any.is(Foo.class)) {
//       foo = any.unpack(Foo.class);
//     }
//
//  Example 3: Pack and unpack a message in Python.
//
//     foo = Foo(...)
//     any = Any()
//     any.Pack(foo)
//     ...
//     if any.Is(Foo.DESCRIPTOR):
//       any.Unpack(foo)
//       ...
//
//  Example 4: Pack and unpack a message in Go
//
//      foo := &pb.Foo{...}
//      any, err := ptypes.MarshalAny(foo)
//      ...
//      foo := &pb.Foo{}
//      if err := ptypes.UnmarshalAny(any, foo); err != nil {
//        ...
//      }
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
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
