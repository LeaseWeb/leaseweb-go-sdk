/*
Public Clouds

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
	"fmt"
)

// NetworkType the model 'NetworkType'
type NetworkType string

// List of networkType
const (
	NETWORKTYPE_INTERNAL NetworkType = "INTERNAL"
	NETWORKTYPE_PUBLIC NetworkType = "PUBLIC"
)

// All allowed values of NetworkType enum
var AllowedNetworkTypeEnumValues = []NetworkType{
	"INTERNAL",
	"PUBLIC",
}

func (v *NetworkType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkType(value)
	for _, existing := range AllowedNetworkTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkType", value)
}

// NewNetworkTypeFromValue returns a pointer to a valid NetworkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkTypeFromValue(v string) (*NetworkType, error) {
	ev := NetworkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkType: valid values are %v", v, AllowedNetworkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkType) IsValid() bool {
	for _, existing := range AllowedNetworkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to networkType value
func (v NetworkType) Ptr() *NetworkType {
	return &v
}

type NullableNetworkType struct {
	value *NetworkType
	isSet bool
}

func (v NullableNetworkType) Get() *NetworkType {
	return v.value
}

func (v *NullableNetworkType) Set(val *NetworkType) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkType) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkType(val *NetworkType) *NullableNetworkType {
	return &NullableNetworkType{value: val, isSet: true}
}

func (v NullableNetworkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

