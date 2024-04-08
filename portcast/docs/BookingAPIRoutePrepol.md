# BookingAPIRoutePrepol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **NullableInt32** | Location ID - Refer to location obj for details | [optional] 
**Date** | Pointer to **NullableString** | Timestamp on when the first event takes place at the pre-POL | [optional] 
**Actual** | Pointer to **NullableBool** | actual or estimated boolean | [optional] 

## Methods

### NewBookingAPIRoutePrepol

`func NewBookingAPIRoutePrepol() *BookingAPIRoutePrepol`

NewBookingAPIRoutePrepol instantiates a new BookingAPIRoutePrepol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingAPIRoutePrepolWithDefaults

`func NewBookingAPIRoutePrepolWithDefaults() *BookingAPIRoutePrepol`

NewBookingAPIRoutePrepolWithDefaults instantiates a new BookingAPIRoutePrepol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *BookingAPIRoutePrepol) GetLocation() int32`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BookingAPIRoutePrepol) GetLocationOk() (*int32, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BookingAPIRoutePrepol) SetLocation(v int32)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BookingAPIRoutePrepol) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *BookingAPIRoutePrepol) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *BookingAPIRoutePrepol) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetDate

`func (o *BookingAPIRoutePrepol) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BookingAPIRoutePrepol) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BookingAPIRoutePrepol) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *BookingAPIRoutePrepol) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *BookingAPIRoutePrepol) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *BookingAPIRoutePrepol) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetActual

`func (o *BookingAPIRoutePrepol) GetActual() bool`

GetActual returns the Actual field if non-nil, zero value otherwise.

### GetActualOk

`func (o *BookingAPIRoutePrepol) GetActualOk() (*bool, bool)`

GetActualOk returns a tuple with the Actual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActual

`func (o *BookingAPIRoutePrepol) SetActual(v bool)`

SetActual sets Actual field to given value.

### HasActual

`func (o *BookingAPIRoutePrepol) HasActual() bool`

HasActual returns a boolean if a field has been set.

### SetActualNil

`func (o *BookingAPIRoutePrepol) SetActualNil(b bool)`

 SetActualNil sets the value for Actual to be an explicit nil

### UnsetActual
`func (o *BookingAPIRoutePrepol) UnsetActual()`

UnsetActual ensures that no value is present for Actual, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


