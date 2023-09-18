/*
FastReport Cloud

Testing ExportsApiService

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

func Test_gofrcloud_ExportsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ExportsApiService ExportFolderAndFileClearRecycleBin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		httpRes, err := apiClient.ExportsApi.ExportFolderAndFileClearRecycleBin(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFolderAndFileDeleteFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		httpRes, err := apiClient.ExportsApi.ExportFolderAndFileDeleteFiles(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFolderAndFileGetCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportFolderAndFileGetCount(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFolderAndFileGetFoldersAndFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportFolderAndFileGetFoldersAndFiles(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFolderAndFileGetRecycleBinFoldersAndFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		resp, httpRes, err := apiClient.ExportsApi.ExportFolderAndFileGetRecycleBinFoldersAndFiles(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFolderAndFileRecoverAllFromRecycleBin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		httpRes, err := apiClient.ExportsApi.ExportFolderAndFileRecoverAllFromRecycleBin(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersCalculateFolderSize", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportFoldersCalculateFolderSize(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersCopyFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var folderId string

		resp, httpRes, err := apiClient.ExportsApi.ExportFoldersCopyFolder(context.Background(), id, folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersDeleteFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ExportsApi.ExportFoldersDeleteFolder(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersGetBreadcrumbs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportFoldersGetBreadcrumbs(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersGetFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportFoldersGetFolder(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersGetFolders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportFoldersGetFolders(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersGetFoldersCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportFoldersGetFoldersCount(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersGetOrCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ExportsApi.ExportFoldersGetOrCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersGetPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportFoldersGetPermissions(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersGetRootFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ExportsApi.ExportFoldersGetRootFolder(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersMoveFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var folderId string

		resp, httpRes, err := apiClient.ExportsApi.ExportFoldersMoveFolder(context.Background(), id, folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersMoveFolderToBin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ExportsApi.ExportFoldersMoveFolderToBin(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersPostFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportFoldersPostFolder(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersRecoverFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ExportsApi.ExportFoldersRecoverFolder(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersRenameFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportFoldersRenameFolder(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersUpdateIcon", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportFoldersUpdateIcon(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersUpdatePermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ExportsApi.ExportFoldersUpdatePermissions(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportFoldersUpdateTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportFoldersUpdateTags(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportsCopyFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var folderId string

		resp, httpRes, err := apiClient.ExportsApi.ExportsCopyFile(context.Background(), id, folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportsDeleteFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ExportsApi.ExportsDeleteFile(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportsGetFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportsGetFile(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportsGetFileHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportsGetFileHistory(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportsGetFilesCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportsGetFilesCount(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportsGetFilesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportsGetFilesList(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportsGetPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportsGetPermissions(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportsMoveFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var folderId string

		resp, httpRes, err := apiClient.ExportsApi.ExportsMoveFile(context.Background(), id, folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportsMoveFileToBin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ExportsApi.ExportsMoveFileToBin(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportsRecoverFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ExportsApi.ExportsRecoverFile(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportsRenameFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportsRenameFile(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportsUpdateIcon", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportsUpdateIcon(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportsUpdatePermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ExportsApi.ExportsUpdatePermissions(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportsApiService ExportsUpdateTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExportsApi.ExportsUpdateTags(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}