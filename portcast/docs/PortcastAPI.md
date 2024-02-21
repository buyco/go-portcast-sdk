# PortcastAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillOfLading** | Pointer to [**PortcastAPIBillOfLading**](PortcastAPIBillOfLading.md) |  | [optional] 
**BillOfLadingBookmark** | Pointer to [**BillOfLadingBookmark**](BillOfLadingBookmark.md) | Summary of the Bookmark ID Metadata | [optional] 
**ContainerEventList** | Pointer to [**[]ContainerEvent**](ContainerEvent.md) | Container Events List for the tracked shipment | [optional] 
**ContainerMetadata** | Pointer to [**ContainerMetadata**](ContainerMetadata.md) |  | [optional] 
**Co2Emissions** | Pointer to [**PortcastAPICo2Emissions**](PortcastAPICo2Emissions.md) |  | [optional] 
**DelayLists** | Pointer to [**[]DelayReasons**](DelayReasons.md) |  | [optional] 
**ResponseId** | Pointer to **string** | Unique ID for the API response generated | [optional] 
**Msg** | Pointer to **string** | Tracking Status Message | [optional] 
**OrgId** | Pointer to **string** | Customer Org (Flow) to which the bookmark ID is uploaded too | [optional] 
**SailingInfoTracking** | Pointer to [**[]SailingInfoTracking**](SailingInfoTracking.md) | Detailed Tracking Information for each leg of the journey | [optional] 
**StatusInfo** | Pointer to [**PortcastAPIStatusInfo**](PortcastAPIStatusInfo.md) |  | [optional] 
**Success** | Pointer to **bool** | Tracking Status | [optional] 

## Methods

### NewPortcastAPI

`func NewPortcastAPI() *PortcastAPI`

NewPortcastAPI instantiates a new PortcastAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortcastAPIWithDefaults

`func NewPortcastAPIWithDefaults() *PortcastAPI`

NewPortcastAPIWithDefaults instantiates a new PortcastAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillOfLading

`func (o *PortcastAPI) GetBillOfLading() PortcastAPIBillOfLading`

GetBillOfLading returns the BillOfLading field if non-nil, zero value otherwise.

### GetBillOfLadingOk

`func (o *PortcastAPI) GetBillOfLadingOk() (*PortcastAPIBillOfLading, bool)`

GetBillOfLadingOk returns a tuple with the BillOfLading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLading

`func (o *PortcastAPI) SetBillOfLading(v PortcastAPIBillOfLading)`

SetBillOfLading sets BillOfLading field to given value.

### HasBillOfLading

`func (o *PortcastAPI) HasBillOfLading() bool`

HasBillOfLading returns a boolean if a field has been set.

### GetBillOfLadingBookmark

`func (o *PortcastAPI) GetBillOfLadingBookmark() BillOfLadingBookmark`

GetBillOfLadingBookmark returns the BillOfLadingBookmark field if non-nil, zero value otherwise.

### GetBillOfLadingBookmarkOk

`func (o *PortcastAPI) GetBillOfLadingBookmarkOk() (*BillOfLadingBookmark, bool)`

GetBillOfLadingBookmarkOk returns a tuple with the BillOfLadingBookmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingBookmark

`func (o *PortcastAPI) SetBillOfLadingBookmark(v BillOfLadingBookmark)`

SetBillOfLadingBookmark sets BillOfLadingBookmark field to given value.

### HasBillOfLadingBookmark

`func (o *PortcastAPI) HasBillOfLadingBookmark() bool`

HasBillOfLadingBookmark returns a boolean if a field has been set.

### GetContainerEventList

`func (o *PortcastAPI) GetContainerEventList() []ContainerEvent`

GetContainerEventList returns the ContainerEventList field if non-nil, zero value otherwise.

### GetContainerEventListOk

`func (o *PortcastAPI) GetContainerEventListOk() (*[]ContainerEvent, bool)`

GetContainerEventListOk returns a tuple with the ContainerEventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerEventList

`func (o *PortcastAPI) SetContainerEventList(v []ContainerEvent)`

SetContainerEventList sets ContainerEventList field to given value.

### HasContainerEventList

`func (o *PortcastAPI) HasContainerEventList() bool`

HasContainerEventList returns a boolean if a field has been set.

### GetContainerMetadata

`func (o *PortcastAPI) GetContainerMetadata() ContainerMetadata`

GetContainerMetadata returns the ContainerMetadata field if non-nil, zero value otherwise.

### GetContainerMetadataOk

`func (o *PortcastAPI) GetContainerMetadataOk() (*ContainerMetadata, bool)`

GetContainerMetadataOk returns a tuple with the ContainerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerMetadata

`func (o *PortcastAPI) SetContainerMetadata(v ContainerMetadata)`

SetContainerMetadata sets ContainerMetadata field to given value.

### HasContainerMetadata

`func (o *PortcastAPI) HasContainerMetadata() bool`

HasContainerMetadata returns a boolean if a field has been set.

### GetCo2Emissions

`func (o *PortcastAPI) GetCo2Emissions() PortcastAPICo2Emissions`

GetCo2Emissions returns the Co2Emissions field if non-nil, zero value otherwise.

### GetCo2EmissionsOk

`func (o *PortcastAPI) GetCo2EmissionsOk() (*PortcastAPICo2Emissions, bool)`

GetCo2EmissionsOk returns a tuple with the Co2Emissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCo2Emissions

`func (o *PortcastAPI) SetCo2Emissions(v PortcastAPICo2Emissions)`

SetCo2Emissions sets Co2Emissions field to given value.

### HasCo2Emissions

`func (o *PortcastAPI) HasCo2Emissions() bool`

HasCo2Emissions returns a boolean if a field has been set.

### GetDelayLists

`func (o *PortcastAPI) GetDelayLists() []DelayReasons`

GetDelayLists returns the DelayLists field if non-nil, zero value otherwise.

### GetDelayListsOk

`func (o *PortcastAPI) GetDelayListsOk() (*[]DelayReasons, bool)`

GetDelayListsOk returns a tuple with the DelayLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayLists

`func (o *PortcastAPI) SetDelayLists(v []DelayReasons)`

SetDelayLists sets DelayLists field to given value.

### HasDelayLists

`func (o *PortcastAPI) HasDelayLists() bool`

HasDelayLists returns a boolean if a field has been set.

### GetResponseId

`func (o *PortcastAPI) GetResponseId() string`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *PortcastAPI) GetResponseIdOk() (*string, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *PortcastAPI) SetResponseId(v string)`

SetResponseId sets ResponseId field to given value.

### HasResponseId

`func (o *PortcastAPI) HasResponseId() bool`

HasResponseId returns a boolean if a field has been set.

### GetMsg

`func (o *PortcastAPI) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *PortcastAPI) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *PortcastAPI) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *PortcastAPI) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetOrgId

`func (o *PortcastAPI) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *PortcastAPI) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *PortcastAPI) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *PortcastAPI) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSailingInfoTracking

`func (o *PortcastAPI) GetSailingInfoTracking() []SailingInfoTracking`

GetSailingInfoTracking returns the SailingInfoTracking field if non-nil, zero value otherwise.

### GetSailingInfoTrackingOk

`func (o *PortcastAPI) GetSailingInfoTrackingOk() (*[]SailingInfoTracking, bool)`

GetSailingInfoTrackingOk returns a tuple with the SailingInfoTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSailingInfoTracking

`func (o *PortcastAPI) SetSailingInfoTracking(v []SailingInfoTracking)`

SetSailingInfoTracking sets SailingInfoTracking field to given value.

### HasSailingInfoTracking

`func (o *PortcastAPI) HasSailingInfoTracking() bool`

HasSailingInfoTracking returns a boolean if a field has been set.

### GetStatusInfo

`func (o *PortcastAPI) GetStatusInfo() PortcastAPIStatusInfo`

GetStatusInfo returns the StatusInfo field if non-nil, zero value otherwise.

### GetStatusInfoOk

`func (o *PortcastAPI) GetStatusInfoOk() (*PortcastAPIStatusInfo, bool)`

GetStatusInfoOk returns a tuple with the StatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusInfo

`func (o *PortcastAPI) SetStatusInfo(v PortcastAPIStatusInfo)`

SetStatusInfo sets StatusInfo field to given value.

### HasStatusInfo

`func (o *PortcastAPI) HasStatusInfo() bool`

HasStatusInfo returns a boolean if a field has been set.

### GetSuccess

`func (o *PortcastAPI) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PortcastAPI) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PortcastAPI) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *PortcastAPI) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


