/*
Dedicated Network Equipments

This is the description of the Dedicated Network Equipment API.  The base url of this API is `https://api.leaseweb.com`.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkequipment

import (
	"encoding/json"
)

// checks if the NetworkInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkInterface{}

// NetworkInterface struct for NetworkInterface
type NetworkInterface struct {
	// Represents a MAC Address in the standard colon delimited format. Eg. `01:23:45:67:89:0A`
	Mac NullableString `json:"mac,omitempty"`
	Ip NullableString `json:"ip,omitempty"`
	NullRouted NullableBool `json:"nullRouted,omitempty"`
	Gateway NullableString `json:"gateway,omitempty"`
	Ports []Port `json:"ports,omitempty"`
	LocationId NullableString `json:"locationId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkInterface NetworkInterface

// NewNetworkInterface instantiates a new NetworkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkInterface() *NetworkInterface {
	this := NetworkInterface{}
	return &this
}

// NewNetworkInterfaceWithDefaults instantiates a new NetworkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkInterfaceWithDefaults() *NetworkInterface {
	this := NetworkInterface{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkInterface) GetMac() string {
	if o == nil || IsNil(o.Mac.Get()) {
		var ret string
		return ret
	}
	return *o.Mac.Get()
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkInterface) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mac.Get(), o.Mac.IsSet()
}

// HasMac returns a boolean if a field has been set.
func (o *NetworkInterface) HasMac() bool {
	if o != nil && o.Mac.IsSet() {
		return true
	}

	return false
}

// SetMac gets a reference to the given NullableString and assigns it to the Mac field.
func (o *NetworkInterface) SetMac(v string) {
	o.Mac.Set(&v)
}
// SetMacNil sets the value for Mac to be an explicit nil
func (o *NetworkInterface) SetMacNil() {
	o.Mac.Set(nil)
}

// UnsetMac ensures that no value is present for Mac, not even an explicit nil
func (o *NetworkInterface) UnsetMac() {
	o.Mac.Unset()
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkInterface) GetIp() string {
	if o == nil || IsNil(o.Ip.Get()) {
		var ret string
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkInterface) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *NetworkInterface) HasIp() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableString and assigns it to the Ip field.
func (o *NetworkInterface) SetIp(v string) {
	o.Ip.Set(&v)
}
// SetIpNil sets the value for Ip to be an explicit nil
func (o *NetworkInterface) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *NetworkInterface) UnsetIp() {
	o.Ip.Unset()
}

// GetNullRouted returns the NullRouted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkInterface) GetNullRouted() bool {
	if o == nil || IsNil(o.NullRouted.Get()) {
		var ret bool
		return ret
	}
	return *o.NullRouted.Get()
}

// GetNullRoutedOk returns a tuple with the NullRouted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkInterface) GetNullRoutedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NullRouted.Get(), o.NullRouted.IsSet()
}

// HasNullRouted returns a boolean if a field has been set.
func (o *NetworkInterface) HasNullRouted() bool {
	if o != nil && o.NullRouted.IsSet() {
		return true
	}

	return false
}

// SetNullRouted gets a reference to the given NullableBool and assigns it to the NullRouted field.
func (o *NetworkInterface) SetNullRouted(v bool) {
	o.NullRouted.Set(&v)
}
// SetNullRoutedNil sets the value for NullRouted to be an explicit nil
func (o *NetworkInterface) SetNullRoutedNil() {
	o.NullRouted.Set(nil)
}

// UnsetNullRouted ensures that no value is present for NullRouted, not even an explicit nil
func (o *NetworkInterface) UnsetNullRouted() {
	o.NullRouted.Unset()
}

// GetGateway returns the Gateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkInterface) GetGateway() string {
	if o == nil || IsNil(o.Gateway.Get()) {
		var ret string
		return ret
	}
	return *o.Gateway.Get()
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkInterface) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gateway.Get(), o.Gateway.IsSet()
}

// HasGateway returns a boolean if a field has been set.
func (o *NetworkInterface) HasGateway() bool {
	if o != nil && o.Gateway.IsSet() {
		return true
	}

	return false
}

// SetGateway gets a reference to the given NullableString and assigns it to the Gateway field.
func (o *NetworkInterface) SetGateway(v string) {
	o.Gateway.Set(&v)
}
// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *NetworkInterface) SetGatewayNil() {
	o.Gateway.Set(nil)
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *NetworkInterface) UnsetGateway() {
	o.Gateway.Unset()
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *NetworkInterface) GetPorts() []Port {
	if o == nil || IsNil(o.Ports) {
		var ret []Port
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterface) GetPortsOk() ([]Port, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *NetworkInterface) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []Port and assigns it to the Ports field.
func (o *NetworkInterface) SetPorts(v []Port) {
	o.Ports = v
}

// GetLocationId returns the LocationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkInterface) GetLocationId() string {
	if o == nil || IsNil(o.LocationId.Get()) {
		var ret string
		return ret
	}
	return *o.LocationId.Get()
}

// GetLocationIdOk returns a tuple with the LocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkInterface) GetLocationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocationId.Get(), o.LocationId.IsSet()
}

// HasLocationId returns a boolean if a field has been set.
func (o *NetworkInterface) HasLocationId() bool {
	if o != nil && o.LocationId.IsSet() {
		return true
	}

	return false
}

// SetLocationId gets a reference to the given NullableString and assigns it to the LocationId field.
func (o *NetworkInterface) SetLocationId(v string) {
	o.LocationId.Set(&v)
}
// SetLocationIdNil sets the value for LocationId to be an explicit nil
func (o *NetworkInterface) SetLocationIdNil() {
	o.LocationId.Set(nil)
}

// UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
func (o *NetworkInterface) UnsetLocationId() {
	o.LocationId.Unset()
}

func (o NetworkInterface) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Mac.IsSet() {
		toSerialize["mac"] = o.Mac.Get()
	}
	if o.Ip.IsSet() {
		toSerialize["ip"] = o.Ip.Get()
	}
	if o.NullRouted.IsSet() {
		toSerialize["nullRouted"] = o.NullRouted.Get()
	}
	if o.Gateway.IsSet() {
		toSerialize["gateway"] = o.Gateway.Get()
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if o.LocationId.IsSet() {
		toSerialize["locationId"] = o.LocationId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkInterface) UnmarshalJSON(data []byte) (err error) {
	varNetworkInterface := _NetworkInterface{}

	err = json.Unmarshal(data, &varNetworkInterface)

	if err != nil {
		return err
	}

	*o = NetworkInterface(varNetworkInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mac")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "nullRouted")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "ports")
		delete(additionalProperties, "locationId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkInterface struct {
	value *NetworkInterface
	isSet bool
}

func (v NullableNetworkInterface) Get() *NetworkInterface {
	return v.value
}

func (v *NullableNetworkInterface) Set(val *NetworkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkInterface(val *NetworkInterface) *NullableNetworkInterface {
	return &NullableNetworkInterface{value: val, isSet: true}
}

func (v NullableNetworkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


