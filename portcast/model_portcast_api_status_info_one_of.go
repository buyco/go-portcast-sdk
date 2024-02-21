/*
Container Tracking API

This documentation defines how the users can utilize the Portcast API to retrieve the latest Track and Trace Data for any container, bill of lading or booking across 100+ carriers and 2500+ ports across the globe.  What's changed in the version 2.0? Find out more about it [here](https://support.portcast.io/support/solutions/articles/151000074289-api-release-notes-april-2023-)!!

API version: 2.0
Contact: support@portcast.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package portcast

import (
	"encoding/json"
)

// checks if the PortcastAPIStatusInfoOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortcastAPIStatusInfoOneOf{}

// PortcastAPIStatusInfoOneOf struct for PortcastAPIStatusInfoOneOf
type PortcastAPIStatusInfoOneOf struct {
	// We didn't encounter any warnings or errors in any of Container Status Code Status, Vessel Status Code and Predictions Status Code or while generating the API Response
	Code     *string  `json:"code,omitempty"`
	Metadata []string `json:"metadata,omitempty"`
}

// NewPortcastAPIStatusInfoOneOf instantiates a new PortcastAPIStatusInfoOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortcastAPIStatusInfoOneOf() *PortcastAPIStatusInfoOneOf {
	this := PortcastAPIStatusInfoOneOf{}
	var code string = "SUCCESS"
	this.Code = &code
	return &this
}

// NewPortcastAPIStatusInfoOneOfWithDefaults instantiates a new PortcastAPIStatusInfoOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortcastAPIStatusInfoOneOfWithDefaults() *PortcastAPIStatusInfoOneOf {
	this := PortcastAPIStatusInfoOneOf{}
	var code string = "SUCCESS"
	this.Code = &code
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *PortcastAPIStatusInfoOneOf) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortcastAPIStatusInfoOneOf) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PortcastAPIStatusInfoOneOf) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *PortcastAPIStatusInfoOneOf) SetCode(v string) {
	o.Code = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PortcastAPIStatusInfoOneOf) GetMetadata() []string {
	if o == nil || IsNil(o.Metadata) {
		var ret []string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortcastAPIStatusInfoOneOf) GetMetadataOk() ([]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PortcastAPIStatusInfoOneOf) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []string and assigns it to the Metadata field.
func (o *PortcastAPIStatusInfoOneOf) SetMetadata(v []string) {
	o.Metadata = v
}

func (o PortcastAPIStatusInfoOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortcastAPIStatusInfoOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullablePortcastAPIStatusInfoOneOf struct {
	value *PortcastAPIStatusInfoOneOf
	isSet bool
}

func (v NullablePortcastAPIStatusInfoOneOf) Get() *PortcastAPIStatusInfoOneOf {
	return v.value
}

func (v *NullablePortcastAPIStatusInfoOneOf) Set(val *PortcastAPIStatusInfoOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePortcastAPIStatusInfoOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePortcastAPIStatusInfoOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortcastAPIStatusInfoOneOf(val *PortcastAPIStatusInfoOneOf) *NullablePortcastAPIStatusInfoOneOf {
	return &NullablePortcastAPIStatusInfoOneOf{value: val, isSet: true}
}

func (v NullablePortcastAPIStatusInfoOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortcastAPIStatusInfoOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
