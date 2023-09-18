# RunFTPUploadTaskVM

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

## Methods

### NewRunFTPUploadTaskVM

`func NewRunFTPUploadTaskVM() *RunFTPUploadTaskVM`

NewRunFTPUploadTaskVM instantiates a new RunFTPUploadTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunFTPUploadTaskVMWithDefaults

`func NewRunFTPUploadTaskVMWithDefaults() *RunFTPUploadTaskVM`

NewRunFTPUploadTaskVMWithDefaults instantiates a new RunFTPUploadTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchive

`func (o *RunFTPUploadTaskVM) GetArchive() bool`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *RunFTPUploadTaskVM) GetArchiveOk() (*bool, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *RunFTPUploadTaskVM) SetArchive(v bool)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *RunFTPUploadTaskVM) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetArchiveName

`func (o *RunFTPUploadTaskVM) GetArchiveName() string`

GetArchiveName returns the ArchiveName field if non-nil, zero value otherwise.

### GetArchiveNameOk

`func (o *RunFTPUploadTaskVM) GetArchiveNameOk() (*string, bool)`

GetArchiveNameOk returns a tuple with the ArchiveName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveName

`func (o *RunFTPUploadTaskVM) SetArchiveName(v string)`

SetArchiveName sets ArchiveName field to given value.

### HasArchiveName

`func (o *RunFTPUploadTaskVM) HasArchiveName() bool`

HasArchiveName returns a boolean if a field has been set.

### SetArchiveNameNil

`func (o *RunFTPUploadTaskVM) SetArchiveNameNil(b bool)`

 SetArchiveNameNil sets the value for ArchiveName to be an explicit nil

### UnsetArchiveName
`func (o *RunFTPUploadTaskVM) UnsetArchiveName()`

UnsetArchiveName ensures that no value is present for ArchiveName, not even an explicit nil
### GetDestinationFolder

`func (o *RunFTPUploadTaskVM) GetDestinationFolder() string`

GetDestinationFolder returns the DestinationFolder field if non-nil, zero value otherwise.

### GetDestinationFolderOk

`func (o *RunFTPUploadTaskVM) GetDestinationFolderOk() (*string, bool)`

GetDestinationFolderOk returns a tuple with the DestinationFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationFolder

`func (o *RunFTPUploadTaskVM) SetDestinationFolder(v string)`

SetDestinationFolder sets DestinationFolder field to given value.

### HasDestinationFolder

`func (o *RunFTPUploadTaskVM) HasDestinationFolder() bool`

HasDestinationFolder returns a boolean if a field has been set.

### SetDestinationFolderNil

`func (o *RunFTPUploadTaskVM) SetDestinationFolderNil(b bool)`

 SetDestinationFolderNil sets the value for DestinationFolder to be an explicit nil

### UnsetDestinationFolder
`func (o *RunFTPUploadTaskVM) UnsetDestinationFolder()`

UnsetDestinationFolder ensures that no value is present for DestinationFolder, not even an explicit nil
### GetFtpHost

`func (o *RunFTPUploadTaskVM) GetFtpHost() string`

GetFtpHost returns the FtpHost field if non-nil, zero value otherwise.

### GetFtpHostOk

`func (o *RunFTPUploadTaskVM) GetFtpHostOk() (*string, bool)`

GetFtpHostOk returns a tuple with the FtpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpHost

`func (o *RunFTPUploadTaskVM) SetFtpHost(v string)`

SetFtpHost sets FtpHost field to given value.

### HasFtpHost

`func (o *RunFTPUploadTaskVM) HasFtpHost() bool`

HasFtpHost returns a boolean if a field has been set.

### SetFtpHostNil

`func (o *RunFTPUploadTaskVM) SetFtpHostNil(b bool)`

 SetFtpHostNil sets the value for FtpHost to be an explicit nil

### UnsetFtpHost
`func (o *RunFTPUploadTaskVM) UnsetFtpHost()`

UnsetFtpHost ensures that no value is present for FtpHost, not even an explicit nil
### GetFtpPassword

`func (o *RunFTPUploadTaskVM) GetFtpPassword() string`

GetFtpPassword returns the FtpPassword field if non-nil, zero value otherwise.

### GetFtpPasswordOk

`func (o *RunFTPUploadTaskVM) GetFtpPasswordOk() (*string, bool)`

GetFtpPasswordOk returns a tuple with the FtpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpPassword

`func (o *RunFTPUploadTaskVM) SetFtpPassword(v string)`

SetFtpPassword sets FtpPassword field to given value.

### HasFtpPassword

`func (o *RunFTPUploadTaskVM) HasFtpPassword() bool`

HasFtpPassword returns a boolean if a field has been set.

### SetFtpPasswordNil

`func (o *RunFTPUploadTaskVM) SetFtpPasswordNil(b bool)`

 SetFtpPasswordNil sets the value for FtpPassword to be an explicit nil

### UnsetFtpPassword
`func (o *RunFTPUploadTaskVM) UnsetFtpPassword()`

UnsetFtpPassword ensures that no value is present for FtpPassword, not even an explicit nil
### GetFtpPort

`func (o *RunFTPUploadTaskVM) GetFtpPort() int32`

GetFtpPort returns the FtpPort field if non-nil, zero value otherwise.

### GetFtpPortOk

`func (o *RunFTPUploadTaskVM) GetFtpPortOk() (*int32, bool)`

GetFtpPortOk returns a tuple with the FtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpPort

`func (o *RunFTPUploadTaskVM) SetFtpPort(v int32)`

SetFtpPort sets FtpPort field to given value.

### HasFtpPort

`func (o *RunFTPUploadTaskVM) HasFtpPort() bool`

HasFtpPort returns a boolean if a field has been set.

### GetFtpUsername

`func (o *RunFTPUploadTaskVM) GetFtpUsername() string`

GetFtpUsername returns the FtpUsername field if non-nil, zero value otherwise.

### GetFtpUsernameOk

`func (o *RunFTPUploadTaskVM) GetFtpUsernameOk() (*string, bool)`

GetFtpUsernameOk returns a tuple with the FtpUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpUsername

`func (o *RunFTPUploadTaskVM) SetFtpUsername(v string)`

SetFtpUsername sets FtpUsername field to given value.

### HasFtpUsername

`func (o *RunFTPUploadTaskVM) HasFtpUsername() bool`

HasFtpUsername returns a boolean if a field has been set.

### SetFtpUsernameNil

`func (o *RunFTPUploadTaskVM) SetFtpUsernameNil(b bool)`

 SetFtpUsernameNil sets the value for FtpUsername to be an explicit nil

### UnsetFtpUsername
`func (o *RunFTPUploadTaskVM) UnsetFtpUsername()`

UnsetFtpUsername ensures that no value is present for FtpUsername, not even an explicit nil
### GetUseSFTP

`func (o *RunFTPUploadTaskVM) GetUseSFTP() bool`

GetUseSFTP returns the UseSFTP field if non-nil, zero value otherwise.

### GetUseSFTPOk

`func (o *RunFTPUploadTaskVM) GetUseSFTPOk() (*bool, bool)`

GetUseSFTPOk returns a tuple with the UseSFTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSFTP

`func (o *RunFTPUploadTaskVM) SetUseSFTP(v bool)`

SetUseSFTP sets UseSFTP field to given value.

### HasUseSFTP

`func (o *RunFTPUploadTaskVM) HasUseSFTP() bool`

HasUseSFTP returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


