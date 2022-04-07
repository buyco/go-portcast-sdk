# BookingBookmarksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Response** | Pointer to [**[]BookingBookmark**](BookingBookmark.md) |  | [optional] 

## Methods

### NewBookingBookmarksResponse

`func NewBookingBookmarksResponse() *BookingBookmarksResponse`

NewBookingBookmarksResponse instantiates a new BookingBookmarksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingBookmarksResponseWithDefaults

`func NewBookingBookmarksResponseWithDefaults() *BookingBookmarksResponse`

NewBookingBookmarksResponseWithDefaults instantiates a new BookingBookmarksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *BookingBookmarksResponse) GetContainers() []string`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *BookingBookmarksResponse) GetContainersOk() (*[]string, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *BookingBookmarksResponse) SetContainers(v []string)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *BookingBookmarksResponse) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetStatus

`func (o *BookingBookmarksResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BookingBookmarksResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BookingBookmarksResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BookingBookmarksResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *BookingBookmarksResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BookingBookmarksResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BookingBookmarksResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BookingBookmarksResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResponse

`func (o *BookingBookmarksResponse) GetResponse() []BookingBookmark`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *BookingBookmarksResponse) GetResponseOk() (*[]BookingBookmark, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *BookingBookmarksResponse) SetResponse(v []BookingBookmark)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *BookingBookmarksResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


