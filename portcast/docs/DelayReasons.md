# DelayReasons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Container Delay Reason Object Created Date | [optional] 
**DelayDescription** | Pointer to **string** | Description of the Delay Reason | [optional] 
**LocationTypeCode** | Pointer to **string** | Location qualifier to define if the delay incident has taken place at the origin port (POL) or transshipment port (POT) | [optional] 
**PortCode** | Pointer to **string** | Relates to the UNLOCODE for the location to which the delay incident is associated with | [optional] 
**PortName** | Pointer to **string** | Relates to the location name to which the delay incident is associated with | [optional] 
**ReasonCode** | Pointer to **string** | Delay Incident Code: RLV(Vessel Rollover), LTD(Late Departure), PRC(Port Rotation Changes) | [optional] 
**RotationChange** | Pointer to [**DelayReasonsRotationChange**](DelayReasonsRotationChange.md) |  | [optional] 
**ScheduleChange** | Pointer to [**[]DelayReasonsScheduleChangeInner**](DelayReasonsScheduleChangeInner.md) | Detailed view into impact of delay incidents into the container&#39;s arrival at POD, loading and departure events. | [optional] 
**VesselChange** | Pointer to [**DelayReasonsVesselChange**](DelayReasonsVesselChange.md) |  | [optional] 

## Methods

### NewDelayReasons

`func NewDelayReasons() *DelayReasons`

NewDelayReasons instantiates a new DelayReasons object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelayReasonsWithDefaults

`func NewDelayReasonsWithDefaults() *DelayReasons`

NewDelayReasonsWithDefaults instantiates a new DelayReasons object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *DelayReasons) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DelayReasons) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DelayReasons) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DelayReasons) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDelayDescription

`func (o *DelayReasons) GetDelayDescription() string`

GetDelayDescription returns the DelayDescription field if non-nil, zero value otherwise.

### GetDelayDescriptionOk

`func (o *DelayReasons) GetDelayDescriptionOk() (*string, bool)`

GetDelayDescriptionOk returns a tuple with the DelayDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayDescription

`func (o *DelayReasons) SetDelayDescription(v string)`

SetDelayDescription sets DelayDescription field to given value.

### HasDelayDescription

`func (o *DelayReasons) HasDelayDescription() bool`

HasDelayDescription returns a boolean if a field has been set.

### GetLocationTypeCode

`func (o *DelayReasons) GetLocationTypeCode() string`

GetLocationTypeCode returns the LocationTypeCode field if non-nil, zero value otherwise.

### GetLocationTypeCodeOk

`func (o *DelayReasons) GetLocationTypeCodeOk() (*string, bool)`

GetLocationTypeCodeOk returns a tuple with the LocationTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationTypeCode

`func (o *DelayReasons) SetLocationTypeCode(v string)`

SetLocationTypeCode sets LocationTypeCode field to given value.

### HasLocationTypeCode

`func (o *DelayReasons) HasLocationTypeCode() bool`

HasLocationTypeCode returns a boolean if a field has been set.

### GetPortCode

`func (o *DelayReasons) GetPortCode() string`

GetPortCode returns the PortCode field if non-nil, zero value otherwise.

### GetPortCodeOk

`func (o *DelayReasons) GetPortCodeOk() (*string, bool)`

GetPortCodeOk returns a tuple with the PortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCode

`func (o *DelayReasons) SetPortCode(v string)`

SetPortCode sets PortCode field to given value.

### HasPortCode

`func (o *DelayReasons) HasPortCode() bool`

HasPortCode returns a boolean if a field has been set.

### GetPortName

`func (o *DelayReasons) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *DelayReasons) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *DelayReasons) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *DelayReasons) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetReasonCode

`func (o *DelayReasons) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *DelayReasons) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *DelayReasons) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *DelayReasons) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetRotationChange

`func (o *DelayReasons) GetRotationChange() DelayReasonsRotationChange`

GetRotationChange returns the RotationChange field if non-nil, zero value otherwise.

### GetRotationChangeOk

`func (o *DelayReasons) GetRotationChangeOk() (*DelayReasonsRotationChange, bool)`

GetRotationChangeOk returns a tuple with the RotationChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationChange

`func (o *DelayReasons) SetRotationChange(v DelayReasonsRotationChange)`

SetRotationChange sets RotationChange field to given value.

### HasRotationChange

`func (o *DelayReasons) HasRotationChange() bool`

HasRotationChange returns a boolean if a field has been set.

### GetScheduleChange

`func (o *DelayReasons) GetScheduleChange() []DelayReasonsScheduleChangeInner`

GetScheduleChange returns the ScheduleChange field if non-nil, zero value otherwise.

### GetScheduleChangeOk

`func (o *DelayReasons) GetScheduleChangeOk() (*[]DelayReasonsScheduleChangeInner, bool)`

GetScheduleChangeOk returns a tuple with the ScheduleChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleChange

`func (o *DelayReasons) SetScheduleChange(v []DelayReasonsScheduleChangeInner)`

SetScheduleChange sets ScheduleChange field to given value.

### HasScheduleChange

`func (o *DelayReasons) HasScheduleChange() bool`

HasScheduleChange returns a boolean if a field has been set.

### GetVesselChange

`func (o *DelayReasons) GetVesselChange() DelayReasonsVesselChange`

GetVesselChange returns the VesselChange field if non-nil, zero value otherwise.

### GetVesselChangeOk

`func (o *DelayReasons) GetVesselChangeOk() (*DelayReasonsVesselChange, bool)`

GetVesselChangeOk returns a tuple with the VesselChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselChange

`func (o *DelayReasons) SetVesselChange(v DelayReasonsVesselChange)`

SetVesselChange sets VesselChange field to given value.

### HasVesselChange

`func (o *DelayReasons) HasVesselChange() bool`

HasVesselChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


