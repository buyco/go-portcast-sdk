# PortcastAPIStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] [default to "SUCCESS"]
**Metadata** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPortcastAPIStatusInfo

`func NewPortcastAPIStatusInfo() *PortcastAPIStatusInfo`

NewPortcastAPIStatusInfo instantiates a new PortcastAPIStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortcastAPIStatusInfoWithDefaults

`func NewPortcastAPIStatusInfoWithDefaults() *PortcastAPIStatusInfo`

NewPortcastAPIStatusInfoWithDefaults instantiates a new PortcastAPIStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PortcastAPIStatusInfo) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PortcastAPIStatusInfo) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PortcastAPIStatusInfo) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PortcastAPIStatusInfo) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMetadata

`func (o *PortcastAPIStatusInfo) GetMetadata() []string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PortcastAPIStatusInfo) GetMetadataOk() (*[]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PortcastAPIStatusInfo) SetMetadata(v []string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PortcastAPIStatusInfo) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


