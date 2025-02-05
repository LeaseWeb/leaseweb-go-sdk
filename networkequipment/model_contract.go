/*
Dedicated Network Equipments

This is the description of the Dedicated Network Equipment API.  The base url of this API is `https://api.leaseweb.com`.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkequipment

import (
	"encoding/json"
	"time"
)

// checks if the Contract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Contract{}

// Contract Contract information
type Contract struct {
	Id *string `json:"id,omitempty"`
	CustomerId *string `json:"customerId,omitempty"`
	SalesOrgId *string `json:"salesOrgId,omitempty"`
	DeliveryStatus *string `json:"deliveryStatus,omitempty"`
	Reference NullableString `json:"reference,omitempty"`
	PrivateNetworkPortSpeed NullableFloat32 `json:"privateNetworkPortSpeed,omitempty"`
	Subnets []Subnet `json:"subnets,omitempty"`
	Status *string `json:"status,omitempty"`
	StartsAt *time.Time `json:"startsAt,omitempty"`
	EndsAt NullableTime `json:"endsAt,omitempty"`
	// Service Level Agreement
	Sla *string `json:"sla,omitempty"`
	ContractTerm *int32 `json:"contractTerm,omitempty"`
	ContractType *string `json:"contractType,omitempty"`
	// Applied billing cycle
	BillingCycle *int32 `json:"billingCycle,omitempty"`
	// The interval for which billing will be done
	BillingFrequency *string `json:"billingFrequency,omitempty"`
	// Price per billing frequency
	PricePerFrequency *string `json:"pricePerFrequency,omitempty"`
	Currency *string `json:"currency,omitempty"`
	NetworkTraffic *NetworkTraffic `json:"networkTraffic,omitempty"`
	SoftwareLicenses []SoftwareLicense `json:"softwareLicenses,omitempty"`
	ManagedServices []string `json:"managedServices,omitempty"`
	AggregationPackId NullableString `json:"aggregationPackId,omitempty"`
	Ipv4Quantity NullableInt32 `json:"ipv4Quantity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Contract Contract

// NewContract instantiates a new Contract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContract() *Contract {
	this := Contract{}
	return &this
}

// NewContractWithDefaults instantiates a new Contract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractWithDefaults() *Contract {
	this := Contract{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Contract) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Contract) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Contract) SetId(v string) {
	o.Id = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *Contract) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *Contract) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *Contract) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetSalesOrgId returns the SalesOrgId field value if set, zero value otherwise.
func (o *Contract) GetSalesOrgId() string {
	if o == nil || IsNil(o.SalesOrgId) {
		var ret string
		return ret
	}
	return *o.SalesOrgId
}

// GetSalesOrgIdOk returns a tuple with the SalesOrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetSalesOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesOrgId) {
		return nil, false
	}
	return o.SalesOrgId, true
}

// HasSalesOrgId returns a boolean if a field has been set.
func (o *Contract) HasSalesOrgId() bool {
	if o != nil && !IsNil(o.SalesOrgId) {
		return true
	}

	return false
}

// SetSalesOrgId gets a reference to the given string and assigns it to the SalesOrgId field.
func (o *Contract) SetSalesOrgId(v string) {
	o.SalesOrgId = &v
}

// GetDeliveryStatus returns the DeliveryStatus field value if set, zero value otherwise.
func (o *Contract) GetDeliveryStatus() string {
	if o == nil || IsNil(o.DeliveryStatus) {
		var ret string
		return ret
	}
	return *o.DeliveryStatus
}

// GetDeliveryStatusOk returns a tuple with the DeliveryStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetDeliveryStatusOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryStatus) {
		return nil, false
	}
	return o.DeliveryStatus, true
}

// HasDeliveryStatus returns a boolean if a field has been set.
func (o *Contract) HasDeliveryStatus() bool {
	if o != nil && !IsNil(o.DeliveryStatus) {
		return true
	}

	return false
}

// SetDeliveryStatus gets a reference to the given string and assigns it to the DeliveryStatus field.
func (o *Contract) SetDeliveryStatus(v string) {
	o.DeliveryStatus = &v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Contract) GetReference() string {
	if o == nil || IsNil(o.Reference.Get()) {
		var ret string
		return ret
	}
	return *o.Reference.Get()
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Contract) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reference.Get(), o.Reference.IsSet()
}

// HasReference returns a boolean if a field has been set.
func (o *Contract) HasReference() bool {
	if o != nil && o.Reference.IsSet() {
		return true
	}

	return false
}

// SetReference gets a reference to the given NullableString and assigns it to the Reference field.
func (o *Contract) SetReference(v string) {
	o.Reference.Set(&v)
}
// SetReferenceNil sets the value for Reference to be an explicit nil
func (o *Contract) SetReferenceNil() {
	o.Reference.Set(nil)
}

// UnsetReference ensures that no value is present for Reference, not even an explicit nil
func (o *Contract) UnsetReference() {
	o.Reference.Unset()
}

// GetPrivateNetworkPortSpeed returns the PrivateNetworkPortSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Contract) GetPrivateNetworkPortSpeed() float32 {
	if o == nil || IsNil(o.PrivateNetworkPortSpeed.Get()) {
		var ret float32
		return ret
	}
	return *o.PrivateNetworkPortSpeed.Get()
}

// GetPrivateNetworkPortSpeedOk returns a tuple with the PrivateNetworkPortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Contract) GetPrivateNetworkPortSpeedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateNetworkPortSpeed.Get(), o.PrivateNetworkPortSpeed.IsSet()
}

// HasPrivateNetworkPortSpeed returns a boolean if a field has been set.
func (o *Contract) HasPrivateNetworkPortSpeed() bool {
	if o != nil && o.PrivateNetworkPortSpeed.IsSet() {
		return true
	}

	return false
}

// SetPrivateNetworkPortSpeed gets a reference to the given NullableFloat32 and assigns it to the PrivateNetworkPortSpeed field.
func (o *Contract) SetPrivateNetworkPortSpeed(v float32) {
	o.PrivateNetworkPortSpeed.Set(&v)
}
// SetPrivateNetworkPortSpeedNil sets the value for PrivateNetworkPortSpeed to be an explicit nil
func (o *Contract) SetPrivateNetworkPortSpeedNil() {
	o.PrivateNetworkPortSpeed.Set(nil)
}

// UnsetPrivateNetworkPortSpeed ensures that no value is present for PrivateNetworkPortSpeed, not even an explicit nil
func (o *Contract) UnsetPrivateNetworkPortSpeed() {
	o.PrivateNetworkPortSpeed.Unset()
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *Contract) GetSubnets() []Subnet {
	if o == nil || IsNil(o.Subnets) {
		var ret []Subnet
		return ret
	}
	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetSubnetsOk() ([]Subnet, bool) {
	if o == nil || IsNil(o.Subnets) {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *Contract) HasSubnets() bool {
	if o != nil && !IsNil(o.Subnets) {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []Subnet and assigns it to the Subnets field.
func (o *Contract) SetSubnets(v []Subnet) {
	o.Subnets = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Contract) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Contract) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Contract) SetStatus(v string) {
	o.Status = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *Contract) GetStartsAt() time.Time {
	if o == nil || IsNil(o.StartsAt) {
		var ret time.Time
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetStartsAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *Contract) HasStartsAt() bool {
	if o != nil && !IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given time.Time and assigns it to the StartsAt field.
func (o *Contract) SetStartsAt(v time.Time) {
	o.StartsAt = &v
}

// GetEndsAt returns the EndsAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Contract) GetEndsAt() time.Time {
	if o == nil || IsNil(o.EndsAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndsAt.Get()
}

// GetEndsAtOk returns a tuple with the EndsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Contract) GetEndsAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndsAt.Get(), o.EndsAt.IsSet()
}

// HasEndsAt returns a boolean if a field has been set.
func (o *Contract) HasEndsAt() bool {
	if o != nil && o.EndsAt.IsSet() {
		return true
	}

	return false
}

// SetEndsAt gets a reference to the given NullableTime and assigns it to the EndsAt field.
func (o *Contract) SetEndsAt(v time.Time) {
	o.EndsAt.Set(&v)
}
// SetEndsAtNil sets the value for EndsAt to be an explicit nil
func (o *Contract) SetEndsAtNil() {
	o.EndsAt.Set(nil)
}

// UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil
func (o *Contract) UnsetEndsAt() {
	o.EndsAt.Unset()
}

// GetSla returns the Sla field value if set, zero value otherwise.
func (o *Contract) GetSla() string {
	if o == nil || IsNil(o.Sla) {
		var ret string
		return ret
	}
	return *o.Sla
}

// GetSlaOk returns a tuple with the Sla field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetSlaOk() (*string, bool) {
	if o == nil || IsNil(o.Sla) {
		return nil, false
	}
	return o.Sla, true
}

// HasSla returns a boolean if a field has been set.
func (o *Contract) HasSla() bool {
	if o != nil && !IsNil(o.Sla) {
		return true
	}

	return false
}

// SetSla gets a reference to the given string and assigns it to the Sla field.
func (o *Contract) SetSla(v string) {
	o.Sla = &v
}

// GetContractTerm returns the ContractTerm field value if set, zero value otherwise.
func (o *Contract) GetContractTerm() int32 {
	if o == nil || IsNil(o.ContractTerm) {
		var ret int32
		return ret
	}
	return *o.ContractTerm
}

// GetContractTermOk returns a tuple with the ContractTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetContractTermOk() (*int32, bool) {
	if o == nil || IsNil(o.ContractTerm) {
		return nil, false
	}
	return o.ContractTerm, true
}

// HasContractTerm returns a boolean if a field has been set.
func (o *Contract) HasContractTerm() bool {
	if o != nil && !IsNil(o.ContractTerm) {
		return true
	}

	return false
}

// SetContractTerm gets a reference to the given int32 and assigns it to the ContractTerm field.
func (o *Contract) SetContractTerm(v int32) {
	o.ContractTerm = &v
}

// GetContractType returns the ContractType field value if set, zero value otherwise.
func (o *Contract) GetContractType() string {
	if o == nil || IsNil(o.ContractType) {
		var ret string
		return ret
	}
	return *o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetContractTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContractType) {
		return nil, false
	}
	return o.ContractType, true
}

// HasContractType returns a boolean if a field has been set.
func (o *Contract) HasContractType() bool {
	if o != nil && !IsNil(o.ContractType) {
		return true
	}

	return false
}

// SetContractType gets a reference to the given string and assigns it to the ContractType field.
func (o *Contract) SetContractType(v string) {
	o.ContractType = &v
}

// GetBillingCycle returns the BillingCycle field value if set, zero value otherwise.
func (o *Contract) GetBillingCycle() int32 {
	if o == nil || IsNil(o.BillingCycle) {
		var ret int32
		return ret
	}
	return *o.BillingCycle
}

// GetBillingCycleOk returns a tuple with the BillingCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetBillingCycleOk() (*int32, bool) {
	if o == nil || IsNil(o.BillingCycle) {
		return nil, false
	}
	return o.BillingCycle, true
}

// HasBillingCycle returns a boolean if a field has been set.
func (o *Contract) HasBillingCycle() bool {
	if o != nil && !IsNil(o.BillingCycle) {
		return true
	}

	return false
}

// SetBillingCycle gets a reference to the given int32 and assigns it to the BillingCycle field.
func (o *Contract) SetBillingCycle(v int32) {
	o.BillingCycle = &v
}

// GetBillingFrequency returns the BillingFrequency field value if set, zero value otherwise.
func (o *Contract) GetBillingFrequency() string {
	if o == nil || IsNil(o.BillingFrequency) {
		var ret string
		return ret
	}
	return *o.BillingFrequency
}

// GetBillingFrequencyOk returns a tuple with the BillingFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetBillingFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.BillingFrequency) {
		return nil, false
	}
	return o.BillingFrequency, true
}

// HasBillingFrequency returns a boolean if a field has been set.
func (o *Contract) HasBillingFrequency() bool {
	if o != nil && !IsNil(o.BillingFrequency) {
		return true
	}

	return false
}

// SetBillingFrequency gets a reference to the given string and assigns it to the BillingFrequency field.
func (o *Contract) SetBillingFrequency(v string) {
	o.BillingFrequency = &v
}

// GetPricePerFrequency returns the PricePerFrequency field value if set, zero value otherwise.
func (o *Contract) GetPricePerFrequency() string {
	if o == nil || IsNil(o.PricePerFrequency) {
		var ret string
		return ret
	}
	return *o.PricePerFrequency
}

// GetPricePerFrequencyOk returns a tuple with the PricePerFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetPricePerFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.PricePerFrequency) {
		return nil, false
	}
	return o.PricePerFrequency, true
}

// HasPricePerFrequency returns a boolean if a field has been set.
func (o *Contract) HasPricePerFrequency() bool {
	if o != nil && !IsNil(o.PricePerFrequency) {
		return true
	}

	return false
}

// SetPricePerFrequency gets a reference to the given string and assigns it to the PricePerFrequency field.
func (o *Contract) SetPricePerFrequency(v string) {
	o.PricePerFrequency = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Contract) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Contract) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Contract) SetCurrency(v string) {
	o.Currency = &v
}

// GetNetworkTraffic returns the NetworkTraffic field value if set, zero value otherwise.
func (o *Contract) GetNetworkTraffic() NetworkTraffic {
	if o == nil || IsNil(o.NetworkTraffic) {
		var ret NetworkTraffic
		return ret
	}
	return *o.NetworkTraffic
}

// GetNetworkTrafficOk returns a tuple with the NetworkTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetNetworkTrafficOk() (*NetworkTraffic, bool) {
	if o == nil || IsNil(o.NetworkTraffic) {
		return nil, false
	}
	return o.NetworkTraffic, true
}

// HasNetworkTraffic returns a boolean if a field has been set.
func (o *Contract) HasNetworkTraffic() bool {
	if o != nil && !IsNil(o.NetworkTraffic) {
		return true
	}

	return false
}

// SetNetworkTraffic gets a reference to the given NetworkTraffic and assigns it to the NetworkTraffic field.
func (o *Contract) SetNetworkTraffic(v NetworkTraffic) {
	o.NetworkTraffic = &v
}

// GetSoftwareLicenses returns the SoftwareLicenses field value if set, zero value otherwise.
func (o *Contract) GetSoftwareLicenses() []SoftwareLicense {
	if o == nil || IsNil(o.SoftwareLicenses) {
		var ret []SoftwareLicense
		return ret
	}
	return o.SoftwareLicenses
}

// GetSoftwareLicensesOk returns a tuple with the SoftwareLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetSoftwareLicensesOk() ([]SoftwareLicense, bool) {
	if o == nil || IsNil(o.SoftwareLicenses) {
		return nil, false
	}
	return o.SoftwareLicenses, true
}

// HasSoftwareLicenses returns a boolean if a field has been set.
func (o *Contract) HasSoftwareLicenses() bool {
	if o != nil && !IsNil(o.SoftwareLicenses) {
		return true
	}

	return false
}

// SetSoftwareLicenses gets a reference to the given []SoftwareLicense and assigns it to the SoftwareLicenses field.
func (o *Contract) SetSoftwareLicenses(v []SoftwareLicense) {
	o.SoftwareLicenses = v
}

// GetManagedServices returns the ManagedServices field value if set, zero value otherwise.
func (o *Contract) GetManagedServices() []string {
	if o == nil || IsNil(o.ManagedServices) {
		var ret []string
		return ret
	}
	return o.ManagedServices
}

// GetManagedServicesOk returns a tuple with the ManagedServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetManagedServicesOk() ([]string, bool) {
	if o == nil || IsNil(o.ManagedServices) {
		return nil, false
	}
	return o.ManagedServices, true
}

// HasManagedServices returns a boolean if a field has been set.
func (o *Contract) HasManagedServices() bool {
	if o != nil && !IsNil(o.ManagedServices) {
		return true
	}

	return false
}

// SetManagedServices gets a reference to the given []string and assigns it to the ManagedServices field.
func (o *Contract) SetManagedServices(v []string) {
	o.ManagedServices = v
}

// GetAggregationPackId returns the AggregationPackId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Contract) GetAggregationPackId() string {
	if o == nil || IsNil(o.AggregationPackId.Get()) {
		var ret string
		return ret
	}
	return *o.AggregationPackId.Get()
}

// GetAggregationPackIdOk returns a tuple with the AggregationPackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Contract) GetAggregationPackIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AggregationPackId.Get(), o.AggregationPackId.IsSet()
}

// HasAggregationPackId returns a boolean if a field has been set.
func (o *Contract) HasAggregationPackId() bool {
	if o != nil && o.AggregationPackId.IsSet() {
		return true
	}

	return false
}

// SetAggregationPackId gets a reference to the given NullableString and assigns it to the AggregationPackId field.
func (o *Contract) SetAggregationPackId(v string) {
	o.AggregationPackId.Set(&v)
}
// SetAggregationPackIdNil sets the value for AggregationPackId to be an explicit nil
func (o *Contract) SetAggregationPackIdNil() {
	o.AggregationPackId.Set(nil)
}

// UnsetAggregationPackId ensures that no value is present for AggregationPackId, not even an explicit nil
func (o *Contract) UnsetAggregationPackId() {
	o.AggregationPackId.Unset()
}

// GetIpv4Quantity returns the Ipv4Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Contract) GetIpv4Quantity() int32 {
	if o == nil || IsNil(o.Ipv4Quantity.Get()) {
		var ret int32
		return ret
	}
	return *o.Ipv4Quantity.Get()
}

// GetIpv4QuantityOk returns a tuple with the Ipv4Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Contract) GetIpv4QuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ipv4Quantity.Get(), o.Ipv4Quantity.IsSet()
}

// HasIpv4Quantity returns a boolean if a field has been set.
func (o *Contract) HasIpv4Quantity() bool {
	if o != nil && o.Ipv4Quantity.IsSet() {
		return true
	}

	return false
}

// SetIpv4Quantity gets a reference to the given NullableInt32 and assigns it to the Ipv4Quantity field.
func (o *Contract) SetIpv4Quantity(v int32) {
	o.Ipv4Quantity.Set(&v)
}
// SetIpv4QuantityNil sets the value for Ipv4Quantity to be an explicit nil
func (o *Contract) SetIpv4QuantityNil() {
	o.Ipv4Quantity.Set(nil)
}

// UnsetIpv4Quantity ensures that no value is present for Ipv4Quantity, not even an explicit nil
func (o *Contract) UnsetIpv4Quantity() {
	o.Ipv4Quantity.Unset()
}

func (o Contract) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Contract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	if !IsNil(o.SalesOrgId) {
		toSerialize["salesOrgId"] = o.SalesOrgId
	}
	if !IsNil(o.DeliveryStatus) {
		toSerialize["deliveryStatus"] = o.DeliveryStatus
	}
	if o.Reference.IsSet() {
		toSerialize["reference"] = o.Reference.Get()
	}
	if o.PrivateNetworkPortSpeed.IsSet() {
		toSerialize["privateNetworkPortSpeed"] = o.PrivateNetworkPortSpeed.Get()
	}
	if !IsNil(o.Subnets) {
		toSerialize["subnets"] = o.Subnets
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StartsAt) {
		toSerialize["startsAt"] = o.StartsAt
	}
	if o.EndsAt.IsSet() {
		toSerialize["endsAt"] = o.EndsAt.Get()
	}
	if !IsNil(o.Sla) {
		toSerialize["sla"] = o.Sla
	}
	if !IsNil(o.ContractTerm) {
		toSerialize["contractTerm"] = o.ContractTerm
	}
	if !IsNil(o.ContractType) {
		toSerialize["contractType"] = o.ContractType
	}
	if !IsNil(o.BillingCycle) {
		toSerialize["billingCycle"] = o.BillingCycle
	}
	if !IsNil(o.BillingFrequency) {
		toSerialize["billingFrequency"] = o.BillingFrequency
	}
	if !IsNil(o.PricePerFrequency) {
		toSerialize["pricePerFrequency"] = o.PricePerFrequency
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.NetworkTraffic) {
		toSerialize["networkTraffic"] = o.NetworkTraffic
	}
	if !IsNil(o.SoftwareLicenses) {
		toSerialize["softwareLicenses"] = o.SoftwareLicenses
	}
	if !IsNil(o.ManagedServices) {
		toSerialize["managedServices"] = o.ManagedServices
	}
	if o.AggregationPackId.IsSet() {
		toSerialize["aggregationPackId"] = o.AggregationPackId.Get()
	}
	if o.Ipv4Quantity.IsSet() {
		toSerialize["ipv4Quantity"] = o.Ipv4Quantity.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Contract) UnmarshalJSON(data []byte) (err error) {
	varContract := _Contract{}

	err = json.Unmarshal(data, &varContract)

	if err != nil {
		return err
	}

	*o = Contract(varContract)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "salesOrgId")
		delete(additionalProperties, "deliveryStatus")
		delete(additionalProperties, "reference")
		delete(additionalProperties, "privateNetworkPortSpeed")
		delete(additionalProperties, "subnets")
		delete(additionalProperties, "status")
		delete(additionalProperties, "startsAt")
		delete(additionalProperties, "endsAt")
		delete(additionalProperties, "sla")
		delete(additionalProperties, "contractTerm")
		delete(additionalProperties, "contractType")
		delete(additionalProperties, "billingCycle")
		delete(additionalProperties, "billingFrequency")
		delete(additionalProperties, "pricePerFrequency")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "networkTraffic")
		delete(additionalProperties, "softwareLicenses")
		delete(additionalProperties, "managedServices")
		delete(additionalProperties, "aggregationPackId")
		delete(additionalProperties, "ipv4Quantity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContract struct {
	value *Contract
	isSet bool
}

func (v NullableContract) Get() *Contract {
	return v.value
}

func (v *NullableContract) Set(val *Contract) {
	v.value = val
	v.isSet = true
}

func (v NullableContract) IsSet() bool {
	return v.isSet
}

func (v *NullableContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContract(val *Contract) *NullableContract {
	return &NullableContract{value: val, isSet: true}
}

func (v NullableContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


