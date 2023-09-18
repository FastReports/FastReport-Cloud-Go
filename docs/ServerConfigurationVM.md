# ServerConfigurationVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**LogoLink** | Pointer to **NullableString** |  | [optional] 
**Copyright** | Pointer to **NullableString** |  | [optional] 
**CorporateServerMode** | Pointer to **bool** |  | [optional] 
**LastSLAVersion** | Pointer to **NullableTime** |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**Frontend** | Pointer to [**FrontendApp**](FrontendApp.md) |  | [optional] 
**InvariantLocale** | Pointer to **NullableString** |  | [optional] 
**Auth** | Pointer to [**AuthConfigVM**](AuthConfigVM.md) |  | [optional] 
**DesignerForAnons** | Pointer to **bool** |  | [optional] 
**SlaLink** | Pointer to **NullableString** |  | [optional] 
**FirstStepsVideoLink** | Pointer to **NullableString** |  | [optional] 
**AboutLink** | Pointer to **NullableString** |  | [optional] 
**HomePageLink** | Pointer to **NullableString** |  | [optional] 
**AuthServerName** | Pointer to **NullableString** |  | [optional] 
**UpdateWorkspaceLink** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewServerConfigurationVM

`func NewServerConfigurationVM() *ServerConfigurationVM`

NewServerConfigurationVM instantiates a new ServerConfigurationVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigurationVMWithDefaults

`func NewServerConfigurationVMWithDefaults() *ServerConfigurationVM`

NewServerConfigurationVMWithDefaults instantiates a new ServerConfigurationVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ServerConfigurationVM) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServerConfigurationVM) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServerConfigurationVM) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ServerConfigurationVM) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ServerConfigurationVM) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ServerConfigurationVM) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetLogoLink

`func (o *ServerConfigurationVM) GetLogoLink() string`

GetLogoLink returns the LogoLink field if non-nil, zero value otherwise.

### GetLogoLinkOk

`func (o *ServerConfigurationVM) GetLogoLinkOk() (*string, bool)`

GetLogoLinkOk returns a tuple with the LogoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoLink

`func (o *ServerConfigurationVM) SetLogoLink(v string)`

SetLogoLink sets LogoLink field to given value.

### HasLogoLink

`func (o *ServerConfigurationVM) HasLogoLink() bool`

HasLogoLink returns a boolean if a field has been set.

### SetLogoLinkNil

`func (o *ServerConfigurationVM) SetLogoLinkNil(b bool)`

 SetLogoLinkNil sets the value for LogoLink to be an explicit nil

### UnsetLogoLink
`func (o *ServerConfigurationVM) UnsetLogoLink()`

UnsetLogoLink ensures that no value is present for LogoLink, not even an explicit nil
### GetCopyright

`func (o *ServerConfigurationVM) GetCopyright() string`

GetCopyright returns the Copyright field if non-nil, zero value otherwise.

### GetCopyrightOk

`func (o *ServerConfigurationVM) GetCopyrightOk() (*string, bool)`

GetCopyrightOk returns a tuple with the Copyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyright

`func (o *ServerConfigurationVM) SetCopyright(v string)`

SetCopyright sets Copyright field to given value.

### HasCopyright

`func (o *ServerConfigurationVM) HasCopyright() bool`

HasCopyright returns a boolean if a field has been set.

### SetCopyrightNil

`func (o *ServerConfigurationVM) SetCopyrightNil(b bool)`

 SetCopyrightNil sets the value for Copyright to be an explicit nil

### UnsetCopyright
`func (o *ServerConfigurationVM) UnsetCopyright()`

UnsetCopyright ensures that no value is present for Copyright, not even an explicit nil
### GetCorporateServerMode

`func (o *ServerConfigurationVM) GetCorporateServerMode() bool`

GetCorporateServerMode returns the CorporateServerMode field if non-nil, zero value otherwise.

### GetCorporateServerModeOk

`func (o *ServerConfigurationVM) GetCorporateServerModeOk() (*bool, bool)`

GetCorporateServerModeOk returns a tuple with the CorporateServerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateServerMode

`func (o *ServerConfigurationVM) SetCorporateServerMode(v bool)`

SetCorporateServerMode sets CorporateServerMode field to given value.

### HasCorporateServerMode

`func (o *ServerConfigurationVM) HasCorporateServerMode() bool`

HasCorporateServerMode returns a boolean if a field has been set.

### GetLastSLAVersion

`func (o *ServerConfigurationVM) GetLastSLAVersion() time.Time`

GetLastSLAVersion returns the LastSLAVersion field if non-nil, zero value otherwise.

### GetLastSLAVersionOk

`func (o *ServerConfigurationVM) GetLastSLAVersionOk() (*time.Time, bool)`

GetLastSLAVersionOk returns a tuple with the LastSLAVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSLAVersion

`func (o *ServerConfigurationVM) SetLastSLAVersion(v time.Time)`

SetLastSLAVersion sets LastSLAVersion field to given value.

### HasLastSLAVersion

`func (o *ServerConfigurationVM) HasLastSLAVersion() bool`

HasLastSLAVersion returns a boolean if a field has been set.

### SetLastSLAVersionNil

`func (o *ServerConfigurationVM) SetLastSLAVersionNil(b bool)`

 SetLastSLAVersionNil sets the value for LastSLAVersion to be an explicit nil

### UnsetLastSLAVersion
`func (o *ServerConfigurationVM) UnsetLastSLAVersion()`

UnsetLastSLAVersion ensures that no value is present for LastSLAVersion, not even an explicit nil
### GetIsDisabled

`func (o *ServerConfigurationVM) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *ServerConfigurationVM) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *ServerConfigurationVM) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *ServerConfigurationVM) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetFrontend

`func (o *ServerConfigurationVM) GetFrontend() FrontendApp`

GetFrontend returns the Frontend field if non-nil, zero value otherwise.

### GetFrontendOk

`func (o *ServerConfigurationVM) GetFrontendOk() (*FrontendApp, bool)`

GetFrontendOk returns a tuple with the Frontend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontend

`func (o *ServerConfigurationVM) SetFrontend(v FrontendApp)`

SetFrontend sets Frontend field to given value.

### HasFrontend

`func (o *ServerConfigurationVM) HasFrontend() bool`

HasFrontend returns a boolean if a field has been set.

### GetInvariantLocale

`func (o *ServerConfigurationVM) GetInvariantLocale() string`

GetInvariantLocale returns the InvariantLocale field if non-nil, zero value otherwise.

### GetInvariantLocaleOk

`func (o *ServerConfigurationVM) GetInvariantLocaleOk() (*string, bool)`

GetInvariantLocaleOk returns a tuple with the InvariantLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvariantLocale

`func (o *ServerConfigurationVM) SetInvariantLocale(v string)`

SetInvariantLocale sets InvariantLocale field to given value.

### HasInvariantLocale

`func (o *ServerConfigurationVM) HasInvariantLocale() bool`

HasInvariantLocale returns a boolean if a field has been set.

### SetInvariantLocaleNil

`func (o *ServerConfigurationVM) SetInvariantLocaleNil(b bool)`

 SetInvariantLocaleNil sets the value for InvariantLocale to be an explicit nil

### UnsetInvariantLocale
`func (o *ServerConfigurationVM) UnsetInvariantLocale()`

UnsetInvariantLocale ensures that no value is present for InvariantLocale, not even an explicit nil
### GetAuth

`func (o *ServerConfigurationVM) GetAuth() AuthConfigVM`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *ServerConfigurationVM) GetAuthOk() (*AuthConfigVM, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *ServerConfigurationVM) SetAuth(v AuthConfigVM)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *ServerConfigurationVM) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetDesignerForAnons

`func (o *ServerConfigurationVM) GetDesignerForAnons() bool`

GetDesignerForAnons returns the DesignerForAnons field if non-nil, zero value otherwise.

### GetDesignerForAnonsOk

`func (o *ServerConfigurationVM) GetDesignerForAnonsOk() (*bool, bool)`

GetDesignerForAnonsOk returns a tuple with the DesignerForAnons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignerForAnons

`func (o *ServerConfigurationVM) SetDesignerForAnons(v bool)`

SetDesignerForAnons sets DesignerForAnons field to given value.

### HasDesignerForAnons

`func (o *ServerConfigurationVM) HasDesignerForAnons() bool`

HasDesignerForAnons returns a boolean if a field has been set.

### GetSlaLink

`func (o *ServerConfigurationVM) GetSlaLink() string`

GetSlaLink returns the SlaLink field if non-nil, zero value otherwise.

### GetSlaLinkOk

`func (o *ServerConfigurationVM) GetSlaLinkOk() (*string, bool)`

GetSlaLinkOk returns a tuple with the SlaLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaLink

`func (o *ServerConfigurationVM) SetSlaLink(v string)`

SetSlaLink sets SlaLink field to given value.

### HasSlaLink

`func (o *ServerConfigurationVM) HasSlaLink() bool`

HasSlaLink returns a boolean if a field has been set.

### SetSlaLinkNil

`func (o *ServerConfigurationVM) SetSlaLinkNil(b bool)`

 SetSlaLinkNil sets the value for SlaLink to be an explicit nil

### UnsetSlaLink
`func (o *ServerConfigurationVM) UnsetSlaLink()`

UnsetSlaLink ensures that no value is present for SlaLink, not even an explicit nil
### GetFirstStepsVideoLink

`func (o *ServerConfigurationVM) GetFirstStepsVideoLink() string`

GetFirstStepsVideoLink returns the FirstStepsVideoLink field if non-nil, zero value otherwise.

### GetFirstStepsVideoLinkOk

`func (o *ServerConfigurationVM) GetFirstStepsVideoLinkOk() (*string, bool)`

GetFirstStepsVideoLinkOk returns a tuple with the FirstStepsVideoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstStepsVideoLink

`func (o *ServerConfigurationVM) SetFirstStepsVideoLink(v string)`

SetFirstStepsVideoLink sets FirstStepsVideoLink field to given value.

### HasFirstStepsVideoLink

`func (o *ServerConfigurationVM) HasFirstStepsVideoLink() bool`

HasFirstStepsVideoLink returns a boolean if a field has been set.

### SetFirstStepsVideoLinkNil

`func (o *ServerConfigurationVM) SetFirstStepsVideoLinkNil(b bool)`

 SetFirstStepsVideoLinkNil sets the value for FirstStepsVideoLink to be an explicit nil

### UnsetFirstStepsVideoLink
`func (o *ServerConfigurationVM) UnsetFirstStepsVideoLink()`

UnsetFirstStepsVideoLink ensures that no value is present for FirstStepsVideoLink, not even an explicit nil
### GetAboutLink

`func (o *ServerConfigurationVM) GetAboutLink() string`

GetAboutLink returns the AboutLink field if non-nil, zero value otherwise.

### GetAboutLinkOk

`func (o *ServerConfigurationVM) GetAboutLinkOk() (*string, bool)`

GetAboutLinkOk returns a tuple with the AboutLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAboutLink

`func (o *ServerConfigurationVM) SetAboutLink(v string)`

SetAboutLink sets AboutLink field to given value.

### HasAboutLink

`func (o *ServerConfigurationVM) HasAboutLink() bool`

HasAboutLink returns a boolean if a field has been set.

### SetAboutLinkNil

`func (o *ServerConfigurationVM) SetAboutLinkNil(b bool)`

 SetAboutLinkNil sets the value for AboutLink to be an explicit nil

### UnsetAboutLink
`func (o *ServerConfigurationVM) UnsetAboutLink()`

UnsetAboutLink ensures that no value is present for AboutLink, not even an explicit nil
### GetHomePageLink

`func (o *ServerConfigurationVM) GetHomePageLink() string`

GetHomePageLink returns the HomePageLink field if non-nil, zero value otherwise.

### GetHomePageLinkOk

`func (o *ServerConfigurationVM) GetHomePageLinkOk() (*string, bool)`

GetHomePageLinkOk returns a tuple with the HomePageLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePageLink

`func (o *ServerConfigurationVM) SetHomePageLink(v string)`

SetHomePageLink sets HomePageLink field to given value.

### HasHomePageLink

`func (o *ServerConfigurationVM) HasHomePageLink() bool`

HasHomePageLink returns a boolean if a field has been set.

### SetHomePageLinkNil

`func (o *ServerConfigurationVM) SetHomePageLinkNil(b bool)`

 SetHomePageLinkNil sets the value for HomePageLink to be an explicit nil

### UnsetHomePageLink
`func (o *ServerConfigurationVM) UnsetHomePageLink()`

UnsetHomePageLink ensures that no value is present for HomePageLink, not even an explicit nil
### GetAuthServerName

`func (o *ServerConfigurationVM) GetAuthServerName() string`

GetAuthServerName returns the AuthServerName field if non-nil, zero value otherwise.

### GetAuthServerNameOk

`func (o *ServerConfigurationVM) GetAuthServerNameOk() (*string, bool)`

GetAuthServerNameOk returns a tuple with the AuthServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServerName

`func (o *ServerConfigurationVM) SetAuthServerName(v string)`

SetAuthServerName sets AuthServerName field to given value.

### HasAuthServerName

`func (o *ServerConfigurationVM) HasAuthServerName() bool`

HasAuthServerName returns a boolean if a field has been set.

### SetAuthServerNameNil

`func (o *ServerConfigurationVM) SetAuthServerNameNil(b bool)`

 SetAuthServerNameNil sets the value for AuthServerName to be an explicit nil

### UnsetAuthServerName
`func (o *ServerConfigurationVM) UnsetAuthServerName()`

UnsetAuthServerName ensures that no value is present for AuthServerName, not even an explicit nil
### GetUpdateWorkspaceLink

`func (o *ServerConfigurationVM) GetUpdateWorkspaceLink() string`

GetUpdateWorkspaceLink returns the UpdateWorkspaceLink field if non-nil, zero value otherwise.

### GetUpdateWorkspaceLinkOk

`func (o *ServerConfigurationVM) GetUpdateWorkspaceLinkOk() (*string, bool)`

GetUpdateWorkspaceLinkOk returns a tuple with the UpdateWorkspaceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateWorkspaceLink

`func (o *ServerConfigurationVM) SetUpdateWorkspaceLink(v string)`

SetUpdateWorkspaceLink sets UpdateWorkspaceLink field to given value.

### HasUpdateWorkspaceLink

`func (o *ServerConfigurationVM) HasUpdateWorkspaceLink() bool`

HasUpdateWorkspaceLink returns a boolean if a field has been set.

### SetUpdateWorkspaceLinkNil

`func (o *ServerConfigurationVM) SetUpdateWorkspaceLinkNil(b bool)`

 SetUpdateWorkspaceLinkNil sets the value for UpdateWorkspaceLink to be an explicit nil

### UnsetUpdateWorkspaceLink
`func (o *ServerConfigurationVM) UnsetUpdateWorkspaceLink()`

UnsetUpdateWorkspaceLink ensures that no value is present for UpdateWorkspaceLink, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


