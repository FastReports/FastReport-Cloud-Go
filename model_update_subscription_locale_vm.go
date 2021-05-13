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
)

// UpdateSubscriptionLocaleVM struct for UpdateSubscriptionLocaleVM
type UpdateSubscriptionLocaleVM struct {
	Locale string `json:"locale"`
}

// NewUpdateSubscriptionLocaleVM instantiates a new UpdateSubscriptionLocaleVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSubscriptionLocaleVM(locale string) *UpdateSubscriptionLocaleVM {
	this := UpdateSubscriptionLocaleVM{}
	this.Locale = locale
	return &this
}

// NewUpdateSubscriptionLocaleVMWithDefaults instantiates a new UpdateSubscriptionLocaleVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSubscriptionLocaleVMWithDefaults() *UpdateSubscriptionLocaleVM {
	this := UpdateSubscriptionLocaleVM{}
	return &this
}

// GetLocale returns the Locale field value
func (o *UpdateSubscriptionLocaleVM) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionLocaleVM) GetLocaleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *UpdateSubscriptionLocaleVM) SetLocale(v string) {
	o.Locale = v
}

func (o UpdateSubscriptionLocaleVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["locale"] = o.Locale
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSubscriptionLocaleVM struct {
	value *UpdateSubscriptionLocaleVM
	isSet bool
}

func (v NullableUpdateSubscriptionLocaleVM) Get() *UpdateSubscriptionLocaleVM {
	return v.value
}

func (v *NullableUpdateSubscriptionLocaleVM) Set(val *UpdateSubscriptionLocaleVM) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSubscriptionLocaleVM) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSubscriptionLocaleVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSubscriptionLocaleVM(val *UpdateSubscriptionLocaleVM) *NullableUpdateSubscriptionLocaleVM {
	return &NullableUpdateSubscriptionLocaleVM{value: val, isSet: true}
}

func (v NullableUpdateSubscriptionLocaleVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSubscriptionLocaleVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


