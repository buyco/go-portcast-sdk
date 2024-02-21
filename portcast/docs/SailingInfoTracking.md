# SailingInfoTracking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ais** | Pointer to [**SailingInfoTrackingAis**](SailingInfoTrackingAis.md) |  | [optional] 
**SailingInfo** | Pointer to [**SailingInfoTrackingSailingInfo**](SailingInfoTrackingSailingInfo.md) |  | [optional] 
**StatusInfo** | Pointer to [**SailingInfoTrackingStatusInfo**](SailingInfoTrackingStatusInfo.md) |  | [optional] 
**VoyageDetails** | Pointer to [**[]VoyageDetails**](VoyageDetails.md) | Sailing Schedule for the vessel between the starting and target port of this specific leg of the journey | [optional] 
**Msg** | Pointer to **string** | Details on AIS Data fetch | [optional] 
**Success** | Pointer to **bool** | Successfully Fetched AIS data or not | [optional] [default to true]

## Methods

### NewSailingInfoTracking

`func NewSailingInfoTracking() *SailingInfoTracking`

NewSailingInfoTracking instantiates a new SailingInfoTracking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSailingInfoTrackingWithDefaults

`func NewSailingInfoTrackingWithDefaults() *SailingInfoTracking`

NewSailingInfoTrackingWithDefaults instantiates a new SailingInfoTracking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAis

`func (o *SailingInfoTracking) GetAis() SailingInfoTrackingAis`

GetAis returns the Ais field if non-nil, zero value otherwise.

### GetAisOk

`func (o *SailingInfoTracking) GetAisOk() (*SailingInfoTrackingAis, bool)`

GetAisOk returns a tuple with the Ais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAis

`func (o *SailingInfoTracking) SetAis(v SailingInfoTrackingAis)`

SetAis sets Ais field to given value.

### HasAis

`func (o *SailingInfoTracking) HasAis() bool`

HasAis returns a boolean if a field has been set.

### GetSailingInfo

`func (o *SailingInfoTracking) GetSailingInfo() SailingInfoTrackingSailingInfo`

GetSailingInfo returns the SailingInfo field if non-nil, zero value otherwise.

### GetSailingInfoOk

`func (o *SailingInfoTracking) GetSailingInfoOk() (*SailingInfoTrackingSailingInfo, bool)`

GetSailingInfoOk returns a tuple with the SailingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSailingInfo

`func (o *SailingInfoTracking) SetSailingInfo(v SailingInfoTrackingSailingInfo)`

SetSailingInfo sets SailingInfo field to given value.

### HasSailingInfo

`func (o *SailingInfoTracking) HasSailingInfo() bool`

HasSailingInfo returns a boolean if a field has been set.

### GetStatusInfo

`func (o *SailingInfoTracking) GetStatusInfo() SailingInfoTrackingStatusInfo`

GetStatusInfo returns the StatusInfo field if non-nil, zero value otherwise.

### GetStatusInfoOk

`func (o *SailingInfoTracking) GetStatusInfoOk() (*SailingInfoTrackingStatusInfo, bool)`

GetStatusInfoOk returns a tuple with the StatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusInfo

`func (o *SailingInfoTracking) SetStatusInfo(v SailingInfoTrackingStatusInfo)`

SetStatusInfo sets StatusInfo field to given value.

### HasStatusInfo

`func (o *SailingInfoTracking) HasStatusInfo() bool`

HasStatusInfo returns a boolean if a field has been set.

### GetVoyageDetails

`func (o *SailingInfoTracking) GetVoyageDetails() []VoyageDetails`

GetVoyageDetails returns the VoyageDetails field if non-nil, zero value otherwise.

### GetVoyageDetailsOk

`func (o *SailingInfoTracking) GetVoyageDetailsOk() (*[]VoyageDetails, bool)`

GetVoyageDetailsOk returns a tuple with the VoyageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageDetails

`func (o *SailingInfoTracking) SetVoyageDetails(v []VoyageDetails)`

SetVoyageDetails sets VoyageDetails field to given value.

### HasVoyageDetails

`func (o *SailingInfoTracking) HasVoyageDetails() bool`

HasVoyageDetails returns a boolean if a field has been set.

### GetMsg

`func (o *SailingInfoTracking) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SailingInfoTracking) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SailingInfoTracking) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *SailingInfoTracking) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetSuccess

`func (o *SailingInfoTracking) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SailingInfoTracking) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SailingInfoTracking) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *SailingInfoTracking) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


