/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
	"time"
)

// checks if the InvitedUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvitedUser{}

// InvitedUser struct for InvitedUser
type InvitedUser struct {
	UserId NullableString `json:"userId,omitempty"`
	InvitedAt *time.Time `json:"invitedAt,omitempty"`
}

// NewInvitedUser instantiates a new InvitedUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitedUser() *InvitedUser {
	this := InvitedUser{}
	return &this
}

// NewInvitedUserWithDefaults instantiates a new InvitedUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitedUserWithDefaults() *InvitedUser {
	this := InvitedUser{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitedUser) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitedUser) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *InvitedUser) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *InvitedUser) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *InvitedUser) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *InvitedUser) UnsetUserId() {
	o.UserId.Unset()
}

// GetInvitedAt returns the InvitedAt field value if set, zero value otherwise.
func (o *InvitedUser) GetInvitedAt() time.Time {
	if o == nil || IsNil(o.InvitedAt) {
		var ret time.Time
		return ret
	}
	return *o.InvitedAt
}

// GetInvitedAtOk returns a tuple with the InvitedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitedUser) GetInvitedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.InvitedAt) {
		return nil, false
	}
	return o.InvitedAt, true
}

// HasInvitedAt returns a boolean if a field has been set.
func (o *InvitedUser) HasInvitedAt() bool {
	if o != nil && !IsNil(o.InvitedAt) {
		return true
	}

	return false
}

// SetInvitedAt gets a reference to the given time.Time and assigns it to the InvitedAt field.
func (o *InvitedUser) SetInvitedAt(v time.Time) {
	o.InvitedAt = &v
}

func (o InvitedUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvitedUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if !IsNil(o.InvitedAt) {
		toSerialize["invitedAt"] = o.InvitedAt
	}
	return toSerialize, nil
}

type NullableInvitedUser struct {
	value *InvitedUser
	isSet bool
}

func (v NullableInvitedUser) Get() *InvitedUser {
	return v.value
}

func (v *NullableInvitedUser) Set(val *InvitedUser) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitedUser) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitedUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitedUser(val *InvitedUser) *NullableInvitedUser {
	return &NullableInvitedUser{value: val, isSet: true}
}

func (v NullableInvitedUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitedUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


