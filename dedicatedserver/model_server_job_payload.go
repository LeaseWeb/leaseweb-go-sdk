/*
Leaseweb API for dedicated servers

This documents the rest api dedicatedserver provides.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedserver

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// ServerJobPayload - struct for ServerJobPayload
type ServerJobPayload struct {
	InstallOperatingSystemPayload *InstallOperatingSystemPayload
	Payload *Payload
}

// InstallOperatingSystemPayloadAsServerJobPayload is a convenience function that returns InstallOperatingSystemPayload wrapped in ServerJobPayload
func InstallOperatingSystemPayloadAsServerJobPayload(v *InstallOperatingSystemPayload) ServerJobPayload {
	return ServerJobPayload{
		InstallOperatingSystemPayload: v,
	}
}

// PayloadAsServerJobPayload is a convenience function that returns Payload wrapped in ServerJobPayload
func PayloadAsServerJobPayload(v *Payload) ServerJobPayload {
	return ServerJobPayload{
		Payload: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ServerJobPayload) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InstallOperatingSystemPayload
	err = newStrictDecoder(data).Decode(&dst.InstallOperatingSystemPayload)
	if err == nil {
		jsonInstallOperatingSystemPayload, _ := json.Marshal(dst.InstallOperatingSystemPayload)
		if string(jsonInstallOperatingSystemPayload) == "{}" { // empty struct
			dst.InstallOperatingSystemPayload = nil
		} else {
			if err = validator.Validate(dst.InstallOperatingSystemPayload); err != nil {
				dst.InstallOperatingSystemPayload = nil
			} else {
				match++
			}
		}
	} else {
		dst.InstallOperatingSystemPayload = nil
	}

	// try to unmarshal data into Payload
	err = newStrictDecoder(data).Decode(&dst.Payload)
	if err == nil {
		jsonPayload, _ := json.Marshal(dst.Payload)
		if string(jsonPayload) == "{}" { // empty struct
			dst.Payload = nil
		} else {
			if err = validator.Validate(dst.Payload); err != nil {
				dst.Payload = nil
			} else {
				match++
			}
		}
	} else {
		dst.Payload = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InstallOperatingSystemPayload = nil
		dst.Payload = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ServerJobPayload)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ServerJobPayload)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ServerJobPayload) MarshalJSON() ([]byte, error) {
	if src.InstallOperatingSystemPayload != nil {
		return json.Marshal(&src.InstallOperatingSystemPayload)
	}

	if src.Payload != nil {
		return json.Marshal(&src.Payload)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ServerJobPayload) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InstallOperatingSystemPayload != nil {
		return obj.InstallOperatingSystemPayload
	}

	if obj.Payload != nil {
		return obj.Payload
	}

	// all schemas are nil
	return nil
}

type NullableServerJobPayload struct {
	value *ServerJobPayload
	isSet bool
}

func (v NullableServerJobPayload) Get() *ServerJobPayload {
	return v.value
}

func (v *NullableServerJobPayload) Set(val *ServerJobPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableServerJobPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableServerJobPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerJobPayload(val *ServerJobPayload) *NullableServerJobPayload {
	return &NullableServerJobPayload{value: val, isSet: true}
}

func (v NullableServerJobPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerJobPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


