// Code generated by go-swagger; DO NOT EDIT.

package locations

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

// ListLocationsReader is a Reader for the ListLocations structure.
type ListLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListLocationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListLocationsOK creates a ListLocationsOK with default headers values
func NewListLocationsOK() *ListLocationsOK {
	return &ListLocationsOK{}
}

/*ListLocationsOK handles this case with default header values.

A successful response.
*/
type ListLocationsOK struct {
	Payload *ListLocationsOKBody
}

func (o *ListLocationsOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Locations/List][%d] listLocationsOk  %+v", 200, o.Payload)
}

func (o *ListLocationsOK) GetPayload() *ListLocationsOKBody {
	return o.Payload
}

func (o *ListLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListLocationsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLocationsDefault creates a ListLocationsDefault with default headers values
func NewListLocationsDefault(code int) *ListLocationsDefault {
	return &ListLocationsDefault{
		_statusCode: code,
	}
}

/*ListLocationsDefault handles this case with default header values.

An unexpected error response.
*/
type ListLocationsDefault struct {
	_statusCode int

	Payload *ListLocationsDefaultBody
}

// Code gets the status code for the list locations default response
func (o *ListLocationsDefault) Code() int {
	return o._statusCode
}

func (o *ListLocationsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Locations/List][%d] ListLocations default  %+v", o._statusCode, o.Payload)
}

func (o *ListLocationsDefault) GetPayload() *ListLocationsDefaultBody {
	return o.Payload
}

func (o *ListLocationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListLocationsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListLocationsDefaultBody list locations default body
swagger:model ListLocationsDefaultBody
*/
type ListLocationsDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this list locations default body
func (o *ListLocationsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListLocationsDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ListLocations default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListLocationsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListLocationsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListLocationsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListLocationsOKBody list locations OK body
swagger:model ListLocationsOKBody
*/
type ListLocationsOKBody struct {

	// locations
	Locations []*LocationsItems0 `json:"locations"`
}

// Validate validates this list locations OK body
func (o *ListLocationsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLocations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListLocationsOKBody) validateLocations(formats strfmt.Registry) error {

	if swag.IsZero(o.Locations) { // not required
		return nil
	}

	for i := 0; i < len(o.Locations); i++ {
		if swag.IsZero(o.Locations[i]) { // not required
			continue
		}

		if o.Locations[i] != nil {
			if err := o.Locations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listLocationsOk" + "." + "locations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListLocationsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListLocationsOKBody) UnmarshalBinary(b []byte) error {
	var res ListLocationsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LocationsItems0 Location represents single Backup Location.
swagger:model LocationsItems0
*/
type LocationsItems0 struct {

	// Machine-readable ID.
	LocationID string `json:"locationId,omitempty"`

	// Location name
	Name string `json:"name,omitempty"`

	// Short description
	Description string `json:"description,omitempty"`

	// pmm client config
	PMMClientConfig *LocationsItems0PMMClientConfig `json:"pmmClientConfig,omitempty"`

	// pmm server config
	PMMServerConfig *LocationsItems0PMMServerConfig `json:"pmmServerConfig,omitempty"`

	// s3 config
	S3Config *LocationsItems0S3Config `json:"s3Config,omitempty"`
}

// Validate validates this locations items0
func (o *LocationsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePMMClientConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePMMServerConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateS3Config(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LocationsItems0) validatePMMClientConfig(formats strfmt.Registry) error {

	if swag.IsZero(o.PMMClientConfig) { // not required
		return nil
	}

	if o.PMMClientConfig != nil {
		if err := o.PMMClientConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pmmClientConfig")
			}
			return err
		}
	}

	return nil
}

func (o *LocationsItems0) validatePMMServerConfig(formats strfmt.Registry) error {

	if swag.IsZero(o.PMMServerConfig) { // not required
		return nil
	}

	if o.PMMServerConfig != nil {
		if err := o.PMMServerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pmmServerConfig")
			}
			return err
		}
	}

	return nil
}

func (o *LocationsItems0) validateS3Config(formats strfmt.Registry) error {

	if swag.IsZero(o.S3Config) { // not required
		return nil
	}

	if o.S3Config != nil {
		if err := o.S3Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3Config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LocationsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LocationsItems0) UnmarshalBinary(b []byte) error {
	var res LocationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LocationsItems0PMMClientConfig PMMClientLocationConfig represents file system config inside pmm-client.
swagger:model LocationsItems0PMMClientConfig
*/
type LocationsItems0PMMClientConfig struct {

	// path
	Path string `json:"path,omitempty"`
}

// Validate validates this locations items0 PMM client config
func (o *LocationsItems0PMMClientConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LocationsItems0PMMClientConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LocationsItems0PMMClientConfig) UnmarshalBinary(b []byte) error {
	var res LocationsItems0PMMClientConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LocationsItems0PMMServerConfig PMMServerLocationConfig represents file system config inside pmm-server.
swagger:model LocationsItems0PMMServerConfig
*/
type LocationsItems0PMMServerConfig struct {

	// path
	Path string `json:"path,omitempty"`
}

// Validate validates this locations items0 PMM server config
func (o *LocationsItems0PMMServerConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LocationsItems0PMMServerConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LocationsItems0PMMServerConfig) UnmarshalBinary(b []byte) error {
	var res LocationsItems0PMMServerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LocationsItems0S3Config S3LocationConfig represents S3 bucket configuration.
swagger:model LocationsItems0S3Config
*/
type LocationsItems0S3Config struct {

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// access key
	AccessKey string `json:"accessKey,omitempty"`

	// secret key
	SecretKey string `json:"secretKey,omitempty"`

	// bucket name
	BucketName string `json:"bucketName,omitempty"`
}

// Validate validates this locations items0 s3 config
func (o *LocationsItems0S3Config) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LocationsItems0S3Config) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LocationsItems0S3Config) UnmarshalBinary(b []byte) error {
	var res LocationsItems0S3Config
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
