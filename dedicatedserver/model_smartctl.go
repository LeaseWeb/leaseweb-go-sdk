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

// checks if the Smartctl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Smartctl{}

// Smartctl struct for Smartctl
type Smartctl struct {
	AtaVersion *string `json:"ata_version,omitempty"`
	Attributes *Attributes `json:"attributes,omitempty"`
	DeviceModel *string `json:"device_model,omitempty"`
	ExecutionStatus *string `json:"execution_status,omitempty"`
	FirmwareVersion *string `json:"firmware_version,omitempty"`
	IsSas *bool `json:"is_sas,omitempty"`
	OverallHealth *string `json:"overall_health,omitempty"`
	Rpm *string `json:"rpm,omitempty"`
	SataVersion *string `json:"sata_version,omitempty"`
	SectorSize *string `json:"sector_size,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty"`
	SmartErrorLog *string `json:"smart_error_log,omitempty"`
	SmartSupport *SmartSupport `json:"smart_support,omitempty"`
	SmartctlVersion *string `json:"smartctl_version,omitempty"`
	UserCapacity *string `json:"user_capacity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Smartctl Smartctl

// NewSmartctl instantiates a new Smartctl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartctl() *Smartctl {
	this := Smartctl{}
	return &this
}

// NewSmartctlWithDefaults instantiates a new Smartctl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartctlWithDefaults() *Smartctl {
	this := Smartctl{}
	return &this
}

// GetAtaVersion returns the AtaVersion field value if set, zero value otherwise.
func (o *Smartctl) GetAtaVersion() string {
	if o == nil || IsNil(o.AtaVersion) {
		var ret string
		return ret
	}
	return *o.AtaVersion
}

// GetAtaVersionOk returns a tuple with the AtaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Smartctl) GetAtaVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AtaVersion) {
		return nil, false
	}
	return o.AtaVersion, true
}

// HasAtaVersion returns a boolean if a field has been set.
func (o *Smartctl) HasAtaVersion() bool {
	if o != nil && !IsNil(o.AtaVersion) {
		return true
	}

	return false
}

// SetAtaVersion gets a reference to the given string and assigns it to the AtaVersion field.
func (o *Smartctl) SetAtaVersion(v string) {
	o.AtaVersion = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Smartctl) GetAttributes() Attributes {
	if o == nil || IsNil(o.Attributes) {
		var ret Attributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Smartctl) GetAttributesOk() (*Attributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Smartctl) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Attributes and assigns it to the Attributes field.
func (o *Smartctl) SetAttributes(v Attributes) {
	o.Attributes = &v
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *Smartctl) GetDeviceModel() string {
	if o == nil || IsNil(o.DeviceModel) {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Smartctl) GetDeviceModelOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceModel) {
		return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *Smartctl) HasDeviceModel() bool {
	if o != nil && !IsNil(o.DeviceModel) {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *Smartctl) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetExecutionStatus returns the ExecutionStatus field value if set, zero value otherwise.
func (o *Smartctl) GetExecutionStatus() string {
	if o == nil || IsNil(o.ExecutionStatus) {
		var ret string
		return ret
	}
	return *o.ExecutionStatus
}

// GetExecutionStatusOk returns a tuple with the ExecutionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Smartctl) GetExecutionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionStatus) {
		return nil, false
	}
	return o.ExecutionStatus, true
}

// HasExecutionStatus returns a boolean if a field has been set.
func (o *Smartctl) HasExecutionStatus() bool {
	if o != nil && !IsNil(o.ExecutionStatus) {
		return true
	}

	return false
}

// SetExecutionStatus gets a reference to the given string and assigns it to the ExecutionStatus field.
func (o *Smartctl) SetExecutionStatus(v string) {
	o.ExecutionStatus = &v
}

// GetFirmwareVersion returns the FirmwareVersion field value if set, zero value otherwise.
func (o *Smartctl) GetFirmwareVersion() string {
	if o == nil || IsNil(o.FirmwareVersion) {
		var ret string
		return ret
	}
	return *o.FirmwareVersion
}

// GetFirmwareVersionOk returns a tuple with the FirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Smartctl) GetFirmwareVersionOk() (*string, bool) {
	if o == nil || IsNil(o.FirmwareVersion) {
		return nil, false
	}
	return o.FirmwareVersion, true
}

// HasFirmwareVersion returns a boolean if a field has been set.
func (o *Smartctl) HasFirmwareVersion() bool {
	if o != nil && !IsNil(o.FirmwareVersion) {
		return true
	}

	return false
}

// SetFirmwareVersion gets a reference to the given string and assigns it to the FirmwareVersion field.
func (o *Smartctl) SetFirmwareVersion(v string) {
	o.FirmwareVersion = &v
}

// GetIsSas returns the IsSas field value if set, zero value otherwise.
func (o *Smartctl) GetIsSas() bool {
	if o == nil || IsNil(o.IsSas) {
		var ret bool
		return ret
	}
	return *o.IsSas
}

// GetIsSasOk returns a tuple with the IsSas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Smartctl) GetIsSasOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSas) {
		return nil, false
	}
	return o.IsSas, true
}

// HasIsSas returns a boolean if a field has been set.
func (o *Smartctl) HasIsSas() bool {
	if o != nil && !IsNil(o.IsSas) {
		return true
	}

	return false
}

// SetIsSas gets a reference to the given bool and assigns it to the IsSas field.
func (o *Smartctl) SetIsSas(v bool) {
	o.IsSas = &v
}

// GetOverallHealth returns the OverallHealth field value if set, zero value otherwise.
func (o *Smartctl) GetOverallHealth() string {
	if o == nil || IsNil(o.OverallHealth) {
		var ret string
		return ret
	}
	return *o.OverallHealth
}

// GetOverallHealthOk returns a tuple with the OverallHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Smartctl) GetOverallHealthOk() (*string, bool) {
	if o == nil || IsNil(o.OverallHealth) {
		return nil, false
	}
	return o.OverallHealth, true
}

// HasOverallHealth returns a boolean if a field has been set.
func (o *Smartctl) HasOverallHealth() bool {
	if o != nil && !IsNil(o.OverallHealth) {
		return true
	}

	return false
}

// SetOverallHealth gets a reference to the given string and assigns it to the OverallHealth field.
func (o *Smartctl) SetOverallHealth(v string) {
	o.OverallHealth = &v
}

// GetRpm returns the Rpm field value if set, zero value otherwise.
func (o *Smartctl) GetRpm() string {
	if o == nil || IsNil(o.Rpm) {
		var ret string
		return ret
	}
	return *o.Rpm
}

// GetRpmOk returns a tuple with the Rpm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Smartctl) GetRpmOk() (*string, bool) {
	if o == nil || IsNil(o.Rpm) {
		return nil, false
	}
	return o.Rpm, true
}

// HasRpm returns a boolean if a field has been set.
func (o *Smartctl) HasRpm() bool {
	if o != nil && !IsNil(o.Rpm) {
		return true
	}

	return false
}

// SetRpm gets a reference to the given string and assigns it to the Rpm field.
func (o *Smartctl) SetRpm(v string) {
	o.Rpm = &v
}

// GetSataVersion returns the SataVersion field value if set, zero value otherwise.
func (o *Smartctl) GetSataVersion() string {
	if o == nil || IsNil(o.SataVersion) {
		var ret string
		return ret
	}
	return *o.SataVersion
}

// GetSataVersionOk returns a tuple with the SataVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Smartctl) GetSataVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SataVersion) {
		return nil, false
	}
	return o.SataVersion, true
}

// HasSataVersion returns a boolean if a field has been set.
func (o *Smartctl) HasSataVersion() bool {
	if o != nil && !IsNil(o.SataVersion) {
		return true
	}

	return false
}

// SetSataVersion gets a reference to the given string and assigns it to the SataVersion field.
func (o *Smartctl) SetSataVersion(v string) {
	o.SataVersion = &v
}

// GetSectorSize returns the SectorSize field value if set, zero value otherwise.
func (o *Smartctl) GetSectorSize() string {
	if o == nil || IsNil(o.SectorSize) {
		var ret string
		return ret
	}
	return *o.SectorSize
}

// GetSectorSizeOk returns a tuple with the SectorSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Smartctl) GetSectorSizeOk() (*string, bool) {
	if o == nil || IsNil(o.SectorSize) {
		return nil, false
	}
	return o.SectorSize, true
}

// HasSectorSize returns a boolean if a field has been set.
func (o *Smartctl) HasSectorSize() bool {
	if o != nil && !IsNil(o.SectorSize) {
		return true
	}

	return false
}

// SetSectorSize gets a reference to the given string and assigns it to the SectorSize field.
func (o *Smartctl) SetSectorSize(v string) {
	o.SectorSize = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *Smartctl) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Smartctl) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *Smartctl) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *Smartctl) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSmartErrorLog returns the SmartErrorLog field value if set, zero value otherwise.
func (o *Smartctl) GetSmartErrorLog() string {
	if o == nil || IsNil(o.SmartErrorLog) {
		var ret string
		return ret
	}
	return *o.SmartErrorLog
}

// GetSmartErrorLogOk returns a tuple with the SmartErrorLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Smartctl) GetSmartErrorLogOk() (*string, bool) {
	if o == nil || IsNil(o.SmartErrorLog) {
		return nil, false
	}
	return o.SmartErrorLog, true
}

// HasSmartErrorLog returns a boolean if a field has been set.
func (o *Smartctl) HasSmartErrorLog() bool {
	if o != nil && !IsNil(o.SmartErrorLog) {
		return true
	}

	return false
}

// SetSmartErrorLog gets a reference to the given string and assigns it to the SmartErrorLog field.
func (o *Smartctl) SetSmartErrorLog(v string) {
	o.SmartErrorLog = &v
}

// GetSmartSupport returns the SmartSupport field value if set, zero value otherwise.
func (o *Smartctl) GetSmartSupport() SmartSupport {
	if o == nil || IsNil(o.SmartSupport) {
		var ret SmartSupport
		return ret
	}
	return *o.SmartSupport
}

// GetSmartSupportOk returns a tuple with the SmartSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Smartctl) GetSmartSupportOk() (*SmartSupport, bool) {
	if o == nil || IsNil(o.SmartSupport) {
		return nil, false
	}
	return o.SmartSupport, true
}

// HasSmartSupport returns a boolean if a field has been set.
func (o *Smartctl) HasSmartSupport() bool {
	if o != nil && !IsNil(o.SmartSupport) {
		return true
	}

	return false
}

// SetSmartSupport gets a reference to the given SmartSupport and assigns it to the SmartSupport field.
func (o *Smartctl) SetSmartSupport(v SmartSupport) {
	o.SmartSupport = &v
}

// GetSmartctlVersion returns the SmartctlVersion field value if set, zero value otherwise.
func (o *Smartctl) GetSmartctlVersion() string {
	if o == nil || IsNil(o.SmartctlVersion) {
		var ret string
		return ret
	}
	return *o.SmartctlVersion
}

// GetSmartctlVersionOk returns a tuple with the SmartctlVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Smartctl) GetSmartctlVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SmartctlVersion) {
		return nil, false
	}
	return o.SmartctlVersion, true
}

// HasSmartctlVersion returns a boolean if a field has been set.
func (o *Smartctl) HasSmartctlVersion() bool {
	if o != nil && !IsNil(o.SmartctlVersion) {
		return true
	}

	return false
}

// SetSmartctlVersion gets a reference to the given string and assigns it to the SmartctlVersion field.
func (o *Smartctl) SetSmartctlVersion(v string) {
	o.SmartctlVersion = &v
}

// GetUserCapacity returns the UserCapacity field value if set, zero value otherwise.
func (o *Smartctl) GetUserCapacity() string {
	if o == nil || IsNil(o.UserCapacity) {
		var ret string
		return ret
	}
	return *o.UserCapacity
}

// GetUserCapacityOk returns a tuple with the UserCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Smartctl) GetUserCapacityOk() (*string, bool) {
	if o == nil || IsNil(o.UserCapacity) {
		return nil, false
	}
	return o.UserCapacity, true
}

// HasUserCapacity returns a boolean if a field has been set.
func (o *Smartctl) HasUserCapacity() bool {
	if o != nil && !IsNil(o.UserCapacity) {
		return true
	}

	return false
}

// SetUserCapacity gets a reference to the given string and assigns it to the UserCapacity field.
func (o *Smartctl) SetUserCapacity(v string) {
	o.UserCapacity = &v
}

func (o Smartctl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Smartctl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AtaVersion) {
		toSerialize["ata_version"] = o.AtaVersion
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.DeviceModel) {
		toSerialize["device_model"] = o.DeviceModel
	}
	if !IsNil(o.ExecutionStatus) {
		toSerialize["execution_status"] = o.ExecutionStatus
	}
	if !IsNil(o.FirmwareVersion) {
		toSerialize["firmware_version"] = o.FirmwareVersion
	}
	if !IsNil(o.IsSas) {
		toSerialize["is_sas"] = o.IsSas
	}
	if !IsNil(o.OverallHealth) {
		toSerialize["overall_health"] = o.OverallHealth
	}
	if !IsNil(o.Rpm) {
		toSerialize["rpm"] = o.Rpm
	}
	if !IsNil(o.SataVersion) {
		toSerialize["sata_version"] = o.SataVersion
	}
	if !IsNil(o.SectorSize) {
		toSerialize["sector_size"] = o.SectorSize
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if !IsNil(o.SmartErrorLog) {
		toSerialize["smart_error_log"] = o.SmartErrorLog
	}
	if !IsNil(o.SmartSupport) {
		toSerialize["smart_support"] = o.SmartSupport
	}
	if !IsNil(o.SmartctlVersion) {
		toSerialize["smartctl_version"] = o.SmartctlVersion
	}
	if !IsNil(o.UserCapacity) {
		toSerialize["user_capacity"] = o.UserCapacity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Smartctl) UnmarshalJSON(data []byte) (err error) {
	varSmartctl := _Smartctl{}

	err = json.Unmarshal(data, &varSmartctl)

	if err != nil {
		return err
	}

	*o = Smartctl(varSmartctl)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ata_version")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "device_model")
		delete(additionalProperties, "execution_status")
		delete(additionalProperties, "firmware_version")
		delete(additionalProperties, "is_sas")
		delete(additionalProperties, "overall_health")
		delete(additionalProperties, "rpm")
		delete(additionalProperties, "sata_version")
		delete(additionalProperties, "sector_size")
		delete(additionalProperties, "serial_number")
		delete(additionalProperties, "smart_error_log")
		delete(additionalProperties, "smart_support")
		delete(additionalProperties, "smartctl_version")
		delete(additionalProperties, "user_capacity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSmartctl struct {
	value *Smartctl
	isSet bool
}

func (v NullableSmartctl) Get() *Smartctl {
	return v.value
}

func (v *NullableSmartctl) Set(val *Smartctl) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartctl) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartctl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartctl(val *Smartctl) *NullableSmartctl {
	return &NullableSmartctl{value: val, isSet: true}
}

func (v NullableSmartctl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartctl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


