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
	"fmt"
)

// checks if the LoadBalancerDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerDetails{}

// LoadBalancerDetails struct for LoadBalancerDetails
type LoadBalancerDetails struct {
	// The load balancer unique identifier
	Id string `json:"id"`
	// Load balancer type
	Type string `json:"type"`
	Resources Resources `json:"resources"`
	// The identifying name set to the load balancer
	Reference NullableString `json:"reference"`
	State State `json:"state"`
	// Date and time when the load balancer was started for the first time, right after launching it
	StartedAt NullableTime `json:"startedAt"`
	Ips []IpDetails `json:"ips"`
	// The region where the load balancer was launched into
	Region string `json:"region"`
	Configuration NullableLoadBalancerConfiguration `json:"configuration"`
	AutoScalingGroup NullableAutoScalingGroup `json:"autoScalingGroup"`
	PrivateNetwork NullablePrivateNetwork `json:"privateNetwork"`
	Contract Contract `json:"contract"`
	AdditionalProperties map[string]interface{}
}

type _LoadBalancerDetails LoadBalancerDetails

// NewLoadBalancerDetails instantiates a new LoadBalancerDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerDetails(id string, type_ string, resources Resources, reference NullableString, state State, startedAt NullableTime, ips []IpDetails, region string, configuration NullableLoadBalancerConfiguration, autoScalingGroup NullableAutoScalingGroup, privateNetwork NullablePrivateNetwork, contract Contract) *LoadBalancerDetails {
	this := LoadBalancerDetails{}
	this.Id = id
	this.Type = type_
	this.Resources = resources
	this.Reference = reference
	this.State = state
	this.StartedAt = startedAt
	this.Ips = ips
	this.Region = region
	this.Configuration = configuration
	this.AutoScalingGroup = autoScalingGroup
	this.PrivateNetwork = privateNetwork
	this.Contract = contract
	return &this
}

// NewLoadBalancerDetailsWithDefaults instantiates a new LoadBalancerDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerDetailsWithDefaults() *LoadBalancerDetails {
	this := LoadBalancerDetails{}
	return &this
}

// GetId returns the Id field value
func (o *LoadBalancerDetails) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerDetails) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LoadBalancerDetails) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *LoadBalancerDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LoadBalancerDetails) SetType(v string) {
	o.Type = v
}

// GetResources returns the Resources field value
func (o *LoadBalancerDetails) GetResources() Resources {
	if o == nil {
		var ret Resources
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerDetails) GetResourcesOk() (*Resources, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *LoadBalancerDetails) SetResources(v Resources) {
	o.Resources = v
}

// GetReference returns the Reference field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LoadBalancerDetails) GetReference() string {
	if o == nil || o.Reference.Get() == nil {
		var ret string
		return ret
	}

	return *o.Reference.Get()
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadBalancerDetails) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reference.Get(), o.Reference.IsSet()
}

// SetReference sets field value
func (o *LoadBalancerDetails) SetReference(v string) {
	o.Reference.Set(&v)
}

// GetState returns the State field value
func (o *LoadBalancerDetails) GetState() State {
	if o == nil {
		var ret State
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerDetails) GetStateOk() (*State, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *LoadBalancerDetails) SetState(v State) {
	o.State = v
}

// GetStartedAt returns the StartedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *LoadBalancerDetails) GetStartedAt() time.Time {
	if o == nil || o.StartedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadBalancerDetails) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// SetStartedAt sets field value
func (o *LoadBalancerDetails) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}

// GetIps returns the Ips field value
func (o *LoadBalancerDetails) GetIps() []IpDetails {
	if o == nil {
		var ret []IpDetails
		return ret
	}

	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerDetails) GetIpsOk() ([]IpDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ips, true
}

// SetIps sets field value
func (o *LoadBalancerDetails) SetIps(v []IpDetails) {
	o.Ips = v
}

// GetRegion returns the Region field value
func (o *LoadBalancerDetails) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerDetails) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *LoadBalancerDetails) SetRegion(v string) {
	o.Region = v
}

// GetConfiguration returns the Configuration field value
// If the value is explicit nil, the zero value for LoadBalancerConfiguration will be returned
func (o *LoadBalancerDetails) GetConfiguration() LoadBalancerConfiguration {
	if o == nil || o.Configuration.Get() == nil {
		var ret LoadBalancerConfiguration
		return ret
	}

	return *o.Configuration.Get()
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadBalancerDetails) GetConfigurationOk() (*LoadBalancerConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Configuration.Get(), o.Configuration.IsSet()
}

// SetConfiguration sets field value
func (o *LoadBalancerDetails) SetConfiguration(v LoadBalancerConfiguration) {
	o.Configuration.Set(&v)
}

// GetAutoScalingGroup returns the AutoScalingGroup field value
// If the value is explicit nil, the zero value for AutoScalingGroup will be returned
func (o *LoadBalancerDetails) GetAutoScalingGroup() AutoScalingGroup {
	if o == nil || o.AutoScalingGroup.Get() == nil {
		var ret AutoScalingGroup
		return ret
	}

	return *o.AutoScalingGroup.Get()
}

// GetAutoScalingGroupOk returns a tuple with the AutoScalingGroup field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadBalancerDetails) GetAutoScalingGroupOk() (*AutoScalingGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoScalingGroup.Get(), o.AutoScalingGroup.IsSet()
}

// SetAutoScalingGroup sets field value
func (o *LoadBalancerDetails) SetAutoScalingGroup(v AutoScalingGroup) {
	o.AutoScalingGroup.Set(&v)
}

// GetPrivateNetwork returns the PrivateNetwork field value
// If the value is explicit nil, the zero value for PrivateNetwork will be returned
func (o *LoadBalancerDetails) GetPrivateNetwork() PrivateNetwork {
	if o == nil || o.PrivateNetwork.Get() == nil {
		var ret PrivateNetwork
		return ret
	}

	return *o.PrivateNetwork.Get()
}

// GetPrivateNetworkOk returns a tuple with the PrivateNetwork field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadBalancerDetails) GetPrivateNetworkOk() (*PrivateNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateNetwork.Get(), o.PrivateNetwork.IsSet()
}

// SetPrivateNetwork sets field value
func (o *LoadBalancerDetails) SetPrivateNetwork(v PrivateNetwork) {
	o.PrivateNetwork.Set(&v)
}

// GetContract returns the Contract field value
func (o *LoadBalancerDetails) GetContract() Contract {
	if o == nil {
		var ret Contract
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerDetails) GetContractOk() (*Contract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *LoadBalancerDetails) SetContract(v Contract) {
	o.Contract = v
}

func (o LoadBalancerDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["resources"] = o.Resources
	toSerialize["reference"] = o.Reference.Get()
	toSerialize["state"] = o.State
	toSerialize["startedAt"] = o.StartedAt.Get()
	toSerialize["ips"] = o.Ips
	toSerialize["region"] = o.Region
	toSerialize["configuration"] = o.Configuration.Get()
	toSerialize["autoScalingGroup"] = o.AutoScalingGroup.Get()
	toSerialize["privateNetwork"] = o.PrivateNetwork.Get()
	toSerialize["contract"] = o.Contract

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadBalancerDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"resources",
		"reference",
		"state",
		"startedAt",
		"ips",
		"region",
		"configuration",
		"autoScalingGroup",
		"privateNetwork",
		"contract",
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

	varLoadBalancerDetails := _LoadBalancerDetails{}

	err = json.Unmarshal(data, &varLoadBalancerDetails)

	if err != nil {
		return err
	}

	*o = LoadBalancerDetails(varLoadBalancerDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "resources")
		delete(additionalProperties, "reference")
		delete(additionalProperties, "state")
		delete(additionalProperties, "startedAt")
		delete(additionalProperties, "ips")
		delete(additionalProperties, "region")
		delete(additionalProperties, "configuration")
		delete(additionalProperties, "autoScalingGroup")
		delete(additionalProperties, "privateNetwork")
		delete(additionalProperties, "contract")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadBalancerDetails struct {
	value *LoadBalancerDetails
	isSet bool
}

func (v NullableLoadBalancerDetails) Get() *LoadBalancerDetails {
	return v.value
}

func (v *NullableLoadBalancerDetails) Set(val *LoadBalancerDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerDetails(val *LoadBalancerDetails) *NullableLoadBalancerDetails {
	return &NullableLoadBalancerDetails{value: val, isSet: true}
}

func (v NullableLoadBalancerDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


