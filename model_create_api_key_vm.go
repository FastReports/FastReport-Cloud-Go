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
	"time"
)

// CreateApiKeyVM struct for CreateApiKeyVM
type CreateApiKeyVM struct {
	Description *string `json:"description,omitempty"`
	Expired time.Time `json:"expired"`
}

// NewCreateApiKeyVM instantiates a new CreateApiKeyVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApiKeyVM(expired time.Time) *CreateApiKeyVM {
	this := CreateApiKeyVM{}
	this.Expired = expired
	return &this
}

// NewCreateApiKeyVMWithDefaults instantiates a new CreateApiKeyVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApiKeyVMWithDefaults() *CreateApiKeyVM {
	this := CreateApiKeyVM{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateApiKeyVM) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiKeyVM) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateApiKeyVM) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateApiKeyVM) SetDescription(v string) {
	o.Description = &v
}

// GetExpired returns the Expired field value
func (o *CreateApiKeyVM) GetExpired() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value
// and a boolean to check if the value has been set.
func (o *CreateApiKeyVM) GetExpiredOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Expired, true
}

// SetExpired sets field value
func (o *CreateApiKeyVM) SetExpired(v time.Time) {
	o.Expired = v
}

func (o CreateApiKeyVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["expired"] = o.Expired
	}
	return json.Marshal(toSerialize)
}

type NullableCreateApiKeyVM struct {
	value *CreateApiKeyVM
	isSet bool
}

func (v NullableCreateApiKeyVM) Get() *CreateApiKeyVM {
	return v.value
}

func (v *NullableCreateApiKeyVM) Set(val *CreateApiKeyVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApiKeyVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApiKeyVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApiKeyVM(val *CreateApiKeyVM) *NullableCreateApiKeyVM {
	return &NullableCreateApiKeyVM{value: val, isSet: true}
}

func (v NullableCreateApiKeyVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApiKeyVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


