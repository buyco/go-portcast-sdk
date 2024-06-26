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

// checks if the BookingAPIResponseInnerAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookingAPIResponseInnerAnyOf1{}

// BookingAPIResponseInnerAnyOf1 Duplicate Upload Metadata Object
type BookingAPIResponseInnerAnyOf1 struct {
	// Message describing error status
	Message *string `json:"message,omitempty"`
	// List of Bookmark IDs registered for this container
	PreviousIds []string `json:"previous_ids,omitempty"`
	// Status of individual container upload
	Status *string `json:"status,omitempty"`
	// Container Number registered under the BL/Booking
	Container *string `json:"container,omitempty"`
}

// NewBookingAPIResponseInnerAnyOf1 instantiates a new BookingAPIResponseInnerAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookingAPIResponseInnerAnyOf1() *BookingAPIResponseInnerAnyOf1 {
	this := BookingAPIResponseInnerAnyOf1{}
	return &this
}

// NewBookingAPIResponseInnerAnyOf1WithDefaults instantiates a new BookingAPIResponseInnerAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookingAPIResponseInnerAnyOf1WithDefaults() *BookingAPIResponseInnerAnyOf1 {
	this := BookingAPIResponseInnerAnyOf1{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BookingAPIResponseInnerAnyOf1) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingAPIResponseInnerAnyOf1) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BookingAPIResponseInnerAnyOf1) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BookingAPIResponseInnerAnyOf1) SetMessage(v string) {
	o.Message = &v
}

// GetPreviousIds returns the PreviousIds field value if set, zero value otherwise.
func (o *BookingAPIResponseInnerAnyOf1) GetPreviousIds() []string {
	if o == nil || IsNil(o.PreviousIds) {
		var ret []string
		return ret
	}
	return o.PreviousIds
}

// GetPreviousIdsOk returns a tuple with the PreviousIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingAPIResponseInnerAnyOf1) GetPreviousIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PreviousIds) {
		return nil, false
	}
	return o.PreviousIds, true
}

// HasPreviousIds returns a boolean if a field has been set.
func (o *BookingAPIResponseInnerAnyOf1) HasPreviousIds() bool {
	if o != nil && !IsNil(o.PreviousIds) {
		return true
	}

	return false
}

// SetPreviousIds gets a reference to the given []string and assigns it to the PreviousIds field.
func (o *BookingAPIResponseInnerAnyOf1) SetPreviousIds(v []string) {
	o.PreviousIds = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BookingAPIResponseInnerAnyOf1) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingAPIResponseInnerAnyOf1) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BookingAPIResponseInnerAnyOf1) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BookingAPIResponseInnerAnyOf1) SetStatus(v string) {
	o.Status = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *BookingAPIResponseInnerAnyOf1) GetContainer() string {
	if o == nil || IsNil(o.Container) {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingAPIResponseInnerAnyOf1) GetContainerOk() (*string, bool) {
	if o == nil || IsNil(o.Container) {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *BookingAPIResponseInnerAnyOf1) HasContainer() bool {
	if o != nil && !IsNil(o.Container) {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *BookingAPIResponseInnerAnyOf1) SetContainer(v string) {
	o.Container = &v
}

func (o BookingAPIResponseInnerAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookingAPIResponseInnerAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.PreviousIds) {
		toSerialize["previous_ids"] = o.PreviousIds
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Container) {
		toSerialize["container"] = o.Container
	}
	return toSerialize, nil
}

type NullableBookingAPIResponseInnerAnyOf1 struct {
	value *BookingAPIResponseInnerAnyOf1
	isSet bool
}

func (v NullableBookingAPIResponseInnerAnyOf1) Get() *BookingAPIResponseInnerAnyOf1 {
	return v.value
}

func (v *NullableBookingAPIResponseInnerAnyOf1) Set(val *BookingAPIResponseInnerAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableBookingAPIResponseInnerAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableBookingAPIResponseInnerAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookingAPIResponseInnerAnyOf1(val *BookingAPIResponseInnerAnyOf1) *NullableBookingAPIResponseInnerAnyOf1 {
	return &NullableBookingAPIResponseInnerAnyOf1{value: val, isSet: true}
}

func (v NullableBookingAPIResponseInnerAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookingAPIResponseInnerAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


