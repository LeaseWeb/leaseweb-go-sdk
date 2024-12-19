/*
DNS

>The base URL for this API is: **https://api.leaseweb.com/hosting/v2**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
	"fmt"
)

// checks if the LdResourceRecordSetDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdResourceRecordSetDetails{}

// LdResourceRecordSetDetails struct for LdResourceRecordSetDetails
type LdResourceRecordSetDetails struct {
	// Name of the resource record set
	Name string `json:"name"`
	Type LdResourceRecordSetType `json:"type"`
	GeoContent GeoContent `json:"geoContent"`
	Ttl Ttl `json:"ttl"`
	// May the set be edited
	Editable bool `json:"editable"`
	Links LdLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _LdResourceRecordSetDetails LdResourceRecordSetDetails

// NewLdResourceRecordSetDetails instantiates a new LdResourceRecordSetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdResourceRecordSetDetails(name string, type_ LdResourceRecordSetType, geoContent GeoContent, ttl Ttl, editable bool, links LdLinks) *LdResourceRecordSetDetails {
	this := LdResourceRecordSetDetails{}
	this.Name = name
	this.Type = type_
	this.GeoContent = geoContent
	this.Ttl = ttl
	this.Editable = editable
	this.Links = links
	return &this
}

// NewLdResourceRecordSetDetailsWithDefaults instantiates a new LdResourceRecordSetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdResourceRecordSetDetailsWithDefaults() *LdResourceRecordSetDetails {
	this := LdResourceRecordSetDetails{}
	return &this
}

// GetName returns the Name field value
func (o *LdResourceRecordSetDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LdResourceRecordSetDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LdResourceRecordSetDetails) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *LdResourceRecordSetDetails) GetType() LdResourceRecordSetType {
	if o == nil {
		var ret LdResourceRecordSetType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LdResourceRecordSetDetails) GetTypeOk() (*LdResourceRecordSetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LdResourceRecordSetDetails) SetType(v LdResourceRecordSetType) {
	o.Type = v
}

// GetGeoContent returns the GeoContent field value
func (o *LdResourceRecordSetDetails) GetGeoContent() GeoContent {
	if o == nil {
		var ret GeoContent
		return ret
	}

	return o.GeoContent
}

// GetGeoContentOk returns a tuple with the GeoContent field value
// and a boolean to check if the value has been set.
func (o *LdResourceRecordSetDetails) GetGeoContentOk() (*GeoContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeoContent, true
}

// SetGeoContent sets field value
func (o *LdResourceRecordSetDetails) SetGeoContent(v GeoContent) {
	o.GeoContent = v
}

// GetTtl returns the Ttl field value
func (o *LdResourceRecordSetDetails) GetTtl() Ttl {
	if o == nil {
		var ret Ttl
		return ret
	}

	return o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value
// and a boolean to check if the value has been set.
func (o *LdResourceRecordSetDetails) GetTtlOk() (*Ttl, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ttl, true
}

// SetTtl sets field value
func (o *LdResourceRecordSetDetails) SetTtl(v Ttl) {
	o.Ttl = v
}

// GetEditable returns the Editable field value
func (o *LdResourceRecordSetDetails) GetEditable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Editable
}

// GetEditableOk returns a tuple with the Editable field value
// and a boolean to check if the value has been set.
func (o *LdResourceRecordSetDetails) GetEditableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Editable, true
}

// SetEditable sets field value
func (o *LdResourceRecordSetDetails) SetEditable(v bool) {
	o.Editable = v
}

// GetLinks returns the Links field value
func (o *LdResourceRecordSetDetails) GetLinks() LdLinks {
	if o == nil {
		var ret LdLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *LdResourceRecordSetDetails) GetLinksOk() (*LdLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *LdResourceRecordSetDetails) SetLinks(v LdLinks) {
	o.Links = v
}

func (o LdResourceRecordSetDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdResourceRecordSetDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["geoContent"] = o.GeoContent
	toSerialize["ttl"] = o.Ttl
	toSerialize["editable"] = o.Editable
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LdResourceRecordSetDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"geoContent",
		"ttl",
		"editable",
		"_links",
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

	varLdResourceRecordSetDetails := _LdResourceRecordSetDetails{}

	err = json.Unmarshal(data, &varLdResourceRecordSetDetails)

	if err != nil {
		return err
	}

	*o = LdResourceRecordSetDetails(varLdResourceRecordSetDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "geoContent")
		delete(additionalProperties, "ttl")
		delete(additionalProperties, "editable")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLdResourceRecordSetDetails struct {
	value *LdResourceRecordSetDetails
	isSet bool
}

func (v NullableLdResourceRecordSetDetails) Get() *LdResourceRecordSetDetails {
	return v.value
}

func (v *NullableLdResourceRecordSetDetails) Set(val *LdResourceRecordSetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableLdResourceRecordSetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableLdResourceRecordSetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdResourceRecordSetDetails(val *LdResourceRecordSetDetails) *NullableLdResourceRecordSetDetails {
	return &NullableLdResourceRecordSetDetails{value: val, isSet: true}
}

func (v NullableLdResourceRecordSetDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdResourceRecordSetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


