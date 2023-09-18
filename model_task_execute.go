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

// TaskExecute the model 'TaskExecute'
type TaskExecute int32

// List of TaskExecute
const (
	_0 TaskExecute = 0
	_1 TaskExecute = 1
	_MINUS_1 TaskExecute = -1
)

// All allowed values of TaskExecute enum
var AllowedTaskExecuteEnumValues = []TaskExecute{
	0,
	1,
	-1,
}

func (v *TaskExecute) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaskExecute(value)
	for _, existing := range AllowedTaskExecuteEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaskExecute", value)
}

// NewTaskExecuteFromValue returns a pointer to a valid TaskExecute
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaskExecuteFromValue(v int32) (*TaskExecute, error) {
	ev := TaskExecute(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TaskExecute: valid values are %v", v, AllowedTaskExecuteEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaskExecute) IsValid() bool {
	for _, existing := range AllowedTaskExecuteEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaskExecute value
func (v TaskExecute) Ptr() *TaskExecute {
	return &v
}

type NullableTaskExecute struct {
	value *TaskExecute
	isSet bool
}

func (v NullableTaskExecute) Get() *TaskExecute {
	return v.value
}

func (v *NullableTaskExecute) Set(val *TaskExecute) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskExecute) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskExecute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskExecute(val *TaskExecute) *NullableTaskExecute {
	return &NullableTaskExecute{value: val, isSet: true}
}

func (v NullableTaskExecute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskExecute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

