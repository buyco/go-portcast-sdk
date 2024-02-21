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

// SailingInfoTrackingStatusInfoVessel - Status Defination related to the [Vessel Schedules](docs/Portcast-Status-Codes.md)
type SailingInfoTrackingStatusInfoVessel struct {
	SailingInfoTrackingStatusInfoVesselOneOf  *SailingInfoTrackingStatusInfoVesselOneOf
	SailingInfoTrackingStatusInfoVesselOneOf1 *SailingInfoTrackingStatusInfoVesselOneOf1
	SailingInfoTrackingStatusInfoVesselOneOf2 *SailingInfoTrackingStatusInfoVesselOneOf2
}

// SailingInfoTrackingStatusInfoVesselOneOfAsSailingInfoTrackingStatusInfoVessel is a convenience function that returns SailingInfoTrackingStatusInfoVesselOneOf wrapped in SailingInfoTrackingStatusInfoVessel
func SailingInfoTrackingStatusInfoVesselOneOfAsSailingInfoTrackingStatusInfoVessel(v *SailingInfoTrackingStatusInfoVesselOneOf) SailingInfoTrackingStatusInfoVessel {
	return SailingInfoTrackingStatusInfoVessel{
		SailingInfoTrackingStatusInfoVesselOneOf: v,
	}
}

// SailingInfoTrackingStatusInfoVesselOneOf1AsSailingInfoTrackingStatusInfoVessel is a convenience function that returns SailingInfoTrackingStatusInfoVesselOneOf1 wrapped in SailingInfoTrackingStatusInfoVessel
func SailingInfoTrackingStatusInfoVesselOneOf1AsSailingInfoTrackingStatusInfoVessel(v *SailingInfoTrackingStatusInfoVesselOneOf1) SailingInfoTrackingStatusInfoVessel {
	return SailingInfoTrackingStatusInfoVessel{
		SailingInfoTrackingStatusInfoVesselOneOf1: v,
	}
}

// SailingInfoTrackingStatusInfoVesselOneOf2AsSailingInfoTrackingStatusInfoVessel is a convenience function that returns SailingInfoTrackingStatusInfoVesselOneOf2 wrapped in SailingInfoTrackingStatusInfoVessel
func SailingInfoTrackingStatusInfoVesselOneOf2AsSailingInfoTrackingStatusInfoVessel(v *SailingInfoTrackingStatusInfoVesselOneOf2) SailingInfoTrackingStatusInfoVessel {
	return SailingInfoTrackingStatusInfoVessel{
		SailingInfoTrackingStatusInfoVesselOneOf2: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SailingInfoTrackingStatusInfoVessel) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SailingInfoTrackingStatusInfoVesselOneOf
	err = newStrictDecoder(data).Decode(&dst.SailingInfoTrackingStatusInfoVesselOneOf)
	if err == nil {
		jsonSailingInfoTrackingStatusInfoVesselOneOf, _ := json.Marshal(dst.SailingInfoTrackingStatusInfoVesselOneOf)
		if string(jsonSailingInfoTrackingStatusInfoVesselOneOf) == "{}" { // empty struct
			dst.SailingInfoTrackingStatusInfoVesselOneOf = nil
		} else {
			match++
		}
	} else {
		dst.SailingInfoTrackingStatusInfoVesselOneOf = nil
	}

	// try to unmarshal data into SailingInfoTrackingStatusInfoVesselOneOf1
	err = newStrictDecoder(data).Decode(&dst.SailingInfoTrackingStatusInfoVesselOneOf1)
	if err == nil {
		jsonSailingInfoTrackingStatusInfoVesselOneOf1, _ := json.Marshal(dst.SailingInfoTrackingStatusInfoVesselOneOf1)
		if string(jsonSailingInfoTrackingStatusInfoVesselOneOf1) == "{}" { // empty struct
			dst.SailingInfoTrackingStatusInfoVesselOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.SailingInfoTrackingStatusInfoVesselOneOf1 = nil
	}

	// try to unmarshal data into SailingInfoTrackingStatusInfoVesselOneOf2
	err = newStrictDecoder(data).Decode(&dst.SailingInfoTrackingStatusInfoVesselOneOf2)
	if err == nil {
		jsonSailingInfoTrackingStatusInfoVesselOneOf2, _ := json.Marshal(dst.SailingInfoTrackingStatusInfoVesselOneOf2)
		if string(jsonSailingInfoTrackingStatusInfoVesselOneOf2) == "{}" { // empty struct
			dst.SailingInfoTrackingStatusInfoVesselOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.SailingInfoTrackingStatusInfoVesselOneOf2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SailingInfoTrackingStatusInfoVesselOneOf = nil
		dst.SailingInfoTrackingStatusInfoVesselOneOf1 = nil
		dst.SailingInfoTrackingStatusInfoVesselOneOf2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SailingInfoTrackingStatusInfoVessel)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SailingInfoTrackingStatusInfoVessel)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SailingInfoTrackingStatusInfoVessel) MarshalJSON() ([]byte, error) {
	if src.SailingInfoTrackingStatusInfoVesselOneOf != nil {
		return json.Marshal(&src.SailingInfoTrackingStatusInfoVesselOneOf)
	}

	if src.SailingInfoTrackingStatusInfoVesselOneOf1 != nil {
		return json.Marshal(&src.SailingInfoTrackingStatusInfoVesselOneOf1)
	}

	if src.SailingInfoTrackingStatusInfoVesselOneOf2 != nil {
		return json.Marshal(&src.SailingInfoTrackingStatusInfoVesselOneOf2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SailingInfoTrackingStatusInfoVessel) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SailingInfoTrackingStatusInfoVesselOneOf != nil {
		return obj.SailingInfoTrackingStatusInfoVesselOneOf
	}

	if obj.SailingInfoTrackingStatusInfoVesselOneOf1 != nil {
		return obj.SailingInfoTrackingStatusInfoVesselOneOf1
	}

	if obj.SailingInfoTrackingStatusInfoVesselOneOf2 != nil {
		return obj.SailingInfoTrackingStatusInfoVesselOneOf2
	}

	// all schemas are nil
	return nil
}

type NullableSailingInfoTrackingStatusInfoVessel struct {
	value *SailingInfoTrackingStatusInfoVessel
	isSet bool
}

func (v NullableSailingInfoTrackingStatusInfoVessel) Get() *SailingInfoTrackingStatusInfoVessel {
	return v.value
}

func (v *NullableSailingInfoTrackingStatusInfoVessel) Set(val *SailingInfoTrackingStatusInfoVessel) {
	v.value = val
	v.isSet = true
}

func (v NullableSailingInfoTrackingStatusInfoVessel) IsSet() bool {
	return v.isSet
}

func (v *NullableSailingInfoTrackingStatusInfoVessel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSailingInfoTrackingStatusInfoVessel(val *SailingInfoTrackingStatusInfoVessel) *NullableSailingInfoTrackingStatusInfoVessel {
	return &NullableSailingInfoTrackingStatusInfoVessel{value: val, isSet: true}
}

func (v NullableSailingInfoTrackingStatusInfoVessel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSailingInfoTrackingStatusInfoVessel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}