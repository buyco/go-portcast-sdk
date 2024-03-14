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

// checks if the ContainerEventTerminalDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerEventTerminalDetails{}

// ContainerEventTerminalDetails Terminal Level Details for the container event
type ContainerEventTerminalDetails struct {
	// BIC/SMDG Defined Code for the Terminal Location
	TerminalCode *string `json:"terminal_code,omitempty"`
	// Source of Terminal Code: BIC/SMDG
	TerminalCodeSource *string `json:"terminal_code_source,omitempty"`
	// Terminal Name where the container event takes place
	TerminalName *string `json:"terminal_name,omitempty"`
}

// NewContainerEventTerminalDetails instantiates a new ContainerEventTerminalDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerEventTerminalDetails() *ContainerEventTerminalDetails {
	this := ContainerEventTerminalDetails{}
	return &this
}

// NewContainerEventTerminalDetailsWithDefaults instantiates a new ContainerEventTerminalDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerEventTerminalDetailsWithDefaults() *ContainerEventTerminalDetails {
	this := ContainerEventTerminalDetails{}
	return &this
}

// GetTerminalCode returns the TerminalCode field value if set, zero value otherwise.
func (o *ContainerEventTerminalDetails) GetTerminalCode() string {
	if o == nil || IsNil(o.TerminalCode) {
		var ret string
		return ret
	}
	return *o.TerminalCode
}

// GetTerminalCodeOk returns a tuple with the TerminalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerEventTerminalDetails) GetTerminalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TerminalCode) {
		return nil, false
	}
	return o.TerminalCode, true
}

// HasTerminalCode returns a boolean if a field has been set.
func (o *ContainerEventTerminalDetails) HasTerminalCode() bool {
	if o != nil && !IsNil(o.TerminalCode) {
		return true
	}

	return false
}

// SetTerminalCode gets a reference to the given string and assigns it to the TerminalCode field.
func (o *ContainerEventTerminalDetails) SetTerminalCode(v string) {
	o.TerminalCode = &v
}

// GetTerminalCodeSource returns the TerminalCodeSource field value if set, zero value otherwise.
func (o *ContainerEventTerminalDetails) GetTerminalCodeSource() string {
	if o == nil || IsNil(o.TerminalCodeSource) {
		var ret string
		return ret
	}
	return *o.TerminalCodeSource
}

// GetTerminalCodeSourceOk returns a tuple with the TerminalCodeSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerEventTerminalDetails) GetTerminalCodeSourceOk() (*string, bool) {
	if o == nil || IsNil(o.TerminalCodeSource) {
		return nil, false
	}
	return o.TerminalCodeSource, true
}

// HasTerminalCodeSource returns a boolean if a field has been set.
func (o *ContainerEventTerminalDetails) HasTerminalCodeSource() bool {
	if o != nil && !IsNil(o.TerminalCodeSource) {
		return true
	}

	return false
}

// SetTerminalCodeSource gets a reference to the given string and assigns it to the TerminalCodeSource field.
func (o *ContainerEventTerminalDetails) SetTerminalCodeSource(v string) {
	o.TerminalCodeSource = &v
}

// GetTerminalName returns the TerminalName field value if set, zero value otherwise.
func (o *ContainerEventTerminalDetails) GetTerminalName() string {
	if o == nil || IsNil(o.TerminalName) {
		var ret string
		return ret
	}
	return *o.TerminalName
}

// GetTerminalNameOk returns a tuple with the TerminalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerEventTerminalDetails) GetTerminalNameOk() (*string, bool) {
	if o == nil || IsNil(o.TerminalName) {
		return nil, false
	}
	return o.TerminalName, true
}

// HasTerminalName returns a boolean if a field has been set.
func (o *ContainerEventTerminalDetails) HasTerminalName() bool {
	if o != nil && !IsNil(o.TerminalName) {
		return true
	}

	return false
}

// SetTerminalName gets a reference to the given string and assigns it to the TerminalName field.
func (o *ContainerEventTerminalDetails) SetTerminalName(v string) {
	o.TerminalName = &v
}

func (o ContainerEventTerminalDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerEventTerminalDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TerminalCode) {
		toSerialize["terminal_code"] = o.TerminalCode
	}
	if !IsNil(o.TerminalCodeSource) {
		toSerialize["terminal_code_source"] = o.TerminalCodeSource
	}
	if !IsNil(o.TerminalName) {
		toSerialize["terminal_name"] = o.TerminalName
	}
	return toSerialize, nil
}

type NullableContainerEventTerminalDetails struct {
	value *ContainerEventTerminalDetails
	isSet bool
}

func (v NullableContainerEventTerminalDetails) Get() *ContainerEventTerminalDetails {
	return v.value
}

func (v *NullableContainerEventTerminalDetails) Set(val *ContainerEventTerminalDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerEventTerminalDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerEventTerminalDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerEventTerminalDetails(val *ContainerEventTerminalDetails) *NullableContainerEventTerminalDetails {
	return &NullableContainerEventTerminalDetails{value: val, isSet: true}
}

func (v NullableContainerEventTerminalDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerEventTerminalDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


