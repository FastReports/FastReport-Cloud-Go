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

// DataSourceConnectionType the model 'DataSourceConnectionType'
type DataSourceConnectionType string

// List of DataSourceConnectionType
const (
	JSON DataSourceConnectionType = "JSON"
	MSSQL DataSourceConnectionType = "MSSQL"
	CSV DataSourceConnectionType = "CSV"
	XML DataSourceConnectionType = "XML"
	MY_SQL DataSourceConnectionType = "MySQL"
	POSTGRES DataSourceConnectionType = "Postgres"
	ORACLE_DB DataSourceConnectionType = "OracleDB"
)

var allowedDataSourceConnectionTypeEnumValues = []DataSourceConnectionType{
	"JSON",
	"MSSQL",
	"CSV",
	"XML",
	"MySQL",
	"Postgres",
	"OracleDB",
}

func (v *DataSourceConnectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataSourceConnectionType(value)
	for _, existing := range allowedDataSourceConnectionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataSourceConnectionType", value)
}

// NewDataSourceConnectionTypeFromValue returns a pointer to a valid DataSourceConnectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataSourceConnectionTypeFromValue(v string) (*DataSourceConnectionType, error) {
	ev := DataSourceConnectionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataSourceConnectionType: valid values are %v", v, allowedDataSourceConnectionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataSourceConnectionType) IsValid() bool {
	for _, existing := range allowedDataSourceConnectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataSourceConnectionType value
func (v DataSourceConnectionType) Ptr() *DataSourceConnectionType {
	return &v
}

type NullableDataSourceConnectionType struct {
	value *DataSourceConnectionType
	isSet bool
}

func (v NullableDataSourceConnectionType) Get() *DataSourceConnectionType {
	return v.value
}

func (v *NullableDataSourceConnectionType) Set(val *DataSourceConnectionType) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceConnectionType) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceConnectionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceConnectionType(val *DataSourceConnectionType) *NullableDataSourceConnectionType {
	return &NullableDataSourceConnectionType{value: val, isSet: true}
}

func (v NullableDataSourceConnectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceConnectionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
