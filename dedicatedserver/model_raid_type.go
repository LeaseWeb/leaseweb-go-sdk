/*
Dedicated Servers

This is the description of the Dedicated Server API.  The base url of this API is `https://api.leaseweb.com`.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedserver

import (
	"encoding/json"
	"fmt"
)

// RaidType RAID type to apply. NONE is the equivalent of pass-through mode on HW RAID equipped servers
type RaidType string

// List of raidType
const (
	RAIDTYPE_HW RaidType = "HW"
	RAIDTYPE_SW RaidType = "SW"
	RAIDTYPE_NONE RaidType = "NONE"
)

// All allowed values of RaidType enum
var AllowedRaidTypeEnumValues = []RaidType{
	"HW",
	"SW",
	"NONE",
}

func (v *RaidType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RaidType(value)
	for _, existing := range AllowedRaidTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RaidType", value)
}

// NewRaidTypeFromValue returns a pointer to a valid RaidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRaidTypeFromValue(v string) (*RaidType, error) {
	ev := RaidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RaidType: valid values are %v", v, AllowedRaidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RaidType) IsValid() bool {
	for _, existing := range AllowedRaidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to raidType value
func (v RaidType) Ptr() *RaidType {
	return &v
}

type NullableRaidType struct {
	value *RaidType
	isSet bool
}

func (v NullableRaidType) Get() *RaidType {
	return v.value
}

func (v *NullableRaidType) Set(val *RaidType) {
	v.value = val
	v.isSet = true
}

func (v NullableRaidType) IsSet() bool {
	return v.isSet
}

func (v *NullableRaidType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRaidType(val *RaidType) *NullableRaidType {
	return &NullableRaidType{value: val, isSet: true}
}

func (v NullableRaidType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRaidType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

