/*
Portcast API (1.1.0) - Bill of Lading Tracking

**This documentation is for the latest version of the Portcast Bill of Lading Tracking API.**  There are two variables used in this documentation: 1. `api-url`: the url to use for accessing the API. The official url is `https://api.portcast.io` 2. `x-api-key`: the access token to send along with every request to the API. This key will be provided to each organisation upon API access activation  The general workflow is as below:  1. Create the bill of lading bookmark if it does not exist already (`POST {{api-url}}/api/v1/eta/bill-of-lading-bookmarks`). 2. A bookmark must contain `carrier_no`, `bl_no` and `cntr_no` information. This will return the bill of lading bookmark information created. Record the `id` field from the response. 3. Wait for predictions to be generated. This could take up to 5 mins. 5. Query for the tracking results based on the `id` recorded earlier (`GET {{api-url}}/api/v1/eta/tracking/bill-of-lading-bookmarks/<id>`) 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package portcast

import (
	"encoding/json"
	"time"
)

// Customer struct for Customer
type Customer struct {
	Created time.Time `json:"created"`
	Name string `json:"name"`
}

// NewCustomer instantiates a new Customer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomer(created time.Time, name string) *Customer {
	this := Customer{}
	this.Created = created
	this.Name = name
	return &this
}

// NewCustomerWithDefaults instantiates a new Customer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerWithDefaults() *Customer {
	this := Customer{}
	return &this
}

// GetCreated returns the Created field value
func (o *Customer) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Customer) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Customer) SetCreated(v time.Time) {
	o.Created = v
}

// GetName returns the Name field value
func (o *Customer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Customer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Customer) SetName(v string) {
	o.Name = v
}

func (o Customer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCustomer struct {
	value *Customer
	isSet bool
}

func (v NullableCustomer) Get() *Customer {
	return v.value
}

func (v *NullableCustomer) Set(val *Customer) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomer(val *Customer) *NullableCustomer {
	return &NullableCustomer{value: val, isSet: true}
}

func (v NullableCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


