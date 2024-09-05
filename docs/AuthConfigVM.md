# AuthConfigVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseLocal** | Pointer to **bool** |  | [optional] 
**UseOpenId** | Pointer to **bool** |  | [optional] 
**Authority** | Pointer to **NullableString** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewAuthConfigVM

`func NewAuthConfigVM(t string, ) *AuthConfigVM`

NewAuthConfigVM instantiates a new AuthConfigVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthConfigVMWithDefaults

`func NewAuthConfigVMWithDefaults() *AuthConfigVM`

NewAuthConfigVMWithDefaults instantiates a new AuthConfigVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseLocal

`func (o *AuthConfigVM) GetUseLocal() bool`

GetUseLocal returns the UseLocal field if non-nil, zero value otherwise.

### GetUseLocalOk

`func (o *AuthConfigVM) GetUseLocalOk() (*bool, bool)`

GetUseLocalOk returns a tuple with the UseLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocal

`func (o *AuthConfigVM) SetUseLocal(v bool)`

SetUseLocal sets UseLocal field to given value.

### HasUseLocal

`func (o *AuthConfigVM) HasUseLocal() bool`

HasUseLocal returns a boolean if a field has been set.

### GetUseOpenId

`func (o *AuthConfigVM) GetUseOpenId() bool`

GetUseOpenId returns the UseOpenId field if non-nil, zero value otherwise.

### GetUseOpenIdOk

`func (o *AuthConfigVM) GetUseOpenIdOk() (*bool, bool)`

GetUseOpenIdOk returns a tuple with the UseOpenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOpenId

`func (o *AuthConfigVM) SetUseOpenId(v bool)`

SetUseOpenId sets UseOpenId field to given value.

### HasUseOpenId

`func (o *AuthConfigVM) HasUseOpenId() bool`

HasUseOpenId returns a boolean if a field has been set.

### GetAuthority

`func (o *AuthConfigVM) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *AuthConfigVM) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *AuthConfigVM) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *AuthConfigVM) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### SetAuthorityNil

`func (o *AuthConfigVM) SetAuthorityNil(b bool)`

 SetAuthorityNil sets the value for Authority to be an explicit nil

### UnsetAuthority
`func (o *AuthConfigVM) UnsetAuthority()`

UnsetAuthority ensures that no value is present for Authority, not even an explicit nil
### GetT

`func (o *AuthConfigVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AuthConfigVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AuthConfigVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


