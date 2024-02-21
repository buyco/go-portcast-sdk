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

// checks if the GetScac200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetScac200ResponseInner{}

// GetScac200ResponseInner struct for GetScac200ResponseInner
type GetScac200ResponseInner struct {
	// BL Prefixes observed for the carrier
	BlPrefixes []string `json:"bl_prefixes,omitempty"`
	// Full Name of the Carrier
	FullName *string `json:"full_name,omitempty"`
	// Name of the Carrier
	Name *string `json:"name,omitempty"`
	// Official Carrier SCAC Code
	Scac *string `json:"scac,omitempty"`
	// Other SCAC Synonyms for the carrier (If upload operation is done using any of these, we will convert it to the official SCAC)
	ScacSynonyms []string `json:"scac_synonyms,omitempty"`
	// Supported Status for the carrier, will highlight here if there are any gaps for a particular carrier's coverage.
	SupportedStatus NullableString `json:"supported_status,omitempty"`
	// Last Updated Date for the carrier scac metadata
	Updated *time.Time `json:"updated,omitempty"`
}

// NewGetScac200ResponseInner instantiates a new GetScac200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetScac200ResponseInner() *GetScac200ResponseInner {
	this := GetScac200ResponseInner{}
	return &this
}

// NewGetScac200ResponseInnerWithDefaults instantiates a new GetScac200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetScac200ResponseInnerWithDefaults() *GetScac200ResponseInner {
	this := GetScac200ResponseInner{}
	return &this
}

// GetBlPrefixes returns the BlPrefixes field value if set, zero value otherwise.
func (o *GetScac200ResponseInner) GetBlPrefixes() []string {
	if o == nil || IsNil(o.BlPrefixes) {
		var ret []string
		return ret
	}
	return o.BlPrefixes
}

// GetBlPrefixesOk returns a tuple with the BlPrefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScac200ResponseInner) GetBlPrefixesOk() ([]string, bool) {
	if o == nil || IsNil(o.BlPrefixes) {
		return nil, false
	}
	return o.BlPrefixes, true
}

// HasBlPrefixes returns a boolean if a field has been set.
func (o *GetScac200ResponseInner) HasBlPrefixes() bool {
	if o != nil && !IsNil(o.BlPrefixes) {
		return true
	}

	return false
}

// SetBlPrefixes gets a reference to the given []string and assigns it to the BlPrefixes field.
func (o *GetScac200ResponseInner) SetBlPrefixes(v []string) {
	o.BlPrefixes = v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *GetScac200ResponseInner) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScac200ResponseInner) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *GetScac200ResponseInner) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *GetScac200ResponseInner) SetFullName(v string) {
	o.FullName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetScac200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScac200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetScac200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetScac200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetScac returns the Scac field value if set, zero value otherwise.
func (o *GetScac200ResponseInner) GetScac() string {
	if o == nil || IsNil(o.Scac) {
		var ret string
		return ret
	}
	return *o.Scac
}

// GetScacOk returns a tuple with the Scac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScac200ResponseInner) GetScacOk() (*string, bool) {
	if o == nil || IsNil(o.Scac) {
		return nil, false
	}
	return o.Scac, true
}

// HasScac returns a boolean if a field has been set.
func (o *GetScac200ResponseInner) HasScac() bool {
	if o != nil && !IsNil(o.Scac) {
		return true
	}

	return false
}

// SetScac gets a reference to the given string and assigns it to the Scac field.
func (o *GetScac200ResponseInner) SetScac(v string) {
	o.Scac = &v
}

// GetScacSynonyms returns the ScacSynonyms field value if set, zero value otherwise.
func (o *GetScac200ResponseInner) GetScacSynonyms() []string {
	if o == nil || IsNil(o.ScacSynonyms) {
		var ret []string
		return ret
	}
	return o.ScacSynonyms
}

// GetScacSynonymsOk returns a tuple with the ScacSynonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScac200ResponseInner) GetScacSynonymsOk() ([]string, bool) {
	if o == nil || IsNil(o.ScacSynonyms) {
		return nil, false
	}
	return o.ScacSynonyms, true
}

// HasScacSynonyms returns a boolean if a field has been set.
func (o *GetScac200ResponseInner) HasScacSynonyms() bool {
	if o != nil && !IsNil(o.ScacSynonyms) {
		return true
	}

	return false
}

// SetScacSynonyms gets a reference to the given []string and assigns it to the ScacSynonyms field.
func (o *GetScac200ResponseInner) SetScacSynonyms(v []string) {
	o.ScacSynonyms = v
}

// GetSupportedStatus returns the SupportedStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetScac200ResponseInner) GetSupportedStatus() string {
	if o == nil || IsNil(o.SupportedStatus.Get()) {
		var ret string
		return ret
	}
	return *o.SupportedStatus.Get()
}

// GetSupportedStatusOk returns a tuple with the SupportedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetScac200ResponseInner) GetSupportedStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportedStatus.Get(), o.SupportedStatus.IsSet()
}

// HasSupportedStatus returns a boolean if a field has been set.
func (o *GetScac200ResponseInner) HasSupportedStatus() bool {
	if o != nil && o.SupportedStatus.IsSet() {
		return true
	}

	return false
}

// SetSupportedStatus gets a reference to the given NullableString and assigns it to the SupportedStatus field.
func (o *GetScac200ResponseInner) SetSupportedStatus(v string) {
	o.SupportedStatus.Set(&v)
}

// SetSupportedStatusNil sets the value for SupportedStatus to be an explicit nil
func (o *GetScac200ResponseInner) SetSupportedStatusNil() {
	o.SupportedStatus.Set(nil)
}

// UnsetSupportedStatus ensures that no value is present for SupportedStatus, not even an explicit nil
func (o *GetScac200ResponseInner) UnsetSupportedStatus() {
	o.SupportedStatus.Unset()
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *GetScac200ResponseInner) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetScac200ResponseInner) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *GetScac200ResponseInner) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *GetScac200ResponseInner) SetUpdated(v time.Time) {
	o.Updated = &v
}

func (o GetScac200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetScac200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlPrefixes) {
		toSerialize["bl_prefixes"] = o.BlPrefixes
	}
	if !IsNil(o.FullName) {
		toSerialize["full_name"] = o.FullName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Scac) {
		toSerialize["scac"] = o.Scac
	}
	if !IsNil(o.ScacSynonyms) {
		toSerialize["scac_synonyms"] = o.ScacSynonyms
	}
	if o.SupportedStatus.IsSet() {
		toSerialize["supported_status"] = o.SupportedStatus.Get()
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	return toSerialize, nil
}

type NullableGetScac200ResponseInner struct {
	value *GetScac200ResponseInner
	isSet bool
}

func (v NullableGetScac200ResponseInner) Get() *GetScac200ResponseInner {
	return v.value
}

func (v *NullableGetScac200ResponseInner) Set(val *GetScac200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetScac200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetScac200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetScac200ResponseInner(val *GetScac200ResponseInner) *NullableGetScac200ResponseInner {
	return &NullableGetScac200ResponseInner{value: val, isSet: true}
}

func (v NullableGetScac200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetScac200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
