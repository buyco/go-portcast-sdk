# ContainerMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CntrNo** | Pointer to **string** | Container Number | [optional] 
**DetailSt** | Pointer to **string** | Container Size and Type Code as defined by [ISO 6346](https://en.wikipedia.org/wiki/ISO_6346#Size_and_type_codes) | [optional] 
**ExternalHeightFt** | Pointer to **float32** | Container External Body Height | [optional] 
**ExternalLengthFt** | Pointer to **float32** | Container External Body Length | [optional] 
**ExternalWidthFt** | Pointer to **float32** | Container External Body Width | [optional] 

## Methods

### NewContainerMetadata

`func NewContainerMetadata() *ContainerMetadata`

NewContainerMetadata instantiates a new ContainerMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerMetadataWithDefaults

`func NewContainerMetadataWithDefaults() *ContainerMetadata`

NewContainerMetadataWithDefaults instantiates a new ContainerMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCntrNo

`func (o *ContainerMetadata) GetCntrNo() string`

GetCntrNo returns the CntrNo field if non-nil, zero value otherwise.

### GetCntrNoOk

`func (o *ContainerMetadata) GetCntrNoOk() (*string, bool)`

GetCntrNoOk returns a tuple with the CntrNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCntrNo

`func (o *ContainerMetadata) SetCntrNo(v string)`

SetCntrNo sets CntrNo field to given value.

### HasCntrNo

`func (o *ContainerMetadata) HasCntrNo() bool`

HasCntrNo returns a boolean if a field has been set.

### GetDetailSt

`func (o *ContainerMetadata) GetDetailSt() string`

GetDetailSt returns the DetailSt field if non-nil, zero value otherwise.

### GetDetailStOk

`func (o *ContainerMetadata) GetDetailStOk() (*string, bool)`

GetDetailStOk returns a tuple with the DetailSt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailSt

`func (o *ContainerMetadata) SetDetailSt(v string)`

SetDetailSt sets DetailSt field to given value.

### HasDetailSt

`func (o *ContainerMetadata) HasDetailSt() bool`

HasDetailSt returns a boolean if a field has been set.

### GetExternalHeightFt

`func (o *ContainerMetadata) GetExternalHeightFt() float32`

GetExternalHeightFt returns the ExternalHeightFt field if non-nil, zero value otherwise.

### GetExternalHeightFtOk

`func (o *ContainerMetadata) GetExternalHeightFtOk() (*float32, bool)`

GetExternalHeightFtOk returns a tuple with the ExternalHeightFt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHeightFt

`func (o *ContainerMetadata) SetExternalHeightFt(v float32)`

SetExternalHeightFt sets ExternalHeightFt field to given value.

### HasExternalHeightFt

`func (o *ContainerMetadata) HasExternalHeightFt() bool`

HasExternalHeightFt returns a boolean if a field has been set.

### GetExternalLengthFt

`func (o *ContainerMetadata) GetExternalLengthFt() float32`

GetExternalLengthFt returns the ExternalLengthFt field if non-nil, zero value otherwise.

### GetExternalLengthFtOk

`func (o *ContainerMetadata) GetExternalLengthFtOk() (*float32, bool)`

GetExternalLengthFtOk returns a tuple with the ExternalLengthFt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLengthFt

`func (o *ContainerMetadata) SetExternalLengthFt(v float32)`

SetExternalLengthFt sets ExternalLengthFt field to given value.

### HasExternalLengthFt

`func (o *ContainerMetadata) HasExternalLengthFt() bool`

HasExternalLengthFt returns a boolean if a field has been set.

### GetExternalWidthFt

`func (o *ContainerMetadata) GetExternalWidthFt() float32`

GetExternalWidthFt returns the ExternalWidthFt field if non-nil, zero value otherwise.

### GetExternalWidthFtOk

`func (o *ContainerMetadata) GetExternalWidthFtOk() (*float32, bool)`

GetExternalWidthFtOk returns a tuple with the ExternalWidthFt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalWidthFt

`func (o *ContainerMetadata) SetExternalWidthFt(v float32)`

SetExternalWidthFt sets ExternalWidthFt field to given value.

### HasExternalWidthFt

`func (o *ContainerMetadata) HasExternalWidthFt() bool`

HasExternalWidthFt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


