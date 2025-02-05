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

// checks if the EnableRescueModeOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnableRescueModeOpts{}

// EnableRescueModeOpts struct for EnableRescueModeOpts
type EnableRescueModeOpts struct {
	// Url which will receive callbacks
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	// Rescue mode password. If not provided, it would be automatically generated
	Password *string `json:"password,omitempty"`
	// Base64 Encoded string containing a valid bash script to be run right after rescue mode is launched
	PostInstallScript *string `json:"postInstallScript,omitempty"`
	// If set to `true`, server will be power cycled in order to complete the operation
	PowerCycle *bool `json:"powerCycle,omitempty"`
	// Rescue image identifier
	RescueImageId string `json:"rescueImageId"`
	// User ssh keys
	SshKeys *string `json:"sshKeys,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnableRescueModeOpts EnableRescueModeOpts

// NewEnableRescueModeOpts instantiates a new EnableRescueModeOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableRescueModeOpts(rescueImageId string) *EnableRescueModeOpts {
	this := EnableRescueModeOpts{}
	var powerCycle bool = true
	this.PowerCycle = &powerCycle
	this.RescueImageId = rescueImageId
	return &this
}

// NewEnableRescueModeOptsWithDefaults instantiates a new EnableRescueModeOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableRescueModeOptsWithDefaults() *EnableRescueModeOpts {
	this := EnableRescueModeOpts{}
	var powerCycle bool = true
	this.PowerCycle = &powerCycle
	var rescueImageId string = "GRML"
	this.RescueImageId = rescueImageId
	return &this
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *EnableRescueModeOpts) GetCallbackUrl() string {
	if o == nil || IsNil(o.CallbackUrl) {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableRescueModeOpts) GetCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackUrl) {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *EnableRescueModeOpts) HasCallbackUrl() bool {
	if o != nil && !IsNil(o.CallbackUrl) {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *EnableRescueModeOpts) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *EnableRescueModeOpts) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableRescueModeOpts) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *EnableRescueModeOpts) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *EnableRescueModeOpts) SetPassword(v string) {
	o.Password = &v
}

// GetPostInstallScript returns the PostInstallScript field value if set, zero value otherwise.
func (o *EnableRescueModeOpts) GetPostInstallScript() string {
	if o == nil || IsNil(o.PostInstallScript) {
		var ret string
		return ret
	}
	return *o.PostInstallScript
}

// GetPostInstallScriptOk returns a tuple with the PostInstallScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableRescueModeOpts) GetPostInstallScriptOk() (*string, bool) {
	if o == nil || IsNil(o.PostInstallScript) {
		return nil, false
	}
	return o.PostInstallScript, true
}

// HasPostInstallScript returns a boolean if a field has been set.
func (o *EnableRescueModeOpts) HasPostInstallScript() bool {
	if o != nil && !IsNil(o.PostInstallScript) {
		return true
	}

	return false
}

// SetPostInstallScript gets a reference to the given string and assigns it to the PostInstallScript field.
func (o *EnableRescueModeOpts) SetPostInstallScript(v string) {
	o.PostInstallScript = &v
}

// GetPowerCycle returns the PowerCycle field value if set, zero value otherwise.
func (o *EnableRescueModeOpts) GetPowerCycle() bool {
	if o == nil || IsNil(o.PowerCycle) {
		var ret bool
		return ret
	}
	return *o.PowerCycle
}

// GetPowerCycleOk returns a tuple with the PowerCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableRescueModeOpts) GetPowerCycleOk() (*bool, bool) {
	if o == nil || IsNil(o.PowerCycle) {
		return nil, false
	}
	return o.PowerCycle, true
}

// HasPowerCycle returns a boolean if a field has been set.
func (o *EnableRescueModeOpts) HasPowerCycle() bool {
	if o != nil && !IsNil(o.PowerCycle) {
		return true
	}

	return false
}

// SetPowerCycle gets a reference to the given bool and assigns it to the PowerCycle field.
func (o *EnableRescueModeOpts) SetPowerCycle(v bool) {
	o.PowerCycle = &v
}

// GetRescueImageId returns the RescueImageId field value
func (o *EnableRescueModeOpts) GetRescueImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RescueImageId
}

// GetRescueImageIdOk returns a tuple with the RescueImageId field value
// and a boolean to check if the value has been set.
func (o *EnableRescueModeOpts) GetRescueImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RescueImageId, true
}

// SetRescueImageId sets field value
func (o *EnableRescueModeOpts) SetRescueImageId(v string) {
	o.RescueImageId = v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *EnableRescueModeOpts) GetSshKeys() string {
	if o == nil || IsNil(o.SshKeys) {
		var ret string
		return ret
	}
	return *o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableRescueModeOpts) GetSshKeysOk() (*string, bool) {
	if o == nil || IsNil(o.SshKeys) {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *EnableRescueModeOpts) HasSshKeys() bool {
	if o != nil && !IsNil(o.SshKeys) {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given string and assigns it to the SshKeys field.
func (o *EnableRescueModeOpts) SetSshKeys(v string) {
	o.SshKeys = &v
}

func (o EnableRescueModeOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableRescueModeOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallbackUrl) {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.PostInstallScript) {
		toSerialize["postInstallScript"] = o.PostInstallScript
	}
	if !IsNil(o.PowerCycle) {
		toSerialize["powerCycle"] = o.PowerCycle
	}
	toSerialize["rescueImageId"] = o.RescueImageId
	if !IsNil(o.SshKeys) {
		toSerialize["sshKeys"] = o.SshKeys
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnableRescueModeOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rescueImageId",
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

	varEnableRescueModeOpts := _EnableRescueModeOpts{}

	err = json.Unmarshal(data, &varEnableRescueModeOpts)

	if err != nil {
		return err
	}

	*o = EnableRescueModeOpts(varEnableRescueModeOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "callbackUrl")
		delete(additionalProperties, "password")
		delete(additionalProperties, "postInstallScript")
		delete(additionalProperties, "powerCycle")
		delete(additionalProperties, "rescueImageId")
		delete(additionalProperties, "sshKeys")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnableRescueModeOpts struct {
	value *EnableRescueModeOpts
	isSet bool
}

func (v NullableEnableRescueModeOpts) Get() *EnableRescueModeOpts {
	return v.value
}

func (v *NullableEnableRescueModeOpts) Set(val *EnableRescueModeOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableRescueModeOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableRescueModeOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableRescueModeOpts(val *EnableRescueModeOpts) *NullableEnableRescueModeOpts {
	return &NullableEnableRescueModeOpts{value: val, isSet: true}
}

func (v NullableEnableRescueModeOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableRescueModeOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


