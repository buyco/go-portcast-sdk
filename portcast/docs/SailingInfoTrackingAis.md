# SailingInfoTrackingAis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Course** | Pointer to **NullableFloat32** | Vessel Course | [optional] 
**Imo** | Pointer to **NullableString** | Vessel IMO | [optional] 
**Lat** | Pointer to **NullableFloat32** | Vessel&#39;s Latitude Position | [optional] 
**Lon** | Pointer to **NullableFloat32** | Vessel&#39;s Longitude Position | [optional] 
**SpeedNm** | Pointer to **NullableFloat32** | Vessel&#39;s Speed in knots | [optional] 
**Status** | Pointer to **NullableInt32** | Vessel&#39;s AIS Navigational Status | [optional] 
**TimestampUtc** | Pointer to **NullableTime** | Timestamp for when the AIS Data was fetched | [optional] 

## Methods

### NewSailingInfoTrackingAis

`func NewSailingInfoTrackingAis() *SailingInfoTrackingAis`

NewSailingInfoTrackingAis instantiates a new SailingInfoTrackingAis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSailingInfoTrackingAisWithDefaults

`func NewSailingInfoTrackingAisWithDefaults() *SailingInfoTrackingAis`

NewSailingInfoTrackingAisWithDefaults instantiates a new SailingInfoTrackingAis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCourse

`func (o *SailingInfoTrackingAis) GetCourse() float32`

GetCourse returns the Course field if non-nil, zero value otherwise.

### GetCourseOk

`func (o *SailingInfoTrackingAis) GetCourseOk() (*float32, bool)`

GetCourseOk returns a tuple with the Course field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourse

`func (o *SailingInfoTrackingAis) SetCourse(v float32)`

SetCourse sets Course field to given value.

### HasCourse

`func (o *SailingInfoTrackingAis) HasCourse() bool`

HasCourse returns a boolean if a field has been set.

### SetCourseNil

`func (o *SailingInfoTrackingAis) SetCourseNil(b bool)`

 SetCourseNil sets the value for Course to be an explicit nil

### UnsetCourse
`func (o *SailingInfoTrackingAis) UnsetCourse()`

UnsetCourse ensures that no value is present for Course, not even an explicit nil
### GetImo

`func (o *SailingInfoTrackingAis) GetImo() string`

GetImo returns the Imo field if non-nil, zero value otherwise.

### GetImoOk

`func (o *SailingInfoTrackingAis) GetImoOk() (*string, bool)`

GetImoOk returns a tuple with the Imo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImo

`func (o *SailingInfoTrackingAis) SetImo(v string)`

SetImo sets Imo field to given value.

### HasImo

`func (o *SailingInfoTrackingAis) HasImo() bool`

HasImo returns a boolean if a field has been set.

### SetImoNil

`func (o *SailingInfoTrackingAis) SetImoNil(b bool)`

 SetImoNil sets the value for Imo to be an explicit nil

### UnsetImo
`func (o *SailingInfoTrackingAis) UnsetImo()`

UnsetImo ensures that no value is present for Imo, not even an explicit nil
### GetLat

`func (o *SailingInfoTrackingAis) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *SailingInfoTrackingAis) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *SailingInfoTrackingAis) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *SailingInfoTrackingAis) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *SailingInfoTrackingAis) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *SailingInfoTrackingAis) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLon

`func (o *SailingInfoTrackingAis) GetLon() float32`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *SailingInfoTrackingAis) GetLonOk() (*float32, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *SailingInfoTrackingAis) SetLon(v float32)`

SetLon sets Lon field to given value.

### HasLon

`func (o *SailingInfoTrackingAis) HasLon() bool`

HasLon returns a boolean if a field has been set.

### SetLonNil

`func (o *SailingInfoTrackingAis) SetLonNil(b bool)`

 SetLonNil sets the value for Lon to be an explicit nil

### UnsetLon
`func (o *SailingInfoTrackingAis) UnsetLon()`

UnsetLon ensures that no value is present for Lon, not even an explicit nil
### GetSpeedNm

`func (o *SailingInfoTrackingAis) GetSpeedNm() float32`

GetSpeedNm returns the SpeedNm field if non-nil, zero value otherwise.

### GetSpeedNmOk

`func (o *SailingInfoTrackingAis) GetSpeedNmOk() (*float32, bool)`

GetSpeedNmOk returns a tuple with the SpeedNm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedNm

`func (o *SailingInfoTrackingAis) SetSpeedNm(v float32)`

SetSpeedNm sets SpeedNm field to given value.

### HasSpeedNm

`func (o *SailingInfoTrackingAis) HasSpeedNm() bool`

HasSpeedNm returns a boolean if a field has been set.

### SetSpeedNmNil

`func (o *SailingInfoTrackingAis) SetSpeedNmNil(b bool)`

 SetSpeedNmNil sets the value for SpeedNm to be an explicit nil

### UnsetSpeedNm
`func (o *SailingInfoTrackingAis) UnsetSpeedNm()`

UnsetSpeedNm ensures that no value is present for SpeedNm, not even an explicit nil
### GetStatus

`func (o *SailingInfoTrackingAis) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SailingInfoTrackingAis) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SailingInfoTrackingAis) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SailingInfoTrackingAis) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SailingInfoTrackingAis) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SailingInfoTrackingAis) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTimestampUtc

`func (o *SailingInfoTrackingAis) GetTimestampUtc() time.Time`

GetTimestampUtc returns the TimestampUtc field if non-nil, zero value otherwise.

### GetTimestampUtcOk

`func (o *SailingInfoTrackingAis) GetTimestampUtcOk() (*time.Time, bool)`

GetTimestampUtcOk returns a tuple with the TimestampUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampUtc

`func (o *SailingInfoTrackingAis) SetTimestampUtc(v time.Time)`

SetTimestampUtc sets TimestampUtc field to given value.

### HasTimestampUtc

`func (o *SailingInfoTrackingAis) HasTimestampUtc() bool`

HasTimestampUtc returns a boolean if a field has been set.

### SetTimestampUtcNil

`func (o *SailingInfoTrackingAis) SetTimestampUtcNil(b bool)`

 SetTimestampUtcNil sets the value for TimestampUtc to be an explicit nil

### UnsetTimestampUtc
`func (o *SailingInfoTrackingAis) UnsetTimestampUtc()`

UnsetTimestampUtc ensures that no value is present for TimestampUtc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


