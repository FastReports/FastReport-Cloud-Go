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

// DataSourceDelete the model 'DataSourceDelete'
type DataSourceDelete int32

// List of DataSourceDelete
const (
	_0 DataSourceDelete = 0
	_1 DataSourceDelete = 1
	_MINUS_1 DataSourceDelete = -1
)

// All allowed values of DataSourceDelete enum
var AllowedDataSourceDeleteEnumValues = []DataSourceDelete{
	0,
	1,
	-1,
}

func (v *DataSourceDelete) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataSourceDelete(value)
	for _, existing := range AllowedDataSourceDeleteEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataSourceDelete", value)
}

// NewDataSourceDeleteFromValue returns a pointer to a valid DataSourceDelete
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataSourceDeleteFromValue(v int32) (*DataSourceDelete, error) {
	ev := DataSourceDelete(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataSourceDelete: valid values are %v", v, AllowedDataSourceDeleteEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataSourceDelete) IsValid() bool {
	for _, existing := range AllowedDataSourceDeleteEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataSourceDelete value
func (v DataSourceDelete) Ptr() *DataSourceDelete {
	return &v
}

type NullableDataSourceDelete struct {
	value *DataSourceDelete
	isSet bool
}

func (v NullableDataSourceDelete) Get() *DataSourceDelete {
	return v.value
}

func (v *NullableDataSourceDelete) Set(val *DataSourceDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceDelete(val *DataSourceDelete) *NullableDataSourceDelete {
	return &NullableDataSourceDelete{value: val, isSet: true}
}

func (v NullableDataSourceDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

