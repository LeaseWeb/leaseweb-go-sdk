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

// checks if the DataTrafficNotificationSettingOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataTrafficNotificationSettingOpts{}

// DataTrafficNotificationSettingOpts struct for DataTrafficNotificationSettingOpts
type DataTrafficNotificationSettingOpts struct {
	// Frequency for the Data Traffic Notification
	Frequency string `json:"frequency"`
	// Threshold Value for the Data Traffic Notification
	Threshold string `json:"threshold"`
	// Unit for the Data Traffic Notification
	Unit string `json:"unit"`
	AdditionalProperties map[string]interface{}
}

type _DataTrafficNotificationSettingOpts DataTrafficNotificationSettingOpts

// NewDataTrafficNotificationSettingOpts instantiates a new DataTrafficNotificationSettingOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataTrafficNotificationSettingOpts(frequency string, threshold string, unit string) *DataTrafficNotificationSettingOpts {
	this := DataTrafficNotificationSettingOpts{}
	this.Frequency = frequency
	this.Threshold = threshold
	this.Unit = unit
	return &this
}

// NewDataTrafficNotificationSettingOptsWithDefaults instantiates a new DataTrafficNotificationSettingOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataTrafficNotificationSettingOptsWithDefaults() *DataTrafficNotificationSettingOpts {
	this := DataTrafficNotificationSettingOpts{}
	return &this
}

// GetFrequency returns the Frequency field value
func (o *DataTrafficNotificationSettingOpts) GetFrequency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *DataTrafficNotificationSettingOpts) GetFrequencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *DataTrafficNotificationSettingOpts) SetFrequency(v string) {
	o.Frequency = v
}

// GetThreshold returns the Threshold field value
func (o *DataTrafficNotificationSettingOpts) GetThreshold() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *DataTrafficNotificationSettingOpts) GetThresholdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *DataTrafficNotificationSettingOpts) SetThreshold(v string) {
	o.Threshold = v
}

// GetUnit returns the Unit field value
func (o *DataTrafficNotificationSettingOpts) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *DataTrafficNotificationSettingOpts) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *DataTrafficNotificationSettingOpts) SetUnit(v string) {
	o.Unit = v
}

func (o DataTrafficNotificationSettingOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataTrafficNotificationSettingOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["frequency"] = o.Frequency
	toSerialize["threshold"] = o.Threshold
	toSerialize["unit"] = o.Unit

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DataTrafficNotificationSettingOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"frequency",
		"threshold",
		"unit",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDataTrafficNotificationSettingOpts := _DataTrafficNotificationSettingOpts{}

	err = json.Unmarshal(data, &varDataTrafficNotificationSettingOpts)

	if err != nil {
		return err
	}

	*o = DataTrafficNotificationSettingOpts(varDataTrafficNotificationSettingOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "frequency")
		delete(additionalProperties, "threshold")
		delete(additionalProperties, "unit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDataTrafficNotificationSettingOpts struct {
	value *DataTrafficNotificationSettingOpts
	isSet bool
}

func (v NullableDataTrafficNotificationSettingOpts) Get() *DataTrafficNotificationSettingOpts {
	return v.value
}

func (v *NullableDataTrafficNotificationSettingOpts) Set(val *DataTrafficNotificationSettingOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableDataTrafficNotificationSettingOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableDataTrafficNotificationSettingOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataTrafficNotificationSettingOpts(val *DataTrafficNotificationSettingOpts) *NullableDataTrafficNotificationSettingOpts {
	return &NullableDataTrafficNotificationSettingOpts{value: val, isSet: true}
}

func (v NullableDataTrafficNotificationSettingOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataTrafficNotificationSettingOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


