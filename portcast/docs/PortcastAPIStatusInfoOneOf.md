# PortcastAPIStatusInfoOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | We didn&#39;t encounter any warnings or errors in any of Container Status Code Status, Vessel Status Code and Predictions Status Code or while generating the API Response | [optional] [default to "SUCCESS"]
**Metadata** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPortcastAPIStatusInfoOneOf

`func NewPortcastAPIStatusInfoOneOf() *PortcastAPIStatusInfoOneOf`

NewPortcastAPIStatusInfoOneOf instantiates a new PortcastAPIStatusInfoOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortcastAPIStatusInfoOneOfWithDefaults

`func NewPortcastAPIStatusInfoOneOfWithDefaults() *PortcastAPIStatusInfoOneOf`

NewPortcastAPIStatusInfoOneOfWithDefaults instantiates a new PortcastAPIStatusInfoOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PortcastAPIStatusInfoOneOf) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PortcastAPIStatusInfoOneOf) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PortcastAPIStatusInfoOneOf) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PortcastAPIStatusInfoOneOf) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMetadata

`func (o *PortcastAPIStatusInfoOneOf) GetMetadata() []string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PortcastAPIStatusInfoOneOf) GetMetadataOk() (*[]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PortcastAPIStatusInfoOneOf) SetMetadata(v []string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PortcastAPIStatusInfoOneOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


