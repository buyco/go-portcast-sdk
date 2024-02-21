# BookingAPIResponseInnerAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Obj** | Pointer to [**BillOfLadingBookmark**](BillOfLadingBookmark.md) | Bill of Lading Bookmark Object | [optional] 
**Status** | Pointer to **string** | Status of the individual container upload | [optional] 
**Container** | Pointer to **string** | Container Number registered under the BL/Booking | [optional] 

## Methods

### NewBookingAPIResponseInnerAnyOf

`func NewBookingAPIResponseInnerAnyOf() *BookingAPIResponseInnerAnyOf`

NewBookingAPIResponseInnerAnyOf instantiates a new BookingAPIResponseInnerAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingAPIResponseInnerAnyOfWithDefaults

`func NewBookingAPIResponseInnerAnyOfWithDefaults() *BookingAPIResponseInnerAnyOf`

NewBookingAPIResponseInnerAnyOfWithDefaults instantiates a new BookingAPIResponseInnerAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObj

`func (o *BookingAPIResponseInnerAnyOf) GetObj() BillOfLadingBookmark`

GetObj returns the Obj field if non-nil, zero value otherwise.

### GetObjOk

`func (o *BookingAPIResponseInnerAnyOf) GetObjOk() (*BillOfLadingBookmark, bool)`

GetObjOk returns a tuple with the Obj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObj

`func (o *BookingAPIResponseInnerAnyOf) SetObj(v BillOfLadingBookmark)`

SetObj sets Obj field to given value.

### HasObj

`func (o *BookingAPIResponseInnerAnyOf) HasObj() bool`

HasObj returns a boolean if a field has been set.

### GetStatus

`func (o *BookingAPIResponseInnerAnyOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BookingAPIResponseInnerAnyOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BookingAPIResponseInnerAnyOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BookingAPIResponseInnerAnyOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContainer

`func (o *BookingAPIResponseInnerAnyOf) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *BookingAPIResponseInnerAnyOf) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *BookingAPIResponseInnerAnyOf) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *BookingAPIResponseInnerAnyOf) HasContainer() bool`

HasContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


