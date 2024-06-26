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

// checks if the GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1{}

// GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1 Error when the Bookmark ID isn't found
type GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1 struct {
	// Error when the Bookmark ID isn't found
	Message *string `json:"message,omitempty"`
}

// NewGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1 instantiates a new GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1() *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1 {
	this := GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1{}
	var message string = "Invalid bookmark id"
	this.Message = &message
	return &this
}

// NewGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1WithDefaults instantiates a new GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1WithDefaults() *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1 {
	this := GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1{}
	var message string = "Invalid bookmark id"
	this.Message = &message
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1) SetMessage(v string) {
	o.Message = &v
}

func (o GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1 struct {
	value *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1
	isSet bool
}

func (v NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1) Get() *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1 {
	return v.value
}

func (v *NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1) Set(val *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1(val *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1) *NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1 {
	return &NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1{value: val, isSet: true}
}

func (v NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


