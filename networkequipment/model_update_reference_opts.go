/*
Dedicated Network Equipment API

This is the description of the Dedicated Network Equipment API.  The base url of this API is `https://api.leaseweb.com`.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkequipment

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateReferenceOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateReferenceOpts{}

// UpdateReferenceOpts struct for UpdateReferenceOpts
type UpdateReferenceOpts struct {
	// The reference for this network equipment
	Reference string `json:"reference"`
	AdditionalProperties map[string]interface{}
}

type _UpdateReferenceOpts UpdateReferenceOpts

// NewUpdateReferenceOpts instantiates a new UpdateReferenceOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateReferenceOpts(reference string) *UpdateReferenceOpts {
	this := UpdateReferenceOpts{}
	this.Reference = reference
	return &this
}

// NewUpdateReferenceOptsWithDefaults instantiates a new UpdateReferenceOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateReferenceOptsWithDefaults() *UpdateReferenceOpts {
	this := UpdateReferenceOpts{}
	return &this
}

// GetReference returns the Reference field value
func (o *UpdateReferenceOpts) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *UpdateReferenceOpts) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *UpdateReferenceOpts) SetReference(v string) {
	o.Reference = v
}

func (o UpdateReferenceOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateReferenceOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reference"] = o.Reference

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateReferenceOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reference",
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

	varUpdateReferenceOpts := _UpdateReferenceOpts{}

	err = json.Unmarshal(data, &varUpdateReferenceOpts)

	if err != nil {
		return err
	}

	*o = UpdateReferenceOpts(varUpdateReferenceOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reference")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateReferenceOpts struct {
	value *UpdateReferenceOpts
	isSet bool
}

func (v NullableUpdateReferenceOpts) Get() *UpdateReferenceOpts {
	return v.value
}

func (v *NullableUpdateReferenceOpts) Set(val *UpdateReferenceOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateReferenceOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateReferenceOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateReferenceOpts(val *UpdateReferenceOpts) *NullableUpdateReferenceOpts {
	return &NullableUpdateReferenceOpts{value: val, isSet: true}
}

func (v NullableUpdateReferenceOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateReferenceOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


