# SailingInfoTrackingSailingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierNo** | Pointer to **string** | Carrier SCAC Code | [optional] 
**Created** | Pointer to **NullableTime** | Sailing Info Creation Date | [optional] 
**Id** | Pointer to **string** | Unique Identifier for the Sailing Info | [optional] 
**Imo** | Pointer to **NullableString** | Vessel IMO for the Vessel associated with this specific journey leg | [optional] 
**IsActive** | Pointer to **bool** | Defines if the sailing info is active/latest or not | [optional] [default to true]
**Pod** | Pointer to **NullableString** | Target Location UNLOCODE for this specific leg of the journey | [optional] 
**PodActualArrivalLt** | Pointer to **NullableTime** | Actual Arrival Time at the target port location - Local Time | [optional] 
**PodActualArrivalLtFromAis** | Pointer to **NullableTime** | Actual Arrival Time at the target port location as per AIS Data - Local Time | [optional] 
**PodActualDepartureLtFromAis** | Pointer to **NullableTime** | Actual Departure Time from the target port location as per AIS Data - Local Time | [optional] 
**PodActualDischargeLt** | Pointer to **NullableTime** | Actual Time of Discharge at the target port location - Local Time | [optional] 
**PodName** | Pointer to **NullableString** | Target Location Name for this specific leg of the journey | [optional] 
**PodPredictedArrivalLt** | Pointer to **NullableTime** | Portcast Predicted Time of Arrival at the target port location - Local Time | [optional] 
**PodPredictedDepartureLt** | Pointer to **NullableTime** | Portcast Predicted Time of Departure from the target port location - Local Time | [optional] 
**PodScheduledArrivalLt** | Pointer to **NullableTime** | Scheduled Time of Arrival at the target port location - Local Time | [optional] 
**PodScheduledArrivalLtFromSchedule** | Pointer to **NullableTime** | Scheduled Time of Arrival at the target port location as per Vessel Schedule - Local Time | [optional] 
**PodScheduledDepartureLtFromSchedule** | Pointer to **NullableTime** | Scheduled Time of Departure from the target port location as per Vessel Schedule - Local Time | [optional] 
**PodScheduledDischargeLt** | Pointer to **NullableTime** | Scheduled Discharge Time at the target port location - Local Time | [optional] 
**PodTerminalName** | Pointer to **NullableString** | Terminal Name for the Target Port for this specific leg of the journey | [optional] 
**Pol** | Pointer to **NullableString** | Starting Location UNLOCODE for this specific leg of the journey | [optional] [default to "USHOU"]
**PolActualArrivalLtFromAis** | Pointer to **NullableTime** | Actual Time of Arrival at the starting port location as per AIS Data - Local Time | [optional] 
**PolActualDepartureLt** | Pointer to **NullableTime** | Actual Time of Departure from the starting port location - Local Time | [optional] 
**PolActualDepartureLtFromAis** | Pointer to **NullableTime** | Actual Time of Departure from the starting port location as per AIS Data - Local Time | [optional] 
**PolActualLoadingLt** | Pointer to **NullableTime** | Actual Loading Time at the starting port location - Local Time | [optional] 
**PolName** | Pointer to **NullableString** | Starting Location Name for this specific leg of the journey | [optional] [default to "HOUSTON"]
**PolPredictedArrivalLt** | Pointer to **NullableTime** | Portcast Predicted Time of Arrival at the starting port location - Local Time | [optional] 
**PolPredictedDepartureLt** | Pointer to **NullableTime** | Portcast Predicted Time of Departure from the starting port location - Local Time | [optional] 
**PolScheduledArrivalLtFromSchedule** | Pointer to **NullableTime** | Scheduled Time of Arrival at the starting port location as per Vessel Schedule- Local Time | [optional] 
**PolScheduledDepartureLt** | Pointer to **NullableTime** | Scheduled Time of Departure from the starting port location - Local Time | [optional] 
**PolScheduledDepartureLtFromSchedule** | Pointer to **NullableTime** | Scheduled Time of Departure from the starting port location as per Vessel Schedule - Local Time | [optional] 
**PolScheduledLoadingLt** | Pointer to **NullableTime** | Scheduled Loading Time at the starting port location - Local Time | [optional] 
**PolTerminalName** | Pointer to **NullableString** | Terminal Name for the Port of Loading (POL) | [optional] 
**Updated** | Pointer to **NullableTime** | Sailing Info Updated Date | [optional] 
**VesselFlag** | Pointer to **NullableString** | Vessel Name for the Vessel Flag for the vessel associated with this specific journey leg | [optional] 
**VesselLeg** | Pointer to **int32** | Defines which leg of the journey this sailing info is for | [optional] 
**VesselName** | Pointer to **NullableString** | Vessel Name for the Vessel associated with this specific journey leg | [optional] 
**VoyageNo** | Pointer to **NullableString** | Voyage Number associated with this specific journey leg (From container track and trace data) [Most reliable]  | [optional] 

## Methods

### NewSailingInfoTrackingSailingInfo

`func NewSailingInfoTrackingSailingInfo() *SailingInfoTrackingSailingInfo`

NewSailingInfoTrackingSailingInfo instantiates a new SailingInfoTrackingSailingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSailingInfoTrackingSailingInfoWithDefaults

`func NewSailingInfoTrackingSailingInfoWithDefaults() *SailingInfoTrackingSailingInfo`

NewSailingInfoTrackingSailingInfoWithDefaults instantiates a new SailingInfoTrackingSailingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierNo

`func (o *SailingInfoTrackingSailingInfo) GetCarrierNo() string`

GetCarrierNo returns the CarrierNo field if non-nil, zero value otherwise.

### GetCarrierNoOk

`func (o *SailingInfoTrackingSailingInfo) GetCarrierNoOk() (*string, bool)`

GetCarrierNoOk returns a tuple with the CarrierNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierNo

`func (o *SailingInfoTrackingSailingInfo) SetCarrierNo(v string)`

SetCarrierNo sets CarrierNo field to given value.

### HasCarrierNo

`func (o *SailingInfoTrackingSailingInfo) HasCarrierNo() bool`

HasCarrierNo returns a boolean if a field has been set.

### GetCreated

`func (o *SailingInfoTrackingSailingInfo) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SailingInfoTrackingSailingInfo) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SailingInfoTrackingSailingInfo) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SailingInfoTrackingSailingInfo) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *SailingInfoTrackingSailingInfo) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *SailingInfoTrackingSailingInfo) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetId

`func (o *SailingInfoTrackingSailingInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SailingInfoTrackingSailingInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SailingInfoTrackingSailingInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SailingInfoTrackingSailingInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImo

`func (o *SailingInfoTrackingSailingInfo) GetImo() string`

GetImo returns the Imo field if non-nil, zero value otherwise.

### GetImoOk

`func (o *SailingInfoTrackingSailingInfo) GetImoOk() (*string, bool)`

GetImoOk returns a tuple with the Imo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImo

`func (o *SailingInfoTrackingSailingInfo) SetImo(v string)`

SetImo sets Imo field to given value.

### HasImo

`func (o *SailingInfoTrackingSailingInfo) HasImo() bool`

HasImo returns a boolean if a field has been set.

### SetImoNil

`func (o *SailingInfoTrackingSailingInfo) SetImoNil(b bool)`

 SetImoNil sets the value for Imo to be an explicit nil

### UnsetImo
`func (o *SailingInfoTrackingSailingInfo) UnsetImo()`

UnsetImo ensures that no value is present for Imo, not even an explicit nil
### GetIsActive

`func (o *SailingInfoTrackingSailingInfo) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SailingInfoTrackingSailingInfo) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SailingInfoTrackingSailingInfo) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *SailingInfoTrackingSailingInfo) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetPod

`func (o *SailingInfoTrackingSailingInfo) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *SailingInfoTrackingSailingInfo) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *SailingInfoTrackingSailingInfo) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *SailingInfoTrackingSailingInfo) HasPod() bool`

HasPod returns a boolean if a field has been set.

### SetPodNil

`func (o *SailingInfoTrackingSailingInfo) SetPodNil(b bool)`

 SetPodNil sets the value for Pod to be an explicit nil

### UnsetPod
`func (o *SailingInfoTrackingSailingInfo) UnsetPod()`

UnsetPod ensures that no value is present for Pod, not even an explicit nil
### GetPodActualArrivalLt

`func (o *SailingInfoTrackingSailingInfo) GetPodActualArrivalLt() time.Time`

GetPodActualArrivalLt returns the PodActualArrivalLt field if non-nil, zero value otherwise.

### GetPodActualArrivalLtOk

`func (o *SailingInfoTrackingSailingInfo) GetPodActualArrivalLtOk() (*time.Time, bool)`

GetPodActualArrivalLtOk returns a tuple with the PodActualArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualArrivalLt

`func (o *SailingInfoTrackingSailingInfo) SetPodActualArrivalLt(v time.Time)`

SetPodActualArrivalLt sets PodActualArrivalLt field to given value.

### HasPodActualArrivalLt

`func (o *SailingInfoTrackingSailingInfo) HasPodActualArrivalLt() bool`

HasPodActualArrivalLt returns a boolean if a field has been set.

### SetPodActualArrivalLtNil

`func (o *SailingInfoTrackingSailingInfo) SetPodActualArrivalLtNil(b bool)`

 SetPodActualArrivalLtNil sets the value for PodActualArrivalLt to be an explicit nil

### UnsetPodActualArrivalLt
`func (o *SailingInfoTrackingSailingInfo) UnsetPodActualArrivalLt()`

UnsetPodActualArrivalLt ensures that no value is present for PodActualArrivalLt, not even an explicit nil
### GetPodActualArrivalLtFromAis

`func (o *SailingInfoTrackingSailingInfo) GetPodActualArrivalLtFromAis() time.Time`

GetPodActualArrivalLtFromAis returns the PodActualArrivalLtFromAis field if non-nil, zero value otherwise.

### GetPodActualArrivalLtFromAisOk

`func (o *SailingInfoTrackingSailingInfo) GetPodActualArrivalLtFromAisOk() (*time.Time, bool)`

GetPodActualArrivalLtFromAisOk returns a tuple with the PodActualArrivalLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualArrivalLtFromAis

`func (o *SailingInfoTrackingSailingInfo) SetPodActualArrivalLtFromAis(v time.Time)`

SetPodActualArrivalLtFromAis sets PodActualArrivalLtFromAis field to given value.

### HasPodActualArrivalLtFromAis

`func (o *SailingInfoTrackingSailingInfo) HasPodActualArrivalLtFromAis() bool`

HasPodActualArrivalLtFromAis returns a boolean if a field has been set.

### SetPodActualArrivalLtFromAisNil

`func (o *SailingInfoTrackingSailingInfo) SetPodActualArrivalLtFromAisNil(b bool)`

 SetPodActualArrivalLtFromAisNil sets the value for PodActualArrivalLtFromAis to be an explicit nil

### UnsetPodActualArrivalLtFromAis
`func (o *SailingInfoTrackingSailingInfo) UnsetPodActualArrivalLtFromAis()`

UnsetPodActualArrivalLtFromAis ensures that no value is present for PodActualArrivalLtFromAis, not even an explicit nil
### GetPodActualDepartureLtFromAis

`func (o *SailingInfoTrackingSailingInfo) GetPodActualDepartureLtFromAis() time.Time`

GetPodActualDepartureLtFromAis returns the PodActualDepartureLtFromAis field if non-nil, zero value otherwise.

### GetPodActualDepartureLtFromAisOk

`func (o *SailingInfoTrackingSailingInfo) GetPodActualDepartureLtFromAisOk() (*time.Time, bool)`

GetPodActualDepartureLtFromAisOk returns a tuple with the PodActualDepartureLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualDepartureLtFromAis

`func (o *SailingInfoTrackingSailingInfo) SetPodActualDepartureLtFromAis(v time.Time)`

SetPodActualDepartureLtFromAis sets PodActualDepartureLtFromAis field to given value.

### HasPodActualDepartureLtFromAis

`func (o *SailingInfoTrackingSailingInfo) HasPodActualDepartureLtFromAis() bool`

HasPodActualDepartureLtFromAis returns a boolean if a field has been set.

### SetPodActualDepartureLtFromAisNil

`func (o *SailingInfoTrackingSailingInfo) SetPodActualDepartureLtFromAisNil(b bool)`

 SetPodActualDepartureLtFromAisNil sets the value for PodActualDepartureLtFromAis to be an explicit nil

### UnsetPodActualDepartureLtFromAis
`func (o *SailingInfoTrackingSailingInfo) UnsetPodActualDepartureLtFromAis()`

UnsetPodActualDepartureLtFromAis ensures that no value is present for PodActualDepartureLtFromAis, not even an explicit nil
### GetPodActualDischargeLt

`func (o *SailingInfoTrackingSailingInfo) GetPodActualDischargeLt() time.Time`

GetPodActualDischargeLt returns the PodActualDischargeLt field if non-nil, zero value otherwise.

### GetPodActualDischargeLtOk

`func (o *SailingInfoTrackingSailingInfo) GetPodActualDischargeLtOk() (*time.Time, bool)`

GetPodActualDischargeLtOk returns a tuple with the PodActualDischargeLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualDischargeLt

`func (o *SailingInfoTrackingSailingInfo) SetPodActualDischargeLt(v time.Time)`

SetPodActualDischargeLt sets PodActualDischargeLt field to given value.

### HasPodActualDischargeLt

`func (o *SailingInfoTrackingSailingInfo) HasPodActualDischargeLt() bool`

HasPodActualDischargeLt returns a boolean if a field has been set.

### SetPodActualDischargeLtNil

`func (o *SailingInfoTrackingSailingInfo) SetPodActualDischargeLtNil(b bool)`

 SetPodActualDischargeLtNil sets the value for PodActualDischargeLt to be an explicit nil

### UnsetPodActualDischargeLt
`func (o *SailingInfoTrackingSailingInfo) UnsetPodActualDischargeLt()`

UnsetPodActualDischargeLt ensures that no value is present for PodActualDischargeLt, not even an explicit nil
### GetPodName

`func (o *SailingInfoTrackingSailingInfo) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *SailingInfoTrackingSailingInfo) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *SailingInfoTrackingSailingInfo) SetPodName(v string)`

SetPodName sets PodName field to given value.

### HasPodName

`func (o *SailingInfoTrackingSailingInfo) HasPodName() bool`

HasPodName returns a boolean if a field has been set.

### SetPodNameNil

`func (o *SailingInfoTrackingSailingInfo) SetPodNameNil(b bool)`

 SetPodNameNil sets the value for PodName to be an explicit nil

### UnsetPodName
`func (o *SailingInfoTrackingSailingInfo) UnsetPodName()`

UnsetPodName ensures that no value is present for PodName, not even an explicit nil
### GetPodPredictedArrivalLt

`func (o *SailingInfoTrackingSailingInfo) GetPodPredictedArrivalLt() time.Time`

GetPodPredictedArrivalLt returns the PodPredictedArrivalLt field if non-nil, zero value otherwise.

### GetPodPredictedArrivalLtOk

`func (o *SailingInfoTrackingSailingInfo) GetPodPredictedArrivalLtOk() (*time.Time, bool)`

GetPodPredictedArrivalLtOk returns a tuple with the PodPredictedArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPredictedArrivalLt

`func (o *SailingInfoTrackingSailingInfo) SetPodPredictedArrivalLt(v time.Time)`

SetPodPredictedArrivalLt sets PodPredictedArrivalLt field to given value.

### HasPodPredictedArrivalLt

`func (o *SailingInfoTrackingSailingInfo) HasPodPredictedArrivalLt() bool`

HasPodPredictedArrivalLt returns a boolean if a field has been set.

### SetPodPredictedArrivalLtNil

`func (o *SailingInfoTrackingSailingInfo) SetPodPredictedArrivalLtNil(b bool)`

 SetPodPredictedArrivalLtNil sets the value for PodPredictedArrivalLt to be an explicit nil

### UnsetPodPredictedArrivalLt
`func (o *SailingInfoTrackingSailingInfo) UnsetPodPredictedArrivalLt()`

UnsetPodPredictedArrivalLt ensures that no value is present for PodPredictedArrivalLt, not even an explicit nil
### GetPodPredictedDepartureLt

`func (o *SailingInfoTrackingSailingInfo) GetPodPredictedDepartureLt() time.Time`

GetPodPredictedDepartureLt returns the PodPredictedDepartureLt field if non-nil, zero value otherwise.

### GetPodPredictedDepartureLtOk

`func (o *SailingInfoTrackingSailingInfo) GetPodPredictedDepartureLtOk() (*time.Time, bool)`

GetPodPredictedDepartureLtOk returns a tuple with the PodPredictedDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPredictedDepartureLt

`func (o *SailingInfoTrackingSailingInfo) SetPodPredictedDepartureLt(v time.Time)`

SetPodPredictedDepartureLt sets PodPredictedDepartureLt field to given value.

### HasPodPredictedDepartureLt

`func (o *SailingInfoTrackingSailingInfo) HasPodPredictedDepartureLt() bool`

HasPodPredictedDepartureLt returns a boolean if a field has been set.

### SetPodPredictedDepartureLtNil

`func (o *SailingInfoTrackingSailingInfo) SetPodPredictedDepartureLtNil(b bool)`

 SetPodPredictedDepartureLtNil sets the value for PodPredictedDepartureLt to be an explicit nil

### UnsetPodPredictedDepartureLt
`func (o *SailingInfoTrackingSailingInfo) UnsetPodPredictedDepartureLt()`

UnsetPodPredictedDepartureLt ensures that no value is present for PodPredictedDepartureLt, not even an explicit nil
### GetPodScheduledArrivalLt

`func (o *SailingInfoTrackingSailingInfo) GetPodScheduledArrivalLt() time.Time`

GetPodScheduledArrivalLt returns the PodScheduledArrivalLt field if non-nil, zero value otherwise.

### GetPodScheduledArrivalLtOk

`func (o *SailingInfoTrackingSailingInfo) GetPodScheduledArrivalLtOk() (*time.Time, bool)`

GetPodScheduledArrivalLtOk returns a tuple with the PodScheduledArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledArrivalLt

`func (o *SailingInfoTrackingSailingInfo) SetPodScheduledArrivalLt(v time.Time)`

SetPodScheduledArrivalLt sets PodScheduledArrivalLt field to given value.

### HasPodScheduledArrivalLt

`func (o *SailingInfoTrackingSailingInfo) HasPodScheduledArrivalLt() bool`

HasPodScheduledArrivalLt returns a boolean if a field has been set.

### SetPodScheduledArrivalLtNil

`func (o *SailingInfoTrackingSailingInfo) SetPodScheduledArrivalLtNil(b bool)`

 SetPodScheduledArrivalLtNil sets the value for PodScheduledArrivalLt to be an explicit nil

### UnsetPodScheduledArrivalLt
`func (o *SailingInfoTrackingSailingInfo) UnsetPodScheduledArrivalLt()`

UnsetPodScheduledArrivalLt ensures that no value is present for PodScheduledArrivalLt, not even an explicit nil
### GetPodScheduledArrivalLtFromSchedule

`func (o *SailingInfoTrackingSailingInfo) GetPodScheduledArrivalLtFromSchedule() time.Time`

GetPodScheduledArrivalLtFromSchedule returns the PodScheduledArrivalLtFromSchedule field if non-nil, zero value otherwise.

### GetPodScheduledArrivalLtFromScheduleOk

`func (o *SailingInfoTrackingSailingInfo) GetPodScheduledArrivalLtFromScheduleOk() (*time.Time, bool)`

GetPodScheduledArrivalLtFromScheduleOk returns a tuple with the PodScheduledArrivalLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledArrivalLtFromSchedule

`func (o *SailingInfoTrackingSailingInfo) SetPodScheduledArrivalLtFromSchedule(v time.Time)`

SetPodScheduledArrivalLtFromSchedule sets PodScheduledArrivalLtFromSchedule field to given value.

### HasPodScheduledArrivalLtFromSchedule

`func (o *SailingInfoTrackingSailingInfo) HasPodScheduledArrivalLtFromSchedule() bool`

HasPodScheduledArrivalLtFromSchedule returns a boolean if a field has been set.

### SetPodScheduledArrivalLtFromScheduleNil

`func (o *SailingInfoTrackingSailingInfo) SetPodScheduledArrivalLtFromScheduleNil(b bool)`

 SetPodScheduledArrivalLtFromScheduleNil sets the value for PodScheduledArrivalLtFromSchedule to be an explicit nil

### UnsetPodScheduledArrivalLtFromSchedule
`func (o *SailingInfoTrackingSailingInfo) UnsetPodScheduledArrivalLtFromSchedule()`

UnsetPodScheduledArrivalLtFromSchedule ensures that no value is present for PodScheduledArrivalLtFromSchedule, not even an explicit nil
### GetPodScheduledDepartureLtFromSchedule

`func (o *SailingInfoTrackingSailingInfo) GetPodScheduledDepartureLtFromSchedule() time.Time`

GetPodScheduledDepartureLtFromSchedule returns the PodScheduledDepartureLtFromSchedule field if non-nil, zero value otherwise.

### GetPodScheduledDepartureLtFromScheduleOk

`func (o *SailingInfoTrackingSailingInfo) GetPodScheduledDepartureLtFromScheduleOk() (*time.Time, bool)`

GetPodScheduledDepartureLtFromScheduleOk returns a tuple with the PodScheduledDepartureLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledDepartureLtFromSchedule

`func (o *SailingInfoTrackingSailingInfo) SetPodScheduledDepartureLtFromSchedule(v time.Time)`

SetPodScheduledDepartureLtFromSchedule sets PodScheduledDepartureLtFromSchedule field to given value.

### HasPodScheduledDepartureLtFromSchedule

`func (o *SailingInfoTrackingSailingInfo) HasPodScheduledDepartureLtFromSchedule() bool`

HasPodScheduledDepartureLtFromSchedule returns a boolean if a field has been set.

### SetPodScheduledDepartureLtFromScheduleNil

`func (o *SailingInfoTrackingSailingInfo) SetPodScheduledDepartureLtFromScheduleNil(b bool)`

 SetPodScheduledDepartureLtFromScheduleNil sets the value for PodScheduledDepartureLtFromSchedule to be an explicit nil

### UnsetPodScheduledDepartureLtFromSchedule
`func (o *SailingInfoTrackingSailingInfo) UnsetPodScheduledDepartureLtFromSchedule()`

UnsetPodScheduledDepartureLtFromSchedule ensures that no value is present for PodScheduledDepartureLtFromSchedule, not even an explicit nil
### GetPodScheduledDischargeLt

`func (o *SailingInfoTrackingSailingInfo) GetPodScheduledDischargeLt() time.Time`

GetPodScheduledDischargeLt returns the PodScheduledDischargeLt field if non-nil, zero value otherwise.

### GetPodScheduledDischargeLtOk

`func (o *SailingInfoTrackingSailingInfo) GetPodScheduledDischargeLtOk() (*time.Time, bool)`

GetPodScheduledDischargeLtOk returns a tuple with the PodScheduledDischargeLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledDischargeLt

`func (o *SailingInfoTrackingSailingInfo) SetPodScheduledDischargeLt(v time.Time)`

SetPodScheduledDischargeLt sets PodScheduledDischargeLt field to given value.

### HasPodScheduledDischargeLt

`func (o *SailingInfoTrackingSailingInfo) HasPodScheduledDischargeLt() bool`

HasPodScheduledDischargeLt returns a boolean if a field has been set.

### SetPodScheduledDischargeLtNil

`func (o *SailingInfoTrackingSailingInfo) SetPodScheduledDischargeLtNil(b bool)`

 SetPodScheduledDischargeLtNil sets the value for PodScheduledDischargeLt to be an explicit nil

### UnsetPodScheduledDischargeLt
`func (o *SailingInfoTrackingSailingInfo) UnsetPodScheduledDischargeLt()`

UnsetPodScheduledDischargeLt ensures that no value is present for PodScheduledDischargeLt, not even an explicit nil
### GetPodTerminalName

`func (o *SailingInfoTrackingSailingInfo) GetPodTerminalName() string`

GetPodTerminalName returns the PodTerminalName field if non-nil, zero value otherwise.

### GetPodTerminalNameOk

`func (o *SailingInfoTrackingSailingInfo) GetPodTerminalNameOk() (*string, bool)`

GetPodTerminalNameOk returns a tuple with the PodTerminalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodTerminalName

`func (o *SailingInfoTrackingSailingInfo) SetPodTerminalName(v string)`

SetPodTerminalName sets PodTerminalName field to given value.

### HasPodTerminalName

`func (o *SailingInfoTrackingSailingInfo) HasPodTerminalName() bool`

HasPodTerminalName returns a boolean if a field has been set.

### SetPodTerminalNameNil

`func (o *SailingInfoTrackingSailingInfo) SetPodTerminalNameNil(b bool)`

 SetPodTerminalNameNil sets the value for PodTerminalName to be an explicit nil

### UnsetPodTerminalName
`func (o *SailingInfoTrackingSailingInfo) UnsetPodTerminalName()`

UnsetPodTerminalName ensures that no value is present for PodTerminalName, not even an explicit nil
### GetPol

`func (o *SailingInfoTrackingSailingInfo) GetPol() string`

GetPol returns the Pol field if non-nil, zero value otherwise.

### GetPolOk

`func (o *SailingInfoTrackingSailingInfo) GetPolOk() (*string, bool)`

GetPolOk returns a tuple with the Pol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPol

`func (o *SailingInfoTrackingSailingInfo) SetPol(v string)`

SetPol sets Pol field to given value.

### HasPol

`func (o *SailingInfoTrackingSailingInfo) HasPol() bool`

HasPol returns a boolean if a field has been set.

### SetPolNil

`func (o *SailingInfoTrackingSailingInfo) SetPolNil(b bool)`

 SetPolNil sets the value for Pol to be an explicit nil

### UnsetPol
`func (o *SailingInfoTrackingSailingInfo) UnsetPol()`

UnsetPol ensures that no value is present for Pol, not even an explicit nil
### GetPolActualArrivalLtFromAis

`func (o *SailingInfoTrackingSailingInfo) GetPolActualArrivalLtFromAis() time.Time`

GetPolActualArrivalLtFromAis returns the PolActualArrivalLtFromAis field if non-nil, zero value otherwise.

### GetPolActualArrivalLtFromAisOk

`func (o *SailingInfoTrackingSailingInfo) GetPolActualArrivalLtFromAisOk() (*time.Time, bool)`

GetPolActualArrivalLtFromAisOk returns a tuple with the PolActualArrivalLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualArrivalLtFromAis

`func (o *SailingInfoTrackingSailingInfo) SetPolActualArrivalLtFromAis(v time.Time)`

SetPolActualArrivalLtFromAis sets PolActualArrivalLtFromAis field to given value.

### HasPolActualArrivalLtFromAis

`func (o *SailingInfoTrackingSailingInfo) HasPolActualArrivalLtFromAis() bool`

HasPolActualArrivalLtFromAis returns a boolean if a field has been set.

### SetPolActualArrivalLtFromAisNil

`func (o *SailingInfoTrackingSailingInfo) SetPolActualArrivalLtFromAisNil(b bool)`

 SetPolActualArrivalLtFromAisNil sets the value for PolActualArrivalLtFromAis to be an explicit nil

### UnsetPolActualArrivalLtFromAis
`func (o *SailingInfoTrackingSailingInfo) UnsetPolActualArrivalLtFromAis()`

UnsetPolActualArrivalLtFromAis ensures that no value is present for PolActualArrivalLtFromAis, not even an explicit nil
### GetPolActualDepartureLt

`func (o *SailingInfoTrackingSailingInfo) GetPolActualDepartureLt() time.Time`

GetPolActualDepartureLt returns the PolActualDepartureLt field if non-nil, zero value otherwise.

### GetPolActualDepartureLtOk

`func (o *SailingInfoTrackingSailingInfo) GetPolActualDepartureLtOk() (*time.Time, bool)`

GetPolActualDepartureLtOk returns a tuple with the PolActualDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualDepartureLt

`func (o *SailingInfoTrackingSailingInfo) SetPolActualDepartureLt(v time.Time)`

SetPolActualDepartureLt sets PolActualDepartureLt field to given value.

### HasPolActualDepartureLt

`func (o *SailingInfoTrackingSailingInfo) HasPolActualDepartureLt() bool`

HasPolActualDepartureLt returns a boolean if a field has been set.

### SetPolActualDepartureLtNil

`func (o *SailingInfoTrackingSailingInfo) SetPolActualDepartureLtNil(b bool)`

 SetPolActualDepartureLtNil sets the value for PolActualDepartureLt to be an explicit nil

### UnsetPolActualDepartureLt
`func (o *SailingInfoTrackingSailingInfo) UnsetPolActualDepartureLt()`

UnsetPolActualDepartureLt ensures that no value is present for PolActualDepartureLt, not even an explicit nil
### GetPolActualDepartureLtFromAis

`func (o *SailingInfoTrackingSailingInfo) GetPolActualDepartureLtFromAis() time.Time`

GetPolActualDepartureLtFromAis returns the PolActualDepartureLtFromAis field if non-nil, zero value otherwise.

### GetPolActualDepartureLtFromAisOk

`func (o *SailingInfoTrackingSailingInfo) GetPolActualDepartureLtFromAisOk() (*time.Time, bool)`

GetPolActualDepartureLtFromAisOk returns a tuple with the PolActualDepartureLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualDepartureLtFromAis

`func (o *SailingInfoTrackingSailingInfo) SetPolActualDepartureLtFromAis(v time.Time)`

SetPolActualDepartureLtFromAis sets PolActualDepartureLtFromAis field to given value.

### HasPolActualDepartureLtFromAis

`func (o *SailingInfoTrackingSailingInfo) HasPolActualDepartureLtFromAis() bool`

HasPolActualDepartureLtFromAis returns a boolean if a field has been set.

### SetPolActualDepartureLtFromAisNil

`func (o *SailingInfoTrackingSailingInfo) SetPolActualDepartureLtFromAisNil(b bool)`

 SetPolActualDepartureLtFromAisNil sets the value for PolActualDepartureLtFromAis to be an explicit nil

### UnsetPolActualDepartureLtFromAis
`func (o *SailingInfoTrackingSailingInfo) UnsetPolActualDepartureLtFromAis()`

UnsetPolActualDepartureLtFromAis ensures that no value is present for PolActualDepartureLtFromAis, not even an explicit nil
### GetPolActualLoadingLt

`func (o *SailingInfoTrackingSailingInfo) GetPolActualLoadingLt() time.Time`

GetPolActualLoadingLt returns the PolActualLoadingLt field if non-nil, zero value otherwise.

### GetPolActualLoadingLtOk

`func (o *SailingInfoTrackingSailingInfo) GetPolActualLoadingLtOk() (*time.Time, bool)`

GetPolActualLoadingLtOk returns a tuple with the PolActualLoadingLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualLoadingLt

`func (o *SailingInfoTrackingSailingInfo) SetPolActualLoadingLt(v time.Time)`

SetPolActualLoadingLt sets PolActualLoadingLt field to given value.

### HasPolActualLoadingLt

`func (o *SailingInfoTrackingSailingInfo) HasPolActualLoadingLt() bool`

HasPolActualLoadingLt returns a boolean if a field has been set.

### SetPolActualLoadingLtNil

`func (o *SailingInfoTrackingSailingInfo) SetPolActualLoadingLtNil(b bool)`

 SetPolActualLoadingLtNil sets the value for PolActualLoadingLt to be an explicit nil

### UnsetPolActualLoadingLt
`func (o *SailingInfoTrackingSailingInfo) UnsetPolActualLoadingLt()`

UnsetPolActualLoadingLt ensures that no value is present for PolActualLoadingLt, not even an explicit nil
### GetPolName

`func (o *SailingInfoTrackingSailingInfo) GetPolName() string`

GetPolName returns the PolName field if non-nil, zero value otherwise.

### GetPolNameOk

`func (o *SailingInfoTrackingSailingInfo) GetPolNameOk() (*string, bool)`

GetPolNameOk returns a tuple with the PolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolName

`func (o *SailingInfoTrackingSailingInfo) SetPolName(v string)`

SetPolName sets PolName field to given value.

### HasPolName

`func (o *SailingInfoTrackingSailingInfo) HasPolName() bool`

HasPolName returns a boolean if a field has been set.

### SetPolNameNil

`func (o *SailingInfoTrackingSailingInfo) SetPolNameNil(b bool)`

 SetPolNameNil sets the value for PolName to be an explicit nil

### UnsetPolName
`func (o *SailingInfoTrackingSailingInfo) UnsetPolName()`

UnsetPolName ensures that no value is present for PolName, not even an explicit nil
### GetPolPredictedArrivalLt

`func (o *SailingInfoTrackingSailingInfo) GetPolPredictedArrivalLt() time.Time`

GetPolPredictedArrivalLt returns the PolPredictedArrivalLt field if non-nil, zero value otherwise.

### GetPolPredictedArrivalLtOk

`func (o *SailingInfoTrackingSailingInfo) GetPolPredictedArrivalLtOk() (*time.Time, bool)`

GetPolPredictedArrivalLtOk returns a tuple with the PolPredictedArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolPredictedArrivalLt

`func (o *SailingInfoTrackingSailingInfo) SetPolPredictedArrivalLt(v time.Time)`

SetPolPredictedArrivalLt sets PolPredictedArrivalLt field to given value.

### HasPolPredictedArrivalLt

`func (o *SailingInfoTrackingSailingInfo) HasPolPredictedArrivalLt() bool`

HasPolPredictedArrivalLt returns a boolean if a field has been set.

### SetPolPredictedArrivalLtNil

`func (o *SailingInfoTrackingSailingInfo) SetPolPredictedArrivalLtNil(b bool)`

 SetPolPredictedArrivalLtNil sets the value for PolPredictedArrivalLt to be an explicit nil

### UnsetPolPredictedArrivalLt
`func (o *SailingInfoTrackingSailingInfo) UnsetPolPredictedArrivalLt()`

UnsetPolPredictedArrivalLt ensures that no value is present for PolPredictedArrivalLt, not even an explicit nil
### GetPolPredictedDepartureLt

`func (o *SailingInfoTrackingSailingInfo) GetPolPredictedDepartureLt() time.Time`

GetPolPredictedDepartureLt returns the PolPredictedDepartureLt field if non-nil, zero value otherwise.

### GetPolPredictedDepartureLtOk

`func (o *SailingInfoTrackingSailingInfo) GetPolPredictedDepartureLtOk() (*time.Time, bool)`

GetPolPredictedDepartureLtOk returns a tuple with the PolPredictedDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolPredictedDepartureLt

`func (o *SailingInfoTrackingSailingInfo) SetPolPredictedDepartureLt(v time.Time)`

SetPolPredictedDepartureLt sets PolPredictedDepartureLt field to given value.

### HasPolPredictedDepartureLt

`func (o *SailingInfoTrackingSailingInfo) HasPolPredictedDepartureLt() bool`

HasPolPredictedDepartureLt returns a boolean if a field has been set.

### SetPolPredictedDepartureLtNil

`func (o *SailingInfoTrackingSailingInfo) SetPolPredictedDepartureLtNil(b bool)`

 SetPolPredictedDepartureLtNil sets the value for PolPredictedDepartureLt to be an explicit nil

### UnsetPolPredictedDepartureLt
`func (o *SailingInfoTrackingSailingInfo) UnsetPolPredictedDepartureLt()`

UnsetPolPredictedDepartureLt ensures that no value is present for PolPredictedDepartureLt, not even an explicit nil
### GetPolScheduledArrivalLtFromSchedule

`func (o *SailingInfoTrackingSailingInfo) GetPolScheduledArrivalLtFromSchedule() time.Time`

GetPolScheduledArrivalLtFromSchedule returns the PolScheduledArrivalLtFromSchedule field if non-nil, zero value otherwise.

### GetPolScheduledArrivalLtFromScheduleOk

`func (o *SailingInfoTrackingSailingInfo) GetPolScheduledArrivalLtFromScheduleOk() (*time.Time, bool)`

GetPolScheduledArrivalLtFromScheduleOk returns a tuple with the PolScheduledArrivalLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledArrivalLtFromSchedule

`func (o *SailingInfoTrackingSailingInfo) SetPolScheduledArrivalLtFromSchedule(v time.Time)`

SetPolScheduledArrivalLtFromSchedule sets PolScheduledArrivalLtFromSchedule field to given value.

### HasPolScheduledArrivalLtFromSchedule

`func (o *SailingInfoTrackingSailingInfo) HasPolScheduledArrivalLtFromSchedule() bool`

HasPolScheduledArrivalLtFromSchedule returns a boolean if a field has been set.

### SetPolScheduledArrivalLtFromScheduleNil

`func (o *SailingInfoTrackingSailingInfo) SetPolScheduledArrivalLtFromScheduleNil(b bool)`

 SetPolScheduledArrivalLtFromScheduleNil sets the value for PolScheduledArrivalLtFromSchedule to be an explicit nil

### UnsetPolScheduledArrivalLtFromSchedule
`func (o *SailingInfoTrackingSailingInfo) UnsetPolScheduledArrivalLtFromSchedule()`

UnsetPolScheduledArrivalLtFromSchedule ensures that no value is present for PolScheduledArrivalLtFromSchedule, not even an explicit nil
### GetPolScheduledDepartureLt

`func (o *SailingInfoTrackingSailingInfo) GetPolScheduledDepartureLt() time.Time`

GetPolScheduledDepartureLt returns the PolScheduledDepartureLt field if non-nil, zero value otherwise.

### GetPolScheduledDepartureLtOk

`func (o *SailingInfoTrackingSailingInfo) GetPolScheduledDepartureLtOk() (*time.Time, bool)`

GetPolScheduledDepartureLtOk returns a tuple with the PolScheduledDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledDepartureLt

`func (o *SailingInfoTrackingSailingInfo) SetPolScheduledDepartureLt(v time.Time)`

SetPolScheduledDepartureLt sets PolScheduledDepartureLt field to given value.

### HasPolScheduledDepartureLt

`func (o *SailingInfoTrackingSailingInfo) HasPolScheduledDepartureLt() bool`

HasPolScheduledDepartureLt returns a boolean if a field has been set.

### SetPolScheduledDepartureLtNil

`func (o *SailingInfoTrackingSailingInfo) SetPolScheduledDepartureLtNil(b bool)`

 SetPolScheduledDepartureLtNil sets the value for PolScheduledDepartureLt to be an explicit nil

### UnsetPolScheduledDepartureLt
`func (o *SailingInfoTrackingSailingInfo) UnsetPolScheduledDepartureLt()`

UnsetPolScheduledDepartureLt ensures that no value is present for PolScheduledDepartureLt, not even an explicit nil
### GetPolScheduledDepartureLtFromSchedule

`func (o *SailingInfoTrackingSailingInfo) GetPolScheduledDepartureLtFromSchedule() time.Time`

GetPolScheduledDepartureLtFromSchedule returns the PolScheduledDepartureLtFromSchedule field if non-nil, zero value otherwise.

### GetPolScheduledDepartureLtFromScheduleOk

`func (o *SailingInfoTrackingSailingInfo) GetPolScheduledDepartureLtFromScheduleOk() (*time.Time, bool)`

GetPolScheduledDepartureLtFromScheduleOk returns a tuple with the PolScheduledDepartureLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledDepartureLtFromSchedule

`func (o *SailingInfoTrackingSailingInfo) SetPolScheduledDepartureLtFromSchedule(v time.Time)`

SetPolScheduledDepartureLtFromSchedule sets PolScheduledDepartureLtFromSchedule field to given value.

### HasPolScheduledDepartureLtFromSchedule

`func (o *SailingInfoTrackingSailingInfo) HasPolScheduledDepartureLtFromSchedule() bool`

HasPolScheduledDepartureLtFromSchedule returns a boolean if a field has been set.

### SetPolScheduledDepartureLtFromScheduleNil

`func (o *SailingInfoTrackingSailingInfo) SetPolScheduledDepartureLtFromScheduleNil(b bool)`

 SetPolScheduledDepartureLtFromScheduleNil sets the value for PolScheduledDepartureLtFromSchedule to be an explicit nil

### UnsetPolScheduledDepartureLtFromSchedule
`func (o *SailingInfoTrackingSailingInfo) UnsetPolScheduledDepartureLtFromSchedule()`

UnsetPolScheduledDepartureLtFromSchedule ensures that no value is present for PolScheduledDepartureLtFromSchedule, not even an explicit nil
### GetPolScheduledLoadingLt

`func (o *SailingInfoTrackingSailingInfo) GetPolScheduledLoadingLt() time.Time`

GetPolScheduledLoadingLt returns the PolScheduledLoadingLt field if non-nil, zero value otherwise.

### GetPolScheduledLoadingLtOk

`func (o *SailingInfoTrackingSailingInfo) GetPolScheduledLoadingLtOk() (*time.Time, bool)`

GetPolScheduledLoadingLtOk returns a tuple with the PolScheduledLoadingLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledLoadingLt

`func (o *SailingInfoTrackingSailingInfo) SetPolScheduledLoadingLt(v time.Time)`

SetPolScheduledLoadingLt sets PolScheduledLoadingLt field to given value.

### HasPolScheduledLoadingLt

`func (o *SailingInfoTrackingSailingInfo) HasPolScheduledLoadingLt() bool`

HasPolScheduledLoadingLt returns a boolean if a field has been set.

### SetPolScheduledLoadingLtNil

`func (o *SailingInfoTrackingSailingInfo) SetPolScheduledLoadingLtNil(b bool)`

 SetPolScheduledLoadingLtNil sets the value for PolScheduledLoadingLt to be an explicit nil

### UnsetPolScheduledLoadingLt
`func (o *SailingInfoTrackingSailingInfo) UnsetPolScheduledLoadingLt()`

UnsetPolScheduledLoadingLt ensures that no value is present for PolScheduledLoadingLt, not even an explicit nil
### GetPolTerminalName

`func (o *SailingInfoTrackingSailingInfo) GetPolTerminalName() string`

GetPolTerminalName returns the PolTerminalName field if non-nil, zero value otherwise.

### GetPolTerminalNameOk

`func (o *SailingInfoTrackingSailingInfo) GetPolTerminalNameOk() (*string, bool)`

GetPolTerminalNameOk returns a tuple with the PolTerminalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolTerminalName

`func (o *SailingInfoTrackingSailingInfo) SetPolTerminalName(v string)`

SetPolTerminalName sets PolTerminalName field to given value.

### HasPolTerminalName

`func (o *SailingInfoTrackingSailingInfo) HasPolTerminalName() bool`

HasPolTerminalName returns a boolean if a field has been set.

### SetPolTerminalNameNil

`func (o *SailingInfoTrackingSailingInfo) SetPolTerminalNameNil(b bool)`

 SetPolTerminalNameNil sets the value for PolTerminalName to be an explicit nil

### UnsetPolTerminalName
`func (o *SailingInfoTrackingSailingInfo) UnsetPolTerminalName()`

UnsetPolTerminalName ensures that no value is present for PolTerminalName, not even an explicit nil
### GetUpdated

`func (o *SailingInfoTrackingSailingInfo) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *SailingInfoTrackingSailingInfo) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *SailingInfoTrackingSailingInfo) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *SailingInfoTrackingSailingInfo) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *SailingInfoTrackingSailingInfo) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *SailingInfoTrackingSailingInfo) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetVesselFlag

`func (o *SailingInfoTrackingSailingInfo) GetVesselFlag() string`

GetVesselFlag returns the VesselFlag field if non-nil, zero value otherwise.

### GetVesselFlagOk

`func (o *SailingInfoTrackingSailingInfo) GetVesselFlagOk() (*string, bool)`

GetVesselFlagOk returns a tuple with the VesselFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselFlag

`func (o *SailingInfoTrackingSailingInfo) SetVesselFlag(v string)`

SetVesselFlag sets VesselFlag field to given value.

### HasVesselFlag

`func (o *SailingInfoTrackingSailingInfo) HasVesselFlag() bool`

HasVesselFlag returns a boolean if a field has been set.

### SetVesselFlagNil

`func (o *SailingInfoTrackingSailingInfo) SetVesselFlagNil(b bool)`

 SetVesselFlagNil sets the value for VesselFlag to be an explicit nil

### UnsetVesselFlag
`func (o *SailingInfoTrackingSailingInfo) UnsetVesselFlag()`

UnsetVesselFlag ensures that no value is present for VesselFlag, not even an explicit nil
### GetVesselLeg

`func (o *SailingInfoTrackingSailingInfo) GetVesselLeg() int32`

GetVesselLeg returns the VesselLeg field if non-nil, zero value otherwise.

### GetVesselLegOk

`func (o *SailingInfoTrackingSailingInfo) GetVesselLegOk() (*int32, bool)`

GetVesselLegOk returns a tuple with the VesselLeg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselLeg

`func (o *SailingInfoTrackingSailingInfo) SetVesselLeg(v int32)`

SetVesselLeg sets VesselLeg field to given value.

### HasVesselLeg

`func (o *SailingInfoTrackingSailingInfo) HasVesselLeg() bool`

HasVesselLeg returns a boolean if a field has been set.

### GetVesselName

`func (o *SailingInfoTrackingSailingInfo) GetVesselName() string`

GetVesselName returns the VesselName field if non-nil, zero value otherwise.

### GetVesselNameOk

`func (o *SailingInfoTrackingSailingInfo) GetVesselNameOk() (*string, bool)`

GetVesselNameOk returns a tuple with the VesselName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselName

`func (o *SailingInfoTrackingSailingInfo) SetVesselName(v string)`

SetVesselName sets VesselName field to given value.

### HasVesselName

`func (o *SailingInfoTrackingSailingInfo) HasVesselName() bool`

HasVesselName returns a boolean if a field has been set.

### SetVesselNameNil

`func (o *SailingInfoTrackingSailingInfo) SetVesselNameNil(b bool)`

 SetVesselNameNil sets the value for VesselName to be an explicit nil

### UnsetVesselName
`func (o *SailingInfoTrackingSailingInfo) UnsetVesselName()`

UnsetVesselName ensures that no value is present for VesselName, not even an explicit nil
### GetVoyageNo

`func (o *SailingInfoTrackingSailingInfo) GetVoyageNo() string`

GetVoyageNo returns the VoyageNo field if non-nil, zero value otherwise.

### GetVoyageNoOk

`func (o *SailingInfoTrackingSailingInfo) GetVoyageNoOk() (*string, bool)`

GetVoyageNoOk returns a tuple with the VoyageNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageNo

`func (o *SailingInfoTrackingSailingInfo) SetVoyageNo(v string)`

SetVoyageNo sets VoyageNo field to given value.

### HasVoyageNo

`func (o *SailingInfoTrackingSailingInfo) HasVoyageNo() bool`

HasVoyageNo returns a boolean if a field has been set.

### SetVoyageNoNil

`func (o *SailingInfoTrackingSailingInfo) SetVoyageNoNil(b bool)`

 SetVoyageNoNil sets the value for VoyageNo to be an explicit nil

### UnsetVoyageNo
`func (o *SailingInfoTrackingSailingInfo) UnsetVoyageNo()`

UnsetVoyageNo ensures that no value is present for VoyageNo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


