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

// ReportInfo struct for ReportInfo
type ReportInfo struct {
	Author *string `json:"author,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	CreatorVersion *string `json:"creatorVersion,omitempty"`
	Description *string `json:"description,omitempty"`
	Modified *time.Time `json:"modified,omitempty"`
	Name *string `json:"name,omitempty"`
	Picture *string `json:"picture,omitempty"`
	PreviewPictureRatio *float32 `json:"previewPictureRatio,omitempty"`
	SaveMode *string `json:"saveMode,omitempty"`
	SavePreviewPicture *bool `json:"savePreviewPicture,omitempty"`
	Tag *string `json:"tag,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewReportInfo instantiates a new ReportInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportInfo() *ReportInfo {
	this := ReportInfo{}
	return &this
}

// NewReportInfoWithDefaults instantiates a new ReportInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportInfoWithDefaults() *ReportInfo {
	this := ReportInfo{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *ReportInfo) GetAuthor() string {
	if o == nil || o.Author == nil {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportInfo) GetAuthorOk() (*string, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *ReportInfo) HasAuthor() bool {
	if o != nil && o.Author != nil {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *ReportInfo) SetAuthor(v string) {
	o.Author = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ReportInfo) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportInfo) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ReportInfo) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ReportInfo) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatorVersion returns the CreatorVersion field value if set, zero value otherwise.
func (o *ReportInfo) GetCreatorVersion() string {
	if o == nil || o.CreatorVersion == nil {
		var ret string
		return ret
	}
	return *o.CreatorVersion
}

// GetCreatorVersionOk returns a tuple with the CreatorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportInfo) GetCreatorVersionOk() (*string, bool) {
	if o == nil || o.CreatorVersion == nil {
		return nil, false
	}
	return o.CreatorVersion, true
}

// HasCreatorVersion returns a boolean if a field has been set.
func (o *ReportInfo) HasCreatorVersion() bool {
	if o != nil && o.CreatorVersion != nil {
		return true
	}

	return false
}

// SetCreatorVersion gets a reference to the given string and assigns it to the CreatorVersion field.
func (o *ReportInfo) SetCreatorVersion(v string) {
	o.CreatorVersion = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReportInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReportInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReportInfo) SetDescription(v string) {
	o.Description = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *ReportInfo) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportInfo) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *ReportInfo) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *ReportInfo) SetModified(v time.Time) {
	o.Modified = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReportInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReportInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReportInfo) SetName(v string) {
	o.Name = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *ReportInfo) GetPicture() string {
	if o == nil || o.Picture == nil {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportInfo) GetPictureOk() (*string, bool) {
	if o == nil || o.Picture == nil {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *ReportInfo) HasPicture() bool {
	if o != nil && o.Picture != nil {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *ReportInfo) SetPicture(v string) {
	o.Picture = &v
}

// GetPreviewPictureRatio returns the PreviewPictureRatio field value if set, zero value otherwise.
func (o *ReportInfo) GetPreviewPictureRatio() float32 {
	if o == nil || o.PreviewPictureRatio == nil {
		var ret float32
		return ret
	}
	return *o.PreviewPictureRatio
}

// GetPreviewPictureRatioOk returns a tuple with the PreviewPictureRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportInfo) GetPreviewPictureRatioOk() (*float32, bool) {
	if o == nil || o.PreviewPictureRatio == nil {
		return nil, false
	}
	return o.PreviewPictureRatio, true
}

// HasPreviewPictureRatio returns a boolean if a field has been set.
func (o *ReportInfo) HasPreviewPictureRatio() bool {
	if o != nil && o.PreviewPictureRatio != nil {
		return true
	}

	return false
}

// SetPreviewPictureRatio gets a reference to the given float32 and assigns it to the PreviewPictureRatio field.
func (o *ReportInfo) SetPreviewPictureRatio(v float32) {
	o.PreviewPictureRatio = &v
}

// GetSaveMode returns the SaveMode field value if set, zero value otherwise.
func (o *ReportInfo) GetSaveMode() string {
	if o == nil || o.SaveMode == nil {
		var ret string
		return ret
	}
	return *o.SaveMode
}

// GetSaveModeOk returns a tuple with the SaveMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportInfo) GetSaveModeOk() (*string, bool) {
	if o == nil || o.SaveMode == nil {
		return nil, false
	}
	return o.SaveMode, true
}

// HasSaveMode returns a boolean if a field has been set.
func (o *ReportInfo) HasSaveMode() bool {
	if o != nil && o.SaveMode != nil {
		return true
	}

	return false
}

// SetSaveMode gets a reference to the given string and assigns it to the SaveMode field.
func (o *ReportInfo) SetSaveMode(v string) {
	o.SaveMode = &v
}

// GetSavePreviewPicture returns the SavePreviewPicture field value if set, zero value otherwise.
func (o *ReportInfo) GetSavePreviewPicture() bool {
	if o == nil || o.SavePreviewPicture == nil {
		var ret bool
		return ret
	}
	return *o.SavePreviewPicture
}

// GetSavePreviewPictureOk returns a tuple with the SavePreviewPicture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportInfo) GetSavePreviewPictureOk() (*bool, bool) {
	if o == nil || o.SavePreviewPicture == nil {
		return nil, false
	}
	return o.SavePreviewPicture, true
}

// HasSavePreviewPicture returns a boolean if a field has been set.
func (o *ReportInfo) HasSavePreviewPicture() bool {
	if o != nil && o.SavePreviewPicture != nil {
		return true
	}

	return false
}

// SetSavePreviewPicture gets a reference to the given bool and assigns it to the SavePreviewPicture field.
func (o *ReportInfo) SetSavePreviewPicture(v bool) {
	o.SavePreviewPicture = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *ReportInfo) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportInfo) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *ReportInfo) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *ReportInfo) SetTag(v string) {
	o.Tag = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ReportInfo) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportInfo) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ReportInfo) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ReportInfo) SetVersion(v string) {
	o.Version = &v
}

func (o ReportInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.CreatorVersion != nil {
		toSerialize["creatorVersion"] = o.CreatorVersion
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Picture != nil {
		toSerialize["picture"] = o.Picture
	}
	if o.PreviewPictureRatio != nil {
		toSerialize["previewPictureRatio"] = o.PreviewPictureRatio
	}
	if o.SaveMode != nil {
		toSerialize["saveMode"] = o.SaveMode
	}
	if o.SavePreviewPicture != nil {
		toSerialize["savePreviewPicture"] = o.SavePreviewPicture
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableReportInfo struct {
	value *ReportInfo
	isSet bool
}

func (v NullableReportInfo) Get() *ReportInfo {
	return v.value
}

func (v *NullableReportInfo) Set(val *ReportInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableReportInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableReportInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportInfo(val *ReportInfo) *NullableReportInfo {
	return &NullableReportInfo{value: val, isSet: true}
}

func (v NullableReportInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

