# BookingAPIVesselsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique object ID | [optional] 
**Name** | Pointer to **NullableString** | Vessel Name | [optional] 
**Imo** | Pointer to **NullableInt32** | Vessel IMO | [optional] 
**CallSign** | Pointer to **NullableString** | Vessel Call Sign | [optional] 
**Mmsi** | Pointer to **NullableInt32** | Vessel MMSI Number | [optional] 
**Flag** | Pointer to **NullableString** | Vessel Country Flag | [optional] 

## Methods

### NewBookingAPIVesselsInner

`func NewBookingAPIVesselsInner() *BookingAPIVesselsInner`

NewBookingAPIVesselsInner instantiates a new BookingAPIVesselsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingAPIVesselsInnerWithDefaults

`func NewBookingAPIVesselsInnerWithDefaults() *BookingAPIVesselsInner`

NewBookingAPIVesselsInnerWithDefaults instantiates a new BookingAPIVesselsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BookingAPIVesselsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BookingAPIVesselsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BookingAPIVesselsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BookingAPIVesselsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BookingAPIVesselsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BookingAPIVesselsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BookingAPIVesselsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BookingAPIVesselsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BookingAPIVesselsInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BookingAPIVesselsInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImo

`func (o *BookingAPIVesselsInner) GetImo() int32`

GetImo returns the Imo field if non-nil, zero value otherwise.

### GetImoOk

`func (o *BookingAPIVesselsInner) GetImoOk() (*int32, bool)`

GetImoOk returns a tuple with the Imo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImo

`func (o *BookingAPIVesselsInner) SetImo(v int32)`

SetImo sets Imo field to given value.

### HasImo

`func (o *BookingAPIVesselsInner) HasImo() bool`

HasImo returns a boolean if a field has been set.

### SetImoNil

`func (o *BookingAPIVesselsInner) SetImoNil(b bool)`

 SetImoNil sets the value for Imo to be an explicit nil

### UnsetImo
`func (o *BookingAPIVesselsInner) UnsetImo()`

UnsetImo ensures that no value is present for Imo, not even an explicit nil
### GetCallSign

`func (o *BookingAPIVesselsInner) GetCallSign() string`

GetCallSign returns the CallSign field if non-nil, zero value otherwise.

### GetCallSignOk

`func (o *BookingAPIVesselsInner) GetCallSignOk() (*string, bool)`

GetCallSignOk returns a tuple with the CallSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSign

`func (o *BookingAPIVesselsInner) SetCallSign(v string)`

SetCallSign sets CallSign field to given value.

### HasCallSign

`func (o *BookingAPIVesselsInner) HasCallSign() bool`

HasCallSign returns a boolean if a field has been set.

### SetCallSignNil

`func (o *BookingAPIVesselsInner) SetCallSignNil(b bool)`

 SetCallSignNil sets the value for CallSign to be an explicit nil

### UnsetCallSign
`func (o *BookingAPIVesselsInner) UnsetCallSign()`

UnsetCallSign ensures that no value is present for CallSign, not even an explicit nil
### GetMmsi

`func (o *BookingAPIVesselsInner) GetMmsi() int32`

GetMmsi returns the Mmsi field if non-nil, zero value otherwise.

### GetMmsiOk

`func (o *BookingAPIVesselsInner) GetMmsiOk() (*int32, bool)`

GetMmsiOk returns a tuple with the Mmsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmsi

`func (o *BookingAPIVesselsInner) SetMmsi(v int32)`

SetMmsi sets Mmsi field to given value.

### HasMmsi

`func (o *BookingAPIVesselsInner) HasMmsi() bool`

HasMmsi returns a boolean if a field has been set.

### SetMmsiNil

`func (o *BookingAPIVesselsInner) SetMmsiNil(b bool)`

 SetMmsiNil sets the value for Mmsi to be an explicit nil

### UnsetMmsi
`func (o *BookingAPIVesselsInner) UnsetMmsi()`

UnsetMmsi ensures that no value is present for Mmsi, not even an explicit nil
### GetFlag

`func (o *BookingAPIVesselsInner) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *BookingAPIVesselsInner) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *BookingAPIVesselsInner) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *BookingAPIVesselsInner) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### SetFlagNil

`func (o *BookingAPIVesselsInner) SetFlagNil(b bool)`

 SetFlagNil sets the value for Flag to be an explicit nil

### UnsetFlag
`func (o *BookingAPIVesselsInner) UnsetFlag()`

UnsetFlag ensures that no value is present for Flag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


