# VoyageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveScac** | Pointer to **NullableString** | Vessel Schedule Provider SCAC | [optional] 
**ActualArrivalLt** | Pointer to **NullableTime** | Actual Time of Arrival - Local Time | [optional] 
**ActualArrivalUtc** | Pointer to **NullableTime** | Actual Time of Arrival - UTC Adjusted Time | [optional] 
**ActualDepartureLt** | Pointer to **NullableTime** | Actual Time of Departure - Local Time | [optional] 
**ActualDepartureUtc** | Pointer to **NullableTime** | Actual Time of Departure - UTC Adjusted Time | [optional] 
**DelayReasons** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** | Unique Identifier for the Voyage Detail | [optional] 
**Index** | Pointer to **int32** | Index of the Journey Position | [optional] 
**Lat** | Pointer to **float32** | Latitude of the Port Location | [optional] 
**Lon** | Pointer to **float32** | Longitude of the Port Location | [optional] 
**PortCode** | Pointer to **string** | Port Location UNLOCODE | [optional] 
**PortName** | Pointer to **string** | Port Location Name | [optional] 
**PredictedArrivalLt** | Pointer to **NullableTime** | Portcast Predicted Arrival Time - Local Time | [optional] 
**PredictedArrivalUtc** | Pointer to **NullableTime** | Portcast Predicted Arrival Time - UTC Adjusted | [optional] 
**PredictedDepartureLt** | Pointer to **NullableTime** | Portcast Predicted Departure Time - Local Time | [optional] 
**PredictedDepartureUtc** | Pointer to **NullableTime** | Portcast Predicted Departure Time - UTC Adjusted Time | [optional] 
**PredictionTimeUtc** | Pointer to **NullableTime** | Prediction Generation Timestamp | [optional] 
**ScheduledArrivalLt** | Pointer to **NullableTime** | Vessel Scheduled Arrival Time - Local Time | [optional] 
**ScheduledArrivalUtc** | Pointer to **NullableTime** | Vessel Scheduled Arrival Time - UTC Adjusted Time | [optional] 
**ScheduledDepartureLt** | Pointer to **NullableTime** | Vessel Scheduled Departure Time - Local Time | [optional] 
**ScheduledDepartureUtc** | Pointer to **NullableTime** | Vessel Scheduled Departure Time - Departure Time | [optional] 
**Timezone** | Pointer to **NullableString** | Timezone of the Port Location | [optional] 
**VoyageNoList** | Pointer to **[]string** | Voyage Numbers at the Port Location (Comes from schedule for that leg ; first one is import voyage number, second one is export voyage no) | [optional] 

## Methods

### NewVoyageDetails

`func NewVoyageDetails() *VoyageDetails`

NewVoyageDetails instantiates a new VoyageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoyageDetailsWithDefaults

`func NewVoyageDetailsWithDefaults() *VoyageDetails`

NewVoyageDetailsWithDefaults instantiates a new VoyageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveScac

`func (o *VoyageDetails) GetActiveScac() string`

GetActiveScac returns the ActiveScac field if non-nil, zero value otherwise.

### GetActiveScacOk

`func (o *VoyageDetails) GetActiveScacOk() (*string, bool)`

GetActiveScacOk returns a tuple with the ActiveScac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveScac

`func (o *VoyageDetails) SetActiveScac(v string)`

SetActiveScac sets ActiveScac field to given value.

### HasActiveScac

`func (o *VoyageDetails) HasActiveScac() bool`

HasActiveScac returns a boolean if a field has been set.

### SetActiveScacNil

`func (o *VoyageDetails) SetActiveScacNil(b bool)`

 SetActiveScacNil sets the value for ActiveScac to be an explicit nil

### UnsetActiveScac
`func (o *VoyageDetails) UnsetActiveScac()`

UnsetActiveScac ensures that no value is present for ActiveScac, not even an explicit nil
### GetActualArrivalLt

`func (o *VoyageDetails) GetActualArrivalLt() time.Time`

GetActualArrivalLt returns the ActualArrivalLt field if non-nil, zero value otherwise.

### GetActualArrivalLtOk

`func (o *VoyageDetails) GetActualArrivalLtOk() (*time.Time, bool)`

GetActualArrivalLtOk returns a tuple with the ActualArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualArrivalLt

`func (o *VoyageDetails) SetActualArrivalLt(v time.Time)`

SetActualArrivalLt sets ActualArrivalLt field to given value.

### HasActualArrivalLt

`func (o *VoyageDetails) HasActualArrivalLt() bool`

HasActualArrivalLt returns a boolean if a field has been set.

### SetActualArrivalLtNil

`func (o *VoyageDetails) SetActualArrivalLtNil(b bool)`

 SetActualArrivalLtNil sets the value for ActualArrivalLt to be an explicit nil

### UnsetActualArrivalLt
`func (o *VoyageDetails) UnsetActualArrivalLt()`

UnsetActualArrivalLt ensures that no value is present for ActualArrivalLt, not even an explicit nil
### GetActualArrivalUtc

`func (o *VoyageDetails) GetActualArrivalUtc() time.Time`

GetActualArrivalUtc returns the ActualArrivalUtc field if non-nil, zero value otherwise.

### GetActualArrivalUtcOk

`func (o *VoyageDetails) GetActualArrivalUtcOk() (*time.Time, bool)`

GetActualArrivalUtcOk returns a tuple with the ActualArrivalUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualArrivalUtc

`func (o *VoyageDetails) SetActualArrivalUtc(v time.Time)`

SetActualArrivalUtc sets ActualArrivalUtc field to given value.

### HasActualArrivalUtc

`func (o *VoyageDetails) HasActualArrivalUtc() bool`

HasActualArrivalUtc returns a boolean if a field has been set.

### SetActualArrivalUtcNil

`func (o *VoyageDetails) SetActualArrivalUtcNil(b bool)`

 SetActualArrivalUtcNil sets the value for ActualArrivalUtc to be an explicit nil

### UnsetActualArrivalUtc
`func (o *VoyageDetails) UnsetActualArrivalUtc()`

UnsetActualArrivalUtc ensures that no value is present for ActualArrivalUtc, not even an explicit nil
### GetActualDepartureLt

`func (o *VoyageDetails) GetActualDepartureLt() time.Time`

GetActualDepartureLt returns the ActualDepartureLt field if non-nil, zero value otherwise.

### GetActualDepartureLtOk

`func (o *VoyageDetails) GetActualDepartureLtOk() (*time.Time, bool)`

GetActualDepartureLtOk returns a tuple with the ActualDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDepartureLt

`func (o *VoyageDetails) SetActualDepartureLt(v time.Time)`

SetActualDepartureLt sets ActualDepartureLt field to given value.

### HasActualDepartureLt

`func (o *VoyageDetails) HasActualDepartureLt() bool`

HasActualDepartureLt returns a boolean if a field has been set.

### SetActualDepartureLtNil

`func (o *VoyageDetails) SetActualDepartureLtNil(b bool)`

 SetActualDepartureLtNil sets the value for ActualDepartureLt to be an explicit nil

### UnsetActualDepartureLt
`func (o *VoyageDetails) UnsetActualDepartureLt()`

UnsetActualDepartureLt ensures that no value is present for ActualDepartureLt, not even an explicit nil
### GetActualDepartureUtc

`func (o *VoyageDetails) GetActualDepartureUtc() time.Time`

GetActualDepartureUtc returns the ActualDepartureUtc field if non-nil, zero value otherwise.

### GetActualDepartureUtcOk

`func (o *VoyageDetails) GetActualDepartureUtcOk() (*time.Time, bool)`

GetActualDepartureUtcOk returns a tuple with the ActualDepartureUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDepartureUtc

`func (o *VoyageDetails) SetActualDepartureUtc(v time.Time)`

SetActualDepartureUtc sets ActualDepartureUtc field to given value.

### HasActualDepartureUtc

`func (o *VoyageDetails) HasActualDepartureUtc() bool`

HasActualDepartureUtc returns a boolean if a field has been set.

### SetActualDepartureUtcNil

`func (o *VoyageDetails) SetActualDepartureUtcNil(b bool)`

 SetActualDepartureUtcNil sets the value for ActualDepartureUtc to be an explicit nil

### UnsetActualDepartureUtc
`func (o *VoyageDetails) UnsetActualDepartureUtc()`

UnsetActualDepartureUtc ensures that no value is present for ActualDepartureUtc, not even an explicit nil
### GetDelayReasons

`func (o *VoyageDetails) GetDelayReasons() []string`

GetDelayReasons returns the DelayReasons field if non-nil, zero value otherwise.

### GetDelayReasonsOk

`func (o *VoyageDetails) GetDelayReasonsOk() (*[]string, bool)`

GetDelayReasonsOk returns a tuple with the DelayReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayReasons

`func (o *VoyageDetails) SetDelayReasons(v []string)`

SetDelayReasons sets DelayReasons field to given value.

### HasDelayReasons

`func (o *VoyageDetails) HasDelayReasons() bool`

HasDelayReasons returns a boolean if a field has been set.

### GetId

`func (o *VoyageDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoyageDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoyageDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoyageDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *VoyageDetails) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *VoyageDetails) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *VoyageDetails) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *VoyageDetails) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetLat

`func (o *VoyageDetails) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *VoyageDetails) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *VoyageDetails) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *VoyageDetails) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLon

`func (o *VoyageDetails) GetLon() float32`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *VoyageDetails) GetLonOk() (*float32, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *VoyageDetails) SetLon(v float32)`

SetLon sets Lon field to given value.

### HasLon

`func (o *VoyageDetails) HasLon() bool`

HasLon returns a boolean if a field has been set.

### GetPortCode

`func (o *VoyageDetails) GetPortCode() string`

GetPortCode returns the PortCode field if non-nil, zero value otherwise.

### GetPortCodeOk

`func (o *VoyageDetails) GetPortCodeOk() (*string, bool)`

GetPortCodeOk returns a tuple with the PortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCode

`func (o *VoyageDetails) SetPortCode(v string)`

SetPortCode sets PortCode field to given value.

### HasPortCode

`func (o *VoyageDetails) HasPortCode() bool`

HasPortCode returns a boolean if a field has been set.

### GetPortName

`func (o *VoyageDetails) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *VoyageDetails) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *VoyageDetails) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *VoyageDetails) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPredictedArrivalLt

`func (o *VoyageDetails) GetPredictedArrivalLt() time.Time`

GetPredictedArrivalLt returns the PredictedArrivalLt field if non-nil, zero value otherwise.

### GetPredictedArrivalLtOk

`func (o *VoyageDetails) GetPredictedArrivalLtOk() (*time.Time, bool)`

GetPredictedArrivalLtOk returns a tuple with the PredictedArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedArrivalLt

`func (o *VoyageDetails) SetPredictedArrivalLt(v time.Time)`

SetPredictedArrivalLt sets PredictedArrivalLt field to given value.

### HasPredictedArrivalLt

`func (o *VoyageDetails) HasPredictedArrivalLt() bool`

HasPredictedArrivalLt returns a boolean if a field has been set.

### SetPredictedArrivalLtNil

`func (o *VoyageDetails) SetPredictedArrivalLtNil(b bool)`

 SetPredictedArrivalLtNil sets the value for PredictedArrivalLt to be an explicit nil

### UnsetPredictedArrivalLt
`func (o *VoyageDetails) UnsetPredictedArrivalLt()`

UnsetPredictedArrivalLt ensures that no value is present for PredictedArrivalLt, not even an explicit nil
### GetPredictedArrivalUtc

`func (o *VoyageDetails) GetPredictedArrivalUtc() time.Time`

GetPredictedArrivalUtc returns the PredictedArrivalUtc field if non-nil, zero value otherwise.

### GetPredictedArrivalUtcOk

`func (o *VoyageDetails) GetPredictedArrivalUtcOk() (*time.Time, bool)`

GetPredictedArrivalUtcOk returns a tuple with the PredictedArrivalUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedArrivalUtc

`func (o *VoyageDetails) SetPredictedArrivalUtc(v time.Time)`

SetPredictedArrivalUtc sets PredictedArrivalUtc field to given value.

### HasPredictedArrivalUtc

`func (o *VoyageDetails) HasPredictedArrivalUtc() bool`

HasPredictedArrivalUtc returns a boolean if a field has been set.

### SetPredictedArrivalUtcNil

`func (o *VoyageDetails) SetPredictedArrivalUtcNil(b bool)`

 SetPredictedArrivalUtcNil sets the value for PredictedArrivalUtc to be an explicit nil

### UnsetPredictedArrivalUtc
`func (o *VoyageDetails) UnsetPredictedArrivalUtc()`

UnsetPredictedArrivalUtc ensures that no value is present for PredictedArrivalUtc, not even an explicit nil
### GetPredictedDepartureLt

`func (o *VoyageDetails) GetPredictedDepartureLt() time.Time`

GetPredictedDepartureLt returns the PredictedDepartureLt field if non-nil, zero value otherwise.

### GetPredictedDepartureLtOk

`func (o *VoyageDetails) GetPredictedDepartureLtOk() (*time.Time, bool)`

GetPredictedDepartureLtOk returns a tuple with the PredictedDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedDepartureLt

`func (o *VoyageDetails) SetPredictedDepartureLt(v time.Time)`

SetPredictedDepartureLt sets PredictedDepartureLt field to given value.

### HasPredictedDepartureLt

`func (o *VoyageDetails) HasPredictedDepartureLt() bool`

HasPredictedDepartureLt returns a boolean if a field has been set.

### SetPredictedDepartureLtNil

`func (o *VoyageDetails) SetPredictedDepartureLtNil(b bool)`

 SetPredictedDepartureLtNil sets the value for PredictedDepartureLt to be an explicit nil

### UnsetPredictedDepartureLt
`func (o *VoyageDetails) UnsetPredictedDepartureLt()`

UnsetPredictedDepartureLt ensures that no value is present for PredictedDepartureLt, not even an explicit nil
### GetPredictedDepartureUtc

`func (o *VoyageDetails) GetPredictedDepartureUtc() time.Time`

GetPredictedDepartureUtc returns the PredictedDepartureUtc field if non-nil, zero value otherwise.

### GetPredictedDepartureUtcOk

`func (o *VoyageDetails) GetPredictedDepartureUtcOk() (*time.Time, bool)`

GetPredictedDepartureUtcOk returns a tuple with the PredictedDepartureUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedDepartureUtc

`func (o *VoyageDetails) SetPredictedDepartureUtc(v time.Time)`

SetPredictedDepartureUtc sets PredictedDepartureUtc field to given value.

### HasPredictedDepartureUtc

`func (o *VoyageDetails) HasPredictedDepartureUtc() bool`

HasPredictedDepartureUtc returns a boolean if a field has been set.

### SetPredictedDepartureUtcNil

`func (o *VoyageDetails) SetPredictedDepartureUtcNil(b bool)`

 SetPredictedDepartureUtcNil sets the value for PredictedDepartureUtc to be an explicit nil

### UnsetPredictedDepartureUtc
`func (o *VoyageDetails) UnsetPredictedDepartureUtc()`

UnsetPredictedDepartureUtc ensures that no value is present for PredictedDepartureUtc, not even an explicit nil
### GetPredictionTimeUtc

`func (o *VoyageDetails) GetPredictionTimeUtc() time.Time`

GetPredictionTimeUtc returns the PredictionTimeUtc field if non-nil, zero value otherwise.

### GetPredictionTimeUtcOk

`func (o *VoyageDetails) GetPredictionTimeUtcOk() (*time.Time, bool)`

GetPredictionTimeUtcOk returns a tuple with the PredictionTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionTimeUtc

`func (o *VoyageDetails) SetPredictionTimeUtc(v time.Time)`

SetPredictionTimeUtc sets PredictionTimeUtc field to given value.

### HasPredictionTimeUtc

`func (o *VoyageDetails) HasPredictionTimeUtc() bool`

HasPredictionTimeUtc returns a boolean if a field has been set.

### SetPredictionTimeUtcNil

`func (o *VoyageDetails) SetPredictionTimeUtcNil(b bool)`

 SetPredictionTimeUtcNil sets the value for PredictionTimeUtc to be an explicit nil

### UnsetPredictionTimeUtc
`func (o *VoyageDetails) UnsetPredictionTimeUtc()`

UnsetPredictionTimeUtc ensures that no value is present for PredictionTimeUtc, not even an explicit nil
### GetScheduledArrivalLt

`func (o *VoyageDetails) GetScheduledArrivalLt() time.Time`

GetScheduledArrivalLt returns the ScheduledArrivalLt field if non-nil, zero value otherwise.

### GetScheduledArrivalLtOk

`func (o *VoyageDetails) GetScheduledArrivalLtOk() (*time.Time, bool)`

GetScheduledArrivalLtOk returns a tuple with the ScheduledArrivalLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledArrivalLt

`func (o *VoyageDetails) SetScheduledArrivalLt(v time.Time)`

SetScheduledArrivalLt sets ScheduledArrivalLt field to given value.

### HasScheduledArrivalLt

`func (o *VoyageDetails) HasScheduledArrivalLt() bool`

HasScheduledArrivalLt returns a boolean if a field has been set.

### SetScheduledArrivalLtNil

`func (o *VoyageDetails) SetScheduledArrivalLtNil(b bool)`

 SetScheduledArrivalLtNil sets the value for ScheduledArrivalLt to be an explicit nil

### UnsetScheduledArrivalLt
`func (o *VoyageDetails) UnsetScheduledArrivalLt()`

UnsetScheduledArrivalLt ensures that no value is present for ScheduledArrivalLt, not even an explicit nil
### GetScheduledArrivalUtc

`func (o *VoyageDetails) GetScheduledArrivalUtc() time.Time`

GetScheduledArrivalUtc returns the ScheduledArrivalUtc field if non-nil, zero value otherwise.

### GetScheduledArrivalUtcOk

`func (o *VoyageDetails) GetScheduledArrivalUtcOk() (*time.Time, bool)`

GetScheduledArrivalUtcOk returns a tuple with the ScheduledArrivalUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledArrivalUtc

`func (o *VoyageDetails) SetScheduledArrivalUtc(v time.Time)`

SetScheduledArrivalUtc sets ScheduledArrivalUtc field to given value.

### HasScheduledArrivalUtc

`func (o *VoyageDetails) HasScheduledArrivalUtc() bool`

HasScheduledArrivalUtc returns a boolean if a field has been set.

### SetScheduledArrivalUtcNil

`func (o *VoyageDetails) SetScheduledArrivalUtcNil(b bool)`

 SetScheduledArrivalUtcNil sets the value for ScheduledArrivalUtc to be an explicit nil

### UnsetScheduledArrivalUtc
`func (o *VoyageDetails) UnsetScheduledArrivalUtc()`

UnsetScheduledArrivalUtc ensures that no value is present for ScheduledArrivalUtc, not even an explicit nil
### GetScheduledDepartureLt

`func (o *VoyageDetails) GetScheduledDepartureLt() time.Time`

GetScheduledDepartureLt returns the ScheduledDepartureLt field if non-nil, zero value otherwise.

### GetScheduledDepartureLtOk

`func (o *VoyageDetails) GetScheduledDepartureLtOk() (*time.Time, bool)`

GetScheduledDepartureLtOk returns a tuple with the ScheduledDepartureLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDepartureLt

`func (o *VoyageDetails) SetScheduledDepartureLt(v time.Time)`

SetScheduledDepartureLt sets ScheduledDepartureLt field to given value.

### HasScheduledDepartureLt

`func (o *VoyageDetails) HasScheduledDepartureLt() bool`

HasScheduledDepartureLt returns a boolean if a field has been set.

### SetScheduledDepartureLtNil

`func (o *VoyageDetails) SetScheduledDepartureLtNil(b bool)`

 SetScheduledDepartureLtNil sets the value for ScheduledDepartureLt to be an explicit nil

### UnsetScheduledDepartureLt
`func (o *VoyageDetails) UnsetScheduledDepartureLt()`

UnsetScheduledDepartureLt ensures that no value is present for ScheduledDepartureLt, not even an explicit nil
### GetScheduledDepartureUtc

`func (o *VoyageDetails) GetScheduledDepartureUtc() time.Time`

GetScheduledDepartureUtc returns the ScheduledDepartureUtc field if non-nil, zero value otherwise.

### GetScheduledDepartureUtcOk

`func (o *VoyageDetails) GetScheduledDepartureUtcOk() (*time.Time, bool)`

GetScheduledDepartureUtcOk returns a tuple with the ScheduledDepartureUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDepartureUtc

`func (o *VoyageDetails) SetScheduledDepartureUtc(v time.Time)`

SetScheduledDepartureUtc sets ScheduledDepartureUtc field to given value.

### HasScheduledDepartureUtc

`func (o *VoyageDetails) HasScheduledDepartureUtc() bool`

HasScheduledDepartureUtc returns a boolean if a field has been set.

### SetScheduledDepartureUtcNil

`func (o *VoyageDetails) SetScheduledDepartureUtcNil(b bool)`

 SetScheduledDepartureUtcNil sets the value for ScheduledDepartureUtc to be an explicit nil

### UnsetScheduledDepartureUtc
`func (o *VoyageDetails) UnsetScheduledDepartureUtc()`

UnsetScheduledDepartureUtc ensures that no value is present for ScheduledDepartureUtc, not even an explicit nil
### GetTimezone

`func (o *VoyageDetails) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *VoyageDetails) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *VoyageDetails) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *VoyageDetails) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *VoyageDetails) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *VoyageDetails) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetVoyageNoList

`func (o *VoyageDetails) GetVoyageNoList() []string`

GetVoyageNoList returns the VoyageNoList field if non-nil, zero value otherwise.

### GetVoyageNoListOk

`func (o *VoyageDetails) GetVoyageNoListOk() (*[]string, bool)`

GetVoyageNoListOk returns a tuple with the VoyageNoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageNoList

`func (o *VoyageDetails) SetVoyageNoList(v []string)`

SetVoyageNoList sets VoyageNoList field to given value.

### HasVoyageNoList

`func (o *VoyageDetails) HasVoyageNoList() bool`

HasVoyageNoList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


