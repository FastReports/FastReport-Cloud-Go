/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GroupsVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsVM{}

// GroupsVM struct for GroupsVM
type GroupsVM struct {
	CloudBaseVM
	Groups []GroupVM `json:"groups,omitempty"`
	Count *int64 `json:"count,omitempty"`
	Skip *int32 `json:"skip,omitempty"`
	Take *int32 `json:"take,omitempty"`
	T string `json:"$t"`
}

type _GroupsVM GroupsVM

// NewGroupsVM instantiates a new GroupsVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsVM(t string) *GroupsVM {
	this := GroupsVM{}
	this.T = t
	return &this
}

// NewGroupsVMWithDefaults instantiates a new GroupsVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsVMWithDefaults() *GroupsVM {
	this := GroupsVM{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupsVM) GetGroups() []GroupVM {
	if o == nil {
		var ret []GroupVM
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsVM) GetGroupsOk() ([]GroupVM, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *GroupsVM) HasGroups() bool {
	if o != nil && IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []GroupVM and assigns it to the Groups field.
func (o *GroupsVM) SetGroups(v []GroupVM) {
	o.Groups = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GroupsVM) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsVM) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GroupsVM) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *GroupsVM) SetCount(v int64) {
	o.Count = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *GroupsVM) GetSkip() int32 {
	if o == nil || IsNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsVM) GetSkipOk() (*int32, bool) {
	if o == nil || IsNil(o.Skip) {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *GroupsVM) HasSkip() bool {
	if o != nil && !IsNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *GroupsVM) SetSkip(v int32) {
	o.Skip = &v
}

// GetTake returns the Take field value if set, zero value otherwise.
func (o *GroupsVM) GetTake() int32 {
	if o == nil || IsNil(o.Take) {
		var ret int32
		return ret
	}
	return *o.Take
}

// GetTakeOk returns a tuple with the Take field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsVM) GetTakeOk() (*int32, bool) {
	if o == nil || IsNil(o.Take) {
		return nil, false
	}
	return o.Take, true
}

// HasTake returns a boolean if a field has been set.
func (o *GroupsVM) HasTake() bool {
	if o != nil && !IsNil(o.Take) {
		return true
	}

	return false
}

// SetTake gets a reference to the given int32 and assigns it to the Take field.
func (o *GroupsVM) SetTake(v int32) {
	o.Take = &v
}

// GetT returns the T field value
func (o *GroupsVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *GroupsVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *GroupsVM) SetT(v string) {
	o.T = v
}

func (o GroupsVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !IsNil(o.Take) {
		toSerialize["take"] = o.Take
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *GroupsVM) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"$t",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGroupsVM := _GroupsVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupsVM)

	if err != nil {
		return err
	}

	*o = GroupsVM(varGroupsVM)

	return err
}

type NullableGroupsVM struct {
	value *GroupsVM
	isSet bool
}

func (v NullableGroupsVM) Get() *GroupsVM {
	return v.value
}

func (v *NullableGroupsVM) Set(val *GroupsVM) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsVM) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsVM(val *GroupsVM) *NullableGroupsVM {
	return &NullableGroupsVM{value: val, isSet: true}
}

func (v NullableGroupsVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


