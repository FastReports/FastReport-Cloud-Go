# FTPUploadTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archive** | Pointer to **bool** |  | [optional] 
**ArchiveName** | Pointer to **NullableString** |  | [optional] 
**DestinationFolder** | Pointer to **NullableString** |  | [optional] 
**FtpHost** | Pointer to **NullableString** |  | [optional] 
**FtpPort** | Pointer to **int32** |  | [optional] 
**FtpUsername** | Pointer to **NullableString** |  | [optional] 
**UseSFTP** | Pointer to **bool** |  | [optional] 

## Methods

### NewFTPUploadTaskVM

`func NewFTPUploadTaskVM() *FTPUploadTaskVM`

NewFTPUploadTaskVM instantiates a new FTPUploadTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFTPUploadTaskVMWithDefaults

`func NewFTPUploadTaskVMWithDefaults() *FTPUploadTaskVM`

NewFTPUploadTaskVMWithDefaults instantiates a new FTPUploadTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchive

`func (o *FTPUploadTaskVM) GetArchive() bool`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *FTPUploadTaskVM) GetArchiveOk() (*bool, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *FTPUploadTaskVM) SetArchive(v bool)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *FTPUploadTaskVM) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetArchiveName

`func (o *FTPUploadTaskVM) GetArchiveName() string`

GetArchiveName returns the ArchiveName field if non-nil, zero value otherwise.

### GetArchiveNameOk

`func (o *FTPUploadTaskVM) GetArchiveNameOk() (*string, bool)`

GetArchiveNameOk returns a tuple with the ArchiveName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveName

`func (o *FTPUploadTaskVM) SetArchiveName(v string)`

SetArchiveName sets ArchiveName field to given value.

### HasArchiveName

`func (o *FTPUploadTaskVM) HasArchiveName() bool`

HasArchiveName returns a boolean if a field has been set.

### SetArchiveNameNil

`func (o *FTPUploadTaskVM) SetArchiveNameNil(b bool)`

 SetArchiveNameNil sets the value for ArchiveName to be an explicit nil

### UnsetArchiveName
`func (o *FTPUploadTaskVM) UnsetArchiveName()`

UnsetArchiveName ensures that no value is present for ArchiveName, not even an explicit nil
### GetDestinationFolder

`func (o *FTPUploadTaskVM) GetDestinationFolder() string`

GetDestinationFolder returns the DestinationFolder field if non-nil, zero value otherwise.

### GetDestinationFolderOk

`func (o *FTPUploadTaskVM) GetDestinationFolderOk() (*string, bool)`

GetDestinationFolderOk returns a tuple with the DestinationFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationFolder

`func (o *FTPUploadTaskVM) SetDestinationFolder(v string)`

SetDestinationFolder sets DestinationFolder field to given value.

### HasDestinationFolder

`func (o *FTPUploadTaskVM) HasDestinationFolder() bool`

HasDestinationFolder returns a boolean if a field has been set.

### SetDestinationFolderNil

`func (o *FTPUploadTaskVM) SetDestinationFolderNil(b bool)`

 SetDestinationFolderNil sets the value for DestinationFolder to be an explicit nil

### UnsetDestinationFolder
`func (o *FTPUploadTaskVM) UnsetDestinationFolder()`

UnsetDestinationFolder ensures that no value is present for DestinationFolder, not even an explicit nil
### GetFtpHost

`func (o *FTPUploadTaskVM) GetFtpHost() string`

GetFtpHost returns the FtpHost field if non-nil, zero value otherwise.

### GetFtpHostOk

`func (o *FTPUploadTaskVM) GetFtpHostOk() (*string, bool)`

GetFtpHostOk returns a tuple with the FtpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpHost

`func (o *FTPUploadTaskVM) SetFtpHost(v string)`

SetFtpHost sets FtpHost field to given value.

### HasFtpHost

`func (o *FTPUploadTaskVM) HasFtpHost() bool`

HasFtpHost returns a boolean if a field has been set.

### SetFtpHostNil

`func (o *FTPUploadTaskVM) SetFtpHostNil(b bool)`

 SetFtpHostNil sets the value for FtpHost to be an explicit nil

### UnsetFtpHost
`func (o *FTPUploadTaskVM) UnsetFtpHost()`

UnsetFtpHost ensures that no value is present for FtpHost, not even an explicit nil
### GetFtpPort

`func (o *FTPUploadTaskVM) GetFtpPort() int32`

GetFtpPort returns the FtpPort field if non-nil, zero value otherwise.

### GetFtpPortOk

`func (o *FTPUploadTaskVM) GetFtpPortOk() (*int32, bool)`

GetFtpPortOk returns a tuple with the FtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpPort

`func (o *FTPUploadTaskVM) SetFtpPort(v int32)`

SetFtpPort sets FtpPort field to given value.

### HasFtpPort

`func (o *FTPUploadTaskVM) HasFtpPort() bool`

HasFtpPort returns a boolean if a field has been set.

### GetFtpUsername

`func (o *FTPUploadTaskVM) GetFtpUsername() string`

GetFtpUsername returns the FtpUsername field if non-nil, zero value otherwise.

### GetFtpUsernameOk

`func (o *FTPUploadTaskVM) GetFtpUsernameOk() (*string, bool)`

GetFtpUsernameOk returns a tuple with the FtpUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpUsername

`func (o *FTPUploadTaskVM) SetFtpUsername(v string)`

SetFtpUsername sets FtpUsername field to given value.

### HasFtpUsername

`func (o *FTPUploadTaskVM) HasFtpUsername() bool`

HasFtpUsername returns a boolean if a field has been set.

### SetFtpUsernameNil

`func (o *FTPUploadTaskVM) SetFtpUsernameNil(b bool)`

 SetFtpUsernameNil sets the value for FtpUsername to be an explicit nil

### UnsetFtpUsername
`func (o *FTPUploadTaskVM) UnsetFtpUsername()`

UnsetFtpUsername ensures that no value is present for FtpUsername, not even an explicit nil
### GetUseSFTP

`func (o *FTPUploadTaskVM) GetUseSFTP() bool`

GetUseSFTP returns the UseSFTP field if non-nil, zero value otherwise.

### GetUseSFTPOk

`func (o *FTPUploadTaskVM) GetUseSFTPOk() (*bool, bool)`

GetUseSFTPOk returns a tuple with the UseSFTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSFTP

`func (o *FTPUploadTaskVM) SetUseSFTP(v bool)`

SetUseSFTP sets UseSFTP field to given value.

### HasUseSFTP

`func (o *FTPUploadTaskVM) HasUseSFTP() bool`

HasUseSFTP returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


