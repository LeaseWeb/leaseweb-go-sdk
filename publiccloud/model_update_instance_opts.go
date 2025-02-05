/*
Public Clouds

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publiccloud

import (
	"encoding/json"
)

// checks if the UpdateInstanceOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInstanceOpts{}

// UpdateInstanceOpts struct for UpdateInstanceOpts
type UpdateInstanceOpts struct {
	Type *TypeName `json:"type,omitempty"`
	// An identifying name you can refer to the instance
	Reference *string `json:"reference,omitempty"`
	ContractType *ContractType `json:"contractType,omitempty"`
	ContractTerm *ContractTerm `json:"contractTerm,omitempty"`
	BillingFrequency *BillingFrequency `json:"billingFrequency,omitempty"`
	// The root disk's size in GB. Must be at least 5 GB for Linux and FreeBSD instances and 50 GB for Windows instances
	RootDiskSize *int32 `json:"rootDiskSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateInstanceOpts UpdateInstanceOpts

// NewUpdateInstanceOpts instantiates a new UpdateInstanceOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInstanceOpts() *UpdateInstanceOpts {
	this := UpdateInstanceOpts{}
	return &this
}

// NewUpdateInstanceOptsWithDefaults instantiates a new UpdateInstanceOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInstanceOptsWithDefaults() *UpdateInstanceOpts {
	this := UpdateInstanceOpts{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateInstanceOpts) GetType() TypeName {
	if o == nil || IsNil(o.Type) {
		var ret TypeName
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceOpts) GetTypeOk() (*TypeName, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateInstanceOpts) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeName and assigns it to the Type field.
func (o *UpdateInstanceOpts) SetType(v TypeName) {
	o.Type = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *UpdateInstanceOpts) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceOpts) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *UpdateInstanceOpts) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *UpdateInstanceOpts) SetReference(v string) {
	o.Reference = &v
}

// GetContractType returns the ContractType field value if set, zero value otherwise.
func (o *UpdateInstanceOpts) GetContractType() ContractType {
	if o == nil || IsNil(o.ContractType) {
		var ret ContractType
		return ret
	}
	return *o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceOpts) GetContractTypeOk() (*ContractType, bool) {
	if o == nil || IsNil(o.ContractType) {
		return nil, false
	}
	return o.ContractType, true
}

// HasContractType returns a boolean if a field has been set.
func (o *UpdateInstanceOpts) HasContractType() bool {
	if o != nil && !IsNil(o.ContractType) {
		return true
	}

	return false
}

// SetContractType gets a reference to the given ContractType and assigns it to the ContractType field.
func (o *UpdateInstanceOpts) SetContractType(v ContractType) {
	o.ContractType = &v
}

// GetContractTerm returns the ContractTerm field value if set, zero value otherwise.
func (o *UpdateInstanceOpts) GetContractTerm() ContractTerm {
	if o == nil || IsNil(o.ContractTerm) {
		var ret ContractTerm
		return ret
	}
	return *o.ContractTerm
}

// GetContractTermOk returns a tuple with the ContractTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceOpts) GetContractTermOk() (*ContractTerm, bool) {
	if o == nil || IsNil(o.ContractTerm) {
		return nil, false
	}
	return o.ContractTerm, true
}

// HasContractTerm returns a boolean if a field has been set.
func (o *UpdateInstanceOpts) HasContractTerm() bool {
	if o != nil && !IsNil(o.ContractTerm) {
		return true
	}

	return false
}

// SetContractTerm gets a reference to the given ContractTerm and assigns it to the ContractTerm field.
func (o *UpdateInstanceOpts) SetContractTerm(v ContractTerm) {
	o.ContractTerm = &v
}

// GetBillingFrequency returns the BillingFrequency field value if set, zero value otherwise.
func (o *UpdateInstanceOpts) GetBillingFrequency() BillingFrequency {
	if o == nil || IsNil(o.BillingFrequency) {
		var ret BillingFrequency
		return ret
	}
	return *o.BillingFrequency
}

// GetBillingFrequencyOk returns a tuple with the BillingFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceOpts) GetBillingFrequencyOk() (*BillingFrequency, bool) {
	if o == nil || IsNil(o.BillingFrequency) {
		return nil, false
	}
	return o.BillingFrequency, true
}

// HasBillingFrequency returns a boolean if a field has been set.
func (o *UpdateInstanceOpts) HasBillingFrequency() bool {
	if o != nil && !IsNil(o.BillingFrequency) {
		return true
	}

	return false
}

// SetBillingFrequency gets a reference to the given BillingFrequency and assigns it to the BillingFrequency field.
func (o *UpdateInstanceOpts) SetBillingFrequency(v BillingFrequency) {
	o.BillingFrequency = &v
}

// GetRootDiskSize returns the RootDiskSize field value if set, zero value otherwise.
func (o *UpdateInstanceOpts) GetRootDiskSize() int32 {
	if o == nil || IsNil(o.RootDiskSize) {
		var ret int32
		return ret
	}
	return *o.RootDiskSize
}

// GetRootDiskSizeOk returns a tuple with the RootDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceOpts) GetRootDiskSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.RootDiskSize) {
		return nil, false
	}
	return o.RootDiskSize, true
}

// HasRootDiskSize returns a boolean if a field has been set.
func (o *UpdateInstanceOpts) HasRootDiskSize() bool {
	if o != nil && !IsNil(o.RootDiskSize) {
		return true
	}

	return false
}

// SetRootDiskSize gets a reference to the given int32 and assigns it to the RootDiskSize field.
func (o *UpdateInstanceOpts) SetRootDiskSize(v int32) {
	o.RootDiskSize = &v
}

func (o UpdateInstanceOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateInstanceOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.ContractType) {
		toSerialize["contractType"] = o.ContractType
	}
	if !IsNil(o.ContractTerm) {
		toSerialize["contractTerm"] = o.ContractTerm
	}
	if !IsNil(o.BillingFrequency) {
		toSerialize["billingFrequency"] = o.BillingFrequency
	}
	if !IsNil(o.RootDiskSize) {
		toSerialize["rootDiskSize"] = o.RootDiskSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateInstanceOpts) UnmarshalJSON(data []byte) (err error) {
	varUpdateInstanceOpts := _UpdateInstanceOpts{}

	err = json.Unmarshal(data, &varUpdateInstanceOpts)

	if err != nil {
		return err
	}

	*o = UpdateInstanceOpts(varUpdateInstanceOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "reference")
		delete(additionalProperties, "contractType")
		delete(additionalProperties, "contractTerm")
		delete(additionalProperties, "billingFrequency")
		delete(additionalProperties, "rootDiskSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateInstanceOpts struct {
	value *UpdateInstanceOpts
	isSet bool
}

func (v NullableUpdateInstanceOpts) Get() *UpdateInstanceOpts {
	return v.value
}

func (v *NullableUpdateInstanceOpts) Set(val *UpdateInstanceOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInstanceOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInstanceOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInstanceOpts(val *UpdateInstanceOpts) *NullableUpdateInstanceOpts {
	return &NullableUpdateInstanceOpts{value: val, isSet: true}
}

func (v NullableUpdateInstanceOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateInstanceOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


