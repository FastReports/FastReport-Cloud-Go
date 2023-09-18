/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
)

// checks if the GroupUsersVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupUsersVM{}

// GroupUsersVM struct for GroupUsersVM
type GroupUsersVM struct {
	Users []GroupUserVM `json:"users,omitempty"`
	Count *int64 `json:"count,omitempty"`
	Take *int32 `json:"take,omitempty"`
	Skip *int32 `json:"skip,omitempty"`
}

// NewGroupUsersVM instantiates a new GroupUsersVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUsersVM() *GroupUsersVM {
	this := GroupUsersVM{}
	return &this
}

// NewGroupUsersVMWithDefaults instantiates a new GroupUsersVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUsersVMWithDefaults() *GroupUsersVM {
	this := GroupUsersVM{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupUsersVM) GetUsers() []GroupUserVM {
	if o == nil {
		var ret []GroupUserVM
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUsersVM) GetUsersOk() ([]GroupUserVM, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *GroupUsersVM) HasUsers() bool {
	if o != nil && IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []GroupUserVM and assigns it to the Users field.
func (o *GroupUsersVM) SetUsers(v []GroupUserVM) {
	o.Users = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GroupUsersVM) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUsersVM) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GroupUsersVM) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *GroupUsersVM) SetCount(v int64) {
	o.Count = &v
}

// GetTake returns the Take field value if set, zero value otherwise.
func (o *GroupUsersVM) GetTake() int32 {
	if o == nil || IsNil(o.Take) {
		var ret int32
		return ret
	}
	return *o.Take
}

// GetTakeOk returns a tuple with the Take field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUsersVM) GetTakeOk() (*int32, bool) {
	if o == nil || IsNil(o.Take) {
		return nil, false
	}
	return o.Take, true
}

// HasTake returns a boolean if a field has been set.
func (o *GroupUsersVM) HasTake() bool {
	if o != nil && !IsNil(o.Take) {
		return true
	}

	return false
}

// SetTake gets a reference to the given int32 and assigns it to the Take field.
func (o *GroupUsersVM) SetTake(v int32) {
	o.Take = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *GroupUsersVM) GetSkip() int32 {
	if o == nil || IsNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUsersVM) GetSkipOk() (*int32, bool) {
	if o == nil || IsNil(o.Skip) {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *GroupUsersVM) HasSkip() bool {
	if o != nil && !IsNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *GroupUsersVM) SetSkip(v int32) {
	o.Skip = &v
}

func (o GroupUsersVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupUsersVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Take) {
		toSerialize["take"] = o.Take
	}
	if !IsNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	return toSerialize, nil
}

type NullableGroupUsersVM struct {
	value *GroupUsersVM
	isSet bool
}

func (v NullableGroupUsersVM) Get() *GroupUsersVM {
	return v.value
}

func (v *NullableGroupUsersVM) Set(val *GroupUsersVM) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUsersVM) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUsersVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUsersVM(val *GroupUsersVM) *NullableGroupUsersVM {
	return &NullableGroupUsersVM{value: val, isSet: true}
}

func (v NullableGroupUsersVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUsersVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


