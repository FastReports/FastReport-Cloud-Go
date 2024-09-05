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

// checks if the DataSourcePermissionCRUDVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSourcePermissionCRUDVM{}

// DataSourcePermissionCRUDVM struct for DataSourcePermissionCRUDVM
type DataSourcePermissionCRUDVM struct {
	CloudBaseVM
	Create *DataSourceCreate `json:"create,omitempty"`
	Delete *DataSourceDelete `json:"delete,omitempty"`
	Execute *DataSourceExecute `json:"execute,omitempty"`
	Get *DataSourceGet `json:"get,omitempty"`
	Update *DataSourceUpdate `json:"update,omitempty"`
	Administrate *DataSourceAdministrate `json:"administrate,omitempty"`
	T string `json:"$t"`
}

type _DataSourcePermissionCRUDVM DataSourcePermissionCRUDVM

// NewDataSourcePermissionCRUDVM instantiates a new DataSourcePermissionCRUDVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourcePermissionCRUDVM(t string) *DataSourcePermissionCRUDVM {
	this := DataSourcePermissionCRUDVM{}
	this.T = t
	return &this
}

// NewDataSourcePermissionCRUDVMWithDefaults instantiates a new DataSourcePermissionCRUDVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourcePermissionCRUDVMWithDefaults() *DataSourcePermissionCRUDVM {
	this := DataSourcePermissionCRUDVM{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *DataSourcePermissionCRUDVM) GetCreate() DataSourceCreate {
	if o == nil || IsNil(o.Create) {
		var ret DataSourceCreate
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourcePermissionCRUDVM) GetCreateOk() (*DataSourceCreate, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *DataSourcePermissionCRUDVM) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given DataSourceCreate and assigns it to the Create field.
func (o *DataSourcePermissionCRUDVM) SetCreate(v DataSourceCreate) {
	o.Create = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *DataSourcePermissionCRUDVM) GetDelete() DataSourceDelete {
	if o == nil || IsNil(o.Delete) {
		var ret DataSourceDelete
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourcePermissionCRUDVM) GetDeleteOk() (*DataSourceDelete, bool) {
	if o == nil || IsNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *DataSourcePermissionCRUDVM) HasDelete() bool {
	if o != nil && !IsNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given DataSourceDelete and assigns it to the Delete field.
func (o *DataSourcePermissionCRUDVM) SetDelete(v DataSourceDelete) {
	o.Delete = &v
}

// GetExecute returns the Execute field value if set, zero value otherwise.
func (o *DataSourcePermissionCRUDVM) GetExecute() DataSourceExecute {
	if o == nil || IsNil(o.Execute) {
		var ret DataSourceExecute
		return ret
	}
	return *o.Execute
}

// GetExecuteOk returns a tuple with the Execute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourcePermissionCRUDVM) GetExecuteOk() (*DataSourceExecute, bool) {
	if o == nil || IsNil(o.Execute) {
		return nil, false
	}
	return o.Execute, true
}

// HasExecute returns a boolean if a field has been set.
func (o *DataSourcePermissionCRUDVM) HasExecute() bool {
	if o != nil && !IsNil(o.Execute) {
		return true
	}

	return false
}

// SetExecute gets a reference to the given DataSourceExecute and assigns it to the Execute field.
func (o *DataSourcePermissionCRUDVM) SetExecute(v DataSourceExecute) {
	o.Execute = &v
}

// GetGet returns the Get field value if set, zero value otherwise.
func (o *DataSourcePermissionCRUDVM) GetGet() DataSourceGet {
	if o == nil || IsNil(o.Get) {
		var ret DataSourceGet
		return ret
	}
	return *o.Get
}

// GetGetOk returns a tuple with the Get field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourcePermissionCRUDVM) GetGetOk() (*DataSourceGet, bool) {
	if o == nil || IsNil(o.Get) {
		return nil, false
	}
	return o.Get, true
}

// HasGet returns a boolean if a field has been set.
func (o *DataSourcePermissionCRUDVM) HasGet() bool {
	if o != nil && !IsNil(o.Get) {
		return true
	}

	return false
}

// SetGet gets a reference to the given DataSourceGet and assigns it to the Get field.
func (o *DataSourcePermissionCRUDVM) SetGet(v DataSourceGet) {
	o.Get = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *DataSourcePermissionCRUDVM) GetUpdate() DataSourceUpdate {
	if o == nil || IsNil(o.Update) {
		var ret DataSourceUpdate
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourcePermissionCRUDVM) GetUpdateOk() (*DataSourceUpdate, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *DataSourcePermissionCRUDVM) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given DataSourceUpdate and assigns it to the Update field.
func (o *DataSourcePermissionCRUDVM) SetUpdate(v DataSourceUpdate) {
	o.Update = &v
}

// GetAdministrate returns the Administrate field value if set, zero value otherwise.
func (o *DataSourcePermissionCRUDVM) GetAdministrate() DataSourceAdministrate {
	if o == nil || IsNil(o.Administrate) {
		var ret DataSourceAdministrate
		return ret
	}
	return *o.Administrate
}

// GetAdministrateOk returns a tuple with the Administrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourcePermissionCRUDVM) GetAdministrateOk() (*DataSourceAdministrate, bool) {
	if o == nil || IsNil(o.Administrate) {
		return nil, false
	}
	return o.Administrate, true
}

// HasAdministrate returns a boolean if a field has been set.
func (o *DataSourcePermissionCRUDVM) HasAdministrate() bool {
	if o != nil && !IsNil(o.Administrate) {
		return true
	}

	return false
}

// SetAdministrate gets a reference to the given DataSourceAdministrate and assigns it to the Administrate field.
func (o *DataSourcePermissionCRUDVM) SetAdministrate(v DataSourceAdministrate) {
	o.Administrate = &v
}

// GetT returns the T field value
func (o *DataSourcePermissionCRUDVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *DataSourcePermissionCRUDVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *DataSourcePermissionCRUDVM) SetT(v string) {
	o.T = v
}

func (o DataSourcePermissionCRUDVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSourcePermissionCRUDVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Delete) {
		toSerialize["delete"] = o.Delete
	}
	if !IsNil(o.Execute) {
		toSerialize["execute"] = o.Execute
	}
	if !IsNil(o.Get) {
		toSerialize["get"] = o.Get
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Administrate) {
		toSerialize["administrate"] = o.Administrate
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *DataSourcePermissionCRUDVM) UnmarshalJSON(data []byte) (err error) {
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

	varDataSourcePermissionCRUDVM := _DataSourcePermissionCRUDVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataSourcePermissionCRUDVM)

	if err != nil {
		return err
	}

	*o = DataSourcePermissionCRUDVM(varDataSourcePermissionCRUDVM)

	return err
}

type NullableDataSourcePermissionCRUDVM struct {
	value *DataSourcePermissionCRUDVM
	isSet bool
}

func (v NullableDataSourcePermissionCRUDVM) Get() *DataSourcePermissionCRUDVM {
	return v.value
}

func (v *NullableDataSourcePermissionCRUDVM) Set(val *DataSourcePermissionCRUDVM) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourcePermissionCRUDVM) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourcePermissionCRUDVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourcePermissionCRUDVM(val *DataSourcePermissionCRUDVM) *NullableDataSourcePermissionCRUDVM {
	return &NullableDataSourcePermissionCRUDVM{value: val, isSet: true}
}

func (v NullableDataSourcePermissionCRUDVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourcePermissionCRUDVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


