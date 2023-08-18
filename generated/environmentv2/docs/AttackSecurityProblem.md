# AttackSecurityProblem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assessment** | Pointer to [**AttackSecurityProblemAssessmentDto**](AttackSecurityProblemAssessmentDto.md) |  | [optional] 
**SecurityProblemId** | Pointer to **string** | The security problem ID. | [optional] [readonly] 

## Methods

### NewAttackSecurityProblem

`func NewAttackSecurityProblem() *AttackSecurityProblem`

NewAttackSecurityProblem instantiates a new AttackSecurityProblem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackSecurityProblemWithDefaults

`func NewAttackSecurityProblemWithDefaults() *AttackSecurityProblem`

NewAttackSecurityProblemWithDefaults instantiates a new AttackSecurityProblem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssessment

`func (o *AttackSecurityProblem) GetAssessment() AttackSecurityProblemAssessmentDto`

GetAssessment returns the Assessment field if non-nil, zero value otherwise.

### GetAssessmentOk

`func (o *AttackSecurityProblem) GetAssessmentOk() (*AttackSecurityProblemAssessmentDto, bool)`

GetAssessmentOk returns a tuple with the Assessment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessment

`func (o *AttackSecurityProblem) SetAssessment(v AttackSecurityProblemAssessmentDto)`

SetAssessment sets Assessment field to given value.

### HasAssessment

`func (o *AttackSecurityProblem) HasAssessment() bool`

HasAssessment returns a boolean if a field has been set.

### GetSecurityProblemId

`func (o *AttackSecurityProblem) GetSecurityProblemId() string`

GetSecurityProblemId returns the SecurityProblemId field if non-nil, zero value otherwise.

### GetSecurityProblemIdOk

`func (o *AttackSecurityProblem) GetSecurityProblemIdOk() (*string, bool)`

GetSecurityProblemIdOk returns a tuple with the SecurityProblemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProblemId

`func (o *AttackSecurityProblem) SetSecurityProblemId(v string)`

SetSecurityProblemId sets SecurityProblemId field to given value.

### HasSecurityProblemId

`func (o *AttackSecurityProblem) HasSecurityProblemId() bool`

HasSecurityProblemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


