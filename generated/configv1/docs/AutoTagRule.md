# AutoTagRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | [**[]EntityRuleEngineCondition**](EntityRuleEngineCondition.md) | A list of matching rules for the auto-tag.   The tag applies only when **all** conditions are fulfilled. | 
**Enabled** | **bool** | Tag rule is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**Normalization** | Pointer to **NullableString** | Changes applied to the value after applying the value format. Default is LEAVE_TEXT_AS_IS. | [optional] 
**PropagationTypes** | Pointer to **[]string** | How to apply the tag to underlying entities:   * &#x60;SERVICE_TO_PROCESS_GROUP_LIKE&#x60;: Apply to underlying process groups of matching services.  * &#x60;SERVICE_TO_HOST_LIKE&#x60;: Apply to underlying hosts of matching services.  * &#x60;PROCESS_GROUP_TO_HOST&#x60;: Apply to underlying hosts of matching process groups.  * &#x60;PROCESS_GROUP_TO_SERVICE&#x60;: Apply to all services provided by the process groups.  * &#x60;HOST_TO_PROCESS_GROUP_INSTANCE&#x60;: Apply to processes running on matching hosts.  *  &#x60;AZURE_TO_PG&#x60;: Apply to process groups connected to matching Azure entities.  * &#x60;AZURE_TO_SERVICE&#x60;: Apply to services provided by matching Azure entities. | [optional] 
**Type** | **string** | Type of entities to which the rule applies. | 
**ValueFormat** | Pointer to **string** | The value of the auto-tag. If specified, the tag is used in the &#x60;name:valueFormat&#x60; format.   For example, you can extend the &#x60;Infrastructure&#x60; tag to &#x60;Infrastructure:Windows&#x60; and &#x60;Infrastructure:Linux&#x60;.   You can use the following placeholders here:   * &#x60;{AwsAutoScalingGroup:Name}&#x60;  * &#x60;{AwsAvailabilityZone:Name}&#x60;  * &#x60;{AwsElasticLoadBalancer:Name}&#x60;  * &#x60;{AwsRelationalDatabaseService:DBName}&#x60;  * &#x60;{AwsRelationalDatabaseService:Endpoint}&#x60;  * &#x60;{AwsRelationalDatabaseService:Engine}&#x60;  * &#x60;{AwsRelationalDatabaseService:InstanceClass}&#x60;  * &#x60;{AwsRelationalDatabaseService:Name}&#x60;  * &#x60;{AwsRelationalDatabaseService:Port}&#x60;  * &#x60;{AzureRegion:Name}&#x60;  * &#x60;{AzureScaleSet:Name}&#x60;  * &#x60;{AzureVm:Name}&#x60;  * &#x60;{CloudFoundryOrganization:Name}&#x60;  * &#x60;{CustomDevice:DetectedName}&#x60;  * &#x60;{CustomDevice:DnsName}&#x60;  * &#x60;{CustomDevice:IpAddress}&#x60;  * &#x60;{CustomDevice:Port}&#x60;  * &#x60;{DockerContainerGroupInstance:ContainerName}&#x60;  * &#x60;{DockerContainerGroupInstance:FullImageName}&#x60;  * &#x60;{DockerContainerGroupInstance:ImageVersion}&#x60;  * &#x60;{ESXIHost:HardwareModel}&#x60;  * &#x60;{ESXIHost:HardwareVendor}&#x60;  * &#x60;{ESXIHost:Name}&#x60;  * &#x60;{ESXIHost:ProductName}&#x60;  * &#x60;{ESXIHost:ProductVersion}&#x60;  * &#x60;{Ec2Instance:AmiId}&#x60;  * &#x60;{Ec2Instance:BeanstalkEnvironmentName}&#x60;  * &#x60;{Ec2Instance:InstanceId}&#x60;  * &#x60;{Ec2Instance:InstanceType}&#x60;  * &#x60;{Ec2Instance:LocalHostName}&#x60;  * &#x60;{Ec2Instance:Name}&#x60;  * &#x60;{Ec2Instance:PublicHostName}&#x60;  * &#x60;{Ec2Instance:SecurityGroup}&#x60;  * &#x60;{GoogleComputeInstance:Id}&#x60;  * &#x60;{GoogleComputeInstance:IpAddresses}&#x60;  * &#x60;{GoogleComputeInstance:MachineType}&#x60;  * &#x60;{GoogleComputeInstance:Name}&#x60;  * &#x60;{GoogleComputeInstance:ProjectId}&#x60;  * &#x60;{GoogleComputeInstance:Project}&#x60;  * &#x60;{Host:AWSNameTag}&#x60;  * &#x60;{Host:AixLogicalCpuCount}&#x60;  * &#x60;{Host:AzureHostName}&#x60;  * &#x60;{Host:AzureSiteName}&#x60;  * &#x60;{Host:BoshDeploymentId}&#x60;  * &#x60;{Host:BoshInstanceId}&#x60;  * &#x60;{Host:BoshInstanceName}&#x60;  * &#x60;{Host:BoshName}&#x60;  * &#x60;{Host:BoshStemcellVersion}&#x60;  * &#x60;{Host:CpuCores}&#x60;  * &#x60;{Host:DetectedName}&#x60;  * &#x60;{Host:Environment:AppName}&#x60;  * &#x60;{Host:Environment:BoshReleaseVersion}&#x60;  * &#x60;{Host:Environment:Environment}&#x60;  * &#x60;{Host:Environment:Link}&#x60;  * &#x60;{Host:Environment:Organization}&#x60;  * &#x60;{Host:Environment:Owner}&#x60;  * &#x60;{Host:Environment:Support}&#x60;  * &#x60;{Host:IpAddress}&#x60;  * &#x60;{Host:LogicalCpuCores}&#x60;  * &#x60;{Host:OneAgentCustomHostName}&#x60;  * &#x60;{Host:OperatingSystemVersion}&#x60;  * &#x60;{Host:PaasMemoryLimit}&#x60;  * &#x60;{HostGroup:Name}&#x60;  * &#x60;{KubernetesCluster:Name}&#x60;  * &#x60;{KubernetesNode:DetectedName}&#x60;  * &#x60;{OpenstackAvailabilityZone:Name}&#x60;  * &#x60;{OpenstackZone:Name}&#x60;  * &#x60;{OpenstackComputeNode:Name}&#x60;  * &#x60;{OpenstackProject:Name}&#x60;  * &#x60;{OpenstackVm:UnstanceType}&#x60;  * &#x60;{OpenstackVm:Name}&#x60;  * &#x60;{OpenstackVm:SecurityGroup}&#x60;  * &#x60;{ProcessGroup:AmazonECRImageAccountId}&#x60;  * &#x60;{ProcessGroup:AmazonECRImageRegion}&#x60;  * &#x60;{ProcessGroup:AmazonECSCluster}&#x60;  * &#x60;{ProcessGroup:AmazonECSContainerName}&#x60;  * &#x60;{ProcessGroup:AmazonECSFamily}&#x60;  * &#x60;{ProcessGroup:AmazonECSRevision}&#x60;  * &#x60;{ProcessGroup:AmazonLambdaFunctionName}&#x60;  * &#x60;{ProcessGroup:AmazonRegion}&#x60;  * &#x60;{ProcessGroup:ApacheConfigPath}&#x60;  * &#x60;{ProcessGroup:ApacheSparkMasterIpAddress}&#x60;  * &#x60;{ProcessGroup:AspDotNetCoreApplicationPath}&#x60;  * &#x60;{ProcessGroup:AspDotNetCoreApplicationPath}&#x60;  * &#x60;{ProcessGroup:AzureHostName}&#x60;  * &#x60;{ProcessGroup:AzureSiteName}&#x60;  * &#x60;{ProcessGroup:CassandraClusterName}&#x60;  * &#x60;{ProcessGroup:CatalinaBase}&#x60;  * &#x60;{ProcessGroup:CatalinaHome}&#x60;  * &#x60;{ProcessGroup:CloudFoundryAppId}&#x60;  * &#x60;{ProcessGroup:CloudFoundryAppName}&#x60;  * &#x60;{ProcessGroup:CloudFoundryInstanceIndex}&#x60;  * &#x60;{ProcessGroup:CloudFoundrySpaceId}&#x60;  * &#x60;{ProcessGroup:CloudFoundrySpaceName}&#x60;  * &#x60;{ProcessGroup:ColdFusionJvmConfigFile}&#x60;  * &#x60;{ProcessGroup:ColdFusionServiceName}&#x60;  * &#x60;{ProcessGroup:CommandLineArgs}&#x60;  * &#x60;{ProcessGroup:DetectedName}&#x60;  * &#x60;{ProcessGroup:DotNetCommandPath}&#x60;  * &#x60;{ProcessGroup:DotNetCommand}&#x60;  * &#x60;{ProcessGroup:DotNetClusterId}&#x60;  * &#x60;{ProcessGroup:DotNetNodeId}&#x60;  * &#x60;{ProcessGroup:ElasticsearchClusterName}&#x60;  * &#x60;{ProcessGroup:ElasticsearchNodeName}&#x60;  * &#x60;{ProcessGroup:EquinoxConfigPath}&#x60;  * &#x60;{ProcessGroup:ExeName}&#x60;  * &#x60;{ProcessGroup:ExePath}&#x60;  * &#x60;{ProcessGroup:GlassFishDomainName}&#x60;  * &#x60;{ProcessGroup:GlassFishInstanceName}&#x60;  * &#x60;{ProcessGroup:GoogleAppEngineInstance}&#x60;  * &#x60;{ProcessGroup:GoogleAppEngineService}&#x60;  * &#x60;{ProcessGroup:GoogleCloudProject}&#x60;  * &#x60;{ProcessGroup:HybrisBinDirectory}&#x60;  * &#x60;{ProcessGroup:HybrisConfigDirectory}&#x60;  * &#x60;{ProcessGroup:HybrisConfigDirectory}&#x60;  * &#x60;{ProcessGroup:HybrisDataDirectory}&#x60;  * &#x60;{ProcessGroup:IBMCicsRegion}&#x60;  * &#x60;{ProcessGroup:IBMCtgName}&#x60;  * &#x60;{ProcessGroup:IBMImsConnectRegion}&#x60;  * &#x60;{ProcessGroup:IBMImsControlRegion}&#x60;  * &#x60;{ProcessGroup:IBMImsMessageProcessingRegion}&#x60;  * &#x60;{ProcessGroup:IBMImsSoapGwName}&#x60;  * &#x60;{ProcessGroup:IBMIntegrationNodeName}&#x60;  * &#x60;{ProcessGroup:IBMIntegrationServerName}&#x60;  * &#x60;{ProcessGroup:IISAppPool}&#x60;  * &#x60;{ProcessGroup:IISRoleName}&#x60;  * &#x60;{ProcessGroup:JbossHome}&#x60;  * &#x60;{ProcessGroup:JbossMode}&#x60;  * &#x60;{ProcessGroup:JbossServerName}&#x60;  * &#x60;{ProcessGroup:JavaJarFile}&#x60;  * &#x60;{ProcessGroup:JavaJarPath}&#x60;  * &#x60;{ProcessGroup:JavaMainCLass}&#x60;  * &#x60;{ProcessGroup:KubernetesBasePodName}&#x60;  * &#x60;{ProcessGroup:KubernetesContainerName}&#x60;  * &#x60;{ProcessGroup:KubernetesFullPodName}&#x60;  * &#x60;{ProcessGroup:KubernetesNamespace}&#x60;  * &#x60;{ProcessGroup:KubernetesPodUid}&#x60;  * &#x60;{ProcessGroup:MssqlInstanceName}&#x60;  * &#x60;{ProcessGroup:NodeJsAppBaseDirectory}&#x60;  * &#x60;{ProcessGroup:NodeJsAppName}&#x60;  * &#x60;{ProcessGroup:NodeJsScriptName}&#x60;  * &#x60;{ProcessGroup:OracleSid}&#x60;  * &#x60;{ProcessGroup:PHPScriptPath}&#x60;  * &#x60;{ProcessGroup:PHPWorkingDirectory}&#x60;  * &#x60;{ProcessGroup:Ports}&#x60;  * &#x60;{ProcessGroup:RubyAppRootPath}&#x60;  * &#x60;{ProcessGroup:RubyScriptPath}&#x60;  * &#x60;{ProcessGroup:SoftwareAGInstallRoot}&#x60;  * &#x60;{ProcessGroup:SoftwareAGProductPropertyName}&#x60;  * &#x60;{ProcessGroup:SpringBootAppName}&#x60;  * &#x60;{ProcessGroup:SpringBootProfileName}&#x60;  * &#x60;{ProcessGroup:SpringBootStartupClass}&#x60;  * &#x60;{ProcessGroup:TIBCOBusinessWorksAppNodeName}&#x60;  * &#x60;{ProcessGroup:TIBCOBusinessWorksAppSpaceName}&#x60;  * &#x60;{ProcessGroup:TIBCOBusinessWorksCeAppName}&#x60;  * &#x60;{ProcessGroup:TIBCOBusinessWorksCeVersion}&#x60;  * &#x60;{ProcessGroup:TIBCOBusinessWorksDomainName}&#x60;  * &#x60;{ProcessGroup:TIBCOBusinessWorksEnginePropertyFilePath}&#x60;  * &#x60;{ProcessGroup:TIBCOBusinessWorksEnginePropertyFile}&#x60;  * &#x60;{ProcessGroup:TIBCOBusinessWorksHome}&#x60;  * &#x60;{ProcessGroup:VarnishInstanceName}&#x60;  * &#x60;{ProcessGroup:WebLogicClusterName}&#x60;  * &#x60;{ProcessGroup:WebLogicDomainName}&#x60;  * &#x60;{ProcessGroup:WebLogicHome}&#x60;  * &#x60;{ProcessGroup:WebLogicName}&#x60;  * &#x60;{ProcessGroup:WebSphereCellName}&#x60;  * &#x60;{ProcessGroup:WebSphereClusterName}&#x60;  * &#x60;{ProcessGroup:WebSphereNodeName}&#x60;  * &#x60;{ProcessGroup:WebSphereServerName}&#x60;  * &#x60;{ProcessGroup:ActorSystem}&#x60;  * &#x60;{Service:STGServerName}&#x60;  * &#x60;{Service:DatabaseHostName}&#x60;  * &#x60;{Service:DatabaseName}&#x60;  * &#x60;{Service:DatabaseVendor}&#x60;  * &#x60;{Service:DetectedName}&#x60;  * &#x60;{Service:EndpointPath}&#x60;  * &#x60;{Service:EndpointPathGatewayUrl}&#x60;  * &#x60;{Service:IIBApplicationName}&#x60;  * &#x60;{Service:MessageListenerClassName}&#x60;  * &#x60;{Service:Port}&#x60;  * &#x60;{Service:PublicDomainName}&#x60;  * &#x60;{Service:RemoteEndpoint}&#x60;  * &#x60;{Service:RemoteName}&#x60;  * &#x60;{Service:WebApplicationId}&#x60;  * &#x60;{Service:WebContextRoot}&#x60;  * &#x60;{Service:WebServerName}&#x60;  * &#x60;{Service:WebServiceNamespace}&#x60;  * &#x60;{Service:WebServiceName}&#x60;  * &#x60;{VmwareDatacenter:Name}&#x60;  * &#x60;{VmwareVm:Name}&#x60;   | [optional] 

## Methods

### NewAutoTagRule

`func NewAutoTagRule(conditions []EntityRuleEngineCondition, enabled bool, type_ string, ) *AutoTagRule`

NewAutoTagRule instantiates a new AutoTagRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTagRuleWithDefaults

`func NewAutoTagRuleWithDefaults() *AutoTagRule`

NewAutoTagRuleWithDefaults instantiates a new AutoTagRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *AutoTagRule) GetConditions() []EntityRuleEngineCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AutoTagRule) GetConditionsOk() (*[]EntityRuleEngineCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AutoTagRule) SetConditions(v []EntityRuleEngineCondition)`

SetConditions sets Conditions field to given value.


### GetEnabled

`func (o *AutoTagRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoTagRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoTagRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetNormalization

`func (o *AutoTagRule) GetNormalization() string`

GetNormalization returns the Normalization field if non-nil, zero value otherwise.

### GetNormalizationOk

`func (o *AutoTagRule) GetNormalizationOk() (*string, bool)`

GetNormalizationOk returns a tuple with the Normalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalization

`func (o *AutoTagRule) SetNormalization(v string)`

SetNormalization sets Normalization field to given value.

### HasNormalization

`func (o *AutoTagRule) HasNormalization() bool`

HasNormalization returns a boolean if a field has been set.

### SetNormalizationNil

`func (o *AutoTagRule) SetNormalizationNil(b bool)`

 SetNormalizationNil sets the value for Normalization to be an explicit nil

### UnsetNormalization
`func (o *AutoTagRule) UnsetNormalization()`

UnsetNormalization ensures that no value is present for Normalization, not even an explicit nil
### GetPropagationTypes

`func (o *AutoTagRule) GetPropagationTypes() []string`

GetPropagationTypes returns the PropagationTypes field if non-nil, zero value otherwise.

### GetPropagationTypesOk

`func (o *AutoTagRule) GetPropagationTypesOk() (*[]string, bool)`

GetPropagationTypesOk returns a tuple with the PropagationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationTypes

`func (o *AutoTagRule) SetPropagationTypes(v []string)`

SetPropagationTypes sets PropagationTypes field to given value.

### HasPropagationTypes

`func (o *AutoTagRule) HasPropagationTypes() bool`

HasPropagationTypes returns a boolean if a field has been set.

### GetType

`func (o *AutoTagRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoTagRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoTagRule) SetType(v string)`

SetType sets Type field to given value.


### GetValueFormat

`func (o *AutoTagRule) GetValueFormat() string`

GetValueFormat returns the ValueFormat field if non-nil, zero value otherwise.

### GetValueFormatOk

`func (o *AutoTagRule) GetValueFormatOk() (*string, bool)`

GetValueFormatOk returns a tuple with the ValueFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFormat

`func (o *AutoTagRule) SetValueFormat(v string)`

SetValueFormat sets ValueFormat field to given value.

### HasValueFormat

`func (o *AutoTagRule) HasValueFormat() bool`

HasValueFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


