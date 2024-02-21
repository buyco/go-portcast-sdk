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
	"time"
)

// checks if the BookingAPIRoutePostpod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookingAPIRoutePostpod{}

// BookingAPIRoutePostpod Post-POD is the location where the container is destuffed
type BookingAPIRoutePostpod struct {
	// Location ID - Refer to location obj for details
	Location NullableInt32 `json:"location,omitempty"`
	// Timestamp on when the first event takes place at the post-POD
	Date NullableTime `json:"date,omitempty"`
	// actual or estimated boolean
	Actual NullableBool `json:"actual,omitempty"`
}

// NewBookingAPIRoutePostpod instantiates a new BookingAPIRoutePostpod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookingAPIRoutePostpod() *BookingAPIRoutePostpod {
	this := BookingAPIRoutePostpod{}
	return &this
}

// NewBookingAPIRoutePostpodWithDefaults instantiates a new BookingAPIRoutePostpod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookingAPIRoutePostpodWithDefaults() *BookingAPIRoutePostpod {
	this := BookingAPIRoutePostpod{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookingAPIRoutePostpod) GetLocation() int32 {
	if o == nil || IsNil(o.Location.Get()) {
		var ret int32
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookingAPIRoutePostpod) GetLocationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *BookingAPIRoutePostpod) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableInt32 and assigns it to the Location field.
func (o *BookingAPIRoutePostpod) SetLocation(v int32) {
	o.Location.Set(&v)
}

// SetLocationNil sets the value for Location to be an explicit nil
func (o *BookingAPIRoutePostpod) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *BookingAPIRoutePostpod) UnsetLocation() {
	o.Location.Unset()
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookingAPIRoutePostpod) GetDate() time.Time {
	if o == nil || IsNil(o.Date.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookingAPIRoutePostpod) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *BookingAPIRoutePostpod) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableTime and assigns it to the Date field.
func (o *BookingAPIRoutePostpod) SetDate(v time.Time) {
	o.Date.Set(&v)
}

// SetDateNil sets the value for Date to be an explicit nil
func (o *BookingAPIRoutePostpod) SetDateNil() {
	o.Date.Set(nil)
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *BookingAPIRoutePostpod) UnsetDate() {
	o.Date.Unset()
}

// GetActual returns the Actual field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookingAPIRoutePostpod) GetActual() bool {
	if o == nil || IsNil(o.Actual.Get()) {
		var ret bool
		return ret
	}
	return *o.Actual.Get()
}

// GetActualOk returns a tuple with the Actual field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookingAPIRoutePostpod) GetActualOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actual.Get(), o.Actual.IsSet()
}

// HasActual returns a boolean if a field has been set.
func (o *BookingAPIRoutePostpod) HasActual() bool {
	if o != nil && o.Actual.IsSet() {
		return true
	}

	return false
}

// SetActual gets a reference to the given NullableBool and assigns it to the Actual field.
func (o *BookingAPIRoutePostpod) SetActual(v bool) {
	o.Actual.Set(&v)
}

// SetActualNil sets the value for Actual to be an explicit nil
func (o *BookingAPIRoutePostpod) SetActualNil() {
	o.Actual.Set(nil)
}

// UnsetActual ensures that no value is present for Actual, not even an explicit nil
func (o *BookingAPIRoutePostpod) UnsetActual() {
	o.Actual.Unset()
}

func (o BookingAPIRoutePostpod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookingAPIRoutePostpod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}
	if o.Actual.IsSet() {
		toSerialize["actual"] = o.Actual.Get()
	}
	return toSerialize, nil
}

type NullableBookingAPIRoutePostpod struct {
	value *BookingAPIRoutePostpod
	isSet bool
}

func (v NullableBookingAPIRoutePostpod) Get() *BookingAPIRoutePostpod {
	return v.value
}

func (v *NullableBookingAPIRoutePostpod) Set(val *BookingAPIRoutePostpod) {
	v.value = val
	v.isSet = true
}

func (v NullableBookingAPIRoutePostpod) IsSet() bool {
	return v.isSet
}

func (v *NullableBookingAPIRoutePostpod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookingAPIRoutePostpod(val *BookingAPIRoutePostpod) *NullableBookingAPIRoutePostpod {
	return &NullableBookingAPIRoutePostpod{value: val, isSet: true}
}

func (v NullableBookingAPIRoutePostpod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookingAPIRoutePostpod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
