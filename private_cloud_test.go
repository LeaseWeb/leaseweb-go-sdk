package leaseweb

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListPrivateClouds(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 1}, "privateClouds": [
			{
				"id": "218030",
				"customerId": "1301178860",
				"dataCenter": "AMS-01",
				"serviceOffering": "FLAT_FEE",
				"sla": "Bronze",
				"contract": {
					"id": "30000775",
					"startsAt": "2015-11-01T00:00:00+02:00",
					"endsAt": "2016-12-30T10:39:27+01:00",
					"billingCycle": 12,
					"billingFrequency": "MONTH",
					"pricePerFrequency": 0,
					"currency": "EUR"
				},
				"hardware": {
					"cpu": {
						"cores": 25
					},
					"memory": {
						"unit": "GB",
						"amount": 50
					},
					"storage": {
						"unit": "GB",
						"amount": 1
					}
				},
				"ips": [
					{
						"ip": "10.12.144.32",
						"version": 4,
						"type": "PUBLIC"
					}
				],
				"networkTraffic": {
					"type": "DATATRAFFIC",
					"trafficType": "PREMIUM",
					"datatrafficUnit": "TB",
					"datatrafficLimit": 6
				}
			}
		]}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.ListPrivateClouds()

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 1)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.PrivateClouds), 1)

	privateCloud1 := response.PrivateClouds[0]
	assert.Equal(privateCloud1.Id, "218030")
	assert.Equal(privateCloud1.CustomerId, "1301178860")
	assert.Equal(privateCloud1.DataCenter, "AMS-01")
	assert.Equal(privateCloud1.ServiceOffering, "FLAT_FEE")
	assert.Equal(privateCloud1.Sla, "Bronze")
	assert.Equal(privateCloud1.Contract.Id, "30000775")
	assert.Equal(privateCloud1.Contract.StartsAt, "2015-11-01T00:00:00+02:00")
	assert.Equal(privateCloud1.Contract.EndsAt, "2016-12-30T10:39:27+01:00")
	assert.Equal(privateCloud1.Contract.BillingCycle, 12)
	assert.Equal(privateCloud1.Contract.BillingFrequency, "MONTH")
	assert.Equal(privateCloud1.Contract.PricePerFrequency, float32(0))
	assert.Equal(privateCloud1.Contract.Currency, "EUR")
	assert.Equal(privateCloud1.Hardware.Cpu.Cores, 25)
	assert.Equal(privateCloud1.Hardware.Memory.Amount, 50)
	assert.Equal(privateCloud1.Hardware.Memory.Unit, "GB")
	assert.Equal(privateCloud1.Hardware.Storage.Amount, 1)
	assert.Equal(privateCloud1.Hardware.Storage.Unit, "GB")
	assert.Equal(privateCloud1.Ips[0].Ip, "10.12.144.32")
	assert.Equal(privateCloud1.Ips[0].Type, "PUBLIC")
	assert.Equal(privateCloud1.Ips[0].Version, 4)
	assert.Equal(privateCloud1.NetworkTraffic.DataTrafficLimit, 6)
	assert.Equal(privateCloud1.NetworkTraffic.DataTrafficUnit, "TB")
	assert.Equal(privateCloud1.NetworkTraffic.TrafficType, "PREMIUM")
	assert.Equal(privateCloud1.NetworkTraffic.Type, "DATATRAFFIC")
}

func TestListPrivateCloudsBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "privateClouds": []}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.ListPrivateClouds()

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.PrivateClouds), 0)
}

func TestListPrivateCloudsPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{"_metadata":{"limit": 20, "offset": 10, "totalCount": 2}, "privateClouds": []}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.ListPrivateClouds(10, 20)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 2)
	assert.Equal(response.Metadata.Offset, 10)
	assert.Equal(response.Metadata.Limit, 20)
	assert.Equal(len(response.PrivateClouds), 0)
}

func TestListPrivateCloudsError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.ListPrivateClouds()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "ACCESS_DENIED")
}

func TestListPrivateCloudsError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.ListPrivateClouds()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "SERVER_ERROR")
}

func TestListPrivateCloudsError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	invoiceApi := InvoiceApi{}
	resp, err := invoiceApi.ListInvoices()

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "TEMPORARILY_UNAVAILABLE")
}

func TestGetPrivateCloud(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"id": "218030",
			"customerId": "1301178860",
			"dataCenter": "AMS-01",
			"serviceOffering": "FLAT_FEE",
			"sla": "Bronze",
			"contract": {
				"id": "30000775",
				"startsAt": "2015-11-01T00:00:00+02:00",
				"endsAt": "2016-12-30T10:39:27+01:00",
				"billingCycle": 12,
				"billingFrequency": "MONTH",
				"pricePerFrequency": 0,
				"currency": "EUR"
			},
			"hardware": {
				"cpu": {
					"cores": 25
				},
				"memory": {
					"unit": "GB",
					"amount": 50
				},
				"storage": {
					"unit": "GB",
					"amount": 1
				}
			},
			"ips": [
				{
					"ip": "10.12.144.32",
					"version": 4,
					"type": "PUBLIC"
				}
			],
			"networkTraffic": {
				"type": "DATATRAFFIC",
				"trafficType": "PREMIUM",
				"datatrafficUnit": "TB",
				"datatrafficLimit": 6
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.GetPrivateCloud("218030")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Id, "218030")
	assert.Equal(response.CustomerId, "1301178860")
	assert.Equal(response.DataCenter, "AMS-01")
	assert.Equal(response.ServiceOffering, "FLAT_FEE")
	assert.Equal(response.Sla, "Bronze")
	assert.Equal(response.Contract.Id, "30000775")
	assert.Equal(response.Contract.StartsAt, "2015-11-01T00:00:00+02:00")
	assert.Equal(response.Contract.EndsAt, "2016-12-30T10:39:27+01:00")
	assert.Equal(response.Contract.BillingCycle, 12)
	assert.Equal(response.Contract.BillingFrequency, "MONTH")
	assert.Equal(response.Contract.PricePerFrequency, float32(0))
	assert.Equal(response.Contract.Currency, "EUR")
	assert.Equal(response.Hardware.Cpu.Cores, 25)
	assert.Equal(response.Hardware.Memory.Amount, 50)
	assert.Equal(response.Hardware.Memory.Unit, "GB")
	assert.Equal(response.Hardware.Storage.Amount, 1)
	assert.Equal(response.Hardware.Storage.Unit, "GB")
	assert.Equal(response.Ips[0].Ip, "10.12.144.32")
	assert.Equal(response.Ips[0].Type, "PUBLIC")
	assert.Equal(response.Ips[0].Version, 4)
	assert.Equal(response.NetworkTraffic.DataTrafficLimit, 6)
	assert.Equal(response.NetworkTraffic.DataTrafficUnit, "TB")
	assert.Equal(response.NetworkTraffic.TrafficType, "PREMIUM")
	assert.Equal(response.NetworkTraffic.Type, "DATATRAFFIC")
}

func TestGetPrivateCloudError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetPrivateCloud("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "ACCESS_DENIED")
}

func TestGetPrivateCloudError404(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"errorCode": "404", "errorMessage": "Resource 218030 was not found", "userMessage": "Resource with id 218030 not found."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetPrivateCloud("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Resource 218030 was not found")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "404")
	assert.Equal(lswErr.UserMessage, "Resource with id 218030 not found.")
}

func TestGetPrivateCloudError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetPrivateCloud("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "SERVER_ERROR")
}

func TestGetPrivateCloudError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetPrivateCloud("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "TEMPORARILY_UNAVAILABLE")
}

func TestListCredentials(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 1}, "credentials": [
			{
				"type": "REMOTE_MANAGEMENT",
				"username": "root",
				"domain": "123456"
			}
		]}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.ListCredentials("12345678", "REMOTE_MANAGEMENT")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 1)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 1)

	credential := response.Credentials[0]
	assert.Equal(credential.Type, "REMOTE_MANAGEMENT")
	assert.Equal(credential.Username, "root")
	assert.Equal(credential.Domain, "123456")
}

func TestListCredentialsBeEmpty(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 0, "totalCount": 0}, "credentials": []}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.ListCredentials("12345678", "REMOTE_MANAGEMENT")

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 0)
	assert.Equal(response.Metadata.Offset, 0)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 0)
}

func TestListCredentialsPaginate(t *testing.T) {
	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{"_metadata":{"limit": 10, "offset": 1, "totalCount": 11}, "credentials": [
			{
				"type": "REMOTE_MANAGEMENT",
				"username": "root",
				"password": "password123",
				"domain": "123456"
			}
		]}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.ListCredentials("12345678", "REMOTE_MANAGEMENT", 1, 10)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(response.Metadata.TotalCount, 11)
	assert.Equal(response.Metadata.Offset, 1)
	assert.Equal(response.Metadata.Limit, 10)
	assert.Equal(len(response.Credentials), 1)

	credential := response.Credentials[0]
	assert.Equal(credential.Type, "REMOTE_MANAGEMENT")
	assert.Equal(credential.Username, "root")
	assert.Equal(credential.Password, "password123")
	assert.Equal(credential.Domain, "123456")
}

func TestListCredentialsError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.ListCredentials("12345678", "REMOTE_MANAGEMENT")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "ACCESS_DENIED")
}

func TestListCredentialsError404(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"errorCode": "404", "errorMessage": "Resource 218030 was not found", "userMessage": "Resource with id 218030 not found."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.ListCredentials("218030", "REMOTE_MANAGEMENT")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Resource 218030 was not found")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "404")
	assert.Equal(lswErr.UserMessage, "Resource with id 218030 not found.")
}

func TestListCredentialsError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.ListCredentials("12345678", "REMOTE_MANAGEMENT")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "SERVER_ERROR")
}

func TestListCredentialsError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.ListCredentials("12345678", "REMOTE_MANAGEMENT")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "TEMPORARILY_UNAVAILABLE")
}

func TestGetCredentials(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"type": "REMOTE_MANAGEMENT",
			"username": "root",
			"password": "password123",
			"domain": "123456"
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.GetCredentials("218030", "REMOTE_MANAGEMENT", "root")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Type, "REMOTE_MANAGEMENT")
	assert.Equal(response.Username, "root")
	assert.Equal(response.Password, "password123")
	assert.Equal(response.Domain, "123456")
}

func TestGetCredentialsError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetCredentials("218030", "REMOTE_MANAGEMENT", "root")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "ACCESS_DENIED")
}

func TestGetCredentialsError404(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"errorCode": "404", "errorMessage": "Resource 218030 was not found", "userMessage": "Resource with id 218030 not found."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetCredentials("218030", "REMOTE_MANAGEMENT", "root")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Resource 218030 was not found")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "404")
	assert.Equal(lswErr.UserMessage, "Resource with id 218030 not found.")
}

func TestGetCredentialsError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetCredentials("218030", "REMOTE_MANAGEMENT", "root")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "SERVER_ERROR")
}

func TestGetCredentialsError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetCredentials("218030", "REMOTE_MANAGEMENT", "root")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "TEMPORARILY_UNAVAILABLE")
}

func TestGetDataTrafficMetrics(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "SUM"
			},
			"metrics": {
				"DATATRAFFIC_UP": {
					"unit": "GB",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 900
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 2500
						}
					]
				},
				"DATATRAFFIC_DOWN": {
					"unit": "GB",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 90
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 250
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.GetDataTrafficMetrics("218030")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "SUM")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.DataTrafficDown.Unit, "GB")
	assert.Equal(response.Metric.DataTrafficDown.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.DataTrafficDown.Values[0].Value, 90)
	assert.Equal(response.Metric.DataTrafficDown.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.DataTrafficDown.Values[1].Value, 250)
	assert.Equal(response.Metric.DataTrafficUp.Unit, "GB")
	assert.Equal(response.Metric.DataTrafficUp.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.DataTrafficUp.Values[0].Value, 900)
	assert.Equal(response.Metric.DataTrafficUp.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.DataTrafficUp.Values[1].Value, 2500)
}

func TestGetDataTrafficMetricsWithFilter(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "SUM"
			},
			"metrics": {
				"DATATRAFFIC_UP": {
					"unit": "GB",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 900
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 2500
						}
					]
				},
				"DATATRAFFIC_DOWN": {
					"unit": "GB",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 90
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 250
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.GetDataTrafficMetrics("218030", "SUM", "MONTH", "2017-07-01T00:00:00+00:00", "2017-07-02T00:00:00+00:00")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "SUM")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.DataTrafficDown.Unit, "GB")
	assert.Equal(response.Metric.DataTrafficDown.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.DataTrafficDown.Values[0].Value, 90)
	assert.Equal(response.Metric.DataTrafficDown.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.DataTrafficDown.Values[1].Value, 250)
	assert.Equal(response.Metric.DataTrafficUp.Unit, "GB")
	assert.Equal(response.Metric.DataTrafficUp.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.DataTrafficUp.Values[0].Value, 900)
	assert.Equal(response.Metric.DataTrafficUp.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.DataTrafficUp.Values[1].Value, 2500)
}

func TestGetDataTrafficMetricsError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetDataTrafficMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "ACCESS_DENIED")
}

func TestGetDataTrafficMetricsError404(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"errorCode": "404", "errorMessage": "Resource 218030 was not found", "userMessage": "Resource with id 218030 not found."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetDataTrafficMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Resource 218030 was not found")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "404")
	assert.Equal(lswErr.UserMessage, "Resource with id 218030 not found.")
}

func TestGetDataTrafficMetricsError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetDataTrafficMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "SERVER_ERROR")
}

func TestGetDataTrafficMetricsError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetDataTrafficMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "TEMPORARILY_UNAVAILABLE")
}

func TestGetBandWidthMetrics(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "AVG"
			},
			"metrics": {
				"DOWN_PUBLIC": {
					"unit": "bps",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 28202556
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 28202557
						}
					]
				},
				"UP_PUBLIC": {
					"unit": "bps",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 158317518
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 158317519
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.GetBandWidthMetrics("218030")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "AVG")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.DownPublic.Unit, "bps")
	assert.Equal(response.Metric.DownPublic.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.DownPublic.Values[0].Value, 28202556)
	assert.Equal(response.Metric.DownPublic.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.DownPublic.Values[1].Value, 28202557)
	assert.Equal(response.Metric.UpPublic.Unit, "bps")
	assert.Equal(response.Metric.UpPublic.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.UpPublic.Values[0].Value, 158317518)
	assert.Equal(response.Metric.UpPublic.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.UpPublic.Values[1].Value, 158317519)
}

func TestGetBandWidthMetricsWithFilter(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "AVG"
			},
			"metrics": {
				"DOWN_PUBLIC": {
					"unit": "bps",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 28202556
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 28202557
						}
					]
				},
				"UP_PUBLIC": {
					"unit": "bps",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 158317518
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 158317519
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.GetBandWidthMetrics("218030", "AVG", "MONTH", "2017-07-01T00:00:00+00:00", "2017-07-02T00:00:00+00:00")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "AVG")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.DownPublic.Unit, "bps")
	assert.Equal(response.Metric.DownPublic.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.DownPublic.Values[0].Value, 28202556)
	assert.Equal(response.Metric.DownPublic.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.DownPublic.Values[1].Value, 28202557)
	assert.Equal(response.Metric.UpPublic.Unit, "bps")
	assert.Equal(response.Metric.UpPublic.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.UpPublic.Values[0].Value, 158317518)
	assert.Equal(response.Metric.UpPublic.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.UpPublic.Values[1].Value, 158317519)
}

func TestGetBandWidthMetricsError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetBandWidthMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "ACCESS_DENIED")
}

func TestGetBandWidthMetricsError404(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"errorCode": "404", "errorMessage": "Resource 218030 was not found", "userMessage": "Resource with id 218030 not found."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetBandWidthMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Resource 218030 was not found")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "404")
	assert.Equal(lswErr.UserMessage, "Resource with id 218030 not found.")
}

func TestGetBandWidthMetricsError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetBandWidthMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "SERVER_ERROR")
}

func TestGetBandWidthMetricsError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetBandWidthMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "TEMPORARILY_UNAVAILABLE")
}

func TestGetCpuMetrics(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-01T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "MAX"
			},
			"metrics": {
				"CPU": {
					"unit": "CORES",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 24
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 24
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.GetCpuMetrics("218030")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "MAX")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.Cpu.Unit, "CORES")
	assert.Equal(response.Metric.Cpu.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.Cpu.Values[0].Value, 24)
	assert.Equal(response.Metric.Cpu.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.Cpu.Values[1].Value, 24)
}

func TestGetCpuMetricsWithFilter(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "MAX"
			},
			"metrics": {
				"CPU": {
					"unit": "CORES",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 24
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 24
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.GetCpuMetrics("218030", "MAX", "MONTH", "2017-07-01T00:00:00+00:00", "2017-07-02T00:00:00+00:00")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "MAX")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.Cpu.Unit, "CORES")
	assert.Equal(response.Metric.Cpu.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.Cpu.Values[0].Value, 24)
	assert.Equal(response.Metric.Cpu.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.Cpu.Values[1].Value, 24)
}

func TestGetCpuMetricsError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetCpuMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "ACCESS_DENIED")
}

func TestGetCpuMetricsError404(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"errorCode": "404", "errorMessage": "Resource 218030 was not found", "userMessage": "Resource with id 218030 not found."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetCpuMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Resource 218030 was not found")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "404")
	assert.Equal(lswErr.UserMessage, "Resource with id 218030 not found.")
}

func TestGetCpuMetricsError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetCpuMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "SERVER_ERROR")
}

func TestGetCpuMetricsError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetCpuMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "TEMPORARILY_UNAVAILABLE")
}

func TestGetMemoryMetrics(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "MAX"
			},
			"metrics": {
				"MEMORY": {
					"unit": "GB",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 8
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 16
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.GetMemoryMetrics("218030")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "MAX")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.Memory.Unit, "GB")
	assert.Equal(response.Metric.Memory.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.Memory.Values[0].Value, 8)
	assert.Equal(response.Metric.Memory.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.Memory.Values[1].Value, 16)
}

func TestGetMemoryMetricsWithFilter(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "MAX"
			},
			"metrics": {
				"MEMORY": {
					"unit": "GB",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 8
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 16
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.GetMemoryMetrics("218030", "MAX", "MONTH", "2017-07-01T00:00:00+00:00", "2017-07-02T00:00:00+00:00")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "MAX")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.Memory.Unit, "GB")
	assert.Equal(response.Metric.Memory.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.Memory.Values[0].Value, 8)
	assert.Equal(response.Metric.Memory.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.Memory.Values[1].Value, 16)
}

func TestGetMemoryMetricsError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetMemoryMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "ACCESS_DENIED")
}

func TestGetMemoryMetricsError404(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"errorCode": "404", "errorMessage": "Resource 218030 was not found", "userMessage": "Resource with id 218030 not found."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetMemoryMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Resource 218030 was not found")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "404")
	assert.Equal(lswErr.UserMessage, "Resource with id 218030 not found.")
}

func TestGetMemoryMetricsError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetMemoryMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "SERVER_ERROR")
}

func TestGetMemoryMetricsError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetMemoryMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "TEMPORARILY_UNAVAILABLE")
}

func TestGetStorageMetrics(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "MAX"
			},
			"metrics": {
				"STORAGE": {
					"unit": "GB",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 900
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 2500
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.GetStorageMetrics("218030")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "MAX")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.Storage.Unit, "GB")
	assert.Equal(response.Metric.Storage.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.Storage.Values[0].Value, 900)
	assert.Equal(response.Metric.Storage.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.Storage.Values[1].Value, 2500)
}

func TestGetStorageMetricsWithFilter(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		fmt.Fprintf(w, `{
			"_metadata": {
				"from": "2017-07-01T00:00:00+00:00",
				"to": "2017-07-02T00:00:00+00:00",
				"granularity": "MONTH",
				"aggregation": "MAX"
			},
			"metrics": {
				"STORAGE": {
					"unit": "GB",
					"values": [
						{
							"timestamp": "2017-07-01T00:00:00+00:00",
							"value": 900
						},
						{
							"timestamp": "2017-07-02T00:00:00+00:00",
							"value": 2500
						}
					]
				}
			}
		}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	response, err := privateCloudApi.GetStorageMetrics("218030", "MAX", "MONTH", "2017-07-01T00:00:00+00:00", "2017-07-02T00:00:00+00:00")

	assert := assert.New(t)
	assert.Nil(err)

	assert.Equal(response.Metadata.From, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metadata.To, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metadata.Aggregation, "MAX")
	assert.Equal(response.Metadata.Granularity, "MONTH")
	assert.Equal(response.Metric.Storage.Unit, "GB")
	assert.Equal(response.Metric.Storage.Values[0].Timestamp, "2017-07-01T00:00:00+00:00")
	assert.Equal(response.Metric.Storage.Values[0].Value, 900)
	assert.Equal(response.Metric.Storage.Values[1].Timestamp, "2017-07-02T00:00:00+00:00")
	assert.Equal(response.Metric.Storage.Values[1].Value, 2500)
}

func TestGetStorageMetricsError403(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"errorCode": "ACCESS_DENIED", "errorMessage": "The access token is expired or invalid."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetStorageMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The access token is expired or invalid.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "ACCESS_DENIED")
}

func TestGetStorageMetricsError404(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"errorCode": "404", "errorMessage": "Resource 218030 was not found", "userMessage": "Resource with id 218030 not found."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetStorageMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Resource 218030 was not found")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "404")
	assert.Equal(lswErr.UserMessage, "Resource with id 218030 not found.")
}

func TestGetStorageMetricsError500(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"errorCode": "SERVER_ERROR", "errorMessage": "The server encountered an unexpected condition that prevented it from fulfilling the request."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetStorageMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "SERVER_ERROR")
}

func TestGetStorageMetricsError503(t *testing.T) {

	setup(func(w http.ResponseWriter, r *http.Request) {
		if h := r.Header.Get("x-lsw-auth"); h != "test-api-key" {
			t.Errorf("request did not have x-lsw-auth header set!")
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"errorCode": "TEMPORARILY_UNAVAILABLE", "errorMessage": "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."}`)
	})
	defer teardown()

	privateCloudApi := PrivateCloudApi{}
	resp, err := privateCloudApi.GetStorageMetrics("218030")

	assert := assert.New(t)
	assert.Empty(resp)
	assert.NotNil(err)
	assert.Equal(err.Error(), "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.")

	lswErr, ok := err.(*LeasewebError)
	assert.Equal(true, ok)
	assert.Equal(lswErr.ErrorCode, "TEMPORARILY_UNAVAILABLE")
}
