/*
Booking API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package portcast

import (
	"encoding/json"
)

// RestfulBookingModel struct for RestfulBookingModel
type RestfulBookingModel struct {
	DocType string `json:"doc_type"`
	DocNo string `json:"doc_no"`
	CarrierNo string `json:"carrier_no"`
}

// NewRestfulBookingModel instantiates a new RestfulBookingModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestfulBookingModel(docType string, docNo string, carrierNo string) *RestfulBookingModel {
	this := RestfulBookingModel{}
	this.DocType = docType
	this.DocNo = docNo
	this.CarrierNo = carrierNo
	return &this
}

// NewRestfulBookingModelWithDefaults instantiates a new RestfulBookingModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestfulBookingModelWithDefaults() *RestfulBookingModel {
	this := RestfulBookingModel{}
	return &this
}

// GetDocType returns the DocType field value
func (o *RestfulBookingModel) GetDocType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocType
}

// GetDocTypeOk returns a tuple with the DocType field value
// and a boolean to check if the value has been set.
func (o *RestfulBookingModel) GetDocTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocType, true
}

// SetDocType sets field value
func (o *RestfulBookingModel) SetDocType(v string) {
	o.DocType = v
}

// GetDocNo returns the DocNo field value
func (o *RestfulBookingModel) GetDocNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocNo
}

// GetDocNoOk returns a tuple with the DocNo field value
// and a boolean to check if the value has been set.
func (o *RestfulBookingModel) GetDocNoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocNo, true
}

// SetDocNo sets field value
func (o *RestfulBookingModel) SetDocNo(v string) {
	o.DocNo = v
}

// GetCarrierNo returns the CarrierNo field value
func (o *RestfulBookingModel) GetCarrierNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierNo
}

// GetCarrierNoOk returns a tuple with the CarrierNo field value
// and a boolean to check if the value has been set.
func (o *RestfulBookingModel) GetCarrierNoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CarrierNo, true
}

// SetCarrierNo sets field value
func (o *RestfulBookingModel) SetCarrierNo(v string) {
	o.CarrierNo = v
}

func (o RestfulBookingModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["doc_type"] = o.DocType
	}
	if true {
		toSerialize["doc_no"] = o.DocNo
	}
	if true {
		toSerialize["carrier_no"] = o.CarrierNo
	}
	return json.Marshal(toSerialize)
}

type NullableRestfulBookingModel struct {
	value *RestfulBookingModel
	isSet bool
}

func (v NullableRestfulBookingModel) Get() *RestfulBookingModel {
	return v.value
}

func (v *NullableRestfulBookingModel) Set(val *RestfulBookingModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRestfulBookingModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRestfulBookingModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestfulBookingModel(val *RestfulBookingModel) *NullableRestfulBookingModel {
	return &NullableRestfulBookingModel{value: val, isSet: true}
}

func (v NullableRestfulBookingModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestfulBookingModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

