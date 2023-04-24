# TrackingEventBillOfLading

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlNo** | Pointer to **string** |  | [optional] 
**CarrierNo** | Pointer to **string** |  | [optional] 
**CntrNo** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PlaceOfDelivery** | Pointer to **string** |  | [optional] 
**PlaceOfDeliveryName** | Pointer to **string** |  | [optional] 
**PlaceOfReceipt** | Pointer to **string** |  | [optional] 
**PlaceOfReceiptName** | Pointer to **string** |  | [optional] 
**Pod** | Pointer to **string** |  | [optional] 
**PodActualArrivalLt** | Pointer to **string** |  | [optional] 
**PodActualArrivalUtc** | Pointer to **string** |  | [optional] 
**PodActualArrivalLtFromAis** | Pointer to **string** |  | [optional] 
**PodActualDepartureLtFromAis** | Pointer to **string** |  | [optional] 
**PodActualDischargeLt** | Pointer to **string** |  | [optional] 
**PodName** | Pointer to **string** |  | [optional] 
**PodPredictedArrivalLt** | Pointer to **string** |  | [optional] 
**PodPredictedArrivalUtc** | Pointer to **string** |  | [optional] 
**PodPredictedDepartureLt** | Pointer to **string** |  | [optional] 
**PodScheduledArrivalLt** | Pointer to **string** |  | [optional] 
**PodScheduledArrivalLtFromSchedule** | Pointer to **string** |  | [optional] 
**PodScheduledDepartureLtFromSchedule** | Pointer to **string** |  | [optional] 
**PodScheduledDischargeLt** | Pointer to **string** |  | [optional] 
**Pol** | Pointer to **string** |  | [optional] 
**PolActualArrivalLt** | Pointer to **string** |  | [optional] 
**PolActualArrivalLtFromAis** | Pointer to **string** |  | [optional] 
**PolActualDepartureLt** | Pointer to **string** |  | [optional] 
**PolActualDepartureLtFromAis** | Pointer to **string** |  | [optional] 
**PolActualLoadingLt** | Pointer to **string** |  | [optional] 
**PolName** | Pointer to **string** |  | [optional] 
**PolPredictedArrivalLt** | Pointer to **string** |  | [optional] 
**PolPredictedDepartureLt** | Pointer to **string** |  | [optional] 
**PolScheduledArrivalLtFromSchedule** | Pointer to **string** |  | [optional] 
**PolScheduledDepartureLt** | Pointer to **string** |  | [optional] 
**PolScheduledDepartureLtFromSchedule** | Pointer to **string** |  | [optional] 
**PolScheduledLoadingLt** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTrackingEventBillOfLading

`func NewTrackingEventBillOfLading() *TrackingEventBillOfLading`

NewTrackingEventBillOfLading instantiates a new TrackingEventBillOfLading object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingEventBillOfLadingWithDefaults

`func NewTrackingEventBillOfLadingWithDefaults() *TrackingEventBillOfLading`

NewTrackingEventBillOfLadingWithDefaults instantiates a new TrackingEventBillOfLading object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlNo

`func (o *TrackingEventBillOfLading) GetBlNo() string`

GetBlNo returns the BlNo field if non-nil, zero value otherwise.

### GetBlNoOk

`func (o *TrackingEventBillOfLading) GetBlNoOk() (*string, bool)`

GetBlNoOk returns a tuple with the BlNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlNo

`func (o *TrackingEventBillOfLading) SetBlNo(v string)`

SetBlNo sets BlNo field to given value.

### HasBlNo

`func (o *TrackingEventBillOfLading) HasBlNo() bool`

HasBlNo returns a boolean if a field has been set.

### GetCarrierNo

`func (o *TrackingEventBillOfLading) GetCarrierNo() string`

GetCarrierNo returns the CarrierNo field if non-nil, zero value otherwise.

### GetCarrierNoOk

`func (o *TrackingEventBillOfLading) GetCarrierNoOk() (*string, bool)`

GetCarrierNoOk returns a tuple with the CarrierNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierNo

`func (o *TrackingEventBillOfLading) SetCarrierNo(v string)`

SetCarrierNo sets CarrierNo field to given value.

### HasCarrierNo

`func (o *TrackingEventBillOfLading) HasCarrierNo() bool`

HasCarrierNo returns a boolean if a field has been set.

### GetCntrNo

`func (o *TrackingEventBillOfLading) GetCntrNo() string`

GetCntrNo returns the CntrNo field if non-nil, zero value otherwise.

### GetCntrNoOk

`func (o *TrackingEventBillOfLading) GetCntrNoOk() (*string, bool)`

GetCntrNoOk returns a tuple with the CntrNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCntrNo

`func (o *TrackingEventBillOfLading) SetCntrNo(v string)`

SetCntrNo sets CntrNo field to given value.

### HasCntrNo

`func (o *TrackingEventBillOfLading) HasCntrNo() bool`

HasCntrNo returns a boolean if a field has been set.

### GetCreated

`func (o *TrackingEventBillOfLading) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TrackingEventBillOfLading) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TrackingEventBillOfLading) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TrackingEventBillOfLading) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *TrackingEventBillOfLading) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackingEventBillOfLading) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackingEventBillOfLading) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrackingEventBillOfLading) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlaceOfDelivery

`func (o *TrackingEventBillOfLading) GetPlaceOfDelivery() string`

GetPlaceOfDelivery returns the PlaceOfDelivery field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryOk

`func (o *TrackingEventBillOfLading) GetPlaceOfDeliveryOk() (*string, bool)`

GetPlaceOfDeliveryOk returns a tuple with the PlaceOfDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDelivery

`func (o *TrackingEventBillOfLading) SetPlaceOfDelivery(v string)`

SetPlaceOfDelivery sets PlaceOfDelivery field to given value.

### HasPlaceOfDelivery

`func (o *TrackingEventBillOfLading) HasPlaceOfDelivery() bool`

HasPlaceOfDelivery returns a boolean if a field has been set.

### GetPlaceOfDeliveryName

`func (o *TrackingEventBillOfLading) GetPlaceOfDeliveryName() string`

GetPlaceOfDeliveryName returns the PlaceOfDeliveryName field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryNameOk

`func (o *TrackingEventBillOfLading) GetPlaceOfDeliveryNameOk() (*string, bool)`

GetPlaceOfDeliveryNameOk returns a tuple with the PlaceOfDeliveryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDeliveryName

`func (o *TrackingEventBillOfLading) SetPlaceOfDeliveryName(v string)`

SetPlaceOfDeliveryName sets PlaceOfDeliveryName field to given value.

### HasPlaceOfDeliveryName

`func (o *TrackingEventBillOfLading) HasPlaceOfDeliveryName() bool`

HasPlaceOfDeliveryName returns a boolean if a field has been set.

### GetPlaceOfReceipt

`func (o *TrackingEventBillOfLading) GetPlaceOfReceipt() string`

GetPlaceOfReceipt returns the PlaceOfReceipt field if non-nil, zero value otherwise.

### GetPlaceOfReceiptOk

`func (o *TrackingEventBillOfLading) GetPlaceOfReceiptOk() (*string, bool)`

GetPlaceOfReceiptOk returns a tuple with the PlaceOfReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfReceipt

`func (o *TrackingEventBillOfLading) SetPlaceOfReceipt(v string)`

SetPlaceOfReceipt sets PlaceOfReceipt field to given value.

### HasPlaceOfReceipt

`func (o *TrackingEventBillOfLading) HasPlaceOfReceipt() bool`

HasPlaceOfReceipt returns a boolean if a field has been set.

### GetPlaceOfReceiptName

`func (o *TrackingEventBillOfLading) GetPlaceOfReceiptName() string`

GetPlaceOfReceiptName returns the PlaceOfReceiptName field if non-nil, zero value otherwise.

### GetPlaceOfReceiptNameOk

`func (o *TrackingEventBillOfLading) GetPlaceOfReceiptNameOk() (*string, bool)`

GetPlaceOfReceiptNameOk returns a tuple with the PlaceOfReceiptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfReceiptName

`func (o *TrackingEventBillOfLading) SetPlaceOfReceiptName(v string)`

SetPlaceOfReceiptName sets PlaceOfReceiptName field to given value.

### HasPlaceOfReceiptName

`func (o *TrackingEventBillOfLading) HasPlaceOfReceiptName() bool`

HasPlaceOfReceiptName returns a boolean if a field has been set.

### GetPod

`func (o *TrackingEventBillOfLading) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *TrackingEventBillOfLading) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *TrackingEventBillOfLading) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *TrackingEventBillOfLading) HasPod() bool`

HasPod returns a boolean if a field has been set.

### GetPodActualArrivalLt

`func (o *TrackingEventBillOfLading) GetPodActualArrivalLt() string`

GetPodActualArrivalLt returns the PodActualArrivalLt field if non-nil, zero value otherwise.

### GetPodActualArrivalLtOk

`func (o *TrackingEventBillOfLading) GetPodActualArrivalLtOk() (*string, bool)`

GetPodActualArrivalLtOk returns a tuple with the PodActualArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualArrivalLt

`func (o *TrackingEventBillOfLading) SetPodActualArrivalLt(v string)`

SetPodActualArrivalLt sets PodActualArrivalLt field to given value.

### HasPodActualArrivalLt

`func (o *TrackingEventBillOfLading) HasPodActualArrivalLt() bool`

HasPodActualArrivalLt returns a boolean if a field has been set.

### GetPodActualArrivalUtc

`func (o *TrackingEventBillOfLading) GetPodActualArrivalUtc() string`

GetPodActualArrivalUtc returns the PodActualArrivalUtc field if non-nil, zero value otherwise.

### GetPodActualArrivalUtcOk

`func (o *TrackingEventBillOfLading) GetPodActualArrivalUtcOk() (*string, bool)`

GetPodActualArrivalUtcOk returns a tuple with the PodActualArrivalUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualArrivalUtc

`func (o *TrackingEventBillOfLading) SetPodActualArrivalUtc(v string)`

SetPodActualArrivalUtc sets PodActualArrivalUtc field to given value.

### HasPodActualArrivalUtc

`func (o *TrackingEventBillOfLading) HasPodActualArrivalUtc() bool`

HasPodActualArrivalUtc returns a boolean if a field has been set.

### GetPodActualArrivalLtFromAis

`func (o *TrackingEventBillOfLading) GetPodActualArrivalLtFromAis() string`

GetPodActualArrivalLtFromAis returns the PodActualArrivalLtFromAis field if non-nil, zero value otherwise.

### GetPodActualArrivalLtFromAisOk

`func (o *TrackingEventBillOfLading) GetPodActualArrivalLtFromAisOk() (*string, bool)`

GetPodActualArrivalLtFromAisOk returns a tuple with the PodActualArrivalLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualArrivalLtFromAis

`func (o *TrackingEventBillOfLading) SetPodActualArrivalLtFromAis(v string)`

SetPodActualArrivalLtFromAis sets PodActualArrivalLtFromAis field to given value.

### HasPodActualArrivalLtFromAis

`func (o *TrackingEventBillOfLading) HasPodActualArrivalLtFromAis() bool`

HasPodActualArrivalLtFromAis returns a boolean if a field has been set.

### GetPodActualDepartureLtFromAis

`func (o *TrackingEventBillOfLading) GetPodActualDepartureLtFromAis() string`

GetPodActualDepartureLtFromAis returns the PodActualDepartureLtFromAis field if non-nil, zero value otherwise.

### GetPodActualDepartureLtFromAisOk

`func (o *TrackingEventBillOfLading) GetPodActualDepartureLtFromAisOk() (*string, bool)`

GetPodActualDepartureLtFromAisOk returns a tuple with the PodActualDepartureLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualDepartureLtFromAis

`func (o *TrackingEventBillOfLading) SetPodActualDepartureLtFromAis(v string)`

SetPodActualDepartureLtFromAis sets PodActualDepartureLtFromAis field to given value.

### HasPodActualDepartureLtFromAis

`func (o *TrackingEventBillOfLading) HasPodActualDepartureLtFromAis() bool`

HasPodActualDepartureLtFromAis returns a boolean if a field has been set.

### GetPodActualDischargeLt

`func (o *TrackingEventBillOfLading) GetPodActualDischargeLt() string`

GetPodActualDischargeLt returns the PodActualDischargeLt field if non-nil, zero value otherwise.

### GetPodActualDischargeLtOk

`func (o *TrackingEventBillOfLading) GetPodActualDischargeLtOk() (*string, bool)`

GetPodActualDischargeLtOk returns a tuple with the PodActualDischargeLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualDischargeLt

`func (o *TrackingEventBillOfLading) SetPodActualDischargeLt(v string)`

SetPodActualDischargeLt sets PodActualDischargeLt field to given value.

### HasPodActualDischargeLt

`func (o *TrackingEventBillOfLading) HasPodActualDischargeLt() bool`

HasPodActualDischargeLt returns a boolean if a field has been set.

### GetPodName

`func (o *TrackingEventBillOfLading) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *TrackingEventBillOfLading) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *TrackingEventBillOfLading) SetPodName(v string)`

SetPodName sets PodName field to given value.

### HasPodName

`func (o *TrackingEventBillOfLading) HasPodName() bool`

HasPodName returns a boolean if a field has been set.

### GetPodPredictedArrivalLt

`func (o *TrackingEventBillOfLading) GetPodPredictedArrivalLt() string`

GetPodPredictedArrivalLt returns the PodPredictedArrivalLt field if non-nil, zero value otherwise.

### GetPodPredictedArrivalLtOk

`func (o *TrackingEventBillOfLading) GetPodPredictedArrivalLtOk() (*string, bool)`

GetPodPredictedArrivalLtOk returns a tuple with the PodPredictedArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPredictedArrivalLt

`func (o *TrackingEventBillOfLading) SetPodPredictedArrivalLt(v string)`

SetPodPredictedArrivalLt sets PodPredictedArrivalLt field to given value.

### HasPodPredictedArrivalLt

`func (o *TrackingEventBillOfLading) HasPodPredictedArrivalLt() bool`

HasPodPredictedArrivalLt returns a boolean if a field has been set.

### GetPodPredictedArrivalUtc

`func (o *TrackingEventBillOfLading) GetPodPredictedArrivalUtc() string`

GetPodPredictedArrivalUtc returns the PodPredictedArrivalUtc field if non-nil, zero value otherwise.

### GetPodPredictedArrivalUtcOk

`func (o *TrackingEventBillOfLading) GetPodPredictedArrivalUtcOk() (*string, bool)`

GetPodPredictedArrivalUtcOk returns a tuple with the PodPredictedArrivalUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPredictedArrivalUtc

`func (o *TrackingEventBillOfLading) SetPodPredictedArrivalUtc(v string)`

SetPodPredictedArrivalUtc sets PodPredictedArrivalUtc field to given value.

### HasPodPredictedArrivalUtc

`func (o *TrackingEventBillOfLading) HasPodPredictedArrivalUtc() bool`

HasPodPredictedArrivalUtc returns a boolean if a field has been set.

### GetPodPredictedDepartureLt

`func (o *TrackingEventBillOfLading) GetPodPredictedDepartureLt() string`

GetPodPredictedDepartureLt returns the PodPredictedDepartureLt field if non-nil, zero value otherwise.

### GetPodPredictedDepartureLtOk

`func (o *TrackingEventBillOfLading) GetPodPredictedDepartureLtOk() (*string, bool)`

GetPodPredictedDepartureLtOk returns a tuple with the PodPredictedDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPredictedDepartureLt

`func (o *TrackingEventBillOfLading) SetPodPredictedDepartureLt(v string)`

SetPodPredictedDepartureLt sets PodPredictedDepartureLt field to given value.

### HasPodPredictedDepartureLt

`func (o *TrackingEventBillOfLading) HasPodPredictedDepartureLt() bool`

HasPodPredictedDepartureLt returns a boolean if a field has been set.

### GetPodScheduledArrivalLt

`func (o *TrackingEventBillOfLading) GetPodScheduledArrivalLt() string`

GetPodScheduledArrivalLt returns the PodScheduledArrivalLt field if non-nil, zero value otherwise.

### GetPodScheduledArrivalLtOk

`func (o *TrackingEventBillOfLading) GetPodScheduledArrivalLtOk() (*string, bool)`

GetPodScheduledArrivalLtOk returns a tuple with the PodScheduledArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledArrivalLt

`func (o *TrackingEventBillOfLading) SetPodScheduledArrivalLt(v string)`

SetPodScheduledArrivalLt sets PodScheduledArrivalLt field to given value.

### HasPodScheduledArrivalLt

`func (o *TrackingEventBillOfLading) HasPodScheduledArrivalLt() bool`

HasPodScheduledArrivalLt returns a boolean if a field has been set.

### GetPodScheduledArrivalLtFromSchedule

`func (o *TrackingEventBillOfLading) GetPodScheduledArrivalLtFromSchedule() string`

GetPodScheduledArrivalLtFromSchedule returns the PodScheduledArrivalLtFromSchedule field if non-nil, zero value otherwise.

### GetPodScheduledArrivalLtFromScheduleOk

`func (o *TrackingEventBillOfLading) GetPodScheduledArrivalLtFromScheduleOk() (*string, bool)`

GetPodScheduledArrivalLtFromScheduleOk returns a tuple with the PodScheduledArrivalLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledArrivalLtFromSchedule

`func (o *TrackingEventBillOfLading) SetPodScheduledArrivalLtFromSchedule(v string)`

SetPodScheduledArrivalLtFromSchedule sets PodScheduledArrivalLtFromSchedule field to given value.

### HasPodScheduledArrivalLtFromSchedule

`func (o *TrackingEventBillOfLading) HasPodScheduledArrivalLtFromSchedule() bool`

HasPodScheduledArrivalLtFromSchedule returns a boolean if a field has been set.

### GetPodScheduledDepartureLtFromSchedule

`func (o *TrackingEventBillOfLading) GetPodScheduledDepartureLtFromSchedule() string`

GetPodScheduledDepartureLtFromSchedule returns the PodScheduledDepartureLtFromSchedule field if non-nil, zero value otherwise.

### GetPodScheduledDepartureLtFromScheduleOk

`func (o *TrackingEventBillOfLading) GetPodScheduledDepartureLtFromScheduleOk() (*string, bool)`

GetPodScheduledDepartureLtFromScheduleOk returns a tuple with the PodScheduledDepartureLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledDepartureLtFromSchedule

`func (o *TrackingEventBillOfLading) SetPodScheduledDepartureLtFromSchedule(v string)`

SetPodScheduledDepartureLtFromSchedule sets PodScheduledDepartureLtFromSchedule field to given value.

### HasPodScheduledDepartureLtFromSchedule

`func (o *TrackingEventBillOfLading) HasPodScheduledDepartureLtFromSchedule() bool`

HasPodScheduledDepartureLtFromSchedule returns a boolean if a field has been set.

### GetPodScheduledDischargeLt

`func (o *TrackingEventBillOfLading) GetPodScheduledDischargeLt() string`

GetPodScheduledDischargeLt returns the PodScheduledDischargeLt field if non-nil, zero value otherwise.

### GetPodScheduledDischargeLtOk

`func (o *TrackingEventBillOfLading) GetPodScheduledDischargeLtOk() (*string, bool)`

GetPodScheduledDischargeLtOk returns a tuple with the PodScheduledDischargeLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledDischargeLt

`func (o *TrackingEventBillOfLading) SetPodScheduledDischargeLt(v string)`

SetPodScheduledDischargeLt sets PodScheduledDischargeLt field to given value.

### HasPodScheduledDischargeLt

`func (o *TrackingEventBillOfLading) HasPodScheduledDischargeLt() bool`

HasPodScheduledDischargeLt returns a boolean if a field has been set.

### GetPol

`func (o *TrackingEventBillOfLading) GetPol() string`

GetPol returns the Pol field if non-nil, zero value otherwise.

### GetPolOk

`func (o *TrackingEventBillOfLading) GetPolOk() (*string, bool)`

GetPolOk returns a tuple with the Pol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPol

`func (o *TrackingEventBillOfLading) SetPol(v string)`

SetPol sets Pol field to given value.

### HasPol

`func (o *TrackingEventBillOfLading) HasPol() bool`

HasPol returns a boolean if a field has been set.

### GetPolActualArrivalLt

`func (o *TrackingEventBillOfLading) GetPolActualArrivalLt() string`

GetPolActualArrivalLt returns the PolActualArrivalLt field if non-nil, zero value otherwise.

### GetPolActualArrivalLtOk

`func (o *TrackingEventBillOfLading) GetPolActualArrivalLtOk() (*string, bool)`

GetPolActualArrivalLtOk returns a tuple with the PolActualArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualArrivalLt

`func (o *TrackingEventBillOfLading) SetPolActualArrivalLt(v string)`

SetPolActualArrivalLt sets PolActualArrivalLt field to given value.

### HasPolActualArrivalLt

`func (o *TrackingEventBillOfLading) HasPolActualArrivalLt() bool`

HasPolActualArrivalLt returns a boolean if a field has been set.

### GetPolActualArrivalLtFromAis

`func (o *TrackingEventBillOfLading) GetPolActualArrivalLtFromAis() string`

GetPolActualArrivalLtFromAis returns the PolActualArrivalLtFromAis field if non-nil, zero value otherwise.

### GetPolActualArrivalLtFromAisOk

`func (o *TrackingEventBillOfLading) GetPolActualArrivalLtFromAisOk() (*string, bool)`

GetPolActualArrivalLtFromAisOk returns a tuple with the PolActualArrivalLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualArrivalLtFromAis

`func (o *TrackingEventBillOfLading) SetPolActualArrivalLtFromAis(v string)`

SetPolActualArrivalLtFromAis sets PolActualArrivalLtFromAis field to given value.

### HasPolActualArrivalLtFromAis

`func (o *TrackingEventBillOfLading) HasPolActualArrivalLtFromAis() bool`

HasPolActualArrivalLtFromAis returns a boolean if a field has been set.

### GetPolActualDepartureLt

`func (o *TrackingEventBillOfLading) GetPolActualDepartureLt() string`

GetPolActualDepartureLt returns the PolActualDepartureLt field if non-nil, zero value otherwise.

### GetPolActualDepartureLtOk

`func (o *TrackingEventBillOfLading) GetPolActualDepartureLtOk() (*string, bool)`

GetPolActualDepartureLtOk returns a tuple with the PolActualDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualDepartureLt

`func (o *TrackingEventBillOfLading) SetPolActualDepartureLt(v string)`

SetPolActualDepartureLt sets PolActualDepartureLt field to given value.

### HasPolActualDepartureLt

`func (o *TrackingEventBillOfLading) HasPolActualDepartureLt() bool`

HasPolActualDepartureLt returns a boolean if a field has been set.

### GetPolActualDepartureLtFromAis

`func (o *TrackingEventBillOfLading) GetPolActualDepartureLtFromAis() string`

GetPolActualDepartureLtFromAis returns the PolActualDepartureLtFromAis field if non-nil, zero value otherwise.

### GetPolActualDepartureLtFromAisOk

`func (o *TrackingEventBillOfLading) GetPolActualDepartureLtFromAisOk() (*string, bool)`

GetPolActualDepartureLtFromAisOk returns a tuple with the PolActualDepartureLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualDepartureLtFromAis

`func (o *TrackingEventBillOfLading) SetPolActualDepartureLtFromAis(v string)`

SetPolActualDepartureLtFromAis sets PolActualDepartureLtFromAis field to given value.

### HasPolActualDepartureLtFromAis

`func (o *TrackingEventBillOfLading) HasPolActualDepartureLtFromAis() bool`

HasPolActualDepartureLtFromAis returns a boolean if a field has been set.

### GetPolActualLoadingLt

`func (o *TrackingEventBillOfLading) GetPolActualLoadingLt() string`

GetPolActualLoadingLt returns the PolActualLoadingLt field if non-nil, zero value otherwise.

### GetPolActualLoadingLtOk

`func (o *TrackingEventBillOfLading) GetPolActualLoadingLtOk() (*string, bool)`

GetPolActualLoadingLtOk returns a tuple with the PolActualLoadingLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualLoadingLt

`func (o *TrackingEventBillOfLading) SetPolActualLoadingLt(v string)`

SetPolActualLoadingLt sets PolActualLoadingLt field to given value.

### HasPolActualLoadingLt

`func (o *TrackingEventBillOfLading) HasPolActualLoadingLt() bool`

HasPolActualLoadingLt returns a boolean if a field has been set.

### GetPolName

`func (o *TrackingEventBillOfLading) GetPolName() string`

GetPolName returns the PolName field if non-nil, zero value otherwise.

### GetPolNameOk

`func (o *TrackingEventBillOfLading) GetPolNameOk() (*string, bool)`

GetPolNameOk returns a tuple with the PolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolName

`func (o *TrackingEventBillOfLading) SetPolName(v string)`

SetPolName sets PolName field to given value.

### HasPolName

`func (o *TrackingEventBillOfLading) HasPolName() bool`

HasPolName returns a boolean if a field has been set.

### GetPolPredictedArrivalLt

`func (o *TrackingEventBillOfLading) GetPolPredictedArrivalLt() string`

GetPolPredictedArrivalLt returns the PolPredictedArrivalLt field if non-nil, zero value otherwise.

### GetPolPredictedArrivalLtOk

`func (o *TrackingEventBillOfLading) GetPolPredictedArrivalLtOk() (*string, bool)`

GetPolPredictedArrivalLtOk returns a tuple with the PolPredictedArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolPredictedArrivalLt

`func (o *TrackingEventBillOfLading) SetPolPredictedArrivalLt(v string)`

SetPolPredictedArrivalLt sets PolPredictedArrivalLt field to given value.

### HasPolPredictedArrivalLt

`func (o *TrackingEventBillOfLading) HasPolPredictedArrivalLt() bool`

HasPolPredictedArrivalLt returns a boolean if a field has been set.

### GetPolPredictedDepartureLt

`func (o *TrackingEventBillOfLading) GetPolPredictedDepartureLt() string`

GetPolPredictedDepartureLt returns the PolPredictedDepartureLt field if non-nil, zero value otherwise.

### GetPolPredictedDepartureLtOk

`func (o *TrackingEventBillOfLading) GetPolPredictedDepartureLtOk() (*string, bool)`

GetPolPredictedDepartureLtOk returns a tuple with the PolPredictedDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolPredictedDepartureLt

`func (o *TrackingEventBillOfLading) SetPolPredictedDepartureLt(v string)`

SetPolPredictedDepartureLt sets PolPredictedDepartureLt field to given value.

### HasPolPredictedDepartureLt

`func (o *TrackingEventBillOfLading) HasPolPredictedDepartureLt() bool`

HasPolPredictedDepartureLt returns a boolean if a field has been set.

### GetPolScheduledArrivalLtFromSchedule

`func (o *TrackingEventBillOfLading) GetPolScheduledArrivalLtFromSchedule() string`

GetPolScheduledArrivalLtFromSchedule returns the PolScheduledArrivalLtFromSchedule field if non-nil, zero value otherwise.

### GetPolScheduledArrivalLtFromScheduleOk

`func (o *TrackingEventBillOfLading) GetPolScheduledArrivalLtFromScheduleOk() (*string, bool)`

GetPolScheduledArrivalLtFromScheduleOk returns a tuple with the PolScheduledArrivalLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledArrivalLtFromSchedule

`func (o *TrackingEventBillOfLading) SetPolScheduledArrivalLtFromSchedule(v string)`

SetPolScheduledArrivalLtFromSchedule sets PolScheduledArrivalLtFromSchedule field to given value.

### HasPolScheduledArrivalLtFromSchedule

`func (o *TrackingEventBillOfLading) HasPolScheduledArrivalLtFromSchedule() bool`

HasPolScheduledArrivalLtFromSchedule returns a boolean if a field has been set.

### GetPolScheduledDepartureLt

`func (o *TrackingEventBillOfLading) GetPolScheduledDepartureLt() string`

GetPolScheduledDepartureLt returns the PolScheduledDepartureLt field if non-nil, zero value otherwise.

### GetPolScheduledDepartureLtOk

`func (o *TrackingEventBillOfLading) GetPolScheduledDepartureLtOk() (*string, bool)`

GetPolScheduledDepartureLtOk returns a tuple with the PolScheduledDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledDepartureLt

`func (o *TrackingEventBillOfLading) SetPolScheduledDepartureLt(v string)`

SetPolScheduledDepartureLt sets PolScheduledDepartureLt field to given value.

### HasPolScheduledDepartureLt

`func (o *TrackingEventBillOfLading) HasPolScheduledDepartureLt() bool`

HasPolScheduledDepartureLt returns a boolean if a field has been set.

### GetPolScheduledDepartureLtFromSchedule

`func (o *TrackingEventBillOfLading) GetPolScheduledDepartureLtFromSchedule() string`

GetPolScheduledDepartureLtFromSchedule returns the PolScheduledDepartureLtFromSchedule field if non-nil, zero value otherwise.

### GetPolScheduledDepartureLtFromScheduleOk

`func (o *TrackingEventBillOfLading) GetPolScheduledDepartureLtFromScheduleOk() (*string, bool)`

GetPolScheduledDepartureLtFromScheduleOk returns a tuple with the PolScheduledDepartureLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledDepartureLtFromSchedule

`func (o *TrackingEventBillOfLading) SetPolScheduledDepartureLtFromSchedule(v string)`

SetPolScheduledDepartureLtFromSchedule sets PolScheduledDepartureLtFromSchedule field to given value.

### HasPolScheduledDepartureLtFromSchedule

`func (o *TrackingEventBillOfLading) HasPolScheduledDepartureLtFromSchedule() bool`

HasPolScheduledDepartureLtFromSchedule returns a boolean if a field has been set.

### GetPolScheduledLoadingLt

`func (o *TrackingEventBillOfLading) GetPolScheduledLoadingLt() string`

GetPolScheduledLoadingLt returns the PolScheduledLoadingLt field if non-nil, zero value otherwise.

### GetPolScheduledLoadingLtOk

`func (o *TrackingEventBillOfLading) GetPolScheduledLoadingLtOk() (*string, bool)`

GetPolScheduledLoadingLtOk returns a tuple with the PolScheduledLoadingLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledLoadingLt

`func (o *TrackingEventBillOfLading) SetPolScheduledLoadingLt(v string)`

SetPolScheduledLoadingLt sets PolScheduledLoadingLt field to given value.

### HasPolScheduledLoadingLt

`func (o *TrackingEventBillOfLading) HasPolScheduledLoadingLt() bool`

HasPolScheduledLoadingLt returns a boolean if a field has been set.

### GetUpdated

`func (o *TrackingEventBillOfLading) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TrackingEventBillOfLading) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TrackingEventBillOfLading) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *TrackingEventBillOfLading) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


