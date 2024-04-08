# BillOfLadingBookmark

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlNo** | Pointer to **NullableString** | Bill of Lading No or Booking No. | [optional] 
**CarrierNo** | Pointer to **string** | Carrier SCAC Code | [optional] 
**CntrNo** | Pointer to **string** | Container Number | [optional] 
**Created** | Pointer to **time.Time** | Bookmark Upload Date | [optional] 
**Deleted** | Pointer to **bool** | Archival Status of the bookmark | [optional] [default to false]
**Id** | Pointer to **string** | Unique Identifier called Bill of Lading Bookmark ID [Primary Key] (Also refered as Bookmark ID) | [optional] 
**OrgId** | Pointer to **string** | Customer Org (Flow) to which the bookmark ID is uploaded too | [optional] 
**Status** | Pointer to **string** | [Deprecated] Old Status codes - Not supported | [optional] 
**StatusCode** | Pointer to **string** | [Depricated] Old Status codes - Not supported | [optional] 
**StatusInfo** | Pointer to [**BillOfLadingBookmarkStatusInfo**](BillOfLadingBookmarkStatusInfo.md) |  | [optional] 
**SystemDeleted** | Pointer to **bool** | System Deletion Status of the bookmark | [optional] [default to false]
**Updated** | Pointer to **time.Time** | Last Updated date for the Bookmark ID (Refers to when we fetched the latest data for the bookmark ID) | [optional] 

## Methods

### NewBillOfLadingBookmark

`func NewBillOfLadingBookmark() *BillOfLadingBookmark`

NewBillOfLadingBookmark instantiates a new BillOfLadingBookmark object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillOfLadingBookmarkWithDefaults

`func NewBillOfLadingBookmarkWithDefaults() *BillOfLadingBookmark`

NewBillOfLadingBookmarkWithDefaults instantiates a new BillOfLadingBookmark object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlNo

`func (o *BillOfLadingBookmark) GetBlNo() string`

GetBlNo returns the BlNo field if non-nil, zero value otherwise.

### GetBlNoOk

`func (o *BillOfLadingBookmark) GetBlNoOk() (*string, bool)`

GetBlNoOk returns a tuple with the BlNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlNo

`func (o *BillOfLadingBookmark) SetBlNo(v string)`

SetBlNo sets BlNo field to given value.

### HasBlNo

`func (o *BillOfLadingBookmark) HasBlNo() bool`

HasBlNo returns a boolean if a field has been set.

### SetBlNoNil

`func (o *BillOfLadingBookmark) SetBlNoNil(b bool)`

 SetBlNoNil sets the value for BlNo to be an explicit nil

### UnsetBlNo
`func (o *BillOfLadingBookmark) UnsetBlNo()`

UnsetBlNo ensures that no value is present for BlNo, not even an explicit nil
### GetCarrierNo

`func (o *BillOfLadingBookmark) GetCarrierNo() string`

GetCarrierNo returns the CarrierNo field if non-nil, zero value otherwise.

### GetCarrierNoOk

`func (o *BillOfLadingBookmark) GetCarrierNoOk() (*string, bool)`

GetCarrierNoOk returns a tuple with the CarrierNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierNo

`func (o *BillOfLadingBookmark) SetCarrierNo(v string)`

SetCarrierNo sets CarrierNo field to given value.

### HasCarrierNo

`func (o *BillOfLadingBookmark) HasCarrierNo() bool`

HasCarrierNo returns a boolean if a field has been set.

### GetCntrNo

`func (o *BillOfLadingBookmark) GetCntrNo() string`

GetCntrNo returns the CntrNo field if non-nil, zero value otherwise.

### GetCntrNoOk

`func (o *BillOfLadingBookmark) GetCntrNoOk() (*string, bool)`

GetCntrNoOk returns a tuple with the CntrNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCntrNo

`func (o *BillOfLadingBookmark) SetCntrNo(v string)`

SetCntrNo sets CntrNo field to given value.

### HasCntrNo

`func (o *BillOfLadingBookmark) HasCntrNo() bool`

HasCntrNo returns a boolean if a field has been set.

### GetCreated

`func (o *BillOfLadingBookmark) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BillOfLadingBookmark) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BillOfLadingBookmark) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BillOfLadingBookmark) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDeleted

`func (o *BillOfLadingBookmark) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *BillOfLadingBookmark) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *BillOfLadingBookmark) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *BillOfLadingBookmark) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetId

`func (o *BillOfLadingBookmark) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillOfLadingBookmark) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillOfLadingBookmark) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillOfLadingBookmark) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *BillOfLadingBookmark) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *BillOfLadingBookmark) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *BillOfLadingBookmark) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *BillOfLadingBookmark) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetStatus

`func (o *BillOfLadingBookmark) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillOfLadingBookmark) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillOfLadingBookmark) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillOfLadingBookmark) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *BillOfLadingBookmark) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *BillOfLadingBookmark) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *BillOfLadingBookmark) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *BillOfLadingBookmark) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetStatusInfo

`func (o *BillOfLadingBookmark) GetStatusInfo() BillOfLadingBookmarkStatusInfo`

GetStatusInfo returns the StatusInfo field if non-nil, zero value otherwise.

### GetStatusInfoOk

`func (o *BillOfLadingBookmark) GetStatusInfoOk() (*BillOfLadingBookmarkStatusInfo, bool)`

GetStatusInfoOk returns a tuple with the StatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusInfo

`func (o *BillOfLadingBookmark) SetStatusInfo(v BillOfLadingBookmarkStatusInfo)`

SetStatusInfo sets StatusInfo field to given value.

### HasStatusInfo

`func (o *BillOfLadingBookmark) HasStatusInfo() bool`

HasStatusInfo returns a boolean if a field has been set.

### GetSystemDeleted

`func (o *BillOfLadingBookmark) GetSystemDeleted() bool`

GetSystemDeleted returns the SystemDeleted field if non-nil, zero value otherwise.

### GetSystemDeletedOk

`func (o *BillOfLadingBookmark) GetSystemDeletedOk() (*bool, bool)`

GetSystemDeletedOk returns a tuple with the SystemDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDeleted

`func (o *BillOfLadingBookmark) SetSystemDeleted(v bool)`

SetSystemDeleted sets SystemDeleted field to given value.

### HasSystemDeleted

`func (o *BillOfLadingBookmark) HasSystemDeleted() bool`

HasSystemDeleted returns a boolean if a field has been set.

### GetUpdated

`func (o *BillOfLadingBookmark) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *BillOfLadingBookmark) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *BillOfLadingBookmark) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *BillOfLadingBookmark) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


