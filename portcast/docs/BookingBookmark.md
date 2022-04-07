# BookingBookmark

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**Obj** | Pointer to [**Bookmark**](Bookmark.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**PreviousIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBookingBookmark

`func NewBookingBookmark() *BookingBookmark`

NewBookingBookmark instantiates a new BookingBookmark object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingBookmarkWithDefaults

`func NewBookingBookmarkWithDefaults() *BookingBookmark`

NewBookingBookmarkWithDefaults instantiates a new BookingBookmark object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BookingBookmark) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BookingBookmark) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BookingBookmark) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BookingBookmark) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContainer

`func (o *BookingBookmark) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *BookingBookmark) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *BookingBookmark) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *BookingBookmark) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetObj

`func (o *BookingBookmark) GetObj() Bookmark`

GetObj returns the Obj field if non-nil, zero value otherwise.

### GetObjOk

`func (o *BookingBookmark) GetObjOk() (*Bookmark, bool)`

GetObjOk returns a tuple with the Obj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObj

`func (o *BookingBookmark) SetObj(v Bookmark)`

SetObj sets Obj field to given value.

### HasObj

`func (o *BookingBookmark) HasObj() bool`

HasObj returns a boolean if a field has been set.

### GetMessage

`func (o *BookingBookmark) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BookingBookmark) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BookingBookmark) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BookingBookmark) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPreviousIds

`func (o *BookingBookmark) GetPreviousIds() []string`

GetPreviousIds returns the PreviousIds field if non-nil, zero value otherwise.

### GetPreviousIdsOk

`func (o *BookingBookmark) GetPreviousIdsOk() (*[]string, bool)`

GetPreviousIdsOk returns a tuple with the PreviousIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousIds

`func (o *BookingBookmark) SetPreviousIds(v []string)`

SetPreviousIds sets PreviousIds field to given value.

### HasPreviousIds

`func (o *BookingBookmark) HasPreviousIds() bool`

HasPreviousIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


