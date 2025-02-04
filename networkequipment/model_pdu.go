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

// checks if the Pdu type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pdu{}

// Pdu Object describing the PDU power information
type Pdu struct {
	// The current power status of the network equipment.
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Pdu Pdu

// NewPdu instantiates a new Pdu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPdu() *Pdu {
	this := Pdu{}
	return &this
}

// NewPduWithDefaults instantiates a new Pdu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduWithDefaults() *Pdu {
	this := Pdu{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Pdu) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pdu) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Pdu) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Pdu) SetStatus(v string) {
	o.Status = &v
}

func (o Pdu) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pdu) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Pdu) UnmarshalJSON(data []byte) (err error) {
	varPdu := _Pdu{}

	err = json.Unmarshal(data, &varPdu)

	if err != nil {
		return err
	}

	*o = Pdu(varPdu)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePdu struct {
	value *Pdu
	isSet bool
}

func (v NullablePdu) Get() *Pdu {
	return v.value
}

func (v *NullablePdu) Set(val *Pdu) {
	v.value = val
	v.isSet = true
}

func (v NullablePdu) IsSet() bool {
	return v.isSet
}

func (v *NullablePdu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePdu(val *Pdu) *NullablePdu {
	return &NullablePdu{value: val, isSet: true}
}

func (v NullablePdu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePdu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


