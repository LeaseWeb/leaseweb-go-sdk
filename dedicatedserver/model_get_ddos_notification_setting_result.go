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

// checks if the GetDdosNotificationSettingResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDdosNotificationSettingResult{}

// GetDdosNotificationSettingResult struct for GetDdosNotificationSettingResult
type GetDdosNotificationSettingResult struct {
	// Email notifications for nulling events
	Nulling *string `json:"nulling,omitempty"`
	// Email notifications for scrubbing events
	Scrubbing *string `json:"scrubbing,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDdosNotificationSettingResult GetDdosNotificationSettingResult

// NewGetDdosNotificationSettingResult instantiates a new GetDdosNotificationSettingResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDdosNotificationSettingResult() *GetDdosNotificationSettingResult {
	this := GetDdosNotificationSettingResult{}
	return &this
}

// NewGetDdosNotificationSettingResultWithDefaults instantiates a new GetDdosNotificationSettingResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDdosNotificationSettingResultWithDefaults() *GetDdosNotificationSettingResult {
	this := GetDdosNotificationSettingResult{}
	return &this
}

// GetNulling returns the Nulling field value if set, zero value otherwise.
func (o *GetDdosNotificationSettingResult) GetNulling() string {
	if o == nil || IsNil(o.Nulling) {
		var ret string
		return ret
	}
	return *o.Nulling
}

// GetNullingOk returns a tuple with the Nulling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDdosNotificationSettingResult) GetNullingOk() (*string, bool) {
	if o == nil || IsNil(o.Nulling) {
		return nil, false
	}
	return o.Nulling, true
}

// HasNulling returns a boolean if a field has been set.
func (o *GetDdosNotificationSettingResult) HasNulling() bool {
	if o != nil && !IsNil(o.Nulling) {
		return true
	}

	return false
}

// SetNulling gets a reference to the given string and assigns it to the Nulling field.
func (o *GetDdosNotificationSettingResult) SetNulling(v string) {
	o.Nulling = &v
}

// GetScrubbing returns the Scrubbing field value if set, zero value otherwise.
func (o *GetDdosNotificationSettingResult) GetScrubbing() string {
	if o == nil || IsNil(o.Scrubbing) {
		var ret string
		return ret
	}
	return *o.Scrubbing
}

// GetScrubbingOk returns a tuple with the Scrubbing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDdosNotificationSettingResult) GetScrubbingOk() (*string, bool) {
	if o == nil || IsNil(o.Scrubbing) {
		return nil, false
	}
	return o.Scrubbing, true
}

// HasScrubbing returns a boolean if a field has been set.
func (o *GetDdosNotificationSettingResult) HasScrubbing() bool {
	if o != nil && !IsNil(o.Scrubbing) {
		return true
	}

	return false
}

// SetScrubbing gets a reference to the given string and assigns it to the Scrubbing field.
func (o *GetDdosNotificationSettingResult) SetScrubbing(v string) {
	o.Scrubbing = &v
}

func (o GetDdosNotificationSettingResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDdosNotificationSettingResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nulling) {
		toSerialize["nulling"] = o.Nulling
	}
	if !IsNil(o.Scrubbing) {
		toSerialize["scrubbing"] = o.Scrubbing
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDdosNotificationSettingResult) UnmarshalJSON(data []byte) (err error) {
	varGetDdosNotificationSettingResult := _GetDdosNotificationSettingResult{}

	err = json.Unmarshal(data, &varGetDdosNotificationSettingResult)

	if err != nil {
		return err
	}

	*o = GetDdosNotificationSettingResult(varGetDdosNotificationSettingResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "nulling")
		delete(additionalProperties, "scrubbing")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDdosNotificationSettingResult struct {
	value *GetDdosNotificationSettingResult
	isSet bool
}

func (v NullableGetDdosNotificationSettingResult) Get() *GetDdosNotificationSettingResult {
	return v.value
}

func (v *NullableGetDdosNotificationSettingResult) Set(val *GetDdosNotificationSettingResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDdosNotificationSettingResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDdosNotificationSettingResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDdosNotificationSettingResult(val *GetDdosNotificationSettingResult) *NullableGetDdosNotificationSettingResult {
	return &NullableGetDdosNotificationSettingResult{value: val, isSet: true}
}

func (v NullableGetDdosNotificationSettingResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDdosNotificationSettingResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


