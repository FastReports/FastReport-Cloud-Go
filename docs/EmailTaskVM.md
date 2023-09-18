# EmailTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **NullableString** |  | [optional] 
**EnableSsl** | Pointer to **bool** |  | [optional] 
**From** | Pointer to **NullableString** |  | [optional] 
**IsBodyHtml** | Pointer to **bool** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Server** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**To** | Pointer to **[]string** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmailTaskVM

`func NewEmailTaskVM() *EmailTaskVM`

NewEmailTaskVM instantiates a new EmailTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTaskVMWithDefaults

`func NewEmailTaskVMWithDefaults() *EmailTaskVM`

NewEmailTaskVMWithDefaults instantiates a new EmailTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *EmailTaskVM) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EmailTaskVM) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EmailTaskVM) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *EmailTaskVM) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *EmailTaskVM) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *EmailTaskVM) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetEnableSsl

`func (o *EmailTaskVM) GetEnableSsl() bool`

GetEnableSsl returns the EnableSsl field if non-nil, zero value otherwise.

### GetEnableSslOk

`func (o *EmailTaskVM) GetEnableSslOk() (*bool, bool)`

GetEnableSslOk returns a tuple with the EnableSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSsl

`func (o *EmailTaskVM) SetEnableSsl(v bool)`

SetEnableSsl sets EnableSsl field to given value.

### HasEnableSsl

`func (o *EmailTaskVM) HasEnableSsl() bool`

HasEnableSsl returns a boolean if a field has been set.

### GetFrom

`func (o *EmailTaskVM) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailTaskVM) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailTaskVM) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EmailTaskVM) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *EmailTaskVM) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *EmailTaskVM) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetIsBodyHtml

`func (o *EmailTaskVM) GetIsBodyHtml() bool`

GetIsBodyHtml returns the IsBodyHtml field if non-nil, zero value otherwise.

### GetIsBodyHtmlOk

`func (o *EmailTaskVM) GetIsBodyHtmlOk() (*bool, bool)`

GetIsBodyHtmlOk returns a tuple with the IsBodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBodyHtml

`func (o *EmailTaskVM) SetIsBodyHtml(v bool)`

SetIsBodyHtml sets IsBodyHtml field to given value.

### HasIsBodyHtml

`func (o *EmailTaskVM) HasIsBodyHtml() bool`

HasIsBodyHtml returns a boolean if a field has been set.

### GetPort

`func (o *EmailTaskVM) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EmailTaskVM) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EmailTaskVM) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *EmailTaskVM) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetServer

`func (o *EmailTaskVM) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *EmailTaskVM) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *EmailTaskVM) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *EmailTaskVM) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *EmailTaskVM) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *EmailTaskVM) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetSubject

`func (o *EmailTaskVM) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailTaskVM) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailTaskVM) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailTaskVM) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *EmailTaskVM) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *EmailTaskVM) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetTo

`func (o *EmailTaskVM) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailTaskVM) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailTaskVM) SetTo(v []string)`

SetTo sets To field to given value.

### HasTo

`func (o *EmailTaskVM) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *EmailTaskVM) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *EmailTaskVM) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetUsername

`func (o *EmailTaskVM) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EmailTaskVM) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EmailTaskVM) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EmailTaskVM) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *EmailTaskVM) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *EmailTaskVM) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


