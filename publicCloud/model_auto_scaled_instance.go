/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
)

// checks if the AutoScaledInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoScaledInstance{}

// AutoScaledInstance struct for AutoScaledInstance
type AutoScaledInstance struct {
	// The unique identifier of the instance
	Id *string `json:"id,omitempty"`
	// The instance type, which determines the amount of resources
	Type *string `json:"type,omitempty"`
	Resources *InstanceResources `json:"resources,omitempty"`
	// The region in which the instance was launched
	Region *string `json:"region,omitempty"`
	// The identifying name set to the instance
	Reference *string `json:"reference,omitempty"`
	State *State `json:"state,omitempty"`
	ProductType *string `json:"productType,omitempty"`
	includesPrivateNetwork *bool `json:"hasPrivateNetwork,omitempty"`
	Ips []Ip `json:"ips,omitempty"`
}

// NewAutoScaledInstance instantiates a new AutoScaledInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoScaledInstance() *AutoScaledInstance {
	this := AutoScaledInstance{}
	return &this
}

// NewAutoScaledInstanceWithDefaults instantiates a new AutoScaledInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoScaledInstanceWithDefaults() *AutoScaledInstance {
	this := AutoScaledInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutoScaledInstance) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScaledInstance) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutoScaledInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AutoScaledInstance) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AutoScaledInstance) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScaledInstance) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AutoScaledInstance) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AutoScaledInstance) SetType(v string) {
	o.Type = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *AutoScaledInstance) GetResources() InstanceResources {
	if o == nil || IsNil(o.Resources) {
		var ret InstanceResources
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScaledInstance) GetResourcesOk() (*InstanceResources, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *AutoScaledInstance) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given InstanceResources and assigns it to the Resources field.
func (o *AutoScaledInstance) SetResources(v InstanceResources) {
	o.Resources = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AutoScaledInstance) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScaledInstance) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AutoScaledInstance) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AutoScaledInstance) SetRegion(v string) {
	o.Region = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *AutoScaledInstance) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScaledInstance) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *AutoScaledInstance) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *AutoScaledInstance) SetReference(v string) {
	o.Reference = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AutoScaledInstance) GetState() State {
	if o == nil || IsNil(o.State) {
		var ret State
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScaledInstance) GetStateOk() (*State, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AutoScaledInstance) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given State and assigns it to the State field.
func (o *AutoScaledInstance) SetState(v State) {
	o.State = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *AutoScaledInstance) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScaledInstance) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *AutoScaledInstance) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *AutoScaledInstance) SetProductType(v string) {
	o.ProductType = &v
}

// GetincludesPrivateNetwork returns the includesPrivateNetwork field value if set, zero value otherwise.
func (o *AutoScaledInstance) GetincludesPrivateNetwork() bool {
	if o == nil || IsNil(o.includesPrivateNetwork) {
		var ret bool
		return ret
	}
	return *o.includesPrivateNetwork
}

// GetincludesPrivateNetworkOk returns a tuple with the includesPrivateNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScaledInstance) GetincludesPrivateNetworkOk() (*bool, bool) {
	if o == nil || IsNil(o.includesPrivateNetwork) {
		return nil, false
	}
	return o.includesPrivateNetwork, true
}

// HasincludesPrivateNetwork returns a boolean if a field has been set.
func (o *AutoScaledInstance) HasincludesPrivateNetwork() bool {
	if o != nil && !IsNil(o.includesPrivateNetwork) {
		return true
	}

	return false
}

// SetincludesPrivateNetwork gets a reference to the given bool and assigns it to the includesPrivateNetwork field.
func (o *AutoScaledInstance) SetincludesPrivateNetwork(v bool) {
	o.includesPrivateNetwork = &v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *AutoScaledInstance) GetIps() []Ip {
	if o == nil || IsNil(o.Ips) {
		var ret []Ip
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoScaledInstance) GetIpsOk() ([]Ip, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *AutoScaledInstance) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given []Ip and assigns it to the Ips field.
func (o *AutoScaledInstance) SetIps(v []Ip) {
	o.Ips = v
}

func (o AutoScaledInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoScaledInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !IsNil(o.includesPrivateNetwork) {
		toSerialize["hasPrivateNetwork"] = o.includesPrivateNetwork
	}
	if !IsNil(o.Ips) {
		toSerialize["ips"] = o.Ips
	}
	return toSerialize, nil
}

type NullableAutoScaledInstance struct {
	value *AutoScaledInstance
	isSet bool
}

func (v NullableAutoScaledInstance) Get() *AutoScaledInstance {
	return v.value
}

func (v *NullableAutoScaledInstance) Set(val *AutoScaledInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoScaledInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoScaledInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoScaledInstance(val *AutoScaledInstance) *NullableAutoScaledInstance {
	return &NullableAutoScaledInstance{value: val, isSet: true}
}

func (v NullableAutoScaledInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoScaledInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


