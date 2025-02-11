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

// checks if the GetDataTrafficNotificationSettingListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDataTrafficNotificationSettingListResult{}

// GetDataTrafficNotificationSettingListResult struct for GetDataTrafficNotificationSettingListResult
type GetDataTrafficNotificationSettingListResult struct {
	Metadata *Metadata `json:"_metadata,omitempty"`
	// An array of Data traffic Notification Settings
	DatatrafficNotificationSettings []DataTrafficNotificationSetting `json:"datatrafficNotificationSettings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDataTrafficNotificationSettingListResult GetDataTrafficNotificationSettingListResult

// NewGetDataTrafficNotificationSettingListResult instantiates a new GetDataTrafficNotificationSettingListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDataTrafficNotificationSettingListResult() *GetDataTrafficNotificationSettingListResult {
	this := GetDataTrafficNotificationSettingListResult{}
	return &this
}

// NewGetDataTrafficNotificationSettingListResultWithDefaults instantiates a new GetDataTrafficNotificationSettingListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDataTrafficNotificationSettingListResultWithDefaults() *GetDataTrafficNotificationSettingListResult {
	this := GetDataTrafficNotificationSettingListResult{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetDataTrafficNotificationSettingListResult) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDataTrafficNotificationSettingListResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetDataTrafficNotificationSettingListResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *GetDataTrafficNotificationSettingListResult) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetDatatrafficNotificationSettings returns the DatatrafficNotificationSettings field value if set, zero value otherwise.
func (o *GetDataTrafficNotificationSettingListResult) GetDatatrafficNotificationSettings() []DataTrafficNotificationSetting {
	if o == nil || IsNil(o.DatatrafficNotificationSettings) {
		var ret []DataTrafficNotificationSetting
		return ret
	}
	return o.DatatrafficNotificationSettings
}

// GetDatatrafficNotificationSettingsOk returns a tuple with the DatatrafficNotificationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDataTrafficNotificationSettingListResult) GetDatatrafficNotificationSettingsOk() ([]DataTrafficNotificationSetting, bool) {
	if o == nil || IsNil(o.DatatrafficNotificationSettings) {
		return nil, false
	}
	return o.DatatrafficNotificationSettings, true
}

// HasDatatrafficNotificationSettings returns a boolean if a field has been set.
func (o *GetDataTrafficNotificationSettingListResult) HasDatatrafficNotificationSettings() bool {
	if o != nil && !IsNil(o.DatatrafficNotificationSettings) {
		return true
	}

	return false
}

// SetDatatrafficNotificationSettings gets a reference to the given []DataTrafficNotificationSetting and assigns it to the DatatrafficNotificationSettings field.
func (o *GetDataTrafficNotificationSettingListResult) SetDatatrafficNotificationSettings(v []DataTrafficNotificationSetting) {
	o.DatatrafficNotificationSettings = v
}

func (o GetDataTrafficNotificationSettingListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDataTrafficNotificationSettingListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}
	if !IsNil(o.DatatrafficNotificationSettings) {
		toSerialize["datatrafficNotificationSettings"] = o.DatatrafficNotificationSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDataTrafficNotificationSettingListResult) UnmarshalJSON(data []byte) (err error) {
	varGetDataTrafficNotificationSettingListResult := _GetDataTrafficNotificationSettingListResult{}

	err = json.Unmarshal(data, &varGetDataTrafficNotificationSettingListResult)

	if err != nil {
		return err
	}

	*o = GetDataTrafficNotificationSettingListResult(varGetDataTrafficNotificationSettingListResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_metadata")
		delete(additionalProperties, "datatrafficNotificationSettings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDataTrafficNotificationSettingListResult struct {
	value *GetDataTrafficNotificationSettingListResult
	isSet bool
}

func (v NullableGetDataTrafficNotificationSettingListResult) Get() *GetDataTrafficNotificationSettingListResult {
	return v.value
}

func (v *NullableGetDataTrafficNotificationSettingListResult) Set(val *GetDataTrafficNotificationSettingListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDataTrafficNotificationSettingListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDataTrafficNotificationSettingListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDataTrafficNotificationSettingListResult(val *GetDataTrafficNotificationSettingListResult) *NullableGetDataTrafficNotificationSettingListResult {
	return &NullableGetDataTrafficNotificationSettingListResult{value: val, isSet: true}
}

func (v NullableGetDataTrafficNotificationSettingListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDataTrafficNotificationSettingListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


