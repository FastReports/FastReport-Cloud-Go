/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
	"time"
)

// checks if the AuditSubscriptionActionVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditSubscriptionActionVM{}

// AuditSubscriptionActionVM struct for AuditSubscriptionActionVM
type AuditSubscriptionActionVM struct {
	PeriodStart *time.Time `json:"periodStart,omitempty"`
	PeriodEnd *time.Time `json:"periodEnd,omitempty"`
	PlanId NullableString `json:"planId,omitempty"`
}

// NewAuditSubscriptionActionVM instantiates a new AuditSubscriptionActionVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditSubscriptionActionVM(t string) *AuditSubscriptionActionVM {
	this := AuditSubscriptionActionVM{}
	this.T = t
	return &this
}

// NewAuditSubscriptionActionVMWithDefaults instantiates a new AuditSubscriptionActionVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditSubscriptionActionVMWithDefaults() *AuditSubscriptionActionVM {
	this := AuditSubscriptionActionVM{}
	return &this
}

// GetPeriodStart returns the PeriodStart field value if set, zero value otherwise.
func (o *AuditSubscriptionActionVM) GetPeriodStart() time.Time {
	if o == nil || IsNil(o.PeriodStart) {
		var ret time.Time
		return ret
	}
	return *o.PeriodStart
}

// GetPeriodStartOk returns a tuple with the PeriodStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSubscriptionActionVM) GetPeriodStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PeriodStart) {
		return nil, false
	}
	return o.PeriodStart, true
}

// HasPeriodStart returns a boolean if a field has been set.
func (o *AuditSubscriptionActionVM) HasPeriodStart() bool {
	if o != nil && !IsNil(o.PeriodStart) {
		return true
	}

	return false
}

// SetPeriodStart gets a reference to the given time.Time and assigns it to the PeriodStart field.
func (o *AuditSubscriptionActionVM) SetPeriodStart(v time.Time) {
	o.PeriodStart = &v
}

// GetPeriodEnd returns the PeriodEnd field value if set, zero value otherwise.
func (o *AuditSubscriptionActionVM) GetPeriodEnd() time.Time {
	if o == nil || IsNil(o.PeriodEnd) {
		var ret time.Time
		return ret
	}
	return *o.PeriodEnd
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditSubscriptionActionVM) GetPeriodEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PeriodEnd) {
		return nil, false
	}
	return o.PeriodEnd, true
}

// HasPeriodEnd returns a boolean if a field has been set.
func (o *AuditSubscriptionActionVM) HasPeriodEnd() bool {
	if o != nil && !IsNil(o.PeriodEnd) {
		return true
	}

	return false
}

// SetPeriodEnd gets a reference to the given time.Time and assigns it to the PeriodEnd field.
func (o *AuditSubscriptionActionVM) SetPeriodEnd(v time.Time) {
	o.PeriodEnd = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditSubscriptionActionVM) GetPlanId() string {
	if o == nil || IsNil(o.PlanId.Get()) {
		var ret string
		return ret
	}
	return *o.PlanId.Get()
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditSubscriptionActionVM) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanId.Get(), o.PlanId.IsSet()
}

// HasPlanId returns a boolean if a field has been set.
func (o *AuditSubscriptionActionVM) HasPlanId() bool {
	if o != nil && o.PlanId.IsSet() {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given NullableString and assigns it to the PlanId field.
func (o *AuditSubscriptionActionVM) SetPlanId(v string) {
	o.PlanId.Set(&v)
}
// SetPlanIdNil sets the value for PlanId to be an explicit nil
func (o *AuditSubscriptionActionVM) SetPlanIdNil() {
	o.PlanId.Set(nil)
}

// UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
func (o *AuditSubscriptionActionVM) UnsetPlanId() {
	o.PlanId.Unset()
}

func (o AuditSubscriptionActionVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditSubscriptionActionVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PeriodStart) {
		toSerialize["periodStart"] = o.PeriodStart
	}
	if !IsNil(o.PeriodEnd) {
		toSerialize["periodEnd"] = o.PeriodEnd
	}
	if o.PlanId.IsSet() {
		toSerialize["planId"] = o.PlanId.Get()
	}
	return toSerialize, nil
}

type NullableAuditSubscriptionActionVM struct {
	value *AuditSubscriptionActionVM
	isSet bool
}

func (v NullableAuditSubscriptionActionVM) Get() *AuditSubscriptionActionVM {
	return v.value
}

func (v *NullableAuditSubscriptionActionVM) Set(val *AuditSubscriptionActionVM) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditSubscriptionActionVM) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditSubscriptionActionVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditSubscriptionActionVM(val *AuditSubscriptionActionVM) *NullableAuditSubscriptionActionVM {
	return &NullableAuditSubscriptionActionVM{value: val, isSet: true}
}

func (v NullableAuditSubscriptionActionVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditSubscriptionActionVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


