/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
	"fmt"
)

// LoadBalancerMetricsGranularity Defines the time interval for data aggregation
type LoadBalancerMetricsGranularity string

// List of loadBalancerMetricsGranularity
const (
	LOADBALANCERMETRICSGRANULARITY__5M LoadBalancerMetricsGranularity = "5m"
	LOADBALANCERMETRICSGRANULARITY__10M LoadBalancerMetricsGranularity = "10m"
	LOADBALANCERMETRICSGRANULARITY__30M LoadBalancerMetricsGranularity = "30m"
	LOADBALANCERMETRICSGRANULARITY__1H LoadBalancerMetricsGranularity = "1h"
	LOADBALANCERMETRICSGRANULARITY__1D LoadBalancerMetricsGranularity = "1d"
	LOADBALANCERMETRICSGRANULARITY__1W LoadBalancerMetricsGranularity = "1w"
)

// All allowed values of LoadBalancerMetricsGranularity enum
var AllowedLoadBalancerMetricsGranularityEnumValues = []LoadBalancerMetricsGranularity{
	"5m",
	"10m",
	"30m",
	"1h",
	"1d",
	"1w",
}

func (v *LoadBalancerMetricsGranularity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LoadBalancerMetricsGranularity(value)
	for _, existing := range AllowedLoadBalancerMetricsGranularityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LoadBalancerMetricsGranularity", value)
}

// NewLoadBalancerMetricsGranularityFromValue returns a pointer to a valid LoadBalancerMetricsGranularity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLoadBalancerMetricsGranularityFromValue(v string) (*LoadBalancerMetricsGranularity, error) {
	ev := LoadBalancerMetricsGranularity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LoadBalancerMetricsGranularity: valid values are %v", v, AllowedLoadBalancerMetricsGranularityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LoadBalancerMetricsGranularity) IsValid() bool {
	for _, existing := range AllowedLoadBalancerMetricsGranularityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to loadBalancerMetricsGranularity value
func (v LoadBalancerMetricsGranularity) Ptr() *LoadBalancerMetricsGranularity {
	return &v
}

type NullableLoadBalancerMetricsGranularity struct {
	value *LoadBalancerMetricsGranularity
	isSet bool
}

func (v NullableLoadBalancerMetricsGranularity) Get() *LoadBalancerMetricsGranularity {
	return v.value
}

func (v *NullableLoadBalancerMetricsGranularity) Set(val *LoadBalancerMetricsGranularity) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerMetricsGranularity) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerMetricsGranularity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerMetricsGranularity(val *LoadBalancerMetricsGranularity) *NullableLoadBalancerMetricsGranularity {
	return &NullableLoadBalancerMetricsGranularity{value: val, isSet: true}
}

func (v NullableLoadBalancerMetricsGranularity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerMetricsGranularity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

