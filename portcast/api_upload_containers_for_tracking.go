/*
Container Tracking API

This documentation defines how the users can utilize the Portcast API to retrieve the latest Track and Trace Data for any container, bill of lading or booking across 100+ carriers and 2500+ ports across the globe.  What's changed in the version 2.0? Find out more about it [here](https://support.portcast.io/support/solutions/articles/151000074289-api-release-notes-april-2023-)!!

API version: 2.0
Contact: support@portcast.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package portcast

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


type UploadContainersForTrackingAPI interface {

	/*
	PostBooking Upload using Booking/Bill of Lading

	To initiate tracking any booking or bill of lading, when you do not have the container numbers, you can use this API to trigger an upload and start tracking the containers within that booking!

This API automatically fetches the container numbers which are a part of the Booking or Bill of Lading provided by you and start tracking the same.

When the containers are fetched, we upload them on our end and share back the [Bill of Lading Bookmark](../reference/PortcastAPI.json/components/schemas/Bill_of_Lading_Bookmark) object for your reference which can then be used to fetch tracking data for those containers.

To initiate an upload, you need to provide three inputs:
- Carrier SCAC Code: The carrier SCAC code which helps us identify which carrier to fetch the data from
- Document Number: Bill of Lading or Booking Number which helps us fetch the containers
- Document Type: This helps us identify if the doc no provided is a booking number or a bill of lading number 

Supported carriers can be found [here](https://docs.google.com/spreadsheets/d/1l7eA1brGaEZwhUS1_xwq1puyK35maPN0T-DWwNJvnhI).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostBookingRequest
	*/
	PostBooking(ctx context.Context) ApiPostBookingRequest

	// PostBookingExecute executes the request
	//  @return BookingAPI
	PostBookingExecute(r ApiPostBookingRequest) (*BookingAPI, *http.Response, error)

	/*
	PostEtaBillOfLadingBookmarks Upload using Container Number

	*To track shipment, you need to provide three inputs:*

- Carrier SCAC code: This helps us to identify which carrier to fetch the data from, use AUTO if you don't know. If AUTO is used the bill of lading number is disregarded, only container number is used and the data return is for the latest journey of the container.
- Bill of Lading or Booking Number: This helps us to fetch the correct journey along with detailed tracking data
- Container Number: This helps us to identify which container you intend to track within the Booking

Incase of any missing inputs, refer to [Input Data Guide](docs/Input_Data_Guide.md) for more details!

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostEtaBillOfLadingBookmarksRequest
	*/
	PostEtaBillOfLadingBookmarks(ctx context.Context) ApiPostEtaBillOfLadingBookmarksRequest

	// PostEtaBillOfLadingBookmarksExecute executes the request
	//  @return PostEtaBillOfLadingBookmarks200Response
	PostEtaBillOfLadingBookmarksExecute(r ApiPostEtaBillOfLadingBookmarksRequest) (*PostEtaBillOfLadingBookmarks200Response, *http.Response, error)
}

// UploadContainersForTrackingAPIService UploadContainersForTrackingAPI service
type UploadContainersForTrackingAPIService service

type ApiPostBookingRequest struct {
	ctx context.Context
	ApiService UploadContainersForTrackingAPI
	xCustomer *string
	postBookingRequest *PostBookingRequest
}

// [Enterprise Customers] Customer Name to indicate which org to upload too
func (r ApiPostBookingRequest) XCustomer(xCustomer string) ApiPostBookingRequest {
	r.xCustomer = &xCustomer
	return r
}

// Upload Parameters
func (r ApiPostBookingRequest) PostBookingRequest(postBookingRequest PostBookingRequest) ApiPostBookingRequest {
	r.postBookingRequest = &postBookingRequest
	return r
}

func (r ApiPostBookingRequest) Execute() (*BookingAPI, *http.Response, error) {
	return r.ApiService.PostBookingExecute(r)
}

/*
PostBooking Upload using Booking/Bill of Lading

To initiate tracking any booking or bill of lading, when you do not have the container numbers, you can use this API to trigger an upload and start tracking the containers within that booking!

This API automatically fetches the container numbers which are a part of the Booking or Bill of Lading provided by you and start tracking the same.

When the containers are fetched, we upload them on our end and share back the [Bill of Lading Bookmark](../reference/PortcastAPI.json/components/schemas/Bill_of_Lading_Bookmark) object for your reference which can then be used to fetch tracking data for those containers.

To initiate an upload, you need to provide three inputs:
- Carrier SCAC Code: The carrier SCAC code which helps us identify which carrier to fetch the data from
- Document Number: Bill of Lading or Booking Number which helps us fetch the containers
- Document Type: This helps us identify if the doc no provided is a booking number or a bill of lading number 

Supported carriers can be found [here](https://docs.google.com/spreadsheets/d/1l7eA1brGaEZwhUS1_xwq1puyK35maPN0T-DWwNJvnhI).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostBookingRequest
*/
func (a *UploadContainersForTrackingAPIService) PostBooking(ctx context.Context) ApiPostBookingRequest {
	return ApiPostBookingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BookingAPI
func (a *UploadContainersForTrackingAPIService) PostBookingExecute(r ApiPostBookingRequest) (*BookingAPI, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BookingAPI
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UploadContainersForTrackingAPIService.PostBooking")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eta/booking"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xCustomer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-customer", r.xCustomer, "")
	}
	// body params
	localVarPostBody = r.postBookingRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x-api-key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v PostBooking400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PostEtaBillOfLadingBookmarks403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v PostBooking422Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostEtaBillOfLadingBookmarksRequest struct {
	ctx context.Context
	ApiService UploadContainersForTrackingAPI
	xCustomer *string
	postEtaBillOfLadingBookmarksRequest *PostEtaBillOfLadingBookmarksRequest
}

// [Enterprise Customers] Customer Name to indicate which org to upload too
func (r ApiPostEtaBillOfLadingBookmarksRequest) XCustomer(xCustomer string) ApiPostEtaBillOfLadingBookmarksRequest {
	r.xCustomer = &xCustomer
	return r
}

// Upload Parameters
func (r ApiPostEtaBillOfLadingBookmarksRequest) PostEtaBillOfLadingBookmarksRequest(postEtaBillOfLadingBookmarksRequest PostEtaBillOfLadingBookmarksRequest) ApiPostEtaBillOfLadingBookmarksRequest {
	r.postEtaBillOfLadingBookmarksRequest = &postEtaBillOfLadingBookmarksRequest
	return r
}

func (r ApiPostEtaBillOfLadingBookmarksRequest) Execute() (*PostEtaBillOfLadingBookmarks200Response, *http.Response, error) {
	return r.ApiService.PostEtaBillOfLadingBookmarksExecute(r)
}

/*
PostEtaBillOfLadingBookmarks Upload using Container Number

*To track shipment, you need to provide three inputs:*

- Carrier SCAC code: This helps us to identify which carrier to fetch the data from, use AUTO if you don't know. If AUTO is used the bill of lading number is disregarded, only container number is used and the data return is for the latest journey of the container.
- Bill of Lading or Booking Number: This helps us to fetch the correct journey along with detailed tracking data
- Container Number: This helps us to identify which container you intend to track within the Booking

Incase of any missing inputs, refer to [Input Data Guide](docs/Input_Data_Guide.md) for more details!

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostEtaBillOfLadingBookmarksRequest
*/
func (a *UploadContainersForTrackingAPIService) PostEtaBillOfLadingBookmarks(ctx context.Context) ApiPostEtaBillOfLadingBookmarksRequest {
	return ApiPostEtaBillOfLadingBookmarksRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostEtaBillOfLadingBookmarks200Response
func (a *UploadContainersForTrackingAPIService) PostEtaBillOfLadingBookmarksExecute(r ApiPostEtaBillOfLadingBookmarksRequest) (*PostEtaBillOfLadingBookmarks200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostEtaBillOfLadingBookmarks200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UploadContainersForTrackingAPIService.PostEtaBillOfLadingBookmarks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eta/bill-of-lading-bookmarks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xCustomer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-customer", r.xCustomer, "")
	}
	// body params
	localVarPostBody = r.postEtaBillOfLadingBookmarksRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x-api-key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v PostEtaBillOfLadingBookmarks400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v PostEtaBillOfLadingBookmarks401Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v PostEtaBillOfLadingBookmarks403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v PostEtaBillOfLadingBookmarks409Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
