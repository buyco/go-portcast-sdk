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

// checks if the PortTerminalAddOnImportPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortTerminalAddOnImportPlan{}

// PortTerminalAddOnImportPlan Summry of key temrinal events and statuses at the import Port of Discharge (POD).
type PortTerminalAddOnImportPlan struct {
	// Unique Identifier for the Import Plan Object.
	Id *string `json:"id,omitempty"`
	// Import Plan Object update date
	Updated *time.Time `json:"updated,omitempty"`
	// Relates to the UNLOCODE for the import port, as standardized by Portcast
	PortCode *string `json:"port_code,omitempty"`
	// Relates to the import port name, as standardized by Portcast.
	PortName *string `json:"port_name,omitempty"`
	// BIC/SMDG Defined Code for the Port Terminal Location.
	FacilityCode *string `json:"facility_code,omitempty"`
	// Import port terminal name.
	FacilityName *string `json:"facility_name,omitempty"`
	// Vessel name of the vessel associated with the import location.
	VesselName *string `json:"vessel_name,omitempty"`
	// Voyage Number associated with the import journey, as reported by the Terminal.
	VoyageNo *string `json:"voyage_no,omitempty"`
	// Container status, as reported by the Terminal.
	ContainerStatus *string `json:"container_status,omitempty"`
	// Date and time when the container was discharged from the vessel, as reported by the Terminal - Local Time.
	DischargeDate *string `json:"discharge_date,omitempty"`
	// Current location of the container in the yard.
	YardLocation *string `json:"yard_location,omitempty"`
	// Identifies any holds that might be applicable to this shipment.
	ContainerHolds *string `json:"container_holds,omitempty"`
	// Last Free Day, as reported by the Terminal. Charges will start to acrue after this date.
	LastFreeDay *string `json:"last_free_day,omitempty"`
	// Dollar amount owed in demurrage fees, as reported by the Terminal. 
	DemurrageOwed *string `json:"demurrage_owed,omitempty"`
	// Date and time when the container gates out of the port terminal, as reported by the Terminal - Local Time.
	GateOutDate *string `json:"gate_out_date,omitempty"`
	// Latest vessel ETA, as reported by the Terminal - Local Time.
	LatestEta *string `json:"latest_eta,omitempty"`
	// Confirms if the container is ready for delivery.
	ReadyForDelivery *string `json:"ready_for_delivery,omitempty"`
	// Confirms if the pick up appointment has been set.
	AppointmentSet *string `json:"appointment_set,omitempty"`
	// Date and time when container is scheduled to be picked up from the port, as reported by the Terminal - Local Time. 
	PickupAppointmentDate *string `json:"pickup_appointment_date,omitempty"`
}

// NewPortTerminalAddOnImportPlan instantiates a new PortTerminalAddOnImportPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortTerminalAddOnImportPlan() *PortTerminalAddOnImportPlan {
	this := PortTerminalAddOnImportPlan{}
	return &this
}

// NewPortTerminalAddOnImportPlanWithDefaults instantiates a new PortTerminalAddOnImportPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortTerminalAddOnImportPlanWithDefaults() *PortTerminalAddOnImportPlan {
	this := PortTerminalAddOnImportPlan{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PortTerminalAddOnImportPlan) SetId(v string) {
	o.Id = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *PortTerminalAddOnImportPlan) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetPortCode returns the PortCode field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetPortCode() string {
	if o == nil || IsNil(o.PortCode) {
		var ret string
		return ret
	}
	return *o.PortCode
}

// GetPortCodeOk returns a tuple with the PortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetPortCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PortCode) {
		return nil, false
	}
	return o.PortCode, true
}

// HasPortCode returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasPortCode() bool {
	if o != nil && !IsNil(o.PortCode) {
		return true
	}

	return false
}

// SetPortCode gets a reference to the given string and assigns it to the PortCode field.
func (o *PortTerminalAddOnImportPlan) SetPortCode(v string) {
	o.PortCode = &v
}

// GetPortName returns the PortName field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetPortName() string {
	if o == nil || IsNil(o.PortName) {
		var ret string
		return ret
	}
	return *o.PortName
}

// GetPortNameOk returns a tuple with the PortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetPortNameOk() (*string, bool) {
	if o == nil || IsNil(o.PortName) {
		return nil, false
	}
	return o.PortName, true
}

// HasPortName returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasPortName() bool {
	if o != nil && !IsNil(o.PortName) {
		return true
	}

	return false
}

// SetPortName gets a reference to the given string and assigns it to the PortName field.
func (o *PortTerminalAddOnImportPlan) SetPortName(v string) {
	o.PortName = &v
}

// GetFacilityCode returns the FacilityCode field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetFacilityCode() string {
	if o == nil || IsNil(o.FacilityCode) {
		var ret string
		return ret
	}
	return *o.FacilityCode
}

// GetFacilityCodeOk returns a tuple with the FacilityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetFacilityCodeOk() (*string, bool) {
	if o == nil || IsNil(o.FacilityCode) {
		return nil, false
	}
	return o.FacilityCode, true
}

// HasFacilityCode returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasFacilityCode() bool {
	if o != nil && !IsNil(o.FacilityCode) {
		return true
	}

	return false
}

// SetFacilityCode gets a reference to the given string and assigns it to the FacilityCode field.
func (o *PortTerminalAddOnImportPlan) SetFacilityCode(v string) {
	o.FacilityCode = &v
}

// GetFacilityName returns the FacilityName field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetFacilityName() string {
	if o == nil || IsNil(o.FacilityName) {
		var ret string
		return ret
	}
	return *o.FacilityName
}

// GetFacilityNameOk returns a tuple with the FacilityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetFacilityNameOk() (*string, bool) {
	if o == nil || IsNil(o.FacilityName) {
		return nil, false
	}
	return o.FacilityName, true
}

// HasFacilityName returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasFacilityName() bool {
	if o != nil && !IsNil(o.FacilityName) {
		return true
	}

	return false
}

// SetFacilityName gets a reference to the given string and assigns it to the FacilityName field.
func (o *PortTerminalAddOnImportPlan) SetFacilityName(v string) {
	o.FacilityName = &v
}

// GetVesselName returns the VesselName field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetVesselName() string {
	if o == nil || IsNil(o.VesselName) {
		var ret string
		return ret
	}
	return *o.VesselName
}

// GetVesselNameOk returns a tuple with the VesselName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetVesselNameOk() (*string, bool) {
	if o == nil || IsNil(o.VesselName) {
		return nil, false
	}
	return o.VesselName, true
}

// HasVesselName returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasVesselName() bool {
	if o != nil && !IsNil(o.VesselName) {
		return true
	}

	return false
}

// SetVesselName gets a reference to the given string and assigns it to the VesselName field.
func (o *PortTerminalAddOnImportPlan) SetVesselName(v string) {
	o.VesselName = &v
}

// GetVoyageNo returns the VoyageNo field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetVoyageNo() string {
	if o == nil || IsNil(o.VoyageNo) {
		var ret string
		return ret
	}
	return *o.VoyageNo
}

// GetVoyageNoOk returns a tuple with the VoyageNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetVoyageNoOk() (*string, bool) {
	if o == nil || IsNil(o.VoyageNo) {
		return nil, false
	}
	return o.VoyageNo, true
}

// HasVoyageNo returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasVoyageNo() bool {
	if o != nil && !IsNil(o.VoyageNo) {
		return true
	}

	return false
}

// SetVoyageNo gets a reference to the given string and assigns it to the VoyageNo field.
func (o *PortTerminalAddOnImportPlan) SetVoyageNo(v string) {
	o.VoyageNo = &v
}

// GetContainerStatus returns the ContainerStatus field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetContainerStatus() string {
	if o == nil || IsNil(o.ContainerStatus) {
		var ret string
		return ret
	}
	return *o.ContainerStatus
}

// GetContainerStatusOk returns a tuple with the ContainerStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetContainerStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerStatus) {
		return nil, false
	}
	return o.ContainerStatus, true
}

// HasContainerStatus returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasContainerStatus() bool {
	if o != nil && !IsNil(o.ContainerStatus) {
		return true
	}

	return false
}

// SetContainerStatus gets a reference to the given string and assigns it to the ContainerStatus field.
func (o *PortTerminalAddOnImportPlan) SetContainerStatus(v string) {
	o.ContainerStatus = &v
}

// GetDischargeDate returns the DischargeDate field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetDischargeDate() string {
	if o == nil || IsNil(o.DischargeDate) {
		var ret string
		return ret
	}
	return *o.DischargeDate
}

// GetDischargeDateOk returns a tuple with the DischargeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetDischargeDateOk() (*string, bool) {
	if o == nil || IsNil(o.DischargeDate) {
		return nil, false
	}
	return o.DischargeDate, true
}

// HasDischargeDate returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasDischargeDate() bool {
	if o != nil && !IsNil(o.DischargeDate) {
		return true
	}

	return false
}

// SetDischargeDate gets a reference to the given string and assigns it to the DischargeDate field.
func (o *PortTerminalAddOnImportPlan) SetDischargeDate(v string) {
	o.DischargeDate = &v
}

// GetYardLocation returns the YardLocation field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetYardLocation() string {
	if o == nil || IsNil(o.YardLocation) {
		var ret string
		return ret
	}
	return *o.YardLocation
}

// GetYardLocationOk returns a tuple with the YardLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetYardLocationOk() (*string, bool) {
	if o == nil || IsNil(o.YardLocation) {
		return nil, false
	}
	return o.YardLocation, true
}

// HasYardLocation returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasYardLocation() bool {
	if o != nil && !IsNil(o.YardLocation) {
		return true
	}

	return false
}

// SetYardLocation gets a reference to the given string and assigns it to the YardLocation field.
func (o *PortTerminalAddOnImportPlan) SetYardLocation(v string) {
	o.YardLocation = &v
}

// GetContainerHolds returns the ContainerHolds field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetContainerHolds() string {
	if o == nil || IsNil(o.ContainerHolds) {
		var ret string
		return ret
	}
	return *o.ContainerHolds
}

// GetContainerHoldsOk returns a tuple with the ContainerHolds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetContainerHoldsOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerHolds) {
		return nil, false
	}
	return o.ContainerHolds, true
}

// HasContainerHolds returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasContainerHolds() bool {
	if o != nil && !IsNil(o.ContainerHolds) {
		return true
	}

	return false
}

// SetContainerHolds gets a reference to the given string and assigns it to the ContainerHolds field.
func (o *PortTerminalAddOnImportPlan) SetContainerHolds(v string) {
	o.ContainerHolds = &v
}

// GetLastFreeDay returns the LastFreeDay field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetLastFreeDay() string {
	if o == nil || IsNil(o.LastFreeDay) {
		var ret string
		return ret
	}
	return *o.LastFreeDay
}

// GetLastFreeDayOk returns a tuple with the LastFreeDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetLastFreeDayOk() (*string, bool) {
	if o == nil || IsNil(o.LastFreeDay) {
		return nil, false
	}
	return o.LastFreeDay, true
}

// HasLastFreeDay returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasLastFreeDay() bool {
	if o != nil && !IsNil(o.LastFreeDay) {
		return true
	}

	return false
}

// SetLastFreeDay gets a reference to the given string and assigns it to the LastFreeDay field.
func (o *PortTerminalAddOnImportPlan) SetLastFreeDay(v string) {
	o.LastFreeDay = &v
}

// GetDemurrageOwed returns the DemurrageOwed field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetDemurrageOwed() string {
	if o == nil || IsNil(o.DemurrageOwed) {
		var ret string
		return ret
	}
	return *o.DemurrageOwed
}

// GetDemurrageOwedOk returns a tuple with the DemurrageOwed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetDemurrageOwedOk() (*string, bool) {
	if o == nil || IsNil(o.DemurrageOwed) {
		return nil, false
	}
	return o.DemurrageOwed, true
}

// HasDemurrageOwed returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasDemurrageOwed() bool {
	if o != nil && !IsNil(o.DemurrageOwed) {
		return true
	}

	return false
}

// SetDemurrageOwed gets a reference to the given string and assigns it to the DemurrageOwed field.
func (o *PortTerminalAddOnImportPlan) SetDemurrageOwed(v string) {
	o.DemurrageOwed = &v
}

// GetGateOutDate returns the GateOutDate field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetGateOutDate() string {
	if o == nil || IsNil(o.GateOutDate) {
		var ret string
		return ret
	}
	return *o.GateOutDate
}

// GetGateOutDateOk returns a tuple with the GateOutDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetGateOutDateOk() (*string, bool) {
	if o == nil || IsNil(o.GateOutDate) {
		return nil, false
	}
	return o.GateOutDate, true
}

// HasGateOutDate returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasGateOutDate() bool {
	if o != nil && !IsNil(o.GateOutDate) {
		return true
	}

	return false
}

// SetGateOutDate gets a reference to the given string and assigns it to the GateOutDate field.
func (o *PortTerminalAddOnImportPlan) SetGateOutDate(v string) {
	o.GateOutDate = &v
}

// GetLatestEta returns the LatestEta field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetLatestEta() string {
	if o == nil || IsNil(o.LatestEta) {
		var ret string
		return ret
	}
	return *o.LatestEta
}

// GetLatestEtaOk returns a tuple with the LatestEta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetLatestEtaOk() (*string, bool) {
	if o == nil || IsNil(o.LatestEta) {
		return nil, false
	}
	return o.LatestEta, true
}

// HasLatestEta returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasLatestEta() bool {
	if o != nil && !IsNil(o.LatestEta) {
		return true
	}

	return false
}

// SetLatestEta gets a reference to the given string and assigns it to the LatestEta field.
func (o *PortTerminalAddOnImportPlan) SetLatestEta(v string) {
	o.LatestEta = &v
}

// GetReadyForDelivery returns the ReadyForDelivery field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetReadyForDelivery() string {
	if o == nil || IsNil(o.ReadyForDelivery) {
		var ret string
		return ret
	}
	return *o.ReadyForDelivery
}

// GetReadyForDeliveryOk returns a tuple with the ReadyForDelivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetReadyForDeliveryOk() (*string, bool) {
	if o == nil || IsNil(o.ReadyForDelivery) {
		return nil, false
	}
	return o.ReadyForDelivery, true
}

// HasReadyForDelivery returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasReadyForDelivery() bool {
	if o != nil && !IsNil(o.ReadyForDelivery) {
		return true
	}

	return false
}

// SetReadyForDelivery gets a reference to the given string and assigns it to the ReadyForDelivery field.
func (o *PortTerminalAddOnImportPlan) SetReadyForDelivery(v string) {
	o.ReadyForDelivery = &v
}

// GetAppointmentSet returns the AppointmentSet field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetAppointmentSet() string {
	if o == nil || IsNil(o.AppointmentSet) {
		var ret string
		return ret
	}
	return *o.AppointmentSet
}

// GetAppointmentSetOk returns a tuple with the AppointmentSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetAppointmentSetOk() (*string, bool) {
	if o == nil || IsNil(o.AppointmentSet) {
		return nil, false
	}
	return o.AppointmentSet, true
}

// HasAppointmentSet returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasAppointmentSet() bool {
	if o != nil && !IsNil(o.AppointmentSet) {
		return true
	}

	return false
}

// SetAppointmentSet gets a reference to the given string and assigns it to the AppointmentSet field.
func (o *PortTerminalAddOnImportPlan) SetAppointmentSet(v string) {
	o.AppointmentSet = &v
}

// GetPickupAppointmentDate returns the PickupAppointmentDate field value if set, zero value otherwise.
func (o *PortTerminalAddOnImportPlan) GetPickupAppointmentDate() string {
	if o == nil || IsNil(o.PickupAppointmentDate) {
		var ret string
		return ret
	}
	return *o.PickupAppointmentDate
}

// GetPickupAppointmentDateOk returns a tuple with the PickupAppointmentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortTerminalAddOnImportPlan) GetPickupAppointmentDateOk() (*string, bool) {
	if o == nil || IsNil(o.PickupAppointmentDate) {
		return nil, false
	}
	return o.PickupAppointmentDate, true
}

// HasPickupAppointmentDate returns a boolean if a field has been set.
func (o *PortTerminalAddOnImportPlan) HasPickupAppointmentDate() bool {
	if o != nil && !IsNil(o.PickupAppointmentDate) {
		return true
	}

	return false
}

// SetPickupAppointmentDate gets a reference to the given string and assigns it to the PickupAppointmentDate field.
func (o *PortTerminalAddOnImportPlan) SetPickupAppointmentDate(v string) {
	o.PickupAppointmentDate = &v
}

func (o PortTerminalAddOnImportPlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortTerminalAddOnImportPlan) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ContainerStatus) {
		toSerialize["container_status"] = o.ContainerStatus
	}
	if !IsNil(o.DischargeDate) {
		toSerialize["discharge_date"] = o.DischargeDate
	}
	if !IsNil(o.YardLocation) {
		toSerialize["yard_location"] = o.YardLocation
	}
	if !IsNil(o.ContainerHolds) {
		toSerialize["container_holds"] = o.ContainerHolds
	}
	if !IsNil(o.LastFreeDay) {
		toSerialize["last_free_day"] = o.LastFreeDay
	}
	if !IsNil(o.DemurrageOwed) {
		toSerialize["demurrage_owed"] = o.DemurrageOwed
	}
	if !IsNil(o.GateOutDate) {
		toSerialize["gate_out_date"] = o.GateOutDate
	}
	if !IsNil(o.LatestEta) {
		toSerialize["latest_eta"] = o.LatestEta
	}
	if !IsNil(o.ReadyForDelivery) {
		toSerialize["ready_for_delivery"] = o.ReadyForDelivery
	}
	if !IsNil(o.AppointmentSet) {
		toSerialize["appointment_set"] = o.AppointmentSet
	}
	if !IsNil(o.PickupAppointmentDate) {
		toSerialize["pickup_appointment_date"] = o.PickupAppointmentDate
	}
	return toSerialize, nil
}

type NullablePortTerminalAddOnImportPlan struct {
	value *PortTerminalAddOnImportPlan
	isSet bool
}

func (v NullablePortTerminalAddOnImportPlan) Get() *PortTerminalAddOnImportPlan {
	return v.value
}

func (v *NullablePortTerminalAddOnImportPlan) Set(val *PortTerminalAddOnImportPlan) {
	v.value = val
	v.isSet = true
}

func (v NullablePortTerminalAddOnImportPlan) IsSet() bool {
	return v.isSet
}

func (v *NullablePortTerminalAddOnImportPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortTerminalAddOnImportPlan(val *PortTerminalAddOnImportPlan) *NullablePortTerminalAddOnImportPlan {
	return &NullablePortTerminalAddOnImportPlan{value: val, isSet: true}
}

func (v NullablePortTerminalAddOnImportPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortTerminalAddOnImportPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


