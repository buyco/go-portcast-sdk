/*
Container Tracking API

Testing GetTrackingDataForUploadedContainersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package portcast

import (
	"context"
	openapiclient "github.com/buyco/go-portcast-sdk/portcast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_portcast_GetTrackingDataForUploadedContainersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GetTrackingDataForUploadedContainersAPIService DeleteTrackingBillOfLadingBookmarksBookmarkId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bookmarkId string

		resp, httpRes, err := apiClient.GetTrackingDataForUploadedContainersAPI.DeleteTrackingBillOfLadingBookmarksBookmarkId(context.Background(), bookmarkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GetTrackingDataForUploadedContainersAPIService GetApiV1EtaBillOfLadingBookmarks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bookmarkId string

		resp, httpRes, err := apiClient.GetTrackingDataForUploadedContainersAPI.GetApiV1EtaBillOfLadingBookmarks(context.Background(), bookmarkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GetTrackingDataForUploadedContainersAPIService ListApiV1EtaBillOfLadingBookmarks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GetTrackingDataForUploadedContainersAPI.ListApiV1EtaBillOfLadingBookmarks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}