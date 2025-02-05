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

// checks if the NetworkCapabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkCapabilities{}

// NetworkCapabilities struct for NetworkCapabilities
type NetworkCapabilities struct {
	Autonegotiation *string `json:"autonegotiation,omitempty"`
	BusMaster *string `json:"bus_master,omitempty"`
	CapList *string `json:"cap_list,omitempty"`
	Ethernet *string `json:"ethernet,omitempty"`
	LinkSpeeds *LinkSpeeds `json:"link_speeds,omitempty"`
	Msi *string `json:"msi,omitempty"`
	Msix *string `json:"msix,omitempty"`
	Pciexpress *string `json:"pciexpress,omitempty"`
	Physical *string `json:"physical,omitempty"`
	Pm *string `json:"pm,omitempty"`
	Tp *string `json:"tp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkCapabilities NetworkCapabilities

// NewNetworkCapabilities instantiates a new NetworkCapabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkCapabilities() *NetworkCapabilities {
	this := NetworkCapabilities{}
	return &this
}

// NewNetworkCapabilitiesWithDefaults instantiates a new NetworkCapabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkCapabilitiesWithDefaults() *NetworkCapabilities {
	this := NetworkCapabilities{}
	return &this
}

// GetAutonegotiation returns the Autonegotiation field value if set, zero value otherwise.
func (o *NetworkCapabilities) GetAutonegotiation() string {
	if o == nil || IsNil(o.Autonegotiation) {
		var ret string
		return ret
	}
	return *o.Autonegotiation
}

// GetAutonegotiationOk returns a tuple with the Autonegotiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkCapabilities) GetAutonegotiationOk() (*string, bool) {
	if o == nil || IsNil(o.Autonegotiation) {
		return nil, false
	}
	return o.Autonegotiation, true
}

// HasAutonegotiation returns a boolean if a field has been set.
func (o *NetworkCapabilities) HasAutonegotiation() bool {
	if o != nil && !IsNil(o.Autonegotiation) {
		return true
	}

	return false
}

// SetAutonegotiation gets a reference to the given string and assigns it to the Autonegotiation field.
func (o *NetworkCapabilities) SetAutonegotiation(v string) {
	o.Autonegotiation = &v
}

// GetBusMaster returns the BusMaster field value if set, zero value otherwise.
func (o *NetworkCapabilities) GetBusMaster() string {
	if o == nil || IsNil(o.BusMaster) {
		var ret string
		return ret
	}
	return *o.BusMaster
}

// GetBusMasterOk returns a tuple with the BusMaster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkCapabilities) GetBusMasterOk() (*string, bool) {
	if o == nil || IsNil(o.BusMaster) {
		return nil, false
	}
	return o.BusMaster, true
}

// HasBusMaster returns a boolean if a field has been set.
func (o *NetworkCapabilities) HasBusMaster() bool {
	if o != nil && !IsNil(o.BusMaster) {
		return true
	}

	return false
}

// SetBusMaster gets a reference to the given string and assigns it to the BusMaster field.
func (o *NetworkCapabilities) SetBusMaster(v string) {
	o.BusMaster = &v
}

// GetCapList returns the CapList field value if set, zero value otherwise.
func (o *NetworkCapabilities) GetCapList() string {
	if o == nil || IsNil(o.CapList) {
		var ret string
		return ret
	}
	return *o.CapList
}

// GetCapListOk returns a tuple with the CapList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkCapabilities) GetCapListOk() (*string, bool) {
	if o == nil || IsNil(o.CapList) {
		return nil, false
	}
	return o.CapList, true
}

// HasCapList returns a boolean if a field has been set.
func (o *NetworkCapabilities) HasCapList() bool {
	if o != nil && !IsNil(o.CapList) {
		return true
	}

	return false
}

// SetCapList gets a reference to the given string and assigns it to the CapList field.
func (o *NetworkCapabilities) SetCapList(v string) {
	o.CapList = &v
}

// GetEthernet returns the Ethernet field value if set, zero value otherwise.
func (o *NetworkCapabilities) GetEthernet() string {
	if o == nil || IsNil(o.Ethernet) {
		var ret string
		return ret
	}
	return *o.Ethernet
}

// GetEthernetOk returns a tuple with the Ethernet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkCapabilities) GetEthernetOk() (*string, bool) {
	if o == nil || IsNil(o.Ethernet) {
		return nil, false
	}
	return o.Ethernet, true
}

// HasEthernet returns a boolean if a field has been set.
func (o *NetworkCapabilities) HasEthernet() bool {
	if o != nil && !IsNil(o.Ethernet) {
		return true
	}

	return false
}

// SetEthernet gets a reference to the given string and assigns it to the Ethernet field.
func (o *NetworkCapabilities) SetEthernet(v string) {
	o.Ethernet = &v
}

// GetLinkSpeeds returns the LinkSpeeds field value if set, zero value otherwise.
func (o *NetworkCapabilities) GetLinkSpeeds() LinkSpeeds {
	if o == nil || IsNil(o.LinkSpeeds) {
		var ret LinkSpeeds
		return ret
	}
	return *o.LinkSpeeds
}

// GetLinkSpeedsOk returns a tuple with the LinkSpeeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkCapabilities) GetLinkSpeedsOk() (*LinkSpeeds, bool) {
	if o == nil || IsNil(o.LinkSpeeds) {
		return nil, false
	}
	return o.LinkSpeeds, true
}

// HasLinkSpeeds returns a boolean if a field has been set.
func (o *NetworkCapabilities) HasLinkSpeeds() bool {
	if o != nil && !IsNil(o.LinkSpeeds) {
		return true
	}

	return false
}

// SetLinkSpeeds gets a reference to the given LinkSpeeds and assigns it to the LinkSpeeds field.
func (o *NetworkCapabilities) SetLinkSpeeds(v LinkSpeeds) {
	o.LinkSpeeds = &v
}

// GetMsi returns the Msi field value if set, zero value otherwise.
func (o *NetworkCapabilities) GetMsi() string {
	if o == nil || IsNil(o.Msi) {
		var ret string
		return ret
	}
	return *o.Msi
}

// GetMsiOk returns a tuple with the Msi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkCapabilities) GetMsiOk() (*string, bool) {
	if o == nil || IsNil(o.Msi) {
		return nil, false
	}
	return o.Msi, true
}

// HasMsi returns a boolean if a field has been set.
func (o *NetworkCapabilities) HasMsi() bool {
	if o != nil && !IsNil(o.Msi) {
		return true
	}

	return false
}

// SetMsi gets a reference to the given string and assigns it to the Msi field.
func (o *NetworkCapabilities) SetMsi(v string) {
	o.Msi = &v
}

// GetMsix returns the Msix field value if set, zero value otherwise.
func (o *NetworkCapabilities) GetMsix() string {
	if o == nil || IsNil(o.Msix) {
		var ret string
		return ret
	}
	return *o.Msix
}

// GetMsixOk returns a tuple with the Msix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkCapabilities) GetMsixOk() (*string, bool) {
	if o == nil || IsNil(o.Msix) {
		return nil, false
	}
	return o.Msix, true
}

// HasMsix returns a boolean if a field has been set.
func (o *NetworkCapabilities) HasMsix() bool {
	if o != nil && !IsNil(o.Msix) {
		return true
	}

	return false
}

// SetMsix gets a reference to the given string and assigns it to the Msix field.
func (o *NetworkCapabilities) SetMsix(v string) {
	o.Msix = &v
}

// GetPciexpress returns the Pciexpress field value if set, zero value otherwise.
func (o *NetworkCapabilities) GetPciexpress() string {
	if o == nil || IsNil(o.Pciexpress) {
		var ret string
		return ret
	}
	return *o.Pciexpress
}

// GetPciexpressOk returns a tuple with the Pciexpress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkCapabilities) GetPciexpressOk() (*string, bool) {
	if o == nil || IsNil(o.Pciexpress) {
		return nil, false
	}
	return o.Pciexpress, true
}

// HasPciexpress returns a boolean if a field has been set.
func (o *NetworkCapabilities) HasPciexpress() bool {
	if o != nil && !IsNil(o.Pciexpress) {
		return true
	}

	return false
}

// SetPciexpress gets a reference to the given string and assigns it to the Pciexpress field.
func (o *NetworkCapabilities) SetPciexpress(v string) {
	o.Pciexpress = &v
}

// GetPhysical returns the Physical field value if set, zero value otherwise.
func (o *NetworkCapabilities) GetPhysical() string {
	if o == nil || IsNil(o.Physical) {
		var ret string
		return ret
	}
	return *o.Physical
}

// GetPhysicalOk returns a tuple with the Physical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkCapabilities) GetPhysicalOk() (*string, bool) {
	if o == nil || IsNil(o.Physical) {
		return nil, false
	}
	return o.Physical, true
}

// HasPhysical returns a boolean if a field has been set.
func (o *NetworkCapabilities) HasPhysical() bool {
	if o != nil && !IsNil(o.Physical) {
		return true
	}

	return false
}

// SetPhysical gets a reference to the given string and assigns it to the Physical field.
func (o *NetworkCapabilities) SetPhysical(v string) {
	o.Physical = &v
}

// GetPm returns the Pm field value if set, zero value otherwise.
func (o *NetworkCapabilities) GetPm() string {
	if o == nil || IsNil(o.Pm) {
		var ret string
		return ret
	}
	return *o.Pm
}

// GetPmOk returns a tuple with the Pm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkCapabilities) GetPmOk() (*string, bool) {
	if o == nil || IsNil(o.Pm) {
		return nil, false
	}
	return o.Pm, true
}

// HasPm returns a boolean if a field has been set.
func (o *NetworkCapabilities) HasPm() bool {
	if o != nil && !IsNil(o.Pm) {
		return true
	}

	return false
}

// SetPm gets a reference to the given string and assigns it to the Pm field.
func (o *NetworkCapabilities) SetPm(v string) {
	o.Pm = &v
}

// GetTp returns the Tp field value if set, zero value otherwise.
func (o *NetworkCapabilities) GetTp() string {
	if o == nil || IsNil(o.Tp) {
		var ret string
		return ret
	}
	return *o.Tp
}

// GetTpOk returns a tuple with the Tp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkCapabilities) GetTpOk() (*string, bool) {
	if o == nil || IsNil(o.Tp) {
		return nil, false
	}
	return o.Tp, true
}

// HasTp returns a boolean if a field has been set.
func (o *NetworkCapabilities) HasTp() bool {
	if o != nil && !IsNil(o.Tp) {
		return true
	}

	return false
}

// SetTp gets a reference to the given string and assigns it to the Tp field.
func (o *NetworkCapabilities) SetTp(v string) {
	o.Tp = &v
}

func (o NetworkCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkCapabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Autonegotiation) {
		toSerialize["autonegotiation"] = o.Autonegotiation
	}
	if !IsNil(o.BusMaster) {
		toSerialize["bus_master"] = o.BusMaster
	}
	if !IsNil(o.CapList) {
		toSerialize["cap_list"] = o.CapList
	}
	if !IsNil(o.Ethernet) {
		toSerialize["ethernet"] = o.Ethernet
	}
	if !IsNil(o.LinkSpeeds) {
		toSerialize["link_speeds"] = o.LinkSpeeds
	}
	if !IsNil(o.Msi) {
		toSerialize["msi"] = o.Msi
	}
	if !IsNil(o.Msix) {
		toSerialize["msix"] = o.Msix
	}
	if !IsNil(o.Pciexpress) {
		toSerialize["pciexpress"] = o.Pciexpress
	}
	if !IsNil(o.Physical) {
		toSerialize["physical"] = o.Physical
	}
	if !IsNil(o.Pm) {
		toSerialize["pm"] = o.Pm
	}
	if !IsNil(o.Tp) {
		toSerialize["tp"] = o.Tp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkCapabilities) UnmarshalJSON(data []byte) (err error) {
	varNetworkCapabilities := _NetworkCapabilities{}

	err = json.Unmarshal(data, &varNetworkCapabilities)

	if err != nil {
		return err
	}

	*o = NetworkCapabilities(varNetworkCapabilities)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "autonegotiation")
		delete(additionalProperties, "bus_master")
		delete(additionalProperties, "cap_list")
		delete(additionalProperties, "ethernet")
		delete(additionalProperties, "link_speeds")
		delete(additionalProperties, "msi")
		delete(additionalProperties, "msix")
		delete(additionalProperties, "pciexpress")
		delete(additionalProperties, "physical")
		delete(additionalProperties, "pm")
		delete(additionalProperties, "tp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkCapabilities struct {
	value *NetworkCapabilities
	isSet bool
}

func (v NullableNetworkCapabilities) Get() *NetworkCapabilities {
	return v.value
}

func (v *NullableNetworkCapabilities) Set(val *NetworkCapabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkCapabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkCapabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkCapabilities(val *NetworkCapabilities) *NullableNetworkCapabilities {
	return &NullableNetworkCapabilities{value: val, isSet: true}
}

func (v NullableNetworkCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkCapabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


