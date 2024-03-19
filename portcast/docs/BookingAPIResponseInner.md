# BookingAPIResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Obj** | Pointer to [**BillOfLadingBookmark**](BillOfLadingBookmark.md) | Bill of Lading Bookmark Object | [optional] 
**Status** | Pointer to **interface{}** | Status of individual container upload | [optional] 
**Container** | Pointer to **interface{}** | Container Number registered under the BL/Booking | [optional] 
**Message** | Pointer to **interface{}** | Message describing error status | [optional] 
**PreviousIds** | Pointer to **interface{}** | List of Bookmark IDs registered for this container | [optional] 

## Methods

### NewBookingAPIResponseInner

`func NewBookingAPIResponseInner() *BookingAPIResponseInner`

NewBookingAPIResponseInner instantiates a new BookingAPIResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingAPIResponseInnerWithDefaults

`func NewBookingAPIResponseInnerWithDefaults() *BookingAPIResponseInner`

NewBookingAPIResponseInnerWithDefaults instantiates a new BookingAPIResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObj

`func (o *BookingAPIResponseInner) GetObj() BillOfLadingBookmark`

GetObj returns the Obj field if non-nil, zero value otherwise.

### GetObjOk

`func (o *BookingAPIResponseInner) GetObjOk() (*BillOfLadingBookmark, bool)`

GetObjOk returns a tuple with the Obj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObj

`func (o *BookingAPIResponseInner) SetObj(v BillOfLadingBookmark)`

SetObj sets Obj field to given value.

### HasObj

`func (o *BookingAPIResponseInner) HasObj() bool`

HasObj returns a boolean if a field has been set.

### GetStatus

`func (o *BookingAPIResponseInner) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BookingAPIResponseInner) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BookingAPIResponseInner) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BookingAPIResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BookingAPIResponseInner) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BookingAPIResponseInner) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetContainer

`func (o *BookingAPIResponseInner) GetContainer() interface{}`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *BookingAPIResponseInner) GetContainerOk() (*interface{}, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *BookingAPIResponseInner) SetContainer(v interface{})`

SetContainer sets Container field to given value.

### HasContainer

`func (o *BookingAPIResponseInner) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *BookingAPIResponseInner) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *BookingAPIResponseInner) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetMessage

`func (o *BookingAPIResponseInner) GetMessage() interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BookingAPIResponseInner) GetMessageOk() (*interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BookingAPIResponseInner) SetMessage(v interface{})`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BookingAPIResponseInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BookingAPIResponseInner) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BookingAPIResponseInner) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetPreviousIds

`func (o *BookingAPIResponseInner) GetPreviousIds() interface{}`

GetPreviousIds returns the PreviousIds field if non-nil, zero value otherwise.

### GetPreviousIdsOk

`func (o *BookingAPIResponseInner) GetPreviousIdsOk() (*interface{}, bool)`

GetPreviousIdsOk returns a tuple with the PreviousIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousIds

`func (o *BookingAPIResponseInner) SetPreviousIds(v interface{})`

SetPreviousIds sets PreviousIds field to given value.

### HasPreviousIds

`func (o *BookingAPIResponseInner) HasPreviousIds() bool`

HasPreviousIds returns a boolean if a field has been set.

### SetPreviousIdsNil

`func (o *BookingAPIResponseInner) SetPreviousIdsNil(b bool)`

 SetPreviousIdsNil sets the value for PreviousIds to be an explicit nil

### UnsetPreviousIds
`func (o *BookingAPIResponseInner) UnsetPreviousIds()`

UnsetPreviousIds ensures that no value is present for PreviousIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


