# BookingBookmarkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**Obj** | Pointer to [**Bookmark**](Bookmark.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**PreviousIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBookingBookmarkResponse

`func NewBookingBookmarkResponse() *BookingBookmarkResponse`

NewBookingBookmarkResponse instantiates a new BookingBookmarkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingBookmarkResponseWithDefaults

`func NewBookingBookmarkResponseWithDefaults() *BookingBookmarkResponse`

NewBookingBookmarkResponseWithDefaults instantiates a new BookingBookmarkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BookingBookmarkResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BookingBookmarkResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BookingBookmarkResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BookingBookmarkResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContainer

`func (o *BookingBookmarkResponse) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *BookingBookmarkResponse) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *BookingBookmarkResponse) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *BookingBookmarkResponse) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetObj

`func (o *BookingBookmarkResponse) GetObj() Bookmark`

GetObj returns the Obj field if non-nil, zero value otherwise.

### GetObjOk

`func (o *BookingBookmarkResponse) GetObjOk() (*Bookmark, bool)`

GetObjOk returns a tuple with the Obj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObj

`func (o *BookingBookmarkResponse) SetObj(v Bookmark)`

SetObj sets Obj field to given value.

### HasObj

`func (o *BookingBookmarkResponse) HasObj() bool`

HasObj returns a boolean if a field has been set.

### GetMessage

`func (o *BookingBookmarkResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BookingBookmarkResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BookingBookmarkResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BookingBookmarkResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPreviousIds

`func (o *BookingBookmarkResponse) GetPreviousIds() []string`

GetPreviousIds returns the PreviousIds field if non-nil, zero value otherwise.

### GetPreviousIdsOk

`func (o *BookingBookmarkResponse) GetPreviousIdsOk() (*[]string, bool)`

GetPreviousIdsOk returns a tuple with the PreviousIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousIds

`func (o *BookingBookmarkResponse) SetPreviousIds(v []string)`

SetPreviousIds sets PreviousIds field to given value.

### HasPreviousIds

`func (o *BookingBookmarkResponse) HasPreviousIds() bool`

HasPreviousIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


