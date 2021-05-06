/*
 * FastReport Cloud
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package FastReport.Cloud.SDK

import (
	"encoding/json"
	"time"
)

// FileVM struct for FileVM
type FileVM struct {
	Name *string `json:"name,omitempty"`
	ParentId *string `json:"parentId,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Type *string `json:"type,omitempty"`
	Size *int64 `json:"size,omitempty"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	Status *string `json:"status,omitempty"`
	StatusReason *string `json:"statusReason,omitempty"`
	Id *string `json:"id,omitempty"`
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	CreatorUserId *string `json:"creatorUserId,omitempty"`
	EditedTime *time.Time `json:"editedTime,omitempty"`
	EditorUserId *string `json:"editorUserId,omitempty"`
}

// NewFileVM instantiates a new FileVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileVM() *FileVM {
	this := FileVM{}
	return &this
}

// NewFileVMWithDefaults instantiates a new FileVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileVMWithDefaults() *FileVM {
	this := FileVM{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FileVM) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVM) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FileVM) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FileVM) SetName(v string) {
	o.Name = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *FileVM) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVM) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *FileVM) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *FileVM) SetParentId(v string) {
	o.ParentId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FileVM) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVM) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FileVM) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FileVM) SetTags(v []string) {
	o.Tags = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *FileVM) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVM) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *FileVM) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *FileVM) SetIcon(v string) {
	o.Icon = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FileVM) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVM) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FileVM) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FileVM) SetType(v string) {
	o.Type = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *FileVM) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVM) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *FileVM) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *FileVM) SetSize(v int64) {
	o.Size = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *FileVM) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *FileVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *FileVM) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FileVM) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVM) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FileVM) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FileVM) SetStatus(v string) {
	o.Status = &v
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise.
func (o *FileVM) GetStatusReason() string {
	if o == nil || o.StatusReason == nil {
		var ret string
		return ret
	}
	return *o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVM) GetStatusReasonOk() (*string, bool) {
	if o == nil || o.StatusReason == nil {
		return nil, false
	}
	return o.StatusReason, true
}

// HasStatusReason returns a boolean if a field has been set.
func (o *FileVM) HasStatusReason() bool {
	if o != nil && o.StatusReason != nil {
		return true
	}

	return false
}

// SetStatusReason gets a reference to the given string and assigns it to the StatusReason field.
func (o *FileVM) SetStatusReason(v string) {
	o.StatusReason = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FileVM) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVM) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FileVM) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FileVM) SetId(v string) {
	o.Id = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *FileVM) GetCreatedTime() time.Time {
	if o == nil || o.CreatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVM) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *FileVM) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *FileVM) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetCreatorUserId returns the CreatorUserId field value if set, zero value otherwise.
func (o *FileVM) GetCreatorUserId() string {
	if o == nil || o.CreatorUserId == nil {
		var ret string
		return ret
	}
	return *o.CreatorUserId
}

// GetCreatorUserIdOk returns a tuple with the CreatorUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVM) GetCreatorUserIdOk() (*string, bool) {
	if o == nil || o.CreatorUserId == nil {
		return nil, false
	}
	return o.CreatorUserId, true
}

// HasCreatorUserId returns a boolean if a field has been set.
func (o *FileVM) HasCreatorUserId() bool {
	if o != nil && o.CreatorUserId != nil {
		return true
	}

	return false
}

// SetCreatorUserId gets a reference to the given string and assigns it to the CreatorUserId field.
func (o *FileVM) SetCreatorUserId(v string) {
	o.CreatorUserId = &v
}

// GetEditedTime returns the EditedTime field value if set, zero value otherwise.
func (o *FileVM) GetEditedTime() time.Time {
	if o == nil || o.EditedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EditedTime
}

// GetEditedTimeOk returns a tuple with the EditedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVM) GetEditedTimeOk() (*time.Time, bool) {
	if o == nil || o.EditedTime == nil {
		return nil, false
	}
	return o.EditedTime, true
}

// HasEditedTime returns a boolean if a field has been set.
func (o *FileVM) HasEditedTime() bool {
	if o != nil && o.EditedTime != nil {
		return true
	}

	return false
}

// SetEditedTime gets a reference to the given time.Time and assigns it to the EditedTime field.
func (o *FileVM) SetEditedTime(v time.Time) {
	o.EditedTime = &v
}

// GetEditorUserId returns the EditorUserId field value if set, zero value otherwise.
func (o *FileVM) GetEditorUserId() string {
	if o == nil || o.EditorUserId == nil {
		var ret string
		return ret
	}
	return *o.EditorUserId
}

// GetEditorUserIdOk returns a tuple with the EditorUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileVM) GetEditorUserIdOk() (*string, bool) {
	if o == nil || o.EditorUserId == nil {
		return nil, false
	}
	return o.EditorUserId, true
}

// HasEditorUserId returns a boolean if a field has been set.
func (o *FileVM) HasEditorUserId() bool {
	if o != nil && o.EditorUserId != nil {
		return true
	}

	return false
}

// SetEditorUserId gets a reference to the given string and assigns it to the EditorUserId field.
func (o *FileVM) SetEditorUserId(v string) {
	o.EditorUserId = &v
}

func (o FileVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.SubscriptionId != nil {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusReason != nil {
		toSerialize["statusReason"] = o.StatusReason
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedTime != nil {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if o.CreatorUserId != nil {
		toSerialize["creatorUserId"] = o.CreatorUserId
	}
	if o.EditedTime != nil {
		toSerialize["editedTime"] = o.EditedTime
	}
	if o.EditorUserId != nil {
		toSerialize["editorUserId"] = o.EditorUserId
	}
	return json.Marshal(toSerialize)
}

type NullableFileVM struct {
	value *FileVM
	isSet bool
}

func (v NullableFileVM) Get() *FileVM {
	return v.value
}

func (v *NullableFileVM) Set(val *FileVM) {
	v.value = val
	v.isSet = true
}

func (v NullableFileVM) IsSet() bool {
	return v.isSet
}

func (v *NullableFileVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileVM(val *FileVM) *NullableFileVM {
	return &NullableFileVM{value: val, isSet: true}
}

func (v NullableFileVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


