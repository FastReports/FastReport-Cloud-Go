# PreviewTemplateVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locale** | Pointer to **NullableString** |  | [optional] 
**ReportParameters** | Pointer to **map[string]string** |  | [optional] 
**CacheTolerance** | Pointer to **float64** |  | [optional] [default to 300]

## Methods

### NewPreviewTemplateVM

`func NewPreviewTemplateVM() *PreviewTemplateVM`

NewPreviewTemplateVM instantiates a new PreviewTemplateVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewTemplateVMWithDefaults

`func NewPreviewTemplateVMWithDefaults() *PreviewTemplateVM`

NewPreviewTemplateVMWithDefaults instantiates a new PreviewTemplateVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocale

`func (o *PreviewTemplateVM) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *PreviewTemplateVM) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *PreviewTemplateVM) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *PreviewTemplateVM) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *PreviewTemplateVM) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *PreviewTemplateVM) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetReportParameters

`func (o *PreviewTemplateVM) GetReportParameters() map[string]string`

GetReportParameters returns the ReportParameters field if non-nil, zero value otherwise.

### GetReportParametersOk

`func (o *PreviewTemplateVM) GetReportParametersOk() (*map[string]string, bool)`

GetReportParametersOk returns a tuple with the ReportParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportParameters

`func (o *PreviewTemplateVM) SetReportParameters(v map[string]string)`

SetReportParameters sets ReportParameters field to given value.

### HasReportParameters

`func (o *PreviewTemplateVM) HasReportParameters() bool`

HasReportParameters returns a boolean if a field has been set.

### SetReportParametersNil

`func (o *PreviewTemplateVM) SetReportParametersNil(b bool)`

 SetReportParametersNil sets the value for ReportParameters to be an explicit nil

### UnsetReportParameters
`func (o *PreviewTemplateVM) UnsetReportParameters()`

UnsetReportParameters ensures that no value is present for ReportParameters, not even an explicit nil
### GetCacheTolerance

`func (o *PreviewTemplateVM) GetCacheTolerance() float64`

GetCacheTolerance returns the CacheTolerance field if non-nil, zero value otherwise.

### GetCacheToleranceOk

`func (o *PreviewTemplateVM) GetCacheToleranceOk() (*float64, bool)`

GetCacheToleranceOk returns a tuple with the CacheTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTolerance

`func (o *PreviewTemplateVM) SetCacheTolerance(v float64)`

SetCacheTolerance sets CacheTolerance field to given value.

### HasCacheTolerance

`func (o *PreviewTemplateVM) HasCacheTolerance() bool`

HasCacheTolerance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


