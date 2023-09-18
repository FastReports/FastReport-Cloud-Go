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

// SubscriptionAdministrate the model 'SubscriptionAdministrate'
type SubscriptionAdministrate int32

// List of SubscriptionAdministrate
const (
	_0 SubscriptionAdministrate = 0
	_1 SubscriptionAdministrate = 1
	_2 SubscriptionAdministrate = 2
	_4 SubscriptionAdministrate = 4
	_8 SubscriptionAdministrate = 8
	_MINUS_1 SubscriptionAdministrate = -1
)

// All allowed values of SubscriptionAdministrate enum
var AllowedSubscriptionAdministrateEnumValues = []SubscriptionAdministrate{
	0,
	1,
	2,
	4,
	8,
	-1,
}

func (v *SubscriptionAdministrate) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscriptionAdministrate(value)
	for _, existing := range AllowedSubscriptionAdministrateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscriptionAdministrate", value)
}

// NewSubscriptionAdministrateFromValue returns a pointer to a valid SubscriptionAdministrate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscriptionAdministrateFromValue(v int32) (*SubscriptionAdministrate, error) {
	ev := SubscriptionAdministrate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscriptionAdministrate: valid values are %v", v, AllowedSubscriptionAdministrateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscriptionAdministrate) IsValid() bool {
	for _, existing := range AllowedSubscriptionAdministrateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubscriptionAdministrate value
func (v SubscriptionAdministrate) Ptr() *SubscriptionAdministrate {
	return &v
}

type NullableSubscriptionAdministrate struct {
	value *SubscriptionAdministrate
	isSet bool
}

func (v NullableSubscriptionAdministrate) Get() *SubscriptionAdministrate {
	return v.value
}

func (v *NullableSubscriptionAdministrate) Set(val *SubscriptionAdministrate) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionAdministrate) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionAdministrate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionAdministrate(val *SubscriptionAdministrate) *NullableSubscriptionAdministrate {
	return &NullableSubscriptionAdministrate{value: val, isSet: true}
}

func (v NullableSubscriptionAdministrate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionAdministrate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

