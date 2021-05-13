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

// DeleteApiKeyVM struct for DeleteApiKeyVM
type DeleteApiKeyVM struct {
	ApiKey string `json:"apiKey"`
}

// NewDeleteApiKeyVM instantiates a new DeleteApiKeyVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteApiKeyVM(apiKey string) *DeleteApiKeyVM {
	this := DeleteApiKeyVM{}
	this.ApiKey = apiKey
	return &this
}

// NewDeleteApiKeyVMWithDefaults instantiates a new DeleteApiKeyVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteApiKeyVMWithDefaults() *DeleteApiKeyVM {
	this := DeleteApiKeyVM{}
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *DeleteApiKeyVM) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *DeleteApiKeyVM) GetApiKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *DeleteApiKeyVM) SetApiKey(v string) {
	o.ApiKey = v
}

func (o DeleteApiKeyVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apiKey"] = o.ApiKey
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteApiKeyVM struct {
	value *DeleteApiKeyVM
	isSet bool
}

func (v NullableDeleteApiKeyVM) Get() *DeleteApiKeyVM {
	return v.value
}

func (v *NullableDeleteApiKeyVM) Set(val *DeleteApiKeyVM) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteApiKeyVM) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteApiKeyVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteApiKeyVM(val *DeleteApiKeyVM) *NullableDeleteApiKeyVM {
	return &NullableDeleteApiKeyVM{value: val, isSet: true}
}

func (v NullableDeleteApiKeyVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteApiKeyVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


