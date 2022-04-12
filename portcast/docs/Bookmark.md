# Bookmark

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlNo** | Pointer to **string** |  | [optional] 
**CarrierNo** | Pointer to **string** |  | [optional] 
**CntrNo** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**CustomerNo** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **string** |  | [optional] 
**SystemDeleted** | Pointer to **bool** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**VesselList** | Pointer to **[]string** |  | [optional] 
**VoyageNo** | Pointer to **string** |  | [optional] 

## Methods

### NewBookmark

`func NewBookmark() *Bookmark`

NewBookmark instantiates a new Bookmark object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarkWithDefaults

`func NewBookmarkWithDefaults() *Bookmark`

NewBookmarkWithDefaults instantiates a new Bookmark object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlNo

`func (o *Bookmark) GetBlNo() string`

GetBlNo returns the BlNo field if non-nil, zero value otherwise.

### GetBlNoOk

`func (o *Bookmark) GetBlNoOk() (*string, bool)`

GetBlNoOk returns a tuple with the BlNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlNo

`func (o *Bookmark) SetBlNo(v string)`

SetBlNo sets BlNo field to given value.

### HasBlNo

`func (o *Bookmark) HasBlNo() bool`

HasBlNo returns a boolean if a field has been set.

### GetCarrierNo

`func (o *Bookmark) GetCarrierNo() string`

GetCarrierNo returns the CarrierNo field if non-nil, zero value otherwise.

### GetCarrierNoOk

`func (o *Bookmark) GetCarrierNoOk() (*string, bool)`

GetCarrierNoOk returns a tuple with the CarrierNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierNo

`func (o *Bookmark) SetCarrierNo(v string)`

SetCarrierNo sets CarrierNo field to given value.

### HasCarrierNo

`func (o *Bookmark) HasCarrierNo() bool`

HasCarrierNo returns a boolean if a field has been set.

### GetCntrNo

`func (o *Bookmark) GetCntrNo() string`

GetCntrNo returns the CntrNo field if non-nil, zero value otherwise.

### GetCntrNoOk

`func (o *Bookmark) GetCntrNoOk() (*string, bool)`

GetCntrNoOk returns a tuple with the CntrNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCntrNo

`func (o *Bookmark) SetCntrNo(v string)`

SetCntrNo sets CntrNo field to given value.

### HasCntrNo

`func (o *Bookmark) HasCntrNo() bool`

HasCntrNo returns a boolean if a field has been set.

### GetCreated

`func (o *Bookmark) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Bookmark) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Bookmark) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Bookmark) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCustomerId

`func (o *Bookmark) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Bookmark) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Bookmark) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Bookmark) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerNo

`func (o *Bookmark) GetCustomerNo() string`

GetCustomerNo returns the CustomerNo field if non-nil, zero value otherwise.

### GetCustomerNoOk

`func (o *Bookmark) GetCustomerNoOk() (*string, bool)`

GetCustomerNoOk returns a tuple with the CustomerNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNo

`func (o *Bookmark) SetCustomerNo(v string)`

SetCustomerNo sets CustomerNo field to given value.

### HasCustomerNo

`func (o *Bookmark) HasCustomerNo() bool`

HasCustomerNo returns a boolean if a field has been set.

### GetDeleted

`func (o *Bookmark) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Bookmark) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Bookmark) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Bookmark) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetId

`func (o *Bookmark) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bookmark) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bookmark) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Bookmark) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *Bookmark) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Bookmark) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Bookmark) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Bookmark) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetStatus

`func (o *Bookmark) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Bookmark) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Bookmark) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Bookmark) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *Bookmark) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *Bookmark) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *Bookmark) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *Bookmark) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetSystemDeleted

`func (o *Bookmark) GetSystemDeleted() bool`

GetSystemDeleted returns the SystemDeleted field if non-nil, zero value otherwise.

### GetSystemDeletedOk

`func (o *Bookmark) GetSystemDeletedOk() (*bool, bool)`

GetSystemDeletedOk returns a tuple with the SystemDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDeleted

`func (o *Bookmark) SetSystemDeleted(v bool)`

SetSystemDeleted sets SystemDeleted field to given value.

### HasSystemDeleted

`func (o *Bookmark) HasSystemDeleted() bool`

HasSystemDeleted returns a boolean if a field has been set.

### GetUpdated

`func (o *Bookmark) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Bookmark) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Bookmark) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Bookmark) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetVesselList

`func (o *Bookmark) GetVesselList() []string`

GetVesselList returns the VesselList field if non-nil, zero value otherwise.

### GetVesselListOk

`func (o *Bookmark) GetVesselListOk() (*[]string, bool)`

GetVesselListOk returns a tuple with the VesselList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselList

`func (o *Bookmark) SetVesselList(v []string)`

SetVesselList sets VesselList field to given value.

### HasVesselList

`func (o *Bookmark) HasVesselList() bool`

HasVesselList returns a boolean if a field has been set.

### GetVoyageNo

`func (o *Bookmark) GetVoyageNo() string`

GetVoyageNo returns the VoyageNo field if non-nil, zero value otherwise.

### GetVoyageNoOk

`func (o *Bookmark) GetVoyageNoOk() (*string, bool)`

GetVoyageNoOk returns a tuple with the VoyageNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageNo

`func (o *Bookmark) SetVoyageNo(v string)`

SetVoyageNo sets VoyageNo field to given value.

### HasVoyageNo

`func (o *Bookmark) HasVoyageNo() bool`

HasVoyageNo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


