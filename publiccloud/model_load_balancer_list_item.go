/*
Public Clouds

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the LoadBalancerListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerListItem{}

// LoadBalancerListItem struct for LoadBalancerListItem
type LoadBalancerListItem struct {
	// The load balancer unique identifier
	Id string `json:"id"`
	Type TypeName `json:"type"`
	Resources Resources `json:"resources"`
	// The identifying name set to the load balancer
	Reference NullableString `json:"reference"`
	State State `json:"state"`
	// Date and time when the load balancer was started for the first time, right after launching it
	StartedAt NullableTime `json:"startedAt"`
	Region RegionName `json:"region"`
	Configuration NullableLoadBalancerConfiguration `json:"configuration"`
	AutoScalingGroup NullableAutoScalingGroup `json:"autoScalingGroup"`
	PrivateNetwork NullablePrivateNetwork `json:"privateNetwork"`
	Contract Contract `json:"contract"`
	Ips []Ip `json:"ips"`
	AdditionalProperties map[string]interface{}
}

type _LoadBalancerListItem LoadBalancerListItem

// NewLoadBalancerListItem instantiates a new LoadBalancerListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerListItem(id string, type_ TypeName, resources Resources, reference NullableString, state State, startedAt NullableTime, region RegionName, configuration NullableLoadBalancerConfiguration, autoScalingGroup NullableAutoScalingGroup, privateNetwork NullablePrivateNetwork, contract Contract, ips []Ip) *LoadBalancerListItem {
	this := LoadBalancerListItem{}
	this.Id = id
	this.Type = type_
	this.Resources = resources
	this.Reference = reference
	this.State = state
	this.StartedAt = startedAt
	this.Region = region
	this.Configuration = configuration
	this.AutoScalingGroup = autoScalingGroup
	this.PrivateNetwork = privateNetwork
	this.Contract = contract
	this.Ips = ips
	return &this
}

// NewLoadBalancerListItemWithDefaults instantiates a new LoadBalancerListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerListItemWithDefaults() *LoadBalancerListItem {
	this := LoadBalancerListItem{}
	return &this
}

// GetId returns the Id field value
func (o *LoadBalancerListItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerListItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LoadBalancerListItem) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *LoadBalancerListItem) GetType() TypeName {
	if o == nil {
		var ret TypeName
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerListItem) GetTypeOk() (*TypeName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LoadBalancerListItem) SetType(v TypeName) {
	o.Type = v
}

// GetResources returns the Resources field value
func (o *LoadBalancerListItem) GetResources() Resources {
	if o == nil {
		var ret Resources
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerListItem) GetResourcesOk() (*Resources, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *LoadBalancerListItem) SetResources(v Resources) {
	o.Resources = v
}

// GetReference returns the Reference field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LoadBalancerListItem) GetReference() string {
	if o == nil || o.Reference.Get() == nil {
		var ret string
		return ret
	}

	return *o.Reference.Get()
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadBalancerListItem) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reference.Get(), o.Reference.IsSet()
}

// SetReference sets field value
func (o *LoadBalancerListItem) SetReference(v string) {
	o.Reference.Set(&v)
}

// GetState returns the State field value
func (o *LoadBalancerListItem) GetState() State {
	if o == nil {
		var ret State
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerListItem) GetStateOk() (*State, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *LoadBalancerListItem) SetState(v State) {
	o.State = v
}

// GetStartedAt returns the StartedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *LoadBalancerListItem) GetStartedAt() time.Time {
	if o == nil || o.StartedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadBalancerListItem) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// SetStartedAt sets field value
func (o *LoadBalancerListItem) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}

// GetRegion returns the Region field value
func (o *LoadBalancerListItem) GetRegion() RegionName {
	if o == nil {
		var ret RegionName
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerListItem) GetRegionOk() (*RegionName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *LoadBalancerListItem) SetRegion(v RegionName) {
	o.Region = v
}

// GetConfiguration returns the Configuration field value
// If the value is explicit nil, the zero value for LoadBalancerConfiguration will be returned
func (o *LoadBalancerListItem) GetConfiguration() LoadBalancerConfiguration {
	if o == nil || o.Configuration.Get() == nil {
		var ret LoadBalancerConfiguration
		return ret
	}

	return *o.Configuration.Get()
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadBalancerListItem) GetConfigurationOk() (*LoadBalancerConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Configuration.Get(), o.Configuration.IsSet()
}

// SetConfiguration sets field value
func (o *LoadBalancerListItem) SetConfiguration(v LoadBalancerConfiguration) {
	o.Configuration.Set(&v)
}

// GetAutoScalingGroup returns the AutoScalingGroup field value
// If the value is explicit nil, the zero value for AutoScalingGroup will be returned
func (o *LoadBalancerListItem) GetAutoScalingGroup() AutoScalingGroup {
	if o == nil || o.AutoScalingGroup.Get() == nil {
		var ret AutoScalingGroup
		return ret
	}

	return *o.AutoScalingGroup.Get()
}

// GetAutoScalingGroupOk returns a tuple with the AutoScalingGroup field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadBalancerListItem) GetAutoScalingGroupOk() (*AutoScalingGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoScalingGroup.Get(), o.AutoScalingGroup.IsSet()
}

// SetAutoScalingGroup sets field value
func (o *LoadBalancerListItem) SetAutoScalingGroup(v AutoScalingGroup) {
	o.AutoScalingGroup.Set(&v)
}

// GetPrivateNetwork returns the PrivateNetwork field value
// If the value is explicit nil, the zero value for PrivateNetwork will be returned
func (o *LoadBalancerListItem) GetPrivateNetwork() PrivateNetwork {
	if o == nil || o.PrivateNetwork.Get() == nil {
		var ret PrivateNetwork
		return ret
	}

	return *o.PrivateNetwork.Get()
}

// GetPrivateNetworkOk returns a tuple with the PrivateNetwork field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadBalancerListItem) GetPrivateNetworkOk() (*PrivateNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateNetwork.Get(), o.PrivateNetwork.IsSet()
}

// SetPrivateNetwork sets field value
func (o *LoadBalancerListItem) SetPrivateNetwork(v PrivateNetwork) {
	o.PrivateNetwork.Set(&v)
}

// GetContract returns the Contract field value
func (o *LoadBalancerListItem) GetContract() Contract {
	if o == nil {
		var ret Contract
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerListItem) GetContractOk() (*Contract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *LoadBalancerListItem) SetContract(v Contract) {
	o.Contract = v
}

// GetIps returns the Ips field value
func (o *LoadBalancerListItem) GetIps() []Ip {
	if o == nil {
		var ret []Ip
		return ret
	}

	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerListItem) GetIpsOk() ([]Ip, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ips, true
}

// SetIps sets field value
func (o *LoadBalancerListItem) SetIps(v []Ip) {
	o.Ips = v
}

func (o LoadBalancerListItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["resources"] = o.Resources
	toSerialize["reference"] = o.Reference.Get()
	toSerialize["state"] = o.State
	toSerialize["startedAt"] = o.StartedAt.Get()
	toSerialize["region"] = o.Region
	toSerialize["configuration"] = o.Configuration.Get()
	toSerialize["autoScalingGroup"] = o.AutoScalingGroup.Get()
	toSerialize["privateNetwork"] = o.PrivateNetwork.Get()
	toSerialize["contract"] = o.Contract
	toSerialize["ips"] = o.Ips

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadBalancerListItem) UnmarshalJSON(data []byte) (err error) {
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
		"region",
		"configuration",
		"autoScalingGroup",
		"privateNetwork",
		"contract",
		"ips",
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

	varLoadBalancerListItem := _LoadBalancerListItem{}

	err = json.Unmarshal(data, &varLoadBalancerListItem)

	if err != nil {
		return err
	}

	*o = LoadBalancerListItem(varLoadBalancerListItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "resources")
		delete(additionalProperties, "reference")
		delete(additionalProperties, "state")
		delete(additionalProperties, "startedAt")
		delete(additionalProperties, "region")
		delete(additionalProperties, "configuration")
		delete(additionalProperties, "autoScalingGroup")
		delete(additionalProperties, "privateNetwork")
		delete(additionalProperties, "contract")
		delete(additionalProperties, "ips")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadBalancerListItem struct {
	value *LoadBalancerListItem
	isSet bool
}

func (v NullableLoadBalancerListItem) Get() *LoadBalancerListItem {
	return v.value
}

func (v *NullableLoadBalancerListItem) Set(val *LoadBalancerListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerListItem(val *LoadBalancerListItem) *NullableLoadBalancerListItem {
	return &NullableLoadBalancerListItem{value: val, isSet: true}
}

func (v NullableLoadBalancerListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


