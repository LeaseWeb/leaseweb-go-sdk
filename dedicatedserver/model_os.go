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

// checks if the Os type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Os{}

// Os struct for Os
type Os struct {
	// The architecture of the operating system
	Architecture *string `json:"architecture,omitempty"`
	// The operating system family
	Family *string `json:"family,omitempty"`
	// A human readable name for the operating system
	Name *string `json:"name,omitempty"`
	// The type of operating system
	Type *string `json:"type,omitempty"`
	// The version of the operating system
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Os Os

// NewOs instantiates a new Os object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOs() *Os {
	this := Os{}
	return &this
}

// NewOsWithDefaults instantiates a new Os object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsWithDefaults() *Os {
	this := Os{}
	return &this
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *Os) GetArchitecture() string {
	if o == nil || IsNil(o.Architecture) {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Os) GetArchitectureOk() (*string, bool) {
	if o == nil || IsNil(o.Architecture) {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *Os) HasArchitecture() bool {
	if o != nil && !IsNil(o.Architecture) {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *Os) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *Os) GetFamily() string {
	if o == nil || IsNil(o.Family) {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Os) GetFamilyOk() (*string, bool) {
	if o == nil || IsNil(o.Family) {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *Os) HasFamily() bool {
	if o != nil && !IsNil(o.Family) {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *Os) SetFamily(v string) {
	o.Family = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Os) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Os) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Os) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Os) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Os) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Os) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Os) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Os) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Os) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Os) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Os) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Os) SetVersion(v string) {
	o.Version = &v
}

func (o Os) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Os) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Architecture) {
		toSerialize["architecture"] = o.Architecture
	}
	if !IsNil(o.Family) {
		toSerialize["family"] = o.Family
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Os) UnmarshalJSON(data []byte) (err error) {
	varOs := _Os{}

	err = json.Unmarshal(data, &varOs)

	if err != nil {
		return err
	}

	*o = Os(varOs)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "architecture")
		delete(additionalProperties, "family")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOs struct {
	value *Os
	isSet bool
}

func (v NullableOs) Get() *Os {
	return v.value
}

func (v *NullableOs) Set(val *Os) {
	v.value = val
	v.isSet = true
}

func (v NullableOs) IsSet() bool {
	return v.isSet
}

func (v *NullableOs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOs(val *Os) *NullableOs {
	return &NullableOs{value: val, isSet: true}
}

func (v NullableOs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


