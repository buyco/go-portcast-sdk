# DelayReasonsScheduleChangeInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTypeCode** | Pointer to **string** | Defines if it&#39;s a loading, departure or arrival event | [optional] 
**LocationTypeCode** | Pointer to **string** | Defines if the change is happening at the POL, POT or POD | [optional] 
**NewDate** | Pointer to **string** | Latest changed datetime information - port local time | [optional] 
**OldDate** | Pointer to **string** | Last recorded datetime information - port local time | [optional] 

## Methods

### NewDelayReasonsScheduleChangeInner

`func NewDelayReasonsScheduleChangeInner() *DelayReasonsScheduleChangeInner`

NewDelayReasonsScheduleChangeInner instantiates a new DelayReasonsScheduleChangeInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelayReasonsScheduleChangeInnerWithDefaults

`func NewDelayReasonsScheduleChangeInnerWithDefaults() *DelayReasonsScheduleChangeInner`

NewDelayReasonsScheduleChangeInnerWithDefaults instantiates a new DelayReasonsScheduleChangeInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTypeCode

`func (o *DelayReasonsScheduleChangeInner) GetEventTypeCode() string`

GetEventTypeCode returns the EventTypeCode field if non-nil, zero value otherwise.

### GetEventTypeCodeOk

`func (o *DelayReasonsScheduleChangeInner) GetEventTypeCodeOk() (*string, bool)`

GetEventTypeCodeOk returns a tuple with the EventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeCode

`func (o *DelayReasonsScheduleChangeInner) SetEventTypeCode(v string)`

SetEventTypeCode sets EventTypeCode field to given value.

### HasEventTypeCode

`func (o *DelayReasonsScheduleChangeInner) HasEventTypeCode() bool`

HasEventTypeCode returns a boolean if a field has been set.

### GetLocationTypeCode

`func (o *DelayReasonsScheduleChangeInner) GetLocationTypeCode() string`

GetLocationTypeCode returns the LocationTypeCode field if non-nil, zero value otherwise.

### GetLocationTypeCodeOk

`func (o *DelayReasonsScheduleChangeInner) GetLocationTypeCodeOk() (*string, bool)`

GetLocationTypeCodeOk returns a tuple with the LocationTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationTypeCode

`func (o *DelayReasonsScheduleChangeInner) SetLocationTypeCode(v string)`

SetLocationTypeCode sets LocationTypeCode field to given value.

### HasLocationTypeCode

`func (o *DelayReasonsScheduleChangeInner) HasLocationTypeCode() bool`

HasLocationTypeCode returns a boolean if a field has been set.

### GetNewDate

`func (o *DelayReasonsScheduleChangeInner) GetNewDate() string`

GetNewDate returns the NewDate field if non-nil, zero value otherwise.

### GetNewDateOk

`func (o *DelayReasonsScheduleChangeInner) GetNewDateOk() (*string, bool)`

GetNewDateOk returns a tuple with the NewDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDate

`func (o *DelayReasonsScheduleChangeInner) SetNewDate(v string)`

SetNewDate sets NewDate field to given value.

### HasNewDate

`func (o *DelayReasonsScheduleChangeInner) HasNewDate() bool`

HasNewDate returns a boolean if a field has been set.

### GetOldDate

`func (o *DelayReasonsScheduleChangeInner) GetOldDate() string`

GetOldDate returns the OldDate field if non-nil, zero value otherwise.

### GetOldDateOk

`func (o *DelayReasonsScheduleChangeInner) GetOldDateOk() (*string, bool)`

GetOldDateOk returns a tuple with the OldDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldDate

`func (o *DelayReasonsScheduleChangeInner) SetOldDate(v string)`

SetOldDate sets OldDate field to given value.

### HasOldDate

`func (o *DelayReasonsScheduleChangeInner) HasOldDate() bool`

HasOldDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


