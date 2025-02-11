/*
IP management

> The base URL for this API is: **https://api.leaseweb.com/ipMgmt/v2/_**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipmgmt

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the NullRoutedIP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NullRoutedIP{}

// NullRoutedIP struct for NullRoutedIP
type NullRoutedIP struct {
	// Null route ID
	Id string `json:"id"`
	// IP address
	Ip string `json:"ip"`
	// Null route date in UTC in format yyyy-mm-ddThh:mm:ssZ
	NulledAt time.Time `json:"nulledAt"`
	// Email address of the user who created the null route or 'LeaseWeb' if null route was created by LeaseWeb.
	NulledBy string `json:"nulledBy"`
	// Null route permission level. If greater than 1 then the null route can only be removed by LeaseWeb.
	NullLevel int32 `json:"nullLevel"`
	// The date and time when the null route is to be automatically removed.
	AutomatedUnnullingAt NullableTime `json:"automatedUnnullingAt"`
	// The date and time when the null route has been removed. If null then the null route is still active.
	UnnulledAt NullableTime `json:"unnulledAt"`
	// Email address of the user who removed the null route or 'LeaseWeb' if null route was removed by LeaseWeb.
	UnnulledBy NullableString `json:"unnulledBy"`
	// Reference stored with the null route.
	TicketId NullableString `json:"ticketId"`
	// Comment stored with the null route.
	Comment NullableString `json:"comment"`
	// ID of the equipment which was assigned to the IP at the time of null route creation.
	EquipmentId string `json:"equipmentId"`
	AssignedContract NullableAssignedContract `json:"assignedContract"`
	AdditionalProperties map[string]interface{}
}

type _NullRoutedIP NullRoutedIP

// NewNullRoutedIP instantiates a new NullRoutedIP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNullRoutedIP(id string, ip string, nulledAt time.Time, nulledBy string, nullLevel int32, automatedUnnullingAt NullableTime, unnulledAt NullableTime, unnulledBy NullableString, ticketId NullableString, comment NullableString, equipmentId string, assignedContract NullableAssignedContract) *NullRoutedIP {
	this := NullRoutedIP{}
	this.Id = id
	this.Ip = ip
	this.NulledAt = nulledAt
	this.NulledBy = nulledBy
	this.NullLevel = nullLevel
	this.AutomatedUnnullingAt = automatedUnnullingAt
	this.UnnulledAt = unnulledAt
	this.UnnulledBy = unnulledBy
	this.TicketId = ticketId
	this.Comment = comment
	this.EquipmentId = equipmentId
	this.AssignedContract = assignedContract
	return &this
}

// NewNullRoutedIPWithDefaults instantiates a new NullRoutedIP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNullRoutedIPWithDefaults() *NullRoutedIP {
	this := NullRoutedIP{}
	var nullLevel int32 = 0
	this.NullLevel = nullLevel
	return &this
}

// GetId returns the Id field value
func (o *NullRoutedIP) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NullRoutedIP) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NullRoutedIP) SetId(v string) {
	o.Id = v
}

// GetIp returns the Ip field value
func (o *NullRoutedIP) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *NullRoutedIP) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *NullRoutedIP) SetIp(v string) {
	o.Ip = v
}

// GetNulledAt returns the NulledAt field value
func (o *NullRoutedIP) GetNulledAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.NulledAt
}

// GetNulledAtOk returns a tuple with the NulledAt field value
// and a boolean to check if the value has been set.
func (o *NullRoutedIP) GetNulledAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NulledAt, true
}

// SetNulledAt sets field value
func (o *NullRoutedIP) SetNulledAt(v time.Time) {
	o.NulledAt = v
}

// GetNulledBy returns the NulledBy field value
func (o *NullRoutedIP) GetNulledBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NulledBy
}

// GetNulledByOk returns a tuple with the NulledBy field value
// and a boolean to check if the value has been set.
func (o *NullRoutedIP) GetNulledByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NulledBy, true
}

// SetNulledBy sets field value
func (o *NullRoutedIP) SetNulledBy(v string) {
	o.NulledBy = v
}

// GetNullLevel returns the NullLevel field value
func (o *NullRoutedIP) GetNullLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NullLevel
}

// GetNullLevelOk returns a tuple with the NullLevel field value
// and a boolean to check if the value has been set.
func (o *NullRoutedIP) GetNullLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NullLevel, true
}

// SetNullLevel sets field value
func (o *NullRoutedIP) SetNullLevel(v int32) {
	o.NullLevel = v
}

// GetAutomatedUnnullingAt returns the AutomatedUnnullingAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *NullRoutedIP) GetAutomatedUnnullingAt() time.Time {
	if o == nil || o.AutomatedUnnullingAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.AutomatedUnnullingAt.Get()
}

// GetAutomatedUnnullingAtOk returns a tuple with the AutomatedUnnullingAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullRoutedIP) GetAutomatedUnnullingAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutomatedUnnullingAt.Get(), o.AutomatedUnnullingAt.IsSet()
}

// SetAutomatedUnnullingAt sets field value
func (o *NullRoutedIP) SetAutomatedUnnullingAt(v time.Time) {
	o.AutomatedUnnullingAt.Set(&v)
}

// GetUnnulledAt returns the UnnulledAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *NullRoutedIP) GetUnnulledAt() time.Time {
	if o == nil || o.UnnulledAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.UnnulledAt.Get()
}

// GetUnnulledAtOk returns a tuple with the UnnulledAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullRoutedIP) GetUnnulledAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnnulledAt.Get(), o.UnnulledAt.IsSet()
}

// SetUnnulledAt sets field value
func (o *NullRoutedIP) SetUnnulledAt(v time.Time) {
	o.UnnulledAt.Set(&v)
}

// GetUnnulledBy returns the UnnulledBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NullRoutedIP) GetUnnulledBy() string {
	if o == nil || o.UnnulledBy.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnnulledBy.Get()
}

// GetUnnulledByOk returns a tuple with the UnnulledBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullRoutedIP) GetUnnulledByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnnulledBy.Get(), o.UnnulledBy.IsSet()
}

// SetUnnulledBy sets field value
func (o *NullRoutedIP) SetUnnulledBy(v string) {
	o.UnnulledBy.Set(&v)
}

// GetTicketId returns the TicketId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NullRoutedIP) GetTicketId() string {
	if o == nil || o.TicketId.Get() == nil {
		var ret string
		return ret
	}

	return *o.TicketId.Get()
}

// GetTicketIdOk returns a tuple with the TicketId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullRoutedIP) GetTicketIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TicketId.Get(), o.TicketId.IsSet()
}

// SetTicketId sets field value
func (o *NullRoutedIP) SetTicketId(v string) {
	o.TicketId.Set(&v)
}

// GetComment returns the Comment field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NullRoutedIP) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}

	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullRoutedIP) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// SetComment sets field value
func (o *NullRoutedIP) SetComment(v string) {
	o.Comment.Set(&v)
}

// GetEquipmentId returns the EquipmentId field value
func (o *NullRoutedIP) GetEquipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EquipmentId
}

// GetEquipmentIdOk returns a tuple with the EquipmentId field value
// and a boolean to check if the value has been set.
func (o *NullRoutedIP) GetEquipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EquipmentId, true
}

// SetEquipmentId sets field value
func (o *NullRoutedIP) SetEquipmentId(v string) {
	o.EquipmentId = v
}

// GetAssignedContract returns the AssignedContract field value
// If the value is explicit nil, the zero value for AssignedContract will be returned
func (o *NullRoutedIP) GetAssignedContract() AssignedContract {
	if o == nil || o.AssignedContract.Get() == nil {
		var ret AssignedContract
		return ret
	}

	return *o.AssignedContract.Get()
}

// GetAssignedContractOk returns a tuple with the AssignedContract field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullRoutedIP) GetAssignedContractOk() (*AssignedContract, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedContract.Get(), o.AssignedContract.IsSet()
}

// SetAssignedContract sets field value
func (o *NullRoutedIP) SetAssignedContract(v AssignedContract) {
	o.AssignedContract.Set(&v)
}

func (o NullRoutedIP) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NullRoutedIP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["ip"] = o.Ip
	toSerialize["nulledAt"] = o.NulledAt
	toSerialize["nulledBy"] = o.NulledBy
	toSerialize["nullLevel"] = o.NullLevel
	toSerialize["automatedUnnullingAt"] = o.AutomatedUnnullingAt.Get()
	toSerialize["unnulledAt"] = o.UnnulledAt.Get()
	toSerialize["unnulledBy"] = o.UnnulledBy.Get()
	toSerialize["ticketId"] = o.TicketId.Get()
	toSerialize["comment"] = o.Comment.Get()
	toSerialize["equipmentId"] = o.EquipmentId
	toSerialize["assignedContract"] = o.AssignedContract.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NullRoutedIP) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"ip",
		"nulledAt",
		"nulledBy",
		"nullLevel",
		"automatedUnnullingAt",
		"unnulledAt",
		"unnulledBy",
		"ticketId",
		"comment",
		"equipmentId",
		"assignedContract",
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

	varNullRoutedIP := _NullRoutedIP{}

	err = json.Unmarshal(data, &varNullRoutedIP)

	if err != nil {
		return err
	}

	*o = NullRoutedIP(varNullRoutedIP)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "nulledAt")
		delete(additionalProperties, "nulledBy")
		delete(additionalProperties, "nullLevel")
		delete(additionalProperties, "automatedUnnullingAt")
		delete(additionalProperties, "unnulledAt")
		delete(additionalProperties, "unnulledBy")
		delete(additionalProperties, "ticketId")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "equipmentId")
		delete(additionalProperties, "assignedContract")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNullRoutedIP struct {
	value *NullRoutedIP
	isSet bool
}

func (v NullableNullRoutedIP) Get() *NullRoutedIP {
	return v.value
}

func (v *NullableNullRoutedIP) Set(val *NullRoutedIP) {
	v.value = val
	v.isSet = true
}

func (v NullableNullRoutedIP) IsSet() bool {
	return v.isSet
}

func (v *NullableNullRoutedIP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNullRoutedIP(val *NullRoutedIP) *NullableNullRoutedIP {
	return &NullableNullRoutedIP{value: val, isSet: true}
}

func (v NullableNullRoutedIP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNullRoutedIP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


