/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
	"time"
)

// checks if the CpuMetricsMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CpuMetricsMetadata{}

// CpuMetricsMetadata struct for CpuMetricsMetadata
type CpuMetricsMetadata struct {
	From *time.Time `json:"from,omitempty"`
	To *time.Time `json:"to,omitempty"`
	Granularity *MetricsGranularity `json:"granularity,omitempty"`
	Summary *CpuMetricsMetadataSummary `json:"summary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CpuMetricsMetadata CpuMetricsMetadata

// NewCpuMetricsMetadata instantiates a new CpuMetricsMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpuMetricsMetadata() *CpuMetricsMetadata {
	this := CpuMetricsMetadata{}
	return &this
}

// NewCpuMetricsMetadataWithDefaults instantiates a new CpuMetricsMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpuMetricsMetadataWithDefaults() *CpuMetricsMetadata {
	this := CpuMetricsMetadata{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *CpuMetricsMetadata) GetFrom() time.Time {
	if o == nil || IsNil(o.From) {
		var ret time.Time
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuMetricsMetadata) GetFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *CpuMetricsMetadata) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given time.Time and assigns it to the From field.
func (o *CpuMetricsMetadata) SetFrom(v time.Time) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *CpuMetricsMetadata) GetTo() time.Time {
	if o == nil || IsNil(o.To) {
		var ret time.Time
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuMetricsMetadata) GetToOk() (*time.Time, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *CpuMetricsMetadata) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given time.Time and assigns it to the To field.
func (o *CpuMetricsMetadata) SetTo(v time.Time) {
	o.To = &v
}

// GetGranularity returns the Granularity field value if set, zero value otherwise.
func (o *CpuMetricsMetadata) GetGranularity() MetricsGranularity {
	if o == nil || IsNil(o.Granularity) {
		var ret MetricsGranularity
		return ret
	}
	return *o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuMetricsMetadata) GetGranularityOk() (*MetricsGranularity, bool) {
	if o == nil || IsNil(o.Granularity) {
		return nil, false
	}
	return o.Granularity, true
}

// HasGranularity returns a boolean if a field has been set.
func (o *CpuMetricsMetadata) HasGranularity() bool {
	if o != nil && !IsNil(o.Granularity) {
		return true
	}

	return false
}

// SetGranularity gets a reference to the given MetricsGranularity and assigns it to the Granularity field.
func (o *CpuMetricsMetadata) SetGranularity(v MetricsGranularity) {
	o.Granularity = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *CpuMetricsMetadata) GetSummary() CpuMetricsMetadataSummary {
	if o == nil || IsNil(o.Summary) {
		var ret CpuMetricsMetadataSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuMetricsMetadata) GetSummaryOk() (*CpuMetricsMetadataSummary, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *CpuMetricsMetadata) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given CpuMetricsMetadataSummary and assigns it to the Summary field.
func (o *CpuMetricsMetadata) SetSummary(v CpuMetricsMetadataSummary) {
	o.Summary = &v
}

func (o CpuMetricsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CpuMetricsMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.Granularity) {
		toSerialize["granularity"] = o.Granularity
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CpuMetricsMetadata) UnmarshalJSON(data []byte) (err error) {
	varCpuMetricsMetadata := _CpuMetricsMetadata{}

	err = json.Unmarshal(data, &varCpuMetricsMetadata)

	if err != nil {
		return err
	}

	*o = CpuMetricsMetadata(varCpuMetricsMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		delete(additionalProperties, "granularity")
		delete(additionalProperties, "summary")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCpuMetricsMetadata struct {
	value *CpuMetricsMetadata
	isSet bool
}

func (v NullableCpuMetricsMetadata) Get() *CpuMetricsMetadata {
	return v.value
}

func (v *NullableCpuMetricsMetadata) Set(val *CpuMetricsMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableCpuMetricsMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableCpuMetricsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpuMetricsMetadata(val *CpuMetricsMetadata) *NullableCpuMetricsMetadata {
	return &NullableCpuMetricsMetadata{value: val, isSet: true}
}

func (v NullableCpuMetricsMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpuMetricsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


