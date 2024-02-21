# PortcastAPIStatusInfoOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | We had a WARNING in one or more of Vessel Status Code and/or Predictions Status Code or no data due to bad input information | [optional] [default to "WARNING"]
**Metadata** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPortcastAPIStatusInfoOneOf1

`func NewPortcastAPIStatusInfoOneOf1() *PortcastAPIStatusInfoOneOf1`

NewPortcastAPIStatusInfoOneOf1 instantiates a new PortcastAPIStatusInfoOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortcastAPIStatusInfoOneOf1WithDefaults

`func NewPortcastAPIStatusInfoOneOf1WithDefaults() *PortcastAPIStatusInfoOneOf1`

NewPortcastAPIStatusInfoOneOf1WithDefaults instantiates a new PortcastAPIStatusInfoOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PortcastAPIStatusInfoOneOf1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PortcastAPIStatusInfoOneOf1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PortcastAPIStatusInfoOneOf1) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PortcastAPIStatusInfoOneOf1) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMetadata

`func (o *PortcastAPIStatusInfoOneOf1) GetMetadata() []string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PortcastAPIStatusInfoOneOf1) GetMetadataOk() (*[]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PortcastAPIStatusInfoOneOf1) SetMetadata(v []string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PortcastAPIStatusInfoOneOf1) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


