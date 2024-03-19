# PortTerminalAddOnImportPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique Identifier for the Import Plan Object. | [optional] 
**Updated** | Pointer to **time.Time** | Import Plan Object update date | [optional] 
**PortCode** | Pointer to **string** | Relates to the UNLOCODE for the import port, as standardized by Portcast | [optional] 
**PortName** | Pointer to **string** | Relates to the import port name, as standardized by Portcast. | [optional] 
**FacilityCode** | Pointer to **string** | BIC/SMDG Defined Code for the Port Terminal Location. | [optional] 
**FacilityName** | Pointer to **string** | Import port terminal name. | [optional] 
**VesselName** | Pointer to **string** | Vessel name of the vessel associated with the import location. | [optional] 
**VoyageNo** | Pointer to **string** | Voyage Number associated with the import journey, as reported by the Terminal. | [optional] 
**ContainerStatus** | Pointer to **string** | Container status, as reported by the Terminal. | [optional] 
**DischargeDate** | Pointer to **string** | Date and time when the container was discharged from the vessel, as reported by the Terminal - Local Time. | [optional] 
**YardLocation** | Pointer to **string** | Current location of the container in the yard. | [optional] 
**ContainerHolds** | Pointer to **string** | Identifies any holds that might be applicable to this shipment. | [optional] 
**LastFreeDay** | Pointer to **string** | Last Free Day, as reported by the Terminal. Charges will start to acrue after this date. | [optional] 
**DemurrageOwed** | Pointer to **string** | Dollar amount owed in demurrage fees, as reported by the Terminal.  | [optional] 
**GateOutDate** | Pointer to **string** | Date and time when the container gates out of the port terminal, as reported by the Terminal - Local Time. | [optional] 
**LatestEta** | Pointer to **string** | Latest vessel ETA, as reported by the Terminal - Local Time. | [optional] 
**ReadyForDelivery** | Pointer to **string** | Confirms if the container is ready for delivery. | [optional] 
**AppointmentSet** | Pointer to **string** | Confirms if the pick up appointment has been set. | [optional] 
**PickupAppointmentDate** | Pointer to **time.Time** | Date and time when container is scheduled to be picked up from the port, as reported by the Terminal - Local Time.  | [optional] 

## Methods

### NewPortTerminalAddOnImportPlan

`func NewPortTerminalAddOnImportPlan() *PortTerminalAddOnImportPlan`

NewPortTerminalAddOnImportPlan instantiates a new PortTerminalAddOnImportPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortTerminalAddOnImportPlanWithDefaults

`func NewPortTerminalAddOnImportPlanWithDefaults() *PortTerminalAddOnImportPlan`

NewPortTerminalAddOnImportPlanWithDefaults instantiates a new PortTerminalAddOnImportPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortTerminalAddOnImportPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortTerminalAddOnImportPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortTerminalAddOnImportPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortTerminalAddOnImportPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdated

`func (o *PortTerminalAddOnImportPlan) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PortTerminalAddOnImportPlan) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PortTerminalAddOnImportPlan) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PortTerminalAddOnImportPlan) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetPortCode

`func (o *PortTerminalAddOnImportPlan) GetPortCode() string`

GetPortCode returns the PortCode field if non-nil, zero value otherwise.

### GetPortCodeOk

`func (o *PortTerminalAddOnImportPlan) GetPortCodeOk() (*string, bool)`

GetPortCodeOk returns a tuple with the PortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCode

`func (o *PortTerminalAddOnImportPlan) SetPortCode(v string)`

SetPortCode sets PortCode field to given value.

### HasPortCode

`func (o *PortTerminalAddOnImportPlan) HasPortCode() bool`

HasPortCode returns a boolean if a field has been set.

### GetPortName

`func (o *PortTerminalAddOnImportPlan) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *PortTerminalAddOnImportPlan) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *PortTerminalAddOnImportPlan) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *PortTerminalAddOnImportPlan) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetFacilityCode

`func (o *PortTerminalAddOnImportPlan) GetFacilityCode() string`

GetFacilityCode returns the FacilityCode field if non-nil, zero value otherwise.

### GetFacilityCodeOk

`func (o *PortTerminalAddOnImportPlan) GetFacilityCodeOk() (*string, bool)`

GetFacilityCodeOk returns a tuple with the FacilityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCode

`func (o *PortTerminalAddOnImportPlan) SetFacilityCode(v string)`

SetFacilityCode sets FacilityCode field to given value.

### HasFacilityCode

`func (o *PortTerminalAddOnImportPlan) HasFacilityCode() bool`

HasFacilityCode returns a boolean if a field has been set.

### GetFacilityName

`func (o *PortTerminalAddOnImportPlan) GetFacilityName() string`

GetFacilityName returns the FacilityName field if non-nil, zero value otherwise.

### GetFacilityNameOk

`func (o *PortTerminalAddOnImportPlan) GetFacilityNameOk() (*string, bool)`

GetFacilityNameOk returns a tuple with the FacilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityName

`func (o *PortTerminalAddOnImportPlan) SetFacilityName(v string)`

SetFacilityName sets FacilityName field to given value.

### HasFacilityName

`func (o *PortTerminalAddOnImportPlan) HasFacilityName() bool`

HasFacilityName returns a boolean if a field has been set.

### GetVesselName

`func (o *PortTerminalAddOnImportPlan) GetVesselName() string`

GetVesselName returns the VesselName field if non-nil, zero value otherwise.

### GetVesselNameOk

`func (o *PortTerminalAddOnImportPlan) GetVesselNameOk() (*string, bool)`

GetVesselNameOk returns a tuple with the VesselName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselName

`func (o *PortTerminalAddOnImportPlan) SetVesselName(v string)`

SetVesselName sets VesselName field to given value.

### HasVesselName

`func (o *PortTerminalAddOnImportPlan) HasVesselName() bool`

HasVesselName returns a boolean if a field has been set.

### GetVoyageNo

`func (o *PortTerminalAddOnImportPlan) GetVoyageNo() string`

GetVoyageNo returns the VoyageNo field if non-nil, zero value otherwise.

### GetVoyageNoOk

`func (o *PortTerminalAddOnImportPlan) GetVoyageNoOk() (*string, bool)`

GetVoyageNoOk returns a tuple with the VoyageNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageNo

`func (o *PortTerminalAddOnImportPlan) SetVoyageNo(v string)`

SetVoyageNo sets VoyageNo field to given value.

### HasVoyageNo

`func (o *PortTerminalAddOnImportPlan) HasVoyageNo() bool`

HasVoyageNo returns a boolean if a field has been set.

### GetContainerStatus

`func (o *PortTerminalAddOnImportPlan) GetContainerStatus() string`

GetContainerStatus returns the ContainerStatus field if non-nil, zero value otherwise.

### GetContainerStatusOk

`func (o *PortTerminalAddOnImportPlan) GetContainerStatusOk() (*string, bool)`

GetContainerStatusOk returns a tuple with the ContainerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerStatus

`func (o *PortTerminalAddOnImportPlan) SetContainerStatus(v string)`

SetContainerStatus sets ContainerStatus field to given value.

### HasContainerStatus

`func (o *PortTerminalAddOnImportPlan) HasContainerStatus() bool`

HasContainerStatus returns a boolean if a field has been set.

### GetDischargeDate

`func (o *PortTerminalAddOnImportPlan) GetDischargeDate() string`

GetDischargeDate returns the DischargeDate field if non-nil, zero value otherwise.

### GetDischargeDateOk

`func (o *PortTerminalAddOnImportPlan) GetDischargeDateOk() (*string, bool)`

GetDischargeDateOk returns a tuple with the DischargeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDischargeDate

`func (o *PortTerminalAddOnImportPlan) SetDischargeDate(v string)`

SetDischargeDate sets DischargeDate field to given value.

### HasDischargeDate

`func (o *PortTerminalAddOnImportPlan) HasDischargeDate() bool`

HasDischargeDate returns a boolean if a field has been set.

### GetYardLocation

`func (o *PortTerminalAddOnImportPlan) GetYardLocation() string`

GetYardLocation returns the YardLocation field if non-nil, zero value otherwise.

### GetYardLocationOk

`func (o *PortTerminalAddOnImportPlan) GetYardLocationOk() (*string, bool)`

GetYardLocationOk returns a tuple with the YardLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYardLocation

`func (o *PortTerminalAddOnImportPlan) SetYardLocation(v string)`

SetYardLocation sets YardLocation field to given value.

### HasYardLocation

`func (o *PortTerminalAddOnImportPlan) HasYardLocation() bool`

HasYardLocation returns a boolean if a field has been set.

### GetContainerHolds

`func (o *PortTerminalAddOnImportPlan) GetContainerHolds() string`

GetContainerHolds returns the ContainerHolds field if non-nil, zero value otherwise.

### GetContainerHoldsOk

`func (o *PortTerminalAddOnImportPlan) GetContainerHoldsOk() (*string, bool)`

GetContainerHoldsOk returns a tuple with the ContainerHolds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerHolds

`func (o *PortTerminalAddOnImportPlan) SetContainerHolds(v string)`

SetContainerHolds sets ContainerHolds field to given value.

### HasContainerHolds

`func (o *PortTerminalAddOnImportPlan) HasContainerHolds() bool`

HasContainerHolds returns a boolean if a field has been set.

### GetLastFreeDay

`func (o *PortTerminalAddOnImportPlan) GetLastFreeDay() string`

GetLastFreeDay returns the LastFreeDay field if non-nil, zero value otherwise.

### GetLastFreeDayOk

`func (o *PortTerminalAddOnImportPlan) GetLastFreeDayOk() (*string, bool)`

GetLastFreeDayOk returns a tuple with the LastFreeDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFreeDay

`func (o *PortTerminalAddOnImportPlan) SetLastFreeDay(v string)`

SetLastFreeDay sets LastFreeDay field to given value.

### HasLastFreeDay

`func (o *PortTerminalAddOnImportPlan) HasLastFreeDay() bool`

HasLastFreeDay returns a boolean if a field has been set.

### GetDemurrageOwed

`func (o *PortTerminalAddOnImportPlan) GetDemurrageOwed() string`

GetDemurrageOwed returns the DemurrageOwed field if non-nil, zero value otherwise.

### GetDemurrageOwedOk

`func (o *PortTerminalAddOnImportPlan) GetDemurrageOwedOk() (*string, bool)`

GetDemurrageOwedOk returns a tuple with the DemurrageOwed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemurrageOwed

`func (o *PortTerminalAddOnImportPlan) SetDemurrageOwed(v string)`

SetDemurrageOwed sets DemurrageOwed field to given value.

### HasDemurrageOwed

`func (o *PortTerminalAddOnImportPlan) HasDemurrageOwed() bool`

HasDemurrageOwed returns a boolean if a field has been set.

### GetGateOutDate

`func (o *PortTerminalAddOnImportPlan) GetGateOutDate() string`

GetGateOutDate returns the GateOutDate field if non-nil, zero value otherwise.

### GetGateOutDateOk

`func (o *PortTerminalAddOnImportPlan) GetGateOutDateOk() (*string, bool)`

GetGateOutDateOk returns a tuple with the GateOutDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateOutDate

`func (o *PortTerminalAddOnImportPlan) SetGateOutDate(v string)`

SetGateOutDate sets GateOutDate field to given value.

### HasGateOutDate

`func (o *PortTerminalAddOnImportPlan) HasGateOutDate() bool`

HasGateOutDate returns a boolean if a field has been set.

### GetLatestEta

`func (o *PortTerminalAddOnImportPlan) GetLatestEta() string`

GetLatestEta returns the LatestEta field if non-nil, zero value otherwise.

### GetLatestEtaOk

`func (o *PortTerminalAddOnImportPlan) GetLatestEtaOk() (*string, bool)`

GetLatestEtaOk returns a tuple with the LatestEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestEta

`func (o *PortTerminalAddOnImportPlan) SetLatestEta(v string)`

SetLatestEta sets LatestEta field to given value.

### HasLatestEta

`func (o *PortTerminalAddOnImportPlan) HasLatestEta() bool`

HasLatestEta returns a boolean if a field has been set.

### GetReadyForDelivery

`func (o *PortTerminalAddOnImportPlan) GetReadyForDelivery() string`

GetReadyForDelivery returns the ReadyForDelivery field if non-nil, zero value otherwise.

### GetReadyForDeliveryOk

`func (o *PortTerminalAddOnImportPlan) GetReadyForDeliveryOk() (*string, bool)`

GetReadyForDeliveryOk returns a tuple with the ReadyForDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyForDelivery

`func (o *PortTerminalAddOnImportPlan) SetReadyForDelivery(v string)`

SetReadyForDelivery sets ReadyForDelivery field to given value.

### HasReadyForDelivery

`func (o *PortTerminalAddOnImportPlan) HasReadyForDelivery() bool`

HasReadyForDelivery returns a boolean if a field has been set.

### GetAppointmentSet

`func (o *PortTerminalAddOnImportPlan) GetAppointmentSet() string`

GetAppointmentSet returns the AppointmentSet field if non-nil, zero value otherwise.

### GetAppointmentSetOk

`func (o *PortTerminalAddOnImportPlan) GetAppointmentSetOk() (*string, bool)`

GetAppointmentSetOk returns a tuple with the AppointmentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppointmentSet

`func (o *PortTerminalAddOnImportPlan) SetAppointmentSet(v string)`

SetAppointmentSet sets AppointmentSet field to given value.

### HasAppointmentSet

`func (o *PortTerminalAddOnImportPlan) HasAppointmentSet() bool`

HasAppointmentSet returns a boolean if a field has been set.

### GetPickupAppointmentDate

`func (o *PortTerminalAddOnImportPlan) GetPickupAppointmentDate() time.Time`

GetPickupAppointmentDate returns the PickupAppointmentDate field if non-nil, zero value otherwise.

### GetPickupAppointmentDateOk

`func (o *PortTerminalAddOnImportPlan) GetPickupAppointmentDateOk() (*time.Time, bool)`

GetPickupAppointmentDateOk returns a tuple with the PickupAppointmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAppointmentDate

`func (o *PortTerminalAddOnImportPlan) SetPickupAppointmentDate(v time.Time)`

SetPickupAppointmentDate sets PickupAppointmentDate field to given value.

### HasPickupAppointmentDate

`func (o *PortTerminalAddOnImportPlan) HasPickupAppointmentDate() bool`

HasPickupAppointmentDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


