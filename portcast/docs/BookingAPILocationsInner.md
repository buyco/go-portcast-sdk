# BookingAPILocationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique object ID | [optional] 
**Name** | Pointer to **NullableString** | Location Name | [optional] 
**State** | Pointer to **NullableString** | Location State | [optional] 
**Country** | Pointer to **NullableString** | Location Country | [optional] 
**CountryCode** | Pointer to **NullableString** | Country Code as per UN LOCODEs | [optional] 
**Locode** | Pointer to **NullableString** | UN assigned LOCODE for the location | [optional] 
**Lat** | Pointer to **NullableFloat32** | Latitude | [optional] 
**Lng** | Pointer to **NullableFloat32** | Longitude | [optional] 

## Methods

### NewBookingAPILocationsInner

`func NewBookingAPILocationsInner() *BookingAPILocationsInner`

NewBookingAPILocationsInner instantiates a new BookingAPILocationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingAPILocationsInnerWithDefaults

`func NewBookingAPILocationsInnerWithDefaults() *BookingAPILocationsInner`

NewBookingAPILocationsInnerWithDefaults instantiates a new BookingAPILocationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BookingAPILocationsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BookingAPILocationsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BookingAPILocationsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BookingAPILocationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BookingAPILocationsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BookingAPILocationsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BookingAPILocationsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BookingAPILocationsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BookingAPILocationsInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BookingAPILocationsInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetState

`func (o *BookingAPILocationsInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BookingAPILocationsInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BookingAPILocationsInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BookingAPILocationsInner) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *BookingAPILocationsInner) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *BookingAPILocationsInner) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCountry

`func (o *BookingAPILocationsInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BookingAPILocationsInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BookingAPILocationsInner) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BookingAPILocationsInner) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *BookingAPILocationsInner) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *BookingAPILocationsInner) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCountryCode

`func (o *BookingAPILocationsInner) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *BookingAPILocationsInner) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *BookingAPILocationsInner) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *BookingAPILocationsInner) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *BookingAPILocationsInner) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *BookingAPILocationsInner) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetLocode

`func (o *BookingAPILocationsInner) GetLocode() string`

GetLocode returns the Locode field if non-nil, zero value otherwise.

### GetLocodeOk

`func (o *BookingAPILocationsInner) GetLocodeOk() (*string, bool)`

GetLocodeOk returns a tuple with the Locode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocode

`func (o *BookingAPILocationsInner) SetLocode(v string)`

SetLocode sets Locode field to given value.

### HasLocode

`func (o *BookingAPILocationsInner) HasLocode() bool`

HasLocode returns a boolean if a field has been set.

### SetLocodeNil

`func (o *BookingAPILocationsInner) SetLocodeNil(b bool)`

 SetLocodeNil sets the value for Locode to be an explicit nil

### UnsetLocode
`func (o *BookingAPILocationsInner) UnsetLocode()`

UnsetLocode ensures that no value is present for Locode, not even an explicit nil
### GetLat

`func (o *BookingAPILocationsInner) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *BookingAPILocationsInner) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *BookingAPILocationsInner) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *BookingAPILocationsInner) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *BookingAPILocationsInner) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *BookingAPILocationsInner) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLng

`func (o *BookingAPILocationsInner) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *BookingAPILocationsInner) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *BookingAPILocationsInner) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *BookingAPILocationsInner) HasLng() bool`

HasLng returns a boolean if a field has been set.

### SetLngNil

`func (o *BookingAPILocationsInner) SetLngNil(b bool)`

 SetLngNil sets the value for Lng to be an explicit nil

### UnsetLng
`func (o *BookingAPILocationsInner) UnsetLng()`

UnsetLng ensures that no value is present for Lng, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


