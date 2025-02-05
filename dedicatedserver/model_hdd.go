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

// checks if the Hdd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Hdd{}

// Hdd A single object of the hard disk drive
type Hdd struct {
	// Id of the hard disk drive
	Id *string `json:"id,omitempty"`
	// The total amount of hard disk drives
	Amount *int32 `json:"amount,omitempty"`
	// The size number of the hard disk drive
	Size *float32 `json:"size,omitempty"`
	// The type of the hard disk drive
	Type *string `json:"type,omitempty"`
	// The unit of the hard disk drive
	Unit *string `json:"unit,omitempty"`
	// Hard disk drive performance type
	PerformanceType NullableString `json:"performanceType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Hdd Hdd

// NewHdd instantiates a new Hdd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHdd() *Hdd {
	this := Hdd{}
	return &this
}

// NewHddWithDefaults instantiates a new Hdd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHddWithDefaults() *Hdd {
	this := Hdd{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Hdd) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hdd) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Hdd) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Hdd) SetId(v string) {
	o.Id = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Hdd) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hdd) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Hdd) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *Hdd) SetAmount(v int32) {
	o.Amount = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Hdd) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hdd) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Hdd) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *Hdd) SetSize(v float32) {
	o.Size = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Hdd) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hdd) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Hdd) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Hdd) SetType(v string) {
	o.Type = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Hdd) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hdd) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Hdd) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Hdd) SetUnit(v string) {
	o.Unit = &v
}

// GetPerformanceType returns the PerformanceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Hdd) GetPerformanceType() string {
	if o == nil || IsNil(o.PerformanceType.Get()) {
		var ret string
		return ret
	}
	return *o.PerformanceType.Get()
}

// GetPerformanceTypeOk returns a tuple with the PerformanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Hdd) GetPerformanceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerformanceType.Get(), o.PerformanceType.IsSet()
}

// HasPerformanceType returns a boolean if a field has been set.
func (o *Hdd) HasPerformanceType() bool {
	if o != nil && o.PerformanceType.IsSet() {
		return true
	}

	return false
}

// SetPerformanceType gets a reference to the given NullableString and assigns it to the PerformanceType field.
func (o *Hdd) SetPerformanceType(v string) {
	o.PerformanceType.Set(&v)
}
// SetPerformanceTypeNil sets the value for PerformanceType to be an explicit nil
func (o *Hdd) SetPerformanceTypeNil() {
	o.PerformanceType.Set(nil)
}

// UnsetPerformanceType ensures that no value is present for PerformanceType, not even an explicit nil
func (o *Hdd) UnsetPerformanceType() {
	o.PerformanceType.Unset()
}

func (o Hdd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Hdd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if o.PerformanceType.IsSet() {
		toSerialize["performanceType"] = o.PerformanceType.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Hdd) UnmarshalJSON(data []byte) (err error) {
	varHdd := _Hdd{}

	err = json.Unmarshal(data, &varHdd)

	if err != nil {
		return err
	}

	*o = Hdd(varHdd)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "size")
		delete(additionalProperties, "type")
		delete(additionalProperties, "unit")
		delete(additionalProperties, "performanceType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHdd struct {
	value *Hdd
	isSet bool
}

func (v NullableHdd) Get() *Hdd {
	return v.value
}

func (v *NullableHdd) Set(val *Hdd) {
	v.value = val
	v.isSet = true
}

func (v NullableHdd) IsSet() bool {
	return v.isSet
}

func (v *NullableHdd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHdd(val *Hdd) *NullableHdd {
	return &NullableHdd{value: val, isSet: true}
}

func (v NullableHdd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHdd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


