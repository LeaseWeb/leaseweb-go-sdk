/*
IP management

> The base URL for this API is: **https://api.leaseweb.com/ipMgmt/v2/_**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipmgmt

import (
	"encoding/json"
	"fmt"
)

// checks if the GetNullRouteHistoryListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNullRouteHistoryListResult{}

// GetNullRouteHistoryListResult struct for GetNullRouteHistoryListResult
type GetNullRouteHistoryListResult struct {
	Nullroutes []NullRoutedIP `json:"nullroutes"`
	Metadata Metadata `json:"_metadata"`
	AdditionalProperties map[string]interface{}
}

type _GetNullRouteHistoryListResult GetNullRouteHistoryListResult

// NewGetNullRouteHistoryListResult instantiates a new GetNullRouteHistoryListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNullRouteHistoryListResult(nullroutes []NullRoutedIP, metadata Metadata) *GetNullRouteHistoryListResult {
	this := GetNullRouteHistoryListResult{}
	this.Nullroutes = nullroutes
	this.Metadata = metadata
	return &this
}

// NewGetNullRouteHistoryListResultWithDefaults instantiates a new GetNullRouteHistoryListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNullRouteHistoryListResultWithDefaults() *GetNullRouteHistoryListResult {
	this := GetNullRouteHistoryListResult{}
	return &this
}

// GetNullroutes returns the Nullroutes field value
func (o *GetNullRouteHistoryListResult) GetNullroutes() []NullRoutedIP {
	if o == nil {
		var ret []NullRoutedIP
		return ret
	}

	return o.Nullroutes
}

// GetNullroutesOk returns a tuple with the Nullroutes field value
// and a boolean to check if the value has been set.
func (o *GetNullRouteHistoryListResult) GetNullroutesOk() ([]NullRoutedIP, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nullroutes, true
}

// SetNullroutes sets field value
func (o *GetNullRouteHistoryListResult) SetNullroutes(v []NullRoutedIP) {
	o.Nullroutes = v
}

// GetMetadata returns the Metadata field value
func (o *GetNullRouteHistoryListResult) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *GetNullRouteHistoryListResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *GetNullRouteHistoryListResult) SetMetadata(v Metadata) {
	o.Metadata = v
}

func (o GetNullRouteHistoryListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNullRouteHistoryListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nullroutes"] = o.Nullroutes
	toSerialize["_metadata"] = o.Metadata

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetNullRouteHistoryListResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nullroutes",
		"_metadata",
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

	varGetNullRouteHistoryListResult := _GetNullRouteHistoryListResult{}

	err = json.Unmarshal(data, &varGetNullRouteHistoryListResult)

	if err != nil {
		return err
	}

	*o = GetNullRouteHistoryListResult(varGetNullRouteHistoryListResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "nullroutes")
		delete(additionalProperties, "_metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetNullRouteHistoryListResult struct {
	value *GetNullRouteHistoryListResult
	isSet bool
}

func (v NullableGetNullRouteHistoryListResult) Get() *GetNullRouteHistoryListResult {
	return v.value
}

func (v *NullableGetNullRouteHistoryListResult) Set(val *GetNullRouteHistoryListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNullRouteHistoryListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNullRouteHistoryListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNullRouteHistoryListResult(val *GetNullRouteHistoryListResult) *NullableGetNullRouteHistoryListResult {
	return &NullableGetNullRouteHistoryListResult{value: val, isSet: true}
}

func (v NullableGetNullRouteHistoryListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNullRouteHistoryListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


