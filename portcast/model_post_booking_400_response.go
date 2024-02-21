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

// checks if the PostBooking400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostBooking400Response{}

// PostBooking400Response struct for PostBooking400Response
type PostBooking400Response struct {
	Message *string `json:"message,omitempty"`
}

// NewPostBooking400Response instantiates a new PostBooking400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostBooking400Response() *PostBooking400Response {
	this := PostBooking400Response{}
	return &this
}

// NewPostBooking400ResponseWithDefaults instantiates a new PostBooking400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostBooking400ResponseWithDefaults() *PostBooking400Response {
	this := PostBooking400Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PostBooking400Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostBooking400Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PostBooking400Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PostBooking400Response) SetMessage(v string) {
	o.Message = &v
}

func (o PostBooking400Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostBooking400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullablePostBooking400Response struct {
	value *PostBooking400Response
	isSet bool
}

func (v NullablePostBooking400Response) Get() *PostBooking400Response {
	return v.value
}

func (v *NullablePostBooking400Response) Set(val *PostBooking400Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostBooking400Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostBooking400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostBooking400Response(val *PostBooking400Response) *NullablePostBooking400Response {
	return &NullablePostBooking400Response{value: val, isSet: true}
}

func (v NullablePostBooking400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostBooking400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}