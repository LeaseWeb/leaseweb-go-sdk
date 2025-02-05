/*
Public Clouds

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
)

// checks if the CpuMetricsMetadataSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CpuMetricsMetadataSummary{}

// CpuMetricsMetadataSummary struct for CpuMetricsMetadataSummary
type CpuMetricsMetadataSummary struct {
	CpuMetrics *CpuMetricsMetadataSummaryCpuMetrics `json:"cpuMetrics,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CpuMetricsMetadataSummary CpuMetricsMetadataSummary

// NewCpuMetricsMetadataSummary instantiates a new CpuMetricsMetadataSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpuMetricsMetadataSummary() *CpuMetricsMetadataSummary {
	this := CpuMetricsMetadataSummary{}
	return &this
}

// NewCpuMetricsMetadataSummaryWithDefaults instantiates a new CpuMetricsMetadataSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpuMetricsMetadataSummaryWithDefaults() *CpuMetricsMetadataSummary {
	this := CpuMetricsMetadataSummary{}
	return &this
}

// GetCpuMetrics returns the CpuMetrics field value if set, zero value otherwise.
func (o *CpuMetricsMetadataSummary) GetCpuMetrics() CpuMetricsMetadataSummaryCpuMetrics {
	if o == nil || IsNil(o.CpuMetrics) {
		var ret CpuMetricsMetadataSummaryCpuMetrics
		return ret
	}
	return *o.CpuMetrics
}

// GetCpuMetricsOk returns a tuple with the CpuMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuMetricsMetadataSummary) GetCpuMetricsOk() (*CpuMetricsMetadataSummaryCpuMetrics, bool) {
	if o == nil || IsNil(o.CpuMetrics) {
		return nil, false
	}
	return o.CpuMetrics, true
}

// HasCpuMetrics returns a boolean if a field has been set.
func (o *CpuMetricsMetadataSummary) HasCpuMetrics() bool {
	if o != nil && !IsNil(o.CpuMetrics) {
		return true
	}

	return false
}

// SetCpuMetrics gets a reference to the given CpuMetricsMetadataSummaryCpuMetrics and assigns it to the CpuMetrics field.
func (o *CpuMetricsMetadataSummary) SetCpuMetrics(v CpuMetricsMetadataSummaryCpuMetrics) {
	o.CpuMetrics = &v
}

func (o CpuMetricsMetadataSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CpuMetricsMetadataSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CpuMetrics) {
		toSerialize["cpuMetrics"] = o.CpuMetrics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CpuMetricsMetadataSummary) UnmarshalJSON(data []byte) (err error) {
	varCpuMetricsMetadataSummary := _CpuMetricsMetadataSummary{}

	err = json.Unmarshal(data, &varCpuMetricsMetadataSummary)

	if err != nil {
		return err
	}

	*o = CpuMetricsMetadataSummary(varCpuMetricsMetadataSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cpuMetrics")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCpuMetricsMetadataSummary struct {
	value *CpuMetricsMetadataSummary
	isSet bool
}

func (v NullableCpuMetricsMetadataSummary) Get() *CpuMetricsMetadataSummary {
	return v.value
}

func (v *NullableCpuMetricsMetadataSummary) Set(val *CpuMetricsMetadataSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCpuMetricsMetadataSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCpuMetricsMetadataSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpuMetricsMetadataSummary(val *CpuMetricsMetadataSummary) *NullableCpuMetricsMetadataSummary {
	return &NullableCpuMetricsMetadataSummary{value: val, isSet: true}
}

func (v NullableCpuMetricsMetadataSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpuMetricsMetadataSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


