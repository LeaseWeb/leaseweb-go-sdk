/*
Dedicated Servers

This is the description of the Dedicated Server API.  The base url of this API is `https://api.leaseweb.com`.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedserver

import (
	"encoding/json"
)

// checks if the HardwareCpu type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HardwareCpu{}

// HardwareCpu struct for HardwareCpu
type HardwareCpu struct {
	Capabilities *CpuCapabilities `json:"capabilities,omitempty"`
	Description *string `json:"description,omitempty"`
	Hz *string `json:"hz,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty"`
	Settings *CpuSettings `json:"settings,omitempty"`
	Slot *string `json:"slot,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HardwareCpu HardwareCpu

// NewHardwareCpu instantiates a new HardwareCpu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHardwareCpu() *HardwareCpu {
	this := HardwareCpu{}
	return &this
}

// NewHardwareCpuWithDefaults instantiates a new HardwareCpu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHardwareCpuWithDefaults() *HardwareCpu {
	this := HardwareCpu{}
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *HardwareCpu) GetCapabilities() CpuCapabilities {
	if o == nil || IsNil(o.Capabilities) {
		var ret CpuCapabilities
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareCpu) GetCapabilitiesOk() (*CpuCapabilities, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *HardwareCpu) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given CpuCapabilities and assigns it to the Capabilities field.
func (o *HardwareCpu) SetCapabilities(v CpuCapabilities) {
	o.Capabilities = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HardwareCpu) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareCpu) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HardwareCpu) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HardwareCpu) SetDescription(v string) {
	o.Description = &v
}

// GetHz returns the Hz field value if set, zero value otherwise.
func (o *HardwareCpu) GetHz() string {
	if o == nil || IsNil(o.Hz) {
		var ret string
		return ret
	}
	return *o.Hz
}

// GetHzOk returns a tuple with the Hz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareCpu) GetHzOk() (*string, bool) {
	if o == nil || IsNil(o.Hz) {
		return nil, false
	}
	return o.Hz, true
}

// HasHz returns a boolean if a field has been set.
func (o *HardwareCpu) HasHz() bool {
	if o != nil && !IsNil(o.Hz) {
		return true
	}

	return false
}

// SetHz gets a reference to the given string and assigns it to the Hz field.
func (o *HardwareCpu) SetHz(v string) {
	o.Hz = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *HardwareCpu) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareCpu) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *HardwareCpu) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *HardwareCpu) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *HardwareCpu) GetSettings() CpuSettings {
	if o == nil || IsNil(o.Settings) {
		var ret CpuSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareCpu) GetSettingsOk() (*CpuSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *HardwareCpu) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given CpuSettings and assigns it to the Settings field.
func (o *HardwareCpu) SetSettings(v CpuSettings) {
	o.Settings = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *HardwareCpu) GetSlot() string {
	if o == nil || IsNil(o.Slot) {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareCpu) GetSlotOk() (*string, bool) {
	if o == nil || IsNil(o.Slot) {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *HardwareCpu) HasSlot() bool {
	if o != nil && !IsNil(o.Slot) {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *HardwareCpu) SetSlot(v string) {
	o.Slot = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *HardwareCpu) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareCpu) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *HardwareCpu) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *HardwareCpu) SetVendor(v string) {
	o.Vendor = &v
}

func (o HardwareCpu) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HardwareCpu) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Hz) {
		toSerialize["hz"] = o.Hz
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Slot) {
		toSerialize["slot"] = o.Slot
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HardwareCpu) UnmarshalJSON(data []byte) (err error) {
	varHardwareCpu := _HardwareCpu{}

	err = json.Unmarshal(data, &varHardwareCpu)

	if err != nil {
		return err
	}

	*o = HardwareCpu(varHardwareCpu)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "capabilities")
		delete(additionalProperties, "description")
		delete(additionalProperties, "hz")
		delete(additionalProperties, "serial_number")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "slot")
		delete(additionalProperties, "vendor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHardwareCpu struct {
	value *HardwareCpu
	isSet bool
}

func (v NullableHardwareCpu) Get() *HardwareCpu {
	return v.value
}

func (v *NullableHardwareCpu) Set(val *HardwareCpu) {
	v.value = val
	v.isSet = true
}

func (v NullableHardwareCpu) IsSet() bool {
	return v.isSet
}

func (v *NullableHardwareCpu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHardwareCpu(val *HardwareCpu) *NullableHardwareCpu {
	return &NullableHardwareCpu{value: val, isSet: true}
}

func (v NullableHardwareCpu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHardwareCpu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


