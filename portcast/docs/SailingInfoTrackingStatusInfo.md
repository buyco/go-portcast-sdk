# SailingInfoTrackingStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prediction** | Pointer to [**SailingInfoTrackingStatusInfoPrediction**](SailingInfoTrackingStatusInfoPrediction.md) |  | [optional] 
**Vessel** | Pointer to [**SailingInfoTrackingStatusInfoVessel**](SailingInfoTrackingStatusInfoVessel.md) |  | [optional] 

## Methods

### NewSailingInfoTrackingStatusInfo

`func NewSailingInfoTrackingStatusInfo() *SailingInfoTrackingStatusInfo`

NewSailingInfoTrackingStatusInfo instantiates a new SailingInfoTrackingStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSailingInfoTrackingStatusInfoWithDefaults

`func NewSailingInfoTrackingStatusInfoWithDefaults() *SailingInfoTrackingStatusInfo`

NewSailingInfoTrackingStatusInfoWithDefaults instantiates a new SailingInfoTrackingStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrediction

`func (o *SailingInfoTrackingStatusInfo) GetPrediction() SailingInfoTrackingStatusInfoPrediction`

GetPrediction returns the Prediction field if non-nil, zero value otherwise.

### GetPredictionOk

`func (o *SailingInfoTrackingStatusInfo) GetPredictionOk() (*SailingInfoTrackingStatusInfoPrediction, bool)`

GetPredictionOk returns a tuple with the Prediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrediction

`func (o *SailingInfoTrackingStatusInfo) SetPrediction(v SailingInfoTrackingStatusInfoPrediction)`

SetPrediction sets Prediction field to given value.

### HasPrediction

`func (o *SailingInfoTrackingStatusInfo) HasPrediction() bool`

HasPrediction returns a boolean if a field has been set.

### GetVessel

`func (o *SailingInfoTrackingStatusInfo) GetVessel() SailingInfoTrackingStatusInfoVessel`

GetVessel returns the Vessel field if non-nil, zero value otherwise.

### GetVesselOk

`func (o *SailingInfoTrackingStatusInfo) GetVesselOk() (*SailingInfoTrackingStatusInfoVessel, bool)`

GetVesselOk returns a tuple with the Vessel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVessel

`func (o *SailingInfoTrackingStatusInfo) SetVessel(v SailingInfoTrackingStatusInfoVessel)`

SetVessel sets Vessel field to given value.

### HasVessel

`func (o *SailingInfoTrackingStatusInfo) HasVessel() bool`

HasVessel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


