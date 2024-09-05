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

// checks if the FileTagsUpdateVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileTagsUpdateVM{}

// FileTagsUpdateVM struct for FileTagsUpdateVM
type FileTagsUpdateVM struct {
	CloudBaseVM
	Tags []string `json:"tags,omitempty"`
	T string `json:"$t"`
}

type _FileTagsUpdateVM FileTagsUpdateVM

// NewFileTagsUpdateVM instantiates a new FileTagsUpdateVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileTagsUpdateVM(t string) *FileTagsUpdateVM {
	this := FileTagsUpdateVM{}
	this.T = t
	return &this
}

// NewFileTagsUpdateVMWithDefaults instantiates a new FileTagsUpdateVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileTagsUpdateVMWithDefaults() *FileTagsUpdateVM {
	this := FileTagsUpdateVM{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileTagsUpdateVM) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileTagsUpdateVM) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FileTagsUpdateVM) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FileTagsUpdateVM) SetTags(v []string) {
	o.Tags = v
}

// GetT returns the T field value
func (o *FileTagsUpdateVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *FileTagsUpdateVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *FileTagsUpdateVM) SetT(v string) {
	o.T = v
}

func (o FileTagsUpdateVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileTagsUpdateVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *FileTagsUpdateVM) UnmarshalJSON(data []byte) (err error) {
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

	varFileTagsUpdateVM := _FileTagsUpdateVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFileTagsUpdateVM)

	if err != nil {
		return err
	}

	*o = FileTagsUpdateVM(varFileTagsUpdateVM)

	return err
}

type NullableFileTagsUpdateVM struct {
	value *FileTagsUpdateVM
	isSet bool
}

func (v NullableFileTagsUpdateVM) Get() *FileTagsUpdateVM {
	return v.value
}

func (v *NullableFileTagsUpdateVM) Set(val *FileTagsUpdateVM) {
	v.value = val
	v.isSet = true
}

func (v NullableFileTagsUpdateVM) IsSet() bool {
	return v.isSet
}

func (v *NullableFileTagsUpdateVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileTagsUpdateVM(val *FileTagsUpdateVM) *NullableFileTagsUpdateVM {
	return &NullableFileTagsUpdateVM{value: val, isSet: true}
}

func (v NullableFileTagsUpdateVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileTagsUpdateVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


