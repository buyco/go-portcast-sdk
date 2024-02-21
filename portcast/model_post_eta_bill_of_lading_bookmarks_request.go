/*
Container Tracking API

This documentation defines how the users can utilize the Portcast API to retrieve the latest Track and Trace Data for any container, bill of lading or booking across 100+ carriers and 2500+ ports across the globe.  What's changed in the version 2.0? Find out more about it [here](https://support.portcast.io/support/solutions/articles/151000074289-api-release-notes-april-2023-)!!

API version: 2.0
Contact: support@portcast.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package portcast

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PostEtaBillOfLadingBookmarksRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostEtaBillOfLadingBookmarksRequest{}

// PostEtaBillOfLadingBookmarksRequest Upload Parameters
type PostEtaBillOfLadingBookmarksRequest struct {
	// Carrier SCAC Code or Carrier Synonym (AUTO Detection Supported)
	CarrierNo string `json:"carrier_no"`
	// Carrier Provided Master Bill of Lading Number or Booking Number [Leave blank if not available]
	BlNo *string `json:"bl_no,omitempty"`
	// Container Number (ISO6346)
	CntrNo string `json:"cntr_no"`
	// A https endpoint for Portcast to push the json object whenever there are updates.
	CallbackUrl *string `json:"callback_url,omitempty"`
}

type _PostEtaBillOfLadingBookmarksRequest PostEtaBillOfLadingBookmarksRequest

// NewPostEtaBillOfLadingBookmarksRequest instantiates a new PostEtaBillOfLadingBookmarksRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostEtaBillOfLadingBookmarksRequest(carrierNo string, cntrNo string) *PostEtaBillOfLadingBookmarksRequest {
	this := PostEtaBillOfLadingBookmarksRequest{}
	this.CarrierNo = carrierNo
	this.CntrNo = cntrNo
	return &this
}

// NewPostEtaBillOfLadingBookmarksRequestWithDefaults instantiates a new PostEtaBillOfLadingBookmarksRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostEtaBillOfLadingBookmarksRequestWithDefaults() *PostEtaBillOfLadingBookmarksRequest {
	this := PostEtaBillOfLadingBookmarksRequest{}
	return &this
}

// GetCarrierNo returns the CarrierNo field value
func (o *PostEtaBillOfLadingBookmarksRequest) GetCarrierNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierNo
}

// GetCarrierNoOk returns a tuple with the CarrierNo field value
// and a boolean to check if the value has been set.
func (o *PostEtaBillOfLadingBookmarksRequest) GetCarrierNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierNo, true
}

// SetCarrierNo sets field value
func (o *PostEtaBillOfLadingBookmarksRequest) SetCarrierNo(v string) {
	o.CarrierNo = v
}

// GetBlNo returns the BlNo field value if set, zero value otherwise.
func (o *PostEtaBillOfLadingBookmarksRequest) GetBlNo() string {
	if o == nil || IsNil(o.BlNo) {
		var ret string
		return ret
	}
	return *o.BlNo
}

// GetBlNoOk returns a tuple with the BlNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostEtaBillOfLadingBookmarksRequest) GetBlNoOk() (*string, bool) {
	if o == nil || IsNil(o.BlNo) {
		return nil, false
	}
	return o.BlNo, true
}

// HasBlNo returns a boolean if a field has been set.
func (o *PostEtaBillOfLadingBookmarksRequest) HasBlNo() bool {
	if o != nil && !IsNil(o.BlNo) {
		return true
	}

	return false
}

// SetBlNo gets a reference to the given string and assigns it to the BlNo field.
func (o *PostEtaBillOfLadingBookmarksRequest) SetBlNo(v string) {
	o.BlNo = &v
}

// GetCntrNo returns the CntrNo field value
func (o *PostEtaBillOfLadingBookmarksRequest) GetCntrNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CntrNo
}

// GetCntrNoOk returns a tuple with the CntrNo field value
// and a boolean to check if the value has been set.
func (o *PostEtaBillOfLadingBookmarksRequest) GetCntrNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CntrNo, true
}

// SetCntrNo sets field value
func (o *PostEtaBillOfLadingBookmarksRequest) SetCntrNo(v string) {
	o.CntrNo = v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *PostEtaBillOfLadingBookmarksRequest) GetCallbackUrl() string {
	if o == nil || IsNil(o.CallbackUrl) {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostEtaBillOfLadingBookmarksRequest) GetCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackUrl) {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *PostEtaBillOfLadingBookmarksRequest) HasCallbackUrl() bool {
	if o != nil && !IsNil(o.CallbackUrl) {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *PostEtaBillOfLadingBookmarksRequest) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

func (o PostEtaBillOfLadingBookmarksRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostEtaBillOfLadingBookmarksRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["carrier_no"] = o.CarrierNo
	if !IsNil(o.BlNo) {
		toSerialize["bl_no"] = o.BlNo
	}
	toSerialize["cntr_no"] = o.CntrNo
	if !IsNil(o.CallbackUrl) {
		toSerialize["callback_url"] = o.CallbackUrl
	}
	return toSerialize, nil
}

func (o *PostEtaBillOfLadingBookmarksRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"carrier_no",
		"cntr_no",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostEtaBillOfLadingBookmarksRequest := _PostEtaBillOfLadingBookmarksRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostEtaBillOfLadingBookmarksRequest)

	if err != nil {
		return err
	}

	*o = PostEtaBillOfLadingBookmarksRequest(varPostEtaBillOfLadingBookmarksRequest)

	return err
}

type NullablePostEtaBillOfLadingBookmarksRequest struct {
	value *PostEtaBillOfLadingBookmarksRequest
	isSet bool
}

func (v NullablePostEtaBillOfLadingBookmarksRequest) Get() *PostEtaBillOfLadingBookmarksRequest {
	return v.value
}

func (v *NullablePostEtaBillOfLadingBookmarksRequest) Set(val *PostEtaBillOfLadingBookmarksRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostEtaBillOfLadingBookmarksRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostEtaBillOfLadingBookmarksRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostEtaBillOfLadingBookmarksRequest(val *PostEtaBillOfLadingBookmarksRequest) *NullablePostEtaBillOfLadingBookmarksRequest {
	return &NullablePostEtaBillOfLadingBookmarksRequest{value: val, isSet: true}
}

func (v NullablePostEtaBillOfLadingBookmarksRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostEtaBillOfLadingBookmarksRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
