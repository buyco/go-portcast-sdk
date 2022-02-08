# TrackingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillOfLading** | Pointer to [**TrackingEventBillOfLading**](TrackingEventBillOfLading.md) |  | [optional] 
**BillOfLadingBookmark** | Pointer to [**TrackingEventBillOfLadingBookmark**](TrackingEventBillOfLadingBookmark.md) |  | [optional] 
**ContainerEventList** | Pointer to [**[]TrackingEventContainerEventList**](TrackingEventContainerEventList.md) |  | [optional] 
**ContainerMetaInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**SailingInfoTracking** | Pointer to [**[]TrackingEventSailingInfoTracking**](TrackingEventSailingInfoTracking.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewTrackingEvent

`func NewTrackingEvent() *TrackingEvent`

NewTrackingEvent instantiates a new TrackingEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingEventWithDefaults

`func NewTrackingEventWithDefaults() *TrackingEvent`

NewTrackingEventWithDefaults instantiates a new TrackingEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillOfLading

`func (o *TrackingEvent) GetBillOfLading() TrackingEventBillOfLading`

GetBillOfLading returns the BillOfLading field if non-nil, zero value otherwise.

### GetBillOfLadingOk

`func (o *TrackingEvent) GetBillOfLadingOk() (*TrackingEventBillOfLading, bool)`

GetBillOfLadingOk returns a tuple with the BillOfLading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLading

`func (o *TrackingEvent) SetBillOfLading(v TrackingEventBillOfLading)`

SetBillOfLading sets BillOfLading field to given value.

### HasBillOfLading

`func (o *TrackingEvent) HasBillOfLading() bool`

HasBillOfLading returns a boolean if a field has been set.

### GetBillOfLadingBookmark

`func (o *TrackingEvent) GetBillOfLadingBookmark() TrackingEventBillOfLadingBookmark`

GetBillOfLadingBookmark returns the BillOfLadingBookmark field if non-nil, zero value otherwise.

### GetBillOfLadingBookmarkOk

`func (o *TrackingEvent) GetBillOfLadingBookmarkOk() (*TrackingEventBillOfLadingBookmark, bool)`

GetBillOfLadingBookmarkOk returns a tuple with the BillOfLadingBookmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingBookmark

`func (o *TrackingEvent) SetBillOfLadingBookmark(v TrackingEventBillOfLadingBookmark)`

SetBillOfLadingBookmark sets BillOfLadingBookmark field to given value.

### HasBillOfLadingBookmark

`func (o *TrackingEvent) HasBillOfLadingBookmark() bool`

HasBillOfLadingBookmark returns a boolean if a field has been set.

### GetContainerEventList

`func (o *TrackingEvent) GetContainerEventList() []TrackingEventContainerEventList`

GetContainerEventList returns the ContainerEventList field if non-nil, zero value otherwise.

### GetContainerEventListOk

`func (o *TrackingEvent) GetContainerEventListOk() (*[]TrackingEventContainerEventList, bool)`

GetContainerEventListOk returns a tuple with the ContainerEventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerEventList

`func (o *TrackingEvent) SetContainerEventList(v []TrackingEventContainerEventList)`

SetContainerEventList sets ContainerEventList field to given value.

### HasContainerEventList

`func (o *TrackingEvent) HasContainerEventList() bool`

HasContainerEventList returns a boolean if a field has been set.

### GetContainerMetaInfo

`func (o *TrackingEvent) GetContainerMetaInfo() map[string]interface{}`

GetContainerMetaInfo returns the ContainerMetaInfo field if non-nil, zero value otherwise.

### GetContainerMetaInfoOk

`func (o *TrackingEvent) GetContainerMetaInfoOk() (*map[string]interface{}, bool)`

GetContainerMetaInfoOk returns a tuple with the ContainerMetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerMetaInfo

`func (o *TrackingEvent) SetContainerMetaInfo(v map[string]interface{})`

SetContainerMetaInfo sets ContainerMetaInfo field to given value.

### HasContainerMetaInfo

`func (o *TrackingEvent) HasContainerMetaInfo() bool`

HasContainerMetaInfo returns a boolean if a field has been set.

### GetId

`func (o *TrackingEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackingEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackingEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrackingEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMsg

`func (o *TrackingEvent) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *TrackingEvent) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *TrackingEvent) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *TrackingEvent) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetOrgId

`func (o *TrackingEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *TrackingEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *TrackingEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *TrackingEvent) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSailingInfoTracking

`func (o *TrackingEvent) GetSailingInfoTracking() []TrackingEventSailingInfoTracking`

GetSailingInfoTracking returns the SailingInfoTracking field if non-nil, zero value otherwise.

### GetSailingInfoTrackingOk

`func (o *TrackingEvent) GetSailingInfoTrackingOk() (*[]TrackingEventSailingInfoTracking, bool)`

GetSailingInfoTrackingOk returns a tuple with the SailingInfoTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSailingInfoTracking

`func (o *TrackingEvent) SetSailingInfoTracking(v []TrackingEventSailingInfoTracking)`

SetSailingInfoTracking sets SailingInfoTracking field to given value.

### HasSailingInfoTracking

`func (o *TrackingEvent) HasSailingInfoTracking() bool`

HasSailingInfoTracking returns a boolean if a field has been set.

### GetSuccess

`func (o *TrackingEvent) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TrackingEvent) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TrackingEvent) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *TrackingEvent) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


