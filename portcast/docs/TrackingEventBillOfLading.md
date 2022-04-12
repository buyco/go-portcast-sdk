# TrackingEventBillOfLading

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualDeliveryTime** | Pointer to **string** |  | [optional] 
**ActualReceiptTime** | Pointer to **string** |  | [optional] 
**BlNo** | Pointer to **string** |  | [optional] 
**CarrierNo** | Pointer to **string** |  | [optional] 
**Cbm** | Pointer to **string** |  | [optional] 
**CntrNo** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kgs** | Pointer to **string** |  | [optional] 
**PlaceOfDelivery** | Pointer to **string** |  | [optional] 
**PlaceOfDeliveryName** | Pointer to **string** |  | [optional] 
**PlaceOfReceipt** | Pointer to **string** |  | [optional] 
**PlaceOfReceiptName** | Pointer to **string** |  | [optional] 
**Pod** | Pointer to **string** |  | [optional] 
**PodActualArrivalLt** | Pointer to **string** |  | [optional] 
**PodActualArrivalUtc** | Pointer to **string** |  | [optional] 
**PodActualDepartureLt** | Pointer to **string** |  | [optional] 
**PodActualDischargeLt** | Pointer to **string** |  | [optional] 
**PodName** | Pointer to **string** |  | [optional] 
**PodPredictedArrivalLt** | Pointer to **time.Time** |  | [optional] 
**PodPredictedArrivalUtc** | Pointer to **time.Time** |  | [optional] 
**PodPredictedDepartureLt** | Pointer to **time.Time** |  | [optional] 
**PodPredictedDischargeLt** | Pointer to **time.Time** |  | [optional] 
**PodScheduledArrivalLt** | Pointer to **time.Time** |  | [optional] 
**PodScheduledDepartureLt** | Pointer to **time.Time** |  | [optional] 
**PodScheduledDischargeLt** | Pointer to **time.Time** |  | [optional] 
**Pol** | Pointer to **string** |  | [optional] 
**PolActualArrivalLt** | Pointer to **time.Time** |  | [optional] 
**PolActualDepartureLt** | Pointer to **time.Time** |  | [optional] 
**PolActualLoadingLt** | Pointer to **time.Time** |  | [optional] 
**PolName** | Pointer to **string** |  | [optional] 
**PolPredictedArrivalLt** | Pointer to **string** |  | [optional] 
**PolPredictedDepartureLt** | Pointer to **string** |  | [optional] 
**PolPredictedLoadingLt** | Pointer to **string** |  | [optional] 
**PolScheduledArrivalLt** | Pointer to **string** |  | [optional] 
**PolScheduledDepartureLt** | Pointer to **string** |  | [optional] 
**PolScheduledLoadingLt** | Pointer to **string** |  | [optional] 
**ScheduledDeliveryTime** | Pointer to **string** |  | [optional] 
**ScheduledReceiptTime** | Pointer to **string** |  | [optional] 
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

### GetActualDeliveryTime

`func (o *TrackingEventBillOfLading) GetActualDeliveryTime() string`

GetActualDeliveryTime returns the ActualDeliveryTime field if non-nil, zero value otherwise.

### GetActualDeliveryTimeOk

`func (o *TrackingEventBillOfLading) GetActualDeliveryTimeOk() (*string, bool)`

GetActualDeliveryTimeOk returns a tuple with the ActualDeliveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDeliveryTime

`func (o *TrackingEventBillOfLading) SetActualDeliveryTime(v string)`

SetActualDeliveryTime sets ActualDeliveryTime field to given value.

### HasActualDeliveryTime

`func (o *TrackingEventBillOfLading) HasActualDeliveryTime() bool`

HasActualDeliveryTime returns a boolean if a field has been set.

### GetActualReceiptTime

`func (o *TrackingEventBillOfLading) GetActualReceiptTime() string`

GetActualReceiptTime returns the ActualReceiptTime field if non-nil, zero value otherwise.

### GetActualReceiptTimeOk

`func (o *TrackingEventBillOfLading) GetActualReceiptTimeOk() (*string, bool)`

GetActualReceiptTimeOk returns a tuple with the ActualReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualReceiptTime

`func (o *TrackingEventBillOfLading) SetActualReceiptTime(v string)`

SetActualReceiptTime sets ActualReceiptTime field to given value.

### HasActualReceiptTime

`func (o *TrackingEventBillOfLading) HasActualReceiptTime() bool`

HasActualReceiptTime returns a boolean if a field has been set.

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

### GetCbm

`func (o *TrackingEventBillOfLading) GetCbm() string`

GetCbm returns the Cbm field if non-nil, zero value otherwise.

### GetCbmOk

`func (o *TrackingEventBillOfLading) GetCbmOk() (*string, bool)`

GetCbmOk returns a tuple with the Cbm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbm

`func (o *TrackingEventBillOfLading) SetCbm(v string)`

SetCbm sets Cbm field to given value.

### HasCbm

`func (o *TrackingEventBillOfLading) HasCbm() bool`

HasCbm returns a boolean if a field has been set.

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

### GetKgs

`func (o *TrackingEventBillOfLading) GetKgs() string`

GetKgs returns the Kgs field if non-nil, zero value otherwise.

### GetKgsOk

`func (o *TrackingEventBillOfLading) GetKgsOk() (*string, bool)`

GetKgsOk returns a tuple with the Kgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKgs

`func (o *TrackingEventBillOfLading) SetKgs(v string)`

SetKgs sets Kgs field to given value.

### HasKgs

`func (o *TrackingEventBillOfLading) HasKgs() bool`

HasKgs returns a boolean if a field has been set.

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

### GetPodActualDepartureLt

`func (o *TrackingEventBillOfLading) GetPodActualDepartureLt() string`

GetPodActualDepartureLt returns the PodActualDepartureLt field if non-nil, zero value otherwise.

### GetPodActualDepartureLtOk

`func (o *TrackingEventBillOfLading) GetPodActualDepartureLtOk() (*string, bool)`

GetPodActualDepartureLtOk returns a tuple with the PodActualDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualDepartureLt

`func (o *TrackingEventBillOfLading) SetPodActualDepartureLt(v string)`

SetPodActualDepartureLt sets PodActualDepartureLt field to given value.

### HasPodActualDepartureLt

`func (o *TrackingEventBillOfLading) HasPodActualDepartureLt() bool`

HasPodActualDepartureLt returns a boolean if a field has been set.

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

`func (o *TrackingEventBillOfLading) GetPodPredictedArrivalLt() time.Time`

GetPodPredictedArrivalLt returns the PodPredictedArrivalLt field if non-nil, zero value otherwise.

### GetPodPredictedArrivalLtOk

`func (o *TrackingEventBillOfLading) GetPodPredictedArrivalLtOk() (*time.Time, bool)`

GetPodPredictedArrivalLtOk returns a tuple with the PodPredictedArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPredictedArrivalLt

`func (o *TrackingEventBillOfLading) SetPodPredictedArrivalLt(v time.Time)`

SetPodPredictedArrivalLt sets PodPredictedArrivalLt field to given value.

### HasPodPredictedArrivalLt

`func (o *TrackingEventBillOfLading) HasPodPredictedArrivalLt() bool`

HasPodPredictedArrivalLt returns a boolean if a field has been set.

### GetPodPredictedArrivalUtc

`func (o *TrackingEventBillOfLading) GetPodPredictedArrivalUtc() time.Time`

GetPodPredictedArrivalUtc returns the PodPredictedArrivalUtc field if non-nil, zero value otherwise.

### GetPodPredictedArrivalUtcOk

`func (o *TrackingEventBillOfLading) GetPodPredictedArrivalUtcOk() (*time.Time, bool)`

GetPodPredictedArrivalUtcOk returns a tuple with the PodPredictedArrivalUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPredictedArrivalUtc

`func (o *TrackingEventBillOfLading) SetPodPredictedArrivalUtc(v time.Time)`

SetPodPredictedArrivalUtc sets PodPredictedArrivalUtc field to given value.

### HasPodPredictedArrivalUtc

`func (o *TrackingEventBillOfLading) HasPodPredictedArrivalUtc() bool`

HasPodPredictedArrivalUtc returns a boolean if a field has been set.

### GetPodPredictedDepartureLt

`func (o *TrackingEventBillOfLading) GetPodPredictedDepartureLt() time.Time`

GetPodPredictedDepartureLt returns the PodPredictedDepartureLt field if non-nil, zero value otherwise.

### GetPodPredictedDepartureLtOk

`func (o *TrackingEventBillOfLading) GetPodPredictedDepartureLtOk() (*time.Time, bool)`

GetPodPredictedDepartureLtOk returns a tuple with the PodPredictedDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPredictedDepartureLt

`func (o *TrackingEventBillOfLading) SetPodPredictedDepartureLt(v time.Time)`

SetPodPredictedDepartureLt sets PodPredictedDepartureLt field to given value.

### HasPodPredictedDepartureLt

`func (o *TrackingEventBillOfLading) HasPodPredictedDepartureLt() bool`

HasPodPredictedDepartureLt returns a boolean if a field has been set.

### GetPodPredictedDischargeLt

`func (o *TrackingEventBillOfLading) GetPodPredictedDischargeLt() time.Time`

GetPodPredictedDischargeLt returns the PodPredictedDischargeLt field if non-nil, zero value otherwise.

### GetPodPredictedDischargeLtOk

`func (o *TrackingEventBillOfLading) GetPodPredictedDischargeLtOk() (*time.Time, bool)`

GetPodPredictedDischargeLtOk returns a tuple with the PodPredictedDischargeLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPredictedDischargeLt

`func (o *TrackingEventBillOfLading) SetPodPredictedDischargeLt(v time.Time)`

SetPodPredictedDischargeLt sets PodPredictedDischargeLt field to given value.

### HasPodPredictedDischargeLt

`func (o *TrackingEventBillOfLading) HasPodPredictedDischargeLt() bool`

HasPodPredictedDischargeLt returns a boolean if a field has been set.

### GetPodScheduledArrivalLt

`func (o *TrackingEventBillOfLading) GetPodScheduledArrivalLt() time.Time`

GetPodScheduledArrivalLt returns the PodScheduledArrivalLt field if non-nil, zero value otherwise.

### GetPodScheduledArrivalLtOk

`func (o *TrackingEventBillOfLading) GetPodScheduledArrivalLtOk() (*time.Time, bool)`

GetPodScheduledArrivalLtOk returns a tuple with the PodScheduledArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledArrivalLt

`func (o *TrackingEventBillOfLading) SetPodScheduledArrivalLt(v time.Time)`

SetPodScheduledArrivalLt sets PodScheduledArrivalLt field to given value.

### HasPodScheduledArrivalLt

`func (o *TrackingEventBillOfLading) HasPodScheduledArrivalLt() bool`

HasPodScheduledArrivalLt returns a boolean if a field has been set.

### GetPodScheduledDepartureLt

`func (o *TrackingEventBillOfLading) GetPodScheduledDepartureLt() time.Time`

GetPodScheduledDepartureLt returns the PodScheduledDepartureLt field if non-nil, zero value otherwise.

### GetPodScheduledDepartureLtOk

`func (o *TrackingEventBillOfLading) GetPodScheduledDepartureLtOk() (*time.Time, bool)`

GetPodScheduledDepartureLtOk returns a tuple with the PodScheduledDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledDepartureLt

`func (o *TrackingEventBillOfLading) SetPodScheduledDepartureLt(v time.Time)`

SetPodScheduledDepartureLt sets PodScheduledDepartureLt field to given value.

### HasPodScheduledDepartureLt

`func (o *TrackingEventBillOfLading) HasPodScheduledDepartureLt() bool`

HasPodScheduledDepartureLt returns a boolean if a field has been set.

### GetPodScheduledDischargeLt

`func (o *TrackingEventBillOfLading) GetPodScheduledDischargeLt() time.Time`

GetPodScheduledDischargeLt returns the PodScheduledDischargeLt field if non-nil, zero value otherwise.

### GetPodScheduledDischargeLtOk

`func (o *TrackingEventBillOfLading) GetPodScheduledDischargeLtOk() (*time.Time, bool)`

GetPodScheduledDischargeLtOk returns a tuple with the PodScheduledDischargeLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledDischargeLt

`func (o *TrackingEventBillOfLading) SetPodScheduledDischargeLt(v time.Time)`

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

`func (o *TrackingEventBillOfLading) GetPolActualArrivalLt() time.Time`

GetPolActualArrivalLt returns the PolActualArrivalLt field if non-nil, zero value otherwise.

### GetPolActualArrivalLtOk

`func (o *TrackingEventBillOfLading) GetPolActualArrivalLtOk() (*time.Time, bool)`

GetPolActualArrivalLtOk returns a tuple with the PolActualArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualArrivalLt

`func (o *TrackingEventBillOfLading) SetPolActualArrivalLt(v time.Time)`

SetPolActualArrivalLt sets PolActualArrivalLt field to given value.

### HasPolActualArrivalLt

`func (o *TrackingEventBillOfLading) HasPolActualArrivalLt() bool`

HasPolActualArrivalLt returns a boolean if a field has been set.

### GetPolActualDepartureLt

`func (o *TrackingEventBillOfLading) GetPolActualDepartureLt() time.Time`

GetPolActualDepartureLt returns the PolActualDepartureLt field if non-nil, zero value otherwise.

### GetPolActualDepartureLtOk

`func (o *TrackingEventBillOfLading) GetPolActualDepartureLtOk() (*time.Time, bool)`

GetPolActualDepartureLtOk returns a tuple with the PolActualDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualDepartureLt

`func (o *TrackingEventBillOfLading) SetPolActualDepartureLt(v time.Time)`

SetPolActualDepartureLt sets PolActualDepartureLt field to given value.

### HasPolActualDepartureLt

`func (o *TrackingEventBillOfLading) HasPolActualDepartureLt() bool`

HasPolActualDepartureLt returns a boolean if a field has been set.

### GetPolActualLoadingLt

`func (o *TrackingEventBillOfLading) GetPolActualLoadingLt() time.Time`

GetPolActualLoadingLt returns the PolActualLoadingLt field if non-nil, zero value otherwise.

### GetPolActualLoadingLtOk

`func (o *TrackingEventBillOfLading) GetPolActualLoadingLtOk() (*time.Time, bool)`

GetPolActualLoadingLtOk returns a tuple with the PolActualLoadingLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualLoadingLt

`func (o *TrackingEventBillOfLading) SetPolActualLoadingLt(v time.Time)`

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

### GetPolPredictedLoadingLt

`func (o *TrackingEventBillOfLading) GetPolPredictedLoadingLt() string`

GetPolPredictedLoadingLt returns the PolPredictedLoadingLt field if non-nil, zero value otherwise.

### GetPolPredictedLoadingLtOk

`func (o *TrackingEventBillOfLading) GetPolPredictedLoadingLtOk() (*string, bool)`

GetPolPredictedLoadingLtOk returns a tuple with the PolPredictedLoadingLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolPredictedLoadingLt

`func (o *TrackingEventBillOfLading) SetPolPredictedLoadingLt(v string)`

SetPolPredictedLoadingLt sets PolPredictedLoadingLt field to given value.

### HasPolPredictedLoadingLt

`func (o *TrackingEventBillOfLading) HasPolPredictedLoadingLt() bool`

HasPolPredictedLoadingLt returns a boolean if a field has been set.

### GetPolScheduledArrivalLt

`func (o *TrackingEventBillOfLading) GetPolScheduledArrivalLt() string`

GetPolScheduledArrivalLt returns the PolScheduledArrivalLt field if non-nil, zero value otherwise.

### GetPolScheduledArrivalLtOk

`func (o *TrackingEventBillOfLading) GetPolScheduledArrivalLtOk() (*string, bool)`

GetPolScheduledArrivalLtOk returns a tuple with the PolScheduledArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledArrivalLt

`func (o *TrackingEventBillOfLading) SetPolScheduledArrivalLt(v string)`

SetPolScheduledArrivalLt sets PolScheduledArrivalLt field to given value.

### HasPolScheduledArrivalLt

`func (o *TrackingEventBillOfLading) HasPolScheduledArrivalLt() bool`

HasPolScheduledArrivalLt returns a boolean if a field has been set.

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

### GetScheduledDeliveryTime

`func (o *TrackingEventBillOfLading) GetScheduledDeliveryTime() string`

GetScheduledDeliveryTime returns the ScheduledDeliveryTime field if non-nil, zero value otherwise.

### GetScheduledDeliveryTimeOk

`func (o *TrackingEventBillOfLading) GetScheduledDeliveryTimeOk() (*string, bool)`

GetScheduledDeliveryTimeOk returns a tuple with the ScheduledDeliveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDeliveryTime

`func (o *TrackingEventBillOfLading) SetScheduledDeliveryTime(v string)`

SetScheduledDeliveryTime sets ScheduledDeliveryTime field to given value.

### HasScheduledDeliveryTime

`func (o *TrackingEventBillOfLading) HasScheduledDeliveryTime() bool`

HasScheduledDeliveryTime returns a boolean if a field has been set.

### GetScheduledReceiptTime

`func (o *TrackingEventBillOfLading) GetScheduledReceiptTime() string`

GetScheduledReceiptTime returns the ScheduledReceiptTime field if non-nil, zero value otherwise.

### GetScheduledReceiptTimeOk

`func (o *TrackingEventBillOfLading) GetScheduledReceiptTimeOk() (*string, bool)`

GetScheduledReceiptTimeOk returns a tuple with the ScheduledReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledReceiptTime

`func (o *TrackingEventBillOfLading) SetScheduledReceiptTime(v string)`

SetScheduledReceiptTime sets ScheduledReceiptTime field to given value.

### HasScheduledReceiptTime

`func (o *TrackingEventBillOfLading) HasScheduledReceiptTime() bool`

HasScheduledReceiptTime returns a boolean if a field has been set.

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


