# PortcastAPIBillOfLading

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlNo** | Pointer to **interface{}** | Bill of Lading or Booking Number | [optional] 
**CarrierNo** | Pointer to **interface{}** | Carrier SCAC Code | [optional] 
**CntrNo** | Pointer to **interface{}** | Container Number | [optional] 
**Created** | Pointer to **interface{}** | Bill of Lading Object Created Date | [optional] 
**Id** | Pointer to **interface{}** | Bill of Lading Object ID | [optional] 
**PlaceOfDelivery** | Pointer to **interface{}** | UNLOCODE for the Location of Container Delivery (After the POD) | [optional] 
**PlaceOfDeliveryName** | Pointer to **interface{}** | Name for the Location of Container Delivery (After the POD) | [optional] 
**PlaceOfReceipt** | Pointer to **interface{}** | UNLOCODE for the Location of Container Receipt (Before the POL) | [optional] 
**PlaceOfReceiptName** | Pointer to **interface{}** | Name for the Location of Container Receipt (Before the POL) | [optional] 
**Pod** | Pointer to **interface{}** | UNLOCODE for the Port of Discharge | [optional] 
**PodActualArrivalLt** | Pointer to **interface{}** | Actual Time of Arrival for the Final Vessel at the Port of Discharge (POD) - Local Time | [optional] 
**PodActualArrivalLtFromAis** | Pointer to **interface{}** | Actual Time of Arrival for the Final Vessel at the Port of Discharge (POD) as per AIS Data- Local Time | [optional] 
**PodActualDepartureLtFromAis** | Pointer to **interface{}** | Actual Time of Departure for the Final Vessel from the Port of Discharge (POD) as per AIS Data - Local Time | [optional] 
**PodActualDischargeLt** | Pointer to **interface{}** | Actual Container Discharge Time from Final Vessel at the Port of Discharge (POD) - Local Time | [optional] 
**PodName** | Pointer to **interface{}** | Port of Discharge Name | [optional] 
**PodPredictedArrivalLt** | Pointer to **interface{}** | Portcast Predicted Time of Arrival for the Final Vessel at the Port of Discharge (POD) - Local Time [Most Reliable Source of ETA] | [optional] 
**PodPredictedDepartureLt** | Pointer to **interface{}** | Portcast Predicted Time of Departure for the Final Vessel from the Port of Discharge (POD) - Local Time | [optional] 
**PodScheduledArrivalLt** | Pointer to **interface{}** | Scheduled Time of Arrival for the Final Vessel at the Port of Discharge (POD) as per carrier T&amp;T - Local Time | [optional] 
**PodScheduledArrivalLtFirstSeen** | Pointer to **interface{}** | First Recorded Scheduled Time of Arrival for the Final Vessel at the Port of Discharge (POD) as per carrier T&amp;T - Local Time | [optional] 
**PodScheduledArrivalLtFromSchedule** | Pointer to **interface{}** | Scheduled Time of Arrival for the Final Vessel at the Port of Discharge (POD) as per vessel schedule - Local Time | [optional] 
**PodScheduledDepartureLtFromSchedule** | Pointer to **interface{}** | Scheduled Time of Departure for the Final Vessel from the Port of Discharge (POD) as per vessel schedule - Local Time | [optional] 
**PodScheduledDischargeLt** | Pointer to **interface{}** | Carrier Scheduled Container Discharge Time from Final Vessel at the Port of Discharge (POD) - Local Time | [optional] 
**PodTerminalName** | Pointer to **interface{}** | Terminal Name for the Port of Discharge (POD) | [optional] 
**Pol** | Pointer to **interface{}** | UNLOCODE for the Port of Loading | [optional] 
**PolActualArrivalLtFromAis** | Pointer to **interface{}** | Actual Time of Arrival for the First Vessel at the Port of Loading (POL) as per AIS Data - Local Time | [optional] 
**PolActualDepartureLt** | Pointer to **interface{}** | Actual Time of Departure for the First Vessel from the Port of Loading (POL) - Local Time | [optional] 
**PolActualDepartureLtFromAis** | Pointer to **interface{}** | Actual Time of Departure for the First Vessel from the Port of Loading (POL) as per AIS Data - Local Time | [optional] 
**PolActualLoadingLt** | Pointer to **interface{}** | Actual Container Loading Time on First Vessel at the Port of Loading (POL) - Local Time | [optional] 
**PolName** | Pointer to **interface{}** | Port of Loading Name | [optional] 
**PolPredictedArrivalLt** | Pointer to **interface{}** | Portcast Predicted Time of Arrival for the First Vessel at the Port of Loading (POL) - Local Time | [optional] 
**PolPredictedDepartureLt** | Pointer to **interface{}** | Portcast Predicted Time of Departure for the First Vessel from the Port of Loading (POL) - Local Time [Most Reliable Source of ETD] | [optional] 
**PolScheduledArrivalLtFromSchedule** | Pointer to **interface{}** | Scheduled Time of Arrival for the First Vessel at the Port of Loading (POL) as per Vessel Schedule - Local Time | [optional] 
**PolScheduledDepartureLt** | Pointer to **interface{}** | Scheduled Time of Departure for the First Vessel from the Port of Loading (POL) as per Container T&amp;T - Local Time | [optional] 
**PolScheduledDepartureLtFirstSeen** | Pointer to **interface{}** | First Recorded Scheduled Time of Departure for the First Vessel at the Port of Loading (POL) as per carrier T&amp;T - Local Time | [optional] 
**PolScheduledDepartureLtFromSchedule** | Pointer to **interface{}** | Scheduled Time of Departure for the First Vessel from the Port of Loading (POL) as per Vessel Schedule - Local Time | [optional] 
**PolScheduledLoadingLt** | Pointer to **interface{}** | Carrier Scheduled Container Loading Time on First Vessel at the Port of Loading (POL) - Local Time | [optional] 
**PolTerminalName** | Pointer to **interface{}** | Terminal Name for the Port of Loading (POL) | [optional] 
**Updated** | Pointer to **interface{}** | Bill of Lading Object Updated Date | [optional] 

## Methods

### NewPortcastAPIBillOfLading

`func NewPortcastAPIBillOfLading() *PortcastAPIBillOfLading`

NewPortcastAPIBillOfLading instantiates a new PortcastAPIBillOfLading object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortcastAPIBillOfLadingWithDefaults

`func NewPortcastAPIBillOfLadingWithDefaults() *PortcastAPIBillOfLading`

NewPortcastAPIBillOfLadingWithDefaults instantiates a new PortcastAPIBillOfLading object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlNo

`func (o *PortcastAPIBillOfLading) GetBlNo() interface{}`

GetBlNo returns the BlNo field if non-nil, zero value otherwise.

### GetBlNoOk

`func (o *PortcastAPIBillOfLading) GetBlNoOk() (*interface{}, bool)`

GetBlNoOk returns a tuple with the BlNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlNo

`func (o *PortcastAPIBillOfLading) SetBlNo(v interface{})`

SetBlNo sets BlNo field to given value.

### HasBlNo

`func (o *PortcastAPIBillOfLading) HasBlNo() bool`

HasBlNo returns a boolean if a field has been set.

### SetBlNoNil

`func (o *PortcastAPIBillOfLading) SetBlNoNil(b bool)`

 SetBlNoNil sets the value for BlNo to be an explicit nil

### UnsetBlNo
`func (o *PortcastAPIBillOfLading) UnsetBlNo()`

UnsetBlNo ensures that no value is present for BlNo, not even an explicit nil
### GetCarrierNo

`func (o *PortcastAPIBillOfLading) GetCarrierNo() interface{}`

GetCarrierNo returns the CarrierNo field if non-nil, zero value otherwise.

### GetCarrierNoOk

`func (o *PortcastAPIBillOfLading) GetCarrierNoOk() (*interface{}, bool)`

GetCarrierNoOk returns a tuple with the CarrierNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierNo

`func (o *PortcastAPIBillOfLading) SetCarrierNo(v interface{})`

SetCarrierNo sets CarrierNo field to given value.

### HasCarrierNo

`func (o *PortcastAPIBillOfLading) HasCarrierNo() bool`

HasCarrierNo returns a boolean if a field has been set.

### SetCarrierNoNil

`func (o *PortcastAPIBillOfLading) SetCarrierNoNil(b bool)`

 SetCarrierNoNil sets the value for CarrierNo to be an explicit nil

### UnsetCarrierNo
`func (o *PortcastAPIBillOfLading) UnsetCarrierNo()`

UnsetCarrierNo ensures that no value is present for CarrierNo, not even an explicit nil
### GetCntrNo

`func (o *PortcastAPIBillOfLading) GetCntrNo() interface{}`

GetCntrNo returns the CntrNo field if non-nil, zero value otherwise.

### GetCntrNoOk

`func (o *PortcastAPIBillOfLading) GetCntrNoOk() (*interface{}, bool)`

GetCntrNoOk returns a tuple with the CntrNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCntrNo

`func (o *PortcastAPIBillOfLading) SetCntrNo(v interface{})`

SetCntrNo sets CntrNo field to given value.

### HasCntrNo

`func (o *PortcastAPIBillOfLading) HasCntrNo() bool`

HasCntrNo returns a boolean if a field has been set.

### SetCntrNoNil

`func (o *PortcastAPIBillOfLading) SetCntrNoNil(b bool)`

 SetCntrNoNil sets the value for CntrNo to be an explicit nil

### UnsetCntrNo
`func (o *PortcastAPIBillOfLading) UnsetCntrNo()`

UnsetCntrNo ensures that no value is present for CntrNo, not even an explicit nil
### GetCreated

`func (o *PortcastAPIBillOfLading) GetCreated() interface{}`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PortcastAPIBillOfLading) GetCreatedOk() (*interface{}, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PortcastAPIBillOfLading) SetCreated(v interface{})`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PortcastAPIBillOfLading) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *PortcastAPIBillOfLading) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *PortcastAPIBillOfLading) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetId

`func (o *PortcastAPIBillOfLading) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortcastAPIBillOfLading) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortcastAPIBillOfLading) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *PortcastAPIBillOfLading) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PortcastAPIBillOfLading) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PortcastAPIBillOfLading) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPlaceOfDelivery

`func (o *PortcastAPIBillOfLading) GetPlaceOfDelivery() interface{}`

GetPlaceOfDelivery returns the PlaceOfDelivery field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryOk

`func (o *PortcastAPIBillOfLading) GetPlaceOfDeliveryOk() (*interface{}, bool)`

GetPlaceOfDeliveryOk returns a tuple with the PlaceOfDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDelivery

`func (o *PortcastAPIBillOfLading) SetPlaceOfDelivery(v interface{})`

SetPlaceOfDelivery sets PlaceOfDelivery field to given value.

### HasPlaceOfDelivery

`func (o *PortcastAPIBillOfLading) HasPlaceOfDelivery() bool`

HasPlaceOfDelivery returns a boolean if a field has been set.

### SetPlaceOfDeliveryNil

`func (o *PortcastAPIBillOfLading) SetPlaceOfDeliveryNil(b bool)`

 SetPlaceOfDeliveryNil sets the value for PlaceOfDelivery to be an explicit nil

### UnsetPlaceOfDelivery
`func (o *PortcastAPIBillOfLading) UnsetPlaceOfDelivery()`

UnsetPlaceOfDelivery ensures that no value is present for PlaceOfDelivery, not even an explicit nil
### GetPlaceOfDeliveryName

`func (o *PortcastAPIBillOfLading) GetPlaceOfDeliveryName() interface{}`

GetPlaceOfDeliveryName returns the PlaceOfDeliveryName field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryNameOk

`func (o *PortcastAPIBillOfLading) GetPlaceOfDeliveryNameOk() (*interface{}, bool)`

GetPlaceOfDeliveryNameOk returns a tuple with the PlaceOfDeliveryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDeliveryName

`func (o *PortcastAPIBillOfLading) SetPlaceOfDeliveryName(v interface{})`

SetPlaceOfDeliveryName sets PlaceOfDeliveryName field to given value.

### HasPlaceOfDeliveryName

`func (o *PortcastAPIBillOfLading) HasPlaceOfDeliveryName() bool`

HasPlaceOfDeliveryName returns a boolean if a field has been set.

### SetPlaceOfDeliveryNameNil

`func (o *PortcastAPIBillOfLading) SetPlaceOfDeliveryNameNil(b bool)`

 SetPlaceOfDeliveryNameNil sets the value for PlaceOfDeliveryName to be an explicit nil

### UnsetPlaceOfDeliveryName
`func (o *PortcastAPIBillOfLading) UnsetPlaceOfDeliveryName()`

UnsetPlaceOfDeliveryName ensures that no value is present for PlaceOfDeliveryName, not even an explicit nil
### GetPlaceOfReceipt

`func (o *PortcastAPIBillOfLading) GetPlaceOfReceipt() interface{}`

GetPlaceOfReceipt returns the PlaceOfReceipt field if non-nil, zero value otherwise.

### GetPlaceOfReceiptOk

`func (o *PortcastAPIBillOfLading) GetPlaceOfReceiptOk() (*interface{}, bool)`

GetPlaceOfReceiptOk returns a tuple with the PlaceOfReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfReceipt

`func (o *PortcastAPIBillOfLading) SetPlaceOfReceipt(v interface{})`

SetPlaceOfReceipt sets PlaceOfReceipt field to given value.

### HasPlaceOfReceipt

`func (o *PortcastAPIBillOfLading) HasPlaceOfReceipt() bool`

HasPlaceOfReceipt returns a boolean if a field has been set.

### SetPlaceOfReceiptNil

`func (o *PortcastAPIBillOfLading) SetPlaceOfReceiptNil(b bool)`

 SetPlaceOfReceiptNil sets the value for PlaceOfReceipt to be an explicit nil

### UnsetPlaceOfReceipt
`func (o *PortcastAPIBillOfLading) UnsetPlaceOfReceipt()`

UnsetPlaceOfReceipt ensures that no value is present for PlaceOfReceipt, not even an explicit nil
### GetPlaceOfReceiptName

`func (o *PortcastAPIBillOfLading) GetPlaceOfReceiptName() interface{}`

GetPlaceOfReceiptName returns the PlaceOfReceiptName field if non-nil, zero value otherwise.

### GetPlaceOfReceiptNameOk

`func (o *PortcastAPIBillOfLading) GetPlaceOfReceiptNameOk() (*interface{}, bool)`

GetPlaceOfReceiptNameOk returns a tuple with the PlaceOfReceiptName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfReceiptName

`func (o *PortcastAPIBillOfLading) SetPlaceOfReceiptName(v interface{})`

SetPlaceOfReceiptName sets PlaceOfReceiptName field to given value.

### HasPlaceOfReceiptName

`func (o *PortcastAPIBillOfLading) HasPlaceOfReceiptName() bool`

HasPlaceOfReceiptName returns a boolean if a field has been set.

### SetPlaceOfReceiptNameNil

`func (o *PortcastAPIBillOfLading) SetPlaceOfReceiptNameNil(b bool)`

 SetPlaceOfReceiptNameNil sets the value for PlaceOfReceiptName to be an explicit nil

### UnsetPlaceOfReceiptName
`func (o *PortcastAPIBillOfLading) UnsetPlaceOfReceiptName()`

UnsetPlaceOfReceiptName ensures that no value is present for PlaceOfReceiptName, not even an explicit nil
### GetPod

`func (o *PortcastAPIBillOfLading) GetPod() interface{}`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *PortcastAPIBillOfLading) GetPodOk() (*interface{}, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *PortcastAPIBillOfLading) SetPod(v interface{})`

SetPod sets Pod field to given value.

### HasPod

`func (o *PortcastAPIBillOfLading) HasPod() bool`

HasPod returns a boolean if a field has been set.

### SetPodNil

`func (o *PortcastAPIBillOfLading) SetPodNil(b bool)`

 SetPodNil sets the value for Pod to be an explicit nil

### UnsetPod
`func (o *PortcastAPIBillOfLading) UnsetPod()`

UnsetPod ensures that no value is present for Pod, not even an explicit nil
### GetPodActualArrivalLt

`func (o *PortcastAPIBillOfLading) GetPodActualArrivalLt() interface{}`

GetPodActualArrivalLt returns the PodActualArrivalLt field if non-nil, zero value otherwise.

### GetPodActualArrivalLtOk

`func (o *PortcastAPIBillOfLading) GetPodActualArrivalLtOk() (*interface{}, bool)`

GetPodActualArrivalLtOk returns a tuple with the PodActualArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualArrivalLt

`func (o *PortcastAPIBillOfLading) SetPodActualArrivalLt(v interface{})`

SetPodActualArrivalLt sets PodActualArrivalLt field to given value.

### HasPodActualArrivalLt

`func (o *PortcastAPIBillOfLading) HasPodActualArrivalLt() bool`

HasPodActualArrivalLt returns a boolean if a field has been set.

### SetPodActualArrivalLtNil

`func (o *PortcastAPIBillOfLading) SetPodActualArrivalLtNil(b bool)`

 SetPodActualArrivalLtNil sets the value for PodActualArrivalLt to be an explicit nil

### UnsetPodActualArrivalLt
`func (o *PortcastAPIBillOfLading) UnsetPodActualArrivalLt()`

UnsetPodActualArrivalLt ensures that no value is present for PodActualArrivalLt, not even an explicit nil
### GetPodActualArrivalLtFromAis

`func (o *PortcastAPIBillOfLading) GetPodActualArrivalLtFromAis() interface{}`

GetPodActualArrivalLtFromAis returns the PodActualArrivalLtFromAis field if non-nil, zero value otherwise.

### GetPodActualArrivalLtFromAisOk

`func (o *PortcastAPIBillOfLading) GetPodActualArrivalLtFromAisOk() (*interface{}, bool)`

GetPodActualArrivalLtFromAisOk returns a tuple with the PodActualArrivalLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualArrivalLtFromAis

`func (o *PortcastAPIBillOfLading) SetPodActualArrivalLtFromAis(v interface{})`

SetPodActualArrivalLtFromAis sets PodActualArrivalLtFromAis field to given value.

### HasPodActualArrivalLtFromAis

`func (o *PortcastAPIBillOfLading) HasPodActualArrivalLtFromAis() bool`

HasPodActualArrivalLtFromAis returns a boolean if a field has been set.

### SetPodActualArrivalLtFromAisNil

`func (o *PortcastAPIBillOfLading) SetPodActualArrivalLtFromAisNil(b bool)`

 SetPodActualArrivalLtFromAisNil sets the value for PodActualArrivalLtFromAis to be an explicit nil

### UnsetPodActualArrivalLtFromAis
`func (o *PortcastAPIBillOfLading) UnsetPodActualArrivalLtFromAis()`

UnsetPodActualArrivalLtFromAis ensures that no value is present for PodActualArrivalLtFromAis, not even an explicit nil
### GetPodActualDepartureLtFromAis

`func (o *PortcastAPIBillOfLading) GetPodActualDepartureLtFromAis() interface{}`

GetPodActualDepartureLtFromAis returns the PodActualDepartureLtFromAis field if non-nil, zero value otherwise.

### GetPodActualDepartureLtFromAisOk

`func (o *PortcastAPIBillOfLading) GetPodActualDepartureLtFromAisOk() (*interface{}, bool)`

GetPodActualDepartureLtFromAisOk returns a tuple with the PodActualDepartureLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualDepartureLtFromAis

`func (o *PortcastAPIBillOfLading) SetPodActualDepartureLtFromAis(v interface{})`

SetPodActualDepartureLtFromAis sets PodActualDepartureLtFromAis field to given value.

### HasPodActualDepartureLtFromAis

`func (o *PortcastAPIBillOfLading) HasPodActualDepartureLtFromAis() bool`

HasPodActualDepartureLtFromAis returns a boolean if a field has been set.

### SetPodActualDepartureLtFromAisNil

`func (o *PortcastAPIBillOfLading) SetPodActualDepartureLtFromAisNil(b bool)`

 SetPodActualDepartureLtFromAisNil sets the value for PodActualDepartureLtFromAis to be an explicit nil

### UnsetPodActualDepartureLtFromAis
`func (o *PortcastAPIBillOfLading) UnsetPodActualDepartureLtFromAis()`

UnsetPodActualDepartureLtFromAis ensures that no value is present for PodActualDepartureLtFromAis, not even an explicit nil
### GetPodActualDischargeLt

`func (o *PortcastAPIBillOfLading) GetPodActualDischargeLt() interface{}`

GetPodActualDischargeLt returns the PodActualDischargeLt field if non-nil, zero value otherwise.

### GetPodActualDischargeLtOk

`func (o *PortcastAPIBillOfLading) GetPodActualDischargeLtOk() (*interface{}, bool)`

GetPodActualDischargeLtOk returns a tuple with the PodActualDischargeLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodActualDischargeLt

`func (o *PortcastAPIBillOfLading) SetPodActualDischargeLt(v interface{})`

SetPodActualDischargeLt sets PodActualDischargeLt field to given value.

### HasPodActualDischargeLt

`func (o *PortcastAPIBillOfLading) HasPodActualDischargeLt() bool`

HasPodActualDischargeLt returns a boolean if a field has been set.

### SetPodActualDischargeLtNil

`func (o *PortcastAPIBillOfLading) SetPodActualDischargeLtNil(b bool)`

 SetPodActualDischargeLtNil sets the value for PodActualDischargeLt to be an explicit nil

### UnsetPodActualDischargeLt
`func (o *PortcastAPIBillOfLading) UnsetPodActualDischargeLt()`

UnsetPodActualDischargeLt ensures that no value is present for PodActualDischargeLt, not even an explicit nil
### GetPodName

`func (o *PortcastAPIBillOfLading) GetPodName() interface{}`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *PortcastAPIBillOfLading) GetPodNameOk() (*interface{}, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *PortcastAPIBillOfLading) SetPodName(v interface{})`

SetPodName sets PodName field to given value.

### HasPodName

`func (o *PortcastAPIBillOfLading) HasPodName() bool`

HasPodName returns a boolean if a field has been set.

### SetPodNameNil

`func (o *PortcastAPIBillOfLading) SetPodNameNil(b bool)`

 SetPodNameNil sets the value for PodName to be an explicit nil

### UnsetPodName
`func (o *PortcastAPIBillOfLading) UnsetPodName()`

UnsetPodName ensures that no value is present for PodName, not even an explicit nil
### GetPodPredictedArrivalLt

`func (o *PortcastAPIBillOfLading) GetPodPredictedArrivalLt() interface{}`

GetPodPredictedArrivalLt returns the PodPredictedArrivalLt field if non-nil, zero value otherwise.

### GetPodPredictedArrivalLtOk

`func (o *PortcastAPIBillOfLading) GetPodPredictedArrivalLtOk() (*interface{}, bool)`

GetPodPredictedArrivalLtOk returns a tuple with the PodPredictedArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPredictedArrivalLt

`func (o *PortcastAPIBillOfLading) SetPodPredictedArrivalLt(v interface{})`

SetPodPredictedArrivalLt sets PodPredictedArrivalLt field to given value.

### HasPodPredictedArrivalLt

`func (o *PortcastAPIBillOfLading) HasPodPredictedArrivalLt() bool`

HasPodPredictedArrivalLt returns a boolean if a field has been set.

### SetPodPredictedArrivalLtNil

`func (o *PortcastAPIBillOfLading) SetPodPredictedArrivalLtNil(b bool)`

 SetPodPredictedArrivalLtNil sets the value for PodPredictedArrivalLt to be an explicit nil

### UnsetPodPredictedArrivalLt
`func (o *PortcastAPIBillOfLading) UnsetPodPredictedArrivalLt()`

UnsetPodPredictedArrivalLt ensures that no value is present for PodPredictedArrivalLt, not even an explicit nil
### GetPodPredictedDepartureLt

`func (o *PortcastAPIBillOfLading) GetPodPredictedDepartureLt() interface{}`

GetPodPredictedDepartureLt returns the PodPredictedDepartureLt field if non-nil, zero value otherwise.

### GetPodPredictedDepartureLtOk

`func (o *PortcastAPIBillOfLading) GetPodPredictedDepartureLtOk() (*interface{}, bool)`

GetPodPredictedDepartureLtOk returns a tuple with the PodPredictedDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPredictedDepartureLt

`func (o *PortcastAPIBillOfLading) SetPodPredictedDepartureLt(v interface{})`

SetPodPredictedDepartureLt sets PodPredictedDepartureLt field to given value.

### HasPodPredictedDepartureLt

`func (o *PortcastAPIBillOfLading) HasPodPredictedDepartureLt() bool`

HasPodPredictedDepartureLt returns a boolean if a field has been set.

### SetPodPredictedDepartureLtNil

`func (o *PortcastAPIBillOfLading) SetPodPredictedDepartureLtNil(b bool)`

 SetPodPredictedDepartureLtNil sets the value for PodPredictedDepartureLt to be an explicit nil

### UnsetPodPredictedDepartureLt
`func (o *PortcastAPIBillOfLading) UnsetPodPredictedDepartureLt()`

UnsetPodPredictedDepartureLt ensures that no value is present for PodPredictedDepartureLt, not even an explicit nil
### GetPodScheduledArrivalLt

`func (o *PortcastAPIBillOfLading) GetPodScheduledArrivalLt() interface{}`

GetPodScheduledArrivalLt returns the PodScheduledArrivalLt field if non-nil, zero value otherwise.

### GetPodScheduledArrivalLtOk

`func (o *PortcastAPIBillOfLading) GetPodScheduledArrivalLtOk() (*interface{}, bool)`

GetPodScheduledArrivalLtOk returns a tuple with the PodScheduledArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledArrivalLt

`func (o *PortcastAPIBillOfLading) SetPodScheduledArrivalLt(v interface{})`

SetPodScheduledArrivalLt sets PodScheduledArrivalLt field to given value.

### HasPodScheduledArrivalLt

`func (o *PortcastAPIBillOfLading) HasPodScheduledArrivalLt() bool`

HasPodScheduledArrivalLt returns a boolean if a field has been set.

### SetPodScheduledArrivalLtNil

`func (o *PortcastAPIBillOfLading) SetPodScheduledArrivalLtNil(b bool)`

 SetPodScheduledArrivalLtNil sets the value for PodScheduledArrivalLt to be an explicit nil

### UnsetPodScheduledArrivalLt
`func (o *PortcastAPIBillOfLading) UnsetPodScheduledArrivalLt()`

UnsetPodScheduledArrivalLt ensures that no value is present for PodScheduledArrivalLt, not even an explicit nil
### GetPodScheduledArrivalLtFirstSeen

`func (o *PortcastAPIBillOfLading) GetPodScheduledArrivalLtFirstSeen() interface{}`

GetPodScheduledArrivalLtFirstSeen returns the PodScheduledArrivalLtFirstSeen field if non-nil, zero value otherwise.

### GetPodScheduledArrivalLtFirstSeenOk

`func (o *PortcastAPIBillOfLading) GetPodScheduledArrivalLtFirstSeenOk() (*interface{}, bool)`

GetPodScheduledArrivalLtFirstSeenOk returns a tuple with the PodScheduledArrivalLtFirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledArrivalLtFirstSeen

`func (o *PortcastAPIBillOfLading) SetPodScheduledArrivalLtFirstSeen(v interface{})`

SetPodScheduledArrivalLtFirstSeen sets PodScheduledArrivalLtFirstSeen field to given value.

### HasPodScheduledArrivalLtFirstSeen

`func (o *PortcastAPIBillOfLading) HasPodScheduledArrivalLtFirstSeen() bool`

HasPodScheduledArrivalLtFirstSeen returns a boolean if a field has been set.

### SetPodScheduledArrivalLtFirstSeenNil

`func (o *PortcastAPIBillOfLading) SetPodScheduledArrivalLtFirstSeenNil(b bool)`

 SetPodScheduledArrivalLtFirstSeenNil sets the value for PodScheduledArrivalLtFirstSeen to be an explicit nil

### UnsetPodScheduledArrivalLtFirstSeen
`func (o *PortcastAPIBillOfLading) UnsetPodScheduledArrivalLtFirstSeen()`

UnsetPodScheduledArrivalLtFirstSeen ensures that no value is present for PodScheduledArrivalLtFirstSeen, not even an explicit nil
### GetPodScheduledArrivalLtFromSchedule

`func (o *PortcastAPIBillOfLading) GetPodScheduledArrivalLtFromSchedule() interface{}`

GetPodScheduledArrivalLtFromSchedule returns the PodScheduledArrivalLtFromSchedule field if non-nil, zero value otherwise.

### GetPodScheduledArrivalLtFromScheduleOk

`func (o *PortcastAPIBillOfLading) GetPodScheduledArrivalLtFromScheduleOk() (*interface{}, bool)`

GetPodScheduledArrivalLtFromScheduleOk returns a tuple with the PodScheduledArrivalLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledArrivalLtFromSchedule

`func (o *PortcastAPIBillOfLading) SetPodScheduledArrivalLtFromSchedule(v interface{})`

SetPodScheduledArrivalLtFromSchedule sets PodScheduledArrivalLtFromSchedule field to given value.

### HasPodScheduledArrivalLtFromSchedule

`func (o *PortcastAPIBillOfLading) HasPodScheduledArrivalLtFromSchedule() bool`

HasPodScheduledArrivalLtFromSchedule returns a boolean if a field has been set.

### SetPodScheduledArrivalLtFromScheduleNil

`func (o *PortcastAPIBillOfLading) SetPodScheduledArrivalLtFromScheduleNil(b bool)`

 SetPodScheduledArrivalLtFromScheduleNil sets the value for PodScheduledArrivalLtFromSchedule to be an explicit nil

### UnsetPodScheduledArrivalLtFromSchedule
`func (o *PortcastAPIBillOfLading) UnsetPodScheduledArrivalLtFromSchedule()`

UnsetPodScheduledArrivalLtFromSchedule ensures that no value is present for PodScheduledArrivalLtFromSchedule, not even an explicit nil
### GetPodScheduledDepartureLtFromSchedule

`func (o *PortcastAPIBillOfLading) GetPodScheduledDepartureLtFromSchedule() interface{}`

GetPodScheduledDepartureLtFromSchedule returns the PodScheduledDepartureLtFromSchedule field if non-nil, zero value otherwise.

### GetPodScheduledDepartureLtFromScheduleOk

`func (o *PortcastAPIBillOfLading) GetPodScheduledDepartureLtFromScheduleOk() (*interface{}, bool)`

GetPodScheduledDepartureLtFromScheduleOk returns a tuple with the PodScheduledDepartureLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledDepartureLtFromSchedule

`func (o *PortcastAPIBillOfLading) SetPodScheduledDepartureLtFromSchedule(v interface{})`

SetPodScheduledDepartureLtFromSchedule sets PodScheduledDepartureLtFromSchedule field to given value.

### HasPodScheduledDepartureLtFromSchedule

`func (o *PortcastAPIBillOfLading) HasPodScheduledDepartureLtFromSchedule() bool`

HasPodScheduledDepartureLtFromSchedule returns a boolean if a field has been set.

### SetPodScheduledDepartureLtFromScheduleNil

`func (o *PortcastAPIBillOfLading) SetPodScheduledDepartureLtFromScheduleNil(b bool)`

 SetPodScheduledDepartureLtFromScheduleNil sets the value for PodScheduledDepartureLtFromSchedule to be an explicit nil

### UnsetPodScheduledDepartureLtFromSchedule
`func (o *PortcastAPIBillOfLading) UnsetPodScheduledDepartureLtFromSchedule()`

UnsetPodScheduledDepartureLtFromSchedule ensures that no value is present for PodScheduledDepartureLtFromSchedule, not even an explicit nil
### GetPodScheduledDischargeLt

`func (o *PortcastAPIBillOfLading) GetPodScheduledDischargeLt() interface{}`

GetPodScheduledDischargeLt returns the PodScheduledDischargeLt field if non-nil, zero value otherwise.

### GetPodScheduledDischargeLtOk

`func (o *PortcastAPIBillOfLading) GetPodScheduledDischargeLtOk() (*interface{}, bool)`

GetPodScheduledDischargeLtOk returns a tuple with the PodScheduledDischargeLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodScheduledDischargeLt

`func (o *PortcastAPIBillOfLading) SetPodScheduledDischargeLt(v interface{})`

SetPodScheduledDischargeLt sets PodScheduledDischargeLt field to given value.

### HasPodScheduledDischargeLt

`func (o *PortcastAPIBillOfLading) HasPodScheduledDischargeLt() bool`

HasPodScheduledDischargeLt returns a boolean if a field has been set.

### SetPodScheduledDischargeLtNil

`func (o *PortcastAPIBillOfLading) SetPodScheduledDischargeLtNil(b bool)`

 SetPodScheduledDischargeLtNil sets the value for PodScheduledDischargeLt to be an explicit nil

### UnsetPodScheduledDischargeLt
`func (o *PortcastAPIBillOfLading) UnsetPodScheduledDischargeLt()`

UnsetPodScheduledDischargeLt ensures that no value is present for PodScheduledDischargeLt, not even an explicit nil
### GetPodTerminalName

`func (o *PortcastAPIBillOfLading) GetPodTerminalName() interface{}`

GetPodTerminalName returns the PodTerminalName field if non-nil, zero value otherwise.

### GetPodTerminalNameOk

`func (o *PortcastAPIBillOfLading) GetPodTerminalNameOk() (*interface{}, bool)`

GetPodTerminalNameOk returns a tuple with the PodTerminalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodTerminalName

`func (o *PortcastAPIBillOfLading) SetPodTerminalName(v interface{})`

SetPodTerminalName sets PodTerminalName field to given value.

### HasPodTerminalName

`func (o *PortcastAPIBillOfLading) HasPodTerminalName() bool`

HasPodTerminalName returns a boolean if a field has been set.

### SetPodTerminalNameNil

`func (o *PortcastAPIBillOfLading) SetPodTerminalNameNil(b bool)`

 SetPodTerminalNameNil sets the value for PodTerminalName to be an explicit nil

### UnsetPodTerminalName
`func (o *PortcastAPIBillOfLading) UnsetPodTerminalName()`

UnsetPodTerminalName ensures that no value is present for PodTerminalName, not even an explicit nil
### GetPol

`func (o *PortcastAPIBillOfLading) GetPol() interface{}`

GetPol returns the Pol field if non-nil, zero value otherwise.

### GetPolOk

`func (o *PortcastAPIBillOfLading) GetPolOk() (*interface{}, bool)`

GetPolOk returns a tuple with the Pol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPol

`func (o *PortcastAPIBillOfLading) SetPol(v interface{})`

SetPol sets Pol field to given value.

### HasPol

`func (o *PortcastAPIBillOfLading) HasPol() bool`

HasPol returns a boolean if a field has been set.

### SetPolNil

`func (o *PortcastAPIBillOfLading) SetPolNil(b bool)`

 SetPolNil sets the value for Pol to be an explicit nil

### UnsetPol
`func (o *PortcastAPIBillOfLading) UnsetPol()`

UnsetPol ensures that no value is present for Pol, not even an explicit nil
### GetPolActualArrivalLtFromAis

`func (o *PortcastAPIBillOfLading) GetPolActualArrivalLtFromAis() interface{}`

GetPolActualArrivalLtFromAis returns the PolActualArrivalLtFromAis field if non-nil, zero value otherwise.

### GetPolActualArrivalLtFromAisOk

`func (o *PortcastAPIBillOfLading) GetPolActualArrivalLtFromAisOk() (*interface{}, bool)`

GetPolActualArrivalLtFromAisOk returns a tuple with the PolActualArrivalLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualArrivalLtFromAis

`func (o *PortcastAPIBillOfLading) SetPolActualArrivalLtFromAis(v interface{})`

SetPolActualArrivalLtFromAis sets PolActualArrivalLtFromAis field to given value.

### HasPolActualArrivalLtFromAis

`func (o *PortcastAPIBillOfLading) HasPolActualArrivalLtFromAis() bool`

HasPolActualArrivalLtFromAis returns a boolean if a field has been set.

### SetPolActualArrivalLtFromAisNil

`func (o *PortcastAPIBillOfLading) SetPolActualArrivalLtFromAisNil(b bool)`

 SetPolActualArrivalLtFromAisNil sets the value for PolActualArrivalLtFromAis to be an explicit nil

### UnsetPolActualArrivalLtFromAis
`func (o *PortcastAPIBillOfLading) UnsetPolActualArrivalLtFromAis()`

UnsetPolActualArrivalLtFromAis ensures that no value is present for PolActualArrivalLtFromAis, not even an explicit nil
### GetPolActualDepartureLt

`func (o *PortcastAPIBillOfLading) GetPolActualDepartureLt() interface{}`

GetPolActualDepartureLt returns the PolActualDepartureLt field if non-nil, zero value otherwise.

### GetPolActualDepartureLtOk

`func (o *PortcastAPIBillOfLading) GetPolActualDepartureLtOk() (*interface{}, bool)`

GetPolActualDepartureLtOk returns a tuple with the PolActualDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualDepartureLt

`func (o *PortcastAPIBillOfLading) SetPolActualDepartureLt(v interface{})`

SetPolActualDepartureLt sets PolActualDepartureLt field to given value.

### HasPolActualDepartureLt

`func (o *PortcastAPIBillOfLading) HasPolActualDepartureLt() bool`

HasPolActualDepartureLt returns a boolean if a field has been set.

### SetPolActualDepartureLtNil

`func (o *PortcastAPIBillOfLading) SetPolActualDepartureLtNil(b bool)`

 SetPolActualDepartureLtNil sets the value for PolActualDepartureLt to be an explicit nil

### UnsetPolActualDepartureLt
`func (o *PortcastAPIBillOfLading) UnsetPolActualDepartureLt()`

UnsetPolActualDepartureLt ensures that no value is present for PolActualDepartureLt, not even an explicit nil
### GetPolActualDepartureLtFromAis

`func (o *PortcastAPIBillOfLading) GetPolActualDepartureLtFromAis() interface{}`

GetPolActualDepartureLtFromAis returns the PolActualDepartureLtFromAis field if non-nil, zero value otherwise.

### GetPolActualDepartureLtFromAisOk

`func (o *PortcastAPIBillOfLading) GetPolActualDepartureLtFromAisOk() (*interface{}, bool)`

GetPolActualDepartureLtFromAisOk returns a tuple with the PolActualDepartureLtFromAis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualDepartureLtFromAis

`func (o *PortcastAPIBillOfLading) SetPolActualDepartureLtFromAis(v interface{})`

SetPolActualDepartureLtFromAis sets PolActualDepartureLtFromAis field to given value.

### HasPolActualDepartureLtFromAis

`func (o *PortcastAPIBillOfLading) HasPolActualDepartureLtFromAis() bool`

HasPolActualDepartureLtFromAis returns a boolean if a field has been set.

### SetPolActualDepartureLtFromAisNil

`func (o *PortcastAPIBillOfLading) SetPolActualDepartureLtFromAisNil(b bool)`

 SetPolActualDepartureLtFromAisNil sets the value for PolActualDepartureLtFromAis to be an explicit nil

### UnsetPolActualDepartureLtFromAis
`func (o *PortcastAPIBillOfLading) UnsetPolActualDepartureLtFromAis()`

UnsetPolActualDepartureLtFromAis ensures that no value is present for PolActualDepartureLtFromAis, not even an explicit nil
### GetPolActualLoadingLt

`func (o *PortcastAPIBillOfLading) GetPolActualLoadingLt() interface{}`

GetPolActualLoadingLt returns the PolActualLoadingLt field if non-nil, zero value otherwise.

### GetPolActualLoadingLtOk

`func (o *PortcastAPIBillOfLading) GetPolActualLoadingLtOk() (*interface{}, bool)`

GetPolActualLoadingLtOk returns a tuple with the PolActualLoadingLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolActualLoadingLt

`func (o *PortcastAPIBillOfLading) SetPolActualLoadingLt(v interface{})`

SetPolActualLoadingLt sets PolActualLoadingLt field to given value.

### HasPolActualLoadingLt

`func (o *PortcastAPIBillOfLading) HasPolActualLoadingLt() bool`

HasPolActualLoadingLt returns a boolean if a field has been set.

### SetPolActualLoadingLtNil

`func (o *PortcastAPIBillOfLading) SetPolActualLoadingLtNil(b bool)`

 SetPolActualLoadingLtNil sets the value for PolActualLoadingLt to be an explicit nil

### UnsetPolActualLoadingLt
`func (o *PortcastAPIBillOfLading) UnsetPolActualLoadingLt()`

UnsetPolActualLoadingLt ensures that no value is present for PolActualLoadingLt, not even an explicit nil
### GetPolName

`func (o *PortcastAPIBillOfLading) GetPolName() interface{}`

GetPolName returns the PolName field if non-nil, zero value otherwise.

### GetPolNameOk

`func (o *PortcastAPIBillOfLading) GetPolNameOk() (*interface{}, bool)`

GetPolNameOk returns a tuple with the PolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolName

`func (o *PortcastAPIBillOfLading) SetPolName(v interface{})`

SetPolName sets PolName field to given value.

### HasPolName

`func (o *PortcastAPIBillOfLading) HasPolName() bool`

HasPolName returns a boolean if a field has been set.

### SetPolNameNil

`func (o *PortcastAPIBillOfLading) SetPolNameNil(b bool)`

 SetPolNameNil sets the value for PolName to be an explicit nil

### UnsetPolName
`func (o *PortcastAPIBillOfLading) UnsetPolName()`

UnsetPolName ensures that no value is present for PolName, not even an explicit nil
### GetPolPredictedArrivalLt

`func (o *PortcastAPIBillOfLading) GetPolPredictedArrivalLt() interface{}`

GetPolPredictedArrivalLt returns the PolPredictedArrivalLt field if non-nil, zero value otherwise.

### GetPolPredictedArrivalLtOk

`func (o *PortcastAPIBillOfLading) GetPolPredictedArrivalLtOk() (*interface{}, bool)`

GetPolPredictedArrivalLtOk returns a tuple with the PolPredictedArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolPredictedArrivalLt

`func (o *PortcastAPIBillOfLading) SetPolPredictedArrivalLt(v interface{})`

SetPolPredictedArrivalLt sets PolPredictedArrivalLt field to given value.

### HasPolPredictedArrivalLt

`func (o *PortcastAPIBillOfLading) HasPolPredictedArrivalLt() bool`

HasPolPredictedArrivalLt returns a boolean if a field has been set.

### SetPolPredictedArrivalLtNil

`func (o *PortcastAPIBillOfLading) SetPolPredictedArrivalLtNil(b bool)`

 SetPolPredictedArrivalLtNil sets the value for PolPredictedArrivalLt to be an explicit nil

### UnsetPolPredictedArrivalLt
`func (o *PortcastAPIBillOfLading) UnsetPolPredictedArrivalLt()`

UnsetPolPredictedArrivalLt ensures that no value is present for PolPredictedArrivalLt, not even an explicit nil
### GetPolPredictedDepartureLt

`func (o *PortcastAPIBillOfLading) GetPolPredictedDepartureLt() interface{}`

GetPolPredictedDepartureLt returns the PolPredictedDepartureLt field if non-nil, zero value otherwise.

### GetPolPredictedDepartureLtOk

`func (o *PortcastAPIBillOfLading) GetPolPredictedDepartureLtOk() (*interface{}, bool)`

GetPolPredictedDepartureLtOk returns a tuple with the PolPredictedDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolPredictedDepartureLt

`func (o *PortcastAPIBillOfLading) SetPolPredictedDepartureLt(v interface{})`

SetPolPredictedDepartureLt sets PolPredictedDepartureLt field to given value.

### HasPolPredictedDepartureLt

`func (o *PortcastAPIBillOfLading) HasPolPredictedDepartureLt() bool`

HasPolPredictedDepartureLt returns a boolean if a field has been set.

### SetPolPredictedDepartureLtNil

`func (o *PortcastAPIBillOfLading) SetPolPredictedDepartureLtNil(b bool)`

 SetPolPredictedDepartureLtNil sets the value for PolPredictedDepartureLt to be an explicit nil

### UnsetPolPredictedDepartureLt
`func (o *PortcastAPIBillOfLading) UnsetPolPredictedDepartureLt()`

UnsetPolPredictedDepartureLt ensures that no value is present for PolPredictedDepartureLt, not even an explicit nil
### GetPolScheduledArrivalLtFromSchedule

`func (o *PortcastAPIBillOfLading) GetPolScheduledArrivalLtFromSchedule() interface{}`

GetPolScheduledArrivalLtFromSchedule returns the PolScheduledArrivalLtFromSchedule field if non-nil, zero value otherwise.

### GetPolScheduledArrivalLtFromScheduleOk

`func (o *PortcastAPIBillOfLading) GetPolScheduledArrivalLtFromScheduleOk() (*interface{}, bool)`

GetPolScheduledArrivalLtFromScheduleOk returns a tuple with the PolScheduledArrivalLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledArrivalLtFromSchedule

`func (o *PortcastAPIBillOfLading) SetPolScheduledArrivalLtFromSchedule(v interface{})`

SetPolScheduledArrivalLtFromSchedule sets PolScheduledArrivalLtFromSchedule field to given value.

### HasPolScheduledArrivalLtFromSchedule

`func (o *PortcastAPIBillOfLading) HasPolScheduledArrivalLtFromSchedule() bool`

HasPolScheduledArrivalLtFromSchedule returns a boolean if a field has been set.

### SetPolScheduledArrivalLtFromScheduleNil

`func (o *PortcastAPIBillOfLading) SetPolScheduledArrivalLtFromScheduleNil(b bool)`

 SetPolScheduledArrivalLtFromScheduleNil sets the value for PolScheduledArrivalLtFromSchedule to be an explicit nil

### UnsetPolScheduledArrivalLtFromSchedule
`func (o *PortcastAPIBillOfLading) UnsetPolScheduledArrivalLtFromSchedule()`

UnsetPolScheduledArrivalLtFromSchedule ensures that no value is present for PolScheduledArrivalLtFromSchedule, not even an explicit nil
### GetPolScheduledDepartureLt

`func (o *PortcastAPIBillOfLading) GetPolScheduledDepartureLt() interface{}`

GetPolScheduledDepartureLt returns the PolScheduledDepartureLt field if non-nil, zero value otherwise.

### GetPolScheduledDepartureLtOk

`func (o *PortcastAPIBillOfLading) GetPolScheduledDepartureLtOk() (*interface{}, bool)`

GetPolScheduledDepartureLtOk returns a tuple with the PolScheduledDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledDepartureLt

`func (o *PortcastAPIBillOfLading) SetPolScheduledDepartureLt(v interface{})`

SetPolScheduledDepartureLt sets PolScheduledDepartureLt field to given value.

### HasPolScheduledDepartureLt

`func (o *PortcastAPIBillOfLading) HasPolScheduledDepartureLt() bool`

HasPolScheduledDepartureLt returns a boolean if a field has been set.

### SetPolScheduledDepartureLtNil

`func (o *PortcastAPIBillOfLading) SetPolScheduledDepartureLtNil(b bool)`

 SetPolScheduledDepartureLtNil sets the value for PolScheduledDepartureLt to be an explicit nil

### UnsetPolScheduledDepartureLt
`func (o *PortcastAPIBillOfLading) UnsetPolScheduledDepartureLt()`

UnsetPolScheduledDepartureLt ensures that no value is present for PolScheduledDepartureLt, not even an explicit nil
### GetPolScheduledDepartureLtFirstSeen

`func (o *PortcastAPIBillOfLading) GetPolScheduledDepartureLtFirstSeen() interface{}`

GetPolScheduledDepartureLtFirstSeen returns the PolScheduledDepartureLtFirstSeen field if non-nil, zero value otherwise.

### GetPolScheduledDepartureLtFirstSeenOk

`func (o *PortcastAPIBillOfLading) GetPolScheduledDepartureLtFirstSeenOk() (*interface{}, bool)`

GetPolScheduledDepartureLtFirstSeenOk returns a tuple with the PolScheduledDepartureLtFirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledDepartureLtFirstSeen

`func (o *PortcastAPIBillOfLading) SetPolScheduledDepartureLtFirstSeen(v interface{})`

SetPolScheduledDepartureLtFirstSeen sets PolScheduledDepartureLtFirstSeen field to given value.

### HasPolScheduledDepartureLtFirstSeen

`func (o *PortcastAPIBillOfLading) HasPolScheduledDepartureLtFirstSeen() bool`

HasPolScheduledDepartureLtFirstSeen returns a boolean if a field has been set.

### SetPolScheduledDepartureLtFirstSeenNil

`func (o *PortcastAPIBillOfLading) SetPolScheduledDepartureLtFirstSeenNil(b bool)`

 SetPolScheduledDepartureLtFirstSeenNil sets the value for PolScheduledDepartureLtFirstSeen to be an explicit nil

### UnsetPolScheduledDepartureLtFirstSeen
`func (o *PortcastAPIBillOfLading) UnsetPolScheduledDepartureLtFirstSeen()`

UnsetPolScheduledDepartureLtFirstSeen ensures that no value is present for PolScheduledDepartureLtFirstSeen, not even an explicit nil
### GetPolScheduledDepartureLtFromSchedule

`func (o *PortcastAPIBillOfLading) GetPolScheduledDepartureLtFromSchedule() interface{}`

GetPolScheduledDepartureLtFromSchedule returns the PolScheduledDepartureLtFromSchedule field if non-nil, zero value otherwise.

### GetPolScheduledDepartureLtFromScheduleOk

`func (o *PortcastAPIBillOfLading) GetPolScheduledDepartureLtFromScheduleOk() (*interface{}, bool)`

GetPolScheduledDepartureLtFromScheduleOk returns a tuple with the PolScheduledDepartureLtFromSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledDepartureLtFromSchedule

`func (o *PortcastAPIBillOfLading) SetPolScheduledDepartureLtFromSchedule(v interface{})`

SetPolScheduledDepartureLtFromSchedule sets PolScheduledDepartureLtFromSchedule field to given value.

### HasPolScheduledDepartureLtFromSchedule

`func (o *PortcastAPIBillOfLading) HasPolScheduledDepartureLtFromSchedule() bool`

HasPolScheduledDepartureLtFromSchedule returns a boolean if a field has been set.

### SetPolScheduledDepartureLtFromScheduleNil

`func (o *PortcastAPIBillOfLading) SetPolScheduledDepartureLtFromScheduleNil(b bool)`

 SetPolScheduledDepartureLtFromScheduleNil sets the value for PolScheduledDepartureLtFromSchedule to be an explicit nil

### UnsetPolScheduledDepartureLtFromSchedule
`func (o *PortcastAPIBillOfLading) UnsetPolScheduledDepartureLtFromSchedule()`

UnsetPolScheduledDepartureLtFromSchedule ensures that no value is present for PolScheduledDepartureLtFromSchedule, not even an explicit nil
### GetPolScheduledLoadingLt

`func (o *PortcastAPIBillOfLading) GetPolScheduledLoadingLt() interface{}`

GetPolScheduledLoadingLt returns the PolScheduledLoadingLt field if non-nil, zero value otherwise.

### GetPolScheduledLoadingLtOk

`func (o *PortcastAPIBillOfLading) GetPolScheduledLoadingLtOk() (*interface{}, bool)`

GetPolScheduledLoadingLtOk returns a tuple with the PolScheduledLoadingLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolScheduledLoadingLt

`func (o *PortcastAPIBillOfLading) SetPolScheduledLoadingLt(v interface{})`

SetPolScheduledLoadingLt sets PolScheduledLoadingLt field to given value.

### HasPolScheduledLoadingLt

`func (o *PortcastAPIBillOfLading) HasPolScheduledLoadingLt() bool`

HasPolScheduledLoadingLt returns a boolean if a field has been set.

### SetPolScheduledLoadingLtNil

`func (o *PortcastAPIBillOfLading) SetPolScheduledLoadingLtNil(b bool)`

 SetPolScheduledLoadingLtNil sets the value for PolScheduledLoadingLt to be an explicit nil

### UnsetPolScheduledLoadingLt
`func (o *PortcastAPIBillOfLading) UnsetPolScheduledLoadingLt()`

UnsetPolScheduledLoadingLt ensures that no value is present for PolScheduledLoadingLt, not even an explicit nil
### GetPolTerminalName

`func (o *PortcastAPIBillOfLading) GetPolTerminalName() interface{}`

GetPolTerminalName returns the PolTerminalName field if non-nil, zero value otherwise.

### GetPolTerminalNameOk

`func (o *PortcastAPIBillOfLading) GetPolTerminalNameOk() (*interface{}, bool)`

GetPolTerminalNameOk returns a tuple with the PolTerminalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolTerminalName

`func (o *PortcastAPIBillOfLading) SetPolTerminalName(v interface{})`

SetPolTerminalName sets PolTerminalName field to given value.

### HasPolTerminalName

`func (o *PortcastAPIBillOfLading) HasPolTerminalName() bool`

HasPolTerminalName returns a boolean if a field has been set.

### SetPolTerminalNameNil

`func (o *PortcastAPIBillOfLading) SetPolTerminalNameNil(b bool)`

 SetPolTerminalNameNil sets the value for PolTerminalName to be an explicit nil

### UnsetPolTerminalName
`func (o *PortcastAPIBillOfLading) UnsetPolTerminalName()`

UnsetPolTerminalName ensures that no value is present for PolTerminalName, not even an explicit nil
### GetUpdated

`func (o *PortcastAPIBillOfLading) GetUpdated() interface{}`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PortcastAPIBillOfLading) GetUpdatedOk() (*interface{}, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PortcastAPIBillOfLading) SetUpdated(v interface{})`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PortcastAPIBillOfLading) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *PortcastAPIBillOfLading) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *PortcastAPIBillOfLading) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


