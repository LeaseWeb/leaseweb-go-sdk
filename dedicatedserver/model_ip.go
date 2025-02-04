/*
Dedicated Servers API

This is the description of the Dedicated Server API.  The base url of this API is `https://api.leaseweb.com`.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dedicatedserver

import (
	"encoding/json"
)

// checks if the Ip type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ip{}

// Ip struct for Ip
type Ip struct {
	Ddos *DDos `json:"ddos,omitempty"`
	// Indicates if the IP is a Floating IP
	FloatingIp *bool `json:"floatingIp,omitempty"`
	// Gateway
	Gateway *string `json:"gateway,omitempty"`
	// IP address in CIDR notation
	Ip *string `json:"ip,omitempty"`
	// IP address is main
	MainIp *bool `json:"mainIp,omitempty"`
	NetworkType *NetworkType `json:"networkType,omitempty"`
	// IP address null routed
	NullRouted *bool `json:"nullRouted,omitempty"`
	// The reverse lookup value
	ReverseLookup *string `json:"reverseLookup,omitempty"`
	// IP version
	Version *int32 `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Ip Ip

// NewIp instantiates a new Ip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIp() *Ip {
	this := Ip{}
	return &this
}

// NewIpWithDefaults instantiates a new Ip object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpWithDefaults() *Ip {
	this := Ip{}
	return &this
}

// GetDdos returns the Ddos field value if set, zero value otherwise.
func (o *Ip) GetDdos() DDos {
	if o == nil || IsNil(o.Ddos) {
		var ret DDos
		return ret
	}
	return *o.Ddos
}

// GetDdosOk returns a tuple with the Ddos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ip) GetDdosOk() (*DDos, bool) {
	if o == nil || IsNil(o.Ddos) {
		return nil, false
	}
	return o.Ddos, true
}

// HasDdos returns a boolean if a field has been set.
func (o *Ip) HasDdos() bool {
	if o != nil && !IsNil(o.Ddos) {
		return true
	}

	return false
}

// SetDdos gets a reference to the given DDos and assigns it to the Ddos field.
func (o *Ip) SetDdos(v DDos) {
	o.Ddos = &v
}

// GetFloatingIp returns the FloatingIp field value if set, zero value otherwise.
func (o *Ip) GetFloatingIp() bool {
	if o == nil || IsNil(o.FloatingIp) {
		var ret bool
		return ret
	}
	return *o.FloatingIp
}

// GetFloatingIpOk returns a tuple with the FloatingIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ip) GetFloatingIpOk() (*bool, bool) {
	if o == nil || IsNil(o.FloatingIp) {
		return nil, false
	}
	return o.FloatingIp, true
}

// HasFloatingIp returns a boolean if a field has been set.
func (o *Ip) HasFloatingIp() bool {
	if o != nil && !IsNil(o.FloatingIp) {
		return true
	}

	return false
}

// SetFloatingIp gets a reference to the given bool and assigns it to the FloatingIp field.
func (o *Ip) SetFloatingIp(v bool) {
	o.FloatingIp = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *Ip) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ip) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *Ip) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *Ip) SetGateway(v string) {
	o.Gateway = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *Ip) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ip) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *Ip) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *Ip) SetIp(v string) {
	o.Ip = &v
}

// GetMainIp returns the MainIp field value if set, zero value otherwise.
func (o *Ip) GetMainIp() bool {
	if o == nil || IsNil(o.MainIp) {
		var ret bool
		return ret
	}
	return *o.MainIp
}

// GetMainIpOk returns a tuple with the MainIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ip) GetMainIpOk() (*bool, bool) {
	if o == nil || IsNil(o.MainIp) {
		return nil, false
	}
	return o.MainIp, true
}

// HasMainIp returns a boolean if a field has been set.
func (o *Ip) HasMainIp() bool {
	if o != nil && !IsNil(o.MainIp) {
		return true
	}

	return false
}

// SetMainIp gets a reference to the given bool and assigns it to the MainIp field.
func (o *Ip) SetMainIp(v bool) {
	o.MainIp = &v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *Ip) GetNetworkType() NetworkType {
	if o == nil || IsNil(o.NetworkType) {
		var ret NetworkType
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ip) GetNetworkTypeOk() (*NetworkType, bool) {
	if o == nil || IsNil(o.NetworkType) {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *Ip) HasNetworkType() bool {
	if o != nil && !IsNil(o.NetworkType) {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given NetworkType and assigns it to the NetworkType field.
func (o *Ip) SetNetworkType(v NetworkType) {
	o.NetworkType = &v
}

// GetNullRouted returns the NullRouted field value if set, zero value otherwise.
func (o *Ip) GetNullRouted() bool {
	if o == nil || IsNil(o.NullRouted) {
		var ret bool
		return ret
	}
	return *o.NullRouted
}

// GetNullRoutedOk returns a tuple with the NullRouted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ip) GetNullRoutedOk() (*bool, bool) {
	if o == nil || IsNil(o.NullRouted) {
		return nil, false
	}
	return o.NullRouted, true
}

// HasNullRouted returns a boolean if a field has been set.
func (o *Ip) HasNullRouted() bool {
	if o != nil && !IsNil(o.NullRouted) {
		return true
	}

	return false
}

// SetNullRouted gets a reference to the given bool and assigns it to the NullRouted field.
func (o *Ip) SetNullRouted(v bool) {
	o.NullRouted = &v
}

// GetReverseLookup returns the ReverseLookup field value if set, zero value otherwise.
func (o *Ip) GetReverseLookup() string {
	if o == nil || IsNil(o.ReverseLookup) {
		var ret string
		return ret
	}
	return *o.ReverseLookup
}

// GetReverseLookupOk returns a tuple with the ReverseLookup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ip) GetReverseLookupOk() (*string, bool) {
	if o == nil || IsNil(o.ReverseLookup) {
		return nil, false
	}
	return o.ReverseLookup, true
}

// HasReverseLookup returns a boolean if a field has been set.
func (o *Ip) HasReverseLookup() bool {
	if o != nil && !IsNil(o.ReverseLookup) {
		return true
	}

	return false
}

// SetReverseLookup gets a reference to the given string and assigns it to the ReverseLookup field.
func (o *Ip) SetReverseLookup(v string) {
	o.ReverseLookup = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Ip) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ip) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Ip) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *Ip) SetVersion(v int32) {
	o.Version = &v
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
	if !IsNil(o.Ddos) {
		toSerialize["ddos"] = o.Ddos
	}
	if !IsNil(o.FloatingIp) {
		toSerialize["floatingIp"] = o.FloatingIp
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.MainIp) {
		toSerialize["mainIp"] = o.MainIp
	}
	if !IsNil(o.NetworkType) {
		toSerialize["networkType"] = o.NetworkType
	}
	if !IsNil(o.NullRouted) {
		toSerialize["nullRouted"] = o.NullRouted
	}
	if !IsNil(o.ReverseLookup) {
		toSerialize["reverseLookup"] = o.ReverseLookup
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Ip) UnmarshalJSON(data []byte) (err error) {
	varIp := _Ip{}

	err = json.Unmarshal(data, &varIp)

	if err != nil {
		return err
	}

	*o = Ip(varIp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ddos")
		delete(additionalProperties, "floatingIp")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "mainIp")
		delete(additionalProperties, "networkType")
		delete(additionalProperties, "nullRouted")
		delete(additionalProperties, "reverseLookup")
		delete(additionalProperties, "version")
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


