# BillOfLading

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlNo** | Pointer to **string** | Bill of Lading or Booking Number | [optional] 
**CarrierNo** | Pointer to **string** | Carrier SCAC Code | [optional] 
**CntrNo** | Pointer to **string** | Container Number | [optional] 
**Created** | Pointer to **time.Time** | Bill of Lading Object Created Date | [optional] 
**Id** | Pointer to **string** | Bill of Lading Object ID | [optional] 
**PlaceOfDelivery** | Pointer to **string** | UNLOCODE for the Location of Container Delivery (After the POD) | [optional] 
**PlaceOfDeliveryName** | Pointer to **string** | Name for the Location of Container Delivery (After the POD) | [optional] 
**PlaceOfReceipt** | Pointer to **string** | UNLOCODE for the Location of Container Receipt (Before the POL) | [optional] 
**PlaceOfReceiptName** | Pointer to **string** | Name for the Location of Container Receipt (Before the POL) | [optional] 
**Pod** | Pointer to **string** | UNLOCODE for the Port of Discharge | [optional] 
**PodActualArrivalLt** | Pointer to **NullableTime** | Actual Time of Arrival for the Final Vessel at the Port of Discharge (POD) - Local Time | [optional] 
**PodActualArrivalLtFromAis** | Pointer to **NullableTime** | Actual Time of Arrival for the Final Vessel at the Port of Discharge (POD) as per AIS Data- Local Time | [optional] 
**PodActualDepartureLtFromAis** | Pointer to **NullableTime** | Actual Time of Departure for the Final Vessel from the Port of Discharge (POD) as per AIS Data - Local Time | [optional] 
**PodActualDischargeLt** | Pointer to **NullableTime** | Actual Container Discharge Time from Final Vessel at the Port of Discharge (POD) - Local Time | [optional] 
**PodName** | Pointer to **string** | Port of Discharge Name | [optional] 
**PodPredictedArrivalLt** | Pointer to **NullableTime** | Portcast Predicted Time of Arrival for the Final Vessel at the Port of Discharge (POD) - Local Time [Most Reliable Source of ETA] | [optional] 
**PodPredictedDepartureLt** | Pointer to **NullableTime** | Portcast Predicted Time of Departure for the Final Vessel from the Port of Discharge (POD) - Local Time | [optional] 
**PodScheduledArrivalLt** | Pointer to **NullableTime** | Scheduled Time of Arrival for the Final Vessel at the Port of Discharge (POD) as per carrier T&amp;T - Local Time | [optional] 
**PodScheduledArrivalLtFirstSeen** | Pointer to **NullableTime** | First Recorded Scheduled Time of Arrival for the Final Vessel at the Port of Discharge (POD) as per carrier T&amp;T - Local Time | [optional] 
**PodScheduledArrivalLtFromSchedule** | Pointer to **NullableTime** | Scheduled Time of Arrival for the Final Vessel at the Port of Discharge (POD) as per vessel schedule - Local Time | [optional] 
**PodScheduledDepartureLtFromSchedule** | Pointer to **NullableTime** | Scheduled Time of Departure for the Final Vessel from the Port of Discharge (POD) as per vessel schedule - Local Time | [optional] 
**PodScheduledDischargeLt** | Pointer to **NullableTime** | Carrier Scheduled Container Discharge Time from Final Vessel at the Port of Discharge (POD) - Local Time | [optional] 
**PodTerminalName** | Pointer to **string** | Terminal Name for the Port of Discharge (POD) | [optional] 
**Pol** | Pointer to **string** | UNLOCODE for the Port of Loading | [optional] 
**PolActualArrivalLtFromAis** | Pointer to **NullableTime** | Actual Time of Arrival for the First Vessel at the Port of Loading (POL) as per AIS Data - Local Time | [optional] 
**PolActualDepartureLt** | Pointer to **NullableTime** | Actual Time of Departure for the First Vessel from the Port of Loading (POL) - Local Time | [optional] 
**PolActualDepartureLtFromAis** | Pointer to **NullableTime** | Actual Time of Departure for the First Vessel from the Port of Loading (POL) as per AIS Data - Local Time | [optional] 
**PolActualLoadingLt** | Pointer to **NullableTime** | Actual Container Loading Time on First Vessel at the Port of Loading (POL) - Local Time | [optional] 
**PolName** | Pointer to **string** | Port of Loading Name | [optional] 
**PolPredictedArrivalLt** | Pointer to **NullableTime** | Portcast Predicted Time of Arrival for the First Vessel at the Port of Loading (POL) - Local Time | [optional] 
**PolPredictedDepartureLt** | Pointer to **NullableTime** | Portcast Predicted Time of Departure for the First Vessel from the Port of Loading (POL) - Local Time [Most Reliable Source of ETD] | [optional] 
**PolScheduledArrivalLtFromSchedule** | Pointer to **NullableTime** | Scheduled Time of Arrival for the First Vessel at the Port of Loading (POL) as per Vessel Schedule - Local Time | [optional] 
**PolScheduledDepartureLt** | Pointer to **NullableTime** | Scheduled Time of Departure for the First Vessel from the Port of Loading (POL) as per Container T&amp;T - Local Time | [optional] 
**PolScheduledDepartureLtFirstSeen** | Pointer to **NullableTime** | First Recorded Scheduled Time of Departure for the First Vessel at the Port of Loading (POL) as per carrier T&amp;T - Local Time | [optional] 
**PolScheduledDepartureLtFromSchedule** | Pointer to **NullableTime** | Scheduled Time of Departure for the First Vessel from the Port of Loading (POL) as per Vessel Schedule - Local Time | [optional] 
**PolScheduledLoadingLt** | Pointer to **NullableTime** | Carrier Scheduled Container Loading Time on First Vessel at the Port of Loading (POL) - Local Time | [optional] 
**PolTerminalName** | Pointer to **string** | Terminal Name for the Port of Loading (POL) | [optional] 
**Updated** | Pointer to **time.Time** | Bill of Lading Object Updated Date | [optional] 

## Methods

### NewBillOfLading

`func NewBillOfLading() *BillOfLading`

NewBillOfLading instantiates a new BillOfLading object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillOfLadingWithDefaults

`func NewBillOfLadingWithDefaults() *BillOfLading`

NewBillOfLadingWithDefaults instantiates a new BillOfLading object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlNo

`func (o *BillOfLading) GetBlNo() string`

GetBlNo returns the BlNo field if non-nil, zero value otherwise.

### GetBlNoOk

`func (o *BillOfLading) GetBlNoOk() (*string, bool)`

GetBlNoOk returns a tuple with the BlNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlNo

`func (o *BillOfLading) SetBlNo(v string)`

SetBlNo sets BlNo field to given value.

### HasBlNo

`func (o *BillOfLading) HasBlNo() bool`

HasBlNo returns a boolean if a field has been set.

### GetCarrierNo

`func (o *BillOfLading) GetCarrierNo() string`

GetCarrierNo returns the CarrierNo field if non-nil, zero value otherwise.

### GetCarrierNoOk

`func (o *BillOfLading) GetCarrierNoOk() (*string, bool)`

GetCarrierNoOk returns a tuple with the CarrierNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierNo

`func (o *BillOfLading) SetCarrierNo(v string)`

SetCarrierNo sets CarrierNo field to given value.

### HasCarrierNo

`func (o *BillOfLading) HasCarrierNo() bool`

HasCarrierNo returns a boolean if a field has been set.

### GetCntrNo

`func (o *BillOfLading) GetCntrNo() string`

GetCntrNo returns the CntrNo field if non-nil, zero value otherwise.

### GetCntrNoOk

`func (o *BillOfLading) GetCntrNoOk() (*string, bool)`

GetCntrNoOk returns a tuple with the CntrNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCntrNo

`func (o *BillOfLading) SetCntrNo(v string)`

SetCntrNo sets CntrNo field to given value.

### HasCntrNo

`func (o *BillOfLading) HasCntrNo() bool`

HasCntrNo returns a boolean if a field has been set.

### GetCreated

`func (o *BillOfLading) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BillOfLading) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BillOfLading) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BillOfLading) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *BillOfLading) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillOfLading) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillOfLading) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillOfLading) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlaceOfDelivery

`func (o *BillOfLading) GetPlaceOfDelivery() string`

GetPlaceOfDelivery returns the PlaceOfDelivery field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryOk

`func (o *BillOfLading) GetPlaceOfDeliveryOk() (*string, bool)`

GetPlaceOfDeliveryOk returns a tuple with the PlaceOfDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDelivery

`func (o *BillOfLading) SetPlaceOfDelivery(v string)`

SetPlaceOfDelivery sets PlaceOfDelivery field to given value.

### HasPlaceOfDelivery

`func (o *BillOfLading) HasPlaceOfDelivery() bool`

HasPlaceOfDelivery returns a boolean if a field has been set.

### GetPlaceOfDeliveryName

`func (o *BillOfLading) GetPlaceOfDeliveryName() string`

GetPlaceOfDeliveryName returns the PlaceOfDeliveryName field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryNameOk

`func (o *BillOfLading) GetPlaceOfDeliveryNameOk() (*string, bool)`

GetPlaceOfDeliveryNameOk returns a tuple with the PlaceOfDeliveryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDeliveryName

`func (o *BillOfLading) SetPlaceOfDeliveryName(v string)`

SetPlaceOfDeliveryName sets PlaceOfDeliveryName field to given value.

### HasPlaceOfDeliveryName

`func (o *BillOfLading) HasPlaceOfDeliveryName() bool`

HasPlaceOfDeliveryName returns a boolean if a field has been set.

### GetPlaceOfReceipt

`func (o *BillOfLading) GetPlaceOfReceipt() string`

GetPlaceOfReceipt returns the PlaceOfReceipt field if non-nil, zero value otherwise.

### GetPlaceOfReceiptOk

`func (o *BillOfLading) GetPlaceOfReceiptOk() (*string, bool)`

GetPlaceOfReceiptOk returns a tuple with the PlaceOfReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfReceipt

`func (o *BillOfLading) SetPlaceOfReceipt(v string)`

SetPlaceOfReceipt sets PlaceOfReceipt field to given value.

### HasPlaceOfReceipt

`func (o *BillOfLading) HasPlaceOfReceipt() bool`

HasPlaceOfReceipt returns a boolean if a field has been set.

### GetPlaceOfReceiptName

`func (o *BillOfLading) GetPlaceOfReceiptName() string`

GetPlaceOfReceiptName returns the PlaceOfReceiptName field if non-nil, zero value otherwise.

### GetPlaceOfReceiptNameOk

`func (o *BillOfLading) GetPlaceOfReceiptNameOk() (*string, bool)`

GetPlaceOfReceiptNameOk returns a tuple with the PlaceOfReceiptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfReceiptName

`func (o *BillOfLading) SetPlaceOfReceiptName(v string)`

SetPlaceOfReceiptName sets PlaceOfReceiptName field to given value.

### HasPlaceOfReceiptName

`func (o *BillOfLading) HasPlaceOfReceiptName() bool`

HasPlaceOfReceiptName returns a boolean if a field has been set.

### GetPod

`func (o *BillOfLading) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *BillOfLading) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *BillOfLading) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *BillOfLading) HasPod() bool`

HasPod returns a boolean if a field has been set.

### GetPodActualArrivalLt

`func (o *BillOfLading) GetPodActualArrivalLt() time.Time`

GetPodActualArrivalLt returns the PodActualArrivalLt field if non-nil, zero value otherwise.

### GetPodActualArrivalLtOk

`func (o *BillOfLading) GetPodActualArrivalLtOk() (*time.Time, bool)`

GetPodActualArrivalLtOk returns a tuple with the PodActualArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualArrivalLt

`func (o *BillOfLading) SetPodActualArrivalLt(v time.Time)`

SetPodActualArrivalLt sets PodActualArrivalLt field to given value.

### HasPodActualArrivalLt

`func (o *BillOfLading) HasPodActualArrivalLt() bool`

HasPodActualArrivalLt returns a boolean if a field has been set.

### SetPodActualArrivalLtNil

`func (o *BillOfLading) SetPodActualArrivalLtNil(b bool)`

 SetPodActualArrivalLtNil sets the value for PodActualArrivalLt to be an explicit nil

### UnsetPodActualArrivalLt
`func (o *BillOfLading) UnsetPodActualArrivalLt()`

UnsetPodActualArrivalLt ensures that no value is present for PodActualArrivalLt, not even an explicit nil
### GetPodActualArrivalLtFromAis

`func (o *BillOfLading) GetPodActualArrivalLtFromAis() time.Time`

GetPodActualArrivalLtFromAis returns the PodActualArrivalLtFromAis field if non-nil, zero value otherwise.

### GetPodActualArrivalLtFromAisOk

`func (o *BillOfLading) GetPodActualArrivalLtFromAisOk() (*time.Time, bool)`

GetPodActualArrivalLtFromAisOk returns a tuple with the PodActualArrivalLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualArrivalLtFromAis

`func (o *BillOfLading) SetPodActualArrivalLtFromAis(v time.Time)`

SetPodActualArrivalLtFromAis sets PodActualArrivalLtFromAis field to given value.

### HasPodActualArrivalLtFromAis

`func (o *BillOfLading) HasPodActualArrivalLtFromAis() bool`

HasPodActualArrivalLtFromAis returns a boolean if a field has been set.

### SetPodActualArrivalLtFromAisNil

`func (o *BillOfLading) SetPodActualArrivalLtFromAisNil(b bool)`

 SetPodActualArrivalLtFromAisNil sets the value for PodActualArrivalLtFromAis to be an explicit nil

### UnsetPodActualArrivalLtFromAis
`func (o *BillOfLading) UnsetPodActualArrivalLtFromAis()`

UnsetPodActualArrivalLtFromAis ensures that no value is present for PodActualArrivalLtFromAis, not even an explicit nil
### GetPodActualDepartureLtFromAis

`func (o *BillOfLading) GetPodActualDepartureLtFromAis() time.Time`

GetPodActualDepartureLtFromAis returns the PodActualDepartureLtFromAis field if non-nil, zero value otherwise.

### GetPodActualDepartureLtFromAisOk

`func (o *BillOfLading) GetPodActualDepartureLtFromAisOk() (*time.Time, bool)`

GetPodActualDepartureLtFromAisOk returns a tuple with the PodActualDepartureLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualDepartureLtFromAis

`func (o *BillOfLading) SetPodActualDepartureLtFromAis(v time.Time)`

SetPodActualDepartureLtFromAis sets PodActualDepartureLtFromAis field to given value.

### HasPodActualDepartureLtFromAis

`func (o *BillOfLading) HasPodActualDepartureLtFromAis() bool`

HasPodActualDepartureLtFromAis returns a boolean if a field has been set.

### SetPodActualDepartureLtFromAisNil

`func (o *BillOfLading) SetPodActualDepartureLtFromAisNil(b bool)`

 SetPodActualDepartureLtFromAisNil sets the value for PodActualDepartureLtFromAis to be an explicit nil

### UnsetPodActualDepartureLtFromAis
`func (o *BillOfLading) UnsetPodActualDepartureLtFromAis()`

UnsetPodActualDepartureLtFromAis ensures that no value is present for PodActualDepartureLtFromAis, not even an explicit nil
### GetPodActualDischargeLt

`func (o *BillOfLading) GetPodActualDischargeLt() time.Time`

GetPodActualDischargeLt returns the PodActualDischargeLt field if non-nil, zero value otherwise.

### GetPodActualDischargeLtOk

`func (o *BillOfLading) GetPodActualDischargeLtOk() (*time.Time, bool)`

GetPodActualDischargeLtOk returns a tuple with the PodActualDischargeLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualDischargeLt

`func (o *BillOfLading) SetPodActualDischargeLt(v time.Time)`

SetPodActualDischargeLt sets PodActualDischargeLt field to given value.

### HasPodActualDischargeLt

`func (o *BillOfLading) HasPodActualDischargeLt() bool`

HasPodActualDischargeLt returns a boolean if a field has been set.

### SetPodActualDischargeLtNil

`func (o *BillOfLading) SetPodActualDischargeLtNil(b bool)`

 SetPodActualDischargeLtNil sets the value for PodActualDischargeLt to be an explicit nil

### UnsetPodActualDischargeLt
`func (o *BillOfLading) UnsetPodActualDischargeLt()`

UnsetPodActualDischargeLt ensures that no value is present for PodActualDischargeLt, not even an explicit nil
### GetPodName

`func (o *BillOfLading) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *BillOfLading) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *BillOfLading) SetPodName(v string)`

SetPodName sets PodName field to given value.

### HasPodName

`func (o *BillOfLading) HasPodName() bool`

HasPodName returns a boolean if a field has been set.

### GetPodPredictedArrivalLt

`func (o *BillOfLading) GetPodPredictedArrivalLt() time.Time`

GetPodPredictedArrivalLt returns the PodPredictedArrivalLt field if non-nil, zero value otherwise.

### GetPodPredictedArrivalLtOk

`func (o *BillOfLading) GetPodPredictedArrivalLtOk() (*time.Time, bool)`

GetPodPredictedArrivalLtOk returns a tuple with the PodPredictedArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPredictedArrivalLt

`func (o *BillOfLading) SetPodPredictedArrivalLt(v time.Time)`

SetPodPredictedArrivalLt sets PodPredictedArrivalLt field to given value.

### HasPodPredictedArrivalLt

`func (o *BillOfLading) HasPodPredictedArrivalLt() bool`

HasPodPredictedArrivalLt returns a boolean if a field has been set.

### SetPodPredictedArrivalLtNil

`func (o *BillOfLading) SetPodPredictedArrivalLtNil(b bool)`

 SetPodPredictedArrivalLtNil sets the value for PodPredictedArrivalLt to be an explicit nil

### UnsetPodPredictedArrivalLt
`func (o *BillOfLading) UnsetPodPredictedArrivalLt()`

UnsetPodPredictedArrivalLt ensures that no value is present for PodPredictedArrivalLt, not even an explicit nil
### GetPodPredictedDepartureLt

`func (o *BillOfLading) GetPodPredictedDepartureLt() time.Time`

GetPodPredictedDepartureLt returns the PodPredictedDepartureLt field if non-nil, zero value otherwise.

### GetPodPredictedDepartureLtOk

`func (o *BillOfLading) GetPodPredictedDepartureLtOk() (*time.Time, bool)`

GetPodPredictedDepartureLtOk returns a tuple with the PodPredictedDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPredictedDepartureLt

`func (o *BillOfLading) SetPodPredictedDepartureLt(v time.Time)`

SetPodPredictedDepartureLt sets PodPredictedDepartureLt field to given value.

### HasPodPredictedDepartureLt

`func (o *BillOfLading) HasPodPredictedDepartureLt() bool`

HasPodPredictedDepartureLt returns a boolean if a field has been set.

### SetPodPredictedDepartureLtNil

`func (o *BillOfLading) SetPodPredictedDepartureLtNil(b bool)`

 SetPodPredictedDepartureLtNil sets the value for PodPredictedDepartureLt to be an explicit nil

### UnsetPodPredictedDepartureLt
`func (o *BillOfLading) UnsetPodPredictedDepartureLt()`

UnsetPodPredictedDepartureLt ensures that no value is present for PodPredictedDepartureLt, not even an explicit nil
### GetPodScheduledArrivalLt

`func (o *BillOfLading) GetPodScheduledArrivalLt() time.Time`

GetPodScheduledArrivalLt returns the PodScheduledArrivalLt field if non-nil, zero value otherwise.

### GetPodScheduledArrivalLtOk

`func (o *BillOfLading) GetPodScheduledArrivalLtOk() (*time.Time, bool)`

GetPodScheduledArrivalLtOk returns a tuple with the PodScheduledArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledArrivalLt

`func (o *BillOfLading) SetPodScheduledArrivalLt(v time.Time)`

SetPodScheduledArrivalLt sets PodScheduledArrivalLt field to given value.

### HasPodScheduledArrivalLt

`func (o *BillOfLading) HasPodScheduledArrivalLt() bool`

HasPodScheduledArrivalLt returns a boolean if a field has been set.

### SetPodScheduledArrivalLtNil

`func (o *BillOfLading) SetPodScheduledArrivalLtNil(b bool)`

 SetPodScheduledArrivalLtNil sets the value for PodScheduledArrivalLt to be an explicit nil

### UnsetPodScheduledArrivalLt
`func (o *BillOfLading) UnsetPodScheduledArrivalLt()`

UnsetPodScheduledArrivalLt ensures that no value is present for PodScheduledArrivalLt, not even an explicit nil
### GetPodScheduledArrivalLtFirstSeen

`func (o *BillOfLading) GetPodScheduledArrivalLtFirstSeen() time.Time`

GetPodScheduledArrivalLtFirstSeen returns the PodScheduledArrivalLtFirstSeen field if non-nil, zero value otherwise.

### GetPodScheduledArrivalLtFirstSeenOk

`func (o *BillOfLading) GetPodScheduledArrivalLtFirstSeenOk() (*time.Time, bool)`

GetPodScheduledArrivalLtFirstSeenOk returns a tuple with the PodScheduledArrivalLtFirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledArrivalLtFirstSeen

`func (o *BillOfLading) SetPodScheduledArrivalLtFirstSeen(v time.Time)`

SetPodScheduledArrivalLtFirstSeen sets PodScheduledArrivalLtFirstSeen field to given value.

### HasPodScheduledArrivalLtFirstSeen

`func (o *BillOfLading) HasPodScheduledArrivalLtFirstSeen() bool`

HasPodScheduledArrivalLtFirstSeen returns a boolean if a field has been set.

### SetPodScheduledArrivalLtFirstSeenNil

`func (o *BillOfLading) SetPodScheduledArrivalLtFirstSeenNil(b bool)`

 SetPodScheduledArrivalLtFirstSeenNil sets the value for PodScheduledArrivalLtFirstSeen to be an explicit nil

### UnsetPodScheduledArrivalLtFirstSeen
`func (o *BillOfLading) UnsetPodScheduledArrivalLtFirstSeen()`

UnsetPodScheduledArrivalLtFirstSeen ensures that no value is present for PodScheduledArrivalLtFirstSeen, not even an explicit nil
### GetPodScheduledArrivalLtFromSchedule

`func (o *BillOfLading) GetPodScheduledArrivalLtFromSchedule() time.Time`

GetPodScheduledArrivalLtFromSchedule returns the PodScheduledArrivalLtFromSchedule field if non-nil, zero value otherwise.

### GetPodScheduledArrivalLtFromScheduleOk

`func (o *BillOfLading) GetPodScheduledArrivalLtFromScheduleOk() (*time.Time, bool)`

GetPodScheduledArrivalLtFromScheduleOk returns a tuple with the PodScheduledArrivalLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledArrivalLtFromSchedule

`func (o *BillOfLading) SetPodScheduledArrivalLtFromSchedule(v time.Time)`

SetPodScheduledArrivalLtFromSchedule sets PodScheduledArrivalLtFromSchedule field to given value.

### HasPodScheduledArrivalLtFromSchedule

`func (o *BillOfLading) HasPodScheduledArrivalLtFromSchedule() bool`

HasPodScheduledArrivalLtFromSchedule returns a boolean if a field has been set.

### SetPodScheduledArrivalLtFromScheduleNil

`func (o *BillOfLading) SetPodScheduledArrivalLtFromScheduleNil(b bool)`

 SetPodScheduledArrivalLtFromScheduleNil sets the value for PodScheduledArrivalLtFromSchedule to be an explicit nil

### UnsetPodScheduledArrivalLtFromSchedule
`func (o *BillOfLading) UnsetPodScheduledArrivalLtFromSchedule()`

UnsetPodScheduledArrivalLtFromSchedule ensures that no value is present for PodScheduledArrivalLtFromSchedule, not even an explicit nil
### GetPodScheduledDepartureLtFromSchedule

`func (o *BillOfLading) GetPodScheduledDepartureLtFromSchedule() time.Time`

GetPodScheduledDepartureLtFromSchedule returns the PodScheduledDepartureLtFromSchedule field if non-nil, zero value otherwise.

### GetPodScheduledDepartureLtFromScheduleOk

`func (o *BillOfLading) GetPodScheduledDepartureLtFromScheduleOk() (*time.Time, bool)`

GetPodScheduledDepartureLtFromScheduleOk returns a tuple with the PodScheduledDepartureLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledDepartureLtFromSchedule

`func (o *BillOfLading) SetPodScheduledDepartureLtFromSchedule(v time.Time)`

SetPodScheduledDepartureLtFromSchedule sets PodScheduledDepartureLtFromSchedule field to given value.

### HasPodScheduledDepartureLtFromSchedule

`func (o *BillOfLading) HasPodScheduledDepartureLtFromSchedule() bool`

HasPodScheduledDepartureLtFromSchedule returns a boolean if a field has been set.

### SetPodScheduledDepartureLtFromScheduleNil

`func (o *BillOfLading) SetPodScheduledDepartureLtFromScheduleNil(b bool)`

 SetPodScheduledDepartureLtFromScheduleNil sets the value for PodScheduledDepartureLtFromSchedule to be an explicit nil

### UnsetPodScheduledDepartureLtFromSchedule
`func (o *BillOfLading) UnsetPodScheduledDepartureLtFromSchedule()`

UnsetPodScheduledDepartureLtFromSchedule ensures that no value is present for PodScheduledDepartureLtFromSchedule, not even an explicit nil
### GetPodScheduledDischargeLt

`func (o *BillOfLading) GetPodScheduledDischargeLt() time.Time`

GetPodScheduledDischargeLt returns the PodScheduledDischargeLt field if non-nil, zero value otherwise.

### GetPodScheduledDischargeLtOk

`func (o *BillOfLading) GetPodScheduledDischargeLtOk() (*time.Time, bool)`

GetPodScheduledDischargeLtOk returns a tuple with the PodScheduledDischargeLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledDischargeLt

`func (o *BillOfLading) SetPodScheduledDischargeLt(v time.Time)`

SetPodScheduledDischargeLt sets PodScheduledDischargeLt field to given value.

### HasPodScheduledDischargeLt

`func (o *BillOfLading) HasPodScheduledDischargeLt() bool`

HasPodScheduledDischargeLt returns a boolean if a field has been set.

### SetPodScheduledDischargeLtNil

`func (o *BillOfLading) SetPodScheduledDischargeLtNil(b bool)`

 SetPodScheduledDischargeLtNil sets the value for PodScheduledDischargeLt to be an explicit nil

### UnsetPodScheduledDischargeLt
`func (o *BillOfLading) UnsetPodScheduledDischargeLt()`

UnsetPodScheduledDischargeLt ensures that no value is present for PodScheduledDischargeLt, not even an explicit nil
### GetPodTerminalName

`func (o *BillOfLading) GetPodTerminalName() string`

GetPodTerminalName returns the PodTerminalName field if non-nil, zero value otherwise.

### GetPodTerminalNameOk

`func (o *BillOfLading) GetPodTerminalNameOk() (*string, bool)`

GetPodTerminalNameOk returns a tuple with the PodTerminalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodTerminalName

`func (o *BillOfLading) SetPodTerminalName(v string)`

SetPodTerminalName sets PodTerminalName field to given value.

### HasPodTerminalName

`func (o *BillOfLading) HasPodTerminalName() bool`

HasPodTerminalName returns a boolean if a field has been set.

### GetPol

`func (o *BillOfLading) GetPol() string`

GetPol returns the Pol field if non-nil, zero value otherwise.

### GetPolOk

`func (o *BillOfLading) GetPolOk() (*string, bool)`

GetPolOk returns a tuple with the Pol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPol

`func (o *BillOfLading) SetPol(v string)`

SetPol sets Pol field to given value.

### HasPol

`func (o *BillOfLading) HasPol() bool`

HasPol returns a boolean if a field has been set.

### GetPolActualArrivalLtFromAis

`func (o *BillOfLading) GetPolActualArrivalLtFromAis() time.Time`

GetPolActualArrivalLtFromAis returns the PolActualArrivalLtFromAis field if non-nil, zero value otherwise.

### GetPolActualArrivalLtFromAisOk

`func (o *BillOfLading) GetPolActualArrivalLtFromAisOk() (*time.Time, bool)`

GetPolActualArrivalLtFromAisOk returns a tuple with the PolActualArrivalLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualArrivalLtFromAis

`func (o *BillOfLading) SetPolActualArrivalLtFromAis(v time.Time)`

SetPolActualArrivalLtFromAis sets PolActualArrivalLtFromAis field to given value.

### HasPolActualArrivalLtFromAis

`func (o *BillOfLading) HasPolActualArrivalLtFromAis() bool`

HasPolActualArrivalLtFromAis returns a boolean if a field has been set.

### SetPolActualArrivalLtFromAisNil

`func (o *BillOfLading) SetPolActualArrivalLtFromAisNil(b bool)`

 SetPolActualArrivalLtFromAisNil sets the value for PolActualArrivalLtFromAis to be an explicit nil

### UnsetPolActualArrivalLtFromAis
`func (o *BillOfLading) UnsetPolActualArrivalLtFromAis()`

UnsetPolActualArrivalLtFromAis ensures that no value is present for PolActualArrivalLtFromAis, not even an explicit nil
### GetPolActualDepartureLt

`func (o *BillOfLading) GetPolActualDepartureLt() time.Time`

GetPolActualDepartureLt returns the PolActualDepartureLt field if non-nil, zero value otherwise.

### GetPolActualDepartureLtOk

`func (o *BillOfLading) GetPolActualDepartureLtOk() (*time.Time, bool)`

GetPolActualDepartureLtOk returns a tuple with the PolActualDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualDepartureLt

`func (o *BillOfLading) SetPolActualDepartureLt(v time.Time)`

SetPolActualDepartureLt sets PolActualDepartureLt field to given value.

### HasPolActualDepartureLt

`func (o *BillOfLading) HasPolActualDepartureLt() bool`

HasPolActualDepartureLt returns a boolean if a field has been set.

### SetPolActualDepartureLtNil

`func (o *BillOfLading) SetPolActualDepartureLtNil(b bool)`

 SetPolActualDepartureLtNil sets the value for PolActualDepartureLt to be an explicit nil

### UnsetPolActualDepartureLt
`func (o *BillOfLading) UnsetPolActualDepartureLt()`

UnsetPolActualDepartureLt ensures that no value is present for PolActualDepartureLt, not even an explicit nil
### GetPolActualDepartureLtFromAis

`func (o *BillOfLading) GetPolActualDepartureLtFromAis() time.Time`

GetPolActualDepartureLtFromAis returns the PolActualDepartureLtFromAis field if non-nil, zero value otherwise.

### GetPolActualDepartureLtFromAisOk

`func (o *BillOfLading) GetPolActualDepartureLtFromAisOk() (*time.Time, bool)`

GetPolActualDepartureLtFromAisOk returns a tuple with the PolActualDepartureLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualDepartureLtFromAis

`func (o *BillOfLading) SetPolActualDepartureLtFromAis(v time.Time)`

SetPolActualDepartureLtFromAis sets PolActualDepartureLtFromAis field to given value.

### HasPolActualDepartureLtFromAis

`func (o *BillOfLading) HasPolActualDepartureLtFromAis() bool`

HasPolActualDepartureLtFromAis returns a boolean if a field has been set.

### SetPolActualDepartureLtFromAisNil

`func (o *BillOfLading) SetPolActualDepartureLtFromAisNil(b bool)`

 SetPolActualDepartureLtFromAisNil sets the value for PolActualDepartureLtFromAis to be an explicit nil

### UnsetPolActualDepartureLtFromAis
`func (o *BillOfLading) UnsetPolActualDepartureLtFromAis()`

UnsetPolActualDepartureLtFromAis ensures that no value is present for PolActualDepartureLtFromAis, not even an explicit nil
### GetPolActualLoadingLt

`func (o *BillOfLading) GetPolActualLoadingLt() time.Time`

GetPolActualLoadingLt returns the PolActualLoadingLt field if non-nil, zero value otherwise.

### GetPolActualLoadingLtOk

`func (o *BillOfLading) GetPolActualLoadingLtOk() (*time.Time, bool)`

GetPolActualLoadingLtOk returns a tuple with the PolActualLoadingLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualLoadingLt

`func (o *BillOfLading) SetPolActualLoadingLt(v time.Time)`

SetPolActualLoadingLt sets PolActualLoadingLt field to given value.

### HasPolActualLoadingLt

`func (o *BillOfLading) HasPolActualLoadingLt() bool`

HasPolActualLoadingLt returns a boolean if a field has been set.

### SetPolActualLoadingLtNil

`func (o *BillOfLading) SetPolActualLoadingLtNil(b bool)`

 SetPolActualLoadingLtNil sets the value for PolActualLoadingLt to be an explicit nil

### UnsetPolActualLoadingLt
`func (o *BillOfLading) UnsetPolActualLoadingLt()`

UnsetPolActualLoadingLt ensures that no value is present for PolActualLoadingLt, not even an explicit nil
### GetPolName

`func (o *BillOfLading) GetPolName() string`

GetPolName returns the PolName field if non-nil, zero value otherwise.

### GetPolNameOk

`func (o *BillOfLading) GetPolNameOk() (*string, bool)`

GetPolNameOk returns a tuple with the PolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolName

`func (o *BillOfLading) SetPolName(v string)`

SetPolName sets PolName field to given value.

### HasPolName

`func (o *BillOfLading) HasPolName() bool`

HasPolName returns a boolean if a field has been set.

### GetPolPredictedArrivalLt

`func (o *BillOfLading) GetPolPredictedArrivalLt() time.Time`

GetPolPredictedArrivalLt returns the PolPredictedArrivalLt field if non-nil, zero value otherwise.

### GetPolPredictedArrivalLtOk

`func (o *BillOfLading) GetPolPredictedArrivalLtOk() (*time.Time, bool)`

GetPolPredictedArrivalLtOk returns a tuple with the PolPredictedArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolPredictedArrivalLt

`func (o *BillOfLading) SetPolPredictedArrivalLt(v time.Time)`

SetPolPredictedArrivalLt sets PolPredictedArrivalLt field to given value.

### HasPolPredictedArrivalLt

`func (o *BillOfLading) HasPolPredictedArrivalLt() bool`

HasPolPredictedArrivalLt returns a boolean if a field has been set.

### SetPolPredictedArrivalLtNil

`func (o *BillOfLading) SetPolPredictedArrivalLtNil(b bool)`

 SetPolPredictedArrivalLtNil sets the value for PolPredictedArrivalLt to be an explicit nil

### UnsetPolPredictedArrivalLt
`func (o *BillOfLading) UnsetPolPredictedArrivalLt()`

UnsetPolPredictedArrivalLt ensures that no value is present for PolPredictedArrivalLt, not even an explicit nil
### GetPolPredictedDepartureLt

`func (o *BillOfLading) GetPolPredictedDepartureLt() time.Time`

GetPolPredictedDepartureLt returns the PolPredictedDepartureLt field if non-nil, zero value otherwise.

### GetPolPredictedDepartureLtOk

`func (o *BillOfLading) GetPolPredictedDepartureLtOk() (*time.Time, bool)`

GetPolPredictedDepartureLtOk returns a tuple with the PolPredictedDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolPredictedDepartureLt

`func (o *BillOfLading) SetPolPredictedDepartureLt(v time.Time)`

SetPolPredictedDepartureLt sets PolPredictedDepartureLt field to given value.

### HasPolPredictedDepartureLt

`func (o *BillOfLading) HasPolPredictedDepartureLt() bool`

HasPolPredictedDepartureLt returns a boolean if a field has been set.

### SetPolPredictedDepartureLtNil

`func (o *BillOfLading) SetPolPredictedDepartureLtNil(b bool)`

 SetPolPredictedDepartureLtNil sets the value for PolPredictedDepartureLt to be an explicit nil

### UnsetPolPredictedDepartureLt
`func (o *BillOfLading) UnsetPolPredictedDepartureLt()`

UnsetPolPredictedDepartureLt ensures that no value is present for PolPredictedDepartureLt, not even an explicit nil
### GetPolScheduledArrivalLtFromSchedule

`func (o *BillOfLading) GetPolScheduledArrivalLtFromSchedule() time.Time`

GetPolScheduledArrivalLtFromSchedule returns the PolScheduledArrivalLtFromSchedule field if non-nil, zero value otherwise.

### GetPolScheduledArrivalLtFromScheduleOk

`func (o *BillOfLading) GetPolScheduledArrivalLtFromScheduleOk() (*time.Time, bool)`

GetPolScheduledArrivalLtFromScheduleOk returns a tuple with the PolScheduledArrivalLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledArrivalLtFromSchedule

`func (o *BillOfLading) SetPolScheduledArrivalLtFromSchedule(v time.Time)`

SetPolScheduledArrivalLtFromSchedule sets PolScheduledArrivalLtFromSchedule field to given value.

### HasPolScheduledArrivalLtFromSchedule

`func (o *BillOfLading) HasPolScheduledArrivalLtFromSchedule() bool`

HasPolScheduledArrivalLtFromSchedule returns a boolean if a field has been set.

### SetPolScheduledArrivalLtFromScheduleNil

`func (o *BillOfLading) SetPolScheduledArrivalLtFromScheduleNil(b bool)`

 SetPolScheduledArrivalLtFromScheduleNil sets the value for PolScheduledArrivalLtFromSchedule to be an explicit nil

### UnsetPolScheduledArrivalLtFromSchedule
`func (o *BillOfLading) UnsetPolScheduledArrivalLtFromSchedule()`

UnsetPolScheduledArrivalLtFromSchedule ensures that no value is present for PolScheduledArrivalLtFromSchedule, not even an explicit nil
### GetPolScheduledDepartureLt

`func (o *BillOfLading) GetPolScheduledDepartureLt() time.Time`

GetPolScheduledDepartureLt returns the PolScheduledDepartureLt field if non-nil, zero value otherwise.

### GetPolScheduledDepartureLtOk

`func (o *BillOfLading) GetPolScheduledDepartureLtOk() (*time.Time, bool)`

GetPolScheduledDepartureLtOk returns a tuple with the PolScheduledDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledDepartureLt

`func (o *BillOfLading) SetPolScheduledDepartureLt(v time.Time)`

SetPolScheduledDepartureLt sets PolScheduledDepartureLt field to given value.

### HasPolScheduledDepartureLt

`func (o *BillOfLading) HasPolScheduledDepartureLt() bool`

HasPolScheduledDepartureLt returns a boolean if a field has been set.

### SetPolScheduledDepartureLtNil

`func (o *BillOfLading) SetPolScheduledDepartureLtNil(b bool)`

 SetPolScheduledDepartureLtNil sets the value for PolScheduledDepartureLt to be an explicit nil

### UnsetPolScheduledDepartureLt
`func (o *BillOfLading) UnsetPolScheduledDepartureLt()`

UnsetPolScheduledDepartureLt ensures that no value is present for PolScheduledDepartureLt, not even an explicit nil
### GetPolScheduledDepartureLtFirstSeen

`func (o *BillOfLading) GetPolScheduledDepartureLtFirstSeen() time.Time`

GetPolScheduledDepartureLtFirstSeen returns the PolScheduledDepartureLtFirstSeen field if non-nil, zero value otherwise.

### GetPolScheduledDepartureLtFirstSeenOk

`func (o *BillOfLading) GetPolScheduledDepartureLtFirstSeenOk() (*time.Time, bool)`

GetPolScheduledDepartureLtFirstSeenOk returns a tuple with the PolScheduledDepartureLtFirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledDepartureLtFirstSeen

`func (o *BillOfLading) SetPolScheduledDepartureLtFirstSeen(v time.Time)`

SetPolScheduledDepartureLtFirstSeen sets PolScheduledDepartureLtFirstSeen field to given value.

### HasPolScheduledDepartureLtFirstSeen

`func (o *BillOfLading) HasPolScheduledDepartureLtFirstSeen() bool`

HasPolScheduledDepartureLtFirstSeen returns a boolean if a field has been set.

### SetPolScheduledDepartureLtFirstSeenNil

`func (o *BillOfLading) SetPolScheduledDepartureLtFirstSeenNil(b bool)`

 SetPolScheduledDepartureLtFirstSeenNil sets the value for PolScheduledDepartureLtFirstSeen to be an explicit nil

### UnsetPolScheduledDepartureLtFirstSeen
`func (o *BillOfLading) UnsetPolScheduledDepartureLtFirstSeen()`

UnsetPolScheduledDepartureLtFirstSeen ensures that no value is present for PolScheduledDepartureLtFirstSeen, not even an explicit nil
### GetPolScheduledDepartureLtFromSchedule

`func (o *BillOfLading) GetPolScheduledDepartureLtFromSchedule() time.Time`

GetPolScheduledDepartureLtFromSchedule returns the PolScheduledDepartureLtFromSchedule field if non-nil, zero value otherwise.

### GetPolScheduledDepartureLtFromScheduleOk

`func (o *BillOfLading) GetPolScheduledDepartureLtFromScheduleOk() (*time.Time, bool)`

GetPolScheduledDepartureLtFromScheduleOk returns a tuple with the PolScheduledDepartureLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledDepartureLtFromSchedule

`func (o *BillOfLading) SetPolScheduledDepartureLtFromSchedule(v time.Time)`

SetPolScheduledDepartureLtFromSchedule sets PolScheduledDepartureLtFromSchedule field to given value.

### HasPolScheduledDepartureLtFromSchedule

`func (o *BillOfLading) HasPolScheduledDepartureLtFromSchedule() bool`

HasPolScheduledDepartureLtFromSchedule returns a boolean if a field has been set.

### SetPolScheduledDepartureLtFromScheduleNil

`func (o *BillOfLading) SetPolScheduledDepartureLtFromScheduleNil(b bool)`

 SetPolScheduledDepartureLtFromScheduleNil sets the value for PolScheduledDepartureLtFromSchedule to be an explicit nil

### UnsetPolScheduledDepartureLtFromSchedule
`func (o *BillOfLading) UnsetPolScheduledDepartureLtFromSchedule()`

UnsetPolScheduledDepartureLtFromSchedule ensures that no value is present for PolScheduledDepartureLtFromSchedule, not even an explicit nil
### GetPolScheduledLoadingLt

`func (o *BillOfLading) GetPolScheduledLoadingLt() time.Time`

GetPolScheduledLoadingLt returns the PolScheduledLoadingLt field if non-nil, zero value otherwise.

### GetPolScheduledLoadingLtOk

`func (o *BillOfLading) GetPolScheduledLoadingLtOk() (*time.Time, bool)`

GetPolScheduledLoadingLtOk returns a tuple with the PolScheduledLoadingLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledLoadingLt

`func (o *BillOfLading) SetPolScheduledLoadingLt(v time.Time)`

SetPolScheduledLoadingLt sets PolScheduledLoadingLt field to given value.

### HasPolScheduledLoadingLt

`func (o *BillOfLading) HasPolScheduledLoadingLt() bool`

HasPolScheduledLoadingLt returns a boolean if a field has been set.

### SetPolScheduledLoadingLtNil

`func (o *BillOfLading) SetPolScheduledLoadingLtNil(b bool)`

 SetPolScheduledLoadingLtNil sets the value for PolScheduledLoadingLt to be an explicit nil

### UnsetPolScheduledLoadingLt
`func (o *BillOfLading) UnsetPolScheduledLoadingLt()`

UnsetPolScheduledLoadingLt ensures that no value is present for PolScheduledLoadingLt, not even an explicit nil
### GetPolTerminalName

`func (o *BillOfLading) GetPolTerminalName() string`

GetPolTerminalName returns the PolTerminalName field if non-nil, zero value otherwise.

### GetPolTerminalNameOk

`func (o *BillOfLading) GetPolTerminalNameOk() (*string, bool)`

GetPolTerminalNameOk returns a tuple with the PolTerminalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolTerminalName

`func (o *BillOfLading) SetPolTerminalName(v string)`

SetPolTerminalName sets PolTerminalName field to given value.

### HasPolTerminalName

`func (o *BillOfLading) HasPolTerminalName() bool`

HasPolTerminalName returns a boolean if a field has been set.

### GetUpdated

`func (o *BillOfLading) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *BillOfLading) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *BillOfLading) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *BillOfLading) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


