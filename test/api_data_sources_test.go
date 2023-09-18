/*
FastReport Cloud

Testing DataSourcesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package gofrcloud

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/fastreports/gofrcloud"
)

func Test_gofrcloud_DataSourcesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DataSourcesApiService DataSourcesCountDataSourcesAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		resp, httpRes, err := apiClient.DataSourcesApi.DataSourcesCountDataSourcesAsync(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataSourcesApiService DataSourcesCreateDataSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DataSourcesApi.DataSourcesCreateDataSource(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataSourcesApiService DataSourcesDeleteDataSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.DataSourcesApi.DataSourcesDeleteDataSource(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataSourcesApiService DataSourcesFetchData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.DataSourcesApi.DataSourcesFetchData(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataSourcesApiService DataSourcesGetAvailableDataSources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DataSourcesApi.DataSourcesGetAvailableDataSources(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataSourcesApiService DataSourcesGetDataSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.DataSourcesApi.DataSourcesGetDataSource(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataSourcesApiService DataSourcesGetPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.DataSourcesApi.DataSourcesGetPermissions(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataSourcesApiService DataSourcesRenameDataSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.DataSourcesApi.DataSourcesRenameDataSource(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataSourcesApiService DataSourcesUpdateConnectionString", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.DataSourcesApi.DataSourcesUpdateConnectionString(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataSourcesApiService DataSourcesUpdatePermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.DataSourcesApi.DataSourcesUpdatePermissions(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataSourcesApiService DataSourcesUpdateSubscriptionDataSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.DataSourcesApi.DataSourcesUpdateSubscriptionDataSource(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
