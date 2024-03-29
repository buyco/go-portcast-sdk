/*
Portcast API (1.1.0) - Bill of Lading Tracking

**This documentation is for the latest version of the Portcast Bill of Lading Tracking API.**  There are two variables used in this documentation: 1. `api-url`: the url to use for accessing the API. The official url is `https://api.portcast.io` 2. `x-api-key`: the access token to send along with every request to the API. This key will be provided to each organisation upon API access activation  The general workflow is as below:  1. Create the bill of lading bookmark if it does not exist already (`POST {{api-url}}/api/v1/eta/bill-of-lading-bookmarks`). 2. A bookmark must contain `carrier_no`, `bl_no` and `cntr_no` information. This will return the bill of lading bookmark information created. Record the `id` field from the response. 3. Wait for predictions to be generated. This could take up to 5 mins. 5. Query for the tracking results based on the `id` recorded earlier (`GET {{api-url}}/api/v1/eta/tracking/bill-of-lading-bookmarks/<id>`)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package portcast

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

type BookingBookmarksApi interface {

	/*
			EtaBookingPost Create new bookmarks for booking

			Please send the bill of lading or booking number with the following fields:
		```
		{
		    "carrier_no": <string, required, scac code of carrier, will use first 4 characters of bl_no as carrier_no if left empty>,
		    "bl_no": <string, required, ocean carrier bill of lading number>,
		    "dock_type": <string, required, type of doc_no. Accept value: BK (booking) , BL (Bill of Lading)>
		}
		```

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @return ApiEtaBookingPostRequest
	*/
	EtaBookingPost(ctx _context.Context) ApiEtaBookingPostRequest

	// EtaBookingPostExecute executes the request
	//  @return BookingBookmarksResponse
	EtaBookingPostExecute(r ApiEtaBookingPostRequest) (BookingBookmarksResponse, *_nethttp.Response, error)
}

// BookingBookmarksApiService BookingBookmarksApi service
type BookingBookmarksApiService service

type ApiEtaBookingPostRequest struct {
	ctx                    _context.Context
	ApiService             BookingBookmarksApi
	bookingBookmarkRequest *BookingBookmarkRequest
}

func (r ApiEtaBookingPostRequest) BookingBookmarkRequest(bookingBookmarkRequest BookingBookmarkRequest) ApiEtaBookingPostRequest {
	r.bookingBookmarkRequest = &bookingBookmarkRequest
	return r
}

func (r ApiEtaBookingPostRequest) Execute() (BookingBookmarksResponse, *_nethttp.Response, error) {
	return r.ApiService.EtaBookingPostExecute(r)
}

/*
EtaBookingPost Create new bookmarks for booking

Please send the bill of lading or booking number with the following fields:
```

	{
	    "carrier_no": <string, required, scac code of carrier, will use first 4 characters of bl_no as carrier_no if left empty>,
	    "bl_no": <string, required, ocean carrier bill of lading number>,
	    "dock_type": <string, required, type of doc_no. Accept value: BK (booking) , BL (Bill of Lading)>
	}

```

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiEtaBookingPostRequest
*/
func (a *BookingBookmarksApiService) EtaBookingPost(ctx _context.Context) ApiEtaBookingPostRequest {
	return ApiEtaBookingPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BookingBookmarksResponse
func (a *BookingBookmarksApiService) EtaBookingPostExecute(r ApiEtaBookingPostRequest) (BookingBookmarksResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BookingBookmarksResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BookingBookmarksApiService.EtaBookingPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eta/booking"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.bookingBookmarkRequest == nil {
		return localVarReturnValue, nil, reportError("bookingBookmarkRequest is required and must be specified")
	}

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
	// body params
	localVarPostBody = r.bookingBookmarkRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["authorization"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["authorization"] = key
			}
		}
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x-customer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-customer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v MessageModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v MessageModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
