# PortTerminalAddOnExportPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique Identifier for the Export Plan Object. | [optional] 
**Updated** | Pointer to **time.Time** | Export Plan Object update date | [optional] 
**PortCode** | Pointer to **string** | Relates to the UNLOCODE for the export port, as standardized by Portcast | [optional] 
**PortName** | Pointer to **string** | Relates to the export port name, as standardized by Portcast | [optional] 
**FacilityCode** | Pointer to **string** | BIC/SMDG Defined Code for the Port Terminal Location | [optional] 
**FacilityName** | Pointer to **string** | Export port terminal name | [optional] 
**VesselName** | Pointer to **string** | Vessel name of the vessel associated with the export location. | [optional] 
**VoyageNo** | Pointer to **string** | Voyage Number associated with the export journey, as reported by the Terminal. | [optional] 
**ErdStandard** | Pointer to **time.Time** | Earliest Receipt Date (ERD) for a standard container, as reported by the terminal - Local Time. | [optional] 
**ErdReefer** | Pointer to **time.Time** | Earliest Receipt Date (ERD) for a reefer container, as reported by the terminal - Local Time. | [optional] 
**GateInDate** | Pointer to **string** | Date a container was received at the terminal through the gate, as reported by the terminal - Local Time. | [optional] 
**PortCutoffStandard** | Pointer to **time.Time** | Last possible date and time for gating-in standard container at the export terminal, as reported by the Terminal - Local Time. | [optional] 
**PortCutoffReefer** | Pointer to **time.Time** | Last possible date and time for gating-in reefer container at the export terminal, as reported by the terminal - Local Time. | [optional] 
**LatestEta** | Pointer to **time.Time** | Latest vessel ETA to the export port, as reported by the terminal - Local Time. | [optional] 
**ActualArrival** | Pointer to **time.Time** | Actual vessel time of arrival, as reported by the terminal - Local Time. | [optional] 
**LatestEtd** | Pointer to **time.Time** | Latest vessel ETD from the export port, as reported by the terminal - Local Time. | [optional] 
**ActualDeparture** | Pointer to **time.Time** | Actual vessel time of departure from the export port, as reported by the terminal - Local Time. | [optional] 

## Methods

### NewPortTerminalAddOnExportPlan

`func NewPortTerminalAddOnExportPlan() *PortTerminalAddOnExportPlan`

NewPortTerminalAddOnExportPlan instantiates a new PortTerminalAddOnExportPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortTerminalAddOnExportPlanWithDefaults

`func NewPortTerminalAddOnExportPlanWithDefaults() *PortTerminalAddOnExportPlan`

NewPortTerminalAddOnExportPlanWithDefaults instantiates a new PortTerminalAddOnExportPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortTerminalAddOnExportPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortTerminalAddOnExportPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortTerminalAddOnExportPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortTerminalAddOnExportPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdated

`func (o *PortTerminalAddOnExportPlan) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PortTerminalAddOnExportPlan) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PortTerminalAddOnExportPlan) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PortTerminalAddOnExportPlan) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetPortCode

`func (o *PortTerminalAddOnExportPlan) GetPortCode() string`

GetPortCode returns the PortCode field if non-nil, zero value otherwise.

### GetPortCodeOk

`func (o *PortTerminalAddOnExportPlan) GetPortCodeOk() (*string, bool)`

GetPortCodeOk returns a tuple with the PortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCode

`func (o *PortTerminalAddOnExportPlan) SetPortCode(v string)`

SetPortCode sets PortCode field to given value.

### HasPortCode

`func (o *PortTerminalAddOnExportPlan) HasPortCode() bool`

HasPortCode returns a boolean if a field has been set.

### GetPortName

`func (o *PortTerminalAddOnExportPlan) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *PortTerminalAddOnExportPlan) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *PortTerminalAddOnExportPlan) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *PortTerminalAddOnExportPlan) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetFacilityCode

`func (o *PortTerminalAddOnExportPlan) GetFacilityCode() string`

GetFacilityCode returns the FacilityCode field if non-nil, zero value otherwise.

### GetFacilityCodeOk

`func (o *PortTerminalAddOnExportPlan) GetFacilityCodeOk() (*string, bool)`

GetFacilityCodeOk returns a tuple with the FacilityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCode

`func (o *PortTerminalAddOnExportPlan) SetFacilityCode(v string)`

SetFacilityCode sets FacilityCode field to given value.

### HasFacilityCode

`func (o *PortTerminalAddOnExportPlan) HasFacilityCode() bool`

HasFacilityCode returns a boolean if a field has been set.

### GetFacilityName

`func (o *PortTerminalAddOnExportPlan) GetFacilityName() string`

GetFacilityName returns the FacilityName field if non-nil, zero value otherwise.

### GetFacilityNameOk

`func (o *PortTerminalAddOnExportPlan) GetFacilityNameOk() (*string, bool)`

GetFacilityNameOk returns a tuple with the FacilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityName

`func (o *PortTerminalAddOnExportPlan) SetFacilityName(v string)`

SetFacilityName sets FacilityName field to given value.

### HasFacilityName

`func (o *PortTerminalAddOnExportPlan) HasFacilityName() bool`

HasFacilityName returns a boolean if a field has been set.

### GetVesselName

`func (o *PortTerminalAddOnExportPlan) GetVesselName() string`

GetVesselName returns the VesselName field if non-nil, zero value otherwise.

### GetVesselNameOk

`func (o *PortTerminalAddOnExportPlan) GetVesselNameOk() (*string, bool)`

GetVesselNameOk returns a tuple with the VesselName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselName

`func (o *PortTerminalAddOnExportPlan) SetVesselName(v string)`

SetVesselName sets VesselName field to given value.

### HasVesselName

`func (o *PortTerminalAddOnExportPlan) HasVesselName() bool`

HasVesselName returns a boolean if a field has been set.

### GetVoyageNo

`func (o *PortTerminalAddOnExportPlan) GetVoyageNo() string`

GetVoyageNo returns the VoyageNo field if non-nil, zero value otherwise.

### GetVoyageNoOk

`func (o *PortTerminalAddOnExportPlan) GetVoyageNoOk() (*string, bool)`

GetVoyageNoOk returns a tuple with the VoyageNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageNo

`func (o *PortTerminalAddOnExportPlan) SetVoyageNo(v string)`

SetVoyageNo sets VoyageNo field to given value.

### HasVoyageNo

`func (o *PortTerminalAddOnExportPlan) HasVoyageNo() bool`

HasVoyageNo returns a boolean if a field has been set.

### GetErdStandard

`func (o *PortTerminalAddOnExportPlan) GetErdStandard() time.Time`

GetErdStandard returns the ErdStandard field if non-nil, zero value otherwise.

### GetErdStandardOk

`func (o *PortTerminalAddOnExportPlan) GetErdStandardOk() (*time.Time, bool)`

GetErdStandardOk returns a tuple with the ErdStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErdStandard

`func (o *PortTerminalAddOnExportPlan) SetErdStandard(v time.Time)`

SetErdStandard sets ErdStandard field to given value.

### HasErdStandard

`func (o *PortTerminalAddOnExportPlan) HasErdStandard() bool`

HasErdStandard returns a boolean if a field has been set.

### GetErdReefer

`func (o *PortTerminalAddOnExportPlan) GetErdReefer() time.Time`

GetErdReefer returns the ErdReefer field if non-nil, zero value otherwise.

### GetErdReeferOk

`func (o *PortTerminalAddOnExportPlan) GetErdReeferOk() (*time.Time, bool)`

GetErdReeferOk returns a tuple with the ErdReefer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErdReefer

`func (o *PortTerminalAddOnExportPlan) SetErdReefer(v time.Time)`

SetErdReefer sets ErdReefer field to given value.

### HasErdReefer

`func (o *PortTerminalAddOnExportPlan) HasErdReefer() bool`

HasErdReefer returns a boolean if a field has been set.

### GetGateInDate

`func (o *PortTerminalAddOnExportPlan) GetGateInDate() string`

GetGateInDate returns the GateInDate field if non-nil, zero value otherwise.

### GetGateInDateOk

`func (o *PortTerminalAddOnExportPlan) GetGateInDateOk() (*string, bool)`

GetGateInDateOk returns a tuple with the GateInDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateInDate

`func (o *PortTerminalAddOnExportPlan) SetGateInDate(v string)`

SetGateInDate sets GateInDate field to given value.

### HasGateInDate

`func (o *PortTerminalAddOnExportPlan) HasGateInDate() bool`

HasGateInDate returns a boolean if a field has been set.

### GetPortCutoffStandard

`func (o *PortTerminalAddOnExportPlan) GetPortCutoffStandard() time.Time`

GetPortCutoffStandard returns the PortCutoffStandard field if non-nil, zero value otherwise.

### GetPortCutoffStandardOk

`func (o *PortTerminalAddOnExportPlan) GetPortCutoffStandardOk() (*time.Time, bool)`

GetPortCutoffStandardOk returns a tuple with the PortCutoffStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCutoffStandard

`func (o *PortTerminalAddOnExportPlan) SetPortCutoffStandard(v time.Time)`

SetPortCutoffStandard sets PortCutoffStandard field to given value.

### HasPortCutoffStandard

`func (o *PortTerminalAddOnExportPlan) HasPortCutoffStandard() bool`

HasPortCutoffStandard returns a boolean if a field has been set.

### GetPortCutoffReefer

`func (o *PortTerminalAddOnExportPlan) GetPortCutoffReefer() time.Time`

GetPortCutoffReefer returns the PortCutoffReefer field if non-nil, zero value otherwise.

### GetPortCutoffReeferOk

`func (o *PortTerminalAddOnExportPlan) GetPortCutoffReeferOk() (*time.Time, bool)`

GetPortCutoffReeferOk returns a tuple with the PortCutoffReefer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCutoffReefer

`func (o *PortTerminalAddOnExportPlan) SetPortCutoffReefer(v time.Time)`

SetPortCutoffReefer sets PortCutoffReefer field to given value.

### HasPortCutoffReefer

`func (o *PortTerminalAddOnExportPlan) HasPortCutoffReefer() bool`

HasPortCutoffReefer returns a boolean if a field has been set.

### GetLatestEta

`func (o *PortTerminalAddOnExportPlan) GetLatestEta() time.Time`

GetLatestEta returns the LatestEta field if non-nil, zero value otherwise.

### GetLatestEtaOk

`func (o *PortTerminalAddOnExportPlan) GetLatestEtaOk() (*time.Time, bool)`

GetLatestEtaOk returns a tuple with the LatestEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestEta

`func (o *PortTerminalAddOnExportPlan) SetLatestEta(v time.Time)`

SetLatestEta sets LatestEta field to given value.

### HasLatestEta

`func (o *PortTerminalAddOnExportPlan) HasLatestEta() bool`

HasLatestEta returns a boolean if a field has been set.

### GetActualArrival

`func (o *PortTerminalAddOnExportPlan) GetActualArrival() time.Time`

GetActualArrival returns the ActualArrival field if non-nil, zero value otherwise.

### GetActualArrivalOk

`func (o *PortTerminalAddOnExportPlan) GetActualArrivalOk() (*time.Time, bool)`

GetActualArrivalOk returns a tuple with the ActualArrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualArrival

`func (o *PortTerminalAddOnExportPlan) SetActualArrival(v time.Time)`

SetActualArrival sets ActualArrival field to given value.

### HasActualArrival

`func (o *PortTerminalAddOnExportPlan) HasActualArrival() bool`

HasActualArrival returns a boolean if a field has been set.

### GetLatestEtd

`func (o *PortTerminalAddOnExportPlan) GetLatestEtd() time.Time`

GetLatestEtd returns the LatestEtd field if non-nil, zero value otherwise.

### GetLatestEtdOk

`func (o *PortTerminalAddOnExportPlan) GetLatestEtdOk() (*time.Time, bool)`

GetLatestEtdOk returns a tuple with the LatestEtd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestEtd

`func (o *PortTerminalAddOnExportPlan) SetLatestEtd(v time.Time)`

SetLatestEtd sets LatestEtd field to given value.

### HasLatestEtd

`func (o *PortTerminalAddOnExportPlan) HasLatestEtd() bool`

HasLatestEtd returns a boolean if a field has been set.

### GetActualDeparture

`func (o *PortTerminalAddOnExportPlan) GetActualDeparture() time.Time`

GetActualDeparture returns the ActualDeparture field if non-nil, zero value otherwise.

### GetActualDepartureOk

`func (o *PortTerminalAddOnExportPlan) GetActualDepartureOk() (*time.Time, bool)`

GetActualDepartureOk returns a tuple with the ActualDeparture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDeparture

`func (o *PortTerminalAddOnExportPlan) SetActualDeparture(v time.Time)`

SetActualDeparture sets ActualDeparture field to given value.

### HasActualDeparture

`func (o *PortTerminalAddOnExportPlan) HasActualDeparture() bool`

HasActualDeparture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


