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

// checks if the GetDhcpReservationListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDhcpReservationListResult{}

// GetDhcpReservationListResult struct for GetDhcpReservationListResult
type GetDhcpReservationListResult struct {
	Metadata *Metadata `json:"_metadata,omitempty"`
	// An array of active DHCP reservations
	Leases []Lease `json:"leases,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDhcpReservationListResult GetDhcpReservationListResult

// NewGetDhcpReservationListResult instantiates a new GetDhcpReservationListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDhcpReservationListResult() *GetDhcpReservationListResult {
	this := GetDhcpReservationListResult{}
	return &this
}

// NewGetDhcpReservationListResultWithDefaults instantiates a new GetDhcpReservationListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDhcpReservationListResultWithDefaults() *GetDhcpReservationListResult {
	this := GetDhcpReservationListResult{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GetDhcpReservationListResult) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDhcpReservationListResult) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GetDhcpReservationListResult) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *GetDhcpReservationListResult) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetLeases returns the Leases field value if set, zero value otherwise.
func (o *GetDhcpReservationListResult) GetLeases() []Lease {
	if o == nil || IsNil(o.Leases) {
		var ret []Lease
		return ret
	}
	return o.Leases
}

// GetLeasesOk returns a tuple with the Leases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDhcpReservationListResult) GetLeasesOk() ([]Lease, bool) {
	if o == nil || IsNil(o.Leases) {
		return nil, false
	}
	return o.Leases, true
}

// HasLeases returns a boolean if a field has been set.
func (o *GetDhcpReservationListResult) HasLeases() bool {
	if o != nil && !IsNil(o.Leases) {
		return true
	}

	return false
}

// SetLeases gets a reference to the given []Lease and assigns it to the Leases field.
func (o *GetDhcpReservationListResult) SetLeases(v []Lease) {
	o.Leases = v
}

func (o GetDhcpReservationListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDhcpReservationListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["_metadata"] = o.Metadata
	}
	if !IsNil(o.Leases) {
		toSerialize["leases"] = o.Leases
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDhcpReservationListResult) UnmarshalJSON(data []byte) (err error) {
	varGetDhcpReservationListResult := _GetDhcpReservationListResult{}

	err = json.Unmarshal(data, &varGetDhcpReservationListResult)

	if err != nil {
		return err
	}

	*o = GetDhcpReservationListResult(varGetDhcpReservationListResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_metadata")
		delete(additionalProperties, "leases")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDhcpReservationListResult struct {
	value *GetDhcpReservationListResult
	isSet bool
}

func (v NullableGetDhcpReservationListResult) Get() *GetDhcpReservationListResult {
	return v.value
}

func (v *NullableGetDhcpReservationListResult) Set(val *GetDhcpReservationListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDhcpReservationListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDhcpReservationListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDhcpReservationListResult(val *GetDhcpReservationListResult) *NullableGetDhcpReservationListResult {
	return &NullableGetDhcpReservationListResult{value: val, isSet: true}
}

func (v NullableGetDhcpReservationListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDhcpReservationListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


