/*
Dedicated Network Equipment API

This is the description of the Dedicated Network Equipment API.  The base url of this API is `https://api.leaseweb.com`.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkequipment

import (
	"encoding/json"
)

// checks if the FeatureAvailability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureAvailability{}

// FeatureAvailability List of features that are available for this device
type FeatureAvailability struct {
	Automation *bool `json:"automation,omitempty"`
	PowerCycle *bool `json:"powerCycle,omitempty"`
	IpmiReboot *bool `json:"ipmiReboot,omitempty"`
	PrivateNetwork *bool `json:"privateNetwork,omitempty"`
	RemoteManagement *bool `json:"remoteManagement,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FeatureAvailability FeatureAvailability

// NewFeatureAvailability instantiates a new FeatureAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureAvailability() *FeatureAvailability {
	this := FeatureAvailability{}
	return &this
}

// NewFeatureAvailabilityWithDefaults instantiates a new FeatureAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureAvailabilityWithDefaults() *FeatureAvailability {
	this := FeatureAvailability{}
	return &this
}

// GetAutomation returns the Automation field value if set, zero value otherwise.
func (o *FeatureAvailability) GetAutomation() bool {
	if o == nil || IsNil(o.Automation) {
		var ret bool
		return ret
	}
	return *o.Automation
}

// GetAutomationOk returns a tuple with the Automation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureAvailability) GetAutomationOk() (*bool, bool) {
	if o == nil || IsNil(o.Automation) {
		return nil, false
	}
	return o.Automation, true
}

// HasAutomation returns a boolean if a field has been set.
func (o *FeatureAvailability) HasAutomation() bool {
	if o != nil && !IsNil(o.Automation) {
		return true
	}

	return false
}

// SetAutomation gets a reference to the given bool and assigns it to the Automation field.
func (o *FeatureAvailability) SetAutomation(v bool) {
	o.Automation = &v
}

// GetPowerCycle returns the PowerCycle field value if set, zero value otherwise.
func (o *FeatureAvailability) GetPowerCycle() bool {
	if o == nil || IsNil(o.PowerCycle) {
		var ret bool
		return ret
	}
	return *o.PowerCycle
}

// GetPowerCycleOk returns a tuple with the PowerCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureAvailability) GetPowerCycleOk() (*bool, bool) {
	if o == nil || IsNil(o.PowerCycle) {
		return nil, false
	}
	return o.PowerCycle, true
}

// HasPowerCycle returns a boolean if a field has been set.
func (o *FeatureAvailability) HasPowerCycle() bool {
	if o != nil && !IsNil(o.PowerCycle) {
		return true
	}

	return false
}

// SetPowerCycle gets a reference to the given bool and assigns it to the PowerCycle field.
func (o *FeatureAvailability) SetPowerCycle(v bool) {
	o.PowerCycle = &v
}

// GetIpmiReboot returns the IpmiReboot field value if set, zero value otherwise.
func (o *FeatureAvailability) GetIpmiReboot() bool {
	if o == nil || IsNil(o.IpmiReboot) {
		var ret bool
		return ret
	}
	return *o.IpmiReboot
}

// GetIpmiRebootOk returns a tuple with the IpmiReboot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureAvailability) GetIpmiRebootOk() (*bool, bool) {
	if o == nil || IsNil(o.IpmiReboot) {
		return nil, false
	}
	return o.IpmiReboot, true
}

// HasIpmiReboot returns a boolean if a field has been set.
func (o *FeatureAvailability) HasIpmiReboot() bool {
	if o != nil && !IsNil(o.IpmiReboot) {
		return true
	}

	return false
}

// SetIpmiReboot gets a reference to the given bool and assigns it to the IpmiReboot field.
func (o *FeatureAvailability) SetIpmiReboot(v bool) {
	o.IpmiReboot = &v
}

// GetPrivateNetwork returns the PrivateNetwork field value if set, zero value otherwise.
func (o *FeatureAvailability) GetPrivateNetwork() bool {
	if o == nil || IsNil(o.PrivateNetwork) {
		var ret bool
		return ret
	}
	return *o.PrivateNetwork
}

// GetPrivateNetworkOk returns a tuple with the PrivateNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureAvailability) GetPrivateNetworkOk() (*bool, bool) {
	if o == nil || IsNil(o.PrivateNetwork) {
		return nil, false
	}
	return o.PrivateNetwork, true
}

// HasPrivateNetwork returns a boolean if a field has been set.
func (o *FeatureAvailability) HasPrivateNetwork() bool {
	if o != nil && !IsNil(o.PrivateNetwork) {
		return true
	}

	return false
}

// SetPrivateNetwork gets a reference to the given bool and assigns it to the PrivateNetwork field.
func (o *FeatureAvailability) SetPrivateNetwork(v bool) {
	o.PrivateNetwork = &v
}

// GetRemoteManagement returns the RemoteManagement field value if set, zero value otherwise.
func (o *FeatureAvailability) GetRemoteManagement() bool {
	if o == nil || IsNil(o.RemoteManagement) {
		var ret bool
		return ret
	}
	return *o.RemoteManagement
}

// GetRemoteManagementOk returns a tuple with the RemoteManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureAvailability) GetRemoteManagementOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoteManagement) {
		return nil, false
	}
	return o.RemoteManagement, true
}

// HasRemoteManagement returns a boolean if a field has been set.
func (o *FeatureAvailability) HasRemoteManagement() bool {
	if o != nil && !IsNil(o.RemoteManagement) {
		return true
	}

	return false
}

// SetRemoteManagement gets a reference to the given bool and assigns it to the RemoteManagement field.
func (o *FeatureAvailability) SetRemoteManagement(v bool) {
	o.RemoteManagement = &v
}

func (o FeatureAvailability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureAvailability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Automation) {
		toSerialize["automation"] = o.Automation
	}
	if !IsNil(o.PowerCycle) {
		toSerialize["powerCycle"] = o.PowerCycle
	}
	if !IsNil(o.IpmiReboot) {
		toSerialize["ipmiReboot"] = o.IpmiReboot
	}
	if !IsNil(o.PrivateNetwork) {
		toSerialize["privateNetwork"] = o.PrivateNetwork
	}
	if !IsNil(o.RemoteManagement) {
		toSerialize["remoteManagement"] = o.RemoteManagement
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FeatureAvailability) UnmarshalJSON(data []byte) (err error) {
	varFeatureAvailability := _FeatureAvailability{}

	err = json.Unmarshal(data, &varFeatureAvailability)

	if err != nil {
		return err
	}

	*o = FeatureAvailability(varFeatureAvailability)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "automation")
		delete(additionalProperties, "powerCycle")
		delete(additionalProperties, "ipmiReboot")
		delete(additionalProperties, "privateNetwork")
		delete(additionalProperties, "remoteManagement")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFeatureAvailability struct {
	value *FeatureAvailability
	isSet bool
}

func (v NullableFeatureAvailability) Get() *FeatureAvailability {
	return v.value
}

func (v *NullableFeatureAvailability) Set(val *FeatureAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureAvailability(val *FeatureAvailability) *NullableFeatureAvailability {
	return &NullableFeatureAvailability{value: val, isSet: true}
}

func (v NullableFeatureAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


