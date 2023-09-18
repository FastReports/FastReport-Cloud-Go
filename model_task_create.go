/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
	"fmt"
)

// TaskCreate the model 'TaskCreate'
type TaskCreate int32

// List of TaskCreate
const (
	_0 TaskCreate = 0
	_MINUS_1 TaskCreate = -1
)

// All allowed values of TaskCreate enum
var AllowedTaskCreateEnumValues = []TaskCreate{
	0,
	-1,
}

func (v *TaskCreate) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaskCreate(value)
	for _, existing := range AllowedTaskCreateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaskCreate", value)
}

// NewTaskCreateFromValue returns a pointer to a valid TaskCreate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaskCreateFromValue(v int32) (*TaskCreate, error) {
	ev := TaskCreate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TaskCreate: valid values are %v", v, AllowedTaskCreateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaskCreate) IsValid() bool {
	for _, existing := range AllowedTaskCreateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaskCreate value
func (v TaskCreate) Ptr() *TaskCreate {
	return &v
}

type NullableTaskCreate struct {
	value *TaskCreate
	isSet bool
}

func (v NullableTaskCreate) Get() *TaskCreate {
	return v.value
}

func (v *NullableTaskCreate) Set(val *TaskCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskCreate(val *TaskCreate) *NullableTaskCreate {
	return &NullableTaskCreate{value: val, isSet: true}
}

func (v NullableTaskCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

