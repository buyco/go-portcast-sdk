/*
Portcast API (1.1.0) - Bill of Lading Tracking

**This documentation is for the latest version of the Portcast Bill of Lading Tracking API.**  There are two variables used in this documentation: 1. `api-url`: the url to use for accessing the API. The official url is `https://api.portcast.io` 2. `x-api-key`: the access token to send along with every request to the API. This key will be provided to each organisation upon API access activation  The general workflow is as below:  1. Create the bill of lading bookmark if it does not exist already (`POST {{api-url}}/api/v1/eta/bill-of-lading-bookmarks`). 2. A bookmark must contain `carrier_no`, `bl_no` and `cntr_no` information. This will return the bill of lading bookmark information created. Record the `id` field from the response. 3. Wait for predictions to be generated. This could take up to 5 mins. 5. Query for the tracking results based on the `id` recorded earlier (`GET {{api-url}}/api/v1/eta/tracking/bill-of-lading-bookmarks/<id>`) 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package portcast

import (
	"encoding/json"
	"time"
)

// TrackingEventVoyageDetails struct for TrackingEventVoyageDetails
type TrackingEventVoyageDetails struct {
	ActiveScac *string `json:"active_scac,omitempty"`
	ActualArrivalLt *time.Time `json:"actual_arrival_lt,omitempty"`
	ActualArrivalUtc *time.Time `json:"actual_arrival_utc,omitempty"`
	ActualDepartureLt *time.Time `json:"actual_departure_lt,omitempty"`
	ActualDepartureUtc *time.Time `json:"actual_departure_utc,omitempty"`
	Id *string `json:"id,omitempty"`
	Index *int32 `json:"index,omitempty"`
	Lat *float32 `json:"lat,omitempty"`
	Lon *float32 `json:"lon,omitempty"`
	OriginalVoyageNo *string `json:"original_voyage_no,omitempty"`
	PortCode *string `json:"port_code,omitempty"`
	PortName *string `json:"port_name,omitempty"`
	PredictedArrivalLt *string `json:"predicted_arrival_lt,omitempty"`
	PredictedArrivalUtc *string `json:"predicted_arrival_utc,omitempty"`
	PredictedDepartureLt *string `json:"predicted_departure_lt,omitempty"`
	PredictedDepartureUtc *string `json:"predicted_departure_utc,omitempty"`
	PredictionTimeUtc *string `json:"prediction_time_utc,omitempty"`
	ScheduledArrivalLt *time.Time `json:"scheduled_arrival_lt,omitempty"`
	ScheduledArrivalOnTimeConfidence *string `json:"scheduled_arrival_on_time_confidence,omitempty"`
	ScheduledArrivalUtc *time.Time `json:"scheduled_arrival_utc,omitempty"`
	ScheduledDepartureLt *time.Time `json:"scheduled_departure_lt,omitempty"`
	ScheduledDepartureOnTimeConfidence *string `json:"scheduled_departure_on_time_confidence,omitempty"`
	ScheduledDepartureUtc *time.Time `json:"scheduled_departure_utc,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	VoyageNoList *[]string `json:"voyage_no_list,omitempty"`
}

// NewTrackingEventVoyageDetails instantiates a new TrackingEventVoyageDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingEventVoyageDetails() *TrackingEventVoyageDetails {
	this := TrackingEventVoyageDetails{}
	return &this
}

// NewTrackingEventVoyageDetailsWithDefaults instantiates a new TrackingEventVoyageDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingEventVoyageDetailsWithDefaults() *TrackingEventVoyageDetails {
	this := TrackingEventVoyageDetails{}
	return &this
}

// GetActiveScac returns the ActiveScac field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetActiveScac() string {
	if o == nil || o.ActiveScac == nil {
		var ret string
		return ret
	}
	return *o.ActiveScac
}

// GetActiveScacOk returns a tuple with the ActiveScac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetActiveScacOk() (*string, bool) {
	if o == nil || o.ActiveScac == nil {
		return nil, false
	}
	return o.ActiveScac, true
}

// HasActiveScac returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasActiveScac() bool {
	if o != nil && o.ActiveScac != nil {
		return true
	}

	return false
}

// SetActiveScac gets a reference to the given string and assigns it to the ActiveScac field.
func (o *TrackingEventVoyageDetails) SetActiveScac(v string) {
	o.ActiveScac = &v
}

// GetActualArrivalLt returns the ActualArrivalLt field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetActualArrivalLt() time.Time {
	if o == nil || o.ActualArrivalLt == nil {
		var ret time.Time
		return ret
	}
	return *o.ActualArrivalLt
}

// GetActualArrivalLtOk returns a tuple with the ActualArrivalLt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetActualArrivalLtOk() (*time.Time, bool) {
	if o == nil || o.ActualArrivalLt == nil {
		return nil, false
	}
	return o.ActualArrivalLt, true
}

// HasActualArrivalLt returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasActualArrivalLt() bool {
	if o != nil && o.ActualArrivalLt != nil {
		return true
	}

	return false
}

// SetActualArrivalLt gets a reference to the given time.Time and assigns it to the ActualArrivalLt field.
func (o *TrackingEventVoyageDetails) SetActualArrivalLt(v time.Time) {
	o.ActualArrivalLt = &v
}

// GetActualArrivalUtc returns the ActualArrivalUtc field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetActualArrivalUtc() time.Time {
	if o == nil || o.ActualArrivalUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.ActualArrivalUtc
}

// GetActualArrivalUtcOk returns a tuple with the ActualArrivalUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetActualArrivalUtcOk() (*time.Time, bool) {
	if o == nil || o.ActualArrivalUtc == nil {
		return nil, false
	}
	return o.ActualArrivalUtc, true
}

// HasActualArrivalUtc returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasActualArrivalUtc() bool {
	if o != nil && o.ActualArrivalUtc != nil {
		return true
	}

	return false
}

// SetActualArrivalUtc gets a reference to the given time.Time and assigns it to the ActualArrivalUtc field.
func (o *TrackingEventVoyageDetails) SetActualArrivalUtc(v time.Time) {
	o.ActualArrivalUtc = &v
}

// GetActualDepartureLt returns the ActualDepartureLt field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetActualDepartureLt() time.Time {
	if o == nil || o.ActualDepartureLt == nil {
		var ret time.Time
		return ret
	}
	return *o.ActualDepartureLt
}

// GetActualDepartureLtOk returns a tuple with the ActualDepartureLt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetActualDepartureLtOk() (*time.Time, bool) {
	if o == nil || o.ActualDepartureLt == nil {
		return nil, false
	}
	return o.ActualDepartureLt, true
}

// HasActualDepartureLt returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasActualDepartureLt() bool {
	if o != nil && o.ActualDepartureLt != nil {
		return true
	}

	return false
}

// SetActualDepartureLt gets a reference to the given time.Time and assigns it to the ActualDepartureLt field.
func (o *TrackingEventVoyageDetails) SetActualDepartureLt(v time.Time) {
	o.ActualDepartureLt = &v
}

// GetActualDepartureUtc returns the ActualDepartureUtc field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetActualDepartureUtc() time.Time {
	if o == nil || o.ActualDepartureUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.ActualDepartureUtc
}

// GetActualDepartureUtcOk returns a tuple with the ActualDepartureUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetActualDepartureUtcOk() (*time.Time, bool) {
	if o == nil || o.ActualDepartureUtc == nil {
		return nil, false
	}
	return o.ActualDepartureUtc, true
}

// HasActualDepartureUtc returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasActualDepartureUtc() bool {
	if o != nil && o.ActualDepartureUtc != nil {
		return true
	}

	return false
}

// SetActualDepartureUtc gets a reference to the given time.Time and assigns it to the ActualDepartureUtc field.
func (o *TrackingEventVoyageDetails) SetActualDepartureUtc(v time.Time) {
	o.ActualDepartureUtc = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TrackingEventVoyageDetails) SetId(v string) {
	o.Id = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *TrackingEventVoyageDetails) SetIndex(v int32) {
	o.Index = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetLat() float32 {
	if o == nil || o.Lat == nil {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetLatOk() (*float32, bool) {
	if o == nil || o.Lat == nil {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasLat() bool {
	if o != nil && o.Lat != nil {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *TrackingEventVoyageDetails) SetLat(v float32) {
	o.Lat = &v
}

// GetLon returns the Lon field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetLon() float32 {
	if o == nil || o.Lon == nil {
		var ret float32
		return ret
	}
	return *o.Lon
}

// GetLonOk returns a tuple with the Lon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetLonOk() (*float32, bool) {
	if o == nil || o.Lon == nil {
		return nil, false
	}
	return o.Lon, true
}

// HasLon returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasLon() bool {
	if o != nil && o.Lon != nil {
		return true
	}

	return false
}

// SetLon gets a reference to the given float32 and assigns it to the Lon field.
func (o *TrackingEventVoyageDetails) SetLon(v float32) {
	o.Lon = &v
}

// GetOriginalVoyageNo returns the OriginalVoyageNo field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetOriginalVoyageNo() string {
	if o == nil || o.OriginalVoyageNo == nil {
		var ret string
		return ret
	}
	return *o.OriginalVoyageNo
}

// GetOriginalVoyageNoOk returns a tuple with the OriginalVoyageNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetOriginalVoyageNoOk() (*string, bool) {
	if o == nil || o.OriginalVoyageNo == nil {
		return nil, false
	}
	return o.OriginalVoyageNo, true
}

// HasOriginalVoyageNo returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasOriginalVoyageNo() bool {
	if o != nil && o.OriginalVoyageNo != nil {
		return true
	}

	return false
}

// SetOriginalVoyageNo gets a reference to the given string and assigns it to the OriginalVoyageNo field.
func (o *TrackingEventVoyageDetails) SetOriginalVoyageNo(v string) {
	o.OriginalVoyageNo = &v
}

// GetPortCode returns the PortCode field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetPortCode() string {
	if o == nil || o.PortCode == nil {
		var ret string
		return ret
	}
	return *o.PortCode
}

// GetPortCodeOk returns a tuple with the PortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetPortCodeOk() (*string, bool) {
	if o == nil || o.PortCode == nil {
		return nil, false
	}
	return o.PortCode, true
}

// HasPortCode returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasPortCode() bool {
	if o != nil && o.PortCode != nil {
		return true
	}

	return false
}

// SetPortCode gets a reference to the given string and assigns it to the PortCode field.
func (o *TrackingEventVoyageDetails) SetPortCode(v string) {
	o.PortCode = &v
}

// GetPortName returns the PortName field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetPortName() string {
	if o == nil || o.PortName == nil {
		var ret string
		return ret
	}
	return *o.PortName
}

// GetPortNameOk returns a tuple with the PortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetPortNameOk() (*string, bool) {
	if o == nil || o.PortName == nil {
		return nil, false
	}
	return o.PortName, true
}

// HasPortName returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasPortName() bool {
	if o != nil && o.PortName != nil {
		return true
	}

	return false
}

// SetPortName gets a reference to the given string and assigns it to the PortName field.
func (o *TrackingEventVoyageDetails) SetPortName(v string) {
	o.PortName = &v
}

// GetPredictedArrivalLt returns the PredictedArrivalLt field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetPredictedArrivalLt() string {
	if o == nil || o.PredictedArrivalLt == nil {
		var ret string
		return ret
	}
	return *o.PredictedArrivalLt
}

// GetPredictedArrivalLtOk returns a tuple with the PredictedArrivalLt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetPredictedArrivalLtOk() (*string, bool) {
	if o == nil || o.PredictedArrivalLt == nil {
		return nil, false
	}
	return o.PredictedArrivalLt, true
}

// HasPredictedArrivalLt returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasPredictedArrivalLt() bool {
	if o != nil && o.PredictedArrivalLt != nil {
		return true
	}

	return false
}

// SetPredictedArrivalLt gets a reference to the given string and assigns it to the PredictedArrivalLt field.
func (o *TrackingEventVoyageDetails) SetPredictedArrivalLt(v string) {
	o.PredictedArrivalLt = &v
}

// GetPredictedArrivalUtc returns the PredictedArrivalUtc field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetPredictedArrivalUtc() string {
	if o == nil || o.PredictedArrivalUtc == nil {
		var ret string
		return ret
	}
	return *o.PredictedArrivalUtc
}

// GetPredictedArrivalUtcOk returns a tuple with the PredictedArrivalUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetPredictedArrivalUtcOk() (*string, bool) {
	if o == nil || o.PredictedArrivalUtc == nil {
		return nil, false
	}
	return o.PredictedArrivalUtc, true
}

// HasPredictedArrivalUtc returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasPredictedArrivalUtc() bool {
	if o != nil && o.PredictedArrivalUtc != nil {
		return true
	}

	return false
}

// SetPredictedArrivalUtc gets a reference to the given string and assigns it to the PredictedArrivalUtc field.
func (o *TrackingEventVoyageDetails) SetPredictedArrivalUtc(v string) {
	o.PredictedArrivalUtc = &v
}

// GetPredictedDepartureLt returns the PredictedDepartureLt field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetPredictedDepartureLt() string {
	if o == nil || o.PredictedDepartureLt == nil {
		var ret string
		return ret
	}
	return *o.PredictedDepartureLt
}

// GetPredictedDepartureLtOk returns a tuple with the PredictedDepartureLt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetPredictedDepartureLtOk() (*string, bool) {
	if o == nil || o.PredictedDepartureLt == nil {
		return nil, false
	}
	return o.PredictedDepartureLt, true
}

// HasPredictedDepartureLt returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasPredictedDepartureLt() bool {
	if o != nil && o.PredictedDepartureLt != nil {
		return true
	}

	return false
}

// SetPredictedDepartureLt gets a reference to the given string and assigns it to the PredictedDepartureLt field.
func (o *TrackingEventVoyageDetails) SetPredictedDepartureLt(v string) {
	o.PredictedDepartureLt = &v
}

// GetPredictedDepartureUtc returns the PredictedDepartureUtc field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetPredictedDepartureUtc() string {
	if o == nil || o.PredictedDepartureUtc == nil {
		var ret string
		return ret
	}
	return *o.PredictedDepartureUtc
}

// GetPredictedDepartureUtcOk returns a tuple with the PredictedDepartureUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetPredictedDepartureUtcOk() (*string, bool) {
	if o == nil || o.PredictedDepartureUtc == nil {
		return nil, false
	}
	return o.PredictedDepartureUtc, true
}

// HasPredictedDepartureUtc returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasPredictedDepartureUtc() bool {
	if o != nil && o.PredictedDepartureUtc != nil {
		return true
	}

	return false
}

// SetPredictedDepartureUtc gets a reference to the given string and assigns it to the PredictedDepartureUtc field.
func (o *TrackingEventVoyageDetails) SetPredictedDepartureUtc(v string) {
	o.PredictedDepartureUtc = &v
}

// GetPredictionTimeUtc returns the PredictionTimeUtc field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetPredictionTimeUtc() string {
	if o == nil || o.PredictionTimeUtc == nil {
		var ret string
		return ret
	}
	return *o.PredictionTimeUtc
}

// GetPredictionTimeUtcOk returns a tuple with the PredictionTimeUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetPredictionTimeUtcOk() (*string, bool) {
	if o == nil || o.PredictionTimeUtc == nil {
		return nil, false
	}
	return o.PredictionTimeUtc, true
}

// HasPredictionTimeUtc returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasPredictionTimeUtc() bool {
	if o != nil && o.PredictionTimeUtc != nil {
		return true
	}

	return false
}

// SetPredictionTimeUtc gets a reference to the given string and assigns it to the PredictionTimeUtc field.
func (o *TrackingEventVoyageDetails) SetPredictionTimeUtc(v string) {
	o.PredictionTimeUtc = &v
}

// GetScheduledArrivalLt returns the ScheduledArrivalLt field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetScheduledArrivalLt() time.Time {
	if o == nil || o.ScheduledArrivalLt == nil {
		var ret time.Time
		return ret
	}
	return *o.ScheduledArrivalLt
}

// GetScheduledArrivalLtOk returns a tuple with the ScheduledArrivalLt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetScheduledArrivalLtOk() (*time.Time, bool) {
	if o == nil || o.ScheduledArrivalLt == nil {
		return nil, false
	}
	return o.ScheduledArrivalLt, true
}

// HasScheduledArrivalLt returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasScheduledArrivalLt() bool {
	if o != nil && o.ScheduledArrivalLt != nil {
		return true
	}

	return false
}

// SetScheduledArrivalLt gets a reference to the given time.Time and assigns it to the ScheduledArrivalLt field.
func (o *TrackingEventVoyageDetails) SetScheduledArrivalLt(v time.Time) {
	o.ScheduledArrivalLt = &v
}

// GetScheduledArrivalOnTimeConfidence returns the ScheduledArrivalOnTimeConfidence field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetScheduledArrivalOnTimeConfidence() string {
	if o == nil || o.ScheduledArrivalOnTimeConfidence == nil {
		var ret string
		return ret
	}
	return *o.ScheduledArrivalOnTimeConfidence
}

// GetScheduledArrivalOnTimeConfidenceOk returns a tuple with the ScheduledArrivalOnTimeConfidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetScheduledArrivalOnTimeConfidenceOk() (*string, bool) {
	if o == nil || o.ScheduledArrivalOnTimeConfidence == nil {
		return nil, false
	}
	return o.ScheduledArrivalOnTimeConfidence, true
}

// HasScheduledArrivalOnTimeConfidence returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasScheduledArrivalOnTimeConfidence() bool {
	if o != nil && o.ScheduledArrivalOnTimeConfidence != nil {
		return true
	}

	return false
}

// SetScheduledArrivalOnTimeConfidence gets a reference to the given string and assigns it to the ScheduledArrivalOnTimeConfidence field.
func (o *TrackingEventVoyageDetails) SetScheduledArrivalOnTimeConfidence(v string) {
	o.ScheduledArrivalOnTimeConfidence = &v
}

// GetScheduledArrivalUtc returns the ScheduledArrivalUtc field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetScheduledArrivalUtc() time.Time {
	if o == nil || o.ScheduledArrivalUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.ScheduledArrivalUtc
}

// GetScheduledArrivalUtcOk returns a tuple with the ScheduledArrivalUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetScheduledArrivalUtcOk() (*time.Time, bool) {
	if o == nil || o.ScheduledArrivalUtc == nil {
		return nil, false
	}
	return o.ScheduledArrivalUtc, true
}

// HasScheduledArrivalUtc returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasScheduledArrivalUtc() bool {
	if o != nil && o.ScheduledArrivalUtc != nil {
		return true
	}

	return false
}

// SetScheduledArrivalUtc gets a reference to the given time.Time and assigns it to the ScheduledArrivalUtc field.
func (o *TrackingEventVoyageDetails) SetScheduledArrivalUtc(v time.Time) {
	o.ScheduledArrivalUtc = &v
}

// GetScheduledDepartureLt returns the ScheduledDepartureLt field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetScheduledDepartureLt() time.Time {
	if o == nil || o.ScheduledDepartureLt == nil {
		var ret time.Time
		return ret
	}
	return *o.ScheduledDepartureLt
}

// GetScheduledDepartureLtOk returns a tuple with the ScheduledDepartureLt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetScheduledDepartureLtOk() (*time.Time, bool) {
	if o == nil || o.ScheduledDepartureLt == nil {
		return nil, false
	}
	return o.ScheduledDepartureLt, true
}

// HasScheduledDepartureLt returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasScheduledDepartureLt() bool {
	if o != nil && o.ScheduledDepartureLt != nil {
		return true
	}

	return false
}

// SetScheduledDepartureLt gets a reference to the given time.Time and assigns it to the ScheduledDepartureLt field.
func (o *TrackingEventVoyageDetails) SetScheduledDepartureLt(v time.Time) {
	o.ScheduledDepartureLt = &v
}

// GetScheduledDepartureOnTimeConfidence returns the ScheduledDepartureOnTimeConfidence field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetScheduledDepartureOnTimeConfidence() string {
	if o == nil || o.ScheduledDepartureOnTimeConfidence == nil {
		var ret string
		return ret
	}
	return *o.ScheduledDepartureOnTimeConfidence
}

// GetScheduledDepartureOnTimeConfidenceOk returns a tuple with the ScheduledDepartureOnTimeConfidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetScheduledDepartureOnTimeConfidenceOk() (*string, bool) {
	if o == nil || o.ScheduledDepartureOnTimeConfidence == nil {
		return nil, false
	}
	return o.ScheduledDepartureOnTimeConfidence, true
}

// HasScheduledDepartureOnTimeConfidence returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasScheduledDepartureOnTimeConfidence() bool {
	if o != nil && o.ScheduledDepartureOnTimeConfidence != nil {
		return true
	}

	return false
}

// SetScheduledDepartureOnTimeConfidence gets a reference to the given string and assigns it to the ScheduledDepartureOnTimeConfidence field.
func (o *TrackingEventVoyageDetails) SetScheduledDepartureOnTimeConfidence(v string) {
	o.ScheduledDepartureOnTimeConfidence = &v
}

// GetScheduledDepartureUtc returns the ScheduledDepartureUtc field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetScheduledDepartureUtc() time.Time {
	if o == nil || o.ScheduledDepartureUtc == nil {
		var ret time.Time
		return ret
	}
	return *o.ScheduledDepartureUtc
}

// GetScheduledDepartureUtcOk returns a tuple with the ScheduledDepartureUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetScheduledDepartureUtcOk() (*time.Time, bool) {
	if o == nil || o.ScheduledDepartureUtc == nil {
		return nil, false
	}
	return o.ScheduledDepartureUtc, true
}

// HasScheduledDepartureUtc returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasScheduledDepartureUtc() bool {
	if o != nil && o.ScheduledDepartureUtc != nil {
		return true
	}

	return false
}

// SetScheduledDepartureUtc gets a reference to the given time.Time and assigns it to the ScheduledDepartureUtc field.
func (o *TrackingEventVoyageDetails) SetScheduledDepartureUtc(v time.Time) {
	o.ScheduledDepartureUtc = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *TrackingEventVoyageDetails) SetTimezone(v string) {
	o.Timezone = &v
}

// GetVoyageNoList returns the VoyageNoList field value if set, zero value otherwise.
func (o *TrackingEventVoyageDetails) GetVoyageNoList() []string {
	if o == nil || o.VoyageNoList == nil {
		var ret []string
		return ret
	}
	return *o.VoyageNoList
}

// GetVoyageNoListOk returns a tuple with the VoyageNoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingEventVoyageDetails) GetVoyageNoListOk() (*[]string, bool) {
	if o == nil || o.VoyageNoList == nil {
		return nil, false
	}
	return o.VoyageNoList, true
}

// HasVoyageNoList returns a boolean if a field has been set.
func (o *TrackingEventVoyageDetails) HasVoyageNoList() bool {
	if o != nil && o.VoyageNoList != nil {
		return true
	}

	return false
}

// SetVoyageNoList gets a reference to the given []string and assigns it to the VoyageNoList field.
func (o *TrackingEventVoyageDetails) SetVoyageNoList(v []string) {
	o.VoyageNoList = &v
}

func (o TrackingEventVoyageDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveScac != nil {
		toSerialize["active_scac"] = o.ActiveScac
	}
	if o.ActualArrivalLt != nil {
		toSerialize["actual_arrival_lt"] = o.ActualArrivalLt
	}
	if o.ActualArrivalUtc != nil {
		toSerialize["actual_arrival_utc"] = o.ActualArrivalUtc
	}
	if o.ActualDepartureLt != nil {
		toSerialize["actual_departure_lt"] = o.ActualDepartureLt
	}
	if o.ActualDepartureUtc != nil {
		toSerialize["actual_departure_utc"] = o.ActualDepartureUtc
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Lat != nil {
		toSerialize["lat"] = o.Lat
	}
	if o.Lon != nil {
		toSerialize["lon"] = o.Lon
	}
	if o.OriginalVoyageNo != nil {
		toSerialize["original_voyage_no"] = o.OriginalVoyageNo
	}
	if o.PortCode != nil {
		toSerialize["port_code"] = o.PortCode
	}
	if o.PortName != nil {
		toSerialize["port_name"] = o.PortName
	}
	if o.PredictedArrivalLt != nil {
		toSerialize["predicted_arrival_lt"] = o.PredictedArrivalLt
	}
	if o.PredictedArrivalUtc != nil {
		toSerialize["predicted_arrival_utc"] = o.PredictedArrivalUtc
	}
	if o.PredictedDepartureLt != nil {
		toSerialize["predicted_departure_lt"] = o.PredictedDepartureLt
	}
	if o.PredictedDepartureUtc != nil {
		toSerialize["predicted_departure_utc"] = o.PredictedDepartureUtc
	}
	if o.PredictionTimeUtc != nil {
		toSerialize["prediction_time_utc"] = o.PredictionTimeUtc
	}
	if o.ScheduledArrivalLt != nil {
		toSerialize["scheduled_arrival_lt"] = o.ScheduledArrivalLt
	}
	if o.ScheduledArrivalOnTimeConfidence != nil {
		toSerialize["scheduled_arrival_on_time_confidence"] = o.ScheduledArrivalOnTimeConfidence
	}
	if o.ScheduledArrivalUtc != nil {
		toSerialize["scheduled_arrival_utc"] = o.ScheduledArrivalUtc
	}
	if o.ScheduledDepartureLt != nil {
		toSerialize["scheduled_departure_lt"] = o.ScheduledDepartureLt
	}
	if o.ScheduledDepartureOnTimeConfidence != nil {
		toSerialize["scheduled_departure_on_time_confidence"] = o.ScheduledDepartureOnTimeConfidence
	}
	if o.ScheduledDepartureUtc != nil {
		toSerialize["scheduled_departure_utc"] = o.ScheduledDepartureUtc
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.VoyageNoList != nil {
		toSerialize["voyage_no_list"] = o.VoyageNoList
	}
	return json.Marshal(toSerialize)
}

type NullableTrackingEventVoyageDetails struct {
	value *TrackingEventVoyageDetails
	isSet bool
}

func (v NullableTrackingEventVoyageDetails) Get() *TrackingEventVoyageDetails {
	return v.value
}

func (v *NullableTrackingEventVoyageDetails) Set(val *TrackingEventVoyageDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingEventVoyageDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingEventVoyageDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingEventVoyageDetails(val *TrackingEventVoyageDetails) *NullableTrackingEventVoyageDetails {
	return &NullableTrackingEventVoyageDetails{value: val, isSet: true}
}

func (v NullableTrackingEventVoyageDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackingEventVoyageDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

