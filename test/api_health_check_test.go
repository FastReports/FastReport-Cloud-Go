/*
FastReport Cloud

Testing HealthCheckAPIService

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

func Test_gofrcloud_HealthCheckAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HealthCheckAPIService HealthCheckDataGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.HealthCheckAPI.HealthCheckDataGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
