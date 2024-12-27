/*
LeaseWeb API for IP address management

> The base URL for this API is: **https://api.leaseweb.com/ipMgmt/v2/_**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ips

import (
	"encoding/json"
	"fmt"
)

// checks if the Ip type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ip{}

// Ip struct for Ip
type Ip struct {
	// IP address
	Ip string `json:"ip"`
	Version ProtocolVersion `json:"version"`
	Type IpType `json:"type"`
	// Prefix length of the IP range represented by the record. Note: this is not the same as `subnet.prefixLength`.
	PrefixLength int32 `json:"prefixLength"`
	// Boolean indicating if this is the primary IP of the assigned equipment
	Primary bool `json:"primary"`
	// Reverse lookup set for the IP. This only applies to IPv4. For IPv6 see [GET /ips/{ip}/reverseLookup](#operation/get/ips/{ip}/reverseLookup).
	ReverseLookup NullableString `json:"reverseLookup"`
	// Boolean to indicate if the IP is null-routed
	NullRouted bool `json:"nullRouted"`
	// Null route level
	NullLevel NullableInt32 `json:"nullLevel"`
	// Boolean indicating if the null route can be removed
	UnnullingAllowed bool `json:"unnullingAllowed"`
	// ID of the equipment using the IP
	EquipmentId string `json:"equipmentId"`
	AssignedContract NullableAssignedContract `json:"assignedContract"`
	Subnet Subnet `json:"subnet"`
	AdditionalProperties map[string]interface{}
}

type _Ip Ip

// NewIp instantiates a new Ip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIp(ip string, version ProtocolVersion, type_ IpType, prefixLength int32, primary bool, reverseLookup NullableString, nullRouted bool, nullLevel NullableInt32, unnullingAllowed bool, equipmentId string, assignedContract NullableAssignedContract, subnet Subnet) *Ip {
	this := Ip{}
	this.Ip = ip
	this.Version = version
	this.Type = type_
	this.PrefixLength = prefixLength
	this.Primary = primary
	this.ReverseLookup = reverseLookup
	this.NullRouted = nullRouted
	this.NullLevel = nullLevel
	this.UnnullingAllowed = unnullingAllowed
	this.EquipmentId = equipmentId
	this.AssignedContract = assignedContract
	this.Subnet = subnet
	return &this
}

// NewIpWithDefaults instantiates a new Ip object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpWithDefaults() *Ip {
	this := Ip{}
	var prefixLength int32 = 0
	this.PrefixLength = prefixLength
	var nullRouted bool = false
	this.NullRouted = nullRouted
	var unnullingAllowed bool = false
	this.UnnullingAllowed = unnullingAllowed
	return &this
}

// GetIp returns the Ip field value
func (o *Ip) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *Ip) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *Ip) SetIp(v string) {
	o.Ip = v
}

// GetVersion returns the Version field value
func (o *Ip) GetVersion() ProtocolVersion {
	if o == nil {
		var ret ProtocolVersion
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Ip) GetVersionOk() (*ProtocolVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Ip) SetVersion(v ProtocolVersion) {
	o.Version = v
}

// GetType returns the Type field value
func (o *Ip) GetType() IpType {
	if o == nil {
		var ret IpType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Ip) GetTypeOk() (*IpType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Ip) SetType(v IpType) {
	o.Type = v
}

// GetPrefixLength returns the PrefixLength field value
func (o *Ip) GetPrefixLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrefixLength
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value
// and a boolean to check if the value has been set.
func (o *Ip) GetPrefixLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixLength, true
}

// SetPrefixLength sets field value
func (o *Ip) SetPrefixLength(v int32) {
	o.PrefixLength = v
}

// GetPrimary returns the Primary field value
func (o *Ip) GetPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value
// and a boolean to check if the value has been set.
func (o *Ip) GetPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Primary, true
}

// SetPrimary sets field value
func (o *Ip) SetPrimary(v bool) {
	o.Primary = v
}

// GetReverseLookup returns the ReverseLookup field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Ip) GetReverseLookup() string {
	if o == nil || o.ReverseLookup.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReverseLookup.Get()
}

// GetReverseLookupOk returns a tuple with the ReverseLookup field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ip) GetReverseLookupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReverseLookup.Get(), o.ReverseLookup.IsSet()
}

// SetReverseLookup sets field value
func (o *Ip) SetReverseLookup(v string) {
	o.ReverseLookup.Set(&v)
}

// GetNullRouted returns the NullRouted field value
func (o *Ip) GetNullRouted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NullRouted
}

// GetNullRoutedOk returns a tuple with the NullRouted field value
// and a boolean to check if the value has been set.
func (o *Ip) GetNullRoutedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NullRouted, true
}

// SetNullRouted sets field value
func (o *Ip) SetNullRouted(v bool) {
	o.NullRouted = v
}

// GetNullLevel returns the NullLevel field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Ip) GetNullLevel() int32 {
	if o == nil || o.NullLevel.Get() == nil {
		var ret int32
		return ret
	}

	return *o.NullLevel.Get()
}

// GetNullLevelOk returns a tuple with the NullLevel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ip) GetNullLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NullLevel.Get(), o.NullLevel.IsSet()
}

// SetNullLevel sets field value
func (o *Ip) SetNullLevel(v int32) {
	o.NullLevel.Set(&v)
}

// GetUnnullingAllowed returns the UnnullingAllowed field value
func (o *Ip) GetUnnullingAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UnnullingAllowed
}

// GetUnnullingAllowedOk returns a tuple with the UnnullingAllowed field value
// and a boolean to check if the value has been set.
func (o *Ip) GetUnnullingAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnnullingAllowed, true
}

// SetUnnullingAllowed sets field value
func (o *Ip) SetUnnullingAllowed(v bool) {
	o.UnnullingAllowed = v
}

// GetEquipmentId returns the EquipmentId field value
func (o *Ip) GetEquipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EquipmentId
}

// GetEquipmentIdOk returns a tuple with the EquipmentId field value
// and a boolean to check if the value has been set.
func (o *Ip) GetEquipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EquipmentId, true
}

// SetEquipmentId sets field value
func (o *Ip) SetEquipmentId(v string) {
	o.EquipmentId = v
}

// GetAssignedContract returns the AssignedContract field value
// If the value is explicit nil, the zero value for AssignedContract will be returned
func (o *Ip) GetAssignedContract() AssignedContract {
	if o == nil || o.AssignedContract.Get() == nil {
		var ret AssignedContract
		return ret
	}

	return *o.AssignedContract.Get()
}

// GetAssignedContractOk returns a tuple with the AssignedContract field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ip) GetAssignedContractOk() (*AssignedContract, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedContract.Get(), o.AssignedContract.IsSet()
}

// SetAssignedContract sets field value
func (o *Ip) SetAssignedContract(v AssignedContract) {
	o.AssignedContract.Set(&v)
}

// GetSubnet returns the Subnet field value
func (o *Ip) GetSubnet() Subnet {
	if o == nil {
		var ret Subnet
		return ret
	}

	return o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value
// and a boolean to check if the value has been set.
func (o *Ip) GetSubnetOk() (*Subnet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subnet, true
}

// SetSubnet sets field value
func (o *Ip) SetSubnet(v Subnet) {
	o.Subnet = v
}

func (o Ip) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ip) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip"] = o.Ip
	toSerialize["version"] = o.Version
	toSerialize["type"] = o.Type
	toSerialize["prefixLength"] = o.PrefixLength
	toSerialize["primary"] = o.Primary
	toSerialize["reverseLookup"] = o.ReverseLookup.Get()
	toSerialize["nullRouted"] = o.NullRouted
	toSerialize["nullLevel"] = o.NullLevel.Get()
	toSerialize["unnullingAllowed"] = o.UnnullingAllowed
	toSerialize["equipmentId"] = o.EquipmentId
	toSerialize["assignedContract"] = o.AssignedContract.Get()
	toSerialize["subnet"] = o.Subnet

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Ip) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ip",
		"version",
		"type",
		"prefixLength",
		"primary",
		"reverseLookup",
		"nullRouted",
		"nullLevel",
		"unnullingAllowed",
		"equipmentId",
		"assignedContract",
		"subnet",
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

	varIp := _Ip{}

	err = json.Unmarshal(data, &varIp)

	if err != nil {
		return err
	}

	*o = Ip(varIp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ip")
		delete(additionalProperties, "version")
		delete(additionalProperties, "type")
		delete(additionalProperties, "prefixLength")
		delete(additionalProperties, "primary")
		delete(additionalProperties, "reverseLookup")
		delete(additionalProperties, "nullRouted")
		delete(additionalProperties, "nullLevel")
		delete(additionalProperties, "unnullingAllowed")
		delete(additionalProperties, "equipmentId")
		delete(additionalProperties, "assignedContract")
		delete(additionalProperties, "subnet")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIp struct {
	value *Ip
	isSet bool
}

func (v NullableIp) Get() *Ip {
	return v.value
}

func (v *NullableIp) Set(val *Ip) {
	v.value = val
	v.isSet = true
}

func (v NullableIp) IsSet() bool {
	return v.isSet
}

func (v *NullableIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIp(val *Ip) *NullableIp {
	return &NullableIp{value: val, isSet: true}
}

func (v NullableIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


