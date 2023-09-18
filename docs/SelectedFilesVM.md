# SelectedFilesVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAllSelected** | Pointer to **bool** |  | [optional] 
**Files** | Pointer to **[]string** |  | [optional] 
**Folders** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSelectedFilesVM

`func NewSelectedFilesVM() *SelectedFilesVM`

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


