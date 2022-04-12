# TrackingEventVoyageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveScac** | Pointer to **string** |  | [optional] 
**ActualArrivalLt** | Pointer to **time.Time** |  | [optional] 
**ActualArrivalUtc** | Pointer to **time.Time** |  | [optional] 
**ActualDepartureLt** | Pointer to **time.Time** |  | [optional] 
**ActualDepartureUtc** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Lat** | Pointer to **float32** |  | [optional] 
**Lon** | Pointer to **float32** |  | [optional] 
**OriginalVoyageNo** | Pointer to **string** |  | [optional] 
**PortCode** | Pointer to **string** |  | [optional] 
**PortName** | Pointer to **string** |  | [optional] 
**PredictedArrivalLt** | Pointer to **string** |  | [optional] 
**PredictedArrivalUtc** | Pointer to **string** |  | [optional] 
**PredictedDepartureLt** | Pointer to **string** |  | [optional] 
**PredictedDepartureUtc** | Pointer to **string** |  | [optional] 
**PredictionTimeUtc** | Pointer to **string** |  | [optional] 
**ScheduledArrivalLt** | Pointer to **time.Time** |  | [optional] 
**ScheduledArrivalOnTimeConfidence** | Pointer to **string** |  | [optional] 
**ScheduledArrivalUtc** | Pointer to **time.Time** |  | [optional] 
**ScheduledDepartureLt** | Pointer to **time.Time** |  | [optional] 
**ScheduledDepartureOnTimeConfidence** | Pointer to **string** |  | [optional] 
**ScheduledDepartureUtc** | Pointer to **time.Time** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**VoyageNoList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTrackingEventVoyageDetails

`func NewTrackingEventVoyageDetails() *TrackingEventVoyageDetails`

NewTrackingEventVoyageDetails instantiates a new TrackingEventVoyageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingEventVoyageDetailsWithDefaults

`func NewTrackingEventVoyageDetailsWithDefaults() *TrackingEventVoyageDetails`

NewTrackingEventVoyageDetailsWithDefaults instantiates a new TrackingEventVoyageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveScac

`func (o *TrackingEventVoyageDetails) GetActiveScac() string`

GetActiveScac returns the ActiveScac field if non-nil, zero value otherwise.

### GetActiveScacOk

`func (o *TrackingEventVoyageDetails) GetActiveScacOk() (*string, bool)`

GetActiveScacOk returns a tuple with the ActiveScac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveScac

`func (o *TrackingEventVoyageDetails) SetActiveScac(v string)`

SetActiveScac sets ActiveScac field to given value.

### HasActiveScac

`func (o *TrackingEventVoyageDetails) HasActiveScac() bool`

HasActiveScac returns a boolean if a field has been set.

### GetActualArrivalLt

`func (o *TrackingEventVoyageDetails) GetActualArrivalLt() time.Time`

GetActualArrivalLt returns the ActualArrivalLt field if non-nil, zero value otherwise.

### GetActualArrivalLtOk

`func (o *TrackingEventVoyageDetails) GetActualArrivalLtOk() (*time.Time, bool)`

GetActualArrivalLtOk returns a tuple with the ActualArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualArrivalLt

`func (o *TrackingEventVoyageDetails) SetActualArrivalLt(v time.Time)`

SetActualArrivalLt sets ActualArrivalLt field to given value.

### HasActualArrivalLt

`func (o *TrackingEventVoyageDetails) HasActualArrivalLt() bool`

HasActualArrivalLt returns a boolean if a field has been set.

### GetActualArrivalUtc

`func (o *TrackingEventVoyageDetails) GetActualArrivalUtc() time.Time`

GetActualArrivalUtc returns the ActualArrivalUtc field if non-nil, zero value otherwise.

### GetActualArrivalUtcOk

`func (o *TrackingEventVoyageDetails) GetActualArrivalUtcOk() (*time.Time, bool)`

GetActualArrivalUtcOk returns a tuple with the ActualArrivalUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualArrivalUtc

`func (o *TrackingEventVoyageDetails) SetActualArrivalUtc(v time.Time)`

SetActualArrivalUtc sets ActualArrivalUtc field to given value.

### HasActualArrivalUtc

`func (o *TrackingEventVoyageDetails) HasActualArrivalUtc() bool`

HasActualArrivalUtc returns a boolean if a field has been set.

### GetActualDepartureLt

`func (o *TrackingEventVoyageDetails) GetActualDepartureLt() time.Time`

GetActualDepartureLt returns the ActualDepartureLt field if non-nil, zero value otherwise.

### GetActualDepartureLtOk

`func (o *TrackingEventVoyageDetails) GetActualDepartureLtOk() (*time.Time, bool)`

GetActualDepartureLtOk returns a tuple with the ActualDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDepartureLt

`func (o *TrackingEventVoyageDetails) SetActualDepartureLt(v time.Time)`

SetActualDepartureLt sets ActualDepartureLt field to given value.

### HasActualDepartureLt

`func (o *TrackingEventVoyageDetails) HasActualDepartureLt() bool`

HasActualDepartureLt returns a boolean if a field has been set.

### GetActualDepartureUtc

`func (o *TrackingEventVoyageDetails) GetActualDepartureUtc() time.Time`

GetActualDepartureUtc returns the ActualDepartureUtc field if non-nil, zero value otherwise.

### GetActualDepartureUtcOk

`func (o *TrackingEventVoyageDetails) GetActualDepartureUtcOk() (*time.Time, bool)`

GetActualDepartureUtcOk returns a tuple with the ActualDepartureUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDepartureUtc

`func (o *TrackingEventVoyageDetails) SetActualDepartureUtc(v time.Time)`

SetActualDepartureUtc sets ActualDepartureUtc field to given value.

### HasActualDepartureUtc

`func (o *TrackingEventVoyageDetails) HasActualDepartureUtc() bool`

HasActualDepartureUtc returns a boolean if a field has been set.

### GetId

`func (o *TrackingEventVoyageDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackingEventVoyageDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackingEventVoyageDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrackingEventVoyageDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *TrackingEventVoyageDetails) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *TrackingEventVoyageDetails) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *TrackingEventVoyageDetails) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *TrackingEventVoyageDetails) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetLat

`func (o *TrackingEventVoyageDetails) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *TrackingEventVoyageDetails) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *TrackingEventVoyageDetails) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *TrackingEventVoyageDetails) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLon

`func (o *TrackingEventVoyageDetails) GetLon() float32`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *TrackingEventVoyageDetails) GetLonOk() (*float32, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *TrackingEventVoyageDetails) SetLon(v float32)`

SetLon sets Lon field to given value.

### HasLon

`func (o *TrackingEventVoyageDetails) HasLon() bool`

HasLon returns a boolean if a field has been set.

### GetOriginalVoyageNo

`func (o *TrackingEventVoyageDetails) GetOriginalVoyageNo() string`

GetOriginalVoyageNo returns the OriginalVoyageNo field if non-nil, zero value otherwise.

### GetOriginalVoyageNoOk

`func (o *TrackingEventVoyageDetails) GetOriginalVoyageNoOk() (*string, bool)`

GetOriginalVoyageNoOk returns a tuple with the OriginalVoyageNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalVoyageNo

`func (o *TrackingEventVoyageDetails) SetOriginalVoyageNo(v string)`

SetOriginalVoyageNo sets OriginalVoyageNo field to given value.

### HasOriginalVoyageNo

`func (o *TrackingEventVoyageDetails) HasOriginalVoyageNo() bool`

HasOriginalVoyageNo returns a boolean if a field has been set.

### GetPortCode

`func (o *TrackingEventVoyageDetails) GetPortCode() string`

GetPortCode returns the PortCode field if non-nil, zero value otherwise.

### GetPortCodeOk

`func (o *TrackingEventVoyageDetails) GetPortCodeOk() (*string, bool)`

GetPortCodeOk returns a tuple with the PortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCode

`func (o *TrackingEventVoyageDetails) SetPortCode(v string)`

SetPortCode sets PortCode field to given value.

### HasPortCode

`func (o *TrackingEventVoyageDetails) HasPortCode() bool`

HasPortCode returns a boolean if a field has been set.

### GetPortName

`func (o *TrackingEventVoyageDetails) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *TrackingEventVoyageDetails) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *TrackingEventVoyageDetails) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *TrackingEventVoyageDetails) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPredictedArrivalLt

`func (o *TrackingEventVoyageDetails) GetPredictedArrivalLt() string`

GetPredictedArrivalLt returns the PredictedArrivalLt field if non-nil, zero value otherwise.

### GetPredictedArrivalLtOk

`func (o *TrackingEventVoyageDetails) GetPredictedArrivalLtOk() (*string, bool)`

GetPredictedArrivalLtOk returns a tuple with the PredictedArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedArrivalLt

`func (o *TrackingEventVoyageDetails) SetPredictedArrivalLt(v string)`

SetPredictedArrivalLt sets PredictedArrivalLt field to given value.

### HasPredictedArrivalLt

`func (o *TrackingEventVoyageDetails) HasPredictedArrivalLt() bool`

HasPredictedArrivalLt returns a boolean if a field has been set.

### GetPredictedArrivalUtc

`func (o *TrackingEventVoyageDetails) GetPredictedArrivalUtc() string`

GetPredictedArrivalUtc returns the PredictedArrivalUtc field if non-nil, zero value otherwise.

### GetPredictedArrivalUtcOk

`func (o *TrackingEventVoyageDetails) GetPredictedArrivalUtcOk() (*string, bool)`

GetPredictedArrivalUtcOk returns a tuple with the PredictedArrivalUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedArrivalUtc

`func (o *TrackingEventVoyageDetails) SetPredictedArrivalUtc(v string)`

SetPredictedArrivalUtc sets PredictedArrivalUtc field to given value.

### HasPredictedArrivalUtc

`func (o *TrackingEventVoyageDetails) HasPredictedArrivalUtc() bool`

HasPredictedArrivalUtc returns a boolean if a field has been set.

### GetPredictedDepartureLt

`func (o *TrackingEventVoyageDetails) GetPredictedDepartureLt() string`

GetPredictedDepartureLt returns the PredictedDepartureLt field if non-nil, zero value otherwise.

### GetPredictedDepartureLtOk

`func (o *TrackingEventVoyageDetails) GetPredictedDepartureLtOk() (*string, bool)`

GetPredictedDepartureLtOk returns a tuple with the PredictedDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedDepartureLt

`func (o *TrackingEventVoyageDetails) SetPredictedDepartureLt(v string)`

SetPredictedDepartureLt sets PredictedDepartureLt field to given value.

### HasPredictedDepartureLt

`func (o *TrackingEventVoyageDetails) HasPredictedDepartureLt() bool`

HasPredictedDepartureLt returns a boolean if a field has been set.

### GetPredictedDepartureUtc

`func (o *TrackingEventVoyageDetails) GetPredictedDepartureUtc() string`

GetPredictedDepartureUtc returns the PredictedDepartureUtc field if non-nil, zero value otherwise.

### GetPredictedDepartureUtcOk

`func (o *TrackingEventVoyageDetails) GetPredictedDepartureUtcOk() (*string, bool)`

GetPredictedDepartureUtcOk returns a tuple with the PredictedDepartureUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedDepartureUtc

`func (o *TrackingEventVoyageDetails) SetPredictedDepartureUtc(v string)`

SetPredictedDepartureUtc sets PredictedDepartureUtc field to given value.

### HasPredictedDepartureUtc

`func (o *TrackingEventVoyageDetails) HasPredictedDepartureUtc() bool`

HasPredictedDepartureUtc returns a boolean if a field has been set.

### GetPredictionTimeUtc

`func (o *TrackingEventVoyageDetails) GetPredictionTimeUtc() string`

GetPredictionTimeUtc returns the PredictionTimeUtc field if non-nil, zero value otherwise.

### GetPredictionTimeUtcOk

`func (o *TrackingEventVoyageDetails) GetPredictionTimeUtcOk() (*string, bool)`

GetPredictionTimeUtcOk returns a tuple with the PredictionTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionTimeUtc

`func (o *TrackingEventVoyageDetails) SetPredictionTimeUtc(v string)`

SetPredictionTimeUtc sets PredictionTimeUtc field to given value.

### HasPredictionTimeUtc

`func (o *TrackingEventVoyageDetails) HasPredictionTimeUtc() bool`

HasPredictionTimeUtc returns a boolean if a field has been set.

### GetScheduledArrivalLt

`func (o *TrackingEventVoyageDetails) GetScheduledArrivalLt() time.Time`

GetScheduledArrivalLt returns the ScheduledArrivalLt field if non-nil, zero value otherwise.

### GetScheduledArrivalLtOk

`func (o *TrackingEventVoyageDetails) GetScheduledArrivalLtOk() (*time.Time, bool)`

GetScheduledArrivalLtOk returns a tuple with the ScheduledArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledArrivalLt

`func (o *TrackingEventVoyageDetails) SetScheduledArrivalLt(v time.Time)`

SetScheduledArrivalLt sets ScheduledArrivalLt field to given value.

### HasScheduledArrivalLt

`func (o *TrackingEventVoyageDetails) HasScheduledArrivalLt() bool`

HasScheduledArrivalLt returns a boolean if a field has been set.

### GetScheduledArrivalOnTimeConfidence

`func (o *TrackingEventVoyageDetails) GetScheduledArrivalOnTimeConfidence() string`

GetScheduledArrivalOnTimeConfidence returns the ScheduledArrivalOnTimeConfidence field if non-nil, zero value otherwise.

### GetScheduledArrivalOnTimeConfidenceOk

`func (o *TrackingEventVoyageDetails) GetScheduledArrivalOnTimeConfidenceOk() (*string, bool)`

GetScheduledArrivalOnTimeConfidenceOk returns a tuple with the ScheduledArrivalOnTimeConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledArrivalOnTimeConfidence

`func (o *TrackingEventVoyageDetails) SetScheduledArrivalOnTimeConfidence(v string)`

SetScheduledArrivalOnTimeConfidence sets ScheduledArrivalOnTimeConfidence field to given value.

### HasScheduledArrivalOnTimeConfidence

`func (o *TrackingEventVoyageDetails) HasScheduledArrivalOnTimeConfidence() bool`

HasScheduledArrivalOnTimeConfidence returns a boolean if a field has been set.

### GetScheduledArrivalUtc

`func (o *TrackingEventVoyageDetails) GetScheduledArrivalUtc() time.Time`

GetScheduledArrivalUtc returns the ScheduledArrivalUtc field if non-nil, zero value otherwise.

### GetScheduledArrivalUtcOk

`func (o *TrackingEventVoyageDetails) GetScheduledArrivalUtcOk() (*time.Time, bool)`

GetScheduledArrivalUtcOk returns a tuple with the ScheduledArrivalUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledArrivalUtc

`func (o *TrackingEventVoyageDetails) SetScheduledArrivalUtc(v time.Time)`

SetScheduledArrivalUtc sets ScheduledArrivalUtc field to given value.

### HasScheduledArrivalUtc

`func (o *TrackingEventVoyageDetails) HasScheduledArrivalUtc() bool`

HasScheduledArrivalUtc returns a boolean if a field has been set.

### GetScheduledDepartureLt

`func (o *TrackingEventVoyageDetails) GetScheduledDepartureLt() time.Time`

GetScheduledDepartureLt returns the ScheduledDepartureLt field if non-nil, zero value otherwise.

### GetScheduledDepartureLtOk

`func (o *TrackingEventVoyageDetails) GetScheduledDepartureLtOk() (*time.Time, bool)`

GetScheduledDepartureLtOk returns a tuple with the ScheduledDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDepartureLt

`func (o *TrackingEventVoyageDetails) SetScheduledDepartureLt(v time.Time)`

SetScheduledDepartureLt sets ScheduledDepartureLt field to given value.

### HasScheduledDepartureLt

`func (o *TrackingEventVoyageDetails) HasScheduledDepartureLt() bool`

HasScheduledDepartureLt returns a boolean if a field has been set.

### GetScheduledDepartureOnTimeConfidence

`func (o *TrackingEventVoyageDetails) GetScheduledDepartureOnTimeConfidence() string`

GetScheduledDepartureOnTimeConfidence returns the ScheduledDepartureOnTimeConfidence field if non-nil, zero value otherwise.

### GetScheduledDepartureOnTimeConfidenceOk

`func (o *TrackingEventVoyageDetails) GetScheduledDepartureOnTimeConfidenceOk() (*string, bool)`

GetScheduledDepartureOnTimeConfidenceOk returns a tuple with the ScheduledDepartureOnTimeConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDepartureOnTimeConfidence

`func (o *TrackingEventVoyageDetails) SetScheduledDepartureOnTimeConfidence(v string)`

SetScheduledDepartureOnTimeConfidence sets ScheduledDepartureOnTimeConfidence field to given value.

### HasScheduledDepartureOnTimeConfidence

`func (o *TrackingEventVoyageDetails) HasScheduledDepartureOnTimeConfidence() bool`

HasScheduledDepartureOnTimeConfidence returns a boolean if a field has been set.

### GetScheduledDepartureUtc

`func (o *TrackingEventVoyageDetails) GetScheduledDepartureUtc() time.Time`

GetScheduledDepartureUtc returns the ScheduledDepartureUtc field if non-nil, zero value otherwise.

### GetScheduledDepartureUtcOk

`func (o *TrackingEventVoyageDetails) GetScheduledDepartureUtcOk() (*time.Time, bool)`

GetScheduledDepartureUtcOk returns a tuple with the ScheduledDepartureUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDepartureUtc

`func (o *TrackingEventVoyageDetails) SetScheduledDepartureUtc(v time.Time)`

SetScheduledDepartureUtc sets ScheduledDepartureUtc field to given value.

### HasScheduledDepartureUtc

`func (o *TrackingEventVoyageDetails) HasScheduledDepartureUtc() bool`

HasScheduledDepartureUtc returns a boolean if a field has been set.

### GetTimezone

`func (o *TrackingEventVoyageDetails) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *TrackingEventVoyageDetails) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *TrackingEventVoyageDetails) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *TrackingEventVoyageDetails) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetVoyageNoList

`func (o *TrackingEventVoyageDetails) GetVoyageNoList() []string`

GetVoyageNoList returns the VoyageNoList field if non-nil, zero value otherwise.

### GetVoyageNoListOk

`func (o *TrackingEventVoyageDetails) GetVoyageNoListOk() (*[]string, bool)`

GetVoyageNoListOk returns a tuple with the VoyageNoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageNoList

`func (o *TrackingEventVoyageDetails) SetVoyageNoList(v []string)`

SetVoyageNoList sets VoyageNoList field to given value.

### HasVoyageNoList

`func (o *TrackingEventVoyageDetails) HasVoyageNoList() bool`

HasVoyageNoList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


