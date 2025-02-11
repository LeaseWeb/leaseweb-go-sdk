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

// checks if the GetNetworkEquipmentListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkEquipmentListResult{}

// GetNetworkEquipmentListResult struct for GetNetworkEquipmentListResult
type GetNetworkEquipmentListResult struct {
	Metadata *Metadata `json:"_metadata,omitempty"`
	// An array of network equipment
	NetworkEquipments []NetworkEquipment `json:"networkEquipments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetNetworkEquipmentListResult GetNetworkEquipmentListResult

// NewGetNetworkEquipmentListResult instantiates a new GetNetworkEquipmentListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkEquipmentListResult() *GetNetworkEquipmentListResult {
	this := GetNetworkEquipmentListResult{}
	return &this
}

// NewGetNetworkEquipmentListResultWithDefaults instantiates a new GetNetworkEquipmentListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkEquipmentListResultWithDefaults() *GetNetworkEquipmentListResult {
	this := GetNetworkEquipmentListResult{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetNetworkEquipmentListResult) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEquipmentListResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetNetworkEquipmentListResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *GetNetworkEquipmentListResult) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetNetworkEquipments returns the NetworkEquipments field value if set, zero value otherwise.
func (o *GetNetworkEquipmentListResult) GetNetworkEquipments() []NetworkEquipment {
	if o == nil || IsNil(o.NetworkEquipments) {
		var ret []NetworkEquipment
		return ret
	}
	return o.NetworkEquipments
}

// GetNetworkEquipmentsOk returns a tuple with the NetworkEquipments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEquipmentListResult) GetNetworkEquipmentsOk() ([]NetworkEquipment, bool) {
	if o == nil || IsNil(o.NetworkEquipments) {
		return nil, false
	}
	return o.NetworkEquipments, true
}

// HasNetworkEquipments returns a boolean if a field has been set.
func (o *GetNetworkEquipmentListResult) HasNetworkEquipments() bool {
	if o != nil && !IsNil(o.NetworkEquipments) {
		return true
	}

	return false
}

// SetNetworkEquipments gets a reference to the given []NetworkEquipment and assigns it to the NetworkEquipments field.
func (o *GetNetworkEquipmentListResult) SetNetworkEquipments(v []NetworkEquipment) {
	o.NetworkEquipments = v
}

func (o GetNetworkEquipmentListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkEquipmentListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}
	if !IsNil(o.NetworkEquipments) {
		toSerialize["networkEquipments"] = o.NetworkEquipments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetNetworkEquipmentListResult) UnmarshalJSON(data []byte) (err error) {
	varGetNetworkEquipmentListResult := _GetNetworkEquipmentListResult{}

	err = json.Unmarshal(data, &varGetNetworkEquipmentListResult)

	if err != nil {
		return err
	}

	*o = GetNetworkEquipmentListResult(varGetNetworkEquipmentListResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_metadata")
		delete(additionalProperties, "networkEquipments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetNetworkEquipmentListResult struct {
	value *GetNetworkEquipmentListResult
	isSet bool
}

func (v NullableGetNetworkEquipmentListResult) Get() *GetNetworkEquipmentListResult {
	return v.value
}

func (v *NullableGetNetworkEquipmentListResult) Set(val *GetNetworkEquipmentListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkEquipmentListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkEquipmentListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkEquipmentListResult(val *GetNetworkEquipmentListResult) *NullableGetNetworkEquipmentListResult {
	return &NullableGetNetworkEquipmentListResult{value: val, isSet: true}
}

func (v NullableGetNetworkEquipmentListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkEquipmentListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


