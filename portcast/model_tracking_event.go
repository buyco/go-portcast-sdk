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

// TrackingEvent struct for TrackingEvent
type TrackingEvent struct {
	BillOfLading         *TrackingEventBillOfLading          `json:"bill_of_lading,omitempty"`
	BillOfLadingBookmark *TrackingEventBillOfLadingBookmark  `json:"bill_of_lading_bookmark,omitempty"`
	ContainerEventList   *[]TrackingEventContainerEventList  `json:"container_event_list,omitempty"`
	ContainerMetaInfo    *map[string]interface{}             `json:"container_meta_info,omitempty"`
	Id                   *string                             `json:"id,omitempty"`
	Msg                  *string                             `json:"msg,omitempty"`
	OrgId                *string                             `json:"org_id,omitempty"`
	SailingInfoTracking  *[]TrackingEventSailingInfoTracking `json:"sailing_info_tracking,omitempty"`
}

// NewTrackingEvent instantiates a new TrackingEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingEvent() *TrackingEvent {
	this := TrackingEvent{}
	return &this
}

// NewTrackingEventWithDefaults instantiates a new TrackingEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingEventWithDefaults() *TrackingEvent {
	this := TrackingEvent{}
	return &this
}

// GetBillOfLading returns the BillOfLading field value if set, zero value otherwise.
func (o *TrackingEvent) GetBillOfLading() TrackingEventBillOfLading {
	if o == nil || o.BillOfLading == nil {
		var ret TrackingEventBillOfLading
		return ret
	}
	return *o.BillOfLading
}

// GetBillOfLadingOk returns a tuple with the BillOfLading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEvent) GetBillOfLadingOk() (*TrackingEventBillOfLading, bool) {
	if o == nil || o.BillOfLading == nil {
		return nil, false
	}
	return o.BillOfLading, true
}

// HasBillOfLading returns a boolean if a field has been set.
func (o *TrackingEvent) HasBillOfLading() bool {
	if o != nil && o.BillOfLading != nil {
		return true
	}

	return false
}

// SetBillOfLading gets a reference to the given TrackingEventBillOfLading and assigns it to the BillOfLading field.
func (o *TrackingEvent) SetBillOfLading(v TrackingEventBillOfLading) {
	o.BillOfLading = &v
}

// GetBillOfLadingBookmark returns the BillOfLadingBookmark field value if set, zero value otherwise.
func (o *TrackingEvent) GetBillOfLadingBookmark() TrackingEventBillOfLadingBookmark {
	if o == nil || o.BillOfLadingBookmark == nil {
		var ret TrackingEventBillOfLadingBookmark
		return ret
	}
	return *o.BillOfLadingBookmark
}

// GetBillOfLadingBookmarkOk returns a tuple with the BillOfLadingBookmark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEvent) GetBillOfLadingBookmarkOk() (*TrackingEventBillOfLadingBookmark, bool) {
	if o == nil || o.BillOfLadingBookmark == nil {
		return nil, false
	}
	return o.BillOfLadingBookmark, true
}

// HasBillOfLadingBookmark returns a boolean if a field has been set.
func (o *TrackingEvent) HasBillOfLadingBookmark() bool {
	if o != nil && o.BillOfLadingBookmark != nil {
		return true
	}

	return false
}

// SetBillOfLadingBookmark gets a reference to the given TrackingEventBillOfLadingBookmark and assigns it to the BillOfLadingBookmark field.
func (o *TrackingEvent) SetBillOfLadingBookmark(v TrackingEventBillOfLadingBookmark) {
	o.BillOfLadingBookmark = &v
}

// GetContainerEventList returns the ContainerEventList field value if set, zero value otherwise.
func (o *TrackingEvent) GetContainerEventList() []TrackingEventContainerEventList {
	if o == nil || o.ContainerEventList == nil {
		var ret []TrackingEventContainerEventList
		return ret
	}
	return *o.ContainerEventList
}

// GetContainerEventListOk returns a tuple with the ContainerEventList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEvent) GetContainerEventListOk() (*[]TrackingEventContainerEventList, bool) {
	if o == nil || o.ContainerEventList == nil {
		return nil, false
	}
	return o.ContainerEventList, true
}

// HasContainerEventList returns a boolean if a field has been set.
func (o *TrackingEvent) HasContainerEventList() bool {
	if o != nil && o.ContainerEventList != nil {
		return true
	}

	return false
}

// SetContainerEventList gets a reference to the given []TrackingEventContainerEventList and assigns it to the ContainerEventList field.
func (o *TrackingEvent) SetContainerEventList(v []TrackingEventContainerEventList) {
	o.ContainerEventList = &v
}

// GetContainerMetaInfo returns the ContainerMetaInfo field value if set, zero value otherwise.
func (o *TrackingEvent) GetContainerMetaInfo() map[string]interface{} {
	if o == nil || o.ContainerMetaInfo == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ContainerMetaInfo
}

// GetContainerMetaInfoOk returns a tuple with the ContainerMetaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEvent) GetContainerMetaInfoOk() (*map[string]interface{}, bool) {
	if o == nil || o.ContainerMetaInfo == nil {
		return nil, false
	}
	return o.ContainerMetaInfo, true
}

// HasContainerMetaInfo returns a boolean if a field has been set.
func (o *TrackingEvent) HasContainerMetaInfo() bool {
	if o != nil && o.ContainerMetaInfo != nil {
		return true
	}

	return false
}

// SetContainerMetaInfo gets a reference to the given map[string]interface{} and assigns it to the ContainerMetaInfo field.
func (o *TrackingEvent) SetContainerMetaInfo(v map[string]interface{}) {
	o.ContainerMetaInfo = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TrackingEvent) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEvent) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TrackingEvent) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TrackingEvent) SetId(v string) {
	o.Id = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *TrackingEvent) GetMsg() string {
	if o == nil || o.Msg == nil {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEvent) GetMsgOk() (*string, bool) {
	if o == nil || o.Msg == nil {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *TrackingEvent) HasMsg() bool {
	if o != nil && o.Msg != nil {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *TrackingEvent) SetMsg(v string) {
	o.Msg = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *TrackingEvent) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEvent) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *TrackingEvent) HasOrgId() bool {
	if o != nil && o.OrgId != nil {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *TrackingEvent) SetOrgId(v string) {
	o.OrgId = &v
}

// GetSailingInfoTracking returns the SailingInfoTracking field value if set, zero value otherwise.
func (o *TrackingEvent) GetSailingInfoTracking() []TrackingEventSailingInfoTracking {
	if o == nil || o.SailingInfoTracking == nil {
		var ret []TrackingEventSailingInfoTracking
		return ret
	}
	return *o.SailingInfoTracking
}

// GetSailingInfoTrackingOk returns a tuple with the SailingInfoTracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEvent) GetSailingInfoTrackingOk() (*[]TrackingEventSailingInfoTracking, bool) {
	if o == nil || o.SailingInfoTracking == nil {
		return nil, false
	}
	return o.SailingInfoTracking, true
}

// HasSailingInfoTracking returns a boolean if a field has been set.
func (o *TrackingEvent) HasSailingInfoTracking() bool {
	if o != nil && o.SailingInfoTracking != nil {
		return true
	}

	return false
}

// SetSailingInfoTracking gets a reference to the given []TrackingEventSailingInfoTracking and assigns it to the SailingInfoTracking field.
func (o *TrackingEvent) SetSailingInfoTracking(v []TrackingEventSailingInfoTracking) {
	o.SailingInfoTracking = &v
}

func (o TrackingEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BillOfLading != nil {
		toSerialize["bill_of_lading"] = o.BillOfLading
	}
	if o.BillOfLadingBookmark != nil {
		toSerialize["bill_of_lading_bookmark"] = o.BillOfLadingBookmark
	}
	if o.ContainerEventList != nil {
		toSerialize["container_event_list"] = o.ContainerEventList
	}
	if o.ContainerMetaInfo != nil {
		toSerialize["container_meta_info"] = o.ContainerMetaInfo
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Msg != nil {
		toSerialize["msg"] = o.Msg
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.SailingInfoTracking != nil {
		toSerialize["sailing_info_tracking"] = o.SailingInfoTracking
	}
	return json.Marshal(toSerialize)
}

type NullableTrackingEvent struct {
	value *TrackingEvent
	isSet bool
}

func (v NullableTrackingEvent) Get() *TrackingEvent {
	return v.value
}

func (v *NullableTrackingEvent) Set(val *TrackingEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingEvent(val *TrackingEvent) *NullableTrackingEvent {
	return &NullableTrackingEvent{value: val, isSet: true}
}

func (v NullableTrackingEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackingEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
