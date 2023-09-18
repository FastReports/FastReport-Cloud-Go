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

// DataSourceAdministrate the model 'DataSourceAdministrate'
type DataSourceAdministrate int32

// List of DataSourceAdministrate
const (
	_0 DataSourceAdministrate = 0
	_1 DataSourceAdministrate = 1
	_2 DataSourceAdministrate = 2
	_4 DataSourceAdministrate = 4
	_8 DataSourceAdministrate = 8
	_MINUS_1 DataSourceAdministrate = -1
)

// All allowed values of DataSourceAdministrate enum
var AllowedDataSourceAdministrateEnumValues = []DataSourceAdministrate{
	0,
	1,
	2,
	4,
	8,
	-1,
}

func (v *DataSourceAdministrate) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataSourceAdministrate(value)
	for _, existing := range AllowedDataSourceAdministrateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataSourceAdministrate", value)
}

// NewDataSourceAdministrateFromValue returns a pointer to a valid DataSourceAdministrate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataSourceAdministrateFromValue(v int32) (*DataSourceAdministrate, error) {
	ev := DataSourceAdministrate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataSourceAdministrate: valid values are %v", v, AllowedDataSourceAdministrateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataSourceAdministrate) IsValid() bool {
	for _, existing := range AllowedDataSourceAdministrateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataSourceAdministrate value
func (v DataSourceAdministrate) Ptr() *DataSourceAdministrate {
	return &v
}

type NullableDataSourceAdministrate struct {
	value *DataSourceAdministrate
	isSet bool
}

func (v NullableDataSourceAdministrate) Get() *DataSourceAdministrate {
	return v.value
}

func (v *NullableDataSourceAdministrate) Set(val *DataSourceAdministrate) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceAdministrate) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceAdministrate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceAdministrate(val *DataSourceAdministrate) *NullableDataSourceAdministrate {
	return &NullableDataSourceAdministrate{value: val, isSet: true}
}

func (v NullableDataSourceAdministrate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceAdministrate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

