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

// checks if the UpdateTargetGroupOpts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTargetGroupOpts{}

// UpdateTargetGroupOpts struct for UpdateTargetGroupOpts
type UpdateTargetGroupOpts struct {
	// The name of the target group
	Name *string `json:"name,omitempty"`
	// The port of the target group
	Port *int32 `json:"port,omitempty"`
	HealthCheck *HealthCheckOpts `json:"healthCheck,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateTargetGroupOpts UpdateTargetGroupOpts

// NewUpdateTargetGroupOpts instantiates a new UpdateTargetGroupOpts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTargetGroupOpts() *UpdateTargetGroupOpts {
	this := UpdateTargetGroupOpts{}
	return &this
}

// NewUpdateTargetGroupOptsWithDefaults instantiates a new UpdateTargetGroupOpts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTargetGroupOptsWithDefaults() *UpdateTargetGroupOpts {
	this := UpdateTargetGroupOpts{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateTargetGroupOpts) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTargetGroupOpts) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateTargetGroupOpts) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateTargetGroupOpts) SetName(v string) {
	o.Name = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *UpdateTargetGroupOpts) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTargetGroupOpts) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *UpdateTargetGroupOpts) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *UpdateTargetGroupOpts) SetPort(v int32) {
	o.Port = &v
}

// GetHealthCheck returns the HealthCheck field value if set, zero value otherwise.
func (o *UpdateTargetGroupOpts) GetHealthCheck() HealthCheckOpts {
	if o == nil || IsNil(o.HealthCheck) {
		var ret HealthCheckOpts
		return ret
	}
	return *o.HealthCheck
}

// GetHealthCheckOk returns a tuple with the HealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTargetGroupOpts) GetHealthCheckOk() (*HealthCheckOpts, bool) {
	if o == nil || IsNil(o.HealthCheck) {
		return nil, false
	}
	return o.HealthCheck, true
}

// HasHealthCheck returns a boolean if a field has been set.
func (o *UpdateTargetGroupOpts) HasHealthCheck() bool {
	if o != nil && !IsNil(o.HealthCheck) {
		return true
	}

	return false
}

// SetHealthCheck gets a reference to the given HealthCheckOpts and assigns it to the HealthCheck field.
func (o *UpdateTargetGroupOpts) SetHealthCheck(v HealthCheckOpts) {
	o.HealthCheck = &v
}

func (o UpdateTargetGroupOpts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTargetGroupOpts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.HealthCheck) {
		toSerialize["healthCheck"] = o.HealthCheck
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateTargetGroupOpts) UnmarshalJSON(data []byte) (err error) {
	varUpdateTargetGroupOpts := _UpdateTargetGroupOpts{}

	err = json.Unmarshal(data, &varUpdateTargetGroupOpts)

	if err != nil {
		return err
	}

	*o = UpdateTargetGroupOpts(varUpdateTargetGroupOpts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "port")
		delete(additionalProperties, "healthCheck")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateTargetGroupOpts struct {
	value *UpdateTargetGroupOpts
	isSet bool
}

func (v NullableUpdateTargetGroupOpts) Get() *UpdateTargetGroupOpts {
	return v.value
}

func (v *NullableUpdateTargetGroupOpts) Set(val *UpdateTargetGroupOpts) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTargetGroupOpts) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTargetGroupOpts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTargetGroupOpts(val *UpdateTargetGroupOpts) *NullableUpdateTargetGroupOpts {
	return &NullableUpdateTargetGroupOpts{value: val, isSet: true}
}

func (v NullableUpdateTargetGroupOpts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTargetGroupOpts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


