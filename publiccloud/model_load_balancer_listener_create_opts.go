/*
Public Clouds

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
	"fmt"
)

// checks if the LoadBalancerListenerCreateOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerListenerCreateOpts{}

// LoadBalancerListenerCreateOpts struct for LoadBalancerListenerCreateOpts
type LoadBalancerListenerCreateOpts struct {
	Protocol Protocol `json:"protocol"`
	// Port that the listener listens to
	Port int32 `json:"port"`
	Certificate *SslCertificate `json:"certificate,omitempty"`
	DefaultRule LoadBalancerListenerDefaultRule `json:"defaultRule"`
	AdditionalProperties map[string]interface{}
}

type _LoadBalancerListenerCreateOpts LoadBalancerListenerCreateOpts

// NewLoadBalancerListenerCreateOpts instantiates a new LoadBalancerListenerCreateOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerListenerCreateOpts(protocol Protocol, port int32, defaultRule LoadBalancerListenerDefaultRule) *LoadBalancerListenerCreateOpts {
	this := LoadBalancerListenerCreateOpts{}
	this.Protocol = protocol
	this.Port = port
	this.DefaultRule = defaultRule
	return &this
}

// NewLoadBalancerListenerCreateOptsWithDefaults instantiates a new LoadBalancerListenerCreateOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerListenerCreateOptsWithDefaults() *LoadBalancerListenerCreateOpts {
	this := LoadBalancerListenerCreateOpts{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *LoadBalancerListenerCreateOpts) GetProtocol() Protocol {
	if o == nil {
		var ret Protocol
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerListenerCreateOpts) GetProtocolOk() (*Protocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *LoadBalancerListenerCreateOpts) SetProtocol(v Protocol) {
	o.Protocol = v
}

// GetPort returns the Port field value
func (o *LoadBalancerListenerCreateOpts) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerListenerCreateOpts) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *LoadBalancerListenerCreateOpts) SetPort(v int32) {
	o.Port = v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *LoadBalancerListenerCreateOpts) GetCertificate() SslCertificate {
	if o == nil || IsNil(o.Certificate) {
		var ret SslCertificate
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerListenerCreateOpts) GetCertificateOk() (*SslCertificate, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *LoadBalancerListenerCreateOpts) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given SslCertificate and assigns it to the Certificate field.
func (o *LoadBalancerListenerCreateOpts) SetCertificate(v SslCertificate) {
	o.Certificate = &v
}

// GetDefaultRule returns the DefaultRule field value
func (o *LoadBalancerListenerCreateOpts) GetDefaultRule() LoadBalancerListenerDefaultRule {
	if o == nil {
		var ret LoadBalancerListenerDefaultRule
		return ret
	}

	return o.DefaultRule
}

// GetDefaultRuleOk returns a tuple with the DefaultRule field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerListenerCreateOpts) GetDefaultRuleOk() (*LoadBalancerListenerDefaultRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultRule, true
}

// SetDefaultRule sets field value
func (o *LoadBalancerListenerCreateOpts) SetDefaultRule(v LoadBalancerListenerDefaultRule) {
	o.DefaultRule = v
}

func (o LoadBalancerListenerCreateOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerListenerCreateOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["protocol"] = o.Protocol
	toSerialize["port"] = o.Port
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	toSerialize["defaultRule"] = o.DefaultRule

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadBalancerListenerCreateOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"protocol",
		"port",
		"defaultRule",
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

	varLoadBalancerListenerCreateOpts := _LoadBalancerListenerCreateOpts{}

	err = json.Unmarshal(data, &varLoadBalancerListenerCreateOpts)

	if err != nil {
		return err
	}

	*o = LoadBalancerListenerCreateOpts(varLoadBalancerListenerCreateOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "port")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "defaultRule")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadBalancerListenerCreateOpts struct {
	value *LoadBalancerListenerCreateOpts
	isSet bool
}

func (v NullableLoadBalancerListenerCreateOpts) Get() *LoadBalancerListenerCreateOpts {
	return v.value
}

func (v *NullableLoadBalancerListenerCreateOpts) Set(val *LoadBalancerListenerCreateOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerListenerCreateOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerListenerCreateOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerListenerCreateOpts(val *LoadBalancerListenerCreateOpts) *NullableLoadBalancerListenerCreateOpts {
	return &NullableLoadBalancerListenerCreateOpts{value: val, isSet: true}
}

func (v NullableLoadBalancerListenerCreateOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerListenerCreateOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


