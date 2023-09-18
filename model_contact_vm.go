/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
	"time"
)

// checks if the ContactVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactVM{}

// ContactVM struct for ContactVM
type ContactVM struct {
	Id NullableString `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Email NullableString `json:"email,omitempty"`
	Groups []string `json:"groups,omitempty"`
	SubscriptionId NullableString `json:"subscriptionId,omitempty"`
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	CreatorUserId NullableString `json:"creatorUserId,omitempty"`
	EditedTime *time.Time `json:"editedTime,omitempty"`
	EditorUserId NullableString `json:"editorUserId,omitempty"`
}

// NewContactVM instantiates a new ContactVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactVM() *ContactVM {
	this := ContactVM{}
	return &this
}

// NewContactVMWithDefaults instantiates a new ContactVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactVMWithDefaults() *ContactVM {
	this := ContactVM{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactVM) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactVM) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ContactVM) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ContactVM) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ContactVM) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ContactVM) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactVM) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactVM) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ContactVM) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ContactVM) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ContactVM) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ContactVM) UnsetName() {
	o.Name.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactVM) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactVM) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *ContactVM) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *ContactVM) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *ContactVM) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *ContactVM) UnsetEmail() {
	o.Email.Unset()
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactVM) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactVM) GetGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ContactVM) HasGroups() bool {
	if o != nil && IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *ContactVM) SetGroups(v []string) {
	o.Groups = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactVM) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *ContactVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *ContactVM) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *ContactVM) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *ContactVM) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *ContactVM) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactVM) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *ContactVM) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *ContactVM) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetCreatorUserId returns the CreatorUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactVM) GetCreatorUserId() string {
	if o == nil || IsNil(o.CreatorUserId.Get()) {
		var ret string
		return ret
	}
	return *o.CreatorUserId.Get()
}

// GetCreatorUserIdOk returns a tuple with the CreatorUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactVM) GetCreatorUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorUserId.Get(), o.CreatorUserId.IsSet()
}

// HasCreatorUserId returns a boolean if a field has been set.
func (o *ContactVM) HasCreatorUserId() bool {
	if o != nil && o.CreatorUserId.IsSet() {
		return true
	}

	return false
}

// SetCreatorUserId gets a reference to the given NullableString and assigns it to the CreatorUserId field.
func (o *ContactVM) SetCreatorUserId(v string) {
	o.CreatorUserId.Set(&v)
}
// SetCreatorUserIdNil sets the value for CreatorUserId to be an explicit nil
func (o *ContactVM) SetCreatorUserIdNil() {
	o.CreatorUserId.Set(nil)
}

// UnsetCreatorUserId ensures that no value is present for CreatorUserId, not even an explicit nil
func (o *ContactVM) UnsetCreatorUserId() {
	o.CreatorUserId.Unset()
}

// GetEditedTime returns the EditedTime field value if set, zero value otherwise.
func (o *ContactVM) GetEditedTime() time.Time {
	if o == nil || IsNil(o.EditedTime) {
		var ret time.Time
		return ret
	}
	return *o.EditedTime
}

// GetEditedTimeOk returns a tuple with the EditedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactVM) GetEditedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EditedTime) {
		return nil, false
	}
	return o.EditedTime, true
}

// HasEditedTime returns a boolean if a field has been set.
func (o *ContactVM) HasEditedTime() bool {
	if o != nil && !IsNil(o.EditedTime) {
		return true
	}

	return false
}

// SetEditedTime gets a reference to the given time.Time and assigns it to the EditedTime field.
func (o *ContactVM) SetEditedTime(v time.Time) {
	o.EditedTime = &v
}

// GetEditorUserId returns the EditorUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactVM) GetEditorUserId() string {
	if o == nil || IsNil(o.EditorUserId.Get()) {
		var ret string
		return ret
	}
	return *o.EditorUserId.Get()
}

// GetEditorUserIdOk returns a tuple with the EditorUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactVM) GetEditorUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EditorUserId.Get(), o.EditorUserId.IsSet()
}

// HasEditorUserId returns a boolean if a field has been set.
func (o *ContactVM) HasEditorUserId() bool {
	if o != nil && o.EditorUserId.IsSet() {
		return true
	}

	return false
}

// SetEditorUserId gets a reference to the given NullableString and assigns it to the EditorUserId field.
func (o *ContactVM) SetEditorUserId(v string) {
	o.EditorUserId.Set(&v)
}
// SetEditorUserIdNil sets the value for EditorUserId to be an explicit nil
func (o *ContactVM) SetEditorUserIdNil() {
	o.EditorUserId.Set(nil)
}

// UnsetEditorUserId ensures that no value is present for EditorUserId, not even an explicit nil
func (o *ContactVM) UnsetEditorUserId() {
	o.EditorUserId.Unset()
}

func (o ContactVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscriptionId"] = o.SubscriptionId.Get()
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if o.CreatorUserId.IsSet() {
		toSerialize["creatorUserId"] = o.CreatorUserId.Get()
	}
	if !IsNil(o.EditedTime) {
		toSerialize["editedTime"] = o.EditedTime
	}
	if o.EditorUserId.IsSet() {
		toSerialize["editorUserId"] = o.EditorUserId.Get()
	}
	return toSerialize, nil
}

type NullableContactVM struct {
	value *ContactVM
	isSet bool
}

func (v NullableContactVM) Get() *ContactVM {
	return v.value
}

func (v *NullableContactVM) Set(val *ContactVM) {
	v.value = val
	v.isSet = true
}

func (v NullableContactVM) IsSet() bool {
	return v.isSet
}

func (v *NullableContactVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactVM(val *ContactVM) *NullableContactVM {
	return &NullableContactVM{value: val, isSet: true}
}

func (v NullableContactVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

