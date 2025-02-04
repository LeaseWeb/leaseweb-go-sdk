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

// checks if the Lease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Lease{}

// Lease A single DHCP reservation
type Lease struct {
	// The PXE bootfile the server will network boot from
	Bootfile *string `json:"bootfile,omitempty"`
	// The time when the DHCP reservation was created
	CreatedAt *string `json:"createdAt,omitempty"`
	// The gateway for this DHCP reservation
	Gateway *string `json:"gateway,omitempty"`
	// The hostname for the server
	Hostname *string `json:"hostname,omitempty"`
	// The IP address this DHCP reservation is for
	Ip *string `json:"ip,omitempty"`
	LastClientRequest *LastClientRequest `json:"lastClientRequest,omitempty"`
	// Represents a MAC Address in the standard colon delimited format. Eg. `01:23:45:67:89:0A`
	Mac *string `json:"mac,omitempty" validate:"regexp=([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})|([0-9a-fA-F]{4}\\\\.[0-9a-fA-F]{4}\\\\.[0-9a-fA-F]{4})$"`
	// The netmask for this DHCP reservation
	Netmask *string `json:"netmask,omitempty"`
	// The site serving this DHCP reservation
	Site *string `json:"site,omitempty"`
	// The time when the DHCP reservation was last updated or used by a client
	UpdatedAt *string `json:"updatedAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Lease Lease

// NewLease instantiates a new Lease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLease() *Lease {
	this := Lease{}
	return &this
}

// NewLeaseWithDefaults instantiates a new Lease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeaseWithDefaults() *Lease {
	this := Lease{}
	return &this
}

// GetBootfile returns the Bootfile field value if set, zero value otherwise.
func (o *Lease) GetBootfile() string {
	if o == nil || IsNil(o.Bootfile) {
		var ret string
		return ret
	}
	return *o.Bootfile
}

// GetBootfileOk returns a tuple with the Bootfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lease) GetBootfileOk() (*string, bool) {
	if o == nil || IsNil(o.Bootfile) {
		return nil, false
	}
	return o.Bootfile, true
}

// HasBootfile returns a boolean if a field has been set.
func (o *Lease) HasBootfile() bool {
	if o != nil && !IsNil(o.Bootfile) {
		return true
	}

	return false
}

// SetBootfile gets a reference to the given string and assigns it to the Bootfile field.
func (o *Lease) SetBootfile(v string) {
	o.Bootfile = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Lease) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lease) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Lease) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Lease) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *Lease) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lease) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *Lease) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *Lease) SetGateway(v string) {
	o.Gateway = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *Lease) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lease) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *Lease) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *Lease) SetHostname(v string) {
	o.Hostname = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *Lease) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lease) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *Lease) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *Lease) SetIp(v string) {
	o.Ip = &v
}

// GetLastClientRequest returns the LastClientRequest field value if set, zero value otherwise.
func (o *Lease) GetLastClientRequest() LastClientRequest {
	if o == nil || IsNil(o.LastClientRequest) {
		var ret LastClientRequest
		return ret
	}
	return *o.LastClientRequest
}

// GetLastClientRequestOk returns a tuple with the LastClientRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lease) GetLastClientRequestOk() (*LastClientRequest, bool) {
	if o == nil || IsNil(o.LastClientRequest) {
		return nil, false
	}
	return o.LastClientRequest, true
}

// HasLastClientRequest returns a boolean if a field has been set.
func (o *Lease) HasLastClientRequest() bool {
	if o != nil && !IsNil(o.LastClientRequest) {
		return true
	}

	return false
}

// SetLastClientRequest gets a reference to the given LastClientRequest and assigns it to the LastClientRequest field.
func (o *Lease) SetLastClientRequest(v LastClientRequest) {
	o.LastClientRequest = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *Lease) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lease) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *Lease) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *Lease) SetMac(v string) {
	o.Mac = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *Lease) GetNetmask() string {
	if o == nil || IsNil(o.Netmask) {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lease) GetNetmaskOk() (*string, bool) {
	if o == nil || IsNil(o.Netmask) {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *Lease) HasNetmask() bool {
	if o != nil && !IsNil(o.Netmask) {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *Lease) SetNetmask(v string) {
	o.Netmask = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *Lease) GetSite() string {
	if o == nil || IsNil(o.Site) {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lease) GetSiteOk() (*string, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *Lease) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *Lease) SetSite(v string) {
	o.Site = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Lease) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lease) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Lease) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Lease) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o Lease) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Lease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bootfile) {
		toSerialize["bootfile"] = o.Bootfile
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.LastClientRequest) {
		toSerialize["lastClientRequest"] = o.LastClientRequest
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Netmask) {
		toSerialize["netmask"] = o.Netmask
	}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Lease) UnmarshalJSON(data []byte) (err error) {
	varLease := _Lease{}

	err = json.Unmarshal(data, &varLease)

	if err != nil {
		return err
	}

	*o = Lease(varLease)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bootfile")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "lastClientRequest")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "netmask")
		delete(additionalProperties, "site")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLease struct {
	value *Lease
	isSet bool
}

func (v NullableLease) Get() *Lease {
	return v.value
}

func (v *NullableLease) Set(val *Lease) {
	v.value = val
	v.isSet = true
}

func (v NullableLease) IsSet() bool {
	return v.isSet
}

func (v *NullableLease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLease(val *Lease) *NullableLease {
	return &NullableLease{value: val, isSet: true}
}

func (v NullableLease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


