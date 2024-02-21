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
	"fmt"
)

// GetApiV1EtaBillOfLadingBookmarksTerminals400Response struct for GetApiV1EtaBillOfLadingBookmarksTerminals400Response
type GetApiV1EtaBillOfLadingBookmarksTerminals400Response struct {
	GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf  *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf
	GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1 *GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GetApiV1EtaBillOfLadingBookmarksTerminals400Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf
	err = json.Unmarshal(data, &dst.GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf)
	if err == nil {
		jsonGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf, _ := json.Marshal(dst.GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf)
		if string(jsonGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf) == "{}" { // empty struct
			dst.GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf, return on the first match
		}
	} else {
		dst.GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf = nil
	}

	// try to unmarshal JSON data into GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1
	err = json.Unmarshal(data, &dst.GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1)
	if err == nil {
		jsonGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1, _ := json.Marshal(dst.GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1)
		if string(jsonGetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1) == "{}" { // empty struct
			dst.GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1 = nil
		} else {
			return nil // data stored in dst.GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1, return on the first match
		}
	} else {
		dst.GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(GetApiV1EtaBillOfLadingBookmarksTerminals400Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GetApiV1EtaBillOfLadingBookmarksTerminals400Response) MarshalJSON() ([]byte, error) {
	if src.GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf != nil {
		return json.Marshal(&src.GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf)
	}

	if src.GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1 != nil {
		return json.Marshal(&src.GetApiV1EtaBillOfLadingBookmarksTerminals400ResponseAnyOf1)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableGetApiV1EtaBillOfLadingBookmarksTerminals400Response struct {
	value *GetApiV1EtaBillOfLadingBookmarksTerminals400Response
	isSet bool
}

func (v NullableGetApiV1EtaBillOfLadingBookmarksTerminals400Response) Get() *GetApiV1EtaBillOfLadingBookmarksTerminals400Response {
	return v.value
}

func (v *NullableGetApiV1EtaBillOfLadingBookmarksTerminals400Response) Set(val *GetApiV1EtaBillOfLadingBookmarksTerminals400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApiV1EtaBillOfLadingBookmarksTerminals400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApiV1EtaBillOfLadingBookmarksTerminals400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApiV1EtaBillOfLadingBookmarksTerminals400Response(val *GetApiV1EtaBillOfLadingBookmarksTerminals400Response) *NullableGetApiV1EtaBillOfLadingBookmarksTerminals400Response {
	return &NullableGetApiV1EtaBillOfLadingBookmarksTerminals400Response{value: val, isSet: true}
}

func (v NullableGetApiV1EtaBillOfLadingBookmarksTerminals400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApiV1EtaBillOfLadingBookmarksTerminals400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
