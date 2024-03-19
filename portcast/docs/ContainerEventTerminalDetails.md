# ContainerEventTerminalDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TerminalCode** | Pointer to **string** | BIC/SMDG Defined Code for the Terminal Location | [optional] 
**TerminalCodeSource** | Pointer to **string** | Source of Terminal Code: BIC/SMDG | [optional] 
**TerminalName** | Pointer to **string** | Terminal Name where the container event takes place | [optional] 

## Methods

### NewContainerEventTerminalDetails

`func NewContainerEventTerminalDetails() *ContainerEventTerminalDetails`

NewContainerEventTerminalDetails instantiates a new ContainerEventTerminalDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerEventTerminalDetailsWithDefaults

`func NewContainerEventTerminalDetailsWithDefaults() *ContainerEventTerminalDetails`

NewContainerEventTerminalDetailsWithDefaults instantiates a new ContainerEventTerminalDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerminalCode

`func (o *ContainerEventTerminalDetails) GetTerminalCode() string`

GetTerminalCode returns the TerminalCode field if non-nil, zero value otherwise.

### GetTerminalCodeOk

`func (o *ContainerEventTerminalDetails) GetTerminalCodeOk() (*string, bool)`

GetTerminalCodeOk returns a tuple with the TerminalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalCode

`func (o *ContainerEventTerminalDetails) SetTerminalCode(v string)`

SetTerminalCode sets TerminalCode field to given value.

### HasTerminalCode

`func (o *ContainerEventTerminalDetails) HasTerminalCode() bool`

HasTerminalCode returns a boolean if a field has been set.

### GetTerminalCodeSource

`func (o *ContainerEventTerminalDetails) GetTerminalCodeSource() string`

GetTerminalCodeSource returns the TerminalCodeSource field if non-nil, zero value otherwise.

### GetTerminalCodeSourceOk

`func (o *ContainerEventTerminalDetails) GetTerminalCodeSourceOk() (*string, bool)`

GetTerminalCodeSourceOk returns a tuple with the TerminalCodeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalCodeSource

`func (o *ContainerEventTerminalDetails) SetTerminalCodeSource(v string)`

SetTerminalCodeSource sets TerminalCodeSource field to given value.

### HasTerminalCodeSource

`func (o *ContainerEventTerminalDetails) HasTerminalCodeSource() bool`

HasTerminalCodeSource returns a boolean if a field has been set.

### GetTerminalName

`func (o *ContainerEventTerminalDetails) GetTerminalName() string`

GetTerminalName returns the TerminalName field if non-nil, zero value otherwise.

### GetTerminalNameOk

`func (o *ContainerEventTerminalDetails) GetTerminalNameOk() (*string, bool)`

GetTerminalNameOk returns a tuple with the TerminalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalName

`func (o *ContainerEventTerminalDetails) SetTerminalName(v string)`

SetTerminalName sets TerminalName field to given value.

### HasTerminalName

`func (o *ContainerEventTerminalDetails) HasTerminalName() bool`

HasTerminalName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


