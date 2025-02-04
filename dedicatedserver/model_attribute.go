/*
Dedicated Servers API

This is the description of the Dedicated Server API.  The base url of this API is `https://api.leaseweb.com`.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedserver

import (
	"encoding/json"
)

// checks if the Attribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Attribute{}

// Attribute struct for Attribute
type Attribute struct {
	Flag *string `json:"flag,omitempty"`
	Id *string `json:"id,omitempty"`
	RawValue *string `json:"raw_value,omitempty"`
	Thresh *string `json:"thresh,omitempty"`
	Type *string `json:"type,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Value *string `json:"value,omitempty"`
	WhenFailed *string `json:"when_failed,omitempty"`
	Worst *string `json:"worst,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Attribute Attribute

// NewAttribute instantiates a new Attribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttribute() *Attribute {
	this := Attribute{}
	return &this
}

// NewAttributeWithDefaults instantiates a new Attribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeWithDefaults() *Attribute {
	this := Attribute{}
	return &this
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *Attribute) GetFlag() string {
	if o == nil || IsNil(o.Flag) {
		var ret string
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetFlagOk() (*string, bool) {
	if o == nil || IsNil(o.Flag) {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *Attribute) HasFlag() bool {
	if o != nil && !IsNil(o.Flag) {
		return true
	}

	return false
}

// SetFlag gets a reference to the given string and assigns it to the Flag field.
func (o *Attribute) SetFlag(v string) {
	o.Flag = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Attribute) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Attribute) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Attribute) SetId(v string) {
	o.Id = &v
}

// GetRawValue returns the RawValue field value if set, zero value otherwise.
func (o *Attribute) GetRawValue() string {
	if o == nil || IsNil(o.RawValue) {
		var ret string
		return ret
	}
	return *o.RawValue
}

// GetRawValueOk returns a tuple with the RawValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetRawValueOk() (*string, bool) {
	if o == nil || IsNil(o.RawValue) {
		return nil, false
	}
	return o.RawValue, true
}

// HasRawValue returns a boolean if a field has been set.
func (o *Attribute) HasRawValue() bool {
	if o != nil && !IsNil(o.RawValue) {
		return true
	}

	return false
}

// SetRawValue gets a reference to the given string and assigns it to the RawValue field.
func (o *Attribute) SetRawValue(v string) {
	o.RawValue = &v
}

// GetThresh returns the Thresh field value if set, zero value otherwise.
func (o *Attribute) GetThresh() string {
	if o == nil || IsNil(o.Thresh) {
		var ret string
		return ret
	}
	return *o.Thresh
}

// GetThreshOk returns a tuple with the Thresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetThreshOk() (*string, bool) {
	if o == nil || IsNil(o.Thresh) {
		return nil, false
	}
	return o.Thresh, true
}

// HasThresh returns a boolean if a field has been set.
func (o *Attribute) HasThresh() bool {
	if o != nil && !IsNil(o.Thresh) {
		return true
	}

	return false
}

// SetThresh gets a reference to the given string and assigns it to the Thresh field.
func (o *Attribute) SetThresh(v string) {
	o.Thresh = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Attribute) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Attribute) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Attribute) SetType(v string) {
	o.Type = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *Attribute) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Attribute) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *Attribute) SetUpdated(v string) {
	o.Updated = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Attribute) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Attribute) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Attribute) SetValue(v string) {
	o.Value = &v
}

// GetWhenFailed returns the WhenFailed field value if set, zero value otherwise.
func (o *Attribute) GetWhenFailed() string {
	if o == nil || IsNil(o.WhenFailed) {
		var ret string
		return ret
	}
	return *o.WhenFailed
}

// GetWhenFailedOk returns a tuple with the WhenFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetWhenFailedOk() (*string, bool) {
	if o == nil || IsNil(o.WhenFailed) {
		return nil, false
	}
	return o.WhenFailed, true
}

// HasWhenFailed returns a boolean if a field has been set.
func (o *Attribute) HasWhenFailed() bool {
	if o != nil && !IsNil(o.WhenFailed) {
		return true
	}

	return false
}

// SetWhenFailed gets a reference to the given string and assigns it to the WhenFailed field.
func (o *Attribute) SetWhenFailed(v string) {
	o.WhenFailed = &v
}

// GetWorst returns the Worst field value if set, zero value otherwise.
func (o *Attribute) GetWorst() string {
	if o == nil || IsNil(o.Worst) {
		var ret string
		return ret
	}
	return *o.Worst
}

// GetWorstOk returns a tuple with the Worst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetWorstOk() (*string, bool) {
	if o == nil || IsNil(o.Worst) {
		return nil, false
	}
	return o.Worst, true
}

// HasWorst returns a boolean if a field has been set.
func (o *Attribute) HasWorst() bool {
	if o != nil && !IsNil(o.Worst) {
		return true
	}

	return false
}

// SetWorst gets a reference to the given string and assigns it to the Worst field.
func (o *Attribute) SetWorst(v string) {
	o.Worst = &v
}

func (o Attribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Attribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Flag) {
		toSerialize["flag"] = o.Flag
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RawValue) {
		toSerialize["raw_value"] = o.RawValue
	}
	if !IsNil(o.Thresh) {
		toSerialize["thresh"] = o.Thresh
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.WhenFailed) {
		toSerialize["when_failed"] = o.WhenFailed
	}
	if !IsNil(o.Worst) {
		toSerialize["worst"] = o.Worst
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Attribute) UnmarshalJSON(data []byte) (err error) {
	varAttribute := _Attribute{}

	err = json.Unmarshal(data, &varAttribute)

	if err != nil {
		return err
	}

	*o = Attribute(varAttribute)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "flag")
		delete(additionalProperties, "id")
		delete(additionalProperties, "raw_value")
		delete(additionalProperties, "thresh")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updated")
		delete(additionalProperties, "value")
		delete(additionalProperties, "when_failed")
		delete(additionalProperties, "worst")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAttribute struct {
	value *Attribute
	isSet bool
}

func (v NullableAttribute) Get() *Attribute {
	return v.value
}

func (v *NullableAttribute) Set(val *Attribute) {
	v.value = val
	v.isSet = true
}

func (v NullableAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttribute(val *Attribute) *NullableAttribute {
	return &NullableAttribute{value: val, isSet: true}
}

func (v NullableAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


