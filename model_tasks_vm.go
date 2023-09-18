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

// checks if the TasksVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TasksVM{}

// TasksVM struct for TasksVM
type TasksVM struct {
	Count *int64 `json:"count,omitempty"`
	Skip *int32 `json:"skip,omitempty"`
	Take *int32 `json:"take,omitempty"`
	Tasks []TaskBaseVM `json:"tasks,omitempty"`
}

// NewTasksVM instantiates a new TasksVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTasksVM() *TasksVM {
	this := TasksVM{}
	return &this
}

// NewTasksVMWithDefaults instantiates a new TasksVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTasksVMWithDefaults() *TasksVM {
	this := TasksVM{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *TasksVM) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksVM) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *TasksVM) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *TasksVM) SetCount(v int64) {
	o.Count = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *TasksVM) GetSkip() int32 {
	if o == nil || IsNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksVM) GetSkipOk() (*int32, bool) {
	if o == nil || IsNil(o.Skip) {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *TasksVM) HasSkip() bool {
	if o != nil && !IsNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *TasksVM) SetSkip(v int32) {
	o.Skip = &v
}

// GetTake returns the Take field value if set, zero value otherwise.
func (o *TasksVM) GetTake() int32 {
	if o == nil || IsNil(o.Take) {
		var ret int32
		return ret
	}
	return *o.Take
}

// GetTakeOk returns a tuple with the Take field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksVM) GetTakeOk() (*int32, bool) {
	if o == nil || IsNil(o.Take) {
		return nil, false
	}
	return o.Take, true
}

// HasTake returns a boolean if a field has been set.
func (o *TasksVM) HasTake() bool {
	if o != nil && !IsNil(o.Take) {
		return true
	}

	return false
}

// SetTake gets a reference to the given int32 and assigns it to the Take field.
func (o *TasksVM) SetTake(v int32) {
	o.Take = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TasksVM) GetTasks() []TaskBaseVM {
	if o == nil {
		var ret []TaskBaseVM
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TasksVM) GetTasksOk() ([]TaskBaseVM, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *TasksVM) HasTasks() bool {
	if o != nil && IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []TaskBaseVM and assigns it to the Tasks field.
func (o *TasksVM) SetTasks(v []TaskBaseVM) {
	o.Tasks = v
}

func (o TasksVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TasksVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !IsNil(o.Take) {
		toSerialize["take"] = o.Take
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return toSerialize, nil
}

type NullableTasksVM struct {
	value *TasksVM
	isSet bool
}

func (v NullableTasksVM) Get() *TasksVM {
	return v.value
}

func (v *NullableTasksVM) Set(val *TasksVM) {
	v.value = val
	v.isSet = true
}

func (v NullableTasksVM) IsSet() bool {
	return v.isSet
}

func (v *NullableTasksVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTasksVM(val *TasksVM) *NullableTasksVM {
	return &NullableTasksVM{value: val, isSet: true}
}

func (v NullableTasksVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTasksVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


