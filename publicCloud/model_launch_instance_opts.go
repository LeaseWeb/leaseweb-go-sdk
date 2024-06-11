/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the LaunchInstanceOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LaunchInstanceOpts{}

// LaunchInstanceOpts struct for LaunchInstanceOpts
type LaunchInstanceOpts struct {
	// Region to launch the instance into
	Region string `json:"region"`
	// Instance type
	Type string `json:"type"`
	OperatingSystemId OperatingSystemId `json:"operatingSystemId"`
	// Market App ID that must be installed into the instance
	MarketAppId NullableString `json:"marketAppId,omitempty"`
	// An identifying name you can refer to the instance
	Reference *string `json:"reference,omitempty"`
	ContractType string `json:"contractType"`
	// Contract commitment. Used only when contract type is MONTHLY
	ContractTerm int32 `json:"contractTerm"`
	// How often you wish to be charged. Used only when contract type is MONTHLY. '1' means every month, '3' every three months and so on.
	BillingFrequency int32 `json:"billingFrequency"`
	// The root disk's size in GB. Must be at least 5 GB for Linux and FreeBSD instances and 50 GB for Windows instances
	RootDiskSize *int32 `json:"rootDiskSize,omitempty"`
	// The root disk's storage type
	RootDiskStorageType string `json:"rootDiskStorageType"`
	// Public SSH key to be installed into the instance. Must be used only on Linux/FreeBSD instances
	SshKey *string `json:"sshKey,omitempty"`
}

type _LaunchInstanceOpts LaunchInstanceOpts

// NewLaunchInstanceOpts instantiates a new LaunchInstanceOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLaunchInstanceOpts(region string, type_ string, operatingSystemId OperatingSystemId, contractType string, contractTerm int32, billingFrequency int32, rootDiskStorageType string) *LaunchInstanceOpts {
	this := LaunchInstanceOpts{}
	this.Region = region
	this.Type = type_
	this.OperatingSystemId = operatingSystemId
	this.ContractType = contractType
	this.ContractTerm = contractTerm
	this.BillingFrequency = billingFrequency
	this.RootDiskStorageType = rootDiskStorageType
	return &this
}

// NewLaunchInstanceOptsWithDefaults instantiates a new LaunchInstanceOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLaunchInstanceOptsWithDefaults() *LaunchInstanceOpts {
	this := LaunchInstanceOpts{}
	return &this
}

// GetRegion returns the Region field value
func (o *LaunchInstanceOpts) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *LaunchInstanceOpts) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *LaunchInstanceOpts) SetRegion(v string) {
	o.Region = v
}

// GetType returns the Type field value
func (o *LaunchInstanceOpts) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LaunchInstanceOpts) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LaunchInstanceOpts) SetType(v string) {
	o.Type = v
}

// GetOperatingSystemId returns the OperatingSystemId field value
func (o *LaunchInstanceOpts) GetOperatingSystemId() OperatingSystemId {
	if o == nil {
		var ret OperatingSystemId
		return ret
	}

	return o.OperatingSystemId
}

// GetOperatingSystemIdOk returns a tuple with the OperatingSystemId field value
// and a boolean to check if the value has been set.
func (o *LaunchInstanceOpts) GetOperatingSystemIdOk() (*OperatingSystemId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperatingSystemId, true
}

// SetOperatingSystemId sets field value
func (o *LaunchInstanceOpts) SetOperatingSystemId(v OperatingSystemId) {
	o.OperatingSystemId = v
}

// GetMarketAppId returns the MarketAppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LaunchInstanceOpts) GetMarketAppId() string {
	if o == nil || IsNil(o.MarketAppId.Get()) {
		var ret string
		return ret
	}
	return *o.MarketAppId.Get()
}

// GetMarketAppIdOk returns a tuple with the MarketAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LaunchInstanceOpts) GetMarketAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarketAppId.Get(), o.MarketAppId.IsSet()
}

// HasMarketAppId returns a boolean if a field has been set.
func (o *LaunchInstanceOpts) HasMarketAppId() bool {
	if o != nil && o.MarketAppId.IsSet() {
		return true
	}

	return false
}

// SetMarketAppId gets a reference to the given NullableString and assigns it to the MarketAppId field.
func (o *LaunchInstanceOpts) SetMarketAppId(v string) {
	o.MarketAppId.Set(&v)
}
// SetMarketAppIdNil sets the value for MarketAppId to be an explicit nil
func (o *LaunchInstanceOpts) SetMarketAppIdNil() {
	o.MarketAppId.Set(nil)
}

// UnsetMarketAppId ensures that no value is present for MarketAppId, not even an explicit nil
func (o *LaunchInstanceOpts) UnsetMarketAppId() {
	o.MarketAppId.Unset()
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *LaunchInstanceOpts) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchInstanceOpts) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *LaunchInstanceOpts) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *LaunchInstanceOpts) SetReference(v string) {
	o.Reference = &v
}

// GetContractType returns the ContractType field value
func (o *LaunchInstanceOpts) GetContractType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value
// and a boolean to check if the value has been set.
func (o *LaunchInstanceOpts) GetContractTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractType, true
}

// SetContractType sets field value
func (o *LaunchInstanceOpts) SetContractType(v string) {
	o.ContractType = v
}

// GetContractTerm returns the ContractTerm field value
func (o *LaunchInstanceOpts) GetContractTerm() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContractTerm
}

// GetContractTermOk returns a tuple with the ContractTerm field value
// and a boolean to check if the value has been set.
func (o *LaunchInstanceOpts) GetContractTermOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractTerm, true
}

// SetContractTerm sets field value
func (o *LaunchInstanceOpts) SetContractTerm(v int32) {
	o.ContractTerm = v
}

// GetBillingFrequency returns the BillingFrequency field value
func (o *LaunchInstanceOpts) GetBillingFrequency() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BillingFrequency
}

// GetBillingFrequencyOk returns a tuple with the BillingFrequency field value
// and a boolean to check if the value has been set.
func (o *LaunchInstanceOpts) GetBillingFrequencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingFrequency, true
}

// SetBillingFrequency sets field value
func (o *LaunchInstanceOpts) SetBillingFrequency(v int32) {
	o.BillingFrequency = v
}

// GetRootDiskSize returns the RootDiskSize field value if set, zero value otherwise.
func (o *LaunchInstanceOpts) GetRootDiskSize() int32 {
	if o == nil || IsNil(o.RootDiskSize) {
		var ret int32
		return ret
	}
	return *o.RootDiskSize
}

// GetRootDiskSizeOk returns a tuple with the RootDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchInstanceOpts) GetRootDiskSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.RootDiskSize) {
		return nil, false
	}
	return o.RootDiskSize, true
}

// HasRootDiskSize returns a boolean if a field has been set.
func (o *LaunchInstanceOpts) HasRootDiskSize() bool {
	if o != nil && !IsNil(o.RootDiskSize) {
		return true
	}

	return false
}

// SetRootDiskSize gets a reference to the given int32 and assigns it to the RootDiskSize field.
func (o *LaunchInstanceOpts) SetRootDiskSize(v int32) {
	o.RootDiskSize = &v
}

// GetRootDiskStorageType returns the RootDiskStorageType field value
func (o *LaunchInstanceOpts) GetRootDiskStorageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootDiskStorageType
}

// GetRootDiskStorageTypeOk returns a tuple with the RootDiskStorageType field value
// and a boolean to check if the value has been set.
func (o *LaunchInstanceOpts) GetRootDiskStorageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootDiskStorageType, true
}

// SetRootDiskStorageType sets field value
func (o *LaunchInstanceOpts) SetRootDiskStorageType(v string) {
	o.RootDiskStorageType = v
}

// GetSshKey returns the SshKey field value if set, zero value otherwise.
func (o *LaunchInstanceOpts) GetSshKey() string {
	if o == nil || IsNil(o.SshKey) {
		var ret string
		return ret
	}
	return *o.SshKey
}

// GetSshKeyOk returns a tuple with the SshKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchInstanceOpts) GetSshKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SshKey) {
		return nil, false
	}
	return o.SshKey, true
}

// HasSshKey returns a boolean if a field has been set.
func (o *LaunchInstanceOpts) HasSshKey() bool {
	if o != nil && !IsNil(o.SshKey) {
		return true
	}

	return false
}

// SetSshKey gets a reference to the given string and assigns it to the SshKey field.
func (o *LaunchInstanceOpts) SetSshKey(v string) {
	o.SshKey = &v
}

func (o LaunchInstanceOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LaunchInstanceOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["region"] = o.Region
	toSerialize["type"] = o.Type
	toSerialize["operatingSystemId"] = o.OperatingSystemId
	if o.MarketAppId.IsSet() {
		toSerialize["marketAppId"] = o.MarketAppId.Get()
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	toSerialize["contractType"] = o.ContractType
	toSerialize["contractTerm"] = o.ContractTerm
	toSerialize["billingFrequency"] = o.BillingFrequency
	if !IsNil(o.RootDiskSize) {
		toSerialize["rootDiskSize"] = o.RootDiskSize
	}
	toSerialize["rootDiskStorageType"] = o.RootDiskStorageType
	if !IsNil(o.SshKey) {
		toSerialize["sshKey"] = o.SshKey
	}
	return toSerialize, nil
}

func (o *LaunchInstanceOpts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"region",
		"type",
		"operatingSystemId",
		"contractType",
		"contractTerm",
		"billingFrequency",
		"rootDiskStorageType",
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

	varLaunchInstanceOpts := _LaunchInstanceOpts{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLaunchInstanceOpts)

	if err != nil {
		return err
	}

	*o = LaunchInstanceOpts(varLaunchInstanceOpts)

	return err
}

type NullableLaunchInstanceOpts struct {
	value *LaunchInstanceOpts
	isSet bool
}

func (v NullableLaunchInstanceOpts) Get() *LaunchInstanceOpts {
	return v.value
}

func (v *NullableLaunchInstanceOpts) Set(val *LaunchInstanceOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableLaunchInstanceOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableLaunchInstanceOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLaunchInstanceOpts(val *LaunchInstanceOpts) *NullableLaunchInstanceOpts {
	return &NullableLaunchInstanceOpts{value: val, isSet: true}
}

func (v NullableLaunchInstanceOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLaunchInstanceOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


