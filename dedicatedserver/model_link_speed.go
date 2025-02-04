/*
Dedicated Servers API

This is the description of the Dedicated Server API.  The base url of this API is `https://api.leaseweb.com`.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedserver

import (
	"encoding/json"
	"fmt"
)

// LinkSpeed The port speed in Mbps
type LinkSpeed int32

// List of linkSpeed
const (
	LINKSPEED__100 LinkSpeed = 100
	LINKSPEED__1000 LinkSpeed = 1000
	LINKSPEED__10000 LinkSpeed = 10000
)

// All allowed values of LinkSpeed enum
var AllowedLinkSpeedEnumValues = []LinkSpeed{
	100,
	1000,
	10000,
}

func (v *LinkSpeed) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LinkSpeed(value)
	for _, existing := range AllowedLinkSpeedEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LinkSpeed", value)
}

// NewLinkSpeedFromValue returns a pointer to a valid LinkSpeed
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLinkSpeedFromValue(v int32) (*LinkSpeed, error) {
	ev := LinkSpeed(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LinkSpeed: valid values are %v", v, AllowedLinkSpeedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LinkSpeed) IsValid() bool {
	for _, existing := range AllowedLinkSpeedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to linkSpeed value
func (v LinkSpeed) Ptr() *LinkSpeed {
	return &v
}

type NullableLinkSpeed struct {
	value *LinkSpeed
	isSet bool
}

func (v NullableLinkSpeed) Get() *LinkSpeed {
	return v.value
}

func (v *NullableLinkSpeed) Set(val *LinkSpeed) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkSpeed) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkSpeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkSpeed(val *LinkSpeed) *NullableLinkSpeed {
	return &NullableLinkSpeed{value: val, isSet: true}
}

func (v NullableLinkSpeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkSpeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

