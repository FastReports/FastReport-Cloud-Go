# ReportInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**CreatorVersion** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Picture** | Pointer to **string** |  | [optional] 
**PreviewPictureRatio** | Pointer to **float32** |  | [optional] 
**SaveMode** | Pointer to **string** |  | [optional] 
**SavePreviewPicture** | Pointer to **bool** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewReportInfo

`func NewReportInfo() *ReportInfo`

NewReportInfo instantiates a new ReportInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportInfoWithDefaults

`func NewReportInfoWithDefaults() *ReportInfo`

NewReportInfoWithDefaults instantiates a new ReportInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *ReportInfo) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ReportInfo) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ReportInfo) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ReportInfo) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCreated

`func (o *ReportInfo) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ReportInfo) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ReportInfo) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ReportInfo) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatorVersion

`func (o *ReportInfo) GetCreatorVersion() string`

GetCreatorVersion returns the CreatorVersion field if non-nil, zero value otherwise.

### GetCreatorVersionOk

`func (o *ReportInfo) GetCreatorVersionOk() (*string, bool)`

GetCreatorVersionOk returns a tuple with the CreatorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorVersion

`func (o *ReportInfo) SetCreatorVersion(v string)`

SetCreatorVersion sets CreatorVersion field to given value.

### HasCreatorVersion

`func (o *ReportInfo) HasCreatorVersion() bool`

HasCreatorVersion returns a boolean if a field has been set.

### GetDescription

`func (o *ReportInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReportInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReportInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReportInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModified

`func (o *ReportInfo) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ReportInfo) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ReportInfo) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ReportInfo) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *ReportInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReportInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPicture

`func (o *ReportInfo) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *ReportInfo) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *ReportInfo) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *ReportInfo) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetPreviewPictureRatio

`func (o *ReportInfo) GetPreviewPictureRatio() float32`

GetPreviewPictureRatio returns the PreviewPictureRatio field if non-nil, zero value otherwise.

### GetPreviewPictureRatioOk

`func (o *ReportInfo) GetPreviewPictureRatioOk() (*float32, bool)`

GetPreviewPictureRatioOk returns a tuple with the PreviewPictureRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewPictureRatio

`func (o *ReportInfo) SetPreviewPictureRatio(v float32)`

SetPreviewPictureRatio sets PreviewPictureRatio field to given value.

### HasPreviewPictureRatio

`func (o *ReportInfo) HasPreviewPictureRatio() bool`

HasPreviewPictureRatio returns a boolean if a field has been set.

### GetSaveMode

`func (o *ReportInfo) GetSaveMode() string`

GetSaveMode returns the SaveMode field if non-nil, zero value otherwise.

### GetSaveModeOk

`func (o *ReportInfo) GetSaveModeOk() (*string, bool)`

GetSaveModeOk returns a tuple with the SaveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveMode

`func (o *ReportInfo) SetSaveMode(v string)`

SetSaveMode sets SaveMode field to given value.

### HasSaveMode

`func (o *ReportInfo) HasSaveMode() bool`

HasSaveMode returns a boolean if a field has been set.

### GetSavePreviewPicture

`func (o *ReportInfo) GetSavePreviewPicture() bool`

GetSavePreviewPicture returns the SavePreviewPicture field if non-nil, zero value otherwise.

### GetSavePreviewPictureOk

`func (o *ReportInfo) GetSavePreviewPictureOk() (*bool, bool)`

GetSavePreviewPictureOk returns a tuple with the SavePreviewPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavePreviewPicture

`func (o *ReportInfo) SetSavePreviewPicture(v bool)`

SetSavePreviewPicture sets SavePreviewPicture field to given value.

### HasSavePreviewPicture

`func (o *ReportInfo) HasSavePreviewPicture() bool`

HasSavePreviewPicture returns a boolean if a field has been set.

### GetTag

`func (o *ReportInfo) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ReportInfo) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ReportInfo) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ReportInfo) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetVersion

`func (o *ReportInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ReportInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ReportInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ReportInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


