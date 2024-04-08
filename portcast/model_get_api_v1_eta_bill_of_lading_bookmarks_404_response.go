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

// checks if the GetApiV1EtaBillOfLadingBookmarks404Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetApiV1EtaBillOfLadingBookmarks404Response{}

// GetApiV1EtaBillOfLadingBookmarks404Response struct for GetApiV1EtaBillOfLadingBookmarks404Response
type GetApiV1EtaBillOfLadingBookmarks404Response struct {
	Error *string `json:"error,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewGetApiV1EtaBillOfLadingBookmarks404Response instantiates a new GetApiV1EtaBillOfLadingBookmarks404Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApiV1EtaBillOfLadingBookmarks404Response() *GetApiV1EtaBillOfLadingBookmarks404Response {
	this := GetApiV1EtaBillOfLadingBookmarks404Response{}
	var error_ string = "Record not found with give Org/Bill Of Lading ID"
	this.Error = &error_
	var message string = "You have requested this URI [/api/v1/eta/tracking/bill-of-lading-bookmarks/e48cb2fb-499b-4af4-a0ef-250eac4eee5f] but did you mean /api/v1/eta/tracking/bill-of-lading-bookmarks/<bill_of_lading_bookmark_id> or /api/v1/eta/tracking/bill-of-lading-bookmarks or /api/v1/orgs/<org_id>/eta/tracking/bill-of-lading-bookmarks ?"
	this.Message = &message
	return &this
}

// NewGetApiV1EtaBillOfLadingBookmarks404ResponseWithDefaults instantiates a new GetApiV1EtaBillOfLadingBookmarks404Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApiV1EtaBillOfLadingBookmarks404ResponseWithDefaults() *GetApiV1EtaBillOfLadingBookmarks404Response {
	this := GetApiV1EtaBillOfLadingBookmarks404Response{}
	var error_ string = "Record not found with give Org/Bill Of Lading ID"
	this.Error = &error_
	var message string = "You have requested this URI [/api/v1/eta/tracking/bill-of-lading-bookmarks/e48cb2fb-499b-4af4-a0ef-250eac4eee5f] but did you mean /api/v1/eta/tracking/bill-of-lading-bookmarks/<bill_of_lading_bookmark_id> or /api/v1/eta/tracking/bill-of-lading-bookmarks or /api/v1/orgs/<org_id>/eta/tracking/bill-of-lading-bookmarks ?"
	this.Message = &message
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetApiV1EtaBillOfLadingBookmarks404Response) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiV1EtaBillOfLadingBookmarks404Response) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetApiV1EtaBillOfLadingBookmarks404Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetApiV1EtaBillOfLadingBookmarks404Response) SetError(v string) {
	o.Error = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetApiV1EtaBillOfLadingBookmarks404Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiV1EtaBillOfLadingBookmarks404Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetApiV1EtaBillOfLadingBookmarks404Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetApiV1EtaBillOfLadingBookmarks404Response) SetMessage(v string) {
	o.Message = &v
}

func (o GetApiV1EtaBillOfLadingBookmarks404Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetApiV1EtaBillOfLadingBookmarks404Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGetApiV1EtaBillOfLadingBookmarks404Response struct {
	value *GetApiV1EtaBillOfLadingBookmarks404Response
	isSet bool
}

func (v NullableGetApiV1EtaBillOfLadingBookmarks404Response) Get() *GetApiV1EtaBillOfLadingBookmarks404Response {
	return v.value
}

func (v *NullableGetApiV1EtaBillOfLadingBookmarks404Response) Set(val *GetApiV1EtaBillOfLadingBookmarks404Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApiV1EtaBillOfLadingBookmarks404Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApiV1EtaBillOfLadingBookmarks404Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApiV1EtaBillOfLadingBookmarks404Response(val *GetApiV1EtaBillOfLadingBookmarks404Response) *NullableGetApiV1EtaBillOfLadingBookmarks404Response {
	return &NullableGetApiV1EtaBillOfLadingBookmarks404Response{value: val, isSet: true}
}

func (v NullableGetApiV1EtaBillOfLadingBookmarks404Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApiV1EtaBillOfLadingBookmarks404Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


