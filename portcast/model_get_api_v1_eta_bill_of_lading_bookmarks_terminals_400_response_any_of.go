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

// checks if the GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf{}

// GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf Error when the Vessel or Port isn't found
type GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf struct {
	// Error when the Vessel or Port isn't found
	Error *string `json:"error,omitempty"`
}

// NewGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf instantiates a new GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf() *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf {
	this := GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf{}
	var error_ string = "Vessel or Port not found"
	this.Error = &error_
	return &this
}

// NewGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOfWithDefaults instantiates a new GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOfWithDefaults() *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf {
	this := GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf{}
	var error_ string = "Vessel or Port not found"
	this.Error = &error_
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf) SetError(v string) {
	o.Error = &v
}

func (o GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf struct {
	value *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf
	isSet bool
}

func (v NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf) Get() *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf {
	return v.value
}

func (v *NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf) Set(val *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf(val *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf) *NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf {
	return &NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf{value: val, isSet: true}
}

func (v NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
