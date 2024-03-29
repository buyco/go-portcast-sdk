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

// TrackingEventSailingInfoTracking struct for TrackingEventSailingInfoTracking
type TrackingEventSailingInfoTracking struct {
	ActualArrivalLt     *string                       `json:"actual_arrival_lt,omitempty"`
	ActualArrivalUtc    *string                       `json:"actual_arrival_utc,omitempty"`
	Ais                 *TrackingEventAis             `json:"ais,omitempty"`
	PredictedArrivalLt  *string                       `json:"predicted_arrival_lt,omitempty"`
	PredictedArrivalUtc *string                       `json:"predicted_arrival_utc,omitempty"`
	PredictionTimeUtc   *string                       `json:"prediction_time_utc,omitempty"`
	SailingInfo         *TrackingEventSailingInfo     `json:"sailing_info,omitempty"`
	ScheduledArrivalLt  *string                       `json:"scheduled_arrival_lt,omitempty"`
	ScheduledArrivalUtc *string                       `json:"scheduled_arrival_utc,omitempty"`
	TargetPortCode      *string                       `json:"target_port_code,omitempty"`
	TargetPortName      *string                       `json:"target_port_name,omitempty"`
	VoyageDetails       *[]TrackingEventVoyageDetails `json:"voyage_details,omitempty"`
	VoyageNoList        *[]string                     `json:"voyage_no_list,omitempty"`
}

// NewTrackingEventSailingInfoTracking instantiates a new TrackingEventSailingInfoTracking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingEventSailingInfoTracking() *TrackingEventSailingInfoTracking {
	this := TrackingEventSailingInfoTracking{}
	return &this
}

// NewTrackingEventSailingInfoTrackingWithDefaults instantiates a new TrackingEventSailingInfoTracking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingEventSailingInfoTrackingWithDefaults() *TrackingEventSailingInfoTracking {
	this := TrackingEventSailingInfoTracking{}
	return &this
}

// GetActualArrivalLt returns the ActualArrivalLt field value if set, zero value otherwise.
func (o *TrackingEventSailingInfoTracking) GetActualArrivalLt() string {
	if o == nil || o.ActualArrivalLt == nil {
		var ret string
		return ret
	}
	return *o.ActualArrivalLt
}

// GetActualArrivalLtOk returns a tuple with the ActualArrivalLt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventSailingInfoTracking) GetActualArrivalLtOk() (*string, bool) {
	if o == nil || o.ActualArrivalLt == nil {
		return nil, false
	}
	return o.ActualArrivalLt, true
}

// HasActualArrivalLt returns a boolean if a field has been set.
func (o *TrackingEventSailingInfoTracking) HasActualArrivalLt() bool {
	if o != nil && o.ActualArrivalLt != nil {
		return true
	}

	return false
}

// SetActualArrivalLt gets a reference to the given string and assigns it to the ActualArrivalLt field.
func (o *TrackingEventSailingInfoTracking) SetActualArrivalLt(v string) {
	o.ActualArrivalLt = &v
}

// GetActualArrivalUtc returns the ActualArrivalUtc field value if set, zero value otherwise.
func (o *TrackingEventSailingInfoTracking) GetActualArrivalUtc() string {
	if o == nil || o.ActualArrivalUtc == nil {
		var ret string
		return ret
	}
	return *o.ActualArrivalUtc
}

// GetActualArrivalUtcOk returns a tuple with the ActualArrivalUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventSailingInfoTracking) GetActualArrivalUtcOk() (*string, bool) {
	if o == nil || o.ActualArrivalUtc == nil {
		return nil, false
	}
	return o.ActualArrivalUtc, true
}

// HasActualArrivalUtc returns a boolean if a field has been set.
func (o *TrackingEventSailingInfoTracking) HasActualArrivalUtc() bool {
	if o != nil && o.ActualArrivalUtc != nil {
		return true
	}

	return false
}

// SetActualArrivalUtc gets a reference to the given string and assigns it to the ActualArrivalUtc field.
func (o *TrackingEventSailingInfoTracking) SetActualArrivalUtc(v string) {
	o.ActualArrivalUtc = &v
}

// GetAis returns the Ais field value if set, zero value otherwise.
func (o *TrackingEventSailingInfoTracking) GetAis() TrackingEventAis {
	if o == nil || o.Ais == nil {
		var ret TrackingEventAis
		return ret
	}
	return *o.Ais
}

// GetAisOk returns a tuple with the Ais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventSailingInfoTracking) GetAisOk() (*TrackingEventAis, bool) {
	if o == nil || o.Ais == nil {
		return nil, false
	}
	return o.Ais, true
}

// HasAis returns a boolean if a field has been set.
func (o *TrackingEventSailingInfoTracking) HasAis() bool {
	if o != nil && o.Ais != nil {
		return true
	}

	return false
}

// SetAis gets a reference to the given TrackingEventAis and assigns it to the Ais field.
func (o *TrackingEventSailingInfoTracking) SetAis(v TrackingEventAis) {
	o.Ais = &v
}

// GetPredictedArrivalLt returns the PredictedArrivalLt field value if set, zero value otherwise.
func (o *TrackingEventSailingInfoTracking) GetPredictedArrivalLt() string {
	if o == nil || o.PredictedArrivalLt == nil {
		var ret string
		return ret
	}
	return *o.PredictedArrivalLt
}

// GetPredictedArrivalLtOk returns a tuple with the PredictedArrivalLt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventSailingInfoTracking) GetPredictedArrivalLtOk() (*string, bool) {
	if o == nil || o.PredictedArrivalLt == nil {
		return nil, false
	}
	return o.PredictedArrivalLt, true
}

// HasPredictedArrivalLt returns a boolean if a field has been set.
func (o *TrackingEventSailingInfoTracking) HasPredictedArrivalLt() bool {
	if o != nil && o.PredictedArrivalLt != nil {
		return true
	}

	return false
}

// SetPredictedArrivalLt gets a reference to the given string and assigns it to the PredictedArrivalLt field.
func (o *TrackingEventSailingInfoTracking) SetPredictedArrivalLt(v string) {
	o.PredictedArrivalLt = &v
}

// GetPredictedArrivalUtc returns the PredictedArrivalUtc field value if set, zero value otherwise.
func (o *TrackingEventSailingInfoTracking) GetPredictedArrivalUtc() string {
	if o == nil || o.PredictedArrivalUtc == nil {
		var ret string
		return ret
	}
	return *o.PredictedArrivalUtc
}

// GetPredictedArrivalUtcOk returns a tuple with the PredictedArrivalUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventSailingInfoTracking) GetPredictedArrivalUtcOk() (*string, bool) {
	if o == nil || o.PredictedArrivalUtc == nil {
		return nil, false
	}
	return o.PredictedArrivalUtc, true
}

// HasPredictedArrivalUtc returns a boolean if a field has been set.
func (o *TrackingEventSailingInfoTracking) HasPredictedArrivalUtc() bool {
	if o != nil && o.PredictedArrivalUtc != nil {
		return true
	}

	return false
}

// SetPredictedArrivalUtc gets a reference to the given string and assigns it to the PredictedArrivalUtc field.
func (o *TrackingEventSailingInfoTracking) SetPredictedArrivalUtc(v string) {
	o.PredictedArrivalUtc = &v
}

// GetPredictionTimeUtc returns the PredictionTimeUtc field value if set, zero value otherwise.
func (o *TrackingEventSailingInfoTracking) GetPredictionTimeUtc() string {
	if o == nil || o.PredictionTimeUtc == nil {
		var ret string
		return ret
	}
	return *o.PredictionTimeUtc
}

// GetPredictionTimeUtcOk returns a tuple with the PredictionTimeUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventSailingInfoTracking) GetPredictionTimeUtcOk() (*string, bool) {
	if o == nil || o.PredictionTimeUtc == nil {
		return nil, false
	}
	return o.PredictionTimeUtc, true
}

// HasPredictionTimeUtc returns a boolean if a field has been set.
func (o *TrackingEventSailingInfoTracking) HasPredictionTimeUtc() bool {
	if o != nil && o.PredictionTimeUtc != nil {
		return true
	}

	return false
}

// SetPredictionTimeUtc gets a reference to the given string and assigns it to the PredictionTimeUtc field.
func (o *TrackingEventSailingInfoTracking) SetPredictionTimeUtc(v string) {
	o.PredictionTimeUtc = &v
}

// GetSailingInfo returns the SailingInfo field value if set, zero value otherwise.
func (o *TrackingEventSailingInfoTracking) GetSailingInfo() TrackingEventSailingInfo {
	if o == nil || o.SailingInfo == nil {
		var ret TrackingEventSailingInfo
		return ret
	}
	return *o.SailingInfo
}

// GetSailingInfoOk returns a tuple with the SailingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventSailingInfoTracking) GetSailingInfoOk() (*TrackingEventSailingInfo, bool) {
	if o == nil || o.SailingInfo == nil {
		return nil, false
	}
	return o.SailingInfo, true
}

// HasSailingInfo returns a boolean if a field has been set.
func (o *TrackingEventSailingInfoTracking) HasSailingInfo() bool {
	if o != nil && o.SailingInfo != nil {
		return true
	}

	return false
}

// SetSailingInfo gets a reference to the given TrackingEventSailingInfo and assigns it to the SailingInfo field.
func (o *TrackingEventSailingInfoTracking) SetSailingInfo(v TrackingEventSailingInfo) {
	o.SailingInfo = &v
}

// GetScheduledArrivalLt returns the ScheduledArrivalLt field value if set, zero value otherwise.
func (o *TrackingEventSailingInfoTracking) GetScheduledArrivalLt() string {
	if o == nil || o.ScheduledArrivalLt == nil {
		var ret string
		return ret
	}
	return *o.ScheduledArrivalLt
}

// GetScheduledArrivalLtOk returns a tuple with the ScheduledArrivalLt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventSailingInfoTracking) GetScheduledArrivalLtOk() (*string, bool) {
	if o == nil || o.ScheduledArrivalLt == nil {
		return nil, false
	}
	return o.ScheduledArrivalLt, true
}

// HasScheduledArrivalLt returns a boolean if a field has been set.
func (o *TrackingEventSailingInfoTracking) HasScheduledArrivalLt() bool {
	if o != nil && o.ScheduledArrivalLt != nil {
		return true
	}

	return false
}

// SetScheduledArrivalLt gets a reference to the given string and assigns it to the ScheduledArrivalLt field.
func (o *TrackingEventSailingInfoTracking) SetScheduledArrivalLt(v string) {
	o.ScheduledArrivalLt = &v
}

// GetScheduledArrivalUtc returns the ScheduledArrivalUtc field value if set, zero value otherwise.
func (o *TrackingEventSailingInfoTracking) GetScheduledArrivalUtc() string {
	if o == nil || o.ScheduledArrivalUtc == nil {
		var ret string
		return ret
	}
	return *o.ScheduledArrivalUtc
}

// GetScheduledArrivalUtcOk returns a tuple with the ScheduledArrivalUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventSailingInfoTracking) GetScheduledArrivalUtcOk() (*string, bool) {
	if o == nil || o.ScheduledArrivalUtc == nil {
		return nil, false
	}
	return o.ScheduledArrivalUtc, true
}

// HasScheduledArrivalUtc returns a boolean if a field has been set.
func (o *TrackingEventSailingInfoTracking) HasScheduledArrivalUtc() bool {
	if o != nil && o.ScheduledArrivalUtc != nil {
		return true
	}

	return false
}

// SetScheduledArrivalUtc gets a reference to the given string and assigns it to the ScheduledArrivalUtc field.
func (o *TrackingEventSailingInfoTracking) SetScheduledArrivalUtc(v string) {
	o.ScheduledArrivalUtc = &v
}

// GetTargetPortCode returns the TargetPortCode field value if set, zero value otherwise.
func (o *TrackingEventSailingInfoTracking) GetTargetPortCode() string {
	if o == nil || o.TargetPortCode == nil {
		var ret string
		return ret
	}
	return *o.TargetPortCode
}

// GetTargetPortCodeOk returns a tuple with the TargetPortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventSailingInfoTracking) GetTargetPortCodeOk() (*string, bool) {
	if o == nil || o.TargetPortCode == nil {
		return nil, false
	}
	return o.TargetPortCode, true
}

// HasTargetPortCode returns a boolean if a field has been set.
func (o *TrackingEventSailingInfoTracking) HasTargetPortCode() bool {
	if o != nil && o.TargetPortCode != nil {
		return true
	}

	return false
}

// SetTargetPortCode gets a reference to the given string and assigns it to the TargetPortCode field.
func (o *TrackingEventSailingInfoTracking) SetTargetPortCode(v string) {
	o.TargetPortCode = &v
}

// GetTargetPortName returns the TargetPortName field value if set, zero value otherwise.
func (o *TrackingEventSailingInfoTracking) GetTargetPortName() string {
	if o == nil || o.TargetPortName == nil {
		var ret string
		return ret
	}
	return *o.TargetPortName
}

// GetTargetPortNameOk returns a tuple with the TargetPortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventSailingInfoTracking) GetTargetPortNameOk() (*string, bool) {
	if o == nil || o.TargetPortName == nil {
		return nil, false
	}
	return o.TargetPortName, true
}

// HasTargetPortName returns a boolean if a field has been set.
func (o *TrackingEventSailingInfoTracking) HasTargetPortName() bool {
	if o != nil && o.TargetPortName != nil {
		return true
	}

	return false
}

// SetTargetPortName gets a reference to the given string and assigns it to the TargetPortName field.
func (o *TrackingEventSailingInfoTracking) SetTargetPortName(v string) {
	o.TargetPortName = &v
}

// GetVoyageDetails returns the VoyageDetails field value if set, zero value otherwise.
func (o *TrackingEventSailingInfoTracking) GetVoyageDetails() []TrackingEventVoyageDetails {
	if o == nil || o.VoyageDetails == nil {
		var ret []TrackingEventVoyageDetails
		return ret
	}
	return *o.VoyageDetails
}

// GetVoyageDetailsOk returns a tuple with the VoyageDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventSailingInfoTracking) GetVoyageDetailsOk() (*[]TrackingEventVoyageDetails, bool) {
	if o == nil || o.VoyageDetails == nil {
		return nil, false
	}
	return o.VoyageDetails, true
}

// HasVoyageDetails returns a boolean if a field has been set.
func (o *TrackingEventSailingInfoTracking) HasVoyageDetails() bool {
	if o != nil && o.VoyageDetails != nil {
		return true
	}

	return false
}

// SetVoyageDetails gets a reference to the given []TrackingEventVoyageDetails and assigns it to the VoyageDetails field.
func (o *TrackingEventSailingInfoTracking) SetVoyageDetails(v []TrackingEventVoyageDetails) {
	o.VoyageDetails = &v
}

// GetVoyageNoList returns the VoyageNoList field value if set, zero value otherwise.
func (o *TrackingEventSailingInfoTracking) GetVoyageNoList() []string {
	if o == nil || o.VoyageNoList == nil {
		var ret []string
		return ret
	}
	return *o.VoyageNoList
}

// GetVoyageNoListOk returns a tuple with the VoyageNoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventSailingInfoTracking) GetVoyageNoListOk() (*[]string, bool) {
	if o == nil || o.VoyageNoList == nil {
		return nil, false
	}
	return o.VoyageNoList, true
}

// HasVoyageNoList returns a boolean if a field has been set.
func (o *TrackingEventSailingInfoTracking) HasVoyageNoList() bool {
	if o != nil && o.VoyageNoList != nil {
		return true
	}

	return false
}

// SetVoyageNoList gets a reference to the given []string and assigns it to the VoyageNoList field.
func (o *TrackingEventSailingInfoTracking) SetVoyageNoList(v []string) {
	o.VoyageNoList = &v
}

func (o TrackingEventSailingInfoTracking) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActualArrivalLt != nil {
		toSerialize["actual_arrival_lt"] = o.ActualArrivalLt
	}
	if o.ActualArrivalUtc != nil {
		toSerialize["actual_arrival_utc"] = o.ActualArrivalUtc
	}
	if o.Ais != nil {
		toSerialize["ais"] = o.Ais
	}
	if o.PredictedArrivalLt != nil {
		toSerialize["predicted_arrival_lt"] = o.PredictedArrivalLt
	}
	if o.PredictedArrivalUtc != nil {
		toSerialize["predicted_arrival_utc"] = o.PredictedArrivalUtc
	}
	if o.PredictionTimeUtc != nil {
		toSerialize["prediction_time_utc"] = o.PredictionTimeUtc
	}
	if o.SailingInfo != nil {
		toSerialize["sailing_info"] = o.SailingInfo
	}
	if o.ScheduledArrivalLt != nil {
		toSerialize["scheduled_arrival_lt"] = o.ScheduledArrivalLt
	}
	if o.ScheduledArrivalUtc != nil {
		toSerialize["scheduled_arrival_utc"] = o.ScheduledArrivalUtc
	}
	if o.TargetPortCode != nil {
		toSerialize["target_port_code"] = o.TargetPortCode
	}
	if o.TargetPortName != nil {
		toSerialize["target_port_name"] = o.TargetPortName
	}
	if o.VoyageDetails != nil {
		toSerialize["voyage_details"] = o.VoyageDetails
	}
	if o.VoyageNoList != nil {
		toSerialize["voyage_no_list"] = o.VoyageNoList
	}
	return json.Marshal(toSerialize)
}

type NullableTrackingEventSailingInfoTracking struct {
	value *TrackingEventSailingInfoTracking
	isSet bool
}

func (v NullableTrackingEventSailingInfoTracking) Get() *TrackingEventSailingInfoTracking {
	return v.value
}

func (v *NullableTrackingEventSailingInfoTracking) Set(val *TrackingEventSailingInfoTracking) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingEventSailingInfoTracking) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingEventSailingInfoTracking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingEventSailingInfoTracking(val *TrackingEventSailingInfoTracking) *NullableTrackingEventSailingInfoTracking {
	return &NullableTrackingEventSailingInfoTracking{value: val, isSet: true}
}

func (v NullableTrackingEventSailingInfoTracking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackingEventSailingInfoTracking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
