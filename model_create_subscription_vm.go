/*
 * FastReport Cloud
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
)

// CreateSubscriptionVM struct for CreateSubscriptionVM
type CreateSubscriptionVM struct {
	PlanId *string `json:"planId,omitempty"`
	Name *string `json:"name,omitempty"`
	UserId *string `json:"userId,omitempty"`
}

// NewCreateSubscriptionVM instantiates a new CreateSubscriptionVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubscriptionVM() *CreateSubscriptionVM {
	this := CreateSubscriptionVM{}
	return &this
}

// NewCreateSubscriptionVMWithDefaults instantiates a new CreateSubscriptionVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubscriptionVMWithDefaults() *CreateSubscriptionVM {
	this := CreateSubscriptionVM{}
	return &this
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *CreateSubscriptionVM) GetPlanId() string {
	if o == nil || o.PlanId == nil {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionVM) GetPlanIdOk() (*string, bool) {
	if o == nil || o.PlanId == nil {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *CreateSubscriptionVM) HasPlanId() bool {
	if o != nil && o.PlanId != nil {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *CreateSubscriptionVM) SetPlanId(v string) {
	o.PlanId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateSubscriptionVM) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionVM) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateSubscriptionVM) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateSubscriptionVM) SetName(v string) {
	o.Name = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *CreateSubscriptionVM) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionVM) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *CreateSubscriptionVM) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *CreateSubscriptionVM) SetUserId(v string) {
	o.UserId = &v
}

func (o CreateSubscriptionVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PlanId != nil {
		toSerialize["planId"] = o.PlanId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSubscriptionVM struct {
	value *CreateSubscriptionVM
	isSet bool
}

func (v NullableCreateSubscriptionVM) Get() *CreateSubscriptionVM {
	return v.value
}

func (v *NullableCreateSubscriptionVM) Set(val *CreateSubscriptionVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubscriptionVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubscriptionVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubscriptionVM(val *CreateSubscriptionVM) *NullableCreateSubscriptionVM {
	return &NullableCreateSubscriptionVM{value: val, isSet: true}
}

func (v NullableCreateSubscriptionVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSubscriptionVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


