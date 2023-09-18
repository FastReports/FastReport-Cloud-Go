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

// checks if the UpdateFetchTaskVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFetchTaskVM{}

// UpdateFetchTaskVM struct for UpdateFetchTaskVM
type UpdateFetchTaskVM struct {
	DataSourceId NullableString `json:"dataSourceId,omitempty"`
}

// NewUpdateFetchTaskVM instantiates a new UpdateFetchTaskVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFetchTaskVM(t string) *UpdateFetchTaskVM {
	this := UpdateFetchTaskVM{}
	this.T = t
	return &this
}

// NewUpdateFetchTaskVMWithDefaults instantiates a new UpdateFetchTaskVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFetchTaskVMWithDefaults() *UpdateFetchTaskVM {
	this := UpdateFetchTaskVM{}
	return &this
}

// GetDataSourceId returns the DataSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFetchTaskVM) GetDataSourceId() string {
	if o == nil || IsNil(o.DataSourceId.Get()) {
		var ret string
		return ret
	}
	return *o.DataSourceId.Get()
}

// GetDataSourceIdOk returns a tuple with the DataSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFetchTaskVM) GetDataSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataSourceId.Get(), o.DataSourceId.IsSet()
}

// HasDataSourceId returns a boolean if a field has been set.
func (o *UpdateFetchTaskVM) HasDataSourceId() bool {
	if o != nil && o.DataSourceId.IsSet() {
		return true
	}

	return false
}

// SetDataSourceId gets a reference to the given NullableString and assigns it to the DataSourceId field.
func (o *UpdateFetchTaskVM) SetDataSourceId(v string) {
	o.DataSourceId.Set(&v)
}
// SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil
func (o *UpdateFetchTaskVM) SetDataSourceIdNil() {
	o.DataSourceId.Set(nil)
}

// UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
func (o *UpdateFetchTaskVM) UnsetDataSourceId() {
	o.DataSourceId.Unset()
}

func (o UpdateFetchTaskVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFetchTaskVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DataSourceId.IsSet() {
		toSerialize["dataSourceId"] = o.DataSourceId.Get()
	}
	return toSerialize, nil
}

type NullableUpdateFetchTaskVM struct {
	value *UpdateFetchTaskVM
	isSet bool
}

func (v NullableUpdateFetchTaskVM) Get() *UpdateFetchTaskVM {
	return v.value
}

func (v *NullableUpdateFetchTaskVM) Set(val *UpdateFetchTaskVM) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFetchTaskVM) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFetchTaskVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFetchTaskVM(val *UpdateFetchTaskVM) *NullableUpdateFetchTaskVM {
	return &NullableUpdateFetchTaskVM{value: val, isSet: true}
}

func (v NullableUpdateFetchTaskVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFetchTaskVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


