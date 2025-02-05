/*
Aggregation packs

This documents the rest api aggregation packs provides.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aggregationpack

import (
	"encoding/json"
)

// checks if the GetAggregationPackListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAggregationPackListResult{}

// GetAggregationPackListResult struct for GetAggregationPackListResult
type GetAggregationPackListResult struct {
	Metadata *Metadata `json:"_metadata,omitempty"`
	// An array of aggregation packs
	ConnectContractItems []AggregationPack `json:"connectContractItems,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetAggregationPackListResult GetAggregationPackListResult

// NewGetAggregationPackListResult instantiates a new GetAggregationPackListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAggregationPackListResult() *GetAggregationPackListResult {
	this := GetAggregationPackListResult{}
	return &this
}

// NewGetAggregationPackListResultWithDefaults instantiates a new GetAggregationPackListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAggregationPackListResultWithDefaults() *GetAggregationPackListResult {
	this := GetAggregationPackListResult{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetAggregationPackListResult) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAggregationPackListResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetAggregationPackListResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *GetAggregationPackListResult) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetConnectContractItems returns the ConnectContractItems field value if set, zero value otherwise.
func (o *GetAggregationPackListResult) GetConnectContractItems() []AggregationPack {
	if o == nil || IsNil(o.ConnectContractItems) {
		var ret []AggregationPack
		return ret
	}
	return o.ConnectContractItems
}

// GetConnectContractItemsOk returns a tuple with the ConnectContractItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAggregationPackListResult) GetConnectContractItemsOk() ([]AggregationPack, bool) {
	if o == nil || IsNil(o.ConnectContractItems) {
		return nil, false
	}
	return o.ConnectContractItems, true
}

// HasConnectContractItems returns a boolean if a field has been set.
func (o *GetAggregationPackListResult) HasConnectContractItems() bool {
	if o != nil && !IsNil(o.ConnectContractItems) {
		return true
	}

	return false
}

// SetConnectContractItems gets a reference to the given []AggregationPack and assigns it to the ConnectContractItems field.
func (o *GetAggregationPackListResult) SetConnectContractItems(v []AggregationPack) {
	o.ConnectContractItems = v
}

func (o GetAggregationPackListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAggregationPackListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}
	if !IsNil(o.ConnectContractItems) {
		toSerialize["connectContractItems"] = o.ConnectContractItems
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAggregationPackListResult) UnmarshalJSON(data []byte) (err error) {
	varGetAggregationPackListResult := _GetAggregationPackListResult{}

	err = json.Unmarshal(data, &varGetAggregationPackListResult)

	if err != nil {
		return err
	}

	*o = GetAggregationPackListResult(varGetAggregationPackListResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_metadata")
		delete(additionalProperties, "connectContractItems")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAggregationPackListResult struct {
	value *GetAggregationPackListResult
	isSet bool
}

func (v NullableGetAggregationPackListResult) Get() *GetAggregationPackListResult {
	return v.value
}

func (v *NullableGetAggregationPackListResult) Set(val *GetAggregationPackListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAggregationPackListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAggregationPackListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAggregationPackListResult(val *GetAggregationPackListResult) *NullableGetAggregationPackListResult {
	return &NullableGetAggregationPackListResult{value: val, isSet: true}
}

func (v NullableGetAggregationPackListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAggregationPackListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


