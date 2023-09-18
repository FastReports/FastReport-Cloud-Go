/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
)

// checks if the ReportsVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportsVM{}

// ReportsVM struct for ReportsVM
type ReportsVM struct {
}

// NewReportsVM instantiates a new ReportsVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportsVM() *ReportsVM {
	this := ReportsVM{}
	return &this
}

// NewReportsVMWithDefaults instantiates a new ReportsVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportsVMWithDefaults() *ReportsVM {
	this := ReportsVM{}
	return &this
}

func (o ReportsVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportsVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableReportsVM struct {
	value *ReportsVM
	isSet bool
}

func (v NullableReportsVM) Get() *ReportsVM {
	return v.value
}

func (v *NullableReportsVM) Set(val *ReportsVM) {
	v.value = val
	v.isSet = true
}

func (v NullableReportsVM) IsSet() bool {
	return v.isSet
}

func (v *NullableReportsVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportsVM(val *ReportsVM) *NullableReportsVM {
	return &NullableReportsVM{value: val, isSet: true}
}

func (v NullableReportsVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportsVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


