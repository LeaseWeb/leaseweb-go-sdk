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

// checks if the NetworkEquipment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkEquipment{}

// NetworkEquipment struct for NetworkEquipment
type NetworkEquipment struct {
	Contract *Contract `json:"contract,omitempty"`
	FeatureAvailability *FeatureAvailability `json:"featureAvailability,omitempty"`
	// Id of the network equipment
	Id *string `json:"id,omitempty"`
	Location *Location `json:"location,omitempty"`
	// The name of the network equipment
	Name *string `json:"name,omitempty"`
	// The network equipment type
	Type *string `json:"type,omitempty"`
	NetworkInterfaces *NetworkInterfaces `json:"networkInterfaces,omitempty"`
	// List of ports that can be used to manage power of the network equipment
	PowerPorts []Port `json:"powerPorts,omitempty"`
	Rack *Rack `json:"rack,omitempty"`
	// Serial number of network equipment
	SerialNumber *string `json:"serialNumber,omitempty"`
	Specs *NetworkEquipmentSpecs `json:"specs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkEquipment NetworkEquipment

// NewNetworkEquipment instantiates a new NetworkEquipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkEquipment() *NetworkEquipment {
	this := NetworkEquipment{}
	return &this
}

// NewNetworkEquipmentWithDefaults instantiates a new NetworkEquipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkEquipmentWithDefaults() *NetworkEquipment {
	this := NetworkEquipment{}
	return &this
}

// GetContract returns the Contract field value if set, zero value otherwise.
func (o *NetworkEquipment) GetContract() Contract {
	if o == nil || IsNil(o.Contract) {
		var ret Contract
		return ret
	}
	return *o.Contract
}

// GetContractOk returns a tuple with the Contract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEquipment) GetContractOk() (*Contract, bool) {
	if o == nil || IsNil(o.Contract) {
		return nil, false
	}
	return o.Contract, true
}

// HasContract returns a boolean if a field has been set.
func (o *NetworkEquipment) HasContract() bool {
	if o != nil && !IsNil(o.Contract) {
		return true
	}

	return false
}

// SetContract gets a reference to the given Contract and assigns it to the Contract field.
func (o *NetworkEquipment) SetContract(v Contract) {
	o.Contract = &v
}

// GetFeatureAvailability returns the FeatureAvailability field value if set, zero value otherwise.
func (o *NetworkEquipment) GetFeatureAvailability() FeatureAvailability {
	if o == nil || IsNil(o.FeatureAvailability) {
		var ret FeatureAvailability
		return ret
	}
	return *o.FeatureAvailability
}

// GetFeatureAvailabilityOk returns a tuple with the FeatureAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEquipment) GetFeatureAvailabilityOk() (*FeatureAvailability, bool) {
	if o == nil || IsNil(o.FeatureAvailability) {
		return nil, false
	}
	return o.FeatureAvailability, true
}

// HasFeatureAvailability returns a boolean if a field has been set.
func (o *NetworkEquipment) HasFeatureAvailability() bool {
	if o != nil && !IsNil(o.FeatureAvailability) {
		return true
	}

	return false
}

// SetFeatureAvailability gets a reference to the given FeatureAvailability and assigns it to the FeatureAvailability field.
func (o *NetworkEquipment) SetFeatureAvailability(v FeatureAvailability) {
	o.FeatureAvailability = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkEquipment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEquipment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkEquipment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworkEquipment) SetId(v string) {
	o.Id = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *NetworkEquipment) GetLocation() Location {
	if o == nil || IsNil(o.Location) {
		var ret Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEquipment) GetLocationOk() (*Location, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *NetworkEquipment) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location and assigns it to the Location field.
func (o *NetworkEquipment) SetLocation(v Location) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkEquipment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEquipment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkEquipment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkEquipment) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkEquipment) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEquipment) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkEquipment) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworkEquipment) SetType(v string) {
	o.Type = &v
}

// GetNetworkInterfaces returns the NetworkInterfaces field value if set, zero value otherwise.
func (o *NetworkEquipment) GetNetworkInterfaces() NetworkInterfaces {
	if o == nil || IsNil(o.NetworkInterfaces) {
		var ret NetworkInterfaces
		return ret
	}
	return *o.NetworkInterfaces
}

// GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEquipment) GetNetworkInterfacesOk() (*NetworkInterfaces, bool) {
	if o == nil || IsNil(o.NetworkInterfaces) {
		return nil, false
	}
	return o.NetworkInterfaces, true
}

// HasNetworkInterfaces returns a boolean if a field has been set.
func (o *NetworkEquipment) HasNetworkInterfaces() bool {
	if o != nil && !IsNil(o.NetworkInterfaces) {
		return true
	}

	return false
}

// SetNetworkInterfaces gets a reference to the given NetworkInterfaces and assigns it to the NetworkInterfaces field.
func (o *NetworkEquipment) SetNetworkInterfaces(v NetworkInterfaces) {
	o.NetworkInterfaces = &v
}

// GetPowerPorts returns the PowerPorts field value if set, zero value otherwise.
func (o *NetworkEquipment) GetPowerPorts() []Port {
	if o == nil || IsNil(o.PowerPorts) {
		var ret []Port
		return ret
	}
	return o.PowerPorts
}

// GetPowerPortsOk returns a tuple with the PowerPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEquipment) GetPowerPortsOk() ([]Port, bool) {
	if o == nil || IsNil(o.PowerPorts) {
		return nil, false
	}
	return o.PowerPorts, true
}

// HasPowerPorts returns a boolean if a field has been set.
func (o *NetworkEquipment) HasPowerPorts() bool {
	if o != nil && !IsNil(o.PowerPorts) {
		return true
	}

	return false
}

// SetPowerPorts gets a reference to the given []Port and assigns it to the PowerPorts field.
func (o *NetworkEquipment) SetPowerPorts(v []Port) {
	o.PowerPorts = v
}

// GetRack returns the Rack field value if set, zero value otherwise.
func (o *NetworkEquipment) GetRack() Rack {
	if o == nil || IsNil(o.Rack) {
		var ret Rack
		return ret
	}
	return *o.Rack
}

// GetRackOk returns a tuple with the Rack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEquipment) GetRackOk() (*Rack, bool) {
	if o == nil || IsNil(o.Rack) {
		return nil, false
	}
	return o.Rack, true
}

// HasRack returns a boolean if a field has been set.
func (o *NetworkEquipment) HasRack() bool {
	if o != nil && !IsNil(o.Rack) {
		return true
	}

	return false
}

// SetRack gets a reference to the given Rack and assigns it to the Rack field.
func (o *NetworkEquipment) SetRack(v Rack) {
	o.Rack = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *NetworkEquipment) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEquipment) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *NetworkEquipment) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *NetworkEquipment) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSpecs returns the Specs field value if set, zero value otherwise.
func (o *NetworkEquipment) GetSpecs() NetworkEquipmentSpecs {
	if o == nil || IsNil(o.Specs) {
		var ret NetworkEquipmentSpecs
		return ret
	}
	return *o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEquipment) GetSpecsOk() (*NetworkEquipmentSpecs, bool) {
	if o == nil || IsNil(o.Specs) {
		return nil, false
	}
	return o.Specs, true
}

// HasSpecs returns a boolean if a field has been set.
func (o *NetworkEquipment) HasSpecs() bool {
	if o != nil && !IsNil(o.Specs) {
		return true
	}

	return false
}

// SetSpecs gets a reference to the given NetworkEquipmentSpecs and assigns it to the Specs field.
func (o *NetworkEquipment) SetSpecs(v NetworkEquipmentSpecs) {
	o.Specs = &v
}

func (o NetworkEquipment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkEquipment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contract) {
		toSerialize["contract"] = o.Contract
	}
	if !IsNil(o.FeatureAvailability) {
		toSerialize["featureAvailability"] = o.FeatureAvailability
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.NetworkInterfaces) {
		toSerialize["networkInterfaces"] = o.NetworkInterfaces
	}
	if !IsNil(o.PowerPorts) {
		toSerialize["powerPorts"] = o.PowerPorts
	}
	if !IsNil(o.Rack) {
		toSerialize["rack"] = o.Rack
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if !IsNil(o.Specs) {
		toSerialize["specs"] = o.Specs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkEquipment) UnmarshalJSON(data []byte) (err error) {
	varNetworkEquipment := _NetworkEquipment{}

	err = json.Unmarshal(data, &varNetworkEquipment)

	if err != nil {
		return err
	}

	*o = NetworkEquipment(varNetworkEquipment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "contract")
		delete(additionalProperties, "featureAvailability")
		delete(additionalProperties, "id")
		delete(additionalProperties, "location")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "networkInterfaces")
		delete(additionalProperties, "powerPorts")
		delete(additionalProperties, "rack")
		delete(additionalProperties, "serialNumber")
		delete(additionalProperties, "specs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkEquipment struct {
	value *NetworkEquipment
	isSet bool
}

func (v NullableNetworkEquipment) Get() *NetworkEquipment {
	return v.value
}

func (v *NullableNetworkEquipment) Set(val *NetworkEquipment) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkEquipment) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkEquipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkEquipment(val *NetworkEquipment) *NullableNetworkEquipment {
	return &NullableNetworkEquipment{value: val, isSet: true}
}

func (v NullableNetworkEquipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkEquipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


