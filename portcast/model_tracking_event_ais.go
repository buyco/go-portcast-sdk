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

// TrackingEventAis struct for TrackingEventAis
type TrackingEventAis struct {
	Course *float32 `json:"course,omitempty"`
	DraughtM *string `json:"draught_m,omitempty"`
	Imo *string `json:"imo,omitempty"`
	Lat *float32 `json:"lat,omitempty"`
	Lon *float32 `json:"lon,omitempty"`
	SpeedNm *int32 `json:"speed_nm,omitempty"`
	Status *string `json:"status,omitempty"`
	TimestampUtc *time.Time `json:"timestamp_utc,omitempty"`
}

// NewTrackingEventAis instantiates a new TrackingEventAis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingEventAis() *TrackingEventAis {
	this := TrackingEventAis{}
	return &this
}

// NewTrackingEventAisWithDefaults instantiates a new TrackingEventAis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingEventAisWithDefaults() *TrackingEventAis {
	this := TrackingEventAis{}
	return &this
}

// GetCourse returns the Course field value if set, zero value otherwise.
func (o *TrackingEventAis) GetCourse() float32 {
	if o == nil || o.Course == nil {
		var ret float32
		return ret
	}
	return *o.Course
}

// GetCourseOk returns a tuple with the Course field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventAis) GetCourseOk() (*float32, bool) {
	if o == nil || o.Course == nil {
		return nil, false
	}
	return o.Course, true
}

// HasCourse returns a boolean if a field has been set.
func (o *TrackingEventAis) HasCourse() bool {
	if o != nil && o.Course != nil {
		return true
	}

	return false
}

// SetCourse gets a reference to the given float32 and assigns it to the Course field.
func (o *TrackingEventAis) SetCourse(v float32) {
	o.Course = &v
}

// GetDraughtM returns the DraughtM field value if set, zero value otherwise.
func (o *TrackingEventAis) GetDraughtM() string {
	if o == nil || o.DraughtM == nil {
		var ret string
		return ret
	}
	return *o.DraughtM
}

// GetDraughtMOk returns a tuple with the DraughtM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventAis) GetDraughtMOk() (*string, bool) {
	if o == nil || o.DraughtM == nil {
		return nil, false
	}
	return o.DraughtM, true
}

// HasDraughtM returns a boolean if a field has been set.
func (o *TrackingEventAis) HasDraughtM() bool {
	if o != nil && o.DraughtM != nil {
		return true
	}

	return false
}

// SetDraughtM gets a reference to the given string and assigns it to the DraughtM field.
func (o *TrackingEventAis) SetDraughtM(v string) {
	o.DraughtM = &v
}

// GetImo returns the Imo field value if set, zero value otherwise.
func (o *TrackingEventAis) GetImo() string {
	if o == nil || o.Imo == nil {
		var ret string
		return ret
	}
	return *o.Imo
}

// GetImoOk returns a tuple with the Imo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventAis) GetImoOk() (*string, bool) {
	if o == nil || o.Imo == nil {
		return nil, false
	}
	return o.Imo, true
}

// HasImo returns a boolean if a field has been set.
func (o *TrackingEventAis) HasImo() bool {
	if o != nil && o.Imo != nil {
		return true
	}

	return false
}

// SetImo gets a reference to the given string and assigns it to the Imo field.
func (o *TrackingEventAis) SetImo(v string) {
	o.Imo = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *TrackingEventAis) GetLat() float32 {
	if o == nil || o.Lat == nil {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventAis) GetLatOk() (*float32, bool) {
	if o == nil || o.Lat == nil {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *TrackingEventAis) HasLat() bool {
	if o != nil && o.Lat != nil {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *TrackingEventAis) SetLat(v float32) {
	o.Lat = &v
}

// GetLon returns the Lon field value if set, zero value otherwise.
func (o *TrackingEventAis) GetLon() float32 {
	if o == nil || o.Lon == nil {
		var ret float32
		return ret
	}
	return *o.Lon
}

// GetLonOk returns a tuple with the Lon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventAis) GetLonOk() (*float32, bool) {
	if o == nil || o.Lon == nil {
		return nil, false
	}
	return o.Lon, true
}

// HasLon returns a boolean if a field has been set.
func (o *TrackingEventAis) HasLon() bool {
	if o != nil && o.Lon != nil {
		return true
	}

	return false
}

// SetLon gets a reference to the given float32 and assigns it to the Lon field.
func (o *TrackingEventAis) SetLon(v float32) {
	o.Lon = &v
}

// GetSpeedNm returns the SpeedNm field value if set, zero value otherwise.
func (o *TrackingEventAis) GetSpeedNm() int32 {
	if o == nil || o.SpeedNm == nil {
		var ret int32
		return ret
	}
	return *o.SpeedNm
}

// GetSpeedNmOk returns a tuple with the SpeedNm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventAis) GetSpeedNmOk() (*int32, bool) {
	if o == nil || o.SpeedNm == nil {
		return nil, false
	}
	return o.SpeedNm, true
}

// HasSpeedNm returns a boolean if a field has been set.
func (o *TrackingEventAis) HasSpeedNm() bool {
	if o != nil && o.SpeedNm != nil {
		return true
	}

	return false
}

// SetSpeedNm gets a reference to the given int32 and assigns it to the SpeedNm field.
func (o *TrackingEventAis) SetSpeedNm(v int32) {
	o.SpeedNm = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TrackingEventAis) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventAis) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TrackingEventAis) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TrackingEventAis) SetStatus(v string) {
	o.Status = &v
}

// GetTimestampUtc returns the TimestampUtc field value if set, zero value otherwise.
func (o *TrackingEventAis) GetTimestampUtc() time.Time {
	if o == nil || o.TimestampUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.TimestampUtc
}

// GetTimestampUtcOk returns a tuple with the TimestampUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventAis) GetTimestampUtcOk() (*time.Time, bool) {
	if o == nil || o.TimestampUtc == nil {
		return nil, false
	}
	return o.TimestampUtc, true
}

// HasTimestampUtc returns a boolean if a field has been set.
func (o *TrackingEventAis) HasTimestampUtc() bool {
	if o != nil && o.TimestampUtc != nil {
		return true
	}

	return false
}

// SetTimestampUtc gets a reference to the given time.Time and assigns it to the TimestampUtc field.
func (o *TrackingEventAis) SetTimestampUtc(v time.Time) {
	o.TimestampUtc = &v
}

func (o TrackingEventAis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Course != nil {
		toSerialize["course"] = o.Course
	}
	if o.DraughtM != nil {
		toSerialize["draught_m"] = o.DraughtM
	}
	if o.Imo != nil {
		toSerialize["imo"] = o.Imo
	}
	if o.Lat != nil {
		toSerialize["lat"] = o.Lat
	}
	if o.Lon != nil {
		toSerialize["lon"] = o.Lon
	}
	if o.SpeedNm != nil {
		toSerialize["speed_nm"] = o.SpeedNm
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TimestampUtc != nil {
		toSerialize["timestamp_utc"] = o.TimestampUtc
	}
	return json.Marshal(toSerialize)
}

type NullableTrackingEventAis struct {
	value *TrackingEventAis
	isSet bool
}

func (v NullableTrackingEventAis) Get() *TrackingEventAis {
	return v.value
}

func (v *NullableTrackingEventAis) Set(val *TrackingEventAis) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingEventAis) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingEventAis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingEventAis(val *TrackingEventAis) *NullableTrackingEventAis {
	return &NullableTrackingEventAis{value: val, isSet: true}
}

func (v NullableTrackingEventAis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackingEventAis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

