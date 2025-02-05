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

// ContractType Select HOURLY for billing based on hourly usage, else MONTHLY for billing per month usage
type ContractType string

// List of contractType
const (
	CONTRACTTYPE_HOURLY ContractType = "HOURLY"
	CONTRACTTYPE_MONTHLY ContractType = "MONTHLY"
)

// All allowed values of ContractType enum
var AllowedContractTypeEnumValues = []ContractType{
	"HOURLY",
	"MONTHLY",
}

func (v *ContractType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContractType(value)
	for _, existing := range AllowedContractTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContractType", value)
}

// NewContractTypeFromValue returns a pointer to a valid ContractType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContractTypeFromValue(v string) (*ContractType, error) {
	ev := ContractType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContractType: valid values are %v", v, AllowedContractTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContractType) IsValid() bool {
	for _, existing := range AllowedContractTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to contractType value
func (v ContractType) Ptr() *ContractType {
	return &v
}

type NullableContractType struct {
	value *ContractType
	isSet bool
}

func (v NullableContractType) Get() *ContractType {
	return v.value
}

func (v *NullableContractType) Set(val *ContractType) {
	v.value = val
	v.isSet = true
}

func (v NullableContractType) IsSet() bool {
	return v.isSet
}

func (v *NullableContractType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractType(val *ContractType) *NullableContractType {
	return &NullableContractType{value: val, isSet: true}
}

func (v NullableContractType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

