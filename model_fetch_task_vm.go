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

// FetchTaskVM struct for FetchTaskVM
type FetchTaskVM struct {
	Name NullableString `json:"name,omitempty"`
	SubscriptionId NullableString `json:"subscriptionId,omitempty"`
	Type *TaskType `json:"type,omitempty"`
}

// NewFetchTaskVM instantiates a new FetchTaskVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchTaskVM() *FetchTaskVM {
	this := FetchTaskVM{}
	return &this
}

// NewFetchTaskVMWithDefaults instantiates a new FetchTaskVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchTaskVMWithDefaults() *FetchTaskVM {
	this := FetchTaskVM{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FetchTaskVM) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchTaskVM) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *FetchTaskVM) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *FetchTaskVM) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *FetchTaskVM) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *FetchTaskVM) UnsetName() {
	o.Name.Unset()
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FetchTaskVM) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FetchTaskVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *FetchTaskVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *FetchTaskVM) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *FetchTaskVM) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *FetchTaskVM) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FetchTaskVM) GetType() TaskType {
	if o == nil || o.Type == nil {
		var ret TaskType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchTaskVM) GetTypeOk() (*TaskType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FetchTaskVM) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TaskType and assigns it to the Type field.
func (o *FetchTaskVM) SetType(v TaskType) {
	o.Type = &v
}

func (o FetchTaskVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscriptionId"] = o.SubscriptionId.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableFetchTaskVM struct {
	value *FetchTaskVM
	isSet bool
}

func (v NullableFetchTaskVM) Get() *FetchTaskVM {
	return v.value
}

func (v *NullableFetchTaskVM) Set(val *FetchTaskVM) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchTaskVM) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchTaskVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchTaskVM(val *FetchTaskVM) *NullableFetchTaskVM {
	return &NullableFetchTaskVM{value: val, isSet: true}
}

func (v NullableFetchTaskVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchTaskVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


