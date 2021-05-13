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

// FilesVM struct for FilesVM
type FilesVM struct {
	Files *[]FileVM `json:"files,omitempty"`
	Count *int64 `json:"count,omitempty"`
	Skip *int32 `json:"skip,omitempty"`
	Take *int32 `json:"take,omitempty"`
}

// NewFilesVM instantiates a new FilesVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesVM() *FilesVM {
	this := FilesVM{}
	return &this
}

// NewFilesVMWithDefaults instantiates a new FilesVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesVMWithDefaults() *FilesVM {
	this := FilesVM{}
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *FilesVM) GetFiles() []FileVM {
	if o == nil || o.Files == nil {
		var ret []FileVM
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesVM) GetFilesOk() (*[]FileVM, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *FilesVM) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FileVM and assigns it to the Files field.
func (o *FilesVM) SetFiles(v []FileVM) {
	o.Files = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *FilesVM) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesVM) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *FilesVM) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *FilesVM) SetCount(v int64) {
	o.Count = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *FilesVM) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesVM) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *FilesVM) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *FilesVM) SetSkip(v int32) {
	o.Skip = &v
}

// GetTake returns the Take field value if set, zero value otherwise.
func (o *FilesVM) GetTake() int32 {
	if o == nil || o.Take == nil {
		var ret int32
		return ret
	}
	return *o.Take
}

// GetTakeOk returns a tuple with the Take field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesVM) GetTakeOk() (*int32, bool) {
	if o == nil || o.Take == nil {
		return nil, false
	}
	return o.Take, true
}

// HasTake returns a boolean if a field has been set.
func (o *FilesVM) HasTake() bool {
	if o != nil && o.Take != nil {
		return true
	}

	return false
}

// SetTake gets a reference to the given int32 and assigns it to the Take field.
func (o *FilesVM) SetTake(v int32) {
	o.Take = &v
}

func (o FilesVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Files != nil {
		toSerialize["files"] = o.Files
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

type NullableFilesVM struct {
	value *FilesVM
	isSet bool
}

func (v NullableFilesVM) Get() *FilesVM {
	return v.value
}

func (v *NullableFilesVM) Set(val *FilesVM) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesVM) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesVM(val *FilesVM) *NullableFilesVM {
	return &NullableFilesVM{value: val, isSet: true}
}

func (v NullableFilesVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


