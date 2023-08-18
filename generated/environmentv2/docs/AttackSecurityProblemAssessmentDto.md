# AttackSecurityProblemAssessmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataAssets** | Pointer to **string** | The reachability of data assets by the attacked target. | [optional] [readonly] 
**Exposure** | Pointer to **string** | The level of exposure of the attacked target | [optional] [readonly] 
**NumberOfReachableDataAssets** | Pointer to **int32** | The number of data assets reachable by the attacked target. | [optional] [readonly] 

## Methods

### NewAttackSecurityProblemAssessmentDto

`func NewAttackSecurityProblemAssessmentDto() *AttackSecurityProblemAssessmentDto`

NewAttackSecurityProblemAssessmentDto instantiates a new AttackSecurityProblemAssessmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackSecurityProblemAssessmentDtoWithDefaults

`func NewAttackSecurityProblemAssessmentDtoWithDefaults() *AttackSecurityProblemAssessmentDto`

NewAttackSecurityProblemAssessmentDtoWithDefaults instantiates a new AttackSecurityProblemAssessmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataAssets

`func (o *AttackSecurityProblemAssessmentDto) GetDataAssets() string`

GetDataAssets returns the DataAssets field if non-nil, zero value otherwise.

### GetDataAssetsOk

`func (o *AttackSecurityProblemAssessmentDto) GetDataAssetsOk() (*string, bool)`

GetDataAssetsOk returns a tuple with the DataAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAssets

`func (o *AttackSecurityProblemAssessmentDto) SetDataAssets(v string)`

SetDataAssets sets DataAssets field to given value.

### HasDataAssets

`func (o *AttackSecurityProblemAssessmentDto) HasDataAssets() bool`

HasDataAssets returns a boolean if a field has been set.

### GetExposure

`func (o *AttackSecurityProblemAssessmentDto) GetExposure() string`

GetExposure returns the Exposure field if non-nil, zero value otherwise.

### GetExposureOk

`func (o *AttackSecurityProblemAssessmentDto) GetExposureOk() (*string, bool)`

GetExposureOk returns a tuple with the Exposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposure

`func (o *AttackSecurityProblemAssessmentDto) SetExposure(v string)`

SetExposure sets Exposure field to given value.

### HasExposure

`func (o *AttackSecurityProblemAssessmentDto) HasExposure() bool`

HasExposure returns a boolean if a field has been set.

### GetNumberOfReachableDataAssets

`func (o *AttackSecurityProblemAssessmentDto) GetNumberOfReachableDataAssets() int32`

GetNumberOfReachableDataAssets returns the NumberOfReachableDataAssets field if non-nil, zero value otherwise.

### GetNumberOfReachableDataAssetsOk

`func (o *AttackSecurityProblemAssessmentDto) GetNumberOfReachableDataAssetsOk() (*int32, bool)`

GetNumberOfReachableDataAssetsOk returns a tuple with the NumberOfReachableDataAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfReachableDataAssets

`func (o *AttackSecurityProblemAssessmentDto) SetNumberOfReachableDataAssets(v int32)`

SetNumberOfReachableDataAssets sets NumberOfReachableDataAssets field to given value.

### HasNumberOfReachableDataAssets

`func (o *AttackSecurityProblemAssessmentDto) HasNumberOfReachableDataAssets() bool`

HasNumberOfReachableDataAssets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


