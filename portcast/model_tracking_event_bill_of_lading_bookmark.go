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

// TrackingEventBillOfLadingBookmark struct for TrackingEventBillOfLadingBookmark
type TrackingEventBillOfLadingBookmark struct {
	BlNo          *string    `json:"bl_no,omitempty"`
	CarrierNo     *string    `json:"carrier_no,omitempty"`
	CntrNo        *string    `json:"cntr_no,omitempty"`
	Created       *time.Time `json:"created,omitempty"`
	Deleted       *bool      `json:"deleted,omitempty"`
	Id            *string    `json:"id,omitempty"`
	OrgId         *string    `json:"org_id,omitempty"`
	Status        *string    `json:"status,omitempty"`
	StatusCode    *string    `json:"status_code,omitempty"`
	SystemDeleted *bool      `json:"system_deleted,omitempty"`
	Updated       *time.Time `json:"updated,omitempty"`
	VoyageNo      *string    `json:"voyage_no,omitempty"`
}

// NewTrackingEventBillOfLadingBookmark instantiates a new TrackingEventBillOfLadingBookmark object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingEventBillOfLadingBookmark() *TrackingEventBillOfLadingBookmark {
	this := TrackingEventBillOfLadingBookmark{}
	return &this
}

// NewTrackingEventBillOfLadingBookmarkWithDefaults instantiates a new TrackingEventBillOfLadingBookmark object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingEventBillOfLadingBookmarkWithDefaults() *TrackingEventBillOfLadingBookmark {
	this := TrackingEventBillOfLadingBookmark{}
	return &this
}

// GetBlNo returns the BlNo field value if set, zero value otherwise.
func (o *TrackingEventBillOfLadingBookmark) GetBlNo() string {
	if o == nil || o.BlNo == nil {
		var ret string
		return ret
	}
	return *o.BlNo
}

// GetBlNoOk returns a tuple with the BlNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventBillOfLadingBookmark) GetBlNoOk() (*string, bool) {
	if o == nil || o.BlNo == nil {
		return nil, false
	}
	return o.BlNo, true
}

// HasBlNo returns a boolean if a field has been set.
func (o *TrackingEventBillOfLadingBookmark) HasBlNo() bool {
	if o != nil && o.BlNo != nil {
		return true
	}

	return false
}

// SetBlNo gets a reference to the given string and assigns it to the BlNo field.
func (o *TrackingEventBillOfLadingBookmark) SetBlNo(v string) {
	o.BlNo = &v
}

// GetCarrierNo returns the CarrierNo field value if set, zero value otherwise.
func (o *TrackingEventBillOfLadingBookmark) GetCarrierNo() string {
	if o == nil || o.CarrierNo == nil {
		var ret string
		return ret
	}
	return *o.CarrierNo
}

// GetCarrierNoOk returns a tuple with the CarrierNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventBillOfLadingBookmark) GetCarrierNoOk() (*string, bool) {
	if o == nil || o.CarrierNo == nil {
		return nil, false
	}
	return o.CarrierNo, true
}

// HasCarrierNo returns a boolean if a field has been set.
func (o *TrackingEventBillOfLadingBookmark) HasCarrierNo() bool {
	if o != nil && o.CarrierNo != nil {
		return true
	}

	return false
}

// SetCarrierNo gets a reference to the given string and assigns it to the CarrierNo field.
func (o *TrackingEventBillOfLadingBookmark) SetCarrierNo(v string) {
	o.CarrierNo = &v
}

// GetCntrNo returns the CntrNo field value if set, zero value otherwise.
func (o *TrackingEventBillOfLadingBookmark) GetCntrNo() string {
	if o == nil || o.CntrNo == nil {
		var ret string
		return ret
	}
	return *o.CntrNo
}

// GetCntrNoOk returns a tuple with the CntrNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventBillOfLadingBookmark) GetCntrNoOk() (*string, bool) {
	if o == nil || o.CntrNo == nil {
		return nil, false
	}
	return o.CntrNo, true
}

// HasCntrNo returns a boolean if a field has been set.
func (o *TrackingEventBillOfLadingBookmark) HasCntrNo() bool {
	if o != nil && o.CntrNo != nil {
		return true
	}

	return false
}

// SetCntrNo gets a reference to the given string and assigns it to the CntrNo field.
func (o *TrackingEventBillOfLadingBookmark) SetCntrNo(v string) {
	o.CntrNo = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *TrackingEventBillOfLadingBookmark) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventBillOfLadingBookmark) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *TrackingEventBillOfLadingBookmark) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *TrackingEventBillOfLadingBookmark) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *TrackingEventBillOfLadingBookmark) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventBillOfLadingBookmark) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *TrackingEventBillOfLadingBookmark) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *TrackingEventBillOfLadingBookmark) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TrackingEventBillOfLadingBookmark) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventBillOfLadingBookmark) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TrackingEventBillOfLadingBookmark) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TrackingEventBillOfLadingBookmark) SetId(v string) {
	o.Id = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *TrackingEventBillOfLadingBookmark) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventBillOfLadingBookmark) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *TrackingEventBillOfLadingBookmark) HasOrgId() bool {
	if o != nil && o.OrgId != nil {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *TrackingEventBillOfLadingBookmark) SetOrgId(v string) {
	o.OrgId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TrackingEventBillOfLadingBookmark) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventBillOfLadingBookmark) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TrackingEventBillOfLadingBookmark) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TrackingEventBillOfLadingBookmark) SetStatus(v string) {
	o.Status = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *TrackingEventBillOfLadingBookmark) GetStatusCode() string {
	if o == nil || o.StatusCode == nil {
		var ret string
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventBillOfLadingBookmark) GetStatusCodeOk() (*string, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *TrackingEventBillOfLadingBookmark) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given string and assigns it to the StatusCode field.
func (o *TrackingEventBillOfLadingBookmark) SetStatusCode(v string) {
	o.StatusCode = &v
}

// GetSystemDeleted returns the SystemDeleted field value if set, zero value otherwise.
func (o *TrackingEventBillOfLadingBookmark) GetSystemDeleted() bool {
	if o == nil || o.SystemDeleted == nil {
		var ret bool
		return ret
	}
	return *o.SystemDeleted
}

// GetSystemDeletedOk returns a tuple with the SystemDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventBillOfLadingBookmark) GetSystemDeletedOk() (*bool, bool) {
	if o == nil || o.SystemDeleted == nil {
		return nil, false
	}
	return o.SystemDeleted, true
}

// HasSystemDeleted returns a boolean if a field has been set.
func (o *TrackingEventBillOfLadingBookmark) HasSystemDeleted() bool {
	if o != nil && o.SystemDeleted != nil {
		return true
	}

	return false
}

// SetSystemDeleted gets a reference to the given bool and assigns it to the SystemDeleted field.
func (o *TrackingEventBillOfLadingBookmark) SetSystemDeleted(v bool) {
	o.SystemDeleted = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *TrackingEventBillOfLadingBookmark) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventBillOfLadingBookmark) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *TrackingEventBillOfLadingBookmark) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *TrackingEventBillOfLadingBookmark) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetVoyageNo returns the VoyageNo field value if set, zero value otherwise.
func (o *TrackingEventBillOfLadingBookmark) GetVoyageNo() string {
	if o == nil || o.VoyageNo == nil {
		var ret string
		return ret
	}
	return *o.VoyageNo
}

// GetVoyageNoOk returns a tuple with the VoyageNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventBillOfLadingBookmark) GetVoyageNoOk() (*string, bool) {
	if o == nil || o.VoyageNo == nil {
		return nil, false
	}
	return o.VoyageNo, true
}

// HasVoyageNo returns a boolean if a field has been set.
func (o *TrackingEventBillOfLadingBookmark) HasVoyageNo() bool {
	if o != nil && o.VoyageNo != nil {
		return true
	}

	return false
}

// SetVoyageNo gets a reference to the given string and assigns it to the VoyageNo field.
func (o *TrackingEventBillOfLadingBookmark) SetVoyageNo(v string) {
	o.VoyageNo = &v
}

func (o TrackingEventBillOfLadingBookmark) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BlNo != nil {
		toSerialize["bl_no"] = o.BlNo
	}
	if o.CarrierNo != nil {
		toSerialize["carrier_no"] = o.CarrierNo
	}
	if o.CntrNo != nil {
		toSerialize["cntr_no"] = o.CntrNo
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusCode != nil {
		toSerialize["status_code"] = o.StatusCode
	}
	if o.SystemDeleted != nil {
		toSerialize["system_deleted"] = o.SystemDeleted
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	if o.VoyageNo != nil {
		toSerialize["voyage_no"] = o.VoyageNo
	}
	return json.Marshal(toSerialize)
}

type NullableTrackingEventBillOfLadingBookmark struct {
	value *TrackingEventBillOfLadingBookmark
	isSet bool
}

func (v NullableTrackingEventBillOfLadingBookmark) Get() *TrackingEventBillOfLadingBookmark {
	return v.value
}

func (v *NullableTrackingEventBillOfLadingBookmark) Set(val *TrackingEventBillOfLadingBookmark) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingEventBillOfLadingBookmark) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingEventBillOfLadingBookmark) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingEventBillOfLadingBookmark(val *TrackingEventBillOfLadingBookmark) *NullableTrackingEventBillOfLadingBookmark {
	return &NullableTrackingEventBillOfLadingBookmark{value: val, isSet: true}
}

func (v NullableTrackingEventBillOfLadingBookmark) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackingEventBillOfLadingBookmark) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
