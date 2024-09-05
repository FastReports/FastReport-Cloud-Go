# CreateFTPUploadTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archive** | Pointer to **bool** |  | [optional] 
**ArchiveName** | Pointer to **NullableString** |  | [optional] 
**DestinationFolder** | Pointer to **NullableString** |  | [optional] 
**FtpHost** | Pointer to **NullableString** |  | [optional] 
**FtpPassword** | Pointer to **NullableString** |  | [optional] 
**FtpPort** | Pointer to **int32** |  | [optional] 
**FtpUsername** | Pointer to **NullableString** |  | [optional] 
**UseSFTP** | Pointer to **bool** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewCreateFTPUploadTaskVM

`func NewCreateFTPUploadTaskVM(t string, ) *CreateFTPUploadTaskVM`

NewCreateFTPUploadTaskVM instantiates a new CreateFTPUploadTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFTPUploadTaskVMWithDefaults

`func NewCreateFTPUploadTaskVMWithDefaults() *CreateFTPUploadTaskVM`

NewCreateFTPUploadTaskVMWithDefaults instantiates a new CreateFTPUploadTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchive

`func (o *CreateFTPUploadTaskVM) GetArchive() bool`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *CreateFTPUploadTaskVM) GetArchiveOk() (*bool, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *CreateFTPUploadTaskVM) SetArchive(v bool)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *CreateFTPUploadTaskVM) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetArchiveName

`func (o *CreateFTPUploadTaskVM) GetArchiveName() string`

GetArchiveName returns the ArchiveName field if non-nil, zero value otherwise.

### GetArchiveNameOk

`func (o *CreateFTPUploadTaskVM) GetArchiveNameOk() (*string, bool)`

GetArchiveNameOk returns a tuple with the ArchiveName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveName

`func (o *CreateFTPUploadTaskVM) SetArchiveName(v string)`

SetArchiveName sets ArchiveName field to given value.

### HasArchiveName

`func (o *CreateFTPUploadTaskVM) HasArchiveName() bool`

HasArchiveName returns a boolean if a field has been set.

### SetArchiveNameNil

`func (o *CreateFTPUploadTaskVM) SetArchiveNameNil(b bool)`

 SetArchiveNameNil sets the value for ArchiveName to be an explicit nil

### UnsetArchiveName
`func (o *CreateFTPUploadTaskVM) UnsetArchiveName()`

UnsetArchiveName ensures that no value is present for ArchiveName, not even an explicit nil
### GetDestinationFolder

`func (o *CreateFTPUploadTaskVM) GetDestinationFolder() string`

GetDestinationFolder returns the DestinationFolder field if non-nil, zero value otherwise.

### GetDestinationFolderOk

`func (o *CreateFTPUploadTaskVM) GetDestinationFolderOk() (*string, bool)`

GetDestinationFolderOk returns a tuple with the DestinationFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationFolder

`func (o *CreateFTPUploadTaskVM) SetDestinationFolder(v string)`

SetDestinationFolder sets DestinationFolder field to given value.

### HasDestinationFolder

`func (o *CreateFTPUploadTaskVM) HasDestinationFolder() bool`

HasDestinationFolder returns a boolean if a field has been set.

### SetDestinationFolderNil

`func (o *CreateFTPUploadTaskVM) SetDestinationFolderNil(b bool)`

 SetDestinationFolderNil sets the value for DestinationFolder to be an explicit nil

### UnsetDestinationFolder
`func (o *CreateFTPUploadTaskVM) UnsetDestinationFolder()`

UnsetDestinationFolder ensures that no value is present for DestinationFolder, not even an explicit nil
### GetFtpHost

`func (o *CreateFTPUploadTaskVM) GetFtpHost() string`

GetFtpHost returns the FtpHost field if non-nil, zero value otherwise.

### GetFtpHostOk

`func (o *CreateFTPUploadTaskVM) GetFtpHostOk() (*string, bool)`

GetFtpHostOk returns a tuple with the FtpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpHost

`func (o *CreateFTPUploadTaskVM) SetFtpHost(v string)`

SetFtpHost sets FtpHost field to given value.

### HasFtpHost

`func (o *CreateFTPUploadTaskVM) HasFtpHost() bool`

HasFtpHost returns a boolean if a field has been set.

### SetFtpHostNil

`func (o *CreateFTPUploadTaskVM) SetFtpHostNil(b bool)`

 SetFtpHostNil sets the value for FtpHost to be an explicit nil

### UnsetFtpHost
`func (o *CreateFTPUploadTaskVM) UnsetFtpHost()`

UnsetFtpHost ensures that no value is present for FtpHost, not even an explicit nil
### GetFtpPassword

`func (o *CreateFTPUploadTaskVM) GetFtpPassword() string`

GetFtpPassword returns the FtpPassword field if non-nil, zero value otherwise.

### GetFtpPasswordOk

`func (o *CreateFTPUploadTaskVM) GetFtpPasswordOk() (*string, bool)`

GetFtpPasswordOk returns a tuple with the FtpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpPassword

`func (o *CreateFTPUploadTaskVM) SetFtpPassword(v string)`

SetFtpPassword sets FtpPassword field to given value.

### HasFtpPassword

`func (o *CreateFTPUploadTaskVM) HasFtpPassword() bool`

HasFtpPassword returns a boolean if a field has been set.

### SetFtpPasswordNil

`func (o *CreateFTPUploadTaskVM) SetFtpPasswordNil(b bool)`

 SetFtpPasswordNil sets the value for FtpPassword to be an explicit nil

### UnsetFtpPassword
`func (o *CreateFTPUploadTaskVM) UnsetFtpPassword()`

UnsetFtpPassword ensures that no value is present for FtpPassword, not even an explicit nil
### GetFtpPort

`func (o *CreateFTPUploadTaskVM) GetFtpPort() int32`

GetFtpPort returns the FtpPort field if non-nil, zero value otherwise.

### GetFtpPortOk

`func (o *CreateFTPUploadTaskVM) GetFtpPortOk() (*int32, bool)`

GetFtpPortOk returns a tuple with the FtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpPort

`func (o *CreateFTPUploadTaskVM) SetFtpPort(v int32)`

SetFtpPort sets FtpPort field to given value.

### HasFtpPort

`func (o *CreateFTPUploadTaskVM) HasFtpPort() bool`

HasFtpPort returns a boolean if a field has been set.

### GetFtpUsername

`func (o *CreateFTPUploadTaskVM) GetFtpUsername() string`

GetFtpUsername returns the FtpUsername field if non-nil, zero value otherwise.

### GetFtpUsernameOk

`func (o *CreateFTPUploadTaskVM) GetFtpUsernameOk() (*string, bool)`

GetFtpUsernameOk returns a tuple with the FtpUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpUsername

`func (o *CreateFTPUploadTaskVM) SetFtpUsername(v string)`

SetFtpUsername sets FtpUsername field to given value.

### HasFtpUsername

`func (o *CreateFTPUploadTaskVM) HasFtpUsername() bool`

HasFtpUsername returns a boolean if a field has been set.

### SetFtpUsernameNil

`func (o *CreateFTPUploadTaskVM) SetFtpUsernameNil(b bool)`

 SetFtpUsernameNil sets the value for FtpUsername to be an explicit nil

### UnsetFtpUsername
`func (o *CreateFTPUploadTaskVM) UnsetFtpUsername()`

UnsetFtpUsername ensures that no value is present for FtpUsername, not even an explicit nil
### GetUseSFTP

`func (o *CreateFTPUploadTaskVM) GetUseSFTP() bool`

GetUseSFTP returns the UseSFTP field if non-nil, zero value otherwise.

### GetUseSFTPOk

`func (o *CreateFTPUploadTaskVM) GetUseSFTPOk() (*bool, bool)`

GetUseSFTPOk returns a tuple with the UseSFTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSFTP

`func (o *CreateFTPUploadTaskVM) SetUseSFTP(v bool)`

SetUseSFTP sets UseSFTP field to given value.

### HasUseSFTP

`func (o *CreateFTPUploadTaskVM) HasUseSFTP() bool`

HasUseSFTP returns a boolean if a field has been set.

### GetT

`func (o *CreateFTPUploadTaskVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *CreateFTPUploadTaskVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *CreateFTPUploadTaskVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


