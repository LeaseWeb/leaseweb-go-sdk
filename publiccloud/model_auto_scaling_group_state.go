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

// AutoScalingGroupState The Auto Scaling Group's current state.
type AutoScalingGroupState string

// List of autoScalingGroupState
const (
	AUTOSCALINGGROUPSTATE_ACTIVE AutoScalingGroupState = "ACTIVE"
	AUTOSCALINGGROUPSTATE_CREATING AutoScalingGroupState = "CREATING"
	AUTOSCALINGGROUPSTATE_CREATED AutoScalingGroupState = "CREATED"
	AUTOSCALINGGROUPSTATE_DESTROYED AutoScalingGroupState = "DESTROYED"
	AUTOSCALINGGROUPSTATE_DESTROYING AutoScalingGroupState = "DESTROYING"
	AUTOSCALINGGROUPSTATE_SCALING AutoScalingGroupState = "SCALING"
	AUTOSCALINGGROUPSTATE_UPDATING AutoScalingGroupState = "UPDATING"
)

// All allowed values of AutoScalingGroupState enum
var AllowedAutoScalingGroupStateEnumValues = []AutoScalingGroupState{
	"ACTIVE",
	"CREATING",
	"CREATED",
	"DESTROYED",
	"DESTROYING",
	"SCALING",
	"UPDATING",
}

func (v *AutoScalingGroupState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AutoScalingGroupState(value)
	for _, existing := range AllowedAutoScalingGroupStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AutoScalingGroupState", value)
}

// NewAutoScalingGroupStateFromValue returns a pointer to a valid AutoScalingGroupState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAutoScalingGroupStateFromValue(v string) (*AutoScalingGroupState, error) {
	ev := AutoScalingGroupState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AutoScalingGroupState: valid values are %v", v, AllowedAutoScalingGroupStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AutoScalingGroupState) IsValid() bool {
	for _, existing := range AllowedAutoScalingGroupStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to autoScalingGroupState value
func (v AutoScalingGroupState) Ptr() *AutoScalingGroupState {
	return &v
}

type NullableAutoScalingGroupState struct {
	value *AutoScalingGroupState
	isSet bool
}

func (v NullableAutoScalingGroupState) Get() *AutoScalingGroupState {
	return v.value
}

func (v *NullableAutoScalingGroupState) Set(val *AutoScalingGroupState) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoScalingGroupState) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoScalingGroupState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoScalingGroupState(val *AutoScalingGroupState) *NullableAutoScalingGroupState {
	return &NullableAutoScalingGroupState{value: val, isSet: true}
}

func (v NullableAutoScalingGroupState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoScalingGroupState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

