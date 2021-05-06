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
)

// SubscriptionFolder struct for SubscriptionFolder
type SubscriptionFolder struct {
	FolderId *string `json:"folderId,omitempty"`
	BytesUsed *int64 `json:"bytesUsed,omitempty"`
}

// NewSubscriptionFolder instantiates a new SubscriptionFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionFolder() *SubscriptionFolder {
	this := SubscriptionFolder{}
	return &this
}

// NewSubscriptionFolderWithDefaults instantiates a new SubscriptionFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionFolderWithDefaults() *SubscriptionFolder {
	this := SubscriptionFolder{}
	return &this
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *SubscriptionFolder) GetFolderId() string {
	if o == nil || o.FolderId == nil {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionFolder) GetFolderIdOk() (*string, bool) {
	if o == nil || o.FolderId == nil {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *SubscriptionFolder) HasFolderId() bool {
	if o != nil && o.FolderId != nil {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *SubscriptionFolder) SetFolderId(v string) {
	o.FolderId = &v
}

// GetBytesUsed returns the BytesUsed field value if set, zero value otherwise.
func (o *SubscriptionFolder) GetBytesUsed() int64 {
	if o == nil || o.BytesUsed == nil {
		var ret int64
		return ret
	}
	return *o.BytesUsed
}

// GetBytesUsedOk returns a tuple with the BytesUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionFolder) GetBytesUsedOk() (*int64, bool) {
	if o == nil || o.BytesUsed == nil {
		return nil, false
	}
	return o.BytesUsed, true
}

// HasBytesUsed returns a boolean if a field has been set.
func (o *SubscriptionFolder) HasBytesUsed() bool {
	if o != nil && o.BytesUsed != nil {
		return true
	}

	return false
}

// SetBytesUsed gets a reference to the given int64 and assigns it to the BytesUsed field.
func (o *SubscriptionFolder) SetBytesUsed(v int64) {
	o.BytesUsed = &v
}

func (o SubscriptionFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FolderId != nil {
		toSerialize["folderId"] = o.FolderId
	}
	if o.BytesUsed != nil {
		toSerialize["bytesUsed"] = o.BytesUsed
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionFolder struct {
	value *SubscriptionFolder
	isSet bool
}

func (v NullableSubscriptionFolder) Get() *SubscriptionFolder {
	return v.value
}

func (v *NullableSubscriptionFolder) Set(val *SubscriptionFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionFolder(val *SubscriptionFolder) *NullableSubscriptionFolder {
	return &NullableSubscriptionFolder{value: val, isSet: true}
}

func (v NullableSubscriptionFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


