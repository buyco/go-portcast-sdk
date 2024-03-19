# PortcastAPICo2Emissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **NullableInt32** | Total emissions for each container:   (gm/ TEU)    Sum of the emissions derived from WTT and TTW analyses. This offers the carbon footprint associated with the entire lifecycle of the shipped goods.   | [optional] 
**Wtt** | Pointer to **NullableInt32** | Well-to-Tank (WTT) monitoring:   (gm/ TEU)    Measurement of carbon emissions across the entire fuel supply chain; that is, the energy required to extract, produce, and transport a fuel from its source.   | [optional] 
**Ttw** | Pointer to **NullableInt32** | Tank-to-Wheels (TTW) monitoring:   (gm/ TEU)    Measurement of emissions generated during the actual vessel operation, providing a real-time understanding of carbon output during voyages.   | [optional] 
**Intensity** | Pointer to **NullableFloat32** | Co2e intensity:   (kg/ ton-km)    Signifies the carbon efficiency for each container, thus empowering stakeholders to define actions on reducing emissions.   | [optional] 

## Methods

### NewPortcastAPICo2Emissions

`func NewPortcastAPICo2Emissions() *PortcastAPICo2Emissions`

NewPortcastAPICo2Emissions instantiates a new PortcastAPICo2Emissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortcastAPICo2EmissionsWithDefaults

`func NewPortcastAPICo2EmissionsWithDefaults() *PortcastAPICo2Emissions`

NewPortcastAPICo2EmissionsWithDefaults instantiates a new PortcastAPICo2Emissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *PortcastAPICo2Emissions) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PortcastAPICo2Emissions) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PortcastAPICo2Emissions) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PortcastAPICo2Emissions) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *PortcastAPICo2Emissions) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *PortcastAPICo2Emissions) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetWtt

`func (o *PortcastAPICo2Emissions) GetWtt() int32`

GetWtt returns the Wtt field if non-nil, zero value otherwise.

### GetWttOk

`func (o *PortcastAPICo2Emissions) GetWttOk() (*int32, bool)`

GetWttOk returns a tuple with the Wtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWtt

`func (o *PortcastAPICo2Emissions) SetWtt(v int32)`

SetWtt sets Wtt field to given value.

### HasWtt

`func (o *PortcastAPICo2Emissions) HasWtt() bool`

HasWtt returns a boolean if a field has been set.

### SetWttNil

`func (o *PortcastAPICo2Emissions) SetWttNil(b bool)`

 SetWttNil sets the value for Wtt to be an explicit nil

### UnsetWtt
`func (o *PortcastAPICo2Emissions) UnsetWtt()`

UnsetWtt ensures that no value is present for Wtt, not even an explicit nil
### GetTtw

`func (o *PortcastAPICo2Emissions) GetTtw() int32`

GetTtw returns the Ttw field if non-nil, zero value otherwise.

### GetTtwOk

`func (o *PortcastAPICo2Emissions) GetTtwOk() (*int32, bool)`

GetTtwOk returns a tuple with the Ttw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtw

`func (o *PortcastAPICo2Emissions) SetTtw(v int32)`

SetTtw sets Ttw field to given value.

### HasTtw

`func (o *PortcastAPICo2Emissions) HasTtw() bool`

HasTtw returns a boolean if a field has been set.

### SetTtwNil

`func (o *PortcastAPICo2Emissions) SetTtwNil(b bool)`

 SetTtwNil sets the value for Ttw to be an explicit nil

### UnsetTtw
`func (o *PortcastAPICo2Emissions) UnsetTtw()`

UnsetTtw ensures that no value is present for Ttw, not even an explicit nil
### GetIntensity

`func (o *PortcastAPICo2Emissions) GetIntensity() float32`

GetIntensity returns the Intensity field if non-nil, zero value otherwise.

### GetIntensityOk

`func (o *PortcastAPICo2Emissions) GetIntensityOk() (*float32, bool)`

GetIntensityOk returns a tuple with the Intensity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntensity

`func (o *PortcastAPICo2Emissions) SetIntensity(v float32)`

SetIntensity sets Intensity field to given value.

### HasIntensity

`func (o *PortcastAPICo2Emissions) HasIntensity() bool`

HasIntensity returns a boolean if a field has been set.

### SetIntensityNil

`func (o *PortcastAPICo2Emissions) SetIntensityNil(b bool)`

 SetIntensityNil sets the value for Intensity to be an explicit nil

### UnsetIntensity
`func (o *PortcastAPICo2Emissions) UnsetIntensity()`

UnsetIntensity ensures that no value is present for Intensity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


