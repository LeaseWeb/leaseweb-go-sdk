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

// checks if the Network type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Network{}

// Network struct for Network
type Network struct {
	Capabilities *NetworkCapabilities `json:"capabilities,omitempty"`
	Lldp *Lldp `json:"lldp,omitempty"`
	LogicalName *string `json:"logical_name,omitempty"`
	// Represents a MAC Address in the standard colon delimited format. Eg. `01:23:45:67:89:0A`
	MacAddress *string `json:"mac_address,omitempty" validate:"regexp=([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})|([0-9a-fA-F]{4}\\\\.[0-9a-fA-F]{4}\\\\.[0-9a-fA-F]{4})$"`
	Product *string `json:"product,omitempty"`
	Settings *NetworkSettings `json:"settings,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Network Network

// NewNetwork instantiates a new Network object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetwork() *Network {
	this := Network{}
	return &this
}

// NewNetworkWithDefaults instantiates a new Network object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkWithDefaults() *Network {
	this := Network{}
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *Network) GetCapabilities() NetworkCapabilities {
	if o == nil || IsNil(o.Capabilities) {
		var ret NetworkCapabilities
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetCapabilitiesOk() (*NetworkCapabilities, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *Network) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given NetworkCapabilities and assigns it to the Capabilities field.
func (o *Network) SetCapabilities(v NetworkCapabilities) {
	o.Capabilities = &v
}

// GetLldp returns the Lldp field value if set, zero value otherwise.
func (o *Network) GetLldp() Lldp {
	if o == nil || IsNil(o.Lldp) {
		var ret Lldp
		return ret
	}
	return *o.Lldp
}

// GetLldpOk returns a tuple with the Lldp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetLldpOk() (*Lldp, bool) {
	if o == nil || IsNil(o.Lldp) {
		return nil, false
	}
	return o.Lldp, true
}

// HasLldp returns a boolean if a field has been set.
func (o *Network) HasLldp() bool {
	if o != nil && !IsNil(o.Lldp) {
		return true
	}

	return false
}

// SetLldp gets a reference to the given Lldp and assigns it to the Lldp field.
func (o *Network) SetLldp(v Lldp) {
	o.Lldp = &v
}

// GetLogicalName returns the LogicalName field value if set, zero value otherwise.
func (o *Network) GetLogicalName() string {
	if o == nil || IsNil(o.LogicalName) {
		var ret string
		return ret
	}
	return *o.LogicalName
}

// GetLogicalNameOk returns a tuple with the LogicalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetLogicalNameOk() (*string, bool) {
	if o == nil || IsNil(o.LogicalName) {
		return nil, false
	}
	return o.LogicalName, true
}

// HasLogicalName returns a boolean if a field has been set.
func (o *Network) HasLogicalName() bool {
	if o != nil && !IsNil(o.LogicalName) {
		return true
	}

	return false
}

// SetLogicalName gets a reference to the given string and assigns it to the LogicalName field.
func (o *Network) SetLogicalName(v string) {
	o.LogicalName = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *Network) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *Network) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *Network) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *Network) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *Network) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *Network) SetProduct(v string) {
	o.Product = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *Network) GetSettings() NetworkSettings {
	if o == nil || IsNil(o.Settings) {
		var ret NetworkSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetSettingsOk() (*NetworkSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Network) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given NetworkSettings and assigns it to the Settings field.
func (o *Network) SetSettings(v NetworkSettings) {
	o.Settings = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *Network) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *Network) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *Network) SetVendor(v string) {
	o.Vendor = &v
}

func (o Network) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Network) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.Lldp) {
		toSerialize["lldp"] = o.Lldp
	}
	if !IsNil(o.LogicalName) {
		toSerialize["logical_name"] = o.LogicalName
	}
	if !IsNil(o.MacAddress) {
		toSerialize["mac_address"] = o.MacAddress
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Network) UnmarshalJSON(data []byte) (err error) {
	varNetwork := _Network{}

	err = json.Unmarshal(data, &varNetwork)

	if err != nil {
		return err
	}

	*o = Network(varNetwork)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "capabilities")
		delete(additionalProperties, "lldp")
		delete(additionalProperties, "logical_name")
		delete(additionalProperties, "mac_address")
		delete(additionalProperties, "product")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "vendor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetwork struct {
	value *Network
	isSet bool
}

func (v NullableNetwork) Get() *Network {
	return v.value
}

func (v *NullableNetwork) Set(val *Network) {
	v.value = val
	v.isSet = true
}

func (v NullableNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetwork(val *Network) *NullableNetwork {
	return &NullableNetwork{value: val, isSet: true}
}

func (v NullableNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


