# BookingAPIResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Obj** | Pointer to [**BillOfLadingBookmark**](BillOfLadingBookmark.md) | Bill of Lading Bookmark Object | [optional] 
**Status** | Pointer to **string** | Status of individual container upload | [optional] 
**Container** | Pointer to **string** | Container Number registered under the BL/Booking | [optional] 
**Message** | Pointer to **string** | Message describing error status | [optional] 
**PreviousIds** | Pointer to **[]string** | List of Bookmark IDs registered for this container | [optional] 

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

`func (o *BookingAPIResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BookingAPIResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BookingAPIResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BookingAPIResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContainer

`func (o *BookingAPIResponseInner) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *BookingAPIResponseInner) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *BookingAPIResponseInner) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *BookingAPIResponseInner) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetMessage

`func (o *BookingAPIResponseInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BookingAPIResponseInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BookingAPIResponseInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BookingAPIResponseInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPreviousIds

`func (o *BookingAPIResponseInner) GetPreviousIds() []string`

GetPreviousIds returns the PreviousIds field if non-nil, zero value otherwise.

### GetPreviousIdsOk

`func (o *BookingAPIResponseInner) GetPreviousIdsOk() (*[]string, bool)`

GetPreviousIdsOk returns a tuple with the PreviousIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousIds

`func (o *BookingAPIResponseInner) SetPreviousIds(v []string)`

SetPreviousIds sets PreviousIds field to given value.

### HasPreviousIds

`func (o *BookingAPIResponseInner) HasPreviousIds() bool`

HasPreviousIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


