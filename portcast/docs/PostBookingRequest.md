# PostBookingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierNo** | **string** | Carrier SCAC Code or Carrier Synonym (AUTO Detection Supported) | 
**DocNo** | **string** | Booking or Bill of Lading number | 
**DocType** | **string** | Document number type: BK (Booking) or BL (Bill of Lading) | 
**CallbackUrl** | Pointer to **string** | A https endpoint for Portcast to push the json object whenever there are updates. | [optional] 

## Methods

### NewPostBookingRequest

`func NewPostBookingRequest(carrierNo string, docNo string, docType string, ) *PostBookingRequest`

NewPostBookingRequest instantiates a new PostBookingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostBookingRequestWithDefaults

`func NewPostBookingRequestWithDefaults() *PostBookingRequest`

NewPostBookingRequestWithDefaults instantiates a new PostBookingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierNo

`func (o *PostBookingRequest) GetCarrierNo() string`

GetCarrierNo returns the CarrierNo field if non-nil, zero value otherwise.

### GetCarrierNoOk

`func (o *PostBookingRequest) GetCarrierNoOk() (*string, bool)`

GetCarrierNoOk returns a tuple with the CarrierNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierNo

`func (o *PostBookingRequest) SetCarrierNo(v string)`

SetCarrierNo sets CarrierNo field to given value.


### GetDocNo

`func (o *PostBookingRequest) GetDocNo() string`

GetDocNo returns the DocNo field if non-nil, zero value otherwise.

### GetDocNoOk

`func (o *PostBookingRequest) GetDocNoOk() (*string, bool)`

GetDocNoOk returns a tuple with the DocNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocNo

`func (o *PostBookingRequest) SetDocNo(v string)`

SetDocNo sets DocNo field to given value.


### GetDocType

`func (o *PostBookingRequest) GetDocType() string`

GetDocType returns the DocType field if non-nil, zero value otherwise.

### GetDocTypeOk

`func (o *PostBookingRequest) GetDocTypeOk() (*string, bool)`

GetDocTypeOk returns a tuple with the DocType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocType

`func (o *PostBookingRequest) SetDocType(v string)`

SetDocType sets DocType field to given value.


### GetCallbackUrl

`func (o *PostBookingRequest) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *PostBookingRequest) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *PostBookingRequest) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *PostBookingRequest) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


