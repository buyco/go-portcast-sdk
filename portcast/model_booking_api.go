/*
Container Tracking API

This documentation defines how the users can utilize the Portcast API to retrieve the latest Track and Trace Data for any container, bill of lading or booking across 100+ carriers and 2500+ ports across the globe.  What's changed in the version 2.0? Find out more about it [here](https://support.portcast.io/support/solutions/articles/151000074289-api-release-notes-april-2023-)!!

API version: 2.0
Contact: support@portcast.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package portcast

import (
	"encoding/json"
)

// checks if the BookingAPI type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookingAPI{}

// BookingAPI Booking API - POST Request Response
type BookingAPI struct {
	// List of Container Numbers
	Containers []string `json:"containers,omitempty"`
	// Status of the call made
	Status *string `json:"status,omitempty"`
	//   Key | Description  ---------|----------   SUCCESS | Successful call   BL_NOT_FOUND | Carrier didn't return any data for the queried document    INVALID_BL_NO | Carrier returned an error for the queried document   AUTO_CANT_DETECT_SEALINE | AUTO Scac detection wasn't able to get the correct carrier   UNEXPECTED_ERROR | There was an error while fetching the document details, please check input parameters
	Message *string `json:"message,omitempty"`
	// Bill of Lading Bookmarks Created List
	Response []BookingAPIResponseInner `json:"response,omitempty"`
	Route *BookingAPIRoute `json:"route,omitempty"`
	// List of Locations the booking will visit
	Locations []BookingAPILocationsInner `json:"locations,omitempty"`
	// List of Vessels the booking will operate on
	Vessels []BookingAPIVesselsInner `json:"vessels,omitempty"`
}

// NewBookingAPI instantiates a new BookingAPI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookingAPI() *BookingAPI {
	this := BookingAPI{}
	return &this
}

// NewBookingAPIWithDefaults instantiates a new BookingAPI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookingAPIWithDefaults() *BookingAPI {
	this := BookingAPI{}
	return &this
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *BookingAPI) GetContainers() []string {
	if o == nil || IsNil(o.Containers) {
		var ret []string
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingAPI) GetContainersOk() ([]string, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *BookingAPI) HasContainers() bool {
	if o != nil && !IsNil(o.Containers) {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []string and assigns it to the Containers field.
func (o *BookingAPI) SetContainers(v []string) {
	o.Containers = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BookingAPI) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingAPI) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BookingAPI) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BookingAPI) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BookingAPI) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingAPI) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BookingAPI) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BookingAPI) SetMessage(v string) {
	o.Message = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *BookingAPI) GetResponse() []BookingAPIResponseInner {
	if o == nil || IsNil(o.Response) {
		var ret []BookingAPIResponseInner
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingAPI) GetResponseOk() ([]BookingAPIResponseInner, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *BookingAPI) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given []BookingAPIResponseInner and assigns it to the Response field.
func (o *BookingAPI) SetResponse(v []BookingAPIResponseInner) {
	o.Response = v
}

// GetRoute returns the Route field value if set, zero value otherwise.
func (o *BookingAPI) GetRoute() BookingAPIRoute {
	if o == nil || IsNil(o.Route) {
		var ret BookingAPIRoute
		return ret
	}
	return *o.Route
}

// GetRouteOk returns a tuple with the Route field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingAPI) GetRouteOk() (*BookingAPIRoute, bool) {
	if o == nil || IsNil(o.Route) {
		return nil, false
	}
	return o.Route, true
}

// HasRoute returns a boolean if a field has been set.
func (o *BookingAPI) HasRoute() bool {
	if o != nil && !IsNil(o.Route) {
		return true
	}

	return false
}

// SetRoute gets a reference to the given BookingAPIRoute and assigns it to the Route field.
func (o *BookingAPI) SetRoute(v BookingAPIRoute) {
	o.Route = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *BookingAPI) GetLocations() []BookingAPILocationsInner {
	if o == nil || IsNil(o.Locations) {
		var ret []BookingAPILocationsInner
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingAPI) GetLocationsOk() ([]BookingAPILocationsInner, bool) {
	if o == nil || IsNil(o.Locations) {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *BookingAPI) HasLocations() bool {
	if o != nil && !IsNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []BookingAPILocationsInner and assigns it to the Locations field.
func (o *BookingAPI) SetLocations(v []BookingAPILocationsInner) {
	o.Locations = v
}

// GetVessels returns the Vessels field value if set, zero value otherwise.
func (o *BookingAPI) GetVessels() []BookingAPIVesselsInner {
	if o == nil || IsNil(o.Vessels) {
		var ret []BookingAPIVesselsInner
		return ret
	}
	return o.Vessels
}

// GetVesselsOk returns a tuple with the Vessels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookingAPI) GetVesselsOk() ([]BookingAPIVesselsInner, bool) {
	if o == nil || IsNil(o.Vessels) {
		return nil, false
	}
	return o.Vessels, true
}

// HasVessels returns a boolean if a field has been set.
func (o *BookingAPI) HasVessels() bool {
	if o != nil && !IsNil(o.Vessels) {
		return true
	}

	return false
}

// SetVessels gets a reference to the given []BookingAPIVesselsInner and assigns it to the Vessels field.
func (o *BookingAPI) SetVessels(v []BookingAPIVesselsInner) {
	o.Vessels = v
}

func (o BookingAPI) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookingAPI) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Containers) {
		toSerialize["containers"] = o.Containers
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	if !IsNil(o.Route) {
		toSerialize["route"] = o.Route
	}
	if !IsNil(o.Locations) {
		toSerialize["locations"] = o.Locations
	}
	if !IsNil(o.Vessels) {
		toSerialize["vessels"] = o.Vessels
	}
	return toSerialize, nil
}

type NullableBookingAPI struct {
	value *BookingAPI
	isSet bool
}

func (v NullableBookingAPI) Get() *BookingAPI {
	return v.value
}

func (v *NullableBookingAPI) Set(val *BookingAPI) {
	v.value = val
	v.isSet = true
}

func (v NullableBookingAPI) IsSet() bool {
	return v.isSet
}

func (v *NullableBookingAPI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookingAPI(val *BookingAPI) *NullableBookingAPI {
	return &NullableBookingAPI{value: val, isSet: true}
}

func (v NullableBookingAPI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookingAPI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


