/*
FastReport Cloud

Testing TasksAPIService

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

func Test_gofrcloud_TasksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TasksAPIService TasksCreateTask", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TasksAPI.TasksCreateTask(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksDeleteTask", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskId string

		httpRes, err := apiClient.TasksAPI.TasksDeleteTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TasksAPI.TasksGet(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksGetList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TasksAPI.TasksGetList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksGetPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TasksAPI.TasksGetPermissions(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksRenameTask", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TasksAPI.TasksRenameTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksRunTask", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.TasksAPI.TasksRunTask(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksRunTaskById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskId string

		httpRes, err := apiClient.TasksAPI.TasksRunTaskById(context.Background(), taskId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksUpdatePermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.TasksAPI.TasksUpdatePermissions(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService TasksUpdateTask", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TasksAPI.TasksUpdateTask(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
