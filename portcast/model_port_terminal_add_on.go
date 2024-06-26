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

// checks if the PortTerminalAddOn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortTerminalAddOn{}

// PortTerminalAddOn Portcast Port Terminal Data Add On
type PortTerminalAddOn struct {
	// Unique Identifier called Bill of Lading Bookmark ID [Primary Key] (Also refered as Bookmark ID)
	BillOfLadingBookmarkId *string `json:"bill_of_lading_bookmark_id,omitempty"`
	// Container number
	CntrNo *string `json:"cntr_no,omitempty"`
	// Bill of Lading or Booking Number
	BlNo *string `json:"bl_no,omitempty"`
	ExportPlan *PortTerminalAddOnExportPlan `json:"export_plan,omitempty"`
	ImportPlan *PortTerminalAddOnImportPlan `json:"import_plan,omitempty"`
}

// NewPortTerminalAddOn instantiates a new PortTerminalAddOn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortTerminalAddOn() *PortTerminalAddOn {
	this := PortTerminalAddOn{}
	return &this
}

// NewPortTerminalAddOnWithDefaults instantiates a new PortTerminalAddOn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortTerminalAddOnWithDefaults() *PortTerminalAddOn {
	this := PortTerminalAddOn{}
	return &this
}

// GetBillOfLadingBookmarkId returns the BillOfLadingBookmarkId field value if set, zero value otherwise.
func (o *PortTerminalAddOn) GetBillOfLadingBookmarkId() string {
	if o == nil || IsNil(o.BillOfLadingBookmarkId) {
		var ret string
		return ret
	}
	return *o.BillOfLadingBookmarkId
}

// GetBillOfLadingBookmarkIdOk returns a tuple with the BillOfLadingBookmarkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOn) GetBillOfLadingBookmarkIdOk() (*string, bool) {
	if o == nil || IsNil(o.BillOfLadingBookmarkId) {
		return nil, false
	}
	return o.BillOfLadingBookmarkId, true
}

// HasBillOfLadingBookmarkId returns a boolean if a field has been set.
func (o *PortTerminalAddOn) HasBillOfLadingBookmarkId() bool {
	if o != nil && !IsNil(o.BillOfLadingBookmarkId) {
		return true
	}

	return false
}

// SetBillOfLadingBookmarkId gets a reference to the given string and assigns it to the BillOfLadingBookmarkId field.
func (o *PortTerminalAddOn) SetBillOfLadingBookmarkId(v string) {
	o.BillOfLadingBookmarkId = &v
}

// GetCntrNo returns the CntrNo field value if set, zero value otherwise.
func (o *PortTerminalAddOn) GetCntrNo() string {
	if o == nil || IsNil(o.CntrNo) {
		var ret string
		return ret
	}
	return *o.CntrNo
}

// GetCntrNoOk returns a tuple with the CntrNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOn) GetCntrNoOk() (*string, bool) {
	if o == nil || IsNil(o.CntrNo) {
		return nil, false
	}
	return o.CntrNo, true
}

// HasCntrNo returns a boolean if a field has been set.
func (o *PortTerminalAddOn) HasCntrNo() bool {
	if o != nil && !IsNil(o.CntrNo) {
		return true
	}

	return false
}

// SetCntrNo gets a reference to the given string and assigns it to the CntrNo field.
func (o *PortTerminalAddOn) SetCntrNo(v string) {
	o.CntrNo = &v
}

// GetBlNo returns the BlNo field value if set, zero value otherwise.
func (o *PortTerminalAddOn) GetBlNo() string {
	if o == nil || IsNil(o.BlNo) {
		var ret string
		return ret
	}
	return *o.BlNo
}

// GetBlNoOk returns a tuple with the BlNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOn) GetBlNoOk() (*string, bool) {
	if o == nil || IsNil(o.BlNo) {
		return nil, false
	}
	return o.BlNo, true
}

// HasBlNo returns a boolean if a field has been set.
func (o *PortTerminalAddOn) HasBlNo() bool {
	if o != nil && !IsNil(o.BlNo) {
		return true
	}

	return false
}

// SetBlNo gets a reference to the given string and assigns it to the BlNo field.
func (o *PortTerminalAddOn) SetBlNo(v string) {
	o.BlNo = &v
}

// GetExportPlan returns the ExportPlan field value if set, zero value otherwise.
func (o *PortTerminalAddOn) GetExportPlan() PortTerminalAddOnExportPlan {
	if o == nil || IsNil(o.ExportPlan) {
		var ret PortTerminalAddOnExportPlan
		return ret
	}
	return *o.ExportPlan
}

// GetExportPlanOk returns a tuple with the ExportPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOn) GetExportPlanOk() (*PortTerminalAddOnExportPlan, bool) {
	if o == nil || IsNil(o.ExportPlan) {
		return nil, false
	}
	return o.ExportPlan, true
}

// HasExportPlan returns a boolean if a field has been set.
func (o *PortTerminalAddOn) HasExportPlan() bool {
	if o != nil && !IsNil(o.ExportPlan) {
		return true
	}

	return false
}

// SetExportPlan gets a reference to the given PortTerminalAddOnExportPlan and assigns it to the ExportPlan field.
func (o *PortTerminalAddOn) SetExportPlan(v PortTerminalAddOnExportPlan) {
	o.ExportPlan = &v
}

// GetImportPlan returns the ImportPlan field value if set, zero value otherwise.
func (o *PortTerminalAddOn) GetImportPlan() PortTerminalAddOnImportPlan {
	if o == nil || IsNil(o.ImportPlan) {
		var ret PortTerminalAddOnImportPlan
		return ret
	}
	return *o.ImportPlan
}

// GetImportPlanOk returns a tuple with the ImportPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOn) GetImportPlanOk() (*PortTerminalAddOnImportPlan, bool) {
	if o == nil || IsNil(o.ImportPlan) {
		return nil, false
	}
	return o.ImportPlan, true
}

// HasImportPlan returns a boolean if a field has been set.
func (o *PortTerminalAddOn) HasImportPlan() bool {
	if o != nil && !IsNil(o.ImportPlan) {
		return true
	}

	return false
}

// SetImportPlan gets a reference to the given PortTerminalAddOnImportPlan and assigns it to the ImportPlan field.
func (o *PortTerminalAddOn) SetImportPlan(v PortTerminalAddOnImportPlan) {
	o.ImportPlan = &v
}

func (o PortTerminalAddOn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortTerminalAddOn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillOfLadingBookmarkId) {
		toSerialize["bill_of_lading_bookmark_id"] = o.BillOfLadingBookmarkId
	}
	if !IsNil(o.CntrNo) {
		toSerialize["cntr_no"] = o.CntrNo
	}
	if !IsNil(o.BlNo) {
		toSerialize["bl_no"] = o.BlNo
	}
	if !IsNil(o.ExportPlan) {
		toSerialize["export_plan"] = o.ExportPlan
	}
	if !IsNil(o.ImportPlan) {
		toSerialize["import_plan"] = o.ImportPlan
	}
	return toSerialize, nil
}

type NullablePortTerminalAddOn struct {
	value *PortTerminalAddOn
	isSet bool
}

func (v NullablePortTerminalAddOn) Get() *PortTerminalAddOn {
	return v.value
}

func (v *NullablePortTerminalAddOn) Set(val *PortTerminalAddOn) {
	v.value = val
	v.isSet = true
}

func (v NullablePortTerminalAddOn) IsSet() bool {
	return v.isSet
}

func (v *NullablePortTerminalAddOn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortTerminalAddOn(val *PortTerminalAddOn) *NullablePortTerminalAddOn {
	return &NullablePortTerminalAddOn{value: val, isSet: true}
}

func (v NullablePortTerminalAddOn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortTerminalAddOn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


