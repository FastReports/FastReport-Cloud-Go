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

// CreateGroupAdminVM struct for CreateGroupAdminVM
type CreateGroupAdminVM struct {
	OwnerId *string `json:"ownerId,omitempty"`
	Name string `json:"name"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
}

// NewCreateGroupAdminVM instantiates a new CreateGroupAdminVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupAdminVM(name string) *CreateGroupAdminVM {
	this := CreateGroupAdminVM{}
	this.Name = name
	return &this
}

// NewCreateGroupAdminVMWithDefaults instantiates a new CreateGroupAdminVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupAdminVMWithDefaults() *CreateGroupAdminVM {
	this := CreateGroupAdminVM{}
	return &this
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *CreateGroupAdminVM) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupAdminVM) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *CreateGroupAdminVM) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *CreateGroupAdminVM) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetName returns the Name field value
func (o *CreateGroupAdminVM) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGroupAdminVM) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGroupAdminVM) SetName(v string) {
	o.Name = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *CreateGroupAdminVM) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupAdminVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *CreateGroupAdminVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *CreateGroupAdminVM) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

func (o CreateGroupAdminVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.SubscriptionId != nil {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateGroupAdminVM struct {
	value *CreateGroupAdminVM
	isSet bool
}

func (v NullableCreateGroupAdminVM) Get() *CreateGroupAdminVM {
	return v.value
}

func (v *NullableCreateGroupAdminVM) Set(val *CreateGroupAdminVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupAdminVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupAdminVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupAdminVM(val *CreateGroupAdminVM) *NullableCreateGroupAdminVM {
	return &NullableCreateGroupAdminVM{value: val, isSet: true}
}

func (v NullableCreateGroupAdminVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupAdminVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


