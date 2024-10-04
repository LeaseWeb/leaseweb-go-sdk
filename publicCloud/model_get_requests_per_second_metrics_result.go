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

// checks if the GetRequestsPerSecondMetricsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRequestsPerSecondMetricsResult{}

// GetRequestsPerSecondMetricsResult struct for GetRequestsPerSecondMetricsResult
type GetRequestsPerSecondMetricsResult struct {
	Metrics MetricsProperties `json:"metrics,omitempty"`
	Metadata MetricsMetadataProperties `json:"_metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRequestsPerSecondMetricsResult GetRequestsPerSecondMetricsResult

// NewGetRequestsPerSecondMetricsResult instantiates a new GetRequestsPerSecondMetricsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRequestsPerSecondMetricsResult() *GetRequestsPerSecondMetricsResult {
	this := GetRequestsPerSecondMetricsResult{}
	return &this
}

// NewGetRequestsPerSecondMetricsResultWithDefaults instantiates a new GetRequestsPerSecondMetricsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRequestsPerSecondMetricsResultWithDefaults() *GetRequestsPerSecondMetricsResult {
	this := GetRequestsPerSecondMetricsResult{}
	return &this
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *GetRequestsPerSecondMetricsResult) GetMetrics() MetricsProperties {
	if o == nil || IsNil(o.Metrics) {
		var ret MetricsProperties
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRequestsPerSecondMetricsResult) GetMetricsOk() (MetricsProperties, bool) {
	if o == nil || IsNil(o.Metrics) {
		return MetricsProperties{}, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *GetRequestsPerSecondMetricsResult) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given MetricsProperties and assigns it to the Metrics field.
func (o *GetRequestsPerSecondMetricsResult) SetMetrics(v MetricsProperties) {
	o.Metrics = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetRequestsPerSecondMetricsResult) GetMetadata() MetricsMetadataProperties {
	if o == nil || IsNil(o.Metadata) {
		var ret MetricsMetadataProperties
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRequestsPerSecondMetricsResult) GetMetadataOk() (MetricsMetadataProperties, bool) {
	if o == nil || IsNil(o.Metadata) {
		return MetricsMetadataProperties{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetRequestsPerSecondMetricsResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given MetricsMetadataProperties and assigns it to the Metadata field.
func (o *GetRequestsPerSecondMetricsResult) SetMetadata(v MetricsMetadataProperties) {
	o.Metadata = v
}

func (o GetRequestsPerSecondMetricsResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRequestsPerSecondMetricsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetRequestsPerSecondMetricsResult) UnmarshalJSON(data []byte) (err error) {
	varGetRequestsPerSecondMetricsResult := _GetRequestsPerSecondMetricsResult{}

	err = json.Unmarshal(data, &varGetRequestsPerSecondMetricsResult)

	if err != nil {
		return err
	}

	*o = GetRequestsPerSecondMetricsResult(varGetRequestsPerSecondMetricsResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "metrics")
		delete(additionalProperties, "_metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRequestsPerSecondMetricsResult struct {
	value *GetRequestsPerSecondMetricsResult
	isSet bool
}

func (v NullableGetRequestsPerSecondMetricsResult) Get() *GetRequestsPerSecondMetricsResult {
	return v.value
}

func (v *NullableGetRequestsPerSecondMetricsResult) Set(val *GetRequestsPerSecondMetricsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRequestsPerSecondMetricsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRequestsPerSecondMetricsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRequestsPerSecondMetricsResult(val *GetRequestsPerSecondMetricsResult) *NullableGetRequestsPerSecondMetricsResult {
	return &NullableGetRequestsPerSecondMetricsResult{value: val, isSet: true}
}

func (v NullableGetRequestsPerSecondMetricsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRequestsPerSecondMetricsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


