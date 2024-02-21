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

// checks if the PostBookingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostBookingRequest{}

// PostBookingRequest struct for PostBookingRequest
type PostBookingRequest struct {
	// Carrier SCAC Code or Carrier Synonym (AUTO Detection Supported)
	CarrierNo string `json:"carrier_no"`
	// Booking or Bill of Lading number
	DocNo string `json:"doc_no"`
	// Document number type: BK (Booking) or BL (Bill of Lading)
	DocType string `json:"doc_type"`
	// A https endpoint for Portcast to push the json object whenever there are updates.
	CallbackUrl *string `json:"callback_url,omitempty"`
}

type _PostBookingRequest PostBookingRequest

// NewPostBookingRequest instantiates a new PostBookingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostBookingRequest(carrierNo string, docNo string, docType string) *PostBookingRequest {
	this := PostBookingRequest{}
	this.CarrierNo = carrierNo
	this.DocNo = docNo
	this.DocType = docType
	return &this
}

// NewPostBookingRequestWithDefaults instantiates a new PostBookingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostBookingRequestWithDefaults() *PostBookingRequest {
	this := PostBookingRequest{}
	return &this
}

// GetCarrierNo returns the CarrierNo field value
func (o *PostBookingRequest) GetCarrierNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierNo
}

// GetCarrierNoOk returns a tuple with the CarrierNo field value
// and a boolean to check if the value has been set.
func (o *PostBookingRequest) GetCarrierNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierNo, true
}

// SetCarrierNo sets field value
func (o *PostBookingRequest) SetCarrierNo(v string) {
	o.CarrierNo = v
}

// GetDocNo returns the DocNo field value
func (o *PostBookingRequest) GetDocNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocNo
}

// GetDocNoOk returns a tuple with the DocNo field value
// and a boolean to check if the value has been set.
func (o *PostBookingRequest) GetDocNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocNo, true
}

// SetDocNo sets field value
func (o *PostBookingRequest) SetDocNo(v string) {
	o.DocNo = v
}

// GetDocType returns the DocType field value
func (o *PostBookingRequest) GetDocType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocType
}

// GetDocTypeOk returns a tuple with the DocType field value
// and a boolean to check if the value has been set.
func (o *PostBookingRequest) GetDocTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocType, true
}

// SetDocType sets field value
func (o *PostBookingRequest) SetDocType(v string) {
	o.DocType = v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *PostBookingRequest) GetCallbackUrl() string {
	if o == nil || IsNil(o.CallbackUrl) {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostBookingRequest) GetCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackUrl) {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *PostBookingRequest) HasCallbackUrl() bool {
	if o != nil && !IsNil(o.CallbackUrl) {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *PostBookingRequest) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

func (o PostBookingRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostBookingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["carrier_no"] = o.CarrierNo
	toSerialize["doc_no"] = o.DocNo
	toSerialize["doc_type"] = o.DocType
	if !IsNil(o.CallbackUrl) {
		toSerialize["callback_url"] = o.CallbackUrl
	}
	return toSerialize, nil
}

func (o *PostBookingRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"carrier_no",
		"doc_no",
		"doc_type",
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

	varPostBookingRequest := _PostBookingRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostBookingRequest)

	if err != nil {
		return err
	}

	*o = PostBookingRequest(varPostBookingRequest)

	return err
}

type NullablePostBookingRequest struct {
	value *PostBookingRequest
	isSet bool
}

func (v NullablePostBookingRequest) Get() *PostBookingRequest {
	return v.value
}

func (v *NullablePostBookingRequest) Set(val *PostBookingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostBookingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostBookingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostBookingRequest(val *PostBookingRequest) *NullablePostBookingRequest {
	return &NullablePostBookingRequest{value: val, isSet: true}
}

func (v NullablePostBookingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostBookingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}