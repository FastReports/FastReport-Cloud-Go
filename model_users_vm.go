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

// UsersVM struct for UsersVM
type UsersVM struct {
	Users *[]UserVM `json:"users,omitempty"`
	Count *int64 `json:"count,omitempty"`
	Skip *int32 `json:"skip,omitempty"`
	Take *int32 `json:"take,omitempty"`
}

// NewUsersVM instantiates a new UsersVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersVM() *UsersVM {
	this := UsersVM{}
	return &this
}

// NewUsersVMWithDefaults instantiates a new UsersVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersVMWithDefaults() *UsersVM {
	this := UsersVM{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *UsersVM) GetUsers() []UserVM {
	if o == nil || o.Users == nil {
		var ret []UserVM
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersVM) GetUsersOk() (*[]UserVM, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UsersVM) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UserVM and assigns it to the Users field.
func (o *UsersVM) SetUsers(v []UserVM) {
	o.Users = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UsersVM) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersVM) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UsersVM) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *UsersVM) SetCount(v int64) {
	o.Count = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *UsersVM) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersVM) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *UsersVM) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *UsersVM) SetSkip(v int32) {
	o.Skip = &v
}

// GetTake returns the Take field value if set, zero value otherwise.
func (o *UsersVM) GetTake() int32 {
	if o == nil || o.Take == nil {
		var ret int32
		return ret
	}
	return *o.Take
}

// GetTakeOk returns a tuple with the Take field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersVM) GetTakeOk() (*int32, bool) {
	if o == nil || o.Take == nil {
		return nil, false
	}
	return o.Take, true
}

// HasTake returns a boolean if a field has been set.
func (o *UsersVM) HasTake() bool {
	if o != nil && o.Take != nil {
		return true
	}

	return false
}

// SetTake gets a reference to the given int32 and assigns it to the Take field.
func (o *UsersVM) SetTake(v int32) {
	o.Take = &v
}

func (o UsersVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Skip != nil {
		toSerialize["skip"] = o.Skip
	}
	if o.Take != nil {
		toSerialize["take"] = o.Take
	}
	return json.Marshal(toSerialize)
}

type NullableUsersVM struct {
	value *UsersVM
	isSet bool
}

func (v NullableUsersVM) Get() *UsersVM {
	return v.value
}

func (v *NullableUsersVM) Set(val *UsersVM) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersVM) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersVM(val *UsersVM) *NullableUsersVM {
	return &NullableUsersVM{value: val, isSet: true}
}

func (v NullableUsersVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

