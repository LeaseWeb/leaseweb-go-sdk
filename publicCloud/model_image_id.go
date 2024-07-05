/*
LeaseWeb API for launching and managing Public Cloud instances

> The base URL for this API is: **https://api.leaseweb.com/publicCloud/v1/_**  This API provides ways to launch and manage Public Cloud instances.  <div class=\"badge\">BETA</div> This API is in BETA. Documentation might be incorrect or incomplete. Functionality might change with the final release.>

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package publicCloud

import (
	"encoding/json"
	"fmt"
)

// ImageId Image ID
type ImageId string

// List of imageId
const (
	ALMALINUX_8_64_BIT ImageId = "ALMALINUX_8_64BIT"
	ALMALINUX_9_64_BIT ImageId = "ALMALINUX_9_64BIT"
	ARCH_LINUX_64_BIT ImageId = "ARCH_LINUX_64BIT"
	CENTOS_7_64_BIT ImageId = "CENTOS_7_64BIT"
	DEBIAN_10_64_BIT ImageId = "DEBIAN_10_64BIT"
	DEBIAN_11_64_BIT ImageId = "DEBIAN_11_64BIT"
	DEBIAN_12_64_BIT ImageId = "DEBIAN_12_64BIT"
	FREEBSD_13_64_BIT ImageId = "FREEBSD_13_64BIT"
	FREEBSD_14_64_BIT ImageId = "FREEBSD_14_64BIT"
	ROCKY_LINUX_8_64_BIT ImageId = "ROCKY_LINUX_8_64BIT"
	ROCKY_LINUX_9_64_BIT ImageId = "ROCKY_LINUX_9_64BIT"
	UBUNTU_20_04_64_BIT ImageId = "UBUNTU_20_04_64BIT"
	UBUNTU_22_04_64_BIT ImageId = "UBUNTU_22_04_64BIT"
	UBUNTU_24_04_64_BIT ImageId = "UBUNTU_24_04_64BIT"
	WINDOWS_SERVER_2016_STANDARD_64_BIT ImageId = "WINDOWS_SERVER_2016_STANDARD_64BIT"
	WINDOWS_SERVER_2019_STANDARD_64_BIT ImageId = "WINDOWS_SERVER_2019_STANDARD_64BIT"
	WINDOWS_SERVER_2022_STANDARD_64_BIT ImageId = "WINDOWS_SERVER_2022_STANDARD_64BIT"
)

// All allowed values of ImageId enum
var AllowedImageIdEnumValues = []ImageId{
	"ALMALINUX_8_64BIT",
	"ALMALINUX_9_64BIT",
	"ARCH_LINUX_64BIT",
	"CENTOS_7_64BIT",
	"DEBIAN_10_64BIT",
	"DEBIAN_11_64BIT",
	"DEBIAN_12_64BIT",
	"FREEBSD_13_64BIT",
	"FREEBSD_14_64BIT",
	"ROCKY_LINUX_8_64BIT",
	"ROCKY_LINUX_9_64BIT",
	"UBUNTU_20_04_64BIT",
	"UBUNTU_22_04_64BIT",
	"UBUNTU_24_04_64BIT",
	"WINDOWS_SERVER_2016_STANDARD_64BIT",
	"WINDOWS_SERVER_2019_STANDARD_64BIT",
	"WINDOWS_SERVER_2022_STANDARD_64BIT",
}

func (v *ImageId) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImageId(value)
	for _, existing := range AllowedImageIdEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImageId", value)
}

// NewImageIdFromValue returns a pointer to a valid ImageId
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImageIdFromValue(v string) (*ImageId, error) {
	ev := ImageId(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImageId: valid values are %v", v, AllowedImageIdEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageId) IsValid() bool {
	for _, existing := range AllowedImageIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to imageId value
func (v ImageId) Ptr() *ImageId {
	return &v
}

type NullableImageId struct {
	value *ImageId
	isSet bool
}

func (v NullableImageId) Get() *ImageId {
	return v.value
}

func (v *NullableImageId) Set(val *ImageId) {
	v.value = val
	v.isSet = true
}

func (v NullableImageId) IsSet() bool {
	return v.isSet
}

func (v *NullableImageId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageId(val *ImageId) *NullableImageId {
	return &NullableImageId{value: val, isSet: true}
}

func (v NullableImageId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

