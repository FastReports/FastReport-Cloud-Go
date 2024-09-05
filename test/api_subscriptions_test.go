/*
FastReport Cloud

Testing SubscriptionsAPIService

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

func Test_gofrcloud_SubscriptionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SubscriptionsAPIService SubscriptionsGetDefaultPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		resp, httpRes, err := apiClient.SubscriptionsAPI.SubscriptionsGetDefaultPermissions(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionsAPIService SubscriptionsGetMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subId string

		resp, httpRes, err := apiClient.SubscriptionsAPI.SubscriptionsGetMyPermissions(context.Background(), subId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionsAPIService SubscriptionsGetPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SubscriptionsAPI.SubscriptionsGetPermissions(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionsAPIService SubscriptionsGetSubscription", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SubscriptionsAPI.SubscriptionsGetSubscription(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionsAPIService SubscriptionsGetSubscriptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SubscriptionsAPI.SubscriptionsGetSubscriptions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionsAPIService SubscriptionsRenameSubscription", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		resp, httpRes, err := apiClient.SubscriptionsAPI.SubscriptionsRenameSubscription(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionsAPIService SubscriptionsUpdateDefaultPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		resp, httpRes, err := apiClient.SubscriptionsAPI.SubscriptionsUpdateDefaultPermissions(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionsAPIService SubscriptionsUpdateLocale", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		resp, httpRes, err := apiClient.SubscriptionsAPI.SubscriptionsUpdateLocale(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubscriptionsAPIService SubscriptionsUpdatePermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.SubscriptionsAPI.SubscriptionsUpdatePermissions(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
