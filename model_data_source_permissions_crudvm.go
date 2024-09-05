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

// checks if the DataSourcePermissionsCRUDVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSourcePermissionsCRUDVM{}

// DataSourcePermissionsCRUDVM struct for DataSourcePermissionsCRUDVM
type DataSourcePermissionsCRUDVM struct {
	CloudBaseVM
	OwnerId NullableString `json:"ownerId,omitempty"`
	Owner *DataSourcePermissionCRUDVM `json:"owner,omitempty"`
	Groups map[string]DataSourcePermissionCRUDVM `json:"groups,omitempty"`
	Other *DataSourcePermissionCRUDVM `json:"other,omitempty"`
	Anon *DataSourcePermissionCRUDVM `json:"anon,omitempty"`
	T string `json:"$t"`
}

type _DataSourcePermissionsCRUDVM DataSourcePermissionsCRUDVM

// NewDataSourcePermissionsCRUDVM instantiates a new DataSourcePermissionsCRUDVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourcePermissionsCRUDVM(t string) *DataSourcePermissionsCRUDVM {
	this := DataSourcePermissionsCRUDVM{}
	this.T = t
	return &this
}

// NewDataSourcePermissionsCRUDVMWithDefaults instantiates a new DataSourcePermissionsCRUDVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourcePermissionsCRUDVMWithDefaults() *DataSourcePermissionsCRUDVM {
	this := DataSourcePermissionsCRUDVM{}
	return &this
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataSourcePermissionsCRUDVM) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerId.Get()
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSourcePermissionsCRUDVM) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerId.Get(), o.OwnerId.IsSet()
}

// HasOwnerId returns a boolean if a field has been set.
func (o *DataSourcePermissionsCRUDVM) HasOwnerId() bool {
	if o != nil && o.OwnerId.IsSet() {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given NullableString and assigns it to the OwnerId field.
func (o *DataSourcePermissionsCRUDVM) SetOwnerId(v string) {
	o.OwnerId.Set(&v)
}
// SetOwnerIdNil sets the value for OwnerId to be an explicit nil
func (o *DataSourcePermissionsCRUDVM) SetOwnerIdNil() {
	o.OwnerId.Set(nil)
}

// UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
func (o *DataSourcePermissionsCRUDVM) UnsetOwnerId() {
	o.OwnerId.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *DataSourcePermissionsCRUDVM) GetOwner() DataSourcePermissionCRUDVM {
	if o == nil || IsNil(o.Owner) {
		var ret DataSourcePermissionCRUDVM
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourcePermissionsCRUDVM) GetOwnerOk() (*DataSourcePermissionCRUDVM, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *DataSourcePermissionsCRUDVM) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given DataSourcePermissionCRUDVM and assigns it to the Owner field.
func (o *DataSourcePermissionsCRUDVM) SetOwner(v DataSourcePermissionCRUDVM) {
	o.Owner = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataSourcePermissionsCRUDVM) GetGroups() map[string]DataSourcePermissionCRUDVM {
	if o == nil {
		var ret map[string]DataSourcePermissionCRUDVM
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSourcePermissionsCRUDVM) GetGroupsOk() (*map[string]DataSourcePermissionCRUDVM, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *DataSourcePermissionsCRUDVM) HasGroups() bool {
	if o != nil && IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given map[string]DataSourcePermissionCRUDVM and assigns it to the Groups field.
func (o *DataSourcePermissionsCRUDVM) SetGroups(v map[string]DataSourcePermissionCRUDVM) {
	o.Groups = v
}

// GetOther returns the Other field value if set, zero value otherwise.
func (o *DataSourcePermissionsCRUDVM) GetOther() DataSourcePermissionCRUDVM {
	if o == nil || IsNil(o.Other) {
		var ret DataSourcePermissionCRUDVM
		return ret
	}
	return *o.Other
}

// GetOtherOk returns a tuple with the Other field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourcePermissionsCRUDVM) GetOtherOk() (*DataSourcePermissionCRUDVM, bool) {
	if o == nil || IsNil(o.Other) {
		return nil, false
	}
	return o.Other, true
}

// HasOther returns a boolean if a field has been set.
func (o *DataSourcePermissionsCRUDVM) HasOther() bool {
	if o != nil && !IsNil(o.Other) {
		return true
	}

	return false
}

// SetOther gets a reference to the given DataSourcePermissionCRUDVM and assigns it to the Other field.
func (o *DataSourcePermissionsCRUDVM) SetOther(v DataSourcePermissionCRUDVM) {
	o.Other = &v
}

// GetAnon returns the Anon field value if set, zero value otherwise.
func (o *DataSourcePermissionsCRUDVM) GetAnon() DataSourcePermissionCRUDVM {
	if o == nil || IsNil(o.Anon) {
		var ret DataSourcePermissionCRUDVM
		return ret
	}
	return *o.Anon
}

// GetAnonOk returns a tuple with the Anon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourcePermissionsCRUDVM) GetAnonOk() (*DataSourcePermissionCRUDVM, bool) {
	if o == nil || IsNil(o.Anon) {
		return nil, false
	}
	return o.Anon, true
}

// HasAnon returns a boolean if a field has been set.
func (o *DataSourcePermissionsCRUDVM) HasAnon() bool {
	if o != nil && !IsNil(o.Anon) {
		return true
	}

	return false
}

// SetAnon gets a reference to the given DataSourcePermissionCRUDVM and assigns it to the Anon field.
func (o *DataSourcePermissionsCRUDVM) SetAnon(v DataSourcePermissionCRUDVM) {
	o.Anon = &v
}

// GetT returns the T field value
func (o *DataSourcePermissionsCRUDVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *DataSourcePermissionsCRUDVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *DataSourcePermissionsCRUDVM) SetT(v string) {
	o.T = v
}

func (o DataSourcePermissionsCRUDVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSourcePermissionsCRUDVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if o.OwnerId.IsSet() {
		toSerialize["ownerId"] = o.OwnerId.Get()
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Other) {
		toSerialize["other"] = o.Other
	}
	if !IsNil(o.Anon) {
		toSerialize["anon"] = o.Anon
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *DataSourcePermissionsCRUDVM) UnmarshalJSON(data []byte) (err error) {
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

	varDataSourcePermissionsCRUDVM := _DataSourcePermissionsCRUDVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataSourcePermissionsCRUDVM)

	if err != nil {
		return err
	}

	*o = DataSourcePermissionsCRUDVM(varDataSourcePermissionsCRUDVM)

	return err
}

type NullableDataSourcePermissionsCRUDVM struct {
	value *DataSourcePermissionsCRUDVM
	isSet bool
}

func (v NullableDataSourcePermissionsCRUDVM) Get() *DataSourcePermissionsCRUDVM {
	return v.value
}

func (v *NullableDataSourcePermissionsCRUDVM) Set(val *DataSourcePermissionsCRUDVM) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourcePermissionsCRUDVM) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourcePermissionsCRUDVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourcePermissionsCRUDVM(val *DataSourcePermissionsCRUDVM) *NullableDataSourcePermissionsCRUDVM {
	return &NullableDataSourcePermissionsCRUDVM{value: val, isSet: true}
}

func (v NullableDataSourcePermissionsCRUDVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourcePermissionsCRUDVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

