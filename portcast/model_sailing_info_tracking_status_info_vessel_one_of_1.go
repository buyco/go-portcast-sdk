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

// checks if the SailingInfoTrackingStatusInfoVesselOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SailingInfoTrackingStatusInfoVesselOneOf1{}

// SailingInfoTrackingStatusInfoVesselOneOf1 struct for SailingInfoTrackingStatusInfoVesselOneOf1
type SailingInfoTrackingStatusInfoVesselOneOf1 struct {
	// Vessel related data (schedules and vessel AIS data) successfully fetched
	Code     *string  `json:"code,omitempty"`
	Metadata []string `json:"metadata,omitempty"`
}

// NewSailingInfoTrackingStatusInfoVesselOneOf1 instantiates a new SailingInfoTrackingStatusInfoVesselOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSailingInfoTrackingStatusInfoVesselOneOf1() *SailingInfoTrackingStatusInfoVesselOneOf1 {
	this := SailingInfoTrackingStatusInfoVesselOneOf1{}
	var code string = "SUCCESS"
	this.Code = &code
	return &this
}

// NewSailingInfoTrackingStatusInfoVesselOneOf1WithDefaults instantiates a new SailingInfoTrackingStatusInfoVesselOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSailingInfoTrackingStatusInfoVesselOneOf1WithDefaults() *SailingInfoTrackingStatusInfoVesselOneOf1 {
	this := SailingInfoTrackingStatusInfoVesselOneOf1{}
	var code string = "SUCCESS"
	this.Code = &code
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SailingInfoTrackingStatusInfoVesselOneOf1) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SailingInfoTrackingStatusInfoVesselOneOf1) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SailingInfoTrackingStatusInfoVesselOneOf1) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SailingInfoTrackingStatusInfoVesselOneOf1) SetCode(v string) {
	o.Code = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SailingInfoTrackingStatusInfoVesselOneOf1) GetMetadata() []string {
	if o == nil || IsNil(o.Metadata) {
		var ret []string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SailingInfoTrackingStatusInfoVesselOneOf1) GetMetadataOk() ([]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SailingInfoTrackingStatusInfoVesselOneOf1) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []string and assigns it to the Metadata field.
func (o *SailingInfoTrackingStatusInfoVesselOneOf1) SetMetadata(v []string) {
	o.Metadata = v
}

func (o SailingInfoTrackingStatusInfoVesselOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SailingInfoTrackingStatusInfoVesselOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableSailingInfoTrackingStatusInfoVesselOneOf1 struct {
	value *SailingInfoTrackingStatusInfoVesselOneOf1
	isSet bool
}

func (v NullableSailingInfoTrackingStatusInfoVesselOneOf1) Get() *SailingInfoTrackingStatusInfoVesselOneOf1 {
	return v.value
}

func (v *NullableSailingInfoTrackingStatusInfoVesselOneOf1) Set(val *SailingInfoTrackingStatusInfoVesselOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableSailingInfoTrackingStatusInfoVesselOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableSailingInfoTrackingStatusInfoVesselOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSailingInfoTrackingStatusInfoVesselOneOf1(val *SailingInfoTrackingStatusInfoVesselOneOf1) *NullableSailingInfoTrackingStatusInfoVesselOneOf1 {
	return &NullableSailingInfoTrackingStatusInfoVesselOneOf1{value: val, isSet: true}
}

func (v NullableSailingInfoTrackingStatusInfoVesselOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSailingInfoTrackingStatusInfoVesselOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}