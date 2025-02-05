/*
Dedicated Network Equipments

This is the description of the Dedicated Network Equipment API.  The base url of this API is `https://api.leaseweb.com`.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkequipment

import (
	"encoding/json"
)

// checks if the UpdateIpOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIpOpts{}

// UpdateIpOpts struct for UpdateIpOpts
type UpdateIpOpts struct {
	// The detection profile value
	DetectionProfile *string `json:"detectionProfile,omitempty"`
	// The reverse lookup value
	ReverseLookup *string `json:"reverseLookup,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateIpOpts UpdateIpOpts

// NewUpdateIpOpts instantiates a new UpdateIpOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIpOpts() *UpdateIpOpts {
	this := UpdateIpOpts{}
	return &this
}

// NewUpdateIpOptsWithDefaults instantiates a new UpdateIpOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIpOptsWithDefaults() *UpdateIpOpts {
	this := UpdateIpOpts{}
	return &this
}

// GetDetectionProfile returns the DetectionProfile field value if set, zero value otherwise.
func (o *UpdateIpOpts) GetDetectionProfile() string {
	if o == nil || IsNil(o.DetectionProfile) {
		var ret string
		return ret
	}
	return *o.DetectionProfile
}

// GetDetectionProfileOk returns a tuple with the DetectionProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpOpts) GetDetectionProfileOk() (*string, bool) {
	if o == nil || IsNil(o.DetectionProfile) {
		return nil, false
	}
	return o.DetectionProfile, true
}

// HasDetectionProfile returns a boolean if a field has been set.
func (o *UpdateIpOpts) HasDetectionProfile() bool {
	if o != nil && !IsNil(o.DetectionProfile) {
		return true
	}

	return false
}

// SetDetectionProfile gets a reference to the given string and assigns it to the DetectionProfile field.
func (o *UpdateIpOpts) SetDetectionProfile(v string) {
	o.DetectionProfile = &v
}

// GetReverseLookup returns the ReverseLookup field value if set, zero value otherwise.
func (o *UpdateIpOpts) GetReverseLookup() string {
	if o == nil || IsNil(o.ReverseLookup) {
		var ret string
		return ret
	}
	return *o.ReverseLookup
}

// GetReverseLookupOk returns a tuple with the ReverseLookup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpOpts) GetReverseLookupOk() (*string, bool) {
	if o == nil || IsNil(o.ReverseLookup) {
		return nil, false
	}
	return o.ReverseLookup, true
}

// HasReverseLookup returns a boolean if a field has been set.
func (o *UpdateIpOpts) HasReverseLookup() bool {
	if o != nil && !IsNil(o.ReverseLookup) {
		return true
	}

	return false
}

// SetReverseLookup gets a reference to the given string and assigns it to the ReverseLookup field.
func (o *UpdateIpOpts) SetReverseLookup(v string) {
	o.ReverseLookup = &v
}

func (o UpdateIpOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIpOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DetectionProfile) {
		toSerialize["detectionProfile"] = o.DetectionProfile
	}
	if !IsNil(o.ReverseLookup) {
		toSerialize["reverseLookup"] = o.ReverseLookup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateIpOpts) UnmarshalJSON(data []byte) (err error) {
	varUpdateIpOpts := _UpdateIpOpts{}

	err = json.Unmarshal(data, &varUpdateIpOpts)

	if err != nil {
		return err
	}

	*o = UpdateIpOpts(varUpdateIpOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "detectionProfile")
		delete(additionalProperties, "reverseLookup")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateIpOpts struct {
	value *UpdateIpOpts
	isSet bool
}

func (v NullableUpdateIpOpts) Get() *UpdateIpOpts {
	return v.value
}

func (v *NullableUpdateIpOpts) Set(val *UpdateIpOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIpOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIpOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIpOpts(val *UpdateIpOpts) *NullableUpdateIpOpts {
	return &NullableUpdateIpOpts{value: val, isSet: true}
}

func (v NullableUpdateIpOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIpOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


