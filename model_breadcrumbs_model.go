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

// checks if the BreadcrumbsModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BreadcrumbsModel{}

// BreadcrumbsModel struct for BreadcrumbsModel
type BreadcrumbsModel struct {
	Id NullableString `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
}

// NewBreadcrumbsModel instantiates a new BreadcrumbsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBreadcrumbsModel() *BreadcrumbsModel {
	this := BreadcrumbsModel{}
	return &this
}

// NewBreadcrumbsModelWithDefaults instantiates a new BreadcrumbsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBreadcrumbsModelWithDefaults() *BreadcrumbsModel {
	this := BreadcrumbsModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BreadcrumbsModel) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BreadcrumbsModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *BreadcrumbsModel) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *BreadcrumbsModel) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *BreadcrumbsModel) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *BreadcrumbsModel) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BreadcrumbsModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BreadcrumbsModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *BreadcrumbsModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *BreadcrumbsModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *BreadcrumbsModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *BreadcrumbsModel) UnsetName() {
	o.Name.Unset()
}

func (o BreadcrumbsModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BreadcrumbsModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableBreadcrumbsModel struct {
	value *BreadcrumbsModel
	isSet bool
}

func (v NullableBreadcrumbsModel) Get() *BreadcrumbsModel {
	return v.value
}

func (v *NullableBreadcrumbsModel) Set(val *BreadcrumbsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBreadcrumbsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBreadcrumbsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBreadcrumbsModel(val *BreadcrumbsModel) *NullableBreadcrumbsModel {
	return &NullableBreadcrumbsModel{value: val, isSet: true}
}

func (v NullableBreadcrumbsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBreadcrumbsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


