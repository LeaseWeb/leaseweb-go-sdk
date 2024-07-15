/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
	"fmt"
)

// Balance Algorithm to be used for load balancer
type Balance string

// List of balance
const (
	BALANCE_ROUNDROBIN Balance = "roundrobin"
	BALANCE_LEASTCONN Balance = "leastconn"
	BALANCE_SOURCE Balance = "source"
)

// All allowed values of Balance enum
var AllowedBalanceEnumValues = []Balance{
	"roundrobin",
	"leastconn",
	"source",
}

func (v *Balance) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Balance(value)
	for _, existing := range AllowedBalanceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Balance", value)
}

// NewBalanceFromValue returns a pointer to a valid Balance
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBalanceFromValue(v string) (*Balance, error) {
	ev := Balance(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Balance: valid values are %v", v, AllowedBalanceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Balance) IsValid() bool {
	for _, existing := range AllowedBalanceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to balance value
func (v Balance) Ptr() *Balance {
	return &v
}

type NullableBalance struct {
	value *Balance
	isSet bool
}

func (v NullableBalance) Get() *Balance {
	return v.value
}

func (v *NullableBalance) Set(val *Balance) {
	v.value = val
	v.isSet = true
}

func (v NullableBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalance(val *Balance) *NullableBalance {
	return &NullableBalance{value: val, isSet: true}
}

func (v NullableBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

