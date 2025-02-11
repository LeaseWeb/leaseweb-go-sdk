/*
DNS

Testing DnsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package dns

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/dns"
)

func Test_dns_DnsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DnsAPIService CreateResourceRecordSet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainName string

		resp, httpRes, err := apiClient.DnsAPI.CreateResourceRecordSet(context.Background(), domainName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsAPIService DeleteResourceRecordSet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainName string
		var name string
		var type_ string

		httpRes, err := apiClient.DnsAPI.DeleteResourceRecordSet(context.Background(), domainName, name, type_).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsAPIService DeleteResourceRecordSets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainName string

		httpRes, err := apiClient.DnsAPI.DeleteResourceRecordSets(context.Background(), domainName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsAPIService ExportResourceRecordSets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainName string

		httpRes, err := apiClient.DnsAPI.ExportResourceRecordSets(context.Background(), domainName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsAPIService GetResourceRecordSet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainName string
		var name string
		var type_ string

		resp, httpRes, err := apiClient.DnsAPI.GetResourceRecordSet(context.Background(), domainName, name, type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsAPIService GetResourceRecordSetList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainName string

		resp, httpRes, err := apiClient.DnsAPI.GetResourceRecordSetList(context.Background(), domainName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsAPIService ImportResourceRecordSets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainName string

		resp, httpRes, err := apiClient.DnsAPI.ImportResourceRecordSets(context.Background(), domainName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsAPIService UpdateResourceRecordSet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainName string
		var name string
		var type_ string

		resp, httpRes, err := apiClient.DnsAPI.UpdateResourceRecordSet(context.Background(), domainName, name, type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsAPIService UpdateResourceRecordSets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainName string

		resp, httpRes, err := apiClient.DnsAPI.UpdateResourceRecordSets(context.Background(), domainName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsAPIService ValidateResourceRecordSet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainName string

		resp, httpRes, err := apiClient.DnsAPI.ValidateResourceRecordSet(context.Background(), domainName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsAPIService ValidateZone", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domainName string

		resp, httpRes, err := apiClient.DnsAPI.ValidateZone(context.Background(), domainName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
