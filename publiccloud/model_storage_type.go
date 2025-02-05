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

// StorageType Storage type
type StorageType string

// List of storageType
const (
	STORAGETYPE_LOCAL StorageType = "LOCAL"
	STORAGETYPE_CENTRAL StorageType = "CENTRAL"
)

// All allowed values of StorageType enum
var AllowedStorageTypeEnumValues = []StorageType{
	"LOCAL",
	"CENTRAL",
}

func (v *StorageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StorageType(value)
	for _, existing := range AllowedStorageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StorageType", value)
}

// NewStorageTypeFromValue returns a pointer to a valid StorageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStorageTypeFromValue(v string) (*StorageType, error) {
	ev := StorageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StorageType: valid values are %v", v, AllowedStorageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StorageType) IsValid() bool {
	for _, existing := range AllowedStorageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to storageType value
func (v StorageType) Ptr() *StorageType {
	return &v
}

type NullableStorageType struct {
	value *StorageType
	isSet bool
}

func (v NullableStorageType) Get() *StorageType {
	return v.value
}

func (v *NullableStorageType) Set(val *StorageType) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageType) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageType(val *StorageType) *NullableStorageType {
	return &NullableStorageType{value: val, isSet: true}
}

func (v NullableStorageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

