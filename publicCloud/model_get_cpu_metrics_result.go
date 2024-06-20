/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
)

// checks if the GetCpuMetricsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCpuMetricsResult{}

// GetCpuMetricsResult struct for GetCpuMetricsResult
type GetCpuMetricsResult struct {
	Metrics *CpuMetrics `json:"metrics,omitempty"`
	Metadata *CpuMetricsMetadata `json:"_metadata,omitempty"`
}

// NewGetCpuMetricsResult instantiates a new GetCpuMetricsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCpuMetricsResult() *GetCpuMetricsResult {
	this := GetCpuMetricsResult{}
	return &this
}

// NewGetCpuMetricsResultWithDefaults instantiates a new GetCpuMetricsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCpuMetricsResultWithDefaults() *GetCpuMetricsResult {
	this := GetCpuMetricsResult{}
	return &this
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *GetCpuMetricsResult) GetMetrics() CpuMetrics {
	if o == nil || IsNil(o.Metrics) {
		var ret CpuMetrics
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCpuMetricsResult) GetMetricsOk() (*CpuMetrics, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *GetCpuMetricsResult) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given CpuMetrics and assigns it to the Metrics field.
func (o *GetCpuMetricsResult) SetMetrics(v CpuMetrics) {
	o.Metrics = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetCpuMetricsResult) GetMetadata() CpuMetricsMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret CpuMetricsMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCpuMetricsResult) GetMetadataOk() (*CpuMetricsMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetCpuMetricsResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given CpuMetricsMetadata and assigns it to the Metadata field.
func (o *GetCpuMetricsResult) SetMetadata(v CpuMetricsMetadata) {
	o.Metadata = &v
}

func (o GetCpuMetricsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCpuMetricsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableGetCpuMetricsResult struct {
	value *GetCpuMetricsResult
	isSet bool
}

func (v NullableGetCpuMetricsResult) Get() *GetCpuMetricsResult {
	return v.value
}

func (v *NullableGetCpuMetricsResult) Set(val *GetCpuMetricsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCpuMetricsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCpuMetricsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCpuMetricsResult(val *GetCpuMetricsResult) *NullableGetCpuMetricsResult {
	return &NullableGetCpuMetricsResult{value: val, isSet: true}
}

func (v NullableGetCpuMetricsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCpuMetricsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


