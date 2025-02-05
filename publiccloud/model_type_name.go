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

// TypeName Instance/Load balancer type
type TypeName string

// List of typeName
const (
	TYPENAME_M3_LARGE TypeName = "lsw.m3.large"
	TYPENAME_M3_XLARGE TypeName = "lsw.m3.xlarge"
	TYPENAME_M3_2XLARGE TypeName = "lsw.m3.2xlarge"
	TYPENAME_M4_LARGE TypeName = "lsw.m4.large"
	TYPENAME_M4_XLARGE TypeName = "lsw.m4.xlarge"
	TYPENAME_M4_2XLARGE TypeName = "lsw.m4.2xlarge"
	TYPENAME_M4_4XLARGE TypeName = "lsw.m4.4xlarge"
	TYPENAME_M5_LARGE TypeName = "lsw.m5.large"
	TYPENAME_M5_XLARGE TypeName = "lsw.m5.xlarge"
	TYPENAME_M5_2XLARGE TypeName = "lsw.m5.2xlarge"
	TYPENAME_M5_4XLARGE TypeName = "lsw.m5.4xlarge"
	TYPENAME_M5A_LARGE TypeName = "lsw.m5a.large"
	TYPENAME_M5A_XLARGE TypeName = "lsw.m5a.xlarge"
	TYPENAME_M5A_2XLARGE TypeName = "lsw.m5a.2xlarge"
	TYPENAME_M5A_4XLARGE TypeName = "lsw.m5a.4xlarge"
	TYPENAME_M5A_8XLARGE TypeName = "lsw.m5a.8xlarge"
	TYPENAME_M5A_12XLARGE TypeName = "lsw.m5a.12xlarge"
	TYPENAME_M6A_LARGE TypeName = "lsw.m6a.large"
	TYPENAME_M6A_XLARGE TypeName = "lsw.m6a.xlarge"
	TYPENAME_M6A_2XLARGE TypeName = "lsw.m6a.2xlarge"
	TYPENAME_M6A_4XLARGE TypeName = "lsw.m6a.4xlarge"
	TYPENAME_M6A_8XLARGE TypeName = "lsw.m6a.8xlarge"
	TYPENAME_M6A_12XLARGE TypeName = "lsw.m6a.12xlarge"
	TYPENAME_M6A_16XLARGE TypeName = "lsw.m6a.16xlarge"
	TYPENAME_M6A_24XLARGE TypeName = "lsw.m6a.24xlarge"
	TYPENAME_C3_LARGE TypeName = "lsw.c3.large"
	TYPENAME_C3_XLARGE TypeName = "lsw.c3.xlarge"
	TYPENAME_C3_2XLARGE TypeName = "lsw.c3.2xlarge"
	TYPENAME_C3_4XLARGE TypeName = "lsw.c3.4xlarge"
	TYPENAME_C4_LARGE TypeName = "lsw.c4.large"
	TYPENAME_C4_XLARGE TypeName = "lsw.c4.xlarge"
	TYPENAME_C4_2XLARGE TypeName = "lsw.c4.2xlarge"
	TYPENAME_C4_4XLARGE TypeName = "lsw.c4.4xlarge"
	TYPENAME_C5_LARGE TypeName = "lsw.c5.large"
	TYPENAME_C5_XLARGE TypeName = "lsw.c5.xlarge"
	TYPENAME_C5_2XLARGE TypeName = "lsw.c5.2xlarge"
	TYPENAME_C5_4XLARGE TypeName = "lsw.c5.4xlarge"
	TYPENAME_C5A_LARGE TypeName = "lsw.c5a.large"
	TYPENAME_C5A_XLARGE TypeName = "lsw.c5a.xlarge"
	TYPENAME_C5A_2XLARGE TypeName = "lsw.c5a.2xlarge"
	TYPENAME_C5A_4XLARGE TypeName = "lsw.c5a.4xlarge"
	TYPENAME_C5A_9XLARGE TypeName = "lsw.c5a.9xlarge"
	TYPENAME_C5A_12XLARGE TypeName = "lsw.c5a.12xlarge"
	TYPENAME_C6A_LARGE TypeName = "lsw.c6a.large"
	TYPENAME_C6A_XLARGE TypeName = "lsw.c6a.xlarge"
	TYPENAME_C6A_2XLARGE TypeName = "lsw.c6a.2xlarge"
	TYPENAME_C6A_4XLARGE TypeName = "lsw.c6a.4xlarge"
	TYPENAME_C6A_8XLARGE TypeName = "lsw.c6a.8xlarge"
	TYPENAME_C6A_12XLARGE TypeName = "lsw.c6a.12xlarge"
	TYPENAME_C6A_16XLARGE TypeName = "lsw.c6a.16xlarge"
	TYPENAME_C6A_24XLARGE TypeName = "lsw.c6a.24xlarge"
	TYPENAME_R3_LARGE TypeName = "lsw.r3.large"
	TYPENAME_R3_XLARGE TypeName = "lsw.r3.xlarge"
	TYPENAME_R3_2XLARGE TypeName = "lsw.r3.2xlarge"
	TYPENAME_R4_LARGE TypeName = "lsw.r4.large"
	TYPENAME_R4_XLARGE TypeName = "lsw.r4.xlarge"
	TYPENAME_R4_2XLARGE TypeName = "lsw.r4.2xlarge"
	TYPENAME_R5_LARGE TypeName = "lsw.r5.large"
	TYPENAME_R5_XLARGE TypeName = "lsw.r5.xlarge"
	TYPENAME_R5_2XLARGE TypeName = "lsw.r5.2xlarge"
	TYPENAME_R5A_LARGE TypeName = "lsw.r5a.large"
	TYPENAME_R5A_XLARGE TypeName = "lsw.r5a.xlarge"
	TYPENAME_R5A_2XLARGE TypeName = "lsw.r5a.2xlarge"
	TYPENAME_R5A_4XLARGE TypeName = "lsw.r5a.4xlarge"
	TYPENAME_R5A_8XLARGE TypeName = "lsw.r5a.8xlarge"
	TYPENAME_R5A_12XLARGE TypeName = "lsw.r5a.12xlarge"
	TYPENAME_R6A_LARGE TypeName = "lsw.r6a.large"
	TYPENAME_R6A_XLARGE TypeName = "lsw.r6a.xlarge"
	TYPENAME_R6A_2XLARGE TypeName = "lsw.r6a.2xlarge"
	TYPENAME_R6A_4XLARGE TypeName = "lsw.r6a.4xlarge"
	TYPENAME_R6A_8XLARGE TypeName = "lsw.r6a.8xlarge"
	TYPENAME_R6A_12XLARGE TypeName = "lsw.r6a.12xlarge"
	TYPENAME_R6A_16XLARGE TypeName = "lsw.r6a.16xlarge"
	TYPENAME_R6A_24XLARGE TypeName = "lsw.r6a.24xlarge"
)

// All allowed values of TypeName enum
var AllowedTypeNameEnumValues = []TypeName{
	"lsw.m3.large",
	"lsw.m3.xlarge",
	"lsw.m3.2xlarge",
	"lsw.m4.large",
	"lsw.m4.xlarge",
	"lsw.m4.2xlarge",
	"lsw.m4.4xlarge",
	"lsw.m5.large",
	"lsw.m5.xlarge",
	"lsw.m5.2xlarge",
	"lsw.m5.4xlarge",
	"lsw.m5a.large",
	"lsw.m5a.xlarge",
	"lsw.m5a.2xlarge",
	"lsw.m5a.4xlarge",
	"lsw.m5a.8xlarge",
	"lsw.m5a.12xlarge",
	"lsw.m6a.large",
	"lsw.m6a.xlarge",
	"lsw.m6a.2xlarge",
	"lsw.m6a.4xlarge",
	"lsw.m6a.8xlarge",
	"lsw.m6a.12xlarge",
	"lsw.m6a.16xlarge",
	"lsw.m6a.24xlarge",
	"lsw.c3.large",
	"lsw.c3.xlarge",
	"lsw.c3.2xlarge",
	"lsw.c3.4xlarge",
	"lsw.c4.large",
	"lsw.c4.xlarge",
	"lsw.c4.2xlarge",
	"lsw.c4.4xlarge",
	"lsw.c5.large",
	"lsw.c5.xlarge",
	"lsw.c5.2xlarge",
	"lsw.c5.4xlarge",
	"lsw.c5a.large",
	"lsw.c5a.xlarge",
	"lsw.c5a.2xlarge",
	"lsw.c5a.4xlarge",
	"lsw.c5a.9xlarge",
	"lsw.c5a.12xlarge",
	"lsw.c6a.large",
	"lsw.c6a.xlarge",
	"lsw.c6a.2xlarge",
	"lsw.c6a.4xlarge",
	"lsw.c6a.8xlarge",
	"lsw.c6a.12xlarge",
	"lsw.c6a.16xlarge",
	"lsw.c6a.24xlarge",
	"lsw.r3.large",
	"lsw.r3.xlarge",
	"lsw.r3.2xlarge",
	"lsw.r4.large",
	"lsw.r4.xlarge",
	"lsw.r4.2xlarge",
	"lsw.r5.large",
	"lsw.r5.xlarge",
	"lsw.r5.2xlarge",
	"lsw.r5a.large",
	"lsw.r5a.xlarge",
	"lsw.r5a.2xlarge",
	"lsw.r5a.4xlarge",
	"lsw.r5a.8xlarge",
	"lsw.r5a.12xlarge",
	"lsw.r6a.large",
	"lsw.r6a.xlarge",
	"lsw.r6a.2xlarge",
	"lsw.r6a.4xlarge",
	"lsw.r6a.8xlarge",
	"lsw.r6a.12xlarge",
	"lsw.r6a.16xlarge",
	"lsw.r6a.24xlarge",
}

func (v *TypeName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeName(value)
	for _, existing := range AllowedTypeNameEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeName", value)
}

// NewTypeNameFromValue returns a pointer to a valid TypeName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeNameFromValue(v string) (*TypeName, error) {
	ev := TypeName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypeName: valid values are %v", v, AllowedTypeNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeName) IsValid() bool {
	for _, existing := range AllowedTypeNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to typeName value
func (v TypeName) Ptr() *TypeName {
	return &v
}

type NullableTypeName struct {
	value *TypeName
	isSet bool
}

func (v NullableTypeName) Get() *TypeName {
	return v.value
}

func (v *NullableTypeName) Set(val *TypeName) {
	v.value = val
	v.isSet = true
}

func (v NullableTypeName) IsSet() bool {
	return v.isSet
}

func (v *NullableTypeName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypeName(val *TypeName) *NullableTypeName {
	return &NullableTypeName{value: val, isSet: true}
}

func (v NullableTypeName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypeName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

