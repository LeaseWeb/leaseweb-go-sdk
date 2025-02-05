/*
IP management

Testing IpmgmtAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ipmgmt

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ipmgmt"
)

func Test_ipmgmt_IpmgmtAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IpmgmtAPIService GetIPList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IpmgmtAPI.GetIPList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IpmgmtAPIService GetNullRouteHistoryList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IpmgmtAPI.GetNullRouteHistoryList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IpmgmtAPIService GetReverseLookupRecordList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ip string

		resp, httpRes, err := apiClient.IpmgmtAPI.GetReverseLookupRecordList(context.Background(), ip).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IpmgmtAPIService InspectIP", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ip string

		resp, httpRes, err := apiClient.IpmgmtAPI.InspectIP(context.Background(), ip).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IpmgmtAPIService InspectNullRouteHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.IpmgmtAPI.InspectNullRouteHistory(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IpmgmtAPIService NullRouteIP", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ip string

		resp, httpRes, err := apiClient.IpmgmtAPI.NullRouteIP(context.Background(), ip).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IpmgmtAPIService RemoveIPNullRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ip string

		httpRes, err := apiClient.IpmgmtAPI.RemoveIPNullRoute(context.Background(), ip).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IpmgmtAPIService UpdateIP", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ip string

		resp, httpRes, err := apiClient.IpmgmtAPI.UpdateIP(context.Background(), ip).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IpmgmtAPIService UpdateNullRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.IpmgmtAPI.UpdateNullRoute(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IpmgmtAPIService UpdateReverseLookupRecords", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ip string

		resp, httpRes, err := apiClient.IpmgmtAPI.UpdateReverseLookupRecords(context.Background(), ip).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
