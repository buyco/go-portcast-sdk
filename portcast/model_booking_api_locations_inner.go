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

// checks if the BookingAPILocationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookingAPILocationsInner{}

// BookingAPILocationsInner struct for BookingAPILocationsInner
type BookingAPILocationsInner struct {
	// Unique object ID
	Id *int32 `json:"id,omitempty"`
	// Location Name
	Name NullableString `json:"name,omitempty"`
	// Location State
	State NullableString `json:"state,omitempty"`
	// Location Country
	Country NullableString `json:"country,omitempty"`
	// Country Code as per UN LOCODEs
	CountryCode NullableString `json:"country_code,omitempty"`
	// UN assigned LOCODE for the location
	Locode NullableString `json:"locode,omitempty"`
	// Latitude
	Lat NullableFloat32 `json:"lat,omitempty"`
	// Longitude
	Lng NullableFloat32 `json:"lng,omitempty"`
}

// NewBookingAPILocationsInner instantiates a new BookingAPILocationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookingAPILocationsInner() *BookingAPILocationsInner {
	this := BookingAPILocationsInner{}
	return &this
}

// NewBookingAPILocationsInnerWithDefaults instantiates a new BookingAPILocationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookingAPILocationsInnerWithDefaults() *BookingAPILocationsInner {
	this := BookingAPILocationsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BookingAPILocationsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingAPILocationsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BookingAPILocationsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BookingAPILocationsInner) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookingAPILocationsInner) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookingAPILocationsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *BookingAPILocationsInner) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *BookingAPILocationsInner) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *BookingAPILocationsInner) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *BookingAPILocationsInner) UnsetName() {
	o.Name.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookingAPILocationsInner) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookingAPILocationsInner) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *BookingAPILocationsInner) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *BookingAPILocationsInner) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *BookingAPILocationsInner) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *BookingAPILocationsInner) UnsetState() {
	o.State.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookingAPILocationsInner) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookingAPILocationsInner) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *BookingAPILocationsInner) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *BookingAPILocationsInner) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *BookingAPILocationsInner) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *BookingAPILocationsInner) UnsetCountry() {
	o.Country.Unset()
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookingAPILocationsInner) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode.Get()) {
		var ret string
		return ret
	}
	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookingAPILocationsInner) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// HasCountryCode returns a boolean if a field has been set.
func (o *BookingAPILocationsInner) HasCountryCode() bool {
	if o != nil && o.CountryCode.IsSet() {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given NullableString and assigns it to the CountryCode field.
func (o *BookingAPILocationsInner) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}
// SetCountryCodeNil sets the value for CountryCode to be an explicit nil
func (o *BookingAPILocationsInner) SetCountryCodeNil() {
	o.CountryCode.Set(nil)
}

// UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
func (o *BookingAPILocationsInner) UnsetCountryCode() {
	o.CountryCode.Unset()
}

// GetLocode returns the Locode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookingAPILocationsInner) GetLocode() string {
	if o == nil || IsNil(o.Locode.Get()) {
		var ret string
		return ret
	}
	return *o.Locode.Get()
}

// GetLocodeOk returns a tuple with the Locode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookingAPILocationsInner) GetLocodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locode.Get(), o.Locode.IsSet()
}

// HasLocode returns a boolean if a field has been set.
func (o *BookingAPILocationsInner) HasLocode() bool {
	if o != nil && o.Locode.IsSet() {
		return true
	}

	return false
}

// SetLocode gets a reference to the given NullableString and assigns it to the Locode field.
func (o *BookingAPILocationsInner) SetLocode(v string) {
	o.Locode.Set(&v)
}
// SetLocodeNil sets the value for Locode to be an explicit nil
func (o *BookingAPILocationsInner) SetLocodeNil() {
	o.Locode.Set(nil)
}

// UnsetLocode ensures that no value is present for Locode, not even an explicit nil
func (o *BookingAPILocationsInner) UnsetLocode() {
	o.Locode.Unset()
}

// GetLat returns the Lat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookingAPILocationsInner) GetLat() float32 {
	if o == nil || IsNil(o.Lat.Get()) {
		var ret float32
		return ret
	}
	return *o.Lat.Get()
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookingAPILocationsInner) GetLatOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lat.Get(), o.Lat.IsSet()
}

// HasLat returns a boolean if a field has been set.
func (o *BookingAPILocationsInner) HasLat() bool {
	if o != nil && o.Lat.IsSet() {
		return true
	}

	return false
}

// SetLat gets a reference to the given NullableFloat32 and assigns it to the Lat field.
func (o *BookingAPILocationsInner) SetLat(v float32) {
	o.Lat.Set(&v)
}
// SetLatNil sets the value for Lat to be an explicit nil
func (o *BookingAPILocationsInner) SetLatNil() {
	o.Lat.Set(nil)
}

// UnsetLat ensures that no value is present for Lat, not even an explicit nil
func (o *BookingAPILocationsInner) UnsetLat() {
	o.Lat.Unset()
}

// GetLng returns the Lng field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookingAPILocationsInner) GetLng() float32 {
	if o == nil || IsNil(o.Lng.Get()) {
		var ret float32
		return ret
	}
	return *o.Lng.Get()
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookingAPILocationsInner) GetLngOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lng.Get(), o.Lng.IsSet()
}

// HasLng returns a boolean if a field has been set.
func (o *BookingAPILocationsInner) HasLng() bool {
	if o != nil && o.Lng.IsSet() {
		return true
	}

	return false
}

// SetLng gets a reference to the given NullableFloat32 and assigns it to the Lng field.
func (o *BookingAPILocationsInner) SetLng(v float32) {
	o.Lng.Set(&v)
}
// SetLngNil sets the value for Lng to be an explicit nil
func (o *BookingAPILocationsInner) SetLngNil() {
	o.Lng.Set(nil)
}

// UnsetLng ensures that no value is present for Lng, not even an explicit nil
func (o *BookingAPILocationsInner) UnsetLng() {
	o.Lng.Unset()
}

func (o BookingAPILocationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookingAPILocationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.CountryCode.IsSet() {
		toSerialize["country_code"] = o.CountryCode.Get()
	}
	if o.Locode.IsSet() {
		toSerialize["locode"] = o.Locode.Get()
	}
	if o.Lat.IsSet() {
		toSerialize["lat"] = o.Lat.Get()
	}
	if o.Lng.IsSet() {
		toSerialize["lng"] = o.Lng.Get()
	}
	return toSerialize, nil
}

type NullableBookingAPILocationsInner struct {
	value *BookingAPILocationsInner
	isSet bool
}

func (v NullableBookingAPILocationsInner) Get() *BookingAPILocationsInner {
	return v.value
}

func (v *NullableBookingAPILocationsInner) Set(val *BookingAPILocationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBookingAPILocationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBookingAPILocationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookingAPILocationsInner(val *BookingAPILocationsInner) *NullableBookingAPILocationsInner {
	return &NullableBookingAPILocationsInner{value: val, isSet: true}
}

func (v NullableBookingAPILocationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookingAPILocationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


