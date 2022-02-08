# BookingBookmarks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Response** | Pointer to [**[]BookingBookmarkResponse**](BookingBookmarkResponse.md) |  | [optional] 

## Methods

### NewBookingBookmarks

`func NewBookingBookmarks() *BookingBookmarks`

NewBookingBookmarks instantiates a new BookingBookmarks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingBookmarksWithDefaults

`func NewBookingBookmarksWithDefaults() *BookingBookmarks`

NewBookingBookmarksWithDefaults instantiates a new BookingBookmarks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *BookingBookmarks) GetContainers() []string`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *BookingBookmarks) GetContainersOk() (*[]string, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *BookingBookmarks) SetContainers(v []string)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *BookingBookmarks) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetStatus

`func (o *BookingBookmarks) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BookingBookmarks) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BookingBookmarks) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BookingBookmarks) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *BookingBookmarks) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BookingBookmarks) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BookingBookmarks) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BookingBookmarks) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResponse

`func (o *BookingBookmarks) GetResponse() []BookingBookmarkResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *BookingBookmarks) GetResponseOk() (*[]BookingBookmarkResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *BookingBookmarks) SetResponse(v []BookingBookmarkResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *BookingBookmarks) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


