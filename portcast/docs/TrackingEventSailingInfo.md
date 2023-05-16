# TrackingEventSailingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierNo** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Imo** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Pod** | Pointer to **string** |  | [optional] 
**PodActualArrivalLt** | Pointer to **string** |  | [optional] 
**PodActualArrivalLtFromAis** | Pointer to **string** |  | [optional] 
**PodActualDepartureLt** | Pointer to **string** |  | [optional] 
**PodActualDepartureLtFromAis** | Pointer to **string** |  | [optional] 
**PodActualDischargeLt** | Pointer to **string** |  | [optional] 
**PodName** | Pointer to **string** |  | [optional] 
**PodPredictedArrivalLt** | Pointer to **string** |  | [optional] 
**PodPredictedDepartureLt** | Pointer to **string** |  | [optional] 
**PodScheduledArrivalLt** | Pointer to **string** |  | [optional] 
**PodScheduledArrivalLtFromSchedule** | Pointer to **string** |  | [optional] 
**PodScheduledDischargeLt** | Pointer to **string** |  | [optional] 
**PodScheduledDepartureLtFromSchedule** | Pointer to **string** |  | [optional] 
**Pol** | Pointer to **string** |  | [optional] 
**PolActualArrivalLtFromAis** | Pointer to **string** |  | [optional] 
**PolActualDepartureLtFromAis** | Pointer to **string** |  | [optional] 
**PolActualLoadingLt** | Pointer to **string** |  | [optional] 
**PolName** | Pointer to **string** |  | [optional] 
**PolPredictedArrivalLt** | Pointer to **string** |  | [optional] 
**PolPredictedDepartureLt** | Pointer to **string** |  | [optional] 
**PolScheduledArrivalLtFromSchedule** | Pointer to **string** |  | [optional] 
**PolScheduledDepartureLt** | Pointer to **string** |  | [optional] 
**PolScheduledLoadingLt** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**VesselLeg** | Pointer to **int32** |  | [optional] 
**VesselName** | Pointer to **string** |  | [optional] 
**VoyageNo** | Pointer to **string** |  | [optional] 

## Methods

### NewTrackingEventSailingInfo

`func NewTrackingEventSailingInfo() *TrackingEventSailingInfo`

NewTrackingEventSailingInfo instantiates a new TrackingEventSailingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingEventSailingInfoWithDefaults

`func NewTrackingEventSailingInfoWithDefaults() *TrackingEventSailingInfo`

NewTrackingEventSailingInfoWithDefaults instantiates a new TrackingEventSailingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierNo

`func (o *TrackingEventSailingInfo) GetCarrierNo() string`

GetCarrierNo returns the CarrierNo field if non-nil, zero value otherwise.

### GetCarrierNoOk

`func (o *TrackingEventSailingInfo) GetCarrierNoOk() (*string, bool)`

GetCarrierNoOk returns a tuple with the CarrierNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierNo

`func (o *TrackingEventSailingInfo) SetCarrierNo(v string)`

SetCarrierNo sets CarrierNo field to given value.

### HasCarrierNo

`func (o *TrackingEventSailingInfo) HasCarrierNo() bool`

HasCarrierNo returns a boolean if a field has been set.

### GetCreated

`func (o *TrackingEventSailingInfo) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TrackingEventSailingInfo) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TrackingEventSailingInfo) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TrackingEventSailingInfo) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *TrackingEventSailingInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackingEventSailingInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackingEventSailingInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrackingEventSailingInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImo

`func (o *TrackingEventSailingInfo) GetImo() string`

GetImo returns the Imo field if non-nil, zero value otherwise.

### GetImoOk

`func (o *TrackingEventSailingInfo) GetImoOk() (*string, bool)`

GetImoOk returns a tuple with the Imo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImo

`func (o *TrackingEventSailingInfo) SetImo(v string)`

SetImo sets Imo field to given value.

### HasImo

`func (o *TrackingEventSailingInfo) HasImo() bool`

HasImo returns a boolean if a field has been set.

### GetIsActive

`func (o *TrackingEventSailingInfo) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *TrackingEventSailingInfo) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *TrackingEventSailingInfo) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *TrackingEventSailingInfo) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetPod

`func (o *TrackingEventSailingInfo) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *TrackingEventSailingInfo) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *TrackingEventSailingInfo) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *TrackingEventSailingInfo) HasPod() bool`

HasPod returns a boolean if a field has been set.

### GetPodActualArrivalLt

`func (o *TrackingEventSailingInfo) GetPodActualArrivalLt() string`

GetPodActualArrivalLt returns the PodActualArrivalLt field if non-nil, zero value otherwise.

### GetPodActualArrivalLtOk

`func (o *TrackingEventSailingInfo) GetPodActualArrivalLtOk() (*string, bool)`

GetPodActualArrivalLtOk returns a tuple with the PodActualArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualArrivalLt

`func (o *TrackingEventSailingInfo) SetPodActualArrivalLt(v string)`

SetPodActualArrivalLt sets PodActualArrivalLt field to given value.

### HasPodActualArrivalLt

`func (o *TrackingEventSailingInfo) HasPodActualArrivalLt() bool`

HasPodActualArrivalLt returns a boolean if a field has been set.

### GetPodActualArrivalLtFromAis

`func (o *TrackingEventSailingInfo) GetPodActualArrivalLtFromAis() string`

GetPodActualArrivalLtFromAis returns the PodActualArrivalLtFromAis field if non-nil, zero value otherwise.

### GetPodActualArrivalLtFromAisOk

`func (o *TrackingEventSailingInfo) GetPodActualArrivalLtFromAisOk() (*string, bool)`

GetPodActualArrivalLtFromAisOk returns a tuple with the PodActualArrivalLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualArrivalLtFromAis

`func (o *TrackingEventSailingInfo) SetPodActualArrivalLtFromAis(v string)`

SetPodActualArrivalLtFromAis sets PodActualArrivalLtFromAis field to given value.

### HasPodActualArrivalLtFromAis

`func (o *TrackingEventSailingInfo) HasPodActualArrivalLtFromAis() bool`

HasPodActualArrivalLtFromAis returns a boolean if a field has been set.

### GetPodActualDepartureLt

`func (o *TrackingEventSailingInfo) GetPodActualDepartureLt() string`

GetPodActualDepartureLt returns the PodActualDepartureLt field if non-nil, zero value otherwise.

### GetPodActualDepartureLtOk

`func (o *TrackingEventSailingInfo) GetPodActualDepartureLtOk() (*string, bool)`

GetPodActualDepartureLtOk returns a tuple with the PodActualDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualDepartureLt

`func (o *TrackingEventSailingInfo) SetPodActualDepartureLt(v string)`

SetPodActualDepartureLt sets PodActualDepartureLt field to given value.

### HasPodActualDepartureLt

`func (o *TrackingEventSailingInfo) HasPodActualDepartureLt() bool`

HasPodActualDepartureLt returns a boolean if a field has been set.

### GetPodActualDepartureLtFromAis

`func (o *TrackingEventSailingInfo) GetPodActualDepartureLtFromAis() string`

GetPodActualDepartureLtFromAis returns the PodActualDepartureLtFromAis field if non-nil, zero value otherwise.

### GetPodActualDepartureLtFromAisOk

`func (o *TrackingEventSailingInfo) GetPodActualDepartureLtFromAisOk() (*string, bool)`

GetPodActualDepartureLtFromAisOk returns a tuple with the PodActualDepartureLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualDepartureLtFromAis

`func (o *TrackingEventSailingInfo) SetPodActualDepartureLtFromAis(v string)`

SetPodActualDepartureLtFromAis sets PodActualDepartureLtFromAis field to given value.

### HasPodActualDepartureLtFromAis

`func (o *TrackingEventSailingInfo) HasPodActualDepartureLtFromAis() bool`

HasPodActualDepartureLtFromAis returns a boolean if a field has been set.

### GetPodActualDischargeLt

`func (o *TrackingEventSailingInfo) GetPodActualDischargeLt() string`

GetPodActualDischargeLt returns the PodActualDischargeLt field if non-nil, zero value otherwise.

### GetPodActualDischargeLtOk

`func (o *TrackingEventSailingInfo) GetPodActualDischargeLtOk() (*string, bool)`

GetPodActualDischargeLtOk returns a tuple with the PodActualDischargeLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualDischargeLt

`func (o *TrackingEventSailingInfo) SetPodActualDischargeLt(v string)`

SetPodActualDischargeLt sets PodActualDischargeLt field to given value.

### HasPodActualDischargeLt

`func (o *TrackingEventSailingInfo) HasPodActualDischargeLt() bool`

HasPodActualDischargeLt returns a boolean if a field has been set.

### GetPodName

`func (o *TrackingEventSailingInfo) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *TrackingEventSailingInfo) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *TrackingEventSailingInfo) SetPodName(v string)`

SetPodName sets PodName field to given value.

### HasPodName

`func (o *TrackingEventSailingInfo) HasPodName() bool`

HasPodName returns a boolean if a field has been set.

### GetPodPredictedArrivalLt

`func (o *TrackingEventSailingInfo) GetPodPredictedArrivalLt() string`

GetPodPredictedArrivalLt returns the PodPredictedArrivalLt field if non-nil, zero value otherwise.

### GetPodPredictedArrivalLtOk

`func (o *TrackingEventSailingInfo) GetPodPredictedArrivalLtOk() (*string, bool)`

GetPodPredictedArrivalLtOk returns a tuple with the PodPredictedArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPredictedArrivalLt

`func (o *TrackingEventSailingInfo) SetPodPredictedArrivalLt(v string)`

SetPodPredictedArrivalLt sets PodPredictedArrivalLt field to given value.

### HasPodPredictedArrivalLt

`func (o *TrackingEventSailingInfo) HasPodPredictedArrivalLt() bool`

HasPodPredictedArrivalLt returns a boolean if a field has been set.

### GetPodPredictedDepartureLt

`func (o *TrackingEventSailingInfo) GetPodPredictedDepartureLt() string`

GetPodPredictedDepartureLt returns the PodPredictedDepartureLt field if non-nil, zero value otherwise.

### GetPodPredictedDepartureLtOk

`func (o *TrackingEventSailingInfo) GetPodPredictedDepartureLtOk() (*string, bool)`

GetPodPredictedDepartureLtOk returns a tuple with the PodPredictedDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPredictedDepartureLt

`func (o *TrackingEventSailingInfo) SetPodPredictedDepartureLt(v string)`

SetPodPredictedDepartureLt sets PodPredictedDepartureLt field to given value.

### HasPodPredictedDepartureLt

`func (o *TrackingEventSailingInfo) HasPodPredictedDepartureLt() bool`

HasPodPredictedDepartureLt returns a boolean if a field has been set.

### GetPodScheduledArrivalLt

`func (o *TrackingEventSailingInfo) GetPodScheduledArrivalLt() string`

GetPodScheduledArrivalLt returns the PodScheduledArrivalLt field if non-nil, zero value otherwise.

### GetPodScheduledArrivalLtOk

`func (o *TrackingEventSailingInfo) GetPodScheduledArrivalLtOk() (*string, bool)`

GetPodScheduledArrivalLtOk returns a tuple with the PodScheduledArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledArrivalLt

`func (o *TrackingEventSailingInfo) SetPodScheduledArrivalLt(v string)`

SetPodScheduledArrivalLt sets PodScheduledArrivalLt field to given value.

### HasPodScheduledArrivalLt

`func (o *TrackingEventSailingInfo) HasPodScheduledArrivalLt() bool`

HasPodScheduledArrivalLt returns a boolean if a field has been set.

### GetPodScheduledArrivalLtFromSchedule

`func (o *TrackingEventSailingInfo) GetPodScheduledArrivalLtFromSchedule() string`

GetPodScheduledArrivalLtFromSchedule returns the PodScheduledArrivalLtFromSchedule field if non-nil, zero value otherwise.

### GetPodScheduledArrivalLtFromScheduleOk

`func (o *TrackingEventSailingInfo) GetPodScheduledArrivalLtFromScheduleOk() (*string, bool)`

GetPodScheduledArrivalLtFromScheduleOk returns a tuple with the PodScheduledArrivalLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledArrivalLtFromSchedule

`func (o *TrackingEventSailingInfo) SetPodScheduledArrivalLtFromSchedule(v string)`

SetPodScheduledArrivalLtFromSchedule sets PodScheduledArrivalLtFromSchedule field to given value.

### HasPodScheduledArrivalLtFromSchedule

`func (o *TrackingEventSailingInfo) HasPodScheduledArrivalLtFromSchedule() bool`

HasPodScheduledArrivalLtFromSchedule returns a boolean if a field has been set.

### GetPodScheduledDischargeLt

`func (o *TrackingEventSailingInfo) GetPodScheduledDischargeLt() string`

GetPodScheduledDischargeLt returns the PodScheduledDischargeLt field if non-nil, zero value otherwise.

### GetPodScheduledDischargeLtOk

`func (o *TrackingEventSailingInfo) GetPodScheduledDischargeLtOk() (*string, bool)`

GetPodScheduledDischargeLtOk returns a tuple with the PodScheduledDischargeLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledDischargeLt

`func (o *TrackingEventSailingInfo) SetPodScheduledDischargeLt(v string)`

SetPodScheduledDischargeLt sets PodScheduledDischargeLt field to given value.

### HasPodScheduledDischargeLt

`func (o *TrackingEventSailingInfo) HasPodScheduledDischargeLt() bool`

HasPodScheduledDischargeLt returns a boolean if a field has been set.

### GetPodScheduledDepartureLtFromSchedule

`func (o *TrackingEventSailingInfo) GetPodScheduledDepartureLtFromSchedule() string`

GetPodScheduledDepartureLtFromSchedule returns the PodScheduledDepartureLtFromSchedule field if non-nil, zero value otherwise.

### GetPodScheduledDepartureLtFromScheduleOk

`func (o *TrackingEventSailingInfo) GetPodScheduledDepartureLtFromScheduleOk() (*string, bool)`

GetPodScheduledDepartureLtFromScheduleOk returns a tuple with the PodScheduledDepartureLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledDepartureLtFromSchedule

`func (o *TrackingEventSailingInfo) SetPodScheduledDepartureLtFromSchedule(v string)`

SetPodScheduledDepartureLtFromSchedule sets PodScheduledDepartureLtFromSchedule field to given value.

### HasPodScheduledDepartureLtFromSchedule

`func (o *TrackingEventSailingInfo) HasPodScheduledDepartureLtFromSchedule() bool`

HasPodScheduledDepartureLtFromSchedule returns a boolean if a field has been set.

### GetPol

`func (o *TrackingEventSailingInfo) GetPol() string`

GetPol returns the Pol field if non-nil, zero value otherwise.

### GetPolOk

`func (o *TrackingEventSailingInfo) GetPolOk() (*string, bool)`

GetPolOk returns a tuple with the Pol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPol

`func (o *TrackingEventSailingInfo) SetPol(v string)`

SetPol sets Pol field to given value.

### HasPol

`func (o *TrackingEventSailingInfo) HasPol() bool`

HasPol returns a boolean if a field has been set.

### GetPolActualArrivalLtFromAis

`func (o *TrackingEventSailingInfo) GetPolActualArrivalLtFromAis() string`

GetPolActualArrivalLtFromAis returns the PolActualArrivalLtFromAis field if non-nil, zero value otherwise.

### GetPolActualArrivalLtFromAisOk

`func (o *TrackingEventSailingInfo) GetPolActualArrivalLtFromAisOk() (*string, bool)`

GetPolActualArrivalLtFromAisOk returns a tuple with the PolActualArrivalLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualArrivalLtFromAis

`func (o *TrackingEventSailingInfo) SetPolActualArrivalLtFromAis(v string)`

SetPolActualArrivalLtFromAis sets PolActualArrivalLtFromAis field to given value.

### HasPolActualArrivalLtFromAis

`func (o *TrackingEventSailingInfo) HasPolActualArrivalLtFromAis() bool`

HasPolActualArrivalLtFromAis returns a boolean if a field has been set.

### GetPolActualDepartureLtFromAis

`func (o *TrackingEventSailingInfo) GetPolActualDepartureLtFromAis() string`

GetPolActualDepartureLtFromAis returns the PolActualDepartureLtFromAis field if non-nil, zero value otherwise.

### GetPolActualDepartureLtFromAisOk

`func (o *TrackingEventSailingInfo) GetPolActualDepartureLtFromAisOk() (*string, bool)`

GetPolActualDepartureLtFromAisOk returns a tuple with the PolActualDepartureLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualDepartureLtFromAis

`func (o *TrackingEventSailingInfo) SetPolActualDepartureLtFromAis(v string)`

SetPolActualDepartureLtFromAis sets PolActualDepartureLtFromAis field to given value.

### HasPolActualDepartureLtFromAis

`func (o *TrackingEventSailingInfo) HasPolActualDepartureLtFromAis() bool`

HasPolActualDepartureLtFromAis returns a boolean if a field has been set.

### GetPolActualLoadingLt

`func (o *TrackingEventSailingInfo) GetPolActualLoadingLt() string`

GetPolActualLoadingLt returns the PolActualLoadingLt field if non-nil, zero value otherwise.

### GetPolActualLoadingLtOk

`func (o *TrackingEventSailingInfo) GetPolActualLoadingLtOk() (*string, bool)`

GetPolActualLoadingLtOk returns a tuple with the PolActualLoadingLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualLoadingLt

`func (o *TrackingEventSailingInfo) SetPolActualLoadingLt(v string)`

SetPolActualLoadingLt sets PolActualLoadingLt field to given value.

### HasPolActualLoadingLt

`func (o *TrackingEventSailingInfo) HasPolActualLoadingLt() bool`

HasPolActualLoadingLt returns a boolean if a field has been set.

### GetPolName

`func (o *TrackingEventSailingInfo) GetPolName() string`

GetPolName returns the PolName field if non-nil, zero value otherwise.

### GetPolNameOk

`func (o *TrackingEventSailingInfo) GetPolNameOk() (*string, bool)`

GetPolNameOk returns a tuple with the PolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolName

`func (o *TrackingEventSailingInfo) SetPolName(v string)`

SetPolName sets PolName field to given value.

### HasPolName

`func (o *TrackingEventSailingInfo) HasPolName() bool`

HasPolName returns a boolean if a field has been set.

### GetPolPredictedArrivalLt

`func (o *TrackingEventSailingInfo) GetPolPredictedArrivalLt() string`

GetPolPredictedArrivalLt returns the PolPredictedArrivalLt field if non-nil, zero value otherwise.

### GetPolPredictedArrivalLtOk

`func (o *TrackingEventSailingInfo) GetPolPredictedArrivalLtOk() (*string, bool)`

GetPolPredictedArrivalLtOk returns a tuple with the PolPredictedArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolPredictedArrivalLt

`func (o *TrackingEventSailingInfo) SetPolPredictedArrivalLt(v string)`

SetPolPredictedArrivalLt sets PolPredictedArrivalLt field to given value.

### HasPolPredictedArrivalLt

`func (o *TrackingEventSailingInfo) HasPolPredictedArrivalLt() bool`

HasPolPredictedArrivalLt returns a boolean if a field has been set.

### GetPolPredictedDepartureLt

`func (o *TrackingEventSailingInfo) GetPolPredictedDepartureLt() string`

GetPolPredictedDepartureLt returns the PolPredictedDepartureLt field if non-nil, zero value otherwise.

### GetPolPredictedDepartureLtOk

`func (o *TrackingEventSailingInfo) GetPolPredictedDepartureLtOk() (*string, bool)`

GetPolPredictedDepartureLtOk returns a tuple with the PolPredictedDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolPredictedDepartureLt

`func (o *TrackingEventSailingInfo) SetPolPredictedDepartureLt(v string)`

SetPolPredictedDepartureLt sets PolPredictedDepartureLt field to given value.

### HasPolPredictedDepartureLt

`func (o *TrackingEventSailingInfo) HasPolPredictedDepartureLt() bool`

HasPolPredictedDepartureLt returns a boolean if a field has been set.

### GetPolScheduledArrivalLtFromSchedule

`func (o *TrackingEventSailingInfo) GetPolScheduledArrivalLtFromSchedule() string`

GetPolScheduledArrivalLtFromSchedule returns the PolScheduledArrivalLtFromSchedule field if non-nil, zero value otherwise.

### GetPolScheduledArrivalLtFromScheduleOk

`func (o *TrackingEventSailingInfo) GetPolScheduledArrivalLtFromScheduleOk() (*string, bool)`

GetPolScheduledArrivalLtFromScheduleOk returns a tuple with the PolScheduledArrivalLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledArrivalLtFromSchedule

`func (o *TrackingEventSailingInfo) SetPolScheduledArrivalLtFromSchedule(v string)`

SetPolScheduledArrivalLtFromSchedule sets PolScheduledArrivalLtFromSchedule field to given value.

### HasPolScheduledArrivalLtFromSchedule

`func (o *TrackingEventSailingInfo) HasPolScheduledArrivalLtFromSchedule() bool`

HasPolScheduledArrivalLtFromSchedule returns a boolean if a field has been set.

### GetPolScheduledDepartureLt

`func (o *TrackingEventSailingInfo) GetPolScheduledDepartureLt() string`

GetPolScheduledDepartureLt returns the PolScheduledDepartureLt field if non-nil, zero value otherwise.

### GetPolScheduledDepartureLtOk

`func (o *TrackingEventSailingInfo) GetPolScheduledDepartureLtOk() (*string, bool)`

GetPolScheduledDepartureLtOk returns a tuple with the PolScheduledDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledDepartureLt

`func (o *TrackingEventSailingInfo) SetPolScheduledDepartureLt(v string)`

SetPolScheduledDepartureLt sets PolScheduledDepartureLt field to given value.

### HasPolScheduledDepartureLt

`func (o *TrackingEventSailingInfo) HasPolScheduledDepartureLt() bool`

HasPolScheduledDepartureLt returns a boolean if a field has been set.

### GetPolScheduledLoadingLt

`func (o *TrackingEventSailingInfo) GetPolScheduledLoadingLt() string`

GetPolScheduledLoadingLt returns the PolScheduledLoadingLt field if non-nil, zero value otherwise.

### GetPolScheduledLoadingLtOk

`func (o *TrackingEventSailingInfo) GetPolScheduledLoadingLtOk() (*string, bool)`

GetPolScheduledLoadingLtOk returns a tuple with the PolScheduledLoadingLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledLoadingLt

`func (o *TrackingEventSailingInfo) SetPolScheduledLoadingLt(v string)`

SetPolScheduledLoadingLt sets PolScheduledLoadingLt field to given value.

### HasPolScheduledLoadingLt

`func (o *TrackingEventSailingInfo) HasPolScheduledLoadingLt() bool`

HasPolScheduledLoadingLt returns a boolean if a field has been set.

### GetStatusCode

`func (o *TrackingEventSailingInfo) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *TrackingEventSailingInfo) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *TrackingEventSailingInfo) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *TrackingEventSailingInfo) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetUpdated

`func (o *TrackingEventSailingInfo) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TrackingEventSailingInfo) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TrackingEventSailingInfo) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *TrackingEventSailingInfo) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetVesselLeg

`func (o *TrackingEventSailingInfo) GetVesselLeg() int32`

GetVesselLeg returns the VesselLeg field if non-nil, zero value otherwise.

### GetVesselLegOk

`func (o *TrackingEventSailingInfo) GetVesselLegOk() (*int32, bool)`

GetVesselLegOk returns a tuple with the VesselLeg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselLeg

`func (o *TrackingEventSailingInfo) SetVesselLeg(v int32)`

SetVesselLeg sets VesselLeg field to given value.

### HasVesselLeg

`func (o *TrackingEventSailingInfo) HasVesselLeg() bool`

HasVesselLeg returns a boolean if a field has been set.

### GetVesselName

`func (o *TrackingEventSailingInfo) GetVesselName() string`

GetVesselName returns the VesselName field if non-nil, zero value otherwise.

### GetVesselNameOk

`func (o *TrackingEventSailingInfo) GetVesselNameOk() (*string, bool)`

GetVesselNameOk returns a tuple with the VesselName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselName

`func (o *TrackingEventSailingInfo) SetVesselName(v string)`

SetVesselName sets VesselName field to given value.

### HasVesselName

`func (o *TrackingEventSailingInfo) HasVesselName() bool`

HasVesselName returns a boolean if a field has been set.

### GetVoyageNo

`func (o *TrackingEventSailingInfo) GetVoyageNo() string`

GetVoyageNo returns the VoyageNo field if non-nil, zero value otherwise.

### GetVoyageNoOk

`func (o *TrackingEventSailingInfo) GetVoyageNoOk() (*string, bool)`

GetVoyageNoOk returns a tuple with the VoyageNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageNo

`func (o *TrackingEventSailingInfo) SetVoyageNo(v string)`

SetVoyageNo sets VoyageNo field to given value.

### HasVoyageNo

`func (o *TrackingEventSailingInfo) HasVoyageNo() bool`

HasVoyageNo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


