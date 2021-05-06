/*
 * FastReport Cloud
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package FastReport.Cloud.SDK

import (
	"encoding/json"
)

// CreateGroupVM struct for CreateGroupVM
type CreateGroupVM struct {
	Name string `json:"name"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
}

// NewCreateGroupVM instantiates a new CreateGroupVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupVM(name string) *CreateGroupVM {
	this := CreateGroupVM{}
	this.Name = name
	return &this
}

// NewCreateGroupVMWithDefaults instantiates a new CreateGroupVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupVMWithDefaults() *CreateGroupVM {
	this := CreateGroupVM{}
	return &this
}

// GetName returns the Name field value
func (o *CreateGroupVM) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGroupVM) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGroupVM) SetName(v string) {
	o.Name = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *CreateGroupVM) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *CreateGroupVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *CreateGroupVM) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

func (o CreateGroupVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.SubscriptionId != nil {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateGroupVM struct {
	value *CreateGroupVM
	isSet bool
}

func (v NullableCreateGroupVM) Get() *CreateGroupVM {
	return v.value
}

func (v *NullableCreateGroupVM) Set(val *CreateGroupVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupVM(val *CreateGroupVM) *NullableCreateGroupVM {
	return &NullableCreateGroupVM{value: val, isSet: true}
}

func (v NullableCreateGroupVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


