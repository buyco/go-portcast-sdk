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

// checks if the BookingAPIVesselsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookingAPIVesselsInner{}

// BookingAPIVesselsInner struct for BookingAPIVesselsInner
type BookingAPIVesselsInner struct {
	// Unique object ID
	Id *int32 `json:"id,omitempty"`
	// Vessel Name
	Name NullableString `json:"name,omitempty"`
	// Vessel IMO
	Imo NullableInt32 `json:"imo,omitempty"`
	// Vessel Call Sign
	CallSign NullableString `json:"call_sign,omitempty"`
	// Vessel MMSI Number
	Mmsi NullableInt32 `json:"mmsi,omitempty"`
	// Vessel Country Flag
	Flag NullableString `json:"flag,omitempty"`
}

// NewBookingAPIVesselsInner instantiates a new BookingAPIVesselsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookingAPIVesselsInner() *BookingAPIVesselsInner {
	this := BookingAPIVesselsInner{}
	return &this
}

// NewBookingAPIVesselsInnerWithDefaults instantiates a new BookingAPIVesselsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookingAPIVesselsInnerWithDefaults() *BookingAPIVesselsInner {
	this := BookingAPIVesselsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BookingAPIVesselsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingAPIVesselsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BookingAPIVesselsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BookingAPIVesselsInner) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookingAPIVesselsInner) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookingAPIVesselsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *BookingAPIVesselsInner) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *BookingAPIVesselsInner) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *BookingAPIVesselsInner) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *BookingAPIVesselsInner) UnsetName() {
	o.Name.Unset()
}

// GetImo returns the Imo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookingAPIVesselsInner) GetImo() int32 {
	if o == nil || IsNil(o.Imo.Get()) {
		var ret int32
		return ret
	}
	return *o.Imo.Get()
}

// GetImoOk returns a tuple with the Imo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookingAPIVesselsInner) GetImoOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Imo.Get(), o.Imo.IsSet()
}

// HasImo returns a boolean if a field has been set.
func (o *BookingAPIVesselsInner) HasImo() bool {
	if o != nil && o.Imo.IsSet() {
		return true
	}

	return false
}

// SetImo gets a reference to the given NullableInt32 and assigns it to the Imo field.
func (o *BookingAPIVesselsInner) SetImo(v int32) {
	o.Imo.Set(&v)
}

// SetImoNil sets the value for Imo to be an explicit nil
func (o *BookingAPIVesselsInner) SetImoNil() {
	o.Imo.Set(nil)
}

// UnsetImo ensures that no value is present for Imo, not even an explicit nil
func (o *BookingAPIVesselsInner) UnsetImo() {
	o.Imo.Unset()
}

// GetCallSign returns the CallSign field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookingAPIVesselsInner) GetCallSign() string {
	if o == nil || IsNil(o.CallSign.Get()) {
		var ret string
		return ret
	}
	return *o.CallSign.Get()
}

// GetCallSignOk returns a tuple with the CallSign field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookingAPIVesselsInner) GetCallSignOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CallSign.Get(), o.CallSign.IsSet()
}

// HasCallSign returns a boolean if a field has been set.
func (o *BookingAPIVesselsInner) HasCallSign() bool {
	if o != nil && o.CallSign.IsSet() {
		return true
	}

	return false
}

// SetCallSign gets a reference to the given NullableString and assigns it to the CallSign field.
func (o *BookingAPIVesselsInner) SetCallSign(v string) {
	o.CallSign.Set(&v)
}

// SetCallSignNil sets the value for CallSign to be an explicit nil
func (o *BookingAPIVesselsInner) SetCallSignNil() {
	o.CallSign.Set(nil)
}

// UnsetCallSign ensures that no value is present for CallSign, not even an explicit nil
func (o *BookingAPIVesselsInner) UnsetCallSign() {
	o.CallSign.Unset()
}

// GetMmsi returns the Mmsi field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookingAPIVesselsInner) GetMmsi() int32 {
	if o == nil || IsNil(o.Mmsi.Get()) {
		var ret int32
		return ret
	}
	return *o.Mmsi.Get()
}

// GetMmsiOk returns a tuple with the Mmsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookingAPIVesselsInner) GetMmsiOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mmsi.Get(), o.Mmsi.IsSet()
}

// HasMmsi returns a boolean if a field has been set.
func (o *BookingAPIVesselsInner) HasMmsi() bool {
	if o != nil && o.Mmsi.IsSet() {
		return true
	}

	return false
}

// SetMmsi gets a reference to the given NullableInt32 and assigns it to the Mmsi field.
func (o *BookingAPIVesselsInner) SetMmsi(v int32) {
	o.Mmsi.Set(&v)
}

// SetMmsiNil sets the value for Mmsi to be an explicit nil
func (o *BookingAPIVesselsInner) SetMmsiNil() {
	o.Mmsi.Set(nil)
}

// UnsetMmsi ensures that no value is present for Mmsi, not even an explicit nil
func (o *BookingAPIVesselsInner) UnsetMmsi() {
	o.Mmsi.Unset()
}

// GetFlag returns the Flag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookingAPIVesselsInner) GetFlag() string {
	if o == nil || IsNil(o.Flag.Get()) {
		var ret string
		return ret
	}
	return *o.Flag.Get()
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookingAPIVesselsInner) GetFlagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Flag.Get(), o.Flag.IsSet()
}

// HasFlag returns a boolean if a field has been set.
func (o *BookingAPIVesselsInner) HasFlag() bool {
	if o != nil && o.Flag.IsSet() {
		return true
	}

	return false
}

// SetFlag gets a reference to the given NullableString and assigns it to the Flag field.
func (o *BookingAPIVesselsInner) SetFlag(v string) {
	o.Flag.Set(&v)
}

// SetFlagNil sets the value for Flag to be an explicit nil
func (o *BookingAPIVesselsInner) SetFlagNil() {
	o.Flag.Set(nil)
}

// UnsetFlag ensures that no value is present for Flag, not even an explicit nil
func (o *BookingAPIVesselsInner) UnsetFlag() {
	o.Flag.Unset()
}

func (o BookingAPIVesselsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookingAPIVesselsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Imo.IsSet() {
		toSerialize["imo"] = o.Imo.Get()
	}
	if o.CallSign.IsSet() {
		toSerialize["call_sign"] = o.CallSign.Get()
	}
	if o.Mmsi.IsSet() {
		toSerialize["mmsi"] = o.Mmsi.Get()
	}
	if o.Flag.IsSet() {
		toSerialize["flag"] = o.Flag.Get()
	}
	return toSerialize, nil
}

type NullableBookingAPIVesselsInner struct {
	value *BookingAPIVesselsInner
	isSet bool
}

func (v NullableBookingAPIVesselsInner) Get() *BookingAPIVesselsInner {
	return v.value
}

func (v *NullableBookingAPIVesselsInner) Set(val *BookingAPIVesselsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBookingAPIVesselsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBookingAPIVesselsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookingAPIVesselsInner(val *BookingAPIVesselsInner) *NullableBookingAPIVesselsInner {
	return &NullableBookingAPIVesselsInner{value: val, isSet: true}
}

func (v NullableBookingAPIVesselsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookingAPIVesselsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
