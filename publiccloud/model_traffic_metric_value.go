/*
Public Clouds

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
	"time"
)

// checks if the TrafficMetricValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrafficMetricValue{}

// TrafficMetricValue struct for TrafficMetricValue
type TrafficMetricValue struct {
	// Bytes consumed
	Value *int32 `json:"value,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TrafficMetricValue TrafficMetricValue

// NewTrafficMetricValue instantiates a new TrafficMetricValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrafficMetricValue() *TrafficMetricValue {
	this := TrafficMetricValue{}
	return &this
}

// NewTrafficMetricValueWithDefaults instantiates a new TrafficMetricValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrafficMetricValueWithDefaults() *TrafficMetricValue {
	this := TrafficMetricValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TrafficMetricValue) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficMetricValue) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TrafficMetricValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *TrafficMetricValue) SetValue(v int32) {
	o.Value = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TrafficMetricValue) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficMetricValue) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TrafficMetricValue) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *TrafficMetricValue) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o TrafficMetricValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrafficMetricValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TrafficMetricValue) UnmarshalJSON(data []byte) (err error) {
	varTrafficMetricValue := _TrafficMetricValue{}

	err = json.Unmarshal(data, &varTrafficMetricValue)

	if err != nil {
		return err
	}

	*o = TrafficMetricValue(varTrafficMetricValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTrafficMetricValue struct {
	value *TrafficMetricValue
	isSet bool
}

func (v NullableTrafficMetricValue) Get() *TrafficMetricValue {
	return v.value
}

func (v *NullableTrafficMetricValue) Set(val *TrafficMetricValue) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficMetricValue) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficMetricValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficMetricValue(val *TrafficMetricValue) *NullableTrafficMetricValue {
	return &NullableTrafficMetricValue{value: val, isSet: true}
}

func (v NullableTrafficMetricValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficMetricValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


