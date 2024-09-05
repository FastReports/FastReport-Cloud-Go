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

// checks if the AnalysisResultVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalysisResultVM{}

// AnalysisResultVM struct for AnalysisResultVM
type AnalysisResultVM struct {
	CloudBaseVM
	Level *ProblemLevel `json:"level,omitempty"`
	Detail NullableString `json:"detail,omitempty"`
	Id NullableString `json:"id,omitempty"`
	SubscriptionId NullableString `json:"subscriptionId,omitempty"`
	FileId NullableString `json:"fileId,omitempty"`
	CollectionName NullableString `json:"collectionName,omitempty"`
	Type *ProblemType `json:"type,omitempty"`
	Signature NullableString `json:"signature,omitempty"`
	NewSize *int64 `json:"newSize,omitempty"`
	T string `json:"$t"`
}

type _AnalysisResultVM AnalysisResultVM

// NewAnalysisResultVM instantiates a new AnalysisResultVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisResultVM(t string) *AnalysisResultVM {
	this := AnalysisResultVM{}
	this.T = t
	return &this
}

// NewAnalysisResultVMWithDefaults instantiates a new AnalysisResultVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisResultVMWithDefaults() *AnalysisResultVM {
	this := AnalysisResultVM{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *AnalysisResultVM) GetLevel() ProblemLevel {
	if o == nil || IsNil(o.Level) {
		var ret ProblemLevel
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisResultVM) GetLevelOk() (*ProblemLevel, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *AnalysisResultVM) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given ProblemLevel and assigns it to the Level field.
func (o *AnalysisResultVM) SetLevel(v ProblemLevel) {
	o.Level = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnalysisResultVM) GetDetail() string {
	if o == nil || IsNil(o.Detail.Get()) {
		var ret string
		return ret
	}
	return *o.Detail.Get()
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnalysisResultVM) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail.Get(), o.Detail.IsSet()
}

// HasDetail returns a boolean if a field has been set.
func (o *AnalysisResultVM) HasDetail() bool {
	if o != nil && o.Detail.IsSet() {
		return true
	}

	return false
}

// SetDetail gets a reference to the given NullableString and assigns it to the Detail field.
func (o *AnalysisResultVM) SetDetail(v string) {
	o.Detail.Set(&v)
}
// SetDetailNil sets the value for Detail to be an explicit nil
func (o *AnalysisResultVM) SetDetailNil() {
	o.Detail.Set(nil)
}

// UnsetDetail ensures that no value is present for Detail, not even an explicit nil
func (o *AnalysisResultVM) UnsetDetail() {
	o.Detail.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnalysisResultVM) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnalysisResultVM) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *AnalysisResultVM) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *AnalysisResultVM) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *AnalysisResultVM) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *AnalysisResultVM) UnsetId() {
	o.Id.Unset()
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnalysisResultVM) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnalysisResultVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *AnalysisResultVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *AnalysisResultVM) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *AnalysisResultVM) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *AnalysisResultVM) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetFileId returns the FileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnalysisResultVM) GetFileId() string {
	if o == nil || IsNil(o.FileId.Get()) {
		var ret string
		return ret
	}
	return *o.FileId.Get()
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnalysisResultVM) GetFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileId.Get(), o.FileId.IsSet()
}

// HasFileId returns a boolean if a field has been set.
func (o *AnalysisResultVM) HasFileId() bool {
	if o != nil && o.FileId.IsSet() {
		return true
	}

	return false
}

// SetFileId gets a reference to the given NullableString and assigns it to the FileId field.
func (o *AnalysisResultVM) SetFileId(v string) {
	o.FileId.Set(&v)
}
// SetFileIdNil sets the value for FileId to be an explicit nil
func (o *AnalysisResultVM) SetFileIdNil() {
	o.FileId.Set(nil)
}

// UnsetFileId ensures that no value is present for FileId, not even an explicit nil
func (o *AnalysisResultVM) UnsetFileId() {
	o.FileId.Unset()
}

// GetCollectionName returns the CollectionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnalysisResultVM) GetCollectionName() string {
	if o == nil || IsNil(o.CollectionName.Get()) {
		var ret string
		return ret
	}
	return *o.CollectionName.Get()
}

// GetCollectionNameOk returns a tuple with the CollectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnalysisResultVM) GetCollectionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CollectionName.Get(), o.CollectionName.IsSet()
}

// HasCollectionName returns a boolean if a field has been set.
func (o *AnalysisResultVM) HasCollectionName() bool {
	if o != nil && o.CollectionName.IsSet() {
		return true
	}

	return false
}

// SetCollectionName gets a reference to the given NullableString and assigns it to the CollectionName field.
func (o *AnalysisResultVM) SetCollectionName(v string) {
	o.CollectionName.Set(&v)
}
// SetCollectionNameNil sets the value for CollectionName to be an explicit nil
func (o *AnalysisResultVM) SetCollectionNameNil() {
	o.CollectionName.Set(nil)
}

// UnsetCollectionName ensures that no value is present for CollectionName, not even an explicit nil
func (o *AnalysisResultVM) UnsetCollectionName() {
	o.CollectionName.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AnalysisResultVM) GetType() ProblemType {
	if o == nil || IsNil(o.Type) {
		var ret ProblemType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisResultVM) GetTypeOk() (*ProblemType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AnalysisResultVM) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ProblemType and assigns it to the Type field.
func (o *AnalysisResultVM) SetType(v ProblemType) {
	o.Type = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnalysisResultVM) GetSignature() string {
	if o == nil || IsNil(o.Signature.Get()) {
		var ret string
		return ret
	}
	return *o.Signature.Get()
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnalysisResultVM) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signature.Get(), o.Signature.IsSet()
}

// HasSignature returns a boolean if a field has been set.
func (o *AnalysisResultVM) HasSignature() bool {
	if o != nil && o.Signature.IsSet() {
		return true
	}

	return false
}

// SetSignature gets a reference to the given NullableString and assigns it to the Signature field.
func (o *AnalysisResultVM) SetSignature(v string) {
	o.Signature.Set(&v)
}
// SetSignatureNil sets the value for Signature to be an explicit nil
func (o *AnalysisResultVM) SetSignatureNil() {
	o.Signature.Set(nil)
}

// UnsetSignature ensures that no value is present for Signature, not even an explicit nil
func (o *AnalysisResultVM) UnsetSignature() {
	o.Signature.Unset()
}

// GetNewSize returns the NewSize field value if set, zero value otherwise.
func (o *AnalysisResultVM) GetNewSize() int64 {
	if o == nil || IsNil(o.NewSize) {
		var ret int64
		return ret
	}
	return *o.NewSize
}

// GetNewSizeOk returns a tuple with the NewSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisResultVM) GetNewSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.NewSize) {
		return nil, false
	}
	return o.NewSize, true
}

// HasNewSize returns a boolean if a field has been set.
func (o *AnalysisResultVM) HasNewSize() bool {
	if o != nil && !IsNil(o.NewSize) {
		return true
	}

	return false
}

// SetNewSize gets a reference to the given int64 and assigns it to the NewSize field.
func (o *AnalysisResultVM) SetNewSize(v int64) {
	o.NewSize = &v
}

// GetT returns the T field value
func (o *AnalysisResultVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *AnalysisResultVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *AnalysisResultVM) SetT(v string) {
	o.T = v
}

func (o AnalysisResultVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalysisResultVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if o.Detail.IsSet() {
		toSerialize["detail"] = o.Detail.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscriptionId"] = o.SubscriptionId.Get()
	}
	if o.FileId.IsSet() {
		toSerialize["fileId"] = o.FileId.Get()
	}
	if o.CollectionName.IsSet() {
		toSerialize["collectionName"] = o.CollectionName.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Signature.IsSet() {
		toSerialize["signature"] = o.Signature.Get()
	}
	if !IsNil(o.NewSize) {
		toSerialize["newSize"] = o.NewSize
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *AnalysisResultVM) UnmarshalJSON(data []byte) (err error) {
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

	varAnalysisResultVM := _AnalysisResultVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAnalysisResultVM)

	if err != nil {
		return err
	}

	*o = AnalysisResultVM(varAnalysisResultVM)

	return err
}

type NullableAnalysisResultVM struct {
	value *AnalysisResultVM
	isSet bool
}

func (v NullableAnalysisResultVM) Get() *AnalysisResultVM {
	return v.value
}

func (v *NullableAnalysisResultVM) Set(val *AnalysisResultVM) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisResultVM) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisResultVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisResultVM(val *AnalysisResultVM) *NullableAnalysisResultVM {
	return &NullableAnalysisResultVM{value: val, isSet: true}
}

func (v NullableAnalysisResultVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisResultVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

