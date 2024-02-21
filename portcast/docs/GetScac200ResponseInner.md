# GetScac200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlPrefixes** | Pointer to **[]string** | BL Prefixes observed for the carrier | [optional] 
**FullName** | Pointer to **string** | Full Name of the Carrier | [optional] 
**Name** | Pointer to **string** | Name of the Carrier | [optional] 
**Scac** | Pointer to **string** | Official Carrier SCAC Code | [optional] 
**ScacSynonyms** | Pointer to **[]string** | Other SCAC Synonyms for the carrier (If upload operation is done using any of these, we will convert it to the official SCAC) | [optional] 
**SupportedStatus** | Pointer to **NullableString** | Supported Status for the carrier, will highlight here if there are any gaps for a particular carrier&#39;s coverage. | [optional] 
**Updated** | Pointer to **time.Time** | Last Updated Date for the carrier scac metadata | [optional] 

## Methods

### NewGetScac200ResponseInner

`func NewGetScac200ResponseInner() *GetScac200ResponseInner`

NewGetScac200ResponseInner instantiates a new GetScac200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetScac200ResponseInnerWithDefaults

`func NewGetScac200ResponseInnerWithDefaults() *GetScac200ResponseInner`

NewGetScac200ResponseInnerWithDefaults instantiates a new GetScac200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlPrefixes

`func (o *GetScac200ResponseInner) GetBlPrefixes() []string`

GetBlPrefixes returns the BlPrefixes field if non-nil, zero value otherwise.

### GetBlPrefixesOk

`func (o *GetScac200ResponseInner) GetBlPrefixesOk() (*[]string, bool)`

GetBlPrefixesOk returns a tuple with the BlPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlPrefixes

`func (o *GetScac200ResponseInner) SetBlPrefixes(v []string)`

SetBlPrefixes sets BlPrefixes field to given value.

### HasBlPrefixes

`func (o *GetScac200ResponseInner) HasBlPrefixes() bool`

HasBlPrefixes returns a boolean if a field has been set.

### GetFullName

`func (o *GetScac200ResponseInner) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GetScac200ResponseInner) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GetScac200ResponseInner) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *GetScac200ResponseInner) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetName

`func (o *GetScac200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetScac200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetScac200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetScac200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScac

`func (o *GetScac200ResponseInner) GetScac() string`

GetScac returns the Scac field if non-nil, zero value otherwise.

### GetScacOk

`func (o *GetScac200ResponseInner) GetScacOk() (*string, bool)`

GetScacOk returns a tuple with the Scac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScac

`func (o *GetScac200ResponseInner) SetScac(v string)`

SetScac sets Scac field to given value.

### HasScac

`func (o *GetScac200ResponseInner) HasScac() bool`

HasScac returns a boolean if a field has been set.

### GetScacSynonyms

`func (o *GetScac200ResponseInner) GetScacSynonyms() []string`

GetScacSynonyms returns the ScacSynonyms field if non-nil, zero value otherwise.

### GetScacSynonymsOk

`func (o *GetScac200ResponseInner) GetScacSynonymsOk() (*[]string, bool)`

GetScacSynonymsOk returns a tuple with the ScacSynonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScacSynonyms

`func (o *GetScac200ResponseInner) SetScacSynonyms(v []string)`

SetScacSynonyms sets ScacSynonyms field to given value.

### HasScacSynonyms

`func (o *GetScac200ResponseInner) HasScacSynonyms() bool`

HasScacSynonyms returns a boolean if a field has been set.

### GetSupportedStatus

`func (o *GetScac200ResponseInner) GetSupportedStatus() string`

GetSupportedStatus returns the SupportedStatus field if non-nil, zero value otherwise.

### GetSupportedStatusOk

`func (o *GetScac200ResponseInner) GetSupportedStatusOk() (*string, bool)`

GetSupportedStatusOk returns a tuple with the SupportedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedStatus

`func (o *GetScac200ResponseInner) SetSupportedStatus(v string)`

SetSupportedStatus sets SupportedStatus field to given value.

### HasSupportedStatus

`func (o *GetScac200ResponseInner) HasSupportedStatus() bool`

HasSupportedStatus returns a boolean if a field has been set.

### SetSupportedStatusNil

`func (o *GetScac200ResponseInner) SetSupportedStatusNil(b bool)`

 SetSupportedStatusNil sets the value for SupportedStatus to be an explicit nil

### UnsetSupportedStatus
`func (o *GetScac200ResponseInner) UnsetSupportedStatus()`

UnsetSupportedStatus ensures that no value is present for SupportedStatus, not even an explicit nil
### GetUpdated

`func (o *GetScac200ResponseInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetScac200ResponseInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetScac200ResponseInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetScac200ResponseInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


