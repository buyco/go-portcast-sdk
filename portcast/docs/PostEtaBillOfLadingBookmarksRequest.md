# PostEtaBillOfLadingBookmarksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierNo** | **string** | Carrier SCAC Code or Carrier Synonym (AUTO Detection Supported) | 
**BlNo** | Pointer to **string** | Carrier Provided Master Bill of Lading Number or Booking Number [Leave blank if not available] | [optional] 
**CntrNo** | **string** | Container Number (ISO6346) | 
**CallbackUrl** | Pointer to **string** | A https endpoint for Portcast to push the json object whenever there are updates. | [optional] 

## Methods

### NewPostEtaBillOfLadingBookmarksRequest

`func NewPostEtaBillOfLadingBookmarksRequest(carrierNo string, cntrNo string, ) *PostEtaBillOfLadingBookmarksRequest`

NewPostEtaBillOfLadingBookmarksRequest instantiates a new PostEtaBillOfLadingBookmarksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostEtaBillOfLadingBookmarksRequestWithDefaults

`func NewPostEtaBillOfLadingBookmarksRequestWithDefaults() *PostEtaBillOfLadingBookmarksRequest`

NewPostEtaBillOfLadingBookmarksRequestWithDefaults instantiates a new PostEtaBillOfLadingBookmarksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierNo

`func (o *PostEtaBillOfLadingBookmarksRequest) GetCarrierNo() string`

GetCarrierNo returns the CarrierNo field if non-nil, zero value otherwise.

### GetCarrierNoOk

`func (o *PostEtaBillOfLadingBookmarksRequest) GetCarrierNoOk() (*string, bool)`

GetCarrierNoOk returns a tuple with the CarrierNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierNo

`func (o *PostEtaBillOfLadingBookmarksRequest) SetCarrierNo(v string)`

SetCarrierNo sets CarrierNo field to given value.


### GetBlNo

`func (o *PostEtaBillOfLadingBookmarksRequest) GetBlNo() string`

GetBlNo returns the BlNo field if non-nil, zero value otherwise.

### GetBlNoOk

`func (o *PostEtaBillOfLadingBookmarksRequest) GetBlNoOk() (*string, bool)`

GetBlNoOk returns a tuple with the BlNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlNo

`func (o *PostEtaBillOfLadingBookmarksRequest) SetBlNo(v string)`

SetBlNo sets BlNo field to given value.

### HasBlNo

`func (o *PostEtaBillOfLadingBookmarksRequest) HasBlNo() bool`

HasBlNo returns a boolean if a field has been set.

### GetCntrNo

`func (o *PostEtaBillOfLadingBookmarksRequest) GetCntrNo() string`

GetCntrNo returns the CntrNo field if non-nil, zero value otherwise.

### GetCntrNoOk

`func (o *PostEtaBillOfLadingBookmarksRequest) GetCntrNoOk() (*string, bool)`

GetCntrNoOk returns a tuple with the CntrNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCntrNo

`func (o *PostEtaBillOfLadingBookmarksRequest) SetCntrNo(v string)`

SetCntrNo sets CntrNo field to given value.


### GetCallbackUrl

`func (o *PostEtaBillOfLadingBookmarksRequest) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *PostEtaBillOfLadingBookmarksRequest) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *PostEtaBillOfLadingBookmarksRequest) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *PostEtaBillOfLadingBookmarksRequest) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


