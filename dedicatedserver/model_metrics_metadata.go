/*
Dedicated Servers

This is the description of the Dedicated Server API.  The base url of this API is `https://api.leaseweb.com`.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedserver

import (
	"encoding/json"
	"time"
)

// checks if the MetricsMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsMetadata{}

// MetricsMetadata Metadata about the collection
type MetricsMetadata struct {
	// The aggregation type for this request
	Aggregation *string `json:"aggregation,omitempty"`
	// Start of date interval in ISO-8601 format
	From *time.Time `json:"from,omitempty"`
	// The interval for each metric
	Granularity *string `json:"granularity,omitempty"`
	// End of date interval in ISO-8601 format
	To *time.Time `json:"to,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetricsMetadata MetricsMetadata

// NewMetricsMetadata instantiates a new MetricsMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsMetadata() *MetricsMetadata {
	this := MetricsMetadata{}
	return &this
}

// NewMetricsMetadataWithDefaults instantiates a new MetricsMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsMetadataWithDefaults() *MetricsMetadata {
	this := MetricsMetadata{}
	return &this
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *MetricsMetadata) GetAggregation() string {
	if o == nil || IsNil(o.Aggregation) {
		var ret string
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetadata) GetAggregationOk() (*string, bool) {
	if o == nil || IsNil(o.Aggregation) {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *MetricsMetadata) HasAggregation() bool {
	if o != nil && !IsNil(o.Aggregation) {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given string and assigns it to the Aggregation field.
func (o *MetricsMetadata) SetAggregation(v string) {
	o.Aggregation = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *MetricsMetadata) GetFrom() time.Time {
	if o == nil || IsNil(o.From) {
		var ret time.Time
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetadata) GetFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *MetricsMetadata) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given time.Time and assigns it to the From field.
func (o *MetricsMetadata) SetFrom(v time.Time) {
	o.From = &v
}

// GetGranularity returns the Granularity field value if set, zero value otherwise.
func (o *MetricsMetadata) GetGranularity() string {
	if o == nil || IsNil(o.Granularity) {
		var ret string
		return ret
	}
	return *o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetadata) GetGranularityOk() (*string, bool) {
	if o == nil || IsNil(o.Granularity) {
		return nil, false
	}
	return o.Granularity, true
}

// HasGranularity returns a boolean if a field has been set.
func (o *MetricsMetadata) HasGranularity() bool {
	if o != nil && !IsNil(o.Granularity) {
		return true
	}

	return false
}

// SetGranularity gets a reference to the given string and assigns it to the Granularity field.
func (o *MetricsMetadata) SetGranularity(v string) {
	o.Granularity = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *MetricsMetadata) GetTo() time.Time {
	if o == nil || IsNil(o.To) {
		var ret time.Time
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsMetadata) GetToOk() (*time.Time, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *MetricsMetadata) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given time.Time and assigns it to the To field.
func (o *MetricsMetadata) SetTo(v time.Time) {
	o.To = &v
}

func (o MetricsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Aggregation) {
		toSerialize["aggregation"] = o.Aggregation
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Granularity) {
		toSerialize["granularity"] = o.Granularity
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MetricsMetadata) UnmarshalJSON(data []byte) (err error) {
	varMetricsMetadata := _MetricsMetadata{}

	err = json.Unmarshal(data, &varMetricsMetadata)

	if err != nil {
		return err
	}

	*o = MetricsMetadata(varMetricsMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "aggregation")
		delete(additionalProperties, "from")
		delete(additionalProperties, "granularity")
		delete(additionalProperties, "to")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetricsMetadata struct {
	value *MetricsMetadata
	isSet bool
}

func (v NullableMetricsMetadata) Get() *MetricsMetadata {
	return v.value
}

func (v *NullableMetricsMetadata) Set(val *MetricsMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsMetadata(val *MetricsMetadata) *NullableMetricsMetadata {
	return &NullableMetricsMetadata{value: val, isSet: true}
}

func (v NullableMetricsMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


