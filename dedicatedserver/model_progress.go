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

// checks if the Progress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Progress{}

// Progress Describes progress of the jobs on the server
type Progress struct {
	Canceled *int32 `json:"canceled,omitempty"`
	Expired *int32 `json:"expired,omitempty"`
	Failed *int32 `json:"failed,omitempty"`
	Finished *int32 `json:"finished,omitempty"`
	Inprogress *int32 `json:"inprogress,omitempty"`
	Pending *int32 `json:"pending,omitempty"`
	Percentage *int32 `json:"percentage,omitempty"`
	Total *int32 `json:"total,omitempty"`
	Waiting *int32 `json:"waiting,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Progress Progress

// NewProgress instantiates a new Progress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgress() *Progress {
	this := Progress{}
	return &this
}

// NewProgressWithDefaults instantiates a new Progress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgressWithDefaults() *Progress {
	this := Progress{}
	return &this
}

// GetCanceled returns the Canceled field value if set, zero value otherwise.
func (o *Progress) GetCanceled() int32 {
	if o == nil || IsNil(o.Canceled) {
		var ret int32
		return ret
	}
	return *o.Canceled
}

// GetCanceledOk returns a tuple with the Canceled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Progress) GetCanceledOk() (*int32, bool) {
	if o == nil || IsNil(o.Canceled) {
		return nil, false
	}
	return o.Canceled, true
}

// HasCanceled returns a boolean if a field has been set.
func (o *Progress) HasCanceled() bool {
	if o != nil && !IsNil(o.Canceled) {
		return true
	}

	return false
}

// SetCanceled gets a reference to the given int32 and assigns it to the Canceled field.
func (o *Progress) SetCanceled(v int32) {
	o.Canceled = &v
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *Progress) GetExpired() int32 {
	if o == nil || IsNil(o.Expired) {
		var ret int32
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Progress) GetExpiredOk() (*int32, bool) {
	if o == nil || IsNil(o.Expired) {
		return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *Progress) HasExpired() bool {
	if o != nil && !IsNil(o.Expired) {
		return true
	}

	return false
}

// SetExpired gets a reference to the given int32 and assigns it to the Expired field.
func (o *Progress) SetExpired(v int32) {
	o.Expired = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *Progress) GetFailed() int32 {
	if o == nil || IsNil(o.Failed) {
		var ret int32
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Progress) GetFailedOk() (*int32, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *Progress) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given int32 and assigns it to the Failed field.
func (o *Progress) SetFailed(v int32) {
	o.Failed = &v
}

// GetFinished returns the Finished field value if set, zero value otherwise.
func (o *Progress) GetFinished() int32 {
	if o == nil || IsNil(o.Finished) {
		var ret int32
		return ret
	}
	return *o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Progress) GetFinishedOk() (*int32, bool) {
	if o == nil || IsNil(o.Finished) {
		return nil, false
	}
	return o.Finished, true
}

// HasFinished returns a boolean if a field has been set.
func (o *Progress) HasFinished() bool {
	if o != nil && !IsNil(o.Finished) {
		return true
	}

	return false
}

// SetFinished gets a reference to the given int32 and assigns it to the Finished field.
func (o *Progress) SetFinished(v int32) {
	o.Finished = &v
}

// GetInprogress returns the Inprogress field value if set, zero value otherwise.
func (o *Progress) GetInprogress() int32 {
	if o == nil || IsNil(o.Inprogress) {
		var ret int32
		return ret
	}
	return *o.Inprogress
}

// GetInprogressOk returns a tuple with the Inprogress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Progress) GetInprogressOk() (*int32, bool) {
	if o == nil || IsNil(o.Inprogress) {
		return nil, false
	}
	return o.Inprogress, true
}

// HasInprogress returns a boolean if a field has been set.
func (o *Progress) HasInprogress() bool {
	if o != nil && !IsNil(o.Inprogress) {
		return true
	}

	return false
}

// SetInprogress gets a reference to the given int32 and assigns it to the Inprogress field.
func (o *Progress) SetInprogress(v int32) {
	o.Inprogress = &v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *Progress) GetPending() int32 {
	if o == nil || IsNil(o.Pending) {
		var ret int32
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Progress) GetPendingOk() (*int32, bool) {
	if o == nil || IsNil(o.Pending) {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *Progress) HasPending() bool {
	if o != nil && !IsNil(o.Pending) {
		return true
	}

	return false
}

// SetPending gets a reference to the given int32 and assigns it to the Pending field.
func (o *Progress) SetPending(v int32) {
	o.Pending = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *Progress) GetPercentage() int32 {
	if o == nil || IsNil(o.Percentage) {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Progress) GetPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *Progress) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *Progress) SetPercentage(v int32) {
	o.Percentage = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *Progress) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Progress) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *Progress) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *Progress) SetTotal(v int32) {
	o.Total = &v
}

// GetWaiting returns the Waiting field value if set, zero value otherwise.
func (o *Progress) GetWaiting() int32 {
	if o == nil || IsNil(o.Waiting) {
		var ret int32
		return ret
	}
	return *o.Waiting
}

// GetWaitingOk returns a tuple with the Waiting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Progress) GetWaitingOk() (*int32, bool) {
	if o == nil || IsNil(o.Waiting) {
		return nil, false
	}
	return o.Waiting, true
}

// HasWaiting returns a boolean if a field has been set.
func (o *Progress) HasWaiting() bool {
	if o != nil && !IsNil(o.Waiting) {
		return true
	}

	return false
}

// SetWaiting gets a reference to the given int32 and assigns it to the Waiting field.
func (o *Progress) SetWaiting(v int32) {
	o.Waiting = &v
}

func (o Progress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Progress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Canceled) {
		toSerialize["canceled"] = o.Canceled
	}
	if !IsNil(o.Expired) {
		toSerialize["expired"] = o.Expired
	}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.Finished) {
		toSerialize["finished"] = o.Finished
	}
	if !IsNil(o.Inprogress) {
		toSerialize["inprogress"] = o.Inprogress
	}
	if !IsNil(o.Pending) {
		toSerialize["pending"] = o.Pending
	}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Waiting) {
		toSerialize["waiting"] = o.Waiting
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Progress) UnmarshalJSON(data []byte) (err error) {
	varProgress := _Progress{}

	err = json.Unmarshal(data, &varProgress)

	if err != nil {
		return err
	}

	*o = Progress(varProgress)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "canceled")
		delete(additionalProperties, "expired")
		delete(additionalProperties, "failed")
		delete(additionalProperties, "finished")
		delete(additionalProperties, "inprogress")
		delete(additionalProperties, "pending")
		delete(additionalProperties, "percentage")
		delete(additionalProperties, "total")
		delete(additionalProperties, "waiting")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProgress struct {
	value *Progress
	isSet bool
}

func (v NullableProgress) Get() *Progress {
	return v.value
}

func (v *NullableProgress) Set(val *Progress) {
	v.value = val
	v.isSet = true
}

func (v NullableProgress) IsSet() bool {
	return v.isSet
}

func (v *NullableProgress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgress(val *Progress) *NullableProgress {
	return &NullableProgress{value: val, isSet: true}
}

func (v NullableProgress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


