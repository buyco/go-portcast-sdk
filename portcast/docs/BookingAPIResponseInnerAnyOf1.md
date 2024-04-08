# BookingAPIResponseInnerAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message describing error status | [optional] 
**PreviousIds** | Pointer to **[]string** | List of Bookmark IDs registered for this container | [optional] 
**Status** | Pointer to **string** | Status of individual container upload | [optional] 
**Container** | Pointer to **string** | Container Number registered under the BL/Booking | [optional] 

## Methods

### NewBookingAPIResponseInnerAnyOf1

`func NewBookingAPIResponseInnerAnyOf1() *BookingAPIResponseInnerAnyOf1`

NewBookingAPIResponseInnerAnyOf1 instantiates a new BookingAPIResponseInnerAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingAPIResponseInnerAnyOf1WithDefaults

`func NewBookingAPIResponseInnerAnyOf1WithDefaults() *BookingAPIResponseInnerAnyOf1`

NewBookingAPIResponseInnerAnyOf1WithDefaults instantiates a new BookingAPIResponseInnerAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BookingAPIResponseInnerAnyOf1) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BookingAPIResponseInnerAnyOf1) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BookingAPIResponseInnerAnyOf1) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BookingAPIResponseInnerAnyOf1) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPreviousIds

`func (o *BookingAPIResponseInnerAnyOf1) GetPreviousIds() []string`

GetPreviousIds returns the PreviousIds field if non-nil, zero value otherwise.

### GetPreviousIdsOk

`func (o *BookingAPIResponseInnerAnyOf1) GetPreviousIdsOk() (*[]string, bool)`

GetPreviousIdsOk returns a tuple with the PreviousIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousIds

`func (o *BookingAPIResponseInnerAnyOf1) SetPreviousIds(v []string)`

SetPreviousIds sets PreviousIds field to given value.

### HasPreviousIds

`func (o *BookingAPIResponseInnerAnyOf1) HasPreviousIds() bool`

HasPreviousIds returns a boolean if a field has been set.

### GetStatus

`func (o *BookingAPIResponseInnerAnyOf1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BookingAPIResponseInnerAnyOf1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BookingAPIResponseInnerAnyOf1) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BookingAPIResponseInnerAnyOf1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContainer

`func (o *BookingAPIResponseInnerAnyOf1) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *BookingAPIResponseInnerAnyOf1) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *BookingAPIResponseInnerAnyOf1) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *BookingAPIResponseInnerAnyOf1) HasContainer() bool`

HasContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


