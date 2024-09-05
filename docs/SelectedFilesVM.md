# SelectedFilesVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAllSelected** | Pointer to **bool** |  | [optional] 
**FolderId** | Pointer to **NullableString** |  | [optional] 
**SearchPattern** | Pointer to **NullableString** |  | [optional] 
**UseRegex** | Pointer to **bool** |  | [optional] 
**Files** | Pointer to **[]string** |  | [optional] 
**Folders** | Pointer to **[]string** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**IsBin** | Pointer to **bool** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewSelectedFilesVM

`func NewSelectedFilesVM(t string, ) *SelectedFilesVM`

NewSelectedFilesVM instantiates a new SelectedFilesVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectedFilesVMWithDefaults

`func NewSelectedFilesVMWithDefaults() *SelectedFilesVM`

NewSelectedFilesVMWithDefaults instantiates a new SelectedFilesVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAllSelected

`func (o *SelectedFilesVM) GetIsAllSelected() bool`

GetIsAllSelected returns the IsAllSelected field if non-nil, zero value otherwise.

### GetIsAllSelectedOk

`func (o *SelectedFilesVM) GetIsAllSelectedOk() (*bool, bool)`

GetIsAllSelectedOk returns a tuple with the IsAllSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllSelected

`func (o *SelectedFilesVM) SetIsAllSelected(v bool)`

SetIsAllSelected sets IsAllSelected field to given value.

### HasIsAllSelected

`func (o *SelectedFilesVM) HasIsAllSelected() bool`

HasIsAllSelected returns a boolean if a field has been set.

### GetFolderId

`func (o *SelectedFilesVM) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *SelectedFilesVM) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *SelectedFilesVM) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *SelectedFilesVM) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *SelectedFilesVM) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *SelectedFilesVM) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetSearchPattern

`func (o *SelectedFilesVM) GetSearchPattern() string`

GetSearchPattern returns the SearchPattern field if non-nil, zero value otherwise.

### GetSearchPatternOk

`func (o *SelectedFilesVM) GetSearchPatternOk() (*string, bool)`

GetSearchPatternOk returns a tuple with the SearchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPattern

`func (o *SelectedFilesVM) SetSearchPattern(v string)`

SetSearchPattern sets SearchPattern field to given value.

### HasSearchPattern

`func (o *SelectedFilesVM) HasSearchPattern() bool`

HasSearchPattern returns a boolean if a field has been set.

### SetSearchPatternNil

`func (o *SelectedFilesVM) SetSearchPatternNil(b bool)`

 SetSearchPatternNil sets the value for SearchPattern to be an explicit nil

### UnsetSearchPattern
`func (o *SelectedFilesVM) UnsetSearchPattern()`

UnsetSearchPattern ensures that no value is present for SearchPattern, not even an explicit nil
### GetUseRegex

`func (o *SelectedFilesVM) GetUseRegex() bool`

GetUseRegex returns the UseRegex field if non-nil, zero value otherwise.

### GetUseRegexOk

`func (o *SelectedFilesVM) GetUseRegexOk() (*bool, bool)`

GetUseRegexOk returns a tuple with the UseRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRegex

`func (o *SelectedFilesVM) SetUseRegex(v bool)`

SetUseRegex sets UseRegex field to given value.

### HasUseRegex

`func (o *SelectedFilesVM) HasUseRegex() bool`

HasUseRegex returns a boolean if a field has been set.

### GetFiles

`func (o *SelectedFilesVM) GetFiles() []string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SelectedFilesVM) GetFilesOk() (*[]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SelectedFilesVM) SetFiles(v []string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *SelectedFilesVM) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *SelectedFilesVM) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *SelectedFilesVM) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetFolders

`func (o *SelectedFilesVM) GetFolders() []string`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *SelectedFilesVM) GetFoldersOk() (*[]string, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *SelectedFilesVM) SetFolders(v []string)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *SelectedFilesVM) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### SetFoldersNil

`func (o *SelectedFilesVM) SetFoldersNil(b bool)`

 SetFoldersNil sets the value for Folders to be an explicit nil

### UnsetFolders
`func (o *SelectedFilesVM) UnsetFolders()`

UnsetFolders ensures that no value is present for Folders, not even an explicit nil
### GetPath

`func (o *SelectedFilesVM) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SelectedFilesVM) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SelectedFilesVM) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SelectedFilesVM) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *SelectedFilesVM) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *SelectedFilesVM) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetIsBin

`func (o *SelectedFilesVM) GetIsBin() bool`

GetIsBin returns the IsBin field if non-nil, zero value otherwise.

### GetIsBinOk

`func (o *SelectedFilesVM) GetIsBinOk() (*bool, bool)`

GetIsBinOk returns a tuple with the IsBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBin

`func (o *SelectedFilesVM) SetIsBin(v bool)`

SetIsBin sets IsBin field to given value.

### HasIsBin

`func (o *SelectedFilesVM) HasIsBin() bool`

HasIsBin returns a boolean if a field has been set.

### GetT

`func (o *SelectedFilesVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *SelectedFilesVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *SelectedFilesVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


