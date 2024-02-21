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

// checks if the DelayReasonsRotationChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DelayReasonsRotationChange{}

// DelayReasonsRotationChange Detailed view into the vessel schedule changes
type DelayReasonsRotationChange struct {
	// Last recorded schedule information - Ports Order
	ComparisonSchedule []string `json:"comparison_schedule,omitempty"`
	// Latest changed schedule information - Ports Order
	UpdatedSchedule []string `json:"updated_schedule,omitempty"`
}

// NewDelayReasonsRotationChange instantiates a new DelayReasonsRotationChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelayReasonsRotationChange() *DelayReasonsRotationChange {
	this := DelayReasonsRotationChange{}
	return &this
}

// NewDelayReasonsRotationChangeWithDefaults instantiates a new DelayReasonsRotationChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelayReasonsRotationChangeWithDefaults() *DelayReasonsRotationChange {
	this := DelayReasonsRotationChange{}
	return &this
}

// GetComparisonSchedule returns the ComparisonSchedule field value if set, zero value otherwise.
func (o *DelayReasonsRotationChange) GetComparisonSchedule() []string {
	if o == nil || IsNil(o.ComparisonSchedule) {
		var ret []string
		return ret
	}
	return o.ComparisonSchedule
}

// GetComparisonScheduleOk returns a tuple with the ComparisonSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayReasonsRotationChange) GetComparisonScheduleOk() ([]string, bool) {
	if o == nil || IsNil(o.ComparisonSchedule) {
		return nil, false
	}
	return o.ComparisonSchedule, true
}

// HasComparisonSchedule returns a boolean if a field has been set.
func (o *DelayReasonsRotationChange) HasComparisonSchedule() bool {
	if o != nil && !IsNil(o.ComparisonSchedule) {
		return true
	}

	return false
}

// SetComparisonSchedule gets a reference to the given []string and assigns it to the ComparisonSchedule field.
func (o *DelayReasonsRotationChange) SetComparisonSchedule(v []string) {
	o.ComparisonSchedule = v
}

// GetUpdatedSchedule returns the UpdatedSchedule field value if set, zero value otherwise.
func (o *DelayReasonsRotationChange) GetUpdatedSchedule() []string {
	if o == nil || IsNil(o.UpdatedSchedule) {
		var ret []string
		return ret
	}
	return o.UpdatedSchedule
}

// GetUpdatedScheduleOk returns a tuple with the UpdatedSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayReasonsRotationChange) GetUpdatedScheduleOk() ([]string, bool) {
	if o == nil || IsNil(o.UpdatedSchedule) {
		return nil, false
	}
	return o.UpdatedSchedule, true
}

// HasUpdatedSchedule returns a boolean if a field has been set.
func (o *DelayReasonsRotationChange) HasUpdatedSchedule() bool {
	if o != nil && !IsNil(o.UpdatedSchedule) {
		return true
	}

	return false
}

// SetUpdatedSchedule gets a reference to the given []string and assigns it to the UpdatedSchedule field.
func (o *DelayReasonsRotationChange) SetUpdatedSchedule(v []string) {
	o.UpdatedSchedule = v
}

func (o DelayReasonsRotationChange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DelayReasonsRotationChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComparisonSchedule) {
		toSerialize["comparison_schedule"] = o.ComparisonSchedule
	}
	if !IsNil(o.UpdatedSchedule) {
		toSerialize["updated_schedule"] = o.UpdatedSchedule
	}
	return toSerialize, nil
}

type NullableDelayReasonsRotationChange struct {
	value *DelayReasonsRotationChange
	isSet bool
}

func (v NullableDelayReasonsRotationChange) Get() *DelayReasonsRotationChange {
	return v.value
}

func (v *NullableDelayReasonsRotationChange) Set(val *DelayReasonsRotationChange) {
	v.value = val
	v.isSet = true
}

func (v NullableDelayReasonsRotationChange) IsSet() bool {
	return v.isSet
}

func (v *NullableDelayReasonsRotationChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelayReasonsRotationChange(val *DelayReasonsRotationChange) *NullableDelayReasonsRotationChange {
	return &NullableDelayReasonsRotationChange{value: val, isSet: true}
}

func (v NullableDelayReasonsRotationChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelayReasonsRotationChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}