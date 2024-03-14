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

// checks if the SailingInfoTrackingStatusInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SailingInfoTrackingStatusInfo{}

// SailingInfoTrackingStatusInfo Status Defination for Vessel Schedules and Portcast Predictions for this leg of the journey
type SailingInfoTrackingStatusInfo struct {
	Prediction *SailingInfoTrackingStatusInfoPrediction `json:"prediction,omitempty"`
	Vessel *SailingInfoTrackingStatusInfoVessel `json:"vessel,omitempty"`
}

// NewSailingInfoTrackingStatusInfo instantiates a new SailingInfoTrackingStatusInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSailingInfoTrackingStatusInfo() *SailingInfoTrackingStatusInfo {
	this := SailingInfoTrackingStatusInfo{}
	return &this
}

// NewSailingInfoTrackingStatusInfoWithDefaults instantiates a new SailingInfoTrackingStatusInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSailingInfoTrackingStatusInfoWithDefaults() *SailingInfoTrackingStatusInfo {
	this := SailingInfoTrackingStatusInfo{}
	return &this
}

// GetPrediction returns the Prediction field value if set, zero value otherwise.
func (o *SailingInfoTrackingStatusInfo) GetPrediction() SailingInfoTrackingStatusInfoPrediction {
	if o == nil || IsNil(o.Prediction) {
		var ret SailingInfoTrackingStatusInfoPrediction
		return ret
	}
	return *o.Prediction
}

// GetPredictionOk returns a tuple with the Prediction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SailingInfoTrackingStatusInfo) GetPredictionOk() (*SailingInfoTrackingStatusInfoPrediction, bool) {
	if o == nil || IsNil(o.Prediction) {
		return nil, false
	}
	return o.Prediction, true
}

// HasPrediction returns a boolean if a field has been set.
func (o *SailingInfoTrackingStatusInfo) HasPrediction() bool {
	if o != nil && !IsNil(o.Prediction) {
		return true
	}

	return false
}

// SetPrediction gets a reference to the given SailingInfoTrackingStatusInfoPrediction and assigns it to the Prediction field.
func (o *SailingInfoTrackingStatusInfo) SetPrediction(v SailingInfoTrackingStatusInfoPrediction) {
	o.Prediction = &v
}

// GetVessel returns the Vessel field value if set, zero value otherwise.
func (o *SailingInfoTrackingStatusInfo) GetVessel() SailingInfoTrackingStatusInfoVessel {
	if o == nil || IsNil(o.Vessel) {
		var ret SailingInfoTrackingStatusInfoVessel
		return ret
	}
	return *o.Vessel
}

// GetVesselOk returns a tuple with the Vessel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SailingInfoTrackingStatusInfo) GetVesselOk() (*SailingInfoTrackingStatusInfoVessel, bool) {
	if o == nil || IsNil(o.Vessel) {
		return nil, false
	}
	return o.Vessel, true
}

// HasVessel returns a boolean if a field has been set.
func (o *SailingInfoTrackingStatusInfo) HasVessel() bool {
	if o != nil && !IsNil(o.Vessel) {
		return true
	}

	return false
}

// SetVessel gets a reference to the given SailingInfoTrackingStatusInfoVessel and assigns it to the Vessel field.
func (o *SailingInfoTrackingStatusInfo) SetVessel(v SailingInfoTrackingStatusInfoVessel) {
	o.Vessel = &v
}

func (o SailingInfoTrackingStatusInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SailingInfoTrackingStatusInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prediction) {
		toSerialize["prediction"] = o.Prediction
	}
	if !IsNil(o.Vessel) {
		toSerialize["vessel"] = o.Vessel
	}
	return toSerialize, nil
}

type NullableSailingInfoTrackingStatusInfo struct {
	value *SailingInfoTrackingStatusInfo
	isSet bool
}

func (v NullableSailingInfoTrackingStatusInfo) Get() *SailingInfoTrackingStatusInfo {
	return v.value
}

func (v *NullableSailingInfoTrackingStatusInfo) Set(val *SailingInfoTrackingStatusInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSailingInfoTrackingStatusInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSailingInfoTrackingStatusInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSailingInfoTrackingStatusInfo(val *SailingInfoTrackingStatusInfo) *NullableSailingInfoTrackingStatusInfo {
	return &NullableSailingInfoTrackingStatusInfo{value: val, isSet: true}
}

func (v NullableSailingInfoTrackingStatusInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSailingInfoTrackingStatusInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


