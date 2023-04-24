# TrackingEventSailingInfoTracking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualArrivalLt** | Pointer to **string** |  | [optional] 
**ActualArrivalUtc** | Pointer to **string** |  | [optional] 
**Ais** | Pointer to [**TrackingEventAis**](TrackingEventAis.md) |  | [optional] 
**PredictedArrivalLt** | Pointer to **string** |  | [optional] 
**PredictedArrivalUtc** | Pointer to **string** |  | [optional] 
**PredictionTimeUtc** | Pointer to **string** |  | [optional] 
**SailingInfo** | Pointer to [**TrackingEventSailingInfo**](TrackingEventSailingInfo.md) |  | [optional] 
**ScheduledArrivalLt** | Pointer to **string** |  | [optional] 
**ScheduledArrivalUtc** | Pointer to **string** |  | [optional] 
**TargetPortCode** | Pointer to **string** |  | [optional] 
**TargetPortName** | Pointer to **string** |  | [optional] 
**VoyageDetails** | Pointer to [**[]TrackingEventVoyageDetails**](TrackingEventVoyageDetails.md) |  | [optional] 
**VoyageNoList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTrackingEventSailingInfoTracking

`func NewTrackingEventSailingInfoTracking() *TrackingEventSailingInfoTracking`

NewTrackingEventSailingInfoTracking instantiates a new TrackingEventSailingInfoTracking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingEventSailingInfoTrackingWithDefaults

`func NewTrackingEventSailingInfoTrackingWithDefaults() *TrackingEventSailingInfoTracking`

NewTrackingEventSailingInfoTrackingWithDefaults instantiates a new TrackingEventSailingInfoTracking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualArrivalLt

`func (o *TrackingEventSailingInfoTracking) GetActualArrivalLt() string`

GetActualArrivalLt returns the ActualArrivalLt field if non-nil, zero value otherwise.

### GetActualArrivalLtOk

`func (o *TrackingEventSailingInfoTracking) GetActualArrivalLtOk() (*string, bool)`

GetActualArrivalLtOk returns a tuple with the ActualArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualArrivalLt

`func (o *TrackingEventSailingInfoTracking) SetActualArrivalLt(v string)`

SetActualArrivalLt sets ActualArrivalLt field to given value.

### HasActualArrivalLt

`func (o *TrackingEventSailingInfoTracking) HasActualArrivalLt() bool`

HasActualArrivalLt returns a boolean if a field has been set.

### GetActualArrivalUtc

`func (o *TrackingEventSailingInfoTracking) GetActualArrivalUtc() string`

GetActualArrivalUtc returns the ActualArrivalUtc field if non-nil, zero value otherwise.

### GetActualArrivalUtcOk

`func (o *TrackingEventSailingInfoTracking) GetActualArrivalUtcOk() (*string, bool)`

GetActualArrivalUtcOk returns a tuple with the ActualArrivalUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualArrivalUtc

`func (o *TrackingEventSailingInfoTracking) SetActualArrivalUtc(v string)`

SetActualArrivalUtc sets ActualArrivalUtc field to given value.

### HasActualArrivalUtc

`func (o *TrackingEventSailingInfoTracking) HasActualArrivalUtc() bool`

HasActualArrivalUtc returns a boolean if a field has been set.

### GetAis

`func (o *TrackingEventSailingInfoTracking) GetAis() TrackingEventAis`

GetAis returns the Ais field if non-nil, zero value otherwise.

### GetAisOk

`func (o *TrackingEventSailingInfoTracking) GetAisOk() (*TrackingEventAis, bool)`

GetAisOk returns a tuple with the Ais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAis

`func (o *TrackingEventSailingInfoTracking) SetAis(v TrackingEventAis)`

SetAis sets Ais field to given value.

### HasAis

`func (o *TrackingEventSailingInfoTracking) HasAis() bool`

HasAis returns a boolean if a field has been set.

### GetPredictedArrivalLt

`func (o *TrackingEventSailingInfoTracking) GetPredictedArrivalLt() string`

GetPredictedArrivalLt returns the PredictedArrivalLt field if non-nil, zero value otherwise.

### GetPredictedArrivalLtOk

`func (o *TrackingEventSailingInfoTracking) GetPredictedArrivalLtOk() (*string, bool)`

GetPredictedArrivalLtOk returns a tuple with the PredictedArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedArrivalLt

`func (o *TrackingEventSailingInfoTracking) SetPredictedArrivalLt(v string)`

SetPredictedArrivalLt sets PredictedArrivalLt field to given value.

### HasPredictedArrivalLt

`func (o *TrackingEventSailingInfoTracking) HasPredictedArrivalLt() bool`

HasPredictedArrivalLt returns a boolean if a field has been set.

### GetPredictedArrivalUtc

`func (o *TrackingEventSailingInfoTracking) GetPredictedArrivalUtc() string`

GetPredictedArrivalUtc returns the PredictedArrivalUtc field if non-nil, zero value otherwise.

### GetPredictedArrivalUtcOk

`func (o *TrackingEventSailingInfoTracking) GetPredictedArrivalUtcOk() (*string, bool)`

GetPredictedArrivalUtcOk returns a tuple with the PredictedArrivalUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedArrivalUtc

`func (o *TrackingEventSailingInfoTracking) SetPredictedArrivalUtc(v string)`

SetPredictedArrivalUtc sets PredictedArrivalUtc field to given value.

### HasPredictedArrivalUtc

`func (o *TrackingEventSailingInfoTracking) HasPredictedArrivalUtc() bool`

HasPredictedArrivalUtc returns a boolean if a field has been set.

### GetPredictionTimeUtc

`func (o *TrackingEventSailingInfoTracking) GetPredictionTimeUtc() string`

GetPredictionTimeUtc returns the PredictionTimeUtc field if non-nil, zero value otherwise.

### GetPredictionTimeUtcOk

`func (o *TrackingEventSailingInfoTracking) GetPredictionTimeUtcOk() (*string, bool)`

GetPredictionTimeUtcOk returns a tuple with the PredictionTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionTimeUtc

`func (o *TrackingEventSailingInfoTracking) SetPredictionTimeUtc(v string)`

SetPredictionTimeUtc sets PredictionTimeUtc field to given value.

### HasPredictionTimeUtc

`func (o *TrackingEventSailingInfoTracking) HasPredictionTimeUtc() bool`

HasPredictionTimeUtc returns a boolean if a field has been set.

### GetSailingInfo

`func (o *TrackingEventSailingInfoTracking) GetSailingInfo() TrackingEventSailingInfo`

GetSailingInfo returns the SailingInfo field if non-nil, zero value otherwise.

### GetSailingInfoOk

`func (o *TrackingEventSailingInfoTracking) GetSailingInfoOk() (*TrackingEventSailingInfo, bool)`

GetSailingInfoOk returns a tuple with the SailingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSailingInfo

`func (o *TrackingEventSailingInfoTracking) SetSailingInfo(v TrackingEventSailingInfo)`

SetSailingInfo sets SailingInfo field to given value.

### HasSailingInfo

`func (o *TrackingEventSailingInfoTracking) HasSailingInfo() bool`

HasSailingInfo returns a boolean if a field has been set.

### GetScheduledArrivalLt

`func (o *TrackingEventSailingInfoTracking) GetScheduledArrivalLt() string`

GetScheduledArrivalLt returns the ScheduledArrivalLt field if non-nil, zero value otherwise.

### GetScheduledArrivalLtOk

`func (o *TrackingEventSailingInfoTracking) GetScheduledArrivalLtOk() (*string, bool)`

GetScheduledArrivalLtOk returns a tuple with the ScheduledArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledArrivalLt

`func (o *TrackingEventSailingInfoTracking) SetScheduledArrivalLt(v string)`

SetScheduledArrivalLt sets ScheduledArrivalLt field to given value.

### HasScheduledArrivalLt

`func (o *TrackingEventSailingInfoTracking) HasScheduledArrivalLt() bool`

HasScheduledArrivalLt returns a boolean if a field has been set.

### GetScheduledArrivalUtc

`func (o *TrackingEventSailingInfoTracking) GetScheduledArrivalUtc() string`

GetScheduledArrivalUtc returns the ScheduledArrivalUtc field if non-nil, zero value otherwise.

### GetScheduledArrivalUtcOk

`func (o *TrackingEventSailingInfoTracking) GetScheduledArrivalUtcOk() (*string, bool)`

GetScheduledArrivalUtcOk returns a tuple with the ScheduledArrivalUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledArrivalUtc

`func (o *TrackingEventSailingInfoTracking) SetScheduledArrivalUtc(v string)`

SetScheduledArrivalUtc sets ScheduledArrivalUtc field to given value.

### HasScheduledArrivalUtc

`func (o *TrackingEventSailingInfoTracking) HasScheduledArrivalUtc() bool`

HasScheduledArrivalUtc returns a boolean if a field has been set.

### GetTargetPortCode

`func (o *TrackingEventSailingInfoTracking) GetTargetPortCode() string`

GetTargetPortCode returns the TargetPortCode field if non-nil, zero value otherwise.

### GetTargetPortCodeOk

`func (o *TrackingEventSailingInfoTracking) GetTargetPortCodeOk() (*string, bool)`

GetTargetPortCodeOk returns a tuple with the TargetPortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPortCode

`func (o *TrackingEventSailingInfoTracking) SetTargetPortCode(v string)`

SetTargetPortCode sets TargetPortCode field to given value.

### HasTargetPortCode

`func (o *TrackingEventSailingInfoTracking) HasTargetPortCode() bool`

HasTargetPortCode returns a boolean if a field has been set.

### GetTargetPortName

`func (o *TrackingEventSailingInfoTracking) GetTargetPortName() string`

GetTargetPortName returns the TargetPortName field if non-nil, zero value otherwise.

### GetTargetPortNameOk

`func (o *TrackingEventSailingInfoTracking) GetTargetPortNameOk() (*string, bool)`

GetTargetPortNameOk returns a tuple with the TargetPortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPortName

`func (o *TrackingEventSailingInfoTracking) SetTargetPortName(v string)`

SetTargetPortName sets TargetPortName field to given value.

### HasTargetPortName

`func (o *TrackingEventSailingInfoTracking) HasTargetPortName() bool`

HasTargetPortName returns a boolean if a field has been set.

### GetVoyageDetails

`func (o *TrackingEventSailingInfoTracking) GetVoyageDetails() []TrackingEventVoyageDetails`

GetVoyageDetails returns the VoyageDetails field if non-nil, zero value otherwise.

### GetVoyageDetailsOk

`func (o *TrackingEventSailingInfoTracking) GetVoyageDetailsOk() (*[]TrackingEventVoyageDetails, bool)`

GetVoyageDetailsOk returns a tuple with the VoyageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageDetails

`func (o *TrackingEventSailingInfoTracking) SetVoyageDetails(v []TrackingEventVoyageDetails)`

SetVoyageDetails sets VoyageDetails field to given value.

### HasVoyageDetails

`func (o *TrackingEventSailingInfoTracking) HasVoyageDetails() bool`

HasVoyageDetails returns a boolean if a field has been set.

### GetVoyageNoList

`func (o *TrackingEventSailingInfoTracking) GetVoyageNoList() []string`

GetVoyageNoList returns the VoyageNoList field if non-nil, zero value otherwise.

### GetVoyageNoListOk

`func (o *TrackingEventSailingInfoTracking) GetVoyageNoListOk() (*[]string, bool)`

GetVoyageNoListOk returns a tuple with the VoyageNoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageNoList

`func (o *TrackingEventSailingInfoTracking) SetVoyageNoList(v []string)`

SetVoyageNoList sets VoyageNoList field to given value.

### HasVoyageNoList

`func (o *TrackingEventSailingInfoTracking) HasVoyageNoList() bool`

HasVoyageNoList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


