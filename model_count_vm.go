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

// CountVM struct for CountVM
type CountVM struct {
	Count *int64 `json:"count,omitempty"`
}

// NewCountVM instantiates a new CountVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountVM() *CountVM {
	this := CountVM{}
	return &this
}

// NewCountVMWithDefaults instantiates a new CountVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountVMWithDefaults() *CountVM {
	this := CountVM{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CountVM) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountVM) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CountVM) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *CountVM) SetCount(v int64) {
	o.Count = &v
}

func (o CountVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableCountVM struct {
	value *CountVM
	isSet bool
}

func (v NullableCountVM) Get() *CountVM {
	return v.value
}

func (v *NullableCountVM) Set(val *CountVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCountVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCountVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountVM(val *CountVM) *NullableCountVM {
	return &NullableCountVM{value: val, isSet: true}
}

func (v NullableCountVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


