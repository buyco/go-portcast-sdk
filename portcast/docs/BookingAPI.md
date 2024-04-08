# BookingAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | Pointer to **[]string** | List of Container Numbers | [optional] 
**Status** | Pointer to **string** | Status of the call made | [optional] 
**Message** | Pointer to **string** |   Key | Description  ---------|----------   SUCCESS | Successful call   BL_NOT_FOUND | Carrier didn&#39;t return any data for the queried document    INVALID_BL_NO | Carrier returned an error for the queried document   AUTO_CANT_DETECT_SEALINE | AUTO Scac detection wasn&#39;t able to get the correct carrier   UNEXPECTED_ERROR | There was an error while fetching the document details, please check input parameters | [optional] 
**Response** | Pointer to [**[]BookingAPIResponseInner**](BookingAPIResponseInner.md) | Bill of Lading Bookmarks Created List | [optional] 
**Route** | Pointer to [**BookingAPIRoute**](BookingAPIRoute.md) |  | [optional] 
**Locations** | Pointer to [**[]BookingAPILocationsInner**](BookingAPILocationsInner.md) | List of Locations the booking will visit | [optional] 
**Vessels** | Pointer to [**[]BookingAPIVesselsInner**](BookingAPIVesselsInner.md) | List of Vessels the booking will operate on | [optional] 

## Methods

### NewBookingAPI

`func NewBookingAPI() *BookingAPI`

NewBookingAPI instantiates a new BookingAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingAPIWithDefaults

`func NewBookingAPIWithDefaults() *BookingAPI`

NewBookingAPIWithDefaults instantiates a new BookingAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *BookingAPI) GetContainers() []string`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *BookingAPI) GetContainersOk() (*[]string, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *BookingAPI) SetContainers(v []string)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *BookingAPI) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetStatus

`func (o *BookingAPI) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BookingAPI) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BookingAPI) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BookingAPI) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *BookingAPI) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BookingAPI) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BookingAPI) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BookingAPI) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResponse

`func (o *BookingAPI) GetResponse() []BookingAPIResponseInner`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *BookingAPI) GetResponseOk() (*[]BookingAPIResponseInner, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *BookingAPI) SetResponse(v []BookingAPIResponseInner)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *BookingAPI) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetRoute

`func (o *BookingAPI) GetRoute() BookingAPIRoute`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *BookingAPI) GetRouteOk() (*BookingAPIRoute, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *BookingAPI) SetRoute(v BookingAPIRoute)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *BookingAPI) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### GetLocations

`func (o *BookingAPI) GetLocations() []BookingAPILocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *BookingAPI) GetLocationsOk() (*[]BookingAPILocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *BookingAPI) SetLocations(v []BookingAPILocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *BookingAPI) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetVessels

`func (o *BookingAPI) GetVessels() []BookingAPIVesselsInner`

GetVessels returns the Vessels field if non-nil, zero value otherwise.

### GetVesselsOk

`func (o *BookingAPI) GetVesselsOk() (*[]BookingAPIVesselsInner, bool)`

GetVesselsOk returns a tuple with the Vessels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVessels

`func (o *BookingAPI) SetVessels(v []BookingAPIVesselsInner)`

SetVessels sets Vessels field to given value.

### HasVessels

`func (o *BookingAPI) HasVessels() bool`

HasVessels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


