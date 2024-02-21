# GetApiV1EtaBillOfLadingBookmarks404Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] [default to "Record not found with give Org/Bill Of Lading ID"]
**Message** | Pointer to **string** |  | [optional] [default to "You have requested this URI [/api/v1/eta/tracking/bill-of-lading-bookmarks/e48cb2fb-499b-4af4-a0ef-250eac4eee5f] but did you mean /api/v1/eta/tracking/bill-of-lading-bookmarks/<bill_of_lading_bookmark_id> or /api/v1/eta/tracking/bill-of-lading-bookmarks or /api/v1/orgs/<org_id>/eta/tracking/bill-of-lading-bookmarks ?"]

## Methods

### NewGetApiV1EtaBillOfLadingBookmarks404Response

`func NewGetApiV1EtaBillOfLadingBookmarks404Response() *GetApiV1EtaBillOfLadingBookmarks404Response`

NewGetApiV1EtaBillOfLadingBookmarks404Response instantiates a new GetApiV1EtaBillOfLadingBookmarks404Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiV1EtaBillOfLadingBookmarks404ResponseWithDefaults

`func NewGetApiV1EtaBillOfLadingBookmarks404ResponseWithDefaults() *GetApiV1EtaBillOfLadingBookmarks404Response`

NewGetApiV1EtaBillOfLadingBookmarks404ResponseWithDefaults instantiates a new GetApiV1EtaBillOfLadingBookmarks404Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetApiV1EtaBillOfLadingBookmarks404Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetApiV1EtaBillOfLadingBookmarks404Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetApiV1EtaBillOfLadingBookmarks404Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetApiV1EtaBillOfLadingBookmarks404Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *GetApiV1EtaBillOfLadingBookmarks404Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetApiV1EtaBillOfLadingBookmarks404Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetApiV1EtaBillOfLadingBookmarks404Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetApiV1EtaBillOfLadingBookmarks404Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


