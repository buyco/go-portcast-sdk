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
	"time"
)

// checks if the PortTerminalAddOnExportPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortTerminalAddOnExportPlan{}

// PortTerminalAddOnExportPlan Summry of key temrinal events and milestones at the export Port of Loading (POL).
type PortTerminalAddOnExportPlan struct {
	// Unique Identifier for the Export Plan Object.
	Id *string `json:"id,omitempty"`
	// Export Plan Object update date
	Updated *time.Time `json:"updated,omitempty"`
	// Relates to the UNLOCODE for the export port, as standardized by Portcast
	PortCode *string `json:"port_code,omitempty"`
	// Relates to the export port name, as standardized by Portcast
	PortName *string `json:"port_name,omitempty"`
	// BIC/SMDG Defined Code for the Port Terminal Location
	FacilityCode *string `json:"facility_code,omitempty"`
	// Export port terminal name
	FacilityName *string `json:"facility_name,omitempty"`
	// Vessel name of the vessel associated with the export location.
	VesselName *string `json:"vessel_name,omitempty"`
	// Voyage Number associated with the export journey, as reported by the Terminal.
	VoyageNo *string `json:"voyage_no,omitempty"`
	// Earliest Receipt Date (ERD) for a standard container, as reported by the terminal - Local Time.
	ErdStandard *time.Time `json:"erd_standard,omitempty"`
	// Earliest Receipt Date (ERD) for a reefer container, as reported by the terminal - Local Time.
	ErdReefer *time.Time `json:"erd_reefer,omitempty"`
	// Date a container was received at the terminal through the gate, as reported by the terminal - Local Time.
	GateInDate *time.Time `json:"gate_in_date,omitempty"`
	// Last possible date and time for gating-in standard container at the export terminal, as reported by the Terminal - Local Time.
	PortCutoffStandard *time.Time `json:"port_cutoff_standard,omitempty"`
	// Last possible date and time for gating-in reefer container at the export terminal, as reported by the terminal - Local Time.
	PortCutoffReefer *time.Time `json:"port_cutoff_reefer,omitempty"`
	// Latest vessel ETA to the export port, as reported by the terminal - Local Time.
	LatestEta *time.Time `json:"latest_eta,omitempty"`
	// Actual vessel time of arrival, as reported by the terminal - Local Time.
	ActualArrival *time.Time `json:"actual_arrival,omitempty"`
	// Latest vessel ETD from the export port, as reported by the terminal - Local Time.
	LatestEtd *time.Time `json:"latest_etd,omitempty"`
	// Actual vessel time of departure from the export port, as reported by the terminal - Local Time.
	ActualDeparture *time.Time `json:"actual_departure,omitempty"`
}

// NewPortTerminalAddOnExportPlan instantiates a new PortTerminalAddOnExportPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortTerminalAddOnExportPlan() *PortTerminalAddOnExportPlan {
	this := PortTerminalAddOnExportPlan{}
	return &this
}

// NewPortTerminalAddOnExportPlanWithDefaults instantiates a new PortTerminalAddOnExportPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortTerminalAddOnExportPlanWithDefaults() *PortTerminalAddOnExportPlan {
	this := PortTerminalAddOnExportPlan{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PortTerminalAddOnExportPlan) SetId(v string) {
	o.Id = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *PortTerminalAddOnExportPlan) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetPortCode returns the PortCode field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetPortCode() string {
	if o == nil || IsNil(o.PortCode) {
		var ret string
		return ret
	}
	return *o.PortCode
}

// GetPortCodeOk returns a tuple with the PortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetPortCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PortCode) {
		return nil, false
	}
	return o.PortCode, true
}

// HasPortCode returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasPortCode() bool {
	if o != nil && !IsNil(o.PortCode) {
		return true
	}

	return false
}

// SetPortCode gets a reference to the given string and assigns it to the PortCode field.
func (o *PortTerminalAddOnExportPlan) SetPortCode(v string) {
	o.PortCode = &v
}

// GetPortName returns the PortName field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetPortName() string {
	if o == nil || IsNil(o.PortName) {
		var ret string
		return ret
	}
	return *o.PortName
}

// GetPortNameOk returns a tuple with the PortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetPortNameOk() (*string, bool) {
	if o == nil || IsNil(o.PortName) {
		return nil, false
	}
	return o.PortName, true
}

// HasPortName returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasPortName() bool {
	if o != nil && !IsNil(o.PortName) {
		return true
	}

	return false
}

// SetPortName gets a reference to the given string and assigns it to the PortName field.
func (o *PortTerminalAddOnExportPlan) SetPortName(v string) {
	o.PortName = &v
}

// GetFacilityCode returns the FacilityCode field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetFacilityCode() string {
	if o == nil || IsNil(o.FacilityCode) {
		var ret string
		return ret
	}
	return *o.FacilityCode
}

// GetFacilityCodeOk returns a tuple with the FacilityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetFacilityCodeOk() (*string, bool) {
	if o == nil || IsNil(o.FacilityCode) {
		return nil, false
	}
	return o.FacilityCode, true
}

// HasFacilityCode returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasFacilityCode() bool {
	if o != nil && !IsNil(o.FacilityCode) {
		return true
	}

	return false
}

// SetFacilityCode gets a reference to the given string and assigns it to the FacilityCode field.
func (o *PortTerminalAddOnExportPlan) SetFacilityCode(v string) {
	o.FacilityCode = &v
}

// GetFacilityName returns the FacilityName field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetFacilityName() string {
	if o == nil || IsNil(o.FacilityName) {
		var ret string
		return ret
	}
	return *o.FacilityName
}

// GetFacilityNameOk returns a tuple with the FacilityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetFacilityNameOk() (*string, bool) {
	if o == nil || IsNil(o.FacilityName) {
		return nil, false
	}
	return o.FacilityName, true
}

// HasFacilityName returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasFacilityName() bool {
	if o != nil && !IsNil(o.FacilityName) {
		return true
	}

	return false
}

// SetFacilityName gets a reference to the given string and assigns it to the FacilityName field.
func (o *PortTerminalAddOnExportPlan) SetFacilityName(v string) {
	o.FacilityName = &v
}

// GetVesselName returns the VesselName field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetVesselName() string {
	if o == nil || IsNil(o.VesselName) {
		var ret string
		return ret
	}
	return *o.VesselName
}

// GetVesselNameOk returns a tuple with the VesselName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetVesselNameOk() (*string, bool) {
	if o == nil || IsNil(o.VesselName) {
		return nil, false
	}
	return o.VesselName, true
}

// HasVesselName returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasVesselName() bool {
	if o != nil && !IsNil(o.VesselName) {
		return true
	}

	return false
}

// SetVesselName gets a reference to the given string and assigns it to the VesselName field.
func (o *PortTerminalAddOnExportPlan) SetVesselName(v string) {
	o.VesselName = &v
}

// GetVoyageNo returns the VoyageNo field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetVoyageNo() string {
	if o == nil || IsNil(o.VoyageNo) {
		var ret string
		return ret
	}
	return *o.VoyageNo
}

// GetVoyageNoOk returns a tuple with the VoyageNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetVoyageNoOk() (*string, bool) {
	if o == nil || IsNil(o.VoyageNo) {
		return nil, false
	}
	return o.VoyageNo, true
}

// HasVoyageNo returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasVoyageNo() bool {
	if o != nil && !IsNil(o.VoyageNo) {
		return true
	}

	return false
}

// SetVoyageNo gets a reference to the given string and assigns it to the VoyageNo field.
func (o *PortTerminalAddOnExportPlan) SetVoyageNo(v string) {
	o.VoyageNo = &v
}

// GetErdStandard returns the ErdStandard field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetErdStandard() time.Time {
	if o == nil || IsNil(o.ErdStandard) {
		var ret time.Time
		return ret
	}
	return *o.ErdStandard
}

// GetErdStandardOk returns a tuple with the ErdStandard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetErdStandardOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ErdStandard) {
		return nil, false
	}
	return o.ErdStandard, true
}

// HasErdStandard returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasErdStandard() bool {
	if o != nil && !IsNil(o.ErdStandard) {
		return true
	}

	return false
}

// SetErdStandard gets a reference to the given time.Time and assigns it to the ErdStandard field.
func (o *PortTerminalAddOnExportPlan) SetErdStandard(v time.Time) {
	o.ErdStandard = &v
}

// GetErdReefer returns the ErdReefer field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetErdReefer() time.Time {
	if o == nil || IsNil(o.ErdReefer) {
		var ret time.Time
		return ret
	}
	return *o.ErdReefer
}

// GetErdReeferOk returns a tuple with the ErdReefer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetErdReeferOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ErdReefer) {
		return nil, false
	}
	return o.ErdReefer, true
}

// HasErdReefer returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasErdReefer() bool {
	if o != nil && !IsNil(o.ErdReefer) {
		return true
	}

	return false
}

// SetErdReefer gets a reference to the given time.Time and assigns it to the ErdReefer field.
func (o *PortTerminalAddOnExportPlan) SetErdReefer(v time.Time) {
	o.ErdReefer = &v
}

// GetGateInDate returns the GateInDate field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetGateInDate() time.Time {
	if o == nil || IsNil(o.GateInDate) {
		var ret time.Time
		return ret
	}
	return *o.GateInDate
}

// GetGateInDateOk returns a tuple with the GateInDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetGateInDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.GateInDate) {
		return nil, false
	}
	return o.GateInDate, true
}

// HasGateInDate returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasGateInDate() bool {
	if o != nil && !IsNil(o.GateInDate) {
		return true
	}

	return false
}

// SetGateInDate gets a reference to the given time.Time and assigns it to the GateInDate field.
func (o *PortTerminalAddOnExportPlan) SetGateInDate(v time.Time) {
	o.GateInDate = &v
}

// GetPortCutoffStandard returns the PortCutoffStandard field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetPortCutoffStandard() time.Time {
	if o == nil || IsNil(o.PortCutoffStandard) {
		var ret time.Time
		return ret
	}
	return *o.PortCutoffStandard
}

// GetPortCutoffStandardOk returns a tuple with the PortCutoffStandard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetPortCutoffStandardOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PortCutoffStandard) {
		return nil, false
	}
	return o.PortCutoffStandard, true
}

// HasPortCutoffStandard returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasPortCutoffStandard() bool {
	if o != nil && !IsNil(o.PortCutoffStandard) {
		return true
	}

	return false
}

// SetPortCutoffStandard gets a reference to the given time.Time and assigns it to the PortCutoffStandard field.
func (o *PortTerminalAddOnExportPlan) SetPortCutoffStandard(v time.Time) {
	o.PortCutoffStandard = &v
}

// GetPortCutoffReefer returns the PortCutoffReefer field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetPortCutoffReefer() time.Time {
	if o == nil || IsNil(o.PortCutoffReefer) {
		var ret time.Time
		return ret
	}
	return *o.PortCutoffReefer
}

// GetPortCutoffReeferOk returns a tuple with the PortCutoffReefer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetPortCutoffReeferOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PortCutoffReefer) {
		return nil, false
	}
	return o.PortCutoffReefer, true
}

// HasPortCutoffReefer returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasPortCutoffReefer() bool {
	if o != nil && !IsNil(o.PortCutoffReefer) {
		return true
	}

	return false
}

// SetPortCutoffReefer gets a reference to the given time.Time and assigns it to the PortCutoffReefer field.
func (o *PortTerminalAddOnExportPlan) SetPortCutoffReefer(v time.Time) {
	o.PortCutoffReefer = &v
}

// GetLatestEta returns the LatestEta field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetLatestEta() time.Time {
	if o == nil || IsNil(o.LatestEta) {
		var ret time.Time
		return ret
	}
	return *o.LatestEta
}

// GetLatestEtaOk returns a tuple with the LatestEta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetLatestEtaOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LatestEta) {
		return nil, false
	}
	return o.LatestEta, true
}

// HasLatestEta returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasLatestEta() bool {
	if o != nil && !IsNil(o.LatestEta) {
		return true
	}

	return false
}

// SetLatestEta gets a reference to the given time.Time and assigns it to the LatestEta field.
func (o *PortTerminalAddOnExportPlan) SetLatestEta(v time.Time) {
	o.LatestEta = &v
}

// GetActualArrival returns the ActualArrival field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetActualArrival() time.Time {
	if o == nil || IsNil(o.ActualArrival) {
		var ret time.Time
		return ret
	}
	return *o.ActualArrival
}

// GetActualArrivalOk returns a tuple with the ActualArrival field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetActualArrivalOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ActualArrival) {
		return nil, false
	}
	return o.ActualArrival, true
}

// HasActualArrival returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasActualArrival() bool {
	if o != nil && !IsNil(o.ActualArrival) {
		return true
	}

	return false
}

// SetActualArrival gets a reference to the given time.Time and assigns it to the ActualArrival field.
func (o *PortTerminalAddOnExportPlan) SetActualArrival(v time.Time) {
	o.ActualArrival = &v
}

// GetLatestEtd returns the LatestEtd field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetLatestEtd() time.Time {
	if o == nil || IsNil(o.LatestEtd) {
		var ret time.Time
		return ret
	}
	return *o.LatestEtd
}

// GetLatestEtdOk returns a tuple with the LatestEtd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetLatestEtdOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LatestEtd) {
		return nil, false
	}
	return o.LatestEtd, true
}

// HasLatestEtd returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasLatestEtd() bool {
	if o != nil && !IsNil(o.LatestEtd) {
		return true
	}

	return false
}

// SetLatestEtd gets a reference to the given time.Time and assigns it to the LatestEtd field.
func (o *PortTerminalAddOnExportPlan) SetLatestEtd(v time.Time) {
	o.LatestEtd = &v
}

// GetActualDeparture returns the ActualDeparture field value if set, zero value otherwise.
func (o *PortTerminalAddOnExportPlan) GetActualDeparture() time.Time {
	if o == nil || IsNil(o.ActualDeparture) {
		var ret time.Time
		return ret
	}
	return *o.ActualDeparture
}

// GetActualDepartureOk returns a tuple with the ActualDeparture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnExportPlan) GetActualDepartureOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ActualDeparture) {
		return nil, false
	}
	return o.ActualDeparture, true
}

// HasActualDeparture returns a boolean if a field has been set.
func (o *PortTerminalAddOnExportPlan) HasActualDeparture() bool {
	if o != nil && !IsNil(o.ActualDeparture) {
		return true
	}

	return false
}

// SetActualDeparture gets a reference to the given time.Time and assigns it to the ActualDeparture field.
func (o *PortTerminalAddOnExportPlan) SetActualDeparture(v time.Time) {
	o.ActualDeparture = &v
}

func (o PortTerminalAddOnExportPlan) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortTerminalAddOnExportPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.PortCode) {
		toSerialize["port_code"] = o.PortCode
	}
	if !IsNil(o.PortName) {
		toSerialize["port_name"] = o.PortName
	}
	if !IsNil(o.FacilityCode) {
		toSerialize["facility_code"] = o.FacilityCode
	}
	if !IsNil(o.FacilityName) {
		toSerialize["facility_name"] = o.FacilityName
	}
	if !IsNil(o.VesselName) {
		toSerialize["vessel_name"] = o.VesselName
	}
	if !IsNil(o.VoyageNo) {
		toSerialize["voyage_no"] = o.VoyageNo
	}
	if !IsNil(o.ErdStandard) {
		toSerialize["erd_standard"] = o.ErdStandard
	}
	if !IsNil(o.ErdReefer) {
		toSerialize["erd_reefer"] = o.ErdReefer
	}
	if !IsNil(o.GateInDate) {
		toSerialize["gate_in_date"] = o.GateInDate
	}
	if !IsNil(o.PortCutoffStandard) {
		toSerialize["port_cutoff_standard"] = o.PortCutoffStandard
	}
	if !IsNil(o.PortCutoffReefer) {
		toSerialize["port_cutoff_reefer"] = o.PortCutoffReefer
	}
	if !IsNil(o.LatestEta) {
		toSerialize["latest_eta"] = o.LatestEta
	}
	if !IsNil(o.ActualArrival) {
		toSerialize["actual_arrival"] = o.ActualArrival
	}
	if !IsNil(o.LatestEtd) {
		toSerialize["latest_etd"] = o.LatestEtd
	}
	if !IsNil(o.ActualDeparture) {
		toSerialize["actual_departure"] = o.ActualDeparture
	}
	return toSerialize, nil
}

type NullablePortTerminalAddOnExportPlan struct {
	value *PortTerminalAddOnExportPlan
	isSet bool
}

func (v NullablePortTerminalAddOnExportPlan) Get() *PortTerminalAddOnExportPlan {
	return v.value
}

func (v *NullablePortTerminalAddOnExportPlan) Set(val *PortTerminalAddOnExportPlan) {
	v.value = val
	v.isSet = true
}

func (v NullablePortTerminalAddOnExportPlan) IsSet() bool {
	return v.isSet
}

func (v *NullablePortTerminalAddOnExportPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortTerminalAddOnExportPlan(val *PortTerminalAddOnExportPlan) *NullablePortTerminalAddOnExportPlan {
	return &NullablePortTerminalAddOnExportPlan{value: val, isSet: true}
}

func (v NullablePortTerminalAddOnExportPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortTerminalAddOnExportPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}