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

// checks if the Vlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Vlan{}

// Vlan struct for Vlan
type Vlan struct {
	Id *string `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Vlan Vlan

// NewVlan instantiates a new Vlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVlan() *Vlan {
	this := Vlan{}
	return &this
}

// NewVlanWithDefaults instantiates a new Vlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVlanWithDefaults() *Vlan {
	this := Vlan{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Vlan) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vlan) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Vlan) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Vlan) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Vlan) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vlan) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Vlan) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Vlan) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Vlan) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vlan) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Vlan) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Vlan) SetName(v string) {
	o.Name = &v
}

func (o Vlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Vlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Vlan) UnmarshalJSON(data []byte) (err error) {
	varVlan := _Vlan{}

	err = json.Unmarshal(data, &varVlan)

	if err != nil {
		return err
	}

	*o = Vlan(varVlan)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVlan struct {
	value *Vlan
	isSet bool
}

func (v NullableVlan) Get() *Vlan {
	return v.value
}

func (v *NullableVlan) Set(val *Vlan) {
	v.value = val
	v.isSet = true
}

func (v NullableVlan) IsSet() bool {
	return v.isSet
}

func (v *NullableVlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVlan(val *Vlan) *NullableVlan {
	return &NullableVlan{value: val, isSet: true}
}

func (v NullableVlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


