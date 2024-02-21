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

// checks if the SailingInfoTrackingStatusInfoPredictionOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SailingInfoTrackingStatusInfoPredictionOneOf{}

// SailingInfoTrackingStatusInfoPredictionOneOf struct for SailingInfoTrackingStatusInfoPredictionOneOf
type SailingInfoTrackingStatusInfoPredictionOneOf struct {
	// No status code is assigned yet ~ Vessel hasn't gone through the prediction engine yet.
	Code     *string  `json:"code,omitempty"`
	Metadata []string `json:"metadata,omitempty"`
}

// NewSailingInfoTrackingStatusInfoPredictionOneOf instantiates a new SailingInfoTrackingStatusInfoPredictionOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSailingInfoTrackingStatusInfoPredictionOneOf() *SailingInfoTrackingStatusInfoPredictionOneOf {
	this := SailingInfoTrackingStatusInfoPredictionOneOf{}
	var code string = "INIT"
	this.Code = &code
	return &this
}

// NewSailingInfoTrackingStatusInfoPredictionOneOfWithDefaults instantiates a new SailingInfoTrackingStatusInfoPredictionOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSailingInfoTrackingStatusInfoPredictionOneOfWithDefaults() *SailingInfoTrackingStatusInfoPredictionOneOf {
	this := SailingInfoTrackingStatusInfoPredictionOneOf{}
	var code string = "INIT"
	this.Code = &code
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SailingInfoTrackingStatusInfoPredictionOneOf) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SailingInfoTrackingStatusInfoPredictionOneOf) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SailingInfoTrackingStatusInfoPredictionOneOf) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SailingInfoTrackingStatusInfoPredictionOneOf) SetCode(v string) {
	o.Code = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SailingInfoTrackingStatusInfoPredictionOneOf) GetMetadata() []string {
	if o == nil || IsNil(o.Metadata) {
		var ret []string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SailingInfoTrackingStatusInfoPredictionOneOf) GetMetadataOk() ([]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SailingInfoTrackingStatusInfoPredictionOneOf) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []string and assigns it to the Metadata field.
func (o *SailingInfoTrackingStatusInfoPredictionOneOf) SetMetadata(v []string) {
	o.Metadata = v
}

func (o SailingInfoTrackingStatusInfoPredictionOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SailingInfoTrackingStatusInfoPredictionOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableSailingInfoTrackingStatusInfoPredictionOneOf struct {
	value *SailingInfoTrackingStatusInfoPredictionOneOf
	isSet bool
}

func (v NullableSailingInfoTrackingStatusInfoPredictionOneOf) Get() *SailingInfoTrackingStatusInfoPredictionOneOf {
	return v.value
}

func (v *NullableSailingInfoTrackingStatusInfoPredictionOneOf) Set(val *SailingInfoTrackingStatusInfoPredictionOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSailingInfoTrackingStatusInfoPredictionOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSailingInfoTrackingStatusInfoPredictionOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSailingInfoTrackingStatusInfoPredictionOneOf(val *SailingInfoTrackingStatusInfoPredictionOneOf) *NullableSailingInfoTrackingStatusInfoPredictionOneOf {
	return &NullableSailingInfoTrackingStatusInfoPredictionOneOf{value: val, isSet: true}
}

func (v NullableSailingInfoTrackingStatusInfoPredictionOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSailingInfoTrackingStatusInfoPredictionOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}