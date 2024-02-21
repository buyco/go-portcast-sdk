# ContainerEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Container Event Object Created Date | [optional] 
**EventRaw** | Pointer to **string** | Description of the event, as reported by the Carrier | [optional] 
**EventTime** | Pointer to **NullableTime** | The latest actual date-time of the event, as reported by the Carrier - Local Time | [optional] 
**EventTimeEstimated** | Pointer to **NullableTime** | The latest estimated date-time of the event, as reported by the Carrier - Local Time | [optional] 
**EventTypeCode** | Pointer to **string** | Portcast Standardized Event Code | [optional] 
**EventTypeName** | Pointer to **string** | Portcast Standardized Event Name | [optional] 
**Id** | Pointer to **string** | Unique Identifier for the Container Event Object | [optional] 
**LocationRaw** | Pointer to **NullableString** | Relates to a physical location where the container event takes place, as reported by the Carrier | [optional] 
**LocationTypeCode** | Pointer to **string** | Describes the stage at which the container event takes place - Code | [optional] 
**LocationTypeName** | Pointer to **string** | Describes the stage at which the container event takes place - Full Description | [optional] 
**ModeOfTransport** | Pointer to **string** | Mode of Transport which is used for the execution of the container event, as interpretted by Portcast | [optional] 
**PortCode** | Pointer to **NullableString** | Relates to the UNLOCODE for the location where the container event takes place, as standardized by Portcast | [optional] 
**PortName** | Pointer to **NullableString** | Relates to the location name where the container event takes place, as standardized by Portcast | [optional] 
**TerminalDetails** | Pointer to [**ContainerEventTerminalDetails**](ContainerEventTerminalDetails.md) |  | [optional] 
**Updated** | Pointer to **time.Time** | Container Event Object Updated Date | [optional] 
**VesselImo** | Pointer to **NullableFloat32** | Vessel IMO of the vessel associated with the event | [optional] 
**VesselName** | Pointer to **NullableString** | Vessel Name of the vessel associated with the event | [optional] 

## Methods

### NewContainerEvent

`func NewContainerEvent() *ContainerEvent`

NewContainerEvent instantiates a new ContainerEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerEventWithDefaults

`func NewContainerEventWithDefaults() *ContainerEvent`

NewContainerEventWithDefaults instantiates a new ContainerEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ContainerEvent) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ContainerEvent) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ContainerEvent) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ContainerEvent) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEventRaw

`func (o *ContainerEvent) GetEventRaw() string`

GetEventRaw returns the EventRaw field if non-nil, zero value otherwise.

### GetEventRawOk

`func (o *ContainerEvent) GetEventRawOk() (*string, bool)`

GetEventRawOk returns a tuple with the EventRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRaw

`func (o *ContainerEvent) SetEventRaw(v string)`

SetEventRaw sets EventRaw field to given value.

### HasEventRaw

`func (o *ContainerEvent) HasEventRaw() bool`

HasEventRaw returns a boolean if a field has been set.

### GetEventTime

`func (o *ContainerEvent) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *ContainerEvent) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *ContainerEvent) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *ContainerEvent) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### SetEventTimeNil

`func (o *ContainerEvent) SetEventTimeNil(b bool)`

 SetEventTimeNil sets the value for EventTime to be an explicit nil

### UnsetEventTime
`func (o *ContainerEvent) UnsetEventTime()`

UnsetEventTime ensures that no value is present for EventTime, not even an explicit nil
### GetEventTimeEstimated

`func (o *ContainerEvent) GetEventTimeEstimated() time.Time`

GetEventTimeEstimated returns the EventTimeEstimated field if non-nil, zero value otherwise.

### GetEventTimeEstimatedOk

`func (o *ContainerEvent) GetEventTimeEstimatedOk() (*time.Time, bool)`

GetEventTimeEstimatedOk returns a tuple with the EventTimeEstimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimeEstimated

`func (o *ContainerEvent) SetEventTimeEstimated(v time.Time)`

SetEventTimeEstimated sets EventTimeEstimated field to given value.

### HasEventTimeEstimated

`func (o *ContainerEvent) HasEventTimeEstimated() bool`

HasEventTimeEstimated returns a boolean if a field has been set.

### SetEventTimeEstimatedNil

`func (o *ContainerEvent) SetEventTimeEstimatedNil(b bool)`

 SetEventTimeEstimatedNil sets the value for EventTimeEstimated to be an explicit nil

### UnsetEventTimeEstimated
`func (o *ContainerEvent) UnsetEventTimeEstimated()`

UnsetEventTimeEstimated ensures that no value is present for EventTimeEstimated, not even an explicit nil
### GetEventTypeCode

`func (o *ContainerEvent) GetEventTypeCode() string`

GetEventTypeCode returns the EventTypeCode field if non-nil, zero value otherwise.

### GetEventTypeCodeOk

`func (o *ContainerEvent) GetEventTypeCodeOk() (*string, bool)`

GetEventTypeCodeOk returns a tuple with the EventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeCode

`func (o *ContainerEvent) SetEventTypeCode(v string)`

SetEventTypeCode sets EventTypeCode field to given value.

### HasEventTypeCode

`func (o *ContainerEvent) HasEventTypeCode() bool`

HasEventTypeCode returns a boolean if a field has been set.

### GetEventTypeName

`func (o *ContainerEvent) GetEventTypeName() string`

GetEventTypeName returns the EventTypeName field if non-nil, zero value otherwise.

### GetEventTypeNameOk

`func (o *ContainerEvent) GetEventTypeNameOk() (*string, bool)`

GetEventTypeNameOk returns a tuple with the EventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeName

`func (o *ContainerEvent) SetEventTypeName(v string)`

SetEventTypeName sets EventTypeName field to given value.

### HasEventTypeName

`func (o *ContainerEvent) HasEventTypeName() bool`

HasEventTypeName returns a boolean if a field has been set.

### GetId

`func (o *ContainerEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContainerEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocationRaw

`func (o *ContainerEvent) GetLocationRaw() string`

GetLocationRaw returns the LocationRaw field if non-nil, zero value otherwise.

### GetLocationRawOk

`func (o *ContainerEvent) GetLocationRawOk() (*string, bool)`

GetLocationRawOk returns a tuple with the LocationRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationRaw

`func (o *ContainerEvent) SetLocationRaw(v string)`

SetLocationRaw sets LocationRaw field to given value.

### HasLocationRaw

`func (o *ContainerEvent) HasLocationRaw() bool`

HasLocationRaw returns a boolean if a field has been set.

### SetLocationRawNil

`func (o *ContainerEvent) SetLocationRawNil(b bool)`

 SetLocationRawNil sets the value for LocationRaw to be an explicit nil

### UnsetLocationRaw
`func (o *ContainerEvent) UnsetLocationRaw()`

UnsetLocationRaw ensures that no value is present for LocationRaw, not even an explicit nil
### GetLocationTypeCode

`func (o *ContainerEvent) GetLocationTypeCode() string`

GetLocationTypeCode returns the LocationTypeCode field if non-nil, zero value otherwise.

### GetLocationTypeCodeOk

`func (o *ContainerEvent) GetLocationTypeCodeOk() (*string, bool)`

GetLocationTypeCodeOk returns a tuple with the LocationTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationTypeCode

`func (o *ContainerEvent) SetLocationTypeCode(v string)`

SetLocationTypeCode sets LocationTypeCode field to given value.

### HasLocationTypeCode

`func (o *ContainerEvent) HasLocationTypeCode() bool`

HasLocationTypeCode returns a boolean if a field has been set.

### GetLocationTypeName

`func (o *ContainerEvent) GetLocationTypeName() string`

GetLocationTypeName returns the LocationTypeName field if non-nil, zero value otherwise.

### GetLocationTypeNameOk

`func (o *ContainerEvent) GetLocationTypeNameOk() (*string, bool)`

GetLocationTypeNameOk returns a tuple with the LocationTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationTypeName

`func (o *ContainerEvent) SetLocationTypeName(v string)`

SetLocationTypeName sets LocationTypeName field to given value.

### HasLocationTypeName

`func (o *ContainerEvent) HasLocationTypeName() bool`

HasLocationTypeName returns a boolean if a field has been set.

### GetModeOfTransport

`func (o *ContainerEvent) GetModeOfTransport() string`

GetModeOfTransport returns the ModeOfTransport field if non-nil, zero value otherwise.

### GetModeOfTransportOk

`func (o *ContainerEvent) GetModeOfTransportOk() (*string, bool)`

GetModeOfTransportOk returns a tuple with the ModeOfTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeOfTransport

`func (o *ContainerEvent) SetModeOfTransport(v string)`

SetModeOfTransport sets ModeOfTransport field to given value.

### HasModeOfTransport

`func (o *ContainerEvent) HasModeOfTransport() bool`

HasModeOfTransport returns a boolean if a field has been set.

### GetPortCode

`func (o *ContainerEvent) GetPortCode() string`

GetPortCode returns the PortCode field if non-nil, zero value otherwise.

### GetPortCodeOk

`func (o *ContainerEvent) GetPortCodeOk() (*string, bool)`

GetPortCodeOk returns a tuple with the PortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCode

`func (o *ContainerEvent) SetPortCode(v string)`

SetPortCode sets PortCode field to given value.

### HasPortCode

`func (o *ContainerEvent) HasPortCode() bool`

HasPortCode returns a boolean if a field has been set.

### SetPortCodeNil

`func (o *ContainerEvent) SetPortCodeNil(b bool)`

 SetPortCodeNil sets the value for PortCode to be an explicit nil

### UnsetPortCode
`func (o *ContainerEvent) UnsetPortCode()`

UnsetPortCode ensures that no value is present for PortCode, not even an explicit nil
### GetPortName

`func (o *ContainerEvent) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *ContainerEvent) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *ContainerEvent) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *ContainerEvent) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### SetPortNameNil

`func (o *ContainerEvent) SetPortNameNil(b bool)`

 SetPortNameNil sets the value for PortName to be an explicit nil

### UnsetPortName
`func (o *ContainerEvent) UnsetPortName()`

UnsetPortName ensures that no value is present for PortName, not even an explicit nil
### GetTerminalDetails

`func (o *ContainerEvent) GetTerminalDetails() ContainerEventTerminalDetails`

GetTerminalDetails returns the TerminalDetails field if non-nil, zero value otherwise.

### GetTerminalDetailsOk

`func (o *ContainerEvent) GetTerminalDetailsOk() (*ContainerEventTerminalDetails, bool)`

GetTerminalDetailsOk returns a tuple with the TerminalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalDetails

`func (o *ContainerEvent) SetTerminalDetails(v ContainerEventTerminalDetails)`

SetTerminalDetails sets TerminalDetails field to given value.

### HasTerminalDetails

`func (o *ContainerEvent) HasTerminalDetails() bool`

HasTerminalDetails returns a boolean if a field has been set.

### GetUpdated

`func (o *ContainerEvent) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ContainerEvent) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ContainerEvent) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ContainerEvent) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetVesselImo

`func (o *ContainerEvent) GetVesselImo() float32`

GetVesselImo returns the VesselImo field if non-nil, zero value otherwise.

### GetVesselImoOk

`func (o *ContainerEvent) GetVesselImoOk() (*float32, bool)`

GetVesselImoOk returns a tuple with the VesselImo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselImo

`func (o *ContainerEvent) SetVesselImo(v float32)`

SetVesselImo sets VesselImo field to given value.

### HasVesselImo

`func (o *ContainerEvent) HasVesselImo() bool`

HasVesselImo returns a boolean if a field has been set.

### SetVesselImoNil

`func (o *ContainerEvent) SetVesselImoNil(b bool)`

 SetVesselImoNil sets the value for VesselImo to be an explicit nil

### UnsetVesselImo
`func (o *ContainerEvent) UnsetVesselImo()`

UnsetVesselImo ensures that no value is present for VesselImo, not even an explicit nil
### GetVesselName

`func (o *ContainerEvent) GetVesselName() string`

GetVesselName returns the VesselName field if non-nil, zero value otherwise.

### GetVesselNameOk

`func (o *ContainerEvent) GetVesselNameOk() (*string, bool)`

GetVesselNameOk returns a tuple with the VesselName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselName

`func (o *ContainerEvent) SetVesselName(v string)`

SetVesselName sets VesselName field to given value.

### HasVesselName

`func (o *ContainerEvent) HasVesselName() bool`

HasVesselName returns a boolean if a field has been set.

### SetVesselNameNil

`func (o *ContainerEvent) SetVesselNameNil(b bool)`

 SetVesselNameNil sets the value for VesselName to be an explicit nil

### UnsetVesselName
`func (o *ContainerEvent) UnsetVesselName()`

UnsetVesselName ensures that no value is present for VesselName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


