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

// FileStatus the model 'FileStatus'
type FileStatus string

// List of FileStatus
const (
	NONE FileStatus = "None"
	IN_QUEUE FileStatus = "InQueue"
	IN_PROCESS FileStatus = "InProcess"
	SUCCESS FileStatus = "Success"
	FAILED FileStatus = "Failed"
)

// All allowed values of FileStatus enum
var AllowedFileStatusEnumValues = []FileStatus{
	"None",
	"InQueue",
	"InProcess",
	"Success",
	"Failed",
}

func (v *FileStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileStatus(value)
	for _, existing := range AllowedFileStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileStatus", value)
}

// NewFileStatusFromValue returns a pointer to a valid FileStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileStatusFromValue(v string) (*FileStatus, error) {
	ev := FileStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileStatus: valid values are %v", v, AllowedFileStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileStatus) IsValid() bool {
	for _, existing := range AllowedFileStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileStatus value
func (v FileStatus) Ptr() *FileStatus {
	return &v
}

type NullableFileStatus struct {
	value *FileStatus
	isSet bool
}

func (v NullableFileStatus) Get() *FileStatus {
	return v.value
}

func (v *NullableFileStatus) Set(val *FileStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableFileStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableFileStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileStatus(val *FileStatus) *NullableFileStatus {
	return &NullableFileStatus{value: val, isSet: true}
}

func (v NullableFileStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

