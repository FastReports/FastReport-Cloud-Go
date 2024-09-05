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

// ProblemType the model 'ProblemType'
type ProblemType string

// List of ProblemType
const (
	ANON_PERMISSIONS ProblemType = "AnonPermissions"
	UNRELATED_FILES ProblemType = "UnrelatedFiles"
	OTHER_WITHOUT_GET ProblemType = "OtherWithoutGet"
	DEAD_SUB_IN_USER ProblemType = "DeadSubInUser"
	EMPTY_SUB ProblemType = "EmptySub"
	LOST_FILE_CHUNKS ProblemType = "LostFileChunks"
	WRONG_SUBSCRIPITON_SIZES ProblemType = "WrongSubscripitonSizes"
	FILES_WITHOUT_CHUNKS ProblemType = "FilesWithoutChunks"
	FILES_WITH_DELETED_PARENTS ProblemType = "FilesWithDeletedParents"
)

// All allowed values of ProblemType enum
var AllowedProblemTypeEnumValues = []ProblemType{
	"AnonPermissions",
	"UnrelatedFiles",
	"OtherWithoutGet",
	"DeadSubInUser",
	"EmptySub",
	"LostFileChunks",
	"WrongSubscripitonSizes",
	"FilesWithoutChunks",
	"FilesWithDeletedParents",
}

func (v *ProblemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProblemType(value)
	for _, existing := range AllowedProblemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProblemType", value)
}

// NewProblemTypeFromValue returns a pointer to a valid ProblemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProblemTypeFromValue(v string) (*ProblemType, error) {
	ev := ProblemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProblemType: valid values are %v", v, AllowedProblemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProblemType) IsValid() bool {
	for _, existing := range AllowedProblemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProblemType value
func (v ProblemType) Ptr() *ProblemType {
	return &v
}

type NullableProblemType struct {
	value *ProblemType
	isSet bool
}

func (v NullableProblemType) Get() *ProblemType {
	return v.value
}

func (v *NullableProblemType) Set(val *ProblemType) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemType) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemType(val *ProblemType) *NullableProblemType {
	return &NullableProblemType{value: val, isSet: true}
}

func (v NullableProblemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

