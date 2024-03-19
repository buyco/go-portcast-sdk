# PostBooking422Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] [default to "Master org is not allowed to perform this action"]

## Methods

### NewPostBooking422Response

`func NewPostBooking422Response() *PostBooking422Response`

NewPostBooking422Response instantiates a new PostBooking422Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostBooking422ResponseWithDefaults

`func NewPostBooking422ResponseWithDefaults() *PostBooking422Response`

NewPostBooking422ResponseWithDefaults instantiates a new PostBooking422Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PostBooking422Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PostBooking422Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PostBooking422Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PostBooking422Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


