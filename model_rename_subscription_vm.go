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

// RenameSubscriptionVM struct for RenameSubscriptionVM
type RenameSubscriptionVM struct {
	Name *string `json:"name,omitempty"`
}

// NewRenameSubscriptionVM instantiates a new RenameSubscriptionVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenameSubscriptionVM() *RenameSubscriptionVM {
	this := RenameSubscriptionVM{}
	return &this
}

// NewRenameSubscriptionVMWithDefaults instantiates a new RenameSubscriptionVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenameSubscriptionVMWithDefaults() *RenameSubscriptionVM {
	this := RenameSubscriptionVM{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RenameSubscriptionVM) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenameSubscriptionVM) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RenameSubscriptionVM) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RenameSubscriptionVM) SetName(v string) {
	o.Name = &v
}

func (o RenameSubscriptionVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableRenameSubscriptionVM struct {
	value *RenameSubscriptionVM
	isSet bool
}

func (v NullableRenameSubscriptionVM) Get() *RenameSubscriptionVM {
	return v.value
}

func (v *NullableRenameSubscriptionVM) Set(val *RenameSubscriptionVM) {
	v.value = val
	v.isSet = true
}

func (v NullableRenameSubscriptionVM) IsSet() bool {
	return v.isSet
}

func (v *NullableRenameSubscriptionVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenameSubscriptionVM(val *RenameSubscriptionVM) *NullableRenameSubscriptionVM {
	return &NullableRenameSubscriptionVM{value: val, isSet: true}
}

func (v NullableRenameSubscriptionVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenameSubscriptionVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


