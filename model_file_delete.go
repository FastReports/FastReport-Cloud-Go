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
	"fmt"
)

// FileDelete the model 'FileDelete'
type FileDelete int32

// List of FileDelete
const (
	_0 FileDelete = 0
	_1 FileDelete = 1
	_MINUS_1 FileDelete = -1
)

var allowedFileDeleteEnumValues = []FileDelete{
	0,
	1,
	-1,
}

func (v *FileDelete) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileDelete(value)
	for _, existing := range allowedFileDeleteEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileDelete", value)
}

// NewFileDeleteFromValue returns a pointer to a valid FileDelete
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileDeleteFromValue(v int32) (*FileDelete, error) {
	ev := FileDelete(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileDelete: valid values are %v", v, allowedFileDeleteEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileDelete) IsValid() bool {
	for _, existing := range allowedFileDeleteEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileDelete value
func (v FileDelete) Ptr() *FileDelete {
	return &v
}

type NullableFileDelete struct {
	value *FileDelete
	isSet bool
}

func (v NullableFileDelete) Get() *FileDelete {
	return v.value
}

func (v *NullableFileDelete) Set(val *FileDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDelete(val *FileDelete) *NullableFileDelete {
	return &NullableFileDelete{value: val, isSet: true}
}

func (v NullableFileDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
