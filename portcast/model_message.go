/*
Portcast API (1.1.0) - Bill of Lading Tracking

**This documentation is for the latest version of the Portcast Bill of Lading Tracking API.**  There are two variables used in this documentation: 1. `api-url`: the url to use for accessing the API. The official url is `https://api.portcast.io` 2. `x-api-key`: the access token to send along with every request to the API. This key will be provided to each organisation upon API access activation  The general workflow is as below:  1. Create the bill of lading bookmark if it does not exist already (`POST {{api-url}}/api/v1/eta/bill-of-lading-bookmarks`). 2. A bookmark must contain `carrier_no`, `bl_no` and `cntr_no` information. This will return the bill of lading bookmark information created. Record the `id` field from the response. 3. Wait for predictions to be generated. This could take up to 5 mins. 5. Query for the tracking results based on the `id` recorded earlier (`GET {{api-url}}/api/v1/eta/tracking/bill-of-lading-bookmarks/<id>`) 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package portcast

import (
	"encoding/json"
)

// Message struct for Message
type Message struct {
	Message *string `json:"message,omitempty"`
}

// NewMessage instantiates a new Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessage() *Message {
	this := Message{}
	return &this
}

// NewMessageWithDefaults instantiates a new Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageWithDefaults() *Message {
	this := Message{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Message) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Message) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Message) SetMessage(v string) {
	o.Message = &v
}

func (o Message) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableMessage struct {
	value *Message
	isSet bool
}

func (v NullableMessage) Get() *Message {
	return v.value
}

func (v *NullableMessage) Set(val *Message) {
	v.value = val
	v.isSet = true
}

func (v NullableMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage(val *Message) *NullableMessage {
	return &NullableMessage{value: val, isSet: true}
}

func (v NullableMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


