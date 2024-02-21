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
	"strings"
)

type GetTrackingDataForUploadedContainersAPI interface {

	/*
		DeleteTrackingBillOfLadingBookmarksBookmarkId Archive Bookmark ID

		Archive a container Bookmark to stop tracking it.

	This will remove the container from the refresh processes and archive the shipment.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param bookmarkId Bill of Lading Bookmark ID
		@return ApiDeleteTrackingBillOfLadingBookmarksBookmarkIdRequest
	*/
	DeleteTrackingBillOfLadingBookmarksBookmarkId(ctx context.Context, bookmarkId string) ApiDeleteTrackingBillOfLadingBookmarksBookmarkIdRequest

	// DeleteTrackingBillOfLadingBookmarksBookmarkIdExecute executes the request
	//  @return DeleteTrackingBillOfLadingBookmarksBookmarkId200Response
	DeleteTrackingBillOfLadingBookmarksBookmarkIdExecute(r ApiDeleteTrackingBillOfLadingBookmarksBookmarkIdRequest) (*DeleteTrackingBillOfLadingBookmarksBookmarkId200Response, *http.Response, error)

	/*
		GetApiV1EtaBillOfLadingBookmarks Fetch tracking data by Bookmark ID

		Fetch the tracking data using the Bookmark ID.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param bookmarkId Bill of Lading Bookmark ID
		@return ApiGetApiV1EtaBillOfLadingBookmarksRequest
	*/
	GetApiV1EtaBillOfLadingBookmarks(ctx context.Context, bookmarkId string) ApiGetApiV1EtaBillOfLadingBookmarksRequest

	// GetApiV1EtaBillOfLadingBookmarksExecute executes the request
	//  @return GetApiV1EtaBillOfLadingBookmarks200Response
	GetApiV1EtaBillOfLadingBookmarksExecute(r ApiGetApiV1EtaBillOfLadingBookmarksRequest) (*GetApiV1EtaBillOfLadingBookmarks200Response, *http.Response, error)

	/*
		ListApiV1EtaBillOfLadingBookmarks List Tracking data by Container, Bill of Lading or Carrier

		Returns a list of your Tracking Data for the Bookmarks based on your query. The tracking requests are returned sorted by creation date, with the most recent tracking request appearing first.

	The list API doesn't return tracking data for more than 5 bookmarks, hence there's pagination which is applied on the same!

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiListApiV1EtaBillOfLadingBookmarksRequest
	*/
	ListApiV1EtaBillOfLadingBookmarks(ctx context.Context) ApiListApiV1EtaBillOfLadingBookmarksRequest

	// ListApiV1EtaBillOfLadingBookmarksExecute executes the request
	//  @return ListApiV1EtaBillOfLadingBookmarks200Response
	ListApiV1EtaBillOfLadingBookmarksExecute(r ApiListApiV1EtaBillOfLadingBookmarksRequest) (*ListApiV1EtaBillOfLadingBookmarks200Response, *http.Response, error)
}

// GetTrackingDataForUploadedContainersAPIService GetTrackingDataForUploadedContainersAPI service
type GetTrackingDataForUploadedContainersAPIService service

type ApiDeleteTrackingBillOfLadingBookmarksBookmarkIdRequest struct {
	ctx        context.Context
	ApiService GetTrackingDataForUploadedContainersAPI
	bookmarkId string
	xCustomer  *string
}

// [Enterprise Customers] Customer Name to indicate which org the bookmark belongs to
func (r ApiDeleteTrackingBillOfLadingBookmarksBookmarkIdRequest) XCustomer(xCustomer string) ApiDeleteTrackingBillOfLadingBookmarksBookmarkIdRequest {
	r.xCustomer = &xCustomer
	return r
}

func (r ApiDeleteTrackingBillOfLadingBookmarksBookmarkIdRequest) Execute() (*DeleteTrackingBillOfLadingBookmarksBookmarkId200Response, *http.Response, error) {
	return r.ApiService.DeleteTrackingBillOfLadingBookmarksBookmarkIdExecute(r)
}

/*
DeleteTrackingBillOfLadingBookmarksBookmarkId Archive Bookmark ID

Archive a container Bookmark to stop tracking it.

This will remove the container from the refresh processes and archive the shipment.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bookmarkId Bill of Lading Bookmark ID
	@return ApiDeleteTrackingBillOfLadingBookmarksBookmarkIdRequest
*/
func (a *GetTrackingDataForUploadedContainersAPIService) DeleteTrackingBillOfLadingBookmarksBookmarkId(ctx context.Context, bookmarkId string) ApiDeleteTrackingBillOfLadingBookmarksBookmarkIdRequest {
	return ApiDeleteTrackingBillOfLadingBookmarksBookmarkIdRequest{
		ApiService: a,
		ctx:        ctx,
		bookmarkId: bookmarkId,
	}
}

// Execute executes the request
//
//	@return DeleteTrackingBillOfLadingBookmarksBookmarkId200Response
func (a *GetTrackingDataForUploadedContainersAPIService) DeleteTrackingBillOfLadingBookmarksBookmarkIdExecute(r ApiDeleteTrackingBillOfLadingBookmarksBookmarkIdRequest) (*DeleteTrackingBillOfLadingBookmarksBookmarkId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeleteTrackingBillOfLadingBookmarksBookmarkId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GetTrackingDataForUploadedContainersAPIService.DeleteTrackingBillOfLadingBookmarksBookmarkId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eta/bill-of-lading-bookmarks/{bookmark_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"bookmark_id"+"}", url.PathEscape(parameterValueToString(r.bookmarkId, "bookmarkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
			var v GetApiV1EtaBillOfLadingBookmarks400Response
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v GetApiV1EtaBillOfLadingBookmarks404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiGetApiV1EtaBillOfLadingBookmarksRequest struct {
	ctx        context.Context
	ApiService GetTrackingDataForUploadedContainersAPI
	bookmarkId string
	xCustomer  *string
}

// [Enterprise Customers] Customer Name to indicate which org to get tracking data from
func (r ApiGetApiV1EtaBillOfLadingBookmarksRequest) XCustomer(xCustomer string) ApiGetApiV1EtaBillOfLadingBookmarksRequest {
	r.xCustomer = &xCustomer
	return r
}

func (r ApiGetApiV1EtaBillOfLadingBookmarksRequest) Execute() (*GetApiV1EtaBillOfLadingBookmarks200Response, *http.Response, error) {
	return r.ApiService.GetApiV1EtaBillOfLadingBookmarksExecute(r)
}

/*
GetApiV1EtaBillOfLadingBookmarks Fetch tracking data by Bookmark ID

Fetch the tracking data using the Bookmark ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bookmarkId Bill of Lading Bookmark ID
	@return ApiGetApiV1EtaBillOfLadingBookmarksRequest
*/
func (a *GetTrackingDataForUploadedContainersAPIService) GetApiV1EtaBillOfLadingBookmarks(ctx context.Context, bookmarkId string) ApiGetApiV1EtaBillOfLadingBookmarksRequest {
	return ApiGetApiV1EtaBillOfLadingBookmarksRequest{
		ApiService: a,
		ctx:        ctx,
		bookmarkId: bookmarkId,
	}
}

// Execute executes the request
//
//	@return GetApiV1EtaBillOfLadingBookmarks200Response
func (a *GetTrackingDataForUploadedContainersAPIService) GetApiV1EtaBillOfLadingBookmarksExecute(r ApiGetApiV1EtaBillOfLadingBookmarksRequest) (*GetApiV1EtaBillOfLadingBookmarks200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetApiV1EtaBillOfLadingBookmarks200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GetTrackingDataForUploadedContainersAPIService.GetApiV1EtaBillOfLadingBookmarks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eta/tracking/bill-of-lading-bookmarks/{bookmark_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"bookmark_id"+"}", url.PathEscape(parameterValueToString(r.bookmarkId, "bookmarkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
			var v GetApiV1EtaBillOfLadingBookmarks400Response
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v GetApiV1EtaBillOfLadingBookmarks404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiListApiV1EtaBillOfLadingBookmarksRequest struct {
	ctx        context.Context
	ApiService GetTrackingDataForUploadedContainersAPI
	carrierNo  *string
	blNo       *string
	cntrNo     *string
	startAfter *string
	xCustomer  *string
}

// Carrier SCAC Code
func (r ApiListApiV1EtaBillOfLadingBookmarksRequest) CarrierNo(carrierNo string) ApiListApiV1EtaBillOfLadingBookmarksRequest {
	r.carrierNo = &carrierNo
	return r
}

// Bill of Lading or Booking Number
func (r ApiListApiV1EtaBillOfLadingBookmarksRequest) BlNo(blNo string) ApiListApiV1EtaBillOfLadingBookmarksRequest {
	r.blNo = &blNo
	return r
}

// Container Number
func (r ApiListApiV1EtaBillOfLadingBookmarksRequest) CntrNo(cntrNo string) ApiListApiV1EtaBillOfLadingBookmarksRequest {
	r.cntrNo = &cntrNo
	return r
}

// Pagination Token
func (r ApiListApiV1EtaBillOfLadingBookmarksRequest) StartAfter(startAfter string) ApiListApiV1EtaBillOfLadingBookmarksRequest {
	r.startAfter = &startAfter
	return r
}

// [Enterprise Customers] Customer Name to indicate which org to get tracking data from
func (r ApiListApiV1EtaBillOfLadingBookmarksRequest) XCustomer(xCustomer string) ApiListApiV1EtaBillOfLadingBookmarksRequest {
	r.xCustomer = &xCustomer
	return r
}

func (r ApiListApiV1EtaBillOfLadingBookmarksRequest) Execute() (*ListApiV1EtaBillOfLadingBookmarks200Response, *http.Response, error) {
	return r.ApiService.ListApiV1EtaBillOfLadingBookmarksExecute(r)
}

/*
ListApiV1EtaBillOfLadingBookmarks List Tracking data by Container, Bill of Lading or Carrier

Returns a list of your Tracking Data for the Bookmarks based on your query. The tracking requests are returned sorted by creation date, with the most recent tracking request appearing first.

The list API doesn't return tracking data for more than 5 bookmarks, hence there's pagination which is applied on the same!

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListApiV1EtaBillOfLadingBookmarksRequest
*/
func (a *GetTrackingDataForUploadedContainersAPIService) ListApiV1EtaBillOfLadingBookmarks(ctx context.Context) ApiListApiV1EtaBillOfLadingBookmarksRequest {
	return ApiListApiV1EtaBillOfLadingBookmarksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListApiV1EtaBillOfLadingBookmarks200Response
func (a *GetTrackingDataForUploadedContainersAPIService) ListApiV1EtaBillOfLadingBookmarksExecute(r ApiListApiV1EtaBillOfLadingBookmarksRequest) (*ListApiV1EtaBillOfLadingBookmarks200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListApiV1EtaBillOfLadingBookmarks200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GetTrackingDataForUploadedContainersAPIService.ListApiV1EtaBillOfLadingBookmarks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eta/tracking/bill-of-lading-bookmarks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.carrierNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "carrier_no", r.carrierNo, "")
	}
	if r.blNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "bl_no", r.blNo, "")
	}
	if r.cntrNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cntr_no", r.cntrNo, "")
	}
	if r.startAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "_start_after", r.startAfter, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
			var v ListApiV1EtaBillOfLadingBookmarks400Response
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